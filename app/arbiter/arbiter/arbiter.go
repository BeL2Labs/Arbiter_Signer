// Copyright (c) 2025 The bel2 developers

package arbiter

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/ecdsa"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/ethereum/go-ethereum/common"
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/BeL2Labs/Arbiter_Signer/app/arbiter/api/mempool"
	"github.com/BeL2Labs/Arbiter_Signer/app/arbiter/config"
	"github.com/BeL2Labs/Arbiter_Signer/app/arbiter/contract"
	"github.com/BeL2Labs/Arbiter_Signer/app/arbiter/contract/events"
)

const DELAY_BLOCK uint64 = 3

type account struct {
	PrivateKey string `json:"privKey"`
}

type Arbiter struct {
	ctx     context.Context
	config  *config.Config
	escNode *contract.ArbitratorContract
	account *account

	mempoolAPI *mempool.API

	logger *log.Logger
}

func NewArbiter(ctx context.Context, config *config.Config) *Arbiter {
	escData, err := os.ReadFile(config.EscKeyFilePath)
	if err != nil {
		g.Log().Fatal(ctx, "get esc keyfile error", err, " keystore path ", config.EscKeyFilePath)
	}
	var escAccount account
	err = json.Unmarshal(escData, &escAccount)
	if err != nil {
		g.Log().Fatal(ctx, "Unmarshal keyfile error", err, " content ", string(escData))
	}

	arbiterData, err := os.ReadFile(config.ArbiterKeyFilePath)
	if err != nil {
		g.Log().Fatal(ctx, "get arbiter keyfile error", err, " keystore path ", config.ArbiterKeyFilePath)
	}
	var arbiterAccount account
	err = json.Unmarshal(arbiterData, &arbiterAccount)
	if err != nil {
		g.Log().Fatal(ctx, "Unmarshal keyfile error", err, " content ", string(arbiterData))
	}

	err = createDir(config)
	if err != nil {
		g.Log().Fatal(ctx, "create dir error", err)
	}

	logFilePath := gfile.Join(config.LoanLogPath, "event.log")
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		g.Log().Fatal(ctx, "create log file error", err)
	}
	logger := log.New(logFile, "", log.Ldate|log.Ltime)

	escNode := newESCNode(ctx, config, escAccount.PrivateKey, logger)

	mempoolAPI := mempool.NewAPI(mempool.Config{Network: config.Network})

	return &Arbiter{
		ctx:        ctx,
		config:     config,
		account:    &arbiterAccount,
		escNode:    escNode,
		mempoolAPI: mempoolAPI,
		logger:     logger,
	}
}

func (v *Arbiter) Start() {
	if v.config.Signer {
		go v.processArbiterSig()
	}

	if v.config.Listener {
		go v.listenESCContract()
	}
}

func (v *Arbiter) listenESCContract() {
	g.Log().Info(v.ctx, "listenESCContract start")

	startHeight, _ := events.GetCurrentBlock(v.config.DataDir)
	if v.config.ESCStartHeight > startHeight {
		startHeight = v.config.ESCStartHeight
	}

	keyfile := v.config.EscKeyFilePath
	data, err := os.ReadFile(keyfile)
	if err != nil {
		g.Log().Fatal(v.ctx, "get keyfile error", err, " private key path ", keyfile)
	}
	var a account
	err = json.Unmarshal(data, &a)
	if err != nil {
		g.Log().Fatal(v.ctx, "Unmarshal keyfile error", err, " content ", string(data))
	}

	v.escNode.Start(startHeight)
}

func (v *Arbiter) processArbiterSig() {
	g.Log().Info(v.ctx, "processArbiterSignature start")

	var netWorkParams = chaincfg.MainNetParams
	if strings.ToLower(v.config.Network) == "testnet" {
		netWorkParams = chaincfg.TestNet3Params
	}

	for {
		// get all deploy file
		files, err := os.ReadDir(v.config.LoanNeedSignReqPath)
		if err != nil {
			g.Log().Error(v.ctx, "read dir error", err)
			continue
		}

		for _, file := range files {
			// read file
			filePath := v.config.LoanNeedSignReqPath + "/" + file.Name()
			fileContent, err := os.ReadFile(filePath)
			if err != nil {
				g.Log().Error(v.ctx, "read file error", err)
				continue
			}
			logEvt, err := v.decodeLogEvtByFileContent(fileContent)
			if err != nil {
				g.Log().Error(v.ctx, "decodeLogEvtByFileContent error", err)
				v.moveToDirectory(filePath, v.config.LoanNeedSignFailedPath+"/"+file.Name()+".failed")
				v.logger.Println("[ERR]  SIGN: decode event failed, file:", filePath)
				continue
			}
			var ev = make(map[string]interface{})
			err = v.escNode.Loan_abi.UnpackIntoMap(ev, "ArbitrationRequested", logEvt.EventData)
			if err != nil {
				g.Log().Error(v.ctx, "UnpackIntoMap error", err)
				v.moveToDirectory(filePath, v.config.LoanNeedSignFailedPath+"/"+file.Name()+".failed")
				v.logger.Println("[ERR]  SIGN: unpack event into map failed, file:", filePath)
				continue
			}
			g.Log().Info(v.ctx, "ev", ev)
			queryId := logEvt.Topics[1]
			dappAddress := logEvt.Topics[2]
			rawData := ev["btcTx"].([]byte)
			script := ev["script"].([]byte)
			arbitratorAddress := ev["arbitrator"].(common.Address)

			g.Log().Info(v.ctx, "dappAddress", dappAddress)
			g.Log().Info(v.ctx, "queryId", hex.EncodeToString(queryId[:]))
			g.Log().Info(v.ctx, "rawData", hex.EncodeToString(rawData))
			g.Log().Info(v.ctx, "script", hex.EncodeToString(script))
			g.Log().Info(v.ctx, "arbitratorAddress", arbitratorAddress)

			// sign btc tx
			tx, err := decodeTx(rawData)
			if err != nil {
				g.Log().Error(v.ctx, "decodeTx error", err, "rawData:", rawData)
				v.moveToDirectory(filePath, v.config.LoanNeedSignFailedPath+"/"+file.Name()+".decodeRawDataFailed")
				v.logger.Println("[ERR]  SIGN: decode event failed, block:", logEvt.Block, "tx:", logEvt.TxHash)
				continue
			}

			// todo check tx outputs, need to have output to arbiter
			if len(tx.TxOut) != 2 {
				g.Log().Error(v.ctx, "invalid tx outputs", len(tx.TxOut))
				v.moveToDirectory(filePath, v.config.LoanNeedSignFailedPath+"/"+file.Name()+".invalidTxOutputs")
				v.logger.Println("[ERR]  SIGN: invalid tx outputs count, block:", logEvt.Block, "tx:", logEvt.TxHash)
				continue
			}

			// get btc arbiter BTC address
			arbiterAddress := common.HexToAddress(arbitratorAddress.String())
			arbitratorBTCAddress, err := v.escNode.GetArbiterBTCAddress(arbiterAddress)
			if err != nil {
				g.Log().Error(v.ctx, "GetArbiterOperatorAddress error", err)
				v.moveToDirectory(filePath, v.config.LoanNeedSignFailedPath+"/"+file.Name()+".getArbiterOperatorAddressFailed")
				v.logger.Println("[ERR]  SIGN: get arbiter operator address failed, block:", logEvt.Block, "tx:", logEvt.TxHash)
				continue
			}
			existArbiterBtcAddress := false
			arbiterFeeVoutIndex := 0
			for i, vout := range tx.TxOut {
				if vout.Value < 546 {
					g.Log().Error(v.ctx, "invalid tx outputs with dust value")
					v.moveToDirectory(filePath, v.config.LoanNeedSignFailedPath+"/"+file.Name()+".invalidTxOutputs")
					v.logger.Println("[ERR]  SIGN: invalid tx outputs with dust value, block:", logEvt.Block, "tx:", logEvt.TxHash)
					continue
				}
				_, addrs, _, err := txscript.ExtractPkScriptAddrs(vout.PkScript, &netWorkParams)
				if err != nil {
					g.Log().Error(v.ctx, "ExtractPkScriptAddrs err:", err)
					v.moveToDirectory(filePath, v.config.LoanNeedSignFailedPath+"/"+file.Name()+".ExtractPkScriptAddrsFailed")
					v.logger.Println("[ERR]  SIGN: extract pk script failed, block:", logEvt.Block, "tx:", logEvt.TxHash)
					continue
				}
				for _, addr := range addrs {
					g.Log().Info(v.ctx, "tx vout addr:", addr.String(), "arbitratorBTCAddress:", arbitratorBTCAddress)
					if addr.String() == arbitratorBTCAddress {
						existArbiterBtcAddress = true
						arbiterFeeVoutIndex = i
					}
				}
			}
			if !existArbiterBtcAddress {
				g.Log().Error(v.ctx, "invalid tx outputs, without fee to arbiter")
				v.moveToDirectory(filePath, v.config.LoanNeedSignFailedPath+"/"+file.Name()+".invalidTxOutputs")
				v.logger.Println("[ERR]  SIGN: invalid tx outputs, without fee to arbiter, block:", logEvt.Block, "tx:", logEvt.TxHash)
				continue
			}

			// get pay address by script
			script1Hash := sha256.Sum256(script)
			wsh, err := btcutil.NewAddressWitnessScriptHash(script1Hash[:], &netWorkParams)
			if err != nil {
				g.Log().Error(v.ctx, "NewAddressWitnessScriptHash err:", err)
				v.moveToDirectory(filePath, v.config.LoanNeedSignFailedPath+"/"+file.Name()+".newAddressWitnessScriptHashFailed")
				v.logger.Println("[ERR]  SIGN: new addr witness sh failed, block:", logEvt.Block, "tx:", logEvt.TxHash)
				continue
			}
			payAddress, err := btcutil.DecodeAddress(wsh.EncodeAddress(), &netWorkParams)
			if err != nil {
				g.Log().Error(v.ctx, "DecodeAddress err:", err)
				v.moveToDirectory(filePath, v.config.LoanNeedSignFailedPath+"/"+file.Name()+".DecodeAddressFailed")
				v.logger.Println("[ERR]  SIGN: decode address failed, block:", logEvt.Block, "tx:", logEvt.TxHash)
				continue
			}
			g.Log().Info(v.ctx, "payAddress", payAddress.String())
			p2wsh, err := txscript.PayToAddrScript(payAddress)
			if err != nil {
				g.Log().Error(v.ctx, "PayToAddrScript err:", err)
				v.moveToDirectory(filePath, v.config.LoanNeedSignFailedPath+"/"+file.Name()+".PayToAddrScriptFailed")
				v.logger.Println("[ERR]  SIGN: get ptaddr script failed, block:", logEvt.Block, "tx:", logEvt.TxHash)
				continue
			}

			// get preOutput by tx.Inputs(idx)
			idx := 1 - arbiterFeeVoutIndex
			input := tx.TxIn[idx]
			g.Log().Info(v.ctx, "input.PreviousOutPoint.Hash", input.PreviousOutPoint.Hash.String())
			preTx, err := v.mempoolAPI.GetRawTransaction(input.PreviousOutPoint.Hash.String())
			if err != nil {
				g.Log().Error(v.ctx, "GetRawTransaction error", err)
				// v.moveToDirectory(filePath, v.config.LoanNeedSignFailedPath+"/"+file.Name()+".GetRawTransactionFailed")
				v.logger.Println("[ERR]  SIGN: get raw tx failed, block:", logEvt.Block, "tx:", logEvt.TxHash)
				continue
			}
			// check utxo tx output counts
			if len(preTx.Vout) <= int(input.PreviousOutPoint.Index) {
				g.Log().Error(v.ctx, "invalid input.PreviousOutPoint.Index", input.PreviousOutPoint.Index)
				v.moveToDirectory(filePath, v.config.LoanNeedSignFailedPath+"/"+file.Name()+".invalidPreviousOutPointIndex")
				v.logger.Println("[ERR]  SIGN: invalid input.PreviousOutPoint.Index, block:", logEvt.Block, "tx:", logEvt.TxHash)
				continue
			}
			preAmount := int64(preTx.Vout[input.PreviousOutPoint.Index].Value)

			// check fee rate
			feeRate, err := v.escNode.GetArbitrationBTCFeeRate()
			if err != nil {
				g.Log().Error(v.ctx, "GetArbitrationBTCFeeRate error", err)
				v.moveToDirectory(filePath, v.config.LoanNeedSignFailedPath+"/"+file.Name()+".GetArbitrationBTCFeeRateFailed")
				v.logger.Println("[ERR]  SIGN: get fee rate failed, block:", logEvt.Block, "tx:", logEvt.TxHash)
				continue
			}
			arbiterFee := preAmount * feeRate.Int64() / 10000
			if tx.TxOut[arbiterFeeVoutIndex].Value < arbiterFee {
				g.Log().Error(v.ctx, "invalid fee rate", feeRate)
				v.moveToDirectory(filePath, v.config.LoanNeedSignFailedPath+"/"+file.Name()+".invalidFeeRate")
				v.logger.Println("[ERR]  SIGN: invalid fee rate, block:", logEvt.Block, "tx:", logEvt.TxHash)
				continue
			}

			// check utxo address
			g.Log().Info(v.ctx, "### preTx Vout:", preTx.Vout[input.PreviousOutPoint.Index])
			utxoAddr := preTx.Vout[input.PreviousOutPoint.Index].ScriptpubkeyAddress
			if utxoAddr != payAddress.EncodeAddress() {
				g.Log().Error(v.ctx, "invalid utxo address:", utxoAddr, "need to be:", payAddress)
				v.moveToDirectory(filePath, v.config.LoanNeedSignFailedPath+"/"+file.Name()+".invalidUtxoAddress")
				v.logger.Println("[ERR]  SIGN: invalid utxo address, block:", logEvt.Block, "tx:", logEvt.TxHash)
				continue
			}

			// calculate sigHash
			prevFetcher := txscript.NewCannedPrevOutputFetcher(
				p2wsh, preAmount,
			)
			sigHashes := txscript.NewTxSigHashes(tx, prevFetcher)
			sigHash, err := txscript.CalcWitnessSigHash(script, sigHashes, txscript.SigHashAll, tx, idx, preAmount)
			if err != nil {
				g.Log().Error(v.ctx, "CalcWitnessSigHash error", err)
				v.moveToDirectory(filePath, v.config.LoanNeedSignFailedPath+"/"+file.Name()+".CalcWitnessSigHashFailed")
				v.logger.Println("[ERR]  SIGN: calculate sigHash failed, block:", logEvt.Block, "tx:", logEvt.TxHash)
				continue
			}
			var sigDataHash [32]byte
			copy(sigDataHash[:], sigHash)
			g.Log().Info(v.ctx, "sigHash", hex.EncodeToString(sigDataHash[:]))
			g.Log().Info(v.ctx, "script", hex.EncodeToString(script))

			// first := sha256.Sum256(rawData)
			// sigHash := sha256.Sum256(first[:])

			// ecdsa sign
			priKeyBytes, _ := hex.DecodeString(v.account.PrivateKey)
			priKey, _ := btcec.PrivKeyFromBytes(priKeyBytes)
			signatureArbiter := ecdsa.Sign(priKey, sigHash[:])
			ok := signatureArbiter.Verify(sigHash[:], priKey.PubKey())
			if !ok {
				g.Log().Error(v.ctx, "self ecdsa sign verify failed")
				v.moveToDirectory(filePath, v.config.LoanNeedSignFailedPath+"/"+file.Name()+".SigVerifyFailed")
				v.logger.Println("[ERR]  SIGN: signature verify failed, block:", logEvt.Block, "tx:", logEvt.TxHash)
				continue
			}
			signatureBytes := signatureArbiter.Serialize()
			// signatureBytes = append(signatureBytes, byte(txscript.SigHashAll))
			g.Log().Info(v.ctx, "arbiter signature:", hex.EncodeToString(signatureBytes))

			// feedback signature to contract
			txhash, err := v.escNode.SubmitArbitrationSignature(signatureBytes, queryId)
			g.Log().Notice(v.ctx, "submitArbitrationSignature", "txhash ", txhash.String(), " error ", err)
			if err != nil {
				v.moveToDirectory(v.config.LoanNeedSignReqPath+"/"+file.Name(), v.config.LoanNeedSignFailedPath+"/"+file.Name()+".SubmitSignatureFailed")
				v.logger.Println("[ERR]  SIGN: SubmitArbitrationSignature failed, block:", logEvt.Block, "tx:", logEvt.TxHash, "err:", err.Error())
			} else {
				v.moveToDirectory(v.config.LoanNeedSignReqPath+"/"+file.Name(), v.config.LoanNeedSignSignedPath+"/"+file.Name()+".Succeed")
				v.logger.Println("[INF]  SIGN: SubmitArbitrationSignature succeed, block:", logEvt.Block, "tx:", logEvt.TxHash)
			}
		}

		// sleep 10s to check and process next files
		time.Sleep(time.Second * 10)
	}
}

func decodeTx(txBytes []byte) (*wire.MsgTx, error) {
	tx := wire.NewMsgTx(2)
	err := tx.Deserialize(bytes.NewReader(txBytes))
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (v *Arbiter) decodeLogEvtByFileContent(content []byte) (*events.ContractLogEvent, error) {
	logEvt := &events.ContractLogEvent{}
	err := gob.NewDecoder(bytes.NewReader(content)).Decode(logEvt)
	if err != nil {
		g.Log().Error(v.ctx, "NewDecoder deployBRC20 error", err)
		return nil, err
	}
	return logEvt, nil
}

func createDir(config *config.Config) error {
	if !gfile.Exists(config.LoanNeedSignReqPath) {
		err := gfile.Mkdir(config.LoanNeedSignReqPath)
		if err != nil {
			return err
		}
	}

	if !gfile.Exists(config.LoanNeedSignFailedPath) {
		err := gfile.Mkdir(config.LoanNeedSignFailedPath)
		if err != nil {
			return err
		}
	}

	if !gfile.Exists(config.LoanNeedSignSignedPath) {
		err := gfile.Mkdir(config.LoanNeedSignSignedPath)
		if err != nil {
			return err
		}
	}

	if !gfile.Exists(config.LoanLogPath) {
		err := gfile.Mkdir(config.LoanLogPath)
		if err != nil {
			return err
		}
	}

	return nil
}

func (v *Arbiter) moveToDirectory(oldPath, newPath string) {
	dir := filepath.Dir(newPath)
	_, err := os.Stat(newPath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			g.Log().Error(v.ctx, " createConfirmDir filePaht ", err)
			return
		}
	}

	g.Log().Info(v.ctx, "move file from:", oldPath, "to:", newPath)
	if err := os.Rename(oldPath, newPath); err != nil {
		g.Log().Error(v.ctx, "moveToDirectory error", err, "from:", oldPath, "to:", newPath)
	}
}

func newESCNode(ctx context.Context, config *config.Config, privateKey string, logger *log.Logger) *contract.ArbitratorContract {
	startHeight, err := events.GetCurrentBlock(config.DataDir)
	if err == nil {
		config.ESCStartHeight = startHeight
	}

	contractNode, err := contract.New(ctx, config, privateKey, logger)
	if err != nil {
		g.Log().Fatal(ctx, err)
	}
	return contractNode
}

func getTxHex(tx *wire.MsgTx) (string, error) {
	var buf bytes.Buffer
	if err := tx.Serialize(&buf); err != nil {
		return "", err
	}
	return hex.EncodeToString(buf.Bytes()), nil
}

func GetPubKey(privKeyStr string) (pk string, err error) {
	priKeyBytes, err := hex.DecodeString(privKeyStr)
	if err != nil {
		return
	}
	_, pubKey := btcec.PrivKeyFromBytes(priKeyBytes)
	pk = hex.EncodeToString(pubKey.SerializeCompressed())

	return
}

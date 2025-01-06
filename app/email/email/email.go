// Copyright (c) 2025 The bel2 developers

package email

// Copyright (c) 2025 The bel2 developers

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/hex"
	"log"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/wire"
	"github.com/ethereum/go-ethereum/common"
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/v2/frame/g"
	"gopkg.in/gomail.v2"

	"github.com/BeL2Labs/Arbiter_Signer/app/email/config"
	"github.com/BeL2Labs/Arbiter_Signer/utility/contract"
	"github.com/BeL2Labs/Arbiter_Signer/utility/events"
)

const DELAY_BLOCK uint64 = 3

type account struct {
	PrivateKey string `json:"privKey"`
}

type Email struct {
	ctx     context.Context
	config  *config.Config
	escNode *contract.ArbitratorContract

	logger *log.Logger
}

func NewEmail(ctx context.Context, config *config.Config) *Email {

	err := createDir(config)
	if err != nil {
		g.Log().Fatal(ctx, "create dir error", err)
	}

	logFilePath := gfile.Join(config.EmailLogPath, "event.log")
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		g.Log().Fatal(ctx, "create log file error", err)
	}
	logger := log.New(logFile, "", log.Ldate|log.Ltime)

	escNode := newESCNode(ctx, config, logger)

	return &Email{
		ctx:     ctx,
		config:  config,
		escNode: escNode,
		logger:  logger,
	}
}

func (e *Email) Start() {
	go func() {
		go e.sendEmail()
	}()

	go func() {
		go e.listenESCContract()
	}()
}

func (e *Email) sendEmail() error {
	g.Log().Info(e.ctx, "start check and send email ...")
	time.Sleep(time.Minute)
	for {
		// get all deploy file
		files, err := os.ReadDir(e.config.LoanNeedSignReqPath)
		if err != nil {
			continue
		}

		eventsMap := make(map[string][]BlockID)
		for _, file := range files {
			// read file
			filePath := e.config.LoanNeedSignReqPath + "/" + file.Name()
			fileContent, err := os.ReadFile(filePath)
			if err != nil {
				g.Log().Error(e.ctx, "read file error", err)
				e.logger.Println("[ERR] unpack event into map failed, file:", filePath)
				continue
			}
			logEvt, err := e.decodeLogEvtByFileContent(fileContent)
			if err != nil {
				g.Log().Error(e.ctx, "decodeLogEvtByFileContent error", err)
				os.Remove(filePath)
				e.logger.Println("[ERR] unpack event into map failed, file:", filePath)
				continue
			}
			var ev = make(map[string]interface{})
			err = e.escNode.Loan_abi.UnpackIntoMap(ev, "ArbitrationRequested", logEvt.EventData)
			if err != nil {
				g.Log().Error(e.ctx, "UnpackIntoMap error", err)
				os.Remove(filePath)
				e.logger.Println("[ERR] unpack event into map failed, file:", filePath)
				continue
			}
			queryId := logEvt.Topics[1]
			arbitratorAddress := ev["arbitrator"].(common.Address)

			g.Log().Info(e.ctx, "queryId", hex.EncodeToString(queryId[:]))

			if _, ok := eventsMap[arbitratorAddress.String()]; !ok {
				eventsMap[arbitratorAddress.String()] = make([]BlockID, 0)
			}
			blockId := BlockID{
				Block:   strconv.FormatUint(logEvt.Block, 10),
				QueryId: queryId.String(),
			}
			eventsMap[arbitratorAddress.String()] = append(eventsMap[arbitratorAddress.String()], blockId)
			os.Remove(filePath)
			g.Log().Info(e.ctx, "remove file", filePath)
		}

		g.Log().Info(e.ctx, "eventsMap", eventsMap)
		subject := "Request Arbitration Events"
		body := ""
		for k, v := range eventsMap {
			sort.Sort(ByBlockID(v))
			body += "Arbitrator: " + k + "\n"
			body += "=======================================================================\n"
			for _, blockID := range v {
				body += blockID.Block + ": " + blockID.QueryId + "\n"
			}
			body += "=======================================================================\n\n"
		}

		// subject := "Request Arbitration Events"
		// body := "<html><body>"
		// for arbitrator, blockIDs := range eventsMap {
		// 	body += "<table border='1' style='margin-bottom: 20px; width: 50%;'>"
		// 	body += "<tr><th>Arbiter Address</th><th>Block Number</th><th>Transaction ID</th></tr>"
		// 	for _, blockID := range blockIDs {
		// 		body += "<tr>"
		// 		body += "<td>" + arbitrator + "</td>"
		// 		body += "<td>" + blockID.Block + "</td>"
		// 		body += "<td>" + blockID.QueryId + "</td>"
		// 		body += "</tr>"
		// 	}
		// 	body += "</table>"
		// }
		// body += "</body></html>"

		if body == "" {
			body = "No events"
		}
		g.Log().Info(e.ctx, "start send email")
		g.Log().Info(e.ctx, "subject", subject)
		g.Log().Info(e.ctx, "body", body)
		err = sendEmail(e.config, subject, body)
		if err != nil {
			g.Log().Error(e.ctx, "send email error", err)
			return err
		}
		g.Log().Info(e.ctx, "send email success")

		// sleep one hour and resend email
		time.Sleep(time.Duration(e.config.Duration * int(time.Second)))
	}
}

func (e *Email) listenESCContract() {
	g.Log().Info(e.ctx, "listenESCContract start")

	startHeight, _ := events.GetCurrentBlock(e.config.DataPath)
	if e.config.ESCStartHeight > startHeight {
		startHeight = e.config.ESCStartHeight
	}

	e.escNode.Start(startHeight)
}

type BlockID struct {
	Block   string
	QueryId string
}

type ByBlockID []BlockID

func (b ByBlockID) Len() int {
	return len(b)
}
func (b ByBlockID) Less(i, j int) bool {
	blockI, errI := strconv.ParseUint(b[i].Block, 10, 64)
	blockJ, errJ := strconv.ParseUint(b[j].Block, 10, 64)
	if errI != nil {
		return true
	}
	if errJ != nil {
		return false
	}
	return blockI < blockJ
}
func (b ByBlockID) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func decodeTx(txBytes []byte) (*wire.MsgTx, error) {
	tx := wire.NewMsgTx(2)
	err := tx.Deserialize(bytes.NewReader(txBytes))
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func createDir(config *config.Config) error {
	if !gfile.Exists(config.EmailLogPath) {
		err := gfile.Mkdir(config.EmailLogPath)
		if err != nil {
			return err
		}
	}

	if !gfile.Exists(config.DataPath) {
		err := gfile.Mkdir(config.DataPath)
		if err != nil {
			return err
		}
	}

	if !gfile.Exists(config.LoanNeedSignReqPath) {
		err := gfile.Mkdir(config.LoanNeedSignReqPath)
		if err != nil {
			return err
		}
	}

	return nil
}

func newESCNode(ctx context.Context, config *config.Config, logger *log.Logger) *contract.ArbitratorContract {
	startHeight, err := events.GetCurrentBlock(config.DataPath)
	if err == nil {
		config.ESCStartHeight = startHeight
	}

	arbiterAddress := make(map[string]struct{})
	for _, addr := range config.ArbiterAddresses {
		arbiterAddress[addr] = struct{}{}
	}
	cfg := &contract.Config{
		Http: config.Http,

		ESCArbiterAddresses:       arbiterAddress,
		ESCArbiterContractAddress: config.ESCArbiterContractAddress,

		DataDir:             config.DataPath,
		LoanNeedSignReqPath: config.LoanNeedSignReqPath,
	}

	contractNode, err := contract.New(ctx, cfg, "", logger)
	if err != nil {
		g.Log().Fatal(ctx, err)
	}
	return contractNode
}

func (e *Email) decodeLogEvtByFileContent(content []byte) (*events.ContractLogEvent, error) {
	logEvt := &events.ContractLogEvent{}
	err := gob.NewDecoder(bytes.NewReader(content)).Decode(logEvt)
	if err != nil {
		g.Log().Error(e.ctx, "NewDecoder deployBRC20 error", err)
		return nil, err
	}
	return logEvt, nil
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

func sendEmail(cfg *config.Config, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", cfg.From)
	m.SetHeader("To", cfg.To...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)
	d := gomail.NewDialer(cfg.Host, cfg.Port, cfg.User, cfg.Password)

	return d.DialAndSend(m)
}

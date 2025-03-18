// Copyright (c) 2025 The bel2 developers

package contract

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/BeL2Labs/Arbiter_Signer/app/arbiter/config"
	"github.com/BeL2Labs/Arbiter_Signer/app/arbiter/contract/contract_abi"
	"github.com/BeL2Labs/Arbiter_Signer/app/arbiter/contract/events"
)

type ArbitratorContract struct {
	listener  *ContractListener
	submitter *ContractSubmitter
	ctx       context.Context

	chan_event          chan *events.ContractLogEvent
	chan_interrupt      chan struct{}
	Loan_abi            abi.ABI
	Arbiter_manager_abi abi.ABI
	Arbiter_config_abi  abi.ABI
	Order_abi           abi.ABI

	loanContract           *common.Address
	arbiterManagerContract *common.Address
	configManagerContract  *common.Address
	orderManagerContract   *common.Address
	cfg                    *config.Config

	logger *log.Logger
}

// ArbitratorStatus should be defined according to the possible statuses you have.
type ArbitratorInfo struct {
	Arbitrator            string   // Arbitrator Ethereum address
	Paused                bool     // Paused
	CurrentFeeRate        *big.Int // Current fee rate
	ActiveTransactionID   [32]byte // Current transaction ID (bytes32)
	EthAmount             *big.Int // ETH stake amount
	Erc20Token            string   // ERC20 token address
	NftContract           string   // NFT contract address
	NftTokenIds           []uint64 // NFT token IDs
	Operator              string   // Operator address
	OperatorBtcPubKey     []byte   // Bitcoin public key
	OperatorBtcAddress    string   // Bitcoin address
	Deadline              *big.Int // Last arbitration time , deadline
	RevenueBtcPubKey      []byte   // Bitcoin public key for receiving arbitrator earnings
	RevenueBtcAddress     string   // Bitcoin address for receiving arbitrator earnings
	RevenueETHAddress     string   // ETH address for receiving arbitrator earnings
	LastSubmittedWorkTime *big.Int // Last submitted work time
}

func New(ctx context.Context, cfg *config.Config, privateKey string, logger *log.Logger) (*ArbitratorContract, error) {
	client, err := ConnectRPC(cfg.Http)
	if err != nil {
		return nil, err
	}
	loanAddress := common.HexToAddress(cfg.ESCArbiterContractAddress)
	arbiterManagerAddress := common.HexToAddress(cfg.ESCArbiterManagerContractAddress)
	configManagerAddress := common.HexToAddress(cfg.ESCConfigManagerContractAddress)
	orderManangerAddress := common.HexToAddress(cfg.ESCOrderManagerContractAddress)
	eventChan := make(chan *events.ContractLogEvent, 3)
	chan_interrupt := make(chan struct{})
	listener, err := NewListener(ctx, client, loanAddress, orderManangerAddress, eventChan)
	if err != nil {
		return nil, err
	}

	loanABI, err := abi.JSON(strings.NewReader(contract_abi.ArbiterABI))
	if err != nil {
		return nil, err
	}
	arbiterManagerABI, err := abi.JSON(strings.NewReader(contract_abi.ArbiterManagerABI))
	if err != nil {
		return nil, err
	}
	arbiterConfigABI, err := abi.JSON(strings.NewReader(contract_abi.ArbiterConfigManagerABI))
	if err != nil {
		return nil, err
	}
	orderABI, err := abi.JSON(strings.NewReader(contract_abi.OrderEventManagerABI))
	if err != nil {
		return nil, err
	}
	submitter, err := NewSubmitter(ctx, client, privateKey)
	if err != nil {
		return nil, err
	}
	c := &ArbitratorContract{
		listener:               listener,
		submitter:              submitter,
		ctx:                    ctx,
		chan_event:             eventChan,
		chan_interrupt:         chan_interrupt,
		Loan_abi:               loanABI,
		Arbiter_manager_abi:    arbiterManagerABI,
		Arbiter_config_abi:     arbiterConfigABI,
		Order_abi:              orderABI,
		loanContract:           &loanAddress,
		arbiterManagerContract: &arbiterManagerAddress,
		configManagerContract:  &configManagerAddress,
		orderManagerContract:   &orderManangerAddress,
		cfg:                    cfg,
		logger:                 logger,
	}
	return c, nil
}

func (c *ArbitratorContract) Start(startHeight uint64) error {

	// get arbitrator operator address
	arbiterAddress := common.HexToAddress(c.cfg.ESCArbiterAddress)
	arbitratorOperatorAddress, err := c.getArbiterOperatorAddress(arbiterAddress)
	if err != nil {
		g.Log().Error(c.ctx, "GetArbiterOperatorAddress error", err)
		panic("invalid arbiter address, err:" + err.Error())
	}
	g.Log().Info(c.ctx, "arbitratorOperatorAddress", arbitratorOperatorAddress)
	// check operator address
	if c.submitter.keypair.Address() != arbitratorOperatorAddress.String() {
		g.Log().Error(c.ctx, "Invalid operator address from arbiter address")
		panic("invalid operator address, " +
			"operator from key file:" + c.submitter.keypair.Address() +
			"operator from config:" + arbitratorOperatorAddress.String())
	}

	go func() {
		for {
			select {
			case evt := <-c.chan_event:
				err := c.parseContractEvent(evt)
				if err != nil {
					g.Log().Error(c.ctx, "parseContractEvent failed ", err)
				}
				// g.Log().Info(c.ctx, "parseContractEvent success:", evt)
			}
		}
	}()

	for {
		endBlock, err := c.listener.Start(startHeight)
		if err == nil {
			err = events.UpdateCurrentBlock(c.cfg.DataDir, endBlock)
			if err != nil {
				g.Log().Error(c.ctx, "UpdateCurrentBlock faield ", err)
			}
			startHeight = endBlock + 1
		}
		time.Sleep(5 * time.Second)
	}

}

func (c *ArbitratorContract) parseContractEvent(event *events.ContractLogEvent) error {
	var err error
	if event.Topics[0].Cmp(events.ArbitrationRequested) == 0 {
		err = c.parseTransferNeedSignEvent(event)
		fmt.Println("ArbitrationRequested  >>>>>>>>>>>>>>>> received")
	} else if event.Topics[0].Cmp(events.ArbitrationResultSubmitted) == 0 {
		err = c.parseTransferSignedEvent(event)
		fmt.Println("ArbitrationResultSubmitted  >>>>>>>>>>>>>>>> received")
	} else if event.Topics[0].Cmp(events.ConfirmTransferToLenderEvent) == 0 {
		err = c.parseConfirmTransferToLenderEvent(event)
		fmt.Println("ConfirmTransferToLenderEvent  >>>>>>>>>>>>>>>> received")
	}
	return err
}

func (c *ArbitratorContract) GetSubmiterAddress() string {
	return c.submitter.keypair.Address()
}

func (c *ArbitratorContract) parseTransferNeedSignEvent(event *events.ContractLogEvent) error {
	var ev = make(map[string]interface{})
	err := c.Loan_abi.UnpackIntoMap(ev, "ArbitrationRequested", event.EventData)
	if err != nil {
		g.Log().Error(c.ctx, "parseTransferNeedSignEvent UnpackIntoMap error", err)
		return err
	}
	if ev["arbitrator"].(common.Address).String() != c.cfg.ESCArbiterAddress {
		g.Log().Debug(c.ctx, "find ArbitrationRequested event, but not mine")
		return nil
	}
	c.logger.Println("[INF] EVENT: ArbitrationRequested, block:", event.Block, "tx:", event.TxHash)

	path := c.cfg.LoanNeedSignReqPath + "/" + event.TxHash.String()
	err = events.SaveContractEvent(path, event)
	if err != nil {
		g.Log().Error(c.ctx, "SaveContractEvent error", err)
	}
	g.Log().Noticef(c.ctx, "find btc tx need sign:%s ", event.TxHash.String())
	return err
}

func (c *ArbitratorContract) parseTransferSignedEvent(event *events.ContractLogEvent) error {
	var ev = make(map[string]interface{})
	err := c.Loan_abi.UnpackIntoMap(ev, "ArbitrationResultSubmitted", event.EventData)
	if err != nil {
		g.Log().Error(c.ctx, "parseTransferSignedEvent UnpackIntoMap error", err)
		return err
	}
	path := c.cfg.LoanSignedEventPath + "/" + event.TxHash.String()
	err = events.SaveContractEvent(path, event)
	if err != nil {
		g.Log().Error(c.ctx, "SaveContractEvent error", err)
	}
	g.Log().Noticef(c.ctx, "find btc tx signed:%s ", event.TxHash.String())
	return err
}

func (c *ArbitratorContract) parseConfirmTransferToLenderEvent(event *events.ContractLogEvent) error {
	var ev = make(map[string]interface{})
	err := c.Order_abi.UnpackIntoMap(ev, "ConfirmTransferToLenderEvent", event.EventData)
	if err != nil {
		g.Log().Error(c.ctx, "parseLoanLenderManuallyConfirmedEvent UnpackIntoMap error", err)
		return err
	}
	// get btc address
	g.Log().Info(c.ctx, "arbiter:", event.Topics[3])
	if len(event.Topics) < 4 {
		return errors.New("invalid event count")
	}
	arbiterAddress, err := c.getArbiterOperatorAddress(common.BytesToAddress(event.Topics[3][:]))
	if err != nil {
		g.Log().Error(c.ctx, "GetArbiterBTCAddress error", err)
		return err
	}
	if arbiterAddress.String() != c.cfg.ESCArbiterAddress {
		g.Log().Debug(c.ctx, "find ConfirmTransferToLenderEvent, but not mine")
		return nil
	}
	c.logger.Println("[INF] EVENT: ConfirmTransferToLenderEvent, block:", event.Block, "tx:", event.TxHash)

	path := c.cfg.LoanManuallyConfirmedPath + "/" + event.TxHash.String()
	err = events.SaveContractEvent(path, event)
	if err != nil {
		g.Log().Error(c.ctx, "SaveContractEvent error", err)
	}
	g.Log().Noticef(c.ctx, "find btc tx lender manually confirmed:%s ", event.TxHash.String())
	return err
}

func (c *ArbitratorContract) SubmitManuallyConfirm(orderContractAddress *common.Address) (common.Hash, error) {
	input, err := c.Order_abi.Pack("confirmTransferToArbitrator")
	if err != nil {
		return common.Hash{}, err
	}
	hash, err := c.submitter.MakeAndSendContractTransaction(input, orderContractAddress, big.NewInt(0))
	return hash, err
}

func (c *ArbitratorContract) SubmitArbitrationSignature(rawData []byte, queryId [32]byte) (common.Hash, error) {
	input, err := c.Loan_abi.Pack("submitArbitration", queryId, rawData)
	if err != nil {
		return common.Hash{}, err
	}
	hash, err := c.submitter.MakeAndSendContractTransaction(input, c.loanContract, big.NewInt(0))
	return hash, err
}

func (c *ArbitratorContract) getArbiterOperatorAddress(arbiter common.Address) (common.Address, error) {
	input, err := c.Arbiter_manager_abi.Pack("getArbitratorInfo", arbiter)
	if err != nil {
		return common.Address{}, err
	}
	// use c.arbiterManagerContract to call get getArbitratorInfo operator address
	msg := ethereum.CallMsg{From: common.Address{}, To: c.arbiterManagerContract, Data: input}
	result, err := c.submitter.CallContract(context.TODO(), msg, nil)
	if err != nil {
		return common.Address{}, err
	}
	ev, err := c.Arbiter_manager_abi.Unpack("getArbitratorInfo", result)
	if err != nil || len(ev) == 0 {
		g.Log().Error(c.ctx, "parse ArbitratorInfo UnpackIntoMap error", err)
		return common.Address{}, err
	}
	info := ArbitratorInfo{}
	data, err := json.Marshal(ev[0])
	if err != nil {
		return common.Address{}, err
	}
	json.Unmarshal(data, &info)

	return common.HexToAddress(info.Operator), nil
}

func (c *ArbitratorContract) GetArbiterBTCAddress(arbiter common.Address) (string, error) {
	input, err := c.Arbiter_manager_abi.Pack("getArbitratorInfo", arbiter)
	if err != nil {
		return "", err
	}
	// use c.arbiterManagerContract to call get getArbitratorInfo operator address
	msg := ethereum.CallMsg{From: common.Address{}, To: c.arbiterManagerContract, Data: input}
	result, err := c.submitter.CallContract(context.TODO(), msg, nil)
	if err != nil {
		return "", err
	}
	ev, err := c.Arbiter_manager_abi.Unpack("getArbitratorInfo", result)
	if err != nil || len(ev) == 0 {
		g.Log().Error(c.ctx, "parse ArbitratorInfo UnpackIntoMap error", err)
		return "", err
	}
	info := ArbitratorInfo{}
	data, err := json.Marshal(ev[0])
	if err != nil {
		return "", err
	}
	json.Unmarshal(data, &info)

	return info.RevenueBtcAddress, nil
}

func (c *ArbitratorContract) GetArbitrationBTCFeeRate() (*big.Int, error) {
	input, err := c.Arbiter_config_abi.Pack("getArbitrationBTCFeeRate")
	if err != nil {
		return nil, err
	}
	msg := ethereum.CallMsg{From: common.Address{}, To: c.configManagerContract, Data: input}
	result, err := c.submitter.CallContract(context.TODO(), msg, nil)
	if err != nil {
		return nil, err
	}
	return big.NewInt(0).SetBytes(result), nil
}

type ManuallyConfirmedBTCFeeRate struct {
	CurrentBTCFeeRate *big.Int `json:"currentBTCFeeRate"`
}

func (c *ArbitratorContract) GetManuallyConfirmedBTCFeeRate(arbiter *common.Address) (*big.Int, error) {
	input, err := c.Arbiter_manager_abi.Pack("getArbitratorInfoExt", arbiter)
	if err != nil {
		return nil, err
	}
	msg := ethereum.CallMsg{From: common.Address{}, To: c.configManagerContract, Data: input}
	result, err := c.submitter.CallContract(context.TODO(), msg, nil)
	if err != nil {
		return nil, err
	}
	ev, err := c.Arbiter_manager_abi.Unpack("getArbitratorInfoExt", result)
	if err != nil || len(ev) == 0 {
		g.Log().Error(c.ctx, "parse ArbitratorInfo UnpackIntoMap error", err)
		return nil, err
	}
	info := ManuallyConfirmedBTCFeeRate{}
	data, err := json.Marshal(ev[0])
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &info)
	g.Log().Info(c.ctx, "ManuallyConfirmedBTCFeeRate", info.CurrentBTCFeeRate, "result:", result)
	return info.CurrentBTCFeeRate, nil
}

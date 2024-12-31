// Copyright (c) 2025 The bel2 developers

package contract

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/BeL2Labs/Arbiter_Signer/app/arbiter/contract/contract_abi"
	"github.com/BeL2Labs/Arbiter_Signer/app/arbiter/crypto"
	"github.com/BeL2Labs/Arbiter_Signer/app/arbiter/crypto/secp256k1"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type ContractSubmitter struct {
	client  *CrossClient
	ctx     context.Context
	keypair crypto.Keypair
}

func NewSubmitter(ctx context.Context, client *CrossClient, privateKey string) (*ContractSubmitter, error) {
	pri, err := hex.DecodeString(privateKey)
	kp, err := secp256k1.NewKeypairFromPrivateKey(pri)
	if err != nil {
		return nil, err
	}
	submitter := &ContractSubmitter{
		client:  client,
		ctx:     ctx,
		keypair: kp,
	}
	return submitter, nil
}

func (s *ContractSubmitter) MakeAndSendContractTransaction(data []byte, to *common.Address, value *big.Int) (common.Hash, error) {
	var hash common.Hash
	var from = s.keypair.CommonAddress()
	ctx := context.Background()
	gasPrice, err := s.client.SuggestGasPrice(ctx)
	if err != nil {
		log.Printf("SuggestGasPrice err: %v", err)
		return hash, err
	}
	msg := ethereum.CallMsg{From: from, To: to, Data: data, GasPrice: gasPrice, Value: value}
	gasLimit, err := s.client.EstimateGas(ctx, msg)
	if err != nil || gasLimit == 0 {
		log.Printf("EstimateGas err: %v", err)
		return hash, err
	}
	gasLimit = gasLimit + gasLimit*10
	nonce, err := s.client.PendingNonceAt(ctx, from)
	if err != nil {
		log.Printf("PendingNonceAt err: %v", err)
		return hash, err
	}

	tx := contract_abi.NewTransaction(nonce, to, value, gasLimit, gasPrice, data)

	return s.SignAndSendTransaction(ctx, tx)
}

func (s *ContractSubmitter) SignAndSendTransaction(ctx context.Context, tx *types.Transaction) (common.Hash, error) {

	id, err := s.client.ChainID(ctx)
	if err != nil {
		return common.Hash{}, err
	}
	rawTX, err := contract_abi.RawWithSignature(s.keypair.PrivateKey(), id, tx)
	if err != nil {
		return common.Hash{}, err
	}

	fmt.Println("SignAndSendTransaction rawTX:", hex.EncodeToString(rawTX))
	hash, err := s.client.SendRawTransaction(ctx, rawTX)
	return hash, err
}

func (s *ContractSubmitter) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	return s.client.CallContract(ctx, msg, blockNumber)
}

// Copyright (c) 2025 The bel2 developers

package events

import (
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	ArbitrationRequested = crypto.Keccak256Hash([]byte("ArbitrationRequested(bytes32,address,address,bytes,bytes,address)"))

	ArbitrationResultSubmitted = crypto.Keccak256Hash([]byte("ArbitrationResultSubmitted(bytes,bytes32)"))

	// ConfirmTransferToLenderEvent(address indexed order, bytes32 indexed txId, address indexed arbitrator, uint32 txIndex, uint256 arbitratorBtcFee);
	ConfirmTransferToLenderEvent = crypto.Keccak256Hash([]byte("ConfirmTransferToLenderEvent(address,bytes32,address,uint32,uint256)"))
)

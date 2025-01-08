// Copyright (c) 2025 The bel2 developers

package events

import (
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	ArbitrationRequested = crypto.Keccak256Hash([]byte("ArbitrationRequested(bytes32,address,address,bytes,bytes,address)"))

	ArbitrationSubmitted = crypto.Keccak256Hash([]byte("ArbitrationSubmitted(bytes32,address,address,bytes)"))
)

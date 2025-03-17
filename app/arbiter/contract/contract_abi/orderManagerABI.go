// Copyright (c) 2025 The bel2 developers

package contract_abi

const OrderEventManagerABI = `[
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "order",
          "type": "address"
        },
        {
          "indexed": true,
          "internalType": "bytes32",
          "name": "txId",
          "type": "bytes32"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "arbitrator",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "uint32",
          "name": "txIndex",
          "type": "uint32"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "arbitratorBtcFee",
          "type": "uint256"
        }
      ],
      "name": "ConfirmTransferToLenderEvent",
      "type": "event"
    },
	{
      "inputs": [],
      "name": "confirmTransferToArbitrator",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    }
  ]`

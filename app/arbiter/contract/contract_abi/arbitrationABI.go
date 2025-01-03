// Copyright (c) 2025 The bel2 developers

package contract_abi

const ArbiterABI = `[
    {
      "inputs": [],
      "stateMutability": "nonpayable",
      "type": "constructor"
    },
    {
      "inputs": [],
      "name": "InvalidInitialization",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "NotInitializing",
      "type": "error"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "owner",
          "type": "address"
        }
      ],
      "name": "OwnableInvalidOwner",
      "type": "error"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "account",
          "type": "address"
        }
      ],
      "name": "OwnableUnauthorizedAccount",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "ReentrancyGuardReentrantCall",
      "type": "error"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "bytes32",
          "name": "txId",
          "type": "bytes32"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "dapp",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "address",
          "name": "arbitrator",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "bytes",
          "name": "btcTx",
          "type": "bytes"
        },
        {
          "indexed": false,
          "internalType": "bytes",
          "name": "script",
          "type": "bytes"
        },
        {
          "indexed": false,
          "internalType": "address",
          "name": "timeoutCompensationReceiver",
          "type": "address"
        }
      ],
      "name": "ArbitrationRequested",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "bytes32",
          "name": "txId",
          "type": "bytes32"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "dapp",
          "type": "address"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "arbitrator",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "bytes",
          "name": "btcTxSignature",
          "type": "bytes"
        }
      ],
      "name": "ArbitrationSubmitted",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "uint64",
          "name": "version",
          "type": "uint64"
        }
      ],
      "name": "Initialized",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "previousOwner",
          "type": "address"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "newOwner",
          "type": "address"
        }
      ],
      "name": "OwnershipTransferred",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "arbitratorManager",
          "type": "address"
        }
      ],
      "name": "SetArbitratorManager",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "bytes32",
          "name": "txId",
          "type": "bytes32"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "dapp",
          "type": "address"
        }
      ],
      "name": "TransactionCompleted",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "bytes32",
          "name": "id",
          "type": "bytes32"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "dapp",
          "type": "address"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "arbitrator",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "deadline",
          "type": "uint256"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "depositFee",
          "type": "uint256"
        },
        {
          "indexed": false,
          "internalType": "address",
          "name": "compensationReceiver",
          "type": "address"
        }
      ],
      "name": "TransactionRegistered",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "bytes32",
          "name": "txId",
          "type": "bytes32"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "dapp",
          "type": "address"
        }
      ],
      "name": "UTXOsUploaded",
      "type": "event"
    },
    {
      "inputs": [],
      "name": "arbitratorManager",
      "outputs": [
        {
          "internalType": "contract IArbitratorManager",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "compensationManager",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes32",
          "name": "id",
          "type": "bytes32"
        }
      ],
      "name": "completeTransaction",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes32",
          "name": "id",
          "type": "bytes32"
        },
        {
          "internalType": "address",
          "name": "receivedCompensationAddress",
          "type": "address"
        }
      ],
      "name": "completeTransactionWithSlash",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "configManager",
      "outputs": [
        {
          "internalType": "contract ConfigManager",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "dappRegistry",
      "outputs": [
        {
          "internalType": "contract IDAppRegistry",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "deadline",
          "type": "uint256"
        },
        {
          "internalType": "address",
          "name": "arbitrator",
          "type": "address"
        }
      ],
      "name": "getRegisterTransactionFee",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "fee",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes32",
          "name": "txHash",
          "type": "bytes32"
        }
      ],
      "name": "getTransaction",
      "outputs": [
        {
          "components": [
            {
              "internalType": "address",
              "name": "dapp",
              "type": "address"
            },
            {
              "internalType": "address",
              "name": "arbitrator",
              "type": "address"
            },
            {
              "internalType": "uint256",
              "name": "startTime",
              "type": "uint256"
            },
            {
              "internalType": "uint256",
              "name": "deadline",
              "type": "uint256"
            },
            {
              "internalType": "bytes",
              "name": "btcTx",
              "type": "bytes"
            },
            {
              "internalType": "bytes32",
              "name": "btcTxHash",
              "type": "bytes32"
            },
            {
              "internalType": "enum DataTypes.TransactionStatus",
              "name": "status",
              "type": "uint8"
            },
            {
              "internalType": "uint256",
              "name": "depositedFee",
              "type": "uint256"
            },
            {
              "internalType": "bytes",
              "name": "signature",
              "type": "bytes"
            },
            {
              "internalType": "address",
              "name": "compensationReceiver",
              "type": "address"
            },
            {
              "internalType": "address",
              "name": "timeoutCompensationReceiver",
              "type": "address"
            },
            {
              "components": [
                {
                  "internalType": "bytes32",
                  "name": "txHash",
                  "type": "bytes32"
                },
                {
                  "internalType": "uint32",
                  "name": "index",
                  "type": "uint32"
                },
                {
                  "internalType": "bytes",
                  "name": "script",
                  "type": "bytes"
                },
                {
                  "internalType": "uint256",
                  "name": "amount",
                  "type": "uint256"
                }
              ],
              "internalType": "struct DataTypes.UTXO[]",
              "name": "utxos",
              "type": "tuple[]"
            },
            {
              "internalType": "bytes",
              "name": "script",
              "type": "bytes"
            }
          ],
          "internalType": "struct DataTypes.Transaction",
          "name": "",
          "type": "tuple"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes32",
          "name": "id",
          "type": "bytes32"
        }
      ],
      "name": "getTransactionById",
      "outputs": [
        {
          "components": [
            {
              "internalType": "address",
              "name": "dapp",
              "type": "address"
            },
            {
              "internalType": "address",
              "name": "arbitrator",
              "type": "address"
            },
            {
              "internalType": "uint256",
              "name": "startTime",
              "type": "uint256"
            },
            {
              "internalType": "uint256",
              "name": "deadline",
              "type": "uint256"
            },
            {
              "internalType": "bytes",
              "name": "btcTx",
              "type": "bytes"
            },
            {
              "internalType": "bytes32",
              "name": "btcTxHash",
              "type": "bytes32"
            },
            {
              "internalType": "enum DataTypes.TransactionStatus",
              "name": "status",
              "type": "uint8"
            },
            {
              "internalType": "uint256",
              "name": "depositedFee",
              "type": "uint256"
            },
            {
              "internalType": "bytes",
              "name": "signature",
              "type": "bytes"
            },
            {
              "internalType": "address",
              "name": "compensationReceiver",
              "type": "address"
            },
            {
              "internalType": "address",
              "name": "timeoutCompensationReceiver",
              "type": "address"
            },
            {
              "components": [
                {
                  "internalType": "bytes32",
                  "name": "txHash",
                  "type": "bytes32"
                },
                {
                  "internalType": "uint32",
                  "name": "index",
                  "type": "uint32"
                },
                {
                  "internalType": "bytes",
                  "name": "script",
                  "type": "bytes"
                },
                {
                  "internalType": "uint256",
                  "name": "amount",
                  "type": "uint256"
                }
              ],
              "internalType": "struct DataTypes.UTXO[]",
              "name": "utxos",
              "type": "tuple[]"
            },
            {
              "internalType": "bytes",
              "name": "script",
              "type": "bytes"
            }
          ],
          "internalType": "struct DataTypes.Transaction",
          "name": "",
          "type": "tuple"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "_arbitratorManager",
          "type": "address"
        },
        {
          "internalType": "address",
          "name": "_dappRegistry",
          "type": "address"
        },
        {
          "internalType": "address",
          "name": "_configManager",
          "type": "address"
        },
        {
          "internalType": "address",
          "name": "_compensationManager",
          "type": "address"
        }
      ],
      "name": "initialize",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes32",
          "name": "id",
          "type": "bytes32"
        }
      ],
      "name": "isAbleCompletedTransaction",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "owner",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "arbitrator",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "deadline",
          "type": "uint256"
        },
        {
          "internalType": "address",
          "name": "compensationReceiver",
          "type": "address"
        }
      ],
      "name": "registerTransaction",
      "outputs": [
        {
          "internalType": "bytes32",
          "name": "",
          "type": "bytes32"
        }
      ],
      "stateMutability": "payable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "renounceOwnership",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes32",
          "name": "id",
          "type": "bytes32"
        },
        {
          "internalType": "bytes",
          "name": "btcTx",
          "type": "bytes"
        },
        {
          "internalType": "bytes",
          "name": "script",
          "type": "bytes"
        },
        {
          "internalType": "address",
          "name": "timeoutCompensationReceiver",
          "type": "address"
        }
      ],
      "name": "requestArbitration",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "_arbitratorManager",
          "type": "address"
        }
      ],
      "name": "setArbitratorManager",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes32",
          "name": "id",
          "type": "bytes32"
        },
        {
          "internalType": "bytes",
          "name": "btcTxSignature",
          "type": "bytes"
        }
      ],
      "name": "submitArbitration",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes32",
          "name": "",
          "type": "bytes32"
        }
      ],
      "name": "transactions",
      "outputs": [
        {
          "internalType": "address",
          "name": "dapp",
          "type": "address"
        },
        {
          "internalType": "address",
          "name": "arbitrator",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "startTime",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "deadline",
          "type": "uint256"
        },
        {
          "internalType": "bytes",
          "name": "btcTx",
          "type": "bytes"
        },
        {
          "internalType": "bytes32",
          "name": "btcTxHash",
          "type": "bytes32"
        },
        {
          "internalType": "enum DataTypes.TransactionStatus",
          "name": "status",
          "type": "uint8"
        },
        {
          "internalType": "uint256",
          "name": "depositedFee",
          "type": "uint256"
        },
        {
          "internalType": "bytes",
          "name": "signature",
          "type": "bytes"
        },
        {
          "internalType": "address",
          "name": "compensationReceiver",
          "type": "address"
        },
        {
          "internalType": "address",
          "name": "timeoutCompensationReceiver",
          "type": "address"
        },
        {
          "internalType": "bytes",
          "name": "script",
          "type": "bytes"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes32",
          "name": "id",
          "type": "bytes32"
        }
      ],
      "name": "transferArbitrationFee",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "arbitratorFee",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "systemFee",
          "type": "uint256"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "newOwner",
          "type": "address"
        }
      ],
      "name": "transferOwnership",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes32",
          "name": "",
          "type": "bytes32"
        }
      ],
      "name": "txHashToId",
      "outputs": [
        {
          "internalType": "bytes32",
          "name": "",
          "type": "bytes32"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes32",
          "name": "id",
          "type": "bytes32"
        },
        {
          "components": [
            {
              "internalType": "bytes32",
              "name": "txHash",
              "type": "bytes32"
            },
            {
              "internalType": "uint32",
              "name": "index",
              "type": "uint32"
            },
            {
              "internalType": "bytes",
              "name": "script",
              "type": "bytes"
            },
            {
              "internalType": "uint256",
              "name": "amount",
              "type": "uint256"
            }
          ],
          "internalType": "struct DataTypes.UTXO[]",
          "name": "utxos",
          "type": "tuple[]"
        }
      ],
      "name": "uploadUTXOs",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    }
  ]`

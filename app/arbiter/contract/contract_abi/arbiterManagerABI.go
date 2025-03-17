// Copyright (c) 2025 The bel2 developers

package contract_abi

const ArbiterManagerABI = `[
{
      "inputs": [
        {
          "internalType": "address",
          "name": "arbitrator",
          "type": "address"
        }
      ],
      "name": "getArbitratorInfoExt",
      "outputs": [
        {
          "components": [
            {
              "internalType": "uint256",
              "name": "currentBTCFeeRate",
              "type": "uint256"
            }
          ],
          "internalType": "struct DataTypes.ArbitratorInfoExt",
          "name": "",
          "type": "tuple"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
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
          "internalType": "address",
          "name": "arbitrator",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "deadline",
          "type": "uint256"
        }
      ],
      "name": "ArbitratorDeadlineUpdated",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "arbitrator",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "feeRate",
          "type": "uint256"
        }
      ],
      "name": "ArbitratorFeeRateUpdated",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "arbitrator",
          "type": "address"
        }
      ],
      "name": "ArbitratorFrozen",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "arbitrator",
          "type": "address"
        }
      ],
      "name": "ArbitratorPaused",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "arbitrator",
          "type": "address"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "operator",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "address",
          "name": "revenueAddress",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "string",
          "name": "btcAddress",
          "type": "string"
        },
        {
          "indexed": false,
          "internalType": "bytes",
          "name": "btcPubKey",
          "type": "bytes"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "feeRate",
          "type": "uint256"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "deadline",
          "type": "uint256"
        }
      ],
      "name": "ArbitratorRegistered",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "arbitrator",
          "type": "address"
        },
        {
          "indexed": true,
          "internalType": "bytes32",
          "name": "transactionId",
          "type": "bytes32"
        }
      ],
      "name": "ArbitratorReleased",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "arbitrator",
          "type": "address"
        }
      ],
      "name": "ArbitratorTerminatedWithSlash",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "arbitrator",
          "type": "address"
        }
      ],
      "name": "ArbitratorUnpaused",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "arbitrator",
          "type": "address"
        },
        {
          "indexed": true,
          "internalType": "bytes32",
          "name": "transactionId",
          "type": "bytes32"
        }
      ],
      "name": "ArbitratorWorking",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "oldManager",
          "type": "address"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "newManager",
          "type": "address"
        }
      ],
      "name": "CompensationManagerUpdated",
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
          "name": "transactionManager",
          "type": "address"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "compensationManager",
          "type": "address"
        }
      ],
      "name": "InitializedManager",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "oldNFTContract",
          "type": "address"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "newNFTContract",
          "type": "address"
        }
      ],
      "name": "NFTContractUpdated",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "arbitrator",
          "type": "address"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "operator",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "bytes",
          "name": "btcPubKey",
          "type": "bytes"
        },
        {
          "indexed": false,
          "internalType": "string",
          "name": "btcAddress",
          "type": "string"
        }
      ],
      "name": "OperatorSet",
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
          "name": "arbitrator",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "address",
          "name": "ethAddress",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "bytes",
          "name": "btcPubKey",
          "type": "bytes"
        },
        {
          "indexed": false,
          "internalType": "string",
          "name": "btcAddress",
          "type": "string"
        }
      ],
      "name": "RevenueAddressesSet",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "arbitrator",
          "type": "address"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "assetAddress",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        },
        {
          "indexed": false,
          "internalType": "uint256[]",
          "name": "nftTokenIds",
          "type": "uint256[]"
        }
      ],
      "name": "StakeAdded",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "arbitrator",
          "type": "address"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "assetAddress",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        }
      ],
      "name": "StakeWithdrawn",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "oldManager",
          "type": "address"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "newManager",
          "type": "address"
        }
      ],
      "name": "TransactionManagerUpdated",
      "type": "event"
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
      "inputs": [
        {
          "internalType": "address",
          "name": "arbitrator",
          "type": "address"
        }
      ],
      "name": "frozenArbitrator",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "arbitratorAddress",
          "type": "address"
        }
      ],
      "name": "getArbitratorInfo",
      "outputs": [
        {
          "components": [
            {
              "internalType": "address",
              "name": "arbitrator",
              "type": "address"
            },
            {
              "internalType": "bool",
              "name": "paused",
              "type": "bool"
            },
            {
              "internalType": "uint256",
              "name": "currentFeeRate",
              "type": "uint256"
            },
            {
              "internalType": "bytes32",
              "name": "activeTransactionId",
              "type": "bytes32"
            },
            {
              "internalType": "uint256",
              "name": "ethAmount",
              "type": "uint256"
            },
            {
              "internalType": "address",
              "name": "erc20Token",
              "type": "address"
            },
            {
              "internalType": "address",
              "name": "nftContract",
              "type": "address"
            },
            {
              "internalType": "uint256[]",
              "name": "nftTokenIds",
              "type": "uint256[]"
            },
            {
              "internalType": "address",
              "name": "operator",
              "type": "address"
            },
            {
              "internalType": "bytes",
              "name": "operatorBtcPubKey",
              "type": "bytes"
            },
            {
              "internalType": "string",
              "name": "operatorBtcAddress",
              "type": "string"
            },
            {
              "internalType": "uint256",
              "name": "deadLine",
              "type": "uint256"
            },
            {
              "internalType": "bytes",
              "name": "revenueBtcPubKey",
              "type": "bytes"
            },
            {
              "internalType": "string",
              "name": "revenueBtcAddress",
              "type": "string"
            },
            {
              "internalType": "address",
              "name": "revenueETHAddress",
              "type": "address"
            },
            {
              "internalType": "uint256",
              "name": "lastSubmittedWorkTime",
              "type": "uint256"
            }
          ],
          "internalType": "struct DataTypes.ArbitratorInfo",
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
          "name": "arbitrator",
          "type": "address"
        }
      ],
      "name": "getAvailableStake",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
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
        }
      ],
      "name": "getTotalNFTStakeValue",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "_transactionManager",
          "type": "address"
        },
        {
          "internalType": "address",
          "name": "_compensationManager",
          "type": "address"
        }
      ],
      "name": "initTransactionAndCompensationManager",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "_configManager",
          "type": "address"
        },
        {
          "internalType": "address",
          "name": "_nftContract",
          "type": "address"
        },
        {
          "internalType": "address",
          "name": "_nftInfo",
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
          "internalType": "address",
          "name": "arbitratorAddress",
          "type": "address"
        }
      ],
      "name": "isActiveArbitrator",
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
      "inputs": [
        {
          "internalType": "address",
          "name": "arbitrator",
          "type": "address"
        }
      ],
      "name": "isConfigModifiable",
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
      "inputs": [
        {
          "internalType": "address",
          "name": "arbitrator",
          "type": "address"
        }
      ],
      "name": "isFrozenStatus",
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
      "inputs": [
        {
          "internalType": "address",
          "name": "arbitrator",
          "type": "address"
        },
        {
          "internalType": "address",
          "name": "operator",
          "type": "address"
        }
      ],
      "name": "isOperatorOf",
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
      "inputs": [
        {
          "internalType": "address",
          "name": "arbitrator",
          "type": "address"
        }
      ],
      "name": "isPaused",
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
      "name": "nftContract",
      "outputs": [
        {
          "internalType": "contract IERC721",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "nftInfo",
      "outputs": [
        {
          "internalType": "contract IBNFTInfo",
          "name": "",
          "type": "address"
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
      "inputs": [],
      "name": "pause",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "string",
          "name": "defaultBtcAddress",
          "type": "string"
        },
        {
          "internalType": "bytes",
          "name": "defaultBtcPubKey",
          "type": "bytes"
        },
        {
          "internalType": "uint256",
          "name": "feeRate",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "deadline",
          "type": "uint256"
        }
      ],
      "name": "registerArbitratorByStakeETH",
      "outputs": [],
      "stateMutability": "payable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256[]",
          "name": "tokenIds",
          "type": "uint256[]"
        },
        {
          "internalType": "string",
          "name": "defaultBtcAddress",
          "type": "string"
        },
        {
          "internalType": "bytes",
          "name": "defaultBtcPubKey",
          "type": "bytes"
        },
        {
          "internalType": "uint256",
          "name": "feeRate",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "deadline",
          "type": "uint256"
        }
      ],
      "name": "registerArbitratorByStakeNFT",
      "outputs": [],
      "stateMutability": "nonpayable",
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
          "internalType": "bytes32",
          "name": "transactionId",
          "type": "bytes32"
        }
      ],
      "name": "releaseArbitrator",
      "outputs": [],
      "stateMutability": "nonpayable",
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
          "internalType": "uint256",
          "name": "deadline",
          "type": "uint256"
        }
      ],
      "name": "setArbitratorDeadline",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "feeRate",
          "type": "uint256"
        }
      ],
      "name": "setArbitratorFeeRate",
      "outputs": [],
      "stateMutability": "nonpayable",
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
          "internalType": "bytes32",
          "name": "transactionId",
          "type": "bytes32"
        }
      ],
      "name": "setArbitratorWorking",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "_compensationManager",
          "type": "address"
        }
      ],
      "name": "setCompensationManager",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "_nftContract",
          "type": "address"
        }
      ],
      "name": "setNFTContract",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "operator",
          "type": "address"
        },
        {
          "internalType": "bytes",
          "name": "btcPubKey",
          "type": "bytes"
        },
        {
          "internalType": "string",
          "name": "btcAddress",
          "type": "string"
        }
      ],
      "name": "setOperator",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "ethAddress",
          "type": "address"
        },
        {
          "internalType": "bytes",
          "name": "btcPubKey",
          "type": "bytes"
        },
        {
          "internalType": "string",
          "name": "btcAddress",
          "type": "string"
        }
      ],
      "name": "setRevenueAddresses",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "_transactionManager",
          "type": "address"
        }
      ],
      "name": "setTransactionManager",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "stakeETH",
      "outputs": [],
      "stateMutability": "payable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256[]",
          "name": "tokenIds",
          "type": "uint256[]"
        }
      ],
      "name": "stakeNFT",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "arbitrator",
          "type": "address"
        }
      ],
      "name": "terminateArbitratorWithSlash",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "transactionManager",
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
      "inputs": [],
      "name": "unpause",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "unstake",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "zeroAddress",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    }
  ]`

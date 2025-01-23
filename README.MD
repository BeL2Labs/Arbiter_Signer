Arbiter Signer
=============
[![Go Report Card](https://goreportcard.com/badge/github.com/BeL2Labs/Arbiter_Signer)](https://goreportcard.com/report/github.com/BeL2Labs/Arbiter_Signer)
[![Build Status](https://github.com/BeL2Labs/Arbiter_Signer/workflows/Build%20and%20Test/badge.svg)](https://github.com/BeL2Labs/Arbiter_Signer/actions)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](http://copyfree.org)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/BeL2Labs/Arbiter_Signer)

## Introduction

The Arbiter Signer is used to provide signing services for the bel2 loan arbiter.

## Requirements

[Go](http://golang.org) 1.20 or newer.

## Build and run
### 1. build arbiter

use **main** branch to build:

```shell
$ git clone git@github.com:BeL2Labs/Arbiter_Signer.git

$ cd Arbiter_Signer

$ go mod tidy

$ make
```

### 2. prepare config file

1. **listener**: listen to esc arbiter contract event

2. **signer**: sign request_arbitration_btc_tx and submit to esc arbiter contract 

3. **network**: use **mainnet** or **testnet** for btc address generation

4. **escStartHeigh**t: listen from this height to get esc arbiter contract event

5. **escArbiterContractAddress**: arbiter contract address, if arbiter contract address changed, need to update it

6. **escArbiterAddress**: arbiter wallet address, not operator address

7. **dataPath**: defines the path to record event file

8. **keyFilePath**: defines the path of arbiter BTC keystore and ESC keystore


   ***config.yaml example***：

   ```yaml
   # Chain api
   chain:
     esc: "https://api.elastos.io/esc"

   # Arbiter
   arbiter:
     listener: true,
     signer: true,
     network: "mainnet"
     escStartHeight: 21205821
     escArbiterContractAddress: "0xD206be45b53Fa5C2065049c7a70B1aa1755a9475"
     escArbiterManagerContractAddress: "0x54eE4fc8951A936b6AA02079B76c497c0471c52A"
     escConfigManagerContractAddress: "0x4421c63241A262C423277FFA82C376953072d25f"
     escArbiterAddress: "0x0262aB0ED65373cC855C34529fDdeAa0e686D913"
     dataPath: "~/loan_arbiter/data"
     keyFilePath: "~/loan_arbiter/keys"
   ```



### 3. prepare keystore to keyFilePath

1. create **btcKey** **escKey** and put it into **keyFilePath**
btcKey is used to sign request_arbitration_btc_tx
escKey is used to submit arbitration signature to esc arbiter contract, gas fee(esc ELA) is needed
 
2. keysotre file need to set hex encoded private key

   use [encrypt.sh](https://github.com/BeL2Labs/Arbiter_Signer/blob/main/docs/encrypt.sh) to create btc keystore:
      ```shell
      ~ chmod a+x encrypt.sh
      ~ ./encrypt.sh
      ~ > YOUR_HEX_PRIV_KEY
      ~ > YOUR_PASSWORD
      ~ > btcKey
      ```
    use [encrypt.sh](https://github.com/BeL2Labs/Arbiter_Signer/blob/main/docs/encrypt.sh) to create esc keystore:
      ```shell
      ~ chmod a+x encrypt.sh
      ~ ./encrypt.sh
      ~ > YOUR_HEX_PRIV_KEY
      ~ > YOUR_PASSWORD
      ~ > escKey
      ```

### 4. run arbiter

run arbiter with go file
```shell
go run app/arbiter/main.go
```

run arbiter with default manifest file
```shell
./arbiter
```

run arbiter with custom manifest file
```shell
./arbiter --gf.gcfg.file=config.yaml
```

## Automated Deployment
The automated deployment process can be referenced in the documentation: 
[deploy_loan_arbiter.md](https://github.com/BeL2Labs/Arbiter_Signer/blob/main/docs/deploy_loan_arbiter.md)

## License
arbiter signer is licensed under the [copyfree](http://copyfree.org)  MIT License.
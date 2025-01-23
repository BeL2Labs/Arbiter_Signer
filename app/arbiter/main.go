// Copyright (c) 2025 The bel2 developers

package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	"github.com/BeL2Labs/Arbiter_Signer/app/arbiter/arbiter"
	"github.com/BeL2Labs/Arbiter_Signer/app/arbiter/config"
)

func main() {
	passwordPtr := flag.String("p", "", "Specify the password")
	flag.Parse()
	var password string
	if *passwordPtr == "" {
		fmt.Print("please input key password: ")
		_, err := fmt.Scanln(&password)
		if err != nil {
			fmt.Println("read password error:", err)
			os.Exit(1)
		}
	} else {
		password = *passwordPtr
	}

	if len(os.Args) > 1 {
		operation := os.Args[1]
		switch strings.ToLower(operation) {
		case "getpk", "gen":
			myPriKey := os.Args[2]
			pk, err := arbiter.GetPubKey(myPriKey)
			if err != nil {
				fmt.Println("need hex encoded format private key, err:", err)
				return
			}
			fmt.Println("publicKey:", pk)
			return
		}
	}

	ctx := gctx.New()
	var wg sync.WaitGroup
	wg.Add(1)
	// start arbiter
	g.Log().Info(ctx, "Starting arbiter...")
	arb := arbiter.NewArbiter(ctx, initConfig(ctx), password)
	arb.Start()
	wg.Wait()
}

func initConfig(ctx context.Context) *config.Config {
	network, err := g.Cfg().Get(ctx, "arbiter.network")
	if err != nil {
		g.Log().Error(ctx, "get network config err:", err)
		os.Exit(1)
	}
	signer, err := g.Cfg().Get(ctx, "arbiter.signer")
	if err != nil {
		g.Log().Error(ctx, "get signer err:", err)
		os.Exit(1)
	}
	listener, err := g.Cfg().Get(ctx, "arbiter.listener")
	if err != nil {
		g.Log().Error(ctx, "get listener config err:", err)
		os.Exit(1)
	}
	http, err := g.Cfg().Get(ctx, "chain.esc")
	if err != nil {
		g.Log().Error(ctx, "get http config err:", err)
		os.Exit(1)
	}
	escStartHeight, err := g.Cfg().Get(ctx, "arbiter.escStartHeight")
	if err != nil {
		g.Log().Error(ctx, "get escStartHeight config err:", err)
		os.Exit(1)
	}
	escArbiterContractAddress, err := g.Cfg().Get(ctx, "arbiter.escArbiterContractAddress")
	if err != nil {
		g.Log().Error(ctx, "get escArbiterAddress config err:", err)
		os.Exit(1)
	}
	escArbiterManagerAddress, err := g.Cfg().Get(ctx, "arbiter.escArbiterManagerContractAddress")
	if err != nil {
		g.Log().Error(ctx, "get escArbiterManagerAddress config err:", err)
		os.Exit(1)
	}
	escConfigManagerAddress, err := g.Cfg().Get(ctx, "arbiter.escConfigManagerContractAddress")
	if err != nil {
		g.Log().Error(ctx, "get escConfigManagerAddress config err:", err)
		os.Exit(1)
	}
	escArbiterAddress, err := g.Cfg().Get(ctx, "arbiter.escArbiterAddress")
	if err != nil {
		g.Log().Error(ctx, "get escArbiterAddress config err:", err)
		os.Exit(1)
	}
	gDataPath, err := g.Cfg().Get(ctx, "arbiter.dataPath")
	if err != nil {
		g.Log().Error(ctx, "get dataPath config err:", err)
		os.Exit(1)
	}
	gKeyFilePath, err := g.Cfg().Get(ctx, "arbiter.keyFilePath")
	if err != nil {
		g.Log().Error(ctx, "get keyFilePath config err:", err)
		os.Exit(1)
	}
	dataPath := getExpandedPath(gDataPath.String())
	keyFilePath := getExpandedPath(gKeyFilePath.String())

	g.Log().Info(ctx, "btcCreator:", signer)
	g.Log().Info(ctx, "listener:", listener)
	g.Log().Info(ctx, "http:", http)
	g.Log().Info(ctx, "escStartHeight:", escStartHeight)
	g.Log().Info(ctx, "escArbiterContractAddress:", escArbiterContractAddress)
	g.Log().Info(ctx, "escArbiterManagerAddress:", escArbiterManagerAddress)
	g.Log().Info(ctx, "escArbiterAddress:", escArbiterAddress)
	g.Log().Info(ctx, "dataPath:", dataPath)
	g.Log().Info(ctx, "keyFilePath:", keyFilePath)

	// if want to submit to ESC contract successfully, need to use esc ela as gas.
	escKeyFilePath := gfile.Join(keyFilePath, "escKey")
	arbiterKeyFilePath := gfile.Join(keyFilePath, "btcKey")
	logPath := gfile.Join(dataPath, "logs/")
	loanPath := gfile.Join(dataPath, "loan/")
	loanNeedSignReqPath := gfile.Join(loanPath, "request/")
	loanNeedSignFailedPath := gfile.Join(loanPath, "failed/")
	loanNeedSignSignedPath := gfile.Join(loanPath, "signed/")
	LoanSignedEventPath := gfile.Join(dataPath, "loan_signed_event/")

	return &config.Config{
		Network:                          network.String(),
		Signer:                           signer.Bool(),
		Listener:                         listener.Bool(),
		Http:                             http.String(),
		ESCStartHeight:                   escStartHeight.Uint64(),
		ESCArbiterContractAddress:        escArbiterContractAddress.String(),
		ESCArbiterManagerContractAddress: escArbiterManagerAddress.String(),
		ESCConfigManagerContractAddress:  escConfigManagerAddress.String(),
		ESCArbiterAddress:                escArbiterAddress.String(),

		DataDir:            dataPath,
		EscKeyFilePath:     escKeyFilePath,
		ArbiterKeyFilePath: arbiterKeyFilePath,

		LoanNeedSignReqPath:    loanNeedSignReqPath,
		LoanNeedSignFailedPath: loanNeedSignFailedPath,
		LoanNeedSignSignedPath: loanNeedSignSignedPath,
		LoanSignedEventPath:    LoanSignedEventPath,
		LoanLogPath:            logPath,
	}
}

func getExpandedPath(path string) string {
	if len(path) > 0 && path[0] == '~' {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error getting home directory:", err)
			return path
		}
		path = filepath.Join(homeDir, path[2:])
	}
	return path
}

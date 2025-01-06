// Copyright (c) 2025 The bel2 developers

package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/BeL2Labs/Arbiter_Signer/app/email/config"
	"github.com/BeL2Labs/Arbiter_Signer/app/email/email"

	"github.com/gogf/gf/os/gctx"
	"github.com/gogf/gf/v2/frame/g"
	"gopkg.in/gomail.v2"
)

func main() {
	ctx := gctx.New()
	cfg := initConfig(ctx)
	var wg sync.WaitGroup
	wg.Add(1)
	g.Log().Info(ctx, "Starting email service...")
	eml := email.NewEmail(ctx, cfg)
	eml.Start()
	g.Log().Info(ctx, "Email service started")
	wg.Wait()
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

func initConfig(ctx context.Context) *config.Config {
	// chain
	http, err := g.Cfg().Get(ctx, "chain.esc")
	if err != nil {
		g.Log().Error(ctx, "get esc api config err:", err)
		os.Exit(1)
	}

	// arbiter
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
	loanNeedSignReqPath, err := g.Cfg().Get(ctx, "arbiter.loanNeedSignReqPath")
	if err != nil {
		g.Log().Error(ctx, "get loanNeedSignReqPath config err:", err)
		os.Exit(1)
	}
	arbiters, err := g.Cfg().Get(ctx, "arbiter.arbiters")
	if err != nil {
		g.Log().Error(ctx, "get arbiters config err:", err)
		os.Exit(1)
	}

	// email
	host, err := g.Cfg().Get(ctx, "email.host")
	if err != nil {
		g.Log().Error(ctx, "get host config err:", err)
		os.Exit(1)
	}
	port, err := g.Cfg().Get(ctx, "email.port")
	if err != nil {
		g.Log().Error(ctx, "get port config err:", err)
		os.Exit(1)
	}
	from, err := g.Cfg().Get(ctx, "email.from")
	if err != nil {
		g.Log().Error(ctx, "get from config err:", err)
		os.Exit(1)
	}
	user, err := g.Cfg().Get(ctx, "email.user")
	if err != nil {
		g.Log().Error(ctx, "get user config err:", err)
		os.Exit(1)
	}
	password, err := g.Cfg().Get(ctx, "email.password")
	if err != nil {
		g.Log().Error(ctx, "get password config err:", err)
		os.Exit(1)
	}
	to, err := g.Cfg().Get(ctx, "email.to")
	if err != nil {
		g.Log().Error(ctx, "get to config err:", err)
		os.Exit(1)
	}
	emailLogPath, err := g.Cfg().Get(ctx, "email.emailLogPath")
	if err != nil {
		g.Log().Error(ctx, "get emailLogPath config err:", err)
		os.Exit(1)
	}
	dataPath, err := g.Cfg().Get(ctx, "email.dataPath")
	if err != nil {
		g.Log().Error(ctx, "get dataPath config err:", err)
		os.Exit(1)
	}
	duration, err := g.Cfg().Get(ctx, "email.duration")
	if err != nil {
		g.Log().Error(ctx, "get duration config err:", err)
		os.Exit(1)
	}

	return &config.Config{
		Http: http.String(),

		ESCStartHeight:            escStartHeight.Uint64(),
		ESCArbiterContractAddress: escArbiterContractAddress.String(),
		LoanNeedSignReqPath:       getExpandedPath(loanNeedSignReqPath.String()),
		ArbiterAddresses:          arbiters.Strings(),

		Host:         host.String(),
		Port:         port.Int(),
		From:         from.String(),
		User:         user.String(),
		Password:     password.String(),
		To:           to.Strings(),
		EmailLogPath: getExpandedPath(emailLogPath.String()),
		DataPath:     getExpandedPath(dataPath.String()),
		Duration:     duration.Int(),
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

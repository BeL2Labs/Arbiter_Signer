// Copyright (c) 2025 The bel2 developers

package config

type Config struct {
	// chain
	Http string

	// arbiter
	ESCStartHeight            uint64
	ESCArbiterContractAddress string
	LoanNeedSignReqPath       string
	ArbiterAddresses          []string

	// email
	Host         string
	Port         int
	From         string
	User         string
	Password     string
	To           []string
	EmailLogPath string
	DataPath     string
	Duration     int
}

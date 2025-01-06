// Copyright (c) 2025 The bel2 developers

package contract

type Config struct {
	Http string

	ESCArbiterAddresses              map[string]struct{}
	ESCArbiterContractAddress        string
	ESCArbiterManagerContractAddress string

	DataDir             string
	LoanNeedSignReqPath string
	LoanSignedEventPath string
}

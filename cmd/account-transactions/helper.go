package main

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
)

var prefixName = "account-transactions"

// UnmarshalEnv unmarshals the environmental variables into the given destination interface
func UnmarshalEnv(r interface{}) error {
	if err := envconfig.Process(os.Getenv(prefixName), r); err != nil {
		return fmt.Errorf("failed to process env: %w", err)
	}

	return nil
}

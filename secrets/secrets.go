package secrets

import (
	"errors"
	"os"
)

func WriteApplicationCredentialsToFile() error {
	creds := os.Getenv("APP_CREDENTIALS_CONTENTS")
	if creds == "" {
		return errors.New("APP_CREDENTIALS_CONTENTS not set")
	}

	secretsFile, err := os.Create("admin_secrets.json")
	if err != nil {
		return errors.New("failed to create secrets file")
	}

	_, err = secretsFile.Write([]byte(creds))
	if err != nil {
		return errors.New("failed to write credentials to secrets file")
	}

	return nil
}

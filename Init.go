package envenc

import (
	"errors"
)

func Init(vaultFilePath string, vaultPassword string) error {
	if vaultFilePath == "" {
		return errors.New("vault file path cannot be empty")
	}

	if vaultPassword == "" {
		return errors.New("vault password cannot be empty")
	}

	if fileExists(vaultFilePath) {
		return errors.New("vault file already exists")
	}

	vault := newStore()

	return vaultSave(vaultFilePath, vaultPassword, *vault)
}

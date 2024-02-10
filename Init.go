package envenc

import (
	"errors"
)

func Init(vaultFilePath string, vaultPassword string) error {

	if fileExists(vaultFilePath) {
		return errors.New("vault file already exists")
	}

	vault := newStore()

	return vaultSave(vaultFilePath, vaultPassword, *vault)
}

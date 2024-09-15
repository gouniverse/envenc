package envenc

import "errors"

func KeyGet(vaultFilePath string, vaultPassword string, key string) (string, error) {
	vault, err := vaultOpenFromFile(vaultFilePath, vaultPassword)

	if err != nil {
		return "", err
	}

	if vault == nil {
		return "", errors.New("vault is nil")
	}

	value := vault.Get(key)

	return value, nil
}

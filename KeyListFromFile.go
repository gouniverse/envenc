package envenc

import "strings"

func KeyListFromFile(vaultFilePath string, vaultPassword string) (map[string]string, error) {
	vaultPassword = strings.TrimSpace(vaultPassword)
	vault, err := vaultOpenFromFile(vaultFilePath, vaultPassword)

	if err != nil {
		return nil, err
	}

	data := vault.Data()

	return data, nil
}

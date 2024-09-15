package envenc

import "strings"

func KeyListFromString(vaultString string, vaultPassword string) (map[string]string, error) {
	vaultPassword = strings.TrimSpace(vaultPassword)
	vault, err := vaultOpenFromString(vaultString, vaultPassword)

	if err != nil {
		return nil, err
	}

	data := vault.Data()

	return data, nil
}

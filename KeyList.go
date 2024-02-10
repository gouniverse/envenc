package envenc

import "strings"

func EnvList(vaultFilePath string, vaultPassword string) (map[string]string, error) {
	vaultPassword = strings.TrimSpace(vaultPassword)
	vault, err := vaultOpen(vaultFilePath, vaultPassword)

	if err != nil {
		return nil, err
	}

	data := vault.Data()

	return data, nil
}

package envenc

import "strings"

// KeyListFromFile lists all keys in the vault
//
// Buisiness logic:
//   - Open the vault file
//   - Get the keys from the vault
//
// Parameters:
//   - vaultFilePath: The path to the vault file
//   - vaultPassword: The password to use for the vault
//
// Returns:
//   - map[string]string: A map of keys and their values
//   - error: An error if the keys could not be retrieved
func KeyListFromFile(vaultFilePath string, vaultPassword string) (map[string]string, error) {
	vaultPassword = strings.TrimSpace(vaultPassword)
	vault, err := vaultOpenFromFile(vaultFilePath, vaultPassword)

	if err != nil {
		return nil, err
	}

	data := vault.Data()

	return data, nil
}

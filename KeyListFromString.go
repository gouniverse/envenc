package envenc

import "strings"

// KeyListFromString lists all keys in the vault
//
// Buisiness logic:
//   - Open the vault from string
//   - Get the keys from the vault
//
// Parameters:
//   - vaultString: The string representation of the vault
//   - vaultPassword: The password to use for the vault
// Returns:
//   - map[string]string: A map of keys and their values
//   - error: An error if the keys could not be retrieved
func KeyListFromString(vaultString string, vaultPassword string) (map[string]string, error) {
	vaultPassword = strings.TrimSpace(vaultPassword)
	vault, err := vaultOpenFromString(vaultString, vaultPassword)

	if err != nil {
		return nil, err
	}

	data := vault.Data()

	return data, nil
}

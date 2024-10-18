package envenc

import "errors"

// KeyExists checks if a key exists in the vault
//
// Buisiness logic:
//   - Open the vault file
//   - Check if the key exists in the vault
//
// Parameters:
//   - vaultFilePath: The path to the vault file
//   - vaultPassword: The password to use for the vault
//   - key: The name of the key to check
//
// Returns:
//   - bool: True if the key exists, false otherwise
//   - error: An error if the key could not be retrieved
func KeyExists(vaultFilePath string, vaultPassword string, key string) (bool, error) {
	store, err := vaultOpenFromFile(vaultFilePath, vaultPassword)

	if err != nil {
		return false, err
	}

	if store == nil {
		return false, errors.New("store is nil")
	}

	data := store.Data()

	if _, ok := data[key]; ok {
		return true, nil
	}

	return false, nil
}

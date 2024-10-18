package envenc

import "errors"

// KeyGet gets a key from the vault
//
// Buisiness logic:
//   - Open the vault file
//   - Get the key from the vault
//   - Save the vault file
//
// Parameters:
//   - vaultFilePath: The path to the vault file
//   - vaultPassword: The password to use for the vault
//   - keyName: The name of the key to get
//
// Returns:
//   - string: The value of the key
//   - error: An error if the key could not be retrieved
func KeyGet(vaultFilePath string, vaultPassword string, keyName string) (string, error) {
	vault, err := vaultOpenFromFile(vaultFilePath, vaultPassword)

	if err != nil {
		return "", err
	}

	if vault == nil {
		return "", errors.New("vault is nil")
	}

	value := vault.Get(keyName)

	return value, nil
}

package envenc

// KeySet sets a key in the vault
//
// Buisiness logic:
//   - Open the vault file
//   - Set the key in the vault (if it doesn't exist, create it, otherwise update it)
//   - Save the vault file
//
// Parameters:
//   - vaultFilePath: The path to the vault file
//   - vaultPassword: The password to use for the vault
//   - keyName: The name of the	key to set
//   - keyValue: The value of the key to set
//
// Returns:
//   - error: An error if the key could not be set
func KeySet(vaultFilePath string, vaultPassword string, keyName string, keyValue string) error {
	vault, err := vaultOpenFromFile(vaultFilePath, vaultPassword)

	if err != nil {
		return err
	}

	vault.Set(keyName, keyValue)

	return vaultSave(vaultFilePath, vaultPassword, *vault)
}

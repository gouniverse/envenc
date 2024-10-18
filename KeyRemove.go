package envenc

// KeyRemove removes a key from the vault
//
// Buisiness logic:
//   - Open the vault file
//   - Remove the key from the vault
//   - Save the vault file
//
// Parameters:
//   - vaultFilePath: The path to the vault file
//   - vaultPassword: The password to use for the vault
//   - keyName: The name of the key to remove
//
// Returns:
//   - error: An error if the key could not be removed
func KeyRemove(vaultFilePath string, vaultPassword string, keyName string) error {
	vault, err := vaultOpenFromFile(vaultFilePath, vaultPassword)

	if err != nil {
		return err
	}

	vault.Remove(keyName)

	return vaultSave(vaultFilePath, vaultPassword, *vault)
}

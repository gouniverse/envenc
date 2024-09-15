package envenc

func KeySet(vaultFilePath string, vaultPassword string, key string, value string) error {
	vault, err := vaultOpenFromFile(vaultFilePath, vaultPassword)

	if err != nil {
		return err
	}

	vault.Set(key, value)

	return vaultSave(vaultFilePath, vaultPassword, *vault)
}

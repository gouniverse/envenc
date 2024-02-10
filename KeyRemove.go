package envenc

func KeyRemove(vaultFilePath string, vaultPassword string, key string) error {
	vault, err := vaultOpen(vaultFilePath, vaultPassword)

	if err != nil {
		return err
	}

	vault.Remove(key)

	return vaultSave(vaultFilePath, vaultPassword, *vault)
}

package envenc

import "errors"

func KeyExists(vaultFilePath string, vaultPassword string, key string) (bool, error) {
	store, err := vaultOpen(vaultFilePath, vaultPassword)

	if err != nil {
		return false, err
	}

	if store == nil {
		return false, errors.New("store is nil")
	}

	value := store.Get(key)

	if value == "" {
		return true, nil
	}

	return false, nil
}

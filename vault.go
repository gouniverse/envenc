package envenc

import "errors"

func vaultOpenFromFile(vaultFilePath string, vaultPassword string) (*store, error) {
	if !fileExists(vaultFilePath) {
		return nil, errors.New("vault file does not exist")
	}

	encString, err := fileGetContents(vaultFilePath)

	if err != nil {
		return nil, err
	}

	if encString == "" {
		return nil, errors.New("vault file is empty")
	}

	return vaultOpenFromString(encString, vaultPassword)
}

func vaultOpenFromString(encodedString string, vaultPassword string) (*store, error) {
	if encodedString == "" {
		return nil, errors.New("vault is empty")
	}

	decodedString, err := Decrypt(encodedString, vaultPassword)

	if err != nil {
		return nil, err
	}

	if decodedString == "" {
		return nil, errors.New("decryption failed")
	}

	st, err := newStoreFromJSON(decodedString)

	if err != nil {
		return nil, err
	}

	return st, nil
}

func vaultSave(vaultFilePath string, password string, st store) error {
	storeJSON, err := st.ToJSON()

	if err != nil {
		return err
	}

	encodedString, err := Encrypt(storeJSON, password)

	if err != nil {
		return err
	}

	return filePutContents(vaultFilePath, encodedString, 0644)
}

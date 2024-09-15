package envenc

import (
	"errors"
	"os"

	"github.com/gouniverse/crypto"
)

// fileGetContents reads entire file into a string
func fileGetContents(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	return string(data), err
}

// fileExists checks if a file exists
func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)

	return os.IsExist(err)
}

// filePutContents adds content to file
func filePutContents(filename string, data string, mode os.FileMode) error {
	return os.WriteFile(filename, []byte(data), mode)
}

func strPadLeft(input string, padLength int, padString string) string {
	output := ""
	inputLen := len(input)
	if inputLen >= padLength {
		return input
	}
	ll := padLength - inputLen
	for i := 1; i <= ll; i = i + len(padString) {
		output += padString
	}
	return output + input
}

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

func vaultOpenFromString(encString string, vaultPassword string) (*store, error) {
	if encString == "" {
		return nil, errors.New("vault is empty")
	}

	aesEncString, err := crypto.AESFortifiedDecrypt(encString, vaultPassword)

	if err != nil {
		return nil, err
	}

	finalDecString, err := crypto.XorFortifiedDecrypt(aesEncString, vaultPassword)

	if err != nil {
		return nil, err
	}

	if finalDecString == "" {
		return nil, errors.New("decryption failed")
	}

	st, err := newStoreFromJSON(finalDecString)

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

	xorEncString := crypto.XorFortifiedEncrypt(storeJSON, password)
	finalEncString, err := crypto.AESFortifiedEncrypt(xorEncString, password)

	if err != nil {
		return err
	}

	return filePutContents(vaultFilePath, finalEncString, 0644)
}

package envenc

import (
	"os"
	"testing"
)

func TestVaultSave(t *testing.T) {
	// Arrange
	vaultFilePath := os.TempDir() + string(os.PathSeparator) + strRandom(10) + ".env.vault"
	password := "password"
	st := store{}

	// Act
	err := vaultSave(vaultFilePath, password, st)

	// Assert
	if err != nil {
		t.Fatal(err)
	}

	// Act
	if !fileExists(vaultFilePath) {
		t.Fatal("Vault file should exist at ", vaultFilePath)
	}
}

func TestVaultOpenFromFileFailsWithMissingFile(t *testing.T) {
	// Arrange
	vaultFilePath := ".env.vault"

	// Act
	st, err := vaultOpenFromFile(vaultFilePath, "password")

	// Assert
	if st != nil {
		t.Fatal("Should fail with missing file")
	}

	if err == nil {
		t.Fatal("Should fail with missing file")
	}
}

package envenc

import (
	"testing"
)

func Test_VaultOpenFromFileFailsWithMissingFile(t *testing.T) {
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

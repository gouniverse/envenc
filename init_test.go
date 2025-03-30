package envenc

import (
	"os"
	"testing"
)

func TestInitSuccess(t *testing.T) {
	// Arrange
	vaultFilePath := "test.vault"
	vaultPassword := "testpassword"

	// Cleanup any existing file
	if err := os.Remove(vaultFilePath); err != nil && !os.IsNotExist(err) {
		t.Fatal(err)
	}

	// Act
	err := Init(vaultFilePath, vaultPassword)

	// Assert
	if err != nil {
		t.Fatalf("Init failed: %v", err)
	}

	// Verify file was created
	if !fileExists(vaultFilePath) {
		t.Fatal("Vault file was not created")
	}

	// Cleanup
	os.Remove(vaultFilePath)
}

func TestInitFileExists(t *testing.T) {
	// Arrange
	vaultFilePath := "test.vault"
	vaultPassword := "testpassword"

	// Create a file to simulate existing vault
	if err := os.WriteFile(vaultFilePath, []byte{}, 0644); err != nil {
		t.Fatal(err)
	}

	// Act
	err := Init(vaultFilePath, vaultPassword)

	// Assert
	if err == nil {
		t.Fatal("Expected error when vault file exists")
	}
	if err.Error() != "vault file already exists" {
		t.Fatalf("Unexpected error message: %v", err)
	}

	// Cleanup
	os.Remove(vaultFilePath)
}

func TestInitEmptyFilePath(t *testing.T) {
	// Arrange
	vaultFilePath := ""
	vaultPassword := "testpassword"

	// Act
	err := Init(vaultFilePath, vaultPassword)

	// Assert
	if err == nil {
		t.Fatal("Expected error with empty file path")
	}
	// Note: The actual error will depend on the implementation of fileExists and vaultSave
	// You may want to adjust this assertion based on your specific implementation
}

func TestInitEmptyPassword(t *testing.T) {
	// Arrange
	vaultFilePath := "test.vault"
	vaultPassword := ""

	// Cleanup any existing file
	if err := os.Remove(vaultFilePath); err != nil && !os.IsNotExist(err) {
		t.Fatal(err)
	}

	// Act
	err := Init(vaultFilePath, vaultPassword)

	// Assert
	if err == nil {
		t.Fatal("Expected error with empty password")
	}
	// Note: The actual error will depend on your vaultSave implementation
	// You may want to adjust this assertion based on your specific implementation

	// Cleanup
	os.Remove(vaultFilePath)
}

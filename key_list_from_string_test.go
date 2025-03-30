package envenc

import (
	"os"
	"path/filepath"
	"testing"
)

func setupTestVaultString(t *testing.T) (string, string, string) {
	tempDir := t.TempDir()
	vaultPath := filepath.Join(tempDir, "test.vault")
	vaultPassword := "testpassword"

	// Create a vault with some keys
	if err := Init(vaultPath, vaultPassword); err != nil {
		t.Fatalf("Failed to create vault file: %v", err)
	}

	// Add some test keys
	testKeys := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	for key, value := range testKeys {
		if err := KeySet(vaultPath, vaultPassword, key, value); err != nil {
			t.Fatalf("Failed to set test key: %v", err)
		}
	}

	// Read the vault contents
	vaultString, err := fileGetContents(vaultPath)
	if err != nil {
		t.Fatalf("Failed to read vault file: %v", err)
	}

	return vaultString, vaultPassword, tempDir
}

func TestKeyListFromString_Success(t *testing.T) {
	vaultString, vaultPassword, tempDir := setupTestVaultString(t)

	// Get the list of keys
	keys, err := KeyListFromString(vaultString, vaultPassword)
	if err != nil {
		t.Errorf("KeyListFromString failed: %v", err)
	}

	delete(keys, "id") // Remove id key, its used by DataObject

	// Verify we got the expected keys
	if len(keys) != 2 {
		t.Errorf("Expected 2 keys, got %d", len(keys))
	}

	if keys["key1"] != "value1" || keys["key2"] != "value2" {
		t.Errorf("Unexpected key values: %v", keys)
	}

	// Clean up
	os.RemoveAll(tempDir)
}

func TestKeyListFromString_InvalidPassword(t *testing.T) {
	vaultString, _, tempDir := setupTestVaultString(t)

	// Try to list keys with invalid password
	invalidPassword := "wrongpassword"
	_, err := KeyListFromString(vaultString, invalidPassword)
	if err == nil {
		t.Error("KeyListFromString should have failed with invalid password")
	}

	// Clean up
	os.RemoveAll(tempDir)
}

func TestKeyListFromString_InvalidVaultString(t *testing.T) {
	invalidVaultString := "invalid vault string"
	vaultPassword := "testpassword"

	// Try to list keys from invalid vault string
	_, err := KeyListFromString(invalidVaultString, vaultPassword)
	if err == nil {
		t.Error("KeyListFromString should have failed with invalid vault string")
	}
}

func TestKeyListFromString_EmptyVault(t *testing.T) {
	tempDir := t.TempDir()
	vaultPath := filepath.Join(tempDir, "test.vault")
	vaultPassword := "testpassword"

	// Create an empty vault
	if err := Init(vaultPath, vaultPassword); err != nil {
		t.Fatalf("Failed to create vault file: %v", err)
	}

	// Read the empty vault
	vaultString, err := fileGetContents(vaultPath)
	if err != nil {
		t.Fatalf("Failed to read vault file: %v", err)
	}

	// Get the list of keys
	keys, err := KeyListFromString(vaultString, vaultPassword)
	if err != nil {
		t.Errorf("KeyListFromString failed: %v", err)
	}

	delete(keys, "id") // Remove id key, its used by DataObject

	// Verify we got an empty map
	if len(keys) != 0 {
		t.Errorf("Expected empty map, got %d keys", len(keys))
	}

	// Clean up
	os.RemoveAll(tempDir)
}

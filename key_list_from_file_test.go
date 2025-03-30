package envenc

import (
	"os"
	"path/filepath"
	"testing"
)

func TestKeyListFromFile_EmptyVault(t *testing.T) {
	vaultPath, vaultPassword := setupTestVault(t)

	// List keys from empty vault
	keys, err := KeyListFromFile(vaultPath, vaultPassword)
	if err != nil {
		t.Errorf("KeyListFromFile failed: %v", err)
	}

	delete(keys, "id") // Remove id key, its used by DataObject

	// Should return empty map for empty vault
	if len(keys) != 0 {
		t.Errorf("Expected empty map, got %v", keys)
	}

	// Clean up
	os.Remove(vaultPath)
}

func TestKeyListFromFile_WithKeys(t *testing.T) {
	vaultPath, vaultPassword := setupTestVault(t)

	// Set some test keys
	testKeys := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}

	for key, value := range testKeys {
		if err := KeySet(vaultPath, vaultPassword, key, value); err != nil {
			t.Errorf("Failed to set test key: %v", err)
			return
		}
	}

	// List keys
	keys, err := KeyListFromFile(vaultPath, vaultPassword)
	if err != nil {
		t.Errorf("KeyListFromFile failed: %v", err)
	}

	delete(keys, "id") // Remove id key, its used by DataObject

	// Verify all keys are present
	if len(keys) != len(testKeys) {
		t.Errorf("Expected %d keys, got %d", len(testKeys), len(keys))
	}

	for key, value := range testKeys {
		if keys[key] != value {
			t.Errorf("Key %s has wrong value: expected %s, got %s", key, value, keys[key])
		}
	}

	// Clean up
	os.Remove(vaultPath)
}

func TestKeyListFromFile_InvalidPassword(t *testing.T) {
	vaultPath, vaultPassword := setupTestVault(t)

	// Set a test key
	if err := KeySet(vaultPath, vaultPassword, "testkey", "testvalue"); err != nil {
		t.Errorf("Failed to set test key: %v", err)
		return
	}

	// Try to list keys with invalid password
	invalidPassword := "wrongpassword"
	_, err := KeyListFromFile(vaultPath, invalidPassword)
	if err == nil {
		t.Error("KeyListFromFile should have failed with invalid password")
	}

	// Clean up
	os.Remove(vaultPath)
}

func TestKeyListFromFile_InvalidFilePath(t *testing.T) {
	vaultPath, vaultPassword := setupTestVault(t)

	// Create a non-existent file path
	invalidPath := filepath.Join(vaultPath, "nonexistent", "test.vault")

	// Try to list keys with invalid path
	_, err := KeyListFromFile(invalidPath, vaultPassword)
	if err == nil {
		t.Error("KeyListFromFile should have failed with invalid file path")
	}

	// Clean up
	os.Remove(vaultPath)
}

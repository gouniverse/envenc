package envenc

import (
	"path/filepath"
	"testing"
)

func TestKeyRemove_ExistingKey(t *testing.T) {
	vaultPath, vaultPassword := setupTestVault(t)
	keyName := "key1"

	// Remove the key
	err := KeyRemove(vaultPath, vaultPassword, keyName)
	if err != nil {
		t.Errorf("KeyRemove failed: %v", err)
	}

	// Verify the key was removed
	exists, err := KeyExists(vaultPath, vaultPassword, keyName)
	if err != nil {
		t.Errorf("KeyExists failed: %v", err)
	}
	if exists {
		t.Error("Key should have been removed but still exists")
	}
}

func TestKeyRemove_NonExistingKey(t *testing.T) {
	vaultPath, vaultPassword := setupTestVault(t)
	keyName := "nonexistentkey"

	// Try to remove a non-existent key
	err := KeyRemove(vaultPath, vaultPassword, keyName)
	if err != nil {
		t.Errorf("KeyRemove failed: %v", err)
	}
}

func TestKeyRemove_InvalidPassword(t *testing.T) {
	vaultPath, _ := setupTestVault(t)
	keyName := "key1"

	// Try to remove key with invalid password
	invalidPassword := "wrongpassword"
	err := KeyRemove(vaultPath, invalidPassword, keyName)
	if err == nil {
		t.Error("KeyRemove should have failed with invalid password")
	}
}

func TestKeyRemove_InvalidFilePath(t *testing.T) {
	vaultPath, vaultPassword := setupTestVault(t)
	keyName := "key1"

	// Create a non-existent file path
	invalidPath := filepath.Join(vaultPath, "nonexistent", "test.vault")

	// Try to remove key from invalid path
	err := KeyRemove(invalidPath, vaultPassword, keyName)
	if err == nil {
		t.Error("KeyRemove should have failed with invalid file path")
	}
}

func TestKeyRemove_LastKey(t *testing.T) {
	vaultPath, vaultPassword := setupTestVault(t)

	// Remove all keys
	if err := KeyRemove(vaultPath, vaultPassword, "key1"); err != nil {
		t.Errorf("Failed to remove key1: %v", err)
	}
	if err := KeyRemove(vaultPath, vaultPassword, "key2"); err != nil {
		t.Errorf("Failed to remove key2: %v", err)
	}

	// Verify the vault is empty
	keys, err := KeyListFromFile(vaultPath, vaultPassword)
	if err != nil {
		t.Errorf("KeyListFromFile failed: %v", err)
	}

	delete(keys, "id") // Remove id key, its used by DataObject

	if len(keys) != 0 {
		t.Errorf("Expected empty vault, got %d keys", len(keys))
	}
}

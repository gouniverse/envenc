package envenc

import (
	"os"
	"path/filepath"
	"testing"
)

func TestKeyGet_ValidKey(t *testing.T) {
	vaultPath, vaultPassword := setupTestVault(t)
	keyName := "testkey"
	keyValue := "testvalue"

	// First set the key
	err := KeySet(vaultPath, vaultPassword, keyName, keyValue)
	if err != nil {
		t.Errorf("Failed to set test key: %v", err)
		return
	}

	// Get the key
	gotValue, err := KeyGet(vaultPath, vaultPassword, keyName)
	if err != nil {
		t.Errorf("KeyGet failed: %v", err)
	}
	if gotValue != keyValue {
		t.Errorf("KeyGet() = %v, want %v", gotValue, keyValue)
	}

	// Clean up
	os.Remove(vaultPath)
}

func TestKeyGet_NonExistingKey(t *testing.T) {
	vaultPath, vaultPassword := setupTestVault(t)
	keyName := "nonexistentkey"

	// Get a non-existent key
	_, err := KeyGet(vaultPath, vaultPassword, keyName)
	if err != nil {
		t.Errorf("KeyGet failed: %v", err)
	}

	// Clean up
	os.Remove(vaultPath)
}

func TestKeyGet_InvalidPassword(t *testing.T) {
	vaultPath, vaultPassword := setupTestVault(t)
	keyName := "testkey"
	keyValue := "testvalue"

	// First set the key
	err := KeySet(vaultPath, vaultPassword, keyName, keyValue)
	if err != nil {
		t.Errorf("Failed to set test key: %v", err)
		return
	}

	// Try to get key with invalid password
	invalidPassword := "wrongpassword"
	_, err = KeyGet(vaultPath, invalidPassword, keyName)
	if err == nil {
		t.Error("KeyGet should have failed with invalid password")
	}

	// Clean up
	os.Remove(vaultPath)
}

func TestKeyGet_InvalidFilePath(t *testing.T) {
	vaultPath, vaultPassword := setupTestVault(t)
	keyName := "testkey"

	// Create a non-existent file path
	invalidPath := filepath.Join(vaultPath, "nonexistent", "test.vault")

	// Try to get key with invalid path
	_, err := KeyGet(invalidPath, vaultPassword, keyName)
	if err == nil {
		t.Error("KeyGet should have failed with invalid file path")
	}

	// Clean up
	os.Remove(vaultPath)
}

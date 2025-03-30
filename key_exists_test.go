package envenc

import (
	"os"
	"path/filepath"
	"testing"
)

func setupTestVault(t *testing.T) (vaultPath string, vaultPassword string) {
	tempDir := t.TempDir()
	vaultPath = filepath.Join(tempDir, "test.vault")
	vaultPassword = "testpassword"

	// Create an empty vault file
	if err := Init(vaultPath, vaultPassword); err != nil {
		t.Fatalf("Failed to create vault file: %v", err)
	}

	return vaultPath, vaultPassword
}

func TestKeyExists_ExistingKey(t *testing.T) {
	vaultPath, vaultPassword := setupTestVault(t)
	keyName := "testkey"
	keyValue := "testvalue"

	// First set the key
	err := KeySet(vaultPath, vaultPassword, keyName, keyValue)
	if err != nil {
		t.Errorf("Failed to set test key: %v", err)
		return
	}

	// Check if the key exists
	exists, err := KeyExists(vaultPath, vaultPassword, keyName)
	if err != nil {
		t.Errorf("KeyExists failed: %v", err)
	}
	if !exists {
		t.Error("KeyExists should have returned true for existing key")
	}

	// Clean up
	os.Remove(vaultPath)
}

func TestKeyExists_NonExistingKey(t *testing.T) {
	vaultPath, vaultPassword := setupTestVault(t)
	keyName := "nonexistentkey"

	// Check if the key exists
	_, err := KeyExists(vaultPath, vaultPassword, keyName)
	if err != nil {
		t.Errorf("KeyExists failed: %v", err)
	}

	// Clean up
	os.Remove(vaultPath)
}

func TestKeyExists_InvalidPassword(t *testing.T) {
	vaultPath, vaultPassword := setupTestVault(t)
	keyName := "testkey"
	keyValue := "testvalue"

	// First set the key
	err := KeySet(vaultPath, vaultPassword, keyName, keyValue)
	if err != nil {
		t.Errorf("Failed to set test key: %v", err)
		return
	}

	// Try to check key existence with invalid password
	invalidPassword := "wrongpassword"
	_, err = KeyExists(vaultPath, invalidPassword, keyName)
	if err == nil {
		t.Error("KeyExists should have failed with invalid password")
	}

	// Clean up
	os.Remove(vaultPath)
}

func TestKeyExists_InvalidFilePath(t *testing.T) {
	vaultPath, vaultPassword := setupTestVault(t)
	keyName := "testkey"

	// Create a non-existent file path
	invalidPath := filepath.Join(vaultPath, "nonexistent", "test.vault")

	// Try to check key existence with invalid path
	_, err := KeyExists(invalidPath, vaultPassword, keyName)
	if err == nil {
		t.Error("KeyExists should have failed with invalid file path")
	}

	// Clean up
	os.Remove(vaultPath)
}

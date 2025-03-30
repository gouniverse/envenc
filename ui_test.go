package envenc

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestUI_ApiKeys_Success(t *testing.T) {
	tempDir := t.TempDir()
	vaultPath := filepath.Join(tempDir, "test.vault")
	vaultPassword := "testpassword"

	// Create a vault with a test key
	if err := Init(vaultPath, vaultPassword); err != nil {
		t.Fatalf("Failed to create vault file: %v", err)
	}

	if err := KeySet(vaultPath, vaultPassword, "testkey", "testvalue"); err != nil {
		t.Fatalf("Failed to set test key: %v", err)
	}

	// Create a test request
	req := httptest.NewRequest("GET", "/api/keys", nil)
	req.Form = map[string][]string{
		"vault":    {vaultPath},
		"password": {vaultPassword},
	}

	// Create a response recorder
	w := httptest.NewRecorder()

	// Create UI instance
	ui := &ui{
		vaultPath: vaultPath,
	}

	// Call the API endpoint
	response := ui.apiKeys(w, req)
	w.Write([]byte(response))

	// Check the response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	// Verify response
	expectedPairs := []struct {
		key   string
		value string
	}{
		{"testkey", "testvalue"},
	}

	for _, pair := range expectedPairs {
		if !strings.Contains(response, fmt.Sprintf(`"%s":"%s"`, pair.key, pair.value)) {
			t.Errorf("Expected key-value pair %s:%s in response, got: %s", pair.key, pair.value, response)
		}
	}

	// Clean up
	os.RemoveAll(tempDir)
}

func TestUI_ApiKeys_MissingVault(t *testing.T) {
	tempDir := t.TempDir()
	vaultPath := filepath.Join(tempDir, "nonexistent.vault")
	vaultPassword := "testpassword"

	// Create a test request with a non-existent vault
	req := httptest.NewRequest("GET", "/api/keys", nil)
	req.Form = map[string][]string{
		"vault":    {vaultPath},
		"password": {vaultPassword},
	}

	// Create a response recorder
	w := httptest.NewRecorder()

	// Create UI instance
	ui := &ui{
		vaultPath: vaultPath,
	}

	// Call the API endpoint
	response := ui.apiKeys(w, req)
	w.Write([]byte(response))

	// Check the response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	// Verify error response
	expectedPairs := []struct {
		key   string
		value string
	}{
		{"status", "error"},
		{"message", "Vault file does not exist"},
	}

	for _, pair := range expectedPairs {
		if !strings.Contains(response, fmt.Sprintf(`"%s":"%s"`, pair.key, pair.value)) {
			t.Errorf("Expected key-value pair %s:%s in response, got: %s", pair.key, pair.value, response)
		}
	}

	// Clean up
	os.RemoveAll(tempDir)
}

package envenc

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

type ApiResponse struct {
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

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
	var apiResponse ApiResponse
	if err := json.Unmarshal([]byte(response), &apiResponse); err != nil {
		t.Errorf("Failed to parse JSON response: %v", err)
	}

	if apiResponse.Status != "success" {
		t.Errorf("Expected status 'success', got '%s'", apiResponse.Status)
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
	var apiResponse ApiResponse
	if err := json.Unmarshal([]byte(response), &apiResponse); err != nil {
		t.Errorf("Failed to parse JSON response: %v", err)
	}

	if apiResponse.Status != "error" {
		t.Errorf("Expected status 'error', got '%s'", apiResponse.Status)
	}

	if apiResponse.Message != "Vault file does not exist" {
		t.Errorf("Expected message 'Vault file does not exist', got '%s'", apiResponse.Message)
	}

	// Clean up
	os.RemoveAll(tempDir)
}

func TestUI_ApiKeyAdd_Success(t *testing.T) {
	tempDir := t.TempDir()
	vaultPath := filepath.Join(tempDir, "test.vault")
	vaultPassword := "testpassword"
	testKey := "newkey"
	testValue := "newvalue"

	// Create a vault
	if err := Init(vaultPath, vaultPassword); err != nil {
		t.Fatalf("Failed to create vault file: %v", err)
	}

	// Create a test request
	req := httptest.NewRequest("GET", "/api/key-add", nil)
	req.Form = map[string][]string{
		"vault":    {vaultPath},
		"password": {vaultPassword},
		"key":      {testKey},
		"value":    {testValue},
	}

	// Create a response recorder
	w := httptest.NewRecorder()

	// Create UI instance
	ui := &ui{
		vaultPath: vaultPath,
	}

	// Call the API endpoint
	response := ui.apiKeyAdd(w, req)
	w.Write([]byte(response))

	// Check the response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	// Verify success response
	var apiResponse ApiResponse
	if err := json.Unmarshal([]byte(response), &apiResponse); err != nil {
		t.Errorf("Failed to parse JSON response: %v", err)
	}

	if apiResponse.Status != "success" {
		t.Errorf("Expected status 'success', got '%s'", apiResponse.Status)
	}

	if apiResponse.Message != "Key added successfully" {
		t.Errorf("Expected message 'Key added successfully', got '%s'", apiResponse.Message)
	}

	// Verify key was actually added
	exists, err := KeyExists(vaultPath, vaultPassword, testKey)
	if err != nil {
		t.Errorf("Failed to verify key existence: %v", err)
	}
	if !exists {
		t.Errorf("Key was not added to the vault")
	}

	// Clean up
	os.RemoveAll(tempDir)
}

func TestUI_ApiKeyAdd_MissingVault(t *testing.T) {
	tempDir := t.TempDir()
	vaultPath := filepath.Join(tempDir, "nonexistent.vault")
	vaultPassword := "testpassword"
	testKey := "newkey"
	testValue := "newvalue"

	// Create a test request with a non-existent vault
	req := httptest.NewRequest("GET", "/api/key-add", nil)
	req.Form = map[string][]string{
		"vault":    {vaultPath},
		"password": {vaultPassword},
		"key":      {testKey},
		"value":    {testValue},
	}

	// Create a response recorder
	w := httptest.NewRecorder()

	// Create UI instance
	ui := &ui{
		vaultPath: vaultPath,
	}

	// Call the API endpoint
	response := ui.apiKeyAdd(w, req)
	w.Write([]byte(response))

	// Check the response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	// Verify error response
	var apiResponse ApiResponse
	if err := json.Unmarshal([]byte(response), &apiResponse); err != nil {
		t.Errorf("Failed to parse JSON response: %v", err)
	}

	if apiResponse.Status != "error" {
		t.Errorf("Expected status 'error', got '%s'", apiResponse.Status)
	}

	if apiResponse.Message != "Vault file does not exist" {
		t.Errorf("Expected message 'Vault file does not exist', got '%s'", apiResponse.Message)
	}

	// Clean up
	os.RemoveAll(tempDir)
}

func TestUI_ApiKeyAdd_KeyExists(t *testing.T) {
	tempDir := t.TempDir()
	vaultPath := filepath.Join(tempDir, "test.vault")
	vaultPassword := "testpassword"
	testKey := "existingkey"
	testValue := "newvalue"

	// Create a vault with an existing key
	if err := Init(vaultPath, vaultPassword); err != nil {
		t.Fatalf("Failed to create vault file: %v", err)
	}
	if err := KeySet(vaultPath, vaultPassword, testKey, "oldvalue"); err != nil {
		t.Fatalf("Failed to set test key: %v", err)
	}

	// Create a test request
	req := httptest.NewRequest("GET", "/api/key-add", nil)
	req.Form = map[string][]string{
		"vault":    {vaultPath},
		"password": {vaultPassword},
		"key":      {testKey},
		"value":    {testValue},
	}

	// Create a response recorder
	w := httptest.NewRecorder()

	// Create UI instance
	ui := &ui{
		vaultPath: vaultPath,
	}

	// Call the API endpoint
	response := ui.apiKeyAdd(w, req)
	w.Write([]byte(response))

	// Check the response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	// Verify error response
	var apiResponse ApiResponse
	if err := json.Unmarshal([]byte(response), &apiResponse); err != nil {
		t.Errorf("Failed to parse JSON response: %v", err)
	}

	if apiResponse.Status != "error" {
		t.Errorf("Expected status 'error', got '%s'", apiResponse.Status)
	}

	if apiResponse.Message != "Key already exists" {
		t.Errorf("Expected message 'Key already exists', got '%s'", apiResponse.Message)
	}

	// Clean up
	os.RemoveAll(tempDir)
}

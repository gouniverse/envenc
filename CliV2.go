package envenc

// !!! LOOKS LIKE THERE IS A PROBLEM WITH PTERM !!!
// In GitBash it repeats the text of the input box multiple times making it
// imposible to enter the values

import (

	// "github.com/gouniverse/utils"

	"path/filepath"
	"sort"
	"strings"

	"github.com/mingrammer/cfmt"
	"github.com/pterm/pterm"
	"github.com/samber/lo"
)

type CliV2 struct {
	commands map[string]func([]string)

	errorMessage  string
	vaultPath     string
	vaultPassword string
}

func NewCliV2() *CliV2 {
	return &CliV2{
		commands: map[string]func([]string){
			"init":        (&CliV2{}).VaultInit,
			"encrypt":     cliEncrypt,
			"decrypt":     cliDecrypt,
			"obfuscate":   cliObfuscate,
			"deobfuscate": cliDeobfuscate,
			"key-set":     (&CliV2{}).VaultKeySet,
			"key-get":     (&CliV2{}).VaultKeyGet,
			"key-list":    (&CliV2{}).VaultKeyList,
			"key-remove":  cliVaultKeyRemove,
		},
	}
}

func (c *CliV2) VaultKeyExists(args []string) {
}

func (c *CliV2) VaultKeyList(args []string) {
	c.vaultPath, c.errorMessage = c.FindVaultPathFromArgs(args)

	if c.errorMessage != "" {
		cfmt.Errorln(c.errorMessage)
		return
	}

	if c.vaultPath == "" {
		c.vaultPath, c.errorMessage = c.AskVaultPath()

		if c.errorMessage != "" {
			cfmt.Errorln(c.errorMessage)
			return
		}
	}

	c.vaultPassword, c.errorMessage = c.AskVaultPassword()

	if c.errorMessage != "" {
		cfmt.Errorln(c.errorMessage)
		return
	}

	cfmt.Infoln("Listing keys in vault: " + c.vaultPath)
	cfmt.Infoln("")

	data, err := KeyListFromFile(c.vaultPath, c.vaultPassword)

	if err != nil {
		cfmt.Errorln(err)
		return
	}

	if len(data) == 0 {
		cfmt.Successln("Key list is empty!")
		cfmt.Successln("")
		return
	}

	cfmt.Successln("Total keys: ", len(data))
	cfmt.Successln("")

	keys := lo.Keys(data)

	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	lo.ForEach(keys, func(key string, _ int) {
		cfmt.Successln(key, " = ", data[key])
	})

	cfmt.Successln("")
}

// VaultKeyGet gets a key from the vault
//
// Buisiness logic:
//   - If the vault file is provided as an argument, use it
//   - If the vault file is not provided, ask for it
//   - Check that the vault file exists
//   - Ask for the password to use for the vault
//   - Open the vault file, to confirm the password is correct
//   - Ask for the key's name to get
//   - Get the key from the vault
func (c *CliV2) VaultKeyGet(args []string) {
	c.vaultPath, c.errorMessage = c.FindVaultPathFromArgs(args)

	if c.errorMessage != "" {
		cfmt.Errorln(c.errorMessage)
		return
	}

	if c.vaultPath == "" {
		c.vaultPath, c.errorMessage = c.AskVaultPath()

		if c.errorMessage != "" {
			cfmt.Errorln(c.errorMessage)
			return
		}
	}

	c.vaultPassword, c.errorMessage = c.AskVaultPassword()

	if c.errorMessage != "" {
		cfmt.Errorln(c.errorMessage)
		return
	}

	keyName, errorMessage := c.AskKeyName()

	if errorMessage != "" {
		cfmt.Errorln(errorMessage)
		return
	}

	_, err := KeyListFromFile(c.vaultPath, c.vaultPassword)

	if err != nil {
		cfmt.Errorln(err)
		return
	}

	keyExists, err := KeyExists(c.vaultPath, c.vaultPassword, keyName)

	if err != nil {
		cfmt.Errorln(err)
		return
	}

	if !keyExists {
		cfmt.Errorln("Key does not exist!")
		return
	}

	keyValue, err := KeyGet(c.vaultPath, c.vaultPassword, keyName)

	if err != nil {
		cfmt.Errorln(err)
		return
	}

	if keyValue == "" {
		cfmt.Successln("Key value is empty!")
		return
	}

	cfmt.Successln("The value of the key is:")
	cfmt.Successln(keyValue)
}

// VaultKeySet sets a key in the vault
//
// Buisiness logic:
//   - If the vault file is provided as an argument, use it
//   - If the vault file is not provided, ask for it
//   - Check that the vault file exists
//   - Ask for the password to use for the vault
//   - Open the vault file, to confirm the password is correct
//   - Ask for the key's name to set
//   - Ask for the key's value to set (must support multiline)
//   - Set the key in the vault
//   - Close the vault file
//   - Ask the user if he wants to add another key
//   - If the user wants to add another key, repeat the process
func (c *CliV2) VaultKeySet(args []string) {
	c.vaultPath, c.errorMessage = c.FindVaultPathFromArgs(args)

	if c.errorMessage != "" {
		cfmt.Errorln(c.errorMessage)
		return
	}

	if c.vaultPath == "" {
		c.vaultPath, c.errorMessage = c.AskVaultPath()

		if c.errorMessage != "" {
			cfmt.Errorln(c.errorMessage)
			return
		}
	}

	c.vaultPassword, c.errorMessage = c.AskVaultPassword()

	if c.errorMessage != "" {
		cfmt.Errorln(c.errorMessage)
		return
	}

	_, err := KeyListFromFile(c.vaultPath, c.vaultPassword)

	if err != nil {
		cfmt.Errorln(err)
		return
	}

	keyName, errorMessageKey := c.AskKeyName()

	if errorMessageKey != "" {
		cfmt.Errorln(errorMessageKey)
		return
	}

	keyValue, errorMessageValue := c.AskKeyValue()

	if errorMessageValue != "" {
		cfmt.Errorln(errorMessageValue)
		return
	}

	err = KeySet(c.vaultPath, c.vaultPassword, keyName, keyValue)

	if err != nil {
		cfmt.Errorln(err)
		return
	}

	cfmt.Successln("key set successfully!")
	cfmt.Successln("")
}

// VaultInit initializes a new vault file
//
// Buisiness logic:
//   - If the vault file is provided as an argument, use it
//   - If the vault file is not provided, ask for it
//   - Check that the vault file does not exist already
//   - Ask for the password to use for the vault
//   - Confirm the password to avoid any spelling mistakes
//   - Create the vault file
func (c *CliV2) VaultInit(args []string) {
	c.vaultPath, c.errorMessage = c.FindVaultPathFromArgs(args)

	if c.errorMessage != "" {
		cfmt.Errorln(c.errorMessage)
		return
	}

	if c.vaultPath == "" {
		c.vaultPath, c.errorMessage = c.AskVaultPath()

		if c.errorMessage != "" {
			cfmt.Errorln(c.errorMessage)
			return
		}
	}

	c.vaultPassword, c.errorMessage = c.AskVaultPasswordWithConfirm()

	if c.errorMessage != "" {
		cfmt.Errorln(c.errorMessage)
		return
	}

	err := Init(c.vaultPath, c.vaultPassword)

	if err != nil {
		cfmt.Errorln(err)
		return
	}

	cfmt.Successln("Vault initialized successfully at " + c.vaultPath)
}

// AskVaultPassword asks the user to enter a password
//
// Buisiness logic:
//   - Ask the user to enter a password
//   - If the user enters an empty password, return an error
//   - Otherwise return the password
func (c *CliV2) AskVaultPassword() (string, errorMessage string) {
	password, err := pterm.DefaultInteractiveTextInput.
		WithMask("*").
		Show("Enter password")

	if err != nil {
		return "", err.Error()
	}

	pterm.Println() // Blank line, for better readability

	if password == "" {
		return "", "- Password cannot be empty"
	}

	return password, ""
}

// AskVaultPasswordWithConfirm asks the user to enter a password and confirm it
//
// Buisiness logic:
//   - Ask the user to enter a password
//   - If the user enters an empty password, return an error
//   - Confirm the password to avoid any spelling mistakes
//   - If the password and confirmation do not match, return an error
//   - Otherwise return the password
func (c *CliV2) AskVaultPasswordWithConfirm() (string, errorMessage string) {
	password, err := pterm.DefaultInteractiveTextInput.
		WithMask("*").
		Show("Enter password")

	if err != nil {
		return "", err.Error()
	}

	pterm.Println() // Blank line, for better readability

	if password == "" {
		return "", "- Password cannot be empty"
	}

	passwordConfirm, err := pterm.DefaultInteractiveTextInput.
		WithMask("*").
		Show("Confirm password")

	if err != nil {
		return "", err.Error()
	}

	if password != passwordConfirm {
		return "", "- Passwords do not match"
	}

	pterm.Println() // Blank line, for better readability

	return password, ""
}

// AskVaultPath asks the user to enter the path to the vault file
//
// Buisiness logic:
//   - Ask the user to enter the path to the vault file
//   - If the user enters an empty path, return an error
//   - To confirm its a .vault file, we check the extension
//   - If the extension is not .vault, return an error
//   - Otherwise return the file path
func (c *CliV2) AskVaultPath() (string, errorMessage string) {
	filePath, err := pterm.DefaultInteractiveTextInput.Show("Enter the path to the vault file")

	if err != nil {
		return "", err.Error()
	}

	pterm.Println() // Blank line, for better readability

	if filePath == "" {
		return "", "- File path cannot be empty"
	}

	extension := filepath.Ext(filePath)

	if extension != ".vault" {
		return "", "- File extension must be .vault"
	}

	return filePath, ""
}

// AskKeyName asks the user to enter the name of the key
//
// Buisiness logic:
//   - Ask the user to enter the name of the key
//   - If the user enters an empty name, return an error
//   - If the name contains spaces, return an error
//   - Otherwise return the name
func (c *CliV2) AskKeyName() (string, errorMessage string) {
	keyName, err := pterm.DefaultInteractiveTextInput.Show("Enter the name of the key (i.e. 'DB_PASSWORD')")

	if err != nil {
		return "", err.Error()
	}

	pterm.Println() // Blank line, for better readability

	if keyName == "" {
		return "", "- Key name cannot be empty"
	}

	if strings.Contains(keyName, " ") {
		return "", "- Key name cannot contain spaces"
	}

	return keyName, ""
}

// AskKeyValue asks the user to enter the value of the key
//
// Buisiness logic:
//   - Ask the user to enter the value of the key (allowing multiline)
//   - If the user enters an empty value, do not return an error, it is ok
//   - Otherwise return the value
func (c *CliV2) AskKeyValue() (string, errorMessage string) {
	keyValue, err := pterm.DefaultInteractiveTextInput.
		WithMultiLine().
		Show("Enter the value of the key")

	if err != nil {
		return "", err.Error()
	}

	pterm.Println() // Blank line, for better readability

	return keyValue, ""
}

// FindVaultPathFromArgs finds the file path from the arguments, if provided
//
// Buisiness logic:
//   - If the arguments are empty, return an empty file path
//   - We expect the first argument to be the file path
//   - To confirm its a .vault file, we check the extension
//   - If the extension is not .vault, return an error
//   - Otherwise return the file path
//
// Parameters:
//   - args: The command line arguments (excluding the executable, and the command names)
//
// Returns:
//   - filePath: The file path
//   - errorMessage: The error message
func (c *CliV2) FindVaultPathFromArgs(args []string) (filePath string, errorMessage string) {
	if len(args) < 1 {
		return "", ""
	}

	filePath = args[0]

	extension := filepath.Ext(filePath)

	if extension != ".vault" {
		return "", "- File extension must be .vault"
	}

	return filePath, ""
}

// Run executes the command
//
// # It expects a command with the second argument being the command
//
// Buisiness logic:
//   - Parse command line arguments
//   - First argument is the name of the executable, ignore it
//   - Second argument is the command
//   - If there is no command, help is shown as default
//   - If the command is unknown, help is shown as default
//   - Otherwise execute the command
//
// Parameters
//   - args: The command line arguments
//
// Returns
//   - None
func (c *CliV2) Run(args []string) {
	if len(args) < 2 {
		cliShowHelp([]string{})
		return
	}

	command := lo.ValueOr(c.commands, args[1], cliShowHelp)

	command(args[2:]) // remove the executable and the command names
}

func CliUi(args []string) {
	commands := map[string]func([]string){
		"init":        cliVaultInit,
		"encrypt":     cliEncrypt,
		"decrypt":     cliDecrypt,
		"obfuscate":   cliObfuscate,
		"deobfuscate": cliDeobfuscate,
		"key-set":     cliVaultKeySet,
		"key-list":    cliVaultKeyList,
		"key-remove":  cliVaultKeyRemove,
	}

	if len(args) < 2 {
		cliShowHelp([]string{})
		return
	}

	command := lo.ValueOr(commands, args[1], cliShowHelp)

	command(args[2:])
}

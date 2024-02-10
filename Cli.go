package envenc

import (
	"errors"
	"fmt"
	"path/filepath"
	"sort"
	"strings"
	"syscall"

	// "github.com/gouniverse/utils"
	"github.com/mingrammer/cfmt"
	"github.com/samber/lo"
	"golang.org/x/crypto/ssh/terminal"
)

func Cli(args []string) {
	if len(args) < 2 {
		cliShowHelp()
		return
	}

	command := args[1]

	allowedCommands := []string{
		"init",
		"key-list",
		"key-set",
		"key-remove",
	}

	if !lo.Contains(allowedCommands, command) {
		cliShowHelp()
		return
	}

	if command == "init" {
		cliVaultInit(args[2:])
		return
	}

	if command == "key-set" {
		cliVaultKeySet(args[2:])
		return
	}

	if command == "key-list" {
		cliVaultKeyList(args[2:])
		return
	}

	if command == "key-remove" {
		cliVaultKeyRemove(args[2:])
		return
	}

	cliShowHelp()
}

func cliVaultInit(args []string) {
	vaultPath, err := cliVaultPathFromArgs(args)

	if err != nil {
		cfmt.Errorln(err)
		return
	}

	vaultPassword, err := cliAskPassword()

	if err != nil {
		cfmt.Errorln(err)
		return
	}

	cfmt.Infoln("Confirm password to continue...")
	cfmt.Infoln("")

	vaultPasswordConfirm, err := cliAskPassword()

	if err != nil {
		cfmt.Errorln(err)
		return
	}

	if vaultPassword != vaultPasswordConfirm {
		cfmt.Errorln("passwords do not match!")
		return
	}

	err = Init(vaultPath, vaultPassword)

	if err != nil {
		cfmt.Errorln(err)
		return
	}

	cfmt.Successln("Vault initialized successfully at " + vaultPath)
}

func cliVaultKeyList(args []string) {
	vaultPath, err := cliVaultPathFromArgs(args)

	if err != nil {
		cfmt.Errorln(err)
		return
	}

	vaultPassword, err := cliAskPassword()

	if err != nil {
		cfmt.Errorln(err)
		return
	}

	data, err := EnvList(vaultPath, vaultPassword)

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

func cliVaultKeyRemove(args []string) {
	vaultPath, err := cliVaultPathFromArgs(args)

	if err != nil {
		cfmt.Errorln(err)
		return
	}

	vaultPassword, err := cliAskPassword()

	if err != nil {
		cfmt.Errorln(err)
		return
	}

	key, err := cliAskKey()

	if err != nil {
		cfmt.Errorln(err)
		return
	}

	err = KeyRemove(vaultPath, vaultPassword, key)

	if err != nil {
		cfmt.Errorln(err)
	}

	cfmt.Successln("Removed key successfully!")
	cfmt.Successln("")
}

func cliVaultKeySet(args []string) {
	vaultPath, err := cliVaultPathFromArgs(args)

	if err != nil {
		cfmt.Errorln(err)
		return
	}

	password, err := cliAskPassword()

	if err != nil {
		cfmt.Errorln(err)
		return
	}

	key, err := cliAskKey()

	if err != nil {
		cfmt.Errorln(err)
		return
	}

	value, err := cliAskValue()

	if err != nil {
		cfmt.Errorln(err)
		return
	}

	err = KeySet(vaultPath, password, key, value)

	if err != nil {
		cfmt.Errorln(err)
		return
	}

	cfmt.Successln("key set successfully!")
	cfmt.Successln("")
}

func cliVaultPathFromArgs(args []string) (string, error) {
	if len(args) < 1 {
		return "", errors.New("- No file path for the vault provided, i.e. env.vault")
	}

	filePath := args[0]

	extension := filepath.Ext(filePath)

	if extension != ".vault" {
		return "", errors.New("- File extension must be .vault")
	}

	return filePath, nil
}

func cliShowHelp() {
	cfmt.Infoln("Usage:")
	cfmt.Infoln(" - init [vaultPath] - Initialize the vault")
	cfmt.Infoln(" - key-list [vaultPath] - Lists all the keys in the vault")
	cfmt.Infoln(" - key-remove [vaultPath] - Removes a key from the vault")
	cfmt.Infoln(" - key-set [vaultPath] - Sets a key in the vault")
	cfmt.Infoln(" - help - Show this help")
	cfmt.Infoln("")
	cfmt.Infoln("envenc is a tool for managing encrypted environment variables")
	cfmt.Infoln("it allows you to store and retrieve environment variables in")
	cfmt.Infoln("a secure password protected vault file")
	cfmt.Infoln("")
	cfmt.Infoln("Example:")
	cfmt.Infoln("$> vault init .env.vault")
	cfmt.Infoln("$> vault key-set .env.vault")
	cfmt.Infoln("$> vault key-list .env.vault")
	cfmt.Infoln("$> vault key-remove .env.vault")
	cfmt.Infoln("")
	cfmt.Infoln("For more information visit:")
	cfmt.Infoln("")
	cfmt.Infoln("https://github.com/gouniverse/envenc")
}

func cliAskPassword() (string, error) {
	cfmt.Infoln("=================== START: Vault Security ===================")
	cfmt.Infoln("Enter vault password:")

	passwordBytes, err := terminal.ReadPassword(int(syscall.Stdin))

	if err != nil {
		return "", err
	}
	password := string(passwordBytes)
	password = strings.TrimSpace(password)

	if password == "" {
		return "", errors.New("password cannot be empty")
	}

	if len(password) < 10 {
		return "", errors.New("password must be at least 10 characters long")
	}

	if len(password) < 10 {
		return "", errors.New("password must be at least 10 characters long")
	}

	cfmt.Infoln("Password: ", strPadLeft("***", 32, "*"))
	cfmt.Infoln("==================== END: Vault Security ====================")
	cfmt.Infoln("")
	fmt.Print("\n")

	return password, nil
}

func cliAskKey() (string, error) {
	cfmt.Infoln("Enter key name:")

	key := ""
	fmt.Scanln(&key)
	key = strings.TrimSpace(key)

	if key == "" {
		return "", errors.New("key name cannot be empty")
	}

	if strings.ToUpper(key) != key {
		return "", errors.New("key name must be uppercase")
	}

	return key, nil
}

func cliAskValue() (string, error) {
	cfmt.Infoln("Enter value:")

	value := ""

	fmt.Scanln(&value)

	if value == "" {
		return "", errors.New("value cannot be empty")
	}

	return value, nil
}

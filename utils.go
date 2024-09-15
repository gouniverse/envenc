package envenc

import (
	"crypto/rand"
	"encoding/base64"
	"os"
)

// fileGetContents reads entire file into a string
func fileGetContents(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	return string(data), err
}

// fileExists checks if a file exists
func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)

	return err == nil
}

// filePutContents adds content to file
func filePutContents(filename string, data string, mode os.FileMode) error {
	return os.WriteFile(filename, []byte(data), mode)
}

func strPadLeft(input string, padLength int, padString string) string {
	output := ""
	inputLen := len(input)
	if inputLen >= padLength {
		return input
	}
	ll := padLength - inputLen
	for i := 1; i <= ll; i = i + len(padString) {
		output += padString
	}
	return output + input
}

// strRandom generates random string of specified length
func strRandom(length int) string {
	buff := make([]byte, length)
	rand.Read(buff)
	str := base64.StdEncoding.EncodeToString(buff)

	// Base 64 can be longer than len
	return str[:length]
}

// strRandFromGamma generates random string of specified length with the characters specified in the gamma string
// func strRandFromGamma(length int, gamma string) string {
// 	if length <= 0 {
// 		return ""
// 	}

// 	if gamma == "" {
// 		return ""
// 	}

// 	inRune := []rune(gamma)
// 	out := make([]rune, length) // Pre-allocate for efficiency

// 	for i := range out {
// 		randomIndex := mathRand.IntN(len(inRune))
// 		out[i] = inRune[randomIndex]
// 	}

// 	return string(out)
// }

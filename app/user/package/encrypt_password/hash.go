package encrypt_password

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

type ArgonParam struct {
	Salt    []byte
	Time    uint32
	Memory  uint32
	Threads uint8
	KeyLen  uint32
}

func genSalt(keyLen uint32) ([]byte, error) {
	salt := make([]byte, keyLen)
	if _, err := rand.Read(salt); err != nil {
		return nil, fmt.Errorf("failed to generate salt: %w", err)
	}
	return salt, nil
}

func DefaultArgonParam() (*ArgonParam, error) {
	salt, err := genSalt(32)
	if err != nil {
		return nil, err
	}
	return &ArgonParam{
		Salt:    salt,
		Time:    1,
		Memory:  64 * 1024,
		Threads: 4,
		KeyLen:  32,
	}, nil
}

func HashPassword(password string) (string, error) {
	argonParam, err := DefaultArgonParam()
	if err != nil {
		return "", err
	}
	hashedBytes := argon2.IDKey([]byte(password), argonParam.Salt, argonParam.Time, argonParam.Memory, argonParam.Threads, argonParam.KeyLen)

	encodedSalt := base64.RawStdEncoding.EncodeToString(argonParam.Salt)
	encodedHash := base64.RawStdEncoding.EncodeToString(hashedBytes)

	return fmt.Sprintf("%s$%s", encodedSalt, encodedHash), nil
}

func VerifyPassword(storedHash, password string) (bool, error) {
	parts := split(storedHash, "$")
	if len(parts) != 2 {
		return false, fmt.Errorf("invalid stored hash format")
	}
	encodedSalt, encodedHash := parts[0], parts[1]

	salt, err := base64.RawStdEncoding.DecodeString(encodedSalt)
	if err != nil {
		return false, fmt.Errorf("failed to decode salt: %w", err)
	}

	expectedHash, err := base64.RawStdEncoding.DecodeString(encodedHash)
	if err != nil {
		return false, fmt.Errorf("failed to decode hash: %w", err)
	}

	argonParam := &ArgonParam{
		Salt:    salt,
		Time:    1,
		Memory:  64 * 1024,
		Threads: 4,
		KeyLen:  uint32(len(expectedHash)),
	}
	computedHash := argon2.IDKey([]byte(password), argonParam.Salt, argonParam.Time, argonParam.Memory, argonParam.Threads, argonParam.KeyLen)

	if len(computedHash) != len(expectedHash) {
		return false, nil
	}
	for i := 0; i < len(computedHash); i++ {
		if computedHash[i] != expectedHash[i] {
			return false, nil
		}
	}
	return true, nil
}

func split(s, sep string) []string {
	var parts []string
	start := 0
	for {
		i := indexOf(s, sep, start)
		if i < 0 {
			parts = append(parts, s[start:])
			break
		}
		parts = append(parts, s[start:i])
		start = i + len(sep)
	}
	return parts
}

func indexOf(s, sep string, start int) int {
	index := -1
	for i := start; i <= len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep {
			index = i
			break
		}
	}
	return index
}

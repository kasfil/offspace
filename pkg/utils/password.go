package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/spf13/viper"
	"golang.org/x/crypto/argon2"
)

type hashConfig struct {
	time    uint32
	memory  uint32
	threads uint8
	keylen  uint32
}

// GeneratePassword : Generate password from user and return
// full string that will be stored in database
func GeneratePassword(password string) (string, error) {
	// generate password salt
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	c := &hashConfig{
		time:    viper.GetUint32("hash_time"),
		memory:  viper.GetUint32("hash_memory") * 1024,
		threads: uint8(viper.GetUint("hash_threads")),
		keylen:  viper.GetUint32("hash_keylen"),
	}

	// encrypt password
	hash := argon2.IDKey([]byte(password), salt, c.time, c.memory, c.threads, c.keylen)

	// convert byte into string
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	format := "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s"
	storedHash := fmt.Sprintf(format, argon2.Version, c.memory, c.time, c.threads, b64Salt, b64Hash)

	return storedHash, nil
}

// ComparePassword : Compare user inputted password to existing hash
func ComparePassword(password, hash string) (bool, error) {
	// Split hash into parts
	parts := strings.Split(hash, "$")

	// getting hash config from existing hash parts
	c := &hashConfig{}
	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &c.memory, &c.time, &c.threads)
	if err != nil {
		return false, err
	}

	// getting hash salt
	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}

	// getting hashed password
	hashedPassword, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}

	c.keylen = uint32(len(hashedPassword))

	// create user inputted hash for comparation
	comparedHash := argon2.IDKey([]byte(password), salt, c.time, c.memory*1024, c.threads, c.keylen)

	return (subtle.ConstantTimeCompare(hashedPassword, comparedHash) == 1), nil
}

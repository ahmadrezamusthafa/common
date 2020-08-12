package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/ahmadrezamusthafa/common/errors"
	"strconv"
	"strings"
)

const HMACSHA256HashType = "HMACSHA256"

type PasswordHasher struct {
	Algo      string
	Iteration int
	Key       string
	Salt      string
}

func hmacSha256Digest(key string, salt string, message string, iteration int) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(salt + message))
	hashedMessage := hex.EncodeToString(mac.Sum(nil))
	if iteration <= 1 {
		return hashedMessage
	}
	return hmacSha256Digest(key, salt, hashedMessage, iteration-1)
}

func digest(algo string, key string, salt string, message string, iteration int) (string, error) {
	switch algo {
	case HMACSHA256HashType:
		return hmacSha256Digest(key, salt, message, iteration), nil
	default:
		return "", errors.AddTrace(errors.New("Invalid hash type"))
	}
}

func (p *PasswordHasher) IsValidHashedPassword(userSuppliedPassword string, hashedPassword string) bool {
	passParams := strings.Split(hashedPassword, "-")
	if len(passParams) == 4 {
		algo := passParams[0]
		iter, err := strconv.Atoi(passParams[1])
		if err != nil {
			return false
		}
		salt := passParams[2]
		content, err := digest(algo, p.Key, salt, userSuppliedPassword, iter)
		if err != nil {
			return false
		}
		hashedUserPassword := fmt.Sprintf("%s-%d-%s-%s", algo, iter, salt, content)
		match := hashedUserPassword == hashedPassword
		return match
	}

	mac := hmac.New(sha256.New, []byte(p.Key))
	mac.Write([]byte(p.Salt + userSuppliedPassword))
	hashedMessage := hex.EncodeToString(mac.Sum(nil))
	match := hashedMessage == hashedPassword
	return match
}

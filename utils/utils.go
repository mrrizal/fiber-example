package utils

import (
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

func GetIDFromURLQuery(next, previous string, previousPage *bool) (int, error) {
	if previous != "" {
		*previousPage = true
		tempID, err := strconv.ParseInt(previous, 10, 32)
		if err != nil {
			return 0, err
		}
		return int(tempID), nil
	} else if next != "" {
		tempID, err := strconv.ParseInt(next, 10, 32)
		if err != nil {
			return 0, err
		}
		return int(tempID), nil
	}
	return 0, nil
}

type Hash struct{}

func (c *Hash) Generate(s string) (string, error) {
	saltedBytes := []byte(s)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hash := string(hashedBytes)
	return hash, nil
}

func (c *Hash) Compare(hash string, s string) error {
	incoming := []byte(s)
	existing := []byte(hash)
	return bcrypt.CompareHashAndPassword(existing, incoming)
}

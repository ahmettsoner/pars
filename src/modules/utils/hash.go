package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

func CalculateHash(content string) (string, error) {
	hash := md5.New()
	_, err := hash.Write([]byte(content))
	if err != nil {
		return "", err
	}
	hashInBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashInBytes)
	return hashString, nil
}

func CalculateHashFromObject(object interface{}) (string, error) {
	hash := md5.New()

	templateJsonData, err := json.Marshal(object)
	if err != nil {
		return "", err
	}

	_, err = hash.Write([]byte(templateJsonData))
	if err != nil {
		return "", err
	}
	hashInBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashInBytes)
	return hashString, nil
}

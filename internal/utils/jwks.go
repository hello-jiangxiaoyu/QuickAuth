package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/google/uuid"
	"os"
	"strings"
)

func GenerateKey(tenant string) (map[string][]byte, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	payload := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	key := uuid.New().String()
	if err = SetJWKS(tenant, key, payload); err != nil {
		return nil, err
	}

	return map[string][]byte{key: payload}, err
}

func SetJWKS(tenant string, key string, payload []byte) error {
	path := "data/jwks/" + tenant
	if _, err := os.ReadDir(path); err != nil {
		if err = os.MkdirAll(path, 0700); err != nil {
			return err
		}
	}

	var err error
	writeFile := fmt.Sprintf("%s/%s.key", path, key)
	if len(payload) == 0 {
		err = os.Remove(writeFile)
	} else {
		err = os.WriteFile(writeFile, payload, 0400)
	}

	return err
}

func GetJWKs(tenant string) (map[string][]byte, error) {
	path := "data/jwks/" + tenant
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	res := make(map[string][]byte)
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".key") == false {
			continue
		}

		name := path + "/" + file.Name()
		pemString, err := os.ReadFile(name)
		if err != nil {
			return nil, err
		}
		kid := strings.Split(file.Name(), ".")[0]
		res[kid] = pemString
	}

	return res, nil
}

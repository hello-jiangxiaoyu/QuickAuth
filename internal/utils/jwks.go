package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/go-jose/go-jose/v3"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"os"
	"strings"
)

func LoadRsaPrivateKeys(tenant string) (map[string]*rsa.PrivateKey, error) {
	res, err := GetJWKs(tenant)
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, errors.New("jwk is nil")
	}

	keys := make(map[string]*rsa.PrivateKey)
	for k, v := range res {
		key, err := jwt.ParseRSAPrivateKeyFromPEM(v)
		if err != nil {
			return nil, err
		}
		keys[k] = key
	}

	return keys, nil
}

func LoadRsaPublicKeys(tenant string) (*jose.JSONWebKeySet, error) {
	if tenant == "" {
		return nil, errors.New("tenant name should not be null")
	}
	var err error
	res := map[string][]byte{}
	if res, err = GetJWKs(tenant); err != nil || len(res) == 0 {
		if res, err = GenerateKey(tenant); err != nil {
			return nil, err
		}
	}

	var jwkSet jose.JSONWebKeySet
	var key *rsa.PrivateKey
	for k, v := range res {
		key, err = jwt.ParseRSAPrivateKeyFromPEM(v)
		if err != nil {
			return nil, err
		}

		jwk := jose.JSONWebKey{
			Key:                       key.Public(),
			KeyID:                     k,
			Algorithm:                 "RS256",
			Use:                       "sig",
			Certificates:              []*x509.Certificate{},
			CertificateThumbprintSHA1: []uint8{},
		}
		jwkSet.Keys = append(jwkSet.Keys, jwk)
	}

	return &jwkSet, nil
}

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
	var err error
	err = SetJWKSFile(tenant, key, payload)

	return err
}

func GetJWKs(tenant string) (res map[string][]byte, err error) {
	res, err = getJWKSFile(tenant)
	return res, err
}

func SetJWKSFile(tenant string, key string, value []byte) error {
	path := "data/jwks/" + tenant
	if _, err := os.ReadDir(path); err != nil {
		if err = os.MkdirAll(path, 0700); err != nil {
			return err
		}
	}

	var err error
	writeFile := fmt.Sprintf("%s/%s.key", path, key)
	if len(value) == 0 {
		err = os.Remove(writeFile)
	} else {
		err = os.WriteFile(writeFile, value, 0400)
	}

	return err
}

func getJWKSFile(tenant string) (map[string][]byte, error) {
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

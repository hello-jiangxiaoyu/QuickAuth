package service

import (
	"QuickAuth/pkg/safe"
	"crypto/rsa"
	"crypto/x509"
	"errors"
	"github.com/go-jose/go-jose/v3"
	"github.com/golang-jwt/jwt/v5"
)

func (s *Service) LoadRsaPublicKeys(tenant string) (*jose.JSONWebKeySet, error) {
	if tenant == "" {
		return nil, errors.New("tenant name should not be null")
	}
	var err error
	res := map[string][]byte{}
	if res, err = safe.GetJWKs(tenant); err != nil || len(res) == 0 {
		if res, err = safe.GenerateKey(tenant); err != nil {
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

func LoadRsaPrivateKeys(tenantId string) (map[string]*rsa.PrivateKey, error) {
	res, err := safe.GetJWKs(tenantId)
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

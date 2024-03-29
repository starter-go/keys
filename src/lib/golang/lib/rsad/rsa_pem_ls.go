package rsad

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"github.com/starter-go/keys"
)

type pemLS struct {
}

func (ls *pemLS) loadPrivateKey(kd *keys.KeyData) (*rsa.PrivateKey, error) {
	block, err := ls.loadPemBlock(kd)
	if err != nil {
		return nil, err
	}
	der := block.Bytes
	return x509.ParsePKCS1PrivateKey(der)
}

func (ls *pemLS) storePrivateKey(key *rsa.PrivateKey) (*keys.KeyData, error) {
	der := ls.getPrivateKeyDER(key)
	block := &pem.Block{
		Type:  "RSA PRIVATE	KEY",
		Bytes: der,
	}
	buffer := new(bytes.Buffer)
	err := pem.Encode(buffer, block)
	if err != nil {
		return nil, err
	}
	kd := &keys.KeyData{
		Algorithm:   "RSA",
		Encoding:    "x509",
		Format:      "PEM",
		ContentType: "application/x-rsa-private-key-pem",
		Content:     buffer.Bytes(),
	}
	return kd, nil
}

func (ls *pemLS) getPrivateKeyDER(key *rsa.PrivateKey) []byte {
	der := x509.MarshalPKCS1PrivateKey(key)
	return der
}

func (ls *pemLS) loadPublicKey(kd *keys.KeyData) (*rsa.PublicKey, error) {
	block, err := ls.loadPemBlock(kd)
	if err != nil {
		return nil, err
	}
	der := block.Bytes
	return x509.ParsePKCS1PublicKey(der)
}

func (ls *pemLS) storePublicKey(key *rsa.PublicKey) (*keys.KeyData, error) {
	der := ls.getPublicKeyDER(key)
	block := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: der,
	}
	buffer := new(bytes.Buffer)
	err := pem.Encode(buffer, block)
	if err != nil {
		return nil, err
	}
	kd := &keys.KeyData{
		Algorithm:   "RSA",
		Encoding:    "x509",
		Format:      "PEM",
		ContentType: "application/x-rsa-public-key-pem",
		Content:     buffer.Bytes(),
	}
	return kd, nil
}

func (ls *pemLS) loadPemBlock(kd *keys.KeyData) (*pem.Block, error) {
	data := kd.Content
	for {
		b, _ := pem.Decode(data)
		if b == nil {
			return nil, fmt.Errorf("no PEM block in data")
		}
		return b, nil
	}
}

func (ls *pemLS) getPublicKeyDER(key *rsa.PublicKey) []byte {
	der := x509.MarshalPKCS1PublicKey(key)
	return der
}

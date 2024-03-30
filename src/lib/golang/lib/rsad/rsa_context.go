package rsad

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"hash"
	"io"
	"strings"

	"github.com/starter-go/keys"
	"github.com/starter-go/keys/src/lib/golang/lib"
)

type rsaContext struct {
	driverPublic  keys.PublicKeyDriver
	driverPrivate keys.PrivateKeyDriver
	driverBase    keys.Driver

	loaderPublic  keys.PublicKeyLoader
	loaderPrivate keys.PrivateKeyLoader

	keygen keys.PrivateKeyGenerator
}

type publicKeyContext struct {
	parent *rsaContext
	raw    *rsa.PublicKey
	driver keys.PublicKeyDriver
	facade keys.PublicKey
}

type privateKeyContext struct {
	parent *rsaContext
	public *publicKeyContext
	raw    *rsa.PrivateKey
	driver keys.PrivateKeyDriver
	facade keys.PrivateKey
}

type cipherContext struct {
	public    *publicKeyContext
	private   *privateKeyContext
	decrypter keys.Decrypter
	encrypter keys.Encrypter
	options   keys.Options
	mode      CipherMode
	hash      hash.Hash
}

type signContext struct {
	public  *publicKeyContext
	private *privateKeyContext
}

////////////////////////////////////////////////////////////////////////////////

func (inst *cipherContext) prepareRandom(e *keys.Encryption) io.Reader {
	r := inst.options.Random
	if r == nil {
		r = rand.Reader
	}
	return r
}

func (inst *cipherContext) prepareHash(e *keys.Encryption) hash.Hash {
	h := inst.hash
	if h == nil {
		h = sha256.New()
	}
	return h
}

func (inst *cipherContext) setOptions(opt *keys.Options) {

	alg := opt.Algorithm.String()
	alg = strings.ToLower(alg)
	mode := CipherModeOAEP

	if strings.Contains(alg, "session") {
		mode = CipherModeSessionKey
	} else if strings.Contains(alg, "pkcs1v15") {
		mode = CipherModePKCS1v15
	} else if strings.Contains(alg, "oaep") {
		mode = CipherModeOAEP
	}

	reader := new(lib.ComplexAlgorithmReader)
	reader.Init(opt.Algorithm)
	h, err := reader.ReadHash()
	if err != nil {
		h = crypto.SHA256
	}

	inst.hash = h.New()
	inst.mode = mode
}

////////////////////////////////////////////////////////////////////////////////

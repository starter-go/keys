package rsad

import (
	"crypto/rsa"

	"github.com/starter-go/keys"
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

// type optionContext struct {
// 	public  *publicKeyContext
// 	private *privateKeyContext
// }

// type cipherContext struct {
// 	optionContext

// 	decrypter keys.Decrypter
// 	encrypter keys.Encrypter
// }

// type signContext struct {
// 	optionContext

// 	signer   keys.Signer
// 	verifier keys.Verifier
// }

////////////////////////////////////////////////////////////////////////////////

// func (inst *optionContext) prepareRandom() io.Reader {
// 	r := inst.options.Random
// 	if r == nil {
// 		r = rand.Reader
// 	}
// 	return r
// }

// func (inst *optionContext) prepareHashIF() hash.Hash {
// 	h := inst.prepareHashID()
// 	return h.New()
// }

// func (inst *optionContext) prepareHashID() crypto.Hash {
// 	h := inst.hash
// 	if h == 0 {
// 		h = crypto.SHA256
// 	}
// 	return h
// }

// func (inst *optionContext) setOptions(opt *keys.Options) {

// alg := opt.Algorithm.String()
// alg = strings.ToLower(alg)
// padding  := opt.Padding

// if strings.Contains(alg, "session") {
// 	mode = CipherModeSessionKey
// } else if strings.Contains(alg, "pkcs1v15") {
// 	mode = CipherModePKCS1v15
// } else if strings.Contains(alg, "oaep") {
// 	mode = CipherModeOAEP
// } else if strings.Contains(alg, "pss") {
// 	mode = CipherModePSS
// }

// reader := new(lib.ComplexAlgorithmReader)
// reader.Init(opt.Algorithm)
// h, err := reader.ReadHash()
// if err != nil {
// 	h = crypto.SHA256
// }

// 	inst.hash = opt.Hash
// 	inst.padding = opt.Padding
// }

////////////////////////////////////////////////////////////////////////////////

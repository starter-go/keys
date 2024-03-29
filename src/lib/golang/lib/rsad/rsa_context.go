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

type configuredContext struct {
	public  *publicKeyContext
	private *privateKeyContext
}

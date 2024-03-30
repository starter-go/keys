package testunits

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/keys"
	"github.com/starter-go/units"
	"github.com/starter-go/vlog"
)

// UnitForRSA ...
type UnitForRSA struct {

	//starter:component

	_as func(units.Units) //starter:as(".")

	DriverManager keys.DriverManager //starter:inject("#")

}

func (inst *UnitForRSA) _impl() units.Units {
	return inst
}

// Units ...
func (inst *UnitForRSA) Units(list []*units.Registration) []*units.Registration {
	u1 := &units.Registration{
		Name:     "unit4rsa",
		Enabled:  true,
		Priority: 0,
		Test:     inst.run,
	}
	list = append(list, u1)
	return list
}

func (inst *UnitForRSA) run() error {

	driver1, err := inst.DriverManager.Find("rsa", keys.ClassPrivateKey)
	if err != nil {
		return err
	}
	driver2 := driver1.(keys.PrivateKeyDriver)

	opt := &keys.Options{Size: 1024 * 2}
	kp1, err := driver2.Generator().Generate(opt)
	if err != nil {
		return err
	}

	rsaPK := kp1.Native().Signer().(*rsa.PrivateKey)
	size := rsaPK.Size()
	vlog.Info("RSA key-size: %d", size*8)

	err = inst.tryExportKP(kp1)
	if err != nil {
		return err
	}

	err = inst.tryExportPublic(kp1)
	if err != nil {
		return err
	}

	err = inst.tryCipher(kp1)
	if err != nil {
		return err
	}

	err = inst.trySign(kp1)
	if err != nil {
		return err
	}

	return nil
}

func (inst *UnitForRSA) tryExportKP(kp1 keys.PrivateKey) error {

	want := new(keys.KeyData)
	have, err := kp1.Export(want)
	if err != nil {
		return err
	}

	kp2, err := kp1.Driver().Loader().Load(have)
	if err != nil {
		return err
	}

	h := crypto.SHA1
	f1 := kp1.Fingerprint(h)
	f2 := kp2.Fingerprint(h)
	if !bytes.Equal(f1, f2) {
		hex1 := lang.HexFromBytes(f1)
		hex2 := lang.HexFromBytes(f2)
		format := "tryExportKeyPair: key1 != key2, key1.fingerprint=%s, key2.fingerprint=%s"
		return fmt.Errorf(format, hex1, hex2)
	}

	return nil
}

func (inst *UnitForRSA) tryExportPublic(kp keys.PrivateKey) error {

	pub1 := kp.PublicKey()
	want := new(keys.KeyData)
	have, err := pub1.Export(want)
	if err != nil {
		return err
	}

	pub2, err := pub1.Driver().Loader().Load(have)
	if err != nil {
		return err
	}

	h := crypto.SHA1
	f1 := pub1.Fingerprint(h)
	f2 := pub2.Fingerprint(h)
	if !bytes.Equal(f1, f2) {
		hex1 := lang.HexFromBytes(f1)
		hex2 := lang.HexFromBytes(f2)
		format := "tryExportPublic: key1 != key2, key1.fingerprint=%s, key2.fingerprint=%s"
		return fmt.Errorf(format, hex1, hex2)
	}

	return nil
}

func (inst *UnitForRSA) tryCipher(kp keys.PrivateKey) error {

	data1 := make([]byte, 245)
	rand.Reader.Read(data1)

	opt := &keys.Options{
		Algorithm: "PKCS1v15",
	}

	encrypter, err := kp.PublicKey().NewEncrypter(opt)
	if err != nil {
		return err
	}

	decrypter, err := kp.NewDecrypter(opt)
	if err != nil {
		return err
	}

	en1 := &keys.Encryption{
		CipherText: nil,
		PlainText:  data1,
		IV:         nil,
	}
	err = encrypter.Encrypt(en1)
	if err != nil {
		return err
	}

	en2 := &keys.Encryption{
		CipherText: en1.CipherText,
		PlainText:  nil,
		IV:         nil,
	}
	err = decrypter.Decrypt(en2)
	if err != nil {
		return err
	}

	data2 := en2.PlainText

	if !bytes.Equal(data1, data2) {
		return fmt.Errorf("tryCipher: data1 != data2")
	}
	return nil
}

func (inst *UnitForRSA) trySign(kp keys.PrivateKey) error {

	data := make([]byte, 0)
	rand.Reader.Read(data)
	sum := sha256.Sum256(data)

	opt := &keys.Options{}
	sig := &keys.Signature{}

	opt.Algorithm = "SHA256"
	sig.Digest = sum[:]
	sig.Signature = nil

	signer, err := kp.NewSigner(opt)
	if err != nil {
		return err
	}

	err = signer.Sign(sig)
	if err != nil {
		return err
	}

	verifier, err := kp.PublicKey().NewVerifier(opt)
	if err != nil {
		return err
	}

	return verifier.Verify(sig)
}

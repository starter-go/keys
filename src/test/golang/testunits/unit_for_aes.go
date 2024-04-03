package testunits

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"fmt"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/keys"
	"github.com/starter-go/units"
	"github.com/starter-go/vlog"
)

// UnitForAES ...
type UnitForAES struct {

	//starter:component

	_as func(units.Units) //starter:as(".")

	DriverManager keys.DriverManager //starter:inject("#")

}

func (inst *UnitForAES) _impl() units.Units {
	return inst
}

// Units ...
func (inst *UnitForAES) Units(list []*units.Registration) []*units.Registration {
	u1 := &units.Registration{
		Name:     "unit4aes",
		Enabled:  true,
		Priority: 0,
		Test:     inst.run,
	}
	list = append(list, u1)
	return list
}

func (inst *UnitForAES) run() error {

	driver1, err := inst.DriverManager.Find("aes", keys.ClassSecretKey)
	if err != nil {
		return err
	}
	driver2 := driver1.(keys.SecretKeyDriver)

	opt := &keys.Options{Size: 256}
	sk1, err := driver2.Generator().Generate(opt)
	if err != nil {
		return err
	}

	block, err := sk1.Native().NewCipher()
	if err != nil {
		return err
	}
	size := block.BlockSize()
	vlog.Info("AES block-size(in bytes): %d", size)

	err = inst.tryExportKey(sk1)
	if err != nil {
		return err
	}

	err = inst.tryCipher(sk1)
	if err != nil {
		return err
	}

	return nil
}

func (inst *UnitForAES) tryExportKey(sk1 keys.SecretKey) error {

	want := new(keys.KeyData)
	want.Format = "pem"

	have, err := sk1.Export(want)
	if err != nil {
		return err
	}

	sk2, err := sk1.Driver().Loader().Load(have)
	if err != nil {
		return err
	}

	h := crypto.MD5
	sum1 := sk1.Fingerprint(h)
	sum2 := sk2.Fingerprint(h)
	hex1 := lang.HexFromBytes(sum1)
	hex2 := lang.HexFromBytes(sum2)

	if hex1 != hex2 {
		f := "Fingerprint1 != Fingerprint2 [f1:%s f2:%s]"
		return fmt.Errorf(f, hex1, hex2)
	}
	return nil
}

func (inst *UnitForAES) tryCipher(sk1 keys.SecretKey) error {

	// make sk2

	want := new(keys.KeyData)
	want.Format = "pem"

	have, err := sk1.Export(want)
	if err != nil {
		return err
	}

	sk2, err := sk1.Driver().Loader().Load(have)
	if err != nil {
		return err
	}

	// make data1

	data1 := make([]byte, 1023*233)
	rand.Reader.Read(data1)
	block, err := sk2.Native().NewCipher()
	if err != nil {
		return err
	}

	opt := &keys.Options{
		Flow:    keys.FlowModeCBC,
		Padding: keys.PaddingPKCS7,
		Random:  rand.Reader,
	}
	iv := make([]byte, block.BlockSize())
	opt.Random.Read(iv)
	opt.IV = iv

	// encrypt
	enc, err := sk1.NewEncrypter(opt)
	if err != nil {
		return err
	}
	e1 := &keys.Encryption{
		PlainText: data1,
	}
	err = enc.Encrypt(e1)
	if err != nil {
		return err
	}

	// decrypt
	dec, err := sk2.NewDecrypter(opt)
	if err != nil {
		return err
	}
	e2 := &keys.Encryption{
		CipherText: e1.CipherText,
		IV:         e1.IV,
	}
	err = dec.Decrypt(e2)
	if err != nil {
		return err
	}
	data2 := e2.PlainText

	// check
	eq := bytes.Equal(data1, data2)
	if !eq {
		return fmt.Errorf("UnitForAES.tryCipher: data1 != data2")
	}
	return nil
}

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

	iv := make([]byte, block.BlockSize())
	random := rand.Reader
	random.Read(iv)

	opt := &keys.Options{
		Flow:    keys.FlowModeCBC,
		Padding: keys.PaddingPKCS7,
		Random:  random,
		IV:      iv,
	}

	// encrypt
	enc := sk1.Encrypter()
	e1 := &keys.Crypt{
		IV:      opt.IV,
		Padding: opt.Padding,
		Flow:    opt.Flow,
		Random:  opt.Random,

		PlainText: data1,
	}
	err = enc.Encrypt(e1)
	if err != nil {
		return err
	}

	// decrypt
	dec := sk2.Decrypter()
	e2 := &keys.Crypt{
		IV:      opt.IV,
		Padding: opt.Padding,
		Flow:    opt.Flow,
		Random:  opt.Random,

		CipherText: e1.CipherText,
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

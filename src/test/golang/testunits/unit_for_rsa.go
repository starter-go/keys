package testunits

import (
	"bytes"
	"crypto"
	"fmt"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/keys"
	"github.com/starter-go/units"
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

	driver1, err := inst.DriverManager.Find("rsa", keys.ClassKeyPair)
	if err != nil {
		return err
	}
	driver2 := driver1.(keys.KeyPairDriver)

	kp1, err := driver2.Generator().Generate(&keys.Options{})
	if err != nil {
		return err
	}

	err = inst.tryExportKP(kp1)
	if err != nil {
		return err
	}

	err = inst.tryExportPublic(kp1)
	if err != nil {
		return err
	}

	err = inst.tryCrypt(kp1)
	if err != nil {
		return err
	}

	err = inst.trySign(kp1)
	if err != nil {
		return err
	}

	return nil
}

func (inst *UnitForRSA) tryExportKP(kp1 keys.KeyPair) error {

	want := new(keys.KeyData)
	have, err := kp1.Export(want)
	if err != nil {
		return err
	}

	kp2, err := kp1.KeyPairDriver().Loader().Load(have)
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

func (inst *UnitForRSA) tryExportPublic(kp keys.KeyPair) error {

	pub1 := kp.PublicKey()
	want := new(keys.KeyData)
	have, err := pub1.Export(want)
	if err != nil {
		return err
	}

	pub2, err := pub1.PublicKeyDriver().Loader().Load(have)
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

func (inst *UnitForRSA) tryCrypt(kp keys.KeyPair) error {

	return fmt.Errorf("no impl: tryCrypt")
}

func (inst *UnitForRSA) trySign(kp keys.KeyPair) error {

	return fmt.Errorf("no impl: trySign")
}

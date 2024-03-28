package testunits

import "github.com/starter-go/units"

// Unit1 ...
type Unit1 struct {

	//starter:component

	_as func(units.Units) //starter:as(".")

}

func (inst *Unit1) _impl() units.Units {
	return inst
}

// Units ...
func (inst *Unit1) Units(list []*units.Registration) []*units.Registration {
	u1 := &units.Registration{
		Name:     "unit1",
		Enabled:  true,
		Priority: 0,
		Test:     inst.run,
	}
	list = append(list, u1)
	return list
}

func (inst *Unit1) run() error {
	return nil
}

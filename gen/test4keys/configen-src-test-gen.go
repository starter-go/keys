package test4keys
import (
    pc38c9ad22 "github.com/starter-go/keys"
    p67199fd98 "github.com/starter-go/keys/src/test/golang/testunits"
     "github.com/starter-go/application"
)

// type p67199fd98.Unit1 in package:github.com/starter-go/keys/src/test/golang/testunits
//
// id:com-67199fd9885079ff-testunits-Unit1
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type p67199fd988_testunits_Unit1 struct {
}

func (inst* p67199fd988_testunits_Unit1) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-67199fd9885079ff-testunits-Unit1"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p67199fd988_testunits_Unit1) new() any {
    return &p67199fd98.Unit1{}
}

func (inst* p67199fd988_testunits_Unit1) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p67199fd98.Unit1)
	nop(ie, com)

	


    return nil
}



// type p67199fd98.UnitForAES in package:github.com/starter-go/keys/src/test/golang/testunits
//
// id:com-67199fd9885079ff-testunits-UnitForAES
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type p67199fd988_testunits_UnitForAES struct {
}

func (inst* p67199fd988_testunits_UnitForAES) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-67199fd9885079ff-testunits-UnitForAES"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p67199fd988_testunits_UnitForAES) new() any {
    return &p67199fd98.UnitForAES{}
}

func (inst* p67199fd988_testunits_UnitForAES) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p67199fd98.UnitForAES)
	nop(ie, com)

	
    com.DriverManager = inst.getDriverManager(ie)


    return nil
}


func (inst*p67199fd988_testunits_UnitForAES) getDriverManager(ie application.InjectionExt)pc38c9ad22.DriverManager{
    return ie.GetComponent("#alias-c38c9ad22b7867d5ce346589e145db9f-DriverManager").(pc38c9ad22.DriverManager)
}



// type p67199fd98.UnitForRSA in package:github.com/starter-go/keys/src/test/golang/testunits
//
// id:com-67199fd9885079ff-testunits-UnitForRSA
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type p67199fd988_testunits_UnitForRSA struct {
}

func (inst* p67199fd988_testunits_UnitForRSA) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-67199fd9885079ff-testunits-UnitForRSA"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p67199fd988_testunits_UnitForRSA) new() any {
    return &p67199fd98.UnitForRSA{}
}

func (inst* p67199fd988_testunits_UnitForRSA) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p67199fd98.UnitForRSA)
	nop(ie, com)

	
    com.DriverManager = inst.getDriverManager(ie)


    return nil
}


func (inst*p67199fd988_testunits_UnitForRSA) getDriverManager(ie application.InjectionExt)pc38c9ad22.DriverManager{
    return ie.GetComponent("#alias-c38c9ad22b7867d5ce346589e145db9f-DriverManager").(pc38c9ad22.DriverManager)
}



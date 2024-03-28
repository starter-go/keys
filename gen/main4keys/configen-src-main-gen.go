package main4keys
import (
    pae7be287e "github.com/starter-go/keys/src/main/golang/lib"
     "github.com/starter-go/application"
)

// type pae7be287e.Example in package:github.com/starter-go/keys/src/main/golang/lib
//
// id:com-ae7be287e0a7d2e2-lib-Example
// class:
// alias:
// scope:singleton
//
type pae7be287e0_lib_Example struct {
}

func (inst* pae7be287e0_lib_Example) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ae7be287e0a7d2e2-lib-Example"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pae7be287e0_lib_Example) new() any {
    return &pae7be287e.Example{}
}

func (inst* pae7be287e0_lib_Example) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pae7be287e.Example)
	nop(ie, com)

	


    return nil
}



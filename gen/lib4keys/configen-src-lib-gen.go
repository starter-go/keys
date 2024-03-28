package lib4keys
import (
    pa311542f0 "github.com/starter-go/keys/src/lib/golang/lib"
     "github.com/starter-go/application"
)

// type pa311542f0.Example in package:github.com/starter-go/keys/src/lib/golang/lib
//
// id:com-a311542f0684a399-lib-Example
// class:
// alias:
// scope:singleton
//
type pa311542f06_lib_Example struct {
}

func (inst* pa311542f06_lib_Example) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a311542f0684a399-lib-Example"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa311542f06_lib_Example) new() any {
    return &pa311542f0.Example{}
}

func (inst* pa311542f06_lib_Example) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa311542f0.Example)
	nop(ie, com)

	


    return nil
}



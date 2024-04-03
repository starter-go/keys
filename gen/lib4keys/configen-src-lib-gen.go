package lib4keys
import (
    pc38c9ad22 "github.com/starter-go/keys"
    pa311542f0 "github.com/starter-go/keys/src/lib/golang/lib"
    p1bab13f97 "github.com/starter-go/keys/src/lib/golang/lib/aesd"
    p3e2e62794 "github.com/starter-go/keys/src/lib/golang/lib/rsad"
     "github.com/starter-go/application"
)

// type pa311542f0.DriverManagerImpl in package:github.com/starter-go/keys/src/lib/golang/lib
//
// id:com-a311542f0684a399-lib-DriverManagerImpl
// class:
// alias:alias-c38c9ad22b7867d5ce346589e145db9f-DriverManager
// scope:singleton
//
type pa311542f06_lib_DriverManagerImpl struct {
}

func (inst* pa311542f06_lib_DriverManagerImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a311542f0684a399-lib-DriverManagerImpl"
	r.Classes = ""
	r.Aliases = "alias-c38c9ad22b7867d5ce346589e145db9f-DriverManager"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa311542f06_lib_DriverManagerImpl) new() any {
    return &pa311542f0.DriverManagerImpl{}
}

func (inst* pa311542f06_lib_DriverManagerImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa311542f0.DriverManagerImpl)
	nop(ie, com)

	
    com.RawDrivers = inst.getRawDrivers(ie)


    return nil
}


func (inst*pa311542f06_lib_DriverManagerImpl) getRawDrivers(ie application.InjectionExt)[]pc38c9ad22.Driver{
    dst := make([]pc38c9ad22.Driver, 0)
    src := ie.ListComponents(".class-c38c9ad22b7867d5ce346589e145db9f-Driver")
    for _, item1 := range src {
        item2 := item1.(pc38c9ad22.Driver)
        dst = append(dst, item2)
    }
    return dst
}



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



// type p1bab13f97.Driver in package:github.com/starter-go/keys/src/lib/golang/lib/aesd
//
// id:com-1bab13f97f46035c-aesd-Driver
// class:class-c38c9ad22b7867d5ce346589e145db9f-Driver
// alias:
// scope:singleton
//
type p1bab13f97f_aesd_Driver struct {
}

func (inst* p1bab13f97f_aesd_Driver) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1bab13f97f46035c-aesd-Driver"
	r.Classes = "class-c38c9ad22b7867d5ce346589e145db9f-Driver"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1bab13f97f_aesd_Driver) new() any {
    return &p1bab13f97.Driver{}
}

func (inst* p1bab13f97f_aesd_Driver) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1bab13f97.Driver)
	nop(ie, com)

	


    return nil
}



// type p3e2e62794.Driver in package:github.com/starter-go/keys/src/lib/golang/lib/rsad
//
// id:com-3e2e62794d06903c-rsad-Driver
// class:class-c38c9ad22b7867d5ce346589e145db9f-Driver
// alias:
// scope:singleton
//
type p3e2e62794d_rsad_Driver struct {
}

func (inst* p3e2e62794d_rsad_Driver) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-3e2e62794d06903c-rsad-Driver"
	r.Classes = "class-c38c9ad22b7867d5ce346589e145db9f-Driver"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p3e2e62794d_rsad_Driver) new() any {
    return &p3e2e62794.Driver{}
}

func (inst* p3e2e62794d_rsad_Driver) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p3e2e62794.Driver)
	nop(ie, com)

	


    return nil
}



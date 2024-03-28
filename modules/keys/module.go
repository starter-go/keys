package keys

import (
	"github.com/starter-go/application"
	"github.com/starter-go/keys"
	"github.com/starter-go/keys/gen/lib4keys"
	"github.com/starter-go/keys/gen/main4keys"
	"github.com/starter-go/keys/gen/test4keys"
)

// ModuleForMain  ...
func ModuleForMain() application.Module {
	mb := keys.NewMainModule()
	mb.Components(main4keys.ExportComponents)

	// mb.Depend(security.Module())

	return mb.Create()
}

// ModuleForTest ...
func ModuleForTest() application.Module {
	mb := keys.NewTestModule()
	mb.Components(test4keys.ExportComponents)
	return mb.Create()
}

// ModuleForLib ...
func ModuleForLib() application.Module {
	mb := keys.NewLibModule()
	mb.Components(lib4keys.ExportComponents)
	return mb.Create()
}

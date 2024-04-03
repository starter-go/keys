package keys

import (
	"embed"

	"github.com/starter-go/application"
)

const (
	theModuleName    = "github.com/starter-go/keys"
	theModuleVersion = "v0.0.1"
	theModuleEdition = 1
)

////////////////////////////////////////////////////////////////////////////////

const (
	theLibModuleResPath  = "src/lib/resources"
	theMainModuleResPath = "src/main/resources"
	theTestModuleResPath = "src/test/resources"
)

//go:embed "src/lib/resources"
var theLibModuleResFS embed.FS

//go:embed "src/main/resources"
var theMainModuleResFS embed.FS

//go:embed "src/test/resources"
var theTestModuleResFS embed.FS

////////////////////////////////////////////////////////////////////////////////

// NewLibModule ...
func NewLibModule() *application.ModuleBuilder {
	mb := new(application.ModuleBuilder)
	mb.Name(theModuleName + "#main")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleEdition)
	mb.EmbedResources(theLibModuleResFS, theLibModuleResPath)
	return mb
}

// NewMainModule ...
func NewMainModule() *application.ModuleBuilder {
	mb := new(application.ModuleBuilder)
	mb.Name(theModuleName + "#lib")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleEdition)
	mb.EmbedResources(theMainModuleResFS, theMainModuleResPath)
	return mb
}

// NewTestModule ...
func NewTestModule() *application.ModuleBuilder {
	mb := new(application.ModuleBuilder)
	mb.Name(theModuleName + "#test")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleEdition)
	mb.EmbedResources(theTestModuleResFS, theTestModuleResPath)
	return mb
}

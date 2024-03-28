package keys

import "embed"

const (
	theModuleName    = "github.com/starter-go/keys"
	theModuleVersion = "v0.0.0"
	theModuleEdition = 0
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
func NewLibModule() {
}

// NewMainModule ...
func NewMainModule() {
}

// NewTestModule ...
func NewTestModule() {
}

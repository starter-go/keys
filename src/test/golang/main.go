package main

import (
	"os"

	"github.com/starter-go/keys/modules/keys"
	"github.com/starter-go/units"
)

func main() {
	m := keys.ModuleForTest()
	r := units.NewRunner()
	r.Dependencies(m)
	r.Run(os.Args)
}

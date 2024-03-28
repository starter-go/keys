package keys

import (
	"testing"

	"github.com/starter-go/vlog"
)

func Test(t *testing.T) {

	list := make([]ComplexAlgorithm, 0)

	list = append(list, "RSA sha256 CBC")
	list = append(list, "RSA_sha256_CBC")
	list = append(list, "RSA.sha256.CBC")
	list = append(list, "RSAwithSHA256andCBC")
	list = append(list, "OAEP-SHA224-MGF1Padding")
	list = append(list, "OAEPwithSHA-224andMGF1Padding")

	for _, item := range list {
		ca1 := item
		ca2 := item.Normalize()
		vlog.Info("ComplexAlgorithm.Normalize: from [%s] to [%s]", ca1, ca2)
	}
}

package lib

import (
	"crypto"
	"fmt"
	"strings"

	"github.com/starter-go/keys"
)

type ComplexAlgorithmReader struct {
	items []string
}

func (inst *ComplexAlgorithmReader) Init(ca keys.ComplexAlgorithm) {
	inst.items = ca.Normalize().ToArray()
}

func (inst *ComplexAlgorithmReader) ReadHash() (crypto.Hash, error) {
	all := inst.items
	for _, item := range all {
		switch item {

		case "md4":
			return crypto.MD4, nil
		case "md5":
			return crypto.MD5, nil

		case "sha1":
			return crypto.SHA1, nil
		case "sha224":
			return crypto.SHA224, nil
		case "sha256":
			return crypto.SHA256, nil
		case "sha384":
			return crypto.SHA384, nil
		case "sha512":
			return crypto.SHA512, nil

		}
	}
	return 0, fmt.Errorf("no hash algorithm info")
}

func (inst *ComplexAlgorithmReader) Contains(keyword string) bool {
	keyword = strings.ToLower(keyword)
	all := inst.items
	for _, item := range all {
		if item == keyword {
			return true
		}
	}
	return false
}

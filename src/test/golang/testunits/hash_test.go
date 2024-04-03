package testunits

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"testing"
)

func TestHash(t *testing.T) {

	data := make([]byte, 1024)
	rand.Reader.Read(data)

	sum10 := sha256.Sum256(data)
	sum1 := sum10[:]

	md := sha256.New()
	// md.Reset()
	md.Write(data)
	sum2 := md.Sum(nil)

	t.Log("sum1 = ", sum1)
	t.Log("sum2 = ", sum2)

	if !bytes.Equal(sum1, sum2) {
		t.Errorf("sum1 != sum2")
	}
}

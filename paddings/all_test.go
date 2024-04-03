package paddings

import (
	"crypto/rand"
	"testing"

	"github.com/starter-go/base/lang"
)

func TestComputePaddingLength(t *testing.T) {
	blockSize := 5
	for i := 0; i < (blockSize * 3); i++ {
		buffer := make([]byte, i)
		plen := computePaddingLength(buffer, blockSize)
		f := "TestComputePaddingLength[block_size:%d data_length:%d padding_length:%d]"
		t.Logf(f, blockSize, i, plen)
	}
}

func TestPKCS7Padding(t *testing.T) {

	const blockSize = 7
	data1 := make([]byte, 21)
	rand.Reader.Read(data1)
	p := NewPKCS7Padding()

	data2, err := p.AddPadding(data1, blockSize)
	if err != nil {
		t.Error(err)
		return
	}

	data3, err := p.RemovePadding(data2, blockSize)
	if err != nil {
		t.Error(err)
		return
	}

	d1 := lang.HexFromBytes(data1)
	d2 := lang.HexFromBytes(data2)
	d3 := lang.HexFromBytes(data3)

	t.Logf("TestPKCS7Padding: blockSize = %d", blockSize)
	t.Logf("TestPKCS7Padding: data1 = %s", d1)
	t.Logf("TestPKCS7Padding: data2 = %s", d2)
	t.Logf("TestPKCS7Padding: data3 = %s", d3)

}

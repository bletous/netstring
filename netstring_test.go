package netstring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIt(t *testing.T) {

	var encoded = [][]byte{
		[]byte("Hello world!"),
		[]byte(""),
		[]byte("Goodbye world"),
	}

	out, err := Encode(encoded[0], encoded[1], encoded[2])
	assert.Nil(t, err)
	assert.Equal(t, "12:Hello world!,0:,13:Goodbye world,", string(out))

	decoded, err := Decode(out)
	assert.Nil(t, err)
	assert.Equal(t, encoded, decoded)

}

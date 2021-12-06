package misc

import (
	"github.com/iotaledger/wasp/packages/cryptolib"
	"strconv"
	"testing"
)

func BenchmarkED25519Sign(b *testing.B) {
	keyPair := cryptolib.NewKeyPair()
	for i := 0; i < b.N; i++ {
		d := []byte("DataToSign" + strconv.Itoa(i))
		_, _ = keyPair.PrivateKey.Sign(nil, d, nil)
	}
	//
	//assert.True(t, publicKey.VerifySignature(data, sig))
}

func BenchmarkED25519SignVerify(b *testing.B) {
	keyPair := cryptolib.NewKeyPair()

	for i := 0; i < b.N; i++ {
		d := []byte("DataToSign" + strconv.Itoa(i))
		sig, _ := keyPair.PrivateKey.Sign(nil, d, nil)
		if !keyPair.Verify(d, sig) {
			panic("very bad")
		}
	}
}

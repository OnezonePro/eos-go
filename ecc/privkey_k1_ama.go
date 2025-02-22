package ecc

import (
	"fmt"
	"github.com/armoniax/eos-go/btcsuite/btcd/btcec"
	"github.com/armoniax/eos-go/btcsuite/btcutil"
)

type innerK1AMPrivateKey struct {
	privKey *btcec.PrivateKey
}

func (k *innerK1AMPrivateKey) publicKey() PublicKey {
	return PublicKey{Curve: CurveK1, Content: k.privKey.PubKey().SerializeCompressed(), inner: &innerK1AMPublicKey{}}
}

func (k *innerK1AMPrivateKey) sign(hash []byte) (out Signature, err error) {
	if len(hash) != 32 {
		return out, fmt.Errorf("hash should be 32 bytes")
	}

	compactSig, err := k.privKey.SignCanonical(btcec.S256(), hash)

	if err != nil {
		return out, fmt.Errorf("canonical, %s", err)
	}

	return Signature{Curve: CurveK1, Content: compactSig, inner: &innerK1AMASignature{}}, nil
}

func (k *innerK1AMPrivateKey) string() string {
	wif, _ := btcutil.NewWIF(k.privKey, '\x80', false) // no error possible
	return wif.String()
}

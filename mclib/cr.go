package mclib

import (
	_ "bytes"
	_ "encoding/binary"
	_ "crypto/sha256"
	"crypto/rand"
	"crypto/ecdsa"
	"crypto/elliptic"
	"errors"
	_ "fmt"
	"math/big"
)

var MyEllipticCurve = elliptic.P224()


type Signature struct{
	R []byte
	S []byte
}

func NewKey()(*ecdsa.PrivateKey,error){
	return ecdsa.GenerateKey(MyEllipticCurve, rand.Reader)
}

func MySign(hashed []byte, priv *ecdsa.PrivateKey)(*Signature,error){
	r, s, err := ecdsa.Sign(rand.Reader, priv, hashed)

	if err != nil {
    	return nil, err
    }
	return &Signature{R:r.Bytes(),S:s.Bytes()},nil
}

func MyVerify(sign *Signature, pub *ecdsa.PublicKey, hashed []byte)(error){
	if !ecdsa.Verify(pub, hashed, (&big.Int{}).SetBytes(sign.R), (&big.Int{}).SetBytes(sign.S)) {
		return errors.New("Verify error")
    }

	return nil
}
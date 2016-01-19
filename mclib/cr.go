package mclib

import (
	_ "bytes"
	_ "encoding/binary"
	_ "crypto/sha256"
	"crypto/rand"
	"crypto/ecdsa"
	"errors"
	_ "fmt"
	"math/big"
)

type Signature struct{
	R []byte
	S []byte
}

func MySign(hashed []byte, priv *ecdsa.PrivateKey)(*Signature,error){
	r, s, err := ecdsa.Sign(rand.Reader, priv, hashed)
	if err != nil {
    	return nil, err
    }
	return &Signature{R:r.Bytes(),S:s.Bytes()},nil
}

func MyVerify(sign *Signature, pub *ecdsa.PublicKey, hashed []byte)(error){
	var r  = &big.Int{}
	r.SetBytes(sign.R) 
	var s  = &big.Int{} 
	s.SetBytes(sign.S) 
	if !ecdsa.Verify(pub, hashed, r, s) {
		return errors.New("Verify error")
    }

	return nil
}
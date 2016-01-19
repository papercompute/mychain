package mclib

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"testing"
	_ "fmt"
)

func TestSignVerify(t *testing.T) {
	priv, err := ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
    if err != nil {
    	t.Errorf("GenerateKey error: %v", err)
    	return
    }
    hashed:=ComputeSha256([]byte("1234567890abcdefg"))
    sign,err:=MySign(hashed,priv)
    if err != nil {
    	t.Errorf("MySign error: %v", err)
    	return
    }
    err=MyVerify(sign,&priv.PublicKey,hashed)
    if err != nil {
    	t.Errorf("MyVerify error: %v", err)
    	return
    }
}


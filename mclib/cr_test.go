package mclib

import (
	"testing"
)

func TestSignVerify(t *testing.T) {
	priv, err := NewKey()
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


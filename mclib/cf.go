package mclib

import (
	"crypto/ecdsa"
	_"crypto/elliptic"
)

var (
	MyPrivateKey *ecdsa.PrivateKey = nil
	MyPrivateKeyFile = "./mykey"
)

func MyPrivateKeyInit(){
	if MyPrivateKey==nil{
		MyPrivateKey,_=NewKey()
	}
}

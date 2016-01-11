package mclib

import (
	_ "bytes"
	_ "encoding/binary"
	_ "errors"
	"time"
	"crypto/sha256"
)
type TxHeader struct{
	CreatedAt     uint32
	To            []byte 
	From          []byte
	Hash          []byte
	Nonce         uint32	
}

type TxBody struct {
	Header        TxHeader
	Data          []byte
	Signature     []byte
}

func TxHash(data []byte) []byte{
	hash := sha256.New()
	hash.Write(data)
	return hash.Sum(nil)
}

func NewTx(to, from, data []byte) *TxBody {
	return &TxBody{
		Header: TxHeader{
			CreatedAt:uint32(time.Now().Unix()),
			To: to,
			From: from,
			Hash: TxHash(data),
		},
		Data: data,
	}
}
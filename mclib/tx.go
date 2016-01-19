package mclib

import (
	"time"
	"bytes"
	"encoding/binary"
	"crypto/sha256"
	"math/rand"
)

type TxHeader struct{
	From          []byte
	To            []byte 
	Time          uint32
	Nonce         uint32
	Hash          []byte
}

func (h *TxHeader) Bytes() []byte{
	buf := new(bytes.Buffer)
	buf.Write(h.From)
	buf.Write(h.To)
	binary.Write(buf, binary.LittleEndian, h.Time)
	binary.Write(buf, binary.LittleEndian, h.Nonce)
	buf.Write(h.Hash)
	return buf.Bytes()
}

type TxBody struct {
	Header        TxHeader
	Data          []byte
	Signature     []byte
}

func ComputeSha256(data []byte) []byte{
	hash := sha256.New()
	hash.Write(data)
	return hash.Sum(nil)
}

func NewTx(to, from, data []byte) *TxBody {
	return &TxBody{
		Header: TxHeader{
			Time:uint32(time.Now().Unix()),
			Nonce: rand.New(rand.NewSource(time.Now().UnixNano())).Uint32(),
			To: to,
			From: from,
			Hash: ComputeSha256(data),
		},
		Data: data,
		
	}
}

func (tx *TxBody) Hash() []byte {
	var b []byte
	b=append(b,tx.Header.Bytes()...)
	b=append(b,tx.Data...)
	return ComputeSha256(b)
}


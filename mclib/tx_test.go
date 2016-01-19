package mclib

import (
	_ "reflect"
	"testing"
	"fmt"
)

func TestNewTx(t *testing.T) {
	tx := NewTx([]byte("to boom"),[]byte("from boom"),[]byte("data data data"))
	fmt.Printf("%s,%s,%s",string(tx.Header.To),string(tx.Header.From),string(tx.Data))
	//t.Error("TestNewTx failed")
}

func TestTxVerify(t *testing.T) {
	//t.Error("TestTxVerify failed")
}
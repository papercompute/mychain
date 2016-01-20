package main 

import (
	_"log"
	"fmt"
	mc "github.com/papercompute/mychain/mclib"
)

func main() {
	fmt.Println("txtest")
	tx := mc.NewTx([]byte("to boom"),[]byte("from boom"),[]byte("data data data"))
	fmt.Printf("%s,%s,%s",string(tx.Header.To),string(tx.Header.From),string(tx.Data))

	priv, _ := mc.NewKey()
	fmt.Println("key %v",priv)
}
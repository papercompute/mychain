/*
 * Generates a private/public key pair in PEM format (not Certificate)
 *
 * The generated private key can be parsed with openssl as follows:
 * > openssl rsa -in key.pem -text
 *
 * The generated public key can be parsed as follows:
 * > openssl rsa -pubin -in pub.pem -text
 */
package main

import (
	"crypto/rsa"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func main() {

	priv, err := rsa.GenerateKey(rand.Reader, 2014);
	if err != nil {
		fmt.Println(err);
		return;
	}
	err = priv.Validate();
	if err != nil {
		fmt.Println("Validation failed.", err);
	}

	priv_der := x509.MarshalPKCS1PrivateKey(priv);

	priv_blk := pem.Block {
	Type: "RSA PRIVATE KEY",
	Headers: nil,
	Bytes: priv_der,
	};

	priv_pem := string(pem.EncodeToMemory(&priv_blk));

	fmt.Printf(priv_pem);


	pub := priv.PublicKey;
	pub_der, err := x509.MarshalPKIXPublicKey(&pub);
	if err != nil {
		fmt.Println("Failed to get der format for PublicKey.", err);
		return;
	}

	pub_blk := pem.Block {
	Type: "PUBLIC KEY",
	Headers: nil,
	Bytes: pub_der,
	}
	pub_pem := string(pem.EncodeToMemory(&pub_blk));
	fmt.Printf(pub_pem);
}
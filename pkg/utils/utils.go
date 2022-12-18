package utils

import (
	"encoding/json"
	"crypto/ecdsa"
	"crypto/x509"
	"crypto/sha256"
	"encoding/pem"
	"errors"
)

type Envelope struct {
	PayloadType string      `json:"payloadType"`
	Payload     string      `json:"payload"`
	Signatures  []Signature `json:"signatures"`
}

type Signature struct {
	KeyID string `json:"keyid"`
	Sig   string `json:"sig"`
	Cert  string `json:"cert"`
}

func GetCert(envelope []byte) ([]byte, error) {
	env := &Envelope{}
	if err := json.Unmarshal(envelope, env); err != nil {
		return nil, err
	}

	return []byte(env.Signatures[0].Cert), nil

}

func GetPubKeyFromCert(certIn string) (*ecdsa.PublicKey, error) {

	block, _ := pem.Decode([]byte(certIn))
	if block == nil {
		return nil, errors.New("failed to parse certificate PEM")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, errors.New("failed to parse certificate: " + err.Error())
	}

	pubKey := cert.PublicKey.(*ecdsa.PublicKey)

	// // print out the public key in PEM format
	// pubKeyBytes, err := x509.MarshalPKIXPublicKey(pubKey)
	// if err != nil {
	// 	return nil, errors.New("unable to marshal public key: " + err.Error())
	// }

	// pubKeyPem := pem.EncodeToMemory(&pem.Block{
	// 	Type:  "PUBLIC KEY",
	// 	Bytes: pubKeyBytes,
	// })

	// fmt.Println(string(pubKeyPem))

	return pubKey, nil
}

func VerifySignature(pubKey *ecdsa.PublicKey, payload []byte, sig []byte) (bool, error) {
	hash := sha256.Sum256(payload)
	verified := ecdsa.VerifyASN1(pubKey, hash[:], sig)
	return verified, nil
}

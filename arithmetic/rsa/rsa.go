// @File     : rsa
// @Author   : Ville
// @Time     : 19-9-17 下午3:15
// rsa
package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"golang.org/x/crypto/ssh"
)

func genRsaKey(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	private, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	return private, &private.PublicKey, nil
}

func encodePrivateKey(private *rsa.PrivateKey) []byte {
	return pem.EncodeToMemory(&pem.Block{
		Bytes:   x509.MarshalPKCS1PrivateKey(private),
		Headers: nil,
		Type:    "RSA PRIVATE KEY",
	})
}

func encodePublicKey(public *rsa.PublicKey) ([]byte, error) {
	publicBytes, err := x509.MarshalPKIXPublicKey(public)
	if err != nil {
		return nil, err
	}
	return pem.EncodeToMemory(&pem.Block{
		Bytes:   publicBytes,
		Type:    "PUBLIC KEY",
		Headers: nil,
	}), nil
}

func encodeSSHKey(public *rsa.PublicKey) ([]byte, error) {
	publicKey, err := ssh.NewPublicKey(public)
	if err != nil {
		return nil, err
	}
	return ssh.MarshalAuthorizedKey(publicKey), nil
}

func GenRsaKey(bits int) (private, public []byte, err error) {
	privKey, pubKey, err := genRsaKey(bits)
	if err != nil {
		return
	}

	priv := encodePrivateKey(privKey)

	pub, err := encodePublicKey(pubKey)
	if err != nil {
		return
	}

	private = make([]byte, len(priv))
	public = make([]byte, len(pub))

	copy(public, pub)
	copy(private, priv)

	return
}

func GenSSHKeyPair() (private []byte, public []byte, err error) {
	prvkey, pubkey, err := genRsaKey(2048)
	if err != nil {
		return
	}
	priv := encodePrivateKey(prvkey)
	pub, err := encodeSSHKey(pubkey)
	if err != nil {
		return
	}

	private = make([]byte, len(priv))
	public = make([]byte, len(pub))
	copy(public, pub)
	copy(private, priv)

	return
}

func SSHToPEM(skey []byte) (pemKey []byte, err error) {
	pub, _, _, _, err := ssh.ParseAuthorizedKey(skey)
	if err != nil {
		return
	}
	cr := pub.(ssh.CryptoPublicKey)
	crPub := cr.CryptoPublicKey()

	rsaPub := crPub.(*rsa.PublicKey)

	return encodePublicKey(rsaPub)
}

func PEMToSSH(pemKey []byte) (skey []byte, err error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(pemKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	return encodeSSHKey(pub)
}

type Cipher interface {
	// origData 明文
	// Encrypt 加密
	Encrypt(origData []byte) ([]byte, error)
	// cipherData 密文
	// Encrypt 解密
	Decrypt(cipherData []byte) ([]byte, error)
}

type RsaCipher struct {
	privateK []byte
	publicK  []byte
}

func NewRsaCipher(privateK []byte, publicK []byte) *RsaCipher {
	return &RsaCipher{privateK: privateK, publicK: publicK}
}

// 加密
func (r *RsaCipher) Encrypt(origData []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(r.publicK)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	switch pubInterface.(type) {
	case *rsa.PublicKey:
		//加密
		return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
	default:
		return nil, errors.New("public key reflect error")
	}
}

// 解密
func (r *RsaCipher) Decrypt(cipherData []byte) ([]byte, error) {
	//解密
	block, _ := pem.Decode(r.privateK)
	if block == nil {
		return nil, errors.New("private key error")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, cipherData)
}

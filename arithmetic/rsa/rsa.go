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

// genRsaKey make  rsa.PrivateKey rsa.PublicKey according to bits
func genRsaKey(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	private, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	return private, &private.PublicKey, nil
}

// encodePrivateKey rsa.PrivateKey converted  to pem data
func encodePrivateKey(private *rsa.PrivateKey) []byte {
	return pem.EncodeToMemory(&pem.Block{
		Bytes:   x509.MarshalPKCS1PrivateKey(private),
		Headers: nil,
		Type:    "RSA PRIVATE KEY",
	})
}

// encodePublicKey rsa.PublicKey converted  to pem data
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

// encodeSSHKey rsa.PublicKey converted  to ssh data
func encodeSSHKey(public *rsa.PublicKey) ([]byte, error) {
	publicKey, err := ssh.NewPublicKey(public)
	if err != nil {
		return nil, err
	}
	return ssh.MarshalAuthorizedKey(publicKey), nil
}

func parseSSHPublicKey(sKey []byte) (*rsa.PublicKey, error) {
	pub, _, _, _, err := ssh.ParseAuthorizedKey(sKey)
	if err != nil {
		return nil, err
	}
	// 根据ssh源码查到 ssh.PublicKey 被 ssh.rsaPublicKey 实现
	// 且 ssh.rsaPublicKey 实现了  ssh.CryptoPublicKey
	// 所以先转换成 ssh.CryptoPublicKey 调用 CryptoPublicKey 得到  crypto.PublicKey
	// *rsa.PublicKey 实现了 crypto.PublicKey
	cr := pub.(ssh.CryptoPublicKey)
	crPub := cr.CryptoPublicKey()
	rsaPub := crPub.(*rsa.PublicKey)
	return rsaPub, nil
}

// GenRsaKey Generating RSA Key Pairs
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

// SSHToPEM key pairs in SSH format converted to PEM format
func SSHToPEM(skey []byte) (pemKey []byte, err error) {
	pub, _, _, _, err := ssh.ParseAuthorizedKey(skey)
	if err != nil {
		return
	}
	// 根据ssh源码查到 ssh.PublicKey 被 ssh.rsaPublicKey 实现
	// 且 ssh.rsaPublicKey 实现了  ssh.CryptoPublicKey
	// 所以先转换成 ssh.CryptoPublicKey 调用 CryptoPublicKey 得到  crypto.PublicKey
	// *rsa.PublicKey 实现了 crypto.PublicKey
	cr := pub.(ssh.CryptoPublicKey)
	crPub := cr.CryptoPublicKey()
	rsaPub := crPub.(*rsa.PublicKey)
	return encodePublicKey(rsaPub)
}

// PEMToSSH key pairs in PEM format converted to SSH format
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

func encrypt(origData []byte, pbk []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(pbk)
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

func decrypt(enData []byte, pvk []byte) ([]byte, error) {
	//解密
	block, _ := pem.Decode(pvk)
	if block == nil {
		return nil, errors.New("private key error")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, enData)
}

type Cipher interface {
	// origData 明文
	// Encrypt 加密
	Encrypt(origData []byte) ([]byte, error)
	// cipherData 密文
	// Encrypt 解密
	Decrypt(cipherData []byte) ([]byte, error)
}

type PEMRsaCipher struct {
	privateK []byte
	publicK  []byte
}

func NewRsaCipher(privateK []byte, publicK []byte) *PEMRsaCipher {
	return &PEMRsaCipher{privateK: privateK, publicK: publicK}
}

// 加密
func (r *PEMRsaCipher) Encrypt(origData []byte) ([]byte, error) {
	return encrypt(origData, r.publicK)
}

// 解密
func (r *PEMRsaCipher) Decrypt(cipherData []byte) ([]byte, error) {
	return decrypt(cipherData, r.privateK)
}

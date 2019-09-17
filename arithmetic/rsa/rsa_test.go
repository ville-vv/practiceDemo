// @File     : rsa_test
// @Author   : Ville
// @Time     : 19-9-17 下午4:27
// rsa
package rsa

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGenRsaKey(t *testing.T) {
	priv, pub, err := GenRsaKey(2048)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(priv))

	fmt.Println(string(pub))

}

func TestRsaCipher_Pem(t *testing.T) {
	priv, pub, err := GenRsaKey(2048)
	if err != nil {
		t.Error(err)
		return
	}
	cip := NewRsaCipher(priv, pub)
	originData := []byte("Hello, I am Chinese")
	cipData, err := cip.Encrypt(originData)
	if err != nil {
		t.Error(err)
		return
	}
	deCryptData, err := cip.Decrypt(cipData)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("原文：", string(originData))
	fmt.Printf("密文：%v\n", cipData)
	fmt.Println("结果：", string(deCryptData))
}

func TestGenSSHKeyPair(t *testing.T) {
	priv, pub, err := GenSSHKeyPair()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(priv))

	fmt.Println(string(pub))
}

func TestSSHToPEM(t *testing.T) {
	_, pub, err := GenSSHKeyPair()
	if err != nil {
		t.Error(err)
		return
	}
	tpub, err := SSHToPEM(pub)
	if err != nil {
		t.Error(err)
		return
	}
	sshPub, err := PEMToSSH(tpub)
	if !reflect.DeepEqual(sshPub, pub) {
		t.Error("反解析数据不对")
		return
	}
	fmt.Println("反解析 OK", string(sshPub))
}

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

func TestSSHToPEMAndPEMToSSH(t *testing.T) {
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

func TestSSHToPEM(t *testing.T) {
	ssh1 := "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDGJUMgmB9jUAwXQCzhvNF3mQYuqAUNIlXSlOBObdSbb73gRwbLhQiyhpbosqBQsMoqTaLRWvd/Z3XEXfG2BvBD2cdT/mdSpgWEKpYlppFoMNqeLU1J+98J+1APSrBfL3DMird7vvhrF2+CVtaNECxGJ3mEWZwamJRV7W02/h9nn0uKMCLGzlh6HFwvZmHIhl8IakXWK0ql0wQyZzpxyKqRBXTSZrLX2apu0ztS8Q/WANZkSWSyQl2KaJcSN8poLfeGWM+EzeBP74upp07aq0wd5j/kShNMD5IFIndCdhKLYM9BzGltAIqSQ6fTXs/f79cxSyOPr4pVGYVs4UM11MYn	xxxx@qq.com"
	ssh2 := "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDGJUMgmB9jUAwXQCzhvNF3mQYuqAUNIlXSlOBObdSbb73gRwbLhQiyhpbosqBQsMoqTaLRWvd/Z3XEXfG2BvBD2cdT/mdSpgWEKpYlppFoMNqeLU1J+98J+1APSrBfL3DMird7vvhrF2+CVtaNECxGJ3mEWZwamJRV7W02/h9nn0uKMCLGzlh6HFwvZmHIhl8IakXWK0ql0wQyZzpxyKqRBXTSZrLX2apu0ztS8Q/WANZkSWSyQl2KaJcSN8poLfeGWM+EzeBP74upp07aq0wd5j/kShNMD5IFIndCdhKLYM9BzGltAIqSQ6fTXs/f79cxSyOPr4pVGYVs4UM11MYn xxxx@163.com"

	tpub1, err := SSHToPEM([]byte(ssh1))
	if err != nil {
		t.Error(err)
		return
	}

	tpub2, err := SSHToPEM([]byte(ssh2))
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(string(tpub1))
	fmt.Println(string(tpub2))

}

func TestRsaCipher_SSH(t *testing.T) {
	priv, pub, err := GenSSHKeyPair()
	if err != nil {
		t.Error(err)
		return
	}
	sPub, err := SSHToPEM(pub)
	if err != nil {
		t.Error(err)
		return
	}
	cip := NewRsaCipher(priv, sPub)
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

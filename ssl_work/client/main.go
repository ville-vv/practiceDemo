package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)
func addTrust(pool*x509.CertPool, path string) {
	aCrt, err := ioutil.ReadFile(path)
	if err!= nil {
		fmt.Println("ReadFile err:",err)
		return
	}
	pool.AppendCertsFromPEM(aCrt)
}



func UnCheckServerSSL()  {
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Post("https://uatsky.yesbank.in:444/app/uat/fundtransfer2R/getbalance","application/json", nil)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func TwoWaySSlCheck(){
	pool := x509.NewCertPool()

	addTrust(pool,"../sslfile/server.crt")

	cliCrt, err := tls.LoadX509KeyPair("../sslfile/ville_pem.crt", "../sslfile/ville.key")


	if err != nil {
		fmt.Println("Loadx509keypair err:", err)
		return
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:      pool,
			Certificates: []tls.Certificate{cliCrt},
		},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://localhost:8080")
	if err != nil {
		fmt.Println("Get error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func main(){
	TwoWaySSlCheck()
	//UnCheckServerSSL()
}

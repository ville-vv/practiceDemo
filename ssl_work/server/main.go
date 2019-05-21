
package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type httpsHandler struct {
}

func addTrust(pool*x509.CertPool, path string) {
	aCrt, err := ioutil.ReadFile(path)
	if err!= nil {
		fmt.Println("ReadFile err:",err)
		return
	}
	pool.AppendCertsFromPEM(aCrt)

}

func (*httpsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "golang https server!!!")
}

func main() {
	pool := x509.NewCertPool()
	// 这里是加载客户端提供的证书，最好是加载客户端提供的根证书
	addTrust(pool,"../sslfile/ca.crt")
	//addTrust(pool,"../sslfile/selfsigned.crt")
	addTrust(pool,"../sslfile/ville_pem.crt")


	s := &http.Server{
		Addr:    ":8080",
		Handler: &httpsHandler{},
		TLSConfig: &tls.Config{
			ClientCAs:  pool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}

	if err := s.ListenAndServeTLS("../sslfile/server.crt", "../sslfile/server.key"); err != nil {
		log.Fatal("ListenAndServeTLS err:", err)
	}
}

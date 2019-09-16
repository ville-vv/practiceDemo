package download_file_work

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"
)

func DownloadFile(remoteFile string) {
	fmt.Println(time.Now().Format("01-02-2006"))
	url3 := remoteFile
	httpReq := &http.Client{}
	body := bytes.NewBuffer([]byte(""))
	request, _ := http.NewRequest("GET", url3, body)
	resp, err := httpReq.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	pt := path.Base(url3)

	f, err := os.OpenFile(pt, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("OpenFile ", err)
	}
	defer f.Close()

	if _, err := io.Copy(f, resp.Body); err != nil {
		fmt.Println("Copy write fail ", err)
	}

}

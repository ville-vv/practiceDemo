
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"time"
)


func main(){
	fmt.Println(time.Now().Format("01-02-2006"))
	url3 := "http://wss.wanlaowo.com/record/recordings/1101001/b535d729-eca5-4296-8f94-24d4205f9f83.mp3"
	httpReq := &http.Client{}
	body := bytes.NewBuffer([]byte(""))
	request, _ := http.NewRequest("GET",url3, body)
	resp, err := httpReq.Do(request)
	if err != nil{
		fmt.Println(err)
	}
	defer resp.Body.Close()
	fd, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println("ReadAll ",err)
		return
	}
	pt := path.Base(url3)
	err = ioutil.WriteFile(pt, fd, 0666)
	if err != nil{
		fmt.Println("WriteFile ", err)
	}
}

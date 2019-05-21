package heapsterapi

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

/**
 * http get请求
 * @return : []byte body 数据
 * @return : error 错误信息
 */
func Get(url string) ([]byte, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("%s: %v", url, resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

/**
 * http Post请求
 * @param  : url  请求地址
 * @param  : body 请求参数
 * @return : []byte body 数据
 * @return : error 错误信息
 */
func PostBytes(url string, body []byte) ([]byte, error) {

	resp, err := http.Post(url, "application/octet-stream", bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("Post to %s: %v", url, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("%s: %v", url, resp.StatusCode)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Read response body: %v", err)
	}

	return b, nil
}

/**
 * http Post请求
 * @param  : url  请求地址
 * @param  : values 请求参数
 * @return : []byte body 数据
 * @return : error 错误信息
 */
func Post(url string, values url.Values) ([]byte, error) {

	resp, err := http.PostForm(url, values)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("%s: %v", url, resp.StatusCode)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

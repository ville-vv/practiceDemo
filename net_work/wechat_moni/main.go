package main

import (
	"bytes"
	"compress/flate"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type deflate struct {
}

// 解压
func (d *deflate) Decode(rd io.Reader) ([]byte, error) {
	return ioutil.ReadAll(flate.NewReader(rd))
}

// 压缩
func (d *deflate) Encode(rd io.Reader) ([]byte, error) {
	buf := bytes.NewBufferString("")
	wt, err := flate.NewWriter(buf, -1)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(wt, rd)
	if err != nil {
		wt.Close()
		return nil, err
	}
	wt.Close()
	return buf.Bytes(), nil
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

type WeChatMo struct {
	httpCli  *http.Client
	tranCli  *http.Transport
	deviceID int
}

type OneRequest struct {
	UID    string
	Method string
	Url    string
	Header map[string]string
}

func (w *WeChatMo) UserAgent() string {
	//                         Mozilla/5.0 (Linux; Android 9; STF-AL00 Build/HUAWEISTF-AL00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/66.0.3359.126 MQQBrowser/6.2 TBS/45016 Mobile Safari/537.36 MMWEBID/3078 MicroMessenger/7.0.9.1560(0x27000935) Process/tools NetType/WIFI Language/zh_CN ABI/arm64
	//						   Mozilla/5.0 (Linux; Android 9; STF-AL00 Build/HUAWEISTF-AL00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/66.0.3359.126 MQQBrowser/6.2 TBS/45016 Mobile Safari/537.36 MMWEBID/8489 MicroMessenger/7.0.9.1560(0x27000935) Process/tools NetType/WIFI Language/zh_CN ABI/arm64
	return fmt.Sprintf("Mozilla/5.0 (Linux; Android 10; STF-AL00 Build/HUAWEISTF-AL00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/66.0.3359.126 MQQBrowser/6.2 TBS/45016 Mobile Safari/537.36 MMWEBID/%d MicroMessenger/7.0.9.1560(0x27000935) Process/tools NetType/WIFI Language/zh_CN ABI/arm64", w.deviceID)
}

func (w *WeChatMo) SetCookie(header http.Header) string {
	return strings.ReplaceAll(header.Get("Set-Cookie"), "; path=/", "")
}

func (w *WeChatMo) DoRequest(one *OneRequest) {
	requestHome, err := http.NewRequest("GET", one.Url, nil)
	if err != nil {
		fmt.Printf(" http.NewRequest requestHome %v", err)
		return
	}
	requestHome.Header.Set("User-Agent", w.UserAgent())
	requestHome.Header.Set("Accept-Language", "zh-CN,en-US;q=0.9")
	for k, v := range one.Header {
		if v != "" {
			requestHome.Header.Set(k, v)
		}
	}
	resp1, err := w.httpCli.Do(requestHome)
	if err != nil {
		fmt.Printf(" httpCli.Do requestHome %v", err)
		return
	}
	defer resp1.Body.Close()
	body, err := ioutil.ReadAll(resp1.Body)
	if err != nil {
		fmt.Printf("ioutil.ReadAll %v", err)
		return
	}
	fmt.Println("response body :", string(body))
}

func (w *WeChatMo) DoTran(one *OneRequest) (local string, err error) {
	requestHome, err := http.NewRequest("GET", one.Url, nil)
	if err != nil {
		fmt.Println("Error http.NewRequest requestHome %v", err)
		return
	}
	requestHome.Header.Set("User-Agent", w.UserAgent())
	requestHome.Header.Set("Accept-Language", "zh-CN,en-US;q=0.9")
	for k, v := range one.Header {
		if v != "" {
			requestHome.Header.Set(k, v)
		}
	}
	resp1, err := w.tranCli.RoundTrip(requestHome)
	if err != nil {
		fmt.Printf("Error tranCli.RoundTrip requestHome %v", err)
		return
	}
	defer resp1.Body.Close()
	local = resp1.Header.Get("Location")
	if !strings.Contains(local, "http") && local != "" {
		local = "http:" + local
	}
	return
}

func main() {

	weChat := &WeChatMo{httpCli: &http.Client{}, deviceID: 4567, tranCli: &http.Transport{}}
	//cookie := fmt.Sprintf("PHPSESSID=%s", Md5(time.Now().String()))
	//weChat.DoRequest(&OneRequest{
	//	UID:    "1",
	//	Method: "GET",
	//	Url:    "http://www.rewfw.top/favicon.ico",
	//	Header: map[string]string{
	//		"Referer": "http://www.shuaizhenkj.cn/app/index.php?c=entry&do=show&m=xiaof_toupiao&i=434&sid=15490&id=5825&wxref=mp.weixin.qq.com",
	//		"Cookie":  cookie,
	//	},
	//})
	//
	//weChat.DoRequest(&OneRequest{
	//	UID:    "2",
	//	Method: "GET",
	//	Url:    "http://www.shuaizhenkj.cn/app/index.php?c=entry&do=statistics&m=xiaof_toupiao&i=434&sid=15490&id=5825&type=click&wxref=mp.weixin.qq.com",
	//	Header: map[string]string{
	//		"Referer":          "http://www.shuaizhenkj.cn/app/index.php?c=entry&do=show&m=xiaof_toupiao&i=434&sid=15490&id=5825&wxref=mp.weixin.qq.com",
	//		"Cookie":           cookie,
	//		"Connection":       "keep-alive",
	//		"X-Requested-With": "XMLHttpRequest",
	//	},
	//})
	//
	//weChat.DoRequest(&OneRequest{
	//	UID:    "2",
	//	Method: "GET",
	//	Url:    "http://www.shuaizhenkj.cn/app/index.php?c=entry&do=statistics&m=xiaof_toupiao&i=434&sid=15490&id=5825&type=click&wxref=mp.weixin.qq.com",
	//	Header: map[string]string{
	//		"Referer":          "http://www.shuaizhenkj.cn/app/index.php?c=entry&do=show&m=xiaof_toupiao&i=434&sid=15490&id=5825&wxref=mp.weixin.qq.com",
	//		"Cookie":           cookie,
	//		"Connection":       "keep-alive",
	//		"X-Requested-With": "XMLHttpRequest",
	//	},
	//})

	// /addons/xiaof_toupiao/index.php?c=entry&do=show&m=xiaof_toupiao&i=434&sid=15490&id=5825&referee=mn-zPky3693twPPckKWPpd6u6c9C&wxref=mp.weixin.qq.com&from=singlemessage
	local, err := weChat.DoTran(&OneRequest{
		UID:    "2",
		Method: "GET",
		Url:    "http://www.rewfw.top/addons/xiaof_toupiao/index.php?c=entry&do=show&m=xiaof_toupiao&i=434&sid=15490&id=5825&referee=mn-zPky3693twPPckKWPpd6u6c9C&wxref=mp.weixin.qq.com&from=groupmessage",
		Header: map[string]string{
			"User-Agent": "Mozilla/5.0 (Linux; Android 9; STF-AL00 Build/HUAWEISTF-AL00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/66.0.3359.126 MQQBrowser/6.2 TBS/45016 Mobile Safari/537.36 MMWEBID/3078 MicroMessenger/7.0.9.1560(0x27000935) Process/tools NetType/WIFI Language/zh_CN ABI/arm64",
			"Connection": "keep-alive",
		},
	})
	if err != nil {
		fmt.Printf("Error weChat.DoTran requestHome %v", err)
		return
	}

	if local == "" {
		return
	}
	localUrl, err := url.Parse(local)
	if err != nil {
		fmt.Printf("Error url.Parse requestHome %v", err)
		return
	}

	statistics := localUrl.Query().Get("statistics")
	if statistics == "" {
		fmt.Printf("Error statistics %v", err)
		return
	}
	fmt.Println("location:", localUrl.Host, "statistics:", statistics)
	host := fmt.Sprintf("http://%s/app/index.php?", "www.shuaizhenkj.cn")
	hostWithstatistics := fmt.Sprintf("http://%s/app/index.php?statistics=%s", "www.shuaizhenkj.cn", statistics)
	cookie := fmt.Sprintf("PHPSESSID=%s", Md5(time.Now().String()))
	// 请求展示页面
	weChat.DoRequest(&OneRequest{
		UID:    "2",
		Method: "GET",
		//Url:    hostWithstatistics + "&c=entry&do=show&m=xiaof_toupiao&i=434&sid=15490&id=5825&referee=mn-zPky3693twPPckKWPpd6u6c9C&wxref=mp.weixin.qq.com",
		Url: "http://www.shuaizhenkj.cn/app/index.php?statistics=675dec4d3230833e6a10ac1be4a77690&c=entry&do=show&m=xiaof_toupiao&i=434&sid=15490&id=5825&referee=mn-zPky3693twPPckKWPpd6u6c9C&wxref=mp.weixin.qq.com",
		Header: map[string]string{
			"Upgrade-Insecure-Requests": "1",
			"Cookie":                    cookie,
			"Connection":                "keep-alive",
			"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,image/wxpic,image/sharpp,image/apng,image/tpg,*/*;q=0.8",
		},
	})

	// xiaoftoupiao
	weChat.DoRequest(&OneRequest{
		UID:    "2",
		Method: "GET",
		Url:    "http://www.shuaizhenkj.cn/app/index.php?i=434&c=utility&a=visit&do=showjs&m=xiaof_toupiao",
		Header: map[string]string{
			"Referer":    hostWithstatistics + "&c=entry&do=show&m=xiaof_toupiao&i=434&sid=15490&id=5825&referee=mn-zPky3693twPPckKWPpd6u6c9C&wxref=mp.weixin.qq.com",
			"Cookie":     cookie,
			"Connection": "keep-alive",
		},
	})

	// xiaoftoupiao
	weChat.DoRequest(&OneRequest{
		UID:    "2",
		Method: "GET",
		Url:    "http://www.shuaizhenkj.cn/app/index.php?c=entry&do=statistics&m=xiaof_toupiao&i=434&sid=15490&id=5825&type=click&wxref=mp.weixin.qq.com",
		Header: map[string]string{
			"Referer":    hostWithstatistics + "&c=entry&do=show&m=xiaof_toupiao&i=434&sid=15490&id=5825&referee=mn-zPky3693twPPckKWPpd6u6c9C&wxref=mp.weixin.qq.com",
			"Cookie":     cookie,
			"Connection": "keep-alive",
		},
	})

	// 真正的点赞接口
	weChat.DoRequest(&OneRequest{
		UID:    "2",
		Method: "GET",
		Url:    host + "c=entry&do=vote&m=xiaof_toupiao&sid=15490&i=434&type=good&id=5825",
		Header: map[string]string{
			//"Referer":          hostWithstatistics + "&c=entry&do=show&m=xiaof_toupiao&i=434&sid=15490&id=5825&referee=mn-zPky3693twPPcKWPpd6u6c9C&wxref=mp.weixin.qq.com&from=groupmessage&wxref=mp.weixin.qq.com",
			"Cookie":           cookie,
			"Connection":       "keep-alive",
			"X-Requested-With": "XMLHttpRequest",
		},
	})
	fmt.Println("cookie:", cookie)
	fmt.Println("UA:", weChat.UserAgent())
}

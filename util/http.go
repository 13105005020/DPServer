package util

import (
	"fmt"
	"github.com/mikemintang/go-curl"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var cookies = "cy=15; cye=xiamen; dper=e7d73042a40d308ff8394c8ac3986fab21f31a0a76d8568ea02c5d1d643ea024b3e5cbec008ea446d89af47474777cb53944b1500d9e59b1fb5c0cc9a19cd8e2; ua=%E5%B0%8F%E6%B3%A2%E6%B3%A2%E8%B6%85%E7%88%B1%E5%90%83; ctu=fd177bee5827455dd08cafe28c495146e2e78f3af9d197538e0eb6296fbf9520; _lxsdk_cuid=16e3ae901b6c8-0148625c07bd8d-541d3410-2a3000-16e3ae901b7c8; _lxsdk=16e3ae901b6c8-0148625c07bd8d-541d3410-2a3000-16e3ae901b7c8; _hc.v=69147b78-7d67-ca22-f428-50099699f1e7.1572946380; ll=7fd06e815b796be3df069dec7836c3df; tg_list_scroll=0; JSESSIONID=F6788F365ABE1FC8621D7CA52DE209B1; _lx_utm=utm_source%3DBaidu%26utm_medium%3Dorganic; Hm_lvt_185e211f4a5af52aaffe6d9c1a2737f4=1572955715,1572959648; Hm_lpvt_185e211f4a5af52aaffe6d9c1a2737f4=1572959648; _lxsdk_s=16e3bb378e2-e0d-141-b48%7C%7C2"

func RequestGet(address string, useProxy int) []byte {
	defer func() {
		if r := recover(); r != nil {
			var errMsg interface{}
			switch err := r.(type) {
			case error:
				errMsg = err.Error()
			default:
				errMsg = err
			}
			Put(errMsg)
			time.Sleep(30 * time.Second)
			RequestGet(address, useProxy)
		}
	}()
	client := &http.Client{}
	if useProxy == 1 {
		proxy := func(_ *http.Request) (*url.URL, error) {
			return url.Parse("http://" + GetIp())
		}
		transport := &http.Transport{Proxy: proxy}
		client = &http.Client{Transport: transport}
	}
	req, err := http.NewRequest("GET", address, nil)
	Check(err)
	req.Header.Set("Cookie", cookies)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.70 Safari/537.36")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
	resp, err := client.Do(req)
	Check(err)
	defer resp.Body.Close()
	ret, _ := ioutil.ReadAll(resp.Body)
	if len(strings.Split(string(ret), "403 Forbidden")) > 1 {
		Put(string(ret))
		time.Sleep(1200 * time.Second)
		RequestGet(address, useProxy)
	}
	return ret
}

func RequestPost(url string, postData map[string]interface{}, headers map[string]string) string {
	queries := map[string]string{}
	cookieMap := map[string]string{}
	// 链式操作
	req := curl.NewRequest()
	for _, val := range strings.Split(cookies, ";") {
		arr := strings.Split(val, "=")
		cookieMap[arr[0]] = arr[1]
	}
	resp, err := req.
		SetUrl(url).
		SetHeaders(headers).
		SetCookies(cookieMap).
		SetQueries(queries).
		SetPostData(postData).
		Post()
	if err != nil {
		fmt.Println(err)
	}
	return resp.Body
}

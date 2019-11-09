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

func RequestGet(address, cookies string, useProxy int) []byte {
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
			RequestGet(address, cookies, useProxy)
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
		RequestGet(address, cookies, useProxy)
	}
	return ret
}

func RequestPost(url, cookies string, postData map[string]interface{}, queries map[string]string, headers map[string]string) string {
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

package util

import (
	"fmt"
	"github.com/mikemintang/go-curl"
	"io/ioutil"
	"net/http"
)

func RequestGet(url string) []byte {
	rep, _ := http.Get(url)
	ret, _ := ioutil.ReadAll(rep.Body)
	return ret
}

func RequestPost(url string, postData map[string]interface{}, cookies, headers map[string]string) string {
	queries := map[string]string{
	}
	// 链式操作
	req := curl.NewRequest()
	resp, err := req.
		SetUrl(url).
		SetHeaders(headers).
		SetCookies(cookies).
		SetQueries(queries).
		SetPostData(postData).
		Post()
	if err != nil {
		fmt.Println(err)
	}
	return resp.Body
}

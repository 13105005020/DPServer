package util

import (
	"fmt"
	"github.com/mikemintang/go-curl"
	"io/ioutil"
	"log"
	"net/http"
)

func RequestGet(url string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Cookie", "_lxsdk_cuid=16e20827ee7c8-0b168eca267105-541d3410-2a3000-16e20827ee763; _lxsdk=16e20827ee7c8-0b168eca267105-541d3410-2a3000-16e20827ee763; _hc.v=9d5a4fef-f784-5f83-0da1-7fb41dee24c0.1572503454; s_ViewType=10; ctu=2ffbcd1e3abf9b58d21d58edac28af0af3d961142549c4f83e7be4432d6f5ac9; cye=xiamen; tg_list_scroll=0; cy=15; ctu=752ce1bff975837c92902ea86da6ff4db4d20f85c36afc687abbfceb5447c84858097b40880291056f7ecba319a5d1ee; dper=3c0b529600c3b7c91e298b135243c2789ebf5ef365eff7f07627ac5ff9ce263fa9295e194ef5022eaf8516660c23235e97bec968bcf1a76257077c7813aab9a3c3410db19e2c32b21a4fa82049b13ca8a946eb80cec81b3ca139f87924a8b11e; ll=7fd06e815b796be3df069dec7836c3df; ua=Arumiy; _lx_utm=utm_source%3DBaidu%26utm_medium%3Dorganic; Hm_lvt_185e211f4a5af52aaffe6d9c1a2737f4=1572593393,1572617623; Hm_lpvt_185e211f4a5af52aaffe6d9c1a2737f4=1572617623; _lxsdk_s=16e2750283c-42f-647-7e7%7C%7C12; JSESSIONID=9C487957A9B243BF1FBAF261FC29C50F")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	ret, _ := ioutil.ReadAll(resp.Body)
	return ret
}

func RequestPost(url string, postData map[string]interface{}, cookies, headers map[string]string) string {
	queries := map[string]string{}
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

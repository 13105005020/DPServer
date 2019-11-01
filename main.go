package main

import (
	"awesomeProject/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	var ch chan int
	go func() {
		for {
			ticker1 := time.NewTimer(86399 * time.Second)
			go GetDraw()
			<-ticker1.C
		}
		ch <- 1
	}()
	//Sign()
	//报名即中通知
	go func() {
		defer func() {
			if r := recover(); r != nil {
				println(r)
			}
		}()
		for {
			ticker2 := time.NewTimer(20 * time.Second)
			println("run")
			GetFree()
			<-ticker2.C
		}
		ch <- 1
	}()
	<-ch
}

//todo报名操作，未实现
func Sign() {
	postData := map[string]interface{}{
		"offlineActivityId": 1172219874,
		"phoneNo":           "153****6717",
		"shippingAddress":   "",
		"extraCount":        "",
		"birthdayStr":       "",
		"email":             "",
		"marryDayStr":       "",
		"babyBirths":        "",
		"pregnant":          "",
		"marryStatus":       0,
		"comboId":           "",
		"branchId":          "",
		"usePassCard":       0,
		"passCardNo":        "",
		"isShareSina":       "false",
		"isShareQQ":         "false",
	}
	cookies := map[string]string{
	}
	headers := map[string]string{
		"Accept":           "application/json, text/javascript",
		"Accept-Encoding":  "gzip, deflate",
		"Accept-Language":  "zh-CN,zh;q=0.9",
		"Connection":       "keep-alive",
		"Content-Length":   "225",
		"Content-Type":     "application/x-www-form-urlencoded;charset=UTF-8;",
		"Cookie":           "_lxsdk_cuid=16e20827ee7c8-0b168eca267105-541d3410-2a3000-16e20827ee763; _lxsdk=16e20827ee7c8-0b168eca267105-541d3410-2a3000-16e20827ee763; _hc.v=9d5a4fef-f784-5f83-0da1-7fb41dee24c0.1572503454; s_ViewType=10; dper=0a60e057db8c68d08c43b4c0b906e85126d5e1fcd8e50497fe4a7e4d06ec6e1c165c283e614ef7a324ce112b637108d71c28de0319b274a86567207181558409e3ecc4c52a3d48f1c41c0596813837b6d9713ad4ffac32afe8d39b37476d2c74; ll=7fd06e815b796be3df069dec7836c3df; ua=%E5%B0%8F%E6%B3%A2%E6%B3%A2%E5%B0%91%E5%90%83%E7%82%B9; ctu=2ffbcd1e3abf9b58d21d58edac28af0af3d961142549c4f83e7be4432d6f5ac9; cye=xiamen; cy=1; _lx_utm=utm_source%3Ddp_pc_event; _lxsdk_s=16e20ee0b64-63b-a9c-73f%7C%7C170",
		"Host":             "s.dianping.com",
		"Origin":           "http://s.dianping.com",
		"Referer":          "http://s.dianping.com/event/1172219874",
		"User-Agent":       "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.70 Safari/537.36",
		"X-Request":        "JSON",
		"X-Requested-With": "XMLHttpRequest",
	}
	resp := util.RequestPost("http://s.dianping.com/ajax/json/activity/offline/saveApplyInfo", postData, cookies, headers)
	fmt.Printf("%+v\n", resp)
}

func GetFree() {
	var (
		data        map[string]interface{}
		sendHistory = make(map[string][]int)
	)
	f, err := os.OpenFile("send_history.txt", os.O_RDWR, 0600)
	content, err := ioutil.ReadAll(f)
	f.Close()
	Check(err)
	json.Unmarshal(content, &sendHistory)
	for cityId, emailSlice := range util.OpenCity {
		url := "http://m.dianping.com/activity/static/pc/ajaxList"
		cookies := map[string]string{
			"cy":          "15",
			"cye":         "xiamen",
			"_lx_utm":     "utm_source%3DBaidu%26utm_medium%3Dorganic",
			"_lxsdk_cuid": "16e20827ee7c8-0b168eca267105-541d3410-2a3000-16e20827ee763",
			"_lxsdk":      "16e20827ee7c8-0b168eca267105-541d3410-2a3000-16e20827ee763",
			"_hc.v":       "9d5a4fef-f784-5f83-0da1-7fb41dee24c0.1572503454",
			"s_ViewType":  "10",
			"_lxsdk_s":    "16e20827ee8-844-a5b-b38%7C%7C42",
		}
		headers := map[string]string{
			"User-Agent":   "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.70 Safari/537.36",
			"Content-Type": "application/json",
		}
		postData := map[string]interface{}{
			"cityId": cityId,
			"mode":   "",
			"page":   1,
			"type":   0,
		}
		req := util.RequestPost(url, postData, cookies, headers)
		json.Unmarshal([]byte(req), &data)
		for _, v := range data["data"].(map[string]interface{})["detail"].([]interface{}) {
			activeId := int(v.(map[string]interface{})["offlineActivityId"].(float64))
			sendEmail := GetSendEmail(activeId, emailSlice, sendHistory)
			if int(v.(map[string]interface{})["mode"].(float64)) == 5 {
				title := v.(map[string]interface{})["activityTitle"].(string)
				util.SendEmail("报名即中", title, sendEmail)
				for _, email := range sendEmail {
					sendHistory[email] = append(sendHistory[email], activeId)
				}
			}
		}
		content, err = json.Marshal(sendHistory)
		Check(err)
		f, err := os.OpenFile("send_history.txt", os.O_RDWR|os.O_TRUNC, 0600)
		_, err = f.Write(content)
		Check(err)
		f.Close()
	}
}

func Put(data interface{}) {
	fmt.Printf("%+v\n", data)
}

func GetSendEmail(activeId int, emailSlice []string, sendHistory map[string][]int) (sendEmail []string) {
	for _, val := range emailSlice {
		flags := 0
		for _, v := range sendHistory[val] {
			if v == activeId {
				flags = 1
			}
		}
		if flags == 0 {
			sendEmail = append(sendEmail, val)
		}
	}
	return
}

func GetMixed(a, b []string) (c []string) {
	for _, v := range a {
		if util.InArray(b, v) {
			c = append(c, v)
		}
	}
	return
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

//获取全国抽奖城市
func GetDraw() {
	var (
		data     map[string]interface{}
		drawCity []string
	)
	url := "http://m.dianping.com/activity/static/pc/ajaxList"
	for id := range util.City {
		postData := map[string]interface{}{
			"cityId": id,
			"mode":   "",
			"page":   1,
			"type":   0,
		}
		cookies := map[string]string{
			"cy":          "15",
			"cye":         "xiamen",
			"_lx_utm":     "utm_source%3DBaidu%26utm_medium%3Dorganic",
			"_lxsdk_cuid": "16e20827ee7c8-0b168eca267105-541d3410-2a3000-16e20827ee763",
			"_lxsdk":      "16e20827ee7c8-0b168eca267105-541d3410-2a3000-16e20827ee763",
			"_hc.v":       "9d5a4fef-f784-5f83-0da1-7fb41dee24c0.1572503454",
			"s_ViewType":  "10",
			"_lxsdk_s":    "16e20827ee8-844-a5b-b38%7C%7C42",
		}
		headers := map[string]string{
			"User-Agent":   "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.70 Safari/537.36",
			"Content-Type": "application/json",
		}
		req := util.RequestPost(url, postData, cookies, headers)
		json.Unmarshal([]byte(req), &data)
		for _, v := range data["data"].(map[string]interface{})["detail"].([]interface{}) {
			if int(v.(map[string]interface{})["mode"].(float64)) == 5 {
				Put(util.City[id][0])
				drawCity = append(drawCity, util.City[id][0])
			}
		}
	}
	content, _ := json.Marshal(util.Unique(drawCity))
	f, err := os.OpenFile("draw_city.txt", os.O_RDWR|os.O_TRUNC, 0600)
	Check(err)
	_, err = f.Write(content)
	Check(err)
	f.Close()
}

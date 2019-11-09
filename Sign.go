package main

import (
	"DPServer/util"
	"fmt"
)

//todo报名操作，未实现
func main() {
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
	resp := util.RequestPost("http://s.dianping.com/ajax/json/activity/offline/saveApplyInfo", postData, headers)
	fmt.Printf("%+v\n", resp)
}

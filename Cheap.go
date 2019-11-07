package main

import (
	"awesomeProject/util"
	"encoding/json"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			var errMsg interface{}
			switch err := r.(type) {
			case error:
				errMsg = err.Error()
			default:
				errMsg = err
			}
			util.Put(errMsg)
			main()
		}
	}()
	var (
		url         string
		detailMap   = make(map[string]interface{})
		goodsMap    = make(map[string]map[string]string)
		contentData string
		result      = make([]map[string]map[string]interface{}, 0)
		maxPage     int
	)
	for {
		for _, v := range util.Category {
			url = "http://t.dianping.com/list/xiamen-category_" + strconv.Itoa(v)
			data := util.RequestGet(url, 0)
			pageArr := util.GetBetween(string(data), "data-page=\"", "\"", 0)
			if len(pageArr) == 0 {
				maxPage = 0
			} else {
				maxPage, _ = strconv.Atoi(pageArr[len(pageArr)-2])
			}
			println(maxPage)
			for page := 0; page <= maxPage; page++ {
				url = "http://t.dianping.com/list/xiamen-category_" + strconv.Itoa(v)
				if page != 0 {
					url += "?pageIndex=" + strconv.Itoa(page)
				}
				data := util.RequestGet(url, 0)
				contentData = strings.Replace(string(data), "\n", "", -1)
				contentData = strings.Replace(contentData, " ", "", -1)
				idArr := util.GetBetween(contentData, "{'dealId':'", "'}\">", 0)
				titleArr := util.GetBetween(contentData, "<h3>", "</h3>", 10)
				detailArr := util.GetBetween(contentData, "<h4>", "</h4>", 0)
				for k, v := range idArr {
					goodsMap[v] = map[string]string{
						"tittle": titleArr[k],
						"detail": detailArr[k],
					}
				}
				ids := ""
				for k, v := range idArr {
					if k != 0 {
						ids += "%2C"
					}
					ids += v
				}
				detailData := util.RequestGet("http://t.dianping.com/jsonp/dealPromo?ids="+ids, 0)
				json.Unmarshal(detailData, &detailMap)
				var cheapData = make(map[string]map[string]interface{})
				if detailMap["msg"].(map[string]interface{})["promo"] == nil {
					continue
				}
				for k, v := range detailMap["msg"].(map[string]interface{})["promo"].(map[string]interface{}) {
					detail := v.([]interface{})[0].(map[string]interface{})
					cheapData[k] = map[string]interface{}{
						"tittle":      goodsMap[k]["tittle"],
						"detail":      goodsMap[k]["detail"],
						"isPriceLine": detail["isPriceLine"],
						"amount":      detail["amount"],
						"isEnable":    detail["isEnable"],
						"desc":        detail["desc"],
						"button":      detail["button"],
						"tag":         detail["tag"],
					}
				}
				util.Put(cheapData)
				if len(cheapData) > 0 {
					result = append(result, cheapData)
				}
				time.Sleep(time.Duration(rand.Intn(10)+3) * time.Second)
			}
			time.Sleep(time.Duration(rand.Intn(10)+3) * time.Second)
		}
		now := time.Now()
		sendDay := now.Add(time.Hour * 24)
		next := time.Date(sendDay.Year(), sendDay.Month(), sendDay.Day(), 12, 0, 0, 0, sendDay.Location())
		ticker1 := time.NewTimer(next.Sub(now))
		<-ticker1.C
	}
	util.Put(result)
	util.WriteTxt(result, "cheap.txt")
}

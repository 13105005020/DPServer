package main

import (
	"awesomeProject/util"
	"encoding/json"
	"time"
)

func main() {
	var (
		ch           chan int
		noDrawCityId []int
	)
	go func() {
		defer util.Recover()
		for {
			now := time.Now()
			sendDay := now.Add(time.Hour * 24)
			next := time.Date(sendDay.Year(), sendDay.Month(), sendDay.Day(), 12, 0, 0, 0, sendDay.Location())
			ticker1 := time.NewTimer(next.Sub(now))
			<-ticker1.C
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
				headers := map[string]string{
					"User-Agent":   "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.70 Safari/537.36",
					"Content-Type": "application/json",
				}
				req := util.RequestPost(url, postData, headers)
				json.Unmarshal([]byte(req), &data)
				if len(data["data"].(map[string]interface{})["detail"].([]interface{})) == 0 {
					noDrawCityId = append(noDrawCityId, id)
				}
				for _, v := range data["data"].(map[string]interface{})["detail"].([]interface{}) {
					if int(v.(map[string]interface{})["mode"].(float64)) == 5 {
						drawCity = append(drawCity, util.City[id][0])
					}
				}
			}
			content, _ := json.Marshal(util.Unique(drawCity))
			util.WriteTxt(util.Unique(drawCity), "draw_city.txt")
			util.WriteTxt(noDrawCityId, "no_draw_city.txt")
			util.SendEmail("DP助手", "今日天天抽城市列表:"+string(content), util.GetUser())
		}
		ch <- 1
	}()
	<-ch
}

package main

import (
	"DPServer/util"
	"encoding/json"
	"strconv"
	"time"
)

//报名即中邮件通知
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
		}
		main()
	}()
	for {
		var (
			data        map[string]interface{}
			sendHistory = make(map[string][]int)
		)
		util.ReadTxt("free_send_history.txt", &sendHistory)
		for cityId, emailSlice := range util.OpenCity {
			url := "https://m.dianping.com/activity/static/list?page=1&cityid=" + strconv.Itoa(cityId)+"&filter=1"
			req := util.RequestGet(url, 0)
			json.Unmarshal([]byte(req), &data)
			for _, v := range data["data"].(map[string]interface{})["mobileActivitys"].([]interface{}) {
				detail := v.(map[string]interface{})
				activeId := int(detail["offlineActivityId"].(float64))
				if int(detail["mode"].(float64)) == 16 {
					sendEmail := util.GetSendEmail(activeId, emailSlice, sendHistory)
					title := detail["title"].(string)
					regionName := detail["regionName"].(string)
					content := "(" + regionName + ")" + title
					content += "\n\n马上领取:https://evt.dianping.com/synthesislink/5518248.html?offlineActivityId=" + strconv.Itoa(activeId) + "&source=mShare&utm_source=mShare"
					//util.SendEmail("报名即中", content, sendEmail)
					for _, email := range sendEmail {
						sendHistory[email] = append(sendHistory[email], activeId)
					}
					if len(sendEmail) > 0 {
						util.PushWeChat("报名即中", "https://evt.dianping.com/synthesislink/5518248.html?offlineActivityId="+strconv.Itoa(activeId)+"&source=mShare&utm_source=mShare")
					}
				}
			}
			util.WriteTxt(sendHistory, "free_send_history.txt")
		}
		ticker := time.NewTimer(10 * time.Second)
		<-ticker.C
	}
}

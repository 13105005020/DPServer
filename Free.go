package main

import (
	"awesomeProject/util"
	"encoding/json"
	"io/ioutil"
	"os"
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
		f, err := os.OpenFile("free_send_history.txt", os.O_RDWR, 0600)
		content, err := ioutil.ReadAll(f)
		f.Close()
		util.Check(err)
		json.Unmarshal(content, &sendHistory)
		for cityId, emailSlice := range util.OpenCity {
			url := "https://m.dianping.com/activity/static/list?page=1&cityid=" + strconv.Itoa(cityId)
			req := util.RequestGet(url, 0)
			json.Unmarshal([]byte(req), &data)
			util.Put(data)
			for _, v := range data["data"].(map[string]interface{})["mobileActivitys"].([]interface{}) {
				detail := v.(map[string]interface{})
				activeId := int(detail["offlineActivityId"].(float64))
				if int(detail["mode"].(float64)) == 16 {
					sendEmail := util.GetSendEmail(activeId, emailSlice, sendHistory)
					title := detail["title"].(string)
					regionName := detail["regionName"].(string)
					content := "(" + regionName + ")" + title
					content += "\n\n马上领取:https://evt.dianping.com/synthesislink/5518248.html?offlineActivityId=" + strconv.Itoa(activeId) + "&source=mShare&utm_source=mShare"
					util.SendEmail("报名即中", content, sendEmail)
					for _, email := range sendEmail {
						sendHistory[email] = append(sendHistory[email], activeId)
					}
					if len(sendEmail) > 0 {
						util.PushWeChat("报名即中", content)
					}
				}
			}
			util.WriteTxt(sendHistory, "free_send_history.txt")
		}
		ticker := time.NewTimer(10 * time.Second)
		<-ticker.C
	}
}

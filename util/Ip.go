package util

import (
	"encoding/json"
	"strconv"
)

func GetIp() string {
	var (
		data map[string]interface{}
	)
	url := "http://localhost:8090/get"
	res := RequestGet(url,"", 0)
	json.Unmarshal(res, &data)
	ip := data["ip"].(string) + ":" + strconv.Itoa(int(data["port"].(float64)))
	Put(ip)
	return ip
}

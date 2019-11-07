package util

import (
	"fmt"
	"gopkg.in/gin-gonic/gin.v1/json"
	"math/rand"
	"os"
	"strings"
	"time"
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func Put(data ...interface{}) {
	for _, v := range data {
		fmt.Printf("%+v ", v)
	}
	fmt.Printf("\n")
}

func ReadTxt(data interface{}, fileName string) {

}

func SliceSlice(data []string, length int) (ret [][]string) {
	var t int
	for i := 0; i < len(data); i += length {
		if i+length > len(data) {
			t = len(data)
		} else {
			t = i + length
		}
		ret = append(ret, data[i:t])
	}
	return ret
}

func WriteTxt(data interface{}, fileName string) {
	content, err := json.Marshal(data)
	Check(err)
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_TRUNC, 0600)
	_, err = f.Write(content)
	Check(err)
	f.Close()
}

func GetBetween(text, start, end string, x int) (data []string) {
	for k, v := range strings.Split(text, start) {
		if k > x {
			data = append(data, strings.Split(v, end)[0])
		}
	}
	return data
}

func Recover() {
}

func GetMixed(a, b []string) (c []string) {
	for _, v := range a {
		if InArray(b, v) {
			c = append(c, v)
		}
	}
	return
}

func GetAgent() string {
	agent := [...]string{
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:50.0) Gecko/20100101 Firefox/50.0",
		"Opera/9.80 (Macintosh; Intel Mac OS X 10.6.8; U; en) Presto/2.8.131 Version/11.11",
		"Opera/9.80 (Windows NT 6.1; U; en) Presto/2.8.131 Version/11.11",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; 360SE)",
		"Mozilla/5.0 (Windows NT 6.1; rv:2.0.1) Gecko/20100101 Firefox/4.0.1",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; The World)",
		"User-Agent,Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_8; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
		"User-Agent, Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Maxthon 2.0)",
		"User-Agent,Mozilla/5.0 (Windows; U; Windows NT 6.1; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len := len(agent)
	return agent[r.Intn(len)]
}

func InArray(arr []string, str string) bool {
	ret := false
	for _, val := range arr {
		if str == val {
			ret = true
		}
	}
	return ret
}

func Unique(m []string) []string {
	s := make([]string, 0)
	sMap := make(map[string]string)
	for _, value := range m {
		if value != "" {
			//计算map长度
			length := len(sMap)
			sMap[value] = "1"
			//比较map长度, 如果map长度不相等， 说明key不存在
			if len(sMap) != length {
				s = append(s, value)
			}
		}
	}
	return s
}

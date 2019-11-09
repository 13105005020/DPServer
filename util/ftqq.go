package util

func PushWeChat(title, content string) {
	url := "https://sc.ftqq.com/SCU65931T7f37ab36afdd7f3f8eaf749e3637e0e85dc375bf42fca.send"
	data := map[string]interface{}{
	}
	queries:=map[string]string{
		"text": title,
		"desp": content,
	}
	rep := RequestPost(url,"", data,queries, map[string]string{})
	println(string(rep))
}

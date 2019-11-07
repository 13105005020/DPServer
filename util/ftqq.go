package util

func PushWeChat(title, content string) {
	url := "https://sc.ftqq.com/SCU65931T7f37ab36afdd7f3f8eaf749e3637e0e85dc375bf42fca.send?text=" + title
	data := map[string]interface{}{
		"desp": content,
	}
	rep := RequestPost(url, data, map[string]string{})
	println(string(rep))
}

package util

var OpenCity = map[int][]string{
	15: { //厦门
		"2754004610@qq.com",
		"1148031762@qq.com",
		"394248551@qq.com",
		"610621514@qq.com",
		"385679909@qq.com",
		"1026655727@qq.com",
		"545481714@qq.com",
		"875927220@qq.com",
		"240496436@qq.com",
		"614408807@qq.com",
		"1679113925@qq.com",
		"304791632@qq.com",
		"478811822@qq.com",
		"263951816@qq.com",
		"43584009@qq.com",
		"931388525@qq.com",
		"297879770@qq.com",
		"363783539@qq.com",
		"417780875@qq.com",
		"1054591767@qq.com",
		"369284609@qq.com",
		"806966581@qq.com",
		"532490908@qq.com",
		"1068168688@qq.com",
		"327605633@qq.com",
		"1296771512@qq.com",
		"17378220@qq.com",
		"380248450@qq.com",
		"147453169@qq.com",
		"854779874@qq.com",
		"374490803@qq.com",
		"857950784@qq.com",
		"498088301@qq.com",
		"1493338695@qq.com",
		"351728749@qq.com",
		"466892477@qq.com",
		"2962511479@qq.com",
		"376640959@qq.com",
		"282506129@qq.com",
		"342805658@qq.com",
		"110235271@qq.com",
		"297003475@qq.com",
		"1836475598@qq.com",
		"695322546@qq.com",
		"925727736@qq.com",
		"46050671@qq.com",
		"229274@qq.com",
		"844409717@qq.com",
		"907338034@qq.com",
		"759042365@qq.com",
		"229274@qq.com",
		"791370194@qq.com",
		"839694365@qq.com",
		"124455257@qq.com",
		"104398855@qq.com",
		"17486000@qq.com",
		"251686891@qq.com",
		"363783639@qq.com",
		"3462764694@qq.com",
		"574189615@qq.com",
		"617884250@qq.com",
		"12875530@qq.com",
		"122623678@qq.com",
		"1621035953@qq.com",
		"1522419289@qq.com",
		"17797246@qq.com",
		"285649243@qq.com",
		"624560691@qq.com",
		"573700968@qq.com",
		"1094705039@qq.com",
		"804231365@qq.com",
		"1341339287@qq.com",
	},
	14: { //福州
		"512932542@qq.com",
		"1148031762@qq.com",
	},
}
var Category = []int{
	34, 20, 19, 26, 166, 14, 18, 15, 25, 12, 24, 21, 22, 37, 132, 35, 17, 134, 27, 30, 133, 33, 130, 29, 23, 31, 32, 36, 137, 38,
}

func GetUser() (userList []string) {
	for _, v := range OpenCity {
		userList = append(userList, v...)
	}
	return Unique(userList)
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

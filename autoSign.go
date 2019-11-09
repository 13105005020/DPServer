package main

import (
	"DPServer/util"
	"encoding/json"
	"strconv"
	"time"
)

func main() {
	history := make(map[string][]int)
	phoneArr := map[string]string{
		"13105005020": "cy=15; cye=xiamen; dper=e7d73042a40d308ff8394c8ac3986fab21f31a0a76d8568ea02c5d1d643ea024b3e5cbec008ea446d89af47474777cb53944b1500d9e59b1fb5c0cc9a19cd8e2; ua=%E5%B0%8F%E6%B3%A2%E6%B3%A2%E8%B6%85%E7%88%B1%E5%90%83; ctu=fd177bee5827455dd08cafe28c495146e2e78f3af9d197538e0eb6296fbf9520; _lxsdk_cuid=16e3ae901b6c8-0148625c07bd8d-541d3410-2a3000-16e3ae901b7c8; _lxsdk=16e3ae901b6c8-0148625c07bd8d-541d3410-2a3000-16e3ae901b7c8; _hc.v=69147b78-7d67-ca22-f428-50099699f1e7.1572946380; ll=7fd06e815b796be3df069dec7836c3df; tg_list_scroll=0; JSESSIONID=F6788F365ABE1FC8621D7CA52DE209B1; _lx_utm=utm_source%3DBaidu%26utm_medium%3Dorganic; Hm_lvt_185e211f4a5af52aaffe6d9c1a2737f4=1572955715,1572959648; Hm_lpvt_185e211f4a5af52aaffe6d9c1a2737f4=1572959648; _lxsdk_s=16e3bb378e2-e0d-141-b48%7C%7C2",
		"18859283601": "cy=15; cye=xiamen; _lxsdk_cuid=16e4f0f3634c8-0befd097f0219-54133310-2a3000-16e4f0f3634c8; _lxsdk=16e4f0f3634c8-0befd097f0219-54133310-2a3000-16e4f0f3634c8; _hc.v=f7eabc07-9607-50be-36a7-95250196e3fd.1573284427; lgtoken=0238f2332-a929-4668-b69f-50457d2aa34b; dper=8239452f31fcdaea803ee87e1ef866ff1f421bddbf42ff1d1b0278440727b9a591a5400623f203f2509a5a26510990fe264688594f3d314fe42e39cc51217efe6b92da57cb6f2e21382d5620b2ab7f179fe0505bbb0542550fdbdfe22868e1df; ll=7fd06e815b796be3df069dec7836c3df; ua=%E5%B0%8F%E6%95%8F%E5%AD%90%E9%A5%BF%E4%BA%86%E8%A6%81%E5%90%83%E5%90%83%E5%90%83; ctu=7d8d8f53451d675a8d075756bbc18f4f80a19d2ff73b0d2945b24d4def258822; _lxsdk_s=16e4f0f235c-4a6-27a-df%7C%7C30",
		"15959035387": "cy=15; cye=xiamen; _lxsdk_cuid=16e4f10df36c8-0bcc8d14ec47e8-54133310-2a3000-16e4f10df36c8; _lxsdk=16e4f10df36c8-0bcc8d14ec47e8-54133310-2a3000-16e4f10df36c8; _hc.v=57e4e0d9-a97d-49a5-3030-00ef65fdd207.1573284536; lgtoken=0266ef76c-1b89-4827-9d85-3bdf21786514; dper=1b2563ecded9c51ad9176fe7f271995c874cf877d3bd257161df38d5d01147454f06296de6da2a9e3e0fd2a6672c50b32146c27d709a35e026adeb0b5e766801e1acd48c88d66f36055ac98e29c0a00edf3c6f10222d23824fdd98c60bc11ad1; ll=7fd06e815b796be3df069dec7836c3df; ua=Julia; ctu=1202c9fef0e271fd06c254716aaf6be79e525054b6b4a5170b56308136743d21; _lxsdk_s=16e4f10cb10-6d4-5f8-52c%7C%7C23",
		"15960849915": "cy=15; cye=xiamen; _lxsdk_cuid=16e4f16371ac8-02fffec785411-54133310-2a3000-16e4f16371bc8; _lxsdk=16e4f16371ac8-02fffec785411-54133310-2a3000-16e4f16371bc8; _hc.v=f7d2bf9e-100b-1720-83a6-bffd682d8698.1573284887; lgtoken=0180772b3-2ea4-4069-93d0-16c8e7b801b7; dper=f2a28104faa5f27e771e7dde250b2a4db84b94fc0d0dae92171c6293c9c55a486e52cc5bab34ab10a65561a8e47ecfddb36a15076c2d505a1ab4c0f16ba9c33a3f8e80a2c97567caa8ae283ae11af3f5d778d4ebb93b5ca2ea42b2cdd7d5fd2d; ll=7fd06e815b796be3df069dec7836c3df; ua=%E6%81%8D%E6%83%9A%E7%9A%84%E5%A4%A7%E9%A5%BC; ctu=30a6dc09fd0a2b7b3e139ac474b0251f9bd5da9d38a9fe99c8261c4680e2bab4; _lxsdk_s=16e4f16320f-8c0-b52-5a8%7C%7C20",
		"15880252258": "cy=15; cye=xiamen; _lxsdk_cuid=16e4f14a02ec8-0751dd5814f9b-54133310-2a3000-16e4f14a02ec8; _lxsdk=16e4f14a02ec8-0751dd5814f9b-54133310-2a3000-16e4f14a02ec8; _hc.v=fc4f967f-5bac-c7d3-44c7-0942d5189591.1573284782; lgtoken=09ebd6d3e-b0fe-4edf-9803-72455382db44; dper=b5fb70cb0f0fdbd6dd5dc99558da24d379619e58166d08f76f768ea7537b333bb295edc3d5669830612667adb95477915ab5240e883b10b6769760bd49f8e38d6190ab25200a30f32f61caca959c8051563b72dc0efdab784814aef65d0bd8fb; ll=7fd06e815b796be3df069dec7836c3df; ua=%E9%92%B1%E5%A4%9A%E5%A4%9A; ctu=def796423c30aa8a0d82801108ef77798e4d7db53f9b7ab3060756aca2f8ce85; _lxsdk_s=16e4f1498d0-9e8-b02-759%7C%7C24",
		"13216493510": "cye=xiamen; cy=15; _lxsdk_cuid=16e3f900236c8-03ec0e92505785-541d3410-2a3000-16e3f900236c8; _lxsdk=16e3f900236c8-03ec0e92505785-541d3410-2a3000-16e3f900236c8; _hc.v=f0588df9-a07d-4139-1e5d-ae75a5b2c37d.1573024433; PHOENIX_ID=0a49a8ba-16e439ad489-534d36; _tr.u=fFPjCYtEeD8LGyn8; tg_vg=-2074804131; cityid=15; logan_custom_report=; switchcityflashtoast=1; ctu=fd177bee5827455dd08cafe28c4951461661429fa16ddf385c65c6d4a7fea670; source=m_browser_test_33; dp_pwa_v_=fe25b31c1e622e79b2c551c1b73fa76ba19c48e7; default_ab=shop%3AA%3A5%7Cindex%3AA%3A3%7CshopList%3AC%3A4; logan_session_token=hm0957loylnx92etwnd8; _lx_utm=utm_source%3DBaidu%26utm_medium%3Dorganic; lgtoken=02ad47853-904e-4da8-8d95-924d400a7243; dper=0a868f4d0d29b8d55bc2097a04a7975a3a761ceb6e744cd698eaedaf7ce7f9e5671e1136d3d019eb1026ab70992777f03bdad2ca14092888826326e5522cf164fa1c7d3a81a2ea1ef3d626eccf76c7787f64b925f320fab9f8b563d60dbd5928; ll=7fd06e815b796be3df069dec7836c3df; ua=%E5%A4%A9%E5%A4%A9%E5%BC%80%E5%BF%83; _lxsdk_s=16e4edbe4eb-22c-51e-6b1%7C%7C183",
		"18106963607": "cy=15; cye=xiamen; _lxsdk_cuid=16e4f18137cc8-09bdb0bcf223ea-54133310-2a3000-16e4f18137cc8; _lxsdk=16e4f18137cc8-09bdb0bcf223ea-54133310-2a3000-16e4f18137cc8; _hc.v=9a136e71-e689-f9e0-5be5-da49cffbc617.1573285008; lgtoken=092e649b2-755f-424b-9ccb-44d7020b5a0a; dper=cd30138539399d9a0d954fd47afe993593e387f652c33e34a80faaad4f45245c9da8ee358c2fb61ad89dabf3c358843e755830c9f7e01784a19af7fc6b04bc8944098469b27ab6a23afd16be4c65f93bd5222f2cdd7b813a6a0594af7339fd87; ll=7fd06e815b796be3df069dec7836c3df; ua=%E9%B9%AD%E5%B2%9B%E5%B0%8Ftutor; ctu=c0bd139966f4c621dba7256a3f6b8894d06a2d74e36f8b73239bb15a10dc14d8; _lxsdk_s=16e4f180c1c-23a-a57-f4a%7C%7C20",
		"13275029989": "cy=15; cye=xiamen; _lxsdk_cuid=16e4f23ee1ec8-00518886d472ce-54133310-2a3000-16e4f23ee1ec8; _lxsdk=16e4f23ee1ec8-00518886d472ce-54133310-2a3000-16e4f23ee1ec8; _hc.v=b1337245-2027-63d5-17dc-f177c59b098e.1573285785; lgtoken=0d3f73c72-daba-48d5-bd81-5637c80176a6; dper=57f921bd5f8242f857c58f654dca7f179456fe2441dd23ff1d7c7f5744a9d498e34931376184a74eafbde136a290b282ac436cc8b11fda5fdf1df2dd2ddf3a257dba56e7db17da763932507d548c2b8cae8929adf098bb34c8ad9935dffc07ea; ll=7fd06e815b796be3df069dec7836c3df; ua=Xdd; ctu=a7932aacd5a97ca529d0289aed90126e6e556da6eff1d69ec0c1d4ab1bd3e08d; _lxsdk_s=16e4f23ecc1-37c-44b-9d0%7C%7C18",
		"15578628884": "cy=15; cye=xiamen; _lxsdk_cuid=16e4f529a6767-0cec6df9bcaa47-54133310-2a3000-16e4f529a68c8; _lxsdk=16e4f529a6767-0cec6df9bcaa47-54133310-2a3000-16e4f529a68c8; _hc.v=6edee61b-a6c2-28f3-0d61-3ff960223054.1573288844; lgtoken=0b8aceb8f-3f17-490d-b228-db1ce2cfefc9; dper=17317eb4b8d6415fb3543373f6335501de7080975b8e2c670f18a52ccfca40b5d0e9ffb0c717ec017d42958730722073e2e67355d8a3cbd2ef2bf4e0ae192429d665ab85022562d3bdc8032789e7b89c8c623a91f40fbda9194af7017d73246b; ll=7fd06e815b796be3df069dec7836c3df; ua=%E5%A4%A9%E5%85%89b; ctu=c38eb4bd91043101968dba2ccc0e4077123bb79a3934fcfe0e84dbc3bd1381e1; _lxsdk_s=16e4f529940-eb9-09f-4bb%7C%7C17",
		"17750014371": "_lxsdk_cuid=16e4f545bf2c8-0e8659cc890499-2d604637-4a574-16e4f545bf2c8; _lxsdk=16e4f545bf2c8-0e8659cc890499-2d604637-4a574-16e4f545bf2c8; _hc.v=4025c527-0fe9-e333-375e-61624c10f1fe.1573288961; cy=15; cye=xiamen; lgtoken=0b4bd1d2b-148b-436f-8941-37360d5f9fc4; dper=0b59d85eb53043907ec9ad1b444bd2ae324fda375109609b8f0c2ab0dd04548d858320fd07717cf2cd7266b9f3fd9669cd7f488794cbb426e247c6a59c289464e5597d0f3fa6ba042ce26e43d78016381c32c20c58f719b65834303d8e7f8480; ll=7fd06e815b796be3df069dec7836c3df; ua=%E6%B4%B2%E6%B4%B2; ctu=e2bd1a29853b3daa97dd29f0d910b678d1a18340d15eca29ddc4f24988ba9373; _lxsdk_s=16e4f545b03-e21-a95-826%7C%7C26",
		"18059892258": "cy=15; cye=xiamen; _lxsdk_cuid=16e505837d5c8-056ec415cee769-54133310-2a3000-16e505837d5c8; _lxsdk=16e505837d5c8-056ec415cee769-54133310-2a3000-16e505837d5c8; _hc.v=8456d520-cc46-a76f-5c63-b3093235e416.1573305989; lgtoken=0a6f89b29-b16c-4c9a-b354-f2e1e1c24607; dper=640ff6a0a5dff4e39cd4ce9f2f809766adde4f262133db600389fbce26a00d7f16c34cbd5c045724d19cc7babe2125760a799af22fbc34968002f88e5625bb6abe0d8e30b8b779f6cdbdd0980b32ec0be426ddda41f611dd123905d7f2d0d6d3; ll=7fd06e815b796be3df069dec7836c3df; ua=**%E6%B2%B9%E8%91%B1%E7%B2%BF%E7%B2%BF%5EO%5E; ctu=b55c40f79ca92ac445e6f9c688d588cb3166c628325969904da4d69e269a44ab; _lxsdk_s=16e50583692-5d7-1ab-c27%7C%7C17",
	}
	defer func() {
		if r := recover(); r != nil {
			println(r)
			main()
		}
	}()
	var historyIds []int
	for page := 1; ; page++ {
		url := "https://m.dianping.com/activity/static/list?page=" + strconv.Itoa(page) + "&filter=1&cityid=15&type=1"
		res := util.RequestGet(url, phoneArr["13105005020"], 0)
		var data map[string]interface{}
		json.Unmarshal(res, &data)
		for _, val := range data["data"].(map[string]interface{})["mobileActivitys"].([]interface{}) {
			id := int(val.(map[string]interface{})["offlineActivityId"].(float64))
			mode := int(val.(map[string]interface{})["mode"].(float64))
			if util.InArrayInt(historyIds, id) || mode != 16 {
				continue
			}
			historyIds = append(historyIds, id)
			util.Putln(val.(map[string]interface{})["title"])
			for phone, cookie := range phoneArr {
				if util.InArrayInt(history[phone], id) {
					continue
				}
				signUrl := "https://m.dianping.com/mobile/dinendish/apply/doApplyActivity"
				postData := map[string]interface{}{
					"env":               1,
					"offlineActivityId": id,
					"phoneNo":           phone,
				}
				header := map[string]string{
					"Accept":       " application/json, text/plain, */*",
					"Content-Type": " application/json;charset=UTF-8",
					"Origin":       " https://m.dianping.com",
					"Referer":      " https://m.dianping.com/mobile/dinendish/apply/1331624928?a=1&source=null&utm_source=null&showShopId=0&token=%2a&uiwebview=1",
					"User-Agent":   " Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1",
				}
				util.Putln(phone)
				go func() {
					flag:=0
					for i := 0; i < 5; i++ {
						resp := util.RequestPost(signUrl, cookie, postData, map[string]string{}, header)
						var respMap map[string]interface{}
						json.Unmarshal([]byte(resp), &respMap)
						util.Putln(string(resp))
						switch int(respMap["data"].(map[string]interface{})["code"].(float64)) {
						case 404:
							flag=1
							util.PushWeChat(phone+"登录过期", "")
						case 400:
							util.Putln("报名失败")
						case 200:
							flag=1
							util.Putln("报名成功")
						}
						if flag==1{
							break
						}
						history[phone] = append(history[phone], id)
					}
				}()
			}
		}
		if data["data"].(map[string]interface{})["pageEnd"].(bool) == true {
			page = 0
			time.Sleep(3 * time.Second)
			continue
		}
		time.Sleep(3 * time.Second)
	}
}

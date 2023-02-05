package snapchat

import (
	"app/src/request"
	"bytes"
	"encoding/json"
	"net/url"

	http "github.com/bogdanfinn/fhttp"
)

func CheckAccount(client request.RequestsClient, data map[string]string) (string, map[string]interface{}) {
	payload := url.Values{}
	for k, v := range data {
		payload.Set(k, v)
	}

	response, body := client.SendRequest("POST", "https://gcp.api.snapchat.com/scauth/login", bytes.NewReader([]byte(payload.Encode())), http.Header{
		"Host":            {"gcp.api.snapchat.com"},
		"Connection":      {"keep-alive"},
		"Accept":          {"application/json"},
		"User-Agent":      {"Snapchat/11.37.0.32 (ZTE BLADE V7 LITE; Android 6.0#20180821.210529#23; gzip) V/MUSHROOM"},
		"x-snapchat-att":  {"CgsgAAgKGAEV-LreBxKUAkoNgejW9naEgZ4Q3UYKsotk2-Dh-nthDwzENDDeniScsVxa6yJvcvJKTkCVBHZKY_taTpXL_tt3PAEOhoMjA6yH0Hw_v7uA29Xuv8Wday83VqVrgoerIvafjoiqILqJhVljXXpYj2W5JZpPoTD7KhVx4hBOsCbFTvSK3nPVd4tFYH-_muvgksgnfIDrB8PDVVdwHomyRdad9LupSLkk-6uBgD0wzCLIj1C5Zs0S-YWtrL-KGaf1KT4-iEnDS320UhwkON8t06nyX-Kscan__7jaG9oTVz6uAGL4ynV7rHTK1DtifJoJ3tQA1ZN4AXAhLOb6caL7fH0vSaBZgdHLAZ-HAWx-JQYxq6tmbH2aEp2MqSBX3Q=="},
		"Accept-Language": {"en-US;q=1, en;q=0.9"},
		"Content-Type":    {"application/x-www-form-urlencoded; charset=utf-8"},
		"Accept-Encoding": {"gzip, deflate, br"},
	})

	if response.StatusCode != 200 {
		// fmt.Println(body)
		// fmt.Println("=== DETECTED ===")
		return "bad", map[string]interface{}{}
	}

	r_payload := map[string]interface{}{}
	json.Unmarshal([]byte(body), &r_payload)

	if data["status"] != "" {
		return "bad", r_payload
	}
	return "good", r_payload
}

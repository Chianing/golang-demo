package scene

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
	"time"
)

func GetResponse() {
	url := "https://i.news.qq.com/trpc.qqnews_web.kv_srv.kv_srv_http_proxy/list?sub_srv_id=24hours&srv_id=pc&offset=60&limit=20&strategy=1&ext={%22pool%22:[%22top%22],%22is_filter%22:7,%22check_type%22:true}"
	respBody := getUrlRespBody(url)

	respBodyJson, err := simplejson.NewJson(respBody)
	if err != nil {
		fmt.Println("new respBodyJson err, err is: ", err)
	}

	dataArray, _ := respBodyJson.Get("data").Get("list").Array()
	for index, data := range dataArray {
		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println("marshal data to json err, err is: ", err)
			continue
		}
		tempJson, err := simplejson.NewJson(jsonData)
		fmt.Println("index is: ", index, ", title is: ", tempJson.Get("title").MustString())
	}

}

func getUrlRespBody(url string) []byte {
	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("get url err, err is: ", err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("io read err, err is: ", err)
	}
	return respBody
}

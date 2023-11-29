package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"gorm.io/gorm"
	"log"
	"mystudy/biz/models"
	"mystudy/global"
	"strings"
	"testing"
)

// id 最小值
func Test_First(t *testing.T) {
	var name models.Hotel
	// 按照id排序，查询第一个值
	//// SELECT * FROM users ORDER BY id LIMIT 1;
	global.DBEngine.First(&name)
	fmt.Print(name)
}

// 获取表中第一个数据
func Test_Take(t *testing.T) {
	var name models.Hotel
	// 获取表中的，第一个值
	global.DBEngine.Take(&name)
	// SELECT * FROM users LIMIT 1;
	fmt.Print(name)
}

// 获取按id排序，最后一个值
func Test_Last(t *testing.T) {
	var name models.Hotel
	// 获取表中的，第一个值
	result := global.DBEngine.Last(&name)
	// // SELECT * FROM users ORDER BY id DESC LIMIT 1;
	fmt.Println(name)
	fmt.Println(result.RowsAffected) // 返回发现到的记录数
	fmt.Println(result.Error)        // 返回错误或者nil

	errors.Is(result.Error, gorm.ErrRecordNotFound)

}

// es测试文件
func Test_Search(t *testing.T) {

	var (
		es = global.ESClient
		r  models.Hotel
	)

	//搜索已编入索引的文档
	// 3. Search for the indexed documents
	//构建请求体
	// Build the request body.
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	}
	// 转化为json数据
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	//执行搜索请求。
	// Perform the search request.
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("hotel"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()
	//
	//if res.IsError() {
	//	var e map[string]interface{}
	//	if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
	//		log.Fatalf("Error parsing the response body: %s", err)
	//	} else {
	//		//打印响应状态和错误信息。
	//		// Print the response status and error information.
	//		log.Fatalf("[%s] %s: %s",
	//			res.Status(),
	//			e["error"].(map[string]interface{})["type"],
	//			e["error"].(map[string]interface{})["reason"],
	//		)
	//	}
	//}
	//
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	//打印响应状态、结果数和请求持续时间。
	// Print the response status, number of results, and request duration.
	//log.Printf(
	//	"[%s] %d hits; took: %dms",
	//	res.Status(),
	//	int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
	//	int(r["took"].(float64)),
	//)
	////打印每次点击的 ID 和文档源。
	//// Print the ID and document source for each hit.
	//for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
	//	log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	//}

	log.Println(strings.Repeat("=", 37))

}
func Test_HSearch(t *testing.T) {
	// es初始化
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://192.168.218.134:9200",
		},
	}
	es, err := elasticsearch.NewTypedClient(cfg)
	log.Println(err)
	if err == nil {
		log.Println(elasticsearch.Version)
		log.Println(es.Info())
	} else {
		log.Println("Something wrong with connection to Elasticsearch")
	}
	res, _ := es.Get("hotel", "60359").Do(context.TODO())
	fmt.Println(res)

}

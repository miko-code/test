package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/olivere/elastic"
)

func (c *CellcomSlice) GetMatrix(add string) (*CellcomSlice, error) {
	//fix error
	proxyUrl, _ := url.Parse("http://proxyIp:proxyPort")

	tr := &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 30 * time.Second,
		Proxy:           http.ProxyURL(proxyUrl),
	}

	client := &http.Client{Transport: tr}
	//handel err
	resp, _ := client.Get(add)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal(body, c)
	if err != nil {
		fmt.Println("Can't deserislize", resp)
	}
	return c, nil
}

func InsertToEla(c *CellcomSlice) {

	///map cellcom slice to elastic type

	// Create a client and connect to http://192.168.2.10:9201
	client, err := elastic.NewClient(elastic.SetURL("http://192.168.2.10:9201"))
	if err != nil {
		// Handle error
	}

	index1Req := elastic.NewBulkIndexRequest().Index("twitter").Type("tweet").Doc(c)
	bulkRequest := client.Bulk()
	bulkRequest = bulkRequest.Add(index1Req)
	bulkResponse, err := bulkRequest.Do(context.Background())
	if err != nil {
		// ...
	}
}

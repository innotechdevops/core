package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/innotechdevops/core/redisx"
)

type Data struct {
	Text string `json:"text"`
	Int  int64  `json:"int"`
}

func toPtr[T any](v T) *T {
	return &v
}

func getRootKey() string {
	return "demo1:data"
}

func getKey(key string) string {
	return getRootKey() + ":" + key
}

func main() {
	ctx := context.Background()
	redisClient, err := redisx.NewClient(&redisx.RedisOptions{
		Addr:        "localhost:6379",
		Password:    "password",
		DB:          0,
		TraceEnable: false,
	})
	if err != nil {
		panic(err.Error())
	}
	defer redisClient.Close()

	data := &Data{
		Text: "First",
		Int:  1,
	}
	appCache := redisx.NewCache(redisClient.Client())
	err = appCache.Set(ctx, getKey("id1"), &redisx.Item{
		Value: data,
		TTL:   toPtr(time.Second * 10),
	})
	if err != nil {
		log.Fatalln(err.Error())
	}

	data1 := &Data{}
	err = appCache.Get(ctx, getKey("id1"), data1)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(data1.Text)
	fmt.Println(data1.Int)

	data2 := &Data{}
	err = appCache.GetCache(ctx, getKey("id12"), data2, toPtr(time.Second*10), func() (interface{}, error) {
		return userRepoGetData("id12"), nil
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(data2.Text)

	err = appCache.Delete(ctx, getKey("id1"))
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = appCache.DeleteByPrefix(ctx, getRootKey())
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func userRepoGetData(id string) *Data {
	return &Data{
		Text: "New Data:" + id,
	}
}

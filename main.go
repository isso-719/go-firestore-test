package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	opt := option.WithCredentialsFile("config/secret.json")

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		fmt.Println(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		fmt.Println(err)
	}

	// データの書き込み
	_, _, err = client.Collection("Scores").Add(ctx, map[string]interface{}{
		"Value": 200,
	})

	if err != nil {
		fmt.Println(err)
	}

	// データの書き込み (キー指定)
	_, err = client.Collection("Scores").Doc("1").Set(ctx, map[string]interface{}{
		"Value": 200,
	})

	if err != nil {
		fmt.Println(err)
	}

	// データの読み込み
	docs := client.Collection("Scores").Documents(ctx)
	for {
		doc, err := docs.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(doc.Data())
	}

	// データの読み込み (キー指定)
	doc, err := client.Collection("Scores").Doc("1").Get(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(doc.Data())

	// もし doc が 300 より小さい場合は、doc を更新
	if doc.Data()["Value"].(int64) < 300 {
		_, err = client.Collection("Scores").Doc("1").Set(ctx, map[string]interface{}{
			"Value": 300,
		})
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("更新しました")
	}

	// データの削除
	_, err = client.Collection("Scores").Doc("1").Delete(ctx)
	if err != nil {
		fmt.Println(err)
	}
}

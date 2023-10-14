package main

import (
	"context"
	"fmt"

	db "github.com/asragi/spajam2023/db"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
)

func Handler(ctx context.Context) (interface{}, error) {
	db, err := db.DB()
	if err != nil {
		return nil, err
	}
	// 取得するレコード一行のデータ形式を構造体で定義する
	type FoodSeason struct {
		Id     int `json:"id"`
		FoodId int `json:"food-id"`
	}
	// DBからレコードを抽出
	rows, err := db.Query("select id, food-id from food-seasons;")
	if err != nil {
		fmt.Println("Fail to query from db " + err.Error())
	}
	// データを構造体へ変換
	var foodSeasonData []FoodSeason
	for rows.Next() {
		var tmp FoodSeason
		err := rows.Scan(&tmp.Id, &tmp.FoodId)
		if err != nil {
			fmt.Println("Fail to scan records " + err.Error())
		}
		foodSeasonData = append(foodSeasonData, FoodSeason{
			Id:     tmp.Id,
			FoodId: tmp.FoodId,
		})
	}

	return foodSeasonData, nil
}

func main() {
	lambda.Start(Handler)
}

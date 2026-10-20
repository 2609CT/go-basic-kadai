package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Entry struct {
	Timestamp time.Time `json:"timestamp"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
}

func saveDiary(e []Entry) error {

	jsonData, err := json.Marshal(e)
	if err != nil {
		return fmt.Errorf("JSONマーシャルエラー：%w", err)
	}

	err = os.WriteFile("diary.json", jsonData, 0644)
	if err != nil {
		return fmt.Errorf("JSONファイルへの書き込みに失敗しました：%w", err)
	}

	return nil
}

func loadDiary(e []Entry, err error) ([]Entry, error) {

	data, err := os.ReadFile("diary.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []Entry{}, nil
		}
		return nil, fmt.Errorf("JSONファイルの読み込みに失敗しました：%w", err)
	}

	err = json.Unmarshal(data, &e)
	if err != nil {
		return nil, fmt.Errorf("JSONアンマーシャルエラー：%w", err)
	}

	return e, nil
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("使い方：go run main.go [add|list] ...")
		return
	}

	command := os.Args[1]

	var entrys []Entry
	var err error

	switch command {
	case "add":
		if len(os.Args) < 4 {
			fmt.Println("使い方：go run main.go add <タイトル> <本文>")
			return
		}
		entrys, err = loadDiary(entrys, err) //JSONファイル内の全件読み込み
		if err != nil {
			fmt.Println("add：loadDiaryエラー：", err)
			return
		}
		entrys = append(entrys, Entry{Timestamp: time.Now(), Title: os.Args[2], Content: os.Args[3]}) //コマンドライン引数追加
		err = saveDiary(entrys)                                                                       //JSONファイルに書き込み
		if err != nil {
			fmt.Println("add：saveDiaryエラー：", err)
			return
		}
		fmt.Println("保存しました")
	case "list":
		entrys, err = loadDiary(entrys, err)
		if err != nil {
			fmt.Println("list：loadDiaryエラー：", err)
			return
		}

		fmt.Println("--- 日記データ ---")

		for _, v := range entrys {
			fmt.Printf("\n日時：%s\nタイトル：%s\n本文：\n%s\n", v.Timestamp.Format("2006-01-02 15:04:05"), v.Title, v.Content)
		}
	default:
		fmt.Println("不明なコマンドです：", command)
	}
}

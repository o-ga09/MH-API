package util

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

func CsvToJson() {
	// CSVファイルのパス
	inputPath := "data/input/data.csv"
	outputPath := "data/output/data.json"

	currentDir, _ := os.Getwd()
	absolutecsvPath := filepath.Join(currentDir, inputPath)
	absolutejsonPath := filepath.Join(currentDir, outputPath)

	// CSVファイルを開く
	file, err := os.Open(absolutecsvPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// CSVデータをパースして読み込む
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// JSONに変換するデータのスライス
	data := make([]map[string]string, 0)

	// CSVの各行を処理してデータを作成
	for i, record := range records {
		if i == 0 {continue}
		row := make(map[string]string)
		for i, value := range record {
			// 各カラムのヘッダーをキーとして、値をマップに格納
			header := records[0][i]
			row[header] = value
		}
		data = append(data, row)
	}

	// JSONに変換
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	// JSONファイルを作成
	jsonFile, err := os.Create(absolutejsonPath)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	// JSONデータをファイルに書き込み
	_, err = jsonFile.Write(jsonData)
	if err != nil {
		log.Fatal(err)
	}
}
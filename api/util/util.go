package util

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type RequestJson struct {
	Req []Json `json:"req"`
}

type Json struct {
  	Name             string       `json:"name"`
    Desc             string       `json:"desc"`
    Location         string   `json:"location"`
    Specify          string    `json:"specify"`
    Weakness_attack  string `json:"weakness_attack"`
    Weakness_element string `json:"weakness_element"`
}

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
	data := RequestJson{}
	// CSVの各行を処理してデータを作成
	for i, record := range records {
		if i == 0 {continue}
		// 各カラムのヘッダーをキーとして、値をマップに格納
		row := Json{
			Name: record[0],
			Desc: record[1],
			Location: record[2],
			Specify: record[3],
			Weakness_attack: record[4],
			Weakness_element: record[5],
		}
		data.Req = append(data.Req, row)
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

func Mapping(str string) map[string]string {
	arr := strings.Split(str, " ")
	result :=  make(map[string]string)
	keys := [5]string{"頭部","前脚","胴体","後脚","尻尾"}

	for i, value := range arr {
		result[keys[i]] = value
	}

	return result
}

func Strtomap(mapped map[string]string) string {
	var str string
	for _, value := range mapped {
		str += value + " "
	}

	return str
}
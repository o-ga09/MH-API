package util

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type RequestJson struct {
	Req []Json `json:"req"`
}

type Json struct {
	MonsterId        string           `json:"monster_id,omitempty"`
	Name             string           `json:"name,omitempty"`
	Desc             string           `json:"desc,omitempty"`
	Location         string           `json:"location,omitempty"`
	Category         string           `json:"category,omitempty"`
	Title            string           `json:"title,omitempty"`
	Weakness_attack  Weakness_attack  `json:"weakness_attack,omitempty"`
	Weakness_element Weakness_element `json:"weakness_element,omitempty"`
}

type ResponseJson struct {
	Id               string           `json:"monster_id,omitempty"`
	Name             string           `json:"name,omitempty"`
	Desc             string           `json:"desc,omitempty"`
	Location         string           `json:"location,omitempty"`
	Category         string           `json:"category,omitempty"`
	Title            string           `json:"title,omitempty"`
	Weakness_attack  Weakness_attack  `json:"weakness_attack,omitempty"`
	Weakness_element Weakness_element `json:"weakness_element,omitempty"`
}

type Weakness_attack struct {
	FrontLegs AttackCatetgory `json:"front_legs,omitempty"`
	Tail      AttackCatetgory `json:"tail,omitempty"`
	HindLegs  AttackCatetgory `json:"hind_legs,omitempty"`
	Body      AttackCatetgory `json:"body,omitempty"`
	Head      AttackCatetgory `json:"head,omitempty"`
}

type Weakness_element struct {
	FrontLegs Elements `json:"front_legs,omitempty"`
	Tail      Elements `json:"tail,omitempty"`
	HindLegs  Elements `json:"hind_legs,omitempty"`
	Body      Elements `json:"body,omitempty"`
	Head      Elements `json:"head,omitempty"`
}

type AttackCatetgory struct {
	Slashing string `json:"slashing,omitempty"`
	Blow     string `json:"blow,omitempty"`
	Bullet   string `json:"bullet,omitempty"`
}

type Elements struct {
	Fire      string `json:"fire,omitempty"`
	Water     string `json:"water,omitempty"`
	Lightning string `json:"lightning,omitempty"`
	Ice       string `json:"ice,omitempty"`
	Dragon    string `json:"dragon,omitempty"`
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
		if i == 0 {
			continue
		}
		// 各カラムのヘッダーをキーとして、値をマップに格納
		row := Json{
			MonsterId: record[0],
			Name:      record[1],
			Desc:      record[2],
			Location:  record[3],
			Category:  record[4],
			Title:     record[5],
			Weakness_attack: Weakness_attack{
				FrontLegs: AttackCatetgory{
					Slashing: record[6],
					Blow:     record[7],
					Bullet:   record[8],
				},
				Tail: AttackCatetgory{
					Slashing: record[9],
					Blow:     record[10],
					Bullet:   record[11],
				},
				HindLegs: AttackCatetgory{
					Slashing: record[12],
					Blow:     record[13],
					Bullet:   record[14],
				},
				Body: AttackCatetgory{
					Slashing: record[15],
					Blow:     record[16],
					Bullet:   record[17],
				},
				Head: AttackCatetgory{
					Slashing: record[18],
					Blow:     record[19],
					Bullet:   record[20],
				},
			},
			Weakness_element: Weakness_element{
				FrontLegs: Elements{
					Fire:      record[21],
					Water:     record[22],
					Lightning: record[23],
					Ice:       record[24],
					Dragon:    record[25],
				},
				Tail: Elements{
					Fire:      record[26],
					Water:     record[27],
					Lightning: record[28],
					Ice:       record[29],
					Dragon:    record[30],
				},
				HindLegs: Elements{
					Fire:      record[31],
					Water:     record[32],
					Lightning: record[33],
					Ice:       record[34],
					Dragon:    record[35],
				},
				Body: Elements{
					Fire:      record[36],
					Water:     record[37],
					Lightning: record[38],
					Ice:       record[39],
					Dragon:    record[40],
				},
				Head: Elements{
					Fire:      record[41],
					Water:     record[42],
					Lightning: record[43],
					Ice:       record[44],
					Dragon:    record[45],
				},
			},
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
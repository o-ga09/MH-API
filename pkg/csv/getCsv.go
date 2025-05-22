package csv

import (
	"context"
	"encoding/csv"
	"io"

	"cloud.google.com/go/storage"
)

// GCSに保存されたcsvを取得する
func GetCSV(ctx context.Context, fileName string) (*[][]string, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	bucketName := "dev-mh-api-batch-data"
	objectName := fileName

	rc, err := client.Bucket(bucketName).Object(objectName).NewReader(ctx)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(rc)

	var data [][]string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		// 各レコードを結合
		data = append(data, record)
	}
	return &data, nil
}

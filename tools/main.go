package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"unicode"
)

type Monster struct {
	Index int
	Name  struct {
		Ja string
		En string
	}
	Tribe struct {
		Ja string
		En string
	}
	MH      int
	MHG     int
	MHP     int
	MH2     int
	P2nd    int
	P2G     int
	P3rd    int
	MH3G    int
	MH3     int
	MH4     int
	MH4G    int
	MHX     int
	MHXX    int
	MHW     int
	MHWI    int
	MHR     int
	MHRS    int
	Ranking Ranking
}

type Ranking struct {
	Rank      int
	Name      string
	Video_url string
}

const (
	MONSTER_DATA_SQL = "INSERT INTO monster (monster_id, name, created_at, updated_at) VALUES (%s, \"%s\", now(), now());"
	TRIBE_DATA_SQL   = "INSERT INTO tribe (tribe_id, name_ja, name_en, monster_id, created_at, updated_at) VALUES (%s, \"%s\", \"%s\", %s, now(), now());"
	PRODUCT_DATA_SQL = "INSERT INTO product (product_id, name, monster_id, created_at, updated_at) VALUES (%s, \"%s\", %s, now(), now());"
	RANKING_DATA_SQL = "INSERT INTO ranking (ranking, vote_year, monster_id, created_at, updated_at) VALUES (%s, \"%s\", %s, now(), now());"
)

func main() {
	// Example usage
	filePath := "MH_DATA_1.json"
	data, err := ReadJsonFile(filePath)
	if err != nil {
		panic(err)
	}

	filePath1 := "MH_DATA_4.json"
	data1, err := ReadJsonFile(filePath1)
	if err != nil {
		panic(err)
	}

	monsters, err := MapToStruct[Monster](data)
	if err != nil {
		panic(err)
	}

	ranking, err := MapToStruct[Ranking](data1)
	if err != nil {
		panic(err)
	}

	for i, r := range monsters {
		for _, v := range ranking {
			if r.Name.Ja == v.Name {
				monsters[i].Ranking.Name = v.Name
				monsters[i].Ranking.Rank = v.Rank
				monsters[i].Ranking.Video_url = v.Video_url
				break
			}
		}
	}

	_, _, productSQL, _ := CreateSQL(monsters)
	for _, sql := range productSQL {
		fmt.Println(sql)
	}
}

func CreateSQL(m []Monster) ([]string, []string, []string, []string) {
	var monsterSQL []string
	var tribeSQL []string
	var productSQL []string
	var rankingSQL []string
	for i, item := range m {
		index := strconv.Itoa(item.Index)
		no := strconv.Itoa(i + 1)
		rank := strconv.Itoa(item.Ranking.Rank)
		monsterSQL = append(monsterSQL, fmt.Sprintf(MONSTER_DATA_SQL, index, item.Name.Ja))
		tribeSQL = append(tribeSQL, fmt.Sprintf(TRIBE_DATA_SQL, no, item.Tribe.Ja, item.Tribe.En, index))
		if item.MH != 0 {
			productSQL = append(productSQL, fmt.Sprintf(PRODUCT_DATA_SQL, "1", "MH", index))
		}
		if item.MHG != 0 {
			productSQL = append(productSQL, fmt.Sprintf(PRODUCT_DATA_SQL, "2", "MHG", index))
		}
		if item.MHP != 0 {
			productSQL = append(productSQL, fmt.Sprintf(PRODUCT_DATA_SQL, "3", "MHP", index))
		}
		if item.MH2 != 0 {
			productSQL = append(productSQL, fmt.Sprintf(PRODUCT_DATA_SQL, "4", "MH2", index))
		}
		if item.P2nd != 0 {
			productSQL = append(productSQL, fmt.Sprintf(PRODUCT_DATA_SQL, "5", "P2nd", index))
		}
		if item.P2G != 0 {
			productSQL = append(productSQL, fmt.Sprintf(PRODUCT_DATA_SQL, "6", "P2G", index))
		}
		if item.P3rd != 0 {
			productSQL = append(productSQL, fmt.Sprintf(PRODUCT_DATA_SQL, "7", "P3rd", index))
		}
		if item.MH3G != 0 {
			productSQL = append(productSQL, fmt.Sprintf(PRODUCT_DATA_SQL, "8", "MH3G", index))
		}
		if item.MH3 != 0 {
			productSQL = append(productSQL, fmt.Sprintf(PRODUCT_DATA_SQL, "9", "MH3", index))
		}
		if item.MH4 != 0 {
			productSQL = append(productSQL, fmt.Sprintf(PRODUCT_DATA_SQL, "10", "MH4", index))
		}
		if item.MH4G != 0 {
			productSQL = append(productSQL, fmt.Sprintf(PRODUCT_DATA_SQL, "11", "MH4G", index))
		}
		if item.MHX != 0 {
			productSQL = append(productSQL, fmt.Sprintf(PRODUCT_DATA_SQL, "12", "MHX", index))
		}
		if item.MHXX != 0 {
			productSQL = append(productSQL, fmt.Sprintf(PRODUCT_DATA_SQL, "13", "MHXX", index))
		}
		if item.MHW != 0 {
			productSQL = append(productSQL, fmt.Sprintf(PRODUCT_DATA_SQL, "14", "MHW", index))
		}
		if item.MHWI != 0 {
			productSQL = append(productSQL, fmt.Sprintf(PRODUCT_DATA_SQL, "15", "MHWI", index))
		}
		if item.MHR != 0 {
			productSQL = append(productSQL, fmt.Sprintf(PRODUCT_DATA_SQL, "16", "MHR", index))
		}
		if item.MHRS != 0 {
			productSQL = append(productSQL, fmt.Sprintf(PRODUCT_DATA_SQL, "17", "MHRS", index))
		}
		rankingSQL = append(rankingSQL, fmt.Sprintf(RANKING_DATA_SQL, rank, "2024", index))
	}
	return monsterSQL, tribeSQL, productSQL, rankingSQL
}

func ReadJsonFile(filePath string) ([]map[string]interface{}, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var data []map[string]interface{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func MapToStruct[V comparable](m []map[string]interface{}) ([]V, error) {
	var results []V

	for _, item := range m {
		result := reflect.New(reflect.TypeOf(*new(V))).Elem()

		for key, value := range item {
			key = UppercaseFirst(key)

			field := result.FieldByName(key)
			if field.IsValid() && field.CanSet() {
				val := reflect.ValueOf(value)
				if field.Type().Kind() == reflect.Struct {
					// 入れ子の構造体の場合
					nestedData, ok := value.(map[string]interface{})
					if !ok {
						return nil, fmt.Errorf("expected map[string]interface{} for nested struct %s, got %T", key, value)
					}

					nestedValue := reflect.New(field.Type()).Elem()
					if err := mapToNestedStruct(nestedData, nestedValue); err != nil {
						return nil, fmt.Errorf("error mapping to nested struct %s: %w", key, err)
					}
					field.Set(nestedValue)
				} else {
					if field.Type().Kind() == val.Type().Kind() {
						field.Set(val)
					} else {
						// 型変換を試みる
						convertedValue, err := convertType(value, field.Type())
						if err != nil {
							return nil, fmt.Errorf("type conversion error for key %s: %w", key, err)
						}
						field.Set(reflect.ValueOf(convertedValue))
					}
				}
			}
		}

		results = append(results, result.Interface().(V))
	}

	return results, nil
}

func convertType(value interface{}, targetType reflect.Type) (interface{}, error) {
	switch targetType.Kind() {
	case reflect.String:
		return fmt.Sprintf("%v", value), nil
	case reflect.Int:
		// float64 から int への変換を試みる
		if floatValue, ok := value.(float64); ok {
			return int(floatValue), nil
		}
		return 0, fmt.Errorf("cannot convert %T to int", value)
	case reflect.Float64:
		// 他の数値型から float64 への変換を試みる
		if intValue, ok := value.(int); ok {
			return float64(intValue), nil
		}
		return 0.0, fmt.Errorf("cannot convert %T to float64", value)
	case reflect.Bool:
		return fmt.Sprintf("%v", value), nil
	default:
		return nil, fmt.Errorf("unsupported type %v", targetType)
	}
}

func UppercaseFirst(s string) string {
	if len(s) == 0 {
		return ""
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func mapToNestedStruct(data map[string]interface{}, result reflect.Value) error {
	for key, value := range data {
		key = UppercaseFirst(key)
		field := result.FieldByName(key)

		if field.IsValid() && field.CanSet() {
			if field.Type().Kind() == reflect.Struct {
				// さらにネストされた構造体の場合、再帰的に処理
				nestedData, ok := value.(map[string]interface{})
				if !ok {
					return fmt.Errorf("expected map[string]interface{} for nested struct %s, got %T", key, value)
				}
				nestedValue := reflect.New(field.Type()).Elem()
				if err := mapToNestedStruct(nestedData, nestedValue); err != nil {
					return fmt.Errorf("error mapping to nested struct %s: %w", key, err)
				}
				field.Set(nestedValue)
			} else {
				// 通常のフィールドの場合
				val := reflect.ValueOf(value)
				if field.Type().Kind() == val.Type().Kind() {
					field.Set(val)
				} else {
					// 型変換を試みる
					convertedValue, err := convertType(value, field.Type())
					if err != nil {
						return fmt.Errorf("type conversion error for key %s: %w", key, err)
					}
					field.Set(reflect.ValueOf(convertedValue))
				}
			}
		}
	}
	return nil
}

## ロール

- あなたは、メガベンチャー企業で働く CTO クラスのエンジニアです

## ルール

- 日本語で回答すること
- このプロジェクトの GitHub 上のリポジトリは、`o-ga09/MH-API`です。
- コミットする前にテスト(`make test`)を実行する
- コードの提案はもちろん、レビューについても CTO の観点からレビューをお願いします

## 使用技術

- Go
- Gin
- Gorm

## プロジェクト構成

- [Go Standard layout](https://github.com/golang-standards/project-layout)に則った構成
- クリーンアーキテクチャっぽい構成

## テストの実装について

- 外部制約を持つテーブルのテストデータの作成は、関連テーブルのデータも作成すること
- `pkg`配下のディレクトリに共通処理があるので適宜使用すること
- テーブルドリブンテストで実装すること
- `controller`パッケージのテストは、ゴールデンテストを用いること
- テストケースは正常ケース、異常ケースは必ず含めること
- `controller`パッケージのテストに限り、400エラーなどの準正常ケースも実装すること
- モックには`moq`ライブラリを使用して自動生成すること
  - `make generate`で自動生成可能

## ディレクトリ構成

```bash
.
├── CLAUDE.md
├── CODE_OF_CONDUCT.md
├── CODE_OF_CONDUCT_JA.md
├── CONTRIBUTING.md
├── CONTRIBUTING_JA.md
├── Dockerfile
├── LICENSE
├── README.md
├── README_EN.md
├── SECURITY.md
├── cmd
│   ├── api
│   │   └── main.go
│   └── migration
│       └── main.go
├── compose.yml
├── contact.txt
├── cover.out
├── coverage.out
├── data
│   ├── element
│   │   ├── 7_1_11_1.jpg
│   │   ├── 7_1_12_1.jpg
│   │   ├── 7_1_13_1.jpg
│   │   ├── 7_1_14_1.jpg
│   │   ├── 7_1_15_1.jpg
│   │   ├── 7_2_11_1.jpg
│   │   ├── 7_2_12_1.jpg
│   │   ├── 7_2_2_1.jpg
│   │   ├── 7_2_3_1.jpg
│   │   ├── 7_2_4_1.jpg
│   │   └── 7_2_5_1.jpg
│   └── monster
│       ├── 1.png
│       ├── 10.png
│       ├── 100.png
│       ├── 101.png
│       ├── 102.png
│       ├── 103.png
│       ├── 104.png
│       ├── 105.png
│       ├── 106.png
│       ├── 107.png
│       ├── 108.png
│       ├── 109.png
│       ├── 11.png
│       ├── 110.png
│       ├── 111.png
│       ├── 112.png
│       ├── 113.png
│       ├── 114.png
│       ├── 115.png
│       ├── 116.png
│       ├── 117.png
│       ├── 118.png
│       ├── 119.png
│       ├── 12.png
│       ├── 120.png
│       ├── 121.png
│       ├── 122.png
│       ├── 123.png
│       ├── 124.png
│       ├── 125.png
│       ├── 126.png
│       ├── 127.png
│       ├── 128.png
│       ├── 129.png
│       ├── 13.png
│       ├── 130.png
│       ├── 131.png
│       ├── 132.png
│       ├── 133.png
│       ├── 134.png
│       ├── 135.png
│       ├── 136.png
│       ├── 137.png
│       ├── 138.png
│       ├── 139.png
│       ├── 14.png
│       ├── 140.png
│       ├── 141.png
│       ├── 142.png
│       ├── 143.png
│       ├── 144.png
│       ├── 145.png
│       ├── 146.png
│       ├── 147.png
│       ├── 148.png
│       ├── 149.png
│       ├── 15.png
│       ├── 150.png
│       ├── 151.png
│       ├── 152.png
│       ├── 153.png
│       ├── 154.png
│       ├── 155.png
│       ├── 156.png
│       ├── 157.png
│       ├── 158.png
│       ├── 159.png
│       ├── 16.png
│       ├── 160.png
│       ├── 161.png
│       ├── 162.png
│       ├── 163.png
│       ├── 164.png
│       ├── 165.png
│       ├── 166.png
│       ├── 167.png
│       ├── 168.png
│       ├── 169.png
│       ├── 17.png
│       ├── 170.png
│       ├── 171.png
│       ├── 172.png
│       ├── 173.png
│       ├── 174.png
│       ├── 175.png
│       ├── 176.png
│       ├── 177.png
│       ├── 178.png
│       ├── 179.png
│       ├── 18.png
│       ├── 180.png
│       ├── 181.png
│       ├── 182.png
│       ├── 183.png
│       ├── 184.png
│       ├── 185.png
│       ├── 186.png
│       ├── 187.png
│       ├── 188.png
│       ├── 189.png
│       ├── 19.png
│       ├── 190.png
│       ├── 191.png
│       ├── 192.png
│       ├── 193.png
│       ├── 194.png
│       ├── 195.png
│       ├── 196.png
│       ├── 197.png
│       ├── 198.png
│       ├── 199.png
│       ├── 2.png
│       ├── 20.png
│       ├── 200.png
│       ├── 201.png
│       ├── 202.png
│       ├── 203.png
│       ├── 204.png
│       ├── 205.png
│       ├── 206.png
│       ├── 207.png
│       ├── 208.png
│       ├── 209.png
│       ├── 21.png
│       ├── 210.png
│       ├── 211.png
│       ├── 212.png
│       ├── 213.png
│       ├── 214.png
│       ├── 215.png
│       ├── 216.png
│       ├── 217.png
│       ├── 218.png
│       ├── 219.png
│       ├── 22.png
│       ├── 220.png
│       ├── 221.png
│       ├── 222.png
│       ├── 223.png
│       ├── 224.png
│       ├── 225.png
│       ├── 226.png
│       ├── 227.png
│       ├── 228.png
│       ├── 229.png
│       ├── 23.png
│       ├── 24.png
│       ├── 25.png
│       ├── 26.png
│       ├── 27.png
│       ├── 28.png
│       ├── 29.png
│       ├── 3.png
│       ├── 30.png
│       ├── 31.png
│       ├── 32.png
│       ├── 33.png
│       ├── 34.png
│       ├── 35.png
│       ├── 36.png
│       ├── 37.png
│       ├── 38.png
│       ├── 39.png
│       ├── 4.png
│       ├── 40.png
│       ├── 41.png
│       ├── 42.png
│       ├── 43.png
│       ├── 44.png
│       ├── 45.png
│       ├── 46.png
│       ├── 47.png
│       ├── 48.png
│       ├── 49.png
│       ├── 5.png
│       ├── 50.png
│       ├── 51.png
│       ├── 52.png
│       ├── 53.png
│       ├── 54.png
│       ├── 55.png
│       ├── 56.png
│       ├── 57.png
│       ├── 58.png
│       ├── 59.png
│       ├── 6.png
│       ├── 60.png
│       ├── 61.png
│       ├── 62.png
│       ├── 63.png
│       ├── 64.png
│       ├── 65.png
│       ├── 66.png
│       ├── 67.png
│       ├── 68.png
│       ├── 69.png
│       ├── 7.png
│       ├── 70.png
│       ├── 71.png
│       ├── 72.png
│       ├── 73.png
│       ├── 74.png
│       ├── 75.png
│       ├── 76.png
│       ├── 77.png
│       ├── 78.png
│       ├── 79.png
│       ├── 8.png
│       ├── 80.png
│       ├── 81.png
│       ├── 82.png
│       ├── 83.png
│       ├── 84.png
│       ├── 85.png
│       ├── 86.png
│       ├── 87.png
│       ├── 88.png
│       ├── 89.png
│       ├── 9.png
│       ├── 90.png
│       ├── 91.png
│       ├── 92.png
│       ├── 93.png
│       ├── 94.png
│       ├── 95.png
│       ├── 96.png
│       ├── 97.png
│       ├── 98.png
│       └── 99.png
├── db
│   ├── migrations
│   │   ├── 20250426164157_initial-scheme.sql
│   │   ├── 20250521002025_add_element_to_monsters.sql
│   │   ├── 20250524222649_modify-item.sql
│   │   └── 20250525151111_add-weapon-skill.sql
│   ├── mysql
│   │   ├── conf.d
│   │   │   └── my.cnf
│   │   └── sql
│   │       └── init.sql
│   └── seed
│       ├── 00_truncate.sql
│       └── 01_seed.sql
├── doc
│   ├── ER
│   │   ├── er.drawio
│   │   └── er.png
│   ├── architecture
│   │   ├── MH-API_アーキテクチャ図.drawio
│   │   └── MH-API_アーキテクチャ図.png
│   └── openapi
│       ├── APIGateway.md
│       ├── apigateway.yml
│       ├── docs.go
│       ├── openapi.json
│       ├── redoc-static.html
│       ├── schema.yml
│       ├── swagger.json
│       └── swagger.yaml
├── go.mod
├── go.sum
├── internal
│   ├── DI
│   │   ├── health.go
│   │   ├── items.go
│   │   ├── monsters.go
│   │   └── weapons.go
│   ├── controller
│   │   ├── item
│   │   │   ├── handler.go
│   │   │   ├── handler_test.go
│   │   │   ├── request.go
│   │   │   ├── response.go
│   │   │   └── testdata
│   │   │       └── items
│   │   │           ├── get_item_bad_request.json
│   │   │           ├── get_item_by_monster_bad_request.json
│   │   │           ├── get_item_by_monster_empty.json
│   │   │           ├── get_item_by_monster_error.json
│   │   │           ├── get_item_by_monster_not_found.json
│   │   │           ├── get_item_by_monster_not_implemented.json
│   │   │           ├── get_item_by_monster_success.json
│   │   │           ├── get_item_error.json
│   │   │           ├── get_item_not_found.json
│   │   │           ├── get_item_not_implemented.json
│   │   │           ├── get_item_success.json
│   │   │           ├── get_items_error.json
│   │   │           └── get_items_success.json
│   │   ├── monster
│   │   │   ├── handler.go
│   │   │   ├── handler_test.go
│   │   │   ├── request.go
│   │   │   ├── response.go
│   │   │   └── testdata
│   │   │       └── monster
│   │   │           ├── monster_get_all_200.json
│   │   │           ├── monster_get_all_400.json
│   │   │           ├── monster_get_all_404.json
│   │   │           ├── monster_get_all_500.json
│   │   │           ├── monster_get_all_empty.json
│   │   │           ├── monster_get_by_id_200.json
│   │   │           ├── monster_get_by_id_400.json
│   │   │           ├── monster_get_by_id_404.json
│   │   │           └── monster_get_by_id_500.json
│   │   ├── system.go
│   │   └── weapon
│   │       ├── handler.go
│   │       ├── handler_test.go
│   │       ├── request.go
│   │       ├── response.go
│   │       └── testdata
│   │           └── weapon
│   │               ├── weapon_search_200.json
│   │               ├── weapon_search_400.json
│   │               └── weapon_search_500.json
│   ├── database
│   │   └── mysql
│   │       ├── db_connect.go
│   │       ├── health.go
│   │       ├── itemQueryService.go
│   │       ├── itemQueryService_test.go
│   │       ├── monsterQueryService.go
│   │       ├── monsterQueryService_test.go
│   │       ├── monsters.go
│   │       ├── mysql_test.go
│   │       ├── schemas.go
│   │       ├── sentry.go
│   │       ├── testHelper.go
│   │       ├── weaponQueryService.go
│   │       └── weaponQueryService_test.go
│   ├── domain
│   │   ├── fields
│   │   │   ├── field.go
│   │   │   ├── repository.go
│   │   │   ├── repository_mock.go
│   │   │   └── type.go
│   │   ├── health
│   │   │   └── repository.go
│   │   ├── items
│   │   │   ├── item.go
│   │   │   ├── repository.go
│   │   │   ├── repository_mock.go
│   │   │   └── type.go
│   │   ├── monsters
│   │   │   ├── monster.go
│   │   │   ├── repository.go
│   │   │   ├── repository_mock.go
│   │   │   └── types.go
│   │   ├── music
│   │   │   ├── music.go
│   │   │   ├── repository.go
│   │   │   ├── repository_mock.go
│   │   │   └── type.go
│   │   ├── part
│   │   │   ├── field.go
│   │   │   ├── repository.go
│   │   │   ├── repository_mock.go
│   │   │   └── type.go
│   │   ├── products
│   │   │   ├── product.go
│   │   │   ├── repository.go
│   │   │   ├── repository_mock.go
│   │   │   └── type.go
│   │   ├── ranking
│   │   │   ├── ranking.go
│   │   │   ├── repository.go
│   │   │   ├── repository_mock.go
│   │   │   └── type.go
│   │   ├── tribes
│   │   │   ├── repository.go
│   │   │   ├── repository_mock.go
│   │   │   ├── tribe.go
│   │   │   └── type.go
│   │   ├── weakness
│   │   │   ├── repository.go
│   │   │   ├── type.go
│   │   │   └── weakness.go
│   │   └── weapons
│   │       ├── repository.go
│   │       ├── repository_mock.go
│   │       ├── type.go
│   │       └── weapon.go
│   ├── presenter
│   │   ├── middleware
│   │   │   ├── context.go
│   │   │   ├── cors.go
│   │   │   ├── error.go
│   │   │   ├── httplogger.go
│   │   │   ├── logger.go
│   │   │   └── sentry.go
│   │   └── server.go
│   └── service
│       ├── health
│       │   └── health.go
│       ├── items
│       │   ├── items.go
│       │   ├── items_mock.go
│       │   └── items_test.go
│       ├── monsters
│       │   ├── dto.go
│       │   ├── monsters.go
│       │   ├── monsters_test.go
│       │   ├── monsterservice_mock.go
│       │   ├── queryService.go
│       │   └── queryService_mock.go
│       └── weapons
│           ├── dto.go
│           ├── mock_weapon_query_service_test.go
│           ├── weapons.go
│           └── weapons_test.go
├── makefile
├── pkg
│   ├── config
│   │   └── config.go
│   ├── constant
│   │   └── constant.go
│   ├── csv
│   │   └── getCsv.go
│   ├── ptr
│   │   └── ptr.go
│   ├── testutil
│   │   └── golden.go
│   ├── uuid
│   │   └── uuid.go
│   └── validator
│       └── validator.go
├── scenario
│   ├── README.md
│   ├── e2e.yml
│   ├── httptest
│   │   ├── DELETE.http
│   │   ├── GET.http
│   │   ├── POST.http
│   │   └── PUT.http
│   ├── junit.xml
│   ├── report.json
│   └── scenarigo.yaml
├── terraform
│   ├── local.tf
│   ├── main.tf
│   └── stg.tf
└── tools
    ├── go.mod
    ├── go.sum
    └── main.go
```


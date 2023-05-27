package store

import (
	"context"
	"mh-api/api/entity"
	"mh-api/api/interface/repository"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func TestNewMonsterRepository(t *testing.T) {
	t.Skip()
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want repository.IMonsterRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMonsterRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMonsterRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMonsterRepository_SelectMonsterAll(t *testing.T) {
	mockData1 := entity.Monster{
		Id: 1,
		Name: "ジンオウガ",
		Desc: "ジンオウガかっこいい",
		Location: "大社跡",
		Specify: "牙竜種",
		Weakness_attack: "10 10 10 10 10",
		Weakness_element: "10 10 10 10 10",
	}
	mockData2 := entity.Monster{
		Id: 2,
		Name: "タマミツネ",
		Desc: "男の娘",
		Location: "大社跡",
		Specify: "海竜種",
		Weakness_attack: "10 10 10 10 10",
		Weakness_element: "10 10 10 10 10",
	}
	//db mock generate
	db ,mock ,_ := NewDbMock()

	type fields struct {
		Db *gorm.DB
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Monsters
		wantErr bool
	}{
		{name: "ok",fields: fields{Db: db},args: args{ctx: context.Background()},want: entity.Monsters{mockData1,mockData2},wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rows := sqlmock.NewRows([]string{"id","name","desc","location","specify","weakness_attack","weakness_element"}).
			AddRow(mockData1.Id,mockData1.Name,mockData1.Desc,mockData1.Location,mockData1.Specify,mockData1.Weakness_attack,mockData1.Weakness_element).
			AddRow(mockData2.Id,mockData2.Name,mockData2.Desc,mockData2.Location,mockData2.Specify,mockData2.Weakness_attack,mockData2.Weakness_element)

			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `monster`")).WillReturnRows(rows)

			m := &MonsterRepository{
				Db: tt.fields.Db,
			}
			got, err := m.SelectMonsterAll(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("MonsterRepository.SelectMonsterAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MonsterRepository.SelectMonsterAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMonsterRepository_SelectMonsterById(t *testing.T) {
	mockData1 := entity.Monster{
		Id: 1,
		Name: "ジンオウガ",
		Desc: "ジンオウガかっこいい",
		Location: "大社跡",
		Specify: "牙竜種",
		Weakness_attack: "10 10 10 10 10",
		Weakness_element: "10 10 10 10 10",
	}

	//db mock generate
	db ,mock ,_ := NewDbMock()
	type fields struct {
		Db *gorm.DB
	}
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Monster
		wantErr bool
	}{
		{name: "ok",fields: fields{Db: db},args: args{ctx: context.Background(),id: 1},want: mockData1,wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rows := sqlmock.NewRows([]string{"id","name","desc","location","specify","weakness_attack","weakness_element"}).
			AddRow(mockData1.Id,mockData1.Name,mockData1.Desc,mockData1.Location,mockData1.Specify,mockData1.Weakness_attack,mockData1.Weakness_element)

			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `monster` WHERE `monster`.`id` = ?")).
			WithArgs(1).
			WillReturnRows(rows)
			
			m := &MonsterRepository{
				Db: tt.fields.Db,
			}
			got, err := m.SelectMonsterById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("MonsterRepository.SelectMonsterById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MonsterRepository.SelectMonsterById() = %v, want %v", got, tt.want)
			}
		})
	}
}

//mock code
func NewDbMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, mock, err
	}
	
	gormDB, err := gorm.Open(mysql.Dialector{
		Config: &mysql.Config{
			DriverName: "mysql",
			Conn: db,
			SkipInitializeWithVersion: true,
		}}, &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})
	if err != nil {
		return gormDB, mock, err
	}
	return gormDB, mock, err
}

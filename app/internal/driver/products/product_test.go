package products

import (
	"context"
	Products "mh-api/app/internal/domain/products"
	"mh-api/app/internal/driver/mysql"
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestNewMonsterRepository(t *testing.T) {
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())
	conn := mysql.New(context.Background())
	type args struct {
		conn *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want *productRepository
	}{
		// test 1
		{name: "TestNewMonsterRepository", args: args{conn: conn}, want: &productRepository{conn: conn}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMonsterRepository(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMonsterRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_productRepository_Save(t *testing.T) {
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())
	conn := mysql.New(context.Background())
	type fields struct {
		conn *gorm.DB
	}
	type args struct {
		ctx context.Context
		p   Products.Product
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// test 1
		{name: "Save product successfully", fields: fields{conn: conn}, args: args{ctx: context.Background(), p: *Products.NewProduct("0000000001", "test", "test", "test")}, wantErr: false},
		// test 2
		{name: "Save product with error", fields: fields{conn: conn}, args: args{ctx: context.Background(), p: *Products.NewProduct("@$%&^#%$&&*%*&)(*)()", "", "", "")}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &productRepository{
				conn: tt.fields.conn,
			}
			if err := r.Save(tt.args.ctx, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("productRepository.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_productRepository_Remove(t *testing.T) {
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())
	conn := mysql.New(context.Background())
	type fields struct {
		conn *gorm.DB
	}
	type args struct {
		ctx       context.Context
		productId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// test 1
		{name: "Remove product successfully", fields: fields{conn: conn}, args: args{ctx: context.Background(), productId: "0000000001"}, wantErr: false},
		// test 2
		{name: "Remove product with error", fields: fields{conn: conn}, args: args{ctx: context.Background(), productId: ""}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &productRepository{
				conn: tt.fields.conn,
			}
			if err := r.Remove(tt.args.ctx, tt.args.productId); (err != nil) != tt.wantErr {
				t.Errorf("productRepository.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

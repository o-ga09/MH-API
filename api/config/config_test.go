package config

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		want    *Config
		wantErr bool
	}{
		{name: "ケース１",want: &Config{Env: "dev",Port: "8080",DBHost: "dbsrv01",DBUser: "mh-api",DBPassword: "password", DBName: "mh-api"},wantErr: false},
	}
	for _, tt := range tests {
		t.Setenv("ENV",tt.want.Env)
		t.Setenv("PORT",tt.want.Port)
		t.Setenv("DB_HOST",tt.want.DBHost)
		t.Setenv("DB_USER",tt.want.DBUser)
		t.Setenv("DB_PASSWORD",tt.want.DBPassword)
		t.Setenv("DB_NAME",tt.want.DBName)

		t.Run(tt.name, func(t *testing.T) {
			got, err := New()
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

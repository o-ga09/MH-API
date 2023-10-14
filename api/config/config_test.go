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
		{name: "ケース１",want: &Config{Env: "dev",Port: "8080",DBHost: "dbsrv01",DBUser: "mh-api",DBPassword: "password", DBName: "mh-api",USER: "admin",PASSWORD: "password",SECRET_KEY: "secret_key",TOKEN_LIFETIME: "3600"},wantErr: false},
	}
	for _, tt := range tests {
		t.Setenv("ENV",tt.want.Env)
		t.Setenv("PORT",tt.want.Port)
		t.Setenv("DB_HOST",tt.want.DBHost)
		t.Setenv("DB_USER",tt.want.DBUser)
		t.Setenv("DB_PASSWORD",tt.want.DBPassword)
		t.Setenv("DB_NAME",tt.want.DBName)
		t.Setenv("USER",tt.want.USER)
		t.Setenv("PASSWORD",tt.want.PASSWORD)
		t.Setenv("SECRET_KEY",tt.want.SECRET_KEY)
		t.Setenv("TOKEN_LIFETIME",tt.want.TOKEN_LIFETIME)

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

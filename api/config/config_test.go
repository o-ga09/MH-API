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
		{name: "正常系", want: &Config{Env: "dev", Port: "8080", Database_url: "user:password@tcp(localhost)/test?charset=utf8&parseTime=True&loc=Local&tls=true", ProjectID: "00000",ALLOW_URL: "*"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Setenv("ENV", tt.want.Env)
		t.Setenv("PORT", tt.want.Port)
		t.Setenv("DATABASE_URL", tt.want.Database_url)
		t.Setenv("PROJECTID", tt.want.ProjectID)
		t.Setenv("ALLOW_URL",tt.want.ALLOW_URL)

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

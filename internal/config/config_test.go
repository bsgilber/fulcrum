package config

import (
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
	type args struct {
		configFilePath string
		env            string
	}
	tests := []struct {
		name string
		args args
		want *Config
	}{
		{
			name: "test local.yaml",
			args: args{
				configFilePath: "../../test",
				env:            "local",
			},
			want: &Config{
				Server: ServerConfig{
					Port: 8080,
				},
				Logging: LoggerConfig{
					Level: "debug",
				},
				Redis: RedisConnectionConfig{
					Host: "redis://localhost",
					Port: "6379",
				},
				Psql: PsqlConnectionConfig{
					Uri: "postgres://postgres:password@psql/postgres",
				},
				Project: ProjectConfig{
					Id:     "fulcrum",
					Region: "us-central1",
				},
				Env: "local",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfig(tt.args.configFilePath, tt.args.env); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

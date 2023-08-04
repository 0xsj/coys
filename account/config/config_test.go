package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name        string
		configFile  string
		expectedErr bool
		expectedCfg Config
	}{
		{
			name:        "config-test",
			configFile:  "config.yaml",
			expectedErr: false,
			expectedCfg: Config{
				ServiceName:     "MyService",
				ServerAddress:   "localhost:8080",
				MongoHost:       "mongodb://localhost",
				MongoPort:       "27017",
				DatabaseName:    "mydb",
				CollectionItems: "items",
				ApiKey:          "myapikey",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg, err := LoadConfig(tt.configFile)

			if tt.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedCfg, cfg)
			}
		})
	}
}

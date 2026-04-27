package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/infrastructure/repository/sqlite"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/interface/colors"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/interface/writer"
	"github.com/stretchr/testify/assert"
)

func TestParseConfig(t *testing.T) {

	testCases := []struct {
		name           string
		input          string
		noFile         bool
		expectedResult *Config
		expectedError  bool
	}{
		// first: empty file. Order does make scence
		{
			name:  "OK - not full config (defaults)",
			input: "",
			expectedResult: &Config{
				sqlite.Config{
					StoragePath: "./storage/task_storage.db",
				},
				writer.Config{
					MaxColumnWidth: 40,
					ExtraColumns:   []string{},
				},
				colors.Config{
					ColoredOutput: false,
				},
			},
			expectedError: false,
		},
		// second: not empty, because cleanenv caches env variables and it
		// requires additional logic to clear them
		{
			name: "OK - full config",
			input: `
SQLITE_STORAGE_PATH=./storage/storage.db
FORMAT_MAX_WIDTH=50
FORMAT_EXTRA_COLUMNS=created,description
COLORED_OUTPUT=1
			`,
			expectedResult: &Config{
				sqlite.Config{
					StoragePath: "./storage/storage.db",
				},
				writer.Config{
					MaxColumnWidth: 50,
					ExtraColumns:   []string{writer.ValidCreatedName, writer.ValidDescriptionName},
				},
				colors.Config{
					ColoredOutput: true,
				},
			},
			expectedError: false,
		},
		{
			name: "Error - invalid values in config",
			input: `
SQLITE_STORAGE_PATH=./storage/storage.db
FORMAT_MAX_WIDTH=invalid
FORMAT_EXTRA_COLUMNS=zebra,description
COLORED_OUTPUT=invalid
			`,
			expectedError: true,
		},
		{
			name:          "Error - no config file",
			noFile:        true,
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			tempDir := t.TempDir()
			configPath := filepath.Join(tempDir, ".env")

			if !tc.noFile {
				err := os.WriteFile(configPath, []byte(tc.input), 0644)
				assert.NoError(t, err)
			}

			cfg, err := ParseConfig(configPath)

			if tc.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedResult, cfg)
			}
		})
	}
}

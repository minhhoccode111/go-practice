package config

import (
	"os"
	"reflect" // Needed for deep comparison of slices
	"testing"
	"time"
)

func TestGetEnvString(t *testing.T) {
	tests := []struct {
		name         string
		key          string
		defaultValue string
		envValue     string
		expected     string
	}{
		{
			name:         "Env var exists",
			key:          "TEST_STRING_ENV",
			defaultValue: "default",
			envValue:     "env_value",
			expected:     "env_value",
		},
		{
			name:         "Env var does not exist",
			key:          "NON_EXISTENT_STRING_ENV",
			defaultValue: "default",
			envValue:     "", // Not set
			expected:     "default",
		},
		{
			name:         "Env var is empty",
			key:          "EMPTY_STRING_ENV",
			defaultValue: "default",
			envValue:     "", // Set to empty string
			expected:     "default",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envValue != "" {
				os.Setenv(tt.key, tt.envValue)
			} else {
				os.Unsetenv(tt.key)
			}
			defer os.Unsetenv(tt.key) // Clean up

			got := getEnvString(tt.key, tt.defaultValue)
			if got != tt.expected {
				t.Errorf(
					"getEnvString(%q, %q) = %q, expected %q",
					tt.key,
					tt.defaultValue,
					got,
					tt.expected,
				)
			}
		})
	}
}

func TestGetEnvInt(t *testing.T) {
	tests := []struct {
		name         string
		key          string
		defaultValue int
		envValue     string
		expected     int
	}{
		{
			name:         "Env var exists and is valid",
			key:          "TEST_INT_ENV",
			defaultValue: 100,
			envValue:     "200",
			expected:     200,
		},
		{
			name:         "Env var does not exist",
			key:          "NON_EXISTENT_INT_ENV",
			defaultValue: 100,
			envValue:     "",
			expected:     100,
		},
		{
			name:         "Env var is empty",
			key:          "EMPTY_INT_ENV",
			defaultValue: 100,
			envValue:     "",
			expected:     100,
		},
		{
			name:         "Env var is invalid",
			key:          "INVALID_INT_ENV",
			defaultValue: 100,
			envValue:     "abc",
			expected:     100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envValue != "" {
				os.Setenv(tt.key, tt.envValue)
			} else {
				os.Unsetenv(tt.key)
			}
			defer os.Unsetenv(tt.key) // Clean up

			got := getEnvInt(tt.key, tt.defaultValue)
			if got != tt.expected {
				t.Errorf(
					"getEnvInt(%q, %d) = %d, expected %d",
					tt.key,
					tt.defaultValue,
					got,
					tt.expected,
				)
			}
		})
	}
}

func TestGetEnvDuration(t *testing.T) {
	tests := []struct {
		name         string
		key          string
		defaultValue time.Duration
		envValue     string
		expected     time.Duration
	}{
		{
			name:         "Env var exists and is valid",
			key:          "TEST_DURATION_ENV",
			defaultValue: 10 * time.Second,
			envValue:     "20s",
			expected:     20 * time.Second,
		},
		{
			name:         "Env var does not exist",
			key:          "NON_EXISTENT_DURATION_ENV",
			defaultValue: 10 * time.Second,
			envValue:     "",
			expected:     10 * time.Second,
		},
		{
			name:         "Env var is empty",
			key:          "EMPTY_DURATION_ENV",
			defaultValue: 10 * time.Second,
			envValue:     "",
			expected:     10 * time.Second,
		},
		{
			name:         "Env var is invalid",
			key:          "INVALID_DURATION_ENV",
			defaultValue: 10 * time.Second,
			envValue:     "abc",
			expected:     10 * time.Second,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envValue != "" {
				os.Setenv(tt.key, tt.envValue)
			} else {
				os.Unsetenv(tt.key)
			}
			defer os.Unsetenv(tt.key) // Clean up

			got := getEnvDuration(tt.key, tt.defaultValue)
			if got != tt.expected {
				t.Errorf(
					"getEnvDuration(%q, %v) = %v, expected %v",
					tt.key,
					tt.defaultValue,
					got,
					tt.expected,
				)
			}
		})
	}
}

func TestGetEnvStringSlice(t *testing.T) {
	tests := []struct {
		name         string
		key          string
		defaultValue []string
		envValue     string
		expected     []string
	}{
		{
			name:         "Env var exists and is valid",
			key:          "TEST_STRING_SLICE_ENV",
			defaultValue: []string{"a", "b"},
			envValue:     "c,d,e",
			expected:     []string{"c", "d", "e"},
		},
		{
			name:         "Env var does not exist",
			key:          "NON_EXISTENT_STRING_SLICE_ENV",
			defaultValue: []string{"a", "b"},
			envValue:     "",
			expected:     []string{"a", "b"},
		},
		{
			name:         "Env var is empty",
			key:          "EMPTY_STRING_SLICE_ENV",
			defaultValue: []string{"a", "b"},
			envValue:     "",
			expected:     []string{"a", "b"},
		},
		{
			name:         "Env var with spaces",
			key:          "SPACED_STRING_SLICE_ENV",
			defaultValue: []string{"a"},
			envValue:     "  x , y  , z  ",
			expected:     []string{"x", "y", "z"},
		},
		{
			name:         "Env var with empty items",
			key:          "EMPTY_ITEMS_STRING_SLICE_ENV",
			defaultValue: []string{"a"},
			envValue:     "x,,y",
			expected:     []string{"x", "y"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envValue != "" {
				os.Setenv(tt.key, tt.envValue)
			} else {
				os.Unsetenv(tt.key)
			}
			defer os.Unsetenv(tt.key) // Clean up

			got := getEnvStringSlice(tt.key, tt.defaultValue)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf(
					"getEnvStringSlice(%q, %v) got %v, expected %v",
					tt.key,
					tt.defaultValue,
					got,
					tt.expected,
				)
			}
		})
	}
}

func TestDatabaseConfig_DatabaseURL(t *testing.T) {
	dbConfig := DatabaseConfig{
		Name:     "testdb",
		Host:     "localhost",
		Port:     5432,
		Username: "testuser",
		Password: "testpass",
		SSLMode:  "disable",
		Schema:   "public",
	}
	expected := "postgres://testuser:testpass@localhost:5432/testdb?sslmode=disable&search_path=public"
	actual := dbConfig.DatabaseURL()

	if actual != expected {
		t.Errorf("DatabaseURL() got %q, expected %q", actual, expected)
	}
}

func TestServerConfig_ServerAddress(t *testing.T) {
	serverConfig := ServerConfig{
		Host: "localhost",
		Port: 8080,
	}
	expected := "localhost:8080"
	actual := serverConfig.ServerAddress()

	if actual != expected {
		t.Errorf("ServerAddress() got %q, expected %q", actual, expected)
	}
}

func TestCORSConfig_IsValidOrigins(t *testing.T) {
	tests := []struct {
		name     string
		origins  []string
		expected bool
	}{
		{
			name:     "Allow all origins",
			origins:  []string{"*"},
			expected: true,
		},
		{
			name:     "Valid HTTP origin",
			origins:  []string{"http://example.com"},
			expected: true,
		},
		{
			name:     "Valid HTTPS origin",
			origins:  []string{"https://secure.example.com"},
			expected: true,
		},
		{
			name:     "Multiple valid origins",
			origins:  []string{"http://example.com", "https://another.com"},
			expected: true,
		},
		{
			name:     "Invalid scheme",
			origins:  []string{"ftp://example.com"},
			expected: false,
		},
		{
			name:     "Missing host",
			origins:  []string{"http://"},
			expected: false,
		},
		{
			name:     "Invalid URL format",
			origins:  []string{"invalid-url"},
			expected: false,
		},
		{
			name:     "Empty origins list",
			origins:  []string{},
			expected: true, // An empty list means no specific restrictions, which is valid.
		},
		{
			name:     "Mix of valid and invalid, should fail",
			origins:  []string{"http://valid.com", "invalid-one"},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CORSConfig{AllowedOrigins: tt.origins}
			actual := c.IsValidOrigins()
			if actual != tt.expected {
				t.Errorf(
					"IsValidOrigins() for %v got %v, expected %v",
					tt.origins,
					actual,
					tt.expected,
				)
			}
		})
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		config  *Config
		wantErr bool
	}{
		{
			name: "Valid Config",
			config: &Config{
				Server:   ServerConfig{Port: 8080},
				Database: DatabaseConfig{Password: "securepass", Port: 5432},
				JWT:      JWTConfig{Secret: "jwtsecret"},
				CORS:     CORSConfig{AllowedOrigins: []string{"*"}},
			},
			wantErr: false,
		},
		{
			name: "Missing JWT Secret",
			config: &Config{
				Server:   ServerConfig{Port: 8080},
				Database: DatabaseConfig{Password: "securepass"},
				JWT:      JWTConfig{Secret: ""},
				CORS:     CORSConfig{AllowedOrigins: []string{"*"}},
			},
			wantErr: true,
		},
		{
			name: "Missing DB Password",
			config: &Config{
				Server:   ServerConfig{Port: 8080},
				Database: DatabaseConfig{Password: ""},
				JWT:      JWTConfig{Secret: "jwtsecret"},
				CORS:     CORSConfig{AllowedOrigins: []string{"*"}},
			},
			wantErr: true,
		},
		{
			name: "Invalid Server Port (too low)",
			config: &Config{
				Server:   ServerConfig{Port: 0},
				Database: DatabaseConfig{Password: "securepass"},
				JWT:      JWTConfig{Secret: "jwtsecret"},
				CORS:     CORSConfig{AllowedOrigins: []string{"*"}},
			},
			wantErr: true,
		},
		{
			name: "Invalid Server Port (too high)",
			config: &Config{
				Server:   ServerConfig{Port: 65536},
				Database: DatabaseConfig{Password: "securepass"},
				JWT:      JWTConfig{Secret: "jwtsecret"},
				CORS:     CORSConfig{AllowedOrigins: []string{"*"}},
			},
			wantErr: true,
		},
		{
			name: "Invalid DB Port (too low)",
			config: &Config{
				Server:   ServerConfig{Port: 8080},
				Database: DatabaseConfig{Password: "securepass", Port: 0},
				JWT:      JWTConfig{Secret: "jwtsecret"},
				CORS:     CORSConfig{AllowedOrigins: []string{"*"}},
			},
			wantErr: true,
		},
		{
			name: "Invalid DB Port (too high)",
			config: &Config{
				Server:   ServerConfig{Port: 8080},
				Database: DatabaseConfig{Password: "securepass", Port: 65536},
				JWT:      JWTConfig{Secret: "jwtsecret"},
				CORS:     CORSConfig{AllowedOrigins: []string{"*"}},
			},
			wantErr: true,
		},
		{
			name: "Invalid CORS Origin",
			config: &Config{
				Server:   ServerConfig{Port: 8080},
				Database: DatabaseConfig{Password: "securepass"},
				JWT:      JWTConfig{Secret: "jwtsecret"},
				CORS:     CORSConfig{AllowedOrigins: []string{"invalid-url"}},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLoad(t *testing.T) {
	// Save original environment variables and restore them after test
	oldEnv := make(map[string]string)
	envKeys := []string{
		"SERVER_HOST", "SERVER_PORT", "SERVER_READ_TIMEOUT", "SERVER_WRITE_TIMEOUT", "SERVER_IDLE_TIMEOUT",
		"DB_NAME", "DB_HOST", "DB_PORT", "DB_USERNAME", "DB_PASSWORD", "DB_SSL_MODE", "DB_SCHEMA",
		"JWT_SECRET", "JWT_EXPIRATION", "JWT_ISSUER",
		"ACCESS_CONTROL_ALLOW_ORIGIN",
	}
	for _, key := range envKeys {
		oldEnv[key] = os.Getenv(key)
		os.Unsetenv(key)
	}
	defer func() {
		for _, key := range envKeys {
			os.Setenv(key, oldEnv[key])
		}
	}()

	tests := []struct {
		name     string
		envVars  map[string]string
		expected *Config
		wantErr  bool
	}{
		{
			name:    "Load with defaults",
			envVars: map[string]string{},
			expected: &Config{
				Server: ServerConfig{
					Host:         "localhost",
					Port:         8000,
					ReadTimeout:  10 * time.Second,
					WriteTimeout: 10 * time.Second,
					IdleTimeout:  60 * time.Second,
				},
				Database: DatabaseConfig{
					Name:     "authz",
					Host:     "localhost",
					Port:     5432,
					Username: "postgres",
					Password: "Bruh0!0!",
					SSLMode:  "disable",
					Schema:   "public",
				},
				JWT: JWTConfig{
					Secret:     "ai33yUUcmRPI64hq06ViG0404On-nMebsCtY4nTFqOg",
					Expiration: 24 * time.Hour,
					Issuer:     "myapp",
				},
				CORS: CORSConfig{
					AllowedOrigins: []string{"*"},
				},
			},
			wantErr: false,
		},
		{
			name: "Load with custom values",
			envVars: map[string]string{
				"SERVER_PORT":                 "9000",
				"DB_NAME":                     "customdb",
				"JWT_SECRET":                  "customsecret",
				"ACCESS_CONTROL_ALLOW_ORIGIN": "http://test.com,https://test.com",
			},
			expected: &Config{
				Server: ServerConfig{
					Host:         "localhost",
					Port:         9000,
					ReadTimeout:  10 * time.Second,
					WriteTimeout: 10 * time.Second,
					IdleTimeout:  60 * time.Second,
				},
				Database: DatabaseConfig{
					Name:     "customdb",
					Host:     "localhost",
					Port:     5432,
					Username: "postgres",
					Password: "Bruh0!0!",
					SSLMode:  "disable",
					Schema:   "public",
				},
				JWT: JWTConfig{
					Secret:     "customsecret",
					Expiration: 24 * time.Hour,
					Issuer:     "myapp",
				},
				CORS: CORSConfig{
					AllowedOrigins: []string{"http://test.com", "https://test.com"},
				},
			},
			wantErr: false,
		},
		{
			name: "Load with invalid server port, should error from validate",
			envVars: map[string]string{
				"SERVER_PORT": "0",
			},
			expected: nil,
			wantErr:  true,
		},
		{
			name: "Load with missing JWT secret, should error from validate",
			envVars: map[string]string{
				"JWT_SECRET": "",
			},
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set environment variables for the current test case
			for key, value := range tt.envVars {
				os.Setenv(key, value)
			}

			got, err := Load()

			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Load() got = %+v, expected %+v", got, tt.expected)
			}

			// Clean up environment variables for the next test case
			for key := range tt.envVars {
				os.Unsetenv(key)
			}
		})
	}
}

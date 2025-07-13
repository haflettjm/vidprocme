package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// env holds the minimal env vars we must set for the happy path.
var env = map[string]string{
	"PROJECT_ID":             "test-project",
	"P_LOCATION":             "us-central1",
	"PORT":                   "8080",
	"LOG_LEVEL":              "debug",
	"P_ENV":                  "test",
	"L_BUCKET_NAME":          "log-bucket",
	"VID_BUCKET_NAME":        "vid-bucket",
	"LLM_TARGET_BUCKET_NAME": "target-bucket",
	"T_OBJECT_PREFIX":        "processed/",
	"VID_OBJECT_PREFIX":      "uploads/",
	"L_OBJECT_PREFIX":        "logs/",
	"SIGNED_URL_EXPIRY":      "15",
	"VIDEO_MAX_MB":           "500",
	"UPLOAD_TIMEOUT_MIN":     "30",
	"SUBSCRIPTION_JSON_PATH": "subscriptions.json",
	"CACHE_HOST":             "redis",
	"CACHE_PORT":             "6379",
	"CACHE_PASS":             "",
	"CACHE_DB":               "0",
	"CACHE_MAX_IDLE":         "10",
	"CACHE_MAX_CONN":         "100",
	"CACHE_CONN_TIMEOUT":     "5",
	"CACHE_READ_TIMEOUT":     "3",
	"CACHE_WRITE_TIMEOUT":    "3",
	"CACHE_UPLOAD_PRE":       "up:",
	"CACHE_TARGET_PRE":       "tg:",
	"CACHE_DEV_PRE":          "dv:",
	"ALLOWED_ORIGINS":        "http://localhost:3000,https://example.com",
	"ENFORCE_HTTPS":          "true",
	"RATE_LIMIT_PER_MIN":     "20",
	"JWT_AUDIENCE":           "test-aud",
}

func setEnv(m map[string]string) func() {
	for k, v := range m {
		os.Setenv(k, v)
	}
	return func() {
		for k := range m {
			os.Unsetenv(k)
		}
	}
}

func TestConfig_Load_ok(t *testing.T) {
	cleanup := setEnv(env)
	defer cleanup()

	// create fake subscriptions.json
	tmp := t.TempDir()
	jsonPath := filepath.Join(tmp, "subscriptions.json")
	require.NoError(t, os.WriteFile(jsonPath, []byte(`{"llm":{}}`), 0644))
	os.Setenv("SUBSCRIPTION_JSON_PATH", jsonPath)

	cfg := &Config{}
	got, err := cfg.Load()
	require.NoError(t, err)
	assert.Equal(t, "test-project", got.ProjectID)
	assert.Equal(t, 8080, got.Port)
	assert.Equal(t, []string{"http://localhost:3000", "https://example.com"}, got.AllowedOrigins)
	assert.True(t, got.EnforceHttps)
	assert.Equal(t, "subscriptions.json", got.SubscriptionJsonPath)
}

func TestConfig_Load_missingRequired(t *testing.T) {
	_ = setEnv(env) // fill env, then remove one required key
	defer func() { _ = setEnv(env) }()
	os.Unsetenv("PROJECT_ID")

	cfg := &Config{}
	_, err := cfg.Load()
	assert.Error(t, err)
}

func TestConfig_Load_badNumeric(t *testing.T) {
	_ = setEnv(env)
	defer func() { _ = setEnv(env) }()
	os.Setenv("PORT", "not-a-number")

	cfg := &Config{}
	_, err := cfg.Load()
	assert.Error(t, err)
}

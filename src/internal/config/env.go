package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	ProjectID            string   `mapstructure:"PROJECT_ID"`
	Location             string   `mapstructure:"P_LOCATION"`
	Port                 int      `mapstructure:"PORT"`
	LogLevel             string   `mapstructure:"LOG_LEVEL"`
	EnvType              string   `mapstructure:"P_ENV"`
	LogBucket            string   `mapstructure:"L_BUCKET_NAME"`
	VidBucket            string   `mapstructure:"VID_BUCKET_NAME"`
	TargetBucket         string   `mapstructure:"LLM_TARGET_BUCKET_NAME"`
	LogBucketPre         string   `mapstructure:"L_OBJECT_PREFIX"`
	VideoBucketPre       string   `mapstructure:"VID_OBJECT_PREFIX"`
	TargetBucketPre      string   `mapstructure:"T_OBJECT_PREFIX"`
	SignedUrlExpiry      int      `mapstructure:"SIGNED_URL_EXPIRY"`
	VideoMaxMB           int      `mapstructure:"VIDEO_MAX_MB"`
	UploadTimeoutMin     int      `mapstructure:"UPLOAD_TIMEOUT_MIN"`
	SubscriptionJsonPath string   `mapstructure:"SUBSCRIPTION_JSON_PATH"`
	CacheHost            string   `mapstructure:"CACHE_HOST"`
	CachePort            int      `mapstructure:"CACHE_PORT"`
	CachePass            string   `mapstructure:"CACHE_PASS"`
	CacheDB              int      `mapstructure:"CACHE_DB"`
	CacheMaxIdle         int      `mapstructure:"CACHE_MAX_IDLE"`
	CacheMaxConn         string   `mapstructure:"CACHE_MAX_CONN"`
	CacheConnTimeout     int      `mapstructure:"CACHE_CONN_TIMEOUT"`
	CacheReadTimeout     int      `mapstructure:"CACHE_READ_TIMEOUT"`
	CacheWriteTimeout    int      `mapstructure:"CACHE_WRITE_TIMEOUT"`
	CacheUploadPre       string   `mapstructure:"CACHE_UPLOAD_PRE"`
	CacheTargetPre       string   `mapstructure:"CACHE_TARGET_PRE"`
	CacheDevPre          string   `mapstructure:"CACHE_DEV_PRE"`
	AllowedOrigins       []string `mapstructure:"ALLOWED_ORIGINS"` // comma-split handled below
	EnforceHttps         bool     `mapstructure:"ENFORCE_HTTPS"`
	RateLimitPerMin      int      `mapstructure:"RATE_LIMIT_PER_MIN"`
	JWTAudience          string   `mapstructure:"JWT_AUDIENCE"`
}

func (c *Config) Load() error {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	_ = viper.ReadInConfig()

	// automatic string→int, string→bool via mapstructure
	if err := viper.Unmarshal(c); err != nil {
		return fmt.Errorf("unmarshal config: %w", err)
	}

	// comma-split string → slice
	if origins := viper.GetString("ALLOWED_ORIGINS"); origins != "" {
		c.AllowedOrigins = strings.Split(origins, ",")
	}

	return nil
}

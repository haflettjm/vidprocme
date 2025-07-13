package config

import (
	"os"
	"strconv"
	"strings"
)

type Config struct {
	// Load configs
	ProjectID            string
	Location             string
	Port                 int
	LogLevel             string
	EnvType              string
	LogBucket            string
	VidBucket            string
	TargetBucket         string
	LogBucketPre         string
	VideoBucketPre       string
	TargetBucketPre      string
	SignedUrlExpiry      int
	VideoMaxMB           int
	UploadTimeoutMin     int
	SubscriptionJsonPath string
	CacheHost            string
	CachePort            int
	CachePass            string
	CacheDB              int
	CacheMaxIdle         int
	CacheMaxConn         string
	CacheConnTimeout     int
	CacheReadTimeout     int
	CacheWriteTimeout    int
	CacheUploadPre       string
	CacheTargetPre       string
	CacheDevPre          string
	AllowedOrigins       []string
	EnforceHttps         bool
	RateLimitPerMin      int
	JWTAudience          string
}

func (c *Config) Load() (*Config, error) {
	var err error

	// CORE
	c.ProjectID = os.Getenv("PROJECT_ID")
	c.Location = os.Getenv("P_LOCATION")
	c.Port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return nil, err
	}
	c.LogLevel = os.Getenv("LOG_LEVEL")
	c.EnvType = os.Getenv("P_ENV")

	// CLOUD OR SERVER LIMITS
	c.LogBucket = os.Getenv("L_BUCKET_NAME")
	c.TargetBucket = os.Getenv("LLM_TARGET_BUCKET_NAME")
	c.VidBucket = os.Getenv("VID_BUCKET_NAME")
	c.TargetBucketPre = os.Getenv("T_OBJECT_PREFIX")
	c.VideoBucketPre = os.Getenv("VID_OBJECT_PREFIX")
	c.LogBucketPre = os.Getenv("L_OBJECT_PREFIX")
	c.SignedUrlExpiry, err = strconv.Atoi(os.Getenv("SIGNED_URL_EXPIRY"))
	if err != nil {
		return nil, err
	}
	c.VideoMaxMB, err = strconv.Atoi(os.Getenv("VIDEO_MAX_MB"))
	if err != nil {
		return nil, err
	}
	c.UploadTimeoutMin, err = strconv.Atoi(os.Getenv("UPLOAD_TIMEOUT_MIN"))
	if err != nil {
		return nil, err
	}

	// PUB SUB JSON
	c.SubscriptionJsonPath = os.Getenv("SUBSCRIPTION_JSON_PATH")

	// REDIS OR CACHE ENV
	c.CacheHost = os.Getenv("CACHE_HOST")
	c.CachePort, err = strconv.Atoi(os.Getenv("CACHE_PORT"))
	if err != nil {
		return nil, err
	}
	c.CachePass = os.Getenv("CACHE_PASS")
	c.CacheDB, err = strconv.Atoi(os.Getenv("CACHE_DB"))
	if err != nil {
		return nil, err
	}
	c.CacheMaxIdle, err = strconv.Atoi(os.Getenv("CACHE_MAX_IDLE"))
	if err != nil {
		return nil, err
	}
	c.CacheMaxConn = os.Getenv("CACHE_MAX_CONN")
	c.CacheConnTimeout, err = strconv.Atoi(os.Getenv("CACHE_CONN_TIMEOUT"))
	if err != nil {
		return nil, err
	}
	c.CacheReadTimeout, err = strconv.Atoi(os.Getenv("CACHE_READ_TIMEOUT"))
	if err != nil {
		return nil, err
	}
	c.CacheWriteTimeout, err = strconv.Atoi(os.Getenv("CACHE_WRITE_TIMEOUT"))
	if err != nil {
		return nil, err
	}
	c.CacheUploadPre = os.Getenv("CACHE_UPLOAD_PRE")
	c.CacheTargetPre = os.Getenv("CACHE_TARGET_PRE")
	c.CacheDevPre = os.Getenv("CACHE_DEV_PRE")

	// SERVER LIMITS OR CORS
	c.AllowedOrigins = strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")
	c.EnforceHttps = os.Getenv("ENFORCE_HTTPS") == "true"
	c.RateLimitPerMin, err = strconv.Atoi(os.Getenv("RATE_LIMIT_PER_MIN"))
	if err != nil {
		return nil, err
	}
	c.JWTAudience = os.Getenv("JWT_AUDIENCE")

	return c, nil
}

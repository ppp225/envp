package envp

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	log "github.com/ppp225/lvlog"
)

// LoadEnvFromEnvFiles reads .env* files as in convention. See:
// https://github.com/joho/godotenv#precedence--conventions
func LoadEnvFromEnvFiles(env string) {
	if env == "" {
		env = "development"
	}
	godotenv.Load(".env." + env + ".local")
	if env != "test" {
		godotenv.Load(".env.local")
	}
	godotenv.Load(".env." + env)
	godotenv.Load() // .env
}

// GetEnvString gets env var or fallback, and logs return value
func GetEnvString(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		log.Infof("$%s=%q", key, value)
		return value
	}

	log.Warnf("$%s=%q (env var not found, using fallback)", key, defaultVal)
	return defaultVal
}

// GetEnvPassword is GetEnvString, but doesn't log full value, just ****12
func GetEnvPassword(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		logVal := "****"
		if len(value) > 3 {
			logVal += value[len(value)-2:]
		}
		log.Infof("$%s=%q", key, logVal)
		return value
	}

	log.Warnf("$%s=%q (env var not found, using fallback)", key, defaultVal)
	return defaultVal
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// GetEnvStringFrom gets env var or fallback, validates if value is allowed and logs return value
func GetEnvStringFrom(key string, defaultVal string, allowedValues []string) string {
	if value, exists := os.LookupEnv(key); exists {
		if stringInSlice(value, allowedValues) {
			log.Infof("$%s=%q", key, value)
			return value
		}
		log.Fatalf("$%s=%q | Error: Read env var value not allowed. Must be one of: %v", key, value, allowedValues)
	}

	log.Warnf("$%s=%q (env var not found, using fallback)", key, defaultVal)
	return defaultVal
}

// GetEnvFloat gets env var or fallback, validates range and logs return value.
// calls log.Fatal when env var does not validate correctly
func GetEnvFloat(key string, defaultVal float64, min float64, max float64) float64 {
	if envValue, exists := os.LookupEnv(key); exists {
		result, err := strconv.ParseFloat(envValue, 8)
		if err != nil {
			log.Fatalf("$%s=%v | Error: %v", key, envValue, err)
		}
		if !(min <= result && result <= max) {
			log.Fatalf("$%[1]s=%[2]v | Error: Outside of range: %[3]v <= %[2]v <= %[4]v", key, result, min, max)
		}
		log.Infof("$%s=%v", key, result)
		return result
	}
	log.Warnf("$%s=%v (env var not found, using fallback)", key, defaultVal)
	return defaultVal
}

// GetEnvInt gets env var or fallback, validates range and logs return value.
// calls log.Fatal when env var does not validate correctly
func GetEnvInt(key string, defaultVal int, min int, max int) int {
	if envValue, exists := os.LookupEnv(key); exists {
		result, err := strconv.Atoi(envValue)
		if err != nil {
			log.Fatalf("$%s=%v | Error: %v", key, envValue, err)
		}
		if !(min <= result && result <= max) {
			log.Fatalf("$%[1]s=%[2]v | Error: Outside of range: %[3]v <= %[2]v <= %[4]v", key, result, min, max)
		}
		log.Infof("$%s=%v", key, result)
		return result
	}
	log.Warnf("$%s=%v (env var not found, using fallback)", key, defaultVal)
	return defaultVal
}

// GetEnvBool gets env var or fallback, validates range and logs return value.
// calls log.Fatal when env var does not validate correctly
func GetEnvBool(key string, defaultVal bool) bool {
	if envValue, exists := os.LookupEnv(key); exists {
		result, err := strconv.ParseBool(envValue)
		if err != nil {
			log.Fatalf("$%s=%v | Error: %v", key, envValue, err)
		}
		log.Infof("$%s=%v", key, result)
		return result
	}
	log.Warnf("$%s=%v (env var not found, using fallback)", key, defaultVal)
	return defaultVal
}

// SetLogLevelFromEnv reads loglevel from env and sets it.
// for envp, set to: info, warn, fatal, none. Defaults to info.
func SetLogLevelFromEnv(key string) {
	lv := GetEnvString(key, "info")
	switch lv {
	case "none":
		log.SetLevel(log.NONE)
	case "fatal":
		log.SetLevel(log.FATAL)
	case "panic":
		log.SetLevel(log.PANIC | log.FATAL)
	case "error":
		log.SetLevel(log.ERROR | log.PANIC | log.FATAL)
	case "warn":
		log.SetLevel(log.WARN | log.ERROR | log.PANIC | log.FATAL)
	case "info":
		log.SetLevel(log.NORM)
	case "debug":
		log.SetLevel(log.NORM | log.DEBUG)
	case "trace":
		log.SetLevel(log.ALL)
	default:
		log.Warnf("loglevel of value=%v does not exist. Leaving default.", lv)
	}
}

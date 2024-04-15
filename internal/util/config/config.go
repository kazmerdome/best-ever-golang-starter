package config

type Config interface {
	GetString(key string) string
	GetBool(key string) bool
	GetInt(key string) int
	SetDefault(key string, value interface{})
	LoadConfigFile(path, configType, configFile string) error
}

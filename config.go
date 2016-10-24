package singularity

import "sync"

type Configuration interface {
	GetBool(string) bool
	CheckBool(string) (bool, bool)
	GetString(string) string
	CheckString(string) (string, bool)
}

type defaultConfig struct {
	sync.Mutex
	config map[string]interface{}
}

//GetBool returns the bool value of key, and defaults to false if it can't find key.
func (config defaultConfig) GetBool(key string) bool {
	if val1, ok := config.config[key]; ok {
		if val2, ok := val1.(bool); ok {
			return val2
		}
	}
	return false
}

//CheckBool returns the bool value of key, and whether or not it actually found key.
func (config defaultConfig) CheckBool(key string) (bool, bool) {
	if val1, ok := config.config[key]; ok {
		if val2, ok := val1.(bool); ok {
			return val2, true
		}
	}
	return false, false
}

//GetString returns the bool value of key, and defaults to false if it can't find key.
func (config defaultConfig) GetString(key string) string {
	if val1, ok := config.config[key]; ok {
		if val2, ok := val1.(string); ok {
			return val2
		}
	}
	return ""
}

//CheckString returns the bool value of key, and whether or not it actually found key.
func (config defaultConfig) CheckString(key string) (string, bool) {
	if val1, ok := config.config[key]; ok {
		if val2, ok := val1.(string); ok {
			return val2, true
		}
	}
	return "", false
}

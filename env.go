package env

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/cloudokyo/cast"
)

func formatKey(v string) string {
	if len(v) == 0 {
		return ""
	}

	v = strings.ToUpper(v)
	return strings.Replace(v, ".", "_", -1)
}

// Get an env value with default support
func Get(path string, value ...string) string {
	v := os.Getenv(formatKey(path))

	// found value from env
	if len(v) > 0 {
		return v
	}

	// check for default value
	if v = ""; len(value) > 0 {
		v = value[0]
	}

	return v
}

// GetBool an env value with default support
func GetBool(path string, value ...bool) bool {
	v := Get(path)

	// found value from env
	if len(v) > 0 {
		if value, err := cast.ToBoolE(v); err != nil {
			log.Fatalf("%s --> env.GetBool(%s)", err, path)
		} else {
			return value
		}
	}

	r := false

	// check for default value
	if len(value) > 0 {
		r = value[0]
	} else {
		log.Fatalf("NOT FOUND --> env.GetBool(%s)", path)
	}

	return r
}

// GetInt an env value with default support
func GetInt(path string, value ...int) int {
	v := Get(path)

	// found value from env
	if len(v) > 0 {
		if value, err := cast.ToIntE(v); err != nil {
			log.Fatalf("%s --> env.GetInt(%s)", err, path)
		} else {
			return value
		}
	}

	r := 0

	// check for default value
	if len(value) > 0 {
		r = value[0]
	} else {
		log.Fatalf("NOT FOUND --> env.GetInt(%s)", path)
	}

	return r
}

// GetInt32 an env value with default support
func GetInt32(path string, value ...int32) int32 {
	v := Get(path)

	// found value from env
	if len(v) > 0 {
		if value, err := cast.ToInt32E(v); err != nil {
			log.Fatalf("%s --> env.GetInt32(%s)", err, path)
		} else {
			return value
		}
	}

	r := int32(0)

	// check for default value
	if len(value) > 0 {
		r = value[0]
	} else {
		log.Fatalf("NOT FOUND --> env.GetInt32(%s)", path)
	}

	return r
}

// GetInt64 an env value with default support
func GetInt64(path string, value ...int64) int64 {
	v := Get(path)

	// found value from env
	if len(v) > 0 {
		if value, err := cast.ToInt64E(v); err != nil {
			log.Fatalf("%s --> env.GetInt64(%s)", err, path)
		} else {
			return value
		}
	}

	r := int64(0)

	// check for default value
	if len(value) > 0 {
		r = value[0]
	} else {
		log.Fatalf("NOT FOUND --> env.GetInt64(%s)", path)
	}

	return r
}

// GetFloat32 an env value with default support
func GetFloat32(path string, value ...float32) float32 {
	v := Get(path)

	// found value from env
	if len(v) > 0 {
		if value, err := cast.ToFloat32E(v); err != nil {
			log.Fatalf("%s --> env.GetFloat32(%s)", err, path)
		} else {
			return value
		}
	}

	r := float32(0)

	// check for default value
	if len(value) > 0 {
		r = value[0]
	} else {
		log.Fatalf("NOT FOUND --> env.GetFloat32(%s)", path)
	}

	return r
}

// GetFloat64 an env value with default support
func GetFloat64(path string, value ...float64) float64 {
	v := Get(path)

	// found value from env
	if len(v) > 0 {
		if value, err := cast.ToFloat64E(v); err != nil {
			log.Fatalf("%s --> env.GetFloat64(%s)", err, path)
		} else {
			return value
		}
	}

	r := float64(0)

	// check for default value
	if len(value) > 0 {
		r = value[0]
	} else {
		log.Fatalf("NOT FOUND --> env.GetFloat64(%s)", path)
	}

	return r
}

// GetTime an env value with default support
func GetTime(path string, value ...time.Time) time.Time {
	v := Get(path)

	// found value from env
	if len(v) > 0 {
		if value, err := cast.ToTimeE(v); err != nil {
			log.Fatalf("%s --> env.GetTime(%s)", err, path)
		} else {
			return value
		}
	}

	r := time.Time{}

	// check for default value
	if len(value) > 0 {
		r = value[0]
	} else {
		log.Fatalf("NOT FOUND --> env.GetTime(%s)", path)
	}

	return r
}

// GetDuration an env value with default support
func GetDuration(path string, value ...time.Duration) time.Duration {
	v := Get(path)

	// found value from env
	if len(v) > 0 {
		if value, err := cast.ToDurationE(v); err != nil {
			log.Fatalf("%s --> env.GetDuration(%s)", err, path)
		} else {
			return value
		}
	}

	// check for default value
	if len(value) > 0 {
		return value[0]
	} else {
		log.Fatalf("NOT FOUND --> env.GetDuration(%s)", path)
	}

	return 0
}

// GetMapStringBool an env value with default support
func GetMapStringBool(path string, defaults ...string) (ret map[string]bool) {
	v := Get(path, defaults...)

	// init the result
	ret = map[string]bool{}

	// found value from env
	if len(v) == 0 {
		return ret
	}

	chunks := strings.Split(v, ",")
	for _, item := range chunks {
		elements := strings.Split(item, "=")
		length := len(elements)

		if length == 1 {
			ret[strings.TrimSpace(elements[0])] = true
		} else if length == 2 {
			ret[strings.TrimSpace(elements[0])] = cast.ToBool(elements[1])
		}
	}

	return ret
}

// GetMapStringInt an env value with default support
func GetMapStringInt(path string, defaults ...string) (ret map[string]int) {
	v := Get(path, defaults...)

	// init the result
	ret = map[string]int{}

	// found value from env
	if len(v) == 0 {
		return ret
	}

	chunks := strings.Split(v, ",")
	for _, item := range chunks {
		elements := strings.Split(item, "=")
		length := len(elements)

		if length == 1 {
			ret[strings.TrimSpace(elements[0])] = 0
		} else if length == 2 {
			ret[strings.TrimSpace(elements[0])] = cast.ToInt(elements[1])
		}
	}

	return ret
}

// Set an env string value
func Set(path string, v string) error {
	return os.Setenv(formatKey(path), v)
}

// Delete an env value
func Delete(path string) error {
	v := formatKey(path)
	return os.Unsetenv(v)
}

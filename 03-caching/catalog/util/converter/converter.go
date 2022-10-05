package converter

import (
	"strconv"
	"strings"
)

// ToInt64 converts any value to int64
func ToInt64(v interface{}) int64 {
	switch v := v.(type) {
	case string:
		str := strings.TrimSpace(v)
		result, err := strconv.Atoi(str)
		if err != nil {
			return 0
		}
		return int64(result)
	case int:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return v
	case float32:
		return int64(v)
	case float64:
		return int64(v)
	case []byte:
		result, err := strconv.Atoi(string(v))
		if err != nil {
			return 0
		}
		return int64(result)
	case bool:
		if v {
			return 1
		}
		return 0
	default:
		return 0
	}
}

func ToInt(v interface{}) int {
	switch v := v.(type) {
	case string:
		str := strings.TrimSpace(v)
		result, err := strconv.Atoi(str)
		if err != nil {
			return 0
		}
		return result
	case int:
		return v
	case int32:
		return int(v)
	case int64:
		return int(v)
	case float32:
		return int(v)
	case float64:
		return int(v)
	case []byte:
		result, err := strconv.Atoi(string(v))
		if err != nil {
			return 0
		}
		return result
	case bool:
		if v {
			return 1
		}
		return 0
	default:
		return 0
	}
}

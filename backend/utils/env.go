package utils

import (
	"os"
	"strconv"
)

func GetEnvString(key string) string {
	return os.Getenv(key)
}

func GetEnvUint(key string) uint {
	v, _ := strconv.ParseUint(os.Getenv(key), 10, 64)
	return uint(v)
}

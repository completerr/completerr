package utils

import (
	"context"
	"strings"
	"time"
)

const DefaultTimeout time.Duration = 120

func GetEnv(setEnv string, defaultEnv string) string {
	if setEnv == "" {
		return defaultEnv
	}

	return setEnv
}
func GetCtx(timeout time.Duration) (context.Context, context.CancelFunc) {

	return context.WithTimeout(context.Background(), time.Second*timeout)
}

func Contains(s []string, str string, caseSensitive bool) bool {
	for _, v := range s {
		if caseSensitive {
			if v == str {
				return true
			}
		} else {
			if strings.ToLower(v) == strings.ToLower(str) {
				return true
			}
		}
	}

	return false
}

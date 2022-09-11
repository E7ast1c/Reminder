package empty

import (
	"github.com/apex/log"
	"os"
)

func EnvMustNotEmpty(key string) (val string) {
	if val = os.Getenv(key); val == "" {
		log.Fatalf("empty env variable:%s", key)
	}
	return val
}

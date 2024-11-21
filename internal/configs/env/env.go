package env

import (
	"fmt"
	"os"
)

func GetEnv(key string) string {
	env, exists := os.LookupEnv(key)
	if !exists {
		panic(fmt.Sprintf("Não foi possível encontrar a variável de ambiente %s", key))
	}

	return env
}

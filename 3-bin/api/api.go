package api

import (
	"app/hw-3/config"
	"fmt"
)

func WorkWithApi() {
	config, err := config.NewCongig()
	if err != nil {
		fmt.Errorf("НЕ_ЗАГРУЗИЛИ_КЛЮЧ")
	}
	fmt.Printf("Key: %s\n", config.Key)
}

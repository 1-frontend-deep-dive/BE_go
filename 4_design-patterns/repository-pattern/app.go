package goDesignPattern

import "fmt"

func PlayWithRepoDP() {
	appConfig := AppConfig{
		Port: 3000,
	}

	SetRepoConfig(&appConfig)

	fmt.Println(*AppRepo.AppConfig)
}
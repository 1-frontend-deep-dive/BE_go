package goDesignPattern

type Repository struct {
	AppConfig *AppConfig
}

var AppRepo *Repository

func SetRepoConfig(appConfig *AppConfig) {
	updatedAppRepo := Repository{
		AppConfig: appConfig,
	}

	AppRepo = &updatedAppRepo
}

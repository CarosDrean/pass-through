package bootstrap

import (
	"context"
	"os"

	tracing "dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-extension-rkgin-tracing.git"
	"github.com/joho/godotenv"
	rkentry "github.com/rookie-ninja/rk-entry/v2/entry"

	"pass-trougth/infraestructure/handler"
	"pass-trougth/model"
)

func Run(boot []byte) {
	// we ignore the error because it loads the .env file, and in production the envs are defined differently
	_ = godotenv.Load()

	ginEntry := newGinEntry(boot)
	logger := newLogger()

	api := ginEntry.Router
	api.Use(tracing.LoggerHandler(logger, getDebugMode(), getMaxCharBodyLogger(), model.ErrorApi{
		Code:        model.InernalServerErrorCode,
		Description: model.InternalServerErrorDescription,
	}))

	handler.InitRoutes(model.RouterSpecification{
		Api:    api,
		Logger: logger,
		Config: model.ServiceConfig{
			Url:             os.Getenv("EBIZ_ENDPOINT_ACCEPTANCEHAS"),
			UrlRetrieve:     os.Getenv("EBIZ_ENDPOINT_ACCEPTANCEHASRETRIEVE"),
			UrlAuth:         os.Getenv("EBIZ_ENDPOINT_AUTH"),
			User:            os.Getenv("EBIZ_OAUTH_USER"),
			Password:        os.Getenv("EBIZ_OAUTH_PASSWORD"),
			UserBasic:       os.Getenv("EBIZ_BASICAUTH_USER"),
			PasswordBasic:   os.Getenv("EBIZ_BASICAUTH_PASSWORD"),
			SubscriptionKey: os.Getenv("EBIZ_SUBSCRIPTION_KEY"),
		},
	})

	ginEntry.Bootstrap(context.Background())
	rkentry.GlobalAppCtx.WaitForShutdownSig()
	ginEntry.Interrupt(context.Background())
}

package injector

import (
	"bbdate/pkg/config"
	"bbdate/pkg/crypter"
	"bbdate/pkg/logging"
	"fmt"

	"go.uber.org/dig"
)

var (
	c = dig.New()
)

func RunAPI() *dig.Container {
	provideCommonInjection()
	provideAPIInjection()
	return c
}

func provideCommonInjection() {
	// crypter
	// TODO: 本番環境ではAWS使うようにしたい
	if err := c.Provide(func() crypter.ICrypter {
		return crypter.NewCrypter(config.EnvStore.HashSecret)
	}); err != nil {
		logging.Fatal("system", fmt.Sprintf("cannot provide cryptor =>%v", err))
	}

	// db connection
	provideDBConnections()
}

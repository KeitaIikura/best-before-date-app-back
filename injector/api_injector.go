package injector

import (
	controller "bbdate/internal/bbdate/adapter/controller"
	"bbdate/internal/bbdate/adapter/gateway"
	bbdate_usecase "bbdate/internal/bbdate/application/usecase"
	"bbdate/internal/bbdate/domain/service"
	"bbdate/pkg/logging"

	"go.uber.org/multierr"
)

func provideAPIInjection() {
	var err error
	// controller
	err = multierr.Append(err, c.Provide(controller.NewAuthController))

	// usecase
	err = multierr.Append(err, c.Provide(bbdate_usecase.NewAuthCheckInteractor))
	err = multierr.Append(err, c.Provide(bbdate_usecase.NewAuthLoginInteractor))
	err = multierr.Append(err, c.Provide(bbdate_usecase.NewAuthLogoutInteractor))
	err = multierr.Append(err, c.Provide(bbdate_usecase.NewAuthUpdatePasswordInteractor))

	// service
	err = multierr.Append(err, c.Provide(service.NewAuthService))

	// repository (gateway)
	err = multierr.Append(err, c.Provide(gateway.NewUserGateway))

	if err != nil {
		errs := multierr.Errors(err)
		for _, openErr := range errs {
			logging.Info("system", openErr)
		}
		logging.Fatal("system", "failed to inject.")
	}
}

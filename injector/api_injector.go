package injector

import (
	controller "bbdate/internal/bbdate/adapter/controller"
	"bbdate/pkg/logging"

	"go.uber.org/multierr"
)

func provideAPIInjection() {
	var err error
	// controller
	err = multierr.Append(err, c.Provide(controller.NewAuthController))

	// usecase
	// err = multierr.Append(err, c.Provide(management_usecase.NewAuthLoginInteractor))

	// service
	if err != nil {
		errs := multierr.Errors(err)
		for _, openErr := range errs {
			logging.Info("system", openErr)
		}
		logging.Fatal("system", "failed to inject.")
	}
}

package server

import (
	"fmt"
	"net/http"

	"github.com/DustyRat/go-webapp/internal/config"
	"github.com/DustyRat/go-webapp/internal/controller"
	"github.com/DustyRat/go-webapp/internal/handler"
	"github.com/DustyRat/go-webapp/internal/service"

	router "github.com/DustyRat/go-metrics/router/mux"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/rs/cors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Run configures and creates a new http.Server to be used for the application to listen on
func Run(info *service.BuildInfo) error {
	conf, err := config.GetConfig()
	if err != nil {
		return err
	}

	info.Debug = conf.Debug
	level, err := zerolog.ParseLevel(conf.LogLevel)
	if err != nil {
		level = zerolog.ErrorLevel
		log.Warn().Err(err).Msgf("unable to parse log level, logging level is set to %s", level.String())
	}
	zerolog.SetGlobalLevel(level)
	log.Logger = log.With().Str("application", conf.Name).Logger()

	ctrl, err := controller.New(conf)
	if err != nil {
		return errors.Wrap(err, "unable to create controller")
	}

	mux := mux.NewRouter()
	r := router.New(mux)
	handler.AddHandlers(r, info, ctrl, conf.Debug)

	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", conf.Port),
		Handler: cors.Default().Handler(r),
	}

	log.Info().Msgf("Server running %v", srv.Addr)
	return srv.ListenAndServe()
}

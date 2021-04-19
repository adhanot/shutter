// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"shutter/database"
	"shutter/handlers/chambrehandlers"
	"shutter/handlers/salonhandlers"
	"shutter/handlers/sdbhandlers"
	"shutter/handlers/usershandlers"
	"shutter/pkg/environment"
	"shutter/restapi/operations"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
)

//go:generate swagger generate server --target ../../shutter --name Shutter --spec ../gen/api/swagger.yaml --principal interface{}

func configureFlags(api *operations.ShutterAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ShutterAPI) http.Handler {
	// configure the api here
	environment.GetEnvironment()
	postgres := database.NewClient()
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()
	api.BearerAuth = usershandlers.ValidateHeader
	api.SalonSalonOpenHandler = salonhandlers.SalonOpenHandler()
	api.SalonSalonCloseHandler = salonhandlers.SalonCloseHandler()
	api.SalonSalonStopHandler = salonhandlers.SalonStopHandler()
	api.ChambreChambreCloseHandler = chambrehandlers.ChambreCloseHandler()
	api.ChambreChambreOpenHandler = chambrehandlers.ChambreOpenHandler()
	api.ChambreChambreStopHandler = chambrehandlers.ChambreStopHandler()
	api.SdbSdbCloseHandler = sdbhandlers.SdbCloseHandler()
	api.SdbSdbOpenHandler = sdbhandlers.SdbOpenHandler()
	api.SdbSdbStopHandler = sdbhandlers.SdbStopHandler()
	api.UserLoginHandler = usershandlers.NewUsersPseudoHandler(postgres)
	api.UserRegisterHandler = usershandlers.NewUsersCreateHandler(postgres)
	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}

// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"codecentric.de/beer-fridge-go-swagger/beer_container"
	"codecentric.de/beer-fridge-go-swagger/models"
	"codecentric.de/beer-fridge-go-swagger/temperature"
	"crypto/tls"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"net/http"

	"codecentric.de/beer-fridge-go-swagger/restapi/operations"
	"codecentric.de/beer-fridge-go-swagger/restapi/operations/beers"
	"codecentric.de/beer-fridge-go-swagger/restapi/operations/fridge"
)

//go:generate swagger generate server --target ../../beer-fridge-gs --name BeerFridge --spec ../swagger.yml --principal interface{}

func configureFlags(api *operations.BeerFridgeAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.BeerFridgeAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.BeersAddOneHandler = beers.AddOneHandlerFunc(func(params beers.AddOneParams) middleware.Responder {
		if err := beer_container.AddBeer(params.Body); err != nil {
			return beers.NewAddOneDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
		}
		return beers.NewAddOneCreated().WithPayload(params.Body)
	})

	api.BeersDestroyOneHandler = beers.DestroyOneHandlerFunc(func(params beers.DestroyOneParams) middleware.Responder {
		beer_container.DeleteBeer(params.ID)
		return beers.NewDestroyOneNoContent()
	})

	api.BeersGetAllBeersHandler = beers.GetAllBeersHandlerFunc(func(params beers.GetAllBeersParams) middleware.Responder {
		mergedParams := beers.NewGetAllBeersParams()
		if params.Limit != nil {
			mergedParams.Limit = params.Limit
		}
		return beers.NewGetAllBeersOK().WithPayload(beer_container.AllBeers(*mergedParams.Limit))
	})

	api.FridgeGetTemperatureHandler = fridge.GetTemperatureHandlerFunc(func(params fridge.GetTemperatureParams) middleware.Responder {
		return fridge.NewGetTemperatureOK().WithPayload(temperature.GetTemperature())
	})

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.BeersAddOneHandler == nil {
		api.BeersAddOneHandler = beers.AddOneHandlerFunc(func(params beers.AddOneParams) middleware.Responder {
			return middleware.NotImplemented("operation beers.AddOne has not yet been implemented")
		})
	}
	if api.BeersDestroyOneHandler == nil {
		api.BeersDestroyOneHandler = beers.DestroyOneHandlerFunc(func(params beers.DestroyOneParams) middleware.Responder {
			return middleware.NotImplemented("operation beers.DestroyOne has not yet been implemented")
		})
	}
	if api.BeersGetAllBeersHandler == nil {
		api.BeersGetAllBeersHandler = beers.GetAllBeersHandlerFunc(func(params beers.GetAllBeersParams) middleware.Responder {
			return middleware.NotImplemented("operation beers.GetAllBeers has not yet been implemented")
		})
	}
	if api.FridgeGetTemperatureHandler == nil {
		api.FridgeGetTemperatureHandler = fridge.GetTemperatureHandlerFunc(func(params fridge.GetTemperatureParams) middleware.Responder {
			return middleware.NotImplemented("operation fridge.GetTemperature has not yet been implemented")
		})
	}

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

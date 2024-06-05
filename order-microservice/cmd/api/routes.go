package main

import (
	"net/http"
	"order-microservice/internal/auth"
	"order-microservice/internal/cart"
	"order-microservice/internal/order"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	// Initialize a new httprouter router instance.
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	cartRepo := cart.NewMockDB()
	cartService := cart.NewService(cartRepo)
	cartHandler := cart.NewHandler(cartService)

	orderRepo := order.NewMockDB()
	orderService := order.NewService(orderRepo)
	orderHandler := order.NewHandler(orderService)

	router.HandlerFunc(http.MethodGet, "/cart", wrapHandler(auth.AuthMiddleware(http.HandlerFunc(cartHandler.HandleCart))))
	router.HandlerFunc(http.MethodPost, "/cart", wrapHandler(auth.AuthMiddleware(http.HandlerFunc(cartHandler.HandleCart))))
	router.HandlerFunc(http.MethodPost, "/order", wrapHandler(auth.AuthMiddleware(http.HandlerFunc(orderHandler.HandleOrder))))
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	// Return the httprouter instance.
	// wrapping the router with rateLimiter() middleware to limit requests' frequency
	return app.recoverPanic(app.rateLimit(router))
}

func wrapHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
}

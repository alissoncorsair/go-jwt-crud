package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/alissoncorsair/goapi/service/cart"
	"github.com/alissoncorsair/goapi/service/order"
	"github.com/alissoncorsair/goapi/service/product"
	"github.com/alissoncorsair/goapi/service/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr,
		db,
	}
}

func (server *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	log.Println("running on ", server.addr)

	//store would be like a repository
	userStore := user.NewStore(server.db)
	orderStore := order.NewStore(server.db)
	productStore := product.NewStore(server.db)
	userHandler := user.NewHandler(userStore)
	productHandler := product.NewHandler(productStore)
	cartHandler := cart.NewHandler(orderStore, productStore, userStore)
	userHandler.RegisterRoutes(subrouter)
	productHandler.RegisterRoutes(subrouter)
	cartHandler.RegisterRoutes(subrouter)

	return http.ListenAndServe(server.addr, subrouter)
}

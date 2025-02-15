package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/willy-r/ecom-example/service/cart"
	"github.com/willy-r/ecom-example/service/order"
	"github.com/willy-r/ecom-example/service/product"
	"github.com/willy-r/ecom-example/service/user"
)

type ApiServer struct {
	addr string
	db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *ApiServer {
	return &ApiServer{addr: addr, db: db}
}

func (s *ApiServer) Start() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	productStore := product.NewStore(s.db)
	productHanlder := product.NewHandler(productStore)
	productHanlder.RegisterRoutes(subrouter)

	orderStore := order.NewStore(s.db)
	cartHandler := cart.NewHandler(orderStore, productStore, userStore)
	cartHandler.RegisterRoutes(subrouter)

	log.Println("starting server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}

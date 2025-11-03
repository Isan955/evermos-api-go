package main

import (
	"log"
	"net/http"
	"os"

	"evermos-api/config"
	"evermos-api/internal/database"
	"evermos-api/internal/entity"
	"evermos-api/internal/handler"
	"evermos-api/internal/middleware"
	"evermos-api/internal/repository"
	"evermos-api/internal/usecase"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	config.InitDB()

	config.DB.AutoMigrate(
		&entity.User{},
		&entity.Store{},
		&entity.Product{},
		&entity.Category{},
		&entity.Transaction{},
		&entity.TransactionItem{},
	)

	database.Seed()

	// Repos
	userRepo := repository.NewUserRepository()
	storeRepo := repository.NewStoreRepository()
	addressRepo := repository.NewAddressRepository()
	categoryRepo := repository.NewCategoryRepository()
	productRepo := repository.NewProductRepository()
	transactionRepo := repository.NewTransactionRepository()

	// Usecases
	authUsecase := usecase.NewAuthUsecase(userRepo, storeRepo)
	addressUsecase := usecase.NewAddressUsecase(addressRepo)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)
	productUsecase := usecase.NewProductUsecase(productRepo)
	transactionUsecase := usecase.NewTransactionUsecase(productRepo, transactionRepo)

	// Handlers
	authHandler := handler.NewAuthHandler(authUsecase)
	storeHandler := handler.NewStoreHandler(storeRepo)
	addressHandler := handler.NewAddressHandler(addressUsecase)
	categoryHandler := handler.NewCategoryHandler(categoryUsecase)
	productHandler := handler.NewProductHandler(productUsecase)
	transactionHandler := handler.NewTransactionHandler(transactionUsecase)

	// Router
	r := mux.NewRouter()

	// Public Routes
	r.HandleFunc("/register", authHandler.Register).Methods("POST")
	r.HandleFunc("/login", authHandler.Login).Methods("POST")
	r.HandleFunc("/categories", categoryHandler.GetAll).Methods("GET")

	// Auth Required Routes
api := r.PathPrefix("/api").Subrouter()
api.Use(middleware.JWTMiddleware)

// Store
api.HandleFunc("/store/me", storeHandler.GetMyStore).Methods("GET")

// Address
api.HandleFunc("/address", addressHandler.Create).Methods("POST")
api.HandleFunc("/address", addressHandler.GetAll).Methods("GET")
api.HandleFunc("/address/{id}", addressHandler.Update).Methods("PUT")
api.HandleFunc("/address/{id}", addressHandler.Delete).Methods("DELETE")

// Product
api.HandleFunc("/products", productHandler.GetAll).Methods("GET")
api.HandleFunc("/products/{id}", productHandler.GetByID).Methods("GET")
api.HandleFunc("/products", productHandler.Create).Methods("POST")
api.HandleFunc("/products/{id}", productHandler.Update).Methods("PUT")
api.HandleFunc("/products/{id}", productHandler.Delete).Methods("DELETE")

// Transaction ✅ FIXED ✅
api.HandleFunc("/transactions", transactionHandler.Create).Methods("POST")
api.HandleFunc("/transactions", transactionHandler.GetMyTransactions).Methods("GET")
api.HandleFunc("/transactions/{id}", transactionHandler.GetByID).Methods("GET")
api.HandleFunc("/transactions/{id}/cancel", transactionHandler.Cancel).Methods("PUT")

// Admin-only
admin := api.PathPrefix("/admin").Subrouter()
admin.Use(middleware.AdminOnly)
admin.HandleFunc("/categories", categoryHandler.Create).Methods("POST")
admin.HandleFunc("/categories/{id}", categoryHandler.Update).Methods("PUT")
admin.HandleFunc("/categories/{id}", categoryHandler.Delete).Methods("DELETE")


	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, _ := route.GetPathTemplate()
		methods, _ := route.GetMethods()
		log.Println("Route:", methods, path)
		return nil
	})

	port := os.Getenv("APP_PORT")
	log.Println("Server running on port", port)
	http.ListenAndServe(":"+port, r)
}

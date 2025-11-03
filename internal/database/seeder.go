package database

import (
	"log"

	"evermos-api/config"
	"evermos-api/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

func Seed() {
	// Seed Users (Admin & Buyer)
	password, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)

	admin := entity.User{
		Name:     "Admin Evermos",
		Email:    "admin@evermos.com",
		Phone:    "081111111111",
		Password: string(password),
		Role:     "admin",
	}

	buyer := entity.User{
		Name:     "Buyer Test",
		Email:    "buyer@evermos.com",
		Phone:    "082222222222",
		Password: string(password),
		Role:     "buyer",
	}

	config.DB.FirstOrCreate(&admin, entity.User{Email: admin.Email})
	config.DB.FirstOrCreate(&buyer, entity.User{Email: buyer.Email})

	log.Println("Seeded users ✅")

	// Seed Categories
	categories := []entity.Category{
		{Name: "Fashion Muslim"},
		{Name: "Kosmetik Halal"},
		{Name: "Perlengkapan Ibadah"},
	}

	for _, c := range categories {
		config.DB.FirstOrCreate(&c, entity.Category{Name: c.Name})
	}

	log.Println("Seeded categories ✅")

	// Seed Products
	products := []entity.Product{
		{Name: "Hijab Premium", Stock: 50, Price: 45000, CategoryID: 1},
		{Name: "Serum Halal", Stock: 30, Price: 95000, CategoryID: 2},
		{Name: "Peci Rajut", Stock: 100, Price: 25000, CategoryID: 3},
	}

	for _, p := range products {
		config.DB.FirstOrCreate(&p, entity.Product{Name: p.Name})
	}

	log.Println("Seeded products ✅")
}

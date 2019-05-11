package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sat0yu/sqlboiler-goose-sandbox/models"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres@database?sslmode=disable")
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}
	fmt.Println("connected the database")
	defer db.Close()

	ctx := context.Background()
	fmt.Println("created a context")

	err = db.Ping()
	if err != nil {
		panic(fmt.Sprintf("failed to send a ping: %v", err))
	}

	// Create
	p1 := &models.Product{
		Code: null.String{String: "L1212", Valid: true},
		Price: null.Int{Int: 1000, Valid: true},
	}
	p1.Insert(ctx, db, boil.Infer())
	p2 := &models.Product{
		Code: null.String{String: "L1213", Valid: true},
		Price: null.Int{Int: 2000, Valid: true},
	}
	p2.Insert(ctx, db, boil.Infer())
	fmt.Println("Inserted two records")

	products, err := models.Products().All(ctx, db)
	for _, p := range products {
		fmt.Println(p)
	}

	// Read
	product, err := models.Products(qm.Where("code = ?","L1213")).One(ctx, db)
	fmt.Println("Fetched the record which has code 'L1213'")
	fmt.Println(product)

	// Update - update product's price to 3000
	product.Price = null.Int{Int: 3000, Valid: true}
	product.Update(ctx, db, boil.Infer())
	fmt.Println("Updated the record which has code 'L1213'")

	products, err = models.Products().All(ctx, db)
	for _, p := range products {
		fmt.Println(p)
	}

	// Delete - delete product
	product.Delete(ctx, db)
	fmt.Println("Deleted the record which has code 'L1213'")

	products, err = models.Products().All(ctx, db)
	for _, p := range products {
		fmt.Println(p)
	}

	rowAffected, err := models.Products().DeleteAll(ctx, db)
	fmt.Println("Deleted", rowAffected, "records")
}

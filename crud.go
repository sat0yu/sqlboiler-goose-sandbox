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

func crud() {
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
		Code: null.StringFrom( "L1212"),
		Price: null.IntFrom( 1000),
	}
	p1.Insert(ctx, db, boil.Infer())
	p2 := &models.Product{
		Code: null.StringFrom( "L1213"),
		Price: null.IntFrom( 2000),
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
	product.Price = null.IntFrom( 3000)
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

	var rowAffected int64
	rowAffected, err = models.Products().DeleteAll(ctx, db)
	fmt.Println("Deleted", rowAffected, "records(products)")
	rowAffected, err = models.Orders().DeleteAll(ctx, db)
	fmt.Println("Deleted", rowAffected, "records(orders)")
}

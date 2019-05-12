package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sat0yu/sqlboiler-goose-sandbox/models"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func relations() {
	db, err := sql.Open("postgres", "postgres://postgres@database?sslmode=disable")
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}
	fmt.Println("connected the database")
	defer db.Close()

	//// global connection setting
	//boil.SetDB(db)
	//boil.DebugMode = true

	ctx := context.Background()
	fmt.Println("created a context")

	// Create
	products := []models.Product{
		models.Product{
			Code: null.StringFrom("L1212"),
			Price: null.IntFrom(1000),
		},
		models.Product{
			Code: null.StringFrom("L1213"),
			Price: null.IntFrom(2000),
		},
	}
	for _, p := range products {
		p.Insert(ctx, db, boil.Infer())
	}
	createdProducts, err := models.Products().All(ctx, db)
	orders := []models.Order{
		models.Order{
			Status: null.StringFrom("paid"),
			ProductID: null.IntFrom(createdProducts[0].ID),
		},
		models.Order{
			Status: null.StringFrom("cancelled"),
			ProductID: null.IntFrom(createdProducts[1].ID),
		},
		models.Order{
			Status: null.StringFrom("unpaid"),
			ProductID: null.IntFrom(createdProducts[1].ID),
		},
	}
	for _, o := range orders {
		o.Insert(ctx, db, boil.Infer())
	}

	fetchedOrders, err := models.Orders(qm.Load(models.OrderRels.Product)).All(ctx, db)
	for _, o := range fetchedOrders {
		fmt.Println(o.Status)
		fmt.Printf("\t%s\n", o.R.Product.Code)
	}

	fetchedProducts, err := models.Products(qm.Load(models.ProductRels.Orders)).All(ctx, db)
	for _, p := range fetchedProducts {
		fmt.Println(p.Code)
		for _, o := range p.R.Orders {
			fmt.Printf("\t%s\n", o.Status)
		}
	}

	var rowAffected int64
	rowAffected, err = models.Orders().DeleteAll(ctx, db)
	fmt.Println("Deleted", rowAffected, "records(orders)")
	rowAffected, err = models.Products().DeleteAll(ctx, db)
	fmt.Println("Deleted", rowAffected, "records(products)")
}

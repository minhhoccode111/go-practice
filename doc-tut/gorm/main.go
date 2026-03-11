package main

import (
	"context"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	ctx := context.Background()

	// migrate the schema
	db.AutoMigrate(&Product{})

	// create
	err = gorm.G[Product](db).Create(ctx, &Product{Code: "taivisao", Price: 69})

	// read
	product, err := gorm.G[Product](db).Where("id = ?", 1).First(ctx)
	products, err := gorm.G[Product](db).Where("code = ?", "taivisao").Find(ctx)
	fmt.Println("product: ", product)
	fmt.Println("products: ", products)

	// update - update product price to 420
	_, err = gorm.G[Product](db).Where("id = ?", product.ID).Update(ctx, "Price", 420)
	// update - update multiple fields
	_, err = gorm.G[Product](
		db,
	).Where("id = ?", product.ID).
		Updates(ctx, Product{Code: "lataiai", Price: 67})

	// delete - delete product
	_, err = gorm.G[Product](db).Where("id = ?", product.ID).Delete(ctx)
}

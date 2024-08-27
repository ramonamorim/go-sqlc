package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/ramonamorim/go-sqlc/internal/db"
	infrastructure "github.com/ramonamorim/go-sqlc/internal/infrastructure/config"
)

func main() {
	env := infrastructure.LoadEnvConfiguration()
	ctx := context.Background()

	dbConn, err := sql.Open(env.Database, env.DbConnection)
	if err != nil {
		panic(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	err = queries.CreateCategory(ctx, db.CreateCategoryParams{
		ID:          uuid.New().String(),
		Name:        "backend",
		Description: sql.NullString{String: "Backend development", Valid: true}},
	)

	if err != nil {
		log.Fatal(err)
	}

	err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		ID:          "3ca36781-5c24-4932-af21-df2d2f5614b5",
		Name:        "backend updated",
		Description: sql.NullString{String: "Backend development updated", Valid: true},
	})
	if err != nil {
		log.Fatal(err)
	}

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, category := range categories {
		log.Println(category)
	}

	err = queries.DeleteCategory(ctx, "3ca36781-5c24-4932-af21-df2d2f5614b5")
	if err != nil {
		log.Fatal(err)
	}
}

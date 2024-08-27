package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/ramonamorim/go-sqlc/internal/db"
	infrastructure "github.com/ramonamorim/go-sqlc/internal/infrastructure/config"
)

type CourseDB struct {
	dbConn *sql.DB
	*db.Queries
}

func NewCourseDB(dbConn *sql.DB) *CourseDB {
	return &CourseDB{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

type CourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
	Price       float64
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

func (c *CourseDB) callTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := db.New(tx)
	err = fn(q)
	if err != nil {
		if errRb := tx.Rollback(); errRb != nil {
			return fmt.Errorf("error on rollback: %v, orifinal error:: %v", errRb, err)
		}
		return err
	}
	return tx.Commit()
}

func (c CourseDB) CreateCourseAndCategory(ctx context.Context, argsCategory CategoryParams, args CourseParams) error {
	err := c.callTx(ctx, func(q *db.Queries) error {

		err := q.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          argsCategory.ID,
			Name:        argsCategory.Name,
			Description: argsCategory.Description,
		})
		if err != nil {
			return err
		}
		err = q.CreateCourse(ctx, db.CreateCourseParams{
			ID:          args.ID,
			Name:        args.Name,
			Description: args.Description,
			CategoryID:  argsCategory.ID,
			Price:       args.Price,
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

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

	courseArgs := CourseParams{
		ID:          uuid.New().String(),
		Name:        "Programming",
		Description: sql.NullString{String: "Learn some language", Valid: true},
		Price:       100.99,
	}

	categoryParams := CategoryParams{
		ID:          uuid.New().String(),
		Name:        "Go",
		Description: sql.NullString{String: "Go programming language", Valid: true},
	}

	courseDB := NewCourseDB(dbConn)

	err = courseDB.CreateCourseAndCategory(ctx, categoryParams, courseArgs)

	if err != nil {
		log.Fatal(err)
	}

	queries := db.New(dbConn)

	courses, err := queries.ListCourses(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, course := range courses {
		fmt.Printf("Category: %s, Course ID: %s, Course Name: %s, Course Description: %s, Course Price: %f \n",
			course.CategoryName, course.ID, course.Name, course.Description.String, course.Price)
	}
}

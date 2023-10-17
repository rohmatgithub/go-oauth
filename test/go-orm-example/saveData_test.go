package go_orm_example

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	go_orm_example "go-oauth/go-orm-example"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"testing"
	"time"
)

func TestSingleSave(t *testing.T) {
	db, err := go_orm_example.ConnectDB()
	if err != nil {
		t.Fatal(err)
	}

	categories := go_orm_example.Categories{
		Code: "code1",
		Name: "category name",
	}

	//result := db.Find()
	result := db.Create(&categories)
	if result.Error != nil {
		t.Fatal(err)
	}
	assert.Equal(t, true, categories.ID > 0)
	// returns inserted data's primary key
}

func TestTransaction(t *testing.T) {
	db, err := go_orm_example.ConnectDB()
	if err != nil {
		t.Fatal(err)
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(&go_orm_example.Categories{Code: "DRINK", Name: "Category Drink"}).Error; err != nil {
			return err
		}

		if err = tx.Create(&go_orm_example.Categories{Code: "FOOD", Name: "Category Food"}).Error; err != nil {
			return err
		}
		return nil
	})
}

func TestUpsert(t *testing.T) {
	db, err := go_orm_example.ConnectDB()
	if err != nil {
		t.Fatal(err)
	}

	categories := go_orm_example.Categories{
		Code:      "SNACK",
		Name:      "Category Snack",
		UpdatedAt: time.Now(),
	}
	var repo go_orm_example.Categories
	err = db.Transaction(func(tx *gorm.DB) error {
		errDB := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("code = ?", categories.Code).Find(&repo).Error
		if errDB != nil {
			return errDB
		}

		categories.ID = repo.ID
		errDB = tx.Save(&categories).Error
		if errDB != nil {
			return errDB
		}

		return nil
	})
	assert.NotZero(t, categories.ID)
	assert.Equal(t, true, categories.ID > 0)
}

func TestInsert(t *testing.T) {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "root"
		dbname   = "db_belajar_golang"
	)
	//psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	//	"password=%s dbname=%s sslmode=disable",
	//	host, port, user, password, dbname)
	//
	//db, err := sql.Open("postgres", psqlInfo)
	//if err != nil {
	//	t.Fatal(err)
	//}

	gormDB, err := go_orm_example.ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	query := "UPDATE categories SET name = $1, updated_at = $2 WHERE id = $3"

	fmt.Println(time.Now())

	db, err := gormDB.DB()
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec(query, []interface{}{"Category Edit", time.Now(), 10}...)
	if err != nil {
		t.Fatal(err)
	}

}

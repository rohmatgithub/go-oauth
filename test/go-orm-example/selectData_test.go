package go_orm_example

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	go_orm_example "go-oauth/go-orm-example"
	"gorm.io/gorm/clause"
	"testing"
)

func TestSelect(t *testing.T) {
	db, err := go_orm_example.ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	categories := go_orm_example.Categories{
		Code: "code1",
		Name: "category name",
	}

	var repo go_orm_example.Categories
	_ = db.
		//Where("code = $1", categories.Code).
		Raw("SELECT code, name FROM categories WHERE code = $1", categories.Code).Row().Scan(
		&repo.Code, &repo.Name)
	//Find(&repo)
	//if result.Error != nil {
	//	t.Fatal(err)
	//}

	fmt.Println(repo)
	assert.Equal(t, true, repo.ID > 0)
}

func TestSelectWithStruct(t *testing.T) {
	db, err := go_orm_example.ConnectDB()
	if err != nil {
		t.Fatal(err)
	}

	var repo go_orm_example.Categories
	result := db.Where(&go_orm_example.Categories{Code: "code1"}).Find(&repo)
	if result.Error != nil {
		t.Fatal(err)
	}

	assert.Equal(t, true, repo.ID > 0)
}

func TestSelectForUpdate(t *testing.T) {
	db, err := go_orm_example.ConnectDB()
	if err != nil {
		t.Fatal(err)
	}

	tx := db.Begin()
	var repo go_orm_example.Categories
	result := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("code = ?", "code1").Find(&repo)
	if result.Error != nil {
		t.Fatal(err)
	}
	txErr := tx.Commit()
	if tx.Error != nil {
		t.Fatal(txErr)
	}

	assert.Equal(t, true, repo.ID > 0)
}

func TestListRows(t *testing.T) {

}

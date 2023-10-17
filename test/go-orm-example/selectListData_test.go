package go_orm_example

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	go_orm_example "go-oauth/go-orm-example"
	"gorm.io/gorm"
	"testing"
)

func TestSelectListToSql(t *testing.T) {
	db, err := go_orm_example.ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	var resultList []go_orm_example.Categories
	sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Model(&resultList).
			//Where("id = ?", 100).
			Limit(10).Order("name desc").
			Find(&resultList)
	})

	assert.Equal(t, "SELECT * FROM \"categories\" ORDER BY name desc LIMIT 10", sql)
}

func TestSelectListRow1(t *testing.T) {
	db, err := go_orm_example.ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	var categories go_orm_example.Categories
	row := db.Table("categories").
		Where("id = ?", 10).
		Select("code", "name").Row()
	err = row.Scan(&categories.Code, &categories.Name)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "SNACK", categories.Code)
	assert.Equal(t, "Category Edit", categories.Name)
}

func TestSelectListRow2(t *testing.T) {
	db, err := go_orm_example.ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	var categories go_orm_example.Categories
	row := db.Raw("select code, name from categories where id = ?", 10).Row()
	err = row.Scan(&categories.Code, &categories.Name)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "SNACK", categories.Code)
	assert.Equal(t, "Category Edit", categories.Name)
}

func TestSelectListRows(t *testing.T) {
	db, err := go_orm_example.ConnectDB()
	rows, err := db.Model(&go_orm_example.Categories{}).
		//Where("name = ?", "jinzhu").
		Select("id, code, name").
		Order("name asc").Rows()
	defer rows.Close()
	for rows.Next() {
		var temp go_orm_example.Categories
		err = rows.Scan(&temp.ID, &temp.Code, &temp.Name)
		if err != nil {
			t.Fatal(err)
		}
		// do something
		fmt.Printf("id : %d, code : %s, name : %s\n", temp.ID, temp.Code, temp.Name)
	}
}

func TestSelectListWithScanRows(t *testing.T) {
	db, err := go_orm_example.ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	rows, err := db.Model(&go_orm_example.Categories{}).
		//Where("name = ?", "jinzhu").
		Select("*").
		Order("code asc").
		Rows() // (*sql.Rows, error)
	defer rows.Close()

	var listCat []go_orm_example.Categories
	var categories go_orm_example.Categories
	for rows.Next() {
		// ScanRows scan a row into user
		err = db.ScanRows(rows, &categories)

		// do something
		listCat = append(listCat, categories)
	}
	for _, temp := range listCat {
		fmt.Printf("id : %d, code : %s, name : %s\n", temp.ID, temp.Code, temp.Name)
	}
}

package go_orm_example

import (
	go_orm_example "go-oauth/go-orm-example"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectDB(t *testing.T) {
	_, err := go_orm_example.ConnectDB()
	assert.Equal(t, nil, err, "connect to db failed")
}

func TestMigrateDB(t *testing.T) {
	db, err := go_orm_example.ConnectDB()
	if err != nil {
		t.Fatal(err)
	}

	err = db.AutoMigrate(go_orm_example.PersonProfile{}, go_orm_example.User{}, go_orm_example.Product{},
		go_orm_example.Order{}, go_orm_example.OrderItem{})
	if err != nil {
		t.Fatal(err)
	}
}

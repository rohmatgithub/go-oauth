package go_orm_example

import (
	"fmt"
	go_orm_example "go-oauth/go-orm-example"
	"gorm.io/gorm"
	"testing"
)

func TestSavePersonProfileUser(t *testing.T) {
	pp := go_orm_example.PersonProfile{
		FirstName:   "David",
		LastName:    "Doe",
		Address:     "Jakarta",
		PhoneNumber: "081199999121",
	}

	user := go_orm_example.User{
		Username:        "david_doe",
		PersonProfileID: 0,
	}

	db, err := go_orm_example.ConnectDB()
	if err != nil {
		t.Fatal(err)
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		err = tx.Save(&pp).Error
		if err != nil {
			return err
		}

		user.PersonProfileID = pp.ID

		err = tx.Save(&user).Error
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		t.Fatal(err)
	}
}

func TestSelectForeignKey(t *testing.T) {
	db, err := go_orm_example.ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	var listUsers []go_orm_example.User
	err = db.Model(&go_orm_example.User{}).Preload("PersonProfile").Find(&listUsers).Error
	if err != nil {
		t.Fatal(err)
	}
	for _, user := range listUsers {
		fmt.Printf("id : %d, username : %s, firstName : %s, lastName : %s, phone : %s\n",
			user.ID, user.Username, user.PersonProfile.FirstName, user.PersonProfile.LastName, user.PersonProfile.PhoneNumber)
	}
	//defer rows.Close()

	//var listUser []go_orm_example.User
	//var user go_orm_example.User
	//for rows.Next() {
	//	err = db.ScanRows(rows, &user)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//	fmt.Printf("id : %d, username : %s, firstName : %s, lastName : %s, phone : %s\n",
	//		user.ID, user.Username, user.PersonProfile.FirstName, user.PersonProfile.LastName, user.PersonProfile.PhoneNumber)
	//	//listUser = append(listUser, user)
	//}

}

package go_orm_example

import (
	"context"
	"fmt"
	go_orm_example "go-oauth/go-orm-example"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"testing"
	"time"
)

func prepareData() error {
	db, err := go_orm_example.ConnectDB()
	if err != nil {
		return err
	}

	products := []go_orm_example.Product{
		{
			Code: "PRO-2",
			Name: "Product Dua",
		},
		{
			Code: "PRO-3",
			Name: "Product Tiga",
		},
		{
			Code: "PRO-4",
			Name: "Product Empat",
		},
	}

	order := go_orm_example.Order{
		OrderNumber: "ORDER-001",
		OrderDate:   time.Now(),
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		err = tx.Save(&products).Error
		if err != nil {
			return err
		}

		err = tx.Save(&order).Error
		if err != nil {
			return err
		}

		var orderItems []go_orm_example.OrderItem
		for i, pro := range products {
			orderItems = append(orderItems,
				go_orm_example.OrderItem{
					OrderID:   order.ID,
					ProductID: pro.ID,
					NetAmount: 500000 + float64(i)*50000,
				})
		}

		err = tx.Save(&orderItems).Error
		if err != nil {
			return err
		}
		return nil
	})

	return err
}

func TestOneToMany(t *testing.T) {
	//err := prepareData()
	//if err != nil {
	//	t.Fatal(err)
	//}

	db, err := go_orm_example.ConnectDB()
	if err != nil {
		t.Fatal(err)
	}

	var orders []go_orm_example.Order
	err = db.Preload("OrderItem.Product").
		Preload(clause.Associations).
		Find(&orders).Error
	if err != nil {
		t.Fatal(err)
	}
	//err = db.Model(&go_orm_example.Order{}).Preload("Product", "OrderItem").Find(&orders).Error

	for _, order := range orders {
		fmt.Printf("orderNumber : %s, orderDate : %s\n", order.OrderNumber, order.OrderDate.Format("2006-01-02"))
		for _, item := range order.OrderItem {
			fmt.Printf("\tproductCode : %s, productName : %s, netAmount : %.f\n",
				item.Product.Code, item.Product.Name, item.NetAmount)
		}
	}
}

func TestContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Menggunakan goroutine untuk melakukan operasi yang membutuhkan waktu lama
	go func() {
		select {
		case <-time.After(3 * time.Second):
			fmt.Println("Operasi selesai")
		case <-ctx.Done():
			fmt.Println("Operasi dibatalkan:", ctx.Err())
		}
	}()

	// Menunggu hingga operasi selesai atau dibatalkan
	select {
	case <-ctx.Done():
		fmt.Println("Menunggu selesai:", ctx.Err())
	}
}

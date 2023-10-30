package coba_coba

import (
	"fmt"
	"go-oauth/util"
	"testing"
)

type AbstractInterface interface {
	Func1() string
	//Func2() string
	//Func3() string
}

type struct1 struct {
	AbstractInterface
}

type struct2 struct {
	AbstractInterface
}

func (a struct1) Func1() string {
	return "Struct 1"
}

func (a struct2) Func1() string {
	return "Struct 2"
}

func TestPrintInterface(t *testing.T) {
	//a := struct1{}
	//fmt.Println(a.Func1())
	fmt.Println(util.GetUUID())
}

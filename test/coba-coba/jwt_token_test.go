package coba_coba

import (
	"fmt"
	"go-oauth/model"
	"testing"
)

func TestParsingJwtToken(t *testing.T) {
	tokenStr := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJzY29wZSI6IiIsImxvY2FsZSI6ImVuLVVTIiwiY2xpZW50X2lkIjoiZjIxOTQyZGJmNzk5NDhhOGJlYTllMGI3NmQ0ZWJiNmYiLCJpc3MiOiJhdXRoIiwiZXhwIjoxNjk4NjAyNzExLCJpYXQiOjE2OTg1NTk1MTF9.Re8JW9XxnLOUGCxMpsr6wrsgIEQIrAkkWVs6c3a1TTiaqGHHeAS-lEXUwgkVY0Pee5mUJZQLwjka4jeGofB-kg"

	result, err := model.JWTToken{}.ParsingJwtTokenInternal(tokenStr)
	if err.Error != nil {
		t.Fatal(err)
	}
	fmt.Println(result.UserID)
}

package main

import (
	"fmt"
	"testing"
)

func TestToken(t *testing.T) {
	token, err := GetToken()

	if err != nil {
		t.Error(err)
	}

	fmt.Println(token)
}

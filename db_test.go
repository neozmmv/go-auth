package main

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/neozmmv/go-auth/utils"
)

func TestConnectToDb(t *testing.T) {
	godotenv.Load()
	t.Log("Connecting to database...")
	utils.ConnectDatabase()
	if utils.DB == nil {
		t.Fatal("Failed to connect to database")
	}
}

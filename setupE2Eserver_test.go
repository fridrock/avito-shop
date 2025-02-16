package main

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/fridrock/avito-shop/testdbsetup"
)

var testServer *httptest.Server

func TestMain(m *testing.M) {
	os.Setenv("SECRET_KEY", "SECRET_FOR_TEST")
	os.Setenv("PORT", "8080")
	conn := testdbsetup.CreateTestConnection(".")
	defer conn.Close()
	router := setupRouter(conn)
	testServer = httptest.NewServer(router)
	defer testServer.Close()
	m.Run()
}

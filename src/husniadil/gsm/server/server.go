package server

import (
	"husniadil/gsm/handler"
	"husniadil/gsm/util/logger"
	"net/http"

	"github.com/justinas/alice"
)

func StartServer() {
	handlers := alice.New(
		handler.LoggingHandler,
	)

	router := createRouter()

	logger.Info("%s", "GSMArena API Server starting: http://localhost:8080")

	logger.Fatal(http.ListenAndServe(":8080", handlers.Then(router)))
}

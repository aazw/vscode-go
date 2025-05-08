package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aazw/vscode-go/playgrounds/openapi_codegen/pkg/api"
	"github.com/aazw/vscode-go/playgrounds/openapi_codegen/pkg/api/openapi"
)

func main() {
	var (
		strict bool
	)
	flag.BoolVar(&strict, "strict", false, "run strict server flag")
	flag.Parse()

	switch {
	case strict:
		fmt.Println("run strict server")
		runStrictServer()
	default:
		fmt.Println("run gin server")
		runGinServer()
	}
}

func runGinServer() {
	handler := api.NewGinServerHandler()

	r := gin.Default()

	openapi.RegisterHandlers(r, handler)

	s := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8080",
	}
	log.Fatal(s.ListenAndServe())
}

func runStrictServer() {
	ssii := api.NewStrictServerInterfaceImpl()
	handler := openapi.NewStrictHandler(ssii, nil)

	r := gin.Default()

	openapi.RegisterHandlers(r, handler)

	s := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8080",
	}
	log.Fatal(s.ListenAndServe())
}

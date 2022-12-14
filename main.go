package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"flag"

	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/database"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/database/schema"
)

var (
	port *string
	URLs []string
)

func main() {
	initEnv()
	initFlags()

	database.TestConnection()

	router := mux.NewRouter()

	routes.SetMainHandleRoutes(router, &URLs)

	prefix := "/testupload/"
	router.PathPrefix(prefix).Handler(
		http.StripPrefix(prefix, http.FileServer(http.Dir("./client/testUploadFile"))),
	)

	http.ListenAndServe(":"+*port, router)
}

func initEnv() {
	err := godotenv.Load("go.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func initFlags() {
	database_to_generate := flag.String("dbgen", "", "Database name, -gen=true")
	usergen := flag.String("usergen", "postgres", "Database user, -gen=true")

	port = flag.String("port", "8080", "Port to use")

	flag.Parse()

	if strings.Compare(*database_to_generate, "") != 0 {
		schema.Generate(*database_to_generate, *usergen)
		os.Exit(0)
	}
}

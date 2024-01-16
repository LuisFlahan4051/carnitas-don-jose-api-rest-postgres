package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"flag"

	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/database"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/routes"
	"github.com/rs/cors"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/database/schema"
)

var (
	port *string
	URLs []string
)

/*
First usage: set into browser: http://localhost:8080/users?admin_username=root&admin_password=root
*/
func main() {
	initEnv()
	initFlags()

	database.TestConnection()

	router := mux.NewRouter()

	// ADDING ROUTES
	routes.SetMainHandleRoutes(router, &URLs)

	// SERVE STATIC FILES
	prefix := "/testupload"
	router.PathPrefix(prefix).Handler(
		http.StripPrefix(prefix, http.FileServer(http.Dir("./client/testUploadFile"))),
	)

	publicUsersFilesDir := "/users/"
	router.PathPrefix(publicUsersFilesDir).Handler(
		http.StripPrefix(publicUsersFilesDir, http.FileServer(http.Dir("./storage/public/users"))),
	)

	publicNotificationsFilesDir := "/notifications/"
	router.PathPrefix(publicNotificationsFilesDir).Handler(
		http.StripPrefix(publicNotificationsFilesDir, http.FileServer(http.Dir("./storage/public/notifications"))),
	)

	// CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
	})
	handler := c.Handler(router)
	// RUN SERVER
	http.ListenAndServe(":"+*port, handler)
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

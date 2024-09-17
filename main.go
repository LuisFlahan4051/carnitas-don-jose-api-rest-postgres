package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"flag"

	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/database"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/models"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/routes"
	"github.com/rs/cors"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/database/schema"
)

var (
	port     *string
	devTools *bool
	URLs     []string
)

/*
First usage: set into browser: http://localhost:8080/users?admin_username=root&admin_password=root
*/
func main() {
	initEnv()
	initFlags()

	// this change with a flag. > go run main.go dev-tools true
	if *devTools {
		displayDevTools()
		return
	}

	database.TestConnection()
	prepareServer()
}

func displayDevTools() {
	for {
		fmt.Println("Chose an option:")
		fmt.Println("1: Exit.")
		fmt.Println("2: Test connection to DB.")
		fmt.Println("3: Generate typescript interfaces with golang structs.")
		fmt.Print("> ")

		var option int
		_, err := fmt.Scan(&option)
		if err != nil {
			log.Fatal(err)
			return
		}

		switch option {
		case 1:
			fmt.Println("Goodbye!")
			return
		case 2:
			database.TestConnection()
		case 3:
			models.GenerateTypescriptFiles()
		}
	}

}

func prepareServer() {
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
	// GENERATE DATABASE, just for a local instalation of postgres.
	// First use usergen, then dbgen.
	usergen := flag.String("usergen", "postgres", "Database user (local instalation of postgres, no docker), -gen=true")
	database_to_generate := flag.String("dbgen", "", "Database name (local instalation of postgres, no docker), -gen=true")

	// PORT, change the port of the api
	port = flag.String("port", "8080", "Port to use")

	devTools = flag.Bool("dev-tools", false, "Use this option for display an options menu")

	flag.Parse()

	if strings.Compare(*database_to_generate, "") != 0 {
		schema.Generate(*database_to_generate, *usergen)
		os.Exit(0)
	}
}

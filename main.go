package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"flag"

	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/database/schema"
)

var (
	port *string
)

func main() {
	initEnv()
	initFlags()

	router := mux.NewRouter()

	router.HandleFunc("/{id}", routes.HomeHandler)
	router.HandleFunc("/crud/foodtype", routes.PostFoodType).Methods("POST")
	router.HandleFunc("/crud/foodtypes", routes.GetFoodTypes).Methods("GET")
	router.HandleFunc("/crud/foodtype/{id}", routes.GetFoodType).Methods("GET")
	router.HandleFunc("/crud/foodtype/{id}", routes.PatchFoodType).Methods("PATCH")
	router.HandleFunc("/root/foodtype/{id}", routes.RootPatchFoodType).Methods("PATCH")
	router.HandleFunc("/crud/foodtype/{id}", routes.DeleteFoodType).Methods("DELETE")
	router.HandleFunc("/root/foodtype/{id}", routes.RootDeleteFoodType).Methods("DELETE")

	router.HandleFunc("/crud/branch", routes.PostBranch).Methods("POST")
	router.HandleFunc("/crud/branches", routes.GetBranches).Methods("GET")
	router.HandleFunc("/crud/branch/{id}", routes.GetBranch).Methods("GET")
	router.HandleFunc("/crud/branch/{id}", routes.DeleteBranch).Methods("DELETE")
	router.HandleFunc("/crud/branch/{id}", routes.PatchBranch).Methods("PATCH")

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

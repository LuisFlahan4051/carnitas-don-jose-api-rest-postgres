package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/database"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/models"
	"github.com/gorilla/mux"
)

func foodTypeIsDeleted(id uint) bool {
	db := database.Connect()

	deleted := "SELECT deleted_at FROM food_types WHERE id = $1"

	var deletedAt *time.Time

	db.QueryRow(deleted, id).Scan(&deletedAt)

	db.Close()

	log.Println("Deleted at: ", deletedAt)

	return deletedAt != nil
}

func getFoodType(id uint) models.FoodType {
	var foodType models.FoodType

	db := database.Connect()

	query := "SELECT id, created_at, updated_at, deleted_at, name FROM food_types WHERE id = $1"
	err := db.QueryRow(query, id).Scan(
		&foodType.Id,
		&foodType.CreatedAt,
		&foodType.UpdatedAt,
		&foodType.DeletedAt,
		&foodType.Name)
	if err != nil {
		log.Println(err)
	}

	queryAsossiated := "SELECT " +
		"id," +
		" created_at," +
		" updated_at," +
		" deleted_at," +
		" name," +
		" description," +
		" photo," +
		" food_type_id," +
		" food_meat_id" +
		" FROM foods WHERE food_type_id = $1"
	rows, err := db.Query(queryAsossiated, id)
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		var food models.Food
		err = rows.Scan(
			&food.Id,
			&food.CreatedAt,
			&food.UpdatedAt,
			&food.DeletedAt,
			&food.Name,
			&food.Description,
			&food.Photo,
			&food.FoodTypeID,
			&food.FoodMeatID)
		if err != nil {
			log.Println(err)
		}
		foodType.Foods = append(foodType.Foods, food)
	}

	db.Close()
	return foodType
}

//--------------

func PostFoodType(writer http.ResponseWriter, request *http.Request) {
	log.Println("Post Food type")

	var foodType models.FoodType
	err := json.NewDecoder(request.Body).Decode(&foodType)
	logcatch(writer, http.StatusBadRequest, err)

	db := database.Connect()

	query := "INSERT INTO food_types (name, created_at, updated_at) VALUES ($1,$2,$3) RETURNING id"

	date := time.Now()
	foodType.CreatedAt = date
	foodType.UpdatedAt = date
	foodType.DeletedAt = nil

	err = db.QueryRow(query, foodType.Name, foodType.CreatedAt, foodType.UpdatedAt).Scan(&foodType.Id)

	logcatch(writer, http.StatusNotModified, err)

	db.Close()

	json.NewEncoder(writer).Encode(&foodType)
}

func GetFoodType(writer http.ResponseWriter, request *http.Request) {
	log.Println("Get Food type")

	params := mux.Vars(request)
	param, err := strconv.Atoi(params["id"])
	if err != nil || param < 1 {
		logcatch(writer, http.StatusBadRequest, err)
		return
	}

	if foodTypeIsDeleted(uint(param)) {
		log.Println("Error: Item deleted")
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Error: Item deleted"))
		return
	}

	var foodType models.FoodType

	db := database.Connect()

	query := "SELECT id, created_at, updated_at, deleted_at, name FROM food_types WHERE id = $1"
	err = db.QueryRow(query, params["id"]).Scan(
		&foodType.Id,
		&foodType.CreatedAt,
		&foodType.UpdatedAt,
		&foodType.DeletedAt,
		&foodType.Name)
	logcatch(writer, http.StatusBadRequest, err)

	queryAsossiated := "SELECT " +
		"id," +
		" created_at," +
		" updated_at," +
		" deleted_at," +
		" name," +
		" description," +
		" photo," +
		" food_type_id," +
		" food_meat_id" +
		" FROM foods WHERE food_type_id = $1"
	rows, err := db.Query(queryAsossiated, params["id"])
	logcatch(writer, http.StatusBadRequest, err)

	for rows.Next() {
		var food models.Food
		err = rows.Scan(
			&food.Id,
			&food.CreatedAt,
			&food.UpdatedAt,
			&food.DeletedAt,
			&food.Name,
			&food.Description,
			&food.Photo,
			&food.FoodTypeID,
			&food.FoodMeatID)
		logcatch(writer, http.StatusBadRequest, err)
		foodType.Foods = append(foodType.Foods, food)
	}

	db.Close()

	json.NewEncoder(writer).Encode(&foodType)
}

func RootGetFoodType(writer http.ResponseWriter, request *http.Request) {
	log.Println("Get Food type ROOT")

	params := mux.Vars(request)
	param, err := strconv.Atoi(params["id"])
	if err != nil || param < 1 {
		logcatch(writer, http.StatusBadRequest, err)
		return
	}

	var foodTypes []models.FoodType

	db := database.Connect()

	query := "SELECT id, created_at, updated_at, deleted_at, name FROM food_types"
	rows, err := db.Query(query)
	logcatch(writer, http.StatusBadRequest, err)

	for rows.Next() {
		var foodType models.FoodType
		err = rows.Scan(
			&foodType.Id,
			&foodType.CreatedAt,
			&foodType.UpdatedAt,
			&foodType.DeletedAt,
			&foodType.Name)
		logcatch(writer, http.StatusBadRequest, err)
		foodTypes = append(foodTypes, foodType)
	}

	db.Close()
	json.NewEncoder(writer).Encode(&foodTypes)
}

func GetFoodTypes(writer http.ResponseWriter, request *http.Request) {
	log.Println("Get Food types")

	var foodTypes []models.FoodType

	db := database.Connect()

	query := "SELECT id, created_at, updated_at, deleted_at, name FROM food_types"
	rows, err := db.Query(query)
	logcatch(writer, http.StatusBadRequest, err)

	for rows.Next() {
		var foodType models.FoodType
		err = rows.Scan(
			&foodType.Id,
			&foodType.CreatedAt,
			&foodType.UpdatedAt,
			&foodType.DeletedAt,
			&foodType.Name)
		logcatch(writer, http.StatusBadRequest, err)
		if foodType.DeletedAt == nil {
			foodTypes = append(foodTypes, foodType)
		}
	}

	queryAsossiated := "SELECT " +
		"id," +
		" created_at," +
		" updated_at," +
		" deleted_at," +
		" name," +
		" description," +
		" photo," +
		" food_type_id," +
		" food_meat_id" +
		" FROM foods"
	rows, err = db.Query(queryAsossiated)
	logcatch(writer, http.StatusBadRequest, err)

	for rows.Next() {
		var food models.Food
		err = rows.Scan(
			&food.Id,
			&food.CreatedAt,
			&food.UpdatedAt,
			&food.DeletedAt,
			&food.Name,
			&food.Description,
			&food.Photo,
			&food.FoodTypeID,
			&food.FoodMeatID)
		logcatch(writer, http.StatusBadRequest, err)
		for i := 0; i < len(foodTypes); i++ {
			if food.FoodTypeID == foodTypes[i].Id {
				foodTypes[i].Foods = append(foodTypes[i].Foods, food)
			}
		}
	}

	db.Close()

	json.NewEncoder(writer).Encode(&foodTypes)
}

func RootGetFoodTypes(writer http.ResponseWriter, request *http.Request) {
	log.Println("Get Food types ROOT")

	var foodTypes []models.FoodType

	db := database.Connect()

	query := "SELECT id, created_at, updated_at, deleted_at, name FROM food_types"
	rows, err := db.Query(query)
	logcatch(writer, http.StatusBadRequest, err)

	for rows.Next() {
		var foodType models.FoodType
		err = rows.Scan(
			&foodType.Id,
			&foodType.CreatedAt,
			&foodType.UpdatedAt,
			&foodType.DeletedAt,
			&foodType.Name)
		logcatch(writer, http.StatusBadRequest, err)
		foodTypes = append(foodTypes, foodType)
	}

	queryAsossiated := "SELECT " +
		"id," +
		" created_at," +
		" updated_at," +
		" deleted_at," +
		" name," +
		" description," +
		" photo," +
		" food_type_id," +
		" food_meat_id" +
		" FROM foods"
	rows, err = db.Query(queryAsossiated)
	logcatch(writer, http.StatusBadRequest, err)

	for rows.Next() {
		var food models.Food
		err = rows.Scan(
			&food.Id,
			&food.CreatedAt,
			&food.UpdatedAt,
			&food.DeletedAt,
			&food.Name,
			&food.Description,
			&food.Photo,
			&food.FoodTypeID,
			&food.FoodMeatID)
		logcatch(writer, http.StatusBadRequest, err)
		for i := 0; i < len(foodTypes); i++ {
			if food.FoodTypeID == foodTypes[i].Id {
				foodTypes[i].Foods = append(foodTypes[i].Foods, food)
			}
		}
	}

	db.Close()

	json.NewEncoder(writer).Encode(&foodTypes)
}

func PatchFoodType(writer http.ResponseWriter, request *http.Request) {
	log.Println("Update Food type")

	params := mux.Vars(request)

	param, err := strconv.Atoi(params["id"])
	if err != nil || param < 1 {
		logcatch(writer, http.StatusBadRequest, err)
		return
	}

	if foodTypeIsDeleted(uint(param)) {
		log.Println("Error: Status deleted")
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Error: Status deleted"))
		return
	}

	var foodType models.FoodType
	err = json.NewDecoder(request.Body).Decode(&foodType)
	logcatch(writer, http.StatusBadRequest, err)

	db := database.Connect()

	deleted := "SELECT deleted_at FROM food_types WHERE id = $1"
	db.QueryRow(deleted, param).Scan(&foodType.DeletedAt)

	if foodType.DeletedAt != nil {
		log.Println("Error: Deleted status")
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Error: Deleted status"))
		return
	}

	update := "UPDATE food_types SET name = $1, updated_at = $2 WHERE id = $3"
	foodType.UpdatedAt = time.Now()

	_, err = db.Exec(update, foodType.Name, foodType.UpdatedAt, params["id"])
	logcatch(writer, http.StatusBadRequest, err)
	db.Close()

	foodTypeUpdated := getFoodType(uint(param))

	json.NewEncoder(writer).Encode(&foodTypeUpdated)
}

func RootPatchFoodType(writer http.ResponseWriter, request *http.Request) {
	log.Println("Update Food type ROOT")

	params := mux.Vars(request)

	param, err := strconv.Atoi(params["id"])
	if err != nil || param < 1 {
		logcatch(writer, http.StatusBadRequest, err)
		return
	}

	var foodType models.FoodType
	err = json.NewDecoder(request.Body).Decode(&foodType)
	logcatch(writer, http.StatusBadRequest, err)

	db := database.Connect()

	update := "UPDATE food_types SET " +
		"name = $1," +
		" updated_at = $2," +
		" created_at = $3," +
		" deleted_at = $4" +
		" WHERE id = $5"
	_, err = db.Exec(update,
		foodType.Name,
		foodType.UpdatedAt,
		foodType.CreatedAt,
		foodType.DeletedAt,
		foodType.Id)
	logcatch(writer, http.StatusBadRequest, err)
	db.Close()

	foodTypeUpdated := getFoodType(foodType.Id)

	json.NewEncoder(writer).Encode(&foodTypeUpdated)
}

func DeleteFoodType(writer http.ResponseWriter, request *http.Request) {
	log.Println("Delete Food type")

	params := mux.Vars(request)

	param, err := strconv.Atoi(params["id"])
	if err != nil || param < 1 {
		logcatch(writer, http.StatusBadRequest, err)
		return
	}

	db := database.Connect()

	if foodTypeIsDeleted(uint(param)) {
		log.Println("Error: Allready deleted")
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Error: Allready deleted"))
		return
	}

	update := "UPDATE food_types SET deleted_at = $1 WHERE id = $2"

	_, err = db.Exec(update, time.Now(), params["id"])
	logcatch(writer, http.StatusBadRequest, err)
	db.Close()

	foodType := getFoodType(uint(param))

	json.NewEncoder(writer).Encode(&foodType)
}

func RootDeleteFoodType(writer http.ResponseWriter, request *http.Request) {
	log.Println("Delete Food type ROOT")

	params := mux.Vars(request)

	param, err := strconv.Atoi(params["id"])
	if err != nil || param < 1 {
		logcatch(writer, http.StatusBadRequest, err)
		return
	}

	db := database.Connect()

	deleted := "DELETE FROM food_types WHERE id = $1"
	_, err = db.Exec(deleted, params["id"])
	logcatch(writer, http.StatusBadRequest, err)

	db.Close()

	log.Println("Delete success!")
	writer.Write([]byte("Delete success!"))
}

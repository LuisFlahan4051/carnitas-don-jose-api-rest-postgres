package crud

import (
	"errors"
	"fmt"
	"strings"
	"time"

	commons "github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/commonFunctions"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/database"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/models"
)

/*func GetBranchSafeboxesTEST(tableName string, root bool, relationalIDs *map[string]uint) ([]models.BranchSafebox, error) {
	db := database.Connect()
	defer db.Close()
	var model models.BranchSafebox
	var data []models.BranchSafebox

	query, _, err := commons.GetQuery(tableName, model, "SELECT", false)
	if err != nil {
		return nil, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	if !root {
		query += " WHERE deleted_at IS NULL"
	}

	if relationalIDs != nil {
		switch root {
		case true:
			query += " WHERE "
		case false:
			query += " AND "
		}

		var relationConditions []string
		for key, value := range *relationalIDs {
			relationConditions = append(relationConditions, fmt.Sprintf("%s = %d", key, value))
		}
		query += strings.Join(relationConditions, " AND ")
	}

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	defer rows.Close()

	//Scan
	jsonStrings, err := commons.DecodeRowsToJson(rows)
	for _, jsonString := range jsonStrings {
		json.Unmarshal([]byte(jsonString), &model)
		data = append(data, model)
	}
	//----

	if err != nil {
		return nil, fmt.Errorf("can't decode the %s query ERROR: %s", tableName, err.Error())
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}

	return data, nil
}*/

func NewFoodType(foodType *models.FoodType) error {
	db := database.Connect()
	defer db.Close()

	currentDate := time.Now()
	foodType.Id = 0
	foodType.CreatedAt = &currentDate
	foodType.UpdatedAt = &currentDate

	tableName := "food_types"
	query, data, err := commons.GetQuery(tableName, *foodType, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&foodType.Id,
		&foodType.CreatedAt,
		&foodType.UpdatedAt,
		&foodType.DeletedAt,
		&foodType.Name,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if foodType.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetFoodTypes(root bool) ([]models.FoodType, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "food_types"
	query, _, err := commons.GetQuery(tableName, models.FoodType{}, "SELECT", false)
	if err != nil {
		return nil, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	if !root {
		query += " WHERE deleted_at IS NULL"
	}

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	defer rows.Close()

	var foodTypes []models.FoodType
	for rows.Next() {
		var foodType models.FoodType
		err = rows.Scan(
			&foodType.Id,
			&foodType.CreatedAt,
			&foodType.UpdatedAt,
			&foodType.DeletedAt,
			&foodType.Name,
		)
		if err != nil {
			return nil, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
		}

		foodTypes = append(foodTypes, foodType)
	}

	if len(foodTypes) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}

	return foodTypes, nil
}

func GetFoodType(id uint, root bool) (models.FoodType, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "food_types"
	foodType := models.FoodType{}
	query, _, err := commons.GetQuery(tableName, foodType, "SELECT", false)
	if err != nil {
		return foodType, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}
	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	err = db.QueryRow(query, id).Scan(
		&foodType.Id,
		&foodType.CreatedAt,
		&foodType.UpdatedAt,
		&foodType.DeletedAt,
		&foodType.Name,
	)
	if err != nil {
		return foodType, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	if foodType.Id == 0 {
		return foodType, fmt.Errorf("%s not found", tableName)
	}

	return foodType, nil
}

func DeleteFoodType(id uint, root bool) error {
	return commons.DeleteFromTableById("food_types", id, root)
}

func UpdateFoodType(updatingFoodType *models.FoodType, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "food_types"
	currentDate := time.Now()

	if !root {
		updatingFoodType.UpdatedAt = &currentDate
		updatingFoodType.CreatedAt = nil // Not allowed to update this field
		updatingFoodType.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingFoodType.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingFoodType, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND deleted_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingFoodType.Id,
		&updatingFoodType.CreatedAt,
		&updatingFoodType.UpdatedAt,
		&updatingFoodType.DeletedAt,
		&updatingFoodType.Name,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingFoodType.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

//

func NewFoodMeat(foodMeat *models.FoodMeat) error {
	db := database.Connect()
	defer db.Close()

	currentDate := time.Now()
	foodMeat.Id = 0
	foodMeat.CreatedAt = &currentDate
	foodMeat.UpdatedAt = &currentDate

	tableName := "food_meats"
	query, data, err := commons.GetQuery(tableName, *foodMeat, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&foodMeat.Id,
		&foodMeat.CreatedAt,
		&foodMeat.UpdatedAt,
		&foodMeat.DeletedAt,
		&foodMeat.Name,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if foodMeat.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetFoodMeats(root bool) ([]models.FoodMeat, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "food_meats"
	query, _, err := commons.GetQuery(tableName, models.FoodMeat{}, "SELECT", false)

	if err != nil {
		return nil, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	if !root {
		query += " WHERE deleted_at IS NULL"
	}

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	defer rows.Close()

	var foodMeats []models.FoodMeat

	for rows.Next() {
		var foodMeat models.FoodMeat
		err := rows.Scan(
			&foodMeat.Id,
			&foodMeat.CreatedAt,
			&foodMeat.UpdatedAt,
			&foodMeat.DeletedAt,
			&foodMeat.Name,
		)
		if err != nil {
			return nil, fmt.Errorf("can't scan the %s query ERROR: %s", tableName, err.Error())
		}
		foodMeats = append(foodMeats, foodMeat)
	}
	if len(foodMeats) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}
	return foodMeats, nil
}

func GetFoodMeat(id uint, root bool) (models.FoodMeat, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "food_meats"
	foodMeat := models.FoodMeat{}
	query, _, err := commons.GetQuery(tableName, foodMeat, "SELECT", false)
	if err != nil {
		return foodMeat, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}
	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	err = db.QueryRow(query, id).Scan(
		&foodMeat.Id,
		&foodMeat.CreatedAt,
		&foodMeat.UpdatedAt,
		&foodMeat.DeletedAt,
		&foodMeat.Name,
	)
	if err != nil {
		return foodMeat, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	if foodMeat.Id == 0 {
		return foodMeat, fmt.Errorf("%s not found", tableName)
	}

	return foodMeat, nil
}

func DeleteFoodMeat(id uint, root bool) error {
	return commons.DeleteFromTableById("food_meats", id, root)
}

func UpdateFoodMeat(updatingFoodMeat *models.FoodMeat, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "food_meats"
	currentDate := time.Now()

	if !root {
		updatingFoodMeat.UpdatedAt = &currentDate
		updatingFoodMeat.CreatedAt = nil // Not allowed to update this field
		updatingFoodMeat.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingFoodMeat.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingFoodMeat, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND deleted_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingFoodMeat.Id,
		&updatingFoodMeat.CreatedAt,
		&updatingFoodMeat.UpdatedAt,
		&updatingFoodMeat.DeletedAt,
		&updatingFoodMeat.Name,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingFoodMeat.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

//

func NewFood(food *models.Food) error {
	db := database.Connect()
	defer db.Close()

	currentDate := time.Now()
	food.Id = 0
	food.CreatedAt = &currentDate
	food.UpdatedAt = &currentDate

	tableName := "foods"
	query, data, err := commons.GetQuery(tableName, *food, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&food.Id,
		&food.CreatedAt,
		&food.UpdatedAt,
		&food.DeletedAt,
		&food.Name,
		&food.Description,
		&food.Photo,
		&food.FoodTypeID,
		&food.FoodMeatID,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if food.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetFoods(root bool, relationalIDs *map[string]uint) ([]models.Food, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "foods"
	query, _, err := commons.GetQuery(tableName, models.Food{}, "SELECT", false)

	if err != nil {
		return nil, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	if !root {
		query += " WHERE deleted_at IS NULL"
	}

	if relationalIDs != nil {
		switch root {
		case true:
			query += " WHERE "
		case false:
			query += " AND "
		}

		var relationConditions []string
		for key, value := range *relationalIDs {
			relationConditions = append(relationConditions, fmt.Sprintf("%s = %d", key, value))
		}
		query += strings.Join(relationConditions, " AND ")
	}

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	defer rows.Close()

	var foods []models.Food

	for rows.Next() {
		var food models.Food
		err := rows.Scan(
			&food.Id,
			&food.CreatedAt,
			&food.UpdatedAt,
			&food.DeletedAt,
			&food.Name,
			&food.Description,
			&food.Photo,
			&food.FoodTypeID,
			&food.FoodMeatID,
		)
		if err != nil {
			return nil, fmt.Errorf("can't scan the %s query ERROR: %s", tableName, err.Error())
		}
		foods = append(foods, food)
	}
	if len(foods) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}
	return foods, nil
}

func GetFood(id uint, root bool) (models.Food, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "foods"
	food := models.Food{}
	query, _, err := commons.GetQuery(tableName, food, "SELECT", false)
	if err != nil {
		return food, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}
	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	err = db.QueryRow(query, id).Scan(
		&food.Id,
		&food.CreatedAt,
		&food.UpdatedAt,
		&food.DeletedAt,
		&food.Name,
		&food.Description,
		&food.Photo,
		&food.FoodTypeID,
		&food.FoodMeatID,
	)
	if err != nil {
		return food, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	if food.Id == 0 {
		return food, fmt.Errorf("%s not found", tableName)
	}

	return food, nil
}

func DeleteFood(id uint, root bool) error {
	return commons.DeleteFromTableById("foods", id, root)
}

func UpdateFood(updatingFood *models.Food, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "foods"
	currentDate := time.Now()

	if !root {
		updatingFood.UpdatedAt = &currentDate
		updatingFood.CreatedAt = nil // Not allowed to update this field
		updatingFood.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingFood.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingFood, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND deleted_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingFood.Id,
		&updatingFood.CreatedAt,
		&updatingFood.UpdatedAt,
		&updatingFood.DeletedAt,
		&updatingFood.Name,
		&updatingFood.Description,
		&updatingFood.Photo,
		&updatingFood.FoodTypeID,
		&updatingFood.FoodMeatID,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingFood.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

//

func NewDrinkSize(drinkSize *models.DrinkSize) error {
	db := database.Connect()
	defer db.Close()

	currentDate := time.Now()
	drinkSize.Id = 0
	drinkSize.CreatedAt = &currentDate
	drinkSize.UpdatedAt = &currentDate

	tableName := "drink_sizes"
	query, data, err := commons.GetQuery(tableName, *drinkSize, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&drinkSize.Id,
		&drinkSize.CreatedAt,
		&drinkSize.UpdatedAt,
		&drinkSize.DeletedAt,
		&drinkSize.Name,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if drinkSize.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetDrinkSizes(root bool) ([]models.DrinkSize, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "drink_sizes"
	query, _, err := commons.GetQuery(tableName, models.DrinkSize{}, "SELECT", false)

	if err != nil {
		return nil, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	if !root {
		query += " WHERE deleted_at IS NULL"
	}

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	defer rows.Close()

	var drinkSizes []models.DrinkSize

	for rows.Next() {
		var drinkSize models.DrinkSize
		err := rows.Scan(
			&drinkSize.Id,
			&drinkSize.CreatedAt,
			&drinkSize.UpdatedAt,
			&drinkSize.DeletedAt,
			&drinkSize.Name,
		)
		if err != nil {
			return nil, fmt.Errorf("can't scan the %s query ERROR: %s", tableName, err.Error())
		}
		drinkSizes = append(drinkSizes, drinkSize)
	}
	if len(drinkSizes) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}
	return drinkSizes, nil
}

func GetDrinkSize(id uint, root bool) (models.DrinkSize, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "drink_sizes"
	drinkSize := models.DrinkSize{}
	query, _, err := commons.GetQuery(tableName, drinkSize, "SELECT", false)
	if err != nil {
		return drinkSize, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}
	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	err = db.QueryRow(query, id).Scan(
		&drinkSize.Id,
		&drinkSize.CreatedAt,
		&drinkSize.UpdatedAt,
		&drinkSize.DeletedAt,
		&drinkSize.Name,
	)
	if err != nil {
		return drinkSize, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	if drinkSize.Id == 0 {
		return drinkSize, fmt.Errorf("%s not found", tableName)
	}

	return drinkSize, nil
}

func DeleteDrinkSize(id uint, root bool) error {
	return commons.DeleteFromTableById("drink_sizes", id, root)
}

func UpdateDrinkSize(updatingDrinkSize *models.DrinkSize, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "drink_sizes"
	currentDate := time.Now()

	if !root {
		updatingDrinkSize.UpdatedAt = &currentDate
		updatingDrinkSize.CreatedAt = nil // Not allowed to update this field
		updatingDrinkSize.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingDrinkSize.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingDrinkSize, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND deleted_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingDrinkSize.Id,
		&updatingDrinkSize.CreatedAt,
		&updatingDrinkSize.UpdatedAt,
		&updatingDrinkSize.DeletedAt,
		&updatingDrinkSize.Name,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingDrinkSize.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

//

func NewDrinkFlavor(drinkFlavor *models.DrinkFlavor) error {
	db := database.Connect()
	defer db.Close()

	currentDate := time.Now()
	drinkFlavor.Id = 0
	drinkFlavor.CreatedAt = &currentDate
	drinkFlavor.UpdatedAt = &currentDate

	tableName := "drink_flavors"
	query, data, err := commons.GetQuery(tableName, *drinkFlavor, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&drinkFlavor.Id,
		&drinkFlavor.CreatedAt,
		&drinkFlavor.UpdatedAt,
		&drinkFlavor.DeletedAt,
		&drinkFlavor.Name,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if drinkFlavor.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetDrinkFlavors(root bool) ([]models.DrinkFlavor, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "food_meats"
	query, _, err := commons.GetQuery(tableName, models.DrinkFlavor{}, "SELECT", false)

	if err != nil {
		return nil, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	if !root {
		query += " WHERE deleted_at IS NULL"
	}

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	defer rows.Close()

	var drinkFlavors []models.DrinkFlavor

	for rows.Next() {
		var drinkFlavor models.DrinkFlavor
		err := rows.Scan(
			&drinkFlavor.Id,
			&drinkFlavor.CreatedAt,
			&drinkFlavor.UpdatedAt,
			&drinkFlavor.DeletedAt,
			&drinkFlavor.Name,
		)
		if err != nil {
			return nil, fmt.Errorf("can't scan the %s query ERROR: %s", tableName, err.Error())
		}
		drinkFlavors = append(drinkFlavors, drinkFlavor)
	}
	if len(drinkFlavors) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}
	return drinkFlavors, nil
}

func GetDrinkFlavor(id uint, root bool) (models.DrinkFlavor, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "food_meats"
	drinkFlavor := models.DrinkFlavor{}
	query, _, err := commons.GetQuery(tableName, drinkFlavor, "SELECT", false)
	if err != nil {
		return drinkFlavor, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}
	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	err = db.QueryRow(query, id).Scan(
		&drinkFlavor.Id,
		&drinkFlavor.CreatedAt,
		&drinkFlavor.UpdatedAt,
		&drinkFlavor.DeletedAt,
		&drinkFlavor.Name,
	)
	if err != nil {
		return drinkFlavor, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	if drinkFlavor.Id == 0 {
		return drinkFlavor, fmt.Errorf("%s not found", tableName)
	}

	return drinkFlavor, nil
}

func DeleteDrinkFlavor(id uint, root bool) error {
	return commons.DeleteFromTableById("drink_flavors", id, root)
}

func UpdateDrinkFlavor(updatingDrinkFlavor *models.DrinkFlavor, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "drink_flavors"
	currentDate := time.Now()

	if !root {
		updatingDrinkFlavor.UpdatedAt = &currentDate
		updatingDrinkFlavor.CreatedAt = nil // Not allowed to update this field
		updatingDrinkFlavor.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingDrinkFlavor.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingDrinkFlavor, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND deleted_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingDrinkFlavor.Id,
		&updatingDrinkFlavor.CreatedAt,
		&updatingDrinkFlavor.UpdatedAt,
		&updatingDrinkFlavor.DeletedAt,
		&updatingDrinkFlavor.Name,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingDrinkFlavor.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

//

func NewDrink(drink *models.Drink) error {
	db := database.Connect()
	defer db.Close()

	currentDate := time.Now()
	drink.Id = 0
	drink.CreatedAt = &currentDate
	drink.UpdatedAt = &currentDate

	tableName := "drinks"
	query, data, err := commons.GetQuery(tableName, *drink, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&drink.Id,
		&drink.CreatedAt,
		&drink.UpdatedAt,
		&drink.DeletedAt,
		&drink.Name,
		&drink.Description,
		&drink.Photo,
		&drink.DrinkSizeID,
		&drink.DrinkFlavorID,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if drink.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetDrinks(root bool, relationalIDs *map[string]uint) ([]models.Drink, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "drinks"
	query, _, err := commons.GetQuery(tableName, models.Drink{}, "SELECT", false)

	if err != nil {
		return nil, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	if !root {
		query += " WHERE deleted_at IS NULL"
	}

	if relationalIDs != nil {
		switch root {
		case true:
			query += " WHERE "
		case false:
			query += " AND "
		}

		var relationConditions []string
		for key, value := range *relationalIDs {
			relationConditions = append(relationConditions, fmt.Sprintf("%s = %d", key, value))
		}
		query += strings.Join(relationConditions, " AND ")
	}

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	defer rows.Close()

	var drinks []models.Drink

	for rows.Next() {
		var drink models.Drink
		err := rows.Scan(
			&drink.Id,
			&drink.CreatedAt,
			&drink.UpdatedAt,
			&drink.DeletedAt,
			&drink.Name,
			&drink.Description,
			&drink.Photo,
			&drink.DrinkSizeID,
			&drink.DrinkFlavorID,
		)
		if err != nil {
			return nil, fmt.Errorf("can't scan the %s query ERROR: %s", tableName, err.Error())
		}
		drinks = append(drinks, drink)
	}
	if len(drinks) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}
	return drinks, nil
}

func GetDrink(id uint, root bool) (models.Drink, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "drinks"
	drink := models.Drink{}
	query, _, err := commons.GetQuery(tableName, drink, "SELECT", false)
	if err != nil {
		return drink, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}
	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	err = db.QueryRow(query, id).Scan(
		&drink.Id,
		&drink.CreatedAt,
		&drink.UpdatedAt,
		&drink.DeletedAt,
		&drink.Name,
		&drink.Description,
		&drink.Photo,
		&drink.DrinkSizeID,
		&drink.DrinkFlavorID,
	)
	if err != nil {
		return drink, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	if drink.Id == 0 {
		return drink, fmt.Errorf("%s not found", tableName)
	}

	return drink, nil
}

func DeleteDrink(id uint, root bool) error {
	return commons.DeleteFromTableById("drinks", id, root)
}

func UpdateDrink(updatingDrink *models.Drink, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "drinks"
	currentDate := time.Now()

	if !root {
		updatingDrink.UpdatedAt = &currentDate
		updatingDrink.CreatedAt = nil // Not allowed to update this field
		updatingDrink.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingDrink.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingDrink, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND deleted_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingDrink.Id,
		&updatingDrink.CreatedAt,
		&updatingDrink.UpdatedAt,
		&updatingDrink.DeletedAt,
		&updatingDrink.Name,
		&updatingDrink.Description,
		&updatingDrink.Photo,
		&updatingDrink.DrinkSizeID,
		&updatingDrink.DrinkFlavorID,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingDrink.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

//

func NewProduct(product *models.Product) error {
	db := database.Connect()
	defer db.Close()

	currentDate := time.Now()
	product.Id = 0
	product.CreatedAt = &currentDate
	product.UpdatedAt = &currentDate

	tableName := "product_foods"
	query, data, err := commons.GetQuery(tableName, *product, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&product.Id,
		&product.CreatedAt,
		&product.UpdatedAt,
		&product.DeletedAt,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Photo,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if product.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetProducts(root bool) ([]models.Product, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "product_foods"
	query, _, err := commons.GetQuery(tableName, models.Product{}, "SELECT", false)

	if err != nil {
		return nil, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	if !root {
		query += " WHERE deleted_at IS NULL"
	}

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var product models.Product
		err := rows.Scan(
			&product.Id,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.DeletedAt,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Photo,
		)
		if err != nil {
			return nil, fmt.Errorf("can't scan the %s query ERROR: %s", tableName, err.Error())
		}
		products = append(products, product)
	}
	if len(products) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}
	return products, nil
}

func GetProduct(id uint, root bool) (models.Product, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "product_foods"
	product := models.Product{}
	query, _, err := commons.GetQuery(tableName, product, "SELECT", false)
	if err != nil {
		return product, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}
	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	err = db.QueryRow(query, id).Scan(
		&product.Id,
		&product.CreatedAt,
		&product.UpdatedAt,
		&product.DeletedAt,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Photo,
	)
	if err != nil {
		return product, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	if product.Id == 0 {
		return product, fmt.Errorf("%s not found", tableName)
	}

	return product, nil
}

func DeleteProduct(id uint, root bool) error {
	return commons.DeleteFromTableById("products", id, root)
}

func UpdateProduct(updatingProduct *models.Product, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "products"
	currentDate := time.Now()

	if !root {
		updatingProduct.UpdatedAt = &currentDate
		updatingProduct.CreatedAt = nil // Not allowed to update this field
		updatingProduct.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingProduct.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingProduct, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND deleted_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingProduct.Id,
		&updatingProduct.CreatedAt,
		&updatingProduct.UpdatedAt,
		&updatingProduct.DeletedAt,
		&updatingProduct.Name,
		&updatingProduct.Description,
		&updatingProduct.Price,
		&updatingProduct.Photo,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingProduct.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

//

func NewProductFood(productFood *models.ProductFood) error {
	db := database.Connect()
	defer db.Close()

	currentDate := time.Now()
	productFood.Id = 0
	productFood.CreatedAt = &currentDate
	productFood.UpdatedAt = &currentDate

	tableName := "product_foods"
	query, data, err := commons.GetQuery(tableName, *productFood, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&productFood.Id,
		&productFood.CreatedAt,
		&productFood.UpdatedAt,
		&productFood.DeletedAt,
		&productFood.UnitQuantity,
		&productFood.GrammageQuantity,
		&productFood.FoodID,
		&productFood.ProductID,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if productFood.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetProductFoods(root bool, relationalIDs *map[string]uint) ([]models.ProductFood, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "product_foods"
	query, _, err := commons.GetQuery(tableName, models.ProductFood{}, "SELECT", false)

	if err != nil {
		return nil, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	if !root {
		query += " WHERE deleted_at IS NULL"
	}

	if relationalIDs != nil {
		switch root {
		case true:
			query += " WHERE "
		case false:
			query += " AND "
		}

		var relationConditions []string
		for key, value := range *relationalIDs {
			relationConditions = append(relationConditions, fmt.Sprintf("%s = %d", key, value))
		}
		query += strings.Join(relationConditions, " AND ")
	}

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	defer rows.Close()

	var productFoods []models.ProductFood

	for rows.Next() {
		var productFood models.ProductFood
		err := rows.Scan(
			&productFood.Id,
			&productFood.CreatedAt,
			&productFood.UpdatedAt,
			&productFood.DeletedAt,
			&productFood.UnitQuantity,
			&productFood.GrammageQuantity,
			&productFood.FoodID,
			&productFood.ProductID,
		)
		if err != nil {
			return nil, fmt.Errorf("can't scan the %s query ERROR: %s", tableName, err.Error())
		}
		productFoods = append(productFoods, productFood)
	}
	if len(productFoods) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}
	return productFoods, nil
}

func GetProductFood(id uint, root bool) (models.ProductFood, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "product_foods"
	productFood := models.ProductFood{}
	query, _, err := commons.GetQuery(tableName, productFood, "SELECT", false)
	if err != nil {
		return productFood, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}
	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	err = db.QueryRow(query, id).Scan(
		&productFood.Id,
		&productFood.CreatedAt,
		&productFood.UpdatedAt,
		&productFood.DeletedAt,
		&productFood.UnitQuantity,
		&productFood.GrammageQuantity,
		&productFood.FoodID,
		&productFood.ProductID,
	)
	if err != nil {
		return productFood, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	if productFood.Id == 0 {
		return productFood, fmt.Errorf("%s not found", tableName)
	}

	return productFood, nil
}

func DeleteProductFood(id uint, root bool) error {
	return commons.DeleteFromTableById("product_foods", id, root)
}

func UpdateProductFood(updatingProductFood *models.ProductFood, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "product_foods"
	currentDate := time.Now()

	if !root {
		updatingProductFood.UpdatedAt = &currentDate
		updatingProductFood.CreatedAt = nil // Not allowed to update this field
		updatingProductFood.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingProductFood.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingProductFood, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND deleted_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingProductFood.Id,
		&updatingProductFood.CreatedAt,
		&updatingProductFood.UpdatedAt,
		&updatingProductFood.DeletedAt,
		&updatingProductFood.UnitQuantity,
		&updatingProductFood.GrammageQuantity,
		&updatingProductFood.FoodID,
		&updatingProductFood.ProductID,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingProductFood.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

//

func NewProductDrink(productDrink *models.ProductDrink) error {
	db := database.Connect()
	defer db.Close()

	currentDate := time.Now()
	productDrink.Id = 0
	productDrink.CreatedAt = &currentDate
	productDrink.UpdatedAt = &currentDate

	tableName := "product_drinks"
	query, data, err := commons.GetQuery(tableName, *productDrink, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&productDrink.Id,
		&productDrink.CreatedAt,
		&productDrink.UpdatedAt,
		&productDrink.DeletedAt,
		&productDrink.UnitQuantity,
		&productDrink.GrammageQuantity,
		&productDrink.DrinkID,
		&productDrink.ProductID,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if productDrink.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetProductDrinks(root bool, relationalIDs *map[string]uint) ([]models.ProductDrink, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "product_drinks"
	query, _, err := commons.GetQuery(tableName, models.ProductDrink{}, "SELECT", false)

	if err != nil {
		return nil, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	if !root {
		query += " WHERE deleted_at IS NULL"
	}

	if relationalIDs != nil {
		switch root {
		case true:
			query += " WHERE "
		case false:
			query += " AND "
		}

		var relationConditions []string
		for key, value := range *relationalIDs {
			relationConditions = append(relationConditions, fmt.Sprintf("%s = %d", key, value))
		}
		query += strings.Join(relationConditions, " AND ")
	}

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	defer rows.Close()

	var productDrinks []models.ProductDrink

	for rows.Next() {
		var productDrink models.ProductDrink
		err := rows.Scan(
			&productDrink.Id,
			&productDrink.CreatedAt,
			&productDrink.UpdatedAt,
			&productDrink.DeletedAt,
			&productDrink.UnitQuantity,
			&productDrink.GrammageQuantity,
			&productDrink.DrinkID,
			&productDrink.ProductID,
		)
		if err != nil {
			return nil, fmt.Errorf("can't scan the %s query ERROR: %s", tableName, err.Error())
		}
		productDrinks = append(productDrinks, productDrink)
	}
	if len(productDrinks) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}
	return productDrinks, nil
}

func GetProductDrink(id uint, root bool) (models.ProductDrink, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "product_drinks"
	productDrink := models.ProductDrink{}
	query, _, err := commons.GetQuery(tableName, productDrink, "SELECT", false)
	if err != nil {
		return productDrink, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}
	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	err = db.QueryRow(query, id).Scan(
		&productDrink.Id,
		&productDrink.CreatedAt,
		&productDrink.UpdatedAt,
		&productDrink.DeletedAt,
		&productDrink.UnitQuantity,
		&productDrink.GrammageQuantity,
		&productDrink.DrinkID,
		&productDrink.ProductID,
	)
	if err != nil {
		return productDrink, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	if productDrink.Id == 0 {
		return productDrink, fmt.Errorf("%s not found", tableName)
	}

	return productDrink, nil
}

func DeleteProductDrink(id uint, root bool) error {
	return commons.DeleteFromTableById("product_drinks", id, root)
}

func UpdateProductDrink(updatingProductDrink *models.ProductDrink, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "prdocut_drinks"
	currentDate := time.Now()

	if !root {
		updatingProductDrink.UpdatedAt = &currentDate
		updatingProductDrink.CreatedAt = nil // Not allowed to update this field
		updatingProductDrink.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingProductDrink.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingProductDrink, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND deleted_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingProductDrink.Id,
		&updatingProductDrink.CreatedAt,
		&updatingProductDrink.UpdatedAt,
		&updatingProductDrink.DeletedAt,
		&updatingProductDrink.UnitQuantity,
		&updatingProductDrink.GrammageQuantity,
		&updatingProductDrink.DrinkID,
		&updatingProductDrink.ProductID,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingProductDrink.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

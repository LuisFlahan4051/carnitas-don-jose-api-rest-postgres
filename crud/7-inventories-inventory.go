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

func NewInventoryType(inventoryType *models.InventoryType) error {
	db := database.Connect()
	defer db.Close()

	tableName := "inventory_types"
	currentDate := time.Now()
	inventoryType.Id = 0 // To avoid the id to be forced
	inventoryType.CreatedAt = &currentDate
	inventoryType.UpdatedAt = &currentDate

	query, data, err := commons.GetQuery(tableName, *inventoryType, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&inventoryType.Id,
		&inventoryType.CreatedAt,
		&inventoryType.UpdatedAt,
		&inventoryType.DeletedAt,
		&inventoryType.Type,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())

	}
	if inventoryType.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetInventoryTypes(root bool) ([]models.InventoryType, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "inventory_types"
	query, _, err := commons.GetQuery(tableName, models.InventoryType{}, "SELECT", false)
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

	var inventoryTypes []models.InventoryType

	for rows.Next() {
		var inventoryType models.InventoryType
		err := rows.Scan(
			&inventoryType.Id,
			&inventoryType.CreatedAt,
			&inventoryType.UpdatedAt,
			&inventoryType.DeletedAt,
			&inventoryType.Type,
		)
		if err != nil {
			return nil, fmt.Errorf("can't scan the %s query ERROR: %s", tableName, err.Error())
		}
		inventoryTypes = append(inventoryTypes, inventoryType)
	}
	if len(inventoryTypes) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}
	return inventoryTypes, nil
}

func GetInventoryType(id uint, root bool) (models.InventoryType, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "inventory_types"
	inventoryType := models.InventoryType{}
	query, _, err := commons.GetQuery(tableName, inventoryType, "SELECT", false)
	if err != nil {
		return inventoryType, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	err = db.QueryRow(query, id).Scan(
		&inventoryType.Id,
		&inventoryType.CreatedAt,
		&inventoryType.UpdatedAt,
		&inventoryType.DeletedAt,
		&inventoryType.Type,
	)
	if err != nil {
		return inventoryType, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if inventoryType.Id == 0 {
		return inventoryType, fmt.Errorf("%s not found", tableName)
	}
	return inventoryType, nil
}

func DeleteInventoryType(id uint, root bool) error {
	return commons.DeleteFromTableById("inventory_types", id, root)
}

func UpdateInventoryType(updatingInventoryType *models.InventoryType, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "inventory_types"
	currentDate := time.Now()

	if !root {
		updatingInventoryType.UpdatedAt = &currentDate
		updatingInventoryType.CreatedAt = nil // Not allowed to update this field
		updatingInventoryType.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingInventoryType.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingInventoryType, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND delete_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingInventoryType.Id,
		&updatingInventoryType.CreatedAt,
		&updatingInventoryType.UpdatedAt,
		&updatingInventoryType.DeletedAt,
		&updatingInventoryType.Type,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingInventoryType.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

//

func NewInventory(inventory *models.Inventory) error {
	db := database.Connect()
	defer db.Close()

	tableName := "inventories"
	currentDate := time.Now()
	inventory.Id = 0 // To avoid the id to be forced
	inventory.CreatedAt = &currentDate
	inventory.UpdatedAt = &currentDate

	query, data, err := commons.GetQuery(tableName, *inventory, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&inventory.Id,
		&inventory.CreatedAt,
		&inventory.UpdatedAt,
		&inventory.DeletedAt,
		&inventory.Acepted,
		&inventory.InventoryTypeID,
		&inventory.BranchID,
		&inventory.TurnID,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())

	}
	if inventory.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetInventories(root bool) ([]models.Inventory, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "inventories"
	query, _, err := commons.GetQuery(tableName, models.Inventory{}, "SELECT", false)
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

	var inventorys []models.Inventory

	for rows.Next() {
		var inventory models.Inventory
		err := rows.Scan(
			&inventory.Id,
			&inventory.CreatedAt,
			&inventory.UpdatedAt,
			&inventory.DeletedAt,
			&inventory.Acepted,
			&inventory.InventoryTypeID,
			&inventory.BranchID,
			&inventory.TurnID,
		)
		if err != nil {
			return nil, fmt.Errorf("can't scan the %s query ERROR: %s", tableName, err.Error())
		}
		inventorys = append(inventorys, inventory)
	}
	if len(inventorys) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}
	return inventorys, nil
}

func GetInventory(id uint, root bool) (models.Inventory, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "inventories"
	inventory := models.Inventory{}
	query, _, err := commons.GetQuery(tableName, inventory, "SELECT", false)
	if err != nil {
		return inventory, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	err = db.QueryRow(query, id).Scan(
		&inventory.Id,
		&inventory.CreatedAt,
		&inventory.UpdatedAt,
		&inventory.DeletedAt,
		&inventory.Acepted,
		&inventory.InventoryTypeID,
		&inventory.BranchID,
		&inventory.TurnID,
	)
	if err != nil {
		return inventory, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if inventory.Id == 0 {
		return inventory, fmt.Errorf("%s not found", tableName)
	}
	return inventory, nil
}

func DeleteInventory(id uint, root bool) error {
	return commons.DeleteFromTableById("inventories", id, root)
}

func UpdateInventory(updatingInventory *models.Inventory, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "inventories"
	currentDate := time.Now()

	if !root {
		updatingInventory.UpdatedAt = &currentDate
		updatingInventory.CreatedAt = nil // Not allowed to update this field
		updatingInventory.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingInventory.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingInventory, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND delete_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingInventory.Id,
		&updatingInventory.CreatedAt,
		&updatingInventory.UpdatedAt,
		&updatingInventory.DeletedAt,
		&updatingInventory.Acepted,
		&updatingInventory.InventoryTypeID,
		&updatingInventory.BranchID,
		&updatingInventory.TurnID,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingInventory.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

//

func NewInventoryProductStock(inventoryProductStock *models.InventoryProductStock) error {
	db := database.Connect()
	defer db.Close()

	tableName := "inventory_products_stock"
	currentDate := time.Now()
	inventoryProductStock.Id = 0 // To avoid the id to be forced
	inventoryProductStock.CreatedAt = &currentDate
	inventoryProductStock.UpdatedAt = &currentDate

	query, data, err := commons.GetQuery(tableName, *inventoryProductStock, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&inventoryProductStock.Id,
		&inventoryProductStock.CreatedAt,
		&inventoryProductStock.UpdatedAt,
		&inventoryProductStock.DeletedAt,
		&inventoryProductStock.UnitQuantity,
		&inventoryProductStock.GrammageQuantity,
		&inventoryProductStock.InUse,
		&inventoryProductStock.InventoryID,
		&inventoryProductStock.ProductID,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())

	}
	if inventoryProductStock.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetInventoryProductsStock(root bool) ([]models.InventoryProductStock, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "inventory_products_stock"
	query, _, err := commons.GetQuery(tableName, models.InventoryProductStock{}, "SELECT", false)
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

	var inventoryProductsStocks []models.InventoryProductStock

	for rows.Next() {
		var inventoryProductStock models.InventoryProductStock
		err := rows.Scan(
			&inventoryProductStock.Id,
			&inventoryProductStock.CreatedAt,
			&inventoryProductStock.UpdatedAt,
			&inventoryProductStock.DeletedAt,
			&inventoryProductStock.UnitQuantity,
			&inventoryProductStock.GrammageQuantity,
			&inventoryProductStock.InUse,
			&inventoryProductStock.InventoryID,
			&inventoryProductStock.ProductID,
		)
		if err != nil {
			return nil, fmt.Errorf("can't scan the %s query ERROR: %s", tableName, err.Error())
		}
		inventoryProductsStocks = append(inventoryProductsStocks, inventoryProductStock)
	}
	if len(inventoryProductsStocks) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}
	return inventoryProductsStocks, nil
}

func GetInventoryProductStock(id uint, root bool) (models.InventoryProductStock, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "inventory_products_stock"
	inventoryProductStock := models.InventoryProductStock{}
	query, _, err := commons.GetQuery(tableName, inventoryProductStock, "SELECT", false)
	if err != nil {
		return inventoryProductStock, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	err = db.QueryRow(query, id).Scan(
		&inventoryProductStock.Id,
		&inventoryProductStock.CreatedAt,
		&inventoryProductStock.UpdatedAt,
		&inventoryProductStock.DeletedAt,
		&inventoryProductStock.UnitQuantity,
		&inventoryProductStock.GrammageQuantity,
		&inventoryProductStock.InUse,
		&inventoryProductStock.InventoryID,
		&inventoryProductStock.ProductID,
	)
	if err != nil {
		return inventoryProductStock, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if inventoryProductStock.Id == 0 {
		return inventoryProductStock, fmt.Errorf("%s not found", tableName)
	}
	return inventoryProductStock, nil
}

func DeleteInventoryProductStock(id uint, root bool) error {
	return commons.DeleteFromTableById("inventory_products_stock", id, root)
}

func UpdateInventoryProductStock(updatingInventoryProductStock *models.InventoryProductStock, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "inventory_products_stock"
	currentDate := time.Now()

	if !root {
		updatingInventoryProductStock.UpdatedAt = &currentDate
		updatingInventoryProductStock.CreatedAt = nil // Not allowed to update this field
		updatingInventoryProductStock.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingInventoryProductStock.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingInventoryProductStock, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND delete_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingInventoryProductStock.Id,
		&updatingInventoryProductStock.CreatedAt,
		&updatingInventoryProductStock.UpdatedAt,
		&updatingInventoryProductStock.DeletedAt,
		&updatingInventoryProductStock.UnitQuantity,
		&updatingInventoryProductStock.GrammageQuantity,
		&updatingInventoryProductStock.InUse,
		&updatingInventoryProductStock.InventoryID,
		&updatingInventoryProductStock.ProductID,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingInventoryProductStock.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

//

func NewInventorySupplyStock(inventorySupplyStock *models.InventorySupplyStock) error {
	db := database.Connect()
	defer db.Close()

	tableName := "inventory_supplies_stock"
	currentDate := time.Now()
	inventorySupplyStock.Id = 0 // To avoid the id to be forced
	inventorySupplyStock.CreatedAt = &currentDate
	inventorySupplyStock.UpdatedAt = &currentDate

	query, data, err := commons.GetQuery(tableName, *inventorySupplyStock, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&inventorySupplyStock.Id,
		&inventorySupplyStock.CreatedAt,
		&inventorySupplyStock.UpdatedAt,
		&inventorySupplyStock.DeletedAt,
		&inventorySupplyStock.UnitQuantity,
		&inventorySupplyStock.GrammageQuantity,
		&inventorySupplyStock.InUse,
		&inventorySupplyStock.InventoryID,
		&inventorySupplyStock.SupplyID,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())

	}
	if inventorySupplyStock.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetInventorySuppliesStock(root bool) ([]models.InventorySupplyStock, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "inventory_supplies_stock"
	query, _, err := commons.GetQuery(tableName, models.InventorySupplyStock{}, "SELECT", false)
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

	var inventorySuppliesStocks []models.InventorySupplyStock

	for rows.Next() {
		var inventorySupplyStock models.InventorySupplyStock
		err := rows.Scan(
			&inventorySupplyStock.Id,
			&inventorySupplyStock.CreatedAt,
			&inventorySupplyStock.UpdatedAt,
			&inventorySupplyStock.DeletedAt,
			&inventorySupplyStock.UnitQuantity,
			&inventorySupplyStock.GrammageQuantity,
			&inventorySupplyStock.InUse,
			&inventorySupplyStock.InventoryID,
			&inventorySupplyStock.SupplyID,
		)
		if err != nil {
			return nil, fmt.Errorf("can't scan the %s query ERROR: %s", tableName, err.Error())
		}
		inventorySuppliesStocks = append(inventorySuppliesStocks, inventorySupplyStock)
	}
	if len(inventorySuppliesStocks) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}
	return inventorySuppliesStocks, nil
}

func GetInventorySupplyStock(id uint, root bool) (models.InventorySupplyStock, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "inventory_supplies_stock"
	inventorySupplyStock := models.InventorySupplyStock{}
	query, _, err := commons.GetQuery(tableName, inventorySupplyStock, "SELECT", false)
	if err != nil {
		return inventorySupplyStock, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	err = db.QueryRow(query, id).Scan(
		&inventorySupplyStock.Id,
		&inventorySupplyStock.CreatedAt,
		&inventorySupplyStock.UpdatedAt,
		&inventorySupplyStock.DeletedAt,
		&inventorySupplyStock.UnitQuantity,
		&inventorySupplyStock.GrammageQuantity,
		&inventorySupplyStock.InUse,
		&inventorySupplyStock.InventoryID,
		&inventorySupplyStock.SupplyID,
	)
	if err != nil {
		return inventorySupplyStock, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if inventorySupplyStock.Id == 0 {
		return inventorySupplyStock, fmt.Errorf("%s not found", tableName)
	}
	return inventorySupplyStock, nil
}

func DeleteInventorySupplyStock(id uint, root bool) error {
	return commons.DeleteFromTableById("inventory_supplies_stock", id, root)
}

func UpdateInventorySupplyStock(updatingInventorySupplyStock *models.InventorySupplyStock, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "inventory_supplies_stock"
	currentDate := time.Now()

	if !root {
		updatingInventorySupplyStock.UpdatedAt = &currentDate
		updatingInventorySupplyStock.CreatedAt = nil // Not allowed to update this field
		updatingInventorySupplyStock.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingInventorySupplyStock.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingInventorySupplyStock, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND delete_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingInventorySupplyStock.Id,
		&updatingInventorySupplyStock.CreatedAt,
		&updatingInventorySupplyStock.UpdatedAt,
		&updatingInventorySupplyStock.DeletedAt,
		&updatingInventorySupplyStock.UnitQuantity,
		&updatingInventorySupplyStock.GrammageQuantity,
		&updatingInventorySupplyStock.InUse,
		&updatingInventorySupplyStock.InventoryID,
		&updatingInventorySupplyStock.SupplyID,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingInventorySupplyStock.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

//

func NewInventoryArticleStock(inventoryArticleStock *models.InventoryArticleStock) error {
	db := database.Connect()
	defer db.Close()

	tableName := "inventory_articles_stock"
	currentDate := time.Now()
	inventoryArticleStock.Id = 0 // To avoid the id to be forced
	inventoryArticleStock.CreatedAt = &currentDate
	inventoryArticleStock.UpdatedAt = &currentDate

	query, data, err := commons.GetQuery(tableName, *inventoryArticleStock, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&inventoryArticleStock.Id,
		&inventoryArticleStock.CreatedAt,
		&inventoryArticleStock.UpdatedAt,
		&inventoryArticleStock.DeletedAt,
		&inventoryArticleStock.UnitQuantity,
		&inventoryArticleStock.GrammageQuantity,
		&inventoryArticleStock.InUse,
		&inventoryArticleStock.InventoryID,
		&inventoryArticleStock.ArticleID,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())

	}
	if inventoryArticleStock.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetInventoryArticlesStock(root bool) ([]models.InventoryArticleStock, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "inventory_articles_stock"
	query, _, err := commons.GetQuery(tableName, models.InventoryArticleStock{}, "SELECT", false)
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

	var inventoryArticlesStocks []models.InventoryArticleStock

	for rows.Next() {
		var inventoryArticleStock models.InventoryArticleStock
		err := rows.Scan(
			&inventoryArticleStock.Id,
			&inventoryArticleStock.CreatedAt,
			&inventoryArticleStock.UpdatedAt,
			&inventoryArticleStock.DeletedAt,
			&inventoryArticleStock.UnitQuantity,
			&inventoryArticleStock.GrammageQuantity,
			&inventoryArticleStock.InUse,
			&inventoryArticleStock.InventoryID,
			&inventoryArticleStock.ArticleID,
		)
		if err != nil {
			return nil, fmt.Errorf("can't scan the %s query ERROR: %s", tableName, err.Error())
		}
		inventoryArticlesStocks = append(inventoryArticlesStocks, inventoryArticleStock)
	}
	if len(inventoryArticlesStocks) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}
	return inventoryArticlesStocks, nil
}

func GetInventoryArticleStock(id uint, root bool) (models.InventoryArticleStock, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "inventory_articles_stock"
	inventoryArticleStock := models.InventoryArticleStock{}
	query, _, err := commons.GetQuery(tableName, inventoryArticleStock, "SELECT", false)
	if err != nil {
		return inventoryArticleStock, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	err = db.QueryRow(query, id).Scan(
		&inventoryArticleStock.Id,
		&inventoryArticleStock.CreatedAt,
		&inventoryArticleStock.UpdatedAt,
		&inventoryArticleStock.DeletedAt,
		&inventoryArticleStock.UnitQuantity,
		&inventoryArticleStock.GrammageQuantity,
		&inventoryArticleStock.InUse,
		&inventoryArticleStock.InventoryID,
		&inventoryArticleStock.ArticleID,
	)
	if err != nil {
		return inventoryArticleStock, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if inventoryArticleStock.Id == 0 {
		return inventoryArticleStock, fmt.Errorf("%s not found", tableName)
	}
	return inventoryArticleStock, nil
}

func DeleteInventoryArticleStock(id uint, root bool) error {
	return commons.DeleteFromTableById("inventory_articles_stock", id, root)
}

func UpdateInventoryArticleStock(updatingInventoryArticleStock *models.InventoryArticleStock, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "inventory_articles_stock"
	currentDate := time.Now()

	if !root {
		updatingInventoryArticleStock.UpdatedAt = &currentDate
		updatingInventoryArticleStock.CreatedAt = nil // Not allowed to update this field
		updatingInventoryArticleStock.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingInventoryArticleStock.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingInventoryArticleStock, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND delete_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingInventoryArticleStock.Id,
		&updatingInventoryArticleStock.CreatedAt,
		&updatingInventoryArticleStock.UpdatedAt,
		&updatingInventoryArticleStock.DeletedAt,
		&updatingInventoryArticleStock.UnitQuantity,
		&updatingInventoryArticleStock.GrammageQuantity,
		&updatingInventoryArticleStock.InUse,
		&updatingInventoryArticleStock.InventoryID,
		&updatingInventoryArticleStock.ArticleID,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingInventoryArticleStock.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

//

func NewInventorySafebox(inventorySafebox *models.InventorySafebox) error {
	db := database.Connect()
	defer db.Close()

	tableName := "inventory_safebox"
	currentDate := time.Now()
	inventorySafebox.Id = 0 // To avoid the id to be forced
	inventorySafebox.CreatedAt = &currentDate
	inventorySafebox.UpdatedAt = &currentDate

	query, data, err := commons.GetQuery(tableName, *inventorySafebox, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&inventorySafebox.Id,
		&inventorySafebox.CreatedAt,
		&inventorySafebox.UpdatedAt,
		&inventorySafebox.DeletedAt,
		&inventorySafebox.InventoryID,
		&inventorySafebox.SafeboxID,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())

	}
	if inventorySafebox.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetInventoriesSafebox(root bool) ([]models.InventorySafebox, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "inventory_safebox"
	query, _, err := commons.GetQuery(tableName, models.InventorySafebox{}, "SELECT", false)
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

	var inventoriesSafebox []models.InventorySafebox

	for rows.Next() {
		var inventorySafebox models.InventorySafebox
		err := rows.Scan(
			&inventorySafebox.Id,
			&inventorySafebox.CreatedAt,
			&inventorySafebox.UpdatedAt,
			&inventorySafebox.DeletedAt,
			&inventorySafebox.InventoryID,
			&inventorySafebox.SafeboxID,
		)
		if err != nil {
			return nil, fmt.Errorf("can't scan the %s query ERROR: %s", tableName, err.Error())
		}
		inventoriesSafebox = append(inventoriesSafebox, inventorySafebox)
	}
	if len(inventoriesSafebox) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}
	return inventoriesSafebox, nil
}

func GetInventorySafebox(id uint, root bool) (models.InventorySafebox, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "inventory_safebox"
	inventorySafebox := models.InventorySafebox{}
	query, _, err := commons.GetQuery(tableName, inventorySafebox, "SELECT", false)
	if err != nil {
		return inventorySafebox, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	err = db.QueryRow(query, id).Scan(
		&inventorySafebox.Id,
		&inventorySafebox.CreatedAt,
		&inventorySafebox.UpdatedAt,
		&inventorySafebox.DeletedAt,
		&inventorySafebox.InventoryID,
		&inventorySafebox.SafeboxID,
	)
	if err != nil {
		return inventorySafebox, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if inventorySafebox.Id == 0 {
		return inventorySafebox, fmt.Errorf("%s not found", tableName)
	}
	return inventorySafebox, nil
}

func DeleteInventorySafebox(id uint, root bool) error {
	return commons.DeleteFromTableById("inventory_safebox", id, root)
}

func UpdateInventorySafebox(updatingInventorySafebox *models.InventorySafebox, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "inventory_safebox"
	currentDate := time.Now()

	if !root {
		updatingInventorySafebox.UpdatedAt = &currentDate
		updatingInventorySafebox.CreatedAt = nil // Not allowed to update this field
		updatingInventorySafebox.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingInventorySafebox.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingInventorySafebox, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND delete_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingInventorySafebox.Id,
		&updatingInventorySafebox.CreatedAt,
		&updatingInventorySafebox.UpdatedAt,
		&updatingInventorySafebox.DeletedAt,
		&updatingInventorySafebox.InventoryID,
		&updatingInventorySafebox.SafeboxID,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingInventorySafebox.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

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

func NewBranch(branch *models.Branch) error {
	db := database.Connect()
	defer db.Close()

	tableName := "branches"
	currentDate := time.Now()
	branch.Id = 0 // To avoid the id to be forced
	branch.CreatedAt = &currentDate
	branch.UpdatedAt = &currentDate

	query, data, err := commons.GetQuery(tableName, *branch, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&branch.Id,
		&branch.CreatedAt,
		&branch.UpdatedAt,
		&branch.DeletedAt,
		&branch.Name,
		&branch.Address,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())

	}
	if branch.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetBranches(root bool) ([]models.Branch, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "branches"
	query, _, err := commons.GetQuery(tableName, models.Branch{}, "SELECT", false)
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

	var branches []models.Branch

	for rows.Next() {
		var branch models.Branch
		err := rows.Scan(
			&branch.Id,
			&branch.CreatedAt,
			&branch.UpdatedAt,
			&branch.DeletedAt,
			&branch.Name,
			&branch.Address,
		)
		if err != nil {
			return nil, fmt.Errorf("can't scan the %s query ERROR: %s", tableName, err.Error())
		}
		branches = append(branches, branch)
	}
	if len(branches) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}
	return branches, nil
}

func GetBranch(id uint, root bool) (models.Branch, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "branches"
	branch := models.Branch{}
	query, _, err := commons.GetQuery(tableName, branch, "SELECT", false)
	if err != nil {
		return branch, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	err = db.QueryRow(query, id).Scan(
		&branch.Id,
		&branch.CreatedAt,
		&branch.UpdatedAt,
		&branch.DeletedAt,
		&branch.Name,
		&branch.Address,
	)
	if err != nil {
		return branch, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if branch.Id == 0 {
		return branch, fmt.Errorf("%s not found", tableName)
	}
	return branch, nil
}

func DeleteBranch(id uint, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "branches"
	if root {
		query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", tableName)
		_, err := db.Exec(query, id)
		if err != nil {
			return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
		}
		return nil
	}

	if commons.IsDeleted(tableName, id) {
		return errors.New("already deleted")
	}

	query := fmt.Sprintf("UPDATE %s SET deleted_at = NOW() WHERE id = $1", tableName)
	result, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to delete one %s ERROR: %s", tableName, err.Error())
	}

	if rowsAffected == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

func UpdateBranch(updatingBranch *models.Branch, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "branches"
	currentDate := time.Now()

	if !root {
		updatingBranch.UpdatedAt = &currentDate
		updatingBranch.CreatedAt = nil // Not allowed to update this field
		updatingBranch.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingBranch.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingBranch, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND delete_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingBranch.Id,
		&updatingBranch.CreatedAt,
		&updatingBranch.UpdatedAt,
		&updatingBranch.DeletedAt,
		&updatingBranch.Name,
		&updatingBranch.Address,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingBranch.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

//

func NewBranchSafebox(branchSafebox *models.BranchSafebox) error {
	db := database.Connect()
	defer db.Close()

	tableName := "branch_safeboxes"
	currentDate := time.Now()
	branchSafebox.Id = 0
	branchSafebox.CreatedAt = &currentDate
	branchSafebox.UpdatedAt = &currentDate

	query, data, err := commons.GetQuery(tableName, *branchSafebox, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&branchSafebox.Id,
		&branchSafebox.CreatedAt,
		&branchSafebox.UpdatedAt,
		&branchSafebox.DeletedAt,
		&branchSafebox.Name,
		&branchSafebox.Content,
		&branchSafebox.BranchID,
		&branchSafebox.SafeboxID,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if branchSafebox.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetBranchSafeboxes(root bool) ([]models.BranchSafebox, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "branch_safeboxes"
	query, _, err := commons.GetQuery(tableName, models.BranchSafebox{}, "SELECT", false)
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

	var branchSafeboxes []models.BranchSafebox
	for rows.Next() {
		var branchSafebox models.BranchSafebox
		err = rows.Scan(
			&branchSafebox.Id,
			&branchSafebox.CreatedAt,
			&branchSafebox.UpdatedAt,
			&branchSafebox.DeletedAt,
			&branchSafebox.Name,
			&branchSafebox.Content,
			&branchSafebox.BranchID,
			&branchSafebox.SafeboxID,
		)
		if err != nil {
			return nil, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
		}

		branchSafeboxes = append(branchSafeboxes, branchSafebox)
	}

	if len(branchSafeboxes) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}

	return branchSafeboxes, nil
}

func GetBranchSafebox(id uint, root bool) (models.BranchSafebox, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "branch_safeboxes"
	branchSafebox := models.BranchSafebox{}
	query, _, err := commons.GetQuery(tableName, branchSafebox, "SELECT", false)
	if err != nil {
		return branchSafebox, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	err = db.QueryRow(query, id).Scan(
		&branchSafebox.Id,
		&branchSafebox.CreatedAt,
		&branchSafebox.UpdatedAt,
		&branchSafebox.DeletedAt,
		&branchSafebox.Name,
		&branchSafebox.Content,
		&branchSafebox.BranchID,
		&branchSafebox.SafeboxID,
	)
	if err != nil {
		return branchSafebox, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if branchSafebox.Id == 0 {
		return branchSafebox, fmt.Errorf("%s not found", tableName)
	}
	return branchSafebox, nil
}

func DeleteBranchSafebox(id uint, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "branch_safeboxes"
	if root {
		query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", tableName)
		_, err := db.Exec(query, id)
		if err != nil {
			return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
		}
		return nil
	}

	if commons.IsDeleted(tableName, id) {
		return errors.New("already deleted")
	}

	query := fmt.Sprintf("UPDATE %s SET deleted_at = NOW() WHERE id = $1", tableName)
	result, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to delete one %s ERROR: %s", tableName, err.Error())
	}

	if rowsAffected == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

func UpdateBranchSafebox(updatingBranchSafebox *models.BranchSafebox, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "branch_safeboxes"
	currentDate := time.Now()

	if !root {
		updatingBranchSafebox.UpdatedAt = &currentDate
		updatingBranchSafebox.CreatedAt = nil // Not allowed to update this field
		updatingBranchSafebox.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingBranchSafebox.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingBranchSafebox, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND delete_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingBranchSafebox.Id,
		&updatingBranchSafebox.CreatedAt,
		&updatingBranchSafebox.UpdatedAt,
		&updatingBranchSafebox.DeletedAt,
		&updatingBranchSafebox.Name,
		&updatingBranchSafebox.Content,
		&updatingBranchSafebox.BranchID,
		&updatingBranchSafebox.SafeboxID,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingBranchSafebox.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

//

func NewBranchProductStock(branchProductStock *models.BranchProductStock) error {
	db := database.Connect()
	defer db.Close()

	currentDate := time.Now()
	branchProductStock.Id = 0
	branchProductStock.CreatedAt = &currentDate
	branchProductStock.UpdatedAt = &currentDate

	tableName := "branch_products_stock"
	query, data, err := commons.GetQuery(tableName, *branchProductStock, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&branchProductStock.Id,
		&branchProductStock.CreatedAt,
		&branchProductStock.UpdatedAt,
		&branchProductStock.DeletedAt,
		&branchProductStock.UnitQuantity,
		&branchProductStock.GrammageQuantity,
		&branchProductStock.InUse,
		&branchProductStock.BranchID,
		&branchProductStock.ProductID,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if branchProductStock.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetBranchProductsStock(root bool) ([]models.BranchProductStock, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "branch_products_stock"
	query, _, err := commons.GetQuery(tableName, models.BranchProductStock{}, "SELECT", false)
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

	var branchProductsStock []models.BranchProductStock
	for rows.Next() {
		var branchProductStock models.BranchProductStock
		err = rows.Scan(
			&branchProductStock.Id,
			&branchProductStock.CreatedAt,
			&branchProductStock.UpdatedAt,
			&branchProductStock.DeletedAt,
			&branchProductStock.UnitQuantity,
			&branchProductStock.GrammageQuantity,
			&branchProductStock.InUse,
			&branchProductStock.BranchID,
			&branchProductStock.ProductID,
		)
		if err != nil {
			return nil, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
		}

		branchProductsStock = append(branchProductsStock, branchProductStock)
	}

	if len(branchProductsStock) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}

	return branchProductsStock, nil
}

func GetBranchProductStock(id uint, root bool) (models.BranchProductStock, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "branch_products_stock"
	query, _, _ := commons.GetQuery(tableName, models.BranchProductStock{}, "SELECT", false)
	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	branchProductStock := models.BranchProductStock{}
	err := db.QueryRow(query, id).Scan(
		&branchProductStock.Id,
		&branchProductStock.CreatedAt,
		&branchProductStock.UpdatedAt,
		&branchProductStock.DeletedAt,
		&branchProductStock.UnitQuantity,
		&branchProductStock.GrammageQuantity,
		&branchProductStock.InUse,
		&branchProductStock.BranchID,
		&branchProductStock.ProductID,
	)
	if err != nil {
		return branchProductStock, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	if branchProductStock.Id == 0 {
		return branchProductStock, fmt.Errorf("%s not found", tableName)

	}

	return branchProductStock, nil
}

func DeleteBranchProductStock(id uint, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "branch_products_stock"
	if root {
		query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", tableName)
		_, err := db.Exec(query, id)
		if err != nil {
			return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
		}
		return nil
	}

	if commons.IsDeleted(tableName, id) {
		return errors.New("already deleted")
	}

	query := fmt.Sprintf("UPDATE %s SET deleted_at = NOW() WHERE id = $1", tableName)
	result, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to delete one %s ERROR: %s", tableName, err.Error())
	}

	if rowsAffected == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

func UpdateBranchProductStock(updatingBranchProductStock *models.BranchProductStock, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "branch_products_stock"
	currentDate := time.Now()

	if !root {
		updatingBranchProductStock.UpdatedAt = &currentDate
		updatingBranchProductStock.CreatedAt = nil // Not allowed to update this field
		updatingBranchProductStock.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingBranchProductStock.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingBranchProductStock, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND delete_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingBranchProductStock.Id,
		&updatingBranchProductStock.CreatedAt,
		&updatingBranchProductStock.UpdatedAt,
		&updatingBranchProductStock.DeletedAt,
		&updatingBranchProductStock.UnitQuantity,
		&updatingBranchProductStock.GrammageQuantity,
		&updatingBranchProductStock.InUse,
		&updatingBranchProductStock.BranchID,
		&updatingBranchProductStock.ProductID,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingBranchProductStock.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

//

func NewBranchSupplyStock(branchSupplyStock *models.BranchSupplyStock) error {
	db := database.Connect()
	defer db.Close()

	currentDate := time.Now()
	branchSupplyStock.Id = 0
	branchSupplyStock.CreatedAt = &currentDate
	branchSupplyStock.UpdatedAt = &currentDate

	tableName := "branch_supplies_stock"
	query, data, err := commons.GetQuery(tableName, *branchSupplyStock, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&branchSupplyStock.Id,
		&branchSupplyStock.CreatedAt,
		&branchSupplyStock.UpdatedAt,
		&branchSupplyStock.DeletedAt,
		&branchSupplyStock.UnitQuantity,
		&branchSupplyStock.GrammageQuantity,
		&branchSupplyStock.InUse,
		&branchSupplyStock.BranchID,
		&branchSupplyStock.SupplyID,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if branchSupplyStock.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetBranchSuppliesStocks(root bool) ([]models.BranchSupplyStock, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "branch_supplies_stock"
	query, _, err := commons.GetQuery(tableName, models.BranchSupplyStock{}, "SELECT", false)
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

	var branchSuppliesStock []models.BranchSupplyStock
	for rows.Next() {
		var branchSupplyStock models.BranchSupplyStock
		err = rows.Scan(
			&branchSupplyStock.Id,
			&branchSupplyStock.CreatedAt,
			&branchSupplyStock.UpdatedAt,
			&branchSupplyStock.DeletedAt,
			&branchSupplyStock.UnitQuantity,
			&branchSupplyStock.GrammageQuantity,
			&branchSupplyStock.InUse,
			&branchSupplyStock.BranchID,
			&branchSupplyStock.SupplyID,
		)
		if err != nil {
			return nil, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
		}

		branchSuppliesStock = append(branchSuppliesStock, branchSupplyStock)
	}

	if len(branchSuppliesStock) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}

	return branchSuppliesStock, nil
}

func GetBranchSupplyStock(id uint, root bool) (models.BranchSupplyStock, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "branch_supplies_stock"
	query, _, _ := commons.GetQuery(tableName, models.BranchSupplyStock{}, "SELECT", false)
	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	branchSupplyStock := models.BranchSupplyStock{}
	err := db.QueryRow(query, id).Scan(
		&branchSupplyStock.Id,
		&branchSupplyStock.CreatedAt,
		&branchSupplyStock.UpdatedAt,
		&branchSupplyStock.DeletedAt,
		&branchSupplyStock.UnitQuantity,
		&branchSupplyStock.GrammageQuantity,
		&branchSupplyStock.InUse,
		&branchSupplyStock.BranchID,
		&branchSupplyStock.SupplyID,
	)
	if err != nil {
		return branchSupplyStock, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	if branchSupplyStock.Id == 0 {
		return branchSupplyStock, fmt.Errorf("%s not found", tableName)

	}

	return branchSupplyStock, nil
}

func DeleteBranchSupplyStock(id uint, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "branch_supplies_stock"
	if root {
		query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", tableName)
		_, err := db.Exec(query, id)
		if err != nil {
			return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
		}
		return nil
	}

	if commons.IsDeleted(tableName, id) {
		return errors.New("already deleted")
	}

	query := fmt.Sprintf("UPDATE %s SET deleted_at = NOW() WHERE id = $1", tableName)
	result, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to delete one %s ERROR: %s", tableName, err.Error())
	}

	if rowsAffected == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

func UpdateBranchSupplyStock(updatingBranchSupplyStock *models.BranchSupplyStock, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "branch_supplies_stock"
	currentDate := time.Now()

	if !root {
		updatingBranchSupplyStock.UpdatedAt = &currentDate
		updatingBranchSupplyStock.CreatedAt = nil // Not allowed to update this field
		updatingBranchSupplyStock.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingBranchSupplyStock.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingBranchSupplyStock, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND delete_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingBranchSupplyStock.Id,
		&updatingBranchSupplyStock.CreatedAt,
		&updatingBranchSupplyStock.UpdatedAt,
		&updatingBranchSupplyStock.DeletedAt,
		&updatingBranchSupplyStock.UnitQuantity,
		&updatingBranchSupplyStock.GrammageQuantity,
		&updatingBranchSupplyStock.InUse,
		&updatingBranchSupplyStock.BranchID,
		&updatingBranchSupplyStock.SupplyID,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingBranchSupplyStock.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

//

func NewBranchArticleStock(branchArticleStock *models.BranchArticleStock) error {
	db := database.Connect()
	defer db.Close()

	currentDate := time.Now()
	branchArticleStock.Id = 0
	branchArticleStock.CreatedAt = &currentDate
	branchArticleStock.UpdatedAt = &currentDate

	tableName := "branch_supplies_stock"
	query, data, err := commons.GetQuery(tableName, *branchArticleStock, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&branchArticleStock.Id,
		&branchArticleStock.CreatedAt,
		&branchArticleStock.UpdatedAt,
		&branchArticleStock.DeletedAt,
		&branchArticleStock.UnitQuantity,
		&branchArticleStock.GrammageQuantity,
		&branchArticleStock.InUse,
		&branchArticleStock.BranchID,
		&branchArticleStock.ArticleID,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if branchArticleStock.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetBranchArticleStocks(root bool) ([]models.BranchArticleStock, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "branch_articles_stock"
	query, _, err := commons.GetQuery(tableName, models.BranchArticleStock{}, "SELECT", false)
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

	var branchArticleStocks []models.BranchArticleStock
	for rows.Next() {
		var branchArticleStock models.BranchArticleStock
		err = rows.Scan(
			&branchArticleStock.Id,
			&branchArticleStock.CreatedAt,
			&branchArticleStock.UpdatedAt,
			&branchArticleStock.DeletedAt,
			&branchArticleStock.UnitQuantity,
			&branchArticleStock.GrammageQuantity,
			&branchArticleStock.InUse,
			&branchArticleStock.BranchID,
			&branchArticleStock.ArticleID,
		)
		if err != nil {
			return nil, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
		}

		branchArticleStocks = append(branchArticleStocks, branchArticleStock)
	}

	if len(branchArticleStocks) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}

	return branchArticleStocks, nil
}

func GetBranchArticleStock(id uint, root bool) (models.BranchArticleStock, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "branch_articles_stock"
	query, _, _ := commons.GetQuery(tableName, models.BranchArticleStock{}, "SELECT", false)
	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	branchArticleStock := models.BranchArticleStock{}
	err := db.QueryRow(query, id).Scan(
		&branchArticleStock.Id,
		&branchArticleStock.CreatedAt,
		&branchArticleStock.UpdatedAt,
		&branchArticleStock.DeletedAt,
		&branchArticleStock.UnitQuantity,
		&branchArticleStock.GrammageQuantity,
		&branchArticleStock.InUse,
		&branchArticleStock.BranchID,
		&branchArticleStock.ArticleID,
	)
	if err != nil {
		return branchArticleStock, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	if branchArticleStock.Id == 0 {
		return branchArticleStock, fmt.Errorf("%s not found", tableName)

	}

	return branchArticleStock, nil
}

func DeleteBranchArticleStock(id uint, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "branch_articles_stock"
	if root {
		query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", tableName)
		_, err := db.Exec(query, id)
		if err != nil {
			return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
		}
		return nil
	}

	if commons.IsDeleted(tableName, id) {
		return errors.New("already deleted")
	}

	query := fmt.Sprintf("UPDATE %s SET deleted_at = NOW() WHERE id = $1", tableName)
	result, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to delete one %s ERROR: %s", tableName, err.Error())
	}

	if rowsAffected == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

func UpdateBranchArticleStock(updatingBranchArticleStock *models.BranchArticleStock, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "branch_articles_stock"
	currentDate := time.Now()

	if !root {
		updatingBranchArticleStock.UpdatedAt = &currentDate
		updatingBranchArticleStock.CreatedAt = nil // Not allowed to update this field
		updatingBranchArticleStock.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingBranchArticleStock.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingBranchArticleStock, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND delete_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingBranchArticleStock.Id,
		&updatingBranchArticleStock.CreatedAt,
		&updatingBranchArticleStock.UpdatedAt,
		&updatingBranchArticleStock.DeletedAt,
		&updatingBranchArticleStock.UnitQuantity,
		&updatingBranchArticleStock.GrammageQuantity,
		&updatingBranchArticleStock.InUse,
		&updatingBranchArticleStock.BranchID,
		&updatingBranchArticleStock.ArticleID,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingBranchArticleStock.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

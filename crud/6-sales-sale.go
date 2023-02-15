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

func NewSale(sale *models.Sale) error {
	db := database.Connect()
	defer db.Close()

	tableName := "sales"
	currentDate := time.Now()
	sale.Id = 0 // To avoid the id to be forced
	sale.CreatedAt = &currentDate
	sale.UpdatedAt = &currentDate

	query, data, err := commons.GetQuery(tableName, *sale, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&sale.Id,
		&sale.CreatedAt,
		&sale.UpdatedAt,
		&sale.DeletedAt,
		&sale.EntryDate,
		&sale.ExitDate,
		&sale.TableNumber,
		&sale.Packed,
		&sale.UserID,
		&sale.BranchID,
		&sale.TurnID,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())

	}
	if sale.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetSales(root bool, relationalIDs *map[string]uint) ([]models.Sale, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "sales"
	query, _, err := commons.GetQuery(tableName, models.Sale{}, "SELECT", false)
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

	var sales []models.Sale

	for rows.Next() {
		var sale models.Sale
		err := rows.Scan(
			&sale.Id,
			&sale.CreatedAt,
			&sale.UpdatedAt,
			&sale.DeletedAt,
			&sale.EntryDate,
			&sale.ExitDate,
			&sale.TableNumber,
			&sale.Packed,
			&sale.UserID,
			&sale.BranchID,
			&sale.TurnID,
		)
		if err != nil {
			return nil, fmt.Errorf("can't scan the %s query ERROR: %s", tableName, err.Error())
		}
		sales = append(sales, sale)
	}
	if len(sales) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}
	return sales, nil
}

func GetSale(id uint, root bool) (models.Sale, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "sales"
	sale := models.Sale{}
	query, _, err := commons.GetQuery(tableName, sale, "SELECT", false)
	if err != nil {
		return sale, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	err = db.QueryRow(query, id).Scan(
		&sale.Id,
		&sale.CreatedAt,
		&sale.UpdatedAt,
		&sale.DeletedAt,
		&sale.EntryDate,
		&sale.ExitDate,
		&sale.TableNumber,
		&sale.Packed,
		&sale.UserID,
		&sale.BranchID,
		&sale.TurnID,
	)
	if err != nil {
		return sale, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if sale.Id == 0 {
		return sale, fmt.Errorf("%s not found", tableName)
	}
	return sale, nil
}

func DeleteSale(id uint, root bool) error {
	return commons.DeleteFromTableById("sales", id, root)
}

func UpdateSale(updatingSale *models.Sale, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "sales"
	currentDate := time.Now()

	if !root {
		updatingSale.UpdatedAt = &currentDate
		updatingSale.CreatedAt = nil // Not allowed to update this field
		updatingSale.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingSale.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingSale, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND deleted_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingSale.Id,
		&updatingSale.CreatedAt,
		&updatingSale.UpdatedAt,
		&updatingSale.DeletedAt,
		&updatingSale.EntryDate,
		&updatingSale.ExitDate,
		&updatingSale.TableNumber,
		&updatingSale.Packed,
		&updatingSale.UserID,
		&updatingSale.BranchID,
		&updatingSale.TurnID,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingSale.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

//

func NewSaleProduct(saleProduct *models.SaleProduct) error {
	db := database.Connect()
	defer db.Close()

	tableName := "sale_products"
	currentDate := time.Now()
	saleProduct.Id = 0 // To avoid the id to be forced
	saleProduct.CreatedAt = &currentDate
	saleProduct.UpdatedAt = &currentDate

	query, data, err := commons.GetQuery(tableName, *saleProduct, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&saleProduct.Id,
		&saleProduct.CreatedAt,
		&saleProduct.UpdatedAt,
		&saleProduct.DeletedAt,
		&saleProduct.Done,
		&saleProduct.GrammageQuantity,
		&saleProduct.KilogrammagePrice,
		&saleProduct.UnitQuantity,
		&saleProduct.UnitPrice,
		&saleProduct.Discount,
		&saleProduct.Tax,
		&saleProduct.SaleID,
		&saleProduct.ProductID,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())

	}
	if saleProduct.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetSaleProducts(root bool) ([]models.SaleProduct, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "sale_products"
	query, _, err := commons.GetQuery(tableName, models.SaleProduct{}, "SELECT", false)
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

	var saleProducts []models.SaleProduct

	for rows.Next() {
		var saleProduct models.SaleProduct
		err := rows.Scan(
			&saleProduct.Id,
			&saleProduct.CreatedAt,
			&saleProduct.UpdatedAt,
			&saleProduct.DeletedAt,
			&saleProduct.Done,
			&saleProduct.GrammageQuantity,
			&saleProduct.KilogrammagePrice,
			&saleProduct.UnitQuantity,
			&saleProduct.UnitPrice,
			&saleProduct.Discount,
			&saleProduct.Tax,
			&saleProduct.SaleID,
			&saleProduct.ProductID,
		)
		if err != nil {
			return nil, fmt.Errorf("can't scan the %s query ERROR: %s", tableName, err.Error())
		}
		saleProducts = append(saleProducts, saleProduct)
	}
	if len(saleProducts) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}
	return saleProducts, nil
}

func GetSaleProduct(id uint, root bool) (models.SaleProduct, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "sale_products"
	saleProduct := models.SaleProduct{}
	query, _, err := commons.GetQuery(tableName, saleProduct, "SELECT", false)
	if err != nil {
		return saleProduct, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	err = db.QueryRow(query, id).Scan(
		&saleProduct.Id,
		&saleProduct.CreatedAt,
		&saleProduct.UpdatedAt,
		&saleProduct.DeletedAt,
		&saleProduct.Done,
		&saleProduct.GrammageQuantity,
		&saleProduct.KilogrammagePrice,
		&saleProduct.UnitQuantity,
		&saleProduct.UnitPrice,
		&saleProduct.Discount,
		&saleProduct.Tax,
		&saleProduct.SaleID,
		&saleProduct.ProductID,
	)
	if err != nil {
		return saleProduct, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if saleProduct.Id == 0 {
		return saleProduct, fmt.Errorf("%s not found", tableName)
	}
	return saleProduct, nil
}

func DeleteSaleProduct(id uint, root bool) error {
	return commons.DeleteFromTableById("sale_products", id, root)
}

func UpdateSaleProduct(updatingSaleProduct *models.SaleProduct, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "sale_products"
	currentDate := time.Now()

	if !root {
		updatingSaleProduct.UpdatedAt = &currentDate
		updatingSaleProduct.CreatedAt = nil // Not allowed to update this field
		updatingSaleProduct.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingSaleProduct.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingSaleProduct, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND deleted_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingSaleProduct.Id,
		&updatingSaleProduct.CreatedAt,
		&updatingSaleProduct.UpdatedAt,
		&updatingSaleProduct.DeletedAt,
		&updatingSaleProduct.Done,
		&updatingSaleProduct.GrammageQuantity,
		&updatingSaleProduct.KilogrammagePrice,
		&updatingSaleProduct.UnitQuantity,
		&updatingSaleProduct.UnitPrice,
		&updatingSaleProduct.Discount,
		&updatingSaleProduct.Tax,
		&updatingSaleProduct.SaleID,
		&updatingSaleProduct.ProductID,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingSaleProduct.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

//

func NewSaleIncome(saleIncome *models.SaleIncome) error {
	db := database.Connect()
	defer db.Close()

	tableName := "sale_incomes"
	currentDate := time.Now()
	saleIncome.Id = 0 // To avoid the id to be forced
	saleIncome.CreatedAt = &currentDate
	saleIncome.UpdatedAt = &currentDate

	query, data, err := commons.GetQuery(tableName, *saleIncome, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&saleIncome.Id,
		&saleIncome.CreatedAt,
		&saleIncome.UpdatedAt,
		&saleIncome.DeletedAt,
		&saleIncome.SaleID,
		&saleIncome.IncomeID,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())

	}
	if saleIncome.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetSaleIncomes(root bool) ([]models.SaleIncome, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "sale_incomes"
	query, _, err := commons.GetQuery(tableName, models.SaleIncome{}, "SELECT", false)
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

	var saleIncomes []models.SaleIncome

	for rows.Next() {
		var saleIncome models.SaleIncome
		err := rows.Scan(
			&saleIncome.Id,
			&saleIncome.CreatedAt,
			&saleIncome.UpdatedAt,
			&saleIncome.DeletedAt,
			&saleIncome.SaleID,
			&saleIncome.IncomeID,
		)
		if err != nil {
			return nil, fmt.Errorf("can't scan the %s query ERROR: %s", tableName, err.Error())
		}
		saleIncomes = append(saleIncomes, saleIncome)
	}
	if len(saleIncomes) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}
	return saleIncomes, nil
}

func GetSaleIncome(id uint, root bool) (models.SaleIncome, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "sale_incomes"
	saleIncome := models.SaleIncome{}
	query, _, err := commons.GetQuery(tableName, saleIncome, "SELECT", false)
	if err != nil {
		return saleIncome, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	err = db.QueryRow(query, id).Scan(
		&saleIncome.Id,
		&saleIncome.CreatedAt,
		&saleIncome.UpdatedAt,
		&saleIncome.DeletedAt,
		&saleIncome.SaleID,
		&saleIncome.IncomeID,
	)
	if err != nil {
		return saleIncome, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if saleIncome.Id == 0 {
		return saleIncome, fmt.Errorf("%s not found", tableName)
	}
	return saleIncome, nil
}

func DeleteSaleIncome(id uint, root bool) error {
	return commons.DeleteFromTableById("sale_incomes", id, root)
}

func UpdateSaleIncome(updatingSaleIncome *models.SaleIncome, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "sale_incomes"
	currentDate := time.Now()

	if !root {
		updatingSaleIncome.UpdatedAt = &currentDate
		updatingSaleIncome.CreatedAt = nil // Not allowed to update this field
		updatingSaleIncome.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingSaleIncome.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingSaleIncome, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND deleted_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingSaleIncome.Id,
		&updatingSaleIncome.CreatedAt,
		&updatingSaleIncome.UpdatedAt,
		&updatingSaleIncome.DeletedAt,
		&updatingSaleIncome.SaleID,
		&updatingSaleIncome.IncomeID,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingSaleIncome.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

//

func NewSaleExpense(saleExpense *models.SaleExpense) error {
	db := database.Connect()
	defer db.Close()

	tableName := "sale_expenses"
	currentDate := time.Now()
	saleExpense.Id = 0 // To avoid the id to be forced
	saleExpense.CreatedAt = &currentDate
	saleExpense.UpdatedAt = &currentDate

	query, data, err := commons.GetQuery(tableName, *saleExpense, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&saleExpense.Id,
		&saleExpense.CreatedAt,
		&saleExpense.UpdatedAt,
		&saleExpense.DeletedAt,
		&saleExpense.SaleID,
		&saleExpense.ExpenseID,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())

	}
	if saleExpense.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetSaleExpenses(root bool) ([]models.SaleExpense, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "sale_expenses"
	query, _, err := commons.GetQuery(tableName, models.SaleExpense{}, "SELECT", false)
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

	var saleExpenses []models.SaleExpense

	for rows.Next() {
		var saleExpense models.SaleExpense
		err := rows.Scan(
			&saleExpense.Id,
			&saleExpense.CreatedAt,
			&saleExpense.UpdatedAt,
			&saleExpense.DeletedAt,
			&saleExpense.SaleID,
			&saleExpense.ExpenseID,
		)
		if err != nil {
			return nil, fmt.Errorf("can't scan the %s query ERROR: %s", tableName, err.Error())
		}
		saleExpenses = append(saleExpenses, saleExpense)
	}
	if len(saleExpenses) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}
	return saleExpenses, nil
}

func GetSaleExpense(id uint, root bool) (models.SaleExpense, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "sale_expenses"
	saleExpense := models.SaleExpense{}
	query, _, err := commons.GetQuery(tableName, saleExpense, "SELECT", false)
	if err != nil {
		return saleExpense, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	err = db.QueryRow(query, id).Scan(
		&saleExpense.Id,
		&saleExpense.CreatedAt,
		&saleExpense.UpdatedAt,
		&saleExpense.DeletedAt,
		&saleExpense.SaleID,
		&saleExpense.ExpenseID,
	)
	if err != nil {
		return saleExpense, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if saleExpense.Id == 0 {
		return saleExpense, fmt.Errorf("%s not found", tableName)
	}
	return saleExpense, nil
}

func DeleteSaleExpense(id uint, root bool) error {
	return commons.DeleteFromTableById("sale_expenses", id, root)
}

func UpdateSaleExpense(updatingSaleExpense *models.SaleExpense, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "sale_expenses"
	currentDate := time.Now()

	if !root {
		updatingSaleExpense.UpdatedAt = &currentDate
		updatingSaleExpense.CreatedAt = nil // Not allowed to update this field
		updatingSaleExpense.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingSaleExpense.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingSaleExpense, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND deleted_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingSaleExpense.Id,
		&updatingSaleExpense.CreatedAt,
		&updatingSaleExpense.UpdatedAt,
		&updatingSaleExpense.DeletedAt,
		&updatingSaleExpense.SaleID,
		&updatingSaleExpense.ExpenseID,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingSaleExpense.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

//

func NewSaleArgument(saleArgument *models.SaleArgument) error {
	db := database.Connect()
	defer db.Close()

	tableName := "sale_arguments"
	currentDate := time.Now()
	saleArgument.Id = 0 // To avoid the id to be forced
	saleArgument.CreatedAt = &currentDate
	saleArgument.UpdatedAt = &currentDate

	query, data, err := commons.GetQuery(tableName, *saleArgument, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&saleArgument.Id,
		&saleArgument.CreatedAt,
		&saleArgument.UpdatedAt,
		&saleArgument.DeletedAt,
		&saleArgument.SaleID,
		&saleArgument.ArgumentID,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())

	}
	if saleArgument.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetSaleArguments(root bool) ([]models.SaleArgument, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "sale_arguments"
	query, _, err := commons.GetQuery(tableName, models.SaleArgument{}, "SELECT", false)
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

	var saleArguments []models.SaleArgument

	for rows.Next() {
		var saleArgument models.SaleArgument
		err := rows.Scan(
			&saleArgument.Id,
			&saleArgument.CreatedAt,
			&saleArgument.UpdatedAt,
			&saleArgument.DeletedAt,
			&saleArgument.SaleID,
			&saleArgument.ArgumentID,
		)
		if err != nil {
			return nil, fmt.Errorf("can't scan the %s query ERROR: %s", tableName, err.Error())
		}
		saleArguments = append(saleArguments, saleArgument)
	}
	if len(saleArguments) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}
	return saleArguments, nil
}

func GetSaleArgument(id uint, root bool) (models.SaleArgument, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "sale_arguments"
	saleArgument := models.SaleArgument{}
	query, _, err := commons.GetQuery(tableName, saleArgument, "SELECT", false)
	if err != nil {
		return saleArgument, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	err = db.QueryRow(query, id).Scan(
		&saleArgument.Id,
		&saleArgument.CreatedAt,
		&saleArgument.UpdatedAt,
		&saleArgument.DeletedAt,
		&saleArgument.SaleID,
		&saleArgument.ArgumentID,
	)
	if err != nil {
		return saleArgument, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if saleArgument.Id == 0 {
		return saleArgument, fmt.Errorf("%s not found", tableName)
	}
	return saleArgument, nil
}

func DeleteSaleArgument(id uint, root bool) error {
	return commons.DeleteFromTableById("sale_arguments", id, root)
}

func UpdateSaleArgument(updatingSaleArgument *models.SaleArgument, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "sale_arguments"
	currentDate := time.Now()

	if !root {
		updatingSaleArgument.UpdatedAt = &currentDate
		updatingSaleArgument.CreatedAt = nil // Not allowed to update this field
		updatingSaleArgument.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingSaleArgument.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingSaleArgument, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND deleted_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingSaleArgument.Id,
		&updatingSaleArgument.CreatedAt,
		&updatingSaleArgument.UpdatedAt,
		&updatingSaleArgument.DeletedAt,
		&updatingSaleArgument.SaleID,
		&updatingSaleArgument.ArgumentID,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingSaleArgument.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

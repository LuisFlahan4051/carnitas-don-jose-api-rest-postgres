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

func NewSupply(supply *models.Supply) error {
	db := database.Connect()
	defer db.Close()

	tableName := "supplies"
	currentDate := time.Now()
	supply.Id = 0 // To avoid the id to be forced
	supply.CreatedAt = &currentDate
	supply.UpdatedAt = &currentDate

	query, data, err := commons.GetQuery(tableName, *supply, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&supply.Id,
		&supply.CreatedAt,
		&supply.UpdatedAt,
		&supply.DeletedAt,
		&supply.Name,
		&supply.Description,
		&supply.Photo,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())

	}
	if supply.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetSupplies(root bool) ([]models.Supply, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "supplies"
	query, _, err := commons.GetQuery(tableName, models.Supply{}, "SELECT", false)
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

	var supplies []models.Supply

	for rows.Next() {
		var supply models.Supply
		err := rows.Scan(
			&supply.Id,
			&supply.CreatedAt,
			&supply.UpdatedAt,
			&supply.DeletedAt,
			&supply.Name,
			&supply.Description,
			&supply.Photo,
		)
		if err != nil {
			return nil, fmt.Errorf("can't scan the %s query ERROR: %s", tableName, err.Error())
		}
		supplies = append(supplies, supply)
	}
	if len(supplies) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}
	return supplies, nil
}

func GetSupply(id uint, root bool) (models.Supply, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "supplies"
	supply := models.Supply{}
	query, _, err := commons.GetQuery(tableName, supply, "SELECT", false)
	if err != nil {
		return supply, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	err = db.QueryRow(query, id).Scan(
		&supply.Id,
		&supply.CreatedAt,
		&supply.UpdatedAt,
		&supply.DeletedAt,
		&supply.Name,
		&supply.Description,
		&supply.Photo,
	)
	if err != nil {
		return supply, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if supply.Id == 0 {
		return supply, fmt.Errorf("%s not found", tableName)
	}
	return supply, nil
}

func DeleteSupply(id uint, root bool) error {
	return commons.DeleteFromTableById("supplies", id, root)
}

func UpdateSupply(updatingSupply *models.Supply, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "supplies"
	currentDate := time.Now()

	if !root {
		updatingSupply.UpdatedAt = &currentDate
		updatingSupply.CreatedAt = nil // Not allowed to update this field
		updatingSupply.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingSupply.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingSupply, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND deleted_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingSupply.Id,
		&updatingSupply.CreatedAt,
		&updatingSupply.UpdatedAt,
		&updatingSupply.DeletedAt,
		&updatingSupply.Name,
		&updatingSupply.Description,
		&updatingSupply.Photo,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingSupply.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

//

func NewArticle(article *models.Article) error {
	db := database.Connect()
	defer db.Close()

	tableName := "articles"
	currentDate := time.Now()
	article.Id = 0 // To avoid the id to be forced
	article.CreatedAt = &currentDate
	article.UpdatedAt = &currentDate

	query, data, err := commons.GetQuery(tableName, *article, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&article.Id,
		&article.CreatedAt,
		&article.UpdatedAt,
		&article.DeletedAt,
		&article.Name,
		&article.Description,
		&article.Photo,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())

	}
	if article.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetArticles(root bool) ([]models.Article, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "articles"
	query, _, err := commons.GetQuery(tableName, models.Article{}, "SELECT", false)
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

	var articles []models.Article

	for rows.Next() {
		var article models.Article
		err := rows.Scan(
			&article.Id,
			&article.CreatedAt,
			&article.UpdatedAt,
			&article.DeletedAt,
			&article.Name,
			&article.Description,
			&article.Photo,
		)
		if err != nil {
			return nil, fmt.Errorf("can't scan the %s query ERROR: %s", tableName, err.Error())
		}
		articles = append(articles, article)
	}
	if len(articles) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}
	return articles, nil
}

func GetArticle(id uint, root bool) (models.Article, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "articles"
	article := models.Article{}
	query, _, err := commons.GetQuery(tableName, article, "SELECT", false)
	if err != nil {
		return article, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	err = db.QueryRow(query, id).Scan(
		&article.Id,
		&article.CreatedAt,
		&article.UpdatedAt,
		&article.DeletedAt,
		&article.Name,
		&article.Description,
		&article.Photo,
	)
	if err != nil {
		return article, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if article.Id == 0 {
		return article, fmt.Errorf("%s not found", tableName)
	}
	return article, nil
}

func DeleteArticle(id uint, root bool) error {
	return commons.DeleteFromTableById("articles", id, root)
}

func UpdateArticle(updatingArticle *models.Article, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "articles"
	currentDate := time.Now()

	if !root {
		updatingArticle.UpdatedAt = &currentDate
		updatingArticle.CreatedAt = nil // Not allowed to update this field
		updatingArticle.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingArticle.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingArticle, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND deleted_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingArticle.Id,
		&updatingArticle.CreatedAt,
		&updatingArticle.UpdatedAt,
		&updatingArticle.DeletedAt,
		&updatingArticle.Name,
		&updatingArticle.Description,
		&updatingArticle.Photo,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingArticle.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

//

func NewSafebox(safebox *models.Safebox) error {
	db := database.Connect()
	defer db.Close()

	tableName := "safeboxes"
	currentDate := time.Now()
	safebox.Id = 0 // To avoid the id to be forced
	safebox.CreatedAt = &currentDate
	safebox.UpdatedAt = &currentDate

	query, data, err := commons.GetQuery(tableName, *safebox, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&safebox.Id,
		&safebox.CreatedAt,
		&safebox.UpdatedAt,
		&safebox.DeletedAt,
		&safebox.Cents10,
		&safebox.Cents50,
		&safebox.Coins1,
		&safebox.Coins2,
		&safebox.Coins5,
		&safebox.Coins10,
		&safebox.Coins20,
		&safebox.Bills20,
		&safebox.Bills50,
		&safebox.Bills100,
		&safebox.Bills200,
		&safebox.Bills500,
		&safebox.Bills1000,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())

	}
	if safebox.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetSafeboxes(root bool) ([]models.Safebox, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "safeboxes"
	query, _, err := commons.GetQuery(tableName, models.Safebox{}, "SELECT", false)
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

	var safeboxes []models.Safebox

	for rows.Next() {
		var safebox models.Safebox
		err := rows.Scan(
			&safebox.Id,
			&safebox.CreatedAt,
			&safebox.UpdatedAt,
			&safebox.DeletedAt,
			&safebox.Cents10,
			&safebox.Cents50,
			&safebox.Coins1,
			&safebox.Coins2,
			&safebox.Coins5,
			&safebox.Coins10,
			&safebox.Coins20,
			&safebox.Bills20,
			&safebox.Bills50,
			&safebox.Bills100,
			&safebox.Bills200,
			&safebox.Bills500,
			&safebox.Bills1000,
		)
		if err != nil {
			return nil, fmt.Errorf("can't scan the %s query ERROR: %s", tableName, err.Error())
		}
		safeboxes = append(safeboxes, safebox)
	}
	if len(safeboxes) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}
	return safeboxes, nil
}

func GetSafebox(id uint, root bool) (models.Safebox, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "safeboxes"
	safebox := models.Safebox{}
	query, _, err := commons.GetQuery(tableName, safebox, "SELECT", false)
	if err != nil {
		return safebox, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	err = db.QueryRow(query, id).Scan(
		&safebox.Id,
		&safebox.CreatedAt,
		&safebox.UpdatedAt,
		&safebox.DeletedAt,
		&safebox.Cents10,
		&safebox.Cents50,
		&safebox.Coins1,
		&safebox.Coins2,
		&safebox.Coins5,
		&safebox.Coins10,
		&safebox.Coins20,
		&safebox.Bills20,
		&safebox.Bills50,
		&safebox.Bills100,
		&safebox.Bills200,
		&safebox.Bills500,
		&safebox.Bills1000,
	)
	if err != nil {
		return safebox, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if safebox.Id == 0 {
		return safebox, fmt.Errorf("%s not found", tableName)
	}
	return safebox, nil
}

func DeleteSafebox(id uint, root bool) error {
	return commons.DeleteFromTableById("safeboxes", id, root)
}

func UpdateSafebox(updatingSafebox *models.Safebox, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "safeboxes"
	currentDate := time.Now()

	if !root {
		updatingSafebox.UpdatedAt = &currentDate
		updatingSafebox.CreatedAt = nil // Not allowed to update this field
		updatingSafebox.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingSafebox.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingSafebox, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND deleted_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingSafebox.Id,
		&updatingSafebox.CreatedAt,
		&updatingSafebox.UpdatedAt,
		&updatingSafebox.DeletedAt,
		&updatingSafebox.Cents10,
		&updatingSafebox.Cents50,
		&updatingSafebox.Coins1,
		&updatingSafebox.Coins2,
		&updatingSafebox.Coins5,
		&updatingSafebox.Coins10,
		&updatingSafebox.Coins20,
		&updatingSafebox.Bills20,
		&updatingSafebox.Bills50,
		&updatingSafebox.Bills100,
		&updatingSafebox.Bills200,
		&updatingSafebox.Bills500,
		&updatingSafebox.Bills1000,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingSafebox.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

//

func NewIncome(income *models.Income) error {
	db := database.Connect()
	defer db.Close()

	tableName := "incomes"
	currentDate := time.Now()
	income.Id = 0 // To avoid the id to be forced
	income.CreatedAt = &currentDate
	income.UpdatedAt = &currentDate

	query, data, err := commons.GetQuery(tableName, *income, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&income.Id,
		&income.CreatedAt,
		&income.UpdatedAt,
		&income.DeletedAt,
		&income.Reason,
		&income.Income,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())

	}
	if income.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetIncomes(root bool) ([]models.Income, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "incomes"
	query, _, err := commons.GetQuery(tableName, models.Income{}, "SELECT", false)
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

	var incomes []models.Income

	for rows.Next() {
		var income models.Income
		err := rows.Scan(
			&income.Id,
			&income.CreatedAt,
			&income.UpdatedAt,
			&income.DeletedAt,
			&income.Reason,
			&income.Income,
		)
		if err != nil {
			return nil, fmt.Errorf("can't scan the %s query ERROR: %s", tableName, err.Error())
		}
		incomes = append(incomes, income)
	}
	if len(incomes) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}
	return incomes, nil
}

func GetIncome(id uint, root bool) (models.Income, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "incomes"
	income := models.Income{}
	query, _, err := commons.GetQuery(tableName, income, "SELECT", false)
	if err != nil {
		return income, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	err = db.QueryRow(query, id).Scan(
		&income.Id,
		&income.CreatedAt,
		&income.UpdatedAt,
		&income.DeletedAt,
		&income.Reason,
		&income.Income,
	)
	if err != nil {
		return income, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if income.Id == 0 {
		return income, fmt.Errorf("%s not found", tableName)
	}
	return income, nil
}

func DeleteIncome(id uint, root bool) error {
	return commons.DeleteFromTableById("incomes", id, root)
}

func UpdateIncome(updatingIncome *models.Income, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "incomes"
	currentDate := time.Now()

	if !root {
		updatingIncome.UpdatedAt = &currentDate
		updatingIncome.CreatedAt = nil // Not allowed to update this field
		updatingIncome.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingIncome.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingIncome, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND deleted_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingIncome.Id,
		&updatingIncome.CreatedAt,
		&updatingIncome.UpdatedAt,
		&updatingIncome.DeletedAt,
		&updatingIncome.Reason,
		&updatingIncome.Income,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingIncome.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

//

func NewExpense(expense *models.Expense) error {
	db := database.Connect()
	defer db.Close()

	tableName := "expenses"
	currentDate := time.Now()
	expense.Id = 0 // To avoid the id to be forced
	expense.CreatedAt = &currentDate
	expense.UpdatedAt = &currentDate

	query, data, err := commons.GetQuery(tableName, *expense, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&expense.Id,
		&expense.CreatedAt,
		&expense.UpdatedAt,
		&expense.DeletedAt,
		&expense.Reason,
		&expense.Expense,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())

	}
	if expense.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetExpenses(root bool) ([]models.Expense, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "expenses"
	query, _, err := commons.GetQuery(tableName, models.Expense{}, "SELECT", false)
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

	var expenses []models.Expense

	for rows.Next() {
		var expense models.Expense
		err := rows.Scan(
			&expense.Id,
			&expense.CreatedAt,
			&expense.UpdatedAt,
			&expense.DeletedAt,
			&expense.Reason,
			&expense.Expense,
		)
		if err != nil {
			return nil, fmt.Errorf("can't scan the %s query ERROR: %s", tableName, err.Error())
		}
		expenses = append(expenses, expense)
	}
	if len(expenses) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}
	return expenses, nil
}

func GetExpense(id uint, root bool) (models.Expense, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "expenses"
	expense := models.Expense{}
	query, _, err := commons.GetQuery(tableName, expense, "SELECT", false)
	if err != nil {
		return expense, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	err = db.QueryRow(query, id).Scan(
		&expense.Id,
		&expense.CreatedAt,
		&expense.UpdatedAt,
		&expense.DeletedAt,
		&expense.Reason,
		&expense.Expense,
	)
	if err != nil {
		return expense, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if expense.Id == 0 {
		return expense, fmt.Errorf("%s not found", tableName)
	}
	return expense, nil
}

func DeleteExpense(id uint, root bool) error {
	return commons.DeleteFromTableById("expenses", id, root)
}

func UpdateExpense(updatingExpense *models.Expense, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "expenses"
	currentDate := time.Now()

	if !root {
		updatingExpense.UpdatedAt = &currentDate
		updatingExpense.CreatedAt = nil // Not allowed to update this field
		updatingExpense.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingExpense.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingExpense, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND deleted_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingExpense.Id,
		&updatingExpense.CreatedAt,
		&updatingExpense.UpdatedAt,
		&updatingExpense.DeletedAt,
		&updatingExpense.Reason,
		&updatingExpense.Expense,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingExpense.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

//

func NewArgument(argument *models.Argument) error {
	db := database.Connect()
	defer db.Close()

	tableName := "arguments"
	currentDate := time.Now()
	argument.Id = 0 // To avoid the id to be forced
	argument.CreatedAt = &currentDate
	argument.UpdatedAt = &currentDate

	query, data, err := commons.GetQuery(tableName, *argument, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&argument.Id,
		&argument.CreatedAt,
		&argument.UpdatedAt,
		&argument.DeletedAt,
		&argument.Complaint,
		&argument.Score,
		&argument.Argument,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())

	}
	if argument.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetArguments(root bool) ([]models.Argument, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "arguments"
	query, _, err := commons.GetQuery(tableName, models.Argument{}, "SELECT", false)
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

	var arguments []models.Argument

	for rows.Next() {
		var argument models.Argument
		err := rows.Scan(
			&argument.Id,
			&argument.CreatedAt,
			&argument.UpdatedAt,
			&argument.DeletedAt,
			&argument.Complaint,
			&argument.Score,
			&argument.Argument,
		)
		if err != nil {
			return nil, fmt.Errorf("can't scan the %s query ERROR: %s", tableName, err.Error())
		}
		arguments = append(arguments, argument)
	}
	if len(arguments) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}
	return arguments, nil
}

func GetArgument(id uint, root bool) (models.Argument, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "arguments"
	argument := models.Argument{}
	query, _, err := commons.GetQuery(tableName, argument, "SELECT", false)
	if err != nil {
		return argument, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	err = db.QueryRow(query, id).Scan(
		&argument.Id,
		&argument.CreatedAt,
		&argument.UpdatedAt,
		&argument.DeletedAt,
		&argument.Complaint,
		&argument.Score,
		&argument.Argument,
	)
	if err != nil {
		return argument, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if argument.Id == 0 {
		return argument, fmt.Errorf("%s not found", tableName)
	}
	return argument, nil
}

func DeleteArgument(id uint, root bool) error {
	return commons.DeleteFromTableById("arguments", id, root)
}

func UpdateArgument(updatingArgument *models.Argument, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "arguments"
	currentDate := time.Now()

	if !root {
		updatingArgument.UpdatedAt = &currentDate
		updatingArgument.CreatedAt = nil // Not allowed to update this field
		updatingArgument.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingArgument.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingArgument, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND deleted_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingArgument.Id,
		&updatingArgument.CreatedAt,
		&updatingArgument.UpdatedAt,
		&updatingArgument.DeletedAt,
		&updatingArgument.Complaint,
		&updatingArgument.Score,
		&updatingArgument.Argument,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingArgument.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

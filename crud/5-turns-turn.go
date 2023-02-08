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

func NewTurn(turn *models.Turn) error {
	db := database.Connect()
	defer db.Close()

	tableName := "turns"
	currentDate := time.Now()
	turn.Id = 0 // To avoid the id to be forced
	turn.CreatedAt = &currentDate
	turn.UpdatedAt = &currentDate

	query, data, err := commons.GetQuery(tableName, *turn, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&turn.Id,
		&turn.CreatedAt,
		&turn.UpdatedAt,
		&turn.DeletedAt,
		&turn.StartDate,
		&turn.EndDate,
		&turn.Active,
		&turn.IncomesCounter,
		&turn.NetIncomesCounter,
		&turn.ExpensesCounter,
		&turn.UserID,
		&turn.BranchID,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())

	}
	if turn.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetTurns(root bool) ([]models.Turn, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "turns"
	query, _, err := commons.GetQuery(tableName, models.Turn{}, "SELECT", false)
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

	var turns []models.Turn

	for rows.Next() {
		var turn models.Turn
		err := rows.Scan(
			&turn.Id,
			&turn.CreatedAt,
			&turn.UpdatedAt,
			&turn.DeletedAt,
			&turn.StartDate,
			&turn.EndDate,
			&turn.Active,
			&turn.IncomesCounter,
			&turn.NetIncomesCounter,
			&turn.ExpensesCounter,
			&turn.UserID,
			&turn.BranchID,
		)
		if err != nil {
			return nil, fmt.Errorf("can't scan the %s query ERROR: %s", tableName, err.Error())
		}
		turns = append(turns, turn)
	}
	if len(turns) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}
	return turns, nil
}

func GetTurn(id uint, root bool) (models.Turn, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "turns"
	turn := models.Turn{}
	query, _, err := commons.GetQuery(tableName, turn, "SELECT", false)
	if err != nil {
		return turn, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	err = db.QueryRow(query, id).Scan(
		&turn.Id,
		&turn.CreatedAt,
		&turn.UpdatedAt,
		&turn.DeletedAt,
		&turn.StartDate,
		&turn.EndDate,
		&turn.Active,
		&turn.IncomesCounter,
		&turn.NetIncomesCounter,
		&turn.ExpensesCounter,
		&turn.UserID,
		&turn.BranchID,
	)
	if err != nil {
		return turn, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if turn.Id == 0 {
		return turn, fmt.Errorf("%s not found", tableName)
	}
	return turn, nil
}

func DeleteTurn(id uint, root bool) error {
	return commons.DeleteFromTableById("turns", id, root)
}

func UpdateTurn(updatingTurn *models.Turn, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "turns"
	currentDate := time.Now()

	if !root {
		updatingTurn.UpdatedAt = &currentDate
		updatingTurn.CreatedAt = nil // Not allowed to update this field
		updatingTurn.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingTurn.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingTurn, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND delete_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingTurn.Id,
		&updatingTurn.CreatedAt,
		&updatingTurn.UpdatedAt,
		&updatingTurn.DeletedAt,
		&updatingTurn.StartDate,
		&updatingTurn.EndDate,
		&updatingTurn.Active,
		&updatingTurn.IncomesCounter,
		&updatingTurn.NetIncomesCounter,
		&updatingTurn.ExpensesCounter,
		&updatingTurn.UserID,
		&updatingTurn.BranchID,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingTurn.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

//

func NewTurnUserRole(turnUserRole *models.TurnUserRole) error {
	db := database.Connect()
	defer db.Close()

	tableName := "turn_user_roles"
	currentDate := time.Now()
	turnUserRole.Id = 0 // To avoid the id to be forced
	turnUserRole.CreatedAt = &currentDate
	turnUserRole.UpdatedAt = &currentDate

	query, data, err := commons.GetQuery(tableName, *turnUserRole, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&turnUserRole.Id,
		&turnUserRole.CreatedAt,
		&turnUserRole.UpdatedAt,
		&turnUserRole.DeletedAt,
		&turnUserRole.LoginDate,
		&turnUserRole.LogoutDate,
		&turnUserRole.UserID,
		&turnUserRole.TurnID,
		&turnUserRole.RoleID,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())

	}
	if turnUserRole.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetTurnUserRoles(root bool) ([]models.TurnUserRole, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "turn_user_roles"
	query, _, err := commons.GetQuery(tableName, models.TurnUserRole{}, "SELECT", false)
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

	var turnUserRoles []models.TurnUserRole

	for rows.Next() {
		var turnUserRole models.TurnUserRole
		err := rows.Scan(
			&turnUserRole.Id,
			&turnUserRole.CreatedAt,
			&turnUserRole.UpdatedAt,
			&turnUserRole.DeletedAt,
			&turnUserRole.LoginDate,
			&turnUserRole.LogoutDate,
			&turnUserRole.UserID,
			&turnUserRole.TurnID,
			&turnUserRole.RoleID,
		)
		if err != nil {
			return nil, fmt.Errorf("can't scan the %s query ERROR: %s", tableName, err.Error())
		}
		turnUserRoles = append(turnUserRoles, turnUserRole)
	}
	if len(turnUserRoles) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}
	return turnUserRoles, nil
}

func GetTurnUserRole(id uint, root bool) (models.TurnUserRole, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "turn_user_roles"
	turnUserRole := models.TurnUserRole{}
	query, _, err := commons.GetQuery(tableName, turnUserRole, "SELECT", false)
	if err != nil {
		return turnUserRole, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	err = db.QueryRow(query, id).Scan(
		&turnUserRole.Id,
		&turnUserRole.CreatedAt,
		&turnUserRole.UpdatedAt,
		&turnUserRole.DeletedAt,
		&turnUserRole.LoginDate,
		&turnUserRole.LogoutDate,
		&turnUserRole.UserID,
		&turnUserRole.TurnID,
		&turnUserRole.RoleID,
	)
	if err != nil {
		return turnUserRole, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if turnUserRole.Id == 0 {
		return turnUserRole, fmt.Errorf("%s not found", tableName)
	}
	return turnUserRole, nil
}

func DeleteTurnUserRole(id uint, root bool) error {
	return commons.DeleteFromTableById("turn_user_roles", id, root)
}

func UpdateTurnUserRole(updatingTurnUserRole *models.TurnUserRole, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "turn_user_roles"
	currentDate := time.Now()

	if !root {
		updatingTurnUserRole.UpdatedAt = &currentDate
		updatingTurnUserRole.CreatedAt = nil // Not allowed to update this field
		updatingTurnUserRole.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingTurnUserRole.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingTurnUserRole, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND delete_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingTurnUserRole.Id,
		&updatingTurnUserRole.CreatedAt,
		&updatingTurnUserRole.UpdatedAt,
		&updatingTurnUserRole.DeletedAt,
		&updatingTurnUserRole.LoginDate,
		&updatingTurnUserRole.LogoutDate,
		&updatingTurnUserRole.UserID,
		&updatingTurnUserRole.TurnID,
		&updatingTurnUserRole.RoleID,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingTurnUserRole.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

//

func NewTurnSafebox(turnSafebox *models.TurnSafebox) error {
	db := database.Connect()
	defer db.Close()

	tableName := "turn_safebox"
	currentDate := time.Now()
	turnSafebox.Id = 0 // To avoid the id to be forced
	turnSafebox.CreatedAt = &currentDate
	turnSafebox.UpdatedAt = &currentDate

	query, data, err := commons.GetQuery(tableName, *turnSafebox, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&turnSafebox.Id,
		&turnSafebox.CreatedAt,
		&turnSafebox.UpdatedAt,
		&turnSafebox.DeletedAt,
		&turnSafebox.TurnID,
		&turnSafebox.SafeboxID,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())

	}
	if turnSafebox.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetTurnsSafebox(root bool) ([]models.TurnSafebox, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "turn_safebox"
	query, _, err := commons.GetQuery(tableName, models.TurnSafebox{}, "SELECT", false)
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

	var turnSafeboxes []models.TurnSafebox

	for rows.Next() {
		var turnSafebox models.TurnSafebox
		err := rows.Scan(
			&turnSafebox.Id,
			&turnSafebox.CreatedAt,
			&turnSafebox.UpdatedAt,
			&turnSafebox.DeletedAt,
			&turnSafebox.TurnID,
			&turnSafebox.SafeboxID,
		)
		if err != nil {
			return nil, fmt.Errorf("can't scan the %s query ERROR: %s", tableName, err.Error())
		}
		turnSafeboxes = append(turnSafeboxes, turnSafebox)
	}
	if len(turnSafeboxes) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}
	return turnSafeboxes, nil
}

func GetTurnSafebox(id uint, root bool) (models.TurnSafebox, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "turn_safebox"
	turnSafebox := models.TurnSafebox{}
	query, _, err := commons.GetQuery(tableName, turnSafebox, "SELECT", false)
	if err != nil {
		return turnSafebox, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	err = db.QueryRow(query, id).Scan(
		&turnSafebox.Id,
		&turnSafebox.CreatedAt,
		&turnSafebox.UpdatedAt,
		&turnSafebox.DeletedAt,
		&turnSafebox.TurnID,
		&turnSafebox.SafeboxID,
	)
	if err != nil {
		return turnSafebox, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if turnSafebox.Id == 0 {
		return turnSafebox, fmt.Errorf("%s not found", tableName)
	}
	return turnSafebox, nil
}

func DeleteTurnSafebox(id uint, root bool) error {
	return commons.DeleteFromTableById("turn_safebox", id, root)
}

func UpdateTurnSafebox(updatingTurnSafebox *models.TurnSafebox, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "turn_safebox"
	currentDate := time.Now()

	if !root {
		updatingTurnSafebox.UpdatedAt = &currentDate
		updatingTurnSafebox.CreatedAt = nil // Not allowed to update this field
		updatingTurnSafebox.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingTurnSafebox.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingTurnSafebox, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND delete_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingTurnSafebox.Id,
		&updatingTurnSafebox.CreatedAt,
		&updatingTurnSafebox.UpdatedAt,
		&updatingTurnSafebox.DeletedAt,
		&updatingTurnSafebox.TurnID,
		&updatingTurnSafebox.SafeboxID,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingTurnSafebox.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

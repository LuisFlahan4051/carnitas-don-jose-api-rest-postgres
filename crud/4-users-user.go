package crud

import (
	"database/sql"
	"errors"

	"time"

	commons "github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/commonFunctions"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/database"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/models"
)

func NewUser(user *models.User) error {
	db := database.Connect()
	defer db.Close()

	user.Username = commons.CleanSpaces(user.Username)
	user.Password = commons.CleanSpaces(user.Password)
	if user.Username == "" || user.Password == "" {
		return errors.New("username and password is required")
	}

	user.Id = 0 // To don't add id field to query
	query, data, err := commons.GetQuery("users", *user, 0, true)
	if err != nil {
		return err
	}

	err = db.QueryRow(query, data...).Scan(
		&user.Id,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
		&user.Name,
		&user.Lastname,
		&user.Username,
		&user.Password,
		&user.Photo,
		&user.Verified,
		&user.Warning,
		&user.Darktheme,
		&user.ActiveContract,
		&user.Address,
		&user.Born,
		&user.DegreeStudy,
		&user.RelationShip,
		&user.Curp,
		&user.Rfc,
		&user.CitizenID,
		&user.CredentialID,
		&user.OriginState,
		&user.Score,
		&user.Qualities,
		&user.Defects,
		&user.OriginBranchID,
		&user.BranchID)
	if err != nil {
		return err
	}
	if user.Id == 0 {
		return errors.New("can't create user")
	}

	return nil
}

func GetUsers(root bool) ([]models.User, error) {

	db := database.Connect()
	defer db.Close()

	query, _, _ := commons.GetQuery("users", models.User{}, 2, false)

	if !root {
		query += " WHERE deleted_at IS NULL"
	}

	var users []models.User

	rows, err := db.Query(query)
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.Id,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt,
			&user.Name,
			&user.Lastname,
			&user.Username,
			&user.Password,
			&user.Photo,
			&user.Verified,
			&user.Warning,
			&user.Darktheme,
			&user.ActiveContract,
			&user.Address,
			&user.Born,
			&user.DegreeStudy,
			&user.RelationShip,
			&user.Curp,
			&user.Rfc,
			&user.CitizenID,
			&user.CredentialID,
			&user.OriginState,
			&user.Score,
			&user.Qualities,
			&user.Defects,
			&user.OriginBranchID,
			&user.BranchID)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	if len(users) == 0 {
		return users, sql.ErrNoRows
	}

	return users, nil
}

func GetUser(id uint, root bool) (models.User, error) {
	var user models.User

	db := database.Connect()
	defer db.Close()

	query, _, _ := commons.GetQuery("users", models.User{}, 2, false)
	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	err := db.QueryRow(query, id).Scan(
		&user.Id,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
		&user.Name,
		&user.Lastname,
		&user.Username,
		&user.Password,
		&user.Photo,
		&user.Verified,
		&user.Warning,
		&user.Darktheme,
		&user.ActiveContract,
		&user.Address,
		&user.Born,
		&user.DegreeStudy,
		&user.RelationShip,
		&user.Curp,
		&user.Rfc,
		&user.CitizenID,
		&user.CredentialID,
		&user.OriginState,
		&user.Score,
		&user.Qualities,
		&user.Defects,
		&user.OriginBranchID,
		&user.BranchID)

	return user, err
}

func userIsDeleted(id uint) bool {
	db := database.Connect()
	defer db.Close()

	var deletedAt *time.Time
	query := "SELECT deleted_at FROM users WHERE id = $1"
	err := db.QueryRow(query, id).Scan(&deletedAt)

	if err != nil {
		return false
	}

	return deletedAt != nil
}

func DeleteUser(id uint, root bool) error {
	db := database.Connect()
	defer db.Close()

	if root {
		query := "DELETE FROM users WHERE id = $1"
		_, err := db.Exec(query, id)
		return err
	}

	if userIsDeleted(id) {
		return errors.New("user already deleted")
	}

	query := "UPDATE users SET deleted_at = $1 WHERE id = $2"
	result, err := db.Exec(query, time.Now(), id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("user doesn't exist")
	}

	return nil
}

func UpdateUser(updatingUser *models.User, root bool) error {
	db := database.Connect()
	defer db.Close()

	currentDate := time.Now()
	updatingUser.UpdatedAt = &currentDate
	updatingUser.CreatedAt = nil
	updatingUser.DeletedAt = nil

	if root {
		query, data, err := commons.GetQuery("users", *updatingUser, 1, true)
		if err != nil {
			return err
		}

		err = db.QueryRow(query, data...).Scan(
			&updatingUser.Id,
			&updatingUser.CreatedAt,
			&updatingUser.UpdatedAt,
			&updatingUser.DeletedAt,
			&updatingUser.Name,
			&updatingUser.Lastname,
			&updatingUser.Username,
			&updatingUser.Password,
			&updatingUser.Photo,
			&updatingUser.Verified,
			&updatingUser.Warning,
			&updatingUser.Darktheme,
			&updatingUser.ActiveContract,
			&updatingUser.Address,
			&updatingUser.Born,
			&updatingUser.DegreeStudy,
			&updatingUser.RelationShip,
			&updatingUser.Curp,
			&updatingUser.Rfc,
			&updatingUser.CitizenID,
			&updatingUser.CredentialID,
			&updatingUser.OriginState,
			&updatingUser.Score,
			&updatingUser.Qualities,
			&updatingUser.Defects,
			&updatingUser.OriginBranchID,
			&updatingUser.BranchID)

		return err
	}

	if userIsDeleted(updatingUser.Id) {
		return errors.New("user is deleted")
	}

	query, data, err := commons.GetQuery("users", *updatingUser, 1, true)

	if err != nil {
		return err
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingUser.Id,
		&updatingUser.CreatedAt,
		&updatingUser.UpdatedAt,
		&updatingUser.DeletedAt,
		&updatingUser.Name,
		&updatingUser.Lastname,
		&updatingUser.Username,
		&updatingUser.Password,
		&updatingUser.Photo,
		&updatingUser.Verified,
		&updatingUser.Warning,
		&updatingUser.Darktheme,
		&updatingUser.ActiveContract,
		&updatingUser.Address,
		&updatingUser.Born,
		&updatingUser.DegreeStudy,
		&updatingUser.RelationShip,
		&updatingUser.Curp,
		&updatingUser.Rfc,
		&updatingUser.CitizenID,
		&updatingUser.CredentialID,
		&updatingUser.OriginState,
		&updatingUser.Score,
		&updatingUser.Qualities,
		&updatingUser.Defects,
		&updatingUser.OriginBranchID,
		&updatingUser.BranchID)

	return err
}

//

func NewRole(role *models.Role) error {
	db := database.Connect()
	defer db.Close()

	if role.AccessLevel <= 1 {
		return errors.New("access level must be greater than 1, new role can't be root")
	}

	query, data, err := commons.GetQuery("roles", *role, 0, true)
	if err != nil {
		return err
	}

	err = db.QueryRow(query, data...).Scan(
		&role.Id,
		&role.CreatedAt,
		&role.UpdatedAt,
		&role.DeletedAt,
		&role.Name,
		&role.AccessLevel)
	if err != nil {
		return err
	}
	if role.Id == 0 {
		return errors.New("can't create role")
	}

	return nil
}

func NewInheritUserRole(inheritUserRole *models.InheritUserRole) error {
	db := database.Connect()
	defer db.Close()

	query, data, err := commons.GetQuery("inherit_user_roles", *inheritUserRole, 0, true)
	if err != nil {
		return err
	}

	err = db.QueryRow(query, data...).Scan(
		&inheritUserRole.Id,
		&inheritUserRole.CreatedAt,
		&inheritUserRole.UpdatedAt,
		&inheritUserRole.DeletedAt,
		&inheritUserRole.UserID,
		&inheritUserRole.RoleID)
	if err != nil {
		return err
	}
	if inheritUserRole.Id == 0 {
		return errors.New("can't create inherit user role")
	}

	return nil
}

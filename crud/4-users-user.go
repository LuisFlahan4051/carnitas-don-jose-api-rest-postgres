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

func NewRole(role *models.Role) error {
	db := database.Connect()
	defer db.Close()

	tableName := "roles"
	currentDate := time.Now()
	role.Id = 0
	role.UpdatedAt = &currentDate
	role.CreatedAt = &currentDate

	query, data, err := commons.GetQuery(tableName, *role, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&role.Id,
		&role.CreatedAt,
		&role.UpdatedAt,
		&role.DeletedAt,
		&role.Name,
		&role.AccessLevel,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if role.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetRoles(root bool) ([]models.Role, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "roles"
	query, _, err := commons.GetQuery(tableName, models.Role{}, "SELECT", false)
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

	var roles []models.Role
	for rows.Next() {
		var role models.Role
		err = rows.Scan(
			&role.Id,
			&role.CreatedAt,
			&role.UpdatedAt,
			&role.DeletedAt,
			&role.Name,
			&role.AccessLevel,
		)
		if err != nil {
			return nil, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
		}

		roles = append(roles, role)
	}

	if len(roles) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}

	return roles, nil
}

func GetRole(id uint, root bool) (models.Role, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "roles"
	query, _, _ := commons.GetQuery(tableName, models.Role{}, "SELECT", false)
	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	role := models.Role{}
	err := db.QueryRow(query, id).Scan(
		&role.Id,
		&role.CreatedAt,
		&role.UpdatedAt,
		&role.DeletedAt,
		&role.Name,
		&role.AccessLevel,
	)
	if err != nil {
		return role, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	if role.Id == 0 {
		return role, fmt.Errorf("%s not found", tableName)

	}

	return role, nil
}

func DeleteRole(id uint, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "roles"
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

func UpdateRole(updatingRole *models.Role, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "roles"
	currentDate := time.Now()

	if !root {
		updatingRole.UpdatedAt = &currentDate
		updatingRole.CreatedAt = nil // Not allowed to update this field
		updatingRole.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingRole.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingRole, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND delete_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingRole.Id,
		&updatingRole.CreatedAt,
		&updatingRole.UpdatedAt,
		&updatingRole.DeletedAt,
		&updatingRole.Name,
		&updatingRole.AccessLevel,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingRole.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

//

func NewUser(user *models.User) error {
	db := database.Connect()
	defer db.Close()

	user.Username = commons.CleanSpaces(user.Username)
	user.Password = commons.CleanSpaces(user.Password)
	if user.Username == "" || user.Password == "" {
		return errors.New("username and password is required")
	}

	tableName := "users"
	user.Id = 0 // To don't add id field to query
	currentDate := time.Now()
	user.UpdatedAt = &currentDate
	user.CreatedAt = &currentDate

	query, data, err := commons.GetQuery(tableName, *user, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
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
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())

	}
	if user.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)

	}

	return nil
}

func GetUsers(root bool) ([]models.User, error) {

	db := database.Connect()
	defer db.Close()

	tableName := "users"
	query, _, err := commons.GetQuery(tableName, models.User{}, "SELECT", false)
	if err != nil {
		return nil, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	if !root {
		query += " WHERE deleted_at IS NULL"
	}

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("can't execute the query %s ERROR: %s", tableName, err.Error())
	}
	defer rows.Close()

	var users []models.User

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
			return nil, fmt.Errorf("error scanning struct %s ERROR: %s", tableName, err.Error())
		}
		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}

	return users, nil
}

func GetUser(id uint, root bool) (models.User, error) {
	var user models.User

	db := database.Connect()
	defer db.Close()

	tableName := "users"
	query, _, err := commons.GetQuery(tableName, user, "SELECT", false)
	if err != nil {
		return user, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	err = db.QueryRow(query, id).Scan(
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
		&user.BranchID,
	)
	if err != nil {
		return user, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if user.Id == 0 {
		return user, fmt.Errorf("%s not found", tableName)
	}
	return user, err
}

func GetAdmins() ([]models.User, error) {
	db := database.Connect()
	defer db.Close()

	tables := make(map[string]interface{})
	tables["inherit_user_roles"] = models.InheritUserRole{}
	tables["roles"] = models.Role{}
	tables["users"] = models.User{}

	query := commons.GetMixSelect(tables)
	query += " WHERE inherit_user_roles_single.role_id = roles_single.id" +
		" AND inherit_user_roles_single.user_id = users_single.id"

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("can't execute the multiple query ERROR: %s", err.Error())
	}

	var users []models.User

	for rows.Next() {
		var inheritUserRole models.InheritUserRole
		var role models.Role
		var user models.User

		err := rows.Scan(
			&inheritUserRole.Id,
			&inheritUserRole.CreatedAt,
			&inheritUserRole.UpdatedAt,
			&inheritUserRole.DeletedAt,
			&inheritUserRole.RoleID,
			&inheritUserRole.UserID,
			&role.Id,
			&role.CreatedAt,
			&role.UpdatedAt,
			&role.DeletedAt,
			&role.Name,
			&role.AccessLevel,
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
			&user.BranchID,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning struct user ERROR: %s", err.Error())
		}

		role.Id = inheritUserRole.Id
		role.CreatedAt = inheritUserRole.CreatedAt
		role.UpdatedAt = inheritUserRole.UpdatedAt
		role.DeletedAt = inheritUserRole.DeletedAt

		user.InheritUserRoles = append(user.InheritUserRoles, role)
		users = append(users, user)
	}

	return users, nil
}

func DeleteUser(id uint, root bool) error {
	db := database.Connect()
	defer db.Close()
	tableName := "users"

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

	query := fmt.Sprintf("UPDATE %s SET deleted_at = $1 WHERE id = $2", tableName)
	result, err := db.Exec(query, time.Now(), id)
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

func UpdateUser(updatingUser *models.User, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "users"
	currentDate := time.Now()

	if !root {
		updatingUser.UpdatedAt = &currentDate
		updatingUser.CreatedAt = nil // Not allowed to update this field
		updatingUser.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingUser.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingUser, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND delete_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
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
		&updatingUser.BranchID,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingUser.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

//

func NewInheritUserRole(inheritUserRole *models.InheritUserRole) error {
	db := database.Connect()
	defer db.Close()

	currentDate := time.Now()
	inheritUserRole.Id = 0
	inheritUserRole.UpdatedAt = &currentDate
	inheritUserRole.CreatedAt = &currentDate

	tableName := "inherit_user_roles"
	query, data, err := commons.GetQuery(tableName, *inheritUserRole, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&inheritUserRole.Id,
		&inheritUserRole.CreatedAt,
		&inheritUserRole.UpdatedAt,
		&inheritUserRole.DeletedAt,
		&inheritUserRole.RoleID,
		&inheritUserRole.UserID,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if inheritUserRole.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetInheritUserRoles(root bool) ([]models.InheritUserRole, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "inherit_user_roles"
	query, _, err := commons.GetQuery(tableName, models.InheritUserRole{}, "SELECT", false)
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

	var inheritUserRoles []models.InheritUserRole
	for rows.Next() {
		var inheritUserRole models.InheritUserRole
		err = rows.Scan(
			&inheritUserRole.Id,
			&inheritUserRole.CreatedAt,
			&inheritUserRole.UpdatedAt,
			&inheritUserRole.DeletedAt,
			&inheritUserRole.RoleID,
			&inheritUserRole.UserID,
		)
		if err != nil {
			return nil, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
		}

		inheritUserRoles = append(inheritUserRoles, inheritUserRole)
	}

	if len(inheritUserRoles) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}

	return inheritUserRoles, nil
}

func GetInheritUserRole(id uint, root bool) (models.InheritUserRole, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "inherit_user_roles"
	query, _, _ := commons.GetQuery(tableName, models.InheritUserRole{}, "SELECT", false)
	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	inheritUserRole := models.InheritUserRole{}
	err := db.QueryRow(query, id).Scan(
		&inheritUserRole.Id,
		&inheritUserRole.CreatedAt,
		&inheritUserRole.UpdatedAt,
		&inheritUserRole.DeletedAt,
		&inheritUserRole.RoleID,
		&inheritUserRole.UserID,
	)
	if err != nil {
		return inheritUserRole, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	if inheritUserRole.Id == 0 {
		return inheritUserRole, fmt.Errorf("%s not found", tableName)

	}

	return inheritUserRole, nil
}

func DeleteInheritUserRole(id uint, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "inherit_user_roles"
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

func UpdateInheritUserRole(updatingInheritUserRole *models.InheritUserRole, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "inherit_user_roles"
	currentDate := time.Now()

	if !root {
		updatingInheritUserRole.UpdatedAt = &currentDate
		updatingInheritUserRole.CreatedAt = nil // Not allowed to update this field
		updatingInheritUserRole.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingInheritUserRole.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingInheritUserRole, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND delete_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingInheritUserRole.Id,
		&updatingInheritUserRole.CreatedAt,
		&updatingInheritUserRole.UpdatedAt,
		&updatingInheritUserRole.DeletedAt,
		&updatingInheritUserRole.RoleID,
		&updatingInheritUserRole.UserID,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingInheritUserRole.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

// TODO: check this case ---------------------
// Relational, use subID
func GetInheritUserRolesTODO(subId uint, root bool) ([]models.Role, error) {
	var userRoles []models.Role
	db := database.Connect()
	defer db.Close()

	tableName := "inherit_user_roles"
	query, _, err := commons.GetQuery(tableName, models.InheritUserRole{}, "SELECT", false)
	if err != nil {
		return nil, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}
	query += " WHERE user_id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	rows, err := db.Query(query, subId)
	if err != nil {
		return nil, fmt.Errorf("can't execute the query %s ERROR: %s", tableName, err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var inheritUserRole models.InheritUserRole
		err = rows.Scan(
			&inheritUserRole.Id,
			&inheritUserRole.CreatedAt,
			&inheritUserRole.UpdatedAt,
			&inheritUserRole.DeletedAt,
			&inheritUserRole.RoleID,
			&inheritUserRole.UserID)
		if err != nil {
			return nil, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
		}

		subTableName := "roles"
		role, err := GetRole(inheritUserRole.RoleID, root)
		if err != nil {
			return nil, fmt.Errorf("%s not found in a relational table ERROR: %s", subTableName, err.Error())
		}
		userRoles = append(userRoles, role)
	}

	if len(userRoles) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}

	return userRoles, nil
}

//
// --------------------------------------------

//

func NewUserPhone(userPhone *models.UserPhone) error {
	db := database.Connect()
	defer db.Close()

	userPhone.Phone = commons.CleanSpaces(userPhone.Phone)
	if !commons.IsANumber(userPhone.Phone) {
		return errors.New("phone is not a valid input, need just numbers")
	}

	currentDate := time.Now()
	userPhone.Id = 0
	userPhone.UpdatedAt = &currentDate
	userPhone.CreatedAt = &currentDate

	tableName := "user_phones"
	query, data, err := commons.GetQuery(tableName, *userPhone, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&userPhone.Id,
		&userPhone.CreatedAt,
		&userPhone.UpdatedAt,
		&userPhone.DeletedAt,
		&userPhone.Phone,
		&userPhone.Main,
		&userPhone.UserID,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if userPhone.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetUserPhones(root bool) ([]models.UserPhone, error) {
	var userPhones []models.UserPhone
	db := database.Connect()
	defer db.Close()

	tableName := "user_phones"
	query, _, err := commons.GetQuery(tableName, models.UserPhone{}, "SELECT", false)
	if err != nil {
		return nil, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	if !root {
		query += " WHERE deleted_at IS NULL"
	}

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("can't execute the query %s ERROR: %s", tableName, err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var userPhone models.UserPhone
		err = rows.Scan(
			&userPhone.Id,
			&userPhone.CreatedAt,
			&userPhone.UpdatedAt,
			&userPhone.DeletedAt,
			&userPhone.Phone,
			&userPhone.Main,
			&userPhone.UserID,
		)
		if err != nil {
			return nil, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
		}

		userPhones = append(userPhones, userPhone)
	}
	if len(userPhones) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}

	return userPhones, nil
}

func GetUserPhone(id uint, root bool) (models.UserPhone, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "user_phones"
	query, _, _ := commons.GetQuery(tableName, models.UserPhone{}, "SELECT", false)
	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	userPhone := models.UserPhone{}
	err := db.QueryRow(query, id).Scan(
		&userPhone.Id,
		&userPhone.CreatedAt,
		&userPhone.UpdatedAt,
		&userPhone.DeletedAt,
		&userPhone.Phone,
		&userPhone.Main,
		&userPhone.UserID,
	)
	if err != nil {
		return userPhone, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	if userPhone.Id == 0 {
		return userPhone, fmt.Errorf("%s not found", tableName)

	}

	return userPhone, nil
}

func DeleteUserPhone(id uint, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "user_phones"
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

func UpdateUserPhone(updatingUserPhone *models.UserPhone, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "user_phones"
	currentDate := time.Now()

	if !root {
		updatingUserPhone.UpdatedAt = &currentDate
		updatingUserPhone.CreatedAt = nil // Not allowed to update this field
		updatingUserPhone.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingUserPhone.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingUserPhone, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND delete_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingUserPhone.Id,
		&updatingUserPhone.CreatedAt,
		&updatingUserPhone.UpdatedAt,
		&updatingUserPhone.DeletedAt,
		&updatingUserPhone.Phone,
		&updatingUserPhone.Main,
		&updatingUserPhone.UserID,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingUserPhone.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

//

func NewUserMail(userMail *models.UserMail) error {
	db := database.Connect()
	defer db.Close()

	userMail.Mail = commons.CleanSpaces(userMail.Mail)
	if !commons.IsAMail(userMail.Mail) {
		return errors.New("mail is not a valid input")
	}

	currentDate := time.Now()
	userMail.Id = 0
	userMail.UpdatedAt = &currentDate
	userMail.CreatedAt = &currentDate

	tableName := "user_mails"
	query, data, err := commons.GetQuery(tableName, *userMail, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&userMail.Id,
		&userMail.CreatedAt,
		&userMail.UpdatedAt,
		&userMail.DeletedAt,
		&userMail.Mail,
		&userMail.Main,
		&userMail.UserID,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if userMail.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetUserMails(root bool) ([]models.UserMail, error) {
	var userMails []models.UserMail
	db := database.Connect()
	defer db.Close()

	tableName := "user_mails"
	query, _, err := commons.GetQuery(tableName, models.UserMail{}, "SELECT", false)
	if err != nil {
		return nil, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	if !root {
		query += " WHERE deleted_at IS NULL"
	}

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("can't execute the query %s ERROR: %s", tableName, err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var userMail models.UserMail
		err = rows.Scan(
			&userMail.Id,
			&userMail.CreatedAt,
			&userMail.UpdatedAt,
			&userMail.DeletedAt,
			&userMail.Mail,
			&userMail.Main,
			&userMail.UserID)
		if err != nil {
			return nil, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
		}

		userMails = append(userMails, userMail)
	}

	if len(userMails) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}

	return userMails, nil
}

func GetUserMail(id uint, root bool) (models.UserMail, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "user_mails"
	query, _, _ := commons.GetQuery(tableName, models.UserMail{}, "SELECT", false)
	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	userMail := models.UserMail{}
	err := db.QueryRow(query, id).Scan(
		&userMail.Id,
		&userMail.CreatedAt,
		&userMail.UpdatedAt,
		&userMail.DeletedAt,
		&userMail.Mail,
		&userMail.Main,
		&userMail.UserID,
	)
	if err != nil {
		return userMail, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	if userMail.Id == 0 {
		return userMail, fmt.Errorf("%s not found", tableName)

	}

	return userMail, nil
}

func DeleteUserMail(id uint, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "user_mails"
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

func UpdateUserMail(updatingUserMail *models.UserMail, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "user_mails"
	currentDate := time.Now()

	if !root {
		updatingUserMail.UpdatedAt = &currentDate
		updatingUserMail.CreatedAt = nil // Not allowed to update this field
		updatingUserMail.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingUserMail.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingUserMail, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND delete_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingUserMail.Id,
		&updatingUserMail.CreatedAt,
		&updatingUserMail.UpdatedAt,
		&updatingUserMail.DeletedAt,
		&updatingUserMail.Mail,
		&updatingUserMail.Main,
		&updatingUserMail.UserID,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingUserMail.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

//

func NewUserReport(userReport *models.UserReport) error {
	db := database.Connect()
	defer db.Close()

	currentDate := time.Now()
	userReport.Id = 0
	userReport.UpdatedAt = &currentDate
	userReport.CreatedAt = &currentDate

	tableName := "user_reports"
	query, data, err := commons.GetQuery(tableName, *userReport, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&userReport.Id,
		&userReport.CreatedAt,
		&userReport.UpdatedAt,
		&userReport.DeletedAt,
		&userReport.Reason,
		&userReport.UserID,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if userReport.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetUserReports(root bool) ([]models.UserReport, error) {
	var userReports []models.UserReport
	db := database.Connect()
	defer db.Close()

	tableName := "user_reports"
	query, _, err := commons.GetQuery(tableName, models.UserReport{}, "SELECT", false)
	if err != nil {
		return nil, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	if !root {
		query += " WHERE deleted_at IS NULL"
	}

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("can't execute the query %s ERROR: %s", tableName, err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var userReport models.UserReport
		err = rows.Scan(
			&userReport.Id,
			&userReport.CreatedAt,
			&userReport.UpdatedAt,
			&userReport.DeletedAt,
			&userReport.Reason,
			&userReport.UserID)
		if err != nil {
			return nil, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
		}

		userReports = append(userReports, userReport)
	}

	if len(userReports) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}

	return userReports, nil
}

func GetUserReport(id uint, root bool) (models.UserReport, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "user_reports"
	query, _, _ := commons.GetQuery(tableName, models.UserReport{}, "SELECT", false)
	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	userReport := models.UserReport{}
	err := db.QueryRow(query, id).Scan(
		&userReport.Id,
		&userReport.CreatedAt,
		&userReport.UpdatedAt,
		&userReport.DeletedAt,
		&userReport.Reason,
		&userReport.UserID,
	)
	if err != nil {
		return userReport, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	if userReport.Id == 0 {
		return userReport, fmt.Errorf("%s not found", tableName)

	}

	return userReport, nil
}

func DeleteUserReport(id uint, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "user_reports"
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

func UpdateUserReport(updatingUserReport *models.UserReport, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "user_reports"
	currentDate := time.Now()

	if !root {
		updatingUserReport.UpdatedAt = &currentDate
		updatingUserReport.CreatedAt = nil // Not allowed to update this field
		updatingUserReport.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingUserReport.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingUserReport, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND delete_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingUserReport.Id,
		&updatingUserReport.CreatedAt,
		&updatingUserReport.UpdatedAt,
		&updatingUserReport.DeletedAt,
		&updatingUserReport.Reason,
		&updatingUserReport.UserID,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingUserReport.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

//

func NewMonetaryBound(monetaryBounds *models.MonetaryBound) error {
	db := database.Connect()
	defer db.Close()

	currentDate := time.Now()
	monetaryBounds.Id = 0
	monetaryBounds.UpdatedAt = &currentDate
	monetaryBounds.CreatedAt = &currentDate

	tableName := "monetary_bounds"
	query, data, err := commons.GetQuery(tableName, *monetaryBounds, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&monetaryBounds.Id,
		&monetaryBounds.CreatedAt,
		&monetaryBounds.UpdatedAt,
		&monetaryBounds.DeletedAt,
		&monetaryBounds.Reason,
		&monetaryBounds.Bound,
		&monetaryBounds.UserID,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if monetaryBounds.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetMonetaryBounds(root bool) ([]models.MonetaryBound, error) {
	var monetaryBounds []models.MonetaryBound
	db := database.Connect()
	defer db.Close()

	tableName := "monetary_bounds"
	query, _, err := commons.GetQuery(tableName, models.MonetaryBound{}, "SELECT", false)
	if err != nil {
		return nil, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	if !root {
		query += " WHERE deleted_at IS NULL"
	}

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("can't execute the query %s ERROR: %s", tableName, err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var monetaryBound models.MonetaryBound
		err = rows.Scan(
			&monetaryBound.Id,
			&monetaryBound.CreatedAt,
			&monetaryBound.UpdatedAt,
			&monetaryBound.DeletedAt,
			&monetaryBound.Reason,
			&monetaryBound.Bound,
			&monetaryBound.UserID)
		if err != nil {
			return nil, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
		}

		monetaryBounds = append(monetaryBounds, monetaryBound)
	}

	if len(monetaryBounds) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}

	return monetaryBounds, nil
}

func GetMonetaryBound(id uint, root bool) (models.MonetaryBound, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "monetary_bounds"
	query, _, _ := commons.GetQuery(tableName, models.MonetaryBound{}, "SELECT", false)
	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	monetaryBound := models.MonetaryBound{}
	err := db.QueryRow(query, id).Scan(
		&monetaryBound.Id,
		&monetaryBound.CreatedAt,
		&monetaryBound.UpdatedAt,
		&monetaryBound.DeletedAt,
		&monetaryBound.Reason,
		&monetaryBound.Bound,
		&monetaryBound.UserID,
	)
	if err != nil {
		return monetaryBound, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	if monetaryBound.Id == 0 {
		return monetaryBound, fmt.Errorf("%s not found", tableName)

	}

	return monetaryBound, nil
}

func DeleteMonetaryBound(id uint, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "monetary_bounds"
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

func UpdateMonetaryBound(updatingMonetaryBound *models.MonetaryBound, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "monetary_bounds"
	currentDate := time.Now()

	if !root {
		updatingMonetaryBound.UpdatedAt = &currentDate
		updatingMonetaryBound.CreatedAt = nil // Not allowed to update this field
		updatingMonetaryBound.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingMonetaryBound.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingMonetaryBound, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND delete_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingMonetaryBound.Id,
		&updatingMonetaryBound.CreatedAt,
		&updatingMonetaryBound.UpdatedAt,
		&updatingMonetaryBound.DeletedAt,
		&updatingMonetaryBound.Reason,
		&updatingMonetaryBound.Bound,
		&updatingMonetaryBound.UserID,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingMonetaryBound.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

//

func NewMonetaryDiscount(monetaryDiscount *models.MonetaryDiscount) error {
	db := database.Connect()
	defer db.Close()

	currentDate := time.Now()
	monetaryDiscount.Id = 0
	monetaryDiscount.UpdatedAt = &currentDate
	monetaryDiscount.CreatedAt = &currentDate

	tableName := "monetary_discounts"
	query, data, err := commons.GetQuery(tableName, *monetaryDiscount, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&monetaryDiscount.Id,
		&monetaryDiscount.CreatedAt,
		&monetaryDiscount.UpdatedAt,
		&monetaryDiscount.DeletedAt,
		&monetaryDiscount.Reason,
		&monetaryDiscount.Discount,
		&monetaryDiscount.UserID,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if monetaryDiscount.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetMonetaryDiscounts(root bool) ([]models.MonetaryDiscount, error) {
	var monetaryDiscounts []models.MonetaryDiscount
	db := database.Connect()
	defer db.Close()

	tableName := "monetary_discounts"
	query, _, err := commons.GetQuery(tableName, models.MonetaryDiscount{}, "SELECT", false)
	if err != nil {
		return nil, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	if !root {
		query += " WHERE deleted_at IS NULL"
	}

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("can't execute the query %s ERROR: %s", tableName, err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var monetaryDiscount models.MonetaryDiscount
		err = rows.Scan(
			&monetaryDiscount.Id,
			&monetaryDiscount.CreatedAt,
			&monetaryDiscount.UpdatedAt,
			&monetaryDiscount.DeletedAt,
			&monetaryDiscount.Reason,
			&monetaryDiscount.Discount,
			&monetaryDiscount.UserID,
		)
		if err != nil {
			return nil, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
		}

		monetaryDiscounts = append(monetaryDiscounts, monetaryDiscount)
	}

	if len(monetaryDiscounts) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}

	return monetaryDiscounts, nil
}

func GetMonetaryDiscount(id uint, root bool) (models.MonetaryDiscount, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "monetary_discounts"
	query, _, _ := commons.GetQuery(tableName, models.MonetaryDiscount{}, "SELECT", false)
	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	monetaryDiscount := models.MonetaryDiscount{}
	err := db.QueryRow(query, id).Scan(
		&monetaryDiscount.Id,
		&monetaryDiscount.CreatedAt,
		&monetaryDiscount.UpdatedAt,
		&monetaryDiscount.DeletedAt,
		&monetaryDiscount.Reason,
		&monetaryDiscount.Discount,
		&monetaryDiscount.UserID,
	)
	if err != nil {
		return monetaryDiscount, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	if monetaryDiscount.Id == 0 {
		return monetaryDiscount, fmt.Errorf("%s not found", tableName)

	}

	return monetaryDiscount, nil
}

func DeleteMonetaryDiscount(id uint, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "monetary_discounts"
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

func UpdateMonetaryDiscount(updatingMonetaryDiscount *models.MonetaryDiscount, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "monetary_discounts"
	currentDate := time.Now()

	if !root {
		updatingMonetaryDiscount.UpdatedAt = &currentDate
		updatingMonetaryDiscount.CreatedAt = nil // Not allowed to update this field
		updatingMonetaryDiscount.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingMonetaryDiscount.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingMonetaryDiscount, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND delete_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingMonetaryDiscount.Id,
		&updatingMonetaryDiscount.CreatedAt,
		&updatingMonetaryDiscount.UpdatedAt,
		&updatingMonetaryDiscount.DeletedAt,
		&updatingMonetaryDiscount.Reason,
		&updatingMonetaryDiscount.Discount,
		&updatingMonetaryDiscount.UserID,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingMonetaryDiscount.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

//

func NewBranchUserRole(branchUserRole *models.BranchUserRole) error {
	db := database.Connect()
	defer db.Close()

	currentDate := time.Now()
	branchUserRole.Id = 0
	branchUserRole.UpdatedAt = &currentDate
	branchUserRole.CreatedAt = &currentDate

	tableName := "branch_user_roles"
	query, data, err := commons.GetQuery(tableName, *branchUserRole, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&branchUserRole.Id,
		&branchUserRole.CreatedAt,
		&branchUserRole.UpdatedAt,
		&branchUserRole.DeletedAt,
		&branchUserRole.BranchID,
		&branchUserRole.UserID,
		&branchUserRole.RoleID,
	)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if branchUserRole.Id == 0 {
		return fmt.Errorf("can't create %s", tableName)
	}

	return nil
}

func GetBranchUserRoles(root bool) ([]models.BranchUserRole, error) {
	var branchUserRoles []models.BranchUserRole
	db := database.Connect()
	defer db.Close()

	tableName := "branch_user_roles"
	query, _, err := commons.GetQuery(tableName, models.BranchUserRole{}, "SELECT", false)
	if err != nil {
		return nil, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	if !root {
		query += " WHERE deleted_at IS NULL"
	}

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("can't execute the query %s ERROR: %s", tableName, err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var branchUserRole models.BranchUserRole
		err = rows.Scan(
			&branchUserRole.Id,
			&branchUserRole.CreatedAt,
			&branchUserRole.UpdatedAt,
			&branchUserRole.DeletedAt,
			&branchUserRole.BranchID,
			&branchUserRole.UserID,
			&branchUserRole.RoleID,
		)
		if err != nil {
			return nil, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
		}

		branchUserRoles = append(branchUserRoles, branchUserRole)
	}

	if len(branchUserRoles) == 0 {
		return nil, fmt.Errorf("%s not found", tableName)
	}

	return branchUserRoles, nil
}

func GetBranchUserRole(id uint, root bool) (models.BranchUserRole, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "branch_user_roles"
	query, _, _ := commons.GetQuery(tableName, models.BranchUserRole{}, "SELECT", false)
	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	branchUserRole := models.BranchUserRole{}
	err := db.QueryRow(query, id).Scan(
		&branchUserRole.Id,
		&branchUserRole.CreatedAt,
		&branchUserRole.UpdatedAt,
		&branchUserRole.DeletedAt,
		&branchUserRole.BranchID,
		&branchUserRole.UserID,
		&branchUserRole.RoleID,
	)
	if err != nil {
		return branchUserRole, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	if branchUserRole.Id == 0 {
		return branchUserRole, fmt.Errorf("%s not found", tableName)

	}

	return branchUserRole, nil
}

func DeleteBranchUserRole(id uint, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "branch_user_roles"
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

func UpdateBranchUserRole(updatingBranchUserRole *models.BranchUserRole, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "branch_user_roles"
	currentDate := time.Now()

	if !root {
		updatingBranchUserRole.UpdatedAt = &currentDate
		updatingBranchUserRole.CreatedAt = nil // Not allowed to update this field
		updatingBranchUserRole.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, updatingBranchUserRole.Id) && !root {
		return errors.New("regist is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingBranchUserRole, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND delete_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&updatingBranchUserRole.Id,
		&updatingBranchUserRole.CreatedAt,
		&updatingBranchUserRole.UpdatedAt,
		&updatingBranchUserRole.DeletedAt,
		&updatingBranchUserRole.BranchID,
		&updatingBranchUserRole.UserID,
		&updatingBranchUserRole.RoleID,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if updatingBranchUserRole.Id == 0 {
		return fmt.Errorf("%s not found", tableName)
	}

	return nil
}

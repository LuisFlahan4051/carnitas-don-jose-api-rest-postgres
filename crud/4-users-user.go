package crud

import (
	"errors"
	"fmt"

	"time"

	commons "github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/commonFunctions"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/database"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/models"
)

func NewRole(role *models.Role) error {
	db := database.Connect()
	defer db.Close()

	if role.AccessLevel <= 1 {
		return errors.New("access level must be greater than 1, new role can't be root")
	}

	tableName := "roles"
	currentDate := time.Now()
	role.UpdatedAt = &currentDate
	role.CreatedAt = &currentDate

	query, data, err := commons.GetQuery(tableName, *role, 0, true)
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
		return errors.New("can't create role")
	}

	return nil
}

func GetRole(id uint, root bool) (models.Role, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "roles"
	query, _, _ := commons.GetQuery(tableName, models.Role{}, 2, false)
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
		return role, errors.New("role not found")
	}

	return role, nil
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

	query, data, err := commons.GetQuery(tableName, *user, 0, true)
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
		return errors.New("can't create user")
	}

	return nil
}

func GetUsers(root bool) ([]models.User, error) {

	db := database.Connect()
	defer db.Close()

	tableName := "users"
	query, _, err := commons.GetQuery(tableName, models.User{}, 2, false)
	if err != nil {
		return nil, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	if !root {
		query += " WHERE deleted_at IS NULL"
	}

	var users []models.User

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("can't execute the query %s ERROR: %s", tableName, err.Error())
	}
	defer rows.Close()

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
		return nil, errors.New("users not found")
	}

	return users, nil
}

func GetUser(id uint, root bool) (models.User, error) {
	var user models.User

	db := database.Connect()
	defer db.Close()

	tableName := "users"
	query, _, err := commons.GetQuery(tableName, models.User{}, 2, false)
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
		return user, errors.New("user not found")
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
		return err
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
		return fmt.Errorf("fail to delete ERROR: %s", err.Error())
	}

	if rowsAffected == 0 {
		return errors.New("user doesn't exist")
	}

	return nil
}

func UpdateUser(updatingUser *models.User, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "users"
	currentDate := time.Now()
	updatingUser.UpdatedAt = &currentDate
	updatingUser.CreatedAt = nil
	updatingUser.DeletedAt = nil

	if root {
		query, data, err := commons.GetQuery(tableName, *updatingUser, 1, true)
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

		return nil
	}

	if commons.IsDeleted(tableName, updatingUser.Id) {
		return errors.New("user is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *updatingUser, 1, true)
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

	return nil
}

//

func NewInheritUserRole(inheritUserRole *models.InheritUserRole) error {
	db := database.Connect()
	defer db.Close()

	currentDate := time.Now()
	inheritUserRole.UpdatedAt = &currentDate
	inheritUserRole.CreatedAt = &currentDate

	tableName := "inherit_user_roles"
	query, data, err := commons.GetQuery(tableName, *inheritUserRole, 0, true)
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
		return errors.New("can't create inherit user role")
	}

	return nil
}

func GetInheritUserRoles(id uint, root bool) ([]models.Role, error) {
	var userRoles []models.Role
	db := database.Connect()
	defer db.Close()

	tableName := "inherit_user_roles"
	query, _, err := commons.GetQuery(tableName, models.InheritUserRole{}, 2, false)
	if err != nil {
		return nil, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}
	query += " WHERE user_id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	rows, err := db.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
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

		role, err := GetRole(inheritUserRole.RoleID, root)
		if err != nil {
			return nil, fmt.Errorf("critical role no found to relational table ERROR: %s", err.Error())
		}
		userRoles = append(userRoles, role)
	}

	if len(userRoles) == 0 {
		return nil, errors.New("user doesn't have roles")
	}

	return userRoles, nil
}

//

func NewUserPhone(userPhone *models.UserPhone) error {
	db := database.Connect()
	defer db.Close()

	userPhone.Phone = commons.CleanSpaces(userPhone.Phone)
	if !commons.IsANumber(userPhone.Phone) {
		return errors.New("phone is not a valid input, need just numbers")
	}

	currentDate := time.Now()
	userPhone.UpdatedAt = &currentDate
	userPhone.CreatedAt = &currentDate

	tableName := "user_phones"
	query, data, err := commons.GetQuery(tableName, *userPhone, 0, true)
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
		return errors.New("can't create user phone")
	}

	return nil
}

func GetUserPhones(id uint, root bool) ([]models.UserPhone, error) {
	var userPhones []models.UserPhone
	db := database.Connect()
	defer db.Close()

	tableName := "user_phones"
	query, _, err := commons.GetQuery(tableName, models.UserPhone{}, 2, false)
	if err != nil {
		return nil, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}
	query += " WHERE user_id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	rows, err := db.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
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
		return nil, errors.New("user doesn't have phones")
	}

	return userPhones, nil
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
	userMail.UpdatedAt = &currentDate
	userMail.CreatedAt = &currentDate

	tableName := "user_mails"
	query, data, err := commons.GetQuery(tableName, *userMail, 0, true)
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
		return errors.New("can't create user mail")
	}

	return nil
}

func GetUserMails(id uint, root bool) ([]models.UserMail, error) {
	var userMails []models.UserMail
	db := database.Connect()
	defer db.Close()

	tableName := "user_mails"
	query, _, err := commons.GetQuery(tableName, models.UserMail{}, 2, false)
	if err != nil {
		return nil, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}
	query += " WHERE user_id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	rows, err := db.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
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
		return nil, errors.New("user doesn't have mails")
	}

	return userMails, nil
}

//

func NewUserReport(userReport *models.UserReport) error {
	db := database.Connect()
	defer db.Close()

	currentDate := time.Now()
	userReport.UpdatedAt = &currentDate
	userReport.CreatedAt = &currentDate

	tableName := "user_reports"
	query, data, err := commons.GetQuery(tableName, *userReport, 0, true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&userReport.Id,
		&userReport.CreatedAt,
		&userReport.UpdatedAt,
		&userReport.DeletedAt,
		&userReport.Reason,
		&userReport.UserID)
	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	if userReport.Id == 0 {
		return errors.New("can't create user report")
	}

	return nil
}

func GetUserReports(id uint, root bool) ([]models.UserReport, error) {
	var userReports []models.UserReport
	db := database.Connect()
	defer db.Close()

	tableName := "user_reports"
	query, _, err := commons.GetQuery(tableName, models.UserReport{}, 2, false)
	if err != nil {
		return nil, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}
	query += " WHERE user_id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	rows, err := db.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
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
		return nil, errors.New("user doesn't have reports")
	}

	return userReports, nil
}

//

func NewMonetaryBound(monetaryBounds *models.MonetaryBound) error {
	db := database.Connect()
	defer db.Close()

	currentDate := time.Now()
	monetaryBounds.UpdatedAt = &currentDate
	monetaryBounds.CreatedAt = &currentDate

	tableName := "monetary_bounds"
	query, data, err := commons.GetQuery(tableName, *monetaryBounds, 0, true)
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
		return errors.New("can't create monetary bound")
	}

	return nil
}

func GetMonetaryBounds(id uint, root bool) ([]models.MonetaryBound, error) {
	var monetaryBounds []models.MonetaryBound
	db := database.Connect()
	defer db.Close()

	tableName := "monetary_bounds"
	query, _, err := commons.GetQuery(tableName, models.MonetaryBound{}, 2, false)
	if err != nil {
		return nil, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}
	query += " WHERE user_id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	rows, err := db.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
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
		return nil, errors.New("user doesn't have monetary bounds")
	}

	return monetaryBounds, nil
}

//

func NewMonetaryDiscount(monetaryDiscount *models.MonetaryDiscount) error {
	db := database.Connect()
	defer db.Close()

	currentDate := time.Now()
	monetaryDiscount.UpdatedAt = &currentDate
	monetaryDiscount.CreatedAt = &currentDate

	tableName := "monetary_discounts"
	query, data, err := commons.GetQuery(tableName, *monetaryDiscount, 0, true)
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
		return errors.New("can't create monetary discount")
	}

	return nil
}

func GetMonetaryDiscounts(id uint, root bool) ([]models.MonetaryDiscount, error) {
	var monetaryDiscounts []models.MonetaryDiscount
	db := database.Connect()
	defer db.Close()

	tableName := "monetary_discounts"
	query, _, err := commons.GetQuery(tableName, models.MonetaryDiscount{}, 2, false)
	if err != nil {
		return nil, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}
	query += " WHERE user_id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	rows, err := db.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
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
		return nil, errors.New("user doesn't have monetary discounts")
	}

	return monetaryDiscounts, nil
}

//

func NewBranchUserRole(branchUserRole *models.BranchUserRole) error {
	db := database.Connect()
	defer db.Close()

	currentDate := time.Now()
	branchUserRole.UpdatedAt = &currentDate
	branchUserRole.CreatedAt = &currentDate

	tableName := "branch_user_roles"
	query, data, err := commons.GetQuery(tableName, *branchUserRole, 0, true)
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
		return errors.New("can't create branch user role")
	}

	return nil
}

func GetBranchUserRoles(id uint, root bool) ([]models.BranchUserRole, error) {
	var branchUserRoles []models.BranchUserRole
	db := database.Connect()
	defer db.Close()

	tableName := "branch_user_roles"
	query, _, err := commons.GetQuery(tableName, models.BranchUserRole{}, 2, false)
	if err != nil {
		return nil, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}
	query += " WHERE user_id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	rows, err := db.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
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
		return nil, errors.New("user doesn't have branch user roles")
	}

	return branchUserRoles, nil
}

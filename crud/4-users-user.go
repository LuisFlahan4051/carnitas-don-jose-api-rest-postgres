package crud

import (
	"errors"
	"log"

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

func GetRole(id uint, root bool) (models.Role, error) {
	db := database.Connect()
	defer db.Close()

	query, _, _ := commons.GetQuery("roles", models.Role{}, 2, false)
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
		&role.AccessLevel)

	if err != nil {
		return role, errors.New("role doesn't exist")
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
			return users, err
		}
		users = append(users, user)
	}

	if len(users) == 0 {
		return users, errors.New("users not found")
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
	log.Println(query)

	query += " WHERE inherit_user_roles_single.role_id = roles_single.id" +
		" AND inherit_user_roles_single.user_id = users_single.id"
	rows, err := db.Query(query)
	if err != nil {
		return []models.User{}, errors.New("admins not found")
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
			return []models.User{}, err
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

	if root {
		query := "DELETE FROM users WHERE id = $1"
		_, err := db.Exec(query, id)
		return err
	}

	if commons.IsDeleted("users", id) {
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

	if commons.IsDeleted("users", updatingUser.Id) {
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
		&inheritUserRole.RoleID,
		&inheritUserRole.UserID)
	if err != nil {
		return err
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

	query, _, _ := commons.GetQuery("inherit_user_roles", models.InheritUserRole{}, 2, false)
	query += " WHERE user_id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
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
			return nil, errors.New("can't get inherit user role")
		}

		role, err := GetRole(inheritUserRole.RoleID, root)

		if err != nil {
			return nil, err
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

	query, data, err := commons.GetQuery("user_phones", *userPhone, 0, true)
	if err != nil {
		return err
	}

	err = db.QueryRow(query, data...).Scan(
		&userPhone.Id,
		&userPhone.CreatedAt,
		&userPhone.UpdatedAt,
		&userPhone.DeletedAt,
		&userPhone.Phone,
		&userPhone.Main,
		&userPhone.UserID)
	if err != nil {
		return err
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

	query, _, _ := commons.GetQuery("user_phones", models.UserPhone{}, 2, false)
	query += " WHERE user_id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
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
			&userPhone.UserID)
		if err != nil {
			return nil, errors.New("can't get user phone")
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

	query, data, err := commons.GetQuery("user_mails", *userMail, 0, true)
	if err != nil {
		return err
	}

	err = db.QueryRow(query, data...).Scan(
		&userMail.Id,
		&userMail.CreatedAt,
		&userMail.UpdatedAt,
		&userMail.DeletedAt,
		&userMail.Mail,
		&userMail.Main,
		&userMail.UserID)
	if err != nil {
		return err
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

	query, _, _ := commons.GetQuery("user_mails", models.UserMail{}, 2, false)
	query += " WHERE user_id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
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
			return nil, errors.New("can't get user mail")
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

	query, data, err := commons.GetQuery("user_reports", *userReport, 0, true)
	if err != nil {
		return err
	}

	err = db.QueryRow(query, data...).Scan(
		&userReport.Id,
		&userReport.CreatedAt,
		&userReport.UpdatedAt,
		&userReport.DeletedAt,
		&userReport.Reason,
		&userReport.UserID)
	if err != nil {
		return err
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

	query, _, _ := commons.GetQuery("user_reports", models.UserReport{}, 2, false)
	query += " WHERE user_id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
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
			return nil, errors.New("can't get user report")
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

	query, data, err := commons.GetQuery("monetary_bounds", *monetaryBounds, 0, true)
	if err != nil {
		return err
	}

	err = db.QueryRow(query, data...).Scan(
		&monetaryBounds.Id,
		&monetaryBounds.CreatedAt,
		&monetaryBounds.UpdatedAt,
		&monetaryBounds.DeletedAt,
		&monetaryBounds.Reason,
		&monetaryBounds.Bound,
		&monetaryBounds.UserID)
	if err != nil {
		return err
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

	query, _, _ := commons.GetQuery("monetary_bounds", models.MonetaryBound{}, 2, false)
	query += " WHERE user_id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
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
			return nil, errors.New("can't get monetary bound")
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

	query, data, err := commons.GetQuery("monetary_discounts", *monetaryDiscount, 0, true)
	if err != nil {
		return err
	}

	err = db.QueryRow(query, data...).Scan(
		&monetaryDiscount.Id,
		&monetaryDiscount.CreatedAt,
		&monetaryDiscount.UpdatedAt,
		&monetaryDiscount.DeletedAt,
		&monetaryDiscount.Reason,
		&monetaryDiscount.Discount,
		&monetaryDiscount.UserID)
	if err != nil {
		return err
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

	query, _, _ := commons.GetQuery("monetary_discounts", models.MonetaryDiscount{}, 2, false)
	query += " WHERE user_id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
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
			&monetaryDiscount.UserID)
		if err != nil {
			return nil, errors.New("can't get monetary discount")
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

	query, data, err := commons.GetQuery("branch_user_roles", *branchUserRole, 0, true)
	if err != nil {
		return err
	}

	err = db.QueryRow(query, data...).Scan(
		&branchUserRole.Id,
		&branchUserRole.CreatedAt,
		&branchUserRole.UpdatedAt,
		&branchUserRole.DeletedAt,
		&branchUserRole.BranchID,
		&branchUserRole.UserID,
		&branchUserRole.RoleID)
	if err != nil {
		return err
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

	query, _, _ := commons.GetQuery("branch_user_roles", models.BranchUserRole{}, 2, false)
	query += " WHERE user_id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
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
			&branchUserRole.RoleID)
		if err != nil {
			return nil, errors.New("can't get branch user role")
		}

		branchUserRoles = append(branchUserRoles, branchUserRole)
	}

	if len(branchUserRoles) == 0 {
		return nil, errors.New("user doesn't have branch user roles")
	}

	return branchUserRoles, nil
}

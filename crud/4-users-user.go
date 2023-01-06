package crud

import (
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"time"

	commons "github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/commonFunctions"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/database"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/models"
)

func NewUser(user *models.User) error {
	db := database.Connect()
	defer db.Close()

	//TODO:
	date := time.Now()
	name := "luis"
	otherUser := models.User{
		ID: models.ID{
			Id:        24,
			CreatedAt: &date,
		},
		Name:     &name,
		Username: "Prueba final",
		Password: "123456",
	}

	query := "INSERT INTO users("
	fieldsSlice, fieldsValuesMap := commons.GetStructFieldsNotNull(otherUser)
	fields := strings.Join(fieldsSlice, ", ")
	query += fields + ") VALUES ("
	for i := 0; i < len(fieldsSlice)-1; i++ {
		query += "$" + strconv.Itoa(i+1) + ", "
	}
	query += "$" + strconv.Itoa(len(fieldsSlice)) + ")"

	allFieldsSlice, _ := commons.GetStructFieldsNotSlices(otherUser)
	allFields := strings.Join(allFieldsSlice, ", ")
	query += " RETURNING " + allFields

	var data []interface{}

	for _, field := range fieldsSlice {
		valueString := fmt.Sprintf("%v", fieldsValuesMap[field])
		regularExpresion := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
		if regularExpresion.MatchString(valueString) {
			date, _ := time.Parse("2006-01-02", valueString)
			data = append(data, date)
		} else {
			data = append(data, valueString)
		}
	}

	err := db.QueryRow(query, data...).Scan(
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

	var query string

	query = "SELECT "
	fieldsSlice, _ := commons.GetStructFieldsNotSlices(models.User{})
	fields := strings.Join(fieldsSlice, ", ")
	query += fields + " FROM users WHERE deleted_at IS NOT NULL"

	if root {
		query = "SELECT " + fields + " FROM users"
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

	query := "SELECT id," +
		" created_at," +
		" updated_at," +
		" deleted_at," +
		" name," +
		" lastname," +
		" username," +
		" password," +
		" photo," +
		" verified," +
		" warning," +
		" darktheme," +
		" active_contract," +
		" address," +
		" born," +
		" degree_study," +
		" relation_ship," +
		" curp," +
		" rfc," +
		" citizen_id," +
		" credential_id," +
		" origin_state," +
		" score," +
		" qualities," +
		" defects," +
		" origin_branch_id," +
		" branch_id FROM users WHERE id = $1 AND deleted_at IS NULL"

	if root {
		query = "SELECT id," +
			" created_at," +
			" updated_at," +
			" deleted_at," +
			" name," +
			" lastname," +
			" username," +
			" password," +
			" photo," +
			" verified," +
			" warning," +
			" darktheme," +
			" active_contract," +
			" address," +
			" born," +
			" degree_study," +
			" relation_ship," +
			" curp," +
			" rfc," +
			" citizen_id," +
			" credential_id," +
			" origin_state," +
			" score," +
			" qualities," +
			" defects," +
			" origin_branch_id," +
			" branch_id FROM users WHERE id = $1"
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
	result, _ := db.Exec(query, time.Now(), id)
	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return errors.New("user doesn't exist")
	}

	return nil
}

// TODO: fix this
func UpdateUser(newUser *models.User, root bool) error {
	db := database.Connect()
	defer db.Close()

	updatingUser, err := GetUser(newUser.Id, root)
	if err != nil {
		return errors.New("user doesn't exist")
	}

	//compare

	if root {
		query := "UPDATE users SET" +
			" created_at = $1," +
			" updated_at = $2," +
			" deleted_at = $3," +
			" name = $4," +
			" lastname = $5," +
			" username = $26," +
			" password = $6," +
			" photo = $7," +
			" verified = $8," +
			" warning = $27," +
			" darktheme = $9," +
			" active_contract = $10," +
			" address = $11," +
			" born = $12," +
			" degree_study = $13," +
			" relation_ship = $14," +
			" curp = $15," +
			" rfc = $16," +
			" citizen_id = $17," +
			" credential_id = $18," +
			" origin_state = $19," +
			" score = $20," +
			" qualities = $21," +
			" defects = $22," +
			" origin_branch_id = $23," +
			" branch_id = $24 WHERE id = $25" +
			" RETURNING id" +
			" created_at," +
			" updated_at," +
			" deleted_at," +
			" name," +
			" lastname," +
			" username," +
			" password," +
			" photo," +
			" verified," +
			" warning," +
			" darktheme," +
			" active_contract," +
			" address," +
			" born," +
			" degree_study," +
			" relation_ship," +
			" curp," +
			" rfc," +
			" citizen_id," +
			" credential_id," +
			" origin_state," +
			" score," +
			" qualities," +
			" defects," +
			" origin_branch_id," +
			" branch_id"
		err = db.QueryRow(query,
			query,
			updatingUser.CreatedAt,
			time.Now(),
			updatingUser.DeletedAt,
			updatingUser.Name,
			updatingUser.Lastname,
			updatingUser.Password,
			updatingUser.Photo,
			updatingUser.Verified,
			updatingUser.Darktheme,
			updatingUser.ActiveContract,
			updatingUser.Address,
			updatingUser.Born,
			updatingUser.DegreeStudy,
			updatingUser.RelationShip,
			updatingUser.Curp,
			updatingUser.Rfc,
			updatingUser.CitizenID,
			updatingUser.CredentialID,
			updatingUser.OriginState,
			updatingUser.Score,
			updatingUser.Qualities,
			updatingUser.Defects,
			updatingUser.OriginBranchID,
			updatingUser.BranchID,
			updatingUser.Username,
			updatingUser.Warning,
		).Scan(
			&newUser.Id,
			&newUser.CreatedAt,
			&newUser.UpdatedAt,
			&newUser.DeletedAt,
			&newUser.Name,
			&newUser.Lastname,
			&newUser.Username,
			&newUser.Password,
			&newUser.Photo,
			&newUser.Verified,
			&newUser.Warning,
			&newUser.Darktheme,
			&newUser.ActiveContract,
			&newUser.Address,
			&newUser.Born,
			&newUser.DegreeStudy,
			&newUser.RelationShip,
			&newUser.Curp,
			&newUser.Rfc,
			&newUser.CitizenID,
			&newUser.CredentialID,
			&newUser.OriginState,
			&newUser.Score,
			&newUser.Qualities,
			&newUser.Defects,
			&newUser.OriginBranchID,
			&newUser.BranchID)

		return err
	}

	if userIsDeleted(newUser.Id) {
		return errors.New("user is deleted")
	}

	query := "UPDATE users SET" +
		" updated_at = $1," +
		" name = $2," +
		" lastname = $3," +
		" username = $24," +
		" password = $4," +
		" photo = $5," +
		" verified = $6," +
		" warning = $25," +
		" darktheme = $7," +
		" active_contract = $8," +
		" address = $9," +
		" born = $10," +
		" degree_study = $11," +
		" relation_ship = $12," +
		" curp = $13," +
		" rfc = $14," +
		" citizen_id = $15," +
		" credential_id = $16," +
		" origin_state = $17," +
		" score = $18," +
		" qualities = $19," +
		" defects = $20," +
		" origin_branch_id = $21," +
		" branch_id = $22 WHERE id = $23" +
		" RETURNING id," +
		" created_at," +
		" updated_at," +
		" deleted_at," +
		" name," +
		" lastname," +
		" username," +
		" password," +
		" photo," +
		" verified," +
		" warning," +
		" darktheme," +
		" active_contract," +
		" address," +
		" born," +
		" degree_study," +
		" relation_ship," +
		" curp," +
		" rfc," +
		" citizen_id," +
		" credential_id," +
		" origin_state," +
		" score," +
		" qualities," +
		" defects," +
		" origin_branch_id," +
		" branch_id"
	err = db.QueryRow(query,
		query,
		time.Now(),
		updatingUser.Name,
		updatingUser.Lastname,
		updatingUser.Password,
		updatingUser.Photo,
		updatingUser.Verified,
		updatingUser.Darktheme,
		updatingUser.ActiveContract,
		updatingUser.Address,
		updatingUser.Born,
		updatingUser.DegreeStudy,
		updatingUser.RelationShip,
		updatingUser.Curp,
		updatingUser.Rfc,
		updatingUser.CitizenID,
		updatingUser.CredentialID,
		updatingUser.OriginState,
		updatingUser.Score,
		updatingUser.Qualities,
		updatingUser.Defects,
		updatingUser.OriginBranchID,
		updatingUser.BranchID,
		newUser.Id,
		updatingUser.Username,
		updatingUser.Warning).Scan(
		&newUser.Id,
		&newUser.CreatedAt,
		&newUser.UpdatedAt,
		&newUser.DeletedAt,
		&newUser.Name,
		&newUser.Username,
		&newUser.Lastname,
		&newUser.Password,
		&newUser.Photo,
		&newUser.Verified,
		&newUser.Warning,
		&newUser.Darktheme,
		&newUser.ActiveContract,
		&newUser.Address,
		&newUser.Born,
		&newUser.DegreeStudy,
		&newUser.RelationShip,
		&newUser.Curp,
		&newUser.Rfc,
		&newUser.CitizenID,
		&newUser.CredentialID,
		&newUser.OriginState,
		&newUser.Score,
		&newUser.Qualities,
		&newUser.Defects,
		&newUser.OriginBranchID,
		&newUser.BranchID)

	return err
}

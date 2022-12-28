package routes

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/database"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/models"
)

// Save a log in the database
func saveServerActionLog(serverLog models.ServerLogs) {
	db := database.Connect()
	defer db.Close()

	query := "INSERT INTO server_logs (transaction, user_id, branch_id, root, created_at) VALUES ($1, $2, $3, $4, $5)"
	_, err := db.Exec(query, serverLog.Transaction, serverLog.UserID, serverLog.BranchID, serverLog.Root, time.Now())

	if err != nil {
		log.Println("Error saving server action log: " + err.Error())
	}

	log.Println(serverLog.Transaction + " Root: " + strconv.FormatBool(*serverLog.Root))
}

// returns userID, accessLevel, error
func validateUser(input string, password string) (uint, uint, error) {

	key := cleanSpaces(input)
	var query string

	var err error

	query = "SELECT id, password, verified FROM users WHERE username = $1"
	err = errors.New("user does not exist: incorrect username")

	if isANumber(key) {
		query = "SELECT user.id, user.password, user.verified FROM users user, user_phones phone " +
			"WHERE user.id = phone.user_id AND phone.main = true AND phone.phone = $1"
		err = errors.New("user does not exist: incorrect number phone")
	}

	if isAMail(key) {
		query = "SELECT user.id, user.password, user.verified FROM users user, user_mails mail " +
			"WHERE user.id = mail.user_id AND mail.main = true AND mail.email = $1"
		err = errors.New("user does not exist: incorrect email")
	}

	var bufferId uint
	var bufferPassword string
	var bufferVerified bool
	db := database.Connect()
	errQuery := db.QueryRow(query, key).Scan(&bufferId, &bufferPassword, &bufferVerified)

	if errQuery != nil {
		return 0, 0, err
	}

	if bufferPassword != password {
		return 0, 0, errors.New("incorrect password")
	}

	if !bufferVerified {
		return 0, 0, errors.New("user is not verified")
	}

	access_level, errRole := getUserMaxAccessLevel(bufferId)

	return bufferId, access_level, errRole
}

func getUserMaxAccessLevel(id uint) (uint, error) {

	db := database.Connect()
	defer db.Close()

	var access_level uint
	query := "SELECT role.access_level, inherit.deleted_at FROM roles role, inherit_user_roles inherit, users us" +
		" WHERE us.id = $1 AND us.id = inherit.user_id AND inherit.role_id = role.id"
	rows, err := db.Query(query, id)

	if err != nil {
		return 0, errors.New("user does not have a role")
	}

	for rows.Next() {
		var level uint
		var deleted_at *time.Time
		rows.Scan(&level, &deleted_at)

		if level > access_level && deleted_at == nil {
			access_level = level
		}
	}

	return access_level, nil
}

func isANumber(input string) bool {
	_, err := strconv.Atoi(input)
	if err != nil {
		return false
	} else {
		return true
	}
}

func isAMail(input string) bool {
	if strings.Contains(input, "@") {
		return true
	} else {
		return false
	}
}

// Prints a log error and sends it to the client
func logcatch(writer http.ResponseWriter, status int, err error) {
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(status)
		writer.Write([]byte(err.Error()))
	}
}

func cleanSpaces(stringToClean string) string {
	result := strings.ReplaceAll(stringToClean, " ", "")
	return result
}

/*func isRoot(id uint) bool {
	level, _ := getUserMaxAccessLevel(id)
	return level == 1
}*/

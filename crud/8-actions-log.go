package crud

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	commons "github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/commonFunctions"

	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/database"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/models"
)

func GetLogs(pagination models.Pagination) ([]models.ServerLogs, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "server_logs"
	query, _, err := commons.GetQuery(tableName, models.ServerLogs{}, "SELECT", false)
	if err != nil {
		return nil, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	// nil pointers to default values
	if pagination.Page == nil {
		pagination.Page = new(int)
	}
	if pagination.To == nil {
		pagination.To = &time.Time{}
	}
	if pagination.Since == nil {
		pagination.Since = &time.Time{}
	}
	if pagination.Today == nil {
		pagination.Today = new(bool)
	}

	// ------------------- CASE CONTROL Today > Page > Since > To

	//Verify intervals and order
	if pagination.Since.After(*pagination.To) && !pagination.To.IsZero() {
		pagination.Since, pagination.To = pagination.To, pagination.Since
	}

	//Set To to today hour 23:59:59 when Since exists and is unknown
	if pagination.To.IsZero() && !pagination.Since.IsZero() {
		today, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
		todayPtr := today.Add(24 * time.Hour)
		pagination.To = &todayPtr
	}

	//Set Since and To to hour 00:00:00 and 23:59:59
	if *pagination.Today {
		sincePtr, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
		pagination.Since = &sincePtr
		toPtr := sincePtr.Add(24 * time.Hour)
		pagination.To = &toPtr
	}

	//Get logs by intervals of 30 items
	if pagination.Page != nil && *pagination.Page > 0 && !*pagination.Today {
		query += fmt.Sprintf(" ORDER BY created_at DESC LIMIT 30 OFFSET %d", (*pagination.Page-1)*30) // if wrong try ORDER BY id DESC
	} else {
		sinceSplit := strings.Split(pagination.Since.String(), " ")
		since := fmt.Sprintf("%s %s", sinceSplit[0], sinceSplit[1])
		toSplit := strings.Split(pagination.To.String(), " ")
		to := fmt.Sprintf("%s %s", toSplit[0], toSplit[1])
		query += fmt.Sprintf(" WHERE created_at BETWEEN '%s' AND '%s' ORDER BY created_at DESC", since, to)
	}
	// ------------------- EXECUTE -------------------

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	var logs []models.ServerLogs

	for rows.Next() {
		var log models.ServerLogs
		rows.Scan(
			&log.Id,
			&log.CreatedAt,
			&log.Transaction,
			&log.UserID,
			&log.BranchID,
			&log.Root,
		)
		logs = append(logs, log)
	}

	if len(logs) == 0 {
		return nil, errors.New("no logs found")
	}

	return logs, nil
}

func NewServerActionLog(serverLog models.ServerLogs) {
	db := database.Connect()
	defer db.Close()

	currentDate := time.Now()
	serverLog.CreatedAt = currentDate

	tableName := "server_logs"
	query, data, err := commons.GetQuery(tableName, serverLog, "INSERT", false)
	if err != nil {
		log.Printf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	_, err = db.Exec(query, data...)
	if err != nil {
		log.Println("error saving server action log: " + err.Error())
	}

	log.Println(serverLog.Transaction + " Root: " + strconv.FormatBool(*serverLog.Root))
}

func DeleteLogs() error {
	db := database.Connect()
	defer db.Close()

	query := "TRUNCATE TABLE server_logs"
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("can't execute the query ERROR: %s", err.Error())
	}

	return nil
}

//

func NewNotification(notification *models.AdminNotification) error {
	db := database.Connect()
	defer db.Close()

	notification.Id = 0 // To don't add id field to query
	currentDate := time.Now()
	notification.UpdatedAt = &currentDate
	notification.CreatedAt = &currentDate

	tableName := "notifications"
	query, data, err := commons.GetQuery(tableName, *notification, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&notification.Id,
		&notification.CreatedAt,
		&notification.UpdatedAt,
		&notification.DeletedAt,
		&notification.Type,
		&notification.Solved,
		&notification.Description,
		&notification.BranchID,
		&notification.UserID,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if notification.Id == 0 {
		return errors.New("notification not found")
	}

	return nil
}

func GetNotifications(pagination models.Pagination, solved string, root bool, relationalIDs *map[string]uint) ([]models.AdminNotification, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "notifications"
	query, _, err := commons.GetQuery(tableName, models.AdminNotification{}, "SELECT", false)
	if err != nil {
		return nil, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	// nil pointers to default values
	if pagination.Page == nil {
		pagination.Page = new(int)
	}
	if pagination.To == nil {
		pagination.To = &time.Time{}
	}
	if pagination.Since == nil {
		pagination.Since = &time.Time{}
	}
	if pagination.Today == nil {
		pagination.Today = new(bool)
	}

	// ------------------- CASE CONTROL Today > Page > Since > To

	//Verify intervals and order
	if pagination.Since.After(*pagination.To) && !pagination.To.IsZero() {
		pagination.Since, pagination.To = pagination.To, pagination.Since
	}

	//Set To to today hour 23:59:59 when Since exists and is unknown
	if pagination.To.IsZero() && !pagination.Since.IsZero() {
		today, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
		todayPtr := today.Add(24 * time.Hour)
		pagination.To = &todayPtr
	}

	//Set Since and To to hour 00:00:00 and 23:59:59
	if *pagination.Today {
		sincePtr, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
		pagination.Since = &sincePtr
		toPtr := sincePtr.Add(24 * time.Hour)
		pagination.To = &toPtr
	}

	//Get logs by intervals of 30 items
	if pagination.Page != nil && *pagination.Page > 0 && !*pagination.Today {
		deleteNull := ""
		if !root {
			deleteNull = "WHERE deleted_at IS NULL"
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

		var solvedQuery string
		switch solved {
		case "true":
			solvedQuery = "AND solved = 'true'"
		case "false":
			solvedQuery = "AND solved = 'false'"
		default:
			solvedQuery = ""
		}

		query += fmt.Sprintf(" %s %s ORDER BY created_at DESC LIMIT 30 OFFSET %d", deleteNull, solvedQuery, (*pagination.Page-1)*30) // if wrong try ORDER BY id DESC
	} else {
		sinceSplit := strings.Split(pagination.Since.String(), " ")
		since := fmt.Sprintf("%s %s", sinceSplit[0], sinceSplit[1])
		toSplit := strings.Split(pagination.To.String(), " ")
		to := fmt.Sprintf("%s %s", toSplit[0], toSplit[1])

		deleteNull := ""
		if !root {
			deleteNull = "deleted_at IS NULL AND"
		}

		relationQuery := ""
		if relationalIDs != nil {
			var relationConditions []string
			for key, value := range *relationalIDs {
				relationConditions = append(relationConditions, fmt.Sprintf("%s = %d", key, value))
			}
			relationQuery += strings.Join(relationConditions, " AND ") + " AND "
		}

		var solvedQuery string
		switch solved {
		case "true":
			solvedQuery = "solved = 'true' AND"
		case "false":
			solvedQuery = "solved = 'false' AND"
		default:
			solvedQuery = ""
		}

		query += fmt.Sprintf(" WHERE %s %s %s created_at BETWEEN '%s' AND '%s' ORDER BY created_at DESC", deleteNull, solvedQuery, relationQuery, since, to)

	}

	// ------------------- EXECUTE -------------------

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	var notifications []models.AdminNotification

	for rows.Next() {
		var notification models.AdminNotification
		err = rows.Scan(
			&notification.Id,
			&notification.CreatedAt,
			&notification.UpdatedAt,
			&notification.DeletedAt,
			&notification.Type,
			&notification.Solved,
			&notification.Description,
			&notification.BranchID,
			&notification.UserID,
		)
		if err != nil {
			return nil, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
		}

		notifications = append(notifications, notification)
	}

	if len(notifications) == 0 {
		return nil, errors.New("notifications not found")
	}

	return notifications, nil
}

func GetNotification(id uint, root bool) (models.AdminNotification, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "notifications"
	query, _, err := commons.GetQuery(tableName, models.AdminNotification{}, "SELECT", false)
	if err != nil {
		return models.AdminNotification{}, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}
	query += " WHERE id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	var notification models.AdminNotification
	err = db.QueryRow(query, id).Scan(
		&notification.Id,
		&notification.CreatedAt,
		&notification.UpdatedAt,
		&notification.DeletedAt,
		&notification.Type,
		&notification.Solved,
		&notification.Description,
		&notification.BranchID,
		&notification.UserID,
	)
	if err != nil {
		return notification, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if notification.Id == 0 {
		return notification, errors.New("notification not found")
	}

	return notification, nil
}

func UpdateNotification(notification *models.AdminNotification, root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "notifications"
	currentDate := time.Now()

	if !root {
		notification.UpdatedAt = &currentDate
		notification.CreatedAt = nil // Not allowed to update this field
		notification.DeletedAt = nil // Not allowed to update this field
	}

	if commons.IsDeleted(tableName, notification.Id) && !root {
		return errors.New("notification is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *notification, "UPDATE", true)
	querySplit := strings.Split(query, "RETURNING") // Separate "UPDATE () SET () WHERE id = ()" + <stringToIntroduce> + "()"
	query = fmt.Sprintf("%s AND deleted_at IS NULL RETURNING %s", querySplit[0], querySplit[1])
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&notification.Id,
		&notification.CreatedAt,
		&notification.UpdatedAt,
		&notification.DeletedAt,
		&notification.Type,
		&notification.Solved,
		&notification.Description,
		&notification.BranchID,
		&notification.UserID,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if notification.Id == 0 {
		return errors.New("notification doesn't exist")
	}

	return nil
}

func DeleteNotifications(root bool) error {
	db := database.Connect()
	defer db.Close()

	tableName := "notifications"
	if !root {
		lastIdQuery := fmt.Sprintf("SELECT id FROM %s ORDER BY id DESC LIMIT 1", tableName)
		var lastId uint
		err := db.QueryRow(lastIdQuery).Scan(&lastId)
		if err != nil {
			return fmt.Errorf("can't execute the query ERROR: %s", "last id not found "+err.Error())
		}

		query := fmt.Sprintf("UPDATE %s SET deleted_at = CURRENT_TIMESTAMP WHERE id <= %d", tableName, lastId)
		_, err = db.Exec(query)
		if err != nil {
			return fmt.Errorf("can't execute the query ERROR: %s", err.Error())
		}

		return nil
	}

	query := "TRUNCATE TABLE notifications CASCADE"
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("can't execute the query ERROR: %s", err.Error())
	}

	return nil
}

func NewNotificationImage(notificationImage *models.AdminNotificationImage) error {
	db := database.Connect()
	defer db.Close()

	notificationImage.Id = 0 // To don't add id field to query
	currentDate := time.Now()
	notificationImage.UpdatedAt = &currentDate
	notificationImage.CreatedAt = &currentDate

	tableName := "notification_images"
	query, data, err := commons.GetQuery(tableName, *notificationImage, "INSERT", true)
	if err != nil {
		return fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}

	err = db.QueryRow(query, data...).Scan(
		&notificationImage.Id,
		&notificationImage.CreatedAt,
		&notificationImage.UpdatedAt,
		&notificationImage.DeletedAt,
		&notificationImage.Image,
		&notificationImage.NotificationID,
	)

	if err != nil {
		return fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}

	if notificationImage.Id == 0 {
		return errors.New("notification not found")
	}

	return nil
}

func GetNotificationImages(id uint, root bool) ([]models.AdminNotificationImage, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "notification_images"
	query, _, err := commons.GetQuery(tableName, models.AdminNotificationImage{}, "SELECT", false)
	if err != nil {
		return nil, fmt.Errorf("can't get the query %s ERROR: %s", tableName, err.Error())
	}
	query += " WHERE notification_id = $1"

	if !root {
		query += " AND deleted_at IS NULL"
	}

	var images []models.AdminNotificationImage
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("can't execute the %s query ERROR: %s", tableName, err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var image models.AdminNotificationImage
		rows.Scan(
			&image.Id,
			&image.CreatedAt,
			&image.UpdatedAt,
			&image.DeletedAt,
			&image.Image,
			&image.NotificationID,
		)
		images = append(images, image)
	}

	if len(images) == 0 {
		return images, fmt.Errorf("%s not found", tableName)
	}

	return images, nil
}

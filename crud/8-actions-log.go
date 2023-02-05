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
	query, _, err := commons.GetQuery(tableName, models.ServerLogs{}, 2, false)
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
	query, data, err := commons.GetQuery(tableName, serverLog, 0, false)
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
	query, data, err := commons.GetQuery(tableName, *notification, 0, true)
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

func GetNotifications(pagination models.Pagination, root bool) ([]models.AdminNotification, error) {
	db := database.Connect()
	defer db.Close()

	tableName := "notifications"
	query, _, err := commons.GetQuery(tableName, models.AdminNotification{}, 2, false)
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

		query += fmt.Sprintf(" %s ORDER BY created_at DESC LIMIT 30 OFFSET %d", deleteNull, (*pagination.Page-1)*30) // if wrong try ORDER BY id DESC
	} else {
		sinceSplit := strings.Split(pagination.Since.String(), " ")
		since := fmt.Sprintf("%s %s", sinceSplit[0], sinceSplit[1])
		toSplit := strings.Split(pagination.To.String(), " ")
		to := fmt.Sprintf("%s %s", toSplit[0], toSplit[1])

		deleteNull := ""
		if !root {
			deleteNull = "deleted_at IS NULL AND"
		}

		query += fmt.Sprintf(" WHERE %s created_at BETWEEN '%s' AND '%s' ORDER BY created_at DESC", deleteNull, since, to)

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
	query, _, err := commons.GetQuery(tableName, models.AdminNotification{}, 2, false)
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
	notification.UpdatedAt = &currentDate
	notification.CreatedAt = nil
	notification.DeletedAt = nil

	//TODO: Reduce the code with commons.IsDeleted and !root. Do this with all the crud
	if root {
		query, data, err := commons.GetQuery(tableName, *notification, 1, true)
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

		return nil
	}

	if commons.IsDeleted(tableName, notification.Id) {
		return errors.New("notification is deleted")
	}

	query, data, err := commons.GetQuery(tableName, *notification, 1, true)

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

	return nil
}

func DeleteNotifications() error {
	db := database.Connect()
	defer db.Close()

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
	query, data, err := commons.GetQuery(tableName, *notificationImage, 0, true)
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
	query, _, err := commons.GetQuery(tableName, models.AdminNotificationImage{}, 2, false)
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
		return images, errors.New("no images found")
	}

	return images, nil
}

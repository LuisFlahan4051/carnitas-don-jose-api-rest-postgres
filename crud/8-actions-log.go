package crud

import (
	"errors"
	"time"

	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/database"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/models"
)

func GetLogs(pagination models.Pagination) ([]models.ServerLogs, error) {
	var err error
	var query string
	var logs []models.ServerLogs

	db := database.Connect()
	defer db.Close()

	// ------------------- CASE CONTROL Today > Page > Since > To

	//Verify intervals and order
	if pagination.Since.After(pagination.To) && !pagination.To.IsZero() {
		pagination.Since, pagination.To = pagination.To, pagination.Since
	}

	//Set To to today hour 23:59:59 when Since exists and is unknown
	if pagination.To.IsZero() && !pagination.Since.IsZero() {
		today, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
		pagination.To = today.Add(24 * time.Hour)
	}

	//Get logs by intervals of 30 items
	if pagination.Page != nil && *pagination.Page > 0 && !pagination.Today {
		query = "SELECT id," +
			" created_at," +
			" transaction," +
			" user_id," +
			" branch_id," +
			" root FROM server_logs" +
			" ORDER BY id DESC LIMIT 30 OFFSET $1"
		rows, err := db.Query(query, (*pagination.Page-1)*30)

		if err != nil {
			return logs, err
		}

		var logs []models.ServerLogs
		for rows.Next() {
			var log models.ServerLogs
			rows.Scan(
				&log.Id,
				&log.CreateAt,
				&log.Transaction,
				&log.UserID,
				&log.BranchID,
				&log.Root,
			)
			logs = append(logs, log)
		}

		if len(logs) == 0 {
			return logs, errors.New("no logs found")
		}
		return logs, nil
	}

	//Set Since and To to hour 00:00:00 and 23:59:59
	if pagination.Today {
		pagination.Since, _ = time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
		pagination.To = pagination.Since.Add(24 * time.Hour)
	}

	// ------------------- QUERY

	query = "SELECT id," +
		" created_at," +
		" transaction," +
		" user_id," +
		" branch_id," +
		" root FROM server_logs" +
		" WHERE created_at BETWEEN $1 AND $2" +
		" ORDER BY created_at DESC"
	rows, err := db.Query(query, pagination.Since, pagination.To)

	if err != nil {
		return logs, err
	}

	for rows.Next() {
		var log models.ServerLogs
		rows.Scan(
			&log.Id,
			&log.CreateAt,
			&log.Transaction,
			&log.UserID,
			&log.BranchID,
			&log.Root,
		)
		logs = append(logs, log)
	}

	if len(logs) == 0 {
		return logs, errors.New("no logs found")
	}

	return logs, nil
}

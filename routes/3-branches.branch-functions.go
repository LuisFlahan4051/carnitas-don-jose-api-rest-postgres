package routes

import (
	"database/sql"
	"errors"
	"time"

	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/database"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/models"
)

func newBranch(branch *models.Branch) error {
	db := database.Connect()
	defer db.Close()

	var err error

	query := "INSERT INTO branches (name, address) VALUES ($1, $2) RETURNING id"

	err = db.QueryRow(query, &branch.Name, &branch.Address).Scan(&branch.Id)

	return err
}

func getBranch(id uint, root bool) (models.Branch, error) {

	db := database.Connect()
	defer db.Close()

	var branch models.Branch
	var query string
	var err error

	query = "SELECT" +
		" id," +
		" created_at," +
		" updated_at," +
		" deleted_at," +
		" name," +
		" address FROM branches WHERE id = $1 AND deleted_at IS NULL"

	if root {
		query = "SELECT" +
			" id," +
			" created_at," +
			" updated_at," +
			" deleted_at," +
			" name," +
			" address FROM branches WHERE id = $1"
	}

	err = db.QueryRow(query, id).Scan(
		&branch.Id,
		&branch.CreatedAt,
		&branch.UpdatedAt,
		&branch.DeletedAt,
		&branch.Name,
		&branch.Address,
	)

	return branch, err
}

func getBranches(root bool) ([]models.Branch, error) {

	db := database.Connect()
	defer db.Close()

	var query string
	var err error
	var rows *sql.Rows
	var branches []models.Branch

	query = "SELECT" +
		" id," +
		" created_at," +
		" updated_at," +
		" deleted_at," +
		" name," +
		" address FROM branches WHERE deleted_at IS NULL"

	if root {
		query = "SELECT" +
			" id," +
			" created_at," +
			" updated_at," +
			" deleted_at," +
			" name," +
			" address FROM branches"
	}

	rows, err = db.Query(query)

	for rows.Next() {
		var branch models.Branch
		rows.Scan(
			&branch.Id,
			&branch.CreatedAt,
			&branch.UpdatedAt,
			&branch.DeletedAt,
			&branch.Name,
			&branch.Address,
		)
		branches = append(branches, branch)
	}

	return branches, err

}

func branchIsDeleted(id uint) bool {
	db := database.Connect()
	defer db.Close()

	deleted := "SELECT deleted_at FROM branches WHERE id = $1"

	var deletedAt *time.Time

	err := db.QueryRow(deleted, id).Scan(&deletedAt)

	if err != nil {
		return true
	}

	return deletedAt != nil
}

func deleteBranch(id uint, root bool) error {
	db := database.Connect()
	defer db.Close()

	var deleteQuery string
	var result sql.Result

	if root {
		deleteQuery = "DELETE FROM branches WHERE id = $1"
		_, err := db.Exec(deleteQuery, id)
		return err
	}

	if branchIsDeleted(id) {
		return errors.New("branch is already deleted")
	}

	deleteQuery = "UPDATE branches SET deleted_at = $1 WHERE id = $2"
	result, _ = db.Exec(deleteQuery, time.Now(), id)
	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return errors.New("can't delete the branch")
	}

	return nil
}

func updateBranch(branch *models.Branch, root bool) error {
	db := database.Connect()
	defer db.Close()

	var updateQuery string

	if root {
		updateQuery = "UPDATE branches SET" +
			" created_at = $1," +
			" updated_at = $2," +
			" deleted_at = $3," +
			" name = $4," +
			" address = $5 WHERE id = $6" +
			" RETURNING id," +
			" created_at," +
			" updated_at," +
			" deleted_at," +
			" name," +
			" address"
		err := db.QueryRow(
			updateQuery,
			branch.CreatedAt,
			time.Now(),
			&branch.DeletedAt,
			branch.Name,
			branch.Address,
			branch.Id,
		).Scan(
			&branch.Id,
			&branch.CreatedAt,
			&branch.UpdatedAt,
			&branch.DeletedAt,
			&branch.Name,
			&branch.Address,
		)

		return err
	}

	if branchIsDeleted(branch.Id) {
		return errors.New("branch doesn't exist")
	}

	updateQuery = "UPDATE branches SET" +
		" updated_at = $1," +
		" name = $2," +
		" address = $3 WHERE id = $4" +
		" RETURNING id," +
		" created_at," +
		" updated_at," +
		" deleted_at," +
		" name," +
		" address"
	err := db.QueryRow(
		updateQuery,
		time.Now(),
		branch.Name,
		branch.Address,
		branch.Id,
	).Scan(
		&branch.Id,
		&branch.CreatedAt,
		&branch.UpdatedAt,
		&branch.DeletedAt,
		&branch.Name,
		&branch.Address,
	)

	return err
}

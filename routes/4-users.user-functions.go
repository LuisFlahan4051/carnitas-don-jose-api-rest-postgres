package routes

import (
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/database"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/models"
)

func getUser(id string) (models.User, error) {
	var user models.User

	db := database.Connect()

	query := "SELECT id," +
		" created_at," +
		" updated_at," +
		" deleted_at," +
		" name," +
		" lastname," +
		" password," +
		" photo," +
		" verified," +
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
	db.QueryRow(query, id).Scan(
		&user.Id,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
		&user.Name,
		&user.Lastname,
		&user.Password,
		&user.Photo,
		&user.Verified,
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

	queryPhonesAssociated := "SELECT" +
		" id," +
		" created_at," +
		" updated_at," +
		" deleted_at," +
		" phone," +
		" main," +
		" user_id FROM user_phones WHERE user_id = $1"
	rows, err := db.Query(queryPhonesAssociated, id)

	for rows.Next() {
		var phone models.UserPhone
		rows.Scan(
			&phone.Id,
			&phone.CreatedAt,
			&phone.UpdatedAt,
			&phone.DeletedAt,
			&phone.Phone,
			&phone.Main,
			&phone.UserID)

		user.UserPhones = append(user.UserPhones, phone)
	}

	queryEmailsAssociated := "SELECT" +
		" id," +
		" created_at," +
		" updated_at," +
		" deleted_at," +
		" mail," +
		" main," +
		" user_id FROM user_mails WHERE user_id = $1"
	rows, _ = db.Query(queryEmailsAssociated, id)

	for rows.Next() {
		var email models.UserMail
		rows.Scan(
			&email.Id,
			&email.CreatedAt,
			&email.UpdatedAt,
			&email.DeletedAt,
			&email.Mail,
			&email.Main,
			&email.UserID)

		user.UserMails = append(user.UserMails, email)
	}

	queryRolesAssociated := "SELECT" +
		" role.id," +
		" role.created_at," +
		" role.updated_at," +
		" role.deleted_at," +
		" role.name," +
		" role.access_level FROM inherit_user_roles inherit," +
		" roles role WHERE inherit.user_id = $1 AND inherit.role_id = role.id"
	rows, _ = db.Query(queryRolesAssociated, id)

	for rows.Next() {
		var role models.Role
		rows.Scan(
			&role.Id,
			&role.CreatedAt,
			&role.UpdatedAt,
			&role.DeletedAt,
			&role.Name,
			&role.AccessLevel)

		user.InheritUserRoles = append(user.InheritUserRoles, role)
	}

	queryUserReports := "SELECT" +
		" id," +
		" created_at," +
		" updated_at," +
		" deleted_at," +
		" reason," +
		" user_id FROM user_reports WHERE user_id = $1"
	rows, _ = db.Query(queryUserReports, id)

	for rows.Next() {
		var report models.UserReport
		rows.Scan(
			&report.Id,
			&report.CreatedAt,
			&report.UpdatedAt,
			&report.DeletedAt,
			&report.Reason,
			&report.UserID)

		user.UserReports = append(user.UserReports, report)
	}

	queryUserMonetaryBoundsAssociated := "SELECT" +
		" id," +
		" created_at," +
		" updated_at," +
		" deleted_at," +
		" reason," +
		" bound," +
		" user_id FROM monetary_bounds WHERE user_id = $1"
	rows, _ = db.Query(queryUserMonetaryBoundsAssociated, id)

	for rows.Next() {
		var monetaryBound models.MonetaryBound
		rows.Scan(
			&monetaryBound.Id,
			&monetaryBound.CreatedAt,
			&monetaryBound.UpdatedAt,
			&monetaryBound.DeletedAt,
			&monetaryBound.Reason,
			&monetaryBound.Bound,
			&monetaryBound.UserID)

		user.MonetaryBounds = append(user.MonetaryBounds, monetaryBound)
	}

	queryUserMonetaryDiscountsAssociated := "SELECT" +
		" id," +
		" created_at," +
		" updated_at," +
		" deleted_at," +
		" reason," +
		" bound," +
		" user_id FROM monetary_discounts WHERE user_id = $1"
	rows, _ = db.Query(queryUserMonetaryDiscountsAssociated, id)

	for rows.Next() {
		var monetaryDiscount models.MonetaryDiscount
		rows.Scan(
			&monetaryDiscount.Id,
			&monetaryDiscount.CreatedAt,
			&monetaryDiscount.UpdatedAt,
			&monetaryDiscount.DeletedAt,
			&monetaryDiscount.Reason,
			&monetaryDiscount.Discount,
			&monetaryDiscount.UserID)

		user.MonetaryDiscounts = append(user.MonetaryDiscounts, monetaryDiscount)
	}

	queryBranchUserRolesAssociated := "SELECT" +
		" role.id," +
		" role.created_at," +
		" role.updated_at," +
		" role.deleted_at," +
		" role.name," +
		" role.access_level FROM branch_user_roles branch_role, roles role" +
		" WHERE branch_role.user_id = $1 AND branch_role.role_id = role.id"
	rows, _ = db.Query(queryBranchUserRolesAssociated, id)

	for rows.Next() {
		var role models.Role
		rows.Scan(
			&role.Id,
			&role.CreatedAt,
			&role.UpdatedAt,
			&role.DeletedAt,
			&role.Name,
			&role.AccessLevel)

		user.BranchUserRoles = append(user.BranchUserRoles, role)
	}

	queryTurnsAssociated := "SELECT" +
		" id," +
		" created_at," +
		" updated_at," +
		" deleted_at," +
		" start_date," +
		" end_date," +
		" active," +
		" user_id," +
		" branch_id FROM turns WHERE user_id = $1"
	rows, _ = db.Query(queryTurnsAssociated, id)

	for rows.Next() {
		var turn models.Turn
		rows.Scan(
			&turn.Id,
			&turn.CreatedAt,
			&turn.UpdatedAt,
			&turn.DeletedAt,
			&turn.StartDate,
			&turn.EndDate,
			&turn.Active,
			&turn.UserID,
			&turn.BranchID)
		user.Turns = append(user.Turns, turn)
	}

	queryTurnUserRolesAssociated := "SELECT" +
		" id," +
		" created_at," +
		" updated_at," +
		" deleted_at," +
		" login_date," +
		" logout_date," +
		" user_id," +
		" branch_id," +
		" role_id FROM turn_user_roles WHERE user_id = $1"
	rows, _ = db.Query(queryTurnUserRolesAssociated, id)

	for rows.Next() {
		var turnUserRole models.TurnUserRole
		rows.Scan(
			&turnUserRole.Id,
			&turnUserRole.CreatedAt,
			&turnUserRole.UpdatedAt,
			&turnUserRole.DeletedAt,
			&turnUserRole.LoginDate,
			&turnUserRole.LogoutDate,
			&turnUserRole.UserID,
			&turnUserRole.TurnID,
			&turnUserRole.RoleID)
		user.TurnUserRoles = append(user.TurnUserRoles, turnUserRole)
	}

	queryUserSalesAssociated := "SELECT" +
		" id," +
		" created_at," +
		" updated_at," +
		" deleted_at," +
		" entry_date," +
		" exit_date," +
		" table_number," +
		" packed," +
		" user_id," +
		" branch_id," +
		" turn_id FROM sales WHERE user_id = $1"
	rows, _ = db.Query(queryUserSalesAssociated, id)

	for rows.Next() {
		var sale models.Sale
		rows.Scan(
			&sale.Id,
			&sale.CreatedAt,
			&sale.UpdatedAt,
			&sale.DeletedAt,
			&sale.EntryDate,
			&sale.ExitDate,
			&sale.TableNumber,
			&sale.Packed,
			&sale.UserID,
			&sale.BranchID,
			&sale.TurnID)
	}

	queryUserSafeboxActionsAssociated := "SELECT" +
		" id," +
		" created_at," +
		" updated_at," +
		" deleted_at," +
		" withdrawal," +
		" user_id FROM safebox_actions WHERE user_id = $1"
	rows, _ = db.Query(queryUserSafeboxActionsAssociated, id)

	for rows.Next() {
		var safeboxAction models.UserSafeboxAction
		rows.Scan(
			&safeboxAction.Id,
			&safeboxAction.CreatedAt,
			&safeboxAction.UpdatedAt,
			&safeboxAction.DeletedAt,
			&safeboxAction.Withdrawal,
			&safeboxAction.UserID)
		user.UserSafeboxActions = append(user.UserSafeboxActions, safeboxAction)
	}

	queryAdminNotificationsAssociated := "SELECT" +
		" id," +
		" created_at," +
		" updated_at," +
		" deleted_at," +
		" type," +
		" solved," +
		" description," +
		" branch_id," +
		" user_id FROM notifications WHERE user_id = $1"
	rows, _ = db.Query(queryAdminNotificationsAssociated, id)

	for rows.Next() {
		var notification models.AdminNotification
		rows.Scan(
			&notification.Id,
			&notification.CreatedAt,
			&notification.UpdatedAt,
			&notification.DeletedAt,
			&notification.Type,
			&notification.Solved,
			&notification.Description,
			&notification.BranchID,
			&notification.UserID)
		user.AdminNotifications = append(user.AdminNotifications, notification)
	}

	db.Close()

	return user, err
}

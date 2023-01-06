package adminActions

import "github.com/gorilla/mux"

func SetAdminHandleActions(router *mux.Router, URLs *[]string) {

	route := "/notifications"
	router.HandleFunc(route, seeNotifications).Methods("GET")
	*URLs = append(*URLs, route)

	route = "/notification/{notification_id}"
	router.HandleFunc(route, resolveNotification).Methods("PATCH")
	*URLs = append(*URLs, route)

	//

	route = "/branches"
	router.HandleFunc(route, seeBranches).Methods("GET")
	*URLs = append(*URLs, route)

	route = "/branch"
	router.HandleFunc(route, createNewBranch).Methods("POST")
	*URLs = append(*URLs, route)

	route = "/branch/{branch_id}"
	router.HandleFunc(route, changeInfoBranch).Methods("PATCH")
	*URLs = append(*URLs, route)

	route = "/branch/{branch_id}"
	router.HandleFunc(route, dropBranch).Methods("DELETE")
	*URLs = append(*URLs, route)

	//

	route = "/supplies"
	router.HandleFunc(route, seeSupplies).Methods("GET")
	*URLs = append(*URLs, route)

	route = "/supply/{supply_id}"
	router.HandleFunc(route, seeSupply).Methods("GET")
	*URLs = append(*URLs, route)

	route = "/supply"
	router.HandleFunc(route, createNewSupply).Methods("POST")
	*URLs = append(*URLs, route)

	route = "/supply/{supply_id}"
	router.HandleFunc(route, changeSupply).Methods("PATCH")
	*URLs = append(*URLs, route)

	route = "/supply/{supply_id}"
	router.HandleFunc(route, dropSupply).Methods("DELETE")
	*URLs = append(*URLs, route)

	//

	route = "/branch/{branch_id}/supply/{supply_id}"
	router.HandleFunc(route, addSuppliesToBranchStock).Methods("POST")
	*URLs = append(*URLs, route)

	route = "/branch/{branch_id}/supply/{supply_id}"
	router.HandleFunc(route, removeSuppliesFromBranchStock).Methods("DELETE")
	*URLs = append(*URLs, route)

	//

	route = "/articles"
	router.HandleFunc(route, seeArticles).Methods("GET")
	*URLs = append(*URLs, route)

	route = "/article/{article_id}"
	router.HandleFunc(route, seeArticle).Methods("GET")
	*URLs = append(*URLs, route)

	route = "/article"
	router.HandleFunc(route, createNewAritcle).Methods("POST")
	*URLs = append(*URLs, route)

	route = "/article/{article_id}"
	router.HandleFunc(route, changeArticle).Methods("PATCH")
	*URLs = append(*URLs, route)

	route = "/article/{article_id}"
	router.HandleFunc(route, dropArticle).Methods("DELETE")
	*URLs = append(*URLs, route)

	//

	route = "/branch/{branch_id}/article/{article_id}"
	router.HandleFunc(route, addArticlesToBranchStock).Methods("POST")
	*URLs = append(*URLs, route)

	route = "/branch/{branch_id}/article/{article_id}"
	router.HandleFunc(route, removeArticlesFromBranchStock).Methods("PATCH")
	*URLs = append(*URLs, route)

	//

	route = "/branch/{branch_id}/menu/item/{item_id}"
	router.HandleFunc(route, removeMenuItemsFromBranchStock).Methods("DELETE")
	*URLs = append(*URLs, route)

	route = "/branch/{branch_id}/menu/item/{item_id}"
	router.HandleFunc(route, changeMenuItemsFromBranchStock).Methods("PATCH")
	*URLs = append(*URLs, route)

	route = "/food"
	router.HandleFunc(route, changeFood).Methods("PATCH")
	*URLs = append(*URLs, route)

	route = "/food"
	router.HandleFunc(route, dropFood).Methods("DELETE")
	*URLs = append(*URLs, route)

	route = "/drink"
	router.HandleFunc(route, changeDrink).Methods("PATCH")
	*URLs = append(*URLs, route)

	route = "/drink"
	router.HandleFunc(route, dropDrink).Methods("DELETE")
	*URLs = append(*URLs, route)

	//

	route = "/safeboxes"
	router.HandleFunc(route, seeSafeboxes).Methods("GET")
	*URLs = append(*URLs, route)

	route = "/safebox/{safebox_id}"
	router.HandleFunc(route, seeSafebox).Methods("GET")
	*URLs = append(*URLs, route)

	route = "/safebox"
	router.HandleFunc(route, createNewSafeboxToBranch).Methods("POST")
	*URLs = append(*URLs, route)

	//

	route = "/branch/{branch_id}/safebox/{safebox_id}"
	router.HandleFunc(route, changeInfoSafeboxOfBranch).Methods("PATCH")
	*URLs = append(*URLs, route)

	route = "/branch/{branch_id}/safebox/{safebox_id}"
	router.HandleFunc(route, dropSeafeboxOfBranch).Methods("DELETE")
	*URLs = append(*URLs, route)

	route = "/withdraw/safebox/{safebox_id}"
	router.HandleFunc(route, withdrawMoneyFromSafebox).Methods("PATCH")
	*URLs = append(*URLs, route)

	route = "/deposit/safebox/{safebox_id}"
	router.HandleFunc(route, depositMoneyToSafebox).Methods("PATCH")

	//

	route = "/users"
	router.HandleFunc(route, seeUsers).Methods("GET")
	*URLs = append(*URLs, route)

	route = "/user/{user_id}"
	router.HandleFunc(route, seeUser).Methods("GET")
	*URLs = append(*URLs, route)

	route = "/user/{user_id}"
	router.HandleFunc(route, dropUser).Methods("DELETE")
	*URLs = append(*URLs, route)

	route = "/user/{user_id}"
	router.HandleFunc(route, changeUser).Methods("PATCH")
	*URLs = append(*URLs, route)

	route = "/user/{user_id}/make/admin"
	router.HandleFunc(route, makeUserAnAdmin).Methods("POST")
	*URLs = append(*URLs, route)

	route = "/user/{user_id}/make/supervisor"
	router.HandleFunc(route, makeUserAnSupervisor).Methods("POST")
	*URLs = append(*URLs, route)

	route = "/user/{user_id}/reports"
	router.HandleFunc(route, seeUserReports).Methods("GET")
	*URLs = append(*URLs, route)

	//

	route = "/branch/{branch_id}/supervisors"
	router.HandleFunc(route, seeBranchSupervisors).Methods("GET")
	*URLs = append(*URLs, route)

	route = "/turn/{turn_id}"
	router.HandleFunc(route, seeTurn).Methods("GET")
	*URLs = append(*URLs, route)

	route = "/turns"
	router.HandleFunc(route, seeTurns).Methods("GET")
	*URLs = append(*URLs, route)

	route = "/turn/{turn_id}"
	router.HandleFunc(route, dropTurn).Methods("DELETE")
	*URLs = append(*URLs, route)

	route = "/turn/{turn_id}"
	router.HandleFunc(route, changeTurn).Methods("PATCH")
	*URLs = append(*URLs, route)

}

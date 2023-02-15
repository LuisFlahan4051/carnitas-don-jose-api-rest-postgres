package supervisorActions

import "github.com/gorilla/mux"

func SetSupervisorHandleActions(router *mux.Router, URLs *[]string) {

	route := "/branch/{branch_id}"
	router.HandleFunc(route, seeBranch).Methods("GET")
	*URLs = append(*URLs, route)

	route = "/branch/{branch_id}/supplies"
	router.HandleFunc(route, seeSuppliesAtBranchStock).Methods("GET")
	*URLs = append(*URLs, route)

	route = "/branch/{branch_id}/articles"
	router.HandleFunc(route, seeArticlesAtBranchStock).Methods("GET")
	*URLs = append(*URLs, route)

	//

	route = "/user"
	router.HandleFunc(route, registNewUser).Methods("POST")
	*URLs = append(*URLs, route)

	route = "/branch/{branch_id}/users"
	router.HandleFunc(route, seeUsersAtBranch).Methods("GET")
	*URLs = append(*URLs, route)

	route = "/branch/{branch_id}/user/{user_id}/contract/active"
	router.HandleFunc(route, activeUserContranctAtBranch).Methods("PATCH")
	*URLs = append(*URLs, route)

	route = "/branch/{branch_id}/user/{user_id}/contract/desactive"
	router.HandleFunc(route, desactiveUserContractAtBranch).Methods("PATCH")
	*URLs = append(*URLs, route)

	route = "/branch/{branch_id}/user/{user_id}/report"
	router.HandleFunc(route, reportUser).Methods("POST")
	*URLs = append(*URLs, route)

	route = "/branch/{branch_id}/user/{user_id}/role"
	router.HandleFunc(route, setUserRoleAtBranch).Methods("POST")
	*URLs = append(*URLs, route)

	route = "/branch/{branch_id}/user/{user_id}/role/{role_id}"
	router.HandleFunc(route, dropUserRoleFromBranch).Methods("DELETE")
	*URLs = append(*URLs, route)

	route = "/branch/{branch_id}/user/{user_id}/role/{role_id}"
	router.HandleFunc(route, changeUserRoleAtBranch).Methods("PATCH")
	*URLs = append(*URLs, route)

	//

	route = "/notification"
	router.HandleFunc(route, sendNotification).Methods("POST")
	*URLs = append(*URLs, route)

	//

	route = "/food"
	router.HandleFunc(route, createNewFood).Methods("POST")
	*URLs = append(*URLs, route)

	route = "/drink"
	router.HandleFunc(route, createNewDrink).Methods("POST")
	*URLs = append(*URLs, route)

	route = "/menu/item"
	router.HandleFunc(route, createNewMenuIntem).Methods("POST")
	*URLs = append(*URLs, route)

	route = "/branch/{branch_id}/menu/item/{item_id}"
	router.HandleFunc(route, addMenuItemsToBranchStock).Methods("POST")
	*URLs = append(*URLs, route)

	//

	route = "/branch/{branch_id}/turns"
	router.HandleFunc(route, seeTurnsAtBranch).Methods("GET")
	*URLs = append(*URLs, route)

	route = "/branch/{branch_id}/turn/{turn_id}"
	router.HandleFunc(route, seeTurnAtBranch).Methods("GET")
	*URLs = append(*URLs, route)

	route = "/branch/{branch_id}/opened/turn/{turn_id}"
	router.HandleFunc(route, changeActiveTurn).Methods("PATCH")
	*URLs = append(*URLs, route)
}

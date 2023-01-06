package cashierActions

import "github.com/gorilla/mux"

func SetCashierHandleActions(router *mux.Router, URLs *[]string) {
	route := "/branch/{branch_id}/opened/turn"
	router.HandleFunc(route, seeActiveTurn).Methods("GET")
	*URLs = append(*URLs, route)

	route = "/branch/{branch_id}/turn"
	router.HandleFunc(route, openNewTurn).Methods("POST")
	*URLs = append(*URLs, route)

	route = "/branch/{branch_id}/opened/turn/close"
	router.HandleFunc(route, closeTurn).Methods("POST")
	*URLs = append(*URLs, route)

	route = "/branch/{branch_id}/turn/{turn_id}/orders"
	router.HandleFunc(route, seeOrdersAtTurn).Methods("POST")
	*URLs = append(*URLs, route)

	route = "/branch/{branch_id}/turn/{turn_id}/order/{order_id}/close"
	router.HandleFunc(route, closeOrder).Methods("POST")
	*URLs = append(*URLs, route)

	//

	route = "/branch/{branch_id}/turn/{turn_id}/user/{user_id}/role"
	router.HandleFunc(route, setUserRoleAtTurn).Methods("POST")
	*URLs = append(*URLs, route)

	route = "/branch/{branch_id}/turn/{turn_id}/user/{user_id}/role"
	router.HandleFunc(route, changeUserRoleAtTurn).Methods("PATCH")
	*URLs = append(*URLs, route)

	//

	route = "/branch/{branch_id}/turn/{turn_id}/report/losses"
	router.HandleFunc(route, createLossReport).Methods("POST")
	*URLs = append(*URLs, route)
}

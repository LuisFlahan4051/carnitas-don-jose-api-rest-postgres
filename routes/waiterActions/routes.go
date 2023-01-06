package waiterActions

import "github.com/gorilla/mux"

func SetWaiterHandleActions(router *mux.Router, URLs *[]string) {

	route := "/branch/{branch_id}/menu"
	router.HandleFunc(route, seeMenuItemsOfBranch).Methods("GET")
	*URLs = append(*URLs, route)

	//

	route = "/branch/{branch_id}/turn/{turn_id}/order"
	router.HandleFunc(route, openOrder).Methods("POST")
	*URLs = append(*URLs, route)

	route = "/branch/{branch_id}/turn/{turn_id}/order/{order_id}"
	router.HandleFunc(route, cancelOrder).Methods("DELETE")
	*URLs = append(*URLs, route)

	//

	route = "/branch/{branch_id}/turn/{turn_id}/waiter/{user_id}/orders"
	router.HandleFunc(route, seeOrdersOfWaiterAtTurn).Methods("GET")
	*URLs = append(*URLs, route)

	route = "/branch/{branch_id}/turn/{turn_id}/order/{order_id}/products"
	router.HandleFunc(route, addProductsToOrder).Methods("POST")
	*URLs = append(*URLs, route)

	route = "/branch/{branch_id}/turn/{turn_id}/order/{order_id}/comments"
	router.HandleFunc(route, addCommentsToOrder).Methods("POST")
	*URLs = append(*URLs, route)

	route = "/branch/{branch_id}/turn/{turn_id}/order/{order_id}/incomes"
	router.HandleFunc(route, addIncomesToOrder).Methods("POST")
	*URLs = append(*URLs, route)

	route = "/branch/{branch_id}/turn/{turn_id}/order/{order_id}/discounts"
	router.HandleFunc(route, addDiscountToOrder).Methods("POST")
	*URLs = append(*URLs, route)
}

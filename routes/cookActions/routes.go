package cookActions

import "github.com/gorilla/mux"

func SetCookHandleActions(router *mux.Router, URLs *[]string) {

	route := "/branch/{branch_id}/turn/{turn_id}/opened/orders"
	router.HandleFunc(route, seeOpenedOrders).Methods("GET")
	*URLs = append(*URLs, route)

	// state: cooking, prepared
	route = "/branch/{branch_id}/turn/{turn_id}/opened/order/{order_id}/{state}"
	router.HandleFunc(route, changeOrderState).Methods("POST")
	*URLs = append(*URLs, route)

	// state: done, waiting
	route = "/branch/{branch_id}/turn/{turn_id}/order/{order_id}/product/{product_id}/{state}"
	router.HandleFunc(route, changeProductState).Methods("POST")
	*URLs = append(*URLs, route)
}

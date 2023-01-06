package rootActions

import (
	"github.com/gorilla/mux"
)

func SetRootHandleActions(router *mux.Router, URLs *[]string) {

	route := "/user/{user_id}/make/root"
	router.HandleFunc(route, makeUserRoot).Methods("POST")
	*URLs = append(*URLs, route)

	route = "/user/{user_id}/verify"
	router.HandleFunc(route, verifyUser).Methods("POST")
	*URLs = append(*URLs, route)

	// Can add ?page=1 to get the a page or send a jsdon body {"page":1, ...}
	route = "/server/logs"
	router.HandleFunc(route, seeSeverLogs).Methods("GET")
	*URLs = append(*URLs, route)

	route = "/server/logs"
	router.HandleFunc(route, cleanServerLogs).Methods("DELETE")
	*URLs = append(*URLs, route)

	route = "/admins"
	router.HandleFunc(route, seeAdmins).Methods("GET")
	*URLs = append(*URLs, route)

}

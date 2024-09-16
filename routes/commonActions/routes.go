package commonActions

import "github.com/gorilla/mux"

func SetCommonHandleActions(router *mux.Router, URLs *[]string) {
	route := "/my/profile"
	router.HandleFunc(route, seeMyProfile).Methods("GET")
	*URLs = append(*URLs, route)

	route = "/users/{user_id}/profile/picture"
	//router.HandleFunc(route, getProfilePicture).Methods("GET")
	*URLs = append(*URLs, route)

	route = "/my/profile/change/credentials"
	router.HandleFunc(route, changeMyMainCredentials).Methods("PATCH")
	*URLs = append(*URLs, route)

	route = "/my/profile/change/profile"
	router.HandleFunc(route, changeMyProfile).Methods("PATCH")
	*URLs = append(*URLs, route)

	route = "/my/profile/change/picture"
	router.HandleFunc(route, changeMyProfilePicture).Methods("PATCH")
	*URLs = append(*URLs, route)

	route = "/login"
	router.HandleFunc(route, validateUser).Methods("POST")
	*URLs = append(*URLs, route)
}

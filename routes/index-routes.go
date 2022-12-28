package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func HomeHandler(writer http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)

	param := params["id"]

	param2 := request.URL.Query().Get("root")

	writer.Write([]byte(param + param2))

	//writer.Write([]byte("Hello World"))

}

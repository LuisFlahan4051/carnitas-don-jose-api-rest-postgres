package rootActions

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	commons "github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/commonFunctions"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/crud"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/models"
	"github.com/gorilla/mux"
)

func makeUserRoot(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(mux.Vars(request)["user_id"])
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, errors.New("invalid user/{id} value"))
		return
	}

	_, _, err = commons.Authentication(request, commons.ROOT)
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}

	//TODO: Check if the user exists and if is already root
	user := models.User{}

	err = crud.NewInheritUserRole(&models.InheritUserRole{
		UserID: uint(id),
		RoleID: 1,
	})
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, err)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(user)
}

func verifyUser(writer http.ResponseWriter, request *http.Request) {}

func seeSeverLogs(writer http.ResponseWriter, request *http.Request) {
	_, _, err := commons.Authentication(request, commons.ROOT)
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}

	var pagination models.Pagination
	if request.URL.Query().Get("page") != "" {
		// Getting the page from URL
		page, err := strconv.Atoi(request.URL.Query().Get("page"))
		if err != nil {
			commons.Logcatch(writer, http.StatusBadRequest, err)
			return
		}
		pagination.Page = &page
	} else {
		// Getting data from body
		err = json.NewDecoder(request.Body).Decode(&pagination)

		if err != nil {
			// If just enter to the route, it will show the logs of today
			today := true
			pagination.Today = &today
		}
	}

	logs, err := crud.GetLogs(pagination)
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, err)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(logs)
}

func cleanServerLogs(writer http.ResponseWriter, request *http.Request) {}

func seeAdmins(writer http.ResponseWriter, request *http.Request) {}

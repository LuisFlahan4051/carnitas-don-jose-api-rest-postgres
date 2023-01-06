package rootActions

import (
	"encoding/json"
	"net/http"
	"strconv"

	commons "github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/commonFunctions"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/crud"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/models"
)

func makeUserRoot(writer http.ResponseWriter, request *http.Request) {}

func verifyUser(writer http.ResponseWriter, request *http.Request) {}

func seeSeverLogs(writer http.ResponseWriter, request *http.Request) {
	_, _, err := commons.Authentication(request, commons.ROOT)
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}

	var pagination models.Pagination

	if request.URL.Query().Get("page") != "" {
		page, err := strconv.Atoi(request.URL.Query().Get("page"))
		if err != nil {
			commons.Logcatch(writer, http.StatusBadRequest, err)
			return
		}
		pagination.Page = &page
	} else {
		err = json.NewDecoder(request.Body).Decode(&pagination)

		if err != nil {
			commons.Logcatch(writer, http.StatusBadRequest, err)
			return
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

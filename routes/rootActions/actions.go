package rootActions

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

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

	adminBufferId, maxAccessLevel, err := commons.Authentication(request, commons.ROOT)
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}
	root := maxAccessLevel == commons.ROOT && request.URL.Query().Get("root") == "true"

	user, err := crud.GetUser(uint(id), root)
	if err != nil {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}

	inheritUserRoles, err := crud.GetInheritUserRoles(root)
	if err != nil && !(strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "no rows in result set")) {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}

	for _, inheritUserRole := range inheritUserRoles {
		if inheritUserRole.UserID == user.Id && inheritUserRole.RoleID == commons.ROOT {
			commons.Logcatch(writer, http.StatusBadRequest, errors.New("user already root"))
			return
		}
	}

	newInheritUserRole := models.InheritUserRole{
		UserID: user.Id,
		RoleID: commons.ROOT,
	}
	err = crud.NewInheritUserRole(&newInheritUserRole)
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, err)
		return
	}

	/*rootRole, err := crud.GetRole(commons.ROOT, root)
	if err != nil {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}

	user.InheritUserRoles = append(user.InheritUserRoles, rootRole)
	*/

	crud.NewServerActionLog(models.ServerLogs{
		UserID:      adminBufferId,
		Root:        &root,
		Transaction: "makeUserRoot",
	})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	// user.Password = ""
	//json.NewEncoder(writer).Encode(user)
	json.NewEncoder(writer).Encode("done")
}

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

		// If just enter to the route, it will show the logs of today
		if err != nil {
			if err.Error() == "EOF" {
				today := true
				pagination.Today = &today
			} else {
				commons.Logcatch(writer, http.StatusBadRequest, err)
				return
			}
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

func cleanServerLogs(writer http.ResponseWriter, request *http.Request) {
	adminBufferId, maxAccessLevel, err := commons.Authentication(request, commons.ROOT)
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}
	root := maxAccessLevel == commons.ROOT && request.URL.Query().Get("root") == "true"
	if !root {
		commons.Logcatch(writer, http.StatusUnauthorized, errors.New("only root can clean logs"))
		return
	}

	err = crud.DeleteLogs()
	if err != nil {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}

	crud.NewServerActionLog(models.ServerLogs{
		UserID:      adminBufferId,
		Root:        &root,
		Transaction: "cleanServerLogs",
	})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode("done")
}

func seeAdmins(writer http.ResponseWriter, request *http.Request) {
	adminBufferId, maxAccessLevel, err := commons.Authentication(request, commons.ROOT)
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}
	root := maxAccessLevel == commons.ROOT && request.URL.Query().Get("root") == "true"
	if !root {
		commons.Logcatch(writer, http.StatusUnauthorized, errors.New("only root can see the admins"))
		return
	}

	admins, err := crud.GetAdmins()
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, err)
		return
	}

	crud.NewServerActionLog(models.ServerLogs{
		UserID:      adminBufferId,
		Root:        &root,
		Transaction: "seeAdmins",
	})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(admins)
}

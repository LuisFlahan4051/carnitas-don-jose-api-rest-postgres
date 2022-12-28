package routes

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/models"
	"github.com/gorilla/mux"
)

func PostBranch(writer http.ResponseWriter, request *http.Request) {
	//Authentication
	bufferID, accessLevel, err := validateUser(request.URL.Query().Get("user"), request.URL.Query().Get("password"))
	if err != nil {
		logcatch(writer, http.StatusBadRequest, err)
		return
	}

	if accessLevel > 2 {
		logcatch(writer, http.StatusBadRequest, errors.New("you don't have access to this resource"))
		return
	}

	//Create branch
	var branch models.Branch
	err = json.NewDecoder(request.Body).Decode(&branch)
	if err != nil {
		logcatch(writer, http.StatusBadRequest, err)
		return
	}

	err = newBranch(&branch)
	if err != nil {
		logcatch(writer, http.StatusBadRequest, err)
		return
	}

	//Log
	isRoot := false
	saveServerActionLog(models.ServerLogs{
		UserID:      bufferID,
		Transaction: "POST /branches",
		Root:        &isRoot,
	})

	//Response
	json.NewEncoder(writer).Encode(branch)
}

func GetBranch(writer http.ResponseWriter, request *http.Request) {
	//Authentication
	params := mux.Vars(request)
	param, err := strconv.Atoi(params["id"])
	if err != nil || param < 1 {
		logcatch(writer, http.StatusBadRequest, err)
		return
	}

	if request.URL.Query().Get("user") == "" || request.URL.Query().Get("password") == "" {
		logcatch(writer, http.StatusBadRequest, errors.New("need credentials to access this resource"))
		return
	}

	bufferID, accessLevel, err := validateUser(request.URL.Query().Get("user"), request.URL.Query().Get("password"))
	if err != nil {
		logcatch(writer, http.StatusBadRequest, err)
		return
	}

	//Get branch
	isRoot := accessLevel == 1 && request.URL.Query().Get("root") == "true"
	branch, err := getBranch(uint(param), isRoot)
	if err != nil {
		logcatch(writer, http.StatusBadRequest, err)
		return
	}

	//Log
	saveServerActionLog(models.ServerLogs{
		UserID:      bufferID,
		Transaction: "GET /branch/" + params["id"],
		Root:        &isRoot,
	})

	//Response
	json.NewEncoder(writer).Encode(branch)
}

func GetBranches(writer http.ResponseWriter, request *http.Request) {
	//Authentication
	if request.URL.Query().Get("user") == "" || request.URL.Query().Get("password") == "" {
		logcatch(writer, http.StatusBadRequest, errors.New("need credentials to access this resource"))
		return
	}

	bufferID, accessLevel, err := validateUser(request.URL.Query().Get("user"), request.URL.Query().Get("password"))
	if err != nil {
		logcatch(writer, http.StatusBadRequest, err)
		return
	}

	//Get branches
	isRoot := accessLevel == 1 && request.URL.Query().Get("root") == "true"
	branches, err := getBranches(isRoot)
	if err != nil {
		logcatch(writer, http.StatusBadRequest, err)
		return
	}

	//Log
	saveServerActionLog(models.ServerLogs{
		UserID:      bufferID,
		Transaction: "GET /branches",
		Root:        &isRoot,
	})

	//Response
	json.NewEncoder(writer).Encode(branches)
}

func DeleteBranch(writer http.ResponseWriter, request *http.Request) {
	//Authentication
	params := mux.Vars(request)
	param, err := strconv.Atoi(params["id"])
	if err != nil || param < 1 {
		logcatch(writer, http.StatusBadRequest, err)
		return
	}

	if request.URL.Query().Get("user") == "" || request.URL.Query().Get("password") == "" {
		logcatch(writer, http.StatusBadRequest, errors.New("need credentials to access this resource"))
		return
	}

	bufferID, accessLevel, err := validateUser(request.URL.Query().Get("user"), request.URL.Query().Get("password"))
	if err != nil {
		logcatch(writer, http.StatusBadRequest, err)
		return
	}

	if accessLevel > 2 {
		logcatch(writer, http.StatusBadRequest, errors.New("you don't have access to this resource"))
		return
	}

	//Delete branch
	isRoot := accessLevel == 1 && request.URL.Query().Get("root") == "true"
	err = deleteBranch(uint(param), isRoot)
	if err != nil {
		logcatch(writer, http.StatusBadRequest, err)
		return
	}

	saveServerActionLog(models.ServerLogs{
		UserID:      bufferID,
		Transaction: "DELETE /branch/" + params["id"],
		Root:        &isRoot,
	})
	writer.WriteHeader(http.StatusOK)
}

func PatchBranch(writer http.ResponseWriter, request *http.Request) {
	//Authentication
	params := mux.Vars(request)
	param, err := strconv.Atoi(params["id"])
	if err != nil || param < 1 {
		logcatch(writer, http.StatusBadRequest, err)
		return
	}

	if request.URL.Query().Get("user") == "" || request.URL.Query().Get("password") == "" {
		logcatch(writer, http.StatusBadRequest, errors.New("need credentials to access this resource"))
		return
	}

	bufferID, accessLevel, err := validateUser(request.URL.Query().Get("user"), request.URL.Query().Get("password"))
	if err != nil {
		logcatch(writer, http.StatusBadRequest, err)
		return
	}

	if accessLevel > 2 {
		logcatch(writer, http.StatusBadRequest, errors.New("you don't have access to this resource"))
		return
	}

	//Patch branch
	var branch models.Branch
	err = json.NewDecoder(request.Body).Decode(&branch)
	if err != nil {
		logcatch(writer, http.StatusBadRequest, err)
		return
	}

	isRoot := accessLevel == 1 && request.URL.Query().Get("root") == "true"
	branch.Id = uint(param)
	err = updateBranch(&branch, isRoot)
	if err != nil {
		logcatch(writer, http.StatusBadRequest, err)
		return
	}

	//Log
	saveServerActionLog(models.ServerLogs{
		UserID:      bufferID,
		Transaction: "PATCH /branch/" + params["id"],
		Root:        &isRoot,
	})

	//Response
	json.NewEncoder(writer).Encode(branch)
}

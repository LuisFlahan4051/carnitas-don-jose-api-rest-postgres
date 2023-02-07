package adminActions

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strconv"
	"strings"

	commons "github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/commonFunctions"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/crud"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/models"
	"github.com/gorilla/mux"
)

func seeNotifications(writer http.ResponseWriter, request *http.Request) {
	adminBufferId, accessLevel, err := commons.Authentication(request, commons.ADMIN)
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}
	root := accessLevel == commons.ROOT && request.URL.Query().Get("root") == "true"

	solved := request.URL.Query().Get("solved")

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

	notifications, err := crud.GetNotifications(pagination, solved, root)
	if err != nil {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}

	for i := 0; i < len(notifications); i++ {
		notifications[i].Images, err = crud.GetNotificationImages(notifications[i].Id, root)
		if err != nil && !strings.Contains(err.Error(), "no images found") {
			commons.Logcatch(writer, http.StatusInternalServerError, err)
			return
		}
	}

	crud.NewServerActionLog(models.ServerLogs{
		UserID:      adminBufferId,
		Root:        &root,
		Transaction: "seeNotifications",
	})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(notifications)
}

func resolveNotification(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(mux.Vars(request)["notification_id"])
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, errors.New("invalid user/{id} value"))
		return
	}

	adminBufferId, accessLevel, err := commons.Authentication(request, commons.ADMIN)
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}
	root := accessLevel == commons.ROOT && request.URL.Query().Get("root") == "true"

	idExists := commons.RegistExists("notifications", uint(id))
	if idExists == 0 {
		commons.Logcatch(writer, http.StatusInternalServerError, errors.New("notification does not exist"))
		return
	}

	notificationToUpdate := models.AdminNotification{
		ID: models.ID{
			Id: uint(id),
		},
		Solved: true,
	}

	err = crud.UpdateNotification(&notificationToUpdate, root)
	if err != nil {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}

	crud.NewServerActionLog(models.ServerLogs{
		UserID:      adminBufferId,
		Root:        &root,
		Transaction: "resolveNotification",
	})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(notificationToUpdate)
}

func dropNotifications(writer http.ResponseWriter, request *http.Request) {
	adminBufferId, accessLevel, err := commons.Authentication(request, commons.ADMIN)
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}
	root := accessLevel == commons.ROOT && request.URL.Query().Get("root") == "true"

	err = crud.DeleteNotifications()
	if err != nil {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}

	if root {
		err = os.RemoveAll("./storage/public/notifications")
		if err != nil {
			commons.Logcatch(writer, http.StatusBadRequest, err)
			return
		}
	}

	crud.NewServerActionLog(models.ServerLogs{
		UserID:      adminBufferId,
		Root:        &root,
		Transaction: "dropNotifications",
	})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode("notifications deleted")
}

//

func seeBranches(writer http.ResponseWriter, request *http.Request) {}

func createNewBranch(writer http.ResponseWriter, request *http.Request) {}

func changeInfoBranch(writer http.ResponseWriter, request *http.Request) {}

func dropBranch(writer http.ResponseWriter, request *http.Request) {}

//

func seeSupplies(writer http.ResponseWriter, request *http.Request) {}

func seeSupply(writer http.ResponseWriter, request *http.Request) {}

func createNewSupply(writer http.ResponseWriter, request *http.Request) {}

func changeSupply(writer http.ResponseWriter, request *http.Request) {}

func dropSupply(writer http.ResponseWriter, request *http.Request) {}

//

func addSuppliesToBranchStock(writer http.ResponseWriter, request *http.Request) {}

func removeSuppliesFromBranchStock(writer http.ResponseWriter, request *http.Request) {}

//

func seeArticles(writer http.ResponseWriter, request *http.Request) {}

func seeArticle(writer http.ResponseWriter, request *http.Request) {}

func createNewAritcle(writer http.ResponseWriter, request *http.Request) {}

func changeArticle(writer http.ResponseWriter, request *http.Request) {}

func dropArticle(writer http.ResponseWriter, request *http.Request) {}

//

func addArticlesToBranchStock(writer http.ResponseWriter, request *http.Request) {}

func removeArticlesFromBranchStock(writer http.ResponseWriter, request *http.Request) {}

//

func removeMenuItemsFromBranchStock(writer http.ResponseWriter, request *http.Request) {}

func changeMenuItemsFromBranchStock(writer http.ResponseWriter, request *http.Request) {}

func changeFood(writer http.ResponseWriter, request *http.Request) {}

func changeDrink(writer http.ResponseWriter, request *http.Request) {}

func dropFood(writer http.ResponseWriter, request *http.Request) {}

func dropDrink(writer http.ResponseWriter, request *http.Request) {}

//

func seeSafeboxes(writer http.ResponseWriter, request *http.Request) {}

func seeSafebox(writer http.ResponseWriter, request *http.Request) {}

func createNewSafeboxToBranch(writer http.ResponseWriter, request *http.Request) {}

//

func changeInfoSafeboxOfBranch(writer http.ResponseWriter, request *http.Request) {}

func dropSeafeboxOfBranch(writer http.ResponseWriter, request *http.Request) {}

func withdrawMoneyFromSafebox(writer http.ResponseWriter, request *http.Request) {}

func depositMoneyToSafebox(writer http.ResponseWriter, request *http.Request) {}

//

func seeUsers(writer http.ResponseWriter, request *http.Request) {
	adminBufferId, accessLevel, err := commons.Authentication(request, commons.ADMIN)
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}

	root := accessLevel == commons.ROOT && request.URL.Query().Get("root") == "true"
	users, err := crud.GetUsers(root)
	if err != nil {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}

	crud.NewServerActionLog(models.ServerLogs{
		UserID:      adminBufferId,
		Root:        &root,
		Transaction: "seeUsers",
	})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(users)
}

func seeUser(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(mux.Vars(request)["user_id"])
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, errors.New("invalid user/{id} value"))
		return
	}

	adminBufferId, accessLevel, err := commons.Authentication(request, commons.ADMIN)
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}

	root := accessLevel == commons.ROOT && request.URL.Query().Get("root") == "true"
	user, err := crud.GetUser(uint(id), root)
	if err != nil {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}

	crud.NewServerActionLog(models.ServerLogs{
		UserID:      adminBufferId,
		Root:        &root,
		Transaction: "seeUser",
	})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(user)
}

func dropUser(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(mux.Vars(request)["user_id"])
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, errors.New("invalid user/{id} value"))
		return
	}

	adminBufferId, accessLevel, err := commons.Authentication(request, commons.ADMIN)
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}

	root := accessLevel == commons.ROOT && request.URL.Query().Get("root") == "true"
	err = crud.DeleteUser(uint(id), root)
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, err)
		return
	}

	if root {
		err = os.RemoveAll("./storage/public/users/" + strconv.Itoa(id))
		if err != nil {
			commons.Logcatch(writer, http.StatusBadRequest, err)
			return
		}
	}

	crud.NewServerActionLog(models.ServerLogs{
		UserID:      adminBufferId,
		Root:        &root,
		Transaction: "dropUser",
	})

	writer.WriteHeader(http.StatusOK)
}

func changeUser(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(mux.Vars(request)["user_id"])
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, errors.New("invalid user/{id} value"))
		return
	}

	adminBufferId, accessLevel, err := commons.Authentication(request, commons.ADMIN)
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}

	root := accessLevel == commons.ROOT && request.URL.Query().Get("root") == "true"

	var user models.User
	err = json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, err)
		return
	}

	user.Id = uint(id)
	user.Username = commons.CleanSpaces(user.Username)
	user.Password = commons.CleanSpaces(user.Password)

	// Allow this if the validation is needed
	/*userExists := commons.UserExists(uint(id))
	if !userExists {
		commons.Logcatch(writer, http.StatusBadRequest, errors.New("user to modify does not exist"))
		return
	}

	userBackUp, _ := crud.GetUser(user.Id, root)

	err = crud.UpdateUser(&user, root)
	if err != nil {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}

	if user.Username == "" || user.Password == "" {
		crud.UpdateUser(&userBackUp, root)
		commons.Logcatch(writer, http.StatusInternalServerError, errors.New("username or password cannot be empty"))
		return
	}*/

	err = crud.UpdateUser(&user, root)
	if err != nil {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}

	crud.NewServerActionLog(models.ServerLogs{
		UserID:      adminBufferId,
		Root:        &root,
		Transaction: "UpdateUser",
	})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(user)
}

func makeUserAnAdmin(writer http.ResponseWriter, request *http.Request) {}

func makeUserAnSupervisor(writer http.ResponseWriter, request *http.Request) {}

func seeBranchSupervisors(writer http.ResponseWriter, request *http.Request) {}

func seeUserReports(writer http.ResponseWriter, request *http.Request) {}

//

func seeTurn(writer http.ResponseWriter, request *http.Request) {}

func seeTurns(writer http.ResponseWriter, request *http.Request) {}

func dropTurn(writer http.ResponseWriter, request *http.Request) {}

func changeTurn(writer http.ResponseWriter, request *http.Request) {}

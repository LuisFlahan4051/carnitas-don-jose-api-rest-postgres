package supervisorActions

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"errors"
	"os"

	"net/http"

	commons "github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/commonFunctions"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/crud"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/models"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

func seeBranch(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(mux.Vars(request)["branch_id"])
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

	branch, err := crud.GetBranch(uint(id), root)
	if err != nil {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}

	relationIDs := make(map[string]uint)
	relationIDs["branch_id"] = branch.Id

	relationBranchSafebox, err := crud.GetBranchSafeboxes(root, &relationIDs)
	if err != nil {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}
	branch.BranchSafeboxes = relationBranchSafebox

	relationBranchProducts, err := crud.GetBranchProductsStock(root, &relationIDs)
	if err != nil && !strings.Contains(err.Error(), "not found") {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}
	branch.BranchProductsStock = relationBranchProducts

	relationBranchSupplies, err := crud.GetBranchSuppliesStock(root, &relationIDs)
	if err != nil && !strings.Contains(err.Error(), "not found") {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}
	branch.BranchSuppliesStock = relationBranchSupplies

	relationBranchArticles, err := crud.GetBranchArticlesStock(root, &relationIDs)
	if err != nil && !strings.Contains(err.Error(), "not found") {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}
	branch.BranchArticlesStock = relationBranchArticles

	branchUsers, err := crud.GetUsers(root, &relationIDs)
	var branchUsersCleaned []models.User
	for _, user := range branchUsers {
		user.Password = ""
		branchUsersCleaned = append(branchUsersCleaned, user)
	}
	if err != nil && !strings.Contains(err.Error(), "not found") {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}
	branch.Users = branchUsersCleaned

	relationBranchUserRoles, err := crud.GetBranchUserRoles(root, &relationIDs)
	if err != nil && !strings.Contains(err.Error(), "not found") {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}
	branch.BranchUserRoles = relationBranchUserRoles

	relationBranchTurns, err := crud.GetTurns(root, &relationIDs)
	if err != nil && !strings.Contains(err.Error(), "not found") {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}
	branch.Turns = relationBranchTurns

	branchSales, err := crud.GetSales(root, &relationIDs)
	if err != nil && !strings.Contains(err.Error(), "not found") {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}
	branch.Sales = branchSales

	branchInventories, err := crud.GetInventories(root, &relationIDs)
	if err != nil && !strings.Contains(err.Error(), "not found") {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}
	branch.Inventories = branchInventories

	var since time.Time
	since = since.Add(time.Hour * 24)
	allNotifications := models.Pagination{
		Since: &since,
	}
	branchAdminNotifications, err := crud.GetNotifications(allNotifications, "", root, &relationIDs)
	if err != nil && !strings.Contains(err.Error(), "not found") {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}
	branch.AdminNotifications = branchAdminNotifications

	crud.NewServerActionLog(models.ServerLogs{
		UserID:      adminBufferId,
		Root:        &root,
		Transaction: "seeBranch",
	})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(branch)
}

func registNewUser(writer http.ResponseWriter, request *http.Request) {
	adminBufferId, maxAccessLevel, err := commons.Authentication(request, commons.SUPERVISOR)
	root := maxAccessLevel == commons.ROOT && request.URL.Query().Get("root") == "true"
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}

	//Need Content-Type: multipart/form-data sending by inputs of a form, Max 33MB
	err = request.ParseMultipartForm(32 << 20) // 32<<20 = 32 * 2^20 = 33,554,432 bits = 32.768 MB
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, err)
		return
	}

	// Getting the data from the form and decode into the user struct
	var user models.User

	username := request.FormValue("username")
	_, _, err = commons.ValidateUser(username, "")
	if !strings.Contains(err.Error(), "user does not exist:") {
		commons.Logcatch(writer, http.StatusBadRequest, errors.New("the username already exist"))
		return
	}

	file, handle, err := request.FormFile("profile_picture")
	existFile := true
	if err != nil {
		if !strings.Contains(err.Error(), "no such file") {
			commons.Logcatch(writer, http.StatusBadRequest, err)
			return
		}
		existFile = false
	}

	adminBuffer, _ := crud.GetUser(adminBufferId, false)
	user.BranchID = adminBuffer.BranchID
	if user.BranchID == nil {
		commons.Logcatch(writer, http.StatusBadRequest, errors.New("root user must have a origin branch"))
		return
	}
	user.OriginBranchID = user.BranchID

	if existFile {
		defer file.Close()

		if !commons.FileIsImage(handle.Filename) {
			commons.Logcatch(writer, http.StatusBadRequest, errors.New("file type not supported, only .jpg, .jpeg and .png are allowed"))
			return
		}

		decoderFormFields := schema.NewDecoder()
		decoderFormFields.SetAliasTag("json")
		decoderFormFields.Decode(&user, request.Form)

		err = crud.NewUser(&user)
		if err != nil {
			commons.Logcatch(writer, http.StatusBadRequest, err)
			return
		}

		user.Id = commons.RegistExists("users", user.Id)
		if user.Id == 0 {
			commons.Logcatch(writer, http.StatusBadRequest, errors.New("user do not created"))
			return
		}

		// Save the file in the server
		newFileName := "profile_picture.webp"
		localPathStorage := fmt.Sprintf("./storage/public/users/%d/profile/picture/", user.Id)

		err = commons.SavePictureAsWebp(file, localPathStorage, newFileName)

		if err != nil {
			crud.DeleteUser(user.Id, true)
			commons.Logcatch(writer, http.StatusBadRequest, err)
			return
		}

		pathStorage := fmt.Sprintf("%s://%s:%s/users/%d/profile/picture/%s", os.Getenv("HTTP"), os.Getenv("SERVERHOST"), os.Getenv("SERVERPORT"), user.Id, newFileName)

		user.Photo = &pathStorage
		err = crud.UpdateUser(&user, false)
		if err != nil {
			crud.DeleteUser(user.Id, true)
			commons.Logcatch(writer, http.StatusBadRequest, err)
			return
		}

		crud.NewServerActionLog(models.ServerLogs{
			UserID:      adminBufferId,
			Root:        &root,
			Transaction: "registNewUser",
		})
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusCreated)
		json.NewEncoder(writer).Encode(user)
		return
	}

	schema.NewDecoder().Decode(&user, request.Form)

	err = crud.NewUser(&user)
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, err)
		return
	}

	crud.NewServerActionLog(models.ServerLogs{
		UserID:      adminBufferId,
		Root:        &root,
		Transaction: "registNewUser",
	})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(user)
}

func seeUsersAtBranch(writer http.ResponseWriter, request *http.Request) {}

func seeSuppliesAtBranchStock(writer http.ResponseWriter, request *http.Request) {}

func seeArticlesAtBranchStock(writer http.ResponseWriter, request *http.Request) {}

//

func activeUserContranctAtBranch(writer http.ResponseWriter, request *http.Request) {}

func desactiveUserContractAtBranch(writer http.ResponseWriter, request *http.Request) {}

func reportUser(writer http.ResponseWriter, request *http.Request) {
}

//

func setUserRoleAtBranch(writer http.ResponseWriter, request *http.Request) {}

func dropUserRoleFromBranch(writer http.ResponseWriter, request *http.Request) {}

func changeUserRoleAtBranch(writer http.ResponseWriter, request *http.Request) {}

//

func sendNotification(writer http.ResponseWriter, request *http.Request) {
	adminBufferId, maxAccessLevel, err := commons.Authentication(request, commons.SUPERVISOR)
	root := maxAccessLevel == commons.ROOT && request.URL.Query().Get("root") == "true"
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}

	//Need Content-Type: multipart/form-data sending by inputs of a form, Max 33MB
	err = request.ParseMultipartForm(32 << 20) // 32<<20 = 32 * 2^20 = 33,554,432 bits = 32.768 MB
	if err != nil && !strings.Contains(err.Error(), "EOF") {
		commons.Logcatch(writer, http.StatusBadRequest, err)
		return
	}

	var notification models.AdminNotification
	var notificationImages []models.AdminNotificationImage
	decoderFormFields := schema.NewDecoder()
	decoderFormFields.SetAliasTag("json")
	decoderFormFields.Decode(&notification, request.Form)

	form := request.MultipartForm
	files := form.File["images"]
	for iterator, file := range files {

		if !commons.FileIsImage(file.Filename) {
			commons.Logcatch(writer, http.StatusBadRequest, errors.New("file type not supported, only .jpg, .jpeg and .png are allowed"))
			return
		}

		fileBuffer, err := file.Open()
		if err != nil {
			commons.Logcatch(writer, http.StatusBadRequest, err)
		}
		defer fileBuffer.Close()

		newFileName := fmt.Sprintf("%d.webp", iterator+1)
		localPathStorage := fmt.Sprintf("./storage/public/notifications/%d/images/", notification.Id)

		err = commons.SavePictureAsWebp(fileBuffer, localPathStorage, newFileName)
		if err != nil {
			commons.Logcatch(writer, http.StatusBadRequest, errors.New("cant create file "+err.Error()))
			return
		}

		pathStorage := fmt.Sprintf("%s://%s:%s/notifications/%d/images/%s", os.Getenv("HTTP"), os.Getenv("SERVERHOST"), os.Getenv("SERVERPORT"), notification.Id, newFileName)

		image := models.AdminNotificationImage{
			Image: pathStorage,
		}

		/*err = crud.NewNotificationImage(&image)
		if err != nil {
			commons.Logcatch(writer, http.StatusBadRequest, err)
			return
		}

		notification.Images = append(notification.Images, image)*/

		notificationImages = append(notificationImages, image)
	}

	err = crud.NewNotification(&notification)
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, err)
		return
	}

	for _, image := range notificationImages {
		image.NotificationID = notification.Id
		err = crud.NewNotificationImage(&image)
		if err != nil {
			commons.Logcatch(writer, http.StatusBadRequest, err)
			return
		}
		notification.Images = append(notification.Images, image)
	}

	crud.NewServerActionLog(models.ServerLogs{
		UserID:      adminBufferId,
		Root:        &root,
		Transaction: "sendNotification",
	})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(notification)
}

//

func createNewFood(writer http.ResponseWriter, request *http.Request) {}

func createNewDrink(writer http.ResponseWriter, request *http.Request) {}

func createNewMenuIntem(writer http.ResponseWriter, request *http.Request) {}

func addMenuItemsToBranchStock(writer http.ResponseWriter, request *http.Request) {}

//

func seeTurnsAtBranch(writer http.ResponseWriter, request *http.Request) {}

func seeTurnAtBranch(writer http.ResponseWriter, request *http.Request) {}

func changeActiveTurn(writer http.ResponseWriter, request *http.Request) {}

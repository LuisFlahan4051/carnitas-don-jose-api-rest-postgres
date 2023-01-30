package supervisorActions

import (
	"encoding/json"
	"fmt"

	"errors"
	"os"

	"net/http"

	commons "github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/commonFunctions"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/crud"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/models"

	"github.com/gorilla/schema"
)

func seeBranch(writer http.ResponseWriter, request *http.Request) {}

func registNewUser(writer http.ResponseWriter, request *http.Request) {
	bufferId, _, err := commons.Authentication(request, commons.SUPERVISOR)
	root := false
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}

	//Need Content-Type: multipart/form-data sending by inputs of a form, Max 33MB
	err = request.ParseMultipartForm(32 << 20)
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, err)
		return
	}

	// Getting the data from the form and decode into the user struct
	var user models.User

	file, handle, err := request.FormFile("profile_picture")
	existFile := true
	if err != nil {
		if err.Error() != "http: no such file" {
			commons.Logcatch(writer, http.StatusBadRequest, err)
			return
		}
		existFile = false
	}

	if existFile {
		defer file.Close()

		if !commons.FileIsImage(handle.Filename) {
			commons.Logcatch(writer, http.StatusBadRequest, errors.New("file type not supported, only .jpg, .jpeg and .png are allowed"))
			return
		}

		_ = schema.NewDecoder().Decode(&user, request.Form)

		err = crud.NewUser(&user)
		if err != nil {
			commons.Logcatch(writer, http.StatusBadRequest, err)
			return
		}

		user.Id = commons.UserExists(user.Id)
		if user.Id == 0 {
			commons.Logcatch(writer, http.StatusBadRequest, errors.New("user do not created"))
			return
		}

		// Save the file in the server
		newFileName := "profile_picture.webp"
		localPathStorage := fmt.Sprintf("./storage/public/users/%d/profile/picture/", user.Id)

		err = commons.SavePictureAsWebp(file, localPathStorage, newFileName)

		if err != nil {
			commons.Logcatch(writer, http.StatusBadRequest, err)
			return
		}

		pathStorage := fmt.Sprintf("%s://%s:%s/users/%d/profile/picture/%s", os.Getenv("HTTP"), os.Getenv("SERVERHOST"), os.Getenv("SERVERPORT"), user.Id, newFileName)

		user.Photo = &pathStorage
		err = crud.UpdateUser(&user, false)
		if err != nil {
			commons.Logcatch(writer, http.StatusBadRequest, err)
			return
		}

		commons.SaveServerActionLog(models.ServerLogs{
			UserID:      bufferId,
			Root:        &root,
			Transaction: "registNewUser",
		})
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusCreated)
		json.NewEncoder(writer).Encode(user)
		return
	}

	_ = schema.NewDecoder().Decode(&user, request.Form)

	err = crud.NewUser(&user)
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, err)
		return
	}

	commons.SaveServerActionLog(models.ServerLogs{
		UserID:      bufferId,
		Root:        &root,
		Transaction: "registNewUser",
	})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(user)
}

func seeUsersAtBranch(writer http.ResponseWriter, request *http.Request) {}

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

func sendNotification(writer http.ResponseWriter, request *http.Request) {}

//

func createNewFood(writer http.ResponseWriter, request *http.Request) {}

func createNewDrink(writer http.ResponseWriter, request *http.Request) {}

func createNewMenuIntem(writer http.ResponseWriter, request *http.Request) {}

func addMenuItemsToBranchStock(writer http.ResponseWriter, request *http.Request) {}

//

func seeTurnsAtBranch(writer http.ResponseWriter, request *http.Request) {}

func seeTurnAtBranch(writer http.ResponseWriter, request *http.Request) {}

func changeActiveTurn(writer http.ResponseWriter, request *http.Request) {}

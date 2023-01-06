package supervisorActions

import (
	"encoding/json"
	"errors"
	"os"

	"strconv"

	"net/http"

	commons "github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/commonFunctions"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/crud"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/models"

	"github.com/gorilla/schema"
)

func seeBranch(writer http.ResponseWriter, request *http.Request) {}

func registNewUser(writer http.ResponseWriter, request *http.Request) {
	//Need Content-Type: multipart/form-data sending by inputs of a form, Max 33MB
	err := request.ParseMultipartForm(32 << 20)
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, err)
		return
	}

	// Getting the data from the form and decode into the user struct
	var user models.User

	file, handle, err := request.FormFile("profile_picture")

	var nonFile bool
	if err != nil {
		if err.Error() != "http: no such file" {
			commons.Logcatch(writer, http.StatusBadRequest, err)
			return
		}
		nonFile = true
	}

	if !nonFile {
		defer file.Close()

		if !commons.FileIsImage(handle.Filename) {
			commons.Logcatch(writer, http.StatusBadRequest, errors.New("file type not supported, only .jpg, .jpeg and .png are allowed"))
			return
		}

		err = schema.NewDecoder().Decode(&user, request.Form)
		if err != nil {
			commons.Logcatch(writer, http.StatusBadRequest, err)
			return
		}

		err = crud.NewUser(&user)
		if err != nil {
			commons.Logcatch(writer, http.StatusBadRequest, err)
			return
		}
		// Save the file in the server
		newFileName := "profile_picture.webp"
		pathStorage := "./storage/users/" + strconv.Itoa(int(user.Id)) + "/profile/picture/"

		err = commons.SavePictureAsWebp(file, pathStorage, newFileName)
		if err != nil {
			commons.Logcatch(writer, http.StatusBadRequest, err)
			return
		}

		storage := os.Getenv("HTTP") + "://" + os.Getenv("HOST") + ":" + os.Getenv("PORT") + "/user/" + strconv.Itoa(int(user.Id)) + "/profile/picture"
		user.Photo = &storage

		//TODO: CORREGIR ESTO
		err = crud.UpdateUser(&models.User{
			ID: models.ID{
				Id: user.Id,
			},
			Photo: &storage,
		}, false)
		if err != nil {
			commons.Logcatch(writer, http.StatusBadRequest, err)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusCreated)
		json.NewEncoder(writer).Encode(user)
		return
	}

	/*err = schema.NewDecoder().Decode(&user, request.Form)

	if err.Error() != "schema: invalid path \"profile_picture\"" {
		commons.Logcatch(writer, http.StatusBadRequest, err)
		return
	}*/

	err = crud.NewUser(&user)
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, err)
		return
	}

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

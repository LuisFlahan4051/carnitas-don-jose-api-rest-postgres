package commonActions

import (
	"encoding/json"
	"net/http"

	commons "github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/commonFunctions"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/crud"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/models"
)

func seeMyProfile(writer http.ResponseWriter, request *http.Request) {
	var loginForm models.LoginForm
	json.NewDecoder(request.Body).Decode(&loginForm)
	userID, accessLevel, err := commons.ValidateUser(loginForm.Username, loginForm.Password)
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, err)
		return
	}

	root := accessLevel == commons.ROOT && request.URL.Query().Get("root") == "true"
	newUser, err := crud.GetUser(userID, root)

	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, err)
		return
	}

	crud.NewServerActionLog(models.ServerLogs{
		UserID:      newUser.Id,
		Root:        &root,
		Transaction: "seeMyProfile",
	})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(newUser)

}

func changeMyMainCredentials(writer http.ResponseWriter, request *http.Request) {}

func changeMyProfile(writer http.ResponseWriter, request *http.Request) {}

func changeMyProfilePicture(writer http.ResponseWriter, request *http.Request) {}

func validateUser(writer http.ResponseWriter, request *http.Request) {

	var loginForm models.LoginForm
	json.NewDecoder(request.Body).Decode(&loginForm)

	userID, _, err := commons.ValidateUser(loginForm.Username, loginForm.Password)

	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, err)
		return
	}

	root := false
	crud.NewServerActionLog(models.ServerLogs{
		UserID:      userID,
		Root:        &root,
		Transaction: "login",
	})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
}

//-------------------------> Fast Calculator <-------------------------//

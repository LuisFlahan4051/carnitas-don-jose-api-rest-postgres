package commonActions

import (
	"encoding/json"
	"net/http"

	commons "github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/commonFunctions"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/crud"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/models"
)

func seeMyProfile(writer http.ResponseWriter, request *http.Request) {}

func changeMyMainCredentials(writer http.ResponseWriter, request *http.Request) {}

func changeMyProfile(writer http.ResponseWriter, request *http.Request) {}

func changeMyProfilePicture(writer http.ResponseWriter, request *http.Request) {}

func validateUser(writer http.ResponseWriter, request *http.Request) {

	var loginForm models.LoginForm
	json.NewDecoder(request.Body).Decode(&loginForm)

	_, _, err := commons.ValidateUser(loginForm.Username, loginForm.Password)

	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, err)
		return
	}

	root := false
	crud.NewServerActionLog(models.ServerLogs{
		UserID:      0,
		Root:        &root,
		Transaction: "login",
	})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
}

//-------------------------> Fast Calculator <-------------------------//

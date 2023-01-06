package routes

import (
	"html/template"
	"net/http"

	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/routes/adminActions"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/routes/cashierActions"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/routes/commonActions"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/routes/cookActions"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/routes/rootActions"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/routes/supervisorActions"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/routes/waiterActions"
	"github.com/gorilla/mux"
)

var thisURLs []string

func SetMainHandleRoutes(router *mux.Router, URLs *[]string) {
	*URLs = append(*URLs, "INDEX:")
	route := "/"
	router.HandleFunc(route, homeHandler)
	*URLs = append(*URLs, route)

	*URLs = append(*URLs, "ROOT:")
	rootActions.SetRootHandleActions(router, URLs)
	*URLs = append(*URLs, "ADMIN:")
	adminActions.SetAdminHandleActions(router, URLs)
	*URLs = append(*URLs, "SUPERVISOR:")
	supervisorActions.SetSupervisorHandleActions(router, URLs)
	*URLs = append(*URLs, "CASHIER:")
	cashierActions.SetCashierHandleActions(router, URLs)
	*URLs = append(*URLs, "WAITER:")
	waiterActions.SetWaiterHandleActions(router, URLs)
	*URLs = append(*URLs, "COOK:")
	cookActions.SetCookHandleActions(router, URLs)
	*URLs = append(*URLs, "COMMON:")
	commonActions.SetCommonHandleActions(router, URLs)

	thisURLs = *URLs
}

func homeHandler(writer http.ResponseWriter, request *http.Request) {
	var template *template.Template

	template, err := template.ParseFiles("./routes/display-routes.html")
	if err != nil {
		writer.Write([]byte(err.Error()))
		return
	}

	template.Execute(writer, thisURLs)
}

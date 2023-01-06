package cashierActions

import "net/http"

func seeActiveTurn(writer http.ResponseWriter, request *http.Request) {}

//

func openNewTurn(writer http.ResponseWriter, request *http.Request) {}

func closeTurn(writer http.ResponseWriter, request *http.Request) {}

//

func seeOrdersAtTurn(writer http.ResponseWriter, request *http.Request) {}

func closeOrder(writer http.ResponseWriter, request *http.Request) {}

//

func setUserRoleAtTurn(writer http.ResponseWriter, request *http.Request) {}

func changeUserRoleAtTurn(writer http.ResponseWriter, request *http.Request) {}

//

func createLossReport(writer http.ResponseWriter, request *http.Request) {}

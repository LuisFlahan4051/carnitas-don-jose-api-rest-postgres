package waiterActions

import "net/http"

func seeMenuItemsOfBranch(writer http.ResponseWriter, request *http.Request) {}

//

func openOrder(writer http.ResponseWriter, request *http.Request) {}

func cancelOrder(writer http.ResponseWriter, request *http.Request) {}

func seeOrdersOfWaiterAtTurn(writer http.ResponseWriter, request *http.Request) {}

//

func addProductsToOrder(writer http.ResponseWriter, request *http.Request) {}

func addCommentsToOrder(writer http.ResponseWriter, request *http.Request) {}

func addIncomesToOrder(writer http.ResponseWriter, request *http.Request) {}

func addDiscountToOrder(writer http.ResponseWriter, request *http.Request) {}

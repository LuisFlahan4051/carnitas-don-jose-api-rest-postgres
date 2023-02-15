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
	adminBufferId, maxAccessLevel, err := commons.Authentication(request, commons.ADMIN)
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}
	root := maxAccessLevel == commons.ROOT && request.URL.Query().Get("root") == "true"

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
			if strings.Contains(err.Error(), "EOF") {
				today := true
				pagination.Today = &today
			} else {
				commons.Logcatch(writer, http.StatusBadRequest, err)
				return
			}
		}
	}

	notifications, err := crud.GetNotifications(pagination, solved, root, nil)
	if err != nil {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}

	for i := 0; i < len(notifications); i++ {
		notifications[i].Images, err = crud.GetNotificationImages(notifications[i].Id, root)
		if err != nil && !strings.Contains(err.Error(), "not found") {
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
		commons.Logcatch(writer, http.StatusBadRequest, errors.New("invalid notification/{id} value"))
		return
	}

	adminBufferId, maxAccessLevel, err := commons.Authentication(request, commons.ADMIN)
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}
	root := maxAccessLevel == commons.ROOT && request.URL.Query().Get("root") == "true"

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
	json.NewEncoder(writer).Encode("done")
}

func dropNotifications(writer http.ResponseWriter, request *http.Request) {
	adminBufferId, maxAccessLevel, err := commons.Authentication(request, commons.ADMIN)
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}
	root := maxAccessLevel == commons.ROOT && request.URL.Query().Get("root") == "true"

	err = crud.DeleteNotifications(root)
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

func seeBranches(writer http.ResponseWriter, request *http.Request) {
	adminBufferId, maxAccessLevel, err := commons.Authentication(request, commons.ADMIN)
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}
	root := maxAccessLevel == commons.ROOT && request.URL.Query().Get("root") == "true"

	branches, err := crud.GetBranches(root)
	if err != nil {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}

	crud.NewServerActionLog(models.ServerLogs{
		UserID:      adminBufferId,
		Root:        &root,
		Transaction: "seeBranches",
	})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(branches)
}

func createNewBranch(writer http.ResponseWriter, request *http.Request) {
	adminBufferId, maxAccessLevel, err := commons.Authentication(request, commons.ADMIN)
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}
	root := maxAccessLevel == commons.ROOT && request.URL.Query().Get("root") == "true"

	var branch models.Branch
	err = json.NewDecoder(request.Body).Decode(&branch)
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, err)
		return
	}

	err = crud.NewBranch(&branch)
	if err != nil {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}

	var mainSafebox models.Safebox
	err = crud.NewSafebox(&mainSafebox)
	if err != nil {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}

	relationBranchSafebox := models.BranchSafebox{
		Name:      "main",
		Content:   new(float64),
		BranchID:  branch.Id,
		SafeboxID: mainSafebox.Id,
	}
	err = crud.NewBranchSafebox(&relationBranchSafebox)
	if err != nil {
		crud.DeleteBranch(branch.Id, true)
		crud.DeleteSafebox(mainSafebox.Id, true)
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}

	branch.BranchSafeboxes = append(branch.BranchSafeboxes, relationBranchSafebox)

	crud.NewServerActionLog(models.ServerLogs{
		UserID:      adminBufferId,
		Root:        &root,
		Transaction: "createNewBranch",
	})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(branch)
}

func changeInfoBranch(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(mux.Vars(request)["notification_id"])
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, errors.New("invalid branch/{id} value"))
		return
	}

	adminBufferId, maxAccessLevel, err := commons.Authentication(request, commons.ADMIN)
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}
	root := maxAccessLevel == commons.ROOT && request.URL.Query().Get("root") == "true"

	var branch models.Branch
	err = json.NewDecoder(request.Body).Decode(&branch)
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, err)
		return
	}
	branch.Id = uint(id)

	err = crud.UpdateBranch(&branch, root)
	if err != nil {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}

	crud.NewServerActionLog(models.ServerLogs{
		UserID:      adminBufferId,
		Root:        &root,
		Transaction: "changeInfoBranch",
	})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(branch)
}

func dropBranch(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(mux.Vars(request)["notification_id"])
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, errors.New("invalid branch/{id} value"))
		return
	}

	adminBufferId, maxAccessLevel, err := commons.Authentication(request, commons.ADMIN)
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}
	root := maxAccessLevel == commons.ROOT && request.URL.Query().Get("root") == "true"

	err = crud.DeleteBranch(uint(id), root)
	if err != nil {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}

	crud.NewServerActionLog(models.ServerLogs{
		UserID:      adminBufferId,
		Root:        &root,
		Transaction: "dropBranch",
	})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode("done")
}

//

func seeSupplies(writer http.ResponseWriter, request *http.Request) {
	adminBufferId, maxAccessLevel, err := commons.Authentication(request, commons.ADMIN)
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}
	root := maxAccessLevel == commons.ROOT && request.URL.Query().Get("root") == "true"

	supplies, err := crud.GetSupplies(root)
	if err != nil {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}

	crud.NewServerActionLog(models.ServerLogs{
		UserID:      adminBufferId,
		Root:        &root,
		Transaction: "seeSupplies",
	})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(supplies)
}

func seeSupply(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(mux.Vars(request)["supply_id"])
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, errors.New("invalid supply/{id} value"))
		return
	}

	adminBufferId, accessLevel, err := commons.Authentication(request, commons.ADMIN)
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}

	root := accessLevel == commons.ROOT && request.URL.Query().Get("root") == "true"

	supply, err := crud.GetSupply(uint(id), root)
	if err != nil {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}

	crud.NewServerActionLog(models.ServerLogs{
		UserID:      adminBufferId,
		Root:        &root,
		Transaction: "seeBranch",
	})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(supply)
}

func createNewSupply(writer http.ResponseWriter, request *http.Request) {
	adminBufferId, accessLevel, err := commons.Authentication(request, commons.ADMIN)
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}
	root := accessLevel == commons.ROOT && request.URL.Query().Get("root") == "true"

	var supply models.Supply
	err = crud.NewSupply(&supply)
	if err != nil {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}

	crud.NewServerActionLog(models.ServerLogs{
		UserID:      adminBufferId,
		Root:        &root,
		Transaction: "newSupply",
	})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(supply)
}

func changeSupply(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(mux.Vars(request)["supply_id"])
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, errors.New("invalid supply/{id} value"))
		return
	}

	adminBufferId, accessLevel, err := commons.Authentication(request, commons.ADMIN)
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}
	root := accessLevel == commons.ROOT && request.URL.Query().Get("root") == "true"

	var supply models.Supply
	err = json.NewDecoder(request.Body).Decode(&supply)
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, err)
		return
	}
	supply.Id = uint(id)

	crud.NewSupply(&supply)

	crud.NewServerActionLog(models.ServerLogs{
		UserID:      adminBufferId,
		Root:        &root,
		Transaction: "changeSupply",
	})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(supply)
}

func dropSupply(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(mux.Vars(request)["supply_id"])
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, errors.New("invalid supply/{id} value"))
		return
	}

	adminBufferId, accessLevel, err := commons.Authentication(request, commons.ADMIN)
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}
	root := accessLevel == commons.ROOT && request.URL.Query().Get("root") == "true"

	err = crud.DeleteSupply(uint(id), root)
	if err != nil {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}

	crud.NewServerActionLog(models.ServerLogs{
		UserID:      adminBufferId,
		Root:        &root,
		Transaction: "dropSupply",
	})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode("done")
}

//

func addSupplyToBranchStock(writer http.ResponseWriter, request *http.Request) {
	idBranch, err := strconv.Atoi(mux.Vars(request)["branch_id"])
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, errors.New("invalid branch/{id} value"))
		return
	}
	idSupply, err := strconv.Atoi(mux.Vars(request)["supply_id"])
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, errors.New("invalid supply/{id} value"))
		return
	}

	adminBufferId, accessLevel, err := commons.Authentication(request, commons.ADMIN)
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}
	root := accessLevel == commons.ROOT && request.URL.Query().Get("root") == "true"

	var branchSupplyStock models.BranchSupplyStock
	err = json.NewDecoder(request.Body).Decode(&branchSupplyStock)
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, err)
		return
	}

	branchSupplyStock.BranchID = uint(idBranch)
	if commons.RegistExists("branches", branchSupplyStock.BranchID) == 0 {
		commons.Logcatch(writer, http.StatusBadRequest, errors.New("branch not found"))
		return
	}

	branchSupplyStock.SupplyID = uint(idSupply)
	if commons.RegistExists("supplies", branchSupplyStock.SupplyID) == 0 {
		commons.Logcatch(writer, http.StatusBadRequest, errors.New("supply not found"))
		return
	}

	relationIDs := make(map[string]uint)
	relationIDs["branch_id"] = branchSupplyStock.BranchID
	relationIDs["supply_id"] = branchSupplyStock.SupplyID
	rows, _ := crud.GetBranchSuppliesStock(root, &relationIDs)
	if rows != nil {
		commons.Logcatch(writer, http.StatusBadRequest, errors.New("supply already in stock"))
		return
	}

	err = crud.NewBranchSupplyStock(&branchSupplyStock)
	if err != nil {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}

	crud.NewServerActionLog(models.ServerLogs{
		UserID:      adminBufferId,
		Root:        &root,
		Transaction: "addSuppliesToBranchStock",
	})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(branchSupplyStock)
}

func removeSupplyFromBranchStock(writer http.ResponseWriter, request *http.Request) {
	idBranch, err := strconv.Atoi(mux.Vars(request)["branch_id"])
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, errors.New("invalid branch/{id} value"))
		return
	}
	idSupply, err := strconv.Atoi(mux.Vars(request)["supply_id"])
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, errors.New("invalid supply/{id} value"))
		return
	}

	adminBufferId, accessLevel, err := commons.Authentication(request, commons.ADMIN)
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}
	root := accessLevel == commons.ROOT && request.URL.Query().Get("root") == "true"

	var branchSupplyStock models.BranchSupplyStock
	relationIDs := make(map[string]uint)
	relationIDs["branch_id"] = uint(idBranch)
	relationIDs["supply_id"] = uint(idSupply)
	rows, _ := crud.GetBranchSuppliesStock(root, &relationIDs)
	if rows == nil {
		commons.Logcatch(writer, http.StatusBadRequest, errors.New("supply is not in stock"))
		return
	}
	for _, row := range rows {
		//expect only one
		branchSupplyStock = row
	}

	err = crud.DeleteBranchSupplyStock(branchSupplyStock.Id, root)
	if err != nil {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}

	crud.NewServerActionLog(models.ServerLogs{
		UserID:      adminBufferId,
		Root:        &root,
		Transaction: "removeSuppliesFromBranchStock",
	})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode("done")
}

//

func seeArticles(writer http.ResponseWriter, request *http.Request) {
	adminBufferId, maxAccessLevel, err := commons.Authentication(request, commons.ADMIN)
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}
	root := maxAccessLevel == commons.ROOT && request.URL.Query().Get("root") == "true"

	articles, err := crud.GetArticles(root)
	if err != nil {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}

	crud.NewServerActionLog(models.ServerLogs{
		UserID:      adminBufferId,
		Root:        &root,
		Transaction: "seeArticles",
	})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(articles)
}

func seeArticle(writer http.ResponseWriter, request *http.Request) {
	idArticle, err := strconv.Atoi(mux.Vars(request)["article_id"])
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, errors.New("invalid article/{id} value"))
		return
	}

	adminBufferId, maxAccessLevel, err := commons.Authentication(request, commons.ADMIN)
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}
	root := maxAccessLevel == commons.ROOT && request.URL.Query().Get("root") == "true"

	article, err := crud.GetArticle(uint(idArticle), root)
	if err != nil {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}

	crud.NewServerActionLog(models.ServerLogs{
		UserID:      adminBufferId,
		Root:        &root,
		Transaction: "seeArticle",
	})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(article)
}

func createNewAritcle(writer http.ResponseWriter, request *http.Request) {
	adminBufferId, accessLevel, err := commons.Authentication(request, commons.ADMIN)
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}
	root := accessLevel == commons.ROOT && request.URL.Query().Get("root") == "true"

	var article models.Article
	err = json.NewDecoder(request.Body).Decode(&article)
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, err)
		return
	}

	err = crud.NewArticle(&article)
	if err != nil {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}

	crud.NewServerActionLog(models.ServerLogs{
		UserID:      adminBufferId,
		Root:        &root,
		Transaction: "createNewAritcle",
	})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(article)
}

func changeArticle(writer http.ResponseWriter, request *http.Request) {
	idArticle, err := strconv.Atoi(mux.Vars(request)["article_id"])
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, errors.New("invalid article/{id} value"))
		return
	}

	adminBufferId, accessLevel, err := commons.Authentication(request, commons.ADMIN)
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}
	root := accessLevel == commons.ROOT && request.URL.Query().Get("root") == "true"

	var article models.Article
	err = json.NewDecoder(request.Body).Decode(&article)
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, err)
		return
	}
	article.Id = uint(idArticle)

	err = crud.UpdateArticle(&article, root)
	if err != nil {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}

	crud.NewServerActionLog(models.ServerLogs{
		UserID:      adminBufferId,
		Root:        &root,
		Transaction: "changeArticle",
	})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(article)
}

func dropArticle(writer http.ResponseWriter, request *http.Request) {
	idArticle, err := strconv.Atoi(mux.Vars(request)["article_id"])
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, errors.New("invalid article/{id} value"))
		return
	}

	adminBufferId, accessLevel, err := commons.Authentication(request, commons.ADMIN)
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}
	root := accessLevel == commons.ROOT && request.URL.Query().Get("root") == "true"

	err = crud.DeleteArticle(uint(idArticle), root)
	if err != nil {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}

	crud.NewServerActionLog(models.ServerLogs{
		UserID:      adminBufferId,
		Root:        &root,
		Transaction: "dropArticle",
	})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode("done")
}

//

func addArticleToBranchStock(writer http.ResponseWriter, request *http.Request) {
	idBranch, err := strconv.Atoi(mux.Vars(request)["branch_id"])
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, errors.New("invalid branch/{id} value"))
		return
	}
	idArticle, err := strconv.Atoi(mux.Vars(request)["article_id"])
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, errors.New("invalid article/{id} value"))
		return
	}

	adminBufferId, accessLevel, err := commons.Authentication(request, commons.ADMIN)
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}
	root := accessLevel == commons.ROOT && request.URL.Query().Get("root") == "true"

	var branchArticleStock models.BranchArticleStock
	err = json.NewDecoder(request.Body).Decode(&branchArticleStock)
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, err)
		return
	}
	branchArticleStock.BranchID = uint(idBranch)
	if commons.RegistExists("branches", branchArticleStock.BranchID) == 0 {
		commons.Logcatch(writer, http.StatusBadRequest, errors.New("branch not found"))
		return
	}

	branchArticleStock.ArticleID = uint(idArticle)
	if commons.RegistExists("articles", branchArticleStock.ArticleID) == 0 {
		commons.Logcatch(writer, http.StatusBadRequest, errors.New("article not found"))
		return
	}

	relationIDs := make(map[string]uint)
	relationIDs["branch_id"] = branchArticleStock.BranchID
	relationIDs["article_id"] = branchArticleStock.ArticleID
	rows, _ := crud.GetBranchArticlesStock(root, &relationIDs)
	if rows != nil {
		commons.Logcatch(writer, http.StatusBadRequest, errors.New("article already in stock"))
		return
	}

	err = crud.NewBranchArticleStock(&branchArticleStock)
	if err != nil {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}

	crud.NewServerActionLog(models.ServerLogs{
		UserID:      adminBufferId,
		Root:        &root,
		Transaction: "addArticlesToBranchStock",
	})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(branchArticleStock)
}

func removeArticleFromBranchStock(writer http.ResponseWriter, request *http.Request) {
	idBranch, err := strconv.Atoi(mux.Vars(request)["branch_id"])
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, errors.New("invalid branch/{id} value"))
		return
	}
	idArticle, err := strconv.Atoi(mux.Vars(request)["article_id"])
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, errors.New("invalid article/{id} value"))
		return
	}

	adminBufferId, accessLevel, err := commons.Authentication(request, commons.ADMIN)
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}
	root := accessLevel == commons.ROOT && request.URL.Query().Get("root") == "true"

	var branchArticleStock models.BranchArticleStock
	relationIDs := make(map[string]uint)
	relationIDs["branch_id"] = uint(idBranch)
	relationIDs["article_id"] = uint(idArticle)
	rows, _ := crud.GetBranchArticlesStock(root, &relationIDs)
	if rows == nil {
		commons.Logcatch(writer, http.StatusBadRequest, errors.New("article is not in stock"))
		return
	}
	for _, row := range rows {
		//expect only one
		branchArticleStock = row
	}

	err = crud.DeleteBranchArticleStock(branchArticleStock.Id, root)
	if err != nil {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}

	crud.NewServerActionLog(models.ServerLogs{
		UserID:      adminBufferId,
		Root:        &root,
		Transaction: "addArticlesToBranchStock",
	})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(branchArticleStock)
}

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
	adminBufferId, maxAccessLevel, err := commons.Authentication(request, commons.ADMIN)
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}

	root := maxAccessLevel == commons.ROOT && request.URL.Query().Get("root") == "true"
	users, err := crud.GetUsers(root, nil)
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

	adminBufferId, maxAccessLevel, err := commons.Authentication(request, commons.ADMIN)
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}

	root := maxAccessLevel == commons.ROOT && request.URL.Query().Get("root") == "true"
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

	adminBufferId, maxAccessLevel, err := commons.Authentication(request, commons.ADMIN)
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}

	root := maxAccessLevel == commons.ROOT && request.URL.Query().Get("root") == "true"

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

func verifyUser(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(mux.Vars(request)["user_id"])
	if err != nil {
		commons.Logcatch(writer, http.StatusBadRequest, errors.New("invalid user/{id} value"))
		return
	}

	adminBufferId, maxAccessLevel, err := commons.Authentication(request, commons.ADMIN)
	if err != nil {
		commons.Logcatch(writer, http.StatusUnauthorized, err)
		return
	}
	root := maxAccessLevel == commons.ROOT && request.URL.Query().Get("root") == "true"

	user, err := crud.GetUser(uint(id), root)
	if err != nil {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}

	*user.Verified = true
	crud.UpdateUser(&user, root)
	if err != nil {
		commons.Logcatch(writer, http.StatusInternalServerError, err)
		return
	}

	crud.NewServerActionLog(models.ServerLogs{
		UserID:      adminBufferId,
		Root:        &root,
		Transaction: "verifyUser",
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

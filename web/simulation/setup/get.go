package setup

import (
	"log"
	"net/http"

	"github.com/Palantir/palantir/db"
	"github.com/Palantir/palantir/web/auth/token"
	"github.com/Palantir/palantir/web/pathvars"
	"github.com/Palantir/palantir/web/server"
	"github.com/Palantir/palantir/web/util"
)

type getSetupHandler struct {
	*server.Context
}

func (h *getSetupHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	accountID := token.ExtractAccountID(r)

	setupID, isValid := pathvars.ExtractSetupID(r)
	if !isValid {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	dbSession := h.Db.Copy()
	defer dbSession.Close()

	setup, err := dbSession.Setup().Fetch(db.SetupID{Account: accountID, Setup: setupID})
	switch {
	case setup == nil:
		w.WriteHeader(http.StatusNotFound)
		return
	case err != nil:
		log.Print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_ = util.WriteJSONResponse(w, http.StatusOK, setup)

}

package organisations

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"invitr.io.com/responder"
	"invitr.io.com/services/organisations/orgstore"
)

const (
	ROOT = "/organisations"
	BASE = ROOT + "/organisation"
	ORG  = BASE + "/{organisationId}"
)

type OrgMgr struct {
	Env       string
	Logger    *zap.Logger
	Router    *mux.Router
	Responder responder.Responder
	OrgStore  orgstore.OrgStorer
}

func NewOrgMgr(router *mux.Router, environment string, logger *zap.Logger, responder responder.Responder, orgstore orgstore.OrgStorer) *OrgMgr {
	e := &OrgMgr{
		Env:       environment,
		Logger:    logger,
		Router:    router,
		Responder: responder,
		OrgStore:  orgstore,
	}
	e.Register()
	return e
}

// GetOrganisation returns an organisation given an organisation id
func (mgr *OrgMgr) GetOrganisation() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		orgId := mux.Vars(req)["organisationId"]

		org, err := mgr.OrgStore.Get(orgId)
		if err != nil {
			mgr.Responder.Error(w, 400, err)
			return
		}

		if org == nil {
			mgr.Responder.Error(w, 404, errors.New("organisation not found"))
			return
		}

		mgr.Responder.Respond(w, http.StatusOK, org)
	}
}

// NewOrganisation adds a new organisation to the db
func (mgr *OrgMgr) NewOrganisation() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		var getorg orgstore.Organisation
		json.NewDecoder(req.Body).Decode(&getorg)

		org, err := mgr.OrgStore.Insert(&getorg)
		if err != nil {
			mgr.Responder.Error(w, 400, err)
			return
		}

		mgr.Responder.Respond(w, http.StatusCreated, org)
	}
}

func (mgr *OrgMgr) Register() {
	mgr.Router.HandleFunc(ORG, mgr.GetOrganisation()).Methods("GET")
	mgr.Router.HandleFunc(ORG, mgr.NewOrganisation()).Methods("POST")
}

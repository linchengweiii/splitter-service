package router

import (
	"encoding/json"
	"net/http"

	"github.com/linchengweiii/splitter/pkg/group"
)

type GroupRouter interface {
	GetGroup(w http.ResponseWriter, r *http.Request)
}

type GroupRouterImpl struct {
    groupId string
	groupService *group.Service
}

func (router *GroupRouterImpl) GetGroup(w http.ResponseWriter, r *http.Request) {
	group, err := router.groupService.Read(router.groupId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	g, err := json.Marshal(group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(g)
}

func NewGroupRouter(groupId string, groupService *group.Service) GroupRouter {
	return &GroupRouterImpl{groupId, groupService}
}

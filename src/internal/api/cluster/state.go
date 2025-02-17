package cluster

import (
	"net/http"

	"github.com/defenseunicorns/zarf/src/internal/api/common"
	"github.com/defenseunicorns/zarf/src/internal/k8s"
	"github.com/defenseunicorns/zarf/src/internal/message"
	"github.com/defenseunicorns/zarf/src/types"
)

// Read the zarf state secret from the cluster, if it exists.
func ReadState(w http.ResponseWriter, r *http.Request) {
	message.Debug("state.Read()")

	data, err := k8s.LoadZarfState()
	if err != nil {
		message.ErrorWebf(err, w, "unable to load zarf state")
	}

	if data.Distro == "" {
		common.WriteEmpty(w)
	} else {
		common.WriteJSONResponse(w, data, http.StatusOK)
	}
}

// Update the zarf state secret in the cluster.
func UpdateState(w http.ResponseWriter, r *http.Request) {
	message.Debug("state.Update()")

	var data types.ZarfState

	if err := k8s.SaveZarfState(data); err != nil {
		message.ErrorWebf(err, w, "unable to update zarf state")
	} else {
		common.WriteJSONResponse(w, data, http.StatusCreated)
	}
}

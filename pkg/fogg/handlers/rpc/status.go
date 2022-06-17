package rpc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/Validat0rs/fogg/pkg/fogg/handlers/rpc/types"
)

func (rp *Rpc) Status(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get(fmt.Sprintf("%s/%s", os.Getenv("RPC_HOST"), "status"))
	if err != nil {
		rp.logger.Error().Msgf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		rp.logger.Error().Msgf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var status types.Status
	if err := json.Unmarshal(body, &status); err != nil {
		rp.logger.Error().Msgf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	status.Result.NodeInfo.Moniker = r.Host

	js, err := json.Marshal(status)
	if err != nil {
		rp.logger.Error().Msgf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}

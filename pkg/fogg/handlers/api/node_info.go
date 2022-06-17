package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/Validat0rs/fogg/pkg/fogg/handlers/api/types"
)

func (a *Api) NodeInfo(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get(fmt.Sprintf("%s/%s", os.Getenv("API_HOST"), "node_info"))
	if err != nil {
		a.logger.Error().Msgf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		a.logger.Error().Msgf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var result types.Result
	if err := json.Unmarshal(body, &result); err != nil {
		a.logger.Error().Msgf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	result.NodeInfo.Moniker = r.Host

	js, err := json.Marshal(result)
	if err != nil {
		a.logger.Error().Msgf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}

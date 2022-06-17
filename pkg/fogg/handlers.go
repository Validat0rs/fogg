package fogg

import (
	"net/http"

	"github.com/Validat0rs/fogg/pkg/fogg/handlers/api"
	"github.com/Validat0rs/fogg/pkg/fogg/handlers/rpc"

	"github.com/urfave/negroni"
)

func (f *Fogg) SetHandlers() {
	f.apiHandler()
	f.rpcHandler()
}

func (f *Fogg) apiHandler() {
	_api := api.NewApi(f.Monitoring.Logger)
	f.HTTP.Router.Handle("/api/node_info", negroni.New(
		negroni.Wrap(http.HandlerFunc(_api.NodeInfo)),
	))
}

func (f *Fogg) rpcHandler() {
	_rpc := rpc.NewRpc(f.Monitoring.Logger)
	f.HTTP.Router.Handle("/rpc/status", negroni.New(
		negroni.Wrap(http.HandlerFunc(_rpc.Status)),
	))
}

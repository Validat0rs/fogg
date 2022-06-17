package rpc

import (
	"encoding/json"
	. "gopkg.in/check.v1"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Validat0rs/fogg/pkg/fogg/handlers/rpc/types"
)

var (
	_      = Suite(&rpcSuite{})
	status = types.Status{
		Jsonrpc: "",
		ID:      0,
		Result: types.Result{
			NodeInfo: types.NodeInfo{
				ProtocolVersion: types.ProtocolVersion{},
				ID:              "4a08b2f5c0052fbb78c74a0db6bcd8111b10e820",
				ListenAddr:      "127.0.0.1:26657",
				Moniker:         "test-node-1",
			},
		},
	}
)

type rpcSuite struct {
	rpc       *Rpc
	rpcServer *httptest.Server
}

func Test(t *testing.T) { TestingT(t) }

func (s *rpcSuite) SetUpSuite(c *C) {
	js, _ := json.Marshal(status)
	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write(js)
	})
	s.rpcServer = httptest.NewServer(handlerFunc)
}

func (s *rpcSuite) TestStatus(c *C) {
	r := httptest.NewRequest(http.MethodGet, "/status", nil)
	w := httptest.NewRecorder()

	_ = os.Setenv("RPC_HOST", s.rpcServer.URL)

	s.rpc.Status(w, r)
	res := w.Result()
	defer res.Body.Close()

	api, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.Error(err)
	}

	var _status types.Status
	if err = json.Unmarshal(api, &_status); err != nil {
		c.Error(err)
	}

	if _status.Result.NodeInfo.ID != "" {
		c.Error("ID is still set")
	}

	if _status.Result.NodeInfo.ListenAddr != "" {
		c.Error("ListenAddr is still set")
	}
}

func (s *rpcSuite) TearDownSuite(c *C) {
	s.rpcServer.Close()
}

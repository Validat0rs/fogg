package api

import (
	"encoding/json"
	. "gopkg.in/check.v1"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Validat0rs/fogg/pkg/fogg/handlers/api/types"
)

var (
	_        = Suite(&apiSuite{})
	nodeInfo = types.Result{
		NodeInfo: types.NodeInfo{
			ID:         "4a08b2f5c0052fbb78c74a0db6bcd8111b10e820",
			ListenAddr: "127.0.0.1:26657",
			Moniker:    "test-node-1",
		},
		ApplicationVersion: types.ApplicationVersion{},
	}
)

type apiSuite struct {
	api       *Api
	apiServer *httptest.Server
}

func Test(t *testing.T) { TestingT(t) }

func (s *apiSuite) SetUpSuite(c *C) {
	js, _ := json.Marshal(nodeInfo)
	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write(js)
	})
	s.apiServer = httptest.NewServer(handlerFunc)
}

func (s *apiSuite) TestNodeInfo(c *C) {
	r := httptest.NewRequest(http.MethodGet, "/node_info", nil)
	w := httptest.NewRecorder()

	_ = os.Setenv("API_HOST", s.apiServer.URL)

	s.api.NodeInfo(w, r)
	res := w.Result()
	defer res.Body.Close()

	api, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.Error(err)
	}

	var result types.Result
	if err = json.Unmarshal(api, &result); err != nil {
		c.Error(err)
	}

	if result.NodeInfo.ID != "" {
		c.Error("ID is still set")
	}

	if result.NodeInfo.ListenAddr != "" {
		c.Error("ListenAddr is still set")
	}
}

func (s *apiSuite) TearDownSuite(c *C) {
	s.apiServer.Close()
}

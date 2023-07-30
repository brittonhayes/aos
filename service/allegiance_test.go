package service_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brittonhayes/aos/api"
	"github.com/brittonhayes/aos/internal/mocks"
	"github.com/brittonhayes/aos/server"
	"github.com/brittonhayes/aos/service"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type mock struct {
	repo   *mocks.MockRepository
	server *server.Server
	ctrl   *gomock.Controller
}

// executeRequest, creates a new ResponseRecorder
// then executes the request by calling ServeHTTP in the router
// after which the handler writes the response to the response recorder
// which we can then inspect.
func executeRequest(req *http.Request, s *server.Server) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)

	return rr
}

func setupMockServer(t *testing.T) *mock {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockRepository(ctrl)
	svc := service.New(repo)
	server := server.New(context.Background(), &server.Config{
		Repository:  repo,
		Service:     svc,
		ServiceName: "test",
		Environment: "test",
		Port:        8080,
		Tracing:     false,
	})

	return &mock{
		repo:   repo,
		server: server,
		ctrl:   ctrl,
	}
}

func TestGetAllegiances(t *testing.T) {
	req, err := http.NewRequest("GET", "/allegiances", nil)
	if err != nil {
		t.Fatal(err)
	}
	got := []api.Allegiance{{Name: "test"}}
	m := setupMockServer(t)
	defer m.ctrl.Finish()

	m.repo.EXPECT().
		GetAllegiances(gomock.Any(), gomock.Any()).
		Return(got, nil).
		Times(1)

	resp := executeRequest(req, m.server)
	if !assert.Equal(t, http.StatusOK, resp.Code) {
		t.Log(resp.Body.String())
	}
}

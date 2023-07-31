package service_test

import (
	"net/http"
	"testing"

	"github.com/brittonhayes/aos/api"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetGrandStrategies(t *testing.T) {
	req, err := http.NewRequest("GET", "/grand-strategies", nil)
	if err != nil {
		t.Fatal(err)
	}
	got := []api.GrandStrategy{{Name: "test"}}

	m := setupMockServer(t)
	defer m.ctrl.Finish()

	m.repo.EXPECT().
		GetGrandStrategies(gomock.Any()).
		Return(got, nil).
		Times(1)

	resp := executeRequest(req, m.server)
	if !assert.Equal(t, http.StatusOK, resp.Code) {
		t.Log(resp.Body.String())
	}
}

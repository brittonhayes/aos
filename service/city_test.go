package service_test

import (
	"net/http"
	"testing"

	"github.com/brittonhayes/aos/api"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetCities(t *testing.T) {
	req, err := http.NewRequest("GET", "/cities", nil)
	if err != nil {
		t.Fatal(err)
	}

	name := "test"
	got := []api.City{{Name: &name}}

	m := setupMockServer(t)
	defer m.ctrl.Finish()

	m.repo.EXPECT().
		GetCities(gomock.Any(), gomock.Any()).
		Return(got, nil).
		Times(1)

	resp := executeRequest(req, m.server)
	if !assert.Equal(t, http.StatusOK, resp.Code) {
		t.Log(resp.Body.String())
	}
}

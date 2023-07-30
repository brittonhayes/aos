package service_test

import (
	"net/http"
	"testing"

	"github.com/brittonhayes/aos/api"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetWarscrolls(t *testing.T) {
	req, err := http.NewRequest("GET", "/warscrolls", nil)
	if err != nil {
		t.Fatal(err)
	}

	got := []api.Unit{{Name: "test"}}

	m := setupMockServer(t)
	defer m.ctrl.Finish()

	m.repo.EXPECT().
		GetWarscrolls(gomock.Any(), gomock.Any()).
		Return(got, nil).
		Times(1)

	resp := executeRequest(req, m.server)
	if !assert.Equal(t, http.StatusOK, resp.Code) {
		t.Log(resp.Body.String())
	}
}

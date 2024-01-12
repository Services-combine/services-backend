package handler_test

import (
	"github.com/b0shka/services/internal/config"
	"github.com/b0shka/services/internal/handler"
	"github.com/b0shka/services/internal/service"
	"github.com/b0shka/services/pkg/auth"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewHandler(t *testing.T) {
	h := handler.NewHandler(&service.Services{}, &auth.PasetoManager{}, config.FoldersConfig{})

	require.IsType(t, &handler.Handler{}, h)
}

func TestNewHandler_InitRoutes(t *testing.T) {
	h := handler.NewHandler(&service.Services{}, &auth.PasetoManager{}, config.FoldersConfig{})
	router := h.InitRoutes(&config.Config{})

	ts := httptest.NewServer(router)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/ping")
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	require.Equal(t, http.StatusOK, res.StatusCode)
}

package apiserver

import (
	"github.com/stretchr/testify/assert"
	"github.com/trad3r/rest-api-go/internal/app/store/teststore"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_HandleUserCreate(t *testing.T) {
	s := NewServer(teststore.New())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/users", nil)
	s.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusOK)
}

package handler

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

func TestHandlerHandleGet(t *testing.T) {

}

func TestHandlerHandlePost(t *testing.T) {
	t.Run("Empty JSON", func(t *testing.T) {
		r, _ := http.Post("localhost:80/link/add", "application/json", strings.NewReader(""))
		defer r.Body.Close()
		assert.Equal(t, http.StatusBadRequest, r.StatusCode)
	})
}

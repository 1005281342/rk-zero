// Copyright (c) 2021 rookie-ninja
//
// Use of this source code is governed by an Apache-style
// license that can be found in the LICENSE file.

package rkzerometa

import (
	rkentry "github.com/rookie-ninja/rk-entry/entry"
	rkmidmeta "github.com/rookie-ninja/rk-entry/middleware/meta"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var userHandler = func(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func TestInterceptor(t *testing.T) {
	defer assertNotPanic(t)

	beforeCtx := rkmidmeta.NewBeforeCtx()
	mock := rkmidmeta.NewOptionSetMock(beforeCtx)

	inter := Interceptor(rkmidmeta.WithMockOptionSet(mock))
	req, w := newReqAndWriter()

	beforeCtx.Input.Event = rkentry.NoopEventLoggerEntry().GetEventFactory().CreateEventNoop()
	beforeCtx.Output.HeadersToReturn["key"] = "value"

	inter(userHandler)(w, req)

	assert.Equal(t, "value", w.Header().Get("key"))
}

func newReqAndWriter() (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(http.MethodGet, "/ut-path", nil)
	req.Header = http.Header{}
	writer := httptest.NewRecorder()
	return req, writer
}

func assertNotPanic(t *testing.T) {
	if r := recover(); r != nil {
		// Expect panic to be called with non nil error
		assert.True(t, false)
	} else {
		// This should never be called in case of a bug
		assert.True(t, true)
	}
}

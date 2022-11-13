// Copyright(c) 2022 individual contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// <https://www.apache.org/licenses/LICENSE-2.0>
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package harwp_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	assertPkg "github.com/stretchr/testify/assert"
	requirePkg "github.com/stretchr/testify/require"

	"github.com/go-raizu/harwp"
)

func Test_ProxyBase_Write(t *testing.T) {
	assert := assertPkg.New(t)

	body := assertPkg.HTTPBody(func(w http.ResponseWriter, r *http.Request) {
		p := harwp.NewResponseProxy(w)

		n, err := p.Write([]byte("hello world!"))
		if assert.NoError(err) {
			assert.Equal(12, n)
		}
	}, "GET", "", nil)

	assert.Equal("hello world!", body)
}

func Test_ProxyBase_StatusCode(t *testing.T) {
	t.Run("initial-no-code", func(t *testing.T) {
		p := harwp.NewResponseProxy(httptest.NewRecorder())

		// Given that no response has been written yet, then there
		// should be no status code set.
		assertPkg.Equal(t, 0, p.StatusCode())
	})

	t.Run("implicit-write", func(t *testing.T) {
		p := harwp.NewResponseProxy(httptest.NewRecorder())

		// Given that write was called directly, ...
		_, err := p.Write([]byte("foo"))
		requirePkg.NoError(t, err)

		// ... then status code shall be implicitly set to 200 ok.
		assertPkg.Equal(t, 200, p.StatusCode())
	})

	t.Run("explicit-status", func(t *testing.T) {
		p := harwp.NewResponseProxy(httptest.NewRecorder())

		// Given that the status code was set explicitly to
		// a given value, ...
		p.WriteHeader(http.StatusBadRequest)

		// ... then the StatusCode return value shall reflect
		// the previously given value.
		assertPkg.Equal(t, 400, p.StatusCode())
	})
}

func Test_ProxyBase_HeadersWritten(t *testing.T) {
	t.Run("initial-no-code", func(t *testing.T) {
		p := harwp.NewResponseProxy(httptest.NewRecorder())

		// Given that no response has been written yet, then there
		// shall be no headers written yet either.
		assertPkg.False(t, p.HeadersWritten())
	})

	t.Run("implicit-write", func(t *testing.T) {
		p := harwp.NewResponseProxy(httptest.NewRecorder())

		// Given that write was called directly, ...
		_, err := p.Write([]byte("foo"))
		requirePkg.NoError(t, err)

		// ... then the headers shall have been written
		// implicitly as well.
		assertPkg.True(t, p.HeadersWritten())
	})

	t.Run("explicit-status", func(t *testing.T) {
		p := harwp.NewResponseProxy(httptest.NewRecorder())

		// Given that the headers have been written out
		// explicitly, ...
		p.WriteHeader(http.StatusBadRequest)

		// then.
		assertPkg.True(t, p.HeadersWritten())
	})
}

func Test_ProxyBase_BytesWritten(t *testing.T) {
	t.Run("initial-nothing-written", func(t *testing.T) {
		p := harwp.NewResponseProxy(httptest.NewRecorder())

		// Given that nothing has been written yet, then
		// BytesWritten shall report zero.
		assertPkg.Zero(t, p.BytesWritten())
	})

	t.Run("initial-headers-written", func(t *testing.T) {
		p := harwp.NewResponseProxy(httptest.NewRecorder())

		// Given that only the headers have been written yet, ...
		p.WriteHeader(http.StatusBadRequest)

		// then BytesWritten shall report zero.
		assertPkg.Zero(t, p.BytesWritten())
	})

	t.Run("implicit-write", func(t *testing.T) {
		p := harwp.NewResponseProxy(httptest.NewRecorder())

		// Given that write was called, ...
		_, err := p.Write([]byte("foo"))
		requirePkg.NoError(t, err)

		// ... then BytesWritten shall report the given bytes
		// written.
		assertPkg.Equal(t, 3, p.BytesWritten())
	})
}

func Test_ProxyBase_Header(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var p http.ResponseWriter = harwp.NewResponseProxy(w)

		// Given that the http handler writes a custom header to the response writer.
		p.Header().Set("X-Request-ID", "custom-request-id")

		_, _ = p.Write([]byte(""))
	}))
	defer srv.Close()

	resp, err := srv.Client().Get(srv.URL)
	requirePkg.NoError(t, err)

	_, err = io.ReadAll(resp.Body)
	requirePkg.NoError(t, err)

	assertPkg.Equal(t, "custom-request-id", resp.Header.Get("X-Request-ID"))
}

func Test_responseProxyBase_Unwrap(t *testing.T) {
	r := httptest.NewRecorder()

	p := harwp.NewResponseProxy(r)

	assertPkg.Same(t, r, p.Unwrap())
}

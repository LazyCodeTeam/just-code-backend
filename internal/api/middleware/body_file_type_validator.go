package middleware

import (
	"bytes"
	"io"
	"net/http"

	"github.com/LazyCodeTeam/just-code-backend/internal/api/util"
	"github.com/LazyCodeTeam/just-code-backend/internal/core/failure"
)

type splitedBody struct {
	head io.Reader
	tail io.ReadCloser
}

func (s *splitedBody) Read(p []byte) (int, error) {
	return io.MultiReader(s.head, s.tail).Read(p)
}

func (s *splitedBody) Close() error {
	return s.tail.Close()
}

func AcceptedBodyFileTypes(
	whitelisted ...string,
) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			buff := make([]byte, 512)
			_, err := r.Body.Read(buff)
			if err != nil {
				util.WriteError(w, err)
				return
			}
			mime := http.DetectContentType(buff)
			r.Body = &splitedBody{
				head: bytes.NewReader(buff),
				tail: r.Body,
			}
			for _, v := range whitelisted {
				if v == mime {
					next.ServeHTTP(w, r)
					return
				}
			}
			e := failure.NewInputFailure(
				failure.FailureTypeUnsupportedMediaType,
				"supported_types",
				whitelisted,
				"detected_type",
				mime,
			)
			util.WriteError(w, e)
		})
	}
}

// Package api JustCode API
//
//			 Documentation of JustCode API
//
//				BasePath: /
//				Version: 0.0.1
//				Contact: Mateusz Ledwoń<mateuszledwon@duck.com> https://github.com/Axot017
//
//				Consumes:
//				- application/json
//
//				Produces:
//				- application/json
//
//		    Security:
//		    - bearer_auth:
//
//		    SecurityDefinitions:
//		      bearer_auth:
//	         type: apiKey
//	         name: Authorization
//	         in: header
//
// swagger:meta
package api

import "github.com/LazyCodeTeam/just-code-backend/internal/api/dto"

// Empty response
// swagger:response emptyResponse
type emptyResponse struct{}

// swagger:response errorResponse
type errorResponse struct {
	// The error message
	// in: body
	Body dto.Error
}

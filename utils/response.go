package utils

import "backend/enums"

type APIResponse struct {
	StatusCode int                  `json:"statusCode"`
	Success    enums.ResponseStatus `json:"success"`
	Message    string               `json:"message"`
	Data       interface{}          `json:"data"`
}

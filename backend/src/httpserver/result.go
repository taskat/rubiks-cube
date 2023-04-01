package httpserver

import (
	eh "github.com/taskat/rubiks-cube/src/errorhandler"
)

type Result[T any] struct {
	Data   T            `json:"data,omitempty"`
	Errors []eh.Message `json:"errors,omitempty"`
}

func NewResult[T any](data T, errors []eh.Message) Result[T] {
	return Result[T]{Data: data, Errors: errors}
}

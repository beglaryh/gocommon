package collection_errors

import "errors"

type CollectionError error

var (
	LimitExceeded    CollectionError = errors.New("limit exceeded")
	IndexOutOfBounds CollectionError = errors.New("index out of bounds")
)

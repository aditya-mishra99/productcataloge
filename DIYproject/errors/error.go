package errors

import "errors"

var InsufficientProduct = errors.New("product quantity insufficient")
var ProductNotFound = errors.New("product does not exist")

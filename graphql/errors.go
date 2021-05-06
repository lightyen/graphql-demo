package graphql

import "errors"

var ErrAuthentication = errors.New("Authentication Error") // 401
var ErrAuthorization = errors.New("Authorization Error")   // 403

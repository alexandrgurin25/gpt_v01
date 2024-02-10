package common

import "errors"

// InternalError представляет ошибку внутреннего сервера.
var InternalError = errors.New("internal error")

// NotFoundError представляет ошибку "не найдено".
var NotFoundError = errors.New("not found")

// LogicError представляет логическую ошибку.
var LogicError = errors.New("logic error")

// ForbiddenError представляет ошибку доступа запрещено.
var ForbiddenError = errors.New("forbidden error")

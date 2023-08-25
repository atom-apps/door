package errorx

import (
	"net/http"

	"github.com/rogeecn/fen"
)

var ErrInvalidRedirectURL = fen.NewBusError(http.StatusBadRequest, http.StatusBadRequest, "跳转地址不合法")

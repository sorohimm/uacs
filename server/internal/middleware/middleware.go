package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
	"uacs/internal/config"
)

type MiddlewareV0 struct {
	Log        *zap.SugaredLogger
	HttpClient *http.Client
	Cfg        *config.Config
}

func (m *MiddlewareV0) AuthRequired(ctx *gin.Context) {
	cookie, err := ctx.Cookie("accessToken")
	if err != nil || cookie == "" {
		ctx.Status(http.StatusUnauthorized)
		return
	}

	statusCode, err := m.validateAccessToken(ctx)
	if err != nil || statusCode != 200 {
		ctx.AbortWithError(statusCode, errors.New("validation failed"))
		return
	}

	ctx.Next()
}

func (m *MiddlewareV0) validateAccessToken(ctx *gin.Context) (int, error) {
	req, err := http.NewRequest(http.MethodGet, m.Cfg.SsoCfg.TokenValidateEndpoint, nil)

	req.Header.Set("Cookie", ctx.Request.Header.Get("Cookie"))
	resp, err := m.HttpClient.Do(req)

	return resp.StatusCode, err
}

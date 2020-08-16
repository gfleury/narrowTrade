package saxo_models

import (
	"context"

	"github.com/gfleury/intensiveTrade/saxo_openapi"
)

type ModeledAPI struct {
	Ctx    context.Context
	Client *saxo_openapi.APIClient
}

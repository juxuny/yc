package handler

import (
	"context"
	"github.com/juxuny/yc/router"
)

var TriggerConfig = router.TriggerConfig{
	router.TriggerSignValidator: func(ctx context.Context, body []byte) error {
		return nil
	},
}

package jobhandler

import (
	"context"
	"github.com/hawkj/my_iot/iot_server/pkg/common"
)

type JobHandler func(ctx context.Context, g *common.Global)

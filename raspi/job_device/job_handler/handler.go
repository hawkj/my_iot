package jobhandler

import (
	"context"
	"github.com/hawkj/my_iot/raspi/pkg/common"
)

type JobHandler func(ctx context.Context, g *common.Global)

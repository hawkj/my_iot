package jobhandler

import (
	"context"
	"fmt"
	"github.com/hawkj/my_iot/raspi/pkg/common"
)

func Bme280(ctx context.Context, g *common.Global) {
	fmt.Println(ctx.Value("params"))
	fmt.Println("Bme280")
	fmt.Println(g.Config.SiteInfo.Name)
}

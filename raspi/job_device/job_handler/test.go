package jobhandler

import (
	"context"
	"fmt"
	"github.com/hawkj/my_iot/raspi/pkg/common"
)

func Test(ctx context.Context, g *common.Global) {
	fmt.Println(ctx.Value("params"))
	fmt.Println("test")
	fmt.Println(g.Config.SiteInfo.Name)
}

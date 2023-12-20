package jobhandler

import (
	"context"
	"fmt"
	"github.com/hawkj/my_iot/iot_server/pkg/common"
)

func Test(ctx context.Context, g *common.Global) {
	fmt.Println(ctx.Value("params"))
	fmt.Println("test")

}

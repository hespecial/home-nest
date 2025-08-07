package listen

import (
	"context"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
	"home-nest/app/order/cmd/mq/internal/config"
	kqMq "home-nest/app/order/cmd/mq/internal/mqs/kq"
	"home-nest/app/order/cmd/mq/internal/svc"
)

func KqMqs(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {

	return []service.Service{
		//Listening for changes in consumption flow status
		kq.MustNewQueue(c.PaymentUpdateStatusConf, kqMq.NewPaymentUpdateStatusMq(ctx, svcContext)),
		//.....
	}

}

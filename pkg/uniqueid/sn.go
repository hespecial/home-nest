package uniqueid

import (
	"fmt"
	"home-nest/pkg/tool"
	"time"
)

// 生成sn单号
type SnPrefix string

const (
	SnPrefixHomestayOrder SnPrefix = "HSO" //民宿订单前缀
	SnPrefixThirdPayment  SnPrefix = "PMT" //第三方支付流水记录前缀
)

func GenSn(snPrefix SnPrefix) string {
	return fmt.Sprintf("%s%s%s", snPrefix, time.Now().Format("20060102150405"), tool.Krand(8, tool.KcRandKindNum))
}

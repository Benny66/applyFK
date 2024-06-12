package main

import (
	"applyFK/visit"
)

var token string = "xx"
var ticket string = "xx-xx"

func main() {
	// 获取访问的token和ticket
	v := visit.NewApply()

	v.SetToken(token)
	v.SetTicket(ticket)

	// 申请一周内的访客码，自动规避节假日和法定节假日；
	// 如果今天是假期则自动调整日期到最近的工作日；
	// 如果今天是工作日则判断是否已经申请过访客码，如果已经申请过则不再申请；
	v.ApplyOneWeek()

	// 申请制定日期的访客码和次数，默认1人
	// v.ApplyDate("2024/01/01", 1)

	// 申请多人的访客码则
	// v.SetApplyNum(2)
	// v.ApplyDate("2024/01/01", 1)

	// 修改访问人的电话
	// v.SetAppliedPhone("135xxxxxxxxx")
	// v.ApplyDate("2024/01/01", 1)

}

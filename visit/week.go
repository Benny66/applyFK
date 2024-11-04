package visit

import (
	"applyFK/holiday"
	"fmt"
)

func (v *Apply) ApplyOneWeek() error {
	fmt.Println("开始申请一周")
	hadApplyDays, err := v.GetHadApplyDays(2) // 获取已申请天数
	if err != nil {
		fmt.Printf("获取已申请天数失败%s", err)
		return err
	}
	firstWorkDay := false
	for i := 0; i < 30; i++ { // 限制30天循环申请
		currentDate := holiday.CurrentDayStr(i)
		applyDays := 3 // 每天申请3次
		if _, ok := hadApplyDays[currentDate]; ok {
			days := hadApplyDays[currentDate]
			if days >= 3 {
				firstWorkDay = true
				fmt.Printf("时间%s 已申请3天，跳过\n", currentDate)
				continue
			}
			applyDays = 3 - days
			fmt.Printf("时间%s 已申请%d次,还需要%d次\n", currentDate, days, applyDays)
		}
		isHoliday := holiday.IsHoliday(currentDate)
		if firstWorkDay && isHoliday {
			fmt.Printf("时间%s 已到达下一周期的节假日,退出申请\n", currentDate)
			break
		} else if isHoliday {
			fmt.Printf("时间%s 跳过当前周期的节假日\n", currentDate)
			continue
		}
		fmt.Printf("时间%s 开始申请\n", currentDate)
		err = v.ApplyDate(currentDate, applyDays)
		if err != nil {
			return err
		}
		firstWorkDay = true
	}
	fmt.Println("一周申请结束")
	return nil
}

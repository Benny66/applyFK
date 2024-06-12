package visit

import (
	"errors"
	"fmt"
)

func (v *Apply) ApplyDate(date string, num int) error {
	if num <= 0 {
		return errors.New("num should be greater than 0")
	}
	for i := 0; i < num; i++ {
		v.SetApplyDate(date)
		err := v.Apply()
		if err != nil {
			fmt.Printf("时间%s 第%d次申请失败\n", date, i+1)
			return err
		}
		fmt.Printf("时间%s 第%d次申请成功\n", date, i+1)
	}
	return nil
}

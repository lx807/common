package common

import "encoding/json"

//通过json tag进行结构体赋值 只要tag相同，就相当于赋值一样
//其实就是和逐一赋值一样
func SwapTo(resq, category interface{}) error {
	data, err := json.Marshal(resq)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, category)
}

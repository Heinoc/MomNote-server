package common

import (
	"reflect"
)

/**
 * Created by Heinoc on 2018/8/30.
 */

type Response struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result"`
}

func NewResponse(result interface{}) (*Response) {
	//如果是数组类型，则手动将当前的result封装给"list"这个jsonObject key
	if reflect.TypeOf(result).Kind() == reflect.Array | reflect.Slice {
		return &Response{
			Code: 0,
			Msg:  "success",
			Result: map[string]interface{}{
				"list": result,
			},
		}
	} else {
		return &Response{
			Code:   0,
			Msg:    "success",
			Result: result,
		}
	}
}

func NewErrorResponse(err string) (*Response) {
	return &Response{
		Code:   -1,
		Msg:    err,
		Result: "",
	}
}

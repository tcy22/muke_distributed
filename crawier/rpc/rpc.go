package rpc

import "github.com/pkg/errors"

type DemoService struct {
}

type Args struct {
	A, B  int
}

//只有满足如下标准的方法才能用于远程访问。
//DemoService是远程访问的对象，Div是远程访问的对象的函数
//args是远程访问所带的参数
//result是远程访问的结果
func (DemoService) Div(args Args,result *float64) error {
	if args.B == 0 {
		return  errors.New("division by zero")
	}

	*result = float64(args.A) / float64(args.B)
	return nil
}
// Code generated by Kitex v0.4.2. DO NOT EDIT.

package detailsservice

import (
	details "github.com/cloudwego/biz-demo/bookinfo/kitex_gen/cwg/bookinfo/details"
	server "github.com/cloudwego/kitex/server"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler details.DetailsService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}

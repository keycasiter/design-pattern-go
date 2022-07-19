package struct_pattern

import "fmt"

type SrvAct interface {
	Do() string
}

type Srv struct {
	app string
}

func NewSrv() *Srv {
	return &Srv{app: "default"}
}

func (s *Srv) Do() string {
	return fmt.Sprintf("%s is working", s.app)
}

type ProxyLogSrv struct {
	srv *Srv
}

func NewProxyLogSrv() *ProxyLogSrv {
	return &ProxyLogSrv{srv: NewSrv()}
}

func (proxy *ProxyLogSrv) Do() string {
	fmt.Println("log->begin")
	res := proxy.srv.Do()
	fmt.Println("log->processing")
	fmt.Println("log->end")
	return res
}

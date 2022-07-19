package struct_pattern

type SrvAction interface {
	Do() string
}

type QuerySrv struct {
	action string
}

func NewQuerySrv() *QuerySrv {
	return &QuerySrv{action: "query"}
}

func (srv *QuerySrv) Do() string {
	return srv.action
}

type CreateSrv struct {
	action string
}

func NewCreateSrv() *CreateSrv {
	return &CreateSrv{action: "create"}
}

func (srv *CreateSrv) Do() string {
	return srv.action
}

type UpdateSrv struct {
	action string
}

func NewUpdateSrv() *UpdateSrv {
	return &UpdateSrv{action: "update"}
}

func (srv *UpdateSrv) Do() string {
	return srv.action
}

type FacadeAction interface {
	Query() string
	Create() string
	Update() string
}

type FacadeSrv struct {
	querySrv  *QuerySrv
	createSrv *CreateSrv
	updateSrv *UpdateSrv
}

func NewFacadeSrv() *FacadeSrv {
	return &FacadeSrv{
		querySrv:  NewQuerySrv(),
		createSrv: NewCreateSrv(),
		updateSrv: NewUpdateSrv(),
	}
}

func (f *FacadeSrv) Query() string {
	return f.querySrv.Do()
}

func (f *FacadeSrv) Create() string {
	return f.createSrv.Do()
}

func (f *FacadeSrv) Update() string {
	return f.updateSrv.Do()
}

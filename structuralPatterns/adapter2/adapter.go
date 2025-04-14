package adapter2

type Adapter struct {
	Server *Server
}

func (a *Adapter) ConvertToStringInterface() string {
	return string(a.Server.Get())
}

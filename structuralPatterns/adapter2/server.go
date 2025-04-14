package adapter2

type Server struct {
	lst []byte
}

func (s *Server) Initialize() {
	s.lst = append(s.lst, []byte("Hello World")...)
	s.lst = append(s.lst, []byte("1")...)
	s.lst = append(s.lst, []byte("2")...)
	s.lst = append(s.lst, []byte("3")...)
}

func (s *Server) Get() []byte {
	return s.lst
}

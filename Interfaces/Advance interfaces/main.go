// advance interface mechanics and type function...

package main

type Storage interface {
	Get(id int) (any, error)
	Put(id int, val any) error
}

type foodStroge struct {
}

func (s *foodStroge) Get(id int) (any, error) {
	return nil, nil
}
func (s *foodStroge) Put(id int, val any) error {
	return nil
}

type Server struct {
	store Storage
}

func main() {
	s := &Server{
		store: &foodStroge{},
	}

	//logic

	s.store.Get(1)
	s.store.Put(1, "foo")

}

package main

type name struct {
	Name int `json:"name"`
}

func main() {
	m := make(map[*name]*name)
	ns := make([]*name, 0)
	for i := 0; i < 1000; i++ {
		n := &name{
			Name: i,
		}
		ns = append(ns, n)
		m[n] = &name{}
	}
	for i := 0; i < 1000; i++ {
		go func(i int) {
			//m[ns[i]] = &name{
			//	Name: 0,
			//}
			m[ns[i]] = &name{}
		}(i)
	}
}

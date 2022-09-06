package main

func main() {
	s, _ := InjectDependencies()
	s.Log.Debugf("%#v", s.Config)

	s.Log.Fatalf("site: %s", s.Router.Run())
}

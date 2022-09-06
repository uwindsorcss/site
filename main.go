package main

func main() {
	s, _ := InjectDependencies()
	s.Log.Debugf("%#v", s.Config)

	s.Router.Run()
}

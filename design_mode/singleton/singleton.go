package singleton

type Singleton struct{}

var singleton *Singleton

func init() {
	singleton = &Singleton{}
}

func Instance() *Singleton {
	return singleton
}

package states

type State interface {
	Load()
	Update()
	Draw()
	Reset()
	Close()
}

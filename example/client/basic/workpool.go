package basic

// WorkPool interface
type WorkPool interface {
	Start() error
	Stop() error
	Submit(Job) error
}

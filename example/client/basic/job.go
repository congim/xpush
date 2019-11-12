package basic

// Job workpool submit job
type Job interface {
	Work() error
}

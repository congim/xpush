package compress

// Compress interface
type Compress interface {
	Compress([]byte) ([]byte, error)
	UnCompress([]byte) ([]byte, error)
}

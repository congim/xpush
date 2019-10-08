package channel

var _ Channel = (*noopChannel)(nil)

type noopChannel struct {
}

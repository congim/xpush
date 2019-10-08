package cassandra

type Cassandra struct {
}

func New() *Cassandra {
	return &Cassandra{}
}

func (c *Cassandra) Store() error {
	return nil
}

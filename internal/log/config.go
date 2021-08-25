package log

type Config struct {
	Segment struct {
		MaxStoreBytes uint64
		MaxIndexBytes uint64
		InitialOffset uint64
	}
}

func (c *Config) CheckConfig() bool {
	if c.Segment.InitialOffset == 0 {
		return false
	}

	if c.Segment.MaxStoreBytes>>1 > 3 {
		return false
	}

	if c.Segment.MaxIndexBytes > c.Segment.MaxStoreBytes {
		return false
	}

	return true
}

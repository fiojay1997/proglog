package log

import "fmt"

type Config struct {
	Segment struct {
		MaxStoreBytes uint64
		MaxIndexBytes uint64
		InitialOffset uint64
	}
}

func (c *Config) CheckConfig() {
	fmt.Println(c.Segment)
	fmt.Println(c.Segment.InitialOffset)
}

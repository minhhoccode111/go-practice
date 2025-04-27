package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	h int
	m int
}

func New(h, m int) Clock {
	modulo(&h, &m)
	return Clock{h, m}
}

func (c Clock) Add(m int) Clock {
	h := c.h
	m = c.m + m
	modulo(&h, &m)
	return Clock{h, m}
}

func (c Clock) Subtract(m int) Clock {
	h := c.h
	m = c.m - m
	modulo(&h, &m)
	return Clock{h, m}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

func modulo(h *int, m *int) {
	for *m < 0 {
		*m += 60
		*h -= 1
	}
	for *h < 0 {
		*h += 24
	}
	for *m >= 60 {
		*m -= 60
		*h += 1
	}
	for *h >= 24 {
		*h -= 24
	}
}

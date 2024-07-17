package main

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Reader2 interface {
	Read() (p []byte, err error)
}

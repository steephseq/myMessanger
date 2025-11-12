package models

type Config struct {
	Address  string
	Services []Service
}

type Service struct {
	Name  string
	Port  string
	Paths []string
}

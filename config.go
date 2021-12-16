package main

type Config struct {
	Port int `env:"PORT" envDefault:"3000"`
}

package main

type config struct {
	Token     string     `fig:"token" validate:"required"`
	LogLevel  string     `fig:"loglevel" validate:"required"`
	Endpoints []endpoint `fig:"endpoints" validate:"required"`
	IDs       []int64    `fig:"ids" validate:"required"`
}

type endpoint struct {
	URL  string `fig:"url" validate:"required"`
	Name string `fig:"name" validate:"required"`
}

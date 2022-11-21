package main

type config struct {
	Token    string  `fig:"token" validate:"required"`
	LogLevel string  `fig:"loglevel" validate:"required"`
	Endpoint string  `fig:"endpoint" validate:"required"`
	IDs      []int64 `fig:"ids" validate:"required"`
	apiKey   string  `fig:"apikey" validate:"required"`
}

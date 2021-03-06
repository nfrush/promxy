package proxyconfig

import "github.com/prometheus/prometheus/config"

type PromReloadable interface {
	ApplyConfig(*config.Config) error
}

type Reloadable interface {
	ApplyConfig(*Config) error
}

type PromReloadableWrap struct {
	R PromReloadable
}

func (p *PromReloadableWrap) ApplyConfig(c *Config) error {
	return p.R.ApplyConfig(&c.PromConfig)
}

func WrapPromReloadable(p PromReloadable) Reloadable {
	return &PromReloadableWrap{p}
}

// ApplyConfigFunc is a struct that wraps a single function that Applys config
// into something that implements the `PromReloadable` interface
type ApplyConfigFunc struct {
	F func(*config.Config) error
}

func (a *ApplyConfigFunc) ApplyConfig(cfg *config.Config) error {
	return a.F(cfg)
}

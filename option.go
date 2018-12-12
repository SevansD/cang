package cang

import "context"

type (
	Option func(*App) error
)

func WithServices(services ...Service) Option {
	return func(a *App) error {
		a.services = make(map[string]Service, 0)
		for _, s := range services {
			a.services[s.GetName()] = s
		}
		return nil
	}
}

func WithContext(ctx context.Context) Option {
	return func(a *App) error {
		a.ctx = ctx
		return nil
	}
}

package database

import (
    "gitlab.com/zenport.io/go-assignment/engine"
    "gitlab.com/zenport.io/go-assignment/domain"
)

type Provider struct {
}

func (provider *Provider) GetKnightRepository() engine.KnightRepository {
	return &knightRepository{make(map[string]domain.Knight)}
}

func (provider *Provider) Close() {

}

func NewProvider() *Provider {
	return &Provider{}
}

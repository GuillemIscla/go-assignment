package engine

import (
	"errors"
	"fmt"

	"gitlab.com/zenport.io/go-assignment/domain"
)

func (engine *arenaEngine) GetKnight(ID string) (*domain.Knight, error) {
	fighter := engine.knightRepository.Find(ID)
	if fighter == nil {
		return nil, errors.New(fmt.Sprintf("fighter with ID '%s' not found!", ID))
	}

	return fighter, nil
}

func (engine *arenaEngine) ListKnights() []domain.Knight {
	return engine.knightRepository.FindAll()
}

func (engine *arenaEngine) Fight(fighter1ID string, fighter2ID string) domain.Fighter {
	fighter1, err1 := engine.GetKnight(fighter1ID)
	if err1 != nil {
	    panic(err1)
	}
    fighter2, err2 := engine.GetKnight(fighter2ID)
    if err2 != nil {
        panic(err2)
    }
    return engine.arena.Fight(fighter1, fighter2)
}

func (engine *arenaEngine) Save(knight domain.Knight) {
    engine.knightRepository.Save(knight)
}

package database

import "gitlab.com/zenport.io/go-assignment/domain"

type knightRepository struct{
    KnightsData map[string]domain.Knight
}

func (repository *knightRepository) Find(ID string) *domain.Knight {
    value, ok  := repository.KnightsData[ID]
    if !ok {
        return nil
    }
	return &value
}

func (repository *knightRepository) FindAll() []domain.Knight {

    values := make([]domain.Knight, 0, len(repository.KnightsData))

    for _, v := range repository.KnightsData {
    	values = append(values, v)
    }

    return values
}

func (repository *knightRepository) Save(knight domain.Knight) {
    repository.KnightsData[knight.Id] = knight
}

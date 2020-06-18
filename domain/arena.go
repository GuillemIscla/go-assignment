package domain

type Arena struct{}

func (arena *Arena) Fight(fighter1 Fighter, fighter2 Fighter) Fighter {
    if fighter1.GetPower() > fighter2.GetPower() {
            return fighter1
    }
    if fighter2.GetPower() > fighter1.GetPower() {
            return fighter2
    }
	return nil
}

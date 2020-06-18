package domain

type Fighter interface {
	GetID() string
	GetPower() float64
}

type Knight struct {
    Id       string `json:"id"`
    Name     string `json:"name"`
    Strength int `json:"strength"`
    WeaponPower float64 `json:"weapon_power"`
}

func (k *Knight) GetID() string {
    return k.Id
}

func (k *Knight) GetPower() float64 {
    return float64(k.Strength) + k.WeaponPower
}

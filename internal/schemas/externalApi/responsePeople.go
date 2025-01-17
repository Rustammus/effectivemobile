package externalApi

type ExResponsePeople struct {
	Surname    *string `json:"surname" example:"Иванов"`
	Name       *string `json:"name" example:"Иван"`
	Patronymic *string `json:"patronymic" example:"Иванович"`
	Address    *string `json:"address" example:"г. Москва, ул. Ленина, д. 5, кв. 1"`
}

func (p ExResponsePeople) Valid() bool {
	ok := p.Surname != nil && p.Name != nil && p.Address != nil
	return ok
}

package user

import "time"

type Usuario struct {
	Id     int
	Nombre string
	Fecha  time.Time
	Status bool
}

func (user *Usuario) Create(id int, nombre string, fecha time.Time, status bool) {
	user.Id = id
	user.Nombre = nombre
	user.Status = status
	user.Fecha = fecha
}

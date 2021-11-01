package main

import "fmt"

type serVivo interface {
	estaVivo() bool
}

type animal interface {
	comer()
	respirar()
	esCarnivoro() bool
}

type vegetal interface {
	clasificacionVegetal() string
}

type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
}

type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	esHombre   bool
	vivo       bool
}

type mujer struct {
	hombre
}

func (h *hombre) respirar() { h.respirando = true }

func (h *hombre) pensar() { h.pensando = true }

func (h *hombre) comer() { h.comiendo = false }

func (h *hombre) estaVivo() bool { return h.vivo }

func (h *hombre) sexo() string {
	if h.esHombre == true {
		return "Hombre"
	} else {
		return "Muejr"
	}
}

func HumanoRespirando(hu humano) {
	hu.respirar()
	fmt.Printf("Soy un/a %s y ya estoy respirando \n", hu.sexo())
}

type perro struct {
	respirando bool
	comiendo   bool
	carnivoro  bool
	vivo       bool
}

func (p *perro) respirar() { p.respirando = true }

func (p *perro) comer() { p.comiendo = false }

func (p *perro) esCarnivoro() bool { return p.carnivoro }

func (p *perro) estaVivo() bool { return p.vivo }

func animalesRespirar(an animal) {
	an.respirar()
	fmt.Println("Soy un animal y estoy respirando")
}

func animalCarnivoro(an animal) int {
	if an.esCarnivoro() == true {
		return 1
	}

	return 0
}

func estoyVivo(ser serVivo) bool {
	return ser.estaVivo()
}

func main() {
	German := new(hombre)
	German.esHombre = true
	German.vivo = true
	HumanoRespirando(German)

	Alejandra := new(mujer)
	Alejandra.vivo = true
	HumanoRespirando(Alejandra)

	estoyVivo(German)
	estoyVivo(Alejandra)

	totalCarnivoros := 0
	labrador := new(perro)
	pincher := new(perro)
	sanbernardo := new(perro)

	labrador.carnivoro = true
	labrador.vivo = true

	pincher.carnivoro = true
	sanbernardo.carnivoro = false

	totalCarnivoros += animalCarnivoro(labrador)
	totalCarnivoros += animalCarnivoro(pincher)
	totalCarnivoros += animalCarnivoro(sanbernardo)

	fmt.Printf("Animales carnivoros: %d \n", totalCarnivoros)
	fmt.Printf("labrador vivo: %t \n", estoyVivo(labrador))
	fmt.Printf("Alejandra vivo: %t \n", estoyVivo(Alejandra))
	fmt.Printf("German vivo: %t \n", estoyVivo(German))
}

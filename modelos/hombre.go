package modelos

type Hombre struct {
	Edad       int
	Altura     float32
	Peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	Vivo       bool
}

func (h *Hombre) Respirar() {
	h.respirando = true
}
func (h *Hombre) Comer() {
	h.comiendo = true
}
func (h *Hombre) Pensar() {
	h.pensando = true
}
func (h *Hombre) Sexo() string {
	return "Hombre"
}
func (h *Hombre) EstaVivo() bool {
	return true
}

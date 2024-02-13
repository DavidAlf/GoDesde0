package modelos

type Mujer struct {
	Hombre
}

func (m *Mujer) Respirar() {
	m.respirando = true
}
func (m *Mujer) Comer() {
	m.comiendo = true
}
func (m *Mujer) Pensar() {
	m.pensando = true
}
func (h *Mujer) Sexo() string {
	return "Mujer"
}
func (h *Mujer) EstaVivo() bool {
	return false
}

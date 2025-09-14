package model

type ComportamentoIA int

const (
	Formal ComportamentoIA = iota
	Informal
	Profissional
	Amigavel
	Educado
	Conciso
	Detalhado
	Divertido
	Respeitoso
)

func (c ComportamentoIA) String() string {
	nomes := [...]string{
		"Formal",
		"Informal",
		"Profissional",
		"Amigável",
		"Educado",
		"Conciso",
		"Detalhado",
		"Divertido",
		"Respeitoso",
	}

	if c < Formal || c > Respeitoso {
		return "ComportamentoIA Desconhecido"
	}

	return nomes[c]
}

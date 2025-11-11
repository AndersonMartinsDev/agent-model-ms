package model

type Behavioural int

const (
	Formal Behavioural = iota
	Informal
	Profissional
	Amigavel
	Educado
	Conciso
	Detalhado
	Divertido
	Respeitoso
)

func (c Behavioural) String() string {
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

	return nomes[c]
}

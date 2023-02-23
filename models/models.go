package models

type Candidato struct {
	Nome         string `json:"nome"`
	CPF          string `json:"cpf"`
	RG           string `json:"rg"`
	TicketNumber string `json:"ticketnumber"`
}

var Candidatos []Candidato

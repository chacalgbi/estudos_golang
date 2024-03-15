package modelos

import (
	"errors"
	"strings"
	"time"
)

type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criado_em,omitempty"`
}

func (u *Usuario) Preparar(etapa string) error {
	u.formatar()
	if erro := u.validar(etapa); erro != nil {
		return erro
	}
	return nil
}

func (u Usuario) validar(etapa string) error {
	if u.Nome == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}

	if u.Nick == "" {
		return errors.New("O nick é obrigatório e não pode estar em branco")
	}

	if u.Email == "" {
		return errors.New("O email é obrigatório e não pode estar em branco")
	}

	if etapa == "cadastro" && u.Senha == "" {
		return errors.New("A senha é obrigatória e não pode estar em branco")
	}

	return nil
}

func (u *Usuario) formatar() {
	u.Nome = strings.TrimSpace(u.Nome)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)
}

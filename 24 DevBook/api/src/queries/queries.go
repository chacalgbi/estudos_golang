package queries

type Queries struct {
	CriarUsuario       string
	BuscarUsuarios     string
	BuscarUsuarioPorID string
	AtualizarUsuario   string
	DeletarUsuario     string
}

var Q = Queries{
	CriarUsuario:       "INSERT INTO usuarios (nome, nick, email, senha) VALUES ($1, $2, $3, $4) RETURNING id",
	BuscarUsuarios:     "SELECT nome, nick, email, Criadoem FROM usuarios WHERE nome LIKE $1 OR nick LIKE $2",
	BuscarUsuarioPorID: "SELECT id, nome, nick, email, Criadoem FROM usuarios WHERE id = $1",
	AtualizarUsuario:   "UPDATE usuarios SET nome = $1, nick = $2, email = $3 WHERE id = $4",
	DeletarUsuario:     "DELETE FROM usuarios WHERE id = $1",
}

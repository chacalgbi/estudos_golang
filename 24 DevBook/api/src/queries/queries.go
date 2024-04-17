package queries

type Queries struct {
	CriarUsuario       string
	BuscarUsuarios     string
	BuscarUsuarioPorID string
	AtualizarUsuario   string
	DeletarUsuario     string
	BuscarPorEmail     string
	SeguirUsuario      string
	PararDeSeguir      string
	BuscarSeguidores   string
	BuscarSeguindo     string
}

// Postgres queries
// var Q = Queries{
// 	CriarUsuario:       "INSERT INTO usuarios (nome, nick, email, senha) VALUES ($1, $2, $3, $4) RETURNING id",
// 	BuscarUsuarios:     "SELECT nome, nick, email, Criadoem FROM usuarios WHERE nome LIKE $1 OR nick LIKE $2",
// 	BuscarUsuarioPorID: "SELECT id, nome, nick, email, Criadoem FROM usuarios WHERE id = $1",
// 	AtualizarUsuario:   "UPDATE usuarios SET nome = $1, nick = $2, email = $3 WHERE id = $4",
// 	DeletarUsuario:     "DELETE FROM usuarios WHERE id = $1",
// 	BuscarPorEmail:     "SELECT id, senha FROM usuarios WHERE email = $1",
// }

// Mysql queries
var Q = Queries{
	CriarUsuario:       "INSERT INTO usuarios (nome, nick, email, senha) VALUES (?, ?, ?, ?)",
	BuscarUsuarios:     "SELECT nome, nick, email, criadoEm FROM usuarios WHERE nome LIKE ? OR nick LIKE ?",
	BuscarUsuarioPorID: "SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE id = ?",
	AtualizarUsuario:   "UPDATE usuarios SET nome = ?, nick = ?, email = ? WHERE id = ?",
	DeletarUsuario:     "DELETE FROM usuarios WHERE id = ?",
	BuscarPorEmail:     "SELECT id, senha FROM usuarios WHERE email = ?",
	SeguirUsuario:      "INSERT ignore INTO seguidores (usuario_id, seguidor_id) VALUES (?, ?)",
	PararDeSeguir:      "DELETE FROM seguidores WHERE usuario_id = ? AND seguidor_id = ?",
	BuscarSeguidores:   "SELECT u.id, u.nome, u.nick, u.email, u.criadoEm FROM usuarios u JOIN seguidores s ON u.id = s.seguidor_id WHERE s.usuario_id = ?",
	BuscarSeguindo:     "SELECT u.id, u.nome, u.nick, u.email, u.criadoEm FROM usuarios u JOIN seguidores s ON u.id = s.usuario_id WHERE s.seguidor_id = ?",
}

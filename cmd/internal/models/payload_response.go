package models

type PayLoadResponse struct {
	Status  string      `json:"status"`  // Status da operação (ex: "success", "error")
	Message string      `json:"message"` // Mensagem de resposta (ex: "User registered successfully")
	Data    interface{} `json:"data"`    // Dados adicionais (opcional, pode ser um usuário ou outro payload)
}

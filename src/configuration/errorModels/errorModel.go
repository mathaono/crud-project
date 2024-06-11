package errormodels

import (
	"net/http"
)

type ErrorModel struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int64    `json:"code"`
	Causes  []Causes `json:"causes,omitempty"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Modelo da estrutura de resposta da mensagem de erro
func NewErrorModel(message, err string, code int64, causes []Causes) *ErrorModel {
	return &ErrorModel{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

// Erro de requisição, não foi enviado algum campo obrigatório, http 400
func BadRequestError(message string) *ErrorModel {
	return &ErrorModel{
		Message: message,
		Err:     "Bad Request",
		Code:    http.StatusBadRequest,
	}
}

// Erro de validação de acesso, token ou outro tipo de validação está bloqueado, http 401
func UnauthorizedError(message string) *ErrorModel {
	return &ErrorModel{
		Message: message,
		Err:     "Unauthorized",
		Code:    http.StatusUnauthorized,
	}
}

// Erro interno de sistema, problema de endpoint e/ou db, http 500
func InternalServerError(message string) *ErrorModel {
	return &ErrorModel{
		Message: message,
		Err:     "Internal Server Error",
		Code:    http.StatusInternalServerError,
	}
}

// Erro genérico, quando não atende nenhuma condição dos erros acima, http 550
func GenericError(message string) *ErrorModel {
	return &ErrorModel{
		Message: message,
		Err:     "Generic Error",
		Code:    550,
	}
}

// função instanciada da struct ErrorModel onde só retorna uma string, usada para tratamento de erros de certas funções e bibliotecas que tem como retorno o resultado + o erro
// não sei se é a melhor abordagem, é só um teste
func (m *ErrorModel) Error() string {
	return m.Message
}

package rooms

// DTOs de Mensagens comuns
type MessageDTO struct {
	Content string `json:"content" validate:"required"`
}

type MessageResponseDTO struct {
	Content string `json:"content" validate:"required"`
}

// DTOS de Comando
type Arg struct {
	Content string `json:"content" validate:"required"`
	Param   string `json:"param" validate:"required"`
}

type CommandDTO struct {
	Type string `json:"type" validate:"required"`
	Args []Arg  `json:"args" validate:"required"`
}

type CommandResponseDTO struct {
	MessageResponseDTO
}

type CreateRoomDTO struct {
	Name string `json:"name" validate:"required"`
}
type UpdateRoomDTO struct {
	Name string `json:"name" validate:"required"`
}

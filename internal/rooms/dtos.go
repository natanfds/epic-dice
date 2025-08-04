package rooms

type MessageDTO struct {
	Content string `json:"content" validate:"required"`
}

type Arg struct {
	Content string `json:"content" validate:"required"`
	Param   string `json:"param" validate:"required"`
}

type CommandDTO struct {
	Type string `json:"type" validate:"required"`
	Args []Arg  `json:"args" validate:"required"`
}

type MessageResponseDTO struct {
	Content string `json:"content" validate:"required"`
}

type CommandResponseDTO struct {
	MessageResponseDTO
}

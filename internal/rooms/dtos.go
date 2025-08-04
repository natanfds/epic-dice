package rooms

type MessageDTO struct {
	Content string `json:"content", binding:"required"`
}

type Arg struct {
	Content string `json:"content", binding:"required"`
	Param   string `json:"param", binding:"required"`
}

type CommandDTO struct {
	Type string `json:"type", binding:"required"`
	Args []Arg  `json:"args", binding:"required"`
}

type MessageResponseDTO struct {
	Content string `json:"content", binding:"required"`
}

type CommandResponseDTO struct {
	MessageResponseDTO
}

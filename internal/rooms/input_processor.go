package rooms

import "encoding/json"

type InputProcessor interface {
	Exec() ([]byte, error)
}

type MessageProcessor struct {
	Message MessageDTO
}

func (r *MessageProcessor) Exec() ([]byte, error) {
	return json.Marshal(MessageResponseDTO{Content: r.Message.Content})
}

func NewMessageProcessor(message MessageDTO) *MessageProcessor {
	return &MessageProcessor{Message: message}
}

type CommandProcessor struct {
	Command CommandDTO
}

func (r *CommandProcessor) Exec() ([]byte, error) {
	return json.Marshal(CommandResponseDTO{
		MessageResponseDTO: MessageResponseDTO{
			Content: r.Command.Type,
		},
	})
}

func NewCommandProcessor(command CommandDTO) *CommandProcessor {
	return &CommandProcessor{Command: command}
}

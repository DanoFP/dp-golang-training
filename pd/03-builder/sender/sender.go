package sender

import (
	"github.com/DanoFP/dp-golang-training/pd/03-builder/builder"
	"github.com/DanoFP/dp-golang-training/pd/03-builder/message"
)

type Sender struct {
	builder builder.MessageBuilder
}

func (s *Sender) SetBuilder(b builder.MessageBuilder) {
	s.builder = b
}

func (s *Sender) BuildMessage(recipient, message string) (*message.MessageRepresented, error) {
	s.builder.SetRecipient(recipient).SetMessage(message)
	return s.builder.Build()
}

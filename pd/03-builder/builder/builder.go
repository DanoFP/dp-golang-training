package builder

import "github.com/DanoFP/dp-golang-training/pd/03-builder/message"

type MessageBuilder interface {
	SetRecipient(recipient string) MessageBuilder
	SetMessage(message string) MessageBuilder
	Build() (*message.MessageRepresented, error)
}

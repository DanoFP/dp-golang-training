package constructor

import (
	"encoding/json"

	"github.com/DanoFP/dp-golang-training/pd/03-builder/builder"
	"github.com/DanoFP/dp-golang-training/pd/03-builder/message"
)

type JSONMessageBuilder struct {
	message message.Message
}

func (b *JSONMessageBuilder) SetRecipient(recipient string) builder.MessageBuilder {
	b.message.Recipient = recipient
	return b
}

func (b *JSONMessageBuilder) SetMessage(text string) builder.MessageBuilder {
	b.message.Text = text
	return b
}

func (b *JSONMessageBuilder) Build() (*message.MessageRepresented, error) {
	data, err := json.Marshal(b.message)
	if err != nil {
		return nil, err
	}

	return &message.MessageRepresented{Body: data, Format: "JSON"}, nil
}

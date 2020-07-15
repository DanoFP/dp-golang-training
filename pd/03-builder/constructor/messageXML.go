package constructor

import (
	"encoding/xml"

	"github.com/DanoFP/dp-golang-training/pd/03-builder/builder"
	"github.com/DanoFP/dp-golang-training/pd/03-builder/message"
)

type XMLMessageBuilder struct {
	message message.Message
}

func (b *XMLMessageBuilder) SetRecipient(recipient string) builder.MessageBuilder {
	b.message.Recipient = recipient
	return b
}

func (b *XMLMessageBuilder) SetMessage(text string) builder.MessageBuilder {
	b.message.Text = text
	return b
}

func (b *XMLMessageBuilder) Build() (*message.MessageRepresented, error) {
	data, err := xml.Marshal(b.message)
	if err != nil {
		return nil, err
	}

	return &message.MessageRepresented{Body: data, Format: "XML"}, nil
}

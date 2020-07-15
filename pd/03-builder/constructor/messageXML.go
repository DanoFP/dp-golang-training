package constructor

import (
	"encoding/xml"

	"github.com/DanoFP/dp-golang-training/pd/03-builder/builder"
	"github.com/DanoFP/dp-golang-training/pd/03-builder/message"
)

// XMLMessageBuilder is concrete builder
type XMLMessageBuilder struct {
	message message.Message
}

// SetRecipient asigna el receptor del mensaje
func (b *XMLMessageBuilder) SetRecipient(recipient string) builder.MessageBuilder {
	b.message.Recipient = recipient
	return b
}

// SetMessage asigna el mensaje a enviar
func (b *XMLMessageBuilder) SetMessage(text string) builder.MessageBuilder {
	b.message.Text = text
	return b
}

// Build construye el objeto y lo representa en XML
func (b *XMLMessageBuilder) Build() (*message.MessageRepresented, error) {
	data, err := xml.Marshal(b.message)
	if err != nil {
		return nil, err
	}

	return &message.MessageRepresented{Body: data, Format: "XML"}, nil
}

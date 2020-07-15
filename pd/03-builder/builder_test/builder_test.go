package builder_test

import (
	"testing"

	constructor "github.com/DanoFP/dp-golang-training/pd/03-builder/constructor"
	sender "github.com/DanoFP/dp-golang-training/pd/03-builder/sender"
)

func TestSender_BuildMessage(t *testing.T) {
	json := &constructor.JSONMessageBuilder{}
	xml := &constructor.XMLMessageBuilder{}
	sender := &sender.Sender{}

	sender.SetBuilder(json)
	jsonMsg, err := sender.BuildMessage("Gopher", "Hello World!")
	if err != nil {
		t.Fatalf("BuildMesage(): JSON: something went wrong we received: %v", err)
	}

	t.Log(string(jsonMsg.Body))

	sender.SetBuilder(xml)
	xmlMsg, err := sender.BuildMessage("Gopher", "Hola mundo")
	if err != nil {
		t.Fatalf("BuildMesage(): XML: something went wrong we received: %v", err)
	}

	t.Log(string(xmlMsg.Body))
}

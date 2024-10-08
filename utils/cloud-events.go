package utils

import (
	"encoding/json"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/google/uuid"
	"github.com/gookit/slog"
)

// CreateCloudEvent creates a CloudEvent JSON.
func CreateCloudEvent(t string, source string, data string) (cloudevents.Event, error) {
	slog.Debugf("Creating CloudEvent: %s", t)
	ce := cloudevents.NewEvent()
	id := uuid.New()
	ce.SetID(id.String())
	ce.SetSource(source)
	ce.SetType(t)
	ce.SetData(cloudevents.ApplicationJSON, data)
	return ce, nil
}

// MarshallCloudEvent converts a CloudEvent struct to a JSON byte array.
func MarshallCloudEvent(event cloudevents.Event) ([]byte, error) {
	bytes, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// UnmarshalCloudEvent converts a JSON byte array to a CloudEvent struct.
func UnmarshalCloudEvent(bytes []byte) (cloudevents.Event, error) {
	event := cloudevents.NewEvent()
	err := json.Unmarshal(bytes, &event)
	if err != nil {
		return cloudevents.NewEvent(), err
	}
	return event, nil
}

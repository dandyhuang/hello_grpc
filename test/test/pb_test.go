package test

import (
	"cloud.google.com/go/dialogflow/apiv2/dialogflowpb"
	"fmt"
	"github.com/aj3423/aproto"
	"github.com/golang/protobuf/proto"
	"testing"
)

func TestPb(t *testing.T) {
	request := &dialogflowpb.DetectIntentRequest{
		Session: "session",
		QueryParams: &dialogflowpb.QueryParameters{
			TimeZone: "UTF+8:00",
			WebhookHeaders: map[string]string{
				"key1": "v1",
				"key2": "v2",
			},
		},
		QueryInput:            nil,
		OutputAudioConfig:     nil,
		OutputAudioConfigMask: nil,
		InputAudio:            []byte{0, 1, 2, 3, 4},
	}
	bytes, err := proto.Marshal(request)
	if err != nil {
		panic(err)
	}
	out, err := aproto.TryDumpEx(bytes, &aproto.ConsoleRenderer{})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}

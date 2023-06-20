package main

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"testing"
)

type stubSelectObjectContentEventStreamReader struct {
	StreamEvents <-chan s3.SelectObjectContentEventStreamEvent
	Error        error
}

func (s stubSelectObjectContentEventStreamReader) Events() <-chan s3.SelectObjectContentEventStreamEvent {
	return s.StreamEvents
}

func (s stubSelectObjectContentEventStreamReader) Close() error {
	return nil
}

func (s stubSelectObjectContentEventStreamReader) Err() error {
	return s.Error
}

func TestAwsEventStream(t *testing.T) {

	type testCase struct {
		arg1 s3.SelectObjectContentEventStreamEvent
		want s3.SelectObjectContentEventStreamEvent
	}
	streamEvents := make(chan s3.SelectObjectContentEventStreamEvent, 3)
	eventStreamReader := stubSelectObjectContentEventStreamReader{
		StreamEvents: streamEvents,
		Error:        nil,
	}

	mockEventStream := s3.NewSelectObjectContentEventStream(func(o *s3.SelectObjectContentEventStream) {
		o.Reader = eventStreamReader
		o.StreamCloser = io.NopCloser(bytes.NewReader(nil))
	})

	mockOutput := &s3.SelectObjectContentOutput{
		EventStream: mockEventStream,
	}
	statsEvent := &s3.StatsEvent{
		Details: &s3.Stats{
			BytesScanned:   aws.Int64(1024),
			BytesProcessed: aws.Int64(2048),
		},
	}
	streamEvents <- &s3.RecordsEvent{Payload: []byte("hello world")}
	streamEvents <- &s3.EndEvent{}
	streamEvents <- statsEvent
	close(streamEvents)

	got := awsEventStream(mockOutput)
	want := "hello world"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

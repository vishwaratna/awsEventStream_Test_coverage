package main

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
)

func awsEventStream(resp *s3.SelectObjectContentOutput) string {

	var awsResponse string

	defer resp.EventStream.Close()

	for event := range resp.EventStream.Events() {

		switch event := event.(type) {

		case *s3.RecordsEvent:

			awsResponse = string(event.Payload)

		case *s3.StatsEvent:

			var LOG_PROCESSED_BYTES string
			log.Printf(LOG_PROCESSED_BYTES, *event.Details.BytesProcessed)

		case *s3.EndEvent:

			var LOG_QUERY_COMPLETE any
			log.Println(LOG_QUERY_COMPLETE)

		}

	}

	return awsResponse

}

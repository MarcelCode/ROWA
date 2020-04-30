
package util
//Code for push to AWS CloudWatch Logs
/*
import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"time"
	log"github.com/sirupsen/logrus"
)

func AwsInitLogs() *cloudwatchlogs.CloudWatchLogs {
	

	mySession := session.Must(session.NewSession())

// Create a CloudWatchLogs client from just a session.

		
	Loggersvc := cloudwatchlogs.New(mySession, aws.NewConfig().WithRegion("eu-central-1"))

	TestLog := &cloudwatchlogs.InputLogEvent{
		Message : aws.String("Test Message"),
		Timestamp : aws.Int64(time.Now().Unix()*1000),
		//time.Now().Unix()

	}

	var s []*cloudwatchlogs.InputLogEvent

        s = append(s, TestLog)
    

		log.Warning(s)
	inputLog := &cloudwatchlogs.PutLogEventsInput {
		LogEvents : s,
		LogGroupName   : aws.String("backendlogs"),
		LogStreamName  : aws.String("warnings"),
		SequenceToken : aws.String("49602826257371217146750534988396430289150909492058190914"),


	}

	
	resp, err := Loggersvc.PutLogEvents(inputLog)
	log.Warning(resp)
	if err != nil {
		
		log.Fatal(err)
	}
	
	

	return Loggersvc
}*/
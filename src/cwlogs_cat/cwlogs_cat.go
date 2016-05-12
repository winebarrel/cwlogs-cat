package cwlogs_cat

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"regexp"
	"strings"
	"time"
)

func putLogsEvents(log_group_name string, log_stream_name string, message string) (err error) {
	svc := cloudwatchlogs.New(session.New())

	params := &cloudwatchlogs.PutLogEventsInput{
		LogEvents: []*cloudwatchlogs.InputLogEvent{
			{
				Message:   aws.String(message),
				Timestamp: aws.Int64(time.Now().UnixNano() / int64(time.Millisecond)),
			},
		},
		LogGroupName:  aws.String(log_group_name),
		LogStreamName: aws.String(log_stream_name),
	}

	_, err = svc.PutLogEvents(params)

	if err != nil {
		matched, _ := regexp.MatchString(`^(DataAlreadyAcceptedException|InvalidSequenceTokenException):`, err.Error())
		if matched {
			re := regexp.MustCompile(`\bsequenceToken(?: is)?: (\S+)\b`)
			md := re.FindStringSubmatch(err.Error())
			sequence_token := ""

			if len(md) == 2 {
				sequence_token = md[1]
			}

			if sequence_token != "" {
				params.SequenceToken = aws.String(sequence_token)
				_, err = svc.PutLogEvents(params)
			}
		}
	}

	return
}

func Cat(params *CWLogsCatParams, message string) (err error) {
	message = strings.TrimSpace(message)
	err = putLogsEvents(params.log_group_name, params.log_stream_name, message)
	return
}

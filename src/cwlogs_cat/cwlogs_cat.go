package cwlogs_cat

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"regexp"
	"strings"
	"time"
)

func createLogStream(svc *cloudwatchlogs.CloudWatchLogs, log_group_name string, log_stream_name string) (err error) {
	params := &cloudwatchlogs.CreateLogStreamInput{
		LogGroupName:  aws.String(log_group_name),
		LogStreamName: aws.String(log_stream_name),
	}

	_, err = svc.CreateLogStream(params)

	return
}

func putLogsEvents(svc *cloudwatchlogs.CloudWatchLogs, log_group_name string, log_stream_name string, message string, auto_create_stream bool) (err error) {
	params := &cloudwatchlogs.PutLogEventsInput{
		LogEvents: []*cloudwatchlogs.InputLogEvent{
			{
				Message: aws.String(message),
				Timestamp: aws.Int64(time.Now().UnixNano() / int64(time.Millisecond))
			},
		},
		LogGroupName:  aws.String(log_group_name),
		LogStreamName: aws.String(log_stream_name),
	}

	for i := 0; i < 10; i++ {
		_, err = svc.PutLogEvents(params)

		if err == nil {
			break
		}

		if matched, _ := regexp.MatchString(`^ResourceNotFoundException: The specified log stream does not exist\.`, err.Error()); matched {
			if auto_create_stream {
				err = createLogStream(svc, log_group_name, log_stream_name)
			}

			if err != nil {
				break
			}
		} else if matched, _ := regexp.MatchString(`^(DataAlreadyAcceptedException|InvalidSequenceTokenException):`, err.Error()); matched {
			re := regexp.MustCompile(`\bsequenceToken(?: is)?: (\S+)\b`)
			md := re.FindStringSubmatch(err.Error())

			if len(md) == 2 {
				sequence_token := md[1]
				params.SequenceToken = aws.String(sequence_token)
			} else {
				break
			}
		}

		time.Sleep(100 * time.Millisecond)
	}

	return
}

func Cat(params *CWLogsCatParams, message string) (err error) {
	message = strings.TrimSpace(message)
	svc := cloudwatchlogs.New(session.New())
	err = putLogsEvents(svc, params.log_group_name, params.log_stream_name, message, params.auto_create_stream)
	return
}

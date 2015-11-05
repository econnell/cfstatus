package main

import "os"
import "fmt"
import "strings"
import "flag"
import "github.com/aws/aws-sdk-go/aws"
import "github.com/aws/aws-sdk-go/aws/session"
import "github.com/aws/aws-sdk-go/service/cloudformation"

func main() {
	var region = flag.String("region", "us-west-2", "the region the CloudFormation stack is in")
	var stackName = flag.String("stack-name", "", "the name of the stack to view")

	flag.Parse()

	if *stackName == "" {
		fmt.Println("error: -stack-name parameter is required")
		os.Exit(1)
	}

	svc := cloudformation.New(session.New(), &aws.Config{Region: aws.String(*region)})
	nextToken := ""
	type eventData struct {
		EventId              string
		LogicalResourceId    string
		ResourceStatus       string
		ResourceStatusReason string
		ResourceType         string
	}
	events := make([]eventData, 0)
	eventStatus := make(map[string]eventData)

	for true {
		params := &cloudformation.DescribeStackEventsInput{
			StackName: aws.String(*stackName),
		}
		if nextToken != "" {
			params.NextToken = aws.String(nextToken)
		}
		resp, err := svc.DescribeStackEvents(params)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if resp.NextToken == nil {
			break
		}
		nextToken = *resp.NextToken
		for ctr := 0; ctr < len(resp.StackEvents); ctr++ {
			event := eventData{
				EventId:           *resp.StackEvents[ctr].EventId,
				LogicalResourceId: *resp.StackEvents[ctr].LogicalResourceId,
				ResourceStatus:    *resp.StackEvents[ctr].ResourceStatus,
				ResourceType:      *resp.StackEvents[ctr].ResourceType,
			}
			events = append(events, event)
		}
	}
	for ctr := len(events) - 1; ctr >= 0; ctr-- {
		eventStatus[events[ctr].LogicalResourceId] = events[ctr]
	}
	// probably should do some sorting at this point...

	eventsInProgress := false
	for _, currentEvent := range eventStatus {
		if strings.Contains(currentEvent.ResourceStatus, "_IN_PROGRESS") {
			eventsInProgress = true
			fmt.Printf("%-60.60s %-20.20s %s\n", currentEvent.LogicalResourceId, currentEvent.ResourceStatus, currentEvent.ResourceType)
		}
	}
	if eventsInProgress == false {
		fmt.Println("No events are currently in progress")
	}
}

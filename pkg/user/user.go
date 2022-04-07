package user

import (
	"os"

	"github.com/aws-skd-go/aws/session"
)

func main() {
	region := os.Getenv("AWS_REGION")
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

}

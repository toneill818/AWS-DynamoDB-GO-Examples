/*
	The code below is an example on how to put data into amazon's dynamodb
	using GO. This is the minimum requirements to get data into a table.
	For more information about how to use the package visit this site
	https://github.com/awslabs/aws-sdk-go/blob/master/service/dynamodb/examples_test.go
	Note you need to have your AWS access key and secret access key saved in your home
	directory in a file called credentials inside of the .aws folder. 
*/

package main
import(
	"fmt"
	
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/dynamodb"
)
/*
	Below is an example on how to insert items into a dynamodb table.
	Note: the key "id" is a primary key on the example table.
*/
func main(){
	// Create a dynamodb service and connect to the Region your DB is at
	svc := dynamodb.New(&aws.Config{Region: "us-east-1"})

	// This holds all of the values that will be inserted into the table
	params := &dynamodb.PutItemInput{
		Item: &map[string]*dynamodb.AttributeValue{
			// This is our primary key for the item being inserted
			"id": &dynamodb.AttributeValue{
				S: aws.String("Primary Key"),
			},
			// key1 maps to a number
			"key1": &dynamodb.AttributeValue{
				N: aws.String("12345"),
			},
			// Holds a boolean value
			"boolean": &dynamodb.AttributeValue{
				BOOL: aws.Boolean(false),
			},
			// Holds an array of strings
			"string array": &dynamodb.AttributeValue{
				SS: []*string{
					aws.String("element 1"),
					aws.String("element 2"),
					aws.String("element 3"),
				},
			},
			"Bytes": &dynamodb.AttributeValue{
				B: []byte("message"),
			},
			"NULL?": &dynamodb.AttributeValue{
				NULL: aws.Boolean(true),
			},
		},
		TableName: aws.String("example_table"),
	}

	// This will put Item map that was created above into the table. err is nil if no errors occured.
	resp, err := svc.PutItem(params)

	if awserr := aws.Error(err); awserr != nil {
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		panic(err)
	}
	fmt.Println(awsutil.StringValue(resp))
}
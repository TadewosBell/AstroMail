package smtpstack

import (
	"context"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// MakeAWSS3BucketNameCompliant makes a string compliant with AWS S3 bucket naming rules
func makeAWSS3BucketNameCompliant(input string) string {
	// Convert the input string to lowercase
	input = strings.ToLower(input)

	// Replace invalid characters with "-"
	invalidCharsRegex := regexp.MustCompile(`[^a-z0-9-.]`)
	input = invalidCharsRegex.ReplaceAllString(input, "-")

	// Replace multiple consecutive "-" with single "-"
	input = strings.ReplaceAll(input, "--", "-")

	// input = strings.ReplaceAll(input, ".", "-")

	// Remove leading and trailing "-"
	input = strings.Trim(input, "-")

	// Limit the length to 63 characters
	if len(input) > 63 {
		input = input[:63]
	}

	return input
}

func CreateEmailBucket(domain string) (string, error) {
	bucketName := makeAWSS3BucketNameCompliant(fmt.Sprintf("AstroMail-%s", domain))

	sdkConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("AstroMailApp"), config.WithRegion("us-east-1"))
	if err != nil {
		return "", fmt.Errorf("failed to load AWS config: %w", err)
	}

	fmt.Println(sdkConfig, err, bucketName)

	s3Client := s3.NewFromConfig(sdkConfig)

	createBucketParams := &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	}

	_, err = s3Client.CreateBucket(context.TODO(), createBucketParams)
	if err != nil {
		fmt.Printf("failed to create bucket: %s", err)
		return "", err
	}

	// Add S3 bucket policy
	policy := fmt.Sprintf(`{
        "Version": "2012-10-17",
        "Statement": [
            {
                "Effect": "Allow",
                "Principal": {
                    "Service": "ses.amazonaws.com"
                },
                "Action": "s3:PutObject",
                "Resource": "arn:aws:s3:::%s/*"
            }
        ]
    }`, bucketName)

	_, err = s3Client.PutBucketPolicy(context.TODO(), &s3.PutBucketPolicyInput{
		Bucket: aws.String(bucketName),
		Policy: aws.String(policy),
	})
	if err != nil {
		fmt.Printf("failed to set bucket policy: %s", err)
		return "", err
	}

	// Start a goroutine to check for bucket existence
	resultCh := make(chan string)
	errCh := make(chan error)
	go func() {
		result, err := waitForBucketExistence(s3Client, bucketName)
		if err != nil {
			errCh <- err
			return
		}
		resultCh <- result
	}()

	select {
	case result := <-resultCh:
		return result, nil
	case err := <-errCh:
		return "", err
	case <-time.After(2 * time.Minute):
		fmt.Println("Timeout: Bucket not found")
		return "", fmt.Errorf("bucket not found")
	}
}

func waitForBucketExistence(s3Client *s3.Client, bucketName string) (string, error) {
	ctx := context.TODO()
	timeout := time.After(2 * time.Minute)
	ticker := time.NewTicker(10 * time.Second)

	for {
		select {
		case <-timeout:
			return "", fmt.Errorf("bucket not found")
		case <-ticker.C:
			fmt.Println("Checking if bucket exists...")
			// List buckets
			resp, err := s3Client.ListBuckets(ctx, &s3.ListBucketsInput{})
			fmt.Println(resp)
			if err != nil {
				return "", fmt.Errorf("failed to list buckets: %w", err)
			}
			// Check if the bucket exists in the list
			for _, bucket := range resp.Buckets {
				if *bucket.Name == bucketName {
					fmt.Println("Bucket found!")
					return bucketName, nil
				}
			}
			fmt.Println("Bucket not found yet...")
		}
	}
}

// ReadBucketFolderContent retrieves the 50 most recently updated objects in an S3 bucket folder.
func ReadBucketFolderContent(bucketName, folderName string, pageNum int) ([]string, error) {
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithSharedConfigProfile("AstroMailApp"), config.WithRegion("us-east-1"))
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg)

	input := &s3.ListObjectsV2Input{
		Bucket:     &bucketName,
		Prefix:     &folderName,
		StartAfter: nil,
	}

	if pageNum > 1 {
		input.ContinuationToken = &folderName
	}

	result, err := client.ListObjectsV2(context.Background(), input)
	if err != nil {
		return nil, err
	}

	objects := make([]string, 0, len(result.Contents))
	for _, obj := range result.Contents {
		objects = append(objects, *obj.Key)
	}

	sort.Slice(objects, func(i, j int) bool {
		return result.Contents[i].LastModified.After(*result.Contents[j].LastModified)
	})

	if len(objects) > pageNum*50 {
		return objects[(pageNum-1)*50 : pageNum*50], nil
	}

	return objects[(pageNum-1)*50:], nil
}

// GetObjectContentAsString retrieves the content of an object in an S3 bucket as a string.
func GetObjectContentAsString(bucketName, objectKey string) (string, error) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return "", err
	}

	client := s3.NewFromConfig(cfg)

	input := &s3.GetObjectInput{
		Bucket: &bucketName,
		Key:    &objectKey,
	}

	result, err := client.GetObject(context.Background(), input)
	if err != nil {
		return "", err
	}

	content, err := io.ReadAll(result.Body)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

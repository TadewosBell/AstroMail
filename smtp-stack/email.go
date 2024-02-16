package smtpstack

import (
	storage "AstroMail/config"
	emailparser "AstroMail/email-parser"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

func SendEmail(sender, subject, body string, recipient, ccAddresses []string) error {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("AstroMailApp"), config.WithRegion("us-east-1"))
	if err != nil {
		return fmt.Errorf("failed to load SDK configuration: %v", err)
	}

	client := ses.NewFromConfig(cfg)

	input := &ses.SendEmailInput{
		Destination: &types.Destination{
			CcAddresses: ccAddresses,
			ToAddresses: recipient,
		},
		Message: &types.Message{
			Body: &types.Body{
				Html: &types.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(body),
				},
			},
			Subject: &types.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(sender),
	}

	sendEmailResponse, err := client.SendEmail(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}
	fmt.Println("Send email response: ", sendEmailResponse)

	fmt.Println("Email sent successfully!")
	sentEmailEml := emailparser.CreateEMLString(sender, subject, body, recipient, ccAddresses, *sendEmailResponse.MessageId)
	err = storage.SaveEmail(*sendEmailResponse.MessageId, sentEmailEml, "sent")
	fmt.Println("Message ID:", *sendEmailResponse.MessageId)
	fmt.Println(sentEmailEml)
	return nil
}

func ConfigureSESReceiptRules(domain, username, roleARN, bucket string) error {
	// Load AWS SDK config with default credentials and shared config profile.
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile("AstroMailApp"),
		config.WithRegion("us-east-1"),
	)
	if err != nil {
		return fmt.Errorf("failed to load SDK configuration: %v", err)
	}

	// Create SES client.
	client := ses.NewFromConfig(cfg)

	// Create receipt rule set.
	_, err = client.CreateReceiptRuleSet(context.TODO(), &ses.CreateReceiptRuleSetInput{
		RuleSetName: aws.String("SESForwardingRuleSet"),
	})
	if err != nil {
		return fmt.Errorf("failed to create receipt rule set: %v", err)
	}

	// Create receipt rule to forward emails to S3 bucket.
	_, err = client.CreateReceiptRule(context.TODO(), &ses.CreateReceiptRuleInput{
		Rule: &types.ReceiptRule{
			Actions: []types.ReceiptAction{
				{
					S3Action: &types.S3Action{
						BucketName:      aws.String(bucket),
						ObjectKeyPrefix: aws.String("emails/"),
						TopicArn:        nil,
					},
				},
			},
			Enabled:     true,
			Name:        aws.String("ForwardToS3Rule"),
			ScanEnabled: false,
			Recipients: []string{
				username + "@" + domain, // Forward all emails sent to the specified domain
			},
		},
		RuleSetName: aws.String("SESForwardingRuleSet"),
	})
	if err != nil {
		return fmt.Errorf("failed to create receipt rule: %v", err)
	}

	// Activate the receipt rule set.
	_, err = client.SetActiveReceiptRuleSet(context.TODO(), &ses.SetActiveReceiptRuleSetInput{
		RuleSetName: aws.String("SESForwardingRuleSet"),
	})
	if err != nil {
		return fmt.Errorf("failed to activate receipt rule set: %v", err)
	}

	fmt.Println("SES receipt rules configured and activated successfully.")
	return nil
}

func CreateSESPolicyAndRole(domain, bucket string) (string, error) {
	// Load AWS SDK config with default credentials and shared config profile.
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("AstroMailApp"), config.WithRegion("us-east-1"))
	if err != nil {
		return "", fmt.Errorf("failed to load SDK configuration: %v", err)
	}

	// Generate SES policy document.
	policy := fmt.Sprintf(`{
        "Version": "2012-10-17",
        "Statement": [{
            "Effect": "Allow",
            "Action": "s3:PutObject",
            "Resource": "arn:aws:s3:::%s/*",
            "Condition": {
                "StringEquals": {
                    "aws:Referer": "%s"
                }
            }
        }]
    }`, bucket, domain)

	// Create IAM client.
	iamClient := iam.NewFromConfig(cfg)

	// Check if IAM role exists, if not create it.
	roleName := "SESS3ForwardingRole"
	roleArn := ""
	roleExists := false
	listRolesOutput, err := iamClient.ListRoles(context.TODO(), &iam.ListRolesInput{})
	if err != nil {
		return "", fmt.Errorf("failed to list IAM roles: %v", err)
	}
	for _, role := range listRolesOutput.Roles {
		if aws.ToString(role.RoleName) == roleName {
			roleArn = aws.ToString(role.Arn)
			roleExists = true
			break
		}
	}
	if !roleExists {
		createRoleOutput, err := iamClient.CreateRole(context.TODO(), &iam.CreateRoleInput{
			AssumeRolePolicyDocument: aws.String(`{
                "Version": "2012-10-17",
                "Statement": [{
                    "Effect": "Allow",
                    "Principal": {
                        "Service": "ses.amazonaws.com"
                    },
                    "Action": "sts:AssumeRole"
                }]
            }`),
			Description: aws.String("IAM role for SES to forward emails to S3"),
			RoleName:    aws.String(roleName),
		})
		if err != nil {
			return "", fmt.Errorf("failed to create IAM role: %v", err)
		}
		roleArn = aws.ToString(createRoleOutput.Role.Arn)
	}

	// Attach policy to the IAM role.
	_, err = iamClient.PutRolePolicy(context.TODO(), &iam.PutRolePolicyInput{
		PolicyDocument: aws.String(policy),
		PolicyName:     aws.String("SESS3ForwardingPolicy"),
		RoleName:       aws.String(roleName),
	})
	if err != nil {
		return "", fmt.Errorf("failed to attach policy to IAM role: %v", err)
	}

	// Return the IAM role ARN.
	return roleArn, nil
}

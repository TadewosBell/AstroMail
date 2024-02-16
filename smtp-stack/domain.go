package smtpstack

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ses"
)

func VerifyDomain(domain string) (string, error) {
	// Load AWS SDK config with default credentials and shared config profile.
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("AstroMailApp"), config.WithRegion("us-east-1"))
	if err != nil {
		return "", fmt.Errorf("failed to load SDK configuration: %v", err)
	}

	// Create SES client.
	client := ses.NewFromConfig(cfg)

	// Call VerifyDomainIdentity API to verify the domain identity.
	resp, err := client.VerifyDomainIdentity(context.TODO(), &ses.VerifyDomainIdentityInput{
		Domain: aws.String(domain),
	})
	if err != nil {
		return "", fmt.Errorf("failed to verify domain identity %s: %v", domain, err)
	}

	verificationToken := aws.ToString(resp.VerificationToken)
	if verificationToken == "" {
		return "", fmt.Errorf("verification token not returned for domain %s", domain)
	}

	return verificationToken, nil
}

func IsDomainVerified(domain string) (string, error) {
	// Load AWS SDK config with default credentials and shared config profile.
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("AstroMailApp"), config.WithRegion("us-east-1"))
	if err != nil {
		return "Failed", fmt.Errorf("failed to load SDK configuration: %v", err)
	}

	// Create SES client.
	client := ses.NewFromConfig(cfg)

	// Call GetIdentityVerificationAttributes API to get the verification status.
	resp, err := client.GetIdentityVerificationAttributes(context.TODO(), &ses.GetIdentityVerificationAttributesInput{
		Identities: []string{domain},
	})
	if err != nil {
		return "Failed", fmt.Errorf("failed to get identity verification attributes for domain %s: %v", domain, err)
	}

	// Check if the domain verification status is "Success".
	attributes := resp.VerificationAttributes
	if attributes == nil {
		return "Failed", fmt.Errorf("no verification attributes found for domain %s", domain)
	}
	return string(attributes[domain].VerificationStatus), nil
}

package backend

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ses"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func VerifyDomain(accessKey string, secretKey string, domain string) string {
	var verificationToken string

	pulumi.Run(func(ctx *pulumi.Context) error {
		// Configure the AWS provider with the credentials loaded from the CSV.
		awsProvider, err := aws.NewProvider(ctx, "aws-provider", &aws.ProviderArgs{
			AccessKey: pulumi.StringPtr(accessKey),
			SecretKey: pulumi.StringPtr(secretKey),
		})
		if err != nil {
			return err
		}

		// Use the configured AWS provider for your resources.
		// Request verification for your domain with SES.
		domainIdentity, err := ses.NewDomainIdentity(ctx, "EmailDomainEntity", &ses.DomainIdentityArgs{
			Domain: pulumi.String(domain),
		}, pulumi.Provider(awsProvider))
		if err != nil {
			return err
		}

		if err != nil {
			return err
		}

		// Use Apply to access the VerificationToken.
		domainIdentity.VerificationToken.ApplyT(func(verificationTokenTemp string) (string, error) {
			fmt.Println("Verification Token:", verificationTokenTemp)
			verificationToken = verificationTokenTemp
			// You can do additional processing here if needed
			return verificationTokenTemp, nil
		})

		// Output the verification token that you'll need to add to your domain's DNS settings.
		ctx.Export("verificationToken", domainIdentity.VerificationToken)
		return nil
	})
	return verificationToken
}

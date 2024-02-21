# README

## About

Astromail is an open source serverless app that allows you to send emails and recieve them for your domain using AWS. It automatically deploys its own stack and creates the connections needed for the domain of your choice. 

It uses Aws-sdk-v2 to create the stack and wails for the email client

#### BE CAREFUL BECAUSE THIS WILL CREATE RESOURCES IN AWS AND IT WILL CREATE AN ACTIVE SES RULE. YOU CAN ONLY HAVE ONE OF THOSE ACTIVE SO IF YOU HAVE A CUSTOM ONE ALREADY YOU NEED TO MAKE SOME MORE CONSIDERATIONS

## IAM Permissions

For this app to work you need to create an aws user with the following permissions and add the key/secret in to Astromail during the setup phase
```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "ses:VerifyDomainIdentity",
                "ses:GetIdentityVerificationAttributes",
                "ses:SetIdentityMailFromDomain",
                "ses:SendEmail",
                "ses:CreateReceiptRuleSet",
                "ses:CreateReceiptRule",
                "ses:PutIdentityPolicy",
                "ses:SetActiveReceiptRuleSet"
            ],
            "Resource": "*"
        },
        {
            "Effect": "Allow",
            "Action": [
                "s3:CreateBucket",
                "s3:GetBucketLocation",
                "s3:ListAllMyBuckets",
                "s3:PutBucketPolicy",
                "s3:ListBucket",
                "s3:GetBucketLocation",
                "s3:GetObject"
            ],
            "Resource": "arn:aws:s3:::astromail-*"
        },
        {
            "Effect": "Allow",
            "Action": [
                "s3:PutBucketTagging",
                "s3:GetBucketTagging",
                "s3:PutEncryptionConfiguration",
                "s3:GetEncryptionConfiguration"
            ],
            "Resource": "arn:aws:s3:::astromail-*"
        },
        {
            "Effect": "Allow",
            "Action": [
                "iam:ListRoles",
                "iam:CreateRole",
                "iam:PutRolePolicy"
            ],
            "Resource": "*"
        }
    ]
}
```

## Follow the development here

https://medium.com/@tadewoswebkreator/follow-me-as-i-develop-an-open-source-email-client-for-hackers-called-astromail-eefc17039f07

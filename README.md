# CloudFront API Code Examples

## Configuration

### Create an IAM user

* Use Root or IAM user with Administrator policy to create an IAM user with below policy. This user will be used for
  * Create IAM User with AKSK
  	* AWS Console -> Service: IAM -> Users -> Add users 
  	* Check Access key - Programmatic access -> Permissions -> Attach exiting policies -> Create policy
  	* JSON -> copy and paste below policy -> Next... -> Name -> Create
  	* Press sync on Add user -> select above policy -> Next... -> Create
  	* `Download the credentials file!`
  * Rotate AKSK: https://aws.amazon.com/blogs/security/how-to-rotate-access-keys-for-iam-users/
  * Take CloudFront actions on Console
  * Call demo APIs

```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "ManageOwnAccessKeys",
            "Effect": "Allow",
            "Action": [
                "iam:CreateAccessKey",
                "iam:DeleteAccessKey",
                "iam:GetAccessKeyLastUsed",
                "iam:GetUser",
                "iam:ListAccessKeys",
                "iam:UpdateAccessKey"
            ],
            "Resource": "arn:aws:iam::*:user/${aws:username}"
        },
        {
            "Sid": "EnableCloudFront",
            "Effect": "Allow",
            "Action": [
            	"cloudfront:*",
                "cloudwatch:*",
                "s3:Get*",
                "s3:List*"
            ],
            "Resource": "*"
        }
    ]
}
```
### Configure the profile
 
Open the downloaded file an put in access key, secret key and region ("us-east-1")

```
aws configure --profile api-demo

AWS Access Key ID [None]: *******************A
AWS Secret Access Key [None]: *******************a
Default region name [None]: us-east-1
Default output format [None]:
```

## Create Distribution

## Install aws-sdk-go-v2

```
go mod init cloudfront-gov2

go get github.com/aws/aws-sdk-go-v2/aws
go get github.com/aws/aws-sdk-go-v2/config
go get github.com/aws/aws-sdk-go-v2/service/cloudfront
```
* If go get timeout

```
go env -w GOPROXY=https://goproxy.cn
```

## Commands

### CloudFront APIs

* ListDistributions, filtered by Alias

```
go run . -action search-distribution-by-alias -alias hello
```

* GetDistributionConfig

```
go run . -action get-distribution-config -dist $DISTRIBUTION_ID
```

* GetCachePolicyConfig

```
go run . -action get-cache-policy-id -cache $CACHE_POLICY_ID
```

* CreateInvitation

```
go run . -action create-invalidation -dist $DISTRIBUTION_ID -paths "/*"
```

* GetInvalidation

```
go run . -action get-invalidation -dist $DISTRIBUTION_ID -invalidation $INVALIDATION_ID
```
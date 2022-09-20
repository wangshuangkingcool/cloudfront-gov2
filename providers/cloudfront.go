package providers

import (
	"context"
    "strings"
    "time"

    "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
)

const (
    MaxBatchSize int32 = 1000
)

type CloudFrontProvider struct {
	Client *cloudfront.Client
}

func NewCloudFrontProvider() *CloudFrontProvider {
    // cfg := GetConfigFromCreds()
    cfg := GetConfigForProfile("api-demo")
 	return &CloudFrontProvider{
		Client: cloudfront.NewFromConfig(cfg),
	}
}

func (cf *CloudFrontProvider) ListDistributions(c context.Context, 
	input *cloudfront.ListDistributionsInput) ([]types.DistributionSummary, error) {
	result, err := cf.Client.ListDistributions(c, input)
    if err != nil {
        return nil, err
    }

    return result.DistributionList.Items, nil
}


// This may be unnecessary if the max-items is unlimited
func (cf *CloudFrontProvider) ListAllDistributions(c context.Context, 
    input *cloudfront.ListDistributionsInput) ([]types.DistributionSummary, error) {

    var distributions []types.DistributionSummary
    for {
        result, err := cf.Client.ListDistributions(c, input)
        if err != nil {
            return []types.DistributionSummary{}, err
        }

        distributionList := result.DistributionList
        distributions = append(distributions, distributionList.Items...)

        if *distributionList.IsTruncated {
            break
        }

        input.Marker = distributionList.NextMarker
    }

    return distributions, nil
} 

func (cf *CloudFrontProvider) GetDistributionByAlias(c context.Context, 
    alias string) (*types.DistributionSummary, error) {

    input := &cloudfront.ListDistributionsInput{}
    distributions, err := cf.ListDistributions(c, input)
    if err != nil {
        return nil, err
    }

    for _, distribution := range distributions {
        aliases := distribution.Aliases.Items
        for _, found := range aliases {
            if strings.Compare(alias, found) == 0 {
                return &distribution, nil
            }
        }
    }

    return nil, nil
}

func (cf *CloudFrontProvider) GetDistributionConfig(c context.Context, 
    distributionId string) (*cloudfront.GetDistributionConfigOutput, error) {
    input := &cloudfront.GetDistributionConfigInput {Id: aws.String(distributionId)}
    return cf.Client.GetDistributionConfig(c, input)
}

func (cf *CloudFrontProvider) GetCachePolicyConfig(c context.Context, 
    cachePolicyId string) (*cloudfront.GetCachePolicyConfigOutput, error) {
    input := &cloudfront.GetCachePolicyConfigInput {Id: aws.String(cachePolicyId)}
    return cf.Client.GetCachePolicyConfig(c, input)
}

func (cf *CloudFrontProvider) CreateInvalidation(c context.Context, 
    distributionId string, paths []string) (*cloudfront.CreateInvalidationOutput, error) {
    input := &cloudfront.CreateInvalidationInput {
        DistributionId: aws.String(distributionId),
        InvalidationBatch: &types.InvalidationBatch {
            Paths: &types.Paths{
                Quantity: aws.Int32(int32(len(paths))),
                Items: paths,
            },
            CallerReference: aws.String(time.Now().String()),
        },
    }
    return cf.Client.CreateInvalidation(c, input)
}

func (cf *CloudFrontProvider) GetInvalidation(c context.Context, 
    distributionId string, invalidationId string) (*cloudfront.GetInvalidationOutput, error) {
    input := &cloudfront.GetInvalidationInput {
        DistributionId: aws.String(distributionId),
        Id: aws.String(invalidationId),
    }
    return cf.Client.GetInvalidation(c, input)
}


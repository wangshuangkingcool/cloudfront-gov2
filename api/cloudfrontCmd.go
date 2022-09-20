package api

import (
	"context"
	"log"
	"fmt"
	"flag"
	"encoding/json"
    "strings"

	"cdn-demo/providers"
)

var (
    alias string
    distributionId string
    cachePolicyId string
    invalidationId string
    paths string
)

func init() {
    flag.StringVar(&alias, "alias", "", "One `Alias` of the Distribution")
    flag.StringVar(&distributionId, "dist", "", "The `Id` of the Distribution")
    flag.StringVar(&cachePolicyId, "cache", "", "The `Id` of the CachePolicy")
    flag.StringVar(&invalidationId, "invalidation", "", "The `Id` of the Invalidation")
    flag.StringVar(&paths, "paths", "", "The `paths`, separated by comma, for which to remove the cache")
}

func SearchDistributionByAlias() {
	flag.Parse()

    if len(alias) == 0 {
        log.Fatalf("You must supply the Alias")
    }

    client := providers.NewCloudFrontProvider();
	result, err := client.GetDistributionByAlias(context.TODO(), alias)
    if err != nil {
        log.Fatalf("Could not search distribution associated with the alias: %v", err)
    }

    output, _ := json.MarshalIndent(result, "", "\t")
    fmt.Println(string(output))
}

func GetDistributionConfig() {
    flag.Parse()

    if len(distributionId) == 0 {
        log.Fatalf("You must supply the Distribution Id")
    }

    client := providers.NewCloudFrontProvider();
    result, err := client.GetDistributionConfig(context.TODO(), distributionId)
    if err != nil {
        log.Fatalf("Could not get Distribution config by the Id: %v", err)
    }

    output, _ := json.MarshalIndent(result, "", "\t")
    fmt.Println(string(output))
}

func GetCachePolicyConfig() {
    flag.Parse()

    if len(cachePolicyId) == 0 {
        log.Fatalf("You must supply the CachePolicy Id")
    }

    client := providers.NewCloudFrontProvider();
    result, err := client.GetCachePolicyConfig(context.TODO(), cachePolicyId)
    if err != nil {
        log.Fatalf("Could not get CachePolicy config by the Id: %v", err)
    }

    output, _ := json.MarshalIndent(result, "", "\t")
    fmt.Println(string(output))
}

func CreateInvalidation() {
    flag.Parse()

    if len(distributionId) == 0 {
        log.Fatalf("You must supply the Distribution Id")
    }

    client := providers.NewCloudFrontProvider();
    result, err := client.CreateInvalidation(context.TODO(), distributionId, strings.Split(paths, ","))
    if err != nil {
        log.Fatalf("Could not create invalidation: %v", err)
    }

    output, _ := json.MarshalIndent(result, "", "\t")
    fmt.Println(string(output))

}

func GetInvalidation() {
        flag.Parse()

    if len(distributionId) == 0 || len(invalidationId) == 0 {
        log.Fatalf("You must supply the Distribution Id and the Invitation Id")
    }

    client := providers.NewCloudFrontProvider();
    result, err := client.GetInvalidation(context.TODO(), distributionId, invalidationId)
    if err != nil {
        log.Fatalf("Could not get invalidation: %v", err)
    }

    output, _ := json.MarshalIndent(result, "", "\t")
    fmt.Println(string(output))
}
package main

import (
	"log"
	"flag"

	"cdn-demo/api"
)

func main() {
	action := flag.String("action", "", "The action to run")
	flag.Parse()

	if *action == "" {
		log.Fatalf("You must supply the Action")
	}

	switch *action {
	case "get-bandwidth-metric":
		api.GetBandwidthMetric()
	case "get-origin-bandwidth-metric":
		api.GetOriginBandwidthMetric()
	case "search-distribution-by-alias":
		api.SearchDistributionByAlias()
	case "get-distribution-config":
		api.GetDistributionConfig()
	case "get-cache-policy-id":
		api.GetCachePolicyConfig()
	case "create-invalidation":
		api.CreateInvalidation()
	case "get-invalidation":
		api.GetInvalidation()
	case "describe-distirbution-certificate":
		api.DescribeCertificateForDistribution()
	case "describe-alias-certificate":
		api.DescribeCertificateByAlias()
	case "get-metrics":
		api.GetMonitoringMetrics()
	case "verify-ip":
		api.VerifyIp()
	case "prewarm":
		api.StartPrewarm()
	case "prewarm-status":
		api.GetPrewarmStatus()
	case "get-log-urls":
		api.GetLogUrls()
	default:
		log.Fatalf("No action matched")
	}
}


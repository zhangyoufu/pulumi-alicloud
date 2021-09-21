// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package arms

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Arms Prometheus Alert Rules of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.136.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/arms"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		ids, err := arms.GetPrometheusAlertRules(ctx, &arms.GetPrometheusAlertRulesArgs{
// 			ClusterId: "example_value",
// 			Ids: []string{
// 				"example_value-1",
// 				"example_value-2",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("armsPrometheusAlertRuleId1", ids.Rules[0].Id)
// 		opt0 := "^my-PrometheusAlertRule"
// 		nameRegex, err := arms.GetPrometheusAlertRules(ctx, &arms.GetPrometheusAlertRulesArgs{
// 			ClusterId: "example_value",
// 			NameRegex: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("armsPrometheusAlertRuleId2", nameRegex.Rules[0].Id)
// 		return nil
// 	})
// }
// ```
func GetPrometheusAlertRules(ctx *pulumi.Context, args *GetPrometheusAlertRulesArgs, opts ...pulumi.InvokeOption) (*GetPrometheusAlertRulesResult, error) {
	var rv GetPrometheusAlertRulesResult
	err := ctx.Invoke("alicloud:arms/getPrometheusAlertRules:getPrometheusAlertRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPrometheusAlertRules.
type GetPrometheusAlertRulesArgs struct {
	// The ID of the cluster.
	ClusterId string `pulumi:"clusterId"`
	// A list of Prometheus Alert Rule IDs.
	Ids              []string `pulumi:"ids"`
	MatchExpressions *string  `pulumi:"matchExpressions"`
	// A regex string to filter results by Prometheus Alert Rule name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The status of the resource. Valid values: `0`, `1`.
	// * `1`: open.
	// * `0`: off.
	Status *int `pulumi:"status"`
	// The type of the alert rule.
	Type *string `pulumi:"type"`
}

// A collection of values returned by getPrometheusAlertRules.
type GetPrometheusAlertRulesResult struct {
	ClusterId string `pulumi:"clusterId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string                        `pulumi:"id"`
	Ids              []string                      `pulumi:"ids"`
	MatchExpressions *string                       `pulumi:"matchExpressions"`
	NameRegex        *string                       `pulumi:"nameRegex"`
	Names            []string                      `pulumi:"names"`
	OutputFile       *string                       `pulumi:"outputFile"`
	Rules            []GetPrometheusAlertRulesRule `pulumi:"rules"`
	Status           *int                          `pulumi:"status"`
	Type             *string                       `pulumi:"type"`
}

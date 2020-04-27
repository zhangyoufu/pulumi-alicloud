// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pvtz

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides Private Zone Records resource information owned by an Alibaba Cloud account.
func GetZoneRecords(ctx *pulumi.Context, args *GetZoneRecordsArgs, opts ...pulumi.InvokeOption) (*GetZoneRecordsResult, error) {
	var rv GetZoneRecordsResult
	err := ctx.Invoke("alicloud:pvtz/getZoneRecords:getZoneRecords", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getZoneRecords.
type GetZoneRecordsArgs struct {
	// A list of Private Zone Record IDs.
	Ids []string `pulumi:"ids"`
	// Keyword for record rr and value.
	Keyword    *string `pulumi:"keyword"`
	OutputFile *string `pulumi:"outputFile"`
	// ID of the Private Zone.
	ZoneId string `pulumi:"zoneId"`
}

// A collection of values returned by getZoneRecords.
type GetZoneRecordsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of Private Zone Record IDs.
	Ids        []string `pulumi:"ids"`
	Keyword    *string  `pulumi:"keyword"`
	OutputFile *string  `pulumi:"outputFile"`
	// A list of zone records. Each element contains the following attributes:
	Records []GetZoneRecordsRecord `pulumi:"records"`
	ZoneId  string                 `pulumi:"zoneId"`
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source operation to query the instance types that are available to specific instances of Alibaba Cloud.
//
// > **NOTE:** Available in v1.196.0+
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/rds"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			resources, err := rds.GetInstanceClassInfos(ctx, &rds.GetInstanceClassInfosArgs{
//				CommodityCode: "bards",
//				OrderType:     "BUY",
//				OutputFile:    pulumi.StringRef("./classes.txt"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstDbInstanceClass", resources.Infos[0])
//			return nil
//		})
//	}
//
// ```
func GetInstanceClassInfos(ctx *pulumi.Context, args *GetInstanceClassInfosArgs, opts ...pulumi.InvokeOption) (*GetInstanceClassInfosResult, error) {
	var rv GetInstanceClassInfosResult
	err := ctx.Invoke("alicloud:rds/getInstanceClassInfos:getInstanceClassInfos", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceClassInfos.
type GetInstanceClassInfosArgs struct {
	// The commodity code of the instance. Valid values:
	// * **bards**: The instance is a pay-as-you-go primary instance. This value is available on the China site (aliyun.com).
	// * **rds**: The instance is a subscription primary instance. This value is available on the China site (aliyun.com).
	// * **rords**: The instance is a pay-as-you-go read-only instance. This value is available on the China site (aliyun.com).
	// * **rds_rordspre_public_cn**: The instance is a subscription read-only instance. This value is available on the China site (aliyun.com).
	// * **bards_intl**: The instance is a pay-as-you-go primary instance. This value is available on the International site (alibabacloud.com).
	// * **rds_intl**: The instance is a subscription primary instance. This value is available on the International site (alibabacloud.com).
	// * **rords_intl**: The instance is a pay-as-you-go read-only instance. This value is available on the International site (alibabacloud.com).
	// * **rds_rordspre_public_intl**: The instance is a subscription read-only instance. This value is available on the International site (alibabacloud.com).
	CommodityCode string `pulumi:"commodityCode"`
	// The ID of the primary instance.
	DbInstanceId *string `pulumi:"dbInstanceId"`
	// A list of Rds available resource. Each element contains the following attributes:
	Infos []GetInstanceClassInfosInfo `pulumi:"infos"`
	// FThe type of order that you want to query. Valid values:
	// * **BUY**: specifies the query orders that are used to purchase instances.
	// * **UPGRADE**: specifies the query orders that are used to change the specifications of instances.
	// * **RENEW**: specifies the query orders that are used to renew instances.
	// * **CONVERT**: specifies the query orders that are used to change the billing methods of instances.
	OrderType string `pulumi:"orderType"`
	// File name where to save data source results (after running `pulumi up`).
	//
	// > **NOTE**: If you use the CommodityCode parameter to query the instance types that are available to read-only instances, you must specify the DBInstanceId parameter.
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getInstanceClassInfos.
type GetInstanceClassInfosResult struct {
	CommodityCode string  `pulumi:"commodityCode"`
	DbInstanceId  *string `pulumi:"dbInstanceId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of Rds instance class codes.
	Ids []string `pulumi:"ids"`
	// A list of Rds available resource. Each element contains the following attributes:
	Infos      []GetInstanceClassInfosInfo `pulumi:"infos"`
	OrderType  string                      `pulumi:"orderType"`
	OutputFile *string                     `pulumi:"outputFile"`
}

func GetInstanceClassInfosOutput(ctx *pulumi.Context, args GetInstanceClassInfosOutputArgs, opts ...pulumi.InvokeOption) GetInstanceClassInfosResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInstanceClassInfosResult, error) {
			args := v.(GetInstanceClassInfosArgs)
			r, err := GetInstanceClassInfos(ctx, &args, opts...)
			var s GetInstanceClassInfosResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetInstanceClassInfosResultOutput)
}

// A collection of arguments for invoking getInstanceClassInfos.
type GetInstanceClassInfosOutputArgs struct {
	// The commodity code of the instance. Valid values:
	// * **bards**: The instance is a pay-as-you-go primary instance. This value is available on the China site (aliyun.com).
	// * **rds**: The instance is a subscription primary instance. This value is available on the China site (aliyun.com).
	// * **rords**: The instance is a pay-as-you-go read-only instance. This value is available on the China site (aliyun.com).
	// * **rds_rordspre_public_cn**: The instance is a subscription read-only instance. This value is available on the China site (aliyun.com).
	// * **bards_intl**: The instance is a pay-as-you-go primary instance. This value is available on the International site (alibabacloud.com).
	// * **rds_intl**: The instance is a subscription primary instance. This value is available on the International site (alibabacloud.com).
	// * **rords_intl**: The instance is a pay-as-you-go read-only instance. This value is available on the International site (alibabacloud.com).
	// * **rds_rordspre_public_intl**: The instance is a subscription read-only instance. This value is available on the International site (alibabacloud.com).
	CommodityCode pulumi.StringInput `pulumi:"commodityCode"`
	// The ID of the primary instance.
	DbInstanceId pulumi.StringPtrInput `pulumi:"dbInstanceId"`
	// A list of Rds available resource. Each element contains the following attributes:
	Infos GetInstanceClassInfosInfoArrayInput `pulumi:"infos"`
	// FThe type of order that you want to query. Valid values:
	// * **BUY**: specifies the query orders that are used to purchase instances.
	// * **UPGRADE**: specifies the query orders that are used to change the specifications of instances.
	// * **RENEW**: specifies the query orders that are used to renew instances.
	// * **CONVERT**: specifies the query orders that are used to change the billing methods of instances.
	OrderType pulumi.StringInput `pulumi:"orderType"`
	// File name where to save data source results (after running `pulumi up`).
	//
	// > **NOTE**: If you use the CommodityCode parameter to query the instance types that are available to read-only instances, you must specify the DBInstanceId parameter.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetInstanceClassInfosOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceClassInfosArgs)(nil)).Elem()
}

// A collection of values returned by getInstanceClassInfos.
type GetInstanceClassInfosResultOutput struct{ *pulumi.OutputState }

func (GetInstanceClassInfosResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceClassInfosResult)(nil)).Elem()
}

func (o GetInstanceClassInfosResultOutput) ToGetInstanceClassInfosResultOutput() GetInstanceClassInfosResultOutput {
	return o
}

func (o GetInstanceClassInfosResultOutput) ToGetInstanceClassInfosResultOutputWithContext(ctx context.Context) GetInstanceClassInfosResultOutput {
	return o
}

func (o GetInstanceClassInfosResultOutput) CommodityCode() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceClassInfosResult) string { return v.CommodityCode }).(pulumi.StringOutput)
}

func (o GetInstanceClassInfosResultOutput) DbInstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceClassInfosResult) *string { return v.DbInstanceId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstanceClassInfosResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceClassInfosResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of Rds instance class codes.
func (o GetInstanceClassInfosResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInstanceClassInfosResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// A list of Rds available resource. Each element contains the following attributes:
func (o GetInstanceClassInfosResultOutput) Infos() GetInstanceClassInfosInfoArrayOutput {
	return o.ApplyT(func(v GetInstanceClassInfosResult) []GetInstanceClassInfosInfo { return v.Infos }).(GetInstanceClassInfosInfoArrayOutput)
}

func (o GetInstanceClassInfosResultOutput) OrderType() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceClassInfosResult) string { return v.OrderType }).(pulumi.StringOutput)
}

func (o GetInstanceClassInfosResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceClassInfosResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstanceClassInfosResultOutput{})
}

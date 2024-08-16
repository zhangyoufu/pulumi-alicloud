// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `rds.getInstances` data source provides a collection of RDS instances available in Alibaba Cloud account.
// Filters support regular expression for the instance name, searches by tags, and other filters which are listed below.
//
// > **NOTE:** Available since v1.7.0+
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
//			dbInstancesDs, err := rds.GetInstances(ctx, &rds.GetInstancesArgs{
//				NameRegex: pulumi.StringRef("data-\\d+"),
//				Status:    pulumi.StringRef("Running"),
//				Tags: map[string]interface{}{
//					"type": "database",
//					"size": "tiny",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstDbInstanceId", dbInstancesDs.Instances[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetInstances(ctx *pulumi.Context, args *GetInstancesArgs, opts ...pulumi.InvokeOption) (*GetInstancesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetInstancesResult
	err := ctx.Invoke("alicloud:rds/getInstances:getInstances", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstances.
type GetInstancesArgs struct {
	// `Standard` for standard access mode and `Safe` for high security access mode.
	ConnectionMode *string `pulumi:"connectionMode"`
	// `Primary` for primary instance, `Readonly` for read-only instance, `Guard` for disaster recovery instance, and `Temp` for temporary instance.
	DbType *string `pulumi:"dbType"`
	// Default to `false`. Set it to `true` can output parameter template about resource attributes.
	EnableDetails *bool `pulumi:"enableDetails"`
	// Database type. Options are `MySQL`, `SQLServer`, `PostgreSQL`, `MariaDB`. If no value is specified, all types are returned.
	Engine *string `pulumi:"engine"`
	// A list of RDS instance IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by instance name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	PageNumber *int    `pulumi:"pageNumber"`
	PageSize   *int    `pulumi:"pageSize"`
	// Status of the instance.
	Status *string `pulumi:"status"`
	// A map of tags assigned to the DB instances.
	// Note: Before 1.60.0, the value's format is a `json` string which including `TagKey` and `TagValue`. `TagKey` cannot be null, and `TagValue` can be empty. Format example `"{\"key1\":\"value1\"}"`
	Tags map[string]string `pulumi:"tags"`
	// Used to retrieve instances belong to specified VPC.
	VpcId *string `pulumi:"vpcId"`
	// Used to retrieve instances belong to specified `vswitch` resources.
	VswitchId *string `pulumi:"vswitchId"`
}

// A collection of values returned by getInstances.
type GetInstancesResult struct {
	// `Standard` for standard access mode and `Safe` for high security access mode.
	ConnectionMode *string `pulumi:"connectionMode"`
	// `Primary` for primary instance, `Readonly` for read-only instance, `Guard` for disaster recovery instance, and `Temp` for temporary instance.
	DbType        *string `pulumi:"dbType"`
	EnableDetails *bool   `pulumi:"enableDetails"`
	// Database type. Options are `MySQL`, `SQLServer`, `PostgreSQL`, `MariaDB`. If no value is specified, all types are returned.
	Engine *string `pulumi:"engine"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of RDS instance IDs.
	Ids []string `pulumi:"ids"`
	// A list of RDS instances. Each element contains the following attributes:
	Instances []GetInstancesInstance `pulumi:"instances"`
	NameRegex *string                `pulumi:"nameRegex"`
	// A list of RDS instance names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	PageNumber *int     `pulumi:"pageNumber"`
	PageSize   *int     `pulumi:"pageSize"`
	// Status of the instance.
	Status     *string           `pulumi:"status"`
	Tags       map[string]string `pulumi:"tags"`
	TotalCount int               `pulumi:"totalCount"`
	// ID of the VPC the instance belongs to.
	VpcId *string `pulumi:"vpcId"`
	// ID of the vSwitch the instance belongs to.
	VswitchId *string `pulumi:"vswitchId"`
}

func GetInstancesOutput(ctx *pulumi.Context, args GetInstancesOutputArgs, opts ...pulumi.InvokeOption) GetInstancesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInstancesResult, error) {
			args := v.(GetInstancesArgs)
			r, err := GetInstances(ctx, &args, opts...)
			var s GetInstancesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetInstancesResultOutput)
}

// A collection of arguments for invoking getInstances.
type GetInstancesOutputArgs struct {
	// `Standard` for standard access mode and `Safe` for high security access mode.
	ConnectionMode pulumi.StringPtrInput `pulumi:"connectionMode"`
	// `Primary` for primary instance, `Readonly` for read-only instance, `Guard` for disaster recovery instance, and `Temp` for temporary instance.
	DbType pulumi.StringPtrInput `pulumi:"dbType"`
	// Default to `false`. Set it to `true` can output parameter template about resource attributes.
	EnableDetails pulumi.BoolPtrInput `pulumi:"enableDetails"`
	// Database type. Options are `MySQL`, `SQLServer`, `PostgreSQL`, `MariaDB`. If no value is specified, all types are returned.
	Engine pulumi.StringPtrInput `pulumi:"engine"`
	// A list of RDS instance IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by instance name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	PageNumber pulumi.IntPtrInput    `pulumi:"pageNumber"`
	PageSize   pulumi.IntPtrInput    `pulumi:"pageSize"`
	// Status of the instance.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// A map of tags assigned to the DB instances.
	// Note: Before 1.60.0, the value's format is a `json` string which including `TagKey` and `TagValue`. `TagKey` cannot be null, and `TagValue` can be empty. Format example `"{\"key1\":\"value1\"}"`
	Tags pulumi.StringMapInput `pulumi:"tags"`
	// Used to retrieve instances belong to specified VPC.
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
	// Used to retrieve instances belong to specified `vswitch` resources.
	VswitchId pulumi.StringPtrInput `pulumi:"vswitchId"`
}

func (GetInstancesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstancesArgs)(nil)).Elem()
}

// A collection of values returned by getInstances.
type GetInstancesResultOutput struct{ *pulumi.OutputState }

func (GetInstancesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstancesResult)(nil)).Elem()
}

func (o GetInstancesResultOutput) ToGetInstancesResultOutput() GetInstancesResultOutput {
	return o
}

func (o GetInstancesResultOutput) ToGetInstancesResultOutputWithContext(ctx context.Context) GetInstancesResultOutput {
	return o
}

// `Standard` for standard access mode and `Safe` for high security access mode.
func (o GetInstancesResultOutput) ConnectionMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.ConnectionMode }).(pulumi.StringPtrOutput)
}

// `Primary` for primary instance, `Readonly` for read-only instance, `Guard` for disaster recovery instance, and `Temp` for temporary instance.
func (o GetInstancesResultOutput) DbType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.DbType }).(pulumi.StringPtrOutput)
}

func (o GetInstancesResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

// Database type. Options are `MySQL`, `SQLServer`, `PostgreSQL`, `MariaDB`. If no value is specified, all types are returned.
func (o GetInstancesResultOutput) Engine() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.Engine }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstancesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstancesResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of RDS instance IDs.
func (o GetInstancesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInstancesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// A list of RDS instances. Each element contains the following attributes:
func (o GetInstancesResultOutput) Instances() GetInstancesInstanceArrayOutput {
	return o.ApplyT(func(v GetInstancesResult) []GetInstancesInstance { return v.Instances }).(GetInstancesInstanceArrayOutput)
}

func (o GetInstancesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of RDS instance names.
func (o GetInstancesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInstancesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetInstancesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetInstancesResultOutput) PageNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *int { return v.PageNumber }).(pulumi.IntPtrOutput)
}

func (o GetInstancesResultOutput) PageSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *int { return v.PageSize }).(pulumi.IntPtrOutput)
}

// Status of the instance.
func (o GetInstancesResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o GetInstancesResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetInstancesResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GetInstancesResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetInstancesResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

// ID of the VPC the instance belongs to.
func (o GetInstancesResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

// ID of the vSwitch the instance belongs to.
func (o GetInstancesResultOutput) VswitchId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.VswitchId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstancesResultOutput{})
}

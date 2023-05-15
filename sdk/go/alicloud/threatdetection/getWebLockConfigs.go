// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package threatdetection

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides Threat Detection Web Lock Config available to the user.[What is Web Lock Config](https://www.alibabacloud.com/help/en/security-center/latest/api-doc-sas-2018-12-03-api-doc-modifyweblockstart)
//
// > **NOTE:** Available in 1.195.0+
func GetWebLockConfigs(ctx *pulumi.Context, args *GetWebLockConfigsArgs, opts ...pulumi.InvokeOption) (*GetWebLockConfigsResult, error) {
	var rv GetWebLockConfigsResult
	err := ctx.Invoke("alicloud:threatdetection/getWebLockConfigs:getWebLockConfigs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getWebLockConfigs.
type GetWebLockConfigsArgs struct {
	// A list of Web Lock Config IDs.
	Ids []string `pulumi:"ids"`
	// The language of the content within the request and the response. Valid values: `zh`, `en`.
	Lang *string `pulumi:"lang"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	PageNumber *int    `pulumi:"pageNumber"`
	PageSize   *int    `pulumi:"pageSize"`
	// The string that allows you to search for servers in fuzzy match mode. You can enter a server name or IP address.
	Remark *string `pulumi:"remark"`
	// The source IP address of the request.
	SourceIp *string `pulumi:"sourceIp"`
	// The protection status of the server that you want to query. Valid values: `on`, `off`.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getWebLockConfigs.
type GetWebLockConfigsResult struct {
	// A list of Web Lock Config Entries. Each element contains the following attributes:
	Configs []GetWebLockConfigsConfig `pulumi:"configs"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of Web Lock Config IDs.
	Ids        []string `pulumi:"ids"`
	Lang       *string  `pulumi:"lang"`
	OutputFile *string  `pulumi:"outputFile"`
	PageNumber *int     `pulumi:"pageNumber"`
	PageSize   *int     `pulumi:"pageSize"`
	Remark     *string  `pulumi:"remark"`
	SourceIp   *string  `pulumi:"sourceIp"`
	Status     *string  `pulumi:"status"`
}

func GetWebLockConfigsOutput(ctx *pulumi.Context, args GetWebLockConfigsOutputArgs, opts ...pulumi.InvokeOption) GetWebLockConfigsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetWebLockConfigsResult, error) {
			args := v.(GetWebLockConfigsArgs)
			r, err := GetWebLockConfigs(ctx, &args, opts...)
			var s GetWebLockConfigsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetWebLockConfigsResultOutput)
}

// A collection of arguments for invoking getWebLockConfigs.
type GetWebLockConfigsOutputArgs struct {
	// A list of Web Lock Config IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The language of the content within the request and the response. Valid values: `zh`, `en`.
	Lang pulumi.StringPtrInput `pulumi:"lang"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	PageNumber pulumi.IntPtrInput    `pulumi:"pageNumber"`
	PageSize   pulumi.IntPtrInput    `pulumi:"pageSize"`
	// The string that allows you to search for servers in fuzzy match mode. You can enter a server name or IP address.
	Remark pulumi.StringPtrInput `pulumi:"remark"`
	// The source IP address of the request.
	SourceIp pulumi.StringPtrInput `pulumi:"sourceIp"`
	// The protection status of the server that you want to query. Valid values: `on`, `off`.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetWebLockConfigsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWebLockConfigsArgs)(nil)).Elem()
}

// A collection of values returned by getWebLockConfigs.
type GetWebLockConfigsResultOutput struct{ *pulumi.OutputState }

func (GetWebLockConfigsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWebLockConfigsResult)(nil)).Elem()
}

func (o GetWebLockConfigsResultOutput) ToGetWebLockConfigsResultOutput() GetWebLockConfigsResultOutput {
	return o
}

func (o GetWebLockConfigsResultOutput) ToGetWebLockConfigsResultOutputWithContext(ctx context.Context) GetWebLockConfigsResultOutput {
	return o
}

// A list of Web Lock Config Entries. Each element contains the following attributes:
func (o GetWebLockConfigsResultOutput) Configs() GetWebLockConfigsConfigArrayOutput {
	return o.ApplyT(func(v GetWebLockConfigsResult) []GetWebLockConfigsConfig { return v.Configs }).(GetWebLockConfigsConfigArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetWebLockConfigsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetWebLockConfigsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of Web Lock Config IDs.
func (o GetWebLockConfigsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetWebLockConfigsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetWebLockConfigsResultOutput) Lang() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWebLockConfigsResult) *string { return v.Lang }).(pulumi.StringPtrOutput)
}

func (o GetWebLockConfigsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWebLockConfigsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetWebLockConfigsResultOutput) PageNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetWebLockConfigsResult) *int { return v.PageNumber }).(pulumi.IntPtrOutput)
}

func (o GetWebLockConfigsResultOutput) PageSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetWebLockConfigsResult) *int { return v.PageSize }).(pulumi.IntPtrOutput)
}

func (o GetWebLockConfigsResultOutput) Remark() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWebLockConfigsResult) *string { return v.Remark }).(pulumi.StringPtrOutput)
}

func (o GetWebLockConfigsResultOutput) SourceIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWebLockConfigsResult) *string { return v.SourceIp }).(pulumi.StringPtrOutput)
}

func (o GetWebLockConfigsResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWebLockConfigsResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetWebLockConfigsResultOutput{})
}

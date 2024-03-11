// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides a list of DNS Domain Records in an Alibaba Cloud account according to the specified filters.
//
// > **NOTE:** Available since v1.0.0.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/dns"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			recordsDs, err := dns.GetRecords(ctx, &dns.GetRecordsArgs{
//				DomainName:      "xiaozhu.top",
//				HostRecordRegex: pulumi.StringRef("^@"),
//				IsLocked:        pulumi.BoolRef(false),
//				OutputFile:      pulumi.StringRef("records.txt"),
//				Type:            pulumi.StringRef("A"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstRecordId", recordsDs.Records[0].RecordId)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetRecords(ctx *pulumi.Context, args *GetRecordsArgs, opts ...pulumi.InvokeOption) (*GetRecordsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRecordsResult
	err := ctx.Invoke("alicloud:dns/getRecords:getRecords", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRecords.
type GetRecordsArgs struct {
	// The domain name associated to the records.
	DomainName string `pulumi:"domainName"`
	// Host record regex.
	HostRecordRegex *string `pulumi:"hostRecordRegex"`
	// A list of record IDs.
	Ids []string `pulumi:"ids"`
	// Whether the record is locked or not.
	IsLocked *bool `pulumi:"isLocked"`
	// ISP line. Valid items are `default`, `telecom`, `unicom`, `mobile`, `oversea`, `edu`, `drpeng`, `btvn`, .etc. For checking all resolution lines enumeration please visit [Alibaba Cloud DNS doc](https://www.alibabacloud.com/help/en/doc-detail/29807.htm)
	Line *string `pulumi:"line"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// Record status. Valid items are `ENABLE` and `DISABLE`.
	Status *string `pulumi:"status"`
	// Record type. Valid items are `A`, `NS`, `MX`, `TXT`, `CNAME`, `SRV`, `AAAA`, `REDIRECT_URL`, `FORWORD_URL` .
	Type *string `pulumi:"type"`
	// Host record value regex.
	ValueRegex *string `pulumi:"valueRegex"`
}

// A collection of values returned by getRecords.
type GetRecordsResult struct {
	// Name of the domain the record belongs to.
	DomainName      string  `pulumi:"domainName"`
	HostRecordRegex *string `pulumi:"hostRecordRegex"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of record IDs.
	Ids      []string `pulumi:"ids"`
	IsLocked *bool    `pulumi:"isLocked"`
	// ISP line of the record.
	Line       *string `pulumi:"line"`
	OutputFile *string `pulumi:"outputFile"`
	// A list of records. Each element contains the following attributes:
	Records []GetRecordsRecord `pulumi:"records"`
	// Status of the record.
	Status *string `pulumi:"status"`
	// Type of the record.
	Type *string `pulumi:"type"`
	// A list of entire URLs. Each item format as `<host_record>.<domain_name>`.
	Urls       []string `pulumi:"urls"`
	ValueRegex *string  `pulumi:"valueRegex"`
}

func GetRecordsOutput(ctx *pulumi.Context, args GetRecordsOutputArgs, opts ...pulumi.InvokeOption) GetRecordsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRecordsResult, error) {
			args := v.(GetRecordsArgs)
			r, err := GetRecords(ctx, &args, opts...)
			var s GetRecordsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRecordsResultOutput)
}

// A collection of arguments for invoking getRecords.
type GetRecordsOutputArgs struct {
	// The domain name associated to the records.
	DomainName pulumi.StringInput `pulumi:"domainName"`
	// Host record regex.
	HostRecordRegex pulumi.StringPtrInput `pulumi:"hostRecordRegex"`
	// A list of record IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// Whether the record is locked or not.
	IsLocked pulumi.BoolPtrInput `pulumi:"isLocked"`
	// ISP line. Valid items are `default`, `telecom`, `unicom`, `mobile`, `oversea`, `edu`, `drpeng`, `btvn`, .etc. For checking all resolution lines enumeration please visit [Alibaba Cloud DNS doc](https://www.alibabacloud.com/help/en/doc-detail/29807.htm)
	Line pulumi.StringPtrInput `pulumi:"line"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// Record status. Valid items are `ENABLE` and `DISABLE`.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// Record type. Valid items are `A`, `NS`, `MX`, `TXT`, `CNAME`, `SRV`, `AAAA`, `REDIRECT_URL`, `FORWORD_URL` .
	Type pulumi.StringPtrInput `pulumi:"type"`
	// Host record value regex.
	ValueRegex pulumi.StringPtrInput `pulumi:"valueRegex"`
}

func (GetRecordsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRecordsArgs)(nil)).Elem()
}

// A collection of values returned by getRecords.
type GetRecordsResultOutput struct{ *pulumi.OutputState }

func (GetRecordsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRecordsResult)(nil)).Elem()
}

func (o GetRecordsResultOutput) ToGetRecordsResultOutput() GetRecordsResultOutput {
	return o
}

func (o GetRecordsResultOutput) ToGetRecordsResultOutputWithContext(ctx context.Context) GetRecordsResultOutput {
	return o
}

// Name of the domain the record belongs to.
func (o GetRecordsResultOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v GetRecordsResult) string { return v.DomainName }).(pulumi.StringOutput)
}

func (o GetRecordsResultOutput) HostRecordRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRecordsResult) *string { return v.HostRecordRegex }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRecordsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRecordsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of record IDs.
func (o GetRecordsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRecordsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetRecordsResultOutput) IsLocked() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetRecordsResult) *bool { return v.IsLocked }).(pulumi.BoolPtrOutput)
}

// ISP line of the record.
func (o GetRecordsResultOutput) Line() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRecordsResult) *string { return v.Line }).(pulumi.StringPtrOutput)
}

func (o GetRecordsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRecordsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// A list of records. Each element contains the following attributes:
func (o GetRecordsResultOutput) Records() GetRecordsRecordArrayOutput {
	return o.ApplyT(func(v GetRecordsResult) []GetRecordsRecord { return v.Records }).(GetRecordsRecordArrayOutput)
}

// Status of the record.
func (o GetRecordsResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRecordsResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// Type of the record.
func (o GetRecordsResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRecordsResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

// A list of entire URLs. Each item format as `<host_record>.<domain_name>`.
func (o GetRecordsResultOutput) Urls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRecordsResult) []string { return v.Urls }).(pulumi.StringArrayOutput)
}

func (o GetRecordsResultOutput) ValueRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRecordsResult) *string { return v.ValueRegex }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRecordsResultOutput{})
}

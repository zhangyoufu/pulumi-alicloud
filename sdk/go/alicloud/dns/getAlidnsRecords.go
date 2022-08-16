// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides a list of Alidns Domain Records in an Alibaba Cloud account according to the specified filters.
//
// > **NOTE:**  Available in 1.86.0+.
//
// ## Example Usage
//
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
//			recordsDs, err := dns.GetAlidnsRecords(ctx, &dns.GetAlidnsRecordsArgs{
//				DomainName: "xiaozhu.top",
//				Ids: []string{
//					"1978593525779****",
//				},
//				OutputFile: pulumi.StringRef("records.txt"),
//				Type:       pulumi.StringRef("A"),
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
func GetAlidnsRecords(ctx *pulumi.Context, args *GetAlidnsRecordsArgs, opts ...pulumi.InvokeOption) (*GetAlidnsRecordsResult, error) {
	var rv GetAlidnsRecordsResult
	err := ctx.Invoke("alicloud:dns/getAlidnsRecords:getAlidnsRecords", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAlidnsRecords.
type GetAlidnsRecordsArgs struct {
	// Sorting direction. Valid values: `DESC`,`ASC`. Default to `AESC`.
	Direction *string `pulumi:"direction"`
	// The domain name associated to the records.
	DomainName string `pulumi:"domainName"`
	// Domain name group ID.
	GroupId *int `pulumi:"groupId"`
	// A list of record IDs.
	Ids []string `pulumi:"ids"`
	// Keywords.
	KeyWord *string `pulumi:"keyWord"`
	// User language.
	Lang *string `pulumi:"lang"`
	// ISP line. For checking all resolution lines enumeration please visit [Alibaba Cloud DNS doc](https://www.alibabacloud.com/help/doc-detail/34339.htm)
	Line *string `pulumi:"line"`
	// Sort by. Sort from newest to oldest according to the time added by resolution.
	OrderBy    *string `pulumi:"orderBy"`
	OutputFile *string `pulumi:"outputFile"`
	// The keywords recorded by the host are searched according to the `%RRKeyWord%` mode, and are not case sensitive.
	RrKeyWord *string `pulumi:"rrKeyWord"`
	// Host record regex.
	RrRegex *string `pulumi:"rrRegex"`
	// Search mode, Valid values: `LIKE`, `EXACT`, `ADVANCED`, `LIKE` (fuzzy), `EXACT` (accurate) search supports KeyWord field, `ADVANCED` (advanced) mode supports other fields.
	SearchMode *string `pulumi:"searchMode"`
	// Record status. Valid values: `ENABLE` and `DISABLE`.
	Status *string `pulumi:"status"`
	// Record type. Valid values: `A`, `NS`, `MX`, `TXT`, `CNAME`, `SRV`, `AAAA`, `REDIRECT_URL`, `FORWORD_URL` .
	Type *string `pulumi:"type"`
	// Analyze type keywords, search by full match, not case sensitive.
	TypeKeyWord *string `pulumi:"typeKeyWord"`
	// The keywords of the recorded value are searched according to the `%ValueKeyWord%` mode, and are not case sensitive.
	ValueKeyWord *string `pulumi:"valueKeyWord"`
	// Host record value regex.
	ValueRegex *string `pulumi:"valueRegex"`
}

// A collection of values returned by getAlidnsRecords.
type GetAlidnsRecordsResult struct {
	Direction *string `pulumi:"direction"`
	// Name of the domain record belongs to.
	DomainName string `pulumi:"domainName"`
	GroupId    *int   `pulumi:"groupId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of record IDs.
	Ids     []string `pulumi:"ids"`
	KeyWord *string  `pulumi:"keyWord"`
	Lang    *string  `pulumi:"lang"`
	// ISP line of the record.
	Line       *string `pulumi:"line"`
	OrderBy    *string `pulumi:"orderBy"`
	OutputFile *string `pulumi:"outputFile"`
	// A list of records. Each element contains the following attributes:
	Records    []GetAlidnsRecordsRecord `pulumi:"records"`
	RrKeyWord  *string                  `pulumi:"rrKeyWord"`
	RrRegex    *string                  `pulumi:"rrRegex"`
	SearchMode *string                  `pulumi:"searchMode"`
	// Status of the record.
	Status *string `pulumi:"status"`
	// Type of the record.
	Type         *string `pulumi:"type"`
	TypeKeyWord  *string `pulumi:"typeKeyWord"`
	ValueKeyWord *string `pulumi:"valueKeyWord"`
	ValueRegex   *string `pulumi:"valueRegex"`
}

func GetAlidnsRecordsOutput(ctx *pulumi.Context, args GetAlidnsRecordsOutputArgs, opts ...pulumi.InvokeOption) GetAlidnsRecordsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAlidnsRecordsResult, error) {
			args := v.(GetAlidnsRecordsArgs)
			r, err := GetAlidnsRecords(ctx, &args, opts...)
			var s GetAlidnsRecordsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAlidnsRecordsResultOutput)
}

// A collection of arguments for invoking getAlidnsRecords.
type GetAlidnsRecordsOutputArgs struct {
	// Sorting direction. Valid values: `DESC`,`ASC`. Default to `AESC`.
	Direction pulumi.StringPtrInput `pulumi:"direction"`
	// The domain name associated to the records.
	DomainName pulumi.StringInput `pulumi:"domainName"`
	// Domain name group ID.
	GroupId pulumi.IntPtrInput `pulumi:"groupId"`
	// A list of record IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// Keywords.
	KeyWord pulumi.StringPtrInput `pulumi:"keyWord"`
	// User language.
	Lang pulumi.StringPtrInput `pulumi:"lang"`
	// ISP line. For checking all resolution lines enumeration please visit [Alibaba Cloud DNS doc](https://www.alibabacloud.com/help/doc-detail/34339.htm)
	Line pulumi.StringPtrInput `pulumi:"line"`
	// Sort by. Sort from newest to oldest according to the time added by resolution.
	OrderBy    pulumi.StringPtrInput `pulumi:"orderBy"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The keywords recorded by the host are searched according to the `%RRKeyWord%` mode, and are not case sensitive.
	RrKeyWord pulumi.StringPtrInput `pulumi:"rrKeyWord"`
	// Host record regex.
	RrRegex pulumi.StringPtrInput `pulumi:"rrRegex"`
	// Search mode, Valid values: `LIKE`, `EXACT`, `ADVANCED`, `LIKE` (fuzzy), `EXACT` (accurate) search supports KeyWord field, `ADVANCED` (advanced) mode supports other fields.
	SearchMode pulumi.StringPtrInput `pulumi:"searchMode"`
	// Record status. Valid values: `ENABLE` and `DISABLE`.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// Record type. Valid values: `A`, `NS`, `MX`, `TXT`, `CNAME`, `SRV`, `AAAA`, `REDIRECT_URL`, `FORWORD_URL` .
	Type pulumi.StringPtrInput `pulumi:"type"`
	// Analyze type keywords, search by full match, not case sensitive.
	TypeKeyWord pulumi.StringPtrInput `pulumi:"typeKeyWord"`
	// The keywords of the recorded value are searched according to the `%ValueKeyWord%` mode, and are not case sensitive.
	ValueKeyWord pulumi.StringPtrInput `pulumi:"valueKeyWord"`
	// Host record value regex.
	ValueRegex pulumi.StringPtrInput `pulumi:"valueRegex"`
}

func (GetAlidnsRecordsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAlidnsRecordsArgs)(nil)).Elem()
}

// A collection of values returned by getAlidnsRecords.
type GetAlidnsRecordsResultOutput struct{ *pulumi.OutputState }

func (GetAlidnsRecordsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAlidnsRecordsResult)(nil)).Elem()
}

func (o GetAlidnsRecordsResultOutput) ToGetAlidnsRecordsResultOutput() GetAlidnsRecordsResultOutput {
	return o
}

func (o GetAlidnsRecordsResultOutput) ToGetAlidnsRecordsResultOutputWithContext(ctx context.Context) GetAlidnsRecordsResultOutput {
	return o
}

func (o GetAlidnsRecordsResultOutput) Direction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlidnsRecordsResult) *string { return v.Direction }).(pulumi.StringPtrOutput)
}

// Name of the domain record belongs to.
func (o GetAlidnsRecordsResultOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v GetAlidnsRecordsResult) string { return v.DomainName }).(pulumi.StringOutput)
}

func (o GetAlidnsRecordsResultOutput) GroupId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetAlidnsRecordsResult) *int { return v.GroupId }).(pulumi.IntPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAlidnsRecordsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAlidnsRecordsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of record IDs.
func (o GetAlidnsRecordsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAlidnsRecordsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetAlidnsRecordsResultOutput) KeyWord() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlidnsRecordsResult) *string { return v.KeyWord }).(pulumi.StringPtrOutput)
}

func (o GetAlidnsRecordsResultOutput) Lang() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlidnsRecordsResult) *string { return v.Lang }).(pulumi.StringPtrOutput)
}

// ISP line of the record.
func (o GetAlidnsRecordsResultOutput) Line() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlidnsRecordsResult) *string { return v.Line }).(pulumi.StringPtrOutput)
}

func (o GetAlidnsRecordsResultOutput) OrderBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlidnsRecordsResult) *string { return v.OrderBy }).(pulumi.StringPtrOutput)
}

func (o GetAlidnsRecordsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlidnsRecordsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// A list of records. Each element contains the following attributes:
func (o GetAlidnsRecordsResultOutput) Records() GetAlidnsRecordsRecordArrayOutput {
	return o.ApplyT(func(v GetAlidnsRecordsResult) []GetAlidnsRecordsRecord { return v.Records }).(GetAlidnsRecordsRecordArrayOutput)
}

func (o GetAlidnsRecordsResultOutput) RrKeyWord() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlidnsRecordsResult) *string { return v.RrKeyWord }).(pulumi.StringPtrOutput)
}

func (o GetAlidnsRecordsResultOutput) RrRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlidnsRecordsResult) *string { return v.RrRegex }).(pulumi.StringPtrOutput)
}

func (o GetAlidnsRecordsResultOutput) SearchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlidnsRecordsResult) *string { return v.SearchMode }).(pulumi.StringPtrOutput)
}

// Status of the record.
func (o GetAlidnsRecordsResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlidnsRecordsResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// Type of the record.
func (o GetAlidnsRecordsResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlidnsRecordsResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o GetAlidnsRecordsResultOutput) TypeKeyWord() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlidnsRecordsResult) *string { return v.TypeKeyWord }).(pulumi.StringPtrOutput)
}

func (o GetAlidnsRecordsResultOutput) ValueKeyWord() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlidnsRecordsResult) *string { return v.ValueKeyWord }).(pulumi.StringPtrOutput)
}

func (o GetAlidnsRecordsResultOutput) ValueRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlidnsRecordsResult) *string { return v.ValueRegex }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAlidnsRecordsResultOutput{})
}

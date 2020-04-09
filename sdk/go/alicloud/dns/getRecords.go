// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package dns

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides a list of DNS Domain Records in an Alibaba Cloud account according to the specified filters.
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/dns_records.html.markdown.
func GetRecords(ctx *pulumi.Context, args *GetRecordsArgs, opts ...pulumi.InvokeOption) (*GetRecordsResult, error) {
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
	// ISP line. Valid items are `default`, `telecom`, `unicom`, `mobile`, `oversea`, `edu`, `drpeng`, `btvn`, .etc. For checking all resolution lines enumeration please visit [Alibaba Cloud DNS doc](https://www.alibabacloud.com/help/doc-detail/34339.htm)
	Line       *string `pulumi:"line"`
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
	// id is the provider-assigned unique ID for this managed resource.
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

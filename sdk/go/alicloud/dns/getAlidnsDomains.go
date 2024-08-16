// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides a list of Alidns Domains in an Alibaba Cloud account according to the specified filters.
//
// > **NOTE:**  Available in 1.95.0+.
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
//			domainsDs, err := dns.GetAlidnsDomains(ctx, &dns.GetAlidnsDomainsArgs{
//				DomainNameRegex: pulumi.StringRef("^hegu"),
//				OutputFile:      pulumi.StringRef("domains.txt"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstDomainId", domainsDs.Domains[0].DomainId)
//			return nil
//		})
//	}
//
// ```
func GetAlidnsDomains(ctx *pulumi.Context, args *GetAlidnsDomainsArgs, opts ...pulumi.InvokeOption) (*GetAlidnsDomainsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAlidnsDomainsResult
	err := ctx.Invoke("alicloud:dns/getAlidnsDomains:getAlidnsDomains", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAlidnsDomains.
type GetAlidnsDomainsArgs struct {
	// Specifies whether the domain is from Alibaba Cloud or not.
	AliDomain *bool `pulumi:"aliDomain"`
	// A regex string to filter results by the domain name.
	DomainNameRegex *string `pulumi:"domainNameRegex"`
	EnableDetails   *bool   `pulumi:"enableDetails"`
	// Domain group ID, if not filled, the default is all groups.
	GroupId *string `pulumi:"groupId"`
	// A regex string to filter results by the group name.
	GroupNameRegex *string `pulumi:"groupNameRegex"`
	// A list of domain IDs.
	Ids []string `pulumi:"ids"`
	// Cloud analysis product ID.
	InstanceId *string `pulumi:"instanceId"`
	// The keywords are searched according to the `%KeyWord%` mode, which is not case sensitive.
	KeyWord *string `pulumi:"keyWord"`
	// User language.
	Lang *string `pulumi:"lang"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The Id of resource group which the dns belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// Search mode, `LIKE` fuzzy search, `EXACT` exact search.
	SearchMode *string `pulumi:"searchMode"`
	// Whether to query the domain name star.
	Starmark *bool `pulumi:"starmark"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Cloud analysis version code.
	VersionCode *string `pulumi:"versionCode"`
}

// A collection of values returned by getAlidnsDomains.
type GetAlidnsDomainsResult struct {
	// Indicates whether the domain is an Alibaba Cloud domain.
	AliDomain       *bool   `pulumi:"aliDomain"`
	DomainNameRegex *string `pulumi:"domainNameRegex"`
	// A list of domains. Each element contains the following attributes:
	Domains       []GetAlidnsDomainsDomain `pulumi:"domains"`
	EnableDetails *bool                    `pulumi:"enableDetails"`
	// Id of group that contains the domain.
	GroupId        *string `pulumi:"groupId"`
	GroupNameRegex *string `pulumi:"groupNameRegex"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of domain IDs.
	Ids []string `pulumi:"ids"`
	// Cloud analysis product ID of the domain.
	InstanceId *string `pulumi:"instanceId"`
	KeyWord    *string `pulumi:"keyWord"`
	Lang       *string `pulumi:"lang"`
	// A list of domain names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// The Id of resource group which the dns belongs.
	ResourceGroupId *string           `pulumi:"resourceGroupId"`
	SearchMode      *string           `pulumi:"searchMode"`
	Starmark        *bool             `pulumi:"starmark"`
	Tags            map[string]string `pulumi:"tags"`
	// Cloud resolution version ID.
	VersionCode *string `pulumi:"versionCode"`
}

func GetAlidnsDomainsOutput(ctx *pulumi.Context, args GetAlidnsDomainsOutputArgs, opts ...pulumi.InvokeOption) GetAlidnsDomainsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAlidnsDomainsResult, error) {
			args := v.(GetAlidnsDomainsArgs)
			r, err := GetAlidnsDomains(ctx, &args, opts...)
			var s GetAlidnsDomainsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAlidnsDomainsResultOutput)
}

// A collection of arguments for invoking getAlidnsDomains.
type GetAlidnsDomainsOutputArgs struct {
	// Specifies whether the domain is from Alibaba Cloud or not.
	AliDomain pulumi.BoolPtrInput `pulumi:"aliDomain"`
	// A regex string to filter results by the domain name.
	DomainNameRegex pulumi.StringPtrInput `pulumi:"domainNameRegex"`
	EnableDetails   pulumi.BoolPtrInput   `pulumi:"enableDetails"`
	// Domain group ID, if not filled, the default is all groups.
	GroupId pulumi.StringPtrInput `pulumi:"groupId"`
	// A regex string to filter results by the group name.
	GroupNameRegex pulumi.StringPtrInput `pulumi:"groupNameRegex"`
	// A list of domain IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// Cloud analysis product ID.
	InstanceId pulumi.StringPtrInput `pulumi:"instanceId"`
	// The keywords are searched according to the `%KeyWord%` mode, which is not case sensitive.
	KeyWord pulumi.StringPtrInput `pulumi:"keyWord"`
	// User language.
	Lang pulumi.StringPtrInput `pulumi:"lang"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The Id of resource group which the dns belongs.
	ResourceGroupId pulumi.StringPtrInput `pulumi:"resourceGroupId"`
	// Search mode, `LIKE` fuzzy search, `EXACT` exact search.
	SearchMode pulumi.StringPtrInput `pulumi:"searchMode"`
	// Whether to query the domain name star.
	Starmark pulumi.BoolPtrInput `pulumi:"starmark"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput `pulumi:"tags"`
	// Cloud analysis version code.
	VersionCode pulumi.StringPtrInput `pulumi:"versionCode"`
}

func (GetAlidnsDomainsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAlidnsDomainsArgs)(nil)).Elem()
}

// A collection of values returned by getAlidnsDomains.
type GetAlidnsDomainsResultOutput struct{ *pulumi.OutputState }

func (GetAlidnsDomainsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAlidnsDomainsResult)(nil)).Elem()
}

func (o GetAlidnsDomainsResultOutput) ToGetAlidnsDomainsResultOutput() GetAlidnsDomainsResultOutput {
	return o
}

func (o GetAlidnsDomainsResultOutput) ToGetAlidnsDomainsResultOutputWithContext(ctx context.Context) GetAlidnsDomainsResultOutput {
	return o
}

// Indicates whether the domain is an Alibaba Cloud domain.
func (o GetAlidnsDomainsResultOutput) AliDomain() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetAlidnsDomainsResult) *bool { return v.AliDomain }).(pulumi.BoolPtrOutput)
}

func (o GetAlidnsDomainsResultOutput) DomainNameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlidnsDomainsResult) *string { return v.DomainNameRegex }).(pulumi.StringPtrOutput)
}

// A list of domains. Each element contains the following attributes:
func (o GetAlidnsDomainsResultOutput) Domains() GetAlidnsDomainsDomainArrayOutput {
	return o.ApplyT(func(v GetAlidnsDomainsResult) []GetAlidnsDomainsDomain { return v.Domains }).(GetAlidnsDomainsDomainArrayOutput)
}

func (o GetAlidnsDomainsResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetAlidnsDomainsResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

// Id of group that contains the domain.
func (o GetAlidnsDomainsResultOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlidnsDomainsResult) *string { return v.GroupId }).(pulumi.StringPtrOutput)
}

func (o GetAlidnsDomainsResultOutput) GroupNameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlidnsDomainsResult) *string { return v.GroupNameRegex }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAlidnsDomainsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAlidnsDomainsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of domain IDs.
func (o GetAlidnsDomainsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAlidnsDomainsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// Cloud analysis product ID of the domain.
func (o GetAlidnsDomainsResultOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlidnsDomainsResult) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

func (o GetAlidnsDomainsResultOutput) KeyWord() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlidnsDomainsResult) *string { return v.KeyWord }).(pulumi.StringPtrOutput)
}

func (o GetAlidnsDomainsResultOutput) Lang() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlidnsDomainsResult) *string { return v.Lang }).(pulumi.StringPtrOutput)
}

// A list of domain names.
func (o GetAlidnsDomainsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAlidnsDomainsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetAlidnsDomainsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlidnsDomainsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The Id of resource group which the dns belongs.
func (o GetAlidnsDomainsResultOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlidnsDomainsResult) *string { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

func (o GetAlidnsDomainsResultOutput) SearchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlidnsDomainsResult) *string { return v.SearchMode }).(pulumi.StringPtrOutput)
}

func (o GetAlidnsDomainsResultOutput) Starmark() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetAlidnsDomainsResult) *bool { return v.Starmark }).(pulumi.BoolPtrOutput)
}

func (o GetAlidnsDomainsResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetAlidnsDomainsResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Cloud resolution version ID.
func (o GetAlidnsDomainsResultOutput) VersionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlidnsDomainsResult) *string { return v.VersionCode }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAlidnsDomainsResultOutput{})
}

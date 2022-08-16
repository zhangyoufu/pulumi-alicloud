// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a VPC Dhcp Options Set resource.
//
// For information about VPC Dhcp Options Set and how to use it, see [What is Dhcp Options Set](https://www.alibabacloud.com/help/doc-detail/174112.htm).
//
// > **NOTE:** Available in v1.134.0+.
//
// ## Example Usage
//
// # Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vpc.NewDhcpOptionsSet(ctx, "example", &vpc.DhcpOptionsSetArgs{
//				DhcpOptionsSetDescription: pulumi.String("example_value"),
//				DhcpOptionsSetName:        pulumi.String("example_value"),
//				DomainName:                pulumi.String("example.com"),
//				DomainNameServers:         pulumi.String("100.100.2.136"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// VPC Dhcp Options Set can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:vpc/dhcpOptionsSet:DhcpOptionsSet example <id>
//
// ```
type DhcpOptionsSet struct {
	pulumi.CustomResourceState

	// AssociateVpcs. Number of VPCs that can be associated with each DHCP options set is 10. Field `associateVpcs` has been deprecated from provider version 1.153.0. It will be removed in the future version. Please use the new resource 'alicloud_vpc_dhcp_options_set_attachment' to attach DhcpOptionsSet and Vpc.
	//
	// Deprecated: Field 'associate_vpcs' has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_dhcp_options_set_attachment' to attach DhcpOptionsSet and Vpc.
	AssociateVpcs DhcpOptionsSetAssociateVpcArrayOutput `pulumi:"associateVpcs"`
	// The description of the DHCP options set. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	DhcpOptionsSetDescription pulumi.StringPtrOutput `pulumi:"dhcpOptionsSetDescription"`
	// The name of the DHCP options set. The name must be 2 to 128 characters in length and can contain letters, Chinese characters, digits, underscores (_), and hyphens (-). It must start with a letter or a Chinese character.
	DhcpOptionsSetName pulumi.StringPtrOutput `pulumi:"dhcpOptionsSetName"`
	// The root domain, for example, example.com. After a DHCP options set is associated with a Virtual Private Cloud (VPC) network, the root domain in the DHCP options set is automatically synchronized to the ECS instances in the VPC network.
	DomainName pulumi.StringPtrOutput `pulumi:"domainName"`
	// The DNS server IP addresses. Up to four DNS server IP addresses can be specified. IP addresses must be separated with commas (,).Before you specify any DNS server IP address, all ECS instances in the associated VPC network use the IP addresses of the Alibaba Cloud DNS servers, which are `100.100.2.136` and `100.100.2.138`.
	DomainNameServers pulumi.StringPtrOutput `pulumi:"domainNameServers"`
	// Specifies whether to precheck this request only. Valid values: `true` or `false`.
	DryRun pulumi.BoolPtrOutput `pulumi:"dryRun"`
	// The ID of the account to which the DHCP options set belongs.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// The status of the DHCP options set. Valid values: `Available`, `InUse` or `Pending`. `Available`: The DHCP options set is available for use. `InUse`: The DHCP options set is in use. `Pending`: The DHCP options set is being configured.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewDhcpOptionsSet registers a new resource with the given unique name, arguments, and options.
func NewDhcpOptionsSet(ctx *pulumi.Context,
	name string, args *DhcpOptionsSetArgs, opts ...pulumi.ResourceOption) (*DhcpOptionsSet, error) {
	if args == nil {
		args = &DhcpOptionsSetArgs{}
	}

	var resource DhcpOptionsSet
	err := ctx.RegisterResource("alicloud:vpc/dhcpOptionsSet:DhcpOptionsSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDhcpOptionsSet gets an existing DhcpOptionsSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDhcpOptionsSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DhcpOptionsSetState, opts ...pulumi.ResourceOption) (*DhcpOptionsSet, error) {
	var resource DhcpOptionsSet
	err := ctx.ReadResource("alicloud:vpc/dhcpOptionsSet:DhcpOptionsSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DhcpOptionsSet resources.
type dhcpOptionsSetState struct {
	// AssociateVpcs. Number of VPCs that can be associated with each DHCP options set is 10. Field `associateVpcs` has been deprecated from provider version 1.153.0. It will be removed in the future version. Please use the new resource 'alicloud_vpc_dhcp_options_set_attachment' to attach DhcpOptionsSet and Vpc.
	//
	// Deprecated: Field 'associate_vpcs' has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_dhcp_options_set_attachment' to attach DhcpOptionsSet and Vpc.
	AssociateVpcs []DhcpOptionsSetAssociateVpc `pulumi:"associateVpcs"`
	// The description of the DHCP options set. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	DhcpOptionsSetDescription *string `pulumi:"dhcpOptionsSetDescription"`
	// The name of the DHCP options set. The name must be 2 to 128 characters in length and can contain letters, Chinese characters, digits, underscores (_), and hyphens (-). It must start with a letter or a Chinese character.
	DhcpOptionsSetName *string `pulumi:"dhcpOptionsSetName"`
	// The root domain, for example, example.com. After a DHCP options set is associated with a Virtual Private Cloud (VPC) network, the root domain in the DHCP options set is automatically synchronized to the ECS instances in the VPC network.
	DomainName *string `pulumi:"domainName"`
	// The DNS server IP addresses. Up to four DNS server IP addresses can be specified. IP addresses must be separated with commas (,).Before you specify any DNS server IP address, all ECS instances in the associated VPC network use the IP addresses of the Alibaba Cloud DNS servers, which are `100.100.2.136` and `100.100.2.138`.
	DomainNameServers *string `pulumi:"domainNameServers"`
	// Specifies whether to precheck this request only. Valid values: `true` or `false`.
	DryRun *bool `pulumi:"dryRun"`
	// The ID of the account to which the DHCP options set belongs.
	OwnerId *string `pulumi:"ownerId"`
	// The status of the DHCP options set. Valid values: `Available`, `InUse` or `Pending`. `Available`: The DHCP options set is available for use. `InUse`: The DHCP options set is in use. `Pending`: The DHCP options set is being configured.
	Status *string `pulumi:"status"`
}

type DhcpOptionsSetState struct {
	// AssociateVpcs. Number of VPCs that can be associated with each DHCP options set is 10. Field `associateVpcs` has been deprecated from provider version 1.153.0. It will be removed in the future version. Please use the new resource 'alicloud_vpc_dhcp_options_set_attachment' to attach DhcpOptionsSet and Vpc.
	//
	// Deprecated: Field 'associate_vpcs' has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_dhcp_options_set_attachment' to attach DhcpOptionsSet and Vpc.
	AssociateVpcs DhcpOptionsSetAssociateVpcArrayInput
	// The description of the DHCP options set. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	DhcpOptionsSetDescription pulumi.StringPtrInput
	// The name of the DHCP options set. The name must be 2 to 128 characters in length and can contain letters, Chinese characters, digits, underscores (_), and hyphens (-). It must start with a letter or a Chinese character.
	DhcpOptionsSetName pulumi.StringPtrInput
	// The root domain, for example, example.com. After a DHCP options set is associated with a Virtual Private Cloud (VPC) network, the root domain in the DHCP options set is automatically synchronized to the ECS instances in the VPC network.
	DomainName pulumi.StringPtrInput
	// The DNS server IP addresses. Up to four DNS server IP addresses can be specified. IP addresses must be separated with commas (,).Before you specify any DNS server IP address, all ECS instances in the associated VPC network use the IP addresses of the Alibaba Cloud DNS servers, which are `100.100.2.136` and `100.100.2.138`.
	DomainNameServers pulumi.StringPtrInput
	// Specifies whether to precheck this request only. Valid values: `true` or `false`.
	DryRun pulumi.BoolPtrInput
	// The ID of the account to which the DHCP options set belongs.
	OwnerId pulumi.StringPtrInput
	// The status of the DHCP options set. Valid values: `Available`, `InUse` or `Pending`. `Available`: The DHCP options set is available for use. `InUse`: The DHCP options set is in use. `Pending`: The DHCP options set is being configured.
	Status pulumi.StringPtrInput
}

func (DhcpOptionsSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*dhcpOptionsSetState)(nil)).Elem()
}

type dhcpOptionsSetArgs struct {
	// AssociateVpcs. Number of VPCs that can be associated with each DHCP options set is 10. Field `associateVpcs` has been deprecated from provider version 1.153.0. It will be removed in the future version. Please use the new resource 'alicloud_vpc_dhcp_options_set_attachment' to attach DhcpOptionsSet and Vpc.
	//
	// Deprecated: Field 'associate_vpcs' has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_dhcp_options_set_attachment' to attach DhcpOptionsSet and Vpc.
	AssociateVpcs []DhcpOptionsSetAssociateVpc `pulumi:"associateVpcs"`
	// The description of the DHCP options set. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	DhcpOptionsSetDescription *string `pulumi:"dhcpOptionsSetDescription"`
	// The name of the DHCP options set. The name must be 2 to 128 characters in length and can contain letters, Chinese characters, digits, underscores (_), and hyphens (-). It must start with a letter or a Chinese character.
	DhcpOptionsSetName *string `pulumi:"dhcpOptionsSetName"`
	// The root domain, for example, example.com. After a DHCP options set is associated with a Virtual Private Cloud (VPC) network, the root domain in the DHCP options set is automatically synchronized to the ECS instances in the VPC network.
	DomainName *string `pulumi:"domainName"`
	// The DNS server IP addresses. Up to four DNS server IP addresses can be specified. IP addresses must be separated with commas (,).Before you specify any DNS server IP address, all ECS instances in the associated VPC network use the IP addresses of the Alibaba Cloud DNS servers, which are `100.100.2.136` and `100.100.2.138`.
	DomainNameServers *string `pulumi:"domainNameServers"`
	// Specifies whether to precheck this request only. Valid values: `true` or `false`.
	DryRun *bool `pulumi:"dryRun"`
}

// The set of arguments for constructing a DhcpOptionsSet resource.
type DhcpOptionsSetArgs struct {
	// AssociateVpcs. Number of VPCs that can be associated with each DHCP options set is 10. Field `associateVpcs` has been deprecated from provider version 1.153.0. It will be removed in the future version. Please use the new resource 'alicloud_vpc_dhcp_options_set_attachment' to attach DhcpOptionsSet and Vpc.
	//
	// Deprecated: Field 'associate_vpcs' has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_dhcp_options_set_attachment' to attach DhcpOptionsSet and Vpc.
	AssociateVpcs DhcpOptionsSetAssociateVpcArrayInput
	// The description of the DHCP options set. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	DhcpOptionsSetDescription pulumi.StringPtrInput
	// The name of the DHCP options set. The name must be 2 to 128 characters in length and can contain letters, Chinese characters, digits, underscores (_), and hyphens (-). It must start with a letter or a Chinese character.
	DhcpOptionsSetName pulumi.StringPtrInput
	// The root domain, for example, example.com. After a DHCP options set is associated with a Virtual Private Cloud (VPC) network, the root domain in the DHCP options set is automatically synchronized to the ECS instances in the VPC network.
	DomainName pulumi.StringPtrInput
	// The DNS server IP addresses. Up to four DNS server IP addresses can be specified. IP addresses must be separated with commas (,).Before you specify any DNS server IP address, all ECS instances in the associated VPC network use the IP addresses of the Alibaba Cloud DNS servers, which are `100.100.2.136` and `100.100.2.138`.
	DomainNameServers pulumi.StringPtrInput
	// Specifies whether to precheck this request only. Valid values: `true` or `false`.
	DryRun pulumi.BoolPtrInput
}

func (DhcpOptionsSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dhcpOptionsSetArgs)(nil)).Elem()
}

type DhcpOptionsSetInput interface {
	pulumi.Input

	ToDhcpOptionsSetOutput() DhcpOptionsSetOutput
	ToDhcpOptionsSetOutputWithContext(ctx context.Context) DhcpOptionsSetOutput
}

func (*DhcpOptionsSet) ElementType() reflect.Type {
	return reflect.TypeOf((**DhcpOptionsSet)(nil)).Elem()
}

func (i *DhcpOptionsSet) ToDhcpOptionsSetOutput() DhcpOptionsSetOutput {
	return i.ToDhcpOptionsSetOutputWithContext(context.Background())
}

func (i *DhcpOptionsSet) ToDhcpOptionsSetOutputWithContext(ctx context.Context) DhcpOptionsSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DhcpOptionsSetOutput)
}

// DhcpOptionsSetArrayInput is an input type that accepts DhcpOptionsSetArray and DhcpOptionsSetArrayOutput values.
// You can construct a concrete instance of `DhcpOptionsSetArrayInput` via:
//
//	DhcpOptionsSetArray{ DhcpOptionsSetArgs{...} }
type DhcpOptionsSetArrayInput interface {
	pulumi.Input

	ToDhcpOptionsSetArrayOutput() DhcpOptionsSetArrayOutput
	ToDhcpOptionsSetArrayOutputWithContext(context.Context) DhcpOptionsSetArrayOutput
}

type DhcpOptionsSetArray []DhcpOptionsSetInput

func (DhcpOptionsSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DhcpOptionsSet)(nil)).Elem()
}

func (i DhcpOptionsSetArray) ToDhcpOptionsSetArrayOutput() DhcpOptionsSetArrayOutput {
	return i.ToDhcpOptionsSetArrayOutputWithContext(context.Background())
}

func (i DhcpOptionsSetArray) ToDhcpOptionsSetArrayOutputWithContext(ctx context.Context) DhcpOptionsSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DhcpOptionsSetArrayOutput)
}

// DhcpOptionsSetMapInput is an input type that accepts DhcpOptionsSetMap and DhcpOptionsSetMapOutput values.
// You can construct a concrete instance of `DhcpOptionsSetMapInput` via:
//
//	DhcpOptionsSetMap{ "key": DhcpOptionsSetArgs{...} }
type DhcpOptionsSetMapInput interface {
	pulumi.Input

	ToDhcpOptionsSetMapOutput() DhcpOptionsSetMapOutput
	ToDhcpOptionsSetMapOutputWithContext(context.Context) DhcpOptionsSetMapOutput
}

type DhcpOptionsSetMap map[string]DhcpOptionsSetInput

func (DhcpOptionsSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DhcpOptionsSet)(nil)).Elem()
}

func (i DhcpOptionsSetMap) ToDhcpOptionsSetMapOutput() DhcpOptionsSetMapOutput {
	return i.ToDhcpOptionsSetMapOutputWithContext(context.Background())
}

func (i DhcpOptionsSetMap) ToDhcpOptionsSetMapOutputWithContext(ctx context.Context) DhcpOptionsSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DhcpOptionsSetMapOutput)
}

type DhcpOptionsSetOutput struct{ *pulumi.OutputState }

func (DhcpOptionsSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DhcpOptionsSet)(nil)).Elem()
}

func (o DhcpOptionsSetOutput) ToDhcpOptionsSetOutput() DhcpOptionsSetOutput {
	return o
}

func (o DhcpOptionsSetOutput) ToDhcpOptionsSetOutputWithContext(ctx context.Context) DhcpOptionsSetOutput {
	return o
}

// AssociateVpcs. Number of VPCs that can be associated with each DHCP options set is 10. Field `associateVpcs` has been deprecated from provider version 1.153.0. It will be removed in the future version. Please use the new resource 'alicloud_vpc_dhcp_options_set_attachment' to attach DhcpOptionsSet and Vpc.
//
// Deprecated: Field 'associate_vpcs' has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_dhcp_options_set_attachment' to attach DhcpOptionsSet and Vpc.
func (o DhcpOptionsSetOutput) AssociateVpcs() DhcpOptionsSetAssociateVpcArrayOutput {
	return o.ApplyT(func(v *DhcpOptionsSet) DhcpOptionsSetAssociateVpcArrayOutput { return v.AssociateVpcs }).(DhcpOptionsSetAssociateVpcArrayOutput)
}

// The description of the DHCP options set. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
func (o DhcpOptionsSetOutput) DhcpOptionsSetDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DhcpOptionsSet) pulumi.StringPtrOutput { return v.DhcpOptionsSetDescription }).(pulumi.StringPtrOutput)
}

// The name of the DHCP options set. The name must be 2 to 128 characters in length and can contain letters, Chinese characters, digits, underscores (_), and hyphens (-). It must start with a letter or a Chinese character.
func (o DhcpOptionsSetOutput) DhcpOptionsSetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DhcpOptionsSet) pulumi.StringPtrOutput { return v.DhcpOptionsSetName }).(pulumi.StringPtrOutput)
}

// The root domain, for example, example.com. After a DHCP options set is associated with a Virtual Private Cloud (VPC) network, the root domain in the DHCP options set is automatically synchronized to the ECS instances in the VPC network.
func (o DhcpOptionsSetOutput) DomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DhcpOptionsSet) pulumi.StringPtrOutput { return v.DomainName }).(pulumi.StringPtrOutput)
}

// The DNS server IP addresses. Up to four DNS server IP addresses can be specified. IP addresses must be separated with commas (,).Before you specify any DNS server IP address, all ECS instances in the associated VPC network use the IP addresses of the Alibaba Cloud DNS servers, which are `100.100.2.136` and `100.100.2.138`.
func (o DhcpOptionsSetOutput) DomainNameServers() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DhcpOptionsSet) pulumi.StringPtrOutput { return v.DomainNameServers }).(pulumi.StringPtrOutput)
}

// Specifies whether to precheck this request only. Valid values: `true` or `false`.
func (o DhcpOptionsSetOutput) DryRun() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DhcpOptionsSet) pulumi.BoolPtrOutput { return v.DryRun }).(pulumi.BoolPtrOutput)
}

// The ID of the account to which the DHCP options set belongs.
func (o DhcpOptionsSetOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpOptionsSet) pulumi.StringOutput { return v.OwnerId }).(pulumi.StringOutput)
}

// The status of the DHCP options set. Valid values: `Available`, `InUse` or `Pending`. `Available`: The DHCP options set is available for use. `InUse`: The DHCP options set is in use. `Pending`: The DHCP options set is being configured.
func (o DhcpOptionsSetOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpOptionsSet) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type DhcpOptionsSetArrayOutput struct{ *pulumi.OutputState }

func (DhcpOptionsSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DhcpOptionsSet)(nil)).Elem()
}

func (o DhcpOptionsSetArrayOutput) ToDhcpOptionsSetArrayOutput() DhcpOptionsSetArrayOutput {
	return o
}

func (o DhcpOptionsSetArrayOutput) ToDhcpOptionsSetArrayOutputWithContext(ctx context.Context) DhcpOptionsSetArrayOutput {
	return o
}

func (o DhcpOptionsSetArrayOutput) Index(i pulumi.IntInput) DhcpOptionsSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DhcpOptionsSet {
		return vs[0].([]*DhcpOptionsSet)[vs[1].(int)]
	}).(DhcpOptionsSetOutput)
}

type DhcpOptionsSetMapOutput struct{ *pulumi.OutputState }

func (DhcpOptionsSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DhcpOptionsSet)(nil)).Elem()
}

func (o DhcpOptionsSetMapOutput) ToDhcpOptionsSetMapOutput() DhcpOptionsSetMapOutput {
	return o
}

func (o DhcpOptionsSetMapOutput) ToDhcpOptionsSetMapOutputWithContext(ctx context.Context) DhcpOptionsSetMapOutput {
	return o
}

func (o DhcpOptionsSetMapOutput) MapIndex(k pulumi.StringInput) DhcpOptionsSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DhcpOptionsSet {
		return vs[0].(map[string]*DhcpOptionsSet)[vs[1].(string)]
	}).(DhcpOptionsSetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DhcpOptionsSetInput)(nil)).Elem(), &DhcpOptionsSet{})
	pulumi.RegisterInputType(reflect.TypeOf((*DhcpOptionsSetArrayInput)(nil)).Elem(), DhcpOptionsSetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DhcpOptionsSetMapInput)(nil)).Elem(), DhcpOptionsSetMap{})
	pulumi.RegisterOutputType(DhcpOptionsSetOutput{})
	pulumi.RegisterOutputType(DhcpOptionsSetArrayOutput{})
	pulumi.RegisterOutputType(DhcpOptionsSetMapOutput{})
}

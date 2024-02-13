// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a VPC Dhcp Options Set resource. DHCP option set.
//
// For information about VPC Dhcp Options Set and how to use it, see [What is Dhcp Options Set](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/dhcp-options-sets-overview).
//
// > **NOTE:** Available since v1.134.0.
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
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "terraform-example"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			domain := "terraform-example.com"
//			if param := cfg.Get("domain"); param != "" {
//				domain = param
//			}
//			_, err := vpc.NewDhcpOptionsSet(ctx, "example", &vpc.DhcpOptionsSetArgs{
//				DhcpOptionsSetName:        pulumi.String(name),
//				DhcpOptionsSetDescription: pulumi.String(name),
//				DomainName:                pulumi.String(domain),
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
// $ pulumi import alicloud:vpc/dhcpOptionsSet:DhcpOptionsSet example <id>
// ```
type DhcpOptionsSet struct {
	pulumi.CustomResourceState

	// Field 'associate_vpcs' has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_dhcp_options_set_attachment' to attach DhcpOptionsSet and Vpc. See `associateVpcs` below.
	//
	// Deprecated: Field 'associate_vpcs' has been deprecated from provider version 1.211.0. Field 'associate_vpcs' has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_dhcp_options_set_attachment' to attach DhcpOptionsSet and Vpc.
	AssociateVpcs DhcpOptionsSetAssociateVpcArrayOutput `pulumi:"associateVpcs"`
	// The description can be blank or contain 1 to 256 characters. It must start with a letter or Chinese character but cannot start with http:// or https://.
	DhcpOptionsSetDescription pulumi.StringPtrOutput `pulumi:"dhcpOptionsSetDescription"`
	// The name must be 2 to 128 characters in length and can contain letters, Chinese characters, digits, underscores (_), and hyphens (-). It must start with a letter or a Chinese character.
	DhcpOptionsSetName pulumi.StringPtrOutput `pulumi:"dhcpOptionsSetName"`
	// The root domain, for example, example.com. After a DHCP options set is associated with a Virtual Private Cloud (VPC) network, the root domain in the DHCP options set is automatically synchronized to the ECS instances in the VPC network.
	DomainName pulumi.StringPtrOutput `pulumi:"domainName"`
	// The DNS server IP addresses. Up to four DNS server IP addresses can be specified. IP addresses must be separated with commas (,).Before you specify any DNS server IP address, all ECS instances in the associated VPC network use the IP addresses of the Alibaba Cloud DNS servers, which are 100.100.2.136 and 100.100.2.138.
	DomainNameServers pulumi.StringPtrOutput `pulumi:"domainNameServers"`
	// Whether to PreCheck only this request, value:
	// - **true**: sends a check request and does not delete the DHCP option set. Check items include whether required parameters are filled in, request format, and restrictions. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.
	// - **false** (default): Sends a normal request and directly deletes the DHCP option set after checking.
	DryRun pulumi.BoolPtrOutput `pulumi:"dryRun"`
	// The lease time of the IPv6 DHCP option set.When the lease time is set to hours: Unit: h. Value range: 24h ~ 1176h,87600h ~ 175200h. Default value: 87600h.When the lease time is set to day: Unit: d. Value range: 1d ~ 49d,3650d ~ 7300d. Default value: 3650d.
	Ipv6LeaseTime pulumi.StringOutput `pulumi:"ipv6LeaseTime"`
	// The lease time of the IPv4 DHCP option set.When the lease time is set to hours: Unit: h. Value range: 24h ~ 1176h,87600h ~ 175200h. Default value: 87600h.When the lease time is set to day: Unit: d. Value range: 1d ~ 49d,3650d ~ 7300d. Default value: 3650d.
	LeaseTime pulumi.StringOutput `pulumi:"leaseTime"`
	// The ID of the account to which the DHCP options set belongs.
	OwnerId pulumi.IntOutput `pulumi:"ownerId"`
	// The ID of the resource group.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// The status of the resource.
	Status pulumi.StringOutput `pulumi:"status"`
	// Tags of the current resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewDhcpOptionsSet registers a new resource with the given unique name, arguments, and options.
func NewDhcpOptionsSet(ctx *pulumi.Context,
	name string, args *DhcpOptionsSetArgs, opts ...pulumi.ResourceOption) (*DhcpOptionsSet, error) {
	if args == nil {
		args = &DhcpOptionsSetArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
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
	// Field 'associate_vpcs' has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_dhcp_options_set_attachment' to attach DhcpOptionsSet and Vpc. See `associateVpcs` below.
	//
	// Deprecated: Field 'associate_vpcs' has been deprecated from provider version 1.211.0. Field 'associate_vpcs' has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_dhcp_options_set_attachment' to attach DhcpOptionsSet and Vpc.
	AssociateVpcs []DhcpOptionsSetAssociateVpc `pulumi:"associateVpcs"`
	// The description can be blank or contain 1 to 256 characters. It must start with a letter or Chinese character but cannot start with http:// or https://.
	DhcpOptionsSetDescription *string `pulumi:"dhcpOptionsSetDescription"`
	// The name must be 2 to 128 characters in length and can contain letters, Chinese characters, digits, underscores (_), and hyphens (-). It must start with a letter or a Chinese character.
	DhcpOptionsSetName *string `pulumi:"dhcpOptionsSetName"`
	// The root domain, for example, example.com. After a DHCP options set is associated with a Virtual Private Cloud (VPC) network, the root domain in the DHCP options set is automatically synchronized to the ECS instances in the VPC network.
	DomainName *string `pulumi:"domainName"`
	// The DNS server IP addresses. Up to four DNS server IP addresses can be specified. IP addresses must be separated with commas (,).Before you specify any DNS server IP address, all ECS instances in the associated VPC network use the IP addresses of the Alibaba Cloud DNS servers, which are 100.100.2.136 and 100.100.2.138.
	DomainNameServers *string `pulumi:"domainNameServers"`
	// Whether to PreCheck only this request, value:
	// - **true**: sends a check request and does not delete the DHCP option set. Check items include whether required parameters are filled in, request format, and restrictions. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.
	// - **false** (default): Sends a normal request and directly deletes the DHCP option set after checking.
	DryRun *bool `pulumi:"dryRun"`
	// The lease time of the IPv6 DHCP option set.When the lease time is set to hours: Unit: h. Value range: 24h ~ 1176h,87600h ~ 175200h. Default value: 87600h.When the lease time is set to day: Unit: d. Value range: 1d ~ 49d,3650d ~ 7300d. Default value: 3650d.
	Ipv6LeaseTime *string `pulumi:"ipv6LeaseTime"`
	// The lease time of the IPv4 DHCP option set.When the lease time is set to hours: Unit: h. Value range: 24h ~ 1176h,87600h ~ 175200h. Default value: 87600h.When the lease time is set to day: Unit: d. Value range: 1d ~ 49d,3650d ~ 7300d. Default value: 3650d.
	LeaseTime *string `pulumi:"leaseTime"`
	// The ID of the account to which the DHCP options set belongs.
	OwnerId *int `pulumi:"ownerId"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The status of the resource.
	Status *string `pulumi:"status"`
	// Tags of the current resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

type DhcpOptionsSetState struct {
	// Field 'associate_vpcs' has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_dhcp_options_set_attachment' to attach DhcpOptionsSet and Vpc. See `associateVpcs` below.
	//
	// Deprecated: Field 'associate_vpcs' has been deprecated from provider version 1.211.0. Field 'associate_vpcs' has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_dhcp_options_set_attachment' to attach DhcpOptionsSet and Vpc.
	AssociateVpcs DhcpOptionsSetAssociateVpcArrayInput
	// The description can be blank or contain 1 to 256 characters. It must start with a letter or Chinese character but cannot start with http:// or https://.
	DhcpOptionsSetDescription pulumi.StringPtrInput
	// The name must be 2 to 128 characters in length and can contain letters, Chinese characters, digits, underscores (_), and hyphens (-). It must start with a letter or a Chinese character.
	DhcpOptionsSetName pulumi.StringPtrInput
	// The root domain, for example, example.com. After a DHCP options set is associated with a Virtual Private Cloud (VPC) network, the root domain in the DHCP options set is automatically synchronized to the ECS instances in the VPC network.
	DomainName pulumi.StringPtrInput
	// The DNS server IP addresses. Up to four DNS server IP addresses can be specified. IP addresses must be separated with commas (,).Before you specify any DNS server IP address, all ECS instances in the associated VPC network use the IP addresses of the Alibaba Cloud DNS servers, which are 100.100.2.136 and 100.100.2.138.
	DomainNameServers pulumi.StringPtrInput
	// Whether to PreCheck only this request, value:
	// - **true**: sends a check request and does not delete the DHCP option set. Check items include whether required parameters are filled in, request format, and restrictions. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.
	// - **false** (default): Sends a normal request and directly deletes the DHCP option set after checking.
	DryRun pulumi.BoolPtrInput
	// The lease time of the IPv6 DHCP option set.When the lease time is set to hours: Unit: h. Value range: 24h ~ 1176h,87600h ~ 175200h. Default value: 87600h.When the lease time is set to day: Unit: d. Value range: 1d ~ 49d,3650d ~ 7300d. Default value: 3650d.
	Ipv6LeaseTime pulumi.StringPtrInput
	// The lease time of the IPv4 DHCP option set.When the lease time is set to hours: Unit: h. Value range: 24h ~ 1176h,87600h ~ 175200h. Default value: 87600h.When the lease time is set to day: Unit: d. Value range: 1d ~ 49d,3650d ~ 7300d. Default value: 3650d.
	LeaseTime pulumi.StringPtrInput
	// The ID of the account to which the DHCP options set belongs.
	OwnerId pulumi.IntPtrInput
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	// The status of the resource.
	Status pulumi.StringPtrInput
	// Tags of the current resource.
	Tags pulumi.MapInput
}

func (DhcpOptionsSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*dhcpOptionsSetState)(nil)).Elem()
}

type dhcpOptionsSetArgs struct {
	// Field 'associate_vpcs' has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_dhcp_options_set_attachment' to attach DhcpOptionsSet and Vpc. See `associateVpcs` below.
	//
	// Deprecated: Field 'associate_vpcs' has been deprecated from provider version 1.211.0. Field 'associate_vpcs' has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_dhcp_options_set_attachment' to attach DhcpOptionsSet and Vpc.
	AssociateVpcs []DhcpOptionsSetAssociateVpc `pulumi:"associateVpcs"`
	// The description can be blank or contain 1 to 256 characters. It must start with a letter or Chinese character but cannot start with http:// or https://.
	DhcpOptionsSetDescription *string `pulumi:"dhcpOptionsSetDescription"`
	// The name must be 2 to 128 characters in length and can contain letters, Chinese characters, digits, underscores (_), and hyphens (-). It must start with a letter or a Chinese character.
	DhcpOptionsSetName *string `pulumi:"dhcpOptionsSetName"`
	// The root domain, for example, example.com. After a DHCP options set is associated with a Virtual Private Cloud (VPC) network, the root domain in the DHCP options set is automatically synchronized to the ECS instances in the VPC network.
	DomainName *string `pulumi:"domainName"`
	// The DNS server IP addresses. Up to four DNS server IP addresses can be specified. IP addresses must be separated with commas (,).Before you specify any DNS server IP address, all ECS instances in the associated VPC network use the IP addresses of the Alibaba Cloud DNS servers, which are 100.100.2.136 and 100.100.2.138.
	DomainNameServers *string `pulumi:"domainNameServers"`
	// Whether to PreCheck only this request, value:
	// - **true**: sends a check request and does not delete the DHCP option set. Check items include whether required parameters are filled in, request format, and restrictions. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.
	// - **false** (default): Sends a normal request and directly deletes the DHCP option set after checking.
	DryRun *bool `pulumi:"dryRun"`
	// The lease time of the IPv6 DHCP option set.When the lease time is set to hours: Unit: h. Value range: 24h ~ 1176h,87600h ~ 175200h. Default value: 87600h.When the lease time is set to day: Unit: d. Value range: 1d ~ 49d,3650d ~ 7300d. Default value: 3650d.
	Ipv6LeaseTime *string `pulumi:"ipv6LeaseTime"`
	// The lease time of the IPv4 DHCP option set.When the lease time is set to hours: Unit: h. Value range: 24h ~ 1176h,87600h ~ 175200h. Default value: 87600h.When the lease time is set to day: Unit: d. Value range: 1d ~ 49d,3650d ~ 7300d. Default value: 3650d.
	LeaseTime *string `pulumi:"leaseTime"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// Tags of the current resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a DhcpOptionsSet resource.
type DhcpOptionsSetArgs struct {
	// Field 'associate_vpcs' has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_dhcp_options_set_attachment' to attach DhcpOptionsSet and Vpc. See `associateVpcs` below.
	//
	// Deprecated: Field 'associate_vpcs' has been deprecated from provider version 1.211.0. Field 'associate_vpcs' has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_dhcp_options_set_attachment' to attach DhcpOptionsSet and Vpc.
	AssociateVpcs DhcpOptionsSetAssociateVpcArrayInput
	// The description can be blank or contain 1 to 256 characters. It must start with a letter or Chinese character but cannot start with http:// or https://.
	DhcpOptionsSetDescription pulumi.StringPtrInput
	// The name must be 2 to 128 characters in length and can contain letters, Chinese characters, digits, underscores (_), and hyphens (-). It must start with a letter or a Chinese character.
	DhcpOptionsSetName pulumi.StringPtrInput
	// The root domain, for example, example.com. After a DHCP options set is associated with a Virtual Private Cloud (VPC) network, the root domain in the DHCP options set is automatically synchronized to the ECS instances in the VPC network.
	DomainName pulumi.StringPtrInput
	// The DNS server IP addresses. Up to four DNS server IP addresses can be specified. IP addresses must be separated with commas (,).Before you specify any DNS server IP address, all ECS instances in the associated VPC network use the IP addresses of the Alibaba Cloud DNS servers, which are 100.100.2.136 and 100.100.2.138.
	DomainNameServers pulumi.StringPtrInput
	// Whether to PreCheck only this request, value:
	// - **true**: sends a check request and does not delete the DHCP option set. Check items include whether required parameters are filled in, request format, and restrictions. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.
	// - **false** (default): Sends a normal request and directly deletes the DHCP option set after checking.
	DryRun pulumi.BoolPtrInput
	// The lease time of the IPv6 DHCP option set.When the lease time is set to hours: Unit: h. Value range: 24h ~ 1176h,87600h ~ 175200h. Default value: 87600h.When the lease time is set to day: Unit: d. Value range: 1d ~ 49d,3650d ~ 7300d. Default value: 3650d.
	Ipv6LeaseTime pulumi.StringPtrInput
	// The lease time of the IPv4 DHCP option set.When the lease time is set to hours: Unit: h. Value range: 24h ~ 1176h,87600h ~ 175200h. Default value: 87600h.When the lease time is set to day: Unit: d. Value range: 1d ~ 49d,3650d ~ 7300d. Default value: 3650d.
	LeaseTime pulumi.StringPtrInput
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	// Tags of the current resource.
	Tags pulumi.MapInput
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

// Field 'associate_vpcs' has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_dhcp_options_set_attachment' to attach DhcpOptionsSet and Vpc. See `associateVpcs` below.
//
// Deprecated: Field 'associate_vpcs' has been deprecated from provider version 1.211.0. Field 'associate_vpcs' has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_dhcp_options_set_attachment' to attach DhcpOptionsSet and Vpc.
func (o DhcpOptionsSetOutput) AssociateVpcs() DhcpOptionsSetAssociateVpcArrayOutput {
	return o.ApplyT(func(v *DhcpOptionsSet) DhcpOptionsSetAssociateVpcArrayOutput { return v.AssociateVpcs }).(DhcpOptionsSetAssociateVpcArrayOutput)
}

// The description can be blank or contain 1 to 256 characters. It must start with a letter or Chinese character but cannot start with http:// or https://.
func (o DhcpOptionsSetOutput) DhcpOptionsSetDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DhcpOptionsSet) pulumi.StringPtrOutput { return v.DhcpOptionsSetDescription }).(pulumi.StringPtrOutput)
}

// The name must be 2 to 128 characters in length and can contain letters, Chinese characters, digits, underscores (_), and hyphens (-). It must start with a letter or a Chinese character.
func (o DhcpOptionsSetOutput) DhcpOptionsSetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DhcpOptionsSet) pulumi.StringPtrOutput { return v.DhcpOptionsSetName }).(pulumi.StringPtrOutput)
}

// The root domain, for example, example.com. After a DHCP options set is associated with a Virtual Private Cloud (VPC) network, the root domain in the DHCP options set is automatically synchronized to the ECS instances in the VPC network.
func (o DhcpOptionsSetOutput) DomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DhcpOptionsSet) pulumi.StringPtrOutput { return v.DomainName }).(pulumi.StringPtrOutput)
}

// The DNS server IP addresses. Up to four DNS server IP addresses can be specified. IP addresses must be separated with commas (,).Before you specify any DNS server IP address, all ECS instances in the associated VPC network use the IP addresses of the Alibaba Cloud DNS servers, which are 100.100.2.136 and 100.100.2.138.
func (o DhcpOptionsSetOutput) DomainNameServers() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DhcpOptionsSet) pulumi.StringPtrOutput { return v.DomainNameServers }).(pulumi.StringPtrOutput)
}

// Whether to PreCheck only this request, value:
// - **true**: sends a check request and does not delete the DHCP option set. Check items include whether required parameters are filled in, request format, and restrictions. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.
// - **false** (default): Sends a normal request and directly deletes the DHCP option set after checking.
func (o DhcpOptionsSetOutput) DryRun() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DhcpOptionsSet) pulumi.BoolPtrOutput { return v.DryRun }).(pulumi.BoolPtrOutput)
}

// The lease time of the IPv6 DHCP option set.When the lease time is set to hours: Unit: h. Value range: 24h ~ 1176h,87600h ~ 175200h. Default value: 87600h.When the lease time is set to day: Unit: d. Value range: 1d ~ 49d,3650d ~ 7300d. Default value: 3650d.
func (o DhcpOptionsSetOutput) Ipv6LeaseTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpOptionsSet) pulumi.StringOutput { return v.Ipv6LeaseTime }).(pulumi.StringOutput)
}

// The lease time of the IPv4 DHCP option set.When the lease time is set to hours: Unit: h. Value range: 24h ~ 1176h,87600h ~ 175200h. Default value: 87600h.When the lease time is set to day: Unit: d. Value range: 1d ~ 49d,3650d ~ 7300d. Default value: 3650d.
func (o DhcpOptionsSetOutput) LeaseTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpOptionsSet) pulumi.StringOutput { return v.LeaseTime }).(pulumi.StringOutput)
}

// The ID of the account to which the DHCP options set belongs.
func (o DhcpOptionsSetOutput) OwnerId() pulumi.IntOutput {
	return o.ApplyT(func(v *DhcpOptionsSet) pulumi.IntOutput { return v.OwnerId }).(pulumi.IntOutput)
}

// The ID of the resource group.
func (o DhcpOptionsSetOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpOptionsSet) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// The status of the resource.
func (o DhcpOptionsSetOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpOptionsSet) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Tags of the current resource.
func (o DhcpOptionsSetOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *DhcpOptionsSet) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
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

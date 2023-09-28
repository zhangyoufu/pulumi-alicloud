// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// > **DEPRECATED:**  This resource has been deprecated and using wafv3.Instance instead.
//
// Provides a WAF Instance resource to create instance in the Web Application Firewall.
//
// For information about WAF and how to use it, see [What is Alibaba Cloud WAF](https://www.alibabacloud.com/help/doc-detail/28517.htm).
//
// > **NOTE:** Available in 1.83.0+ .
//
// ## Import
//
// WAF instance can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:waf/instance:Instance default waf-cn-132435
//
// ```
type Instance struct {
	pulumi.CustomResourceState

	// Specify whether big screen is supported. Valid values: ["0", "1"]. "0" for false and "1" for true.
	BigScreen pulumi.StringOutput `pulumi:"bigScreen"`
	// Specify the number of exclusive WAF IP addresses.
	ExclusiveIpPackage pulumi.StringOutput `pulumi:"exclusiveIpPackage"`
	// The extra bandwidth. Unit: Mbit/s.
	ExtBandwidth pulumi.StringOutput `pulumi:"extBandwidth"`
	// The number of extra domains.
	ExtDomainPackage pulumi.StringOutput `pulumi:"extDomainPackage"`
	// Log storage size. Unit: T. Valid values: [3, 5, 10, 20, 50].
	LogStorage pulumi.StringOutput `pulumi:"logStorage"`
	// Log storage period. Unit: day. Valid values: [180, 360].
	LogTime pulumi.StringOutput `pulumi:"logTime"`
	// Type of configuration change. Valid value: Upgrade.
	ModifyType pulumi.StringPtrOutput `pulumi:"modifyType"`
	// Subscription plan:
	// * China site customers can purchase the following versions of China Mainland region, valid values: ["version3", "version4", "version5"].
	// * China site customers can purchase the following versions of International region, valid values: ["versionProAsia", "versionBusinessAsia", "versionEnterpriseAsia"]
	// * International site customers can purchase the following versions of China Mainland region: ["versionProChina", "versionBusinessChina", "versionEnterpriseChina"]
	// * International site customers can purchase the following versions of International region: ["versionPro", "versionBusiness", "versionEnterprise"].
	PackageCode pulumi.StringOutput `pulumi:"packageCode"`
	// Service time of Web Application Firewall.
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// Specify whether professional service is supported. Valid values: ["true", "false"]
	PrefessionalService pulumi.StringOutput `pulumi:"prefessionalService"`
	// The instance region ID.
	Region pulumi.StringPtrOutput `pulumi:"region"`
	// Renewal period of WAF service. Unit: month
	RenewPeriod pulumi.IntPtrOutput `pulumi:"renewPeriod"`
	// Renewal status of WAF service. Valid values:
	// * AutoRenewal: The service time of WAF is renewed automatically.
	// * ManualRenewal (default): The service time of WAF is renewed manually.Specifies whether to configure a Layer-7 proxy, such as Anti-DDoS Pro or CDN, to filter the inbound traffic before it is forwarded to WAF. Valid values: "On" and "Off". Default to "Off".
	RenewalStatus pulumi.StringPtrOutput `pulumi:"renewalStatus"`
	// The resource group ID.
	ResourceGroupId pulumi.StringPtrOutput `pulumi:"resourceGroupId"`
	// The status of the instance.
	Status pulumi.IntOutput `pulumi:"status"`
	// Subscription of WAF service. Valid values: ["Subscription", "PayAsYouGo"].
	SubscriptionType pulumi.StringOutput `pulumi:"subscriptionType"`
	// Specify whether Log service is supported. Valid values: ["true", "false"]
	WafLog pulumi.StringOutput `pulumi:"wafLog"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BigScreen == nil {
		return nil, errors.New("invalid value for required argument 'BigScreen'")
	}
	if args.ExclusiveIpPackage == nil {
		return nil, errors.New("invalid value for required argument 'ExclusiveIpPackage'")
	}
	if args.ExtBandwidth == nil {
		return nil, errors.New("invalid value for required argument 'ExtBandwidth'")
	}
	if args.ExtDomainPackage == nil {
		return nil, errors.New("invalid value for required argument 'ExtDomainPackage'")
	}
	if args.LogStorage == nil {
		return nil, errors.New("invalid value for required argument 'LogStorage'")
	}
	if args.LogTime == nil {
		return nil, errors.New("invalid value for required argument 'LogTime'")
	}
	if args.PackageCode == nil {
		return nil, errors.New("invalid value for required argument 'PackageCode'")
	}
	if args.PrefessionalService == nil {
		return nil, errors.New("invalid value for required argument 'PrefessionalService'")
	}
	if args.SubscriptionType == nil {
		return nil, errors.New("invalid value for required argument 'SubscriptionType'")
	}
	if args.WafLog == nil {
		return nil, errors.New("invalid value for required argument 'WafLog'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Instance
	err := ctx.RegisterResource("alicloud:waf/instance:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("alicloud:waf/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// Specify whether big screen is supported. Valid values: ["0", "1"]. "0" for false and "1" for true.
	BigScreen *string `pulumi:"bigScreen"`
	// Specify the number of exclusive WAF IP addresses.
	ExclusiveIpPackage *string `pulumi:"exclusiveIpPackage"`
	// The extra bandwidth. Unit: Mbit/s.
	ExtBandwidth *string `pulumi:"extBandwidth"`
	// The number of extra domains.
	ExtDomainPackage *string `pulumi:"extDomainPackage"`
	// Log storage size. Unit: T. Valid values: [3, 5, 10, 20, 50].
	LogStorage *string `pulumi:"logStorage"`
	// Log storage period. Unit: day. Valid values: [180, 360].
	LogTime *string `pulumi:"logTime"`
	// Type of configuration change. Valid value: Upgrade.
	ModifyType *string `pulumi:"modifyType"`
	// Subscription plan:
	// * China site customers can purchase the following versions of China Mainland region, valid values: ["version3", "version4", "version5"].
	// * China site customers can purchase the following versions of International region, valid values: ["versionProAsia", "versionBusinessAsia", "versionEnterpriseAsia"]
	// * International site customers can purchase the following versions of China Mainland region: ["versionProChina", "versionBusinessChina", "versionEnterpriseChina"]
	// * International site customers can purchase the following versions of International region: ["versionPro", "versionBusiness", "versionEnterprise"].
	PackageCode *string `pulumi:"packageCode"`
	// Service time of Web Application Firewall.
	Period *int `pulumi:"period"`
	// Specify whether professional service is supported. Valid values: ["true", "false"]
	PrefessionalService *string `pulumi:"prefessionalService"`
	// The instance region ID.
	Region *string `pulumi:"region"`
	// Renewal period of WAF service. Unit: month
	RenewPeriod *int `pulumi:"renewPeriod"`
	// Renewal status of WAF service. Valid values:
	// * AutoRenewal: The service time of WAF is renewed automatically.
	// * ManualRenewal (default): The service time of WAF is renewed manually.Specifies whether to configure a Layer-7 proxy, such as Anti-DDoS Pro or CDN, to filter the inbound traffic before it is forwarded to WAF. Valid values: "On" and "Off". Default to "Off".
	RenewalStatus *string `pulumi:"renewalStatus"`
	// The resource group ID.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The status of the instance.
	Status *int `pulumi:"status"`
	// Subscription of WAF service. Valid values: ["Subscription", "PayAsYouGo"].
	SubscriptionType *string `pulumi:"subscriptionType"`
	// Specify whether Log service is supported. Valid values: ["true", "false"]
	WafLog *string `pulumi:"wafLog"`
}

type InstanceState struct {
	// Specify whether big screen is supported. Valid values: ["0", "1"]. "0" for false and "1" for true.
	BigScreen pulumi.StringPtrInput
	// Specify the number of exclusive WAF IP addresses.
	ExclusiveIpPackage pulumi.StringPtrInput
	// The extra bandwidth. Unit: Mbit/s.
	ExtBandwidth pulumi.StringPtrInput
	// The number of extra domains.
	ExtDomainPackage pulumi.StringPtrInput
	// Log storage size. Unit: T. Valid values: [3, 5, 10, 20, 50].
	LogStorage pulumi.StringPtrInput
	// Log storage period. Unit: day. Valid values: [180, 360].
	LogTime pulumi.StringPtrInput
	// Type of configuration change. Valid value: Upgrade.
	ModifyType pulumi.StringPtrInput
	// Subscription plan:
	// * China site customers can purchase the following versions of China Mainland region, valid values: ["version3", "version4", "version5"].
	// * China site customers can purchase the following versions of International region, valid values: ["versionProAsia", "versionBusinessAsia", "versionEnterpriseAsia"]
	// * International site customers can purchase the following versions of China Mainland region: ["versionProChina", "versionBusinessChina", "versionEnterpriseChina"]
	// * International site customers can purchase the following versions of International region: ["versionPro", "versionBusiness", "versionEnterprise"].
	PackageCode pulumi.StringPtrInput
	// Service time of Web Application Firewall.
	Period pulumi.IntPtrInput
	// Specify whether professional service is supported. Valid values: ["true", "false"]
	PrefessionalService pulumi.StringPtrInput
	// The instance region ID.
	Region pulumi.StringPtrInput
	// Renewal period of WAF service. Unit: month
	RenewPeriod pulumi.IntPtrInput
	// Renewal status of WAF service. Valid values:
	// * AutoRenewal: The service time of WAF is renewed automatically.
	// * ManualRenewal (default): The service time of WAF is renewed manually.Specifies whether to configure a Layer-7 proxy, such as Anti-DDoS Pro or CDN, to filter the inbound traffic before it is forwarded to WAF. Valid values: "On" and "Off". Default to "Off".
	RenewalStatus pulumi.StringPtrInput
	// The resource group ID.
	ResourceGroupId pulumi.StringPtrInput
	// The status of the instance.
	Status pulumi.IntPtrInput
	// Subscription of WAF service. Valid values: ["Subscription", "PayAsYouGo"].
	SubscriptionType pulumi.StringPtrInput
	// Specify whether Log service is supported. Valid values: ["true", "false"]
	WafLog pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// Specify whether big screen is supported. Valid values: ["0", "1"]. "0" for false and "1" for true.
	BigScreen string `pulumi:"bigScreen"`
	// Specify the number of exclusive WAF IP addresses.
	ExclusiveIpPackage string `pulumi:"exclusiveIpPackage"`
	// The extra bandwidth. Unit: Mbit/s.
	ExtBandwidth string `pulumi:"extBandwidth"`
	// The number of extra domains.
	ExtDomainPackage string `pulumi:"extDomainPackage"`
	// Log storage size. Unit: T. Valid values: [3, 5, 10, 20, 50].
	LogStorage string `pulumi:"logStorage"`
	// Log storage period. Unit: day. Valid values: [180, 360].
	LogTime string `pulumi:"logTime"`
	// Type of configuration change. Valid value: Upgrade.
	ModifyType *string `pulumi:"modifyType"`
	// Subscription plan:
	// * China site customers can purchase the following versions of China Mainland region, valid values: ["version3", "version4", "version5"].
	// * China site customers can purchase the following versions of International region, valid values: ["versionProAsia", "versionBusinessAsia", "versionEnterpriseAsia"]
	// * International site customers can purchase the following versions of China Mainland region: ["versionProChina", "versionBusinessChina", "versionEnterpriseChina"]
	// * International site customers can purchase the following versions of International region: ["versionPro", "versionBusiness", "versionEnterprise"].
	PackageCode string `pulumi:"packageCode"`
	// Service time of Web Application Firewall.
	Period *int `pulumi:"period"`
	// Specify whether professional service is supported. Valid values: ["true", "false"]
	PrefessionalService string `pulumi:"prefessionalService"`
	// The instance region ID.
	Region *string `pulumi:"region"`
	// Renewal period of WAF service. Unit: month
	RenewPeriod *int `pulumi:"renewPeriod"`
	// Renewal status of WAF service. Valid values:
	// * AutoRenewal: The service time of WAF is renewed automatically.
	// * ManualRenewal (default): The service time of WAF is renewed manually.Specifies whether to configure a Layer-7 proxy, such as Anti-DDoS Pro or CDN, to filter the inbound traffic before it is forwarded to WAF. Valid values: "On" and "Off". Default to "Off".
	RenewalStatus *string `pulumi:"renewalStatus"`
	// The resource group ID.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// Subscription of WAF service. Valid values: ["Subscription", "PayAsYouGo"].
	SubscriptionType string `pulumi:"subscriptionType"`
	// Specify whether Log service is supported. Valid values: ["true", "false"]
	WafLog string `pulumi:"wafLog"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// Specify whether big screen is supported. Valid values: ["0", "1"]. "0" for false and "1" for true.
	BigScreen pulumi.StringInput
	// Specify the number of exclusive WAF IP addresses.
	ExclusiveIpPackage pulumi.StringInput
	// The extra bandwidth. Unit: Mbit/s.
	ExtBandwidth pulumi.StringInput
	// The number of extra domains.
	ExtDomainPackage pulumi.StringInput
	// Log storage size. Unit: T. Valid values: [3, 5, 10, 20, 50].
	LogStorage pulumi.StringInput
	// Log storage period. Unit: day. Valid values: [180, 360].
	LogTime pulumi.StringInput
	// Type of configuration change. Valid value: Upgrade.
	ModifyType pulumi.StringPtrInput
	// Subscription plan:
	// * China site customers can purchase the following versions of China Mainland region, valid values: ["version3", "version4", "version5"].
	// * China site customers can purchase the following versions of International region, valid values: ["versionProAsia", "versionBusinessAsia", "versionEnterpriseAsia"]
	// * International site customers can purchase the following versions of China Mainland region: ["versionProChina", "versionBusinessChina", "versionEnterpriseChina"]
	// * International site customers can purchase the following versions of International region: ["versionPro", "versionBusiness", "versionEnterprise"].
	PackageCode pulumi.StringInput
	// Service time of Web Application Firewall.
	Period pulumi.IntPtrInput
	// Specify whether professional service is supported. Valid values: ["true", "false"]
	PrefessionalService pulumi.StringInput
	// The instance region ID.
	Region pulumi.StringPtrInput
	// Renewal period of WAF service. Unit: month
	RenewPeriod pulumi.IntPtrInput
	// Renewal status of WAF service. Valid values:
	// * AutoRenewal: The service time of WAF is renewed automatically.
	// * ManualRenewal (default): The service time of WAF is renewed manually.Specifies whether to configure a Layer-7 proxy, such as Anti-DDoS Pro or CDN, to filter the inbound traffic before it is forwarded to WAF. Valid values: "On" and "Off". Default to "Off".
	RenewalStatus pulumi.StringPtrInput
	// The resource group ID.
	ResourceGroupId pulumi.StringPtrInput
	// Subscription of WAF service. Valid values: ["Subscription", "PayAsYouGo"].
	SubscriptionType pulumi.StringInput
	// Specify whether Log service is supported. Valid values: ["true", "false"]
	WafLog pulumi.StringInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

func (i *Instance) ToOutput(ctx context.Context) pulumix.Output[*Instance] {
	return pulumix.Output[*Instance]{
		OutputState: i.ToInstanceOutputWithContext(ctx).OutputState,
	}
}

// InstanceArrayInput is an input type that accepts InstanceArray and InstanceArrayOutput values.
// You can construct a concrete instance of `InstanceArrayInput` via:
//
//	InstanceArray{ InstanceArgs{...} }
type InstanceArrayInput interface {
	pulumi.Input

	ToInstanceArrayOutput() InstanceArrayOutput
	ToInstanceArrayOutputWithContext(context.Context) InstanceArrayOutput
}

type InstanceArray []InstanceInput

func (InstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (i InstanceArray) ToInstanceArrayOutput() InstanceArrayOutput {
	return i.ToInstanceArrayOutputWithContext(context.Background())
}

func (i InstanceArray) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceArrayOutput)
}

func (i InstanceArray) ToOutput(ctx context.Context) pulumix.Output[[]*Instance] {
	return pulumix.Output[[]*Instance]{
		OutputState: i.ToInstanceArrayOutputWithContext(ctx).OutputState,
	}
}

// InstanceMapInput is an input type that accepts InstanceMap and InstanceMapOutput values.
// You can construct a concrete instance of `InstanceMapInput` via:
//
//	InstanceMap{ "key": InstanceArgs{...} }
type InstanceMapInput interface {
	pulumi.Input

	ToInstanceMapOutput() InstanceMapOutput
	ToInstanceMapOutputWithContext(context.Context) InstanceMapOutput
}

type InstanceMap map[string]InstanceInput

func (InstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (i InstanceMap) ToInstanceMapOutput() InstanceMapOutput {
	return i.ToInstanceMapOutputWithContext(context.Background())
}

func (i InstanceMap) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMapOutput)
}

func (i InstanceMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Instance] {
	return pulumix.Output[map[string]*Instance]{
		OutputState: i.ToInstanceMapOutputWithContext(ctx).OutputState,
	}
}

type InstanceOutput struct{ *pulumi.OutputState }

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

func (o InstanceOutput) ToOutput(ctx context.Context) pulumix.Output[*Instance] {
	return pulumix.Output[*Instance]{
		OutputState: o.OutputState,
	}
}

// Specify whether big screen is supported. Valid values: ["0", "1"]. "0" for false and "1" for true.
func (o InstanceOutput) BigScreen() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.BigScreen }).(pulumi.StringOutput)
}

// Specify the number of exclusive WAF IP addresses.
func (o InstanceOutput) ExclusiveIpPackage() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ExclusiveIpPackage }).(pulumi.StringOutput)
}

// The extra bandwidth. Unit: Mbit/s.
func (o InstanceOutput) ExtBandwidth() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ExtBandwidth }).(pulumi.StringOutput)
}

// The number of extra domains.
func (o InstanceOutput) ExtDomainPackage() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ExtDomainPackage }).(pulumi.StringOutput)
}

// Log storage size. Unit: T. Valid values: [3, 5, 10, 20, 50].
func (o InstanceOutput) LogStorage() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.LogStorage }).(pulumi.StringOutput)
}

// Log storage period. Unit: day. Valid values: [180, 360].
func (o InstanceOutput) LogTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.LogTime }).(pulumi.StringOutput)
}

// Type of configuration change. Valid value: Upgrade.
func (o InstanceOutput) ModifyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.ModifyType }).(pulumi.StringPtrOutput)
}

// Subscription plan:
// * China site customers can purchase the following versions of China Mainland region, valid values: ["version3", "version4", "version5"].
// * China site customers can purchase the following versions of International region, valid values: ["versionProAsia", "versionBusinessAsia", "versionEnterpriseAsia"]
// * International site customers can purchase the following versions of China Mainland region: ["versionProChina", "versionBusinessChina", "versionEnterpriseChina"]
// * International site customers can purchase the following versions of International region: ["versionPro", "versionBusiness", "versionEnterprise"].
func (o InstanceOutput) PackageCode() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.PackageCode }).(pulumi.StringOutput)
}

// Service time of Web Application Firewall.
func (o InstanceOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.Period }).(pulumi.IntPtrOutput)
}

// Specify whether professional service is supported. Valid values: ["true", "false"]
func (o InstanceOutput) PrefessionalService() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.PrefessionalService }).(pulumi.StringOutput)
}

// The instance region ID.
func (o InstanceOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.Region }).(pulumi.StringPtrOutput)
}

// Renewal period of WAF service. Unit: month
func (o InstanceOutput) RenewPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.RenewPeriod }).(pulumi.IntPtrOutput)
}

// Renewal status of WAF service. Valid values:
// * AutoRenewal: The service time of WAF is renewed automatically.
// * ManualRenewal (default): The service time of WAF is renewed manually.Specifies whether to configure a Layer-7 proxy, such as Anti-DDoS Pro or CDN, to filter the inbound traffic before it is forwarded to WAF. Valid values: "On" and "Off". Default to "Off".
func (o InstanceOutput) RenewalStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.RenewalStatus }).(pulumi.StringPtrOutput)
}

// The resource group ID.
func (o InstanceOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

// The status of the instance.
func (o InstanceOutput) Status() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.Status }).(pulumi.IntOutput)
}

// Subscription of WAF service. Valid values: ["Subscription", "PayAsYouGo"].
func (o InstanceOutput) SubscriptionType() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.SubscriptionType }).(pulumi.StringOutput)
}

// Specify whether Log service is supported. Valid values: ["true", "false"]
func (o InstanceOutput) WafLog() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.WafLog }).(pulumi.StringOutput)
}

type InstanceArrayOutput struct{ *pulumi.OutputState }

func (InstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (o InstanceArrayOutput) ToInstanceArrayOutput() InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Instance] {
	return pulumix.Output[[]*Instance]{
		OutputState: o.OutputState,
	}
}

func (o InstanceArrayOutput) Index(i pulumi.IntInput) InstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].([]*Instance)[vs[1].(int)]
	}).(InstanceOutput)
}

type InstanceMapOutput struct{ *pulumi.OutputState }

func (InstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (o InstanceMapOutput) ToInstanceMapOutput() InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Instance] {
	return pulumix.Output[map[string]*Instance]{
		OutputState: o.OutputState,
	}
}

func (o InstanceMapOutput) MapIndex(k pulumi.StringInput) InstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].(map[string]*Instance)[vs[1].(string)]
	}).(InstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceInput)(nil)).Elem(), &Instance{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceArrayInput)(nil)).Elem(), InstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMapInput)(nil)).Elem(), InstanceMap{})
	pulumi.RegisterOutputType(InstanceOutput{})
	pulumi.RegisterOutputType(InstanceArrayOutput{})
	pulumi.RegisterOutputType(InstanceMapOutput{})
}

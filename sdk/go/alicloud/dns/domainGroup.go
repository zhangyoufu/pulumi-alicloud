// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Alidns Domain Group resource. For information about Alidns Domain Group and how to use it, see [What is Resource Alidns Domain Group](https://www.alibabacloud.com/help/en/doc-detail/29762.htm).
//
// > **NOTE:** Available since v1.84.0.
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
//			_, err := dns.NewDomainGroup(ctx, "default", &dns.DomainGroupArgs{
//				DomainGroupName: pulumi.String("tf-example"),
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
// Alidns domain group can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:dns/domainGroup:DomainGroup example 0932eb3ddee7499085c4d13d45*****
// ```
type DomainGroup struct {
	pulumi.CustomResourceState

	// The Name of the domain group. The `domainGroupName` is required when the value of the `groupName`  is Empty.
	DomainGroupName pulumi.StringOutput `pulumi:"domainGroupName"`
	// The Name of the domain group. The `groupName` has been deprecated from provider version 1.97.0. Please use `domainGroupName` instead.
	//
	// Deprecated: Field 'group_name' has been deprecated from version 1.97.0. Use 'domain_group_name' instead.
	GroupName pulumi.StringOutput `pulumi:"groupName"`
	// User language.
	Lang pulumi.StringPtrOutput `pulumi:"lang"`
}

// NewDomainGroup registers a new resource with the given unique name, arguments, and options.
func NewDomainGroup(ctx *pulumi.Context,
	name string, args *DomainGroupArgs, opts ...pulumi.ResourceOption) (*DomainGroup, error) {
	if args == nil {
		args = &DomainGroupArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DomainGroup
	err := ctx.RegisterResource("alicloud:dns/domainGroup:DomainGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainGroup gets an existing DomainGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainGroupState, opts ...pulumi.ResourceOption) (*DomainGroup, error) {
	var resource DomainGroup
	err := ctx.ReadResource("alicloud:dns/domainGroup:DomainGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainGroup resources.
type domainGroupState struct {
	// The Name of the domain group. The `domainGroupName` is required when the value of the `groupName`  is Empty.
	DomainGroupName *string `pulumi:"domainGroupName"`
	// The Name of the domain group. The `groupName` has been deprecated from provider version 1.97.0. Please use `domainGroupName` instead.
	//
	// Deprecated: Field 'group_name' has been deprecated from version 1.97.0. Use 'domain_group_name' instead.
	GroupName *string `pulumi:"groupName"`
	// User language.
	Lang *string `pulumi:"lang"`
}

type DomainGroupState struct {
	// The Name of the domain group. The `domainGroupName` is required when the value of the `groupName`  is Empty.
	DomainGroupName pulumi.StringPtrInput
	// The Name of the domain group. The `groupName` has been deprecated from provider version 1.97.0. Please use `domainGroupName` instead.
	//
	// Deprecated: Field 'group_name' has been deprecated from version 1.97.0. Use 'domain_group_name' instead.
	GroupName pulumi.StringPtrInput
	// User language.
	Lang pulumi.StringPtrInput
}

func (DomainGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainGroupState)(nil)).Elem()
}

type domainGroupArgs struct {
	// The Name of the domain group. The `domainGroupName` is required when the value of the `groupName`  is Empty.
	DomainGroupName *string `pulumi:"domainGroupName"`
	// The Name of the domain group. The `groupName` has been deprecated from provider version 1.97.0. Please use `domainGroupName` instead.
	//
	// Deprecated: Field 'group_name' has been deprecated from version 1.97.0. Use 'domain_group_name' instead.
	GroupName *string `pulumi:"groupName"`
	// User language.
	Lang *string `pulumi:"lang"`
}

// The set of arguments for constructing a DomainGroup resource.
type DomainGroupArgs struct {
	// The Name of the domain group. The `domainGroupName` is required when the value of the `groupName`  is Empty.
	DomainGroupName pulumi.StringPtrInput
	// The Name of the domain group. The `groupName` has been deprecated from provider version 1.97.0. Please use `domainGroupName` instead.
	//
	// Deprecated: Field 'group_name' has been deprecated from version 1.97.0. Use 'domain_group_name' instead.
	GroupName pulumi.StringPtrInput
	// User language.
	Lang pulumi.StringPtrInput
}

func (DomainGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainGroupArgs)(nil)).Elem()
}

type DomainGroupInput interface {
	pulumi.Input

	ToDomainGroupOutput() DomainGroupOutput
	ToDomainGroupOutputWithContext(ctx context.Context) DomainGroupOutput
}

func (*DomainGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainGroup)(nil)).Elem()
}

func (i *DomainGroup) ToDomainGroupOutput() DomainGroupOutput {
	return i.ToDomainGroupOutputWithContext(context.Background())
}

func (i *DomainGroup) ToDomainGroupOutputWithContext(ctx context.Context) DomainGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainGroupOutput)
}

// DomainGroupArrayInput is an input type that accepts DomainGroupArray and DomainGroupArrayOutput values.
// You can construct a concrete instance of `DomainGroupArrayInput` via:
//
//	DomainGroupArray{ DomainGroupArgs{...} }
type DomainGroupArrayInput interface {
	pulumi.Input

	ToDomainGroupArrayOutput() DomainGroupArrayOutput
	ToDomainGroupArrayOutputWithContext(context.Context) DomainGroupArrayOutput
}

type DomainGroupArray []DomainGroupInput

func (DomainGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DomainGroup)(nil)).Elem()
}

func (i DomainGroupArray) ToDomainGroupArrayOutput() DomainGroupArrayOutput {
	return i.ToDomainGroupArrayOutputWithContext(context.Background())
}

func (i DomainGroupArray) ToDomainGroupArrayOutputWithContext(ctx context.Context) DomainGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainGroupArrayOutput)
}

// DomainGroupMapInput is an input type that accepts DomainGroupMap and DomainGroupMapOutput values.
// You can construct a concrete instance of `DomainGroupMapInput` via:
//
//	DomainGroupMap{ "key": DomainGroupArgs{...} }
type DomainGroupMapInput interface {
	pulumi.Input

	ToDomainGroupMapOutput() DomainGroupMapOutput
	ToDomainGroupMapOutputWithContext(context.Context) DomainGroupMapOutput
}

type DomainGroupMap map[string]DomainGroupInput

func (DomainGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DomainGroup)(nil)).Elem()
}

func (i DomainGroupMap) ToDomainGroupMapOutput() DomainGroupMapOutput {
	return i.ToDomainGroupMapOutputWithContext(context.Background())
}

func (i DomainGroupMap) ToDomainGroupMapOutputWithContext(ctx context.Context) DomainGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainGroupMapOutput)
}

type DomainGroupOutput struct{ *pulumi.OutputState }

func (DomainGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainGroup)(nil)).Elem()
}

func (o DomainGroupOutput) ToDomainGroupOutput() DomainGroupOutput {
	return o
}

func (o DomainGroupOutput) ToDomainGroupOutputWithContext(ctx context.Context) DomainGroupOutput {
	return o
}

// The Name of the domain group. The `domainGroupName` is required when the value of the `groupName`  is Empty.
func (o DomainGroupOutput) DomainGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainGroup) pulumi.StringOutput { return v.DomainGroupName }).(pulumi.StringOutput)
}

// The Name of the domain group. The `groupName` has been deprecated from provider version 1.97.0. Please use `domainGroupName` instead.
//
// Deprecated: Field 'group_name' has been deprecated from version 1.97.0. Use 'domain_group_name' instead.
func (o DomainGroupOutput) GroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainGroup) pulumi.StringOutput { return v.GroupName }).(pulumi.StringOutput)
}

// User language.
func (o DomainGroupOutput) Lang() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainGroup) pulumi.StringPtrOutput { return v.Lang }).(pulumi.StringPtrOutput)
}

type DomainGroupArrayOutput struct{ *pulumi.OutputState }

func (DomainGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DomainGroup)(nil)).Elem()
}

func (o DomainGroupArrayOutput) ToDomainGroupArrayOutput() DomainGroupArrayOutput {
	return o
}

func (o DomainGroupArrayOutput) ToDomainGroupArrayOutputWithContext(ctx context.Context) DomainGroupArrayOutput {
	return o
}

func (o DomainGroupArrayOutput) Index(i pulumi.IntInput) DomainGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DomainGroup {
		return vs[0].([]*DomainGroup)[vs[1].(int)]
	}).(DomainGroupOutput)
}

type DomainGroupMapOutput struct{ *pulumi.OutputState }

func (DomainGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DomainGroup)(nil)).Elem()
}

func (o DomainGroupMapOutput) ToDomainGroupMapOutput() DomainGroupMapOutput {
	return o
}

func (o DomainGroupMapOutput) ToDomainGroupMapOutputWithContext(ctx context.Context) DomainGroupMapOutput {
	return o
}

func (o DomainGroupMapOutput) MapIndex(k pulumi.StringInput) DomainGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DomainGroup {
		return vs[0].(map[string]*DomainGroup)[vs[1].(string)]
	}).(DomainGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DomainGroupInput)(nil)).Elem(), &DomainGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainGroupArrayInput)(nil)).Elem(), DomainGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainGroupMapInput)(nil)).Elem(), DomainGroupMap{})
	pulumi.RegisterOutputType(DomainGroupOutput{})
	pulumi.RegisterOutputType(DomainGroupArrayOutput{})
	pulumi.RegisterOutputType(DomainGroupMapOutput{})
}

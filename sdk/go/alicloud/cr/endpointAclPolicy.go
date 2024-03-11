// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cr

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a CR Endpoint Acl Policy resource.
//
// For information about CR Endpoint Acl Policy and how to use it, see [What is Endpoint Acl Policy](https://www.alibabacloud.com/help/doc-detail/145275.htm).
//
// > **NOTE:** Available since v1.139.0.
//
// ## Example Usage
//
// # Basic Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cr"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "tf-example"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			defaultRegistryEnterpriseInstance, err := cr.NewRegistryEnterpriseInstance(ctx, "defaultRegistryEnterpriseInstance", &cr.RegistryEnterpriseInstanceArgs{
//				PaymentType:   pulumi.String("Subscription"),
//				Period:        pulumi.Int(1),
//				RenewalStatus: pulumi.String("ManualRenewal"),
//				InstanceType:  pulumi.String("Advanced"),
//				InstanceName:  pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			defaultEndpointAclService := cr.GetEndpointAclServiceOutput(ctx, cr.GetEndpointAclServiceOutputArgs{
//				EndpointType: pulumi.String("internet"),
//				Enable:       pulumi.Bool(true),
//				InstanceId:   defaultRegistryEnterpriseInstance.ID(),
//				ModuleName:   pulumi.String("Registry"),
//			}, nil)
//			_, err = cr.NewEndpointAclPolicy(ctx, "defaultEndpointAclPolicy", &cr.EndpointAclPolicyArgs{
//				InstanceId: defaultEndpointAclService.ApplyT(func(defaultEndpointAclService cr.GetEndpointAclServiceResult) (*string, error) {
//					return &defaultEndpointAclService.InstanceId, nil
//				}).(pulumi.StringPtrOutput),
//				Entry:        pulumi.String("192.168.1.0/24"),
//				Description:  pulumi.String(name),
//				ModuleName:   pulumi.String("Registry"),
//				EndpointType: pulumi.String("internet"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// CR Endpoint Acl Policy can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:cr/endpointAclPolicy:EndpointAclPolicy example <instance_id>:<endpoint_type>:<entry>
// ```
type EndpointAclPolicy struct {
	pulumi.CustomResourceState

	// The description of the entry.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The type of endpoint. Valid values: `internet`.
	EndpointType pulumi.StringOutput `pulumi:"endpointType"`
	// The IP segment that allowed to access.
	Entry pulumi.StringOutput `pulumi:"entry"`
	// The ID of the CR Instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The module that needs to set the access policy. Valid values: `Registry`.
	ModuleName pulumi.StringPtrOutput `pulumi:"moduleName"`
}

// NewEndpointAclPolicy registers a new resource with the given unique name, arguments, and options.
func NewEndpointAclPolicy(ctx *pulumi.Context,
	name string, args *EndpointAclPolicyArgs, opts ...pulumi.ResourceOption) (*EndpointAclPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndpointType == nil {
		return nil, errors.New("invalid value for required argument 'EndpointType'")
	}
	if args.Entry == nil {
		return nil, errors.New("invalid value for required argument 'Entry'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EndpointAclPolicy
	err := ctx.RegisterResource("alicloud:cr/endpointAclPolicy:EndpointAclPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpointAclPolicy gets an existing EndpointAclPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpointAclPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointAclPolicyState, opts ...pulumi.ResourceOption) (*EndpointAclPolicy, error) {
	var resource EndpointAclPolicy
	err := ctx.ReadResource("alicloud:cr/endpointAclPolicy:EndpointAclPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EndpointAclPolicy resources.
type endpointAclPolicyState struct {
	// The description of the entry.
	Description *string `pulumi:"description"`
	// The type of endpoint. Valid values: `internet`.
	EndpointType *string `pulumi:"endpointType"`
	// The IP segment that allowed to access.
	Entry *string `pulumi:"entry"`
	// The ID of the CR Instance.
	InstanceId *string `pulumi:"instanceId"`
	// The module that needs to set the access policy. Valid values: `Registry`.
	ModuleName *string `pulumi:"moduleName"`
}

type EndpointAclPolicyState struct {
	// The description of the entry.
	Description pulumi.StringPtrInput
	// The type of endpoint. Valid values: `internet`.
	EndpointType pulumi.StringPtrInput
	// The IP segment that allowed to access.
	Entry pulumi.StringPtrInput
	// The ID of the CR Instance.
	InstanceId pulumi.StringPtrInput
	// The module that needs to set the access policy. Valid values: `Registry`.
	ModuleName pulumi.StringPtrInput
}

func (EndpointAclPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointAclPolicyState)(nil)).Elem()
}

type endpointAclPolicyArgs struct {
	// The description of the entry.
	Description *string `pulumi:"description"`
	// The type of endpoint. Valid values: `internet`.
	EndpointType string `pulumi:"endpointType"`
	// The IP segment that allowed to access.
	Entry string `pulumi:"entry"`
	// The ID of the CR Instance.
	InstanceId string `pulumi:"instanceId"`
	// The module that needs to set the access policy. Valid values: `Registry`.
	ModuleName *string `pulumi:"moduleName"`
}

// The set of arguments for constructing a EndpointAclPolicy resource.
type EndpointAclPolicyArgs struct {
	// The description of the entry.
	Description pulumi.StringPtrInput
	// The type of endpoint. Valid values: `internet`.
	EndpointType pulumi.StringInput
	// The IP segment that allowed to access.
	Entry pulumi.StringInput
	// The ID of the CR Instance.
	InstanceId pulumi.StringInput
	// The module that needs to set the access policy. Valid values: `Registry`.
	ModuleName pulumi.StringPtrInput
}

func (EndpointAclPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointAclPolicyArgs)(nil)).Elem()
}

type EndpointAclPolicyInput interface {
	pulumi.Input

	ToEndpointAclPolicyOutput() EndpointAclPolicyOutput
	ToEndpointAclPolicyOutputWithContext(ctx context.Context) EndpointAclPolicyOutput
}

func (*EndpointAclPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointAclPolicy)(nil)).Elem()
}

func (i *EndpointAclPolicy) ToEndpointAclPolicyOutput() EndpointAclPolicyOutput {
	return i.ToEndpointAclPolicyOutputWithContext(context.Background())
}

func (i *EndpointAclPolicy) ToEndpointAclPolicyOutputWithContext(ctx context.Context) EndpointAclPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointAclPolicyOutput)
}

// EndpointAclPolicyArrayInput is an input type that accepts EndpointAclPolicyArray and EndpointAclPolicyArrayOutput values.
// You can construct a concrete instance of `EndpointAclPolicyArrayInput` via:
//
//	EndpointAclPolicyArray{ EndpointAclPolicyArgs{...} }
type EndpointAclPolicyArrayInput interface {
	pulumi.Input

	ToEndpointAclPolicyArrayOutput() EndpointAclPolicyArrayOutput
	ToEndpointAclPolicyArrayOutputWithContext(context.Context) EndpointAclPolicyArrayOutput
}

type EndpointAclPolicyArray []EndpointAclPolicyInput

func (EndpointAclPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EndpointAclPolicy)(nil)).Elem()
}

func (i EndpointAclPolicyArray) ToEndpointAclPolicyArrayOutput() EndpointAclPolicyArrayOutput {
	return i.ToEndpointAclPolicyArrayOutputWithContext(context.Background())
}

func (i EndpointAclPolicyArray) ToEndpointAclPolicyArrayOutputWithContext(ctx context.Context) EndpointAclPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointAclPolicyArrayOutput)
}

// EndpointAclPolicyMapInput is an input type that accepts EndpointAclPolicyMap and EndpointAclPolicyMapOutput values.
// You can construct a concrete instance of `EndpointAclPolicyMapInput` via:
//
//	EndpointAclPolicyMap{ "key": EndpointAclPolicyArgs{...} }
type EndpointAclPolicyMapInput interface {
	pulumi.Input

	ToEndpointAclPolicyMapOutput() EndpointAclPolicyMapOutput
	ToEndpointAclPolicyMapOutputWithContext(context.Context) EndpointAclPolicyMapOutput
}

type EndpointAclPolicyMap map[string]EndpointAclPolicyInput

func (EndpointAclPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EndpointAclPolicy)(nil)).Elem()
}

func (i EndpointAclPolicyMap) ToEndpointAclPolicyMapOutput() EndpointAclPolicyMapOutput {
	return i.ToEndpointAclPolicyMapOutputWithContext(context.Background())
}

func (i EndpointAclPolicyMap) ToEndpointAclPolicyMapOutputWithContext(ctx context.Context) EndpointAclPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointAclPolicyMapOutput)
}

type EndpointAclPolicyOutput struct{ *pulumi.OutputState }

func (EndpointAclPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointAclPolicy)(nil)).Elem()
}

func (o EndpointAclPolicyOutput) ToEndpointAclPolicyOutput() EndpointAclPolicyOutput {
	return o
}

func (o EndpointAclPolicyOutput) ToEndpointAclPolicyOutputWithContext(ctx context.Context) EndpointAclPolicyOutput {
	return o
}

// The description of the entry.
func (o EndpointAclPolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointAclPolicy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The type of endpoint. Valid values: `internet`.
func (o EndpointAclPolicyOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointAclPolicy) pulumi.StringOutput { return v.EndpointType }).(pulumi.StringOutput)
}

// The IP segment that allowed to access.
func (o EndpointAclPolicyOutput) Entry() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointAclPolicy) pulumi.StringOutput { return v.Entry }).(pulumi.StringOutput)
}

// The ID of the CR Instance.
func (o EndpointAclPolicyOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointAclPolicy) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The module that needs to set the access policy. Valid values: `Registry`.
func (o EndpointAclPolicyOutput) ModuleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointAclPolicy) pulumi.StringPtrOutput { return v.ModuleName }).(pulumi.StringPtrOutput)
}

type EndpointAclPolicyArrayOutput struct{ *pulumi.OutputState }

func (EndpointAclPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EndpointAclPolicy)(nil)).Elem()
}

func (o EndpointAclPolicyArrayOutput) ToEndpointAclPolicyArrayOutput() EndpointAclPolicyArrayOutput {
	return o
}

func (o EndpointAclPolicyArrayOutput) ToEndpointAclPolicyArrayOutputWithContext(ctx context.Context) EndpointAclPolicyArrayOutput {
	return o
}

func (o EndpointAclPolicyArrayOutput) Index(i pulumi.IntInput) EndpointAclPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EndpointAclPolicy {
		return vs[0].([]*EndpointAclPolicy)[vs[1].(int)]
	}).(EndpointAclPolicyOutput)
}

type EndpointAclPolicyMapOutput struct{ *pulumi.OutputState }

func (EndpointAclPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EndpointAclPolicy)(nil)).Elem()
}

func (o EndpointAclPolicyMapOutput) ToEndpointAclPolicyMapOutput() EndpointAclPolicyMapOutput {
	return o
}

func (o EndpointAclPolicyMapOutput) ToEndpointAclPolicyMapOutputWithContext(ctx context.Context) EndpointAclPolicyMapOutput {
	return o
}

func (o EndpointAclPolicyMapOutput) MapIndex(k pulumi.StringInput) EndpointAclPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EndpointAclPolicy {
		return vs[0].(map[string]*EndpointAclPolicy)[vs[1].(string)]
	}).(EndpointAclPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointAclPolicyInput)(nil)).Elem(), &EndpointAclPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointAclPolicyArrayInput)(nil)).Elem(), EndpointAclPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointAclPolicyMapInput)(nil)).Elem(), EndpointAclPolicyMap{})
	pulumi.RegisterOutputType(EndpointAclPolicyOutput{})
	pulumi.RegisterOutputType(EndpointAclPolicyArrayOutput{})
	pulumi.RegisterOutputType(EndpointAclPolicyMapOutput{})
}

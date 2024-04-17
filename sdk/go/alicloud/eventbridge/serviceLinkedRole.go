// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eventbridge

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Using this data source can create Event Bridge service-linked roles(SLR). EventBridge may need to access another Alibaba Cloud service to implement a specific feature. In this case, EventBridge must assume a specific service-linked role, which is a Resource Access Management (RAM) role, to obtain permissions to access another Alibaba Cloud service.
//
// For information about Event Bridge service-linked roles(SLR) and how to use it, see [What is service-linked roles](https://www.alibabacloud.com/help/doc-detail/181425.htm).
//
// > **NOTE:** Available in v1.129.0+. After the version 1.142.0, the resource is renamed as `eventbridge.ServiceLinkedRole`.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/eventbridge"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := eventbridge.NewServiceLinkedRole(ctx, "service_linked_role", &eventbridge.ServiceLinkedRoleArgs{
//				ProductName: pulumi.String("AliyunServiceRoleForEventBridgeSourceRocketMQ"),
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
// Event Bridge service-linked roles(SLR) can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:eventbridge/serviceLinkedRole:ServiceLinkedRole example <product_name>
// ```
type ServiceLinkedRole struct {
	pulumi.CustomResourceState

	// The product name for SLR. EventBridge can automatically create the following service-linked roles:
	// Event source related: `AliyunServiceRoleForEventBridgeSendToMNS`,`AliyunServiceRoleForEventBridgeSourceRocketMQ`, `AliyunServiceRoleForEventBridgeSourceActionTrail`, `AliyunServiceRoleForEventBridgeSourceRabbitMQ`
	// Target related: `AliyunServiceRoleForEventBridgeConnectVPC`, `AliyunServiceRoleForEventBridgeSendToFC`, `AliyunServiceRoleForEventBridgeSendToSMS`, `AliyunServiceRoleForEventBridgeSendToDirectMail`, `AliyunServiceRoleForEventBridgeSendToRabbitMQ`, `AliyunServiceRoleForEventBridgeSendToRocketMQ`
	ProductName pulumi.StringOutput `pulumi:"productName"`
}

// NewServiceLinkedRole registers a new resource with the given unique name, arguments, and options.
func NewServiceLinkedRole(ctx *pulumi.Context,
	name string, args *ServiceLinkedRoleArgs, opts ...pulumi.ResourceOption) (*ServiceLinkedRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProductName == nil {
		return nil, errors.New("invalid value for required argument 'ProductName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceLinkedRole
	err := ctx.RegisterResource("alicloud:eventbridge/serviceLinkedRole:ServiceLinkedRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceLinkedRole gets an existing ServiceLinkedRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceLinkedRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceLinkedRoleState, opts ...pulumi.ResourceOption) (*ServiceLinkedRole, error) {
	var resource ServiceLinkedRole
	err := ctx.ReadResource("alicloud:eventbridge/serviceLinkedRole:ServiceLinkedRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceLinkedRole resources.
type serviceLinkedRoleState struct {
	// The product name for SLR. EventBridge can automatically create the following service-linked roles:
	// Event source related: `AliyunServiceRoleForEventBridgeSendToMNS`,`AliyunServiceRoleForEventBridgeSourceRocketMQ`, `AliyunServiceRoleForEventBridgeSourceActionTrail`, `AliyunServiceRoleForEventBridgeSourceRabbitMQ`
	// Target related: `AliyunServiceRoleForEventBridgeConnectVPC`, `AliyunServiceRoleForEventBridgeSendToFC`, `AliyunServiceRoleForEventBridgeSendToSMS`, `AliyunServiceRoleForEventBridgeSendToDirectMail`, `AliyunServiceRoleForEventBridgeSendToRabbitMQ`, `AliyunServiceRoleForEventBridgeSendToRocketMQ`
	ProductName *string `pulumi:"productName"`
}

type ServiceLinkedRoleState struct {
	// The product name for SLR. EventBridge can automatically create the following service-linked roles:
	// Event source related: `AliyunServiceRoleForEventBridgeSendToMNS`,`AliyunServiceRoleForEventBridgeSourceRocketMQ`, `AliyunServiceRoleForEventBridgeSourceActionTrail`, `AliyunServiceRoleForEventBridgeSourceRabbitMQ`
	// Target related: `AliyunServiceRoleForEventBridgeConnectVPC`, `AliyunServiceRoleForEventBridgeSendToFC`, `AliyunServiceRoleForEventBridgeSendToSMS`, `AliyunServiceRoleForEventBridgeSendToDirectMail`, `AliyunServiceRoleForEventBridgeSendToRabbitMQ`, `AliyunServiceRoleForEventBridgeSendToRocketMQ`
	ProductName pulumi.StringPtrInput
}

func (ServiceLinkedRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceLinkedRoleState)(nil)).Elem()
}

type serviceLinkedRoleArgs struct {
	// The product name for SLR. EventBridge can automatically create the following service-linked roles:
	// Event source related: `AliyunServiceRoleForEventBridgeSendToMNS`,`AliyunServiceRoleForEventBridgeSourceRocketMQ`, `AliyunServiceRoleForEventBridgeSourceActionTrail`, `AliyunServiceRoleForEventBridgeSourceRabbitMQ`
	// Target related: `AliyunServiceRoleForEventBridgeConnectVPC`, `AliyunServiceRoleForEventBridgeSendToFC`, `AliyunServiceRoleForEventBridgeSendToSMS`, `AliyunServiceRoleForEventBridgeSendToDirectMail`, `AliyunServiceRoleForEventBridgeSendToRabbitMQ`, `AliyunServiceRoleForEventBridgeSendToRocketMQ`
	ProductName string `pulumi:"productName"`
}

// The set of arguments for constructing a ServiceLinkedRole resource.
type ServiceLinkedRoleArgs struct {
	// The product name for SLR. EventBridge can automatically create the following service-linked roles:
	// Event source related: `AliyunServiceRoleForEventBridgeSendToMNS`,`AliyunServiceRoleForEventBridgeSourceRocketMQ`, `AliyunServiceRoleForEventBridgeSourceActionTrail`, `AliyunServiceRoleForEventBridgeSourceRabbitMQ`
	// Target related: `AliyunServiceRoleForEventBridgeConnectVPC`, `AliyunServiceRoleForEventBridgeSendToFC`, `AliyunServiceRoleForEventBridgeSendToSMS`, `AliyunServiceRoleForEventBridgeSendToDirectMail`, `AliyunServiceRoleForEventBridgeSendToRabbitMQ`, `AliyunServiceRoleForEventBridgeSendToRocketMQ`
	ProductName pulumi.StringInput
}

func (ServiceLinkedRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceLinkedRoleArgs)(nil)).Elem()
}

type ServiceLinkedRoleInput interface {
	pulumi.Input

	ToServiceLinkedRoleOutput() ServiceLinkedRoleOutput
	ToServiceLinkedRoleOutputWithContext(ctx context.Context) ServiceLinkedRoleOutput
}

func (*ServiceLinkedRole) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceLinkedRole)(nil)).Elem()
}

func (i *ServiceLinkedRole) ToServiceLinkedRoleOutput() ServiceLinkedRoleOutput {
	return i.ToServiceLinkedRoleOutputWithContext(context.Background())
}

func (i *ServiceLinkedRole) ToServiceLinkedRoleOutputWithContext(ctx context.Context) ServiceLinkedRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceLinkedRoleOutput)
}

// ServiceLinkedRoleArrayInput is an input type that accepts ServiceLinkedRoleArray and ServiceLinkedRoleArrayOutput values.
// You can construct a concrete instance of `ServiceLinkedRoleArrayInput` via:
//
//	ServiceLinkedRoleArray{ ServiceLinkedRoleArgs{...} }
type ServiceLinkedRoleArrayInput interface {
	pulumi.Input

	ToServiceLinkedRoleArrayOutput() ServiceLinkedRoleArrayOutput
	ToServiceLinkedRoleArrayOutputWithContext(context.Context) ServiceLinkedRoleArrayOutput
}

type ServiceLinkedRoleArray []ServiceLinkedRoleInput

func (ServiceLinkedRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceLinkedRole)(nil)).Elem()
}

func (i ServiceLinkedRoleArray) ToServiceLinkedRoleArrayOutput() ServiceLinkedRoleArrayOutput {
	return i.ToServiceLinkedRoleArrayOutputWithContext(context.Background())
}

func (i ServiceLinkedRoleArray) ToServiceLinkedRoleArrayOutputWithContext(ctx context.Context) ServiceLinkedRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceLinkedRoleArrayOutput)
}

// ServiceLinkedRoleMapInput is an input type that accepts ServiceLinkedRoleMap and ServiceLinkedRoleMapOutput values.
// You can construct a concrete instance of `ServiceLinkedRoleMapInput` via:
//
//	ServiceLinkedRoleMap{ "key": ServiceLinkedRoleArgs{...} }
type ServiceLinkedRoleMapInput interface {
	pulumi.Input

	ToServiceLinkedRoleMapOutput() ServiceLinkedRoleMapOutput
	ToServiceLinkedRoleMapOutputWithContext(context.Context) ServiceLinkedRoleMapOutput
}

type ServiceLinkedRoleMap map[string]ServiceLinkedRoleInput

func (ServiceLinkedRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceLinkedRole)(nil)).Elem()
}

func (i ServiceLinkedRoleMap) ToServiceLinkedRoleMapOutput() ServiceLinkedRoleMapOutput {
	return i.ToServiceLinkedRoleMapOutputWithContext(context.Background())
}

func (i ServiceLinkedRoleMap) ToServiceLinkedRoleMapOutputWithContext(ctx context.Context) ServiceLinkedRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceLinkedRoleMapOutput)
}

type ServiceLinkedRoleOutput struct{ *pulumi.OutputState }

func (ServiceLinkedRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceLinkedRole)(nil)).Elem()
}

func (o ServiceLinkedRoleOutput) ToServiceLinkedRoleOutput() ServiceLinkedRoleOutput {
	return o
}

func (o ServiceLinkedRoleOutput) ToServiceLinkedRoleOutputWithContext(ctx context.Context) ServiceLinkedRoleOutput {
	return o
}

// The product name for SLR. EventBridge can automatically create the following service-linked roles:
// Event source related: `AliyunServiceRoleForEventBridgeSendToMNS`,`AliyunServiceRoleForEventBridgeSourceRocketMQ`, `AliyunServiceRoleForEventBridgeSourceActionTrail`, `AliyunServiceRoleForEventBridgeSourceRabbitMQ`
// Target related: `AliyunServiceRoleForEventBridgeConnectVPC`, `AliyunServiceRoleForEventBridgeSendToFC`, `AliyunServiceRoleForEventBridgeSendToSMS`, `AliyunServiceRoleForEventBridgeSendToDirectMail`, `AliyunServiceRoleForEventBridgeSendToRabbitMQ`, `AliyunServiceRoleForEventBridgeSendToRocketMQ`
func (o ServiceLinkedRoleOutput) ProductName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceLinkedRole) pulumi.StringOutput { return v.ProductName }).(pulumi.StringOutput)
}

type ServiceLinkedRoleArrayOutput struct{ *pulumi.OutputState }

func (ServiceLinkedRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceLinkedRole)(nil)).Elem()
}

func (o ServiceLinkedRoleArrayOutput) ToServiceLinkedRoleArrayOutput() ServiceLinkedRoleArrayOutput {
	return o
}

func (o ServiceLinkedRoleArrayOutput) ToServiceLinkedRoleArrayOutputWithContext(ctx context.Context) ServiceLinkedRoleArrayOutput {
	return o
}

func (o ServiceLinkedRoleArrayOutput) Index(i pulumi.IntInput) ServiceLinkedRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceLinkedRole {
		return vs[0].([]*ServiceLinkedRole)[vs[1].(int)]
	}).(ServiceLinkedRoleOutput)
}

type ServiceLinkedRoleMapOutput struct{ *pulumi.OutputState }

func (ServiceLinkedRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceLinkedRole)(nil)).Elem()
}

func (o ServiceLinkedRoleMapOutput) ToServiceLinkedRoleMapOutput() ServiceLinkedRoleMapOutput {
	return o
}

func (o ServiceLinkedRoleMapOutput) ToServiceLinkedRoleMapOutputWithContext(ctx context.Context) ServiceLinkedRoleMapOutput {
	return o
}

func (o ServiceLinkedRoleMapOutput) MapIndex(k pulumi.StringInput) ServiceLinkedRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceLinkedRole {
		return vs[0].(map[string]*ServiceLinkedRole)[vs[1].(string)]
	}).(ServiceLinkedRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceLinkedRoleInput)(nil)).Elem(), &ServiceLinkedRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceLinkedRoleArrayInput)(nil)).Elem(), ServiceLinkedRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceLinkedRoleMapInput)(nil)).Elem(), ServiceLinkedRoleMap{})
	pulumi.RegisterOutputType(ServiceLinkedRoleOutput{})
	pulumi.RegisterOutputType(ServiceLinkedRoleArrayOutput{})
	pulumi.RegisterOutputType(ServiceLinkedRoleMapOutput{})
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package slb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Add a group of backend servers (ECS or ENI instance) to the Server Load Balancer or remove them from it.
//
// > **NOTE:** Available in 1.53.0+
//
// ## Block servers
//
// The servers mapping supports the following:
//
// * `serverId` - (Required) A list backend server ID (ECS instance ID).
// * `weight` - (Optional) Weight of the backend server. Valid value range: [0-100].
// * `type` - (Optional) Type of the backend server. Valid value `ecs`, `eni`, `eci`. Default to `ecs`. **NOTE:** From 1.170.0+, The `eci` is valid.
// * `serverIp` - (Optional, Computed, Available in 1.93.0+) ServerIp of the backend server. This parameter can be specified when the type is `eni`. `ecs` type currently does not support adding `serverIp` parameter.
//
// ## Import
//
// Load balancer backend server can be imported using the load balancer id.
//
// ```sh
// $ pulumi import alicloud:slb/backendServer:BackendServer example <load_balancer_id>
// ```
type BackendServer struct {
	pulumi.CustomResourceState

	// A list of instances to added backend server in the SLB. It contains three sub-fields as `Block server` follows.
	BackendServers BackendServerBackendServerArrayOutput `pulumi:"backendServers"`
	// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
	DeleteProtectionValidation pulumi.BoolPtrOutput `pulumi:"deleteProtectionValidation"`
	// ID of the load balancer.
	LoadBalancerId pulumi.StringOutput `pulumi:"loadBalancerId"`
}

// NewBackendServer registers a new resource with the given unique name, arguments, and options.
func NewBackendServer(ctx *pulumi.Context,
	name string, args *BackendServerArgs, opts ...pulumi.ResourceOption) (*BackendServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LoadBalancerId == nil {
		return nil, errors.New("invalid value for required argument 'LoadBalancerId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BackendServer
	err := ctx.RegisterResource("alicloud:slb/backendServer:BackendServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackendServer gets an existing BackendServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackendServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackendServerState, opts ...pulumi.ResourceOption) (*BackendServer, error) {
	var resource BackendServer
	err := ctx.ReadResource("alicloud:slb/backendServer:BackendServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackendServer resources.
type backendServerState struct {
	// A list of instances to added backend server in the SLB. It contains three sub-fields as `Block server` follows.
	BackendServers []BackendServerBackendServer `pulumi:"backendServers"`
	// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
	DeleteProtectionValidation *bool `pulumi:"deleteProtectionValidation"`
	// ID of the load balancer.
	LoadBalancerId *string `pulumi:"loadBalancerId"`
}

type BackendServerState struct {
	// A list of instances to added backend server in the SLB. It contains three sub-fields as `Block server` follows.
	BackendServers BackendServerBackendServerArrayInput
	// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
	DeleteProtectionValidation pulumi.BoolPtrInput
	// ID of the load balancer.
	LoadBalancerId pulumi.StringPtrInput
}

func (BackendServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*backendServerState)(nil)).Elem()
}

type backendServerArgs struct {
	// A list of instances to added backend server in the SLB. It contains three sub-fields as `Block server` follows.
	BackendServers []BackendServerBackendServer `pulumi:"backendServers"`
	// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
	DeleteProtectionValidation *bool `pulumi:"deleteProtectionValidation"`
	// ID of the load balancer.
	LoadBalancerId string `pulumi:"loadBalancerId"`
}

// The set of arguments for constructing a BackendServer resource.
type BackendServerArgs struct {
	// A list of instances to added backend server in the SLB. It contains three sub-fields as `Block server` follows.
	BackendServers BackendServerBackendServerArrayInput
	// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
	DeleteProtectionValidation pulumi.BoolPtrInput
	// ID of the load balancer.
	LoadBalancerId pulumi.StringInput
}

func (BackendServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backendServerArgs)(nil)).Elem()
}

type BackendServerInput interface {
	pulumi.Input

	ToBackendServerOutput() BackendServerOutput
	ToBackendServerOutputWithContext(ctx context.Context) BackendServerOutput
}

func (*BackendServer) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendServer)(nil)).Elem()
}

func (i *BackendServer) ToBackendServerOutput() BackendServerOutput {
	return i.ToBackendServerOutputWithContext(context.Background())
}

func (i *BackendServer) ToBackendServerOutputWithContext(ctx context.Context) BackendServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendServerOutput)
}

// BackendServerArrayInput is an input type that accepts BackendServerArray and BackendServerArrayOutput values.
// You can construct a concrete instance of `BackendServerArrayInput` via:
//
//	BackendServerArray{ BackendServerArgs{...} }
type BackendServerArrayInput interface {
	pulumi.Input

	ToBackendServerArrayOutput() BackendServerArrayOutput
	ToBackendServerArrayOutputWithContext(context.Context) BackendServerArrayOutput
}

type BackendServerArray []BackendServerInput

func (BackendServerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackendServer)(nil)).Elem()
}

func (i BackendServerArray) ToBackendServerArrayOutput() BackendServerArrayOutput {
	return i.ToBackendServerArrayOutputWithContext(context.Background())
}

func (i BackendServerArray) ToBackendServerArrayOutputWithContext(ctx context.Context) BackendServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendServerArrayOutput)
}

// BackendServerMapInput is an input type that accepts BackendServerMap and BackendServerMapOutput values.
// You can construct a concrete instance of `BackendServerMapInput` via:
//
//	BackendServerMap{ "key": BackendServerArgs{...} }
type BackendServerMapInput interface {
	pulumi.Input

	ToBackendServerMapOutput() BackendServerMapOutput
	ToBackendServerMapOutputWithContext(context.Context) BackendServerMapOutput
}

type BackendServerMap map[string]BackendServerInput

func (BackendServerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackendServer)(nil)).Elem()
}

func (i BackendServerMap) ToBackendServerMapOutput() BackendServerMapOutput {
	return i.ToBackendServerMapOutputWithContext(context.Background())
}

func (i BackendServerMap) ToBackendServerMapOutputWithContext(ctx context.Context) BackendServerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendServerMapOutput)
}

type BackendServerOutput struct{ *pulumi.OutputState }

func (BackendServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendServer)(nil)).Elem()
}

func (o BackendServerOutput) ToBackendServerOutput() BackendServerOutput {
	return o
}

func (o BackendServerOutput) ToBackendServerOutputWithContext(ctx context.Context) BackendServerOutput {
	return o
}

// A list of instances to added backend server in the SLB. It contains three sub-fields as `Block server` follows.
func (o BackendServerOutput) BackendServers() BackendServerBackendServerArrayOutput {
	return o.ApplyT(func(v *BackendServer) BackendServerBackendServerArrayOutput { return v.BackendServers }).(BackendServerBackendServerArrayOutput)
}

// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
func (o BackendServerOutput) DeleteProtectionValidation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BackendServer) pulumi.BoolPtrOutput { return v.DeleteProtectionValidation }).(pulumi.BoolPtrOutput)
}

// ID of the load balancer.
func (o BackendServerOutput) LoadBalancerId() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendServer) pulumi.StringOutput { return v.LoadBalancerId }).(pulumi.StringOutput)
}

type BackendServerArrayOutput struct{ *pulumi.OutputState }

func (BackendServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackendServer)(nil)).Elem()
}

func (o BackendServerArrayOutput) ToBackendServerArrayOutput() BackendServerArrayOutput {
	return o
}

func (o BackendServerArrayOutput) ToBackendServerArrayOutputWithContext(ctx context.Context) BackendServerArrayOutput {
	return o
}

func (o BackendServerArrayOutput) Index(i pulumi.IntInput) BackendServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BackendServer {
		return vs[0].([]*BackendServer)[vs[1].(int)]
	}).(BackendServerOutput)
}

type BackendServerMapOutput struct{ *pulumi.OutputState }

func (BackendServerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackendServer)(nil)).Elem()
}

func (o BackendServerMapOutput) ToBackendServerMapOutput() BackendServerMapOutput {
	return o
}

func (o BackendServerMapOutput) ToBackendServerMapOutputWithContext(ctx context.Context) BackendServerMapOutput {
	return o
}

func (o BackendServerMapOutput) MapIndex(k pulumi.StringInput) BackendServerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BackendServer {
		return vs[0].(map[string]*BackendServer)[vs[1].(string)]
	}).(BackendServerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackendServerInput)(nil)).Elem(), &BackendServer{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendServerArrayInput)(nil)).Elem(), BackendServerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendServerMapInput)(nil)).Elem(), BackendServerMap{})
	pulumi.RegisterOutputType(BackendServerOutput{})
	pulumi.RegisterOutputType(BackendServerArrayOutput{})
	pulumi.RegisterOutputType(BackendServerMapOutput{})
}

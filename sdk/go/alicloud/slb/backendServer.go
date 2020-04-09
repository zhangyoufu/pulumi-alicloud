// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package slb

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Add a group of backend servers (ECS or ENI instance) to the Server Load Balancer or remove them from it.
//
// > **NOTE:** Available in 1.53.0+
//
//
// ## Block servers
//
// The servers mapping supports the following:
//
// * `serverId` - (Required) A list backend server ID (ECS instance ID).
// * `weight` - (Optional) Weight of the backend server. Valid value range: [0-100].
// * `type` - (Optional) Type of the backend server. Valid value ecs, eni. Default to eni.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/slb_backend_server.html.markdown.
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
	if args == nil || args.LoadBalancerId == nil {
		return nil, errors.New("missing required argument 'LoadBalancerId'")
	}
	if args == nil {
		args = &BackendServerArgs{}
	}
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

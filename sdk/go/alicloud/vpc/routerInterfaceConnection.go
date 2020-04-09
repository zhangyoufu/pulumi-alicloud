// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package vpc

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a VPC router interface connection resource to connect two router interfaces which are in two different VPCs.
// After that, all of the two router interfaces will be active.
//
// > **NOTE:** At present, Router interface does not support changing opposite router interface, the connection delete action is only deactivating it to inactive, not modifying the connection to empty.
//
// > **NOTE:** If you want to changing opposite router interface, you can delete router interface and re-build them.
//
// > **NOTE:** A integrated router interface connection tunnel requires both InitiatingSide and AcceptingSide configuring opposite router interface.
//
// > **NOTE:** Please remember to add a `dependsOn` clause in the router interface connection from the InitiatingSide to the AcceptingSide, because the connection from the AcceptingSide to the InitiatingSide must be done first.
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/router_interface_connection.html.markdown.
type RouterInterfaceConnection struct {
	pulumi.CustomResourceState

	// One side router interface ID.
	InterfaceId pulumi.StringOutput `pulumi:"interfaceId"`
	// Another side router interface ID. It must belong the specified "oppositeInterfaceOwnerId" account.
	OppositeInterfaceId      pulumi.StringOutput `pulumi:"oppositeInterfaceId"`
	OppositeInterfaceOwnerId pulumi.StringOutput `pulumi:"oppositeInterfaceOwnerId"`
	// Another side router ID. It must belong the specified "oppositeInterfaceOwnerId" account. It is valid when field "oppositeInterfaceOwnerId" is specified.
	OppositeRouterId pulumi.StringOutput `pulumi:"oppositeRouterId"`
	// Another side router Type. Optional value: VRouter, VBR. It is valid when field "oppositeInterfaceOwnerId" is specified.
	OppositeRouterType pulumi.StringPtrOutput `pulumi:"oppositeRouterType"`
}

// NewRouterInterfaceConnection registers a new resource with the given unique name, arguments, and options.
func NewRouterInterfaceConnection(ctx *pulumi.Context,
	name string, args *RouterInterfaceConnectionArgs, opts ...pulumi.ResourceOption) (*RouterInterfaceConnection, error) {
	if args == nil || args.InterfaceId == nil {
		return nil, errors.New("missing required argument 'InterfaceId'")
	}
	if args == nil || args.OppositeInterfaceId == nil {
		return nil, errors.New("missing required argument 'OppositeInterfaceId'")
	}
	if args == nil {
		args = &RouterInterfaceConnectionArgs{}
	}
	var resource RouterInterfaceConnection
	err := ctx.RegisterResource("alicloud:vpc/routerInterfaceConnection:RouterInterfaceConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouterInterfaceConnection gets an existing RouterInterfaceConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouterInterfaceConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouterInterfaceConnectionState, opts ...pulumi.ResourceOption) (*RouterInterfaceConnection, error) {
	var resource RouterInterfaceConnection
	err := ctx.ReadResource("alicloud:vpc/routerInterfaceConnection:RouterInterfaceConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouterInterfaceConnection resources.
type routerInterfaceConnectionState struct {
	// One side router interface ID.
	InterfaceId *string `pulumi:"interfaceId"`
	// Another side router interface ID. It must belong the specified "oppositeInterfaceOwnerId" account.
	OppositeInterfaceId      *string `pulumi:"oppositeInterfaceId"`
	OppositeInterfaceOwnerId *string `pulumi:"oppositeInterfaceOwnerId"`
	// Another side router ID. It must belong the specified "oppositeInterfaceOwnerId" account. It is valid when field "oppositeInterfaceOwnerId" is specified.
	OppositeRouterId *string `pulumi:"oppositeRouterId"`
	// Another side router Type. Optional value: VRouter, VBR. It is valid when field "oppositeInterfaceOwnerId" is specified.
	OppositeRouterType *string `pulumi:"oppositeRouterType"`
}

type RouterInterfaceConnectionState struct {
	// One side router interface ID.
	InterfaceId pulumi.StringPtrInput
	// Another side router interface ID. It must belong the specified "oppositeInterfaceOwnerId" account.
	OppositeInterfaceId      pulumi.StringPtrInput
	OppositeInterfaceOwnerId pulumi.StringPtrInput
	// Another side router ID. It must belong the specified "oppositeInterfaceOwnerId" account. It is valid when field "oppositeInterfaceOwnerId" is specified.
	OppositeRouterId pulumi.StringPtrInput
	// Another side router Type. Optional value: VRouter, VBR. It is valid when field "oppositeInterfaceOwnerId" is specified.
	OppositeRouterType pulumi.StringPtrInput
}

func (RouterInterfaceConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*routerInterfaceConnectionState)(nil)).Elem()
}

type routerInterfaceConnectionArgs struct {
	// One side router interface ID.
	InterfaceId string `pulumi:"interfaceId"`
	// Another side router interface ID. It must belong the specified "oppositeInterfaceOwnerId" account.
	OppositeInterfaceId      string  `pulumi:"oppositeInterfaceId"`
	OppositeInterfaceOwnerId *string `pulumi:"oppositeInterfaceOwnerId"`
	// Another side router ID. It must belong the specified "oppositeInterfaceOwnerId" account. It is valid when field "oppositeInterfaceOwnerId" is specified.
	OppositeRouterId *string `pulumi:"oppositeRouterId"`
	// Another side router Type. Optional value: VRouter, VBR. It is valid when field "oppositeInterfaceOwnerId" is specified.
	OppositeRouterType *string `pulumi:"oppositeRouterType"`
}

// The set of arguments for constructing a RouterInterfaceConnection resource.
type RouterInterfaceConnectionArgs struct {
	// One side router interface ID.
	InterfaceId pulumi.StringInput
	// Another side router interface ID. It must belong the specified "oppositeInterfaceOwnerId" account.
	OppositeInterfaceId      pulumi.StringInput
	OppositeInterfaceOwnerId pulumi.StringPtrInput
	// Another side router ID. It must belong the specified "oppositeInterfaceOwnerId" account. It is valid when field "oppositeInterfaceOwnerId" is specified.
	OppositeRouterId pulumi.StringPtrInput
	// Another side router Type. Optional value: VRouter, VBR. It is valid when field "oppositeInterfaceOwnerId" is specified.
	OppositeRouterType pulumi.StringPtrInput
}

func (RouterInterfaceConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routerInterfaceConnectionArgs)(nil)).Elem()
}

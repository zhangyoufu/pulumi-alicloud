// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cen

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Cen Transit Router Multicast Domain Peer Member resource.
//
// For information about Cen Transit Router Multicast Domain Peer Member and how to use it, see [What is Transit Router Multicast Domain Peer Member](https://www.alibabacloud.com/help/en/cloud-enterprise-network/latest/api-cbn-2017-09-12-deregistertransitroutermulticastgroupmembers).
//
// > **NOTE:** Available since v1.195.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cen"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "tf_example"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			defaultRegion := "cn-hangzhou"
//			if param := cfg.Get("defaultRegion"); param != "" {
//				defaultRegion = param
//			}
//			peerRegion := "cn-beijing"
//			if param := cfg.Get("peerRegion"); param != "" {
//				peerRegion = param
//			}
//			_, err := alicloud.NewProvider(ctx, "hz", &alicloud.ProviderArgs{
//				Region: pulumi.String(defaultRegion),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = alicloud.NewProvider(ctx, "bj", &alicloud.ProviderArgs{
//				Region: pulumi.String(peerRegion),
//			})
//			if err != nil {
//				return err
//			}
//			defaultInstance, err := cen.NewInstance(ctx, "defaultInstance", &cen.InstanceArgs{
//				CenInstanceName: pulumi.String(name),
//				ProtectionLevel: pulumi.String("REDUCED"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultTransitRouter, err := cen.NewTransitRouter(ctx, "defaultTransitRouter", &cen.TransitRouterArgs{
//				CenId: defaultInstance.ID(),
//			}, pulumi.Provider(alicloud.Hz))
//			if err != nil {
//				return err
//			}
//			peerTransitRouter, err := cen.NewTransitRouter(ctx, "peerTransitRouter", &cen.TransitRouterArgs{
//				CenId: defaultTransitRouter.CenId,
//			}, pulumi.Provider(alicloud.Bj))
//			if err != nil {
//				return err
//			}
//			defaultTransitRouterMulticastDomain, err := cen.NewTransitRouterMulticastDomain(ctx, "defaultTransitRouterMulticastDomain", &cen.TransitRouterMulticastDomainArgs{
//				TransitRouterId:                  defaultTransitRouter.TransitRouterId,
//				TransitRouterMulticastDomainName: pulumi.String(name),
//			}, pulumi.Provider(alicloud.Hz))
//			if err != nil {
//				return err
//			}
//			peerTransitRouterMulticastDomain, err := cen.NewTransitRouterMulticastDomain(ctx, "peerTransitRouterMulticastDomain", &cen.TransitRouterMulticastDomainArgs{
//				TransitRouterId:                  peerTransitRouter.TransitRouterId,
//				TransitRouterMulticastDomainName: pulumi.String(name),
//			}, pulumi.Provider(alicloud.Bj))
//			if err != nil {
//				return err
//			}
//			_, err = cen.NewTransitRouterMulticastDomainPeerMember(ctx, "defaultTransitRouterMulticastDomainPeerMember", &cen.TransitRouterMulticastDomainPeerMemberArgs{
//				PeerTransitRouterMulticastDomainId: peerTransitRouterMulticastDomain.ID(),
//				TransitRouterMulticastDomainId:     defaultTransitRouterMulticastDomain.ID(),
//				GroupIpAddress:                     pulumi.String("239.1.1.1"),
//			}, pulumi.Provider(alicloud.Hz))
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
// Cen Transit Router Multicast Domain Peer Member can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:cen/transitRouterMulticastDomainPeerMember:TransitRouterMulticastDomainPeerMember example <transit_router_multicast_domain_id>:<group_ip_address>:<peer_transit_router_multicast_domain_id>
//
// ```
type TransitRouterMulticastDomainPeerMember struct {
	pulumi.CustomResourceState

	// Specifies whether only to precheck the request.
	DryRun pulumi.BoolPtrOutput `pulumi:"dryRun"`
	// The IP address of the multicast group to which the multicast member belongs. Value range: **224.0.0.1** to **239.255.255.254**.If the multicast group you specified does not exist in the current multicast domain, the system will automatically create a new multicast group for you in the current multicast domain.
	GroupIpAddress pulumi.StringOutput `pulumi:"groupIpAddress"`
	// The IDs of the inter-region multicast domains.
	PeerTransitRouterMulticastDomainId pulumi.StringOutput `pulumi:"peerTransitRouterMulticastDomainId"`
	// The status of the multicast resource. Valid values:
	// - Registering: being created
	// - Registered: available
	// - Deregistering: being deleted
	Status pulumi.StringOutput `pulumi:"status"`
	// The ID of the multicast domain to which the multicast member belongs.
	TransitRouterMulticastDomainId pulumi.StringOutput `pulumi:"transitRouterMulticastDomainId"`
}

// NewTransitRouterMulticastDomainPeerMember registers a new resource with the given unique name, arguments, and options.
func NewTransitRouterMulticastDomainPeerMember(ctx *pulumi.Context,
	name string, args *TransitRouterMulticastDomainPeerMemberArgs, opts ...pulumi.ResourceOption) (*TransitRouterMulticastDomainPeerMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupIpAddress == nil {
		return nil, errors.New("invalid value for required argument 'GroupIpAddress'")
	}
	if args.PeerTransitRouterMulticastDomainId == nil {
		return nil, errors.New("invalid value for required argument 'PeerTransitRouterMulticastDomainId'")
	}
	if args.TransitRouterMulticastDomainId == nil {
		return nil, errors.New("invalid value for required argument 'TransitRouterMulticastDomainId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TransitRouterMulticastDomainPeerMember
	err := ctx.RegisterResource("alicloud:cen/transitRouterMulticastDomainPeerMember:TransitRouterMulticastDomainPeerMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTransitRouterMulticastDomainPeerMember gets an existing TransitRouterMulticastDomainPeerMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTransitRouterMulticastDomainPeerMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TransitRouterMulticastDomainPeerMemberState, opts ...pulumi.ResourceOption) (*TransitRouterMulticastDomainPeerMember, error) {
	var resource TransitRouterMulticastDomainPeerMember
	err := ctx.ReadResource("alicloud:cen/transitRouterMulticastDomainPeerMember:TransitRouterMulticastDomainPeerMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TransitRouterMulticastDomainPeerMember resources.
type transitRouterMulticastDomainPeerMemberState struct {
	// Specifies whether only to precheck the request.
	DryRun *bool `pulumi:"dryRun"`
	// The IP address of the multicast group to which the multicast member belongs. Value range: **224.0.0.1** to **239.255.255.254**.If the multicast group you specified does not exist in the current multicast domain, the system will automatically create a new multicast group for you in the current multicast domain.
	GroupIpAddress *string `pulumi:"groupIpAddress"`
	// The IDs of the inter-region multicast domains.
	PeerTransitRouterMulticastDomainId *string `pulumi:"peerTransitRouterMulticastDomainId"`
	// The status of the multicast resource. Valid values:
	// - Registering: being created
	// - Registered: available
	// - Deregistering: being deleted
	Status *string `pulumi:"status"`
	// The ID of the multicast domain to which the multicast member belongs.
	TransitRouterMulticastDomainId *string `pulumi:"transitRouterMulticastDomainId"`
}

type TransitRouterMulticastDomainPeerMemberState struct {
	// Specifies whether only to precheck the request.
	DryRun pulumi.BoolPtrInput
	// The IP address of the multicast group to which the multicast member belongs. Value range: **224.0.0.1** to **239.255.255.254**.If the multicast group you specified does not exist in the current multicast domain, the system will automatically create a new multicast group for you in the current multicast domain.
	GroupIpAddress pulumi.StringPtrInput
	// The IDs of the inter-region multicast domains.
	PeerTransitRouterMulticastDomainId pulumi.StringPtrInput
	// The status of the multicast resource. Valid values:
	// - Registering: being created
	// - Registered: available
	// - Deregistering: being deleted
	Status pulumi.StringPtrInput
	// The ID of the multicast domain to which the multicast member belongs.
	TransitRouterMulticastDomainId pulumi.StringPtrInput
}

func (TransitRouterMulticastDomainPeerMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*transitRouterMulticastDomainPeerMemberState)(nil)).Elem()
}

type transitRouterMulticastDomainPeerMemberArgs struct {
	// Specifies whether only to precheck the request.
	DryRun *bool `pulumi:"dryRun"`
	// The IP address of the multicast group to which the multicast member belongs. Value range: **224.0.0.1** to **239.255.255.254**.If the multicast group you specified does not exist in the current multicast domain, the system will automatically create a new multicast group for you in the current multicast domain.
	GroupIpAddress string `pulumi:"groupIpAddress"`
	// The IDs of the inter-region multicast domains.
	PeerTransitRouterMulticastDomainId string `pulumi:"peerTransitRouterMulticastDomainId"`
	// The ID of the multicast domain to which the multicast member belongs.
	TransitRouterMulticastDomainId string `pulumi:"transitRouterMulticastDomainId"`
}

// The set of arguments for constructing a TransitRouterMulticastDomainPeerMember resource.
type TransitRouterMulticastDomainPeerMemberArgs struct {
	// Specifies whether only to precheck the request.
	DryRun pulumi.BoolPtrInput
	// The IP address of the multicast group to which the multicast member belongs. Value range: **224.0.0.1** to **239.255.255.254**.If the multicast group you specified does not exist in the current multicast domain, the system will automatically create a new multicast group for you in the current multicast domain.
	GroupIpAddress pulumi.StringInput
	// The IDs of the inter-region multicast domains.
	PeerTransitRouterMulticastDomainId pulumi.StringInput
	// The ID of the multicast domain to which the multicast member belongs.
	TransitRouterMulticastDomainId pulumi.StringInput
}

func (TransitRouterMulticastDomainPeerMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*transitRouterMulticastDomainPeerMemberArgs)(nil)).Elem()
}

type TransitRouterMulticastDomainPeerMemberInput interface {
	pulumi.Input

	ToTransitRouterMulticastDomainPeerMemberOutput() TransitRouterMulticastDomainPeerMemberOutput
	ToTransitRouterMulticastDomainPeerMemberOutputWithContext(ctx context.Context) TransitRouterMulticastDomainPeerMemberOutput
}

func (*TransitRouterMulticastDomainPeerMember) ElementType() reflect.Type {
	return reflect.TypeOf((**TransitRouterMulticastDomainPeerMember)(nil)).Elem()
}

func (i *TransitRouterMulticastDomainPeerMember) ToTransitRouterMulticastDomainPeerMemberOutput() TransitRouterMulticastDomainPeerMemberOutput {
	return i.ToTransitRouterMulticastDomainPeerMemberOutputWithContext(context.Background())
}

func (i *TransitRouterMulticastDomainPeerMember) ToTransitRouterMulticastDomainPeerMemberOutputWithContext(ctx context.Context) TransitRouterMulticastDomainPeerMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransitRouterMulticastDomainPeerMemberOutput)
}

// TransitRouterMulticastDomainPeerMemberArrayInput is an input type that accepts TransitRouterMulticastDomainPeerMemberArray and TransitRouterMulticastDomainPeerMemberArrayOutput values.
// You can construct a concrete instance of `TransitRouterMulticastDomainPeerMemberArrayInput` via:
//
//	TransitRouterMulticastDomainPeerMemberArray{ TransitRouterMulticastDomainPeerMemberArgs{...} }
type TransitRouterMulticastDomainPeerMemberArrayInput interface {
	pulumi.Input

	ToTransitRouterMulticastDomainPeerMemberArrayOutput() TransitRouterMulticastDomainPeerMemberArrayOutput
	ToTransitRouterMulticastDomainPeerMemberArrayOutputWithContext(context.Context) TransitRouterMulticastDomainPeerMemberArrayOutput
}

type TransitRouterMulticastDomainPeerMemberArray []TransitRouterMulticastDomainPeerMemberInput

func (TransitRouterMulticastDomainPeerMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TransitRouterMulticastDomainPeerMember)(nil)).Elem()
}

func (i TransitRouterMulticastDomainPeerMemberArray) ToTransitRouterMulticastDomainPeerMemberArrayOutput() TransitRouterMulticastDomainPeerMemberArrayOutput {
	return i.ToTransitRouterMulticastDomainPeerMemberArrayOutputWithContext(context.Background())
}

func (i TransitRouterMulticastDomainPeerMemberArray) ToTransitRouterMulticastDomainPeerMemberArrayOutputWithContext(ctx context.Context) TransitRouterMulticastDomainPeerMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransitRouterMulticastDomainPeerMemberArrayOutput)
}

// TransitRouterMulticastDomainPeerMemberMapInput is an input type that accepts TransitRouterMulticastDomainPeerMemberMap and TransitRouterMulticastDomainPeerMemberMapOutput values.
// You can construct a concrete instance of `TransitRouterMulticastDomainPeerMemberMapInput` via:
//
//	TransitRouterMulticastDomainPeerMemberMap{ "key": TransitRouterMulticastDomainPeerMemberArgs{...} }
type TransitRouterMulticastDomainPeerMemberMapInput interface {
	pulumi.Input

	ToTransitRouterMulticastDomainPeerMemberMapOutput() TransitRouterMulticastDomainPeerMemberMapOutput
	ToTransitRouterMulticastDomainPeerMemberMapOutputWithContext(context.Context) TransitRouterMulticastDomainPeerMemberMapOutput
}

type TransitRouterMulticastDomainPeerMemberMap map[string]TransitRouterMulticastDomainPeerMemberInput

func (TransitRouterMulticastDomainPeerMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TransitRouterMulticastDomainPeerMember)(nil)).Elem()
}

func (i TransitRouterMulticastDomainPeerMemberMap) ToTransitRouterMulticastDomainPeerMemberMapOutput() TransitRouterMulticastDomainPeerMemberMapOutput {
	return i.ToTransitRouterMulticastDomainPeerMemberMapOutputWithContext(context.Background())
}

func (i TransitRouterMulticastDomainPeerMemberMap) ToTransitRouterMulticastDomainPeerMemberMapOutputWithContext(ctx context.Context) TransitRouterMulticastDomainPeerMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransitRouterMulticastDomainPeerMemberMapOutput)
}

type TransitRouterMulticastDomainPeerMemberOutput struct{ *pulumi.OutputState }

func (TransitRouterMulticastDomainPeerMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TransitRouterMulticastDomainPeerMember)(nil)).Elem()
}

func (o TransitRouterMulticastDomainPeerMemberOutput) ToTransitRouterMulticastDomainPeerMemberOutput() TransitRouterMulticastDomainPeerMemberOutput {
	return o
}

func (o TransitRouterMulticastDomainPeerMemberOutput) ToTransitRouterMulticastDomainPeerMemberOutputWithContext(ctx context.Context) TransitRouterMulticastDomainPeerMemberOutput {
	return o
}

// Specifies whether only to precheck the request.
func (o TransitRouterMulticastDomainPeerMemberOutput) DryRun() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TransitRouterMulticastDomainPeerMember) pulumi.BoolPtrOutput { return v.DryRun }).(pulumi.BoolPtrOutput)
}

// The IP address of the multicast group to which the multicast member belongs. Value range: **224.0.0.1** to **239.255.255.254**.If the multicast group you specified does not exist in the current multicast domain, the system will automatically create a new multicast group for you in the current multicast domain.
func (o TransitRouterMulticastDomainPeerMemberOutput) GroupIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitRouterMulticastDomainPeerMember) pulumi.StringOutput { return v.GroupIpAddress }).(pulumi.StringOutput)
}

// The IDs of the inter-region multicast domains.
func (o TransitRouterMulticastDomainPeerMemberOutput) PeerTransitRouterMulticastDomainId() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitRouterMulticastDomainPeerMember) pulumi.StringOutput {
		return v.PeerTransitRouterMulticastDomainId
	}).(pulumi.StringOutput)
}

// The status of the multicast resource. Valid values:
// - Registering: being created
// - Registered: available
// - Deregistering: being deleted
func (o TransitRouterMulticastDomainPeerMemberOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitRouterMulticastDomainPeerMember) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The ID of the multicast domain to which the multicast member belongs.
func (o TransitRouterMulticastDomainPeerMemberOutput) TransitRouterMulticastDomainId() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitRouterMulticastDomainPeerMember) pulumi.StringOutput {
		return v.TransitRouterMulticastDomainId
	}).(pulumi.StringOutput)
}

type TransitRouterMulticastDomainPeerMemberArrayOutput struct{ *pulumi.OutputState }

func (TransitRouterMulticastDomainPeerMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TransitRouterMulticastDomainPeerMember)(nil)).Elem()
}

func (o TransitRouterMulticastDomainPeerMemberArrayOutput) ToTransitRouterMulticastDomainPeerMemberArrayOutput() TransitRouterMulticastDomainPeerMemberArrayOutput {
	return o
}

func (o TransitRouterMulticastDomainPeerMemberArrayOutput) ToTransitRouterMulticastDomainPeerMemberArrayOutputWithContext(ctx context.Context) TransitRouterMulticastDomainPeerMemberArrayOutput {
	return o
}

func (o TransitRouterMulticastDomainPeerMemberArrayOutput) Index(i pulumi.IntInput) TransitRouterMulticastDomainPeerMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TransitRouterMulticastDomainPeerMember {
		return vs[0].([]*TransitRouterMulticastDomainPeerMember)[vs[1].(int)]
	}).(TransitRouterMulticastDomainPeerMemberOutput)
}

type TransitRouterMulticastDomainPeerMemberMapOutput struct{ *pulumi.OutputState }

func (TransitRouterMulticastDomainPeerMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TransitRouterMulticastDomainPeerMember)(nil)).Elem()
}

func (o TransitRouterMulticastDomainPeerMemberMapOutput) ToTransitRouterMulticastDomainPeerMemberMapOutput() TransitRouterMulticastDomainPeerMemberMapOutput {
	return o
}

func (o TransitRouterMulticastDomainPeerMemberMapOutput) ToTransitRouterMulticastDomainPeerMemberMapOutputWithContext(ctx context.Context) TransitRouterMulticastDomainPeerMemberMapOutput {
	return o
}

func (o TransitRouterMulticastDomainPeerMemberMapOutput) MapIndex(k pulumi.StringInput) TransitRouterMulticastDomainPeerMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TransitRouterMulticastDomainPeerMember {
		return vs[0].(map[string]*TransitRouterMulticastDomainPeerMember)[vs[1].(string)]
	}).(TransitRouterMulticastDomainPeerMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TransitRouterMulticastDomainPeerMemberInput)(nil)).Elem(), &TransitRouterMulticastDomainPeerMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransitRouterMulticastDomainPeerMemberArrayInput)(nil)).Elem(), TransitRouterMulticastDomainPeerMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransitRouterMulticastDomainPeerMemberMapInput)(nil)).Elem(), TransitRouterMulticastDomainPeerMemberMap{})
	pulumi.RegisterOutputType(TransitRouterMulticastDomainPeerMemberOutput{})
	pulumi.RegisterOutputType(TransitRouterMulticastDomainPeerMemberArrayOutput{})
	pulumi.RegisterOutputType(TransitRouterMulticastDomainPeerMemberMapOutput{})
}

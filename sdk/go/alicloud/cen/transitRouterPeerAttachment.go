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

// Provides a CEN transit router peer attachment resource that associate the transit router with the CEN instance. [What is CEN transit router peer attachment](https://www.alibabacloud.com/help/en/cen/developer-reference/api-cbn-2017-09-12-createtransitrouterpeerattachment)
//
// > **NOTE:** Available since v1.128.0.
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
//			region := "cn-hangzhou"
//			if param := cfg.Get("region"); param != "" {
//				region = param
//			}
//			peerRegion := "cn-beijing"
//			if param := cfg.Get("peerRegion"); param != "" {
//				peerRegion = param
//			}
//			_, err := alicloud.NewProvider(ctx, "hz", &alicloud.ProviderArgs{
//				Region: pulumi.String(region),
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
//			exampleInstance, err := cen.NewInstance(ctx, "exampleInstance", &cen.InstanceArgs{
//				CenInstanceName: pulumi.String(name),
//				ProtectionLevel: pulumi.String("REDUCED"),
//			}, pulumi.Provider(alicloud.Bj))
//			if err != nil {
//				return err
//			}
//			exampleBandwidthPackage, err := cen.NewBandwidthPackage(ctx, "exampleBandwidthPackage", &cen.BandwidthPackageArgs{
//				Bandwidth:               pulumi.Int(5),
//				CenBandwidthPackageName: pulumi.String("tf_example"),
//				GeographicRegionAId:     pulumi.String("China"),
//				GeographicRegionBId:     pulumi.String("China"),
//			}, pulumi.Provider(alicloud.Bj))
//			if err != nil {
//				return err
//			}
//			exampleBandwidthPackageAttachment, err := cen.NewBandwidthPackageAttachment(ctx, "exampleBandwidthPackageAttachment", &cen.BandwidthPackageAttachmentArgs{
//				InstanceId:         exampleInstance.ID(),
//				BandwidthPackageId: exampleBandwidthPackage.ID(),
//			}, pulumi.Provider(alicloud.Bj))
//			if err != nil {
//				return err
//			}
//			exampleTransitRouter, err := cen.NewTransitRouter(ctx, "exampleTransitRouter", &cen.TransitRouterArgs{
//				CenId: exampleBandwidthPackageAttachment.InstanceId,
//			}, pulumi.Provider(alicloud.Hz))
//			if err != nil {
//				return err
//			}
//			peer, err := cen.NewTransitRouter(ctx, "peer", &cen.TransitRouterArgs{
//				CenId: exampleTransitRouter.CenId,
//			}, pulumi.Provider(alicloud.Bj))
//			if err != nil {
//				return err
//			}
//			_, err = cen.NewTransitRouterPeerAttachment(ctx, "exampleTransitRouterPeerAttachment", &cen.TransitRouterPeerAttachmentArgs{
//				CenId:                              exampleInstance.ID(),
//				TransitRouterId:                    exampleTransitRouter.TransitRouterId,
//				PeerTransitRouterRegionId:          pulumi.String(peerRegion),
//				PeerTransitRouterId:                peer.TransitRouterId,
//				CenBandwidthPackageId:              exampleBandwidthPackageAttachment.BandwidthPackageId,
//				Bandwidth:                          pulumi.Int(5),
//				TransitRouterAttachmentDescription: pulumi.String(name),
//				TransitRouterAttachmentName:        pulumi.String(name),
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
// CEN instance can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:cen/transitRouterPeerAttachment:TransitRouterPeerAttachment example tr-********:tr-attach-*******
// ```
type TransitRouterPeerAttachment struct {
	pulumi.CustomResourceState

	// Auto publish route enabled. The system default value is `false`.
	AutoPublishRouteEnabled pulumi.BoolPtrOutput `pulumi:"autoPublishRouteEnabled"`
	// The bandwidth of the bandwidth package.
	Bandwidth pulumi.IntPtrOutput `pulumi:"bandwidth"`
	// The method that is used to allocate bandwidth to the cross-region connection. Valid values: `BandwidthPackage` and `DataTransfer`.
	BandwidthType pulumi.StringOutput `pulumi:"bandwidthType"`
	// The ID of the bandwidth package. If you do not enter the ID of the package, it means you are using the test. The system default test is 1bps, demonstrating that you test network connectivity
	CenBandwidthPackageId pulumi.StringPtrOutput `pulumi:"cenBandwidthPackageId"`
	// The ID of the CEN.
	CenId pulumi.StringOutput `pulumi:"cenId"`
	// Whether to perform pre-check for this request, including permission, instance status verification, etc.
	DryRun pulumi.BoolPtrOutput `pulumi:"dryRun"`
	// The ID of the peer transit router.
	PeerTransitRouterId pulumi.StringOutput `pulumi:"peerTransitRouterId"`
	// The region ID of peer transit router.
	PeerTransitRouterRegionId pulumi.StringOutput `pulumi:"peerTransitRouterRegionId"`
	// The resource type to attachment. Only support `VR` and default value is `VR`.
	ResourceType pulumi.StringPtrOutput `pulumi:"resourceType"`
	// Whether to association route table. System default is `false`.
	RouteTableAssociationEnabled pulumi.BoolPtrOutput `pulumi:"routeTableAssociationEnabled"`
	// Whether to propagation route table. System default is `false`.
	RouteTablePropagationEnabled pulumi.BoolPtrOutput `pulumi:"routeTablePropagationEnabled"`
	// The associating status of the network.
	Status pulumi.StringOutput `pulumi:"status"`
	// The description of transit router attachment. The description is 2~256 characters long and must start with a letter or Chinese, but cannot start with `http://` or `https://`.
	TransitRouterAttachmentDescription pulumi.StringPtrOutput `pulumi:"transitRouterAttachmentDescription"`
	// The ID of transit router attachment id.
	TransitRouterAttachmentId pulumi.StringOutput `pulumi:"transitRouterAttachmentId"`
	// The name of transit router attachment. The name is 2~128 characters in length, starts with uppercase and lowercase letters or Chinese, and can contain numbers, underscores (_) and dashes (-)
	TransitRouterAttachmentName pulumi.StringPtrOutput `pulumi:"transitRouterAttachmentName"`
	// The ID of the transit router to attach.
	TransitRouterId pulumi.StringPtrOutput `pulumi:"transitRouterId"`
}

// NewTransitRouterPeerAttachment registers a new resource with the given unique name, arguments, and options.
func NewTransitRouterPeerAttachment(ctx *pulumi.Context,
	name string, args *TransitRouterPeerAttachmentArgs, opts ...pulumi.ResourceOption) (*TransitRouterPeerAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CenId == nil {
		return nil, errors.New("invalid value for required argument 'CenId'")
	}
	if args.PeerTransitRouterId == nil {
		return nil, errors.New("invalid value for required argument 'PeerTransitRouterId'")
	}
	if args.PeerTransitRouterRegionId == nil {
		return nil, errors.New("invalid value for required argument 'PeerTransitRouterRegionId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TransitRouterPeerAttachment
	err := ctx.RegisterResource("alicloud:cen/transitRouterPeerAttachment:TransitRouterPeerAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTransitRouterPeerAttachment gets an existing TransitRouterPeerAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTransitRouterPeerAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TransitRouterPeerAttachmentState, opts ...pulumi.ResourceOption) (*TransitRouterPeerAttachment, error) {
	var resource TransitRouterPeerAttachment
	err := ctx.ReadResource("alicloud:cen/transitRouterPeerAttachment:TransitRouterPeerAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TransitRouterPeerAttachment resources.
type transitRouterPeerAttachmentState struct {
	// Auto publish route enabled. The system default value is `false`.
	AutoPublishRouteEnabled *bool `pulumi:"autoPublishRouteEnabled"`
	// The bandwidth of the bandwidth package.
	Bandwidth *int `pulumi:"bandwidth"`
	// The method that is used to allocate bandwidth to the cross-region connection. Valid values: `BandwidthPackage` and `DataTransfer`.
	BandwidthType *string `pulumi:"bandwidthType"`
	// The ID of the bandwidth package. If you do not enter the ID of the package, it means you are using the test. The system default test is 1bps, demonstrating that you test network connectivity
	CenBandwidthPackageId *string `pulumi:"cenBandwidthPackageId"`
	// The ID of the CEN.
	CenId *string `pulumi:"cenId"`
	// Whether to perform pre-check for this request, including permission, instance status verification, etc.
	DryRun *bool `pulumi:"dryRun"`
	// The ID of the peer transit router.
	PeerTransitRouterId *string `pulumi:"peerTransitRouterId"`
	// The region ID of peer transit router.
	PeerTransitRouterRegionId *string `pulumi:"peerTransitRouterRegionId"`
	// The resource type to attachment. Only support `VR` and default value is `VR`.
	ResourceType *string `pulumi:"resourceType"`
	// Whether to association route table. System default is `false`.
	RouteTableAssociationEnabled *bool `pulumi:"routeTableAssociationEnabled"`
	// Whether to propagation route table. System default is `false`.
	RouteTablePropagationEnabled *bool `pulumi:"routeTablePropagationEnabled"`
	// The associating status of the network.
	Status *string `pulumi:"status"`
	// The description of transit router attachment. The description is 2~256 characters long and must start with a letter or Chinese, but cannot start with `http://` or `https://`.
	TransitRouterAttachmentDescription *string `pulumi:"transitRouterAttachmentDescription"`
	// The ID of transit router attachment id.
	TransitRouterAttachmentId *string `pulumi:"transitRouterAttachmentId"`
	// The name of transit router attachment. The name is 2~128 characters in length, starts with uppercase and lowercase letters or Chinese, and can contain numbers, underscores (_) and dashes (-)
	TransitRouterAttachmentName *string `pulumi:"transitRouterAttachmentName"`
	// The ID of the transit router to attach.
	TransitRouterId *string `pulumi:"transitRouterId"`
}

type TransitRouterPeerAttachmentState struct {
	// Auto publish route enabled. The system default value is `false`.
	AutoPublishRouteEnabled pulumi.BoolPtrInput
	// The bandwidth of the bandwidth package.
	Bandwidth pulumi.IntPtrInput
	// The method that is used to allocate bandwidth to the cross-region connection. Valid values: `BandwidthPackage` and `DataTransfer`.
	BandwidthType pulumi.StringPtrInput
	// The ID of the bandwidth package. If you do not enter the ID of the package, it means you are using the test. The system default test is 1bps, demonstrating that you test network connectivity
	CenBandwidthPackageId pulumi.StringPtrInput
	// The ID of the CEN.
	CenId pulumi.StringPtrInput
	// Whether to perform pre-check for this request, including permission, instance status verification, etc.
	DryRun pulumi.BoolPtrInput
	// The ID of the peer transit router.
	PeerTransitRouterId pulumi.StringPtrInput
	// The region ID of peer transit router.
	PeerTransitRouterRegionId pulumi.StringPtrInput
	// The resource type to attachment. Only support `VR` and default value is `VR`.
	ResourceType pulumi.StringPtrInput
	// Whether to association route table. System default is `false`.
	RouteTableAssociationEnabled pulumi.BoolPtrInput
	// Whether to propagation route table. System default is `false`.
	RouteTablePropagationEnabled pulumi.BoolPtrInput
	// The associating status of the network.
	Status pulumi.StringPtrInput
	// The description of transit router attachment. The description is 2~256 characters long and must start with a letter or Chinese, but cannot start with `http://` or `https://`.
	TransitRouterAttachmentDescription pulumi.StringPtrInput
	// The ID of transit router attachment id.
	TransitRouterAttachmentId pulumi.StringPtrInput
	// The name of transit router attachment. The name is 2~128 characters in length, starts with uppercase and lowercase letters or Chinese, and can contain numbers, underscores (_) and dashes (-)
	TransitRouterAttachmentName pulumi.StringPtrInput
	// The ID of the transit router to attach.
	TransitRouterId pulumi.StringPtrInput
}

func (TransitRouterPeerAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*transitRouterPeerAttachmentState)(nil)).Elem()
}

type transitRouterPeerAttachmentArgs struct {
	// Auto publish route enabled. The system default value is `false`.
	AutoPublishRouteEnabled *bool `pulumi:"autoPublishRouteEnabled"`
	// The bandwidth of the bandwidth package.
	Bandwidth *int `pulumi:"bandwidth"`
	// The method that is used to allocate bandwidth to the cross-region connection. Valid values: `BandwidthPackage` and `DataTransfer`.
	BandwidthType *string `pulumi:"bandwidthType"`
	// The ID of the bandwidth package. If you do not enter the ID of the package, it means you are using the test. The system default test is 1bps, demonstrating that you test network connectivity
	CenBandwidthPackageId *string `pulumi:"cenBandwidthPackageId"`
	// The ID of the CEN.
	CenId string `pulumi:"cenId"`
	// Whether to perform pre-check for this request, including permission, instance status verification, etc.
	DryRun *bool `pulumi:"dryRun"`
	// The ID of the peer transit router.
	PeerTransitRouterId string `pulumi:"peerTransitRouterId"`
	// The region ID of peer transit router.
	PeerTransitRouterRegionId string `pulumi:"peerTransitRouterRegionId"`
	// The resource type to attachment. Only support `VR` and default value is `VR`.
	ResourceType *string `pulumi:"resourceType"`
	// Whether to association route table. System default is `false`.
	RouteTableAssociationEnabled *bool `pulumi:"routeTableAssociationEnabled"`
	// Whether to propagation route table. System default is `false`.
	RouteTablePropagationEnabled *bool `pulumi:"routeTablePropagationEnabled"`
	// The description of transit router attachment. The description is 2~256 characters long and must start with a letter or Chinese, but cannot start with `http://` or `https://`.
	TransitRouterAttachmentDescription *string `pulumi:"transitRouterAttachmentDescription"`
	// The name of transit router attachment. The name is 2~128 characters in length, starts with uppercase and lowercase letters or Chinese, and can contain numbers, underscores (_) and dashes (-)
	TransitRouterAttachmentName *string `pulumi:"transitRouterAttachmentName"`
	// The ID of the transit router to attach.
	TransitRouterId *string `pulumi:"transitRouterId"`
}

// The set of arguments for constructing a TransitRouterPeerAttachment resource.
type TransitRouterPeerAttachmentArgs struct {
	// Auto publish route enabled. The system default value is `false`.
	AutoPublishRouteEnabled pulumi.BoolPtrInput
	// The bandwidth of the bandwidth package.
	Bandwidth pulumi.IntPtrInput
	// The method that is used to allocate bandwidth to the cross-region connection. Valid values: `BandwidthPackage` and `DataTransfer`.
	BandwidthType pulumi.StringPtrInput
	// The ID of the bandwidth package. If you do not enter the ID of the package, it means you are using the test. The system default test is 1bps, demonstrating that you test network connectivity
	CenBandwidthPackageId pulumi.StringPtrInput
	// The ID of the CEN.
	CenId pulumi.StringInput
	// Whether to perform pre-check for this request, including permission, instance status verification, etc.
	DryRun pulumi.BoolPtrInput
	// The ID of the peer transit router.
	PeerTransitRouterId pulumi.StringInput
	// The region ID of peer transit router.
	PeerTransitRouterRegionId pulumi.StringInput
	// The resource type to attachment. Only support `VR` and default value is `VR`.
	ResourceType pulumi.StringPtrInput
	// Whether to association route table. System default is `false`.
	RouteTableAssociationEnabled pulumi.BoolPtrInput
	// Whether to propagation route table. System default is `false`.
	RouteTablePropagationEnabled pulumi.BoolPtrInput
	// The description of transit router attachment. The description is 2~256 characters long and must start with a letter or Chinese, but cannot start with `http://` or `https://`.
	TransitRouterAttachmentDescription pulumi.StringPtrInput
	// The name of transit router attachment. The name is 2~128 characters in length, starts with uppercase and lowercase letters or Chinese, and can contain numbers, underscores (_) and dashes (-)
	TransitRouterAttachmentName pulumi.StringPtrInput
	// The ID of the transit router to attach.
	TransitRouterId pulumi.StringPtrInput
}

func (TransitRouterPeerAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*transitRouterPeerAttachmentArgs)(nil)).Elem()
}

type TransitRouterPeerAttachmentInput interface {
	pulumi.Input

	ToTransitRouterPeerAttachmentOutput() TransitRouterPeerAttachmentOutput
	ToTransitRouterPeerAttachmentOutputWithContext(ctx context.Context) TransitRouterPeerAttachmentOutput
}

func (*TransitRouterPeerAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**TransitRouterPeerAttachment)(nil)).Elem()
}

func (i *TransitRouterPeerAttachment) ToTransitRouterPeerAttachmentOutput() TransitRouterPeerAttachmentOutput {
	return i.ToTransitRouterPeerAttachmentOutputWithContext(context.Background())
}

func (i *TransitRouterPeerAttachment) ToTransitRouterPeerAttachmentOutputWithContext(ctx context.Context) TransitRouterPeerAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransitRouterPeerAttachmentOutput)
}

// TransitRouterPeerAttachmentArrayInput is an input type that accepts TransitRouterPeerAttachmentArray and TransitRouterPeerAttachmentArrayOutput values.
// You can construct a concrete instance of `TransitRouterPeerAttachmentArrayInput` via:
//
//	TransitRouterPeerAttachmentArray{ TransitRouterPeerAttachmentArgs{...} }
type TransitRouterPeerAttachmentArrayInput interface {
	pulumi.Input

	ToTransitRouterPeerAttachmentArrayOutput() TransitRouterPeerAttachmentArrayOutput
	ToTransitRouterPeerAttachmentArrayOutputWithContext(context.Context) TransitRouterPeerAttachmentArrayOutput
}

type TransitRouterPeerAttachmentArray []TransitRouterPeerAttachmentInput

func (TransitRouterPeerAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TransitRouterPeerAttachment)(nil)).Elem()
}

func (i TransitRouterPeerAttachmentArray) ToTransitRouterPeerAttachmentArrayOutput() TransitRouterPeerAttachmentArrayOutput {
	return i.ToTransitRouterPeerAttachmentArrayOutputWithContext(context.Background())
}

func (i TransitRouterPeerAttachmentArray) ToTransitRouterPeerAttachmentArrayOutputWithContext(ctx context.Context) TransitRouterPeerAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransitRouterPeerAttachmentArrayOutput)
}

// TransitRouterPeerAttachmentMapInput is an input type that accepts TransitRouterPeerAttachmentMap and TransitRouterPeerAttachmentMapOutput values.
// You can construct a concrete instance of `TransitRouterPeerAttachmentMapInput` via:
//
//	TransitRouterPeerAttachmentMap{ "key": TransitRouterPeerAttachmentArgs{...} }
type TransitRouterPeerAttachmentMapInput interface {
	pulumi.Input

	ToTransitRouterPeerAttachmentMapOutput() TransitRouterPeerAttachmentMapOutput
	ToTransitRouterPeerAttachmentMapOutputWithContext(context.Context) TransitRouterPeerAttachmentMapOutput
}

type TransitRouterPeerAttachmentMap map[string]TransitRouterPeerAttachmentInput

func (TransitRouterPeerAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TransitRouterPeerAttachment)(nil)).Elem()
}

func (i TransitRouterPeerAttachmentMap) ToTransitRouterPeerAttachmentMapOutput() TransitRouterPeerAttachmentMapOutput {
	return i.ToTransitRouterPeerAttachmentMapOutputWithContext(context.Background())
}

func (i TransitRouterPeerAttachmentMap) ToTransitRouterPeerAttachmentMapOutputWithContext(ctx context.Context) TransitRouterPeerAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransitRouterPeerAttachmentMapOutput)
}

type TransitRouterPeerAttachmentOutput struct{ *pulumi.OutputState }

func (TransitRouterPeerAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TransitRouterPeerAttachment)(nil)).Elem()
}

func (o TransitRouterPeerAttachmentOutput) ToTransitRouterPeerAttachmentOutput() TransitRouterPeerAttachmentOutput {
	return o
}

func (o TransitRouterPeerAttachmentOutput) ToTransitRouterPeerAttachmentOutputWithContext(ctx context.Context) TransitRouterPeerAttachmentOutput {
	return o
}

// Auto publish route enabled. The system default value is `false`.
func (o TransitRouterPeerAttachmentOutput) AutoPublishRouteEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TransitRouterPeerAttachment) pulumi.BoolPtrOutput { return v.AutoPublishRouteEnabled }).(pulumi.BoolPtrOutput)
}

// The bandwidth of the bandwidth package.
func (o TransitRouterPeerAttachmentOutput) Bandwidth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TransitRouterPeerAttachment) pulumi.IntPtrOutput { return v.Bandwidth }).(pulumi.IntPtrOutput)
}

// The method that is used to allocate bandwidth to the cross-region connection. Valid values: `BandwidthPackage` and `DataTransfer`.
func (o TransitRouterPeerAttachmentOutput) BandwidthType() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitRouterPeerAttachment) pulumi.StringOutput { return v.BandwidthType }).(pulumi.StringOutput)
}

// The ID of the bandwidth package. If you do not enter the ID of the package, it means you are using the test. The system default test is 1bps, demonstrating that you test network connectivity
func (o TransitRouterPeerAttachmentOutput) CenBandwidthPackageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransitRouterPeerAttachment) pulumi.StringPtrOutput { return v.CenBandwidthPackageId }).(pulumi.StringPtrOutput)
}

// The ID of the CEN.
func (o TransitRouterPeerAttachmentOutput) CenId() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitRouterPeerAttachment) pulumi.StringOutput { return v.CenId }).(pulumi.StringOutput)
}

// Whether to perform pre-check for this request, including permission, instance status verification, etc.
func (o TransitRouterPeerAttachmentOutput) DryRun() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TransitRouterPeerAttachment) pulumi.BoolPtrOutput { return v.DryRun }).(pulumi.BoolPtrOutput)
}

// The ID of the peer transit router.
func (o TransitRouterPeerAttachmentOutput) PeerTransitRouterId() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitRouterPeerAttachment) pulumi.StringOutput { return v.PeerTransitRouterId }).(pulumi.StringOutput)
}

// The region ID of peer transit router.
func (o TransitRouterPeerAttachmentOutput) PeerTransitRouterRegionId() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitRouterPeerAttachment) pulumi.StringOutput { return v.PeerTransitRouterRegionId }).(pulumi.StringOutput)
}

// The resource type to attachment. Only support `VR` and default value is `VR`.
func (o TransitRouterPeerAttachmentOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransitRouterPeerAttachment) pulumi.StringPtrOutput { return v.ResourceType }).(pulumi.StringPtrOutput)
}

// Whether to association route table. System default is `false`.
func (o TransitRouterPeerAttachmentOutput) RouteTableAssociationEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TransitRouterPeerAttachment) pulumi.BoolPtrOutput { return v.RouteTableAssociationEnabled }).(pulumi.BoolPtrOutput)
}

// Whether to propagation route table. System default is `false`.
func (o TransitRouterPeerAttachmentOutput) RouteTablePropagationEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TransitRouterPeerAttachment) pulumi.BoolPtrOutput { return v.RouteTablePropagationEnabled }).(pulumi.BoolPtrOutput)
}

// The associating status of the network.
func (o TransitRouterPeerAttachmentOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitRouterPeerAttachment) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The description of transit router attachment. The description is 2~256 characters long and must start with a letter or Chinese, but cannot start with `http://` or `https://`.
func (o TransitRouterPeerAttachmentOutput) TransitRouterAttachmentDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransitRouterPeerAttachment) pulumi.StringPtrOutput {
		return v.TransitRouterAttachmentDescription
	}).(pulumi.StringPtrOutput)
}

// The ID of transit router attachment id.
func (o TransitRouterPeerAttachmentOutput) TransitRouterAttachmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitRouterPeerAttachment) pulumi.StringOutput { return v.TransitRouterAttachmentId }).(pulumi.StringOutput)
}

// The name of transit router attachment. The name is 2~128 characters in length, starts with uppercase and lowercase letters or Chinese, and can contain numbers, underscores (_) and dashes (-)
func (o TransitRouterPeerAttachmentOutput) TransitRouterAttachmentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransitRouterPeerAttachment) pulumi.StringPtrOutput { return v.TransitRouterAttachmentName }).(pulumi.StringPtrOutput)
}

// The ID of the transit router to attach.
func (o TransitRouterPeerAttachmentOutput) TransitRouterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransitRouterPeerAttachment) pulumi.StringPtrOutput { return v.TransitRouterId }).(pulumi.StringPtrOutput)
}

type TransitRouterPeerAttachmentArrayOutput struct{ *pulumi.OutputState }

func (TransitRouterPeerAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TransitRouterPeerAttachment)(nil)).Elem()
}

func (o TransitRouterPeerAttachmentArrayOutput) ToTransitRouterPeerAttachmentArrayOutput() TransitRouterPeerAttachmentArrayOutput {
	return o
}

func (o TransitRouterPeerAttachmentArrayOutput) ToTransitRouterPeerAttachmentArrayOutputWithContext(ctx context.Context) TransitRouterPeerAttachmentArrayOutput {
	return o
}

func (o TransitRouterPeerAttachmentArrayOutput) Index(i pulumi.IntInput) TransitRouterPeerAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TransitRouterPeerAttachment {
		return vs[0].([]*TransitRouterPeerAttachment)[vs[1].(int)]
	}).(TransitRouterPeerAttachmentOutput)
}

type TransitRouterPeerAttachmentMapOutput struct{ *pulumi.OutputState }

func (TransitRouterPeerAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TransitRouterPeerAttachment)(nil)).Elem()
}

func (o TransitRouterPeerAttachmentMapOutput) ToTransitRouterPeerAttachmentMapOutput() TransitRouterPeerAttachmentMapOutput {
	return o
}

func (o TransitRouterPeerAttachmentMapOutput) ToTransitRouterPeerAttachmentMapOutputWithContext(ctx context.Context) TransitRouterPeerAttachmentMapOutput {
	return o
}

func (o TransitRouterPeerAttachmentMapOutput) MapIndex(k pulumi.StringInput) TransitRouterPeerAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TransitRouterPeerAttachment {
		return vs[0].(map[string]*TransitRouterPeerAttachment)[vs[1].(string)]
	}).(TransitRouterPeerAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TransitRouterPeerAttachmentInput)(nil)).Elem(), &TransitRouterPeerAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransitRouterPeerAttachmentArrayInput)(nil)).Elem(), TransitRouterPeerAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransitRouterPeerAttachmentMapInput)(nil)).Elem(), TransitRouterPeerAttachmentMap{})
	pulumi.RegisterOutputType(TransitRouterPeerAttachmentOutput{})
	pulumi.RegisterOutputType(TransitRouterPeerAttachmentArrayOutput{})
	pulumi.RegisterOutputType(TransitRouterPeerAttachmentMapOutput{})
}

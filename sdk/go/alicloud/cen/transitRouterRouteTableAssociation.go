// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cen

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides a CEN transit router route table association resource.[What is Cen Transit Router Route Table Association](https://www.alibabacloud.com/help/en/cloud-enterprise-network/latest/api-doc-cbn-2017-09-12-api-doc-createtransitroutetableaggregation)
//
// > **NOTE:** Available since v1.126.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cen"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
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
//			_default, err := cen.GetTransitRouterAvailableResources(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			masterZone := _default.Resources[0].MasterZones[0]
//			slaveZone := _default.Resources[0].SlaveZones[1]
//			exampleNetwork, err := vpc.NewNetwork(ctx, "exampleNetwork", &vpc.NetworkArgs{
//				VpcName:   pulumi.String(name),
//				CidrBlock: pulumi.String("192.168.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleMaster, err := vpc.NewSwitch(ctx, "exampleMaster", &vpc.SwitchArgs{
//				VswitchName: pulumi.String(name),
//				CidrBlock:   pulumi.String("192.168.1.0/24"),
//				VpcId:       exampleNetwork.ID(),
//				ZoneId:      *pulumi.String(masterZone),
//			})
//			if err != nil {
//				return err
//			}
//			exampleSlave, err := vpc.NewSwitch(ctx, "exampleSlave", &vpc.SwitchArgs{
//				VswitchName: pulumi.String(name),
//				CidrBlock:   pulumi.String("192.168.2.0/24"),
//				VpcId:       exampleNetwork.ID(),
//				ZoneId:      *pulumi.String(slaveZone),
//			})
//			if err != nil {
//				return err
//			}
//			exampleInstance, err := cen.NewInstance(ctx, "exampleInstance", &cen.InstanceArgs{
//				CenInstanceName: pulumi.String(name),
//				ProtectionLevel: pulumi.String("REDUCED"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleTransitRouter, err := cen.NewTransitRouter(ctx, "exampleTransitRouter", &cen.TransitRouterArgs{
//				TransitRouterName: pulumi.String(name),
//				CenId:             exampleInstance.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			exampleTransitRouterVpcAttachment, err := cen.NewTransitRouterVpcAttachment(ctx, "exampleTransitRouterVpcAttachment", &cen.TransitRouterVpcAttachmentArgs{
//				CenId:           exampleInstance.ID(),
//				TransitRouterId: exampleTransitRouter.TransitRouterId,
//				VpcId:           exampleNetwork.ID(),
//				ZoneMappings: cen.TransitRouterVpcAttachmentZoneMappingArray{
//					&cen.TransitRouterVpcAttachmentZoneMappingArgs{
//						ZoneId:    *pulumi.String(masterZone),
//						VswitchId: exampleMaster.ID(),
//					},
//					&cen.TransitRouterVpcAttachmentZoneMappingArgs{
//						ZoneId:    *pulumi.String(slaveZone),
//						VswitchId: exampleSlave.ID(),
//					},
//				},
//				TransitRouterAttachmentName:        pulumi.String(name),
//				TransitRouterAttachmentDescription: pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			exampleTransitRouterRouteTable, err := cen.NewTransitRouterRouteTable(ctx, "exampleTransitRouterRouteTable", &cen.TransitRouterRouteTableArgs{
//				TransitRouterId: exampleTransitRouter.TransitRouterId,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cen.NewTransitRouterRouteTableAssociation(ctx, "exampleTransitRouterRouteTableAssociation", &cen.TransitRouterRouteTableAssociationArgs{
//				TransitRouterRouteTableId: exampleTransitRouterRouteTable.TransitRouterRouteTableId,
//				TransitRouterAttachmentId: exampleTransitRouterVpcAttachment.TransitRouterAttachmentId,
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
// CEN transit router route table association can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:cen/transitRouterRouteTableAssociation:TransitRouterRouteTableAssociation default tr-********:tr-attach-********
//
// ```
type TransitRouterRouteTableAssociation struct {
	pulumi.CustomResourceState

	// The dry run.
	//
	// > **NOTE:** The Zone of CEN has MasterZone and SlaveZone, first zoneId of zoneMapping need be MasterZone. We have a API to describeZones[API](https://help.aliyun.com/document_detail/261356.html)
	DryRun pulumi.BoolPtrOutput `pulumi:"dryRun"`
	// The associating status of the network.
	Status pulumi.StringOutput `pulumi:"status"`
	// The ID the transit router attachment.
	TransitRouterAttachmentId pulumi.StringOutput `pulumi:"transitRouterAttachmentId"`
	// The ID of the transit router route table.
	TransitRouterRouteTableId pulumi.StringOutput `pulumi:"transitRouterRouteTableId"`
}

// NewTransitRouterRouteTableAssociation registers a new resource with the given unique name, arguments, and options.
func NewTransitRouterRouteTableAssociation(ctx *pulumi.Context,
	name string, args *TransitRouterRouteTableAssociationArgs, opts ...pulumi.ResourceOption) (*TransitRouterRouteTableAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TransitRouterAttachmentId == nil {
		return nil, errors.New("invalid value for required argument 'TransitRouterAttachmentId'")
	}
	if args.TransitRouterRouteTableId == nil {
		return nil, errors.New("invalid value for required argument 'TransitRouterRouteTableId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TransitRouterRouteTableAssociation
	err := ctx.RegisterResource("alicloud:cen/transitRouterRouteTableAssociation:TransitRouterRouteTableAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTransitRouterRouteTableAssociation gets an existing TransitRouterRouteTableAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTransitRouterRouteTableAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TransitRouterRouteTableAssociationState, opts ...pulumi.ResourceOption) (*TransitRouterRouteTableAssociation, error) {
	var resource TransitRouterRouteTableAssociation
	err := ctx.ReadResource("alicloud:cen/transitRouterRouteTableAssociation:TransitRouterRouteTableAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TransitRouterRouteTableAssociation resources.
type transitRouterRouteTableAssociationState struct {
	// The dry run.
	//
	// > **NOTE:** The Zone of CEN has MasterZone and SlaveZone, first zoneId of zoneMapping need be MasterZone. We have a API to describeZones[API](https://help.aliyun.com/document_detail/261356.html)
	DryRun *bool `pulumi:"dryRun"`
	// The associating status of the network.
	Status *string `pulumi:"status"`
	// The ID the transit router attachment.
	TransitRouterAttachmentId *string `pulumi:"transitRouterAttachmentId"`
	// The ID of the transit router route table.
	TransitRouterRouteTableId *string `pulumi:"transitRouterRouteTableId"`
}

type TransitRouterRouteTableAssociationState struct {
	// The dry run.
	//
	// > **NOTE:** The Zone of CEN has MasterZone and SlaveZone, first zoneId of zoneMapping need be MasterZone. We have a API to describeZones[API](https://help.aliyun.com/document_detail/261356.html)
	DryRun pulumi.BoolPtrInput
	// The associating status of the network.
	Status pulumi.StringPtrInput
	// The ID the transit router attachment.
	TransitRouterAttachmentId pulumi.StringPtrInput
	// The ID of the transit router route table.
	TransitRouterRouteTableId pulumi.StringPtrInput
}

func (TransitRouterRouteTableAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*transitRouterRouteTableAssociationState)(nil)).Elem()
}

type transitRouterRouteTableAssociationArgs struct {
	// The dry run.
	//
	// > **NOTE:** The Zone of CEN has MasterZone and SlaveZone, first zoneId of zoneMapping need be MasterZone. We have a API to describeZones[API](https://help.aliyun.com/document_detail/261356.html)
	DryRun *bool `pulumi:"dryRun"`
	// The ID the transit router attachment.
	TransitRouterAttachmentId string `pulumi:"transitRouterAttachmentId"`
	// The ID of the transit router route table.
	TransitRouterRouteTableId string `pulumi:"transitRouterRouteTableId"`
}

// The set of arguments for constructing a TransitRouterRouteTableAssociation resource.
type TransitRouterRouteTableAssociationArgs struct {
	// The dry run.
	//
	// > **NOTE:** The Zone of CEN has MasterZone and SlaveZone, first zoneId of zoneMapping need be MasterZone. We have a API to describeZones[API](https://help.aliyun.com/document_detail/261356.html)
	DryRun pulumi.BoolPtrInput
	// The ID the transit router attachment.
	TransitRouterAttachmentId pulumi.StringInput
	// The ID of the transit router route table.
	TransitRouterRouteTableId pulumi.StringInput
}

func (TransitRouterRouteTableAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*transitRouterRouteTableAssociationArgs)(nil)).Elem()
}

type TransitRouterRouteTableAssociationInput interface {
	pulumi.Input

	ToTransitRouterRouteTableAssociationOutput() TransitRouterRouteTableAssociationOutput
	ToTransitRouterRouteTableAssociationOutputWithContext(ctx context.Context) TransitRouterRouteTableAssociationOutput
}

func (*TransitRouterRouteTableAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**TransitRouterRouteTableAssociation)(nil)).Elem()
}

func (i *TransitRouterRouteTableAssociation) ToTransitRouterRouteTableAssociationOutput() TransitRouterRouteTableAssociationOutput {
	return i.ToTransitRouterRouteTableAssociationOutputWithContext(context.Background())
}

func (i *TransitRouterRouteTableAssociation) ToTransitRouterRouteTableAssociationOutputWithContext(ctx context.Context) TransitRouterRouteTableAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransitRouterRouteTableAssociationOutput)
}

func (i *TransitRouterRouteTableAssociation) ToOutput(ctx context.Context) pulumix.Output[*TransitRouterRouteTableAssociation] {
	return pulumix.Output[*TransitRouterRouteTableAssociation]{
		OutputState: i.ToTransitRouterRouteTableAssociationOutputWithContext(ctx).OutputState,
	}
}

// TransitRouterRouteTableAssociationArrayInput is an input type that accepts TransitRouterRouteTableAssociationArray and TransitRouterRouteTableAssociationArrayOutput values.
// You can construct a concrete instance of `TransitRouterRouteTableAssociationArrayInput` via:
//
//	TransitRouterRouteTableAssociationArray{ TransitRouterRouteTableAssociationArgs{...} }
type TransitRouterRouteTableAssociationArrayInput interface {
	pulumi.Input

	ToTransitRouterRouteTableAssociationArrayOutput() TransitRouterRouteTableAssociationArrayOutput
	ToTransitRouterRouteTableAssociationArrayOutputWithContext(context.Context) TransitRouterRouteTableAssociationArrayOutput
}

type TransitRouterRouteTableAssociationArray []TransitRouterRouteTableAssociationInput

func (TransitRouterRouteTableAssociationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TransitRouterRouteTableAssociation)(nil)).Elem()
}

func (i TransitRouterRouteTableAssociationArray) ToTransitRouterRouteTableAssociationArrayOutput() TransitRouterRouteTableAssociationArrayOutput {
	return i.ToTransitRouterRouteTableAssociationArrayOutputWithContext(context.Background())
}

func (i TransitRouterRouteTableAssociationArray) ToTransitRouterRouteTableAssociationArrayOutputWithContext(ctx context.Context) TransitRouterRouteTableAssociationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransitRouterRouteTableAssociationArrayOutput)
}

func (i TransitRouterRouteTableAssociationArray) ToOutput(ctx context.Context) pulumix.Output[[]*TransitRouterRouteTableAssociation] {
	return pulumix.Output[[]*TransitRouterRouteTableAssociation]{
		OutputState: i.ToTransitRouterRouteTableAssociationArrayOutputWithContext(ctx).OutputState,
	}
}

// TransitRouterRouteTableAssociationMapInput is an input type that accepts TransitRouterRouteTableAssociationMap and TransitRouterRouteTableAssociationMapOutput values.
// You can construct a concrete instance of `TransitRouterRouteTableAssociationMapInput` via:
//
//	TransitRouterRouteTableAssociationMap{ "key": TransitRouterRouteTableAssociationArgs{...} }
type TransitRouterRouteTableAssociationMapInput interface {
	pulumi.Input

	ToTransitRouterRouteTableAssociationMapOutput() TransitRouterRouteTableAssociationMapOutput
	ToTransitRouterRouteTableAssociationMapOutputWithContext(context.Context) TransitRouterRouteTableAssociationMapOutput
}

type TransitRouterRouteTableAssociationMap map[string]TransitRouterRouteTableAssociationInput

func (TransitRouterRouteTableAssociationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TransitRouterRouteTableAssociation)(nil)).Elem()
}

func (i TransitRouterRouteTableAssociationMap) ToTransitRouterRouteTableAssociationMapOutput() TransitRouterRouteTableAssociationMapOutput {
	return i.ToTransitRouterRouteTableAssociationMapOutputWithContext(context.Background())
}

func (i TransitRouterRouteTableAssociationMap) ToTransitRouterRouteTableAssociationMapOutputWithContext(ctx context.Context) TransitRouterRouteTableAssociationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransitRouterRouteTableAssociationMapOutput)
}

func (i TransitRouterRouteTableAssociationMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*TransitRouterRouteTableAssociation] {
	return pulumix.Output[map[string]*TransitRouterRouteTableAssociation]{
		OutputState: i.ToTransitRouterRouteTableAssociationMapOutputWithContext(ctx).OutputState,
	}
}

type TransitRouterRouteTableAssociationOutput struct{ *pulumi.OutputState }

func (TransitRouterRouteTableAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TransitRouterRouteTableAssociation)(nil)).Elem()
}

func (o TransitRouterRouteTableAssociationOutput) ToTransitRouterRouteTableAssociationOutput() TransitRouterRouteTableAssociationOutput {
	return o
}

func (o TransitRouterRouteTableAssociationOutput) ToTransitRouterRouteTableAssociationOutputWithContext(ctx context.Context) TransitRouterRouteTableAssociationOutput {
	return o
}

func (o TransitRouterRouteTableAssociationOutput) ToOutput(ctx context.Context) pulumix.Output[*TransitRouterRouteTableAssociation] {
	return pulumix.Output[*TransitRouterRouteTableAssociation]{
		OutputState: o.OutputState,
	}
}

// The dry run.
//
// > **NOTE:** The Zone of CEN has MasterZone and SlaveZone, first zoneId of zoneMapping need be MasterZone. We have a API to describeZones[API](https://help.aliyun.com/document_detail/261356.html)
func (o TransitRouterRouteTableAssociationOutput) DryRun() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TransitRouterRouteTableAssociation) pulumi.BoolPtrOutput { return v.DryRun }).(pulumi.BoolPtrOutput)
}

// The associating status of the network.
func (o TransitRouterRouteTableAssociationOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitRouterRouteTableAssociation) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The ID the transit router attachment.
func (o TransitRouterRouteTableAssociationOutput) TransitRouterAttachmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitRouterRouteTableAssociation) pulumi.StringOutput { return v.TransitRouterAttachmentId }).(pulumi.StringOutput)
}

// The ID of the transit router route table.
func (o TransitRouterRouteTableAssociationOutput) TransitRouterRouteTableId() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitRouterRouteTableAssociation) pulumi.StringOutput { return v.TransitRouterRouteTableId }).(pulumi.StringOutput)
}

type TransitRouterRouteTableAssociationArrayOutput struct{ *pulumi.OutputState }

func (TransitRouterRouteTableAssociationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TransitRouterRouteTableAssociation)(nil)).Elem()
}

func (o TransitRouterRouteTableAssociationArrayOutput) ToTransitRouterRouteTableAssociationArrayOutput() TransitRouterRouteTableAssociationArrayOutput {
	return o
}

func (o TransitRouterRouteTableAssociationArrayOutput) ToTransitRouterRouteTableAssociationArrayOutputWithContext(ctx context.Context) TransitRouterRouteTableAssociationArrayOutput {
	return o
}

func (o TransitRouterRouteTableAssociationArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*TransitRouterRouteTableAssociation] {
	return pulumix.Output[[]*TransitRouterRouteTableAssociation]{
		OutputState: o.OutputState,
	}
}

func (o TransitRouterRouteTableAssociationArrayOutput) Index(i pulumi.IntInput) TransitRouterRouteTableAssociationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TransitRouterRouteTableAssociation {
		return vs[0].([]*TransitRouterRouteTableAssociation)[vs[1].(int)]
	}).(TransitRouterRouteTableAssociationOutput)
}

type TransitRouterRouteTableAssociationMapOutput struct{ *pulumi.OutputState }

func (TransitRouterRouteTableAssociationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TransitRouterRouteTableAssociation)(nil)).Elem()
}

func (o TransitRouterRouteTableAssociationMapOutput) ToTransitRouterRouteTableAssociationMapOutput() TransitRouterRouteTableAssociationMapOutput {
	return o
}

func (o TransitRouterRouteTableAssociationMapOutput) ToTransitRouterRouteTableAssociationMapOutputWithContext(ctx context.Context) TransitRouterRouteTableAssociationMapOutput {
	return o
}

func (o TransitRouterRouteTableAssociationMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*TransitRouterRouteTableAssociation] {
	return pulumix.Output[map[string]*TransitRouterRouteTableAssociation]{
		OutputState: o.OutputState,
	}
}

func (o TransitRouterRouteTableAssociationMapOutput) MapIndex(k pulumi.StringInput) TransitRouterRouteTableAssociationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TransitRouterRouteTableAssociation {
		return vs[0].(map[string]*TransitRouterRouteTableAssociation)[vs[1].(string)]
	}).(TransitRouterRouteTableAssociationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TransitRouterRouteTableAssociationInput)(nil)).Elem(), &TransitRouterRouteTableAssociation{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransitRouterRouteTableAssociationArrayInput)(nil)).Elem(), TransitRouterRouteTableAssociationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransitRouterRouteTableAssociationMapInput)(nil)).Elem(), TransitRouterRouteTableAssociationMap{})
	pulumi.RegisterOutputType(TransitRouterRouteTableAssociationOutput{})
	pulumi.RegisterOutputType(TransitRouterRouteTableAssociationArrayOutput{})
	pulumi.RegisterOutputType(TransitRouterRouteTableAssociationMapOutput{})
}

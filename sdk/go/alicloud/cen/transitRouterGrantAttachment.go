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

// Provides a Cloud Enterprise Network (CEN) Transit Router Grant Attachment resource.
//
// For information about Cloud Enterprise Network (CEN) Transit Router Grant Attachment and how to use it, see [What is Transit Router Grant Attachment](https://www.alibabacloud.com/help/en/cloud-enterprise-network/latest/grantinstancetotransitrouter).
//
// > **NOTE:** Available since v1.187.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cen"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_default, err := alicloud.GetAccount(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			example, err := vpc.NewNetwork(ctx, "example", &vpc.NetworkArgs{
//				VpcName:   pulumi.String("tf_example"),
//				CidrBlock: pulumi.String("172.17.3.0/24"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleInstance, err := cen.NewInstance(ctx, "example", &cen.InstanceArgs{
//				CenInstanceName: pulumi.String("tf_example"),
//				Description:     pulumi.String("an example for cen"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cen.NewTransitRouterGrantAttachment(ctx, "example", &cen.TransitRouterGrantAttachmentArgs{
//				CenId:        exampleInstance.ID(),
//				CenOwnerId:   pulumi.String(_default.Id),
//				InstanceId:   example.ID(),
//				InstanceType: pulumi.String("VPC"),
//				OrderType:    pulumi.String("PayByCenOwner"),
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
// Cloud Enterprise Network (CEN) Transit Router Grant Attachment can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:cen/transitRouterGrantAttachment:TransitRouterGrantAttachment example <instance_type>:<instance_id>:<cen_owner_id>:<cen_id>
// ```
type TransitRouterGrantAttachment struct {
	pulumi.CustomResourceState

	// The ID of the Cloud Enterprise Network (CEN) instance to which the transit router belongs.
	CenId pulumi.StringOutput `pulumi:"cenId"`
	// The ID of the Alibaba Cloud account to which the CEN instance belongs.
	CenOwnerId pulumi.StringOutput `pulumi:"cenOwnerId"`
	// The ID of the network instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The type of the network instance. Valid values: `VPC`, `ExpressConnect`, `VPN`.
	InstanceType pulumi.StringOutput `pulumi:"instanceType"`
	// The entity that pays the fees of the network instance. Valid values: `PayByResourceOwner`, `PayByCenOwner`.
	OrderType pulumi.StringOutput `pulumi:"orderType"`
}

// NewTransitRouterGrantAttachment registers a new resource with the given unique name, arguments, and options.
func NewTransitRouterGrantAttachment(ctx *pulumi.Context,
	name string, args *TransitRouterGrantAttachmentArgs, opts ...pulumi.ResourceOption) (*TransitRouterGrantAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CenId == nil {
		return nil, errors.New("invalid value for required argument 'CenId'")
	}
	if args.CenOwnerId == nil {
		return nil, errors.New("invalid value for required argument 'CenOwnerId'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.InstanceType == nil {
		return nil, errors.New("invalid value for required argument 'InstanceType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TransitRouterGrantAttachment
	err := ctx.RegisterResource("alicloud:cen/transitRouterGrantAttachment:TransitRouterGrantAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTransitRouterGrantAttachment gets an existing TransitRouterGrantAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTransitRouterGrantAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TransitRouterGrantAttachmentState, opts ...pulumi.ResourceOption) (*TransitRouterGrantAttachment, error) {
	var resource TransitRouterGrantAttachment
	err := ctx.ReadResource("alicloud:cen/transitRouterGrantAttachment:TransitRouterGrantAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TransitRouterGrantAttachment resources.
type transitRouterGrantAttachmentState struct {
	// The ID of the Cloud Enterprise Network (CEN) instance to which the transit router belongs.
	CenId *string `pulumi:"cenId"`
	// The ID of the Alibaba Cloud account to which the CEN instance belongs.
	CenOwnerId *string `pulumi:"cenOwnerId"`
	// The ID of the network instance.
	InstanceId *string `pulumi:"instanceId"`
	// The type of the network instance. Valid values: `VPC`, `ExpressConnect`, `VPN`.
	InstanceType *string `pulumi:"instanceType"`
	// The entity that pays the fees of the network instance. Valid values: `PayByResourceOwner`, `PayByCenOwner`.
	OrderType *string `pulumi:"orderType"`
}

type TransitRouterGrantAttachmentState struct {
	// The ID of the Cloud Enterprise Network (CEN) instance to which the transit router belongs.
	CenId pulumi.StringPtrInput
	// The ID of the Alibaba Cloud account to which the CEN instance belongs.
	CenOwnerId pulumi.StringPtrInput
	// The ID of the network instance.
	InstanceId pulumi.StringPtrInput
	// The type of the network instance. Valid values: `VPC`, `ExpressConnect`, `VPN`.
	InstanceType pulumi.StringPtrInput
	// The entity that pays the fees of the network instance. Valid values: `PayByResourceOwner`, `PayByCenOwner`.
	OrderType pulumi.StringPtrInput
}

func (TransitRouterGrantAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*transitRouterGrantAttachmentState)(nil)).Elem()
}

type transitRouterGrantAttachmentArgs struct {
	// The ID of the Cloud Enterprise Network (CEN) instance to which the transit router belongs.
	CenId string `pulumi:"cenId"`
	// The ID of the Alibaba Cloud account to which the CEN instance belongs.
	CenOwnerId string `pulumi:"cenOwnerId"`
	// The ID of the network instance.
	InstanceId string `pulumi:"instanceId"`
	// The type of the network instance. Valid values: `VPC`, `ExpressConnect`, `VPN`.
	InstanceType string `pulumi:"instanceType"`
	// The entity that pays the fees of the network instance. Valid values: `PayByResourceOwner`, `PayByCenOwner`.
	OrderType *string `pulumi:"orderType"`
}

// The set of arguments for constructing a TransitRouterGrantAttachment resource.
type TransitRouterGrantAttachmentArgs struct {
	// The ID of the Cloud Enterprise Network (CEN) instance to which the transit router belongs.
	CenId pulumi.StringInput
	// The ID of the Alibaba Cloud account to which the CEN instance belongs.
	CenOwnerId pulumi.StringInput
	// The ID of the network instance.
	InstanceId pulumi.StringInput
	// The type of the network instance. Valid values: `VPC`, `ExpressConnect`, `VPN`.
	InstanceType pulumi.StringInput
	// The entity that pays the fees of the network instance. Valid values: `PayByResourceOwner`, `PayByCenOwner`.
	OrderType pulumi.StringPtrInput
}

func (TransitRouterGrantAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*transitRouterGrantAttachmentArgs)(nil)).Elem()
}

type TransitRouterGrantAttachmentInput interface {
	pulumi.Input

	ToTransitRouterGrantAttachmentOutput() TransitRouterGrantAttachmentOutput
	ToTransitRouterGrantAttachmentOutputWithContext(ctx context.Context) TransitRouterGrantAttachmentOutput
}

func (*TransitRouterGrantAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**TransitRouterGrantAttachment)(nil)).Elem()
}

func (i *TransitRouterGrantAttachment) ToTransitRouterGrantAttachmentOutput() TransitRouterGrantAttachmentOutput {
	return i.ToTransitRouterGrantAttachmentOutputWithContext(context.Background())
}

func (i *TransitRouterGrantAttachment) ToTransitRouterGrantAttachmentOutputWithContext(ctx context.Context) TransitRouterGrantAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransitRouterGrantAttachmentOutput)
}

// TransitRouterGrantAttachmentArrayInput is an input type that accepts TransitRouterGrantAttachmentArray and TransitRouterGrantAttachmentArrayOutput values.
// You can construct a concrete instance of `TransitRouterGrantAttachmentArrayInput` via:
//
//	TransitRouterGrantAttachmentArray{ TransitRouterGrantAttachmentArgs{...} }
type TransitRouterGrantAttachmentArrayInput interface {
	pulumi.Input

	ToTransitRouterGrantAttachmentArrayOutput() TransitRouterGrantAttachmentArrayOutput
	ToTransitRouterGrantAttachmentArrayOutputWithContext(context.Context) TransitRouterGrantAttachmentArrayOutput
}

type TransitRouterGrantAttachmentArray []TransitRouterGrantAttachmentInput

func (TransitRouterGrantAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TransitRouterGrantAttachment)(nil)).Elem()
}

func (i TransitRouterGrantAttachmentArray) ToTransitRouterGrantAttachmentArrayOutput() TransitRouterGrantAttachmentArrayOutput {
	return i.ToTransitRouterGrantAttachmentArrayOutputWithContext(context.Background())
}

func (i TransitRouterGrantAttachmentArray) ToTransitRouterGrantAttachmentArrayOutputWithContext(ctx context.Context) TransitRouterGrantAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransitRouterGrantAttachmentArrayOutput)
}

// TransitRouterGrantAttachmentMapInput is an input type that accepts TransitRouterGrantAttachmentMap and TransitRouterGrantAttachmentMapOutput values.
// You can construct a concrete instance of `TransitRouterGrantAttachmentMapInput` via:
//
//	TransitRouterGrantAttachmentMap{ "key": TransitRouterGrantAttachmentArgs{...} }
type TransitRouterGrantAttachmentMapInput interface {
	pulumi.Input

	ToTransitRouterGrantAttachmentMapOutput() TransitRouterGrantAttachmentMapOutput
	ToTransitRouterGrantAttachmentMapOutputWithContext(context.Context) TransitRouterGrantAttachmentMapOutput
}

type TransitRouterGrantAttachmentMap map[string]TransitRouterGrantAttachmentInput

func (TransitRouterGrantAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TransitRouterGrantAttachment)(nil)).Elem()
}

func (i TransitRouterGrantAttachmentMap) ToTransitRouterGrantAttachmentMapOutput() TransitRouterGrantAttachmentMapOutput {
	return i.ToTransitRouterGrantAttachmentMapOutputWithContext(context.Background())
}

func (i TransitRouterGrantAttachmentMap) ToTransitRouterGrantAttachmentMapOutputWithContext(ctx context.Context) TransitRouterGrantAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransitRouterGrantAttachmentMapOutput)
}

type TransitRouterGrantAttachmentOutput struct{ *pulumi.OutputState }

func (TransitRouterGrantAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TransitRouterGrantAttachment)(nil)).Elem()
}

func (o TransitRouterGrantAttachmentOutput) ToTransitRouterGrantAttachmentOutput() TransitRouterGrantAttachmentOutput {
	return o
}

func (o TransitRouterGrantAttachmentOutput) ToTransitRouterGrantAttachmentOutputWithContext(ctx context.Context) TransitRouterGrantAttachmentOutput {
	return o
}

// The ID of the Cloud Enterprise Network (CEN) instance to which the transit router belongs.
func (o TransitRouterGrantAttachmentOutput) CenId() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitRouterGrantAttachment) pulumi.StringOutput { return v.CenId }).(pulumi.StringOutput)
}

// The ID of the Alibaba Cloud account to which the CEN instance belongs.
func (o TransitRouterGrantAttachmentOutput) CenOwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitRouterGrantAttachment) pulumi.StringOutput { return v.CenOwnerId }).(pulumi.StringOutput)
}

// The ID of the network instance.
func (o TransitRouterGrantAttachmentOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitRouterGrantAttachment) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The type of the network instance. Valid values: `VPC`, `ExpressConnect`, `VPN`.
func (o TransitRouterGrantAttachmentOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitRouterGrantAttachment) pulumi.StringOutput { return v.InstanceType }).(pulumi.StringOutput)
}

// The entity that pays the fees of the network instance. Valid values: `PayByResourceOwner`, `PayByCenOwner`.
func (o TransitRouterGrantAttachmentOutput) OrderType() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitRouterGrantAttachment) pulumi.StringOutput { return v.OrderType }).(pulumi.StringOutput)
}

type TransitRouterGrantAttachmentArrayOutput struct{ *pulumi.OutputState }

func (TransitRouterGrantAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TransitRouterGrantAttachment)(nil)).Elem()
}

func (o TransitRouterGrantAttachmentArrayOutput) ToTransitRouterGrantAttachmentArrayOutput() TransitRouterGrantAttachmentArrayOutput {
	return o
}

func (o TransitRouterGrantAttachmentArrayOutput) ToTransitRouterGrantAttachmentArrayOutputWithContext(ctx context.Context) TransitRouterGrantAttachmentArrayOutput {
	return o
}

func (o TransitRouterGrantAttachmentArrayOutput) Index(i pulumi.IntInput) TransitRouterGrantAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TransitRouterGrantAttachment {
		return vs[0].([]*TransitRouterGrantAttachment)[vs[1].(int)]
	}).(TransitRouterGrantAttachmentOutput)
}

type TransitRouterGrantAttachmentMapOutput struct{ *pulumi.OutputState }

func (TransitRouterGrantAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TransitRouterGrantAttachment)(nil)).Elem()
}

func (o TransitRouterGrantAttachmentMapOutput) ToTransitRouterGrantAttachmentMapOutput() TransitRouterGrantAttachmentMapOutput {
	return o
}

func (o TransitRouterGrantAttachmentMapOutput) ToTransitRouterGrantAttachmentMapOutputWithContext(ctx context.Context) TransitRouterGrantAttachmentMapOutput {
	return o
}

func (o TransitRouterGrantAttachmentMapOutput) MapIndex(k pulumi.StringInput) TransitRouterGrantAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TransitRouterGrantAttachment {
		return vs[0].(map[string]*TransitRouterGrantAttachment)[vs[1].(string)]
	}).(TransitRouterGrantAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TransitRouterGrantAttachmentInput)(nil)).Elem(), &TransitRouterGrantAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransitRouterGrantAttachmentArrayInput)(nil)).Elem(), TransitRouterGrantAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransitRouterGrantAttachmentMapInput)(nil)).Elem(), TransitRouterGrantAttachmentMap{})
	pulumi.RegisterOutputType(TransitRouterGrantAttachmentOutput{})
	pulumi.RegisterOutputType(TransitRouterGrantAttachmentArrayOutput{})
	pulumi.RegisterOutputType(TransitRouterGrantAttachmentMapOutput{})
}

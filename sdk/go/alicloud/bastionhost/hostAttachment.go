// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bastionhost

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Bastion Host Host Attachment resource to add host into one host group.
//
// > **NOTE:** Available since v1.135.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/bastionhost"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
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
//			defaultZones, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
//				AvailableResourceCreation: pulumi.StringRef("VSwitch"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultNetworks, err := vpc.GetNetworks(ctx, &vpc.GetNetworksArgs{
//				NameRegex: pulumi.StringRef("^default-NODELETING$"),
//				CidrBlock: pulumi.StringRef("10.4.0.0/16"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultSwitches, err := vpc.GetSwitches(ctx, &vpc.GetSwitchesArgs{
//				CidrBlock: pulumi.StringRef("10.4.0.0/24"),
//				VpcId:     pulumi.StringRef(defaultNetworks.Ids[0]),
//				ZoneId:    pulumi.StringRef(defaultZones.Zones[0].Id),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultSecurityGroup, err := ecs.NewSecurityGroup(ctx, "defaultSecurityGroup", &ecs.SecurityGroupArgs{
//				VpcId: *pulumi.String(defaultNetworks.Ids[0]),
//			})
//			if err != nil {
//				return err
//			}
//			defaultInstance, err := bastionhost.NewInstance(ctx, "defaultInstance", &bastionhost.InstanceArgs{
//				Description: pulumi.String(name),
//				LicenseCode: pulumi.String("bhah_ent_50_asset"),
//				PlanCode:    pulumi.String("cloudbastion"),
//				Storage:     pulumi.String("5"),
//				Bandwidth:   pulumi.String("5"),
//				Period:      pulumi.Int(1),
//				VswitchId:   *pulumi.String(defaultSwitches.Ids[0]),
//				SecurityGroupIds: pulumi.StringArray{
//					defaultSecurityGroup.ID(),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			defaultHostGroup, err := bastionhost.NewHostGroup(ctx, "defaultHostGroup", &bastionhost.HostGroupArgs{
//				HostGroupName: pulumi.String(name),
//				InstanceId:    defaultInstance.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			defaultHost, err := bastionhost.NewHost(ctx, "defaultHost", &bastionhost.HostArgs{
//				InstanceId:         defaultInstance.ID(),
//				HostName:           pulumi.String(name),
//				ActiveAddressType:  pulumi.String("Private"),
//				HostPrivateAddress: pulumi.String("172.16.0.10"),
//				OsType:             pulumi.String("Linux"),
//				Source:             pulumi.String("Local"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = bastionhost.NewHostAttachment(ctx, "defaultHostAttachment", &bastionhost.HostAttachmentArgs{
//				HostGroupId: defaultHostGroup.HostGroupId,
//				HostId:      defaultHost.HostId,
//				InstanceId:  defaultInstance.ID(),
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
// Bastion Host Host Attachment can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:bastionhost/hostAttachment:HostAttachment example <instance_id>:<host_group_id>:<host_id>
// ```
type HostAttachment struct {
	pulumi.CustomResourceState

	// Specifies the added to the host group ID.
	HostGroupId pulumi.StringOutput `pulumi:"hostGroupId"`
	// Specified to be part of a host group of host ID.
	HostId pulumi.StringOutput `pulumi:"hostId"`
	// The bastion host instance id.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
}

// NewHostAttachment registers a new resource with the given unique name, arguments, and options.
func NewHostAttachment(ctx *pulumi.Context,
	name string, args *HostAttachmentArgs, opts ...pulumi.ResourceOption) (*HostAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HostGroupId == nil {
		return nil, errors.New("invalid value for required argument 'HostGroupId'")
	}
	if args.HostId == nil {
		return nil, errors.New("invalid value for required argument 'HostId'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource HostAttachment
	err := ctx.RegisterResource("alicloud:bastionhost/hostAttachment:HostAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHostAttachment gets an existing HostAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHostAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HostAttachmentState, opts ...pulumi.ResourceOption) (*HostAttachment, error) {
	var resource HostAttachment
	err := ctx.ReadResource("alicloud:bastionhost/hostAttachment:HostAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HostAttachment resources.
type hostAttachmentState struct {
	// Specifies the added to the host group ID.
	HostGroupId *string `pulumi:"hostGroupId"`
	// Specified to be part of a host group of host ID.
	HostId *string `pulumi:"hostId"`
	// The bastion host instance id.
	InstanceId *string `pulumi:"instanceId"`
}

type HostAttachmentState struct {
	// Specifies the added to the host group ID.
	HostGroupId pulumi.StringPtrInput
	// Specified to be part of a host group of host ID.
	HostId pulumi.StringPtrInput
	// The bastion host instance id.
	InstanceId pulumi.StringPtrInput
}

func (HostAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*hostAttachmentState)(nil)).Elem()
}

type hostAttachmentArgs struct {
	// Specifies the added to the host group ID.
	HostGroupId string `pulumi:"hostGroupId"`
	// Specified to be part of a host group of host ID.
	HostId string `pulumi:"hostId"`
	// The bastion host instance id.
	InstanceId string `pulumi:"instanceId"`
}

// The set of arguments for constructing a HostAttachment resource.
type HostAttachmentArgs struct {
	// Specifies the added to the host group ID.
	HostGroupId pulumi.StringInput
	// Specified to be part of a host group of host ID.
	HostId pulumi.StringInput
	// The bastion host instance id.
	InstanceId pulumi.StringInput
}

func (HostAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hostAttachmentArgs)(nil)).Elem()
}

type HostAttachmentInput interface {
	pulumi.Input

	ToHostAttachmentOutput() HostAttachmentOutput
	ToHostAttachmentOutputWithContext(ctx context.Context) HostAttachmentOutput
}

func (*HostAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**HostAttachment)(nil)).Elem()
}

func (i *HostAttachment) ToHostAttachmentOutput() HostAttachmentOutput {
	return i.ToHostAttachmentOutputWithContext(context.Background())
}

func (i *HostAttachment) ToHostAttachmentOutputWithContext(ctx context.Context) HostAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostAttachmentOutput)
}

// HostAttachmentArrayInput is an input type that accepts HostAttachmentArray and HostAttachmentArrayOutput values.
// You can construct a concrete instance of `HostAttachmentArrayInput` via:
//
//	HostAttachmentArray{ HostAttachmentArgs{...} }
type HostAttachmentArrayInput interface {
	pulumi.Input

	ToHostAttachmentArrayOutput() HostAttachmentArrayOutput
	ToHostAttachmentArrayOutputWithContext(context.Context) HostAttachmentArrayOutput
}

type HostAttachmentArray []HostAttachmentInput

func (HostAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HostAttachment)(nil)).Elem()
}

func (i HostAttachmentArray) ToHostAttachmentArrayOutput() HostAttachmentArrayOutput {
	return i.ToHostAttachmentArrayOutputWithContext(context.Background())
}

func (i HostAttachmentArray) ToHostAttachmentArrayOutputWithContext(ctx context.Context) HostAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostAttachmentArrayOutput)
}

// HostAttachmentMapInput is an input type that accepts HostAttachmentMap and HostAttachmentMapOutput values.
// You can construct a concrete instance of `HostAttachmentMapInput` via:
//
//	HostAttachmentMap{ "key": HostAttachmentArgs{...} }
type HostAttachmentMapInput interface {
	pulumi.Input

	ToHostAttachmentMapOutput() HostAttachmentMapOutput
	ToHostAttachmentMapOutputWithContext(context.Context) HostAttachmentMapOutput
}

type HostAttachmentMap map[string]HostAttachmentInput

func (HostAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HostAttachment)(nil)).Elem()
}

func (i HostAttachmentMap) ToHostAttachmentMapOutput() HostAttachmentMapOutput {
	return i.ToHostAttachmentMapOutputWithContext(context.Background())
}

func (i HostAttachmentMap) ToHostAttachmentMapOutputWithContext(ctx context.Context) HostAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostAttachmentMapOutput)
}

type HostAttachmentOutput struct{ *pulumi.OutputState }

func (HostAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostAttachment)(nil)).Elem()
}

func (o HostAttachmentOutput) ToHostAttachmentOutput() HostAttachmentOutput {
	return o
}

func (o HostAttachmentOutput) ToHostAttachmentOutputWithContext(ctx context.Context) HostAttachmentOutput {
	return o
}

// Specifies the added to the host group ID.
func (o HostAttachmentOutput) HostGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *HostAttachment) pulumi.StringOutput { return v.HostGroupId }).(pulumi.StringOutput)
}

// Specified to be part of a host group of host ID.
func (o HostAttachmentOutput) HostId() pulumi.StringOutput {
	return o.ApplyT(func(v *HostAttachment) pulumi.StringOutput { return v.HostId }).(pulumi.StringOutput)
}

// The bastion host instance id.
func (o HostAttachmentOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *HostAttachment) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

type HostAttachmentArrayOutput struct{ *pulumi.OutputState }

func (HostAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HostAttachment)(nil)).Elem()
}

func (o HostAttachmentArrayOutput) ToHostAttachmentArrayOutput() HostAttachmentArrayOutput {
	return o
}

func (o HostAttachmentArrayOutput) ToHostAttachmentArrayOutputWithContext(ctx context.Context) HostAttachmentArrayOutput {
	return o
}

func (o HostAttachmentArrayOutput) Index(i pulumi.IntInput) HostAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HostAttachment {
		return vs[0].([]*HostAttachment)[vs[1].(int)]
	}).(HostAttachmentOutput)
}

type HostAttachmentMapOutput struct{ *pulumi.OutputState }

func (HostAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HostAttachment)(nil)).Elem()
}

func (o HostAttachmentMapOutput) ToHostAttachmentMapOutput() HostAttachmentMapOutput {
	return o
}

func (o HostAttachmentMapOutput) ToHostAttachmentMapOutputWithContext(ctx context.Context) HostAttachmentMapOutput {
	return o
}

func (o HostAttachmentMapOutput) MapIndex(k pulumi.StringInput) HostAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HostAttachment {
		return vs[0].(map[string]*HostAttachment)[vs[1].(string)]
	}).(HostAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HostAttachmentInput)(nil)).Elem(), &HostAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostAttachmentArrayInput)(nil)).Elem(), HostAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostAttachmentMapInput)(nil)).Elem(), HostAttachmentMap{})
	pulumi.RegisterOutputType(HostAttachmentOutput{})
	pulumi.RegisterOutputType(HostAttachmentArrayOutput{})
	pulumi.RegisterOutputType(HostAttachmentMapOutput{})
}

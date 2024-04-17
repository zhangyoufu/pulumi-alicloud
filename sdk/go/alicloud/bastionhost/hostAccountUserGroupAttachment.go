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

// Provides a Bastion Host Host Account Attachment resource to add list host accounts into one user group.
//
// > **NOTE:** Available since v1.135.0.
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
//			_default, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
//				AvailableResourceCreation: pulumi.StringRef("VSwitch"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultGetNetworks, err := vpc.GetNetworks(ctx, &vpc.GetNetworksArgs{
//				NameRegex: pulumi.StringRef("^default-NODELETING$"),
//				CidrBlock: pulumi.StringRef("10.4.0.0/16"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultGetSwitches, err := vpc.GetSwitches(ctx, &vpc.GetSwitchesArgs{
//				CidrBlock: pulumi.StringRef("10.4.0.0/24"),
//				VpcId:     pulumi.StringRef(defaultGetNetworks.Ids[0]),
//				ZoneId:    pulumi.StringRef(_default.Zones[0].Id),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultSecurityGroup, err := ecs.NewSecurityGroup(ctx, "default", &ecs.SecurityGroupArgs{
//				VpcId: pulumi.String(defaultGetNetworks.Ids[0]),
//			})
//			if err != nil {
//				return err
//			}
//			defaultInstance, err := bastionhost.NewInstance(ctx, "default", &bastionhost.InstanceArgs{
//				Description: pulumi.String(name),
//				LicenseCode: pulumi.String("bhah_ent_50_asset"),
//				PlanCode:    pulumi.String("cloudbastion"),
//				Storage:     pulumi.String("5"),
//				Bandwidth:   pulumi.String("5"),
//				Period:      pulumi.Int(1),
//				VswitchId:   pulumi.String(defaultGetSwitches.Ids[0]),
//				SecurityGroupIds: pulumi.StringArray{
//					defaultSecurityGroup.ID(),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			defaultHost, err := bastionhost.NewHost(ctx, "default", &bastionhost.HostArgs{
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
//			defaultHostAccount, err := bastionhost.NewHostAccount(ctx, "default", &bastionhost.HostAccountArgs{
//				HostAccountName: pulumi.String(name),
//				HostId:          defaultHost.HostId,
//				InstanceId:      defaultHost.InstanceId,
//				ProtocolName:    pulumi.String("SSH"),
//				Password:        pulumi.String("YourPassword12345"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultUserGroup, err := bastionhost.NewUserGroup(ctx, "default", &bastionhost.UserGroupArgs{
//				InstanceId:    defaultHost.InstanceId,
//				UserGroupName: pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = bastionhost.NewHostAccountUserGroupAttachment(ctx, "default", &bastionhost.HostAccountUserGroupAttachmentArgs{
//				InstanceId:  defaultHost.InstanceId,
//				UserGroupId: defaultUserGroup.UserGroupId,
//				HostId:      defaultHost.HostId,
//				HostAccountIds: pulumi.StringArray{
//					defaultHostAccount.HostAccountId,
//				},
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
// Bastion Host Host Account can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:bastionhost/hostAccountUserGroupAttachment:HostAccountUserGroupAttachment example <instance_id>:<user_group_id>:<host_id>
// ```
type HostAccountUserGroupAttachment struct {
	pulumi.CustomResourceState

	// A list IDs of the host account.
	HostAccountIds pulumi.StringArrayOutput `pulumi:"hostAccountIds"`
	// The ID of the host.
	HostId pulumi.StringOutput `pulumi:"hostId"`
	// The ID of the Bastionhost instance where you want to authorize the user group to manage the specified hosts and host accounts.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The ID of the user group that you want to authorize to manage the specified hosts and host accounts.
	UserGroupId pulumi.StringOutput `pulumi:"userGroupId"`
}

// NewHostAccountUserGroupAttachment registers a new resource with the given unique name, arguments, and options.
func NewHostAccountUserGroupAttachment(ctx *pulumi.Context,
	name string, args *HostAccountUserGroupAttachmentArgs, opts ...pulumi.ResourceOption) (*HostAccountUserGroupAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HostAccountIds == nil {
		return nil, errors.New("invalid value for required argument 'HostAccountIds'")
	}
	if args.HostId == nil {
		return nil, errors.New("invalid value for required argument 'HostId'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.UserGroupId == nil {
		return nil, errors.New("invalid value for required argument 'UserGroupId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource HostAccountUserGroupAttachment
	err := ctx.RegisterResource("alicloud:bastionhost/hostAccountUserGroupAttachment:HostAccountUserGroupAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHostAccountUserGroupAttachment gets an existing HostAccountUserGroupAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHostAccountUserGroupAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HostAccountUserGroupAttachmentState, opts ...pulumi.ResourceOption) (*HostAccountUserGroupAttachment, error) {
	var resource HostAccountUserGroupAttachment
	err := ctx.ReadResource("alicloud:bastionhost/hostAccountUserGroupAttachment:HostAccountUserGroupAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HostAccountUserGroupAttachment resources.
type hostAccountUserGroupAttachmentState struct {
	// A list IDs of the host account.
	HostAccountIds []string `pulumi:"hostAccountIds"`
	// The ID of the host.
	HostId *string `pulumi:"hostId"`
	// The ID of the Bastionhost instance where you want to authorize the user group to manage the specified hosts and host accounts.
	InstanceId *string `pulumi:"instanceId"`
	// The ID of the user group that you want to authorize to manage the specified hosts and host accounts.
	UserGroupId *string `pulumi:"userGroupId"`
}

type HostAccountUserGroupAttachmentState struct {
	// A list IDs of the host account.
	HostAccountIds pulumi.StringArrayInput
	// The ID of the host.
	HostId pulumi.StringPtrInput
	// The ID of the Bastionhost instance where you want to authorize the user group to manage the specified hosts and host accounts.
	InstanceId pulumi.StringPtrInput
	// The ID of the user group that you want to authorize to manage the specified hosts and host accounts.
	UserGroupId pulumi.StringPtrInput
}

func (HostAccountUserGroupAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*hostAccountUserGroupAttachmentState)(nil)).Elem()
}

type hostAccountUserGroupAttachmentArgs struct {
	// A list IDs of the host account.
	HostAccountIds []string `pulumi:"hostAccountIds"`
	// The ID of the host.
	HostId string `pulumi:"hostId"`
	// The ID of the Bastionhost instance where you want to authorize the user group to manage the specified hosts and host accounts.
	InstanceId string `pulumi:"instanceId"`
	// The ID of the user group that you want to authorize to manage the specified hosts and host accounts.
	UserGroupId string `pulumi:"userGroupId"`
}

// The set of arguments for constructing a HostAccountUserGroupAttachment resource.
type HostAccountUserGroupAttachmentArgs struct {
	// A list IDs of the host account.
	HostAccountIds pulumi.StringArrayInput
	// The ID of the host.
	HostId pulumi.StringInput
	// The ID of the Bastionhost instance where you want to authorize the user group to manage the specified hosts and host accounts.
	InstanceId pulumi.StringInput
	// The ID of the user group that you want to authorize to manage the specified hosts and host accounts.
	UserGroupId pulumi.StringInput
}

func (HostAccountUserGroupAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hostAccountUserGroupAttachmentArgs)(nil)).Elem()
}

type HostAccountUserGroupAttachmentInput interface {
	pulumi.Input

	ToHostAccountUserGroupAttachmentOutput() HostAccountUserGroupAttachmentOutput
	ToHostAccountUserGroupAttachmentOutputWithContext(ctx context.Context) HostAccountUserGroupAttachmentOutput
}

func (*HostAccountUserGroupAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**HostAccountUserGroupAttachment)(nil)).Elem()
}

func (i *HostAccountUserGroupAttachment) ToHostAccountUserGroupAttachmentOutput() HostAccountUserGroupAttachmentOutput {
	return i.ToHostAccountUserGroupAttachmentOutputWithContext(context.Background())
}

func (i *HostAccountUserGroupAttachment) ToHostAccountUserGroupAttachmentOutputWithContext(ctx context.Context) HostAccountUserGroupAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostAccountUserGroupAttachmentOutput)
}

// HostAccountUserGroupAttachmentArrayInput is an input type that accepts HostAccountUserGroupAttachmentArray and HostAccountUserGroupAttachmentArrayOutput values.
// You can construct a concrete instance of `HostAccountUserGroupAttachmentArrayInput` via:
//
//	HostAccountUserGroupAttachmentArray{ HostAccountUserGroupAttachmentArgs{...} }
type HostAccountUserGroupAttachmentArrayInput interface {
	pulumi.Input

	ToHostAccountUserGroupAttachmentArrayOutput() HostAccountUserGroupAttachmentArrayOutput
	ToHostAccountUserGroupAttachmentArrayOutputWithContext(context.Context) HostAccountUserGroupAttachmentArrayOutput
}

type HostAccountUserGroupAttachmentArray []HostAccountUserGroupAttachmentInput

func (HostAccountUserGroupAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HostAccountUserGroupAttachment)(nil)).Elem()
}

func (i HostAccountUserGroupAttachmentArray) ToHostAccountUserGroupAttachmentArrayOutput() HostAccountUserGroupAttachmentArrayOutput {
	return i.ToHostAccountUserGroupAttachmentArrayOutputWithContext(context.Background())
}

func (i HostAccountUserGroupAttachmentArray) ToHostAccountUserGroupAttachmentArrayOutputWithContext(ctx context.Context) HostAccountUserGroupAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostAccountUserGroupAttachmentArrayOutput)
}

// HostAccountUserGroupAttachmentMapInput is an input type that accepts HostAccountUserGroupAttachmentMap and HostAccountUserGroupAttachmentMapOutput values.
// You can construct a concrete instance of `HostAccountUserGroupAttachmentMapInput` via:
//
//	HostAccountUserGroupAttachmentMap{ "key": HostAccountUserGroupAttachmentArgs{...} }
type HostAccountUserGroupAttachmentMapInput interface {
	pulumi.Input

	ToHostAccountUserGroupAttachmentMapOutput() HostAccountUserGroupAttachmentMapOutput
	ToHostAccountUserGroupAttachmentMapOutputWithContext(context.Context) HostAccountUserGroupAttachmentMapOutput
}

type HostAccountUserGroupAttachmentMap map[string]HostAccountUserGroupAttachmentInput

func (HostAccountUserGroupAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HostAccountUserGroupAttachment)(nil)).Elem()
}

func (i HostAccountUserGroupAttachmentMap) ToHostAccountUserGroupAttachmentMapOutput() HostAccountUserGroupAttachmentMapOutput {
	return i.ToHostAccountUserGroupAttachmentMapOutputWithContext(context.Background())
}

func (i HostAccountUserGroupAttachmentMap) ToHostAccountUserGroupAttachmentMapOutputWithContext(ctx context.Context) HostAccountUserGroupAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostAccountUserGroupAttachmentMapOutput)
}

type HostAccountUserGroupAttachmentOutput struct{ *pulumi.OutputState }

func (HostAccountUserGroupAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostAccountUserGroupAttachment)(nil)).Elem()
}

func (o HostAccountUserGroupAttachmentOutput) ToHostAccountUserGroupAttachmentOutput() HostAccountUserGroupAttachmentOutput {
	return o
}

func (o HostAccountUserGroupAttachmentOutput) ToHostAccountUserGroupAttachmentOutputWithContext(ctx context.Context) HostAccountUserGroupAttachmentOutput {
	return o
}

// A list IDs of the host account.
func (o HostAccountUserGroupAttachmentOutput) HostAccountIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *HostAccountUserGroupAttachment) pulumi.StringArrayOutput { return v.HostAccountIds }).(pulumi.StringArrayOutput)
}

// The ID of the host.
func (o HostAccountUserGroupAttachmentOutput) HostId() pulumi.StringOutput {
	return o.ApplyT(func(v *HostAccountUserGroupAttachment) pulumi.StringOutput { return v.HostId }).(pulumi.StringOutput)
}

// The ID of the Bastionhost instance where you want to authorize the user group to manage the specified hosts and host accounts.
func (o HostAccountUserGroupAttachmentOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *HostAccountUserGroupAttachment) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The ID of the user group that you want to authorize to manage the specified hosts and host accounts.
func (o HostAccountUserGroupAttachmentOutput) UserGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *HostAccountUserGroupAttachment) pulumi.StringOutput { return v.UserGroupId }).(pulumi.StringOutput)
}

type HostAccountUserGroupAttachmentArrayOutput struct{ *pulumi.OutputState }

func (HostAccountUserGroupAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HostAccountUserGroupAttachment)(nil)).Elem()
}

func (o HostAccountUserGroupAttachmentArrayOutput) ToHostAccountUserGroupAttachmentArrayOutput() HostAccountUserGroupAttachmentArrayOutput {
	return o
}

func (o HostAccountUserGroupAttachmentArrayOutput) ToHostAccountUserGroupAttachmentArrayOutputWithContext(ctx context.Context) HostAccountUserGroupAttachmentArrayOutput {
	return o
}

func (o HostAccountUserGroupAttachmentArrayOutput) Index(i pulumi.IntInput) HostAccountUserGroupAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HostAccountUserGroupAttachment {
		return vs[0].([]*HostAccountUserGroupAttachment)[vs[1].(int)]
	}).(HostAccountUserGroupAttachmentOutput)
}

type HostAccountUserGroupAttachmentMapOutput struct{ *pulumi.OutputState }

func (HostAccountUserGroupAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HostAccountUserGroupAttachment)(nil)).Elem()
}

func (o HostAccountUserGroupAttachmentMapOutput) ToHostAccountUserGroupAttachmentMapOutput() HostAccountUserGroupAttachmentMapOutput {
	return o
}

func (o HostAccountUserGroupAttachmentMapOutput) ToHostAccountUserGroupAttachmentMapOutputWithContext(ctx context.Context) HostAccountUserGroupAttachmentMapOutput {
	return o
}

func (o HostAccountUserGroupAttachmentMapOutput) MapIndex(k pulumi.StringInput) HostAccountUserGroupAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HostAccountUserGroupAttachment {
		return vs[0].(map[string]*HostAccountUserGroupAttachment)[vs[1].(string)]
	}).(HostAccountUserGroupAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HostAccountUserGroupAttachmentInput)(nil)).Elem(), &HostAccountUserGroupAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostAccountUserGroupAttachmentArrayInput)(nil)).Elem(), HostAccountUserGroupAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostAccountUserGroupAttachmentMapInput)(nil)).Elem(), HostAccountUserGroupAttachmentMap{})
	pulumi.RegisterOutputType(HostAccountUserGroupAttachmentOutput{})
	pulumi.RegisterOutputType(HostAccountUserGroupAttachmentArrayOutput{})
	pulumi.RegisterOutputType(HostAccountUserGroupAttachmentMapOutput{})
}

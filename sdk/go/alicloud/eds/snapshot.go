// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eds

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a ECD Snapshot resource.
//
// For information about ECD Snapshot and how to use it, see [What is Snapshot](https://www.alibabacloud.com/help/en/elastic-desktop-service/latest/createsnapshot).
//
// > **NOTE:** Available in v1.169.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/eds"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "example_value"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			defaultSimpleOfficeSite, err := eds.NewSimpleOfficeSite(ctx, "defaultSimpleOfficeSite", &eds.SimpleOfficeSiteArgs{
//				CidrBlock:            pulumi.String("172.16.0.0/12"),
//				DesktopAccessType:    pulumi.String("Internet"),
//				OfficeSiteName:       pulumi.String(name),
//				EnableInternetAccess: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			defaultBundles, err := eds.GetBundles(ctx, &eds.GetBundlesArgs{
//				BundleType: pulumi.StringRef("SYSTEM"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultEcdPolicyGroup, err := eds.NewEcdPolicyGroup(ctx, "defaultEcdPolicyGroup", &eds.EcdPolicyGroupArgs{
//				PolicyGroupName: pulumi.String(name),
//				Clipboard:       pulumi.String("readwrite"),
//				LocalDrive:      pulumi.String("read"),
//				AuthorizeAccessPolicyRules: eds.EcdPolicyGroupAuthorizeAccessPolicyRuleArray{
//					&eds.EcdPolicyGroupAuthorizeAccessPolicyRuleArgs{
//						Description: pulumi.String("example_value"),
//						CidrIp:      pulumi.String("1.2.3.4/24"),
//					},
//				},
//				AuthorizeSecurityPolicyRules: eds.EcdPolicyGroupAuthorizeSecurityPolicyRuleArray{
//					&eds.EcdPolicyGroupAuthorizeSecurityPolicyRuleArgs{
//						Type:        pulumi.String("inflow"),
//						Policy:      pulumi.String("accept"),
//						Description: pulumi.String("example_value"),
//						PortRange:   pulumi.String("80/80"),
//						IpProtocol:  pulumi.String("TCP"),
//						Priority:    pulumi.String("1"),
//						CidrIp:      pulumi.String("0.0.0.0/0"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			defaultDesktop, err := eds.NewDesktop(ctx, "defaultDesktop", &eds.DesktopArgs{
//				OfficeSiteId:  defaultSimpleOfficeSite.ID(),
//				PolicyGroupId: defaultEcdPolicyGroup.ID(),
//				BundleId:      pulumi.String(defaultBundles.Bundles[0].Id),
//				DesktopName:   pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = eds.NewSnapshot(ctx, "defaultSnapshot", &eds.SnapshotArgs{
//				Description:    pulumi.String(name),
//				DesktopId:      defaultDesktop.ID(),
//				SnapshotName:   pulumi.String(name),
//				SourceDiskType: pulumi.String("SYSTEM"),
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
// ECD Snapshot can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:eds/snapshot:Snapshot example <id>
//
// ```
type Snapshot struct {
	pulumi.CustomResourceState

	// The description of the Snapshot.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The ID of the Desktop.
	DesktopId pulumi.StringOutput `pulumi:"desktopId"`
	// The name of the Snapshot.
	SnapshotName pulumi.StringOutput `pulumi:"snapshotName"`
	// The type of the disk for which to create a snapshot. Valid values: `SYSTEM`, `DATA`.
	SourceDiskType pulumi.StringOutput `pulumi:"sourceDiskType"`
	// The status of the snapshot.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewSnapshot registers a new resource with the given unique name, arguments, and options.
func NewSnapshot(ctx *pulumi.Context,
	name string, args *SnapshotArgs, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DesktopId == nil {
		return nil, errors.New("invalid value for required argument 'DesktopId'")
	}
	if args.SnapshotName == nil {
		return nil, errors.New("invalid value for required argument 'SnapshotName'")
	}
	if args.SourceDiskType == nil {
		return nil, errors.New("invalid value for required argument 'SourceDiskType'")
	}
	var resource Snapshot
	err := ctx.RegisterResource("alicloud:eds/snapshot:Snapshot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshot gets an existing Snapshot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotState, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	var resource Snapshot
	err := ctx.ReadResource("alicloud:eds/snapshot:Snapshot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Snapshot resources.
type snapshotState struct {
	// The description of the Snapshot.
	Description *string `pulumi:"description"`
	// The ID of the Desktop.
	DesktopId *string `pulumi:"desktopId"`
	// The name of the Snapshot.
	SnapshotName *string `pulumi:"snapshotName"`
	// The type of the disk for which to create a snapshot. Valid values: `SYSTEM`, `DATA`.
	SourceDiskType *string `pulumi:"sourceDiskType"`
	// The status of the snapshot.
	Status *string `pulumi:"status"`
}

type SnapshotState struct {
	// The description of the Snapshot.
	Description pulumi.StringPtrInput
	// The ID of the Desktop.
	DesktopId pulumi.StringPtrInput
	// The name of the Snapshot.
	SnapshotName pulumi.StringPtrInput
	// The type of the disk for which to create a snapshot. Valid values: `SYSTEM`, `DATA`.
	SourceDiskType pulumi.StringPtrInput
	// The status of the snapshot.
	Status pulumi.StringPtrInput
}

func (SnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotState)(nil)).Elem()
}

type snapshotArgs struct {
	// The description of the Snapshot.
	Description *string `pulumi:"description"`
	// The ID of the Desktop.
	DesktopId string `pulumi:"desktopId"`
	// The name of the Snapshot.
	SnapshotName string `pulumi:"snapshotName"`
	// The type of the disk for which to create a snapshot. Valid values: `SYSTEM`, `DATA`.
	SourceDiskType string `pulumi:"sourceDiskType"`
}

// The set of arguments for constructing a Snapshot resource.
type SnapshotArgs struct {
	// The description of the Snapshot.
	Description pulumi.StringPtrInput
	// The ID of the Desktop.
	DesktopId pulumi.StringInput
	// The name of the Snapshot.
	SnapshotName pulumi.StringInput
	// The type of the disk for which to create a snapshot. Valid values: `SYSTEM`, `DATA`.
	SourceDiskType pulumi.StringInput
}

func (SnapshotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotArgs)(nil)).Elem()
}

type SnapshotInput interface {
	pulumi.Input

	ToSnapshotOutput() SnapshotOutput
	ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput
}

func (*Snapshot) ElementType() reflect.Type {
	return reflect.TypeOf((**Snapshot)(nil)).Elem()
}

func (i *Snapshot) ToSnapshotOutput() SnapshotOutput {
	return i.ToSnapshotOutputWithContext(context.Background())
}

func (i *Snapshot) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotOutput)
}

// SnapshotArrayInput is an input type that accepts SnapshotArray and SnapshotArrayOutput values.
// You can construct a concrete instance of `SnapshotArrayInput` via:
//
//	SnapshotArray{ SnapshotArgs{...} }
type SnapshotArrayInput interface {
	pulumi.Input

	ToSnapshotArrayOutput() SnapshotArrayOutput
	ToSnapshotArrayOutputWithContext(context.Context) SnapshotArrayOutput
}

type SnapshotArray []SnapshotInput

func (SnapshotArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Snapshot)(nil)).Elem()
}

func (i SnapshotArray) ToSnapshotArrayOutput() SnapshotArrayOutput {
	return i.ToSnapshotArrayOutputWithContext(context.Background())
}

func (i SnapshotArray) ToSnapshotArrayOutputWithContext(ctx context.Context) SnapshotArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotArrayOutput)
}

// SnapshotMapInput is an input type that accepts SnapshotMap and SnapshotMapOutput values.
// You can construct a concrete instance of `SnapshotMapInput` via:
//
//	SnapshotMap{ "key": SnapshotArgs{...} }
type SnapshotMapInput interface {
	pulumi.Input

	ToSnapshotMapOutput() SnapshotMapOutput
	ToSnapshotMapOutputWithContext(context.Context) SnapshotMapOutput
}

type SnapshotMap map[string]SnapshotInput

func (SnapshotMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Snapshot)(nil)).Elem()
}

func (i SnapshotMap) ToSnapshotMapOutput() SnapshotMapOutput {
	return i.ToSnapshotMapOutputWithContext(context.Background())
}

func (i SnapshotMap) ToSnapshotMapOutputWithContext(ctx context.Context) SnapshotMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotMapOutput)
}

type SnapshotOutput struct{ *pulumi.OutputState }

func (SnapshotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Snapshot)(nil)).Elem()
}

func (o SnapshotOutput) ToSnapshotOutput() SnapshotOutput {
	return o
}

func (o SnapshotOutput) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return o
}

// The description of the Snapshot.
func (o SnapshotOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The ID of the Desktop.
func (o SnapshotOutput) DesktopId() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.DesktopId }).(pulumi.StringOutput)
}

// The name of the Snapshot.
func (o SnapshotOutput) SnapshotName() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.SnapshotName }).(pulumi.StringOutput)
}

// The type of the disk for which to create a snapshot. Valid values: `SYSTEM`, `DATA`.
func (o SnapshotOutput) SourceDiskType() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.SourceDiskType }).(pulumi.StringOutput)
}

// The status of the snapshot.
func (o SnapshotOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type SnapshotArrayOutput struct{ *pulumi.OutputState }

func (SnapshotArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Snapshot)(nil)).Elem()
}

func (o SnapshotArrayOutput) ToSnapshotArrayOutput() SnapshotArrayOutput {
	return o
}

func (o SnapshotArrayOutput) ToSnapshotArrayOutputWithContext(ctx context.Context) SnapshotArrayOutput {
	return o
}

func (o SnapshotArrayOutput) Index(i pulumi.IntInput) SnapshotOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Snapshot {
		return vs[0].([]*Snapshot)[vs[1].(int)]
	}).(SnapshotOutput)
}

type SnapshotMapOutput struct{ *pulumi.OutputState }

func (SnapshotMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Snapshot)(nil)).Elem()
}

func (o SnapshotMapOutput) ToSnapshotMapOutput() SnapshotMapOutput {
	return o
}

func (o SnapshotMapOutput) ToSnapshotMapOutputWithContext(ctx context.Context) SnapshotMapOutput {
	return o
}

func (o SnapshotMapOutput) MapIndex(k pulumi.StringInput) SnapshotOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Snapshot {
		return vs[0].(map[string]*Snapshot)[vs[1].(string)]
	}).(SnapshotOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotInput)(nil)).Elem(), &Snapshot{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotArrayInput)(nil)).Elem(), SnapshotArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotMapInput)(nil)).Elem(), SnapshotMap{})
	pulumi.RegisterOutputType(SnapshotOutput{})
	pulumi.RegisterOutputType(SnapshotArrayOutput{})
	pulumi.RegisterOutputType(SnapshotMapOutput{})
}

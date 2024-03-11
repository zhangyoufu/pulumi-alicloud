// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hbr

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Hybrid Backup Recovery (HBR) Hana Backup Client resource.
//
// For information about Hybrid Backup Recovery (HBR) Hana Backup Client and how to use it, see [What is Hana Backup Client](https://www.alibabacloud.com/help/en/hybrid-backup-recovery/latest/api-hbr-2017-09-08-createclients).
//
// > **NOTE:** Available in v1.198.0+.
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
//	"fmt"
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/hbr"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/resourcemanager"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleZones, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
//				AvailableResourceCreation: pulumi.StringRef("Instance"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleInstanceTypes, err := ecs.GetInstanceTypes(ctx, &ecs.GetInstanceTypesArgs{
//				AvailabilityZone: pulumi.StringRef(exampleZones.Zones[0].Id),
//				CpuCoreCount:     pulumi.IntRef(1),
//				MemorySize:       pulumi.Float64Ref(2),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleImages, err := ecs.GetImages(ctx, &ecs.GetImagesArgs{
//				NameRegex: pulumi.StringRef("^ubuntu_[0-9]+_[0-9]+_x64*"),
//				Owners:    pulumi.StringRef("system"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleNetwork, err := vpc.NewNetwork(ctx, "exampleNetwork", &vpc.NetworkArgs{
//				VpcName:   pulumi.String("terraform-example"),
//				CidrBlock: pulumi.String("172.17.3.0/24"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleSwitch, err := vpc.NewSwitch(ctx, "exampleSwitch", &vpc.SwitchArgs{
//				VswitchName: pulumi.String("terraform-example"),
//				CidrBlock:   pulumi.String("172.17.3.0/24"),
//				VpcId:       exampleNetwork.ID(),
//				ZoneId:      *pulumi.String(exampleZones.Zones[0].Id),
//			})
//			if err != nil {
//				return err
//			}
//			exampleSecurityGroup, err := ecs.NewSecurityGroup(ctx, "exampleSecurityGroup", &ecs.SecurityGroupArgs{
//				VpcId: exampleNetwork.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			exampleInstance, err := ecs.NewInstance(ctx, "exampleInstance", &ecs.InstanceArgs{
//				ImageId:          *pulumi.String(exampleImages.Images[0].Id),
//				InstanceType:     *pulumi.String(exampleInstanceTypes.InstanceTypes[0].Id),
//				AvailabilityZone: *pulumi.String(exampleZones.Zones[0].Id),
//				SecurityGroups: pulumi.StringArray{
//					exampleSecurityGroup.ID(),
//				},
//				InstanceName:       pulumi.String("terraform-example"),
//				InternetChargeType: pulumi.String("PayByBandwidth"),
//				VswitchId:          exampleSwitch.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			exampleResourceGroups, err := resourcemanager.GetResourceGroups(ctx, &resourcemanager.GetResourceGroupsArgs{
//				Status: pulumi.StringRef("OK"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleVault, err := hbr.NewVault(ctx, "exampleVault", &hbr.VaultArgs{
//				VaultName: pulumi.String("terraform-example"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleHanaInstance, err := hbr.NewHanaInstance(ctx, "exampleHanaInstance", &hbr.HanaInstanceArgs{
//				AlertSetting:        pulumi.String("INHERITED"),
//				HanaName:            pulumi.String("terraform-example"),
//				Host:                pulumi.String("1.1.1.1"),
//				InstanceNumber:      pulumi.Int(1),
//				Password:            pulumi.String("YouPassword123"),
//				ResourceGroupId:     *pulumi.String(exampleResourceGroups.Groups[0].Id),
//				Sid:                 pulumi.String("HXE"),
//				UseSsl:              pulumi.Bool(false),
//				UserName:            pulumi.String("admin"),
//				ValidateCertificate: pulumi.Bool(false),
//				VaultId:             exampleVault.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = hbr.NewHanaBackupClient(ctx, "default", &hbr.HanaBackupClientArgs{
//				VaultId: exampleVault.ID(),
//				ClientInfo: pulumi.All(exampleInstance.ID(), exampleHanaInstance.HanaInstanceId).ApplyT(func(_args []interface{}) (string, error) {
//					id := _args[0].(string)
//					hanaInstanceId := _args[1].(string)
//					return fmt.Sprintf("[ { \"instanceId\": \"%v\", \"clusterId\": \"%v\", \"sourceTypes\": [ \"HANA\" ]  }]", id, hanaInstanceId), nil
//				}).(pulumi.StringOutput),
//				AlertSetting: pulumi.String("INHERITED"),
//				UseHttps:     pulumi.Bool(true),
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
// Hybrid Backup Recovery (HBR) Hana Backup Client can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:hbr/hanaBackupClient:HanaBackupClient example <vault_id>:<client_id>
// ```
type HanaBackupClient struct {
	pulumi.CustomResourceState

	// The alert settings. Valid value: `INHERITED`.
	AlertSetting pulumi.StringOutput `pulumi:"alertSetting"`
	// The ID of the backup client.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// The installation information of the HBR clients.
	ClientInfo pulumi.StringPtrOutput `pulumi:"clientInfo"`
	// The ID of the SAP HANA instance.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The ID of the instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The status of the Hana Backup Client.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies whether to transmit data over HTTPS. Valid values: `true`, `false`.
	UseHttps pulumi.BoolPtrOutput `pulumi:"useHttps"`
	// The ID of the backup vault.
	VaultId pulumi.StringOutput `pulumi:"vaultId"`
}

// NewHanaBackupClient registers a new resource with the given unique name, arguments, and options.
func NewHanaBackupClient(ctx *pulumi.Context,
	name string, args *HanaBackupClientArgs, opts ...pulumi.ResourceOption) (*HanaBackupClient, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VaultId == nil {
		return nil, errors.New("invalid value for required argument 'VaultId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource HanaBackupClient
	err := ctx.RegisterResource("alicloud:hbr/hanaBackupClient:HanaBackupClient", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHanaBackupClient gets an existing HanaBackupClient resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHanaBackupClient(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HanaBackupClientState, opts ...pulumi.ResourceOption) (*HanaBackupClient, error) {
	var resource HanaBackupClient
	err := ctx.ReadResource("alicloud:hbr/hanaBackupClient:HanaBackupClient", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HanaBackupClient resources.
type hanaBackupClientState struct {
	// The alert settings. Valid value: `INHERITED`.
	AlertSetting *string `pulumi:"alertSetting"`
	// The ID of the backup client.
	ClientId *string `pulumi:"clientId"`
	// The installation information of the HBR clients.
	ClientInfo *string `pulumi:"clientInfo"`
	// The ID of the SAP HANA instance.
	ClusterId *string `pulumi:"clusterId"`
	// The ID of the instance.
	InstanceId *string `pulumi:"instanceId"`
	// The status of the Hana Backup Client.
	Status *string `pulumi:"status"`
	// Specifies whether to transmit data over HTTPS. Valid values: `true`, `false`.
	UseHttps *bool `pulumi:"useHttps"`
	// The ID of the backup vault.
	VaultId *string `pulumi:"vaultId"`
}

type HanaBackupClientState struct {
	// The alert settings. Valid value: `INHERITED`.
	AlertSetting pulumi.StringPtrInput
	// The ID of the backup client.
	ClientId pulumi.StringPtrInput
	// The installation information of the HBR clients.
	ClientInfo pulumi.StringPtrInput
	// The ID of the SAP HANA instance.
	ClusterId pulumi.StringPtrInput
	// The ID of the instance.
	InstanceId pulumi.StringPtrInput
	// The status of the Hana Backup Client.
	Status pulumi.StringPtrInput
	// Specifies whether to transmit data over HTTPS. Valid values: `true`, `false`.
	UseHttps pulumi.BoolPtrInput
	// The ID of the backup vault.
	VaultId pulumi.StringPtrInput
}

func (HanaBackupClientState) ElementType() reflect.Type {
	return reflect.TypeOf((*hanaBackupClientState)(nil)).Elem()
}

type hanaBackupClientArgs struct {
	// The alert settings. Valid value: `INHERITED`.
	AlertSetting *string `pulumi:"alertSetting"`
	// The installation information of the HBR clients.
	ClientInfo *string `pulumi:"clientInfo"`
	// Specifies whether to transmit data over HTTPS. Valid values: `true`, `false`.
	UseHttps *bool `pulumi:"useHttps"`
	// The ID of the backup vault.
	VaultId string `pulumi:"vaultId"`
}

// The set of arguments for constructing a HanaBackupClient resource.
type HanaBackupClientArgs struct {
	// The alert settings. Valid value: `INHERITED`.
	AlertSetting pulumi.StringPtrInput
	// The installation information of the HBR clients.
	ClientInfo pulumi.StringPtrInput
	// Specifies whether to transmit data over HTTPS. Valid values: `true`, `false`.
	UseHttps pulumi.BoolPtrInput
	// The ID of the backup vault.
	VaultId pulumi.StringInput
}

func (HanaBackupClientArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hanaBackupClientArgs)(nil)).Elem()
}

type HanaBackupClientInput interface {
	pulumi.Input

	ToHanaBackupClientOutput() HanaBackupClientOutput
	ToHanaBackupClientOutputWithContext(ctx context.Context) HanaBackupClientOutput
}

func (*HanaBackupClient) ElementType() reflect.Type {
	return reflect.TypeOf((**HanaBackupClient)(nil)).Elem()
}

func (i *HanaBackupClient) ToHanaBackupClientOutput() HanaBackupClientOutput {
	return i.ToHanaBackupClientOutputWithContext(context.Background())
}

func (i *HanaBackupClient) ToHanaBackupClientOutputWithContext(ctx context.Context) HanaBackupClientOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HanaBackupClientOutput)
}

// HanaBackupClientArrayInput is an input type that accepts HanaBackupClientArray and HanaBackupClientArrayOutput values.
// You can construct a concrete instance of `HanaBackupClientArrayInput` via:
//
//	HanaBackupClientArray{ HanaBackupClientArgs{...} }
type HanaBackupClientArrayInput interface {
	pulumi.Input

	ToHanaBackupClientArrayOutput() HanaBackupClientArrayOutput
	ToHanaBackupClientArrayOutputWithContext(context.Context) HanaBackupClientArrayOutput
}

type HanaBackupClientArray []HanaBackupClientInput

func (HanaBackupClientArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HanaBackupClient)(nil)).Elem()
}

func (i HanaBackupClientArray) ToHanaBackupClientArrayOutput() HanaBackupClientArrayOutput {
	return i.ToHanaBackupClientArrayOutputWithContext(context.Background())
}

func (i HanaBackupClientArray) ToHanaBackupClientArrayOutputWithContext(ctx context.Context) HanaBackupClientArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HanaBackupClientArrayOutput)
}

// HanaBackupClientMapInput is an input type that accepts HanaBackupClientMap and HanaBackupClientMapOutput values.
// You can construct a concrete instance of `HanaBackupClientMapInput` via:
//
//	HanaBackupClientMap{ "key": HanaBackupClientArgs{...} }
type HanaBackupClientMapInput interface {
	pulumi.Input

	ToHanaBackupClientMapOutput() HanaBackupClientMapOutput
	ToHanaBackupClientMapOutputWithContext(context.Context) HanaBackupClientMapOutput
}

type HanaBackupClientMap map[string]HanaBackupClientInput

func (HanaBackupClientMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HanaBackupClient)(nil)).Elem()
}

func (i HanaBackupClientMap) ToHanaBackupClientMapOutput() HanaBackupClientMapOutput {
	return i.ToHanaBackupClientMapOutputWithContext(context.Background())
}

func (i HanaBackupClientMap) ToHanaBackupClientMapOutputWithContext(ctx context.Context) HanaBackupClientMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HanaBackupClientMapOutput)
}

type HanaBackupClientOutput struct{ *pulumi.OutputState }

func (HanaBackupClientOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HanaBackupClient)(nil)).Elem()
}

func (o HanaBackupClientOutput) ToHanaBackupClientOutput() HanaBackupClientOutput {
	return o
}

func (o HanaBackupClientOutput) ToHanaBackupClientOutputWithContext(ctx context.Context) HanaBackupClientOutput {
	return o
}

// The alert settings. Valid value: `INHERITED`.
func (o HanaBackupClientOutput) AlertSetting() pulumi.StringOutput {
	return o.ApplyT(func(v *HanaBackupClient) pulumi.StringOutput { return v.AlertSetting }).(pulumi.StringOutput)
}

// The ID of the backup client.
func (o HanaBackupClientOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *HanaBackupClient) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

// The installation information of the HBR clients.
func (o HanaBackupClientOutput) ClientInfo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HanaBackupClient) pulumi.StringPtrOutput { return v.ClientInfo }).(pulumi.StringPtrOutput)
}

// The ID of the SAP HANA instance.
func (o HanaBackupClientOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *HanaBackupClient) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// The ID of the instance.
func (o HanaBackupClientOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *HanaBackupClient) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The status of the Hana Backup Client.
func (o HanaBackupClientOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *HanaBackupClient) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies whether to transmit data over HTTPS. Valid values: `true`, `false`.
func (o HanaBackupClientOutput) UseHttps() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HanaBackupClient) pulumi.BoolPtrOutput { return v.UseHttps }).(pulumi.BoolPtrOutput)
}

// The ID of the backup vault.
func (o HanaBackupClientOutput) VaultId() pulumi.StringOutput {
	return o.ApplyT(func(v *HanaBackupClient) pulumi.StringOutput { return v.VaultId }).(pulumi.StringOutput)
}

type HanaBackupClientArrayOutput struct{ *pulumi.OutputState }

func (HanaBackupClientArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HanaBackupClient)(nil)).Elem()
}

func (o HanaBackupClientArrayOutput) ToHanaBackupClientArrayOutput() HanaBackupClientArrayOutput {
	return o
}

func (o HanaBackupClientArrayOutput) ToHanaBackupClientArrayOutputWithContext(ctx context.Context) HanaBackupClientArrayOutput {
	return o
}

func (o HanaBackupClientArrayOutput) Index(i pulumi.IntInput) HanaBackupClientOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HanaBackupClient {
		return vs[0].([]*HanaBackupClient)[vs[1].(int)]
	}).(HanaBackupClientOutput)
}

type HanaBackupClientMapOutput struct{ *pulumi.OutputState }

func (HanaBackupClientMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HanaBackupClient)(nil)).Elem()
}

func (o HanaBackupClientMapOutput) ToHanaBackupClientMapOutput() HanaBackupClientMapOutput {
	return o
}

func (o HanaBackupClientMapOutput) ToHanaBackupClientMapOutputWithContext(ctx context.Context) HanaBackupClientMapOutput {
	return o
}

func (o HanaBackupClientMapOutput) MapIndex(k pulumi.StringInput) HanaBackupClientOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HanaBackupClient {
		return vs[0].(map[string]*HanaBackupClient)[vs[1].(string)]
	}).(HanaBackupClientOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HanaBackupClientInput)(nil)).Elem(), &HanaBackupClient{})
	pulumi.RegisterInputType(reflect.TypeOf((*HanaBackupClientArrayInput)(nil)).Elem(), HanaBackupClientArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HanaBackupClientMapInput)(nil)).Elem(), HanaBackupClientMap{})
	pulumi.RegisterOutputType(HanaBackupClientOutput{})
	pulumi.RegisterOutputType(HanaBackupClientArrayOutput{})
	pulumi.RegisterOutputType(HanaBackupClientMapOutput{})
}

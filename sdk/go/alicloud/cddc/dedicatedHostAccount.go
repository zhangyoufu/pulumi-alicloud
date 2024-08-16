// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cddc

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a ApsaraDB for MyBase Dedicated Host Account resource.
//
// For information about ApsaraDB for MyBase Dedicated Host Account and how to use it, see [What is Dedicated Host Account](https://www.alibabacloud.com/help/en/apsaradb-for-mybase/latest/creatededicatedhostaccount).
//
// > **NOTE:** Available since v1.148.0.
//
// > **NOTE:** Each Dedicated host can have only one account. Before you create an account for a host, make sure that the existing account is deleted.
//
// > **DEPRECATED:**  This resource has been [deprecated](https://www.alibabacloud.com/help/en/apsaradb-for-mybase/latest/notice-stop-selling-mybase-hosted-instances-from-august-31-2023) from version `1.225.1`.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cddc"
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
//			_default, err := cddc.GetZones(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			defaultNetwork, err := vpc.NewNetwork(ctx, "default", &vpc.NetworkArgs{
//				VpcName:   pulumi.String(name),
//				CidrBlock: pulumi.String("10.4.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultSwitch, err := vpc.NewSwitch(ctx, "default", &vpc.SwitchArgs{
//				VswitchName: pulumi.String(name),
//				CidrBlock:   pulumi.String("10.4.0.0/24"),
//				VpcId:       defaultNetwork.ID(),
//				ZoneId:      pulumi.String(_default.Ids[0]),
//			})
//			if err != nil {
//				return err
//			}
//			defaultDedicatedHostGroup, err := cddc.NewDedicatedHostGroup(ctx, "default", &cddc.DedicatedHostGroupArgs{
//				Engine:                 pulumi.String("MySQL"),
//				VpcId:                  defaultNetwork.ID(),
//				CpuAllocationRatio:     pulumi.Int(101),
//				MemAllocationRatio:     pulumi.Int(50),
//				DiskAllocationRatio:    pulumi.Int(200),
//				AllocationPolicy:       pulumi.String("Evenly"),
//				HostReplacePolicy:      pulumi.String("Manual"),
//				DedicatedHostGroupDesc: pulumi.String(name),
//				OpenPermission:         pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			defaultGetHostEcsLevelInfos, err := cddc.GetHostEcsLevelInfos(ctx, &cddc.GetHostEcsLevelInfosArgs{
//				DbType:      "mysql",
//				ZoneId:      _default.Ids[0],
//				StorageType: "cloud_essd",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultDedicatedHost, err := cddc.NewDedicatedHost(ctx, "default", &cddc.DedicatedHostArgs{
//				HostName:             pulumi.String(name),
//				DedicatedHostGroupId: defaultDedicatedHostGroup.ID(),
//				HostClass:            pulumi.String(defaultGetHostEcsLevelInfos.Infos[0].ResClassCode),
//				ZoneId:               pulumi.String(_default.Ids[0]),
//				VswitchId:            defaultSwitch.ID(),
//				PaymentType:          pulumi.String("Subscription"),
//				Tags: pulumi.StringMap{
//					"Created": pulumi.String("TF"),
//					"For":     pulumi.String("CDDC_DEDICATED"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cddc.NewDedicatedHostAccount(ctx, "default", &cddc.DedicatedHostAccountArgs{
//				AccountName:     pulumi.String(name),
//				AccountPassword: pulumi.String("Password1234"),
//				DedicatedHostId: defaultDedicatedHost.DedicatedHostId,
//				AccountType:     pulumi.String("Normal"),
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
// ApsaraDB for MyBase Dedicated Host Account can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:cddc/dedicatedHostAccount:DedicatedHostAccount example <dedicated_host_id>:<account_name>
// ```
type DedicatedHostAccount struct {
	pulumi.CustomResourceState

	// The name of the Dedicated host account. The account name must be 2 to 16 characters in length, contain lower case letters, digits, and underscore(_). At the same time, the name must start with a letter and end with a letter or number.
	AccountName pulumi.StringOutput `pulumi:"accountName"`
	// The password of the Dedicated host account. The account password must be 6 to 32 characters in length, and can contain letters, digits, and special characters `!@#$%^&*()_+-=`.
	AccountPassword pulumi.StringOutput `pulumi:"accountPassword"`
	// The type of the Dedicated host account. Valid values: `Admin`, `Normal`.
	AccountType pulumi.StringPtrOutput `pulumi:"accountType"`
	// The ID of Dedicated the host.
	DedicatedHostId pulumi.StringOutput `pulumi:"dedicatedHostId"`
}

// NewDedicatedHostAccount registers a new resource with the given unique name, arguments, and options.
func NewDedicatedHostAccount(ctx *pulumi.Context,
	name string, args *DedicatedHostAccountArgs, opts ...pulumi.ResourceOption) (*DedicatedHostAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.AccountPassword == nil {
		return nil, errors.New("invalid value for required argument 'AccountPassword'")
	}
	if args.DedicatedHostId == nil {
		return nil, errors.New("invalid value for required argument 'DedicatedHostId'")
	}
	if args.AccountPassword != nil {
		args.AccountPassword = pulumi.ToSecret(args.AccountPassword).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"accountPassword",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DedicatedHostAccount
	err := ctx.RegisterResource("alicloud:cddc/dedicatedHostAccount:DedicatedHostAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDedicatedHostAccount gets an existing DedicatedHostAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDedicatedHostAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DedicatedHostAccountState, opts ...pulumi.ResourceOption) (*DedicatedHostAccount, error) {
	var resource DedicatedHostAccount
	err := ctx.ReadResource("alicloud:cddc/dedicatedHostAccount:DedicatedHostAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DedicatedHostAccount resources.
type dedicatedHostAccountState struct {
	// The name of the Dedicated host account. The account name must be 2 to 16 characters in length, contain lower case letters, digits, and underscore(_). At the same time, the name must start with a letter and end with a letter or number.
	AccountName *string `pulumi:"accountName"`
	// The password of the Dedicated host account. The account password must be 6 to 32 characters in length, and can contain letters, digits, and special characters `!@#$%^&*()_+-=`.
	AccountPassword *string `pulumi:"accountPassword"`
	// The type of the Dedicated host account. Valid values: `Admin`, `Normal`.
	AccountType *string `pulumi:"accountType"`
	// The ID of Dedicated the host.
	DedicatedHostId *string `pulumi:"dedicatedHostId"`
}

type DedicatedHostAccountState struct {
	// The name of the Dedicated host account. The account name must be 2 to 16 characters in length, contain lower case letters, digits, and underscore(_). At the same time, the name must start with a letter and end with a letter or number.
	AccountName pulumi.StringPtrInput
	// The password of the Dedicated host account. The account password must be 6 to 32 characters in length, and can contain letters, digits, and special characters `!@#$%^&*()_+-=`.
	AccountPassword pulumi.StringPtrInput
	// The type of the Dedicated host account. Valid values: `Admin`, `Normal`.
	AccountType pulumi.StringPtrInput
	// The ID of Dedicated the host.
	DedicatedHostId pulumi.StringPtrInput
}

func (DedicatedHostAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedHostAccountState)(nil)).Elem()
}

type dedicatedHostAccountArgs struct {
	// The name of the Dedicated host account. The account name must be 2 to 16 characters in length, contain lower case letters, digits, and underscore(_). At the same time, the name must start with a letter and end with a letter or number.
	AccountName string `pulumi:"accountName"`
	// The password of the Dedicated host account. The account password must be 6 to 32 characters in length, and can contain letters, digits, and special characters `!@#$%^&*()_+-=`.
	AccountPassword string `pulumi:"accountPassword"`
	// The type of the Dedicated host account. Valid values: `Admin`, `Normal`.
	AccountType *string `pulumi:"accountType"`
	// The ID of Dedicated the host.
	DedicatedHostId string `pulumi:"dedicatedHostId"`
}

// The set of arguments for constructing a DedicatedHostAccount resource.
type DedicatedHostAccountArgs struct {
	// The name of the Dedicated host account. The account name must be 2 to 16 characters in length, contain lower case letters, digits, and underscore(_). At the same time, the name must start with a letter and end with a letter or number.
	AccountName pulumi.StringInput
	// The password of the Dedicated host account. The account password must be 6 to 32 characters in length, and can contain letters, digits, and special characters `!@#$%^&*()_+-=`.
	AccountPassword pulumi.StringInput
	// The type of the Dedicated host account. Valid values: `Admin`, `Normal`.
	AccountType pulumi.StringPtrInput
	// The ID of Dedicated the host.
	DedicatedHostId pulumi.StringInput
}

func (DedicatedHostAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedHostAccountArgs)(nil)).Elem()
}

type DedicatedHostAccountInput interface {
	pulumi.Input

	ToDedicatedHostAccountOutput() DedicatedHostAccountOutput
	ToDedicatedHostAccountOutputWithContext(ctx context.Context) DedicatedHostAccountOutput
}

func (*DedicatedHostAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedHostAccount)(nil)).Elem()
}

func (i *DedicatedHostAccount) ToDedicatedHostAccountOutput() DedicatedHostAccountOutput {
	return i.ToDedicatedHostAccountOutputWithContext(context.Background())
}

func (i *DedicatedHostAccount) ToDedicatedHostAccountOutputWithContext(ctx context.Context) DedicatedHostAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedHostAccountOutput)
}

// DedicatedHostAccountArrayInput is an input type that accepts DedicatedHostAccountArray and DedicatedHostAccountArrayOutput values.
// You can construct a concrete instance of `DedicatedHostAccountArrayInput` via:
//
//	DedicatedHostAccountArray{ DedicatedHostAccountArgs{...} }
type DedicatedHostAccountArrayInput interface {
	pulumi.Input

	ToDedicatedHostAccountArrayOutput() DedicatedHostAccountArrayOutput
	ToDedicatedHostAccountArrayOutputWithContext(context.Context) DedicatedHostAccountArrayOutput
}

type DedicatedHostAccountArray []DedicatedHostAccountInput

func (DedicatedHostAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DedicatedHostAccount)(nil)).Elem()
}

func (i DedicatedHostAccountArray) ToDedicatedHostAccountArrayOutput() DedicatedHostAccountArrayOutput {
	return i.ToDedicatedHostAccountArrayOutputWithContext(context.Background())
}

func (i DedicatedHostAccountArray) ToDedicatedHostAccountArrayOutputWithContext(ctx context.Context) DedicatedHostAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedHostAccountArrayOutput)
}

// DedicatedHostAccountMapInput is an input type that accepts DedicatedHostAccountMap and DedicatedHostAccountMapOutput values.
// You can construct a concrete instance of `DedicatedHostAccountMapInput` via:
//
//	DedicatedHostAccountMap{ "key": DedicatedHostAccountArgs{...} }
type DedicatedHostAccountMapInput interface {
	pulumi.Input

	ToDedicatedHostAccountMapOutput() DedicatedHostAccountMapOutput
	ToDedicatedHostAccountMapOutputWithContext(context.Context) DedicatedHostAccountMapOutput
}

type DedicatedHostAccountMap map[string]DedicatedHostAccountInput

func (DedicatedHostAccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DedicatedHostAccount)(nil)).Elem()
}

func (i DedicatedHostAccountMap) ToDedicatedHostAccountMapOutput() DedicatedHostAccountMapOutput {
	return i.ToDedicatedHostAccountMapOutputWithContext(context.Background())
}

func (i DedicatedHostAccountMap) ToDedicatedHostAccountMapOutputWithContext(ctx context.Context) DedicatedHostAccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedHostAccountMapOutput)
}

type DedicatedHostAccountOutput struct{ *pulumi.OutputState }

func (DedicatedHostAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedHostAccount)(nil)).Elem()
}

func (o DedicatedHostAccountOutput) ToDedicatedHostAccountOutput() DedicatedHostAccountOutput {
	return o
}

func (o DedicatedHostAccountOutput) ToDedicatedHostAccountOutputWithContext(ctx context.Context) DedicatedHostAccountOutput {
	return o
}

// The name of the Dedicated host account. The account name must be 2 to 16 characters in length, contain lower case letters, digits, and underscore(_). At the same time, the name must start with a letter and end with a letter or number.
func (o DedicatedHostAccountOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedHostAccount) pulumi.StringOutput { return v.AccountName }).(pulumi.StringOutput)
}

// The password of the Dedicated host account. The account password must be 6 to 32 characters in length, and can contain letters, digits, and special characters `!@#$%^&*()_+-=`.
func (o DedicatedHostAccountOutput) AccountPassword() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedHostAccount) pulumi.StringOutput { return v.AccountPassword }).(pulumi.StringOutput)
}

// The type of the Dedicated host account. Valid values: `Admin`, `Normal`.
func (o DedicatedHostAccountOutput) AccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DedicatedHostAccount) pulumi.StringPtrOutput { return v.AccountType }).(pulumi.StringPtrOutput)
}

// The ID of Dedicated the host.
func (o DedicatedHostAccountOutput) DedicatedHostId() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedHostAccount) pulumi.StringOutput { return v.DedicatedHostId }).(pulumi.StringOutput)
}

type DedicatedHostAccountArrayOutput struct{ *pulumi.OutputState }

func (DedicatedHostAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DedicatedHostAccount)(nil)).Elem()
}

func (o DedicatedHostAccountArrayOutput) ToDedicatedHostAccountArrayOutput() DedicatedHostAccountArrayOutput {
	return o
}

func (o DedicatedHostAccountArrayOutput) ToDedicatedHostAccountArrayOutputWithContext(ctx context.Context) DedicatedHostAccountArrayOutput {
	return o
}

func (o DedicatedHostAccountArrayOutput) Index(i pulumi.IntInput) DedicatedHostAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DedicatedHostAccount {
		return vs[0].([]*DedicatedHostAccount)[vs[1].(int)]
	}).(DedicatedHostAccountOutput)
}

type DedicatedHostAccountMapOutput struct{ *pulumi.OutputState }

func (DedicatedHostAccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DedicatedHostAccount)(nil)).Elem()
}

func (o DedicatedHostAccountMapOutput) ToDedicatedHostAccountMapOutput() DedicatedHostAccountMapOutput {
	return o
}

func (o DedicatedHostAccountMapOutput) ToDedicatedHostAccountMapOutputWithContext(ctx context.Context) DedicatedHostAccountMapOutput {
	return o
}

func (o DedicatedHostAccountMapOutput) MapIndex(k pulumi.StringInput) DedicatedHostAccountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DedicatedHostAccount {
		return vs[0].(map[string]*DedicatedHostAccount)[vs[1].(string)]
	}).(DedicatedHostAccountOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedHostAccountInput)(nil)).Elem(), &DedicatedHostAccount{})
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedHostAccountArrayInput)(nil)).Elem(), DedicatedHostAccountArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedHostAccountMapInput)(nil)).Elem(), DedicatedHostAccountMap{})
	pulumi.RegisterOutputType(DedicatedHostAccountOutput{})
	pulumi.RegisterOutputType(DedicatedHostAccountArrayOutput{})
	pulumi.RegisterOutputType(DedicatedHostAccountMapOutput{})
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mongodb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a MongoDB Serverless Instance resource.
//
// For information about MongoDB Serverless Instance and how to use it, see [What is Serverless Instance](https://www.alibabacloud.com/help/doc-detail/26558.html).
//
// > **NOTE:** Deprecated since v1.214.0.
//
// > **DEPRECATED:**  This resource has been deprecated from version `1.214.0`.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/mongodb"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/resourcemanager"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaultZones, err := mongodb.GetZones(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			defaultNetworks, err := vpc.GetNetworks(ctx, &vpc.GetNetworksArgs{
//				NameRegex: pulumi.StringRef("default-NODELETING"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultSwitches, err := vpc.GetSwitches(ctx, &vpc.GetSwitchesArgs{
//				VpcId:  pulumi.StringRef(defaultNetworks.Ids[0]),
//				ZoneId: pulumi.StringRef(defaultZones.Zones[0].Id),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultResourceGroups, err := resourcemanager.GetResourceGroups(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = mongodb.NewServerlessInstance(ctx, "example", &mongodb.ServerlessInstanceArgs{
//				AccountPassword:       pulumi.String("Abc12345"),
//				DbInstanceDescription: pulumi.String("example_value"),
//				DbInstanceStorage:     pulumi.Int(5),
//				StorageEngine:         pulumi.String("WiredTiger"),
//				CapacityUnit:          pulumi.Int(100),
//				Engine:                pulumi.String("MongoDB"),
//				ResourceGroupId:       *pulumi.String(defaultResourceGroups.Groups[0].Id),
//				EngineVersion:         pulumi.String("4.2"),
//				Period:                pulumi.Int(1),
//				PeriodPriceType:       pulumi.String("Month"),
//				VpcId:                 *pulumi.String(defaultNetworks.Ids[0]),
//				ZoneId:                *pulumi.String(defaultZones.Zones[0].Id),
//				VswitchId:             *pulumi.String(defaultSwitches.Ids[0]),
//				Tags: pulumi.Map{
//					"Created": pulumi.Any("MongodbServerlessInstance"),
//					"For":     pulumi.Any("TF"),
//				},
//				SecurityIpGroups: mongodb.ServerlessInstanceSecurityIpGroupArray{
//					&mongodb.ServerlessInstanceSecurityIpGroupArgs{
//						SecurityIpGroupAttribute: pulumi.String("example_value"),
//						SecurityIpGroupName:      pulumi.String("example_value"),
//						SecurityIpList:           pulumi.String("192.168.0.1"),
//					},
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
// MongoDB Serverless Instance can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:mongodb/serverlessInstance:ServerlessInstance example <id>
// ```
type ServerlessInstance struct {
	pulumi.CustomResourceState

	// The password of the database logon account.
	// * The password length is `8` to `32` bits.
	// * The password consists of at least any three of uppercase letters, lowercase letters, numbers, and special characters. The special character is `!#$%!^(MISSING)&*()_+-=`. The MongoDB Serverless instance provides a default database login account. This account cannot be modified. You can only set or modify the password for this account.
	AccountPassword pulumi.StringOutput `pulumi:"accountPassword"`
	// Set whether the instance is automatically renewed.
	AutoRenew pulumi.BoolPtrOutput `pulumi:"autoRenew"`
	// The I/O throughput consumed by the instance. Valid values: `100` to `8000`.
	CapacityUnit pulumi.IntOutput `pulumi:"capacityUnit"`
	// The db instance description.
	DbInstanceDescription pulumi.StringPtrOutput `pulumi:"dbInstanceDescription"`
	// The db instance storage. Valid values: `1` to `100`.
	DbInstanceStorage pulumi.IntOutput `pulumi:"dbInstanceStorage"`
	// The database engine of the instance. Valid values: `MongoDB`.
	Engine pulumi.StringOutput `pulumi:"engine"`
	// The database version number. Valid values: `4.2`.
	EngineVersion pulumi.StringOutput `pulumi:"engineVersion"`
	// The end time of the maintenance window. Specify the time in the `HH:mmZ` format. The time must be in UTC. **NOTE:** The difference between the start time and end time must be one hour. For example, if `maintainStartTime` is `01:00Z`, `maintainEndTime` must be `02:00Z`.
	MaintainEndTime pulumi.StringOutput `pulumi:"maintainEndTime"`
	// The start time of the maintenance window. Specify the time in the `HH:mmZ` format. The time must be in UTC.
	MaintainStartTime pulumi.StringOutput `pulumi:"maintainStartTime"`
	// The purchase duration of the instance, in months. Valid values: `1` to `9`, `12`, `24`, `36`, `60`.
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// The period price type. Valid values: `Day`, `Month`.
	PeriodPriceType pulumi.StringPtrOutput `pulumi:"periodPriceType"`
	// The ID of the resource group.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// An array that consists of the information of IP whitelists.
	SecurityIpGroups ServerlessInstanceSecurityIpGroupArrayOutput `pulumi:"securityIpGroups"`
	// The instance status. For more information, see the instance Status Table.
	Status pulumi.StringOutput `pulumi:"status"`
	// The storage engine used by the instance. Valid values: `WiredTiger`.
	StorageEngine pulumi.StringOutput `pulumi:"storageEngine"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The ID of the VPC network.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// The of the vswitch.
	VswitchId pulumi.StringOutput `pulumi:"vswitchId"`
	// The ID of the zone. Use this parameter to specify the zone created by the instance.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewServerlessInstance registers a new resource with the given unique name, arguments, and options.
func NewServerlessInstance(ctx *pulumi.Context,
	name string, args *ServerlessInstanceArgs, opts ...pulumi.ResourceOption) (*ServerlessInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountPassword == nil {
		return nil, errors.New("invalid value for required argument 'AccountPassword'")
	}
	if args.CapacityUnit == nil {
		return nil, errors.New("invalid value for required argument 'CapacityUnit'")
	}
	if args.DbInstanceStorage == nil {
		return nil, errors.New("invalid value for required argument 'DbInstanceStorage'")
	}
	if args.EngineVersion == nil {
		return nil, errors.New("invalid value for required argument 'EngineVersion'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	if args.VswitchId == nil {
		return nil, errors.New("invalid value for required argument 'VswitchId'")
	}
	if args.ZoneId == nil {
		return nil, errors.New("invalid value for required argument 'ZoneId'")
	}
	if args.AccountPassword != nil {
		args.AccountPassword = pulumi.ToSecret(args.AccountPassword).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"accountPassword",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServerlessInstance
	err := ctx.RegisterResource("alicloud:mongodb/serverlessInstance:ServerlessInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerlessInstance gets an existing ServerlessInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerlessInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerlessInstanceState, opts ...pulumi.ResourceOption) (*ServerlessInstance, error) {
	var resource ServerlessInstance
	err := ctx.ReadResource("alicloud:mongodb/serverlessInstance:ServerlessInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerlessInstance resources.
type serverlessInstanceState struct {
	// The password of the database logon account.
	// * The password length is `8` to `32` bits.
	// * The password consists of at least any three of uppercase letters, lowercase letters, numbers, and special characters. The special character is `!#$%!^(MISSING)&*()_+-=`. The MongoDB Serverless instance provides a default database login account. This account cannot be modified. You can only set or modify the password for this account.
	AccountPassword *string `pulumi:"accountPassword"`
	// Set whether the instance is automatically renewed.
	AutoRenew *bool `pulumi:"autoRenew"`
	// The I/O throughput consumed by the instance. Valid values: `100` to `8000`.
	CapacityUnit *int `pulumi:"capacityUnit"`
	// The db instance description.
	DbInstanceDescription *string `pulumi:"dbInstanceDescription"`
	// The db instance storage. Valid values: `1` to `100`.
	DbInstanceStorage *int `pulumi:"dbInstanceStorage"`
	// The database engine of the instance. Valid values: `MongoDB`.
	Engine *string `pulumi:"engine"`
	// The database version number. Valid values: `4.2`.
	EngineVersion *string `pulumi:"engineVersion"`
	// The end time of the maintenance window. Specify the time in the `HH:mmZ` format. The time must be in UTC. **NOTE:** The difference between the start time and end time must be one hour. For example, if `maintainStartTime` is `01:00Z`, `maintainEndTime` must be `02:00Z`.
	MaintainEndTime *string `pulumi:"maintainEndTime"`
	// The start time of the maintenance window. Specify the time in the `HH:mmZ` format. The time must be in UTC.
	MaintainStartTime *string `pulumi:"maintainStartTime"`
	// The purchase duration of the instance, in months. Valid values: `1` to `9`, `12`, `24`, `36`, `60`.
	Period *int `pulumi:"period"`
	// The period price type. Valid values: `Day`, `Month`.
	PeriodPriceType *string `pulumi:"periodPriceType"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// An array that consists of the information of IP whitelists.
	SecurityIpGroups []ServerlessInstanceSecurityIpGroup `pulumi:"securityIpGroups"`
	// The instance status. For more information, see the instance Status Table.
	Status *string `pulumi:"status"`
	// The storage engine used by the instance. Valid values: `WiredTiger`.
	StorageEngine *string `pulumi:"storageEngine"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The ID of the VPC network.
	VpcId *string `pulumi:"vpcId"`
	// The of the vswitch.
	VswitchId *string `pulumi:"vswitchId"`
	// The ID of the zone. Use this parameter to specify the zone created by the instance.
	ZoneId *string `pulumi:"zoneId"`
}

type ServerlessInstanceState struct {
	// The password of the database logon account.
	// * The password length is `8` to `32` bits.
	// * The password consists of at least any three of uppercase letters, lowercase letters, numbers, and special characters. The special character is `!#$%!^(MISSING)&*()_+-=`. The MongoDB Serverless instance provides a default database login account. This account cannot be modified. You can only set or modify the password for this account.
	AccountPassword pulumi.StringPtrInput
	// Set whether the instance is automatically renewed.
	AutoRenew pulumi.BoolPtrInput
	// The I/O throughput consumed by the instance. Valid values: `100` to `8000`.
	CapacityUnit pulumi.IntPtrInput
	// The db instance description.
	DbInstanceDescription pulumi.StringPtrInput
	// The db instance storage. Valid values: `1` to `100`.
	DbInstanceStorage pulumi.IntPtrInput
	// The database engine of the instance. Valid values: `MongoDB`.
	Engine pulumi.StringPtrInput
	// The database version number. Valid values: `4.2`.
	EngineVersion pulumi.StringPtrInput
	// The end time of the maintenance window. Specify the time in the `HH:mmZ` format. The time must be in UTC. **NOTE:** The difference between the start time and end time must be one hour. For example, if `maintainStartTime` is `01:00Z`, `maintainEndTime` must be `02:00Z`.
	MaintainEndTime pulumi.StringPtrInput
	// The start time of the maintenance window. Specify the time in the `HH:mmZ` format. The time must be in UTC.
	MaintainStartTime pulumi.StringPtrInput
	// The purchase duration of the instance, in months. Valid values: `1` to `9`, `12`, `24`, `36`, `60`.
	Period pulumi.IntPtrInput
	// The period price type. Valid values: `Day`, `Month`.
	PeriodPriceType pulumi.StringPtrInput
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	// An array that consists of the information of IP whitelists.
	SecurityIpGroups ServerlessInstanceSecurityIpGroupArrayInput
	// The instance status. For more information, see the instance Status Table.
	Status pulumi.StringPtrInput
	// The storage engine used by the instance. Valid values: `WiredTiger`.
	StorageEngine pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The ID of the VPC network.
	VpcId pulumi.StringPtrInput
	// The of the vswitch.
	VswitchId pulumi.StringPtrInput
	// The ID of the zone. Use this parameter to specify the zone created by the instance.
	ZoneId pulumi.StringPtrInput
}

func (ServerlessInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverlessInstanceState)(nil)).Elem()
}

type serverlessInstanceArgs struct {
	// The password of the database logon account.
	// * The password length is `8` to `32` bits.
	// * The password consists of at least any three of uppercase letters, lowercase letters, numbers, and special characters. The special character is `!#$%!^(MISSING)&*()_+-=`. The MongoDB Serverless instance provides a default database login account. This account cannot be modified. You can only set or modify the password for this account.
	AccountPassword string `pulumi:"accountPassword"`
	// Set whether the instance is automatically renewed.
	AutoRenew *bool `pulumi:"autoRenew"`
	// The I/O throughput consumed by the instance. Valid values: `100` to `8000`.
	CapacityUnit int `pulumi:"capacityUnit"`
	// The db instance description.
	DbInstanceDescription *string `pulumi:"dbInstanceDescription"`
	// The db instance storage. Valid values: `1` to `100`.
	DbInstanceStorage int `pulumi:"dbInstanceStorage"`
	// The database engine of the instance. Valid values: `MongoDB`.
	Engine *string `pulumi:"engine"`
	// The database version number. Valid values: `4.2`.
	EngineVersion string `pulumi:"engineVersion"`
	// The end time of the maintenance window. Specify the time in the `HH:mmZ` format. The time must be in UTC. **NOTE:** The difference between the start time and end time must be one hour. For example, if `maintainStartTime` is `01:00Z`, `maintainEndTime` must be `02:00Z`.
	MaintainEndTime *string `pulumi:"maintainEndTime"`
	// The start time of the maintenance window. Specify the time in the `HH:mmZ` format. The time must be in UTC.
	MaintainStartTime *string `pulumi:"maintainStartTime"`
	// The purchase duration of the instance, in months. Valid values: `1` to `9`, `12`, `24`, `36`, `60`.
	Period *int `pulumi:"period"`
	// The period price type. Valid values: `Day`, `Month`.
	PeriodPriceType *string `pulumi:"periodPriceType"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// An array that consists of the information of IP whitelists.
	SecurityIpGroups []ServerlessInstanceSecurityIpGroup `pulumi:"securityIpGroups"`
	// The storage engine used by the instance. Valid values: `WiredTiger`.
	StorageEngine *string `pulumi:"storageEngine"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The ID of the VPC network.
	VpcId string `pulumi:"vpcId"`
	// The of the vswitch.
	VswitchId string `pulumi:"vswitchId"`
	// The ID of the zone. Use this parameter to specify the zone created by the instance.
	ZoneId string `pulumi:"zoneId"`
}

// The set of arguments for constructing a ServerlessInstance resource.
type ServerlessInstanceArgs struct {
	// The password of the database logon account.
	// * The password length is `8` to `32` bits.
	// * The password consists of at least any three of uppercase letters, lowercase letters, numbers, and special characters. The special character is `!#$%!^(MISSING)&*()_+-=`. The MongoDB Serverless instance provides a default database login account. This account cannot be modified. You can only set or modify the password for this account.
	AccountPassword pulumi.StringInput
	// Set whether the instance is automatically renewed.
	AutoRenew pulumi.BoolPtrInput
	// The I/O throughput consumed by the instance. Valid values: `100` to `8000`.
	CapacityUnit pulumi.IntInput
	// The db instance description.
	DbInstanceDescription pulumi.StringPtrInput
	// The db instance storage. Valid values: `1` to `100`.
	DbInstanceStorage pulumi.IntInput
	// The database engine of the instance. Valid values: `MongoDB`.
	Engine pulumi.StringPtrInput
	// The database version number. Valid values: `4.2`.
	EngineVersion pulumi.StringInput
	// The end time of the maintenance window. Specify the time in the `HH:mmZ` format. The time must be in UTC. **NOTE:** The difference between the start time and end time must be one hour. For example, if `maintainStartTime` is `01:00Z`, `maintainEndTime` must be `02:00Z`.
	MaintainEndTime pulumi.StringPtrInput
	// The start time of the maintenance window. Specify the time in the `HH:mmZ` format. The time must be in UTC.
	MaintainStartTime pulumi.StringPtrInput
	// The purchase duration of the instance, in months. Valid values: `1` to `9`, `12`, `24`, `36`, `60`.
	Period pulumi.IntPtrInput
	// The period price type. Valid values: `Day`, `Month`.
	PeriodPriceType pulumi.StringPtrInput
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	// An array that consists of the information of IP whitelists.
	SecurityIpGroups ServerlessInstanceSecurityIpGroupArrayInput
	// The storage engine used by the instance. Valid values: `WiredTiger`.
	StorageEngine pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The ID of the VPC network.
	VpcId pulumi.StringInput
	// The of the vswitch.
	VswitchId pulumi.StringInput
	// The ID of the zone. Use this parameter to specify the zone created by the instance.
	ZoneId pulumi.StringInput
}

func (ServerlessInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverlessInstanceArgs)(nil)).Elem()
}

type ServerlessInstanceInput interface {
	pulumi.Input

	ToServerlessInstanceOutput() ServerlessInstanceOutput
	ToServerlessInstanceOutputWithContext(ctx context.Context) ServerlessInstanceOutput
}

func (*ServerlessInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerlessInstance)(nil)).Elem()
}

func (i *ServerlessInstance) ToServerlessInstanceOutput() ServerlessInstanceOutput {
	return i.ToServerlessInstanceOutputWithContext(context.Background())
}

func (i *ServerlessInstance) ToServerlessInstanceOutputWithContext(ctx context.Context) ServerlessInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerlessInstanceOutput)
}

// ServerlessInstanceArrayInput is an input type that accepts ServerlessInstanceArray and ServerlessInstanceArrayOutput values.
// You can construct a concrete instance of `ServerlessInstanceArrayInput` via:
//
//	ServerlessInstanceArray{ ServerlessInstanceArgs{...} }
type ServerlessInstanceArrayInput interface {
	pulumi.Input

	ToServerlessInstanceArrayOutput() ServerlessInstanceArrayOutput
	ToServerlessInstanceArrayOutputWithContext(context.Context) ServerlessInstanceArrayOutput
}

type ServerlessInstanceArray []ServerlessInstanceInput

func (ServerlessInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerlessInstance)(nil)).Elem()
}

func (i ServerlessInstanceArray) ToServerlessInstanceArrayOutput() ServerlessInstanceArrayOutput {
	return i.ToServerlessInstanceArrayOutputWithContext(context.Background())
}

func (i ServerlessInstanceArray) ToServerlessInstanceArrayOutputWithContext(ctx context.Context) ServerlessInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerlessInstanceArrayOutput)
}

// ServerlessInstanceMapInput is an input type that accepts ServerlessInstanceMap and ServerlessInstanceMapOutput values.
// You can construct a concrete instance of `ServerlessInstanceMapInput` via:
//
//	ServerlessInstanceMap{ "key": ServerlessInstanceArgs{...} }
type ServerlessInstanceMapInput interface {
	pulumi.Input

	ToServerlessInstanceMapOutput() ServerlessInstanceMapOutput
	ToServerlessInstanceMapOutputWithContext(context.Context) ServerlessInstanceMapOutput
}

type ServerlessInstanceMap map[string]ServerlessInstanceInput

func (ServerlessInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerlessInstance)(nil)).Elem()
}

func (i ServerlessInstanceMap) ToServerlessInstanceMapOutput() ServerlessInstanceMapOutput {
	return i.ToServerlessInstanceMapOutputWithContext(context.Background())
}

func (i ServerlessInstanceMap) ToServerlessInstanceMapOutputWithContext(ctx context.Context) ServerlessInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerlessInstanceMapOutput)
}

type ServerlessInstanceOutput struct{ *pulumi.OutputState }

func (ServerlessInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerlessInstance)(nil)).Elem()
}

func (o ServerlessInstanceOutput) ToServerlessInstanceOutput() ServerlessInstanceOutput {
	return o
}

func (o ServerlessInstanceOutput) ToServerlessInstanceOutputWithContext(ctx context.Context) ServerlessInstanceOutput {
	return o
}

// The password of the database logon account.
// * The password length is `8` to `32` bits.
// * The password consists of at least any three of uppercase letters, lowercase letters, numbers, and special characters. The special character is `!#$%!^(MISSING)&*()_+-=`. The MongoDB Serverless instance provides a default database login account. This account cannot be modified. You can only set or modify the password for this account.
func (o ServerlessInstanceOutput) AccountPassword() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerlessInstance) pulumi.StringOutput { return v.AccountPassword }).(pulumi.StringOutput)
}

// Set whether the instance is automatically renewed.
func (o ServerlessInstanceOutput) AutoRenew() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServerlessInstance) pulumi.BoolPtrOutput { return v.AutoRenew }).(pulumi.BoolPtrOutput)
}

// The I/O throughput consumed by the instance. Valid values: `100` to `8000`.
func (o ServerlessInstanceOutput) CapacityUnit() pulumi.IntOutput {
	return o.ApplyT(func(v *ServerlessInstance) pulumi.IntOutput { return v.CapacityUnit }).(pulumi.IntOutput)
}

// The db instance description.
func (o ServerlessInstanceOutput) DbInstanceDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerlessInstance) pulumi.StringPtrOutput { return v.DbInstanceDescription }).(pulumi.StringPtrOutput)
}

// The db instance storage. Valid values: `1` to `100`.
func (o ServerlessInstanceOutput) DbInstanceStorage() pulumi.IntOutput {
	return o.ApplyT(func(v *ServerlessInstance) pulumi.IntOutput { return v.DbInstanceStorage }).(pulumi.IntOutput)
}

// The database engine of the instance. Valid values: `MongoDB`.
func (o ServerlessInstanceOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerlessInstance) pulumi.StringOutput { return v.Engine }).(pulumi.StringOutput)
}

// The database version number. Valid values: `4.2`.
func (o ServerlessInstanceOutput) EngineVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerlessInstance) pulumi.StringOutput { return v.EngineVersion }).(pulumi.StringOutput)
}

// The end time of the maintenance window. Specify the time in the `HH:mmZ` format. The time must be in UTC. **NOTE:** The difference between the start time and end time must be one hour. For example, if `maintainStartTime` is `01:00Z`, `maintainEndTime` must be `02:00Z`.
func (o ServerlessInstanceOutput) MaintainEndTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerlessInstance) pulumi.StringOutput { return v.MaintainEndTime }).(pulumi.StringOutput)
}

// The start time of the maintenance window. Specify the time in the `HH:mmZ` format. The time must be in UTC.
func (o ServerlessInstanceOutput) MaintainStartTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerlessInstance) pulumi.StringOutput { return v.MaintainStartTime }).(pulumi.StringOutput)
}

// The purchase duration of the instance, in months. Valid values: `1` to `9`, `12`, `24`, `36`, `60`.
func (o ServerlessInstanceOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServerlessInstance) pulumi.IntPtrOutput { return v.Period }).(pulumi.IntPtrOutput)
}

// The period price type. Valid values: `Day`, `Month`.
func (o ServerlessInstanceOutput) PeriodPriceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerlessInstance) pulumi.StringPtrOutput { return v.PeriodPriceType }).(pulumi.StringPtrOutput)
}

// The ID of the resource group.
func (o ServerlessInstanceOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerlessInstance) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// An array that consists of the information of IP whitelists.
func (o ServerlessInstanceOutput) SecurityIpGroups() ServerlessInstanceSecurityIpGroupArrayOutput {
	return o.ApplyT(func(v *ServerlessInstance) ServerlessInstanceSecurityIpGroupArrayOutput { return v.SecurityIpGroups }).(ServerlessInstanceSecurityIpGroupArrayOutput)
}

// The instance status. For more information, see the instance Status Table.
func (o ServerlessInstanceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerlessInstance) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The storage engine used by the instance. Valid values: `WiredTiger`.
func (o ServerlessInstanceOutput) StorageEngine() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerlessInstance) pulumi.StringOutput { return v.StorageEngine }).(pulumi.StringOutput)
}

// A mapping of tags to assign to the resource.
func (o ServerlessInstanceOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *ServerlessInstance) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// The ID of the VPC network.
func (o ServerlessInstanceOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerlessInstance) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// The of the vswitch.
func (o ServerlessInstanceOutput) VswitchId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerlessInstance) pulumi.StringOutput { return v.VswitchId }).(pulumi.StringOutput)
}

// The ID of the zone. Use this parameter to specify the zone created by the instance.
func (o ServerlessInstanceOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerlessInstance) pulumi.StringOutput { return v.ZoneId }).(pulumi.StringOutput)
}

type ServerlessInstanceArrayOutput struct{ *pulumi.OutputState }

func (ServerlessInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerlessInstance)(nil)).Elem()
}

func (o ServerlessInstanceArrayOutput) ToServerlessInstanceArrayOutput() ServerlessInstanceArrayOutput {
	return o
}

func (o ServerlessInstanceArrayOutput) ToServerlessInstanceArrayOutputWithContext(ctx context.Context) ServerlessInstanceArrayOutput {
	return o
}

func (o ServerlessInstanceArrayOutput) Index(i pulumi.IntInput) ServerlessInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServerlessInstance {
		return vs[0].([]*ServerlessInstance)[vs[1].(int)]
	}).(ServerlessInstanceOutput)
}

type ServerlessInstanceMapOutput struct{ *pulumi.OutputState }

func (ServerlessInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerlessInstance)(nil)).Elem()
}

func (o ServerlessInstanceMapOutput) ToServerlessInstanceMapOutput() ServerlessInstanceMapOutput {
	return o
}

func (o ServerlessInstanceMapOutput) ToServerlessInstanceMapOutputWithContext(ctx context.Context) ServerlessInstanceMapOutput {
	return o
}

func (o ServerlessInstanceMapOutput) MapIndex(k pulumi.StringInput) ServerlessInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServerlessInstance {
		return vs[0].(map[string]*ServerlessInstance)[vs[1].(string)]
	}).(ServerlessInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServerlessInstanceInput)(nil)).Elem(), &ServerlessInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerlessInstanceArrayInput)(nil)).Elem(), ServerlessInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerlessInstanceMapInput)(nil)).Elem(), ServerlessInstanceMap{})
	pulumi.RegisterOutputType(ServerlessInstanceOutput{})
	pulumi.RegisterOutputType(ServerlessInstanceArrayOutput{})
	pulumi.RegisterOutputType(ServerlessInstanceMapOutput{})
}

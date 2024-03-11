// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package adb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a AnalyticDB for MySQL (ADB) DB Cluster Lake Version resource.
//
// For information about AnalyticDB for MySQL (ADB) DB Cluster Lake Version and how to use it, see [What is DB Cluster Lake Version](https://www.alibabacloud.com/help/en/analyticdb-for-mysql/developer-reference/api-adb-2021-12-01-createdbcluster).
//
// > **NOTE:** Available since v1.190.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/adb"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaultZones, err := adb.GetZones(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			defaultNetworks, err := vpc.GetNetworks(ctx, &vpc.GetNetworksArgs{
//				NameRegex: pulumi.StringRef("^default-NODELETING$"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultSwitches, err := vpc.GetSwitches(ctx, &vpc.GetSwitchesArgs{
//				VpcId:  pulumi.StringRef(defaultNetworks.Ids[0]),
//				ZoneId: pulumi.StringRef(defaultZones.Ids[0]),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = adb.NewDBClusterLakeVersion(ctx, "defaultDBClusterLakeVersion", &adb.DBClusterLakeVersionArgs{
//				DbClusterVersion:           pulumi.String("5.0"),
//				VpcId:                      *pulumi.String(defaultNetworks.Ids[0]),
//				VswitchId:                  *pulumi.String(defaultSwitches.Ids[0]),
//				ZoneId:                     *pulumi.String(defaultZones.Ids[0]),
//				ComputeResource:            pulumi.String("16ACU"),
//				StorageResource:            pulumi.String("0ACU"),
//				PaymentType:                pulumi.String("PayAsYouGo"),
//				EnableDefaultResourceGroup: pulumi.Bool(false),
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
// AnalyticDB for MySQL (ADB) DB Cluster Lake Version can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:adb/dBClusterLakeVersion:DBClusterLakeVersion example <id>
// ```
type DBClusterLakeVersion struct {
	pulumi.CustomResourceState

	// The ID of the backup set that you want to use to restore data.
	BackupSetId pulumi.StringPtrOutput `pulumi:"backupSetId"`
	// The name of the service.
	CommodityCode pulumi.StringOutput `pulumi:"commodityCode"`
	// The computing resources of the cluster.
	ComputeResource pulumi.StringOutput `pulumi:"computeResource"`
	// The endpoint of the cluster.
	ConnectionString pulumi.StringOutput `pulumi:"connectionString"`
	// The createTime of the cluster.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The description of the cluster.
	DbClusterDescription pulumi.StringOutput `pulumi:"dbClusterDescription"`
	// The version of the cluster. Valid values: `5.0`.
	DbClusterVersion pulumi.StringOutput `pulumi:"dbClusterVersion"`
	// Whether to enable default allocation of resources to userDefault resource groups.
	EnableDefaultResourceGroup pulumi.BoolPtrOutput `pulumi:"enableDefaultResourceGroup"`
	// The engine of the database.
	Engine pulumi.StringOutput `pulumi:"engine"`
	// The engine version of the database.
	EngineVersion pulumi.StringOutput `pulumi:"engineVersion"`
	// The time when the cluster expires.
	ExpireTime pulumi.StringOutput `pulumi:"expireTime"`
	// Indicates whether the cluster has expired.
	Expired pulumi.StringOutput `pulumi:"expired"`
	// The lock mode of the cluster.
	LockMode pulumi.StringOutput `pulumi:"lockMode"`
	// The reason why the cluster is locked.
	LockReason pulumi.StringOutput `pulumi:"lockReason"`
	// The payment type of the resource. Valid values: `PayAsYouGo`.
	PaymentType pulumi.StringOutput `pulumi:"paymentType"`
	// The port that is used to access the cluster.
	Port pulumi.StringOutput `pulumi:"port"`
	// The ID of the resource group.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// The point in time to which you want to restore data from the backup set.
	RestoreToTime pulumi.StringPtrOutput `pulumi:"restoreToTime"`
	// The method that you want to use to restore data. Valid values:
	RestoreType pulumi.StringPtrOutput `pulumi:"restoreType"`
	// The IP addresses in an IP address whitelist of a cluster. Separate multiple IP addresses with commas (,). You can add a maximum of 500 different IP addresses to a whitelist. The entries in the IP address whitelist must be in one of the following formats:
	// - IP addresses, such as 10.23.XX.XX.
	// - CIDR blocks, such as 10.23.xx.xx/24. In this example, 24 indicates that the prefix of each IP address in the IP whitelist is 24 bits in length. You can replace 24 with a value within the range of 1 to 32.
	SecurityIps pulumi.StringOutput `pulumi:"securityIps"`
	// The ID of the source AnalyticDB for MySQL Data Warehouse Edition cluster.
	SourceDbClusterId pulumi.StringPtrOutput `pulumi:"sourceDbClusterId"`
	// The status of the resource.
	Status pulumi.StringOutput `pulumi:"status"`
	// The storage resources of the cluster.
	StorageResource pulumi.StringOutput `pulumi:"storageResource"`
	// The vpc ID of the resource.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// The ID of the vSwitch.
	VswitchId pulumi.StringOutput `pulumi:"vswitchId"`
	// The zone ID of the resource.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewDBClusterLakeVersion registers a new resource with the given unique name, arguments, and options.
func NewDBClusterLakeVersion(ctx *pulumi.Context,
	name string, args *DBClusterLakeVersionArgs, opts ...pulumi.ResourceOption) (*DBClusterLakeVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ComputeResource == nil {
		return nil, errors.New("invalid value for required argument 'ComputeResource'")
	}
	if args.DbClusterVersion == nil {
		return nil, errors.New("invalid value for required argument 'DbClusterVersion'")
	}
	if args.PaymentType == nil {
		return nil, errors.New("invalid value for required argument 'PaymentType'")
	}
	if args.StorageResource == nil {
		return nil, errors.New("invalid value for required argument 'StorageResource'")
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
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DBClusterLakeVersion
	err := ctx.RegisterResource("alicloud:adb/dBClusterLakeVersion:DBClusterLakeVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDBClusterLakeVersion gets an existing DBClusterLakeVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDBClusterLakeVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DBClusterLakeVersionState, opts ...pulumi.ResourceOption) (*DBClusterLakeVersion, error) {
	var resource DBClusterLakeVersion
	err := ctx.ReadResource("alicloud:adb/dBClusterLakeVersion:DBClusterLakeVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DBClusterLakeVersion resources.
type dbclusterLakeVersionState struct {
	// The ID of the backup set that you want to use to restore data.
	BackupSetId *string `pulumi:"backupSetId"`
	// The name of the service.
	CommodityCode *string `pulumi:"commodityCode"`
	// The computing resources of the cluster.
	ComputeResource *string `pulumi:"computeResource"`
	// The endpoint of the cluster.
	ConnectionString *string `pulumi:"connectionString"`
	// The createTime of the cluster.
	CreateTime *string `pulumi:"createTime"`
	// The description of the cluster.
	DbClusterDescription *string `pulumi:"dbClusterDescription"`
	// The version of the cluster. Valid values: `5.0`.
	DbClusterVersion *string `pulumi:"dbClusterVersion"`
	// Whether to enable default allocation of resources to userDefault resource groups.
	EnableDefaultResourceGroup *bool `pulumi:"enableDefaultResourceGroup"`
	// The engine of the database.
	Engine *string `pulumi:"engine"`
	// The engine version of the database.
	EngineVersion *string `pulumi:"engineVersion"`
	// The time when the cluster expires.
	ExpireTime *string `pulumi:"expireTime"`
	// Indicates whether the cluster has expired.
	Expired *string `pulumi:"expired"`
	// The lock mode of the cluster.
	LockMode *string `pulumi:"lockMode"`
	// The reason why the cluster is locked.
	LockReason *string `pulumi:"lockReason"`
	// The payment type of the resource. Valid values: `PayAsYouGo`.
	PaymentType *string `pulumi:"paymentType"`
	// The port that is used to access the cluster.
	Port *string `pulumi:"port"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The point in time to which you want to restore data from the backup set.
	RestoreToTime *string `pulumi:"restoreToTime"`
	// The method that you want to use to restore data. Valid values:
	RestoreType *string `pulumi:"restoreType"`
	// The IP addresses in an IP address whitelist of a cluster. Separate multiple IP addresses with commas (,). You can add a maximum of 500 different IP addresses to a whitelist. The entries in the IP address whitelist must be in one of the following formats:
	// - IP addresses, such as 10.23.XX.XX.
	// - CIDR blocks, such as 10.23.xx.xx/24. In this example, 24 indicates that the prefix of each IP address in the IP whitelist is 24 bits in length. You can replace 24 with a value within the range of 1 to 32.
	SecurityIps *string `pulumi:"securityIps"`
	// The ID of the source AnalyticDB for MySQL Data Warehouse Edition cluster.
	SourceDbClusterId *string `pulumi:"sourceDbClusterId"`
	// The status of the resource.
	Status *string `pulumi:"status"`
	// The storage resources of the cluster.
	StorageResource *string `pulumi:"storageResource"`
	// The vpc ID of the resource.
	VpcId *string `pulumi:"vpcId"`
	// The ID of the vSwitch.
	VswitchId *string `pulumi:"vswitchId"`
	// The zone ID of the resource.
	ZoneId *string `pulumi:"zoneId"`
}

type DBClusterLakeVersionState struct {
	// The ID of the backup set that you want to use to restore data.
	BackupSetId pulumi.StringPtrInput
	// The name of the service.
	CommodityCode pulumi.StringPtrInput
	// The computing resources of the cluster.
	ComputeResource pulumi.StringPtrInput
	// The endpoint of the cluster.
	ConnectionString pulumi.StringPtrInput
	// The createTime of the cluster.
	CreateTime pulumi.StringPtrInput
	// The description of the cluster.
	DbClusterDescription pulumi.StringPtrInput
	// The version of the cluster. Valid values: `5.0`.
	DbClusterVersion pulumi.StringPtrInput
	// Whether to enable default allocation of resources to userDefault resource groups.
	EnableDefaultResourceGroup pulumi.BoolPtrInput
	// The engine of the database.
	Engine pulumi.StringPtrInput
	// The engine version of the database.
	EngineVersion pulumi.StringPtrInput
	// The time when the cluster expires.
	ExpireTime pulumi.StringPtrInput
	// Indicates whether the cluster has expired.
	Expired pulumi.StringPtrInput
	// The lock mode of the cluster.
	LockMode pulumi.StringPtrInput
	// The reason why the cluster is locked.
	LockReason pulumi.StringPtrInput
	// The payment type of the resource. Valid values: `PayAsYouGo`.
	PaymentType pulumi.StringPtrInput
	// The port that is used to access the cluster.
	Port pulumi.StringPtrInput
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	// The point in time to which you want to restore data from the backup set.
	RestoreToTime pulumi.StringPtrInput
	// The method that you want to use to restore data. Valid values:
	RestoreType pulumi.StringPtrInput
	// The IP addresses in an IP address whitelist of a cluster. Separate multiple IP addresses with commas (,). You can add a maximum of 500 different IP addresses to a whitelist. The entries in the IP address whitelist must be in one of the following formats:
	// - IP addresses, such as 10.23.XX.XX.
	// - CIDR blocks, such as 10.23.xx.xx/24. In this example, 24 indicates that the prefix of each IP address in the IP whitelist is 24 bits in length. You can replace 24 with a value within the range of 1 to 32.
	SecurityIps pulumi.StringPtrInput
	// The ID of the source AnalyticDB for MySQL Data Warehouse Edition cluster.
	SourceDbClusterId pulumi.StringPtrInput
	// The status of the resource.
	Status pulumi.StringPtrInput
	// The storage resources of the cluster.
	StorageResource pulumi.StringPtrInput
	// The vpc ID of the resource.
	VpcId pulumi.StringPtrInput
	// The ID of the vSwitch.
	VswitchId pulumi.StringPtrInput
	// The zone ID of the resource.
	ZoneId pulumi.StringPtrInput
}

func (DBClusterLakeVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*dbclusterLakeVersionState)(nil)).Elem()
}

type dbclusterLakeVersionArgs struct {
	// The ID of the backup set that you want to use to restore data.
	BackupSetId *string `pulumi:"backupSetId"`
	// The computing resources of the cluster.
	ComputeResource string `pulumi:"computeResource"`
	// The description of the cluster.
	DbClusterDescription *string `pulumi:"dbClusterDescription"`
	// The version of the cluster. Valid values: `5.0`.
	DbClusterVersion string `pulumi:"dbClusterVersion"`
	// Whether to enable default allocation of resources to userDefault resource groups.
	EnableDefaultResourceGroup *bool `pulumi:"enableDefaultResourceGroup"`
	// The payment type of the resource. Valid values: `PayAsYouGo`.
	PaymentType string `pulumi:"paymentType"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The point in time to which you want to restore data from the backup set.
	RestoreToTime *string `pulumi:"restoreToTime"`
	// The method that you want to use to restore data. Valid values:
	RestoreType *string `pulumi:"restoreType"`
	// The IP addresses in an IP address whitelist of a cluster. Separate multiple IP addresses with commas (,). You can add a maximum of 500 different IP addresses to a whitelist. The entries in the IP address whitelist must be in one of the following formats:
	// - IP addresses, such as 10.23.XX.XX.
	// - CIDR blocks, such as 10.23.xx.xx/24. In this example, 24 indicates that the prefix of each IP address in the IP whitelist is 24 bits in length. You can replace 24 with a value within the range of 1 to 32.
	SecurityIps *string `pulumi:"securityIps"`
	// The ID of the source AnalyticDB for MySQL Data Warehouse Edition cluster.
	SourceDbClusterId *string `pulumi:"sourceDbClusterId"`
	// The storage resources of the cluster.
	StorageResource string `pulumi:"storageResource"`
	// The vpc ID of the resource.
	VpcId string `pulumi:"vpcId"`
	// The ID of the vSwitch.
	VswitchId string `pulumi:"vswitchId"`
	// The zone ID of the resource.
	ZoneId string `pulumi:"zoneId"`
}

// The set of arguments for constructing a DBClusterLakeVersion resource.
type DBClusterLakeVersionArgs struct {
	// The ID of the backup set that you want to use to restore data.
	BackupSetId pulumi.StringPtrInput
	// The computing resources of the cluster.
	ComputeResource pulumi.StringInput
	// The description of the cluster.
	DbClusterDescription pulumi.StringPtrInput
	// The version of the cluster. Valid values: `5.0`.
	DbClusterVersion pulumi.StringInput
	// Whether to enable default allocation of resources to userDefault resource groups.
	EnableDefaultResourceGroup pulumi.BoolPtrInput
	// The payment type of the resource. Valid values: `PayAsYouGo`.
	PaymentType pulumi.StringInput
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	// The point in time to which you want to restore data from the backup set.
	RestoreToTime pulumi.StringPtrInput
	// The method that you want to use to restore data. Valid values:
	RestoreType pulumi.StringPtrInput
	// The IP addresses in an IP address whitelist of a cluster. Separate multiple IP addresses with commas (,). You can add a maximum of 500 different IP addresses to a whitelist. The entries in the IP address whitelist must be in one of the following formats:
	// - IP addresses, such as 10.23.XX.XX.
	// - CIDR blocks, such as 10.23.xx.xx/24. In this example, 24 indicates that the prefix of each IP address in the IP whitelist is 24 bits in length. You can replace 24 with a value within the range of 1 to 32.
	SecurityIps pulumi.StringPtrInput
	// The ID of the source AnalyticDB for MySQL Data Warehouse Edition cluster.
	SourceDbClusterId pulumi.StringPtrInput
	// The storage resources of the cluster.
	StorageResource pulumi.StringInput
	// The vpc ID of the resource.
	VpcId pulumi.StringInput
	// The ID of the vSwitch.
	VswitchId pulumi.StringInput
	// The zone ID of the resource.
	ZoneId pulumi.StringInput
}

func (DBClusterLakeVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dbclusterLakeVersionArgs)(nil)).Elem()
}

type DBClusterLakeVersionInput interface {
	pulumi.Input

	ToDBClusterLakeVersionOutput() DBClusterLakeVersionOutput
	ToDBClusterLakeVersionOutputWithContext(ctx context.Context) DBClusterLakeVersionOutput
}

func (*DBClusterLakeVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**DBClusterLakeVersion)(nil)).Elem()
}

func (i *DBClusterLakeVersion) ToDBClusterLakeVersionOutput() DBClusterLakeVersionOutput {
	return i.ToDBClusterLakeVersionOutputWithContext(context.Background())
}

func (i *DBClusterLakeVersion) ToDBClusterLakeVersionOutputWithContext(ctx context.Context) DBClusterLakeVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DBClusterLakeVersionOutput)
}

// DBClusterLakeVersionArrayInput is an input type that accepts DBClusterLakeVersionArray and DBClusterLakeVersionArrayOutput values.
// You can construct a concrete instance of `DBClusterLakeVersionArrayInput` via:
//
//	DBClusterLakeVersionArray{ DBClusterLakeVersionArgs{...} }
type DBClusterLakeVersionArrayInput interface {
	pulumi.Input

	ToDBClusterLakeVersionArrayOutput() DBClusterLakeVersionArrayOutput
	ToDBClusterLakeVersionArrayOutputWithContext(context.Context) DBClusterLakeVersionArrayOutput
}

type DBClusterLakeVersionArray []DBClusterLakeVersionInput

func (DBClusterLakeVersionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DBClusterLakeVersion)(nil)).Elem()
}

func (i DBClusterLakeVersionArray) ToDBClusterLakeVersionArrayOutput() DBClusterLakeVersionArrayOutput {
	return i.ToDBClusterLakeVersionArrayOutputWithContext(context.Background())
}

func (i DBClusterLakeVersionArray) ToDBClusterLakeVersionArrayOutputWithContext(ctx context.Context) DBClusterLakeVersionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DBClusterLakeVersionArrayOutput)
}

// DBClusterLakeVersionMapInput is an input type that accepts DBClusterLakeVersionMap and DBClusterLakeVersionMapOutput values.
// You can construct a concrete instance of `DBClusterLakeVersionMapInput` via:
//
//	DBClusterLakeVersionMap{ "key": DBClusterLakeVersionArgs{...} }
type DBClusterLakeVersionMapInput interface {
	pulumi.Input

	ToDBClusterLakeVersionMapOutput() DBClusterLakeVersionMapOutput
	ToDBClusterLakeVersionMapOutputWithContext(context.Context) DBClusterLakeVersionMapOutput
}

type DBClusterLakeVersionMap map[string]DBClusterLakeVersionInput

func (DBClusterLakeVersionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DBClusterLakeVersion)(nil)).Elem()
}

func (i DBClusterLakeVersionMap) ToDBClusterLakeVersionMapOutput() DBClusterLakeVersionMapOutput {
	return i.ToDBClusterLakeVersionMapOutputWithContext(context.Background())
}

func (i DBClusterLakeVersionMap) ToDBClusterLakeVersionMapOutputWithContext(ctx context.Context) DBClusterLakeVersionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DBClusterLakeVersionMapOutput)
}

type DBClusterLakeVersionOutput struct{ *pulumi.OutputState }

func (DBClusterLakeVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DBClusterLakeVersion)(nil)).Elem()
}

func (o DBClusterLakeVersionOutput) ToDBClusterLakeVersionOutput() DBClusterLakeVersionOutput {
	return o
}

func (o DBClusterLakeVersionOutput) ToDBClusterLakeVersionOutputWithContext(ctx context.Context) DBClusterLakeVersionOutput {
	return o
}

// The ID of the backup set that you want to use to restore data.
func (o DBClusterLakeVersionOutput) BackupSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DBClusterLakeVersion) pulumi.StringPtrOutput { return v.BackupSetId }).(pulumi.StringPtrOutput)
}

// The name of the service.
func (o DBClusterLakeVersionOutput) CommodityCode() pulumi.StringOutput {
	return o.ApplyT(func(v *DBClusterLakeVersion) pulumi.StringOutput { return v.CommodityCode }).(pulumi.StringOutput)
}

// The computing resources of the cluster.
func (o DBClusterLakeVersionOutput) ComputeResource() pulumi.StringOutput {
	return o.ApplyT(func(v *DBClusterLakeVersion) pulumi.StringOutput { return v.ComputeResource }).(pulumi.StringOutput)
}

// The endpoint of the cluster.
func (o DBClusterLakeVersionOutput) ConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v *DBClusterLakeVersion) pulumi.StringOutput { return v.ConnectionString }).(pulumi.StringOutput)
}

// The createTime of the cluster.
func (o DBClusterLakeVersionOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DBClusterLakeVersion) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The description of the cluster.
func (o DBClusterLakeVersionOutput) DbClusterDescription() pulumi.StringOutput {
	return o.ApplyT(func(v *DBClusterLakeVersion) pulumi.StringOutput { return v.DbClusterDescription }).(pulumi.StringOutput)
}

// The version of the cluster. Valid values: `5.0`.
func (o DBClusterLakeVersionOutput) DbClusterVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *DBClusterLakeVersion) pulumi.StringOutput { return v.DbClusterVersion }).(pulumi.StringOutput)
}

// Whether to enable default allocation of resources to userDefault resource groups.
func (o DBClusterLakeVersionOutput) EnableDefaultResourceGroup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DBClusterLakeVersion) pulumi.BoolPtrOutput { return v.EnableDefaultResourceGroup }).(pulumi.BoolPtrOutput)
}

// The engine of the database.
func (o DBClusterLakeVersionOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v *DBClusterLakeVersion) pulumi.StringOutput { return v.Engine }).(pulumi.StringOutput)
}

// The engine version of the database.
func (o DBClusterLakeVersionOutput) EngineVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *DBClusterLakeVersion) pulumi.StringOutput { return v.EngineVersion }).(pulumi.StringOutput)
}

// The time when the cluster expires.
func (o DBClusterLakeVersionOutput) ExpireTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DBClusterLakeVersion) pulumi.StringOutput { return v.ExpireTime }).(pulumi.StringOutput)
}

// Indicates whether the cluster has expired.
func (o DBClusterLakeVersionOutput) Expired() pulumi.StringOutput {
	return o.ApplyT(func(v *DBClusterLakeVersion) pulumi.StringOutput { return v.Expired }).(pulumi.StringOutput)
}

// The lock mode of the cluster.
func (o DBClusterLakeVersionOutput) LockMode() pulumi.StringOutput {
	return o.ApplyT(func(v *DBClusterLakeVersion) pulumi.StringOutput { return v.LockMode }).(pulumi.StringOutput)
}

// The reason why the cluster is locked.
func (o DBClusterLakeVersionOutput) LockReason() pulumi.StringOutput {
	return o.ApplyT(func(v *DBClusterLakeVersion) pulumi.StringOutput { return v.LockReason }).(pulumi.StringOutput)
}

// The payment type of the resource. Valid values: `PayAsYouGo`.
func (o DBClusterLakeVersionOutput) PaymentType() pulumi.StringOutput {
	return o.ApplyT(func(v *DBClusterLakeVersion) pulumi.StringOutput { return v.PaymentType }).(pulumi.StringOutput)
}

// The port that is used to access the cluster.
func (o DBClusterLakeVersionOutput) Port() pulumi.StringOutput {
	return o.ApplyT(func(v *DBClusterLakeVersion) pulumi.StringOutput { return v.Port }).(pulumi.StringOutput)
}

// The ID of the resource group.
func (o DBClusterLakeVersionOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *DBClusterLakeVersion) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// The point in time to which you want to restore data from the backup set.
func (o DBClusterLakeVersionOutput) RestoreToTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DBClusterLakeVersion) pulumi.StringPtrOutput { return v.RestoreToTime }).(pulumi.StringPtrOutput)
}

// The method that you want to use to restore data. Valid values:
func (o DBClusterLakeVersionOutput) RestoreType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DBClusterLakeVersion) pulumi.StringPtrOutput { return v.RestoreType }).(pulumi.StringPtrOutput)
}

// The IP addresses in an IP address whitelist of a cluster. Separate multiple IP addresses with commas (,). You can add a maximum of 500 different IP addresses to a whitelist. The entries in the IP address whitelist must be in one of the following formats:
// - IP addresses, such as 10.23.XX.XX.
// - CIDR blocks, such as 10.23.xx.xx/24. In this example, 24 indicates that the prefix of each IP address in the IP whitelist is 24 bits in length. You can replace 24 with a value within the range of 1 to 32.
func (o DBClusterLakeVersionOutput) SecurityIps() pulumi.StringOutput {
	return o.ApplyT(func(v *DBClusterLakeVersion) pulumi.StringOutput { return v.SecurityIps }).(pulumi.StringOutput)
}

// The ID of the source AnalyticDB for MySQL Data Warehouse Edition cluster.
func (o DBClusterLakeVersionOutput) SourceDbClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DBClusterLakeVersion) pulumi.StringPtrOutput { return v.SourceDbClusterId }).(pulumi.StringPtrOutput)
}

// The status of the resource.
func (o DBClusterLakeVersionOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *DBClusterLakeVersion) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The storage resources of the cluster.
func (o DBClusterLakeVersionOutput) StorageResource() pulumi.StringOutput {
	return o.ApplyT(func(v *DBClusterLakeVersion) pulumi.StringOutput { return v.StorageResource }).(pulumi.StringOutput)
}

// The vpc ID of the resource.
func (o DBClusterLakeVersionOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *DBClusterLakeVersion) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// The ID of the vSwitch.
func (o DBClusterLakeVersionOutput) VswitchId() pulumi.StringOutput {
	return o.ApplyT(func(v *DBClusterLakeVersion) pulumi.StringOutput { return v.VswitchId }).(pulumi.StringOutput)
}

// The zone ID of the resource.
func (o DBClusterLakeVersionOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *DBClusterLakeVersion) pulumi.StringOutput { return v.ZoneId }).(pulumi.StringOutput)
}

type DBClusterLakeVersionArrayOutput struct{ *pulumi.OutputState }

func (DBClusterLakeVersionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DBClusterLakeVersion)(nil)).Elem()
}

func (o DBClusterLakeVersionArrayOutput) ToDBClusterLakeVersionArrayOutput() DBClusterLakeVersionArrayOutput {
	return o
}

func (o DBClusterLakeVersionArrayOutput) ToDBClusterLakeVersionArrayOutputWithContext(ctx context.Context) DBClusterLakeVersionArrayOutput {
	return o
}

func (o DBClusterLakeVersionArrayOutput) Index(i pulumi.IntInput) DBClusterLakeVersionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DBClusterLakeVersion {
		return vs[0].([]*DBClusterLakeVersion)[vs[1].(int)]
	}).(DBClusterLakeVersionOutput)
}

type DBClusterLakeVersionMapOutput struct{ *pulumi.OutputState }

func (DBClusterLakeVersionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DBClusterLakeVersion)(nil)).Elem()
}

func (o DBClusterLakeVersionMapOutput) ToDBClusterLakeVersionMapOutput() DBClusterLakeVersionMapOutput {
	return o
}

func (o DBClusterLakeVersionMapOutput) ToDBClusterLakeVersionMapOutputWithContext(ctx context.Context) DBClusterLakeVersionMapOutput {
	return o
}

func (o DBClusterLakeVersionMapOutput) MapIndex(k pulumi.StringInput) DBClusterLakeVersionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DBClusterLakeVersion {
		return vs[0].(map[string]*DBClusterLakeVersion)[vs[1].(string)]
	}).(DBClusterLakeVersionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DBClusterLakeVersionInput)(nil)).Elem(), &DBClusterLakeVersion{})
	pulumi.RegisterInputType(reflect.TypeOf((*DBClusterLakeVersionArrayInput)(nil)).Elem(), DBClusterLakeVersionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DBClusterLakeVersionMapInput)(nil)).Elem(), DBClusterLakeVersionMap{})
	pulumi.RegisterOutputType(DBClusterLakeVersionOutput{})
	pulumi.RegisterOutputType(DBClusterLakeVersionArrayOutput{})
	pulumi.RegisterOutputType(DBClusterLakeVersionMapOutput{})
}

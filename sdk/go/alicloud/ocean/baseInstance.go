// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ocean

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Ocean Base Instance resource.
//
// For information about Ocean Base Instance and how to use it, see [What is Instance](https://www.alibabacloud.com/help/en/apsaradb-for-oceanbase/latest/what-is-oceanbase-database).
//
// > **NOTE:** Available since v1.203.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ocean"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/resourcemanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "terraform-example"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			_default, err := alicloud.GetZones(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			defaultGetResourceGroups, err := resourcemanager.GetResourceGroups(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = ocean.NewBaseInstance(ctx, "default", &ocean.BaseInstanceArgs{
//				ResourceGroupId: pulumi.String(defaultGetResourceGroups.Ids[0]),
//				Zones: pulumi.StringArray{
//					_default.Ids[len(_default.Ids)-2],
//					_default.Ids[len(_default.Ids)-3],
//					_default.Ids[len(_default.Ids)-4],
//				},
//				AutoRenew:        pulumi.Bool(false),
//				DiskSize:         pulumi.Int(100),
//				PaymentType:      pulumi.String("PayAsYouGo"),
//				InstanceClass:    pulumi.String("8C32GB"),
//				BackupRetainMode: pulumi.String("delete_all"),
//				Series:           pulumi.String("normal"),
//				InstanceName:     pulumi.String(name),
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
// Ocean Base Instance can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:ocean/baseInstance:BaseInstance example <id>
// ```
type BaseInstance struct {
	pulumi.CustomResourceState

	// Whether to automatically renew.
	// It takes effect when the parameter ChargeType is PrePaid. Value range:
	// - true: automatic renewal.
	// - false (default): no automatic renewal.
	AutoRenew pulumi.BoolPtrOutput `pulumi:"autoRenew"`
	// The duration of each auto-renewal. When the value of the AutoRenew parameter is True, this parameter is required.
	// - PeriodUnit is Week, AutoRenewPeriod is {"1", "2", "3"}.
	// - PeriodUnit is Month, AutoRenewPeriod is {"1", "2", "3", "6", "12"}.
	AutoRenewPeriod pulumi.IntPtrOutput `pulumi:"autoRenewPeriod"`
	// The backup retention policy after the cluster is deleted. The values are as follows:
	// - receive_all: Keep all backup sets;
	// - delete_all: delete all backup sets;
	// - receive_last: Keep the last backup set.
	// > **NOTE:**   The default value is delete_all.
	BackupRetainMode pulumi.StringPtrOutput `pulumi:"backupRetainMode"`
	// The product code of the OceanBase cluster._oceanbasepre_public_cn: Domestic station cloud database package Year-to-month package._oceanbasepost_public_cn: The domestic station cloud database is paid by the hour._obpre_public_intl: International Station Cloud Database Package Monthly Package.
	CommodityCode pulumi.StringOutput `pulumi:"commodityCode"`
	// The number of CPU cores of the cluster.
	Cpu pulumi.IntOutput `pulumi:"cpu"`
	// The creation time of the resource.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The size of the storage space, in GB.
	// The limits of storage space vary according to the cluster specifications, as follows:
	// - 8C32GB:100GB ~ 10000GB
	// - 14C70GB:200GB ~ 10000GB
	// - 30C180GB:400GB ~ 10000GB
	// - 62C400G:800GB ~ 10000GB.
	//   The default value of each package is its minimum value.
	DiskSize pulumi.IntOutput `pulumi:"diskSize"`
	// The storage type of the cluster. Effective only in the standard cluster version (cloud disk).
	// Two types are currently supported:
	// - cloud_essd_pl1: cloud disk ESSD pl1.
	// - cloud_essd_pl0: cloud disk ESSD pl0. The default value is cloud_essd_pl1.
	DiskType pulumi.StringOutput `pulumi:"diskType"`
	// Cluster specification information.
	// Four packages are currently supported:
	// - 4C16GB：4cores 16GB
	// - 8C32GB：8cores 32GB
	// - 14C70GB：14cores 70GB
	// - 24C120GB：24cores 120GB
	// - 30C180GB：30cores 180GB
	// - 62C400GB：62cores 400GB
	// - 104C600GB：104cores 600GB
	// - 16C70GB：16cores 70GB
	// - 32C160GB：32cores 160GB
	// - 64C380GB：64cores 380GB
	// - 20C32GB：20cores 32GB
	// - 40C64GB：40cores 64GB
	// - 16C32GB：16cores 32GB
	// - 32C70GB：32cores 70GB
	// - 64C180GB：64cores 180GB
	// - 32C180GB：32cores 180GB
	// - 64C400GB：64cores 400GB.
	InstanceClass pulumi.StringOutput `pulumi:"instanceClass"`
	// OceanBase cluster name.The length is 1 to 20 English or Chinese characters.If this parameter is not specified, the default value is the InstanceId of the cluster.
	InstanceName pulumi.StringOutput `pulumi:"instanceName"`
	// The number of nodes in the cluster. If the deployment mode is n-n-n, the number of nodes is n * 3.
	NodeNum pulumi.StringOutput `pulumi:"nodeNum"`
	// The OceanBase Server version number.
	ObVersion pulumi.StringOutput `pulumi:"obVersion"`
	// The payment method of the instance. Value range:
	// - Subscription: Package year and month. When you select this type of payment method, you must make sure that your account supports balance payment or credit payment. Otherwise, an InvalidPayMethod error message will be returned.
	// - PayAsYouGo (default): Pay-as-you-go (default hourly billing).
	PaymentType pulumi.StringOutput `pulumi:"paymentType"`
	// The duration of the resource purchase. The unit is specified by the PeriodUnit. The parameter InstanceChargeType takes effect only when the value is PrePaid and is required. Once the DedicatedHostId is specified, the value cannot exceed the subscription duration of the dedicated host. When PeriodUnit = Week, Period values: {"1", "2", "3", "4"}. When PeriodUnit = Month, Period values: {"1", "2", "3", "4", "5", "6", "7", "8", "9", "12", "24", "36", "48", "60"}.
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// The duration of the purchase of resources.Package year and Month value range: Month.Default value: Month of the package, which is billed by volume. The default period is Hour.
	PeriodUnit pulumi.StringPtrOutput `pulumi:"periodUnit"`
	// The ID of the enterprise resource group to which the instance resides.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// Series of OceanBase cluster instances-normal (default): Standard cluster version (cloud disk)-normal_SSD: Standard cluster version (local disk)-history: history Library cluster version.
	Series pulumi.StringOutput `pulumi:"series"`
	// The status of the resource.
	Status pulumi.StringOutput `pulumi:"status"`
	// Information about the zone where the cluster is deployed.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewBaseInstance registers a new resource with the given unique name, arguments, and options.
func NewBaseInstance(ctx *pulumi.Context,
	name string, args *BaseInstanceArgs, opts ...pulumi.ResourceOption) (*BaseInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DiskSize == nil {
		return nil, errors.New("invalid value for required argument 'DiskSize'")
	}
	if args.InstanceClass == nil {
		return nil, errors.New("invalid value for required argument 'InstanceClass'")
	}
	if args.PaymentType == nil {
		return nil, errors.New("invalid value for required argument 'PaymentType'")
	}
	if args.Series == nil {
		return nil, errors.New("invalid value for required argument 'Series'")
	}
	if args.Zones == nil {
		return nil, errors.New("invalid value for required argument 'Zones'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BaseInstance
	err := ctx.RegisterResource("alicloud:ocean/baseInstance:BaseInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBaseInstance gets an existing BaseInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBaseInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BaseInstanceState, opts ...pulumi.ResourceOption) (*BaseInstance, error) {
	var resource BaseInstance
	err := ctx.ReadResource("alicloud:ocean/baseInstance:BaseInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BaseInstance resources.
type baseInstanceState struct {
	// Whether to automatically renew.
	// It takes effect when the parameter ChargeType is PrePaid. Value range:
	// - true: automatic renewal.
	// - false (default): no automatic renewal.
	AutoRenew *bool `pulumi:"autoRenew"`
	// The duration of each auto-renewal. When the value of the AutoRenew parameter is True, this parameter is required.
	// - PeriodUnit is Week, AutoRenewPeriod is {"1", "2", "3"}.
	// - PeriodUnit is Month, AutoRenewPeriod is {"1", "2", "3", "6", "12"}.
	AutoRenewPeriod *int `pulumi:"autoRenewPeriod"`
	// The backup retention policy after the cluster is deleted. The values are as follows:
	// - receive_all: Keep all backup sets;
	// - delete_all: delete all backup sets;
	// - receive_last: Keep the last backup set.
	// > **NOTE:**   The default value is delete_all.
	BackupRetainMode *string `pulumi:"backupRetainMode"`
	// The product code of the OceanBase cluster._oceanbasepre_public_cn: Domestic station cloud database package Year-to-month package._oceanbasepost_public_cn: The domestic station cloud database is paid by the hour._obpre_public_intl: International Station Cloud Database Package Monthly Package.
	CommodityCode *string `pulumi:"commodityCode"`
	// The number of CPU cores of the cluster.
	Cpu *int `pulumi:"cpu"`
	// The creation time of the resource.
	CreateTime *string `pulumi:"createTime"`
	// The size of the storage space, in GB.
	// The limits of storage space vary according to the cluster specifications, as follows:
	// - 8C32GB:100GB ~ 10000GB
	// - 14C70GB:200GB ~ 10000GB
	// - 30C180GB:400GB ~ 10000GB
	// - 62C400G:800GB ~ 10000GB.
	//   The default value of each package is its minimum value.
	DiskSize *int `pulumi:"diskSize"`
	// The storage type of the cluster. Effective only in the standard cluster version (cloud disk).
	// Two types are currently supported:
	// - cloud_essd_pl1: cloud disk ESSD pl1.
	// - cloud_essd_pl0: cloud disk ESSD pl0. The default value is cloud_essd_pl1.
	DiskType *string `pulumi:"diskType"`
	// Cluster specification information.
	// Four packages are currently supported:
	// - 4C16GB：4cores 16GB
	// - 8C32GB：8cores 32GB
	// - 14C70GB：14cores 70GB
	// - 24C120GB：24cores 120GB
	// - 30C180GB：30cores 180GB
	// - 62C400GB：62cores 400GB
	// - 104C600GB：104cores 600GB
	// - 16C70GB：16cores 70GB
	// - 32C160GB：32cores 160GB
	// - 64C380GB：64cores 380GB
	// - 20C32GB：20cores 32GB
	// - 40C64GB：40cores 64GB
	// - 16C32GB：16cores 32GB
	// - 32C70GB：32cores 70GB
	// - 64C180GB：64cores 180GB
	// - 32C180GB：32cores 180GB
	// - 64C400GB：64cores 400GB.
	InstanceClass *string `pulumi:"instanceClass"`
	// OceanBase cluster name.The length is 1 to 20 English or Chinese characters.If this parameter is not specified, the default value is the InstanceId of the cluster.
	InstanceName *string `pulumi:"instanceName"`
	// The number of nodes in the cluster. If the deployment mode is n-n-n, the number of nodes is n * 3.
	NodeNum *string `pulumi:"nodeNum"`
	// The OceanBase Server version number.
	ObVersion *string `pulumi:"obVersion"`
	// The payment method of the instance. Value range:
	// - Subscription: Package year and month. When you select this type of payment method, you must make sure that your account supports balance payment or credit payment. Otherwise, an InvalidPayMethod error message will be returned.
	// - PayAsYouGo (default): Pay-as-you-go (default hourly billing).
	PaymentType *string `pulumi:"paymentType"`
	// The duration of the resource purchase. The unit is specified by the PeriodUnit. The parameter InstanceChargeType takes effect only when the value is PrePaid and is required. Once the DedicatedHostId is specified, the value cannot exceed the subscription duration of the dedicated host. When PeriodUnit = Week, Period values: {"1", "2", "3", "4"}. When PeriodUnit = Month, Period values: {"1", "2", "3", "4", "5", "6", "7", "8", "9", "12", "24", "36", "48", "60"}.
	Period *int `pulumi:"period"`
	// The duration of the purchase of resources.Package year and Month value range: Month.Default value: Month of the package, which is billed by volume. The default period is Hour.
	PeriodUnit *string `pulumi:"periodUnit"`
	// The ID of the enterprise resource group to which the instance resides.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// Series of OceanBase cluster instances-normal (default): Standard cluster version (cloud disk)-normal_SSD: Standard cluster version (local disk)-history: history Library cluster version.
	Series *string `pulumi:"series"`
	// The status of the resource.
	Status *string `pulumi:"status"`
	// Information about the zone where the cluster is deployed.
	Zones []string `pulumi:"zones"`
}

type BaseInstanceState struct {
	// Whether to automatically renew.
	// It takes effect when the parameter ChargeType is PrePaid. Value range:
	// - true: automatic renewal.
	// - false (default): no automatic renewal.
	AutoRenew pulumi.BoolPtrInput
	// The duration of each auto-renewal. When the value of the AutoRenew parameter is True, this parameter is required.
	// - PeriodUnit is Week, AutoRenewPeriod is {"1", "2", "3"}.
	// - PeriodUnit is Month, AutoRenewPeriod is {"1", "2", "3", "6", "12"}.
	AutoRenewPeriod pulumi.IntPtrInput
	// The backup retention policy after the cluster is deleted. The values are as follows:
	// - receive_all: Keep all backup sets;
	// - delete_all: delete all backup sets;
	// - receive_last: Keep the last backup set.
	// > **NOTE:**   The default value is delete_all.
	BackupRetainMode pulumi.StringPtrInput
	// The product code of the OceanBase cluster._oceanbasepre_public_cn: Domestic station cloud database package Year-to-month package._oceanbasepost_public_cn: The domestic station cloud database is paid by the hour._obpre_public_intl: International Station Cloud Database Package Monthly Package.
	CommodityCode pulumi.StringPtrInput
	// The number of CPU cores of the cluster.
	Cpu pulumi.IntPtrInput
	// The creation time of the resource.
	CreateTime pulumi.StringPtrInput
	// The size of the storage space, in GB.
	// The limits of storage space vary according to the cluster specifications, as follows:
	// - 8C32GB:100GB ~ 10000GB
	// - 14C70GB:200GB ~ 10000GB
	// - 30C180GB:400GB ~ 10000GB
	// - 62C400G:800GB ~ 10000GB.
	//   The default value of each package is its minimum value.
	DiskSize pulumi.IntPtrInput
	// The storage type of the cluster. Effective only in the standard cluster version (cloud disk).
	// Two types are currently supported:
	// - cloud_essd_pl1: cloud disk ESSD pl1.
	// - cloud_essd_pl0: cloud disk ESSD pl0. The default value is cloud_essd_pl1.
	DiskType pulumi.StringPtrInput
	// Cluster specification information.
	// Four packages are currently supported:
	// - 4C16GB：4cores 16GB
	// - 8C32GB：8cores 32GB
	// - 14C70GB：14cores 70GB
	// - 24C120GB：24cores 120GB
	// - 30C180GB：30cores 180GB
	// - 62C400GB：62cores 400GB
	// - 104C600GB：104cores 600GB
	// - 16C70GB：16cores 70GB
	// - 32C160GB：32cores 160GB
	// - 64C380GB：64cores 380GB
	// - 20C32GB：20cores 32GB
	// - 40C64GB：40cores 64GB
	// - 16C32GB：16cores 32GB
	// - 32C70GB：32cores 70GB
	// - 64C180GB：64cores 180GB
	// - 32C180GB：32cores 180GB
	// - 64C400GB：64cores 400GB.
	InstanceClass pulumi.StringPtrInput
	// OceanBase cluster name.The length is 1 to 20 English or Chinese characters.If this parameter is not specified, the default value is the InstanceId of the cluster.
	InstanceName pulumi.StringPtrInput
	// The number of nodes in the cluster. If the deployment mode is n-n-n, the number of nodes is n * 3.
	NodeNum pulumi.StringPtrInput
	// The OceanBase Server version number.
	ObVersion pulumi.StringPtrInput
	// The payment method of the instance. Value range:
	// - Subscription: Package year and month. When you select this type of payment method, you must make sure that your account supports balance payment or credit payment. Otherwise, an InvalidPayMethod error message will be returned.
	// - PayAsYouGo (default): Pay-as-you-go (default hourly billing).
	PaymentType pulumi.StringPtrInput
	// The duration of the resource purchase. The unit is specified by the PeriodUnit. The parameter InstanceChargeType takes effect only when the value is PrePaid and is required. Once the DedicatedHostId is specified, the value cannot exceed the subscription duration of the dedicated host. When PeriodUnit = Week, Period values: {"1", "2", "3", "4"}. When PeriodUnit = Month, Period values: {"1", "2", "3", "4", "5", "6", "7", "8", "9", "12", "24", "36", "48", "60"}.
	Period pulumi.IntPtrInput
	// The duration of the purchase of resources.Package year and Month value range: Month.Default value: Month of the package, which is billed by volume. The default period is Hour.
	PeriodUnit pulumi.StringPtrInput
	// The ID of the enterprise resource group to which the instance resides.
	ResourceGroupId pulumi.StringPtrInput
	// Series of OceanBase cluster instances-normal (default): Standard cluster version (cloud disk)-normal_SSD: Standard cluster version (local disk)-history: history Library cluster version.
	Series pulumi.StringPtrInput
	// The status of the resource.
	Status pulumi.StringPtrInput
	// Information about the zone where the cluster is deployed.
	Zones pulumi.StringArrayInput
}

func (BaseInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*baseInstanceState)(nil)).Elem()
}

type baseInstanceArgs struct {
	// Whether to automatically renew.
	// It takes effect when the parameter ChargeType is PrePaid. Value range:
	// - true: automatic renewal.
	// - false (default): no automatic renewal.
	AutoRenew *bool `pulumi:"autoRenew"`
	// The duration of each auto-renewal. When the value of the AutoRenew parameter is True, this parameter is required.
	// - PeriodUnit is Week, AutoRenewPeriod is {"1", "2", "3"}.
	// - PeriodUnit is Month, AutoRenewPeriod is {"1", "2", "3", "6", "12"}.
	AutoRenewPeriod *int `pulumi:"autoRenewPeriod"`
	// The backup retention policy after the cluster is deleted. The values are as follows:
	// - receive_all: Keep all backup sets;
	// - delete_all: delete all backup sets;
	// - receive_last: Keep the last backup set.
	// > **NOTE:**   The default value is delete_all.
	BackupRetainMode *string `pulumi:"backupRetainMode"`
	// The size of the storage space, in GB.
	// The limits of storage space vary according to the cluster specifications, as follows:
	// - 8C32GB:100GB ~ 10000GB
	// - 14C70GB:200GB ~ 10000GB
	// - 30C180GB:400GB ~ 10000GB
	// - 62C400G:800GB ~ 10000GB.
	//   The default value of each package is its minimum value.
	DiskSize int `pulumi:"diskSize"`
	// The storage type of the cluster. Effective only in the standard cluster version (cloud disk).
	// Two types are currently supported:
	// - cloud_essd_pl1: cloud disk ESSD pl1.
	// - cloud_essd_pl0: cloud disk ESSD pl0. The default value is cloud_essd_pl1.
	DiskType *string `pulumi:"diskType"`
	// Cluster specification information.
	// Four packages are currently supported:
	// - 4C16GB：4cores 16GB
	// - 8C32GB：8cores 32GB
	// - 14C70GB：14cores 70GB
	// - 24C120GB：24cores 120GB
	// - 30C180GB：30cores 180GB
	// - 62C400GB：62cores 400GB
	// - 104C600GB：104cores 600GB
	// - 16C70GB：16cores 70GB
	// - 32C160GB：32cores 160GB
	// - 64C380GB：64cores 380GB
	// - 20C32GB：20cores 32GB
	// - 40C64GB：40cores 64GB
	// - 16C32GB：16cores 32GB
	// - 32C70GB：32cores 70GB
	// - 64C180GB：64cores 180GB
	// - 32C180GB：32cores 180GB
	// - 64C400GB：64cores 400GB.
	InstanceClass string `pulumi:"instanceClass"`
	// OceanBase cluster name.The length is 1 to 20 English or Chinese characters.If this parameter is not specified, the default value is the InstanceId of the cluster.
	InstanceName *string `pulumi:"instanceName"`
	// The number of nodes in the cluster. If the deployment mode is n-n-n, the number of nodes is n * 3.
	NodeNum *string `pulumi:"nodeNum"`
	// The OceanBase Server version number.
	ObVersion *string `pulumi:"obVersion"`
	// The payment method of the instance. Value range:
	// - Subscription: Package year and month. When you select this type of payment method, you must make sure that your account supports balance payment or credit payment. Otherwise, an InvalidPayMethod error message will be returned.
	// - PayAsYouGo (default): Pay-as-you-go (default hourly billing).
	PaymentType string `pulumi:"paymentType"`
	// The duration of the resource purchase. The unit is specified by the PeriodUnit. The parameter InstanceChargeType takes effect only when the value is PrePaid and is required. Once the DedicatedHostId is specified, the value cannot exceed the subscription duration of the dedicated host. When PeriodUnit = Week, Period values: {"1", "2", "3", "4"}. When PeriodUnit = Month, Period values: {"1", "2", "3", "4", "5", "6", "7", "8", "9", "12", "24", "36", "48", "60"}.
	Period *int `pulumi:"period"`
	// The duration of the purchase of resources.Package year and Month value range: Month.Default value: Month of the package, which is billed by volume. The default period is Hour.
	PeriodUnit *string `pulumi:"periodUnit"`
	// The ID of the enterprise resource group to which the instance resides.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// Series of OceanBase cluster instances-normal (default): Standard cluster version (cloud disk)-normal_SSD: Standard cluster version (local disk)-history: history Library cluster version.
	Series string `pulumi:"series"`
	// Information about the zone where the cluster is deployed.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a BaseInstance resource.
type BaseInstanceArgs struct {
	// Whether to automatically renew.
	// It takes effect when the parameter ChargeType is PrePaid. Value range:
	// - true: automatic renewal.
	// - false (default): no automatic renewal.
	AutoRenew pulumi.BoolPtrInput
	// The duration of each auto-renewal. When the value of the AutoRenew parameter is True, this parameter is required.
	// - PeriodUnit is Week, AutoRenewPeriod is {"1", "2", "3"}.
	// - PeriodUnit is Month, AutoRenewPeriod is {"1", "2", "3", "6", "12"}.
	AutoRenewPeriod pulumi.IntPtrInput
	// The backup retention policy after the cluster is deleted. The values are as follows:
	// - receive_all: Keep all backup sets;
	// - delete_all: delete all backup sets;
	// - receive_last: Keep the last backup set.
	// > **NOTE:**   The default value is delete_all.
	BackupRetainMode pulumi.StringPtrInput
	// The size of the storage space, in GB.
	// The limits of storage space vary according to the cluster specifications, as follows:
	// - 8C32GB:100GB ~ 10000GB
	// - 14C70GB:200GB ~ 10000GB
	// - 30C180GB:400GB ~ 10000GB
	// - 62C400G:800GB ~ 10000GB.
	//   The default value of each package is its minimum value.
	DiskSize pulumi.IntInput
	// The storage type of the cluster. Effective only in the standard cluster version (cloud disk).
	// Two types are currently supported:
	// - cloud_essd_pl1: cloud disk ESSD pl1.
	// - cloud_essd_pl0: cloud disk ESSD pl0. The default value is cloud_essd_pl1.
	DiskType pulumi.StringPtrInput
	// Cluster specification information.
	// Four packages are currently supported:
	// - 4C16GB：4cores 16GB
	// - 8C32GB：8cores 32GB
	// - 14C70GB：14cores 70GB
	// - 24C120GB：24cores 120GB
	// - 30C180GB：30cores 180GB
	// - 62C400GB：62cores 400GB
	// - 104C600GB：104cores 600GB
	// - 16C70GB：16cores 70GB
	// - 32C160GB：32cores 160GB
	// - 64C380GB：64cores 380GB
	// - 20C32GB：20cores 32GB
	// - 40C64GB：40cores 64GB
	// - 16C32GB：16cores 32GB
	// - 32C70GB：32cores 70GB
	// - 64C180GB：64cores 180GB
	// - 32C180GB：32cores 180GB
	// - 64C400GB：64cores 400GB.
	InstanceClass pulumi.StringInput
	// OceanBase cluster name.The length is 1 to 20 English or Chinese characters.If this parameter is not specified, the default value is the InstanceId of the cluster.
	InstanceName pulumi.StringPtrInput
	// The number of nodes in the cluster. If the deployment mode is n-n-n, the number of nodes is n * 3.
	NodeNum pulumi.StringPtrInput
	// The OceanBase Server version number.
	ObVersion pulumi.StringPtrInput
	// The payment method of the instance. Value range:
	// - Subscription: Package year and month. When you select this type of payment method, you must make sure that your account supports balance payment or credit payment. Otherwise, an InvalidPayMethod error message will be returned.
	// - PayAsYouGo (default): Pay-as-you-go (default hourly billing).
	PaymentType pulumi.StringInput
	// The duration of the resource purchase. The unit is specified by the PeriodUnit. The parameter InstanceChargeType takes effect only when the value is PrePaid and is required. Once the DedicatedHostId is specified, the value cannot exceed the subscription duration of the dedicated host. When PeriodUnit = Week, Period values: {"1", "2", "3", "4"}. When PeriodUnit = Month, Period values: {"1", "2", "3", "4", "5", "6", "7", "8", "9", "12", "24", "36", "48", "60"}.
	Period pulumi.IntPtrInput
	// The duration of the purchase of resources.Package year and Month value range: Month.Default value: Month of the package, which is billed by volume. The default period is Hour.
	PeriodUnit pulumi.StringPtrInput
	// The ID of the enterprise resource group to which the instance resides.
	ResourceGroupId pulumi.StringPtrInput
	// Series of OceanBase cluster instances-normal (default): Standard cluster version (cloud disk)-normal_SSD: Standard cluster version (local disk)-history: history Library cluster version.
	Series pulumi.StringInput
	// Information about the zone where the cluster is deployed.
	Zones pulumi.StringArrayInput
}

func (BaseInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*baseInstanceArgs)(nil)).Elem()
}

type BaseInstanceInput interface {
	pulumi.Input

	ToBaseInstanceOutput() BaseInstanceOutput
	ToBaseInstanceOutputWithContext(ctx context.Context) BaseInstanceOutput
}

func (*BaseInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**BaseInstance)(nil)).Elem()
}

func (i *BaseInstance) ToBaseInstanceOutput() BaseInstanceOutput {
	return i.ToBaseInstanceOutputWithContext(context.Background())
}

func (i *BaseInstance) ToBaseInstanceOutputWithContext(ctx context.Context) BaseInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BaseInstanceOutput)
}

// BaseInstanceArrayInput is an input type that accepts BaseInstanceArray and BaseInstanceArrayOutput values.
// You can construct a concrete instance of `BaseInstanceArrayInput` via:
//
//	BaseInstanceArray{ BaseInstanceArgs{...} }
type BaseInstanceArrayInput interface {
	pulumi.Input

	ToBaseInstanceArrayOutput() BaseInstanceArrayOutput
	ToBaseInstanceArrayOutputWithContext(context.Context) BaseInstanceArrayOutput
}

type BaseInstanceArray []BaseInstanceInput

func (BaseInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BaseInstance)(nil)).Elem()
}

func (i BaseInstanceArray) ToBaseInstanceArrayOutput() BaseInstanceArrayOutput {
	return i.ToBaseInstanceArrayOutputWithContext(context.Background())
}

func (i BaseInstanceArray) ToBaseInstanceArrayOutputWithContext(ctx context.Context) BaseInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BaseInstanceArrayOutput)
}

// BaseInstanceMapInput is an input type that accepts BaseInstanceMap and BaseInstanceMapOutput values.
// You can construct a concrete instance of `BaseInstanceMapInput` via:
//
//	BaseInstanceMap{ "key": BaseInstanceArgs{...} }
type BaseInstanceMapInput interface {
	pulumi.Input

	ToBaseInstanceMapOutput() BaseInstanceMapOutput
	ToBaseInstanceMapOutputWithContext(context.Context) BaseInstanceMapOutput
}

type BaseInstanceMap map[string]BaseInstanceInput

func (BaseInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BaseInstance)(nil)).Elem()
}

func (i BaseInstanceMap) ToBaseInstanceMapOutput() BaseInstanceMapOutput {
	return i.ToBaseInstanceMapOutputWithContext(context.Background())
}

func (i BaseInstanceMap) ToBaseInstanceMapOutputWithContext(ctx context.Context) BaseInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BaseInstanceMapOutput)
}

type BaseInstanceOutput struct{ *pulumi.OutputState }

func (BaseInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BaseInstance)(nil)).Elem()
}

func (o BaseInstanceOutput) ToBaseInstanceOutput() BaseInstanceOutput {
	return o
}

func (o BaseInstanceOutput) ToBaseInstanceOutputWithContext(ctx context.Context) BaseInstanceOutput {
	return o
}

// Whether to automatically renew.
// It takes effect when the parameter ChargeType is PrePaid. Value range:
// - true: automatic renewal.
// - false (default): no automatic renewal.
func (o BaseInstanceOutput) AutoRenew() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BaseInstance) pulumi.BoolPtrOutput { return v.AutoRenew }).(pulumi.BoolPtrOutput)
}

// The duration of each auto-renewal. When the value of the AutoRenew parameter is True, this parameter is required.
// - PeriodUnit is Week, AutoRenewPeriod is {"1", "2", "3"}.
// - PeriodUnit is Month, AutoRenewPeriod is {"1", "2", "3", "6", "12"}.
func (o BaseInstanceOutput) AutoRenewPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BaseInstance) pulumi.IntPtrOutput { return v.AutoRenewPeriod }).(pulumi.IntPtrOutput)
}

// The backup retention policy after the cluster is deleted. The values are as follows:
// - receive_all: Keep all backup sets;
// - delete_all: delete all backup sets;
// - receive_last: Keep the last backup set.
// > **NOTE:**   The default value is delete_all.
func (o BaseInstanceOutput) BackupRetainMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BaseInstance) pulumi.StringPtrOutput { return v.BackupRetainMode }).(pulumi.StringPtrOutput)
}

// The product code of the OceanBase cluster._oceanbasepre_public_cn: Domestic station cloud database package Year-to-month package._oceanbasepost_public_cn: The domestic station cloud database is paid by the hour._obpre_public_intl: International Station Cloud Database Package Monthly Package.
func (o BaseInstanceOutput) CommodityCode() pulumi.StringOutput {
	return o.ApplyT(func(v *BaseInstance) pulumi.StringOutput { return v.CommodityCode }).(pulumi.StringOutput)
}

// The number of CPU cores of the cluster.
func (o BaseInstanceOutput) Cpu() pulumi.IntOutput {
	return o.ApplyT(func(v *BaseInstance) pulumi.IntOutput { return v.Cpu }).(pulumi.IntOutput)
}

// The creation time of the resource.
func (o BaseInstanceOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *BaseInstance) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The size of the storage space, in GB.
// The limits of storage space vary according to the cluster specifications, as follows:
//   - 8C32GB:100GB ~ 10000GB
//   - 14C70GB:200GB ~ 10000GB
//   - 30C180GB:400GB ~ 10000GB
//   - 62C400G:800GB ~ 10000GB.
//     The default value of each package is its minimum value.
func (o BaseInstanceOutput) DiskSize() pulumi.IntOutput {
	return o.ApplyT(func(v *BaseInstance) pulumi.IntOutput { return v.DiskSize }).(pulumi.IntOutput)
}

// The storage type of the cluster. Effective only in the standard cluster version (cloud disk).
// Two types are currently supported:
// - cloud_essd_pl1: cloud disk ESSD pl1.
// - cloud_essd_pl0: cloud disk ESSD pl0. The default value is cloud_essd_pl1.
func (o BaseInstanceOutput) DiskType() pulumi.StringOutput {
	return o.ApplyT(func(v *BaseInstance) pulumi.StringOutput { return v.DiskType }).(pulumi.StringOutput)
}

// Cluster specification information.
// Four packages are currently supported:
// - 4C16GB：4cores 16GB
// - 8C32GB：8cores 32GB
// - 14C70GB：14cores 70GB
// - 24C120GB：24cores 120GB
// - 30C180GB：30cores 180GB
// - 62C400GB：62cores 400GB
// - 104C600GB：104cores 600GB
// - 16C70GB：16cores 70GB
// - 32C160GB：32cores 160GB
// - 64C380GB：64cores 380GB
// - 20C32GB：20cores 32GB
// - 40C64GB：40cores 64GB
// - 16C32GB：16cores 32GB
// - 32C70GB：32cores 70GB
// - 64C180GB：64cores 180GB
// - 32C180GB：32cores 180GB
// - 64C400GB：64cores 400GB.
func (o BaseInstanceOutput) InstanceClass() pulumi.StringOutput {
	return o.ApplyT(func(v *BaseInstance) pulumi.StringOutput { return v.InstanceClass }).(pulumi.StringOutput)
}

// OceanBase cluster name.The length is 1 to 20 English or Chinese characters.If this parameter is not specified, the default value is the InstanceId of the cluster.
func (o BaseInstanceOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *BaseInstance) pulumi.StringOutput { return v.InstanceName }).(pulumi.StringOutput)
}

// The number of nodes in the cluster. If the deployment mode is n-n-n, the number of nodes is n * 3.
func (o BaseInstanceOutput) NodeNum() pulumi.StringOutput {
	return o.ApplyT(func(v *BaseInstance) pulumi.StringOutput { return v.NodeNum }).(pulumi.StringOutput)
}

// The OceanBase Server version number.
func (o BaseInstanceOutput) ObVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *BaseInstance) pulumi.StringOutput { return v.ObVersion }).(pulumi.StringOutput)
}

// The payment method of the instance. Value range:
// - Subscription: Package year and month. When you select this type of payment method, you must make sure that your account supports balance payment or credit payment. Otherwise, an InvalidPayMethod error message will be returned.
// - PayAsYouGo (default): Pay-as-you-go (default hourly billing).
func (o BaseInstanceOutput) PaymentType() pulumi.StringOutput {
	return o.ApplyT(func(v *BaseInstance) pulumi.StringOutput { return v.PaymentType }).(pulumi.StringOutput)
}

// The duration of the resource purchase. The unit is specified by the PeriodUnit. The parameter InstanceChargeType takes effect only when the value is PrePaid and is required. Once the DedicatedHostId is specified, the value cannot exceed the subscription duration of the dedicated host. When PeriodUnit = Week, Period values: {"1", "2", "3", "4"}. When PeriodUnit = Month, Period values: {"1", "2", "3", "4", "5", "6", "7", "8", "9", "12", "24", "36", "48", "60"}.
func (o BaseInstanceOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BaseInstance) pulumi.IntPtrOutput { return v.Period }).(pulumi.IntPtrOutput)
}

// The duration of the purchase of resources.Package year and Month value range: Month.Default value: Month of the package, which is billed by volume. The default period is Hour.
func (o BaseInstanceOutput) PeriodUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BaseInstance) pulumi.StringPtrOutput { return v.PeriodUnit }).(pulumi.StringPtrOutput)
}

// The ID of the enterprise resource group to which the instance resides.
func (o BaseInstanceOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *BaseInstance) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// Series of OceanBase cluster instances-normal (default): Standard cluster version (cloud disk)-normal_SSD: Standard cluster version (local disk)-history: history Library cluster version.
func (o BaseInstanceOutput) Series() pulumi.StringOutput {
	return o.ApplyT(func(v *BaseInstance) pulumi.StringOutput { return v.Series }).(pulumi.StringOutput)
}

// The status of the resource.
func (o BaseInstanceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *BaseInstance) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Information about the zone where the cluster is deployed.
func (o BaseInstanceOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BaseInstance) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

type BaseInstanceArrayOutput struct{ *pulumi.OutputState }

func (BaseInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BaseInstance)(nil)).Elem()
}

func (o BaseInstanceArrayOutput) ToBaseInstanceArrayOutput() BaseInstanceArrayOutput {
	return o
}

func (o BaseInstanceArrayOutput) ToBaseInstanceArrayOutputWithContext(ctx context.Context) BaseInstanceArrayOutput {
	return o
}

func (o BaseInstanceArrayOutput) Index(i pulumi.IntInput) BaseInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BaseInstance {
		return vs[0].([]*BaseInstance)[vs[1].(int)]
	}).(BaseInstanceOutput)
}

type BaseInstanceMapOutput struct{ *pulumi.OutputState }

func (BaseInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BaseInstance)(nil)).Elem()
}

func (o BaseInstanceMapOutput) ToBaseInstanceMapOutput() BaseInstanceMapOutput {
	return o
}

func (o BaseInstanceMapOutput) ToBaseInstanceMapOutputWithContext(ctx context.Context) BaseInstanceMapOutput {
	return o
}

func (o BaseInstanceMapOutput) MapIndex(k pulumi.StringInput) BaseInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BaseInstance {
		return vs[0].(map[string]*BaseInstance)[vs[1].(string)]
	}).(BaseInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BaseInstanceInput)(nil)).Elem(), &BaseInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*BaseInstanceArrayInput)(nil)).Elem(), BaseInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BaseInstanceMapInput)(nil)).Elem(), BaseInstanceMap{})
	pulumi.RegisterOutputType(BaseInstanceOutput{})
	pulumi.RegisterOutputType(BaseInstanceArrayOutput{})
	pulumi.RegisterOutputType(BaseInstanceMapOutput{})
}

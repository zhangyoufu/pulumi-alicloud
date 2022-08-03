// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lindorm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Lindorm Instance resource.
//
// For information about Lindorm Instance and how to use it, see [What is Instance](https://www.alibabacloud.com/help/zh/doc-detail/174640.html).
//
// > **NOTE:** Available in v1.132.0+.
//
// > **NOTE:**  The Lindorm Instance does not support updating the specifications of multiple different engines or the number of nodes at the same time.
//
// ## Import
//
// Lindorm Instance can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:lindorm/instance:Instance example <id>
// ```
type Instance struct {
	pulumi.CustomResourceState

	// The cold storage capacity of the instance. Unit: GB.
	ColdStorage pulumi.IntOutput `pulumi:"coldStorage"`
	// The core num.
	CoreNum pulumi.IntPtrOutput `pulumi:"coreNum"`
	// The core spec.
	CoreSpec pulumi.StringPtrOutput `pulumi:"coreSpec"`
	// The deletion protection of instance.
	DeletionProection pulumi.BoolOutput `pulumi:"deletionProection"`
	// The disk type of instance. Valid values: `capacityCloudStorage`, `cloudEfficiency`, `cloudEssd`, `cloudSsd`.
	DiskCategory pulumi.StringOutput `pulumi:"diskCategory"`
	// The duration of paid. Valid when the `paymentType` is `Subscription`.  When `pricingCycle` set to `Month`, the valid value id `1` to `9`.  When `pricingCycle` set to `Year`, the valid value id `1` to `3`.
	Duration pulumi.StringPtrOutput `pulumi:"duration"`
	// (Available in v1.163.0+) Whether to enable file engine.
	EnabledFileEngine pulumi.BoolOutput `pulumi:"enabledFileEngine"`
	// (Available in v1.163.0+) Whether to enable lts engine.
	EnabledLtsEngine pulumi.BoolOutput `pulumi:"enabledLtsEngine"`
	// (Available in v1.163.0+) Whether to enable search engine.
	EnabledSearchEngine pulumi.BoolOutput `pulumi:"enabledSearchEngine"`
	// (Available in v1.163.0+) Whether to enable table engine.
	EnabledTableEngine pulumi.BoolOutput `pulumi:"enabledTableEngine"`
	// (Available in v1.163.0+) Whether to enable time serires engine.
	EnabledTimeSeriresEngine pulumi.BoolOutput `pulumi:"enabledTimeSeriresEngine"`
	// The count of file engine.
	FileEngineNodeCount pulumi.IntOutput `pulumi:"fileEngineNodeCount"`
	// The specification of file engine. Valid values: `lindorm.c.xlarge`.
	FileEngineSpecification pulumi.StringOutput `pulumi:"fileEngineSpecification"`
	// The group name.
	GroupName pulumi.StringPtrOutput `pulumi:"groupName"`
	// The name of the instance.
	InstanceName pulumi.StringPtrOutput `pulumi:"instanceName"`
	// The storage capacity of the instance. Unit: GB. For example, the value 50 indicates 50 GB.
	InstanceStorage pulumi.StringPtrOutput `pulumi:"instanceStorage"`
	// The ip white list of instance.
	IpWhiteLists pulumi.StringArrayOutput `pulumi:"ipWhiteLists"`
	// The count of lindorm tunnel service.
	LtsNodeCount pulumi.IntOutput `pulumi:"ltsNodeCount"`
	// The specification of lindorm tunnel service. Valid values: `lindorm.g.2xlarge`, `lindorm.g.xlarge`.
	LtsNodeSpecification pulumi.StringOutput `pulumi:"ltsNodeSpecification"`
	// The billing method. Valid values: `PayAsYouGo` and `Subscription`.
	PaymentType pulumi.StringOutput `pulumi:"paymentType"`
	// The count of phoenix.
	PhoenixNodeCount pulumi.IntOutput `pulumi:"phoenixNodeCount"`
	// The specification of phoenix. Valid values: `lindorm.c.2xlarge`, `lindorm.c.4xlarge`, `lindorm.c.8xlarge`, `lindorm.c.xlarge`, `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`, `lindorm.g.xlarge`.
	PhoenixNodeSpecification pulumi.StringOutput `pulumi:"phoenixNodeSpecification"`
	// The pricing cycle. Valid when the `paymentType` is `Subscription`. Valid values: `Month` and `Year`.
	PricingCycle pulumi.StringPtrOutput `pulumi:"pricingCycle"`
	// The ID of the resource group.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// The count of search engine.
	SearchEngineNodeCount pulumi.IntOutput `pulumi:"searchEngineNodeCount"`
	// The specification of search engine. Valid values: `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`, `lindorm.g.xlarge`.
	SearchEngineSpecification pulumi.StringOutput `pulumi:"searchEngineSpecification"`
	// The status of Instance, enumerative: Valid values: `ACTIVATION`, `DELETED`, `CREATING`, `CLASS_CHANGING`, `LOCKED`, `INSTANCE_LEVEL_MODIFY`, `NET_MODIFYING`, `RESIZING`, `RESTARTING`, `MINOR_VERSION_TRANSING`.
	Status pulumi.StringOutput `pulumi:"status"`
	// The count of table engine.
	TableEngineNodeCount pulumi.IntOutput `pulumi:"tableEngineNodeCount"`
	// The specification of  table engine. Valid values: `lindorm.c.2xlarge`, `lindorm.c.4xlarge`, `lindorm.c.8xlarge`, `lindorm.c.xlarge`, `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`, `lindorm.g.xlarge`.
	TableEngineSpecification pulumi.StringOutput `pulumi:"tableEngineSpecification"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The count of time series engine.
	TimeSeriesEngineNodeCount pulumi.IntOutput `pulumi:"timeSeriesEngineNodeCount"`
	// The specification of time series engine. Valid values: `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`, `lindorm.g.xlarge`.
	TimeSeriresEngineSpecification pulumi.StringOutput `pulumi:"timeSeriresEngineSpecification"`
	// The upgrade type. **NOTE:** Field 'upgrade_type' has been deprecated from provider version 1.163.0 and it will be removed in the future version. Valid values:  `open-lindorm-engine`, `open-phoenix-engine`, `open-search-engine`, `open-tsdb-engine`,  `upgrade-cold-storage`, `upgrade-disk-size`,  `upgrade-lindorm-core-num`, `upgrade-lindorm-engine`,  `upgrade-search-core-num`, `upgrade-search-engine`, `upgrade-tsdb-core-num`, `upgrade-tsdb-engine`.
	//
	// Deprecated: Field 'upgrade_type' has been deprecated from provider version 1.163.0 and it will be removed in the future version.
	UpgradeType pulumi.StringPtrOutput `pulumi:"upgradeType"`
	// The vswitch id.
	VswitchId pulumi.StringOutput `pulumi:"vswitchId"`
	// The zone ID of the instance.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DiskCategory == nil {
		return nil, errors.New("invalid value for required argument 'DiskCategory'")
	}
	if args.PaymentType == nil {
		return nil, errors.New("invalid value for required argument 'PaymentType'")
	}
	if args.VswitchId == nil {
		return nil, errors.New("invalid value for required argument 'VswitchId'")
	}
	var resource Instance
	err := ctx.RegisterResource("alicloud:lindorm/instance:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("alicloud:lindorm/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// The cold storage capacity of the instance. Unit: GB.
	ColdStorage *int `pulumi:"coldStorage"`
	// The core num.
	CoreNum *int `pulumi:"coreNum"`
	// The core spec.
	CoreSpec *string `pulumi:"coreSpec"`
	// The deletion protection of instance.
	DeletionProection *bool `pulumi:"deletionProection"`
	// The disk type of instance. Valid values: `capacityCloudStorage`, `cloudEfficiency`, `cloudEssd`, `cloudSsd`.
	DiskCategory *string `pulumi:"diskCategory"`
	// The duration of paid. Valid when the `paymentType` is `Subscription`.  When `pricingCycle` set to `Month`, the valid value id `1` to `9`.  When `pricingCycle` set to `Year`, the valid value id `1` to `3`.
	Duration *string `pulumi:"duration"`
	// (Available in v1.163.0+) Whether to enable file engine.
	EnabledFileEngine *bool `pulumi:"enabledFileEngine"`
	// (Available in v1.163.0+) Whether to enable lts engine.
	EnabledLtsEngine *bool `pulumi:"enabledLtsEngine"`
	// (Available in v1.163.0+) Whether to enable search engine.
	EnabledSearchEngine *bool `pulumi:"enabledSearchEngine"`
	// (Available in v1.163.0+) Whether to enable table engine.
	EnabledTableEngine *bool `pulumi:"enabledTableEngine"`
	// (Available in v1.163.0+) Whether to enable time serires engine.
	EnabledTimeSeriresEngine *bool `pulumi:"enabledTimeSeriresEngine"`
	// The count of file engine.
	FileEngineNodeCount *int `pulumi:"fileEngineNodeCount"`
	// The specification of file engine. Valid values: `lindorm.c.xlarge`.
	FileEngineSpecification *string `pulumi:"fileEngineSpecification"`
	// The group name.
	GroupName *string `pulumi:"groupName"`
	// The name of the instance.
	InstanceName *string `pulumi:"instanceName"`
	// The storage capacity of the instance. Unit: GB. For example, the value 50 indicates 50 GB.
	InstanceStorage *string `pulumi:"instanceStorage"`
	// The ip white list of instance.
	IpWhiteLists []string `pulumi:"ipWhiteLists"`
	// The count of lindorm tunnel service.
	LtsNodeCount *int `pulumi:"ltsNodeCount"`
	// The specification of lindorm tunnel service. Valid values: `lindorm.g.2xlarge`, `lindorm.g.xlarge`.
	LtsNodeSpecification *string `pulumi:"ltsNodeSpecification"`
	// The billing method. Valid values: `PayAsYouGo` and `Subscription`.
	PaymentType *string `pulumi:"paymentType"`
	// The count of phoenix.
	PhoenixNodeCount *int `pulumi:"phoenixNodeCount"`
	// The specification of phoenix. Valid values: `lindorm.c.2xlarge`, `lindorm.c.4xlarge`, `lindorm.c.8xlarge`, `lindorm.c.xlarge`, `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`, `lindorm.g.xlarge`.
	PhoenixNodeSpecification *string `pulumi:"phoenixNodeSpecification"`
	// The pricing cycle. Valid when the `paymentType` is `Subscription`. Valid values: `Month` and `Year`.
	PricingCycle *string `pulumi:"pricingCycle"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The count of search engine.
	SearchEngineNodeCount *int `pulumi:"searchEngineNodeCount"`
	// The specification of search engine. Valid values: `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`, `lindorm.g.xlarge`.
	SearchEngineSpecification *string `pulumi:"searchEngineSpecification"`
	// The status of Instance, enumerative: Valid values: `ACTIVATION`, `DELETED`, `CREATING`, `CLASS_CHANGING`, `LOCKED`, `INSTANCE_LEVEL_MODIFY`, `NET_MODIFYING`, `RESIZING`, `RESTARTING`, `MINOR_VERSION_TRANSING`.
	Status *string `pulumi:"status"`
	// The count of table engine.
	TableEngineNodeCount *int `pulumi:"tableEngineNodeCount"`
	// The specification of  table engine. Valid values: `lindorm.c.2xlarge`, `lindorm.c.4xlarge`, `lindorm.c.8xlarge`, `lindorm.c.xlarge`, `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`, `lindorm.g.xlarge`.
	TableEngineSpecification *string `pulumi:"tableEngineSpecification"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The count of time series engine.
	TimeSeriesEngineNodeCount *int `pulumi:"timeSeriesEngineNodeCount"`
	// The specification of time series engine. Valid values: `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`, `lindorm.g.xlarge`.
	TimeSeriresEngineSpecification *string `pulumi:"timeSeriresEngineSpecification"`
	// The upgrade type. **NOTE:** Field 'upgrade_type' has been deprecated from provider version 1.163.0 and it will be removed in the future version. Valid values:  `open-lindorm-engine`, `open-phoenix-engine`, `open-search-engine`, `open-tsdb-engine`,  `upgrade-cold-storage`, `upgrade-disk-size`,  `upgrade-lindorm-core-num`, `upgrade-lindorm-engine`,  `upgrade-search-core-num`, `upgrade-search-engine`, `upgrade-tsdb-core-num`, `upgrade-tsdb-engine`.
	//
	// Deprecated: Field 'upgrade_type' has been deprecated from provider version 1.163.0 and it will be removed in the future version.
	UpgradeType *string `pulumi:"upgradeType"`
	// The vswitch id.
	VswitchId *string `pulumi:"vswitchId"`
	// The zone ID of the instance.
	ZoneId *string `pulumi:"zoneId"`
}

type InstanceState struct {
	// The cold storage capacity of the instance. Unit: GB.
	ColdStorage pulumi.IntPtrInput
	// The core num.
	CoreNum pulumi.IntPtrInput
	// The core spec.
	CoreSpec pulumi.StringPtrInput
	// The deletion protection of instance.
	DeletionProection pulumi.BoolPtrInput
	// The disk type of instance. Valid values: `capacityCloudStorage`, `cloudEfficiency`, `cloudEssd`, `cloudSsd`.
	DiskCategory pulumi.StringPtrInput
	// The duration of paid. Valid when the `paymentType` is `Subscription`.  When `pricingCycle` set to `Month`, the valid value id `1` to `9`.  When `pricingCycle` set to `Year`, the valid value id `1` to `3`.
	Duration pulumi.StringPtrInput
	// (Available in v1.163.0+) Whether to enable file engine.
	EnabledFileEngine pulumi.BoolPtrInput
	// (Available in v1.163.0+) Whether to enable lts engine.
	EnabledLtsEngine pulumi.BoolPtrInput
	// (Available in v1.163.0+) Whether to enable search engine.
	EnabledSearchEngine pulumi.BoolPtrInput
	// (Available in v1.163.0+) Whether to enable table engine.
	EnabledTableEngine pulumi.BoolPtrInput
	// (Available in v1.163.0+) Whether to enable time serires engine.
	EnabledTimeSeriresEngine pulumi.BoolPtrInput
	// The count of file engine.
	FileEngineNodeCount pulumi.IntPtrInput
	// The specification of file engine. Valid values: `lindorm.c.xlarge`.
	FileEngineSpecification pulumi.StringPtrInput
	// The group name.
	GroupName pulumi.StringPtrInput
	// The name of the instance.
	InstanceName pulumi.StringPtrInput
	// The storage capacity of the instance. Unit: GB. For example, the value 50 indicates 50 GB.
	InstanceStorage pulumi.StringPtrInput
	// The ip white list of instance.
	IpWhiteLists pulumi.StringArrayInput
	// The count of lindorm tunnel service.
	LtsNodeCount pulumi.IntPtrInput
	// The specification of lindorm tunnel service. Valid values: `lindorm.g.2xlarge`, `lindorm.g.xlarge`.
	LtsNodeSpecification pulumi.StringPtrInput
	// The billing method. Valid values: `PayAsYouGo` and `Subscription`.
	PaymentType pulumi.StringPtrInput
	// The count of phoenix.
	PhoenixNodeCount pulumi.IntPtrInput
	// The specification of phoenix. Valid values: `lindorm.c.2xlarge`, `lindorm.c.4xlarge`, `lindorm.c.8xlarge`, `lindorm.c.xlarge`, `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`, `lindorm.g.xlarge`.
	PhoenixNodeSpecification pulumi.StringPtrInput
	// The pricing cycle. Valid when the `paymentType` is `Subscription`. Valid values: `Month` and `Year`.
	PricingCycle pulumi.StringPtrInput
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	// The count of search engine.
	SearchEngineNodeCount pulumi.IntPtrInput
	// The specification of search engine. Valid values: `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`, `lindorm.g.xlarge`.
	SearchEngineSpecification pulumi.StringPtrInput
	// The status of Instance, enumerative: Valid values: `ACTIVATION`, `DELETED`, `CREATING`, `CLASS_CHANGING`, `LOCKED`, `INSTANCE_LEVEL_MODIFY`, `NET_MODIFYING`, `RESIZING`, `RESTARTING`, `MINOR_VERSION_TRANSING`.
	Status pulumi.StringPtrInput
	// The count of table engine.
	TableEngineNodeCount pulumi.IntPtrInput
	// The specification of  table engine. Valid values: `lindorm.c.2xlarge`, `lindorm.c.4xlarge`, `lindorm.c.8xlarge`, `lindorm.c.xlarge`, `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`, `lindorm.g.xlarge`.
	TableEngineSpecification pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The count of time series engine.
	TimeSeriesEngineNodeCount pulumi.IntPtrInput
	// The specification of time series engine. Valid values: `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`, `lindorm.g.xlarge`.
	TimeSeriresEngineSpecification pulumi.StringPtrInput
	// The upgrade type. **NOTE:** Field 'upgrade_type' has been deprecated from provider version 1.163.0 and it will be removed in the future version. Valid values:  `open-lindorm-engine`, `open-phoenix-engine`, `open-search-engine`, `open-tsdb-engine`,  `upgrade-cold-storage`, `upgrade-disk-size`,  `upgrade-lindorm-core-num`, `upgrade-lindorm-engine`,  `upgrade-search-core-num`, `upgrade-search-engine`, `upgrade-tsdb-core-num`, `upgrade-tsdb-engine`.
	//
	// Deprecated: Field 'upgrade_type' has been deprecated from provider version 1.163.0 and it will be removed in the future version.
	UpgradeType pulumi.StringPtrInput
	// The vswitch id.
	VswitchId pulumi.StringPtrInput
	// The zone ID of the instance.
	ZoneId pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// The cold storage capacity of the instance. Unit: GB.
	ColdStorage *int `pulumi:"coldStorage"`
	// The core num.
	CoreNum *int `pulumi:"coreNum"`
	// The core spec.
	CoreSpec *string `pulumi:"coreSpec"`
	// The deletion protection of instance.
	DeletionProection *bool `pulumi:"deletionProection"`
	// The disk type of instance. Valid values: `capacityCloudStorage`, `cloudEfficiency`, `cloudEssd`, `cloudSsd`.
	DiskCategory string `pulumi:"diskCategory"`
	// The duration of paid. Valid when the `paymentType` is `Subscription`.  When `pricingCycle` set to `Month`, the valid value id `1` to `9`.  When `pricingCycle` set to `Year`, the valid value id `1` to `3`.
	Duration *string `pulumi:"duration"`
	// The count of file engine.
	FileEngineNodeCount *int `pulumi:"fileEngineNodeCount"`
	// The specification of file engine. Valid values: `lindorm.c.xlarge`.
	FileEngineSpecification *string `pulumi:"fileEngineSpecification"`
	// The group name.
	GroupName *string `pulumi:"groupName"`
	// The name of the instance.
	InstanceName *string `pulumi:"instanceName"`
	// The storage capacity of the instance. Unit: GB. For example, the value 50 indicates 50 GB.
	InstanceStorage *string `pulumi:"instanceStorage"`
	// The ip white list of instance.
	IpWhiteLists []string `pulumi:"ipWhiteLists"`
	// The count of lindorm tunnel service.
	LtsNodeCount *int `pulumi:"ltsNodeCount"`
	// The specification of lindorm tunnel service. Valid values: `lindorm.g.2xlarge`, `lindorm.g.xlarge`.
	LtsNodeSpecification *string `pulumi:"ltsNodeSpecification"`
	// The billing method. Valid values: `PayAsYouGo` and `Subscription`.
	PaymentType string `pulumi:"paymentType"`
	// The count of phoenix.
	PhoenixNodeCount *int `pulumi:"phoenixNodeCount"`
	// The specification of phoenix. Valid values: `lindorm.c.2xlarge`, `lindorm.c.4xlarge`, `lindorm.c.8xlarge`, `lindorm.c.xlarge`, `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`, `lindorm.g.xlarge`.
	PhoenixNodeSpecification *string `pulumi:"phoenixNodeSpecification"`
	// The pricing cycle. Valid when the `paymentType` is `Subscription`. Valid values: `Month` and `Year`.
	PricingCycle *string `pulumi:"pricingCycle"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The count of search engine.
	SearchEngineNodeCount *int `pulumi:"searchEngineNodeCount"`
	// The specification of search engine. Valid values: `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`, `lindorm.g.xlarge`.
	SearchEngineSpecification *string `pulumi:"searchEngineSpecification"`
	// The count of table engine.
	TableEngineNodeCount *int `pulumi:"tableEngineNodeCount"`
	// The specification of  table engine. Valid values: `lindorm.c.2xlarge`, `lindorm.c.4xlarge`, `lindorm.c.8xlarge`, `lindorm.c.xlarge`, `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`, `lindorm.g.xlarge`.
	TableEngineSpecification *string `pulumi:"tableEngineSpecification"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The count of time series engine.
	TimeSeriesEngineNodeCount *int `pulumi:"timeSeriesEngineNodeCount"`
	// The specification of time series engine. Valid values: `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`, `lindorm.g.xlarge`.
	TimeSeriresEngineSpecification *string `pulumi:"timeSeriresEngineSpecification"`
	// The upgrade type. **NOTE:** Field 'upgrade_type' has been deprecated from provider version 1.163.0 and it will be removed in the future version. Valid values:  `open-lindorm-engine`, `open-phoenix-engine`, `open-search-engine`, `open-tsdb-engine`,  `upgrade-cold-storage`, `upgrade-disk-size`,  `upgrade-lindorm-core-num`, `upgrade-lindorm-engine`,  `upgrade-search-core-num`, `upgrade-search-engine`, `upgrade-tsdb-core-num`, `upgrade-tsdb-engine`.
	//
	// Deprecated: Field 'upgrade_type' has been deprecated from provider version 1.163.0 and it will be removed in the future version.
	UpgradeType *string `pulumi:"upgradeType"`
	// The vswitch id.
	VswitchId string `pulumi:"vswitchId"`
	// The zone ID of the instance.
	ZoneId *string `pulumi:"zoneId"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// The cold storage capacity of the instance. Unit: GB.
	ColdStorage pulumi.IntPtrInput
	// The core num.
	CoreNum pulumi.IntPtrInput
	// The core spec.
	CoreSpec pulumi.StringPtrInput
	// The deletion protection of instance.
	DeletionProection pulumi.BoolPtrInput
	// The disk type of instance. Valid values: `capacityCloudStorage`, `cloudEfficiency`, `cloudEssd`, `cloudSsd`.
	DiskCategory pulumi.StringInput
	// The duration of paid. Valid when the `paymentType` is `Subscription`.  When `pricingCycle` set to `Month`, the valid value id `1` to `9`.  When `pricingCycle` set to `Year`, the valid value id `1` to `3`.
	Duration pulumi.StringPtrInput
	// The count of file engine.
	FileEngineNodeCount pulumi.IntPtrInput
	// The specification of file engine. Valid values: `lindorm.c.xlarge`.
	FileEngineSpecification pulumi.StringPtrInput
	// The group name.
	GroupName pulumi.StringPtrInput
	// The name of the instance.
	InstanceName pulumi.StringPtrInput
	// The storage capacity of the instance. Unit: GB. For example, the value 50 indicates 50 GB.
	InstanceStorage pulumi.StringPtrInput
	// The ip white list of instance.
	IpWhiteLists pulumi.StringArrayInput
	// The count of lindorm tunnel service.
	LtsNodeCount pulumi.IntPtrInput
	// The specification of lindorm tunnel service. Valid values: `lindorm.g.2xlarge`, `lindorm.g.xlarge`.
	LtsNodeSpecification pulumi.StringPtrInput
	// The billing method. Valid values: `PayAsYouGo` and `Subscription`.
	PaymentType pulumi.StringInput
	// The count of phoenix.
	PhoenixNodeCount pulumi.IntPtrInput
	// The specification of phoenix. Valid values: `lindorm.c.2xlarge`, `lindorm.c.4xlarge`, `lindorm.c.8xlarge`, `lindorm.c.xlarge`, `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`, `lindorm.g.xlarge`.
	PhoenixNodeSpecification pulumi.StringPtrInput
	// The pricing cycle. Valid when the `paymentType` is `Subscription`. Valid values: `Month` and `Year`.
	PricingCycle pulumi.StringPtrInput
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	// The count of search engine.
	SearchEngineNodeCount pulumi.IntPtrInput
	// The specification of search engine. Valid values: `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`, `lindorm.g.xlarge`.
	SearchEngineSpecification pulumi.StringPtrInput
	// The count of table engine.
	TableEngineNodeCount pulumi.IntPtrInput
	// The specification of  table engine. Valid values: `lindorm.c.2xlarge`, `lindorm.c.4xlarge`, `lindorm.c.8xlarge`, `lindorm.c.xlarge`, `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`, `lindorm.g.xlarge`.
	TableEngineSpecification pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The count of time series engine.
	TimeSeriesEngineNodeCount pulumi.IntPtrInput
	// The specification of time series engine. Valid values: `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`, `lindorm.g.xlarge`.
	TimeSeriresEngineSpecification pulumi.StringPtrInput
	// The upgrade type. **NOTE:** Field 'upgrade_type' has been deprecated from provider version 1.163.0 and it will be removed in the future version. Valid values:  `open-lindorm-engine`, `open-phoenix-engine`, `open-search-engine`, `open-tsdb-engine`,  `upgrade-cold-storage`, `upgrade-disk-size`,  `upgrade-lindorm-core-num`, `upgrade-lindorm-engine`,  `upgrade-search-core-num`, `upgrade-search-engine`, `upgrade-tsdb-core-num`, `upgrade-tsdb-engine`.
	//
	// Deprecated: Field 'upgrade_type' has been deprecated from provider version 1.163.0 and it will be removed in the future version.
	UpgradeType pulumi.StringPtrInput
	// The vswitch id.
	VswitchId pulumi.StringInput
	// The zone ID of the instance.
	ZoneId pulumi.StringPtrInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

// InstanceArrayInput is an input type that accepts InstanceArray and InstanceArrayOutput values.
// You can construct a concrete instance of `InstanceArrayInput` via:
//
//          InstanceArray{ InstanceArgs{...} }
type InstanceArrayInput interface {
	pulumi.Input

	ToInstanceArrayOutput() InstanceArrayOutput
	ToInstanceArrayOutputWithContext(context.Context) InstanceArrayOutput
}

type InstanceArray []InstanceInput

func (InstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (i InstanceArray) ToInstanceArrayOutput() InstanceArrayOutput {
	return i.ToInstanceArrayOutputWithContext(context.Background())
}

func (i InstanceArray) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceArrayOutput)
}

// InstanceMapInput is an input type that accepts InstanceMap and InstanceMapOutput values.
// You can construct a concrete instance of `InstanceMapInput` via:
//
//          InstanceMap{ "key": InstanceArgs{...} }
type InstanceMapInput interface {
	pulumi.Input

	ToInstanceMapOutput() InstanceMapOutput
	ToInstanceMapOutputWithContext(context.Context) InstanceMapOutput
}

type InstanceMap map[string]InstanceInput

func (InstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (i InstanceMap) ToInstanceMapOutput() InstanceMapOutput {
	return i.ToInstanceMapOutputWithContext(context.Background())
}

func (i InstanceMap) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMapOutput)
}

type InstanceOutput struct{ *pulumi.OutputState }

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

// The cold storage capacity of the instance. Unit: GB.
func (o InstanceOutput) ColdStorage() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.ColdStorage }).(pulumi.IntOutput)
}

// The core num.
func (o InstanceOutput) CoreNum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.CoreNum }).(pulumi.IntPtrOutput)
}

// The core spec.
func (o InstanceOutput) CoreSpec() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.CoreSpec }).(pulumi.StringPtrOutput)
}

// The deletion protection of instance.
func (o InstanceOutput) DeletionProection() pulumi.BoolOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolOutput { return v.DeletionProection }).(pulumi.BoolOutput)
}

// The disk type of instance. Valid values: `capacityCloudStorage`, `cloudEfficiency`, `cloudEssd`, `cloudSsd`.
func (o InstanceOutput) DiskCategory() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.DiskCategory }).(pulumi.StringOutput)
}

// The duration of paid. Valid when the `paymentType` is `Subscription`.  When `pricingCycle` set to `Month`, the valid value id `1` to `9`.  When `pricingCycle` set to `Year`, the valid value id `1` to `3`.
func (o InstanceOutput) Duration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.Duration }).(pulumi.StringPtrOutput)
}

// (Available in v1.163.0+) Whether to enable file engine.
func (o InstanceOutput) EnabledFileEngine() pulumi.BoolOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolOutput { return v.EnabledFileEngine }).(pulumi.BoolOutput)
}

// (Available in v1.163.0+) Whether to enable lts engine.
func (o InstanceOutput) EnabledLtsEngine() pulumi.BoolOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolOutput { return v.EnabledLtsEngine }).(pulumi.BoolOutput)
}

// (Available in v1.163.0+) Whether to enable search engine.
func (o InstanceOutput) EnabledSearchEngine() pulumi.BoolOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolOutput { return v.EnabledSearchEngine }).(pulumi.BoolOutput)
}

// (Available in v1.163.0+) Whether to enable table engine.
func (o InstanceOutput) EnabledTableEngine() pulumi.BoolOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolOutput { return v.EnabledTableEngine }).(pulumi.BoolOutput)
}

// (Available in v1.163.0+) Whether to enable time serires engine.
func (o InstanceOutput) EnabledTimeSeriresEngine() pulumi.BoolOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolOutput { return v.EnabledTimeSeriresEngine }).(pulumi.BoolOutput)
}

// The count of file engine.
func (o InstanceOutput) FileEngineNodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.FileEngineNodeCount }).(pulumi.IntOutput)
}

// The specification of file engine. Valid values: `lindorm.c.xlarge`.
func (o InstanceOutput) FileEngineSpecification() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.FileEngineSpecification }).(pulumi.StringOutput)
}

// The group name.
func (o InstanceOutput) GroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.GroupName }).(pulumi.StringPtrOutput)
}

// The name of the instance.
func (o InstanceOutput) InstanceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.InstanceName }).(pulumi.StringPtrOutput)
}

// The storage capacity of the instance. Unit: GB. For example, the value 50 indicates 50 GB.
func (o InstanceOutput) InstanceStorage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.InstanceStorage }).(pulumi.StringPtrOutput)
}

// The ip white list of instance.
func (o InstanceOutput) IpWhiteLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringArrayOutput { return v.IpWhiteLists }).(pulumi.StringArrayOutput)
}

// The count of lindorm tunnel service.
func (o InstanceOutput) LtsNodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.LtsNodeCount }).(pulumi.IntOutput)
}

// The specification of lindorm tunnel service. Valid values: `lindorm.g.2xlarge`, `lindorm.g.xlarge`.
func (o InstanceOutput) LtsNodeSpecification() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.LtsNodeSpecification }).(pulumi.StringOutput)
}

// The billing method. Valid values: `PayAsYouGo` and `Subscription`.
func (o InstanceOutput) PaymentType() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.PaymentType }).(pulumi.StringOutput)
}

// The count of phoenix.
func (o InstanceOutput) PhoenixNodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.PhoenixNodeCount }).(pulumi.IntOutput)
}

// The specification of phoenix. Valid values: `lindorm.c.2xlarge`, `lindorm.c.4xlarge`, `lindorm.c.8xlarge`, `lindorm.c.xlarge`, `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`, `lindorm.g.xlarge`.
func (o InstanceOutput) PhoenixNodeSpecification() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.PhoenixNodeSpecification }).(pulumi.StringOutput)
}

// The pricing cycle. Valid when the `paymentType` is `Subscription`. Valid values: `Month` and `Year`.
func (o InstanceOutput) PricingCycle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.PricingCycle }).(pulumi.StringPtrOutput)
}

// The ID of the resource group.
func (o InstanceOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// The count of search engine.
func (o InstanceOutput) SearchEngineNodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.SearchEngineNodeCount }).(pulumi.IntOutput)
}

// The specification of search engine. Valid values: `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`, `lindorm.g.xlarge`.
func (o InstanceOutput) SearchEngineSpecification() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.SearchEngineSpecification }).(pulumi.StringOutput)
}

// The status of Instance, enumerative: Valid values: `ACTIVATION`, `DELETED`, `CREATING`, `CLASS_CHANGING`, `LOCKED`, `INSTANCE_LEVEL_MODIFY`, `NET_MODIFYING`, `RESIZING`, `RESTARTING`, `MINOR_VERSION_TRANSING`.
func (o InstanceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The count of table engine.
func (o InstanceOutput) TableEngineNodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.TableEngineNodeCount }).(pulumi.IntOutput)
}

// The specification of  table engine. Valid values: `lindorm.c.2xlarge`, `lindorm.c.4xlarge`, `lindorm.c.8xlarge`, `lindorm.c.xlarge`, `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`, `lindorm.g.xlarge`.
func (o InstanceOutput) TableEngineSpecification() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.TableEngineSpecification }).(pulumi.StringOutput)
}

// A mapping of tags to assign to the resource.
func (o InstanceOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *Instance) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// The count of time series engine.
func (o InstanceOutput) TimeSeriesEngineNodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.TimeSeriesEngineNodeCount }).(pulumi.IntOutput)
}

// The specification of time series engine. Valid values: `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`, `lindorm.g.xlarge`.
func (o InstanceOutput) TimeSeriresEngineSpecification() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.TimeSeriresEngineSpecification }).(pulumi.StringOutput)
}

// The upgrade type. **NOTE:** Field 'upgrade_type' has been deprecated from provider version 1.163.0 and it will be removed in the future version. Valid values:  `open-lindorm-engine`, `open-phoenix-engine`, `open-search-engine`, `open-tsdb-engine`,  `upgrade-cold-storage`, `upgrade-disk-size`,  `upgrade-lindorm-core-num`, `upgrade-lindorm-engine`,  `upgrade-search-core-num`, `upgrade-search-engine`, `upgrade-tsdb-core-num`, `upgrade-tsdb-engine`.
//
// Deprecated: Field 'upgrade_type' has been deprecated from provider version 1.163.0 and it will be removed in the future version.
func (o InstanceOutput) UpgradeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.UpgradeType }).(pulumi.StringPtrOutput)
}

// The vswitch id.
func (o InstanceOutput) VswitchId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.VswitchId }).(pulumi.StringOutput)
}

// The zone ID of the instance.
func (o InstanceOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ZoneId }).(pulumi.StringOutput)
}

type InstanceArrayOutput struct{ *pulumi.OutputState }

func (InstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (o InstanceArrayOutput) ToInstanceArrayOutput() InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) Index(i pulumi.IntInput) InstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].([]*Instance)[vs[1].(int)]
	}).(InstanceOutput)
}

type InstanceMapOutput struct{ *pulumi.OutputState }

func (InstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (o InstanceMapOutput) ToInstanceMapOutput() InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) MapIndex(k pulumi.StringInput) InstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].(map[string]*Instance)[vs[1].(string)]
	}).(InstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceInput)(nil)).Elem(), &Instance{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceArrayInput)(nil)).Elem(), InstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMapInput)(nil)).Elem(), InstanceMap{})
	pulumi.RegisterOutputType(InstanceOutput{})
	pulumi.RegisterOutputType(InstanceArrayOutput{})
	pulumi.RegisterOutputType(InstanceMapOutput{})
}

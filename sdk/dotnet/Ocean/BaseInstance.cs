// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ocean
{
    /// <summary>
    /// Provides a Ocean Base Instance resource.
    /// 
    /// For information about Ocean Base Instance and how to use it, see [What is Instance](https://www.alibabacloud.com/help/en/apsaradb-for-oceanbase/latest/what-is-oceanbase-database).
    /// 
    /// &gt; **NOTE:** Available since v1.203.0.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "terraform-example";
    ///     var defaultZones = AliCloud.GetZones.Invoke();
    /// 
    ///     var defaultResourceGroups = AliCloud.ResourceManager.GetResourceGroups.Invoke();
    /// 
    ///     var defaultBaseInstance = new AliCloud.Ocean.BaseInstance("defaultBaseInstance", new()
    ///     {
    ///         ResourceGroupId = defaultResourceGroups.Apply(getResourceGroupsResult =&gt; getResourceGroupsResult.Ids[0]),
    ///         Zones = new[]
    ///         {
    ///             Output.Tuple(defaultZones, defaultZones.Apply(getZonesResult =&gt; getZonesResult.Ids).Length).Apply(values =&gt;
    ///             {
    ///                 var defaultZones = values.Item1;
    ///                 var length = values.Item2;
    ///                 return defaultZones.Apply(getZonesResult =&gt; getZonesResult.Ids)[length - 2];
    ///             }),
    ///             Output.Tuple(defaultZones, defaultZones.Apply(getZonesResult =&gt; getZonesResult.Ids).Length).Apply(values =&gt;
    ///             {
    ///                 var defaultZones = values.Item1;
    ///                 var length = values.Item2;
    ///                 return defaultZones.Apply(getZonesResult =&gt; getZonesResult.Ids)[length - 3];
    ///             }),
    ///             Output.Tuple(defaultZones, defaultZones.Apply(getZonesResult =&gt; getZonesResult.Ids).Length).Apply(values =&gt;
    ///             {
    ///                 var defaultZones = values.Item1;
    ///                 var length = values.Item2;
    ///                 return defaultZones.Apply(getZonesResult =&gt; getZonesResult.Ids)[length - 4];
    ///             }),
    ///         },
    ///         AutoRenew = false,
    ///         DiskSize = 100,
    ///         PaymentType = "PayAsYouGo",
    ///         InstanceClass = "8C32GB",
    ///         BackupRetainMode = "delete_all",
    ///         Series = "normal",
    ///         InstanceName = name,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Ocean Base Instance can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:ocean/baseInstance:BaseInstance example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:ocean/baseInstance:BaseInstance")]
    public partial class BaseInstance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether to automatically renew.
        /// It takes effect when the parameter ChargeType is PrePaid. Value range:
        /// - true: automatic renewal.
        /// - false (default): no automatic renewal.
        /// </summary>
        [Output("autoRenew")]
        public Output<bool?> AutoRenew { get; private set; } = null!;

        /// <summary>
        /// The duration of each auto-renewal. When the value of the AutoRenew parameter is True, this parameter is required.
        /// - PeriodUnit is Week, AutoRenewPeriod is {"1", "2", "3"}.
        /// - PeriodUnit is Month, AutoRenewPeriod is {"1", "2", "3", "6", "12"}.
        /// </summary>
        [Output("autoRenewPeriod")]
        public Output<int?> AutoRenewPeriod { get; private set; } = null!;

        /// <summary>
        /// The backup retention policy after the cluster is deleted. The values are as follows:
        /// - receive_all: Keep all backup sets;
        /// - delete_all: delete all backup sets;
        /// - receive_last: Keep the last backup set.
        /// &gt; **NOTE:**   The default value is delete_all.
        /// </summary>
        [Output("backupRetainMode")]
        public Output<string?> BackupRetainMode { get; private set; } = null!;

        /// <summary>
        /// The product code of the OceanBase cluster._oceanbasepre_public_cn: Domestic station cloud database package Year-to-month package._oceanbasepost_public_cn: The domestic station cloud database is paid by the hour._obpre_public_intl: International Station Cloud Database Package Monthly Package.
        /// </summary>
        [Output("commodityCode")]
        public Output<string> CommodityCode { get; private set; } = null!;

        /// <summary>
        /// The number of CPU cores of the cluster.
        /// </summary>
        [Output("cpu")]
        public Output<int> Cpu { get; private set; } = null!;

        /// <summary>
        /// The creation time of the resource.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The size of the storage space, in GB.
        /// The limits of storage space vary according to the cluster specifications, as follows:
        /// - 8C32GB:100GB ~ 10000GB
        /// - 14C70GB:200GB ~ 10000GB
        /// - 30C180GB:400GB ~ 10000GB
        /// - 62C400G:800GB ~ 10000GB.
        /// The default value of each package is its minimum value.
        /// </summary>
        [Output("diskSize")]
        public Output<int> DiskSize { get; private set; } = null!;

        /// <summary>
        /// The storage type of the cluster. Effective only in the standard cluster version (cloud disk).
        /// Two types are currently supported:
        /// - cloud_essd_pl1: cloud disk ESSD pl1.
        /// - cloud_essd_pl0: cloud disk ESSD pl0. The default value is cloud_essd_pl1.
        /// </summary>
        [Output("diskType")]
        public Output<string> DiskType { get; private set; } = null!;

        /// <summary>
        /// Cluster specification information.
        /// Four packages are currently supported:
        /// - 4C16GB：4cores 16GB
        /// - 8C32GB：8cores 32GB
        /// - 14C70GB：14cores 70GB
        /// - 24C120GB：24cores 120GB
        /// - 30C180GB：30cores 180GB
        /// - 62C400GB：62cores 400GB
        /// - 104C600GB：104cores 600GB
        /// - 16C70GB：16cores 70GB
        /// - 32C160GB：32cores 160GB
        /// - 64C380GB：64cores 380GB
        /// - 20C32GB：20cores 32GB
        /// - 40C64GB：40cores 64GB
        /// - 16C32GB：16cores 32GB
        /// - 32C70GB：32cores 70GB
        /// - 64C180GB：64cores 180GB
        /// - 32C180GB：32cores 180GB
        /// - 64C400GB：64cores 400GB.
        /// </summary>
        [Output("instanceClass")]
        public Output<string> InstanceClass { get; private set; } = null!;

        /// <summary>
        /// OceanBase cluster name.The length is 1 to 20 English or Chinese characters.If this parameter is not specified, the default value is the InstanceId of the cluster.
        /// </summary>
        [Output("instanceName")]
        public Output<string> InstanceName { get; private set; } = null!;

        /// <summary>
        /// The number of nodes in the cluster. If the deployment mode is n-n-n, the number of nodes is n * 3.
        /// </summary>
        [Output("nodeNum")]
        public Output<string> NodeNum { get; private set; } = null!;

        /// <summary>
        /// The OceanBase Server version number.
        /// </summary>
        [Output("obVersion")]
        public Output<string> ObVersion { get; private set; } = null!;

        /// <summary>
        /// The payment method of the instance. Value range:
        /// - Subscription: Package year and month. When you select this type of payment method, you must make sure that your account supports balance payment or credit payment. Otherwise, an InvalidPayMethod error message will be returned.
        /// - PayAsYouGo (default): Pay-as-you-go (default hourly billing).
        /// </summary>
        [Output("paymentType")]
        public Output<string> PaymentType { get; private set; } = null!;

        /// <summary>
        /// The duration of the resource purchase. The unit is specified by the PeriodUnit. The parameter InstanceChargeType takes effect only when the value is PrePaid and is required. Once the DedicatedHostId is specified, the value cannot exceed the subscription duration of the dedicated host. When PeriodUnit = Week, Period values: {"1", "2", "3", "4"}. When PeriodUnit = Month, Period values: {"1", "2", "3", "4", "5", "6", "7", "8", "9", "12", "24", "36", "48", "60"}.
        /// </summary>
        [Output("period")]
        public Output<int?> Period { get; private set; } = null!;

        /// <summary>
        /// The duration of the purchase of resources.Package year and Month value range: Month.Default value: Month of the package, which is billed by volume. The default period is Hour.
        /// </summary>
        [Output("periodUnit")]
        public Output<string?> PeriodUnit { get; private set; } = null!;

        /// <summary>
        /// The ID of the enterprise resource group to which the instance resides.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// Series of OceanBase cluster instances-normal (default): Standard cluster version (cloud disk)-normal_SSD: Standard cluster version (local disk)-history: history Library cluster version.
        /// </summary>
        [Output("series")]
        public Output<string> Series { get; private set; } = null!;

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Information about the zone where the cluster is deployed.
        /// </summary>
        [Output("zones")]
        public Output<ImmutableArray<string>> Zones { get; private set; } = null!;


        /// <summary>
        /// Create a BaseInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BaseInstance(string name, BaseInstanceArgs args, CustomResourceOptions? options = null)
            : base("alicloud:ocean/baseInstance:BaseInstance", name, args ?? new BaseInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BaseInstance(string name, Input<string> id, BaseInstanceState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ocean/baseInstance:BaseInstance", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing BaseInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BaseInstance Get(string name, Input<string> id, BaseInstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new BaseInstance(name, id, state, options);
        }
    }

    public sealed class BaseInstanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to automatically renew.
        /// It takes effect when the parameter ChargeType is PrePaid. Value range:
        /// - true: automatic renewal.
        /// - false (default): no automatic renewal.
        /// </summary>
        [Input("autoRenew")]
        public Input<bool>? AutoRenew { get; set; }

        /// <summary>
        /// The duration of each auto-renewal. When the value of the AutoRenew parameter is True, this parameter is required.
        /// - PeriodUnit is Week, AutoRenewPeriod is {"1", "2", "3"}.
        /// - PeriodUnit is Month, AutoRenewPeriod is {"1", "2", "3", "6", "12"}.
        /// </summary>
        [Input("autoRenewPeriod")]
        public Input<int>? AutoRenewPeriod { get; set; }

        /// <summary>
        /// The backup retention policy after the cluster is deleted. The values are as follows:
        /// - receive_all: Keep all backup sets;
        /// - delete_all: delete all backup sets;
        /// - receive_last: Keep the last backup set.
        /// &gt; **NOTE:**   The default value is delete_all.
        /// </summary>
        [Input("backupRetainMode")]
        public Input<string>? BackupRetainMode { get; set; }

        /// <summary>
        /// The size of the storage space, in GB.
        /// The limits of storage space vary according to the cluster specifications, as follows:
        /// - 8C32GB:100GB ~ 10000GB
        /// - 14C70GB:200GB ~ 10000GB
        /// - 30C180GB:400GB ~ 10000GB
        /// - 62C400G:800GB ~ 10000GB.
        /// The default value of each package is its minimum value.
        /// </summary>
        [Input("diskSize", required: true)]
        public Input<int> DiskSize { get; set; } = null!;

        /// <summary>
        /// The storage type of the cluster. Effective only in the standard cluster version (cloud disk).
        /// Two types are currently supported:
        /// - cloud_essd_pl1: cloud disk ESSD pl1.
        /// - cloud_essd_pl0: cloud disk ESSD pl0. The default value is cloud_essd_pl1.
        /// </summary>
        [Input("diskType")]
        public Input<string>? DiskType { get; set; }

        /// <summary>
        /// Cluster specification information.
        /// Four packages are currently supported:
        /// - 4C16GB：4cores 16GB
        /// - 8C32GB：8cores 32GB
        /// - 14C70GB：14cores 70GB
        /// - 24C120GB：24cores 120GB
        /// - 30C180GB：30cores 180GB
        /// - 62C400GB：62cores 400GB
        /// - 104C600GB：104cores 600GB
        /// - 16C70GB：16cores 70GB
        /// - 32C160GB：32cores 160GB
        /// - 64C380GB：64cores 380GB
        /// - 20C32GB：20cores 32GB
        /// - 40C64GB：40cores 64GB
        /// - 16C32GB：16cores 32GB
        /// - 32C70GB：32cores 70GB
        /// - 64C180GB：64cores 180GB
        /// - 32C180GB：32cores 180GB
        /// - 64C400GB：64cores 400GB.
        /// </summary>
        [Input("instanceClass", required: true)]
        public Input<string> InstanceClass { get; set; } = null!;

        /// <summary>
        /// OceanBase cluster name.The length is 1 to 20 English or Chinese characters.If this parameter is not specified, the default value is the InstanceId of the cluster.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        /// <summary>
        /// The number of nodes in the cluster. If the deployment mode is n-n-n, the number of nodes is n * 3.
        /// </summary>
        [Input("nodeNum")]
        public Input<string>? NodeNum { get; set; }

        /// <summary>
        /// The OceanBase Server version number.
        /// </summary>
        [Input("obVersion")]
        public Input<string>? ObVersion { get; set; }

        /// <summary>
        /// The payment method of the instance. Value range:
        /// - Subscription: Package year and month. When you select this type of payment method, you must make sure that your account supports balance payment or credit payment. Otherwise, an InvalidPayMethod error message will be returned.
        /// - PayAsYouGo (default): Pay-as-you-go (default hourly billing).
        /// </summary>
        [Input("paymentType", required: true)]
        public Input<string> PaymentType { get; set; } = null!;

        /// <summary>
        /// The duration of the resource purchase. The unit is specified by the PeriodUnit. The parameter InstanceChargeType takes effect only when the value is PrePaid and is required. Once the DedicatedHostId is specified, the value cannot exceed the subscription duration of the dedicated host. When PeriodUnit = Week, Period values: {"1", "2", "3", "4"}. When PeriodUnit = Month, Period values: {"1", "2", "3", "4", "5", "6", "7", "8", "9", "12", "24", "36", "48", "60"}.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// The duration of the purchase of resources.Package year and Month value range: Month.Default value: Month of the package, which is billed by volume. The default period is Hour.
        /// </summary>
        [Input("periodUnit")]
        public Input<string>? PeriodUnit { get; set; }

        /// <summary>
        /// The ID of the enterprise resource group to which the instance resides.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// Series of OceanBase cluster instances-normal (default): Standard cluster version (cloud disk)-normal_SSD: Standard cluster version (local disk)-history: history Library cluster version.
        /// </summary>
        [Input("series", required: true)]
        public Input<string> Series { get; set; } = null!;

        [Input("zones", required: true)]
        private InputList<string>? _zones;

        /// <summary>
        /// Information about the zone where the cluster is deployed.
        /// </summary>
        public InputList<string> Zones
        {
            get => _zones ?? (_zones = new InputList<string>());
            set => _zones = value;
        }

        public BaseInstanceArgs()
        {
        }
        public static new BaseInstanceArgs Empty => new BaseInstanceArgs();
    }

    public sealed class BaseInstanceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to automatically renew.
        /// It takes effect when the parameter ChargeType is PrePaid. Value range:
        /// - true: automatic renewal.
        /// - false (default): no automatic renewal.
        /// </summary>
        [Input("autoRenew")]
        public Input<bool>? AutoRenew { get; set; }

        /// <summary>
        /// The duration of each auto-renewal. When the value of the AutoRenew parameter is True, this parameter is required.
        /// - PeriodUnit is Week, AutoRenewPeriod is {"1", "2", "3"}.
        /// - PeriodUnit is Month, AutoRenewPeriod is {"1", "2", "3", "6", "12"}.
        /// </summary>
        [Input("autoRenewPeriod")]
        public Input<int>? AutoRenewPeriod { get; set; }

        /// <summary>
        /// The backup retention policy after the cluster is deleted. The values are as follows:
        /// - receive_all: Keep all backup sets;
        /// - delete_all: delete all backup sets;
        /// - receive_last: Keep the last backup set.
        /// &gt; **NOTE:**   The default value is delete_all.
        /// </summary>
        [Input("backupRetainMode")]
        public Input<string>? BackupRetainMode { get; set; }

        /// <summary>
        /// The product code of the OceanBase cluster._oceanbasepre_public_cn: Domestic station cloud database package Year-to-month package._oceanbasepost_public_cn: The domestic station cloud database is paid by the hour._obpre_public_intl: International Station Cloud Database Package Monthly Package.
        /// </summary>
        [Input("commodityCode")]
        public Input<string>? CommodityCode { get; set; }

        /// <summary>
        /// The number of CPU cores of the cluster.
        /// </summary>
        [Input("cpu")]
        public Input<int>? Cpu { get; set; }

        /// <summary>
        /// The creation time of the resource.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The size of the storage space, in GB.
        /// The limits of storage space vary according to the cluster specifications, as follows:
        /// - 8C32GB:100GB ~ 10000GB
        /// - 14C70GB:200GB ~ 10000GB
        /// - 30C180GB:400GB ~ 10000GB
        /// - 62C400G:800GB ~ 10000GB.
        /// The default value of each package is its minimum value.
        /// </summary>
        [Input("diskSize")]
        public Input<int>? DiskSize { get; set; }

        /// <summary>
        /// The storage type of the cluster. Effective only in the standard cluster version (cloud disk).
        /// Two types are currently supported:
        /// - cloud_essd_pl1: cloud disk ESSD pl1.
        /// - cloud_essd_pl0: cloud disk ESSD pl0. The default value is cloud_essd_pl1.
        /// </summary>
        [Input("diskType")]
        public Input<string>? DiskType { get; set; }

        /// <summary>
        /// Cluster specification information.
        /// Four packages are currently supported:
        /// - 4C16GB：4cores 16GB
        /// - 8C32GB：8cores 32GB
        /// - 14C70GB：14cores 70GB
        /// - 24C120GB：24cores 120GB
        /// - 30C180GB：30cores 180GB
        /// - 62C400GB：62cores 400GB
        /// - 104C600GB：104cores 600GB
        /// - 16C70GB：16cores 70GB
        /// - 32C160GB：32cores 160GB
        /// - 64C380GB：64cores 380GB
        /// - 20C32GB：20cores 32GB
        /// - 40C64GB：40cores 64GB
        /// - 16C32GB：16cores 32GB
        /// - 32C70GB：32cores 70GB
        /// - 64C180GB：64cores 180GB
        /// - 32C180GB：32cores 180GB
        /// - 64C400GB：64cores 400GB.
        /// </summary>
        [Input("instanceClass")]
        public Input<string>? InstanceClass { get; set; }

        /// <summary>
        /// OceanBase cluster name.The length is 1 to 20 English or Chinese characters.If this parameter is not specified, the default value is the InstanceId of the cluster.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        /// <summary>
        /// The number of nodes in the cluster. If the deployment mode is n-n-n, the number of nodes is n * 3.
        /// </summary>
        [Input("nodeNum")]
        public Input<string>? NodeNum { get; set; }

        /// <summary>
        /// The OceanBase Server version number.
        /// </summary>
        [Input("obVersion")]
        public Input<string>? ObVersion { get; set; }

        /// <summary>
        /// The payment method of the instance. Value range:
        /// - Subscription: Package year and month. When you select this type of payment method, you must make sure that your account supports balance payment or credit payment. Otherwise, an InvalidPayMethod error message will be returned.
        /// - PayAsYouGo (default): Pay-as-you-go (default hourly billing).
        /// </summary>
        [Input("paymentType")]
        public Input<string>? PaymentType { get; set; }

        /// <summary>
        /// The duration of the resource purchase. The unit is specified by the PeriodUnit. The parameter InstanceChargeType takes effect only when the value is PrePaid and is required. Once the DedicatedHostId is specified, the value cannot exceed the subscription duration of the dedicated host. When PeriodUnit = Week, Period values: {"1", "2", "3", "4"}. When PeriodUnit = Month, Period values: {"1", "2", "3", "4", "5", "6", "7", "8", "9", "12", "24", "36", "48", "60"}.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// The duration of the purchase of resources.Package year and Month value range: Month.Default value: Month of the package, which is billed by volume. The default period is Hour.
        /// </summary>
        [Input("periodUnit")]
        public Input<string>? PeriodUnit { get; set; }

        /// <summary>
        /// The ID of the enterprise resource group to which the instance resides.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// Series of OceanBase cluster instances-normal (default): Standard cluster version (cloud disk)-normal_SSD: Standard cluster version (local disk)-history: history Library cluster version.
        /// </summary>
        [Input("series")]
        public Input<string>? Series { get; set; }

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("zones")]
        private InputList<string>? _zones;

        /// <summary>
        /// Information about the zone where the cluster is deployed.
        /// </summary>
        public InputList<string> Zones
        {
            get => _zones ?? (_zones = new InputList<string>());
            set => _zones = value;
        }

        public BaseInstanceState()
        {
        }
        public static new BaseInstanceState Empty => new BaseInstanceState();
    }
}

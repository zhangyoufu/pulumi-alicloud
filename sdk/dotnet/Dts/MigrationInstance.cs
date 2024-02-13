// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Dts
{
    /// <summary>
    /// Provides a DTS Migration Instance resource.
    /// 
    /// For information about DTS Migration Instance and how to use it, see [What is Synchronization Instance](https://www.alibabacloud.com/help/en/doc-detail/208270.html).
    /// 
    /// &gt; **NOTE:** Available since v1.157.0.
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
    ///     var defaultRegions = AliCloud.GetRegions.Invoke(new()
    ///     {
    ///         Current = true,
    ///     });
    /// 
    ///     var defaultMigrationInstance = new AliCloud.Dts.MigrationInstance("defaultMigrationInstance", new()
    ///     {
    ///         PaymentType = "PayAsYouGo",
    ///         SourceEndpointEngineName = "MySQL",
    ///         SourceEndpointRegion = defaultRegions.Apply(getRegionsResult =&gt; getRegionsResult.Regions[0]?.Id),
    ///         DestinationEndpointEngineName = "MySQL",
    ///         DestinationEndpointRegion = defaultRegions.Apply(getRegionsResult =&gt; getRegionsResult.Regions[0]?.Id),
    ///         InstanceClass = "small",
    ///         SyncArchitecture = "oneway",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// DTS Migration Instance can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:dts/migrationInstance:MigrationInstance example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:dts/migrationInstance:MigrationInstance")]
    public partial class MigrationInstance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// [ETL specifications](https://help.aliyun.com/document_detail/212324.html). The unit is the computing unit ComputeUnit (CU), 1CU=1vCPU+4 GB memory. The value range is an integer greater than or equal to 2.
        /// </summary>
        [Output("computeUnit")]
        public Output<int?> ComputeUnit { get; private set; } = null!;

        /// <summary>
        /// The number of private customized RDS instances under PolarDB-X. The default value is 1. This parameter needs to be passed only when `source_endpoint_engine_name` equals `drds`.
        /// </summary>
        [Output("databaseCount")]
        public Output<int?> DatabaseCount { get; private set; } = null!;

        /// <summary>
        /// The type of destination engine. Valid values: `ADS`, `DB2`, `DRDS`, `DataHub`, `Greenplum`, `MSSQL`, `MySQL`, `PolarDB`, `PostgreSQL`, `Redis`, `Tablestore`, `as400`, `clickhouse`, `kafka`, `mongodb`, `odps`, `oracle`, `polardb_o`, `polardb_pg`, `tidb`. For the correspondence between the supported source and target libraries, see [Supported Databases, Synchronization Initialization Types and Synchronization Topologies](https://help.aliyun.com/document_detail/130744.html), [Supported Databases and Migration Types](https://help.aliyun.com/document_detail/26618.html).
        /// </summary>
        [Output("destinationEndpointEngineName")]
        public Output<string> DestinationEndpointEngineName { get; private set; } = null!;

        /// <summary>
        /// The region of destination instance. List of [supported regions](https://help.aliyun.com/document_detail/141033.html).
        /// </summary>
        [Output("destinationEndpointRegion")]
        public Output<string> DestinationEndpointRegion { get; private set; } = null!;

        /// <summary>
        /// The ID of the Migration Instance.
        /// </summary>
        [Output("dtsInstanceId")]
        public Output<string> DtsInstanceId { get; private set; } = null!;

        /// <summary>
        /// The instance class. Valid values: `large`, `medium`, `small`, `xlarge`, `xxlarge`. You can only upgrade the configuration, not downgrade the configuration. If you downgrade the instance, you need to [submit a ticket](https://selfservice.console.aliyun.com/ticket/category/dts/today).
        /// </summary>
        [Output("instanceClass")]
        public Output<string> InstanceClass { get; private set; } = null!;

        /// <summary>
        /// The payment type of the resource. Valid values: `PayAsYouGo`.
        /// </summary>
        [Output("paymentType")]
        public Output<string> PaymentType { get; private set; } = null!;

        /// <summary>
        /// The type of source endpoint engine. Valid values: `ADS`, `DB2`, `DRDS`, `DataHub`, `Greenplum`, `MSSQL`, `MySQL`, `PolarDB`, `PostgreSQL`, `Redis`, `Tablestore`, `as400`, `clickhouse`, `kafka`, `mongodb`, `odps`, `oracle`, `polardb_o`, `polardb_pg`, `tidb`. For the correspondence between the supported source and target libraries, see [Supported Databases, Synchronization Initialization Types and Synchronization Topologies](https://help.aliyun.com/document_detail/130744.html), [Supported Databases and Migration Types](https://help.aliyun.com/document_detail/26618.html).
        /// </summary>
        [Output("sourceEndpointEngineName")]
        public Output<string> SourceEndpointEngineName { get; private set; } = null!;

        /// <summary>
        /// The region of source instance.
        /// </summary>
        [Output("sourceEndpointRegion")]
        public Output<string> SourceEndpointRegion { get; private set; } = null!;

        /// <summary>
        /// The status.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The sync architecture. Valid values: `oneway`.
        /// </summary>
        [Output("syncArchitecture")]
        public Output<string?> SyncArchitecture { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a MigrationInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MigrationInstance(string name, MigrationInstanceArgs args, CustomResourceOptions? options = null)
            : base("alicloud:dts/migrationInstance:MigrationInstance", name, args ?? new MigrationInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MigrationInstance(string name, Input<string> id, MigrationInstanceState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:dts/migrationInstance:MigrationInstance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MigrationInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MigrationInstance Get(string name, Input<string> id, MigrationInstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new MigrationInstance(name, id, state, options);
        }
    }

    public sealed class MigrationInstanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// [ETL specifications](https://help.aliyun.com/document_detail/212324.html). The unit is the computing unit ComputeUnit (CU), 1CU=1vCPU+4 GB memory. The value range is an integer greater than or equal to 2.
        /// </summary>
        [Input("computeUnit")]
        public Input<int>? ComputeUnit { get; set; }

        /// <summary>
        /// The number of private customized RDS instances under PolarDB-X. The default value is 1. This parameter needs to be passed only when `source_endpoint_engine_name` equals `drds`.
        /// </summary>
        [Input("databaseCount")]
        public Input<int>? DatabaseCount { get; set; }

        /// <summary>
        /// The type of destination engine. Valid values: `ADS`, `DB2`, `DRDS`, `DataHub`, `Greenplum`, `MSSQL`, `MySQL`, `PolarDB`, `PostgreSQL`, `Redis`, `Tablestore`, `as400`, `clickhouse`, `kafka`, `mongodb`, `odps`, `oracle`, `polardb_o`, `polardb_pg`, `tidb`. For the correspondence between the supported source and target libraries, see [Supported Databases, Synchronization Initialization Types and Synchronization Topologies](https://help.aliyun.com/document_detail/130744.html), [Supported Databases and Migration Types](https://help.aliyun.com/document_detail/26618.html).
        /// </summary>
        [Input("destinationEndpointEngineName", required: true)]
        public Input<string> DestinationEndpointEngineName { get; set; } = null!;

        /// <summary>
        /// The region of destination instance. List of [supported regions](https://help.aliyun.com/document_detail/141033.html).
        /// </summary>
        [Input("destinationEndpointRegion", required: true)]
        public Input<string> DestinationEndpointRegion { get; set; } = null!;

        /// <summary>
        /// The instance class. Valid values: `large`, `medium`, `small`, `xlarge`, `xxlarge`. You can only upgrade the configuration, not downgrade the configuration. If you downgrade the instance, you need to [submit a ticket](https://selfservice.console.aliyun.com/ticket/category/dts/today).
        /// </summary>
        [Input("instanceClass")]
        public Input<string>? InstanceClass { get; set; }

        /// <summary>
        /// The payment type of the resource. Valid values: `PayAsYouGo`.
        /// </summary>
        [Input("paymentType", required: true)]
        public Input<string> PaymentType { get; set; } = null!;

        /// <summary>
        /// The type of source endpoint engine. Valid values: `ADS`, `DB2`, `DRDS`, `DataHub`, `Greenplum`, `MSSQL`, `MySQL`, `PolarDB`, `PostgreSQL`, `Redis`, `Tablestore`, `as400`, `clickhouse`, `kafka`, `mongodb`, `odps`, `oracle`, `polardb_o`, `polardb_pg`, `tidb`. For the correspondence between the supported source and target libraries, see [Supported Databases, Synchronization Initialization Types and Synchronization Topologies](https://help.aliyun.com/document_detail/130744.html), [Supported Databases and Migration Types](https://help.aliyun.com/document_detail/26618.html).
        /// </summary>
        [Input("sourceEndpointEngineName", required: true)]
        public Input<string> SourceEndpointEngineName { get; set; } = null!;

        /// <summary>
        /// The region of source instance.
        /// </summary>
        [Input("sourceEndpointRegion", required: true)]
        public Input<string> SourceEndpointRegion { get; set; } = null!;

        /// <summary>
        /// The sync architecture. Valid values: `oneway`.
        /// </summary>
        [Input("syncArchitecture")]
        public Input<string>? SyncArchitecture { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public MigrationInstanceArgs()
        {
        }
        public static new MigrationInstanceArgs Empty => new MigrationInstanceArgs();
    }

    public sealed class MigrationInstanceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// [ETL specifications](https://help.aliyun.com/document_detail/212324.html). The unit is the computing unit ComputeUnit (CU), 1CU=1vCPU+4 GB memory. The value range is an integer greater than or equal to 2.
        /// </summary>
        [Input("computeUnit")]
        public Input<int>? ComputeUnit { get; set; }

        /// <summary>
        /// The number of private customized RDS instances under PolarDB-X. The default value is 1. This parameter needs to be passed only when `source_endpoint_engine_name` equals `drds`.
        /// </summary>
        [Input("databaseCount")]
        public Input<int>? DatabaseCount { get; set; }

        /// <summary>
        /// The type of destination engine. Valid values: `ADS`, `DB2`, `DRDS`, `DataHub`, `Greenplum`, `MSSQL`, `MySQL`, `PolarDB`, `PostgreSQL`, `Redis`, `Tablestore`, `as400`, `clickhouse`, `kafka`, `mongodb`, `odps`, `oracle`, `polardb_o`, `polardb_pg`, `tidb`. For the correspondence between the supported source and target libraries, see [Supported Databases, Synchronization Initialization Types and Synchronization Topologies](https://help.aliyun.com/document_detail/130744.html), [Supported Databases and Migration Types](https://help.aliyun.com/document_detail/26618.html).
        /// </summary>
        [Input("destinationEndpointEngineName")]
        public Input<string>? DestinationEndpointEngineName { get; set; }

        /// <summary>
        /// The region of destination instance. List of [supported regions](https://help.aliyun.com/document_detail/141033.html).
        /// </summary>
        [Input("destinationEndpointRegion")]
        public Input<string>? DestinationEndpointRegion { get; set; }

        /// <summary>
        /// The ID of the Migration Instance.
        /// </summary>
        [Input("dtsInstanceId")]
        public Input<string>? DtsInstanceId { get; set; }

        /// <summary>
        /// The instance class. Valid values: `large`, `medium`, `small`, `xlarge`, `xxlarge`. You can only upgrade the configuration, not downgrade the configuration. If you downgrade the instance, you need to [submit a ticket](https://selfservice.console.aliyun.com/ticket/category/dts/today).
        /// </summary>
        [Input("instanceClass")]
        public Input<string>? InstanceClass { get; set; }

        /// <summary>
        /// The payment type of the resource. Valid values: `PayAsYouGo`.
        /// </summary>
        [Input("paymentType")]
        public Input<string>? PaymentType { get; set; }

        /// <summary>
        /// The type of source endpoint engine. Valid values: `ADS`, `DB2`, `DRDS`, `DataHub`, `Greenplum`, `MSSQL`, `MySQL`, `PolarDB`, `PostgreSQL`, `Redis`, `Tablestore`, `as400`, `clickhouse`, `kafka`, `mongodb`, `odps`, `oracle`, `polardb_o`, `polardb_pg`, `tidb`. For the correspondence between the supported source and target libraries, see [Supported Databases, Synchronization Initialization Types and Synchronization Topologies](https://help.aliyun.com/document_detail/130744.html), [Supported Databases and Migration Types](https://help.aliyun.com/document_detail/26618.html).
        /// </summary>
        [Input("sourceEndpointEngineName")]
        public Input<string>? SourceEndpointEngineName { get; set; }

        /// <summary>
        /// The region of source instance.
        /// </summary>
        [Input("sourceEndpointRegion")]
        public Input<string>? SourceEndpointRegion { get; set; }

        /// <summary>
        /// The status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The sync architecture. Valid values: `oneway`.
        /// </summary>
        [Input("syncArchitecture")]
        public Input<string>? SyncArchitecture { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public MigrationInstanceState()
        {
        }
        public static new MigrationInstanceState Empty => new MigrationInstanceState();
    }
}

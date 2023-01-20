// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.KVStore
{
    public static class GetInstanceClasses
    {
        /// <summary>
        /// This data source provides the KVStore instance classes resource available info of Alibaba Cloud.
        /// 
        /// &gt; **NOTE:** Available in v1.49.0+
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var resourcesZones = AliCloud.GetZones.Invoke(new()
        ///     {
        ///         AvailableResourceCreation = "KVStore",
        ///     });
        /// 
        ///     var resourcesInstanceClasses = AliCloud.KVStore.GetInstanceClasses.Invoke(new()
        ///     {
        ///         Engine = "Redis",
        ///         EngineVersion = "5.0",
        ///         InstanceChargeType = "PrePaid",
        ///         OutputFile = "./classes.txt",
        ///         ZoneId = resourcesZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstKvstoreInstanceClass"] = resourcesInstanceClasses.Apply(getInstanceClassesResult =&gt; getInstanceClassesResult.InstanceClasses),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetInstanceClassesResult> InvokeAsync(GetInstanceClassesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInstanceClassesResult>("alicloud:kvstore/getInstanceClasses:getInstanceClasses", args ?? new GetInstanceClassesArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the KVStore instance classes resource available info of Alibaba Cloud.
        /// 
        /// &gt; **NOTE:** Available in v1.49.0+
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var resourcesZones = AliCloud.GetZones.Invoke(new()
        ///     {
        ///         AvailableResourceCreation = "KVStore",
        ///     });
        /// 
        ///     var resourcesInstanceClasses = AliCloud.KVStore.GetInstanceClasses.Invoke(new()
        ///     {
        ///         Engine = "Redis",
        ///         EngineVersion = "5.0",
        ///         InstanceChargeType = "PrePaid",
        ///         OutputFile = "./classes.txt",
        ///         ZoneId = resourcesZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstKvstoreInstanceClass"] = resourcesInstanceClasses.Apply(getInstanceClassesResult =&gt; getInstanceClassesResult.InstanceClasses),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetInstanceClassesResult> Invoke(GetInstanceClassesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstanceClassesResult>("alicloud:kvstore/getInstanceClasses:getInstanceClasses", args ?? new GetInstanceClassesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceClassesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The KVStore instance system architecture required by the user. Valid values: `standard`, `cluster` and `rwsplit`.
        /// </summary>
        [Input("architecture")]
        public string? Architecture { get; set; }

        /// <summary>
        /// The KVStore instance edition type required by the user. Valid values: `Community` and `Enterprise`.
        /// </summary>
        [Input("editionType")]
        public string? EditionType { get; set; }

        /// <summary>
        /// Database type. Options are `Redis`, `Memcache`. Default to `Redis`.
        /// </summary>
        [Input("engine")]
        public string? Engine { get; set; }

        /// <summary>
        /// Database version required by the user. Value options of Redis can refer to the latest docs [detail info](https://www.alibabacloud.com/help/doc-detail/60873.htm) `EngineVersion`. Value of Memcache should be empty.
        /// </summary>
        [Input("engineVersion")]
        public string? EngineVersion { get; set; }

        /// <summary>
        /// Filter the results by charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PrePaid`.
        /// </summary>
        [Input("instanceChargeType")]
        public string? InstanceChargeType { get; set; }

        /// <summary>
        /// The KVStore instance node type required by the user. Valid values: `double`, `single`, `readone`, `readthree` and `readfive`.
        /// </summary>
        [Input("nodeType")]
        public string? NodeType { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// It has been deprecated from 1.68.0.
        /// </summary>
        [Input("packageType")]
        public string? PackageType { get; set; }

        /// <summary>
        /// It has been deprecated from 1.68.0.
        /// </summary>
        [Input("performanceType")]
        public string? PerformanceType { get; set; }

        [Input("productType")]
        public string? ProductType { get; set; }

        /// <summary>
        /// The KVStore instance series type required by the user. Valid values: `enhanced_performance_type` and `hybrid_storage`.
        /// </summary>
        [Input("seriesType")]
        public string? SeriesType { get; set; }

        /// <summary>
        /// The number of shard.Valid values: `1`, `2`, `4`, `8`, `16`, `32`, `64`, `128`, `256`.
        /// * product_type - (Optional, Available in v1.130.0+) The type of the service. Valid values:
        /// * Local: an ApsaraDB for Redis instance with a local disk.
        /// * OnECS: an ApsaraDB for Redis instance with a standard disk. This type is available only on the Alibaba Cloud China site.
        /// </summary>
        [Input("shardNumber")]
        public int? ShardNumber { get; set; }

        [Input("sortedBy")]
        public string? SortedBy { get; set; }

        /// <summary>
        /// It has been deprecated from 1.68.0.
        /// </summary>
        [Input("storageType")]
        public string? StorageType { get; set; }

        /// <summary>
        /// The Zone to launch the KVStore instance.
        /// </summary>
        [Input("zoneId", required: true)]
        public string ZoneId { get; set; } = null!;

        public GetInstanceClassesArgs()
        {
        }
        public static new GetInstanceClassesArgs Empty => new GetInstanceClassesArgs();
    }

    public sealed class GetInstanceClassesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The KVStore instance system architecture required by the user. Valid values: `standard`, `cluster` and `rwsplit`.
        /// </summary>
        [Input("architecture")]
        public Input<string>? Architecture { get; set; }

        /// <summary>
        /// The KVStore instance edition type required by the user. Valid values: `Community` and `Enterprise`.
        /// </summary>
        [Input("editionType")]
        public Input<string>? EditionType { get; set; }

        /// <summary>
        /// Database type. Options are `Redis`, `Memcache`. Default to `Redis`.
        /// </summary>
        [Input("engine")]
        public Input<string>? Engine { get; set; }

        /// <summary>
        /// Database version required by the user. Value options of Redis can refer to the latest docs [detail info](https://www.alibabacloud.com/help/doc-detail/60873.htm) `EngineVersion`. Value of Memcache should be empty.
        /// </summary>
        [Input("engineVersion")]
        public Input<string>? EngineVersion { get; set; }

        /// <summary>
        /// Filter the results by charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PrePaid`.
        /// </summary>
        [Input("instanceChargeType")]
        public Input<string>? InstanceChargeType { get; set; }

        /// <summary>
        /// The KVStore instance node type required by the user. Valid values: `double`, `single`, `readone`, `readthree` and `readfive`.
        /// </summary>
        [Input("nodeType")]
        public Input<string>? NodeType { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// It has been deprecated from 1.68.0.
        /// </summary>
        [Input("packageType")]
        public Input<string>? PackageType { get; set; }

        /// <summary>
        /// It has been deprecated from 1.68.0.
        /// </summary>
        [Input("performanceType")]
        public Input<string>? PerformanceType { get; set; }

        [Input("productType")]
        public Input<string>? ProductType { get; set; }

        /// <summary>
        /// The KVStore instance series type required by the user. Valid values: `enhanced_performance_type` and `hybrid_storage`.
        /// </summary>
        [Input("seriesType")]
        public Input<string>? SeriesType { get; set; }

        /// <summary>
        /// The number of shard.Valid values: `1`, `2`, `4`, `8`, `16`, `32`, `64`, `128`, `256`.
        /// * product_type - (Optional, Available in v1.130.0+) The type of the service. Valid values:
        /// * Local: an ApsaraDB for Redis instance with a local disk.
        /// * OnECS: an ApsaraDB for Redis instance with a standard disk. This type is available only on the Alibaba Cloud China site.
        /// </summary>
        [Input("shardNumber")]
        public Input<int>? ShardNumber { get; set; }

        [Input("sortedBy")]
        public Input<string>? SortedBy { get; set; }

        /// <summary>
        /// It has been deprecated from 1.68.0.
        /// </summary>
        [Input("storageType")]
        public Input<string>? StorageType { get; set; }

        /// <summary>
        /// The Zone to launch the KVStore instance.
        /// </summary>
        [Input("zoneId", required: true)]
        public Input<string> ZoneId { get; set; } = null!;

        public GetInstanceClassesInvokeArgs()
        {
        }
        public static new GetInstanceClassesInvokeArgs Empty => new GetInstanceClassesInvokeArgs();
    }


    [OutputType]
    public sealed class GetInstanceClassesResult
    {
        public readonly string? Architecture;
        /// <summary>
        /// A list of KVStore available instance classes when the `sorted_by` is "Price". include:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceClassesClassResult> Classes;
        public readonly string? EditionType;
        public readonly string? Engine;
        public readonly string? EngineVersion;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? InstanceChargeType;
        /// <summary>
        /// A list of KVStore available instance classes.
        /// </summary>
        public readonly ImmutableArray<string> InstanceClasses;
        public readonly string? NodeType;
        public readonly string? OutputFile;
        public readonly string? PackageType;
        public readonly string? PerformanceType;
        public readonly string? ProductType;
        public readonly string? SeriesType;
        public readonly int? ShardNumber;
        public readonly string? SortedBy;
        public readonly string? StorageType;
        public readonly string ZoneId;

        [OutputConstructor]
        private GetInstanceClassesResult(
            string? architecture,

            ImmutableArray<Outputs.GetInstanceClassesClassResult> classes,

            string? editionType,

            string? engine,

            string? engineVersion,

            string id,

            string? instanceChargeType,

            ImmutableArray<string> instanceClasses,

            string? nodeType,

            string? outputFile,

            string? packageType,

            string? performanceType,

            string? productType,

            string? seriesType,

            int? shardNumber,

            string? sortedBy,

            string? storageType,

            string zoneId)
        {
            Architecture = architecture;
            Classes = classes;
            EditionType = editionType;
            Engine = engine;
            EngineVersion = engineVersion;
            Id = id;
            InstanceChargeType = instanceChargeType;
            InstanceClasses = instanceClasses;
            NodeType = nodeType;
            OutputFile = outputFile;
            PackageType = packageType;
            PerformanceType = performanceType;
            ProductType = productType;
            SeriesType = seriesType;
            ShardNumber = shardNumber;
            SortedBy = sortedBy;
            StorageType = storageType;
            ZoneId = zoneId;
        }
    }
}

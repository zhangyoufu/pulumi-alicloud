// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Lindorm
{
    public static class GetInstances
    {
        /// <summary>
        /// This data source provides the Lindorm Instances of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.132.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var ids = AliCloud.Lindorm.GetInstances.Invoke();
        /// 
        ///     var nameRegex = AliCloud.Lindorm.GetInstances.Invoke(new()
        ///     {
        ///         NameRegex = "^my-Instance",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["lindormInstanceId1"] = ids.Apply(getInstancesResult =&gt; getInstancesResult.Instances[0]?.Id),
        ///         ["lindormInstanceId2"] = nameRegex.Apply(getInstancesResult =&gt; getInstancesResult.Instances[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetInstancesResult> InvokeAsync(GetInstancesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInstancesResult>("alicloud:lindorm/getInstances:getInstances", args ?? new GetInstancesArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Lindorm Instances of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.132.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var ids = AliCloud.Lindorm.GetInstances.Invoke();
        /// 
        ///     var nameRegex = AliCloud.Lindorm.GetInstances.Invoke(new()
        ///     {
        ///         NameRegex = "^my-Instance",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["lindormInstanceId1"] = ids.Apply(getInstancesResult =&gt; getInstancesResult.Instances[0]?.Id),
        ///         ["lindormInstanceId2"] = nameRegex.Apply(getInstancesResult =&gt; getInstancesResult.Instances[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetInstancesResult> Invoke(GetInstancesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstancesResult>("alicloud:lindorm/getInstances:getInstances", args ?? new GetInstancesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstancesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Default to `false`. Set it to `true` can output more details about resource attributes.
        /// </summary>
        [Input("enableDetails")]
        public bool? EnableDetails { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Instance IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Instance name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The query str, which can use `instance_name` keyword for fuzzy search.
        /// </summary>
        [Input("queryStr")]
        public string? QueryStr { get; set; }

        /// <summary>
        /// The status of Instance, enumerative: Valid values: `ACTIVATION`, `DELETED`, `CREATING`, `CLASS_CHANGING`, `LOCKED`, `INSTANCE_LEVEL_MODIFY`, `NET_MODIFYING`, `RESIZING`, `RESTARTING`, `MINOR_VERSION_TRANSING`.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        /// <summary>
        /// The support engine. Valid values: `1` to `7`.
        /// </summary>
        [Input("supportEngine")]
        public int? SupportEngine { get; set; }

        public GetInstancesArgs()
        {
        }
        public static new GetInstancesArgs Empty => new GetInstancesArgs();
    }

    public sealed class GetInstancesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Default to `false`. Set it to `true` can output more details about resource attributes.
        /// </summary>
        [Input("enableDetails")]
        public Input<bool>? EnableDetails { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Instance IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Instance name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The query str, which can use `instance_name` keyword for fuzzy search.
        /// </summary>
        [Input("queryStr")]
        public Input<string>? QueryStr { get; set; }

        /// <summary>
        /// The status of Instance, enumerative: Valid values: `ACTIVATION`, `DELETED`, `CREATING`, `CLASS_CHANGING`, `LOCKED`, `INSTANCE_LEVEL_MODIFY`, `NET_MODIFYING`, `RESIZING`, `RESTARTING`, `MINOR_VERSION_TRANSING`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The support engine. Valid values: `1` to `7`.
        /// </summary>
        [Input("supportEngine")]
        public Input<int>? SupportEngine { get; set; }

        public GetInstancesInvokeArgs()
        {
        }
        public static new GetInstancesInvokeArgs Empty => new GetInstancesInvokeArgs();
    }


    [OutputType]
    public sealed class GetInstancesResult
    {
        public readonly bool? EnableDetails;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly ImmutableArray<Outputs.GetInstancesInstanceResult> Instances;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly string? QueryStr;
        public readonly string? Status;
        public readonly int? SupportEngine;

        [OutputConstructor]
        private GetInstancesResult(
            bool? enableDetails,

            string id,

            ImmutableArray<string> ids,

            ImmutableArray<Outputs.GetInstancesInstanceResult> instances,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? queryStr,

            string? status,

            int? supportEngine)
        {
            EnableDetails = enableDetails;
            Id = id;
            Ids = ids;
            Instances = instances;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            QueryStr = queryStr;
            Status = status;
            SupportEngine = supportEngine;
        }
    }
}

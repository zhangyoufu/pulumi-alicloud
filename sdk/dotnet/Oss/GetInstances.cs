// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Oss
{
    [Obsolete(@"alicloud.oss.getInstances has been deprecated in favor of alicloud.ots.getInstances")]
    public static class GetInstances
    {
        /// <summary>
        /// This data source provides the ots instances of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available since v1.40.0.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var instancesDs = AliCloud.Ots.GetInstances.Invoke(new()
        ///     {
        ///         OutputFile = "instances.txt",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstInstanceId"] = instancesDs.Apply(getInstancesResult =&gt; getInstancesResult.Instances[0]?.Id),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetInstancesResult> InvokeAsync(GetInstancesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInstancesResult>("alicloud:oss/getInstances:getInstances", args ?? new GetInstancesArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the ots instances of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available since v1.40.0.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var instancesDs = AliCloud.Ots.GetInstances.Invoke(new()
        ///     {
        ///         OutputFile = "instances.txt",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstInstanceId"] = instancesDs.Apply(getInstancesResult =&gt; getInstancesResult.Instances[0]?.Id),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetInstancesResult> Invoke(GetInstancesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstancesResult>("alicloud:oss/getInstances:getInstances", args ?? new GetInstancesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstancesArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of instance IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by instance name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// A map of tags assigned to the instance. It must be in the format:
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var instancesDs = AliCloud.Ots.GetInstances.Invoke(new()
        ///     {
        ///         Tags = 
        ///         {
        ///             { "tagKey1", "tagValue1" },
        ///             { "tagKey2", "tagValue2" },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        public GetInstancesArgs()
        {
        }
        public static new GetInstancesArgs Empty => new GetInstancesArgs();
    }

    public sealed class GetInstancesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of instance IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by instance name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A map of tags assigned to the instance. It must be in the format:
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var instancesDs = AliCloud.Ots.GetInstances.Invoke(new()
        ///     {
        ///         Tags = 
        ///         {
        ///             { "tagKey1", "tagValue1" },
        ///             { "tagKey2", "tagValue2" },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public GetInstancesInvokeArgs()
        {
        }
        public static new GetInstancesInvokeArgs Empty => new GetInstancesInvokeArgs();
    }


    [OutputType]
    public sealed class GetInstancesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of instance IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// A list of instances. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstancesInstanceResult> Instances;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of instance names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        /// <summary>
        /// The tags of the instance.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Tags;

        [OutputConstructor]
        private GetInstancesResult(
            string id,

            ImmutableArray<string> ids,

            ImmutableArray<Outputs.GetInstancesInstanceResult> instances,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            ImmutableDictionary<string, object>? tags)
        {
            Id = id;
            Ids = ids;
            Instances = instances;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Tags = tags;
        }
    }
}

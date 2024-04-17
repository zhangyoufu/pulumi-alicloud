// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Drds
{
    public static class GetInstances
    {
        /// <summary>
        /// The `alicloud.drds.Instance` data source provides a collection of DRDS instances available in Alibaba Cloud account.
        /// Filters support regular expression for the instance name, searches by tags, and other filters which are listed below.
        /// 
        /// &gt; **NOTE:** Available in 1.35.0+.
        /// 
        /// ## Example Usage
        /// 
        ///  &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var drdsInstancesDs = AliCloud.Drds.GetInstances.Invoke(new()
        ///     {
        ///         NameRegex = "drds-\\d+",
        ///         Ids = new[]
        ///         {
        ///             "drdsabc123456",
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstDbInstanceId"] = drdsInstancesDs.Apply(getInstancesResult =&gt; getInstancesResult.Instances[0]?.Id),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetInstancesResult> InvokeAsync(GetInstancesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInstancesResult>("alicloud:drds/getInstances:getInstances", args ?? new GetInstancesArgs(), options.WithDefaults());

        /// <summary>
        /// The `alicloud.drds.Instance` data source provides a collection of DRDS instances available in Alibaba Cloud account.
        /// Filters support regular expression for the instance name, searches by tags, and other filters which are listed below.
        /// 
        /// &gt; **NOTE:** Available in 1.35.0+.
        /// 
        /// ## Example Usage
        /// 
        ///  &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var drdsInstancesDs = AliCloud.Drds.GetInstances.Invoke(new()
        ///     {
        ///         NameRegex = "drds-\\d+",
        ///         Ids = new[]
        ///         {
        ///             "drdsabc123456",
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstDbInstanceId"] = drdsInstancesDs.Apply(getInstancesResult =&gt; getInstancesResult.Instances[0]?.Id),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetInstancesResult> Invoke(GetInstancesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstancesResult>("alicloud:drds/getInstances:getInstances", args ?? new GetInstancesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstancesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A regex string to filter results by instance description.
        /// </summary>
        [Input("descriptionRegex")]
        public string? DescriptionRegex { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of DRDS instance IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by instance description. It is deprecated since v1.91.0 and will be removed in a future release, please use 'description_regex' instead.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetInstancesArgs()
        {
        }
        public static new GetInstancesArgs Empty => new GetInstancesArgs();
    }

    public sealed class GetInstancesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A regex string to filter results by instance description.
        /// </summary>
        [Input("descriptionRegex")]
        public Input<string>? DescriptionRegex { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of DRDS instance IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by instance description. It is deprecated since v1.91.0 and will be removed in a future release, please use 'description_regex' instead.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetInstancesInvokeArgs()
        {
        }
        public static new GetInstancesInvokeArgs Empty => new GetInstancesInvokeArgs();
    }


    [OutputType]
    public sealed class GetInstancesResult
    {
        public readonly string? DescriptionRegex;
        /// <summary>
        /// A list of DRDS descriptions.
        /// </summary>
        public readonly ImmutableArray<string> Descriptions;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of DRDS instance IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// A list of DRDS instances.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstancesInstanceResult> Instances;
        public readonly string? NameRegex;
        public readonly string? OutputFile;

        [OutputConstructor]
        private GetInstancesResult(
            string? descriptionRegex,

            ImmutableArray<string> descriptions,

            string id,

            ImmutableArray<string> ids,

            ImmutableArray<Outputs.GetInstancesInstanceResult> instances,

            string? nameRegex,

            string? outputFile)
        {
            DescriptionRegex = descriptionRegex;
            Descriptions = descriptions;
            Id = id;
            Ids = ids;
            Instances = instances;
            NameRegex = nameRegex;
            OutputFile = outputFile;
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Oos
{
    public static class GetStateConfigurations
    {
        /// <summary>
        /// This data source provides the Oos State Configurations of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.147.0+.
        /// 
        /// ## Example Usage
        /// 
        /// Basic Usage
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
        ///     var ids = AliCloud.Oos.GetStateConfigurations.Invoke();
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["oosStateConfigurationId1"] = ids.Apply(getStateConfigurationsResult =&gt; getStateConfigurationsResult.Configurations[0]?.Id),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetStateConfigurationsResult> InvokeAsync(GetStateConfigurationsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetStateConfigurationsResult>("alicloud:oos/getStateConfigurations:getStateConfigurations", args ?? new GetStateConfigurationsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Oos State Configurations of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.147.0+.
        /// 
        /// ## Example Usage
        /// 
        /// Basic Usage
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
        ///     var ids = AliCloud.Oos.GetStateConfigurations.Invoke();
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["oosStateConfigurationId1"] = ids.Apply(getStateConfigurationsResult =&gt; getStateConfigurationsResult.Configurations[0]?.Id),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetStateConfigurationsResult> Invoke(GetStateConfigurationsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetStateConfigurationsResult>("alicloud:oos/getStateConfigurations:getStateConfigurations", args ?? new GetStateConfigurationsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetStateConfigurationsArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of State Configuration IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public string? ResourceGroupId { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// The tag of the resource.
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        public GetStateConfigurationsArgs()
        {
        }
        public static new GetStateConfigurationsArgs Empty => new GetStateConfigurationsArgs();
    }

    public sealed class GetStateConfigurationsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of State Configuration IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The tag of the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public GetStateConfigurationsInvokeArgs()
        {
        }
        public static new GetStateConfigurationsInvokeArgs Empty => new GetStateConfigurationsInvokeArgs();
    }


    [OutputType]
    public sealed class GetStateConfigurationsResult
    {
        public readonly ImmutableArray<Outputs.GetStateConfigurationsConfigurationResult> Configurations;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? OutputFile;
        public readonly string? ResourceGroupId;
        public readonly ImmutableDictionary<string, object>? Tags;

        [OutputConstructor]
        private GetStateConfigurationsResult(
            ImmutableArray<Outputs.GetStateConfigurationsConfigurationResult> configurations,

            string id,

            ImmutableArray<string> ids,

            string? outputFile,

            string? resourceGroupId,

            ImmutableDictionary<string, object>? tags)
        {
            Configurations = configurations;
            Id = id;
            Ids = ids;
            OutputFile = outputFile;
            ResourceGroupId = resourceGroupId;
            Tags = tags;
        }
    }
}

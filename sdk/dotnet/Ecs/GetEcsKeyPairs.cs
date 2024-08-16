// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ecs
{
    public static class GetEcsKeyPairs
    {
        /// <summary>
        /// This data source provides the Ecs Key Pairs of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.121.0+.
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
        ///     var example = AliCloud.Ecs.GetEcsKeyPairs.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "key_pair_name",
        ///         },
        ///         NameRegex = "key_pair_name",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstEcsKeyPairId"] = example.Apply(getEcsKeyPairsResult =&gt; getEcsKeyPairsResult.Pairs[0]?.Id),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Task<GetEcsKeyPairsResult> InvokeAsync(GetEcsKeyPairsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetEcsKeyPairsResult>("alicloud:ecs/getEcsKeyPairs:getEcsKeyPairs", args ?? new GetEcsKeyPairsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Ecs Key Pairs of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.121.0+.
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
        ///     var example = AliCloud.Ecs.GetEcsKeyPairs.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "key_pair_name",
        ///         },
        ///         NameRegex = "key_pair_name",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstEcsKeyPairId"] = example.Apply(getEcsKeyPairsResult =&gt; getEcsKeyPairsResult.Pairs[0]?.Id),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetEcsKeyPairsResult> Invoke(GetEcsKeyPairsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetEcsKeyPairsResult>("alicloud:ecs/getEcsKeyPairs:getEcsKeyPairs", args ?? new GetEcsKeyPairsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEcsKeyPairsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The finger print of the key pair.
        /// </summary>
        [Input("fingerPrint")]
        public string? FingerPrint { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Key Pair IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Key Pair name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The resource group Id.
        /// </summary>
        [Input("resourceGroupId")]
        public string? ResourceGroupId { get; set; }

        [Input("tags")]
        private Dictionary<string, string>? _tags;
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetEcsKeyPairsArgs()
        {
        }
        public static new GetEcsKeyPairsArgs Empty => new GetEcsKeyPairsArgs();
    }

    public sealed class GetEcsKeyPairsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The finger print of the key pair.
        /// </summary>
        [Input("fingerPrint")]
        public Input<string>? FingerPrint { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Key Pair IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Key Pair name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The resource group Id.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public GetEcsKeyPairsInvokeArgs()
        {
        }
        public static new GetEcsKeyPairsInvokeArgs Empty => new GetEcsKeyPairsInvokeArgs();
    }


    [OutputType]
    public sealed class GetEcsKeyPairsResult
    {
        public readonly string? FingerPrint;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly ImmutableArray<Outputs.GetEcsKeyPairsKeyPairResult> KeyPairs;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly ImmutableArray<Outputs.GetEcsKeyPairsPairResult> Pairs;
        public readonly string? ResourceGroupId;
        public readonly ImmutableDictionary<string, string>? Tags;

        [OutputConstructor]
        private GetEcsKeyPairsResult(
            string? fingerPrint,

            string id,

            ImmutableArray<string> ids,

            ImmutableArray<Outputs.GetEcsKeyPairsKeyPairResult> keyPairs,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            ImmutableArray<Outputs.GetEcsKeyPairsPairResult> pairs,

            string? resourceGroupId,

            ImmutableDictionary<string, string>? tags)
        {
            FingerPrint = fingerPrint;
            Id = id;
            Ids = ids;
            KeyPairs = keyPairs;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Pairs = pairs;
            ResourceGroupId = resourceGroupId;
            Tags = tags;
        }
    }
}

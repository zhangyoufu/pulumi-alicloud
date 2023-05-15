// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CS
{
    public static class GetRegistryEnterpriseRepos
    {
        /// <summary>
        /// This data source provides a list Container Registry Enterprise Edition repositories on Alibaba Cloud.
        /// 
        /// &gt; **NOTE:** Available in v1.87.0+
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myRepos = AliCloud.CS.GetRegistryEnterpriseRepos.Invoke(new()
        ///     {
        ///         InstanceId = "cri-xx",
        ///         NameRegex = "my-repos",
        ///         OutputFile = "my-repo-json",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["output"] = myRepos.Apply(getRegistryEnterpriseReposResult =&gt; getRegistryEnterpriseReposResult.Repos),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRegistryEnterpriseReposResult> InvokeAsync(GetRegistryEnterpriseReposArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRegistryEnterpriseReposResult>("alicloud:cs/getRegistryEnterpriseRepos:getRegistryEnterpriseRepos", args ?? new GetRegistryEnterpriseReposArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides a list Container Registry Enterprise Edition repositories on Alibaba Cloud.
        /// 
        /// &gt; **NOTE:** Available in v1.87.0+
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myRepos = AliCloud.CS.GetRegistryEnterpriseRepos.Invoke(new()
        ///     {
        ///         InstanceId = "cri-xx",
        ///         NameRegex = "my-repos",
        ///         OutputFile = "my-repo-json",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["output"] = myRepos.Apply(getRegistryEnterpriseReposResult =&gt; getRegistryEnterpriseReposResult.Repos),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetRegistryEnterpriseReposResult> Invoke(GetRegistryEnterpriseReposInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRegistryEnterpriseReposResult>("alicloud:cs/getRegistryEnterpriseRepos:getRegistryEnterpriseRepos", args ?? new GetRegistryEnterpriseReposInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRegistryEnterpriseReposArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Boolean, false by default, only repository attributes are exported. Set to true if tags belong to this repository are needed. See `tags` in attributes.
        /// </summary>
        [Input("enableDetails")]
        public bool? EnableDetails { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of ids to filter results by repository id.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// ID of Container Registry Enterprise Edition instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// A regex string to filter results by repository name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// Name of Container Registry Enterprise Edition namespace where the repositories are located in.
        /// </summary>
        [Input("namespace")]
        public string? Namespace { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetRegistryEnterpriseReposArgs()
        {
        }
        public static new GetRegistryEnterpriseReposArgs Empty => new GetRegistryEnterpriseReposArgs();
    }

    public sealed class GetRegistryEnterpriseReposInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Boolean, false by default, only repository attributes are exported. Set to true if tags belong to this repository are needed. See `tags` in attributes.
        /// </summary>
        [Input("enableDetails")]
        public Input<bool>? EnableDetails { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of ids to filter results by repository id.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// ID of Container Registry Enterprise Edition instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// A regex string to filter results by repository name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// Name of Container Registry Enterprise Edition namespace where the repositories are located in.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetRegistryEnterpriseReposInvokeArgs()
        {
        }
        public static new GetRegistryEnterpriseReposInvokeArgs Empty => new GetRegistryEnterpriseReposInvokeArgs();
    }


    [OutputType]
    public sealed class GetRegistryEnterpriseReposResult
    {
        public readonly bool? EnableDetails;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of matched Container Registry Enterprise Edition repositories. Its element is a repository id.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// ID of Container Registry Enterprise Edition instance.
        /// </summary>
        public readonly string InstanceId;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of repository names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        /// <summary>
        /// Name of Container Registry Enterprise Edition namespace where repo is located.
        /// </summary>
        public readonly string? Namespace;
        public readonly string? OutputFile;
        /// <summary>
        /// A list of matched Container Registry Enterprise Edition namespaces. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRegistryEnterpriseReposRepoResult> Repos;

        [OutputConstructor]
        private GetRegistryEnterpriseReposResult(
            bool? enableDetails,

            string id,

            ImmutableArray<string> ids,

            string instanceId,

            string? nameRegex,

            ImmutableArray<string> names,

            string? @namespace,

            string? outputFile,

            ImmutableArray<Outputs.GetRegistryEnterpriseReposRepoResult> repos)
        {
            EnableDetails = enableDetails;
            Id = id;
            Ids = ids;
            InstanceId = instanceId;
            NameRegex = nameRegex;
            Names = names;
            Namespace = @namespace;
            OutputFile = outputFile;
            Repos = repos;
        }
    }
}

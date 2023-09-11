// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CR
{
    /// <summary>
    /// This resource will help you to manager Container Registry repositories, see [What is Repository](https://www.alibabacloud.com/help/en/acr/developer-reference/api-cr-2018-12-01-createrepository).
    /// 
    /// &gt; **NOTE:** Available since v1.35.0.
    /// 
    /// &gt; **NOTE:** You need to set your registry password in Container Registry console before use this resource.
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
    ///     var name = config.Get("name") ?? "tf-example";
    ///     var exampleNamespace = new AliCloud.CR.Namespace("exampleNamespace", new()
    ///     {
    ///         AutoCreate = false,
    ///         DefaultVisibility = "PUBLIC",
    ///     });
    /// 
    ///     var exampleRepo = new AliCloud.CR.Repo("exampleRepo", new()
    ///     {
    ///         Namespace = exampleNamespace.Name,
    ///         Summary = "this is summary of my new repo",
    ///         RepoType = "PUBLIC",
    ///         Detail = "this is a public repo",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Container Registry repository can be imported using the `namespace/repository`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:cr/repo:Repo default `my-namespace/my-repo`
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cr/repo:Repo")]
    public partial class Repo : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The repository specific information. MarkDown format is supported, and the length limit is 2000.
        /// </summary>
        [Output("detail")]
        public Output<string?> Detail { get; private set; } = null!;

        /// <summary>
        /// The repository domain list.
        /// </summary>
        [Output("domainList")]
        public Output<Outputs.RepoDomainList> DomainList { get; private set; } = null!;

        /// <summary>
        /// Name of container registry repository.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Name of container registry namespace where repository is located.
        /// </summary>
        [Output("namespace")]
        public Output<string> Namespace { get; private set; } = null!;

        /// <summary>
        /// `PUBLIC` or `PRIVATE`, repo's visibility.
        /// </summary>
        [Output("repoType")]
        public Output<string> RepoType { get; private set; } = null!;

        /// <summary>
        /// The repository general information. It can contain 1 to 80 characters.
        /// </summary>
        [Output("summary")]
        public Output<string> Summary { get; private set; } = null!;


        /// <summary>
        /// Create a Repo resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Repo(string name, RepoArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cr/repo:Repo", name, args ?? new RepoArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Repo(string name, Input<string> id, RepoState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cr/repo:Repo", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Repo resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Repo Get(string name, Input<string> id, RepoState? state = null, CustomResourceOptions? options = null)
        {
            return new Repo(name, id, state, options);
        }
    }

    public sealed class RepoArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The repository specific information. MarkDown format is supported, and the length limit is 2000.
        /// </summary>
        [Input("detail")]
        public Input<string>? Detail { get; set; }

        /// <summary>
        /// Name of container registry repository.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Name of container registry namespace where repository is located.
        /// </summary>
        [Input("namespace", required: true)]
        public Input<string> Namespace { get; set; } = null!;

        /// <summary>
        /// `PUBLIC` or `PRIVATE`, repo's visibility.
        /// </summary>
        [Input("repoType", required: true)]
        public Input<string> RepoType { get; set; } = null!;

        /// <summary>
        /// The repository general information. It can contain 1 to 80 characters.
        /// </summary>
        [Input("summary", required: true)]
        public Input<string> Summary { get; set; } = null!;

        public RepoArgs()
        {
        }
        public static new RepoArgs Empty => new RepoArgs();
    }

    public sealed class RepoState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The repository specific information. MarkDown format is supported, and the length limit is 2000.
        /// </summary>
        [Input("detail")]
        public Input<string>? Detail { get; set; }

        /// <summary>
        /// The repository domain list.
        /// </summary>
        [Input("domainList")]
        public Input<Inputs.RepoDomainListGetArgs>? DomainList { get; set; }

        /// <summary>
        /// Name of container registry repository.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Name of container registry namespace where repository is located.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// `PUBLIC` or `PRIVATE`, repo's visibility.
        /// </summary>
        [Input("repoType")]
        public Input<string>? RepoType { get; set; }

        /// <summary>
        /// The repository general information. It can contain 1 to 80 characters.
        /// </summary>
        [Input("summary")]
        public Input<string>? Summary { get; set; }

        public RepoState()
        {
        }
        public static new RepoState Empty => new RepoState();
    }
}

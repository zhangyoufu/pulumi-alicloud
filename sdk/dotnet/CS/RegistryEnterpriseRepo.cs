// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CS
{
    /// <summary>
    /// This resource will help you to manager Container Registry Enterprise Edition repositories.
    /// 
    /// For information about Container Registry Enterprise Edition repository and how to use it, see [Create a Repository](https://www.alibabacloud.com/help/en/acr/developer-reference/api-cr-2018-12-01-createrepository)
    /// 
    /// &gt; **NOTE:** Available since v1.86.0.
    /// 
    /// &gt; **NOTE:** You need to set your registry password in Container Registry Enterprise Edition console before use this resource.
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
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "terraform-example";
    ///     var exampleRegistryEnterpriseInstance = new AliCloud.CR.RegistryEnterpriseInstance("exampleRegistryEnterpriseInstance", new()
    ///     {
    ///         PaymentType = "Subscription",
    ///         Period = 1,
    ///         RenewPeriod = 0,
    ///         RenewalStatus = "ManualRenewal",
    ///         InstanceType = "Advanced",
    ///         InstanceName = name,
    ///     });
    /// 
    ///     var exampleRegistryEnterpriseNamespace = new AliCloud.CS.RegistryEnterpriseNamespace("exampleRegistryEnterpriseNamespace", new()
    ///     {
    ///         InstanceId = exampleRegistryEnterpriseInstance.Id,
    ///         AutoCreate = false,
    ///         DefaultVisibility = "PUBLIC",
    ///     });
    /// 
    ///     var exampleRegistryEnterpriseRepo = new AliCloud.CS.RegistryEnterpriseRepo("exampleRegistryEnterpriseRepo", new()
    ///     {
    ///         InstanceId = exampleRegistryEnterpriseInstance.Id,
    ///         Namespace = exampleRegistryEnterpriseNamespace.Name,
    ///         Summary = "this is summary of my new repo",
    ///         RepoType = "PUBLIC",
    ///         Detail = "this is a public repo",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Container Registry Enterprise Edition repository can be imported using the `{instance_id}:{namespace}:{repository}`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:cs/registryEnterpriseRepo:RegistryEnterpriseRepo default `cri-xxx:my-namespace:my-repo`
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cs/registryEnterpriseRepo:RegistryEnterpriseRepo")]
    public partial class RegistryEnterpriseRepo : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The repository specific information. MarkDown format is supported, and the length limit is 2000.
        /// </summary>
        [Output("detail")]
        public Output<string?> Detail { get; private set; } = null!;

        /// <summary>
        /// ID of Container Registry Enterprise Edition instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Name of Container Registry Enterprise Edition repository. It can contain 2 to 64 characters.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Name of Container Registry Enterprise Edition namespace where repository is located. It can contain 2 to 30 characters.
        /// </summary>
        [Output("namespace")]
        public Output<string> Namespace { get; private set; } = null!;

        /// <summary>
        /// The uuid of Container Registry Enterprise Edition repository.
        /// </summary>
        [Output("repoId")]
        public Output<string> RepoId { get; private set; } = null!;

        /// <summary>
        /// `PUBLIC` or `PRIVATE`, repo's visibility.
        /// </summary>
        [Output("repoType")]
        public Output<string> RepoType { get; private set; } = null!;

        /// <summary>
        /// The repository general information. It can contain 1 to 100 characters.
        /// </summary>
        [Output("summary")]
        public Output<string> Summary { get; private set; } = null!;


        /// <summary>
        /// Create a RegistryEnterpriseRepo resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RegistryEnterpriseRepo(string name, RegistryEnterpriseRepoArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cs/registryEnterpriseRepo:RegistryEnterpriseRepo", name, args ?? new RegistryEnterpriseRepoArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RegistryEnterpriseRepo(string name, Input<string> id, RegistryEnterpriseRepoState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cs/registryEnterpriseRepo:RegistryEnterpriseRepo", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RegistryEnterpriseRepo resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RegistryEnterpriseRepo Get(string name, Input<string> id, RegistryEnterpriseRepoState? state = null, CustomResourceOptions? options = null)
        {
            return new RegistryEnterpriseRepo(name, id, state, options);
        }
    }

    public sealed class RegistryEnterpriseRepoArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The repository specific information. MarkDown format is supported, and the length limit is 2000.
        /// </summary>
        [Input("detail")]
        public Input<string>? Detail { get; set; }

        /// <summary>
        /// ID of Container Registry Enterprise Edition instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Name of Container Registry Enterprise Edition repository. It can contain 2 to 64 characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Name of Container Registry Enterprise Edition namespace where repository is located. It can contain 2 to 30 characters.
        /// </summary>
        [Input("namespace", required: true)]
        public Input<string> Namespace { get; set; } = null!;

        /// <summary>
        /// `PUBLIC` or `PRIVATE`, repo's visibility.
        /// </summary>
        [Input("repoType", required: true)]
        public Input<string> RepoType { get; set; } = null!;

        /// <summary>
        /// The repository general information. It can contain 1 to 100 characters.
        /// </summary>
        [Input("summary", required: true)]
        public Input<string> Summary { get; set; } = null!;

        public RegistryEnterpriseRepoArgs()
        {
        }
        public static new RegistryEnterpriseRepoArgs Empty => new RegistryEnterpriseRepoArgs();
    }

    public sealed class RegistryEnterpriseRepoState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The repository specific information. MarkDown format is supported, and the length limit is 2000.
        /// </summary>
        [Input("detail")]
        public Input<string>? Detail { get; set; }

        /// <summary>
        /// ID of Container Registry Enterprise Edition instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Name of Container Registry Enterprise Edition repository. It can contain 2 to 64 characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Name of Container Registry Enterprise Edition namespace where repository is located. It can contain 2 to 30 characters.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The uuid of Container Registry Enterprise Edition repository.
        /// </summary>
        [Input("repoId")]
        public Input<string>? RepoId { get; set; }

        /// <summary>
        /// `PUBLIC` or `PRIVATE`, repo's visibility.
        /// </summary>
        [Input("repoType")]
        public Input<string>? RepoType { get; set; }

        /// <summary>
        /// The repository general information. It can contain 1 to 100 characters.
        /// </summary>
        [Input("summary")]
        public Input<string>? Summary { get; set; }

        public RegistryEnterpriseRepoState()
        {
        }
        public static new RegistryEnterpriseRepoState Empty => new RegistryEnterpriseRepoState();
    }
}

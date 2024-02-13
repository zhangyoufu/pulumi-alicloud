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
    /// Provides a CR Chain resource.
    /// 
    /// For information about CR Chain and how to use it, see [What is Chain](https://www.alibabacloud.com/help/en/acr/developer-reference/api-cr-2018-12-01-createchain).
    /// 
    /// &gt; **NOTE:** Available since v1.161.0.
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
    ///     var defaultRegistryEnterpriseInstance = new AliCloud.CR.RegistryEnterpriseInstance("defaultRegistryEnterpriseInstance", new()
    ///     {
    ///         PaymentType = "Subscription",
    ///         Period = 1,
    ///         RenewPeriod = 0,
    ///         RenewalStatus = "ManualRenewal",
    ///         InstanceType = "Advanced",
    ///         InstanceName = name,
    ///     });
    /// 
    ///     var defaultRegistryEnterpriseNamespace = new AliCloud.CS.RegistryEnterpriseNamespace("defaultRegistryEnterpriseNamespace", new()
    ///     {
    ///         InstanceId = defaultRegistryEnterpriseInstance.Id,
    ///         AutoCreate = false,
    ///         DefaultVisibility = "PUBLIC",
    ///     });
    /// 
    ///     var defaultRegistryEnterpriseRepo = new AliCloud.CS.RegistryEnterpriseRepo("defaultRegistryEnterpriseRepo", new()
    ///     {
    ///         InstanceId = defaultRegistryEnterpriseInstance.Id,
    ///         Namespace = defaultRegistryEnterpriseNamespace.Name,
    ///         Summary = "this is summary of my new repo",
    ///         RepoType = "PUBLIC",
    ///         Detail = "this is a public repo",
    ///     });
    /// 
    ///     var defaultChain = new AliCloud.CR.Chain("defaultChain", new()
    ///     {
    ///         ChainName = name,
    ///         Description = name,
    ///         InstanceId = defaultRegistryEnterpriseNamespace.InstanceId,
    ///         RepoName = defaultRegistryEnterpriseRepo.Name,
    ///         RepoNamespaceName = defaultRegistryEnterpriseNamespace.Name,
    ///         ChainConfigs = new[]
    ///         {
    ///             new AliCloud.CR.Inputs.ChainChainConfigArgs
    ///             {
    ///                 Routers = new[]
    ///                 {
    ///                     new AliCloud.CR.Inputs.ChainChainConfigRouterArgs
    ///                     {
    ///                         Froms = new[]
    ///                         {
    ///                             new AliCloud.CR.Inputs.ChainChainConfigRouterFromArgs
    ///                             {
    ///                                 NodeName = "DOCKER_IMAGE_BUILD",
    ///                             },
    ///                         },
    ///                         Tos = new[]
    ///                         {
    ///                             new AliCloud.CR.Inputs.ChainChainConfigRouterToArgs
    ///                             {
    ///                                 NodeName = "DOCKER_IMAGE_PUSH",
    ///                             },
    ///                         },
    ///                     },
    ///                     new AliCloud.CR.Inputs.ChainChainConfigRouterArgs
    ///                     {
    ///                         Froms = new[]
    ///                         {
    ///                             new AliCloud.CR.Inputs.ChainChainConfigRouterFromArgs
    ///                             {
    ///                                 NodeName = "DOCKER_IMAGE_PUSH",
    ///                             },
    ///                         },
    ///                         Tos = new[]
    ///                         {
    ///                             new AliCloud.CR.Inputs.ChainChainConfigRouterToArgs
    ///                             {
    ///                                 NodeName = "VULNERABILITY_SCANNING",
    ///                             },
    ///                         },
    ///                     },
    ///                     new AliCloud.CR.Inputs.ChainChainConfigRouterArgs
    ///                     {
    ///                         Froms = new[]
    ///                         {
    ///                             new AliCloud.CR.Inputs.ChainChainConfigRouterFromArgs
    ///                             {
    ///                                 NodeName = "VULNERABILITY_SCANNING",
    ///                             },
    ///                         },
    ///                         Tos = new[]
    ///                         {
    ///                             new AliCloud.CR.Inputs.ChainChainConfigRouterToArgs
    ///                             {
    ///                                 NodeName = "ACTIVATE_REPLICATION",
    ///                             },
    ///                         },
    ///                     },
    ///                     new AliCloud.CR.Inputs.ChainChainConfigRouterArgs
    ///                     {
    ///                         Froms = new[]
    ///                         {
    ///                             new AliCloud.CR.Inputs.ChainChainConfigRouterFromArgs
    ///                             {
    ///                                 NodeName = "ACTIVATE_REPLICATION",
    ///                             },
    ///                         },
    ///                         Tos = new[]
    ///                         {
    ///                             new AliCloud.CR.Inputs.ChainChainConfigRouterToArgs
    ///                             {
    ///                                 NodeName = "TRIGGER",
    ///                             },
    ///                         },
    ///                     },
    ///                     new AliCloud.CR.Inputs.ChainChainConfigRouterArgs
    ///                     {
    ///                         Froms = new[]
    ///                         {
    ///                             new AliCloud.CR.Inputs.ChainChainConfigRouterFromArgs
    ///                             {
    ///                                 NodeName = "VULNERABILITY_SCANNING",
    ///                             },
    ///                         },
    ///                         Tos = new[]
    ///                         {
    ///                             new AliCloud.CR.Inputs.ChainChainConfigRouterToArgs
    ///                             {
    ///                                 NodeName = "SNAPSHOT",
    ///                             },
    ///                         },
    ///                     },
    ///                     new AliCloud.CR.Inputs.ChainChainConfigRouterArgs
    ///                     {
    ///                         Froms = new[]
    ///                         {
    ///                             new AliCloud.CR.Inputs.ChainChainConfigRouterFromArgs
    ///                             {
    ///                                 NodeName = "SNAPSHOT",
    ///                             },
    ///                         },
    ///                         Tos = new[]
    ///                         {
    ///                             new AliCloud.CR.Inputs.ChainChainConfigRouterToArgs
    ///                             {
    ///                                 NodeName = "TRIGGER_SNAPSHOT",
    ///                             },
    ///                         },
    ///                     },
    ///                 },
    ///                 Nodes = new[]
    ///                 {
    ///                     new AliCloud.CR.Inputs.ChainChainConfigNodeArgs
    ///                     {
    ///                         Enable = true,
    ///                         NodeName = "DOCKER_IMAGE_BUILD",
    ///                         NodeConfigs = new[]
    ///                         {
    ///                             new AliCloud.CR.Inputs.ChainChainConfigNodeNodeConfigArgs
    ///                             {
    ///                                 DenyPolicies = new[]
    ///                                 {
    ///                                     null,
    ///                                 },
    ///                             },
    ///                         },
    ///                     },
    ///                     new AliCloud.CR.Inputs.ChainChainConfigNodeArgs
    ///                     {
    ///                         Enable = true,
    ///                         NodeName = "DOCKER_IMAGE_PUSH",
    ///                         NodeConfigs = new[]
    ///                         {
    ///                             new AliCloud.CR.Inputs.ChainChainConfigNodeNodeConfigArgs
    ///                             {
    ///                                 DenyPolicies = new[]
    ///                                 {
    ///                                     null,
    ///                                 },
    ///                             },
    ///                         },
    ///                     },
    ///                     new AliCloud.CR.Inputs.ChainChainConfigNodeArgs
    ///                     {
    ///                         Enable = true,
    ///                         NodeName = "VULNERABILITY_SCANNING",
    ///                         NodeConfigs = new[]
    ///                         {
    ///                             new AliCloud.CR.Inputs.ChainChainConfigNodeNodeConfigArgs
    ///                             {
    ///                                 DenyPolicies = new[]
    ///                                 {
    ///                                     new AliCloud.CR.Inputs.ChainChainConfigNodeNodeConfigDenyPolicyArgs
    ///                                     {
    ///                                         IssueLevel = "MEDIUM",
    ///                                         IssueCount = "1",
    ///                                         Action = "BLOCK_DELETE_TAG",
    ///                                         Logic = "AND",
    ///                                     },
    ///                                 },
    ///                             },
    ///                         },
    ///                     },
    ///                     new AliCloud.CR.Inputs.ChainChainConfigNodeArgs
    ///                     {
    ///                         Enable = true,
    ///                         NodeName = "ACTIVATE_REPLICATION",
    ///                         NodeConfigs = new[]
    ///                         {
    ///                             new AliCloud.CR.Inputs.ChainChainConfigNodeNodeConfigArgs
    ///                             {
    ///                                 DenyPolicies = new[]
    ///                                 {
    ///                                     null,
    ///                                 },
    ///                             },
    ///                         },
    ///                     },
    ///                     new AliCloud.CR.Inputs.ChainChainConfigNodeArgs
    ///                     {
    ///                         Enable = true,
    ///                         NodeName = "TRIGGER",
    ///                         NodeConfigs = new[]
    ///                         {
    ///                             new AliCloud.CR.Inputs.ChainChainConfigNodeNodeConfigArgs
    ///                             {
    ///                                 DenyPolicies = new[]
    ///                                 {
    ///                                     null,
    ///                                 },
    ///                             },
    ///                         },
    ///                     },
    ///                     new AliCloud.CR.Inputs.ChainChainConfigNodeArgs
    ///                     {
    ///                         Enable = false,
    ///                         NodeName = "SNAPSHOT",
    ///                         NodeConfigs = new[]
    ///                         {
    ///                             new AliCloud.CR.Inputs.ChainChainConfigNodeNodeConfigArgs
    ///                             {
    ///                                 DenyPolicies = new[]
    ///                                 {
    ///                                     null,
    ///                                 },
    ///                             },
    ///                         },
    ///                     },
    ///                     new AliCloud.CR.Inputs.ChainChainConfigNodeArgs
    ///                     {
    ///                         Enable = false,
    ///                         NodeName = "TRIGGER_SNAPSHOT",
    ///                         NodeConfigs = new[]
    ///                         {
    ///                             new AliCloud.CR.Inputs.ChainChainConfigNodeNodeConfigArgs
    ///                             {
    ///                                 DenyPolicies = new[]
    ///                                 {
    ///                                     null,
    ///                                 },
    ///                             },
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// CR Chain can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:cr/chain:Chain example &lt;instance_id&gt;:&lt;chain_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cr/chain:Chain")]
    public partial class Chain : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The configuration of delivery chain. See `chain_config` below. **NOTE:** This parameter must specify the correct value, otherwise the created resource will be incorrect.
        /// </summary>
        [Output("chainConfigs")]
        public Output<ImmutableArray<Outputs.ChainChainConfig>> ChainConfigs { get; private set; } = null!;

        /// <summary>
        /// Delivery chain ID.
        /// </summary>
        [Output("chainId")]
        public Output<string> ChainId { get; private set; } = null!;

        /// <summary>
        /// The name of delivery chain. The length of the name is 1-64 characters, lowercase English letters and numbers, and the separators "_", "-", "." can be used, noted that the separator cannot be at the first or last position.
        /// </summary>
        [Output("chainName")]
        public Output<string> ChainName { get; private set; } = null!;

        /// <summary>
        /// The description delivery chain.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The ID of CR Enterprise Edition instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The name of CR Enterprise Edition repository. **NOTE:** This parameter must specify a correct value, otherwise the created resource will be incorrect.
        /// </summary>
        [Output("repoName")]
        public Output<string?> RepoName { get; private set; } = null!;

        /// <summary>
        /// The name of CR Enterprise Edition namespace. **NOTE:** This parameter must specify the correct value, otherwise the created resource will be incorrect.
        /// </summary>
        [Output("repoNamespaceName")]
        public Output<string?> RepoNamespaceName { get; private set; } = null!;


        /// <summary>
        /// Create a Chain resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Chain(string name, ChainArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cr/chain:Chain", name, args ?? new ChainArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Chain(string name, Input<string> id, ChainState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cr/chain:Chain", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Chain resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Chain Get(string name, Input<string> id, ChainState? state = null, CustomResourceOptions? options = null)
        {
            return new Chain(name, id, state, options);
        }
    }

    public sealed class ChainArgs : global::Pulumi.ResourceArgs
    {
        [Input("chainConfigs")]
        private InputList<Inputs.ChainChainConfigArgs>? _chainConfigs;

        /// <summary>
        /// The configuration of delivery chain. See `chain_config` below. **NOTE:** This parameter must specify the correct value, otherwise the created resource will be incorrect.
        /// </summary>
        public InputList<Inputs.ChainChainConfigArgs> ChainConfigs
        {
            get => _chainConfigs ?? (_chainConfigs = new InputList<Inputs.ChainChainConfigArgs>());
            set => _chainConfigs = value;
        }

        /// <summary>
        /// The name of delivery chain. The length of the name is 1-64 characters, lowercase English letters and numbers, and the separators "_", "-", "." can be used, noted that the separator cannot be at the first or last position.
        /// </summary>
        [Input("chainName", required: true)]
        public Input<string> ChainName { get; set; } = null!;

        /// <summary>
        /// The description delivery chain.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of CR Enterprise Edition instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The name of CR Enterprise Edition repository. **NOTE:** This parameter must specify a correct value, otherwise the created resource will be incorrect.
        /// </summary>
        [Input("repoName")]
        public Input<string>? RepoName { get; set; }

        /// <summary>
        /// The name of CR Enterprise Edition namespace. **NOTE:** This parameter must specify the correct value, otherwise the created resource will be incorrect.
        /// </summary>
        [Input("repoNamespaceName")]
        public Input<string>? RepoNamespaceName { get; set; }

        public ChainArgs()
        {
        }
        public static new ChainArgs Empty => new ChainArgs();
    }

    public sealed class ChainState : global::Pulumi.ResourceArgs
    {
        [Input("chainConfigs")]
        private InputList<Inputs.ChainChainConfigGetArgs>? _chainConfigs;

        /// <summary>
        /// The configuration of delivery chain. See `chain_config` below. **NOTE:** This parameter must specify the correct value, otherwise the created resource will be incorrect.
        /// </summary>
        public InputList<Inputs.ChainChainConfigGetArgs> ChainConfigs
        {
            get => _chainConfigs ?? (_chainConfigs = new InputList<Inputs.ChainChainConfigGetArgs>());
            set => _chainConfigs = value;
        }

        /// <summary>
        /// Delivery chain ID.
        /// </summary>
        [Input("chainId")]
        public Input<string>? ChainId { get; set; }

        /// <summary>
        /// The name of delivery chain. The length of the name is 1-64 characters, lowercase English letters and numbers, and the separators "_", "-", "." can be used, noted that the separator cannot be at the first or last position.
        /// </summary>
        [Input("chainName")]
        public Input<string>? ChainName { get; set; }

        /// <summary>
        /// The description delivery chain.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of CR Enterprise Edition instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The name of CR Enterprise Edition repository. **NOTE:** This parameter must specify a correct value, otherwise the created resource will be incorrect.
        /// </summary>
        [Input("repoName")]
        public Input<string>? RepoName { get; set; }

        /// <summary>
        /// The name of CR Enterprise Edition namespace. **NOTE:** This parameter must specify the correct value, otherwise the created resource will be incorrect.
        /// </summary>
        [Input("repoNamespaceName")]
        public Input<string>? RepoNamespaceName { get; set; }

        public ChainState()
        {
        }
        public static new ChainState Empty => new ChainState();
    }
}

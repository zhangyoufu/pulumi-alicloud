// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ServiceCatalog
{
    /// <summary>
    /// Provides a Service Catalog Provisioned Product resource.
    /// 
    /// For information about Service Catalog Provisioned Product and how to use it, see [What is Provisioned Product](https://www.alibabacloud.com/help/en/service-catalog/developer-reference/api-servicecatalog-2021-09-01-launchproduct).
    /// 
    /// &gt; **NOTE:** Available in v1.196.0+.
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
    ///     var name = config.Get("name") ?? "tf-testAccServiceCatalogProvisionedProduct";
    ///     var @default = new AliCloud.ServiceCatalog.ProvisionedProduct("default", new()
    ///     {
    ///         ProvisionedProductName = name,
    ///         StackRegionId = "cn-hangzhou",
    ///         ProductVersionId = "pv-bp1d7dxy2pcc1g",
    ///         ProductId = "prod-bp1u3dkc282cwd",
    ///         PortfolioId = "port-bp119dvn27jccw",
    ///         Tags = 
    ///         {
    ///             { "v1", "tf-test" },
    ///         },
    ///         Parameters = new[]
    ///         {
    ///             new AliCloud.ServiceCatalog.Inputs.ProvisionedProductParameterArgs
    ///             {
    ///                 ParameterKey = "role_name",
    ///                 ParameterValue = name,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Service Catalog Provisioned Product can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:servicecatalog/provisionedProduct:ProvisionedProduct example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:servicecatalog/provisionedProduct:ProvisionedProduct")]
    public partial class ProvisionedProduct : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The creation time of the product instance
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The ID of the last instance operation task
        /// </summary>
        [Output("lastProvisioningTaskId")]
        public Output<string> LastProvisioningTaskId { get; private set; } = null!;

        /// <summary>
        /// The ID of the last successful instance operation task
        /// </summary>
        [Output("lastSuccessfulProvisioningTaskId")]
        public Output<string> LastSuccessfulProvisioningTaskId { get; private set; } = null!;

        /// <summary>
        /// The ID of the last task
        /// </summary>
        [Output("lastTaskId")]
        public Output<string> LastTaskId { get; private set; } = null!;

        /// <summary>
        /// The output value of the template.
        /// </summary>
        [Output("outputs")]
        public Output<ImmutableArray<Outputs.ProvisionedProductOutput>> Outputs { get; private set; } = null!;

        /// <summary>
        /// The RAM entity ID of the owner
        /// </summary>
        [Output("ownerPrincipalId")]
        public Output<string> OwnerPrincipalId { get; private set; } = null!;

        /// <summary>
        /// The RAM entity type of the owner
        /// </summary>
        [Output("ownerPrincipalType")]
        public Output<string> OwnerPrincipalType { get; private set; } = null!;

        /// <summary>
        /// Template parameters entered by the user.The maximum value of N is 200.See the following `Block Parameters`.
        /// </summary>
        [Output("parameters")]
        public Output<ImmutableArray<Outputs.ProvisionedProductParameter>> Parameters { get; private set; } = null!;

        /// <summary>
        /// Product mix ID.&gt; When there is a default Startup option, there is no need to fill in the portfolio. When there is no default Startup option, you must fill in the portfolio.
        /// </summary>
        [Output("portfolioId")]
        public Output<string?> PortfolioId { get; private set; } = null!;

        /// <summary>
        /// Product ID.
        /// </summary>
        [Output("productId")]
        public Output<string> ProductId { get; private set; } = null!;

        /// <summary>
        /// The name of the product
        /// </summary>
        [Output("productName")]
        public Output<string> ProductName { get; private set; } = null!;

        /// <summary>
        /// Product version ID.
        /// </summary>
        [Output("productVersionId")]
        public Output<string> ProductVersionId { get; private set; } = null!;

        /// <summary>
        /// The name of the product version
        /// </summary>
        [Output("productVersionName")]
        public Output<string> ProductVersionName { get; private set; } = null!;

        /// <summary>
        /// The ARN of the product instance
        /// </summary>
        [Output("provisionedProductArn")]
        public Output<string> ProvisionedProductArn { get; private set; } = null!;

        /// <summary>
        /// The ID of the instance.
        /// </summary>
        [Output("provisionedProductId")]
        public Output<string> ProvisionedProductId { get; private set; } = null!;

        /// <summary>
        /// The name of the instance.The length is 1~128 characters.
        /// </summary>
        [Output("provisionedProductName")]
        public Output<string> ProvisionedProductName { get; private set; } = null!;

        /// <summary>
        /// Instance type.The value is RosStack, which indicates the stack of Alibaba Cloud resource orchestration service (ROS).
        /// </summary>
        [Output("provisionedProductType")]
        public Output<string> ProvisionedProductType { get; private set; } = null!;

        /// <summary>
        /// The ID of the ROS stack
        /// </summary>
        [Output("stackId")]
        public Output<string> StackId { get; private set; } = null!;

        /// <summary>
        /// The ID of the region to which the resource stack of the Alibaba Cloud resource orchestration service (ROS) belongs.
        /// </summary>
        [Output("stackRegionId")]
        public Output<string> StackRegionId { get; private set; } = null!;

        /// <summary>
        /// Instance status
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The status message of the product instance
        /// </summary>
        [Output("statusMessage")]
        public Output<string> StatusMessage { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a ProvisionedProduct resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProvisionedProduct(string name, ProvisionedProductArgs args, CustomResourceOptions? options = null)
            : base("alicloud:servicecatalog/provisionedProduct:ProvisionedProduct", name, args ?? new ProvisionedProductArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProvisionedProduct(string name, Input<string> id, ProvisionedProductState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:servicecatalog/provisionedProduct:ProvisionedProduct", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProvisionedProduct resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProvisionedProduct Get(string name, Input<string> id, ProvisionedProductState? state = null, CustomResourceOptions? options = null)
        {
            return new ProvisionedProduct(name, id, state, options);
        }
    }

    public sealed class ProvisionedProductArgs : global::Pulumi.ResourceArgs
    {
        [Input("parameters")]
        private InputList<Inputs.ProvisionedProductParameterArgs>? _parameters;

        /// <summary>
        /// Template parameters entered by the user.The maximum value of N is 200.See the following `Block Parameters`.
        /// </summary>
        public InputList<Inputs.ProvisionedProductParameterArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.ProvisionedProductParameterArgs>());
            set => _parameters = value;
        }

        /// <summary>
        /// Product mix ID.&gt; When there is a default Startup option, there is no need to fill in the portfolio. When there is no default Startup option, you must fill in the portfolio.
        /// </summary>
        [Input("portfolioId")]
        public Input<string>? PortfolioId { get; set; }

        /// <summary>
        /// Product ID.
        /// </summary>
        [Input("productId", required: true)]
        public Input<string> ProductId { get; set; } = null!;

        /// <summary>
        /// Product version ID.
        /// </summary>
        [Input("productVersionId", required: true)]
        public Input<string> ProductVersionId { get; set; } = null!;

        /// <summary>
        /// The ID of the instance.
        /// </summary>
        [Input("provisionedProductId")]
        public Input<string>? ProvisionedProductId { get; set; }

        /// <summary>
        /// The name of the instance.The length is 1~128 characters.
        /// </summary>
        [Input("provisionedProductName", required: true)]
        public Input<string> ProvisionedProductName { get; set; } = null!;

        /// <summary>
        /// The ID of the region to which the resource stack of the Alibaba Cloud resource orchestration service (ROS) belongs.
        /// </summary>
        [Input("stackRegionId", required: true)]
        public Input<string> StackRegionId { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ProvisionedProductArgs()
        {
        }
        public static new ProvisionedProductArgs Empty => new ProvisionedProductArgs();
    }

    public sealed class ProvisionedProductState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The creation time of the product instance
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The ID of the last instance operation task
        /// </summary>
        [Input("lastProvisioningTaskId")]
        public Input<string>? LastProvisioningTaskId { get; set; }

        /// <summary>
        /// The ID of the last successful instance operation task
        /// </summary>
        [Input("lastSuccessfulProvisioningTaskId")]
        public Input<string>? LastSuccessfulProvisioningTaskId { get; set; }

        /// <summary>
        /// The ID of the last task
        /// </summary>
        [Input("lastTaskId")]
        public Input<string>? LastTaskId { get; set; }

        [Input("outputs")]
        private InputList<Inputs.ProvisionedProductOutputGetArgs>? _outputs;

        /// <summary>
        /// The output value of the template.
        /// </summary>
        public InputList<Inputs.ProvisionedProductOutputGetArgs> Outputs
        {
            get => _outputs ?? (_outputs = new InputList<Inputs.ProvisionedProductOutputGetArgs>());
            set => _outputs = value;
        }

        /// <summary>
        /// The RAM entity ID of the owner
        /// </summary>
        [Input("ownerPrincipalId")]
        public Input<string>? OwnerPrincipalId { get; set; }

        /// <summary>
        /// The RAM entity type of the owner
        /// </summary>
        [Input("ownerPrincipalType")]
        public Input<string>? OwnerPrincipalType { get; set; }

        [Input("parameters")]
        private InputList<Inputs.ProvisionedProductParameterGetArgs>? _parameters;

        /// <summary>
        /// Template parameters entered by the user.The maximum value of N is 200.See the following `Block Parameters`.
        /// </summary>
        public InputList<Inputs.ProvisionedProductParameterGetArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.ProvisionedProductParameterGetArgs>());
            set => _parameters = value;
        }

        /// <summary>
        /// Product mix ID.&gt; When there is a default Startup option, there is no need to fill in the portfolio. When there is no default Startup option, you must fill in the portfolio.
        /// </summary>
        [Input("portfolioId")]
        public Input<string>? PortfolioId { get; set; }

        /// <summary>
        /// Product ID.
        /// </summary>
        [Input("productId")]
        public Input<string>? ProductId { get; set; }

        /// <summary>
        /// The name of the product
        /// </summary>
        [Input("productName")]
        public Input<string>? ProductName { get; set; }

        /// <summary>
        /// Product version ID.
        /// </summary>
        [Input("productVersionId")]
        public Input<string>? ProductVersionId { get; set; }

        /// <summary>
        /// The name of the product version
        /// </summary>
        [Input("productVersionName")]
        public Input<string>? ProductVersionName { get; set; }

        /// <summary>
        /// The ARN of the product instance
        /// </summary>
        [Input("provisionedProductArn")]
        public Input<string>? ProvisionedProductArn { get; set; }

        /// <summary>
        /// The ID of the instance.
        /// </summary>
        [Input("provisionedProductId")]
        public Input<string>? ProvisionedProductId { get; set; }

        /// <summary>
        /// The name of the instance.The length is 1~128 characters.
        /// </summary>
        [Input("provisionedProductName")]
        public Input<string>? ProvisionedProductName { get; set; }

        /// <summary>
        /// Instance type.The value is RosStack, which indicates the stack of Alibaba Cloud resource orchestration service (ROS).
        /// </summary>
        [Input("provisionedProductType")]
        public Input<string>? ProvisionedProductType { get; set; }

        /// <summary>
        /// The ID of the ROS stack
        /// </summary>
        [Input("stackId")]
        public Input<string>? StackId { get; set; }

        /// <summary>
        /// The ID of the region to which the resource stack of the Alibaba Cloud resource orchestration service (ROS) belongs.
        /// </summary>
        [Input("stackRegionId")]
        public Input<string>? StackRegionId { get; set; }

        /// <summary>
        /// Instance status
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The status message of the product instance
        /// </summary>
        [Input("statusMessage")]
        public Input<string>? StatusMessage { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ProvisionedProductState()
        {
        }
        public static new ProvisionedProductState Empty => new ProvisionedProductState();
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Arms
{
    /// <summary>
    /// Provides a ARMS Environment resource. The arms environment.
    /// 
    /// For information about ARMS Environment and how to use it, see [What is Environment](https://www.alibabacloud.com/help/en/arms/developer-reference/api-arms-2019-08-08-createenvironment).
    /// 
    /// &gt; **NOTE:** Available since v1.212.0.
    /// 
    /// ## Import
    /// 
    /// ARMS Environment can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:arms/environment:Environment example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:arms/environment:Environment")]
    public partial class Environment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The locale. The default is Chinese zh | en.
        /// </summary>
        [Output("aliyunLang")]
        public Output<string?> AliyunLang { get; private set; } = null!;

        /// <summary>
        /// The id or vpcId of the bound container instance.
        /// </summary>
        [Output("bindResourceId")]
        public Output<string?> BindResourceId { get; private set; } = null!;

        /// <summary>
        /// List of abandoned indicators.
        /// </summary>
        [Output("dropMetrics")]
        public Output<string?> DropMetrics { get; private set; } = null!;

        /// <summary>
        /// The first ID of the resource.
        /// </summary>
        [Output("environmentId")]
        public Output<string> EnvironmentId { get; private set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Output("environmentName")]
        public Output<string?> EnvironmentName { get; private set; } = null!;

        /// <summary>
        /// Subtype of environment:
        /// - Type of CS: ACK is currently supported.
        /// - Type of ECS: currently supports ECS.
        /// - Type of Cloud: currently supports Cloud.
        /// </summary>
        [Output("environmentSubType")]
        public Output<string> EnvironmentSubType { get; private set; } = null!;

        /// <summary>
        /// Type of environment.
        /// </summary>
        [Output("environmentType")]
        public Output<string> EnvironmentType { get; private set; } = null!;

        /// <summary>
        /// Hosting type:
        /// - none: unmanaged. The default value of the ACK cluster.
        /// - agent: Managed agent (including ksm). Default values of ASK, ACS, and Acone clusters.
        /// - agent-exproter: Managed agent and exporter. The default value of the cloud service type.
        /// </summary>
        [Output("managedType")]
        public Output<string> ManagedType { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// The tag of the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Environment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Environment(string name, EnvironmentArgs args, CustomResourceOptions? options = null)
            : base("alicloud:arms/environment:Environment", name, args ?? new EnvironmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Environment(string name, Input<string> id, EnvironmentState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:arms/environment:Environment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Environment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Environment Get(string name, Input<string> id, EnvironmentState? state = null, CustomResourceOptions? options = null)
        {
            return new Environment(name, id, state, options);
        }
    }

    public sealed class EnvironmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The locale. The default is Chinese zh | en.
        /// </summary>
        [Input("aliyunLang")]
        public Input<string>? AliyunLang { get; set; }

        /// <summary>
        /// The id or vpcId of the bound container instance.
        /// </summary>
        [Input("bindResourceId")]
        public Input<string>? BindResourceId { get; set; }

        /// <summary>
        /// List of abandoned indicators.
        /// </summary>
        [Input("dropMetrics")]
        public Input<string>? DropMetrics { get; set; }

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Input("environmentName")]
        public Input<string>? EnvironmentName { get; set; }

        /// <summary>
        /// Subtype of environment:
        /// - Type of CS: ACK is currently supported.
        /// - Type of ECS: currently supports ECS.
        /// - Type of Cloud: currently supports Cloud.
        /// </summary>
        [Input("environmentSubType", required: true)]
        public Input<string> EnvironmentSubType { get; set; } = null!;

        /// <summary>
        /// Type of environment.
        /// </summary>
        [Input("environmentType", required: true)]
        public Input<string> EnvironmentType { get; set; } = null!;

        /// <summary>
        /// Hosting type:
        /// - none: unmanaged. The default value of the ACK cluster.
        /// - agent: Managed agent (including ksm). Default values of ASK, ACS, and Acone clusters.
        /// - agent-exproter: Managed agent and exporter. The default value of the cloud service type.
        /// </summary>
        [Input("managedType")]
        public Input<string>? ManagedType { get; set; }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The tag of the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public EnvironmentArgs()
        {
        }
        public static new EnvironmentArgs Empty => new EnvironmentArgs();
    }

    public sealed class EnvironmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The locale. The default is Chinese zh | en.
        /// </summary>
        [Input("aliyunLang")]
        public Input<string>? AliyunLang { get; set; }

        /// <summary>
        /// The id or vpcId of the bound container instance.
        /// </summary>
        [Input("bindResourceId")]
        public Input<string>? BindResourceId { get; set; }

        /// <summary>
        /// List of abandoned indicators.
        /// </summary>
        [Input("dropMetrics")]
        public Input<string>? DropMetrics { get; set; }

        /// <summary>
        /// The first ID of the resource.
        /// </summary>
        [Input("environmentId")]
        public Input<string>? EnvironmentId { get; set; }

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Input("environmentName")]
        public Input<string>? EnvironmentName { get; set; }

        /// <summary>
        /// Subtype of environment:
        /// - Type of CS: ACK is currently supported.
        /// - Type of ECS: currently supports ECS.
        /// - Type of Cloud: currently supports Cloud.
        /// </summary>
        [Input("environmentSubType")]
        public Input<string>? EnvironmentSubType { get; set; }

        /// <summary>
        /// Type of environment.
        /// </summary>
        [Input("environmentType")]
        public Input<string>? EnvironmentType { get; set; }

        /// <summary>
        /// Hosting type:
        /// - none: unmanaged. The default value of the ACK cluster.
        /// - agent: Managed agent (including ksm). Default values of ASK, ACS, and Acone clusters.
        /// - agent-exproter: Managed agent and exporter. The default value of the cloud service type.
        /// </summary>
        [Input("managedType")]
        public Input<string>? ManagedType { get; set; }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The tag of the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public EnvironmentState()
        {
        }
        public static new EnvironmentState Empty => new EnvironmentState();
    }
}

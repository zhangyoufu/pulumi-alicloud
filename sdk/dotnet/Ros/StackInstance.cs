// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ros
{
    /// <summary>
    /// Provides a ROS Stack Instance resource.
    /// 
    /// For information about ROS Stack Instance and how to use it, see [What is Stack Instance](https://www.alibabacloud.com/help/en/doc-detail/151338.html).
    /// 
    /// &gt; **NOTE:** Available in v1.145.0+.
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
    ///     var name = config.Get("name") ?? "tf-example";
    ///     var @this = AliCloud.GetAccount.Invoke();
    /// 
    ///     var defaultRegions = AliCloud.Ros.GetRegions.Invoke();
    /// 
    ///     var defaultStackGroup = new AliCloud.Ros.StackGroup("defaultStackGroup", new()
    ///     {
    ///         StackGroupName = name,
    ///         TemplateBody = "{\"ROSTemplateFormatVersion\":\"2015-09-01\", \"Parameters\": {\"VpcName\": {\"Type\": \"String\"},\"InstanceType\": {\"Type\": \"String\"}}}",
    ///         Description = "test for stack groups",
    ///         Parameters = new[]
    ///         {
    ///             new AliCloud.Ros.Inputs.StackGroupParameterArgs
    ///             {
    ///                 ParameterKey = "VpcName",
    ///                 ParameterValue = "VpcName",
    ///             },
    ///             new AliCloud.Ros.Inputs.StackGroupParameterArgs
    ///             {
    ///                 ParameterKey = "InstanceType",
    ///                 ParameterValue = "InstanceType",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var example = new AliCloud.Ros.StackInstance("example", new()
    ///     {
    ///         StackGroupName = defaultStackGroup.StackGroupName,
    ///         StackInstanceAccountId = @this.Apply(@this =&gt; @this.Apply(getAccountResult =&gt; getAccountResult.Id)),
    ///         StackInstanceRegionId = defaultRegions.Apply(getRegionsResult =&gt; getRegionsResult.Regions[0]?.RegionId),
    ///         OperationPreferences = "{\"FailureToleranceCount\": 1, \"MaxConcurrentCount\": 2}",
    ///         TimeoutInMinutes = "60",
    ///         OperationDescription = "tf-example",
    ///         RetainStacks = true,
    ///         ParameterOverrides = new[]
    ///         {
    ///             new AliCloud.Ros.Inputs.StackInstanceParameterOverrideArgs
    ///             {
    ///                 ParameterValue = "VpcName",
    ///                 ParameterKey = "VpcName",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// ROS Stack Instance can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:ros/stackInstance:StackInstance example &lt;stack_group_name&gt;:&lt;stack_instance_account_id&gt;:&lt;stack_instance_region_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:ros/stackInstance:StackInstance")]
    public partial class StackInstance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The operation description.
        /// </summary>
        [Output("operationDescription")]
        public Output<string?> OperationDescription { get; private set; } = null!;

        /// <summary>
        /// The operation preferences. The operation settings. The following fields are supported:
        /// </summary>
        [Output("operationPreferences")]
        public Output<string?> OperationPreferences { get; private set; } = null!;

        /// <summary>
        /// ParameterOverrides. See the following `Block parameter_overrides`.
        /// </summary>
        [Output("parameterOverrides")]
        public Output<ImmutableArray<Outputs.StackInstanceParameterOverride>> ParameterOverrides { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to retain the stack corresponding to the stack instance.Default value `false`. **NOTE:** When `retain_stacks` is `true`, the stack is retained. If the stack is retained, the corresponding stack is not deleted when the stack instance is deleted from the stack group.
        /// </summary>
        [Output("retainStacks")]
        public Output<bool?> RetainStacks { get; private set; } = null!;

        /// <summary>
        /// The name of the stack group.
        /// </summary>
        [Output("stackGroupName")]
        public Output<string> StackGroupName { get; private set; } = null!;

        /// <summary>
        /// The account to which the stack instance belongs.
        /// </summary>
        [Output("stackInstanceAccountId")]
        public Output<string> StackInstanceAccountId { get; private set; } = null!;

        /// <summary>
        /// The region of the stack instance.
        /// </summary>
        [Output("stackInstanceRegionId")]
        public Output<string> StackInstanceRegionId { get; private set; } = null!;

        /// <summary>
        /// The status of the stack instance. Valid values: `CURRENT` or `OUTDATED`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The timeout period that is specified for the stack creation request. Default value: `60`. Unit: `minutes`.
        /// </summary>
        [Output("timeoutInMinutes")]
        public Output<string?> TimeoutInMinutes { get; private set; } = null!;


        /// <summary>
        /// Create a StackInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StackInstance(string name, StackInstanceArgs args, CustomResourceOptions? options = null)
            : base("alicloud:ros/stackInstance:StackInstance", name, args ?? new StackInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StackInstance(string name, Input<string> id, StackInstanceState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ros/stackInstance:StackInstance", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "parameterOverrides",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing StackInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StackInstance Get(string name, Input<string> id, StackInstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new StackInstance(name, id, state, options);
        }
    }

    public sealed class StackInstanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The operation description.
        /// </summary>
        [Input("operationDescription")]
        public Input<string>? OperationDescription { get; set; }

        /// <summary>
        /// The operation preferences. The operation settings. The following fields are supported:
        /// </summary>
        [Input("operationPreferences")]
        public Input<string>? OperationPreferences { get; set; }

        [Input("parameterOverrides")]
        private InputList<Inputs.StackInstanceParameterOverrideArgs>? _parameterOverrides;

        /// <summary>
        /// ParameterOverrides. See the following `Block parameter_overrides`.
        /// </summary>
        public InputList<Inputs.StackInstanceParameterOverrideArgs> ParameterOverrides
        {
            get => _parameterOverrides ?? (_parameterOverrides = new InputList<Inputs.StackInstanceParameterOverrideArgs>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableArray.Create<Inputs.StackInstanceParameterOverrideArgs>());
                _parameterOverrides = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        /// <summary>
        /// Specifies whether to retain the stack corresponding to the stack instance.Default value `false`. **NOTE:** When `retain_stacks` is `true`, the stack is retained. If the stack is retained, the corresponding stack is not deleted when the stack instance is deleted from the stack group.
        /// </summary>
        [Input("retainStacks")]
        public Input<bool>? RetainStacks { get; set; }

        /// <summary>
        /// The name of the stack group.
        /// </summary>
        [Input("stackGroupName", required: true)]
        public Input<string> StackGroupName { get; set; } = null!;

        /// <summary>
        /// The account to which the stack instance belongs.
        /// </summary>
        [Input("stackInstanceAccountId", required: true)]
        public Input<string> StackInstanceAccountId { get; set; } = null!;

        /// <summary>
        /// The region of the stack instance.
        /// </summary>
        [Input("stackInstanceRegionId", required: true)]
        public Input<string> StackInstanceRegionId { get; set; } = null!;

        /// <summary>
        /// The timeout period that is specified for the stack creation request. Default value: `60`. Unit: `minutes`.
        /// </summary>
        [Input("timeoutInMinutes")]
        public Input<string>? TimeoutInMinutes { get; set; }

        public StackInstanceArgs()
        {
        }
        public static new StackInstanceArgs Empty => new StackInstanceArgs();
    }

    public sealed class StackInstanceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The operation description.
        /// </summary>
        [Input("operationDescription")]
        public Input<string>? OperationDescription { get; set; }

        /// <summary>
        /// The operation preferences. The operation settings. The following fields are supported:
        /// </summary>
        [Input("operationPreferences")]
        public Input<string>? OperationPreferences { get; set; }

        [Input("parameterOverrides")]
        private InputList<Inputs.StackInstanceParameterOverrideGetArgs>? _parameterOverrides;

        /// <summary>
        /// ParameterOverrides. See the following `Block parameter_overrides`.
        /// </summary>
        public InputList<Inputs.StackInstanceParameterOverrideGetArgs> ParameterOverrides
        {
            get => _parameterOverrides ?? (_parameterOverrides = new InputList<Inputs.StackInstanceParameterOverrideGetArgs>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableArray.Create<Inputs.StackInstanceParameterOverrideGetArgs>());
                _parameterOverrides = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        /// <summary>
        /// Specifies whether to retain the stack corresponding to the stack instance.Default value `false`. **NOTE:** When `retain_stacks` is `true`, the stack is retained. If the stack is retained, the corresponding stack is not deleted when the stack instance is deleted from the stack group.
        /// </summary>
        [Input("retainStacks")]
        public Input<bool>? RetainStacks { get; set; }

        /// <summary>
        /// The name of the stack group.
        /// </summary>
        [Input("stackGroupName")]
        public Input<string>? StackGroupName { get; set; }

        /// <summary>
        /// The account to which the stack instance belongs.
        /// </summary>
        [Input("stackInstanceAccountId")]
        public Input<string>? StackInstanceAccountId { get; set; }

        /// <summary>
        /// The region of the stack instance.
        /// </summary>
        [Input("stackInstanceRegionId")]
        public Input<string>? StackInstanceRegionId { get; set; }

        /// <summary>
        /// The status of the stack instance. Valid values: `CURRENT` or `OUTDATED`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The timeout period that is specified for the stack creation request. Default value: `60`. Unit: `minutes`.
        /// </summary>
        [Input("timeoutInMinutes")]
        public Input<string>? TimeoutInMinutes { get; set; }

        public StackInstanceState()
        {
        }
        public static new StackInstanceState Empty => new StackInstanceState();
    }
}

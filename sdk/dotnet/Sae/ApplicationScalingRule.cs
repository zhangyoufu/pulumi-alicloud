// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Sae
{
    /// <summary>
    /// Provides a Serverless App Engine (SAE) Application Scaling Rule resource.
    /// 
    /// For information about Serverless App Engine (SAE) Application Scaling Rule and how to use it, see [What is Application Scaling Rule](https://www.alibabacloud.com/help/en/sae/latest/create-application-scaling-rule).
    /// 
    /// &gt; **NOTE:** Available since v1.159.0.
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
    /// using Random = Pulumi.Random;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "tf-example";
    ///     var defaultRegions = AliCloud.GetRegions.Invoke(new()
    ///     {
    ///         Current = true,
    ///     });
    /// 
    ///     var defaultRandomInteger = new Random.RandomInteger("defaultRandomInteger", new()
    ///     {
    ///         Max = 99999,
    ///         Min = 10000,
    ///     });
    /// 
    ///     var defaultZones = AliCloud.GetZones.Invoke(new()
    ///     {
    ///         AvailableResourceCreation = "VSwitch",
    ///     });
    /// 
    ///     var defaultNetwork = new AliCloud.Vpc.Network("defaultNetwork", new()
    ///     {
    ///         VpcName = name,
    ///         CidrBlock = "10.4.0.0/16",
    ///     });
    /// 
    ///     var defaultSwitch = new AliCloud.Vpc.Switch("defaultSwitch", new()
    ///     {
    ///         VswitchName = name,
    ///         CidrBlock = "10.4.0.0/24",
    ///         VpcId = defaultNetwork.Id,
    ///         ZoneId = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///     });
    /// 
    ///     var defaultSecurityGroup = new AliCloud.Ecs.SecurityGroup("defaultSecurityGroup", new()
    ///     {
    ///         VpcId = defaultNetwork.Id,
    ///     });
    /// 
    ///     var defaultNamespace = new AliCloud.Sae.Namespace("defaultNamespace", new()
    ///     {
    ///         NamespaceId = Output.Tuple(defaultRegions, defaultRandomInteger.Result).Apply(values =&gt;
    ///         {
    ///             var defaultRegions = values.Item1;
    ///             var result = values.Item2;
    ///             return $"{defaultRegions.Apply(getRegionsResult =&gt; getRegionsResult.Regions[0]?.Id)}:example{result}";
    ///         }),
    ///         NamespaceName = name,
    ///         NamespaceDescription = name,
    ///         EnableMicroRegistration = false,
    ///     });
    /// 
    ///     var defaultApplication = new AliCloud.Sae.Application("defaultApplication", new()
    ///     {
    ///         AppDescription = name,
    ///         AppName = name,
    ///         NamespaceId = defaultNamespace.Id,
    ///         ImageUrl = $"registry-vpc.{defaultRegions.Apply(getRegionsResult =&gt; getRegionsResult.Regions[0]?.Id)}.aliyuncs.com/sae-demo-image/consumer:1.0",
    ///         PackageType = "Image",
    ///         SecurityGroupId = defaultSecurityGroup.Id,
    ///         VpcId = defaultNetwork.Id,
    ///         VswitchId = defaultSwitch.Id,
    ///         Timezone = "Asia/Beijing",
    ///         Replicas = 5,
    ///         Cpu = 500,
    ///         Memory = 2048,
    ///     });
    /// 
    ///     var defaultApplicationScalingRule = new AliCloud.Sae.ApplicationScalingRule("defaultApplicationScalingRule", new()
    ///     {
    ///         AppId = defaultApplication.Id,
    ///         ScalingRuleName = name,
    ///         ScalingRuleEnable = true,
    ///         ScalingRuleType = "mix",
    ///         MinReadyInstances = 3,
    ///         MinReadyInstanceRatio = -1,
    ///         ScalingRuleTimer = new AliCloud.Sae.Inputs.ApplicationScalingRuleScalingRuleTimerArgs
    ///         {
    ///             Period = "* * *",
    ///             Schedules = new[]
    ///             {
    ///                 new AliCloud.Sae.Inputs.ApplicationScalingRuleScalingRuleTimerScheduleArgs
    ///                 {
    ///                     AtTime = "08:00",
    ///                     MaxReplicas = 10,
    ///                     MinReplicas = 3,
    ///                 },
    ///                 new AliCloud.Sae.Inputs.ApplicationScalingRuleScalingRuleTimerScheduleArgs
    ///                 {
    ///                     AtTime = "20:00",
    ///                     MaxReplicas = 50,
    ///                     MinReplicas = 3,
    ///                 },
    ///             },
    ///         },
    ///         ScalingRuleMetric = new AliCloud.Sae.Inputs.ApplicationScalingRuleScalingRuleMetricArgs
    ///         {
    ///             MaxReplicas = 50,
    ///             MinReplicas = 3,
    ///             Metrics = new[]
    ///             {
    ///                 new AliCloud.Sae.Inputs.ApplicationScalingRuleScalingRuleMetricMetricArgs
    ///                 {
    ///                     MetricType = "CPU",
    ///                     MetricTargetAverageUtilization = 20,
    ///                 },
    ///                 new AliCloud.Sae.Inputs.ApplicationScalingRuleScalingRuleMetricMetricArgs
    ///                 {
    ///                     MetricType = "MEMORY",
    ///                     MetricTargetAverageUtilization = 30,
    ///                 },
    ///                 new AliCloud.Sae.Inputs.ApplicationScalingRuleScalingRuleMetricMetricArgs
    ///                 {
    ///                     MetricType = "tcpActiveConn",
    ///                     MetricTargetAverageUtilization = 20,
    ///                 },
    ///             },
    ///             ScaleUpRules = new AliCloud.Sae.Inputs.ApplicationScalingRuleScalingRuleMetricScaleUpRulesArgs
    ///             {
    ///                 Step = 10,
    ///                 Disabled = false,
    ///                 StabilizationWindowSeconds = 0,
    ///             },
    ///             ScaleDownRules = new AliCloud.Sae.Inputs.ApplicationScalingRuleScalingRuleMetricScaleDownRulesArgs
    ///             {
    ///                 Step = 10,
    ///                 Disabled = false,
    ///                 StabilizationWindowSeconds = 10,
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
    /// Serverless App Engine (SAE) Application Scaling Rule can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:sae/applicationScalingRule:ApplicationScalingRule example &lt;app_id&gt;:&lt;scaling_rule_name&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:sae/applicationScalingRule:ApplicationScalingRule")]
    public partial class ApplicationScalingRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Application ID.
        /// </summary>
        [Output("appId")]
        public Output<string> AppId { get; private set; } = null!;

        /// <summary>
        /// The min ready instance ratio.
        /// </summary>
        [Output("minReadyInstanceRatio")]
        public Output<int?> MinReadyInstanceRatio { get; private set; } = null!;

        /// <summary>
        /// The min ready instances.
        /// </summary>
        [Output("minReadyInstances")]
        public Output<int?> MinReadyInstances { get; private set; } = null!;

        /// <summary>
        /// True whether the auto scaling policy is enabled. The value description is as follows: true: enabled state. false: disabled status. Valid values: `false`, `true`.
        /// </summary>
        [Output("scalingRuleEnable")]
        public Output<bool> ScalingRuleEnable { get; private set; } = null!;

        /// <summary>
        /// Monitor the configuration of the indicator elasticity strategy. See `scaling_rule_metric` below.
        /// </summary>
        [Output("scalingRuleMetric")]
        public Output<Outputs.ApplicationScalingRuleScalingRuleMetric?> ScalingRuleMetric { get; private set; } = null!;

        /// <summary>
        /// The name of a custom elastic scaling policy. In the application, the policy name cannot be repeated. It must start with a lowercase letter, and can only contain lowercase letters, numbers, and dashes (-), and no more than 32 characters. After the scaling policy is successfully created, the policy name cannot be modified.
        /// </summary>
        [Output("scalingRuleName")]
        public Output<string> ScalingRuleName { get; private set; } = null!;

        /// <summary>
        /// Configuration of Timing Resilient Policies. See `scaling_rule_timer` below.
        /// </summary>
        [Output("scalingRuleTimer")]
        public Output<Outputs.ApplicationScalingRuleScalingRuleTimer?> ScalingRuleTimer { get; private set; } = null!;

        /// <summary>
        /// Flexible strategy type. Valid values: `mix`, `timing` and `metric`.
        /// </summary>
        [Output("scalingRuleType")]
        public Output<string> ScalingRuleType { get; private set; } = null!;


        /// <summary>
        /// Create a ApplicationScalingRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApplicationScalingRule(string name, ApplicationScalingRuleArgs args, CustomResourceOptions? options = null)
            : base("alicloud:sae/applicationScalingRule:ApplicationScalingRule", name, args ?? new ApplicationScalingRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApplicationScalingRule(string name, Input<string> id, ApplicationScalingRuleState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:sae/applicationScalingRule:ApplicationScalingRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ApplicationScalingRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApplicationScalingRule Get(string name, Input<string> id, ApplicationScalingRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new ApplicationScalingRule(name, id, state, options);
        }
    }

    public sealed class ApplicationScalingRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Application ID.
        /// </summary>
        [Input("appId", required: true)]
        public Input<string> AppId { get; set; } = null!;

        /// <summary>
        /// The min ready instance ratio.
        /// </summary>
        [Input("minReadyInstanceRatio")]
        public Input<int>? MinReadyInstanceRatio { get; set; }

        /// <summary>
        /// The min ready instances.
        /// </summary>
        [Input("minReadyInstances")]
        public Input<int>? MinReadyInstances { get; set; }

        /// <summary>
        /// True whether the auto scaling policy is enabled. The value description is as follows: true: enabled state. false: disabled status. Valid values: `false`, `true`.
        /// </summary>
        [Input("scalingRuleEnable")]
        public Input<bool>? ScalingRuleEnable { get; set; }

        /// <summary>
        /// Monitor the configuration of the indicator elasticity strategy. See `scaling_rule_metric` below.
        /// </summary>
        [Input("scalingRuleMetric")]
        public Input<Inputs.ApplicationScalingRuleScalingRuleMetricArgs>? ScalingRuleMetric { get; set; }

        /// <summary>
        /// The name of a custom elastic scaling policy. In the application, the policy name cannot be repeated. It must start with a lowercase letter, and can only contain lowercase letters, numbers, and dashes (-), and no more than 32 characters. After the scaling policy is successfully created, the policy name cannot be modified.
        /// </summary>
        [Input("scalingRuleName", required: true)]
        public Input<string> ScalingRuleName { get; set; } = null!;

        /// <summary>
        /// Configuration of Timing Resilient Policies. See `scaling_rule_timer` below.
        /// </summary>
        [Input("scalingRuleTimer")]
        public Input<Inputs.ApplicationScalingRuleScalingRuleTimerArgs>? ScalingRuleTimer { get; set; }

        /// <summary>
        /// Flexible strategy type. Valid values: `mix`, `timing` and `metric`.
        /// </summary>
        [Input("scalingRuleType", required: true)]
        public Input<string> ScalingRuleType { get; set; } = null!;

        public ApplicationScalingRuleArgs()
        {
        }
        public static new ApplicationScalingRuleArgs Empty => new ApplicationScalingRuleArgs();
    }

    public sealed class ApplicationScalingRuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Application ID.
        /// </summary>
        [Input("appId")]
        public Input<string>? AppId { get; set; }

        /// <summary>
        /// The min ready instance ratio.
        /// </summary>
        [Input("minReadyInstanceRatio")]
        public Input<int>? MinReadyInstanceRatio { get; set; }

        /// <summary>
        /// The min ready instances.
        /// </summary>
        [Input("minReadyInstances")]
        public Input<int>? MinReadyInstances { get; set; }

        /// <summary>
        /// True whether the auto scaling policy is enabled. The value description is as follows: true: enabled state. false: disabled status. Valid values: `false`, `true`.
        /// </summary>
        [Input("scalingRuleEnable")]
        public Input<bool>? ScalingRuleEnable { get; set; }

        /// <summary>
        /// Monitor the configuration of the indicator elasticity strategy. See `scaling_rule_metric` below.
        /// </summary>
        [Input("scalingRuleMetric")]
        public Input<Inputs.ApplicationScalingRuleScalingRuleMetricGetArgs>? ScalingRuleMetric { get; set; }

        /// <summary>
        /// The name of a custom elastic scaling policy. In the application, the policy name cannot be repeated. It must start with a lowercase letter, and can only contain lowercase letters, numbers, and dashes (-), and no more than 32 characters. After the scaling policy is successfully created, the policy name cannot be modified.
        /// </summary>
        [Input("scalingRuleName")]
        public Input<string>? ScalingRuleName { get; set; }

        /// <summary>
        /// Configuration of Timing Resilient Policies. See `scaling_rule_timer` below.
        /// </summary>
        [Input("scalingRuleTimer")]
        public Input<Inputs.ApplicationScalingRuleScalingRuleTimerGetArgs>? ScalingRuleTimer { get; set; }

        /// <summary>
        /// Flexible strategy type. Valid values: `mix`, `timing` and `metric`.
        /// </summary>
        [Input("scalingRuleType")]
        public Input<string>? ScalingRuleType { get; set; }

        public ApplicationScalingRuleState()
        {
        }
        public static new ApplicationScalingRuleState Empty => new ApplicationScalingRuleState();
    }
}

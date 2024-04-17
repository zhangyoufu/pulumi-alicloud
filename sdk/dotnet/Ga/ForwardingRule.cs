// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ga
{
    /// <summary>
    /// Provides a Global Accelerator (GA) Forwarding Rule resource.
    /// 
    /// For information about Global Accelerator (GA) Forwarding Rule and how to use it, see [What is Forwarding Rule](https://www.alibabacloud.com/help/en/global-accelerator/latest/api-ga-2019-11-20-createforwardingrules).
    /// 
    /// &gt; **NOTE:** Available since v1.120.0.
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
    ///     var region = config.Get("region") ?? "cn-hangzhou";
    ///     var name = config.Get("name") ?? "tf-example";
    ///     var @default = AliCloud.GetRegions.Invoke(new()
    ///     {
    ///         Current = true,
    ///     });
    /// 
    ///     var example = new AliCloud.Ga.Accelerator("example", new()
    ///     {
    ///         Duration = 3,
    ///         Spec = "2",
    ///         AcceleratorName = name,
    ///         AutoUseCoupon = false,
    ///         Description = name,
    ///         AutoRenewDuration = 2,
    ///         RenewalStatus = "AutoRenewal",
    ///     });
    /// 
    ///     var exampleBandwidthPackage = new AliCloud.Ga.BandwidthPackage("example", new()
    ///     {
    ///         Type = "Basic",
    ///         Bandwidth = 20,
    ///         BandwidthType = "Basic",
    ///         Duration = "1",
    ///         AutoPay = true,
    ///         PaymentType = "Subscription",
    ///         AutoUseCoupon = false,
    ///         BandwidthPackageName = name,
    ///         Description = name,
    ///     });
    /// 
    ///     var exampleBandwidthPackageAttachment = new AliCloud.Ga.BandwidthPackageAttachment("example", new()
    ///     {
    ///         AcceleratorId = example.Id,
    ///         BandwidthPackageId = exampleBandwidthPackage.Id,
    ///     });
    /// 
    ///     var exampleListener = new AliCloud.Ga.Listener("example", new()
    ///     {
    ///         AcceleratorId = exampleBandwidthPackageAttachment.AcceleratorId,
    ///         ClientAffinity = "SOURCE_IP",
    ///         Description = name,
    ///         Name = name,
    ///         Protocol = "HTTP",
    ///         ProxyProtocol = true,
    ///         PortRanges = new[]
    ///         {
    ///             new AliCloud.Ga.Inputs.ListenerPortRangeArgs
    ///             {
    ///                 FromPort = 60,
    ///                 ToPort = 60,
    ///             },
    ///         },
    ///     });
    /// 
    ///     var exampleEipAddress = new AliCloud.Ecs.EipAddress("example", new()
    ///     {
    ///         Bandwidth = "10",
    ///         InternetChargeType = "PayByBandwidth",
    ///     });
    /// 
    ///     var @virtual = new AliCloud.Ga.EndpointGroup("virtual", new()
    ///     {
    ///         AcceleratorId = example.Id,
    ///         EndpointConfigurations = new[]
    ///         {
    ///             new AliCloud.Ga.Inputs.EndpointGroupEndpointConfigurationArgs
    ///             {
    ///                 Endpoint = exampleEipAddress.IpAddress,
    ///                 Type = "PublicIp",
    ///                 Weight = 20,
    ///                 EnableClientipPreservation = true,
    ///             },
    ///         },
    ///         EndpointGroupRegion = @default.Apply(@default =&gt; @default.Apply(getRegionsResult =&gt; getRegionsResult.Regions[0]?.Id)),
    ///         ListenerId = exampleListener.Id,
    ///         Description = name,
    ///         EndpointGroupType = "virtual",
    ///         EndpointRequestProtocol = "HTTPS",
    ///         HealthCheckIntervalSeconds = 4,
    ///         HealthCheckPath = "/path",
    ///         Name = name,
    ///         ThresholdCount = 4,
    ///         TrafficPercentage = 20,
    ///         PortOverrides = new AliCloud.Ga.Inputs.EndpointGroupPortOverridesArgs
    ///         {
    ///             EndpointPort = 80,
    ///             ListenerPort = 60,
    ///         },
    ///     });
    /// 
    ///     var exampleForwardingRule = new AliCloud.Ga.ForwardingRule("example", new()
    ///     {
    ///         AcceleratorId = example.Id,
    ///         ListenerId = exampleListener.Id,
    ///         RuleConditions = new[]
    ///         {
    ///             new AliCloud.Ga.Inputs.ForwardingRuleRuleConditionArgs
    ///             {
    ///                 RuleConditionType = "Path",
    ///                 PathConfig = new AliCloud.Ga.Inputs.ForwardingRuleRuleConditionPathConfigArgs
    ///                 {
    ///                     Values = new[]
    ///                     {
    ///                         "/testpathconfig",
    ///                     },
    ///                 },
    ///             },
    ///             new AliCloud.Ga.Inputs.ForwardingRuleRuleConditionArgs
    ///             {
    ///                 RuleConditionType = "Host",
    ///                 HostConfigs = new[]
    ///                 {
    ///                     new AliCloud.Ga.Inputs.ForwardingRuleRuleConditionHostConfigArgs
    ///                     {
    ///                         Values = new[]
    ///                         {
    ///                             "www.test.com",
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///         RuleActions = new[]
    ///         {
    ///             new AliCloud.Ga.Inputs.ForwardingRuleRuleActionArgs
    ///             {
    ///                 Order = 40,
    ///                 RuleActionType = "ForwardGroup",
    ///                 ForwardGroupConfig = new AliCloud.Ga.Inputs.ForwardingRuleRuleActionForwardGroupConfigArgs
    ///                 {
    ///                     ServerGroupTuples = new[]
    ///                     {
    ///                         new AliCloud.Ga.Inputs.ForwardingRuleRuleActionForwardGroupConfigServerGroupTupleArgs
    ///                         {
    ///                             EndpointGroupId = @virtual.Id,
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///         Priority = 2,
    ///         ForwardingRuleName = name,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Ga Forwarding Rule can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:ga/forwardingRule:ForwardingRule example &lt;accelerator_id&gt;:&lt;listener_id&gt;:&lt;forwarding_rule_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:ga/forwardingRule:ForwardingRule")]
    public partial class ForwardingRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the Global Accelerator instance.
        /// </summary>
        [Output("acceleratorId")]
        public Output<string> AcceleratorId { get; private set; } = null!;

        /// <summary>
        /// Forwarding Policy ID.
        /// </summary>
        [Output("forwardingRuleId")]
        public Output<string> ForwardingRuleId { get; private set; } = null!;

        /// <summary>
        /// Forwarding policy name. The length of the name is 2-128 English or Chinese characters. It must start with uppercase and lowercase letters or Chinese characters. It can contain numbers, half width period (.), underscores (_) And dash (-).
        /// </summary>
        [Output("forwardingRuleName")]
        public Output<string?> ForwardingRuleName { get; private set; } = null!;

        /// <summary>
        /// Forwarding Policy Status.
        /// </summary>
        [Output("forwardingRuleStatus")]
        public Output<string> ForwardingRuleStatus { get; private set; } = null!;

        /// <summary>
        /// The ID of the listener.
        /// </summary>
        [Output("listenerId")]
        public Output<string> ListenerId { get; private set; } = null!;

        /// <summary>
        /// Forwarding policy priority.
        /// </summary>
        [Output("priority")]
        public Output<int> Priority { get; private set; } = null!;

        /// <summary>
        /// Forward action. See `rule_actions` below.
        /// </summary>
        [Output("ruleActions")]
        public Output<ImmutableArray<Outputs.ForwardingRuleRuleAction>> RuleActions { get; private set; } = null!;

        /// <summary>
        /// Forwarding condition list. See `rule_conditions` below.
        /// </summary>
        [Output("ruleConditions")]
        public Output<ImmutableArray<Outputs.ForwardingRuleRuleCondition>> RuleConditions { get; private set; } = null!;


        /// <summary>
        /// Create a ForwardingRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ForwardingRule(string name, ForwardingRuleArgs args, CustomResourceOptions? options = null)
            : base("alicloud:ga/forwardingRule:ForwardingRule", name, args ?? new ForwardingRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ForwardingRule(string name, Input<string> id, ForwardingRuleState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ga/forwardingRule:ForwardingRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ForwardingRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ForwardingRule Get(string name, Input<string> id, ForwardingRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new ForwardingRule(name, id, state, options);
        }
    }

    public sealed class ForwardingRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Global Accelerator instance.
        /// </summary>
        [Input("acceleratorId", required: true)]
        public Input<string> AcceleratorId { get; set; } = null!;

        /// <summary>
        /// Forwarding policy name. The length of the name is 2-128 English or Chinese characters. It must start with uppercase and lowercase letters or Chinese characters. It can contain numbers, half width period (.), underscores (_) And dash (-).
        /// </summary>
        [Input("forwardingRuleName")]
        public Input<string>? ForwardingRuleName { get; set; }

        /// <summary>
        /// The ID of the listener.
        /// </summary>
        [Input("listenerId", required: true)]
        public Input<string> ListenerId { get; set; } = null!;

        /// <summary>
        /// Forwarding policy priority.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        [Input("ruleActions", required: true)]
        private InputList<Inputs.ForwardingRuleRuleActionArgs>? _ruleActions;

        /// <summary>
        /// Forward action. See `rule_actions` below.
        /// </summary>
        public InputList<Inputs.ForwardingRuleRuleActionArgs> RuleActions
        {
            get => _ruleActions ?? (_ruleActions = new InputList<Inputs.ForwardingRuleRuleActionArgs>());
            set => _ruleActions = value;
        }

        [Input("ruleConditions", required: true)]
        private InputList<Inputs.ForwardingRuleRuleConditionArgs>? _ruleConditions;

        /// <summary>
        /// Forwarding condition list. See `rule_conditions` below.
        /// </summary>
        public InputList<Inputs.ForwardingRuleRuleConditionArgs> RuleConditions
        {
            get => _ruleConditions ?? (_ruleConditions = new InputList<Inputs.ForwardingRuleRuleConditionArgs>());
            set => _ruleConditions = value;
        }

        public ForwardingRuleArgs()
        {
        }
        public static new ForwardingRuleArgs Empty => new ForwardingRuleArgs();
    }

    public sealed class ForwardingRuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Global Accelerator instance.
        /// </summary>
        [Input("acceleratorId")]
        public Input<string>? AcceleratorId { get; set; }

        /// <summary>
        /// Forwarding Policy ID.
        /// </summary>
        [Input("forwardingRuleId")]
        public Input<string>? ForwardingRuleId { get; set; }

        /// <summary>
        /// Forwarding policy name. The length of the name is 2-128 English or Chinese characters. It must start with uppercase and lowercase letters or Chinese characters. It can contain numbers, half width period (.), underscores (_) And dash (-).
        /// </summary>
        [Input("forwardingRuleName")]
        public Input<string>? ForwardingRuleName { get; set; }

        /// <summary>
        /// Forwarding Policy Status.
        /// </summary>
        [Input("forwardingRuleStatus")]
        public Input<string>? ForwardingRuleStatus { get; set; }

        /// <summary>
        /// The ID of the listener.
        /// </summary>
        [Input("listenerId")]
        public Input<string>? ListenerId { get; set; }

        /// <summary>
        /// Forwarding policy priority.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        [Input("ruleActions")]
        private InputList<Inputs.ForwardingRuleRuleActionGetArgs>? _ruleActions;

        /// <summary>
        /// Forward action. See `rule_actions` below.
        /// </summary>
        public InputList<Inputs.ForwardingRuleRuleActionGetArgs> RuleActions
        {
            get => _ruleActions ?? (_ruleActions = new InputList<Inputs.ForwardingRuleRuleActionGetArgs>());
            set => _ruleActions = value;
        }

        [Input("ruleConditions")]
        private InputList<Inputs.ForwardingRuleRuleConditionGetArgs>? _ruleConditions;

        /// <summary>
        /// Forwarding condition list. See `rule_conditions` below.
        /// </summary>
        public InputList<Inputs.ForwardingRuleRuleConditionGetArgs> RuleConditions
        {
            get => _ruleConditions ?? (_ruleConditions = new InputList<Inputs.ForwardingRuleRuleConditionGetArgs>());
            set => _ruleConditions = value;
        }

        public ForwardingRuleState()
        {
        }
        public static new ForwardingRuleState Empty => new ForwardingRuleState();
    }
}

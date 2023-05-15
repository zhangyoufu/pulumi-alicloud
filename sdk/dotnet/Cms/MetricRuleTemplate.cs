// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cms
{
    /// <summary>
    /// Provides a Cloud Monitor Service Metric Rule Template resource.
    /// 
    /// For information about Cloud Monitor Service Metric Rule Template and how to use it, see [What is Metric Rule Template](https://www.alibabacloud.com/help/doc-detail/114984.html).
    /// 
    /// &gt; **NOTE:** Available in v1.134.0+.
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
    ///     var example = new AliCloud.Cms.MetricRuleTemplate("example", new()
    ///     {
    ///         AlertTemplates = new[]
    ///         {
    ///             new AliCloud.Cms.Inputs.MetricRuleTemplateAlertTemplateArgs
    ///             {
    ///                 Category = "ecs",
    ///                 Escalations = new AliCloud.Cms.Inputs.MetricRuleTemplateAlertTemplateEscalationsArgs
    ///                 {
    ///                     Critical = new AliCloud.Cms.Inputs.MetricRuleTemplateAlertTemplateEscalationsCriticalArgs
    ///                     {
    ///                         ComparisonOperator = "GreaterThanThreshold",
    ///                         Statistics = "Average",
    ///                         Threshold = "90",
    ///                         Times = "3",
    ///                     },
    ///                 },
    ///                 MetricName = "cpu_total",
    ///                 Namespace = "acs_ecs_dashboard",
    ///                 RuleName = "tf_testAcc_new",
    ///             },
    ///         },
    ///         MetricRuleTemplateName = "example_value",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Cloud Monitor Service Metric Rule Template can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:cms/metricRuleTemplate:MetricRuleTemplate example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cms/metricRuleTemplate:MetricRuleTemplate")]
    public partial class MetricRuleTemplate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The details of alert rules that are generated based on the alert template. See the following `Block alert_templates`.
        /// </summary>
        [Output("alertTemplates")]
        public Output<ImmutableArray<Outputs.MetricRuleTemplateAlertTemplate>> AlertTemplates { get; private set; } = null!;

        /// <summary>
        /// The mode in which the alert template is applied. Valid values:`GROUP_INSTANCE_FIRST`or `ALARM_TEMPLATE_FIRST`. GROUP_INSTANCE_FIRST: The metrics in the application group take precedence. If a metric specified in the alert template does not exist in the application group, the system does not generate an alert rule for the metric based on the alert template. ALARM_TEMPLATE_FIRST: The metrics specified in the alert template take precedence. If a metric specified in the alert template does not exist in the application group, the system still generates an alert rule for the metric based on the alert template.
        /// </summary>
        [Output("applyMode")]
        public Output<string?> ApplyMode { get; private set; } = null!;

        /// <summary>
        /// The description of the alert template.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The end of the time period during which the alert rule is effective. Valid values: 00 to 23. The value 00 indicates 00:59 and the value 23 indicates 23:59.
        /// </summary>
        [Output("enableEndTime")]
        public Output<string?> EnableEndTime { get; private set; } = null!;

        /// <summary>
        /// The beginning of the time period during which the alert rule is effective. Valid values: 00 to 23. The value 00 indicates 00:00 and the value 23 indicates 23:00.
        /// </summary>
        [Output("enableStartTime")]
        public Output<string?> EnableStartTime { get; private set; } = null!;

        /// <summary>
        /// The ID of the application group.
        /// </summary>
        [Output("groupId")]
        public Output<string?> GroupId { get; private set; } = null!;

        /// <summary>
        /// The name of the alert template.
        /// </summary>
        [Output("metricRuleTemplateName")]
        public Output<string> MetricRuleTemplateName { get; private set; } = null!;

        /// <summary>
        /// The alert notification method. Valid values:Set the value to 4. The value 4 indicates that alert notifications are sent by using TradeManager and DingTalk chatbots.
        /// </summary>
        [Output("notifyLevel")]
        public Output<string?> NotifyLevel { get; private set; } = null!;

        /// <summary>
        /// The version of the alert template to be modified.
        /// 
        /// &gt; **NOTE:** The version changes with the number of times that the alert template is modified.
        /// </summary>
        [Output("restVersion")]
        public Output<string> RestVersion { get; private set; } = null!;

        /// <summary>
        /// The mute period during which notifications are not repeatedly sent for an alert.Valid values: 0 to 86400. Unit: seconds. Default value: `86400`.
        /// 
        /// &gt; **NOTE:** Only one alert notification is sent during each mute period even if the metric value exceeds the alert threshold several times.
        /// </summary>
        [Output("silenceTime")]
        public Output<int?> SilenceTime { get; private set; } = null!;

        /// <summary>
        /// The callback URL to which a POST request is sent when an alert is triggered based on the alert rule.
        /// </summary>
        [Output("webhook")]
        public Output<string?> Webhook { get; private set; } = null!;


        /// <summary>
        /// Create a MetricRuleTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MetricRuleTemplate(string name, MetricRuleTemplateArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cms/metricRuleTemplate:MetricRuleTemplate", name, args ?? new MetricRuleTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MetricRuleTemplate(string name, Input<string> id, MetricRuleTemplateState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cms/metricRuleTemplate:MetricRuleTemplate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MetricRuleTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MetricRuleTemplate Get(string name, Input<string> id, MetricRuleTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new MetricRuleTemplate(name, id, state, options);
        }
    }

    public sealed class MetricRuleTemplateArgs : global::Pulumi.ResourceArgs
    {
        [Input("alertTemplates")]
        private InputList<Inputs.MetricRuleTemplateAlertTemplateArgs>? _alertTemplates;

        /// <summary>
        /// The details of alert rules that are generated based on the alert template. See the following `Block alert_templates`.
        /// </summary>
        public InputList<Inputs.MetricRuleTemplateAlertTemplateArgs> AlertTemplates
        {
            get => _alertTemplates ?? (_alertTemplates = new InputList<Inputs.MetricRuleTemplateAlertTemplateArgs>());
            set => _alertTemplates = value;
        }

        /// <summary>
        /// The mode in which the alert template is applied. Valid values:`GROUP_INSTANCE_FIRST`or `ALARM_TEMPLATE_FIRST`. GROUP_INSTANCE_FIRST: The metrics in the application group take precedence. If a metric specified in the alert template does not exist in the application group, the system does not generate an alert rule for the metric based on the alert template. ALARM_TEMPLATE_FIRST: The metrics specified in the alert template take precedence. If a metric specified in the alert template does not exist in the application group, the system still generates an alert rule for the metric based on the alert template.
        /// </summary>
        [Input("applyMode")]
        public Input<string>? ApplyMode { get; set; }

        /// <summary>
        /// The description of the alert template.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The end of the time period during which the alert rule is effective. Valid values: 00 to 23. The value 00 indicates 00:59 and the value 23 indicates 23:59.
        /// </summary>
        [Input("enableEndTime")]
        public Input<string>? EnableEndTime { get; set; }

        /// <summary>
        /// The beginning of the time period during which the alert rule is effective. Valid values: 00 to 23. The value 00 indicates 00:00 and the value 23 indicates 23:00.
        /// </summary>
        [Input("enableStartTime")]
        public Input<string>? EnableStartTime { get; set; }

        /// <summary>
        /// The ID of the application group.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// The name of the alert template.
        /// </summary>
        [Input("metricRuleTemplateName", required: true)]
        public Input<string> MetricRuleTemplateName { get; set; } = null!;

        /// <summary>
        /// The alert notification method. Valid values:Set the value to 4. The value 4 indicates that alert notifications are sent by using TradeManager and DingTalk chatbots.
        /// </summary>
        [Input("notifyLevel")]
        public Input<string>? NotifyLevel { get; set; }

        /// <summary>
        /// The version of the alert template to be modified.
        /// 
        /// &gt; **NOTE:** The version changes with the number of times that the alert template is modified.
        /// </summary>
        [Input("restVersion")]
        public Input<string>? RestVersion { get; set; }

        /// <summary>
        /// The mute period during which notifications are not repeatedly sent for an alert.Valid values: 0 to 86400. Unit: seconds. Default value: `86400`.
        /// 
        /// &gt; **NOTE:** Only one alert notification is sent during each mute period even if the metric value exceeds the alert threshold several times.
        /// </summary>
        [Input("silenceTime")]
        public Input<int>? SilenceTime { get; set; }

        /// <summary>
        /// The callback URL to which a POST request is sent when an alert is triggered based on the alert rule.
        /// </summary>
        [Input("webhook")]
        public Input<string>? Webhook { get; set; }

        public MetricRuleTemplateArgs()
        {
        }
        public static new MetricRuleTemplateArgs Empty => new MetricRuleTemplateArgs();
    }

    public sealed class MetricRuleTemplateState : global::Pulumi.ResourceArgs
    {
        [Input("alertTemplates")]
        private InputList<Inputs.MetricRuleTemplateAlertTemplateGetArgs>? _alertTemplates;

        /// <summary>
        /// The details of alert rules that are generated based on the alert template. See the following `Block alert_templates`.
        /// </summary>
        public InputList<Inputs.MetricRuleTemplateAlertTemplateGetArgs> AlertTemplates
        {
            get => _alertTemplates ?? (_alertTemplates = new InputList<Inputs.MetricRuleTemplateAlertTemplateGetArgs>());
            set => _alertTemplates = value;
        }

        /// <summary>
        /// The mode in which the alert template is applied. Valid values:`GROUP_INSTANCE_FIRST`or `ALARM_TEMPLATE_FIRST`. GROUP_INSTANCE_FIRST: The metrics in the application group take precedence. If a metric specified in the alert template does not exist in the application group, the system does not generate an alert rule for the metric based on the alert template. ALARM_TEMPLATE_FIRST: The metrics specified in the alert template take precedence. If a metric specified in the alert template does not exist in the application group, the system still generates an alert rule for the metric based on the alert template.
        /// </summary>
        [Input("applyMode")]
        public Input<string>? ApplyMode { get; set; }

        /// <summary>
        /// The description of the alert template.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The end of the time period during which the alert rule is effective. Valid values: 00 to 23. The value 00 indicates 00:59 and the value 23 indicates 23:59.
        /// </summary>
        [Input("enableEndTime")]
        public Input<string>? EnableEndTime { get; set; }

        /// <summary>
        /// The beginning of the time period during which the alert rule is effective. Valid values: 00 to 23. The value 00 indicates 00:00 and the value 23 indicates 23:00.
        /// </summary>
        [Input("enableStartTime")]
        public Input<string>? EnableStartTime { get; set; }

        /// <summary>
        /// The ID of the application group.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// The name of the alert template.
        /// </summary>
        [Input("metricRuleTemplateName")]
        public Input<string>? MetricRuleTemplateName { get; set; }

        /// <summary>
        /// The alert notification method. Valid values:Set the value to 4. The value 4 indicates that alert notifications are sent by using TradeManager and DingTalk chatbots.
        /// </summary>
        [Input("notifyLevel")]
        public Input<string>? NotifyLevel { get; set; }

        /// <summary>
        /// The version of the alert template to be modified.
        /// 
        /// &gt; **NOTE:** The version changes with the number of times that the alert template is modified.
        /// </summary>
        [Input("restVersion")]
        public Input<string>? RestVersion { get; set; }

        /// <summary>
        /// The mute period during which notifications are not repeatedly sent for an alert.Valid values: 0 to 86400. Unit: seconds. Default value: `86400`.
        /// 
        /// &gt; **NOTE:** Only one alert notification is sent during each mute period even if the metric value exceeds the alert threshold several times.
        /// </summary>
        [Input("silenceTime")]
        public Input<int>? SilenceTime { get; set; }

        /// <summary>
        /// The callback URL to which a POST request is sent when an alert is triggered based on the alert rule.
        /// </summary>
        [Input("webhook")]
        public Input<string>? Webhook { get; set; }

        public MetricRuleTemplateState()
        {
        }
        public static new MetricRuleTemplateState Empty => new MetricRuleTemplateState();
    }
}

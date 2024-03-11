// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.SimpleApplicationServer
{
    /// <summary>
    /// Provides a Simple Application Server Firewall Rule resource.
    /// 
    /// For information about Simple Application Server Firewall Rule and how to use it, see [What is Firewall Rule](https://www.alibabacloud.com/help/doc-detail/190449.htm).
    /// 
    /// &gt; **NOTE:** Available since v1.143.0.
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
    ///     var name = config.Get("name") ?? "tf_example";
    ///     var defaultImages = AliCloud.SimpleApplicationServer.GetImages.Invoke();
    /// 
    ///     var defaultServerPlans = AliCloud.SimpleApplicationServer.GetServerPlans.Invoke();
    /// 
    ///     var defaultInstance = new AliCloud.SimpleApplicationServer.Instance("defaultInstance", new()
    ///     {
    ///         PaymentType = "Subscription",
    ///         PlanId = defaultServerPlans.Apply(getServerPlansResult =&gt; getServerPlansResult.Plans[0]?.Id),
    ///         InstanceName = name,
    ///         ImageId = defaultImages.Apply(getImagesResult =&gt; getImagesResult.Images[0]?.Id),
    ///         Period = 1,
    ///         DataDiskSize = 100,
    ///     });
    /// 
    ///     var defaultFirewallRule = new AliCloud.SimpleApplicationServer.FirewallRule("defaultFirewallRule", new()
    ///     {
    ///         InstanceId = defaultInstance.Id,
    ///         RuleProtocol = "Tcp",
    ///         Port = "9999",
    ///         Remark = name,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Simple Application Server Firewall Rule can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:simpleapplicationserver/firewallRule:FirewallRule example &lt;instance_id&gt;:&lt;firewall_rule_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:simpleapplicationserver/firewallRule:FirewallRule")]
    public partial class FirewallRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the firewall rule.
        /// </summary>
        [Output("firewallRuleId")]
        public Output<string> FirewallRuleId { get; private set; } = null!;

        /// <summary>
        /// Alibaba Cloud simple application server instance ID.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The port range. Valid values of port numbers: `1` to `65535`. Specify a port range in the format of `&lt;start port number&gt;/&lt;end port number&gt;`. Example: `1024/1055`, which indicates the port range of `1024` through `1055`.
        /// </summary>
        [Output("port")]
        public Output<string> Port { get; private set; } = null!;

        /// <summary>
        /// The remarks of the firewall rule.
        /// </summary>
        [Output("remark")]
        public Output<string?> Remark { get; private set; } = null!;

        /// <summary>
        /// The transport layer protocol. Valid values: `Tcp`, `Udp`, `TcpAndUdp`.
        /// </summary>
        [Output("ruleProtocol")]
        public Output<string> RuleProtocol { get; private set; } = null!;


        /// <summary>
        /// Create a FirewallRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FirewallRule(string name, FirewallRuleArgs args, CustomResourceOptions? options = null)
            : base("alicloud:simpleapplicationserver/firewallRule:FirewallRule", name, args ?? new FirewallRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FirewallRule(string name, Input<string> id, FirewallRuleState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:simpleapplicationserver/firewallRule:FirewallRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FirewallRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FirewallRule Get(string name, Input<string> id, FirewallRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new FirewallRule(name, id, state, options);
        }
    }

    public sealed class FirewallRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Alibaba Cloud simple application server instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The port range. Valid values of port numbers: `1` to `65535`. Specify a port range in the format of `&lt;start port number&gt;/&lt;end port number&gt;`. Example: `1024/1055`, which indicates the port range of `1024` through `1055`.
        /// </summary>
        [Input("port", required: true)]
        public Input<string> Port { get; set; } = null!;

        /// <summary>
        /// The remarks of the firewall rule.
        /// </summary>
        [Input("remark")]
        public Input<string>? Remark { get; set; }

        /// <summary>
        /// The transport layer protocol. Valid values: `Tcp`, `Udp`, `TcpAndUdp`.
        /// </summary>
        [Input("ruleProtocol", required: true)]
        public Input<string> RuleProtocol { get; set; } = null!;

        public FirewallRuleArgs()
        {
        }
        public static new FirewallRuleArgs Empty => new FirewallRuleArgs();
    }

    public sealed class FirewallRuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the firewall rule.
        /// </summary>
        [Input("firewallRuleId")]
        public Input<string>? FirewallRuleId { get; set; }

        /// <summary>
        /// Alibaba Cloud simple application server instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The port range. Valid values of port numbers: `1` to `65535`. Specify a port range in the format of `&lt;start port number&gt;/&lt;end port number&gt;`. Example: `1024/1055`, which indicates the port range of `1024` through `1055`.
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        /// <summary>
        /// The remarks of the firewall rule.
        /// </summary>
        [Input("remark")]
        public Input<string>? Remark { get; set; }

        /// <summary>
        /// The transport layer protocol. Valid values: `Tcp`, `Udp`, `TcpAndUdp`.
        /// </summary>
        [Input("ruleProtocol")]
        public Input<string>? RuleProtocol { get; set; }

        public FirewallRuleState()
        {
        }
        public static new FirewallRuleState Empty => new FirewallRuleState();
    }
}

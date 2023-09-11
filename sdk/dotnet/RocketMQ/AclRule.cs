// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.RocketMQ
{
    /// <summary>
    /// Provides a Sag Acl Rule resource. This topic describes how to configure an access control list (ACL) rule for a target Smart Access Gateway instance to permit or deny access to or from specified IP addresses in the ACL rule.
    /// 
    /// For information about Sag Acl Rule and how to use it, see [What is access control list (ACL) rule](https://www.alibabacloud.com/help/en/smart-access-gateway/latest/addaclrule).
    /// 
    /// &gt; **NOTE:** Available since v1.60.0.
    /// 
    /// &gt; **NOTE:** Only the following regions support create Cloud Connect Network. [`cn-shanghai`, `cn-shanghai-finance-1`, `cn-hongkong`, `ap-southeast-1`, `ap-southeast-2`, `ap-southeast-3`, `ap-southeast-5`, `ap-northeast-1`, `eu-central-1`]
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
    ///     var name = config.Get("name") ?? "tf_example";
    ///     var defaultAcl = new AliCloud.RocketMQ.Acl("defaultAcl");
    /// 
    ///     var defaultAclRule = new AliCloud.RocketMQ.AclRule("defaultAclRule", new()
    ///     {
    ///         AclId = defaultAcl.Id,
    ///         Description = name,
    ///         Policy = "accept",
    ///         IpProtocol = "ALL",
    ///         Direction = "in",
    ///         SourceCidr = "10.10.1.0/24",
    ///         SourcePortRange = "-1/-1",
    ///         DestCidr = "192.168.1.0/24",
    ///         DestPortRange = "-1/-1",
    ///         Priority = 1,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// The Sag Acl Rule can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:rocketmq/aclRule:AclRule example acr-abc123456
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:rocketmq/aclRule:AclRule")]
    public partial class AclRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the ACL.
        /// </summary>
        [Output("aclId")]
        public Output<string> AclId { get; private set; } = null!;

        /// <summary>
        /// The description of the ACL rule. It must be 1 to 512 characters in length.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The destination address. It is an IPv4 address range in CIDR format. Default value: 0.0.0.0/0.
        /// </summary>
        [Output("destCidr")]
        public Output<string> DestCidr { get; private set; } = null!;

        /// <summary>
        /// The range of the destination port. Valid value: 80/80.
        /// </summary>
        [Output("destPortRange")]
        public Output<string> DestPortRange { get; private set; } = null!;

        /// <summary>
        /// The direction of the ACL rule. Valid values: in|out.
        /// </summary>
        [Output("direction")]
        public Output<string> Direction { get; private set; } = null!;

        /// <summary>
        /// The protocol used by the ACL rule. The value is not case sensitive.
        /// </summary>
        [Output("ipProtocol")]
        public Output<string> IpProtocol { get; private set; } = null!;

        /// <summary>
        /// The policy used by the ACL rule. Valid values: accept|drop.
        /// </summary>
        [Output("policy")]
        public Output<string> Policy { get; private set; } = null!;

        /// <summary>
        /// The priority of the ACL rule. Value range: 1 to 100.
        /// </summary>
        [Output("priority")]
        public Output<int?> Priority { get; private set; } = null!;

        /// <summary>
        /// The source address. It is an IPv4 address range in the CIDR format. Default value: 0.0.0.0/0.
        /// </summary>
        [Output("sourceCidr")]
        public Output<string> SourceCidr { get; private set; } = null!;

        /// <summary>
        /// The range of the source port. Valid value: 80/80.
        /// </summary>
        [Output("sourcePortRange")]
        public Output<string> SourcePortRange { get; private set; } = null!;


        /// <summary>
        /// Create a AclRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AclRule(string name, AclRuleArgs args, CustomResourceOptions? options = null)
            : base("alicloud:rocketmq/aclRule:AclRule", name, args ?? new AclRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AclRule(string name, Input<string> id, AclRuleState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:rocketmq/aclRule:AclRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AclRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AclRule Get(string name, Input<string> id, AclRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new AclRule(name, id, state, options);
        }
    }

    public sealed class AclRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the ACL.
        /// </summary>
        [Input("aclId", required: true)]
        public Input<string> AclId { get; set; } = null!;

        /// <summary>
        /// The description of the ACL rule. It must be 1 to 512 characters in length.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The destination address. It is an IPv4 address range in CIDR format. Default value: 0.0.0.0/0.
        /// </summary>
        [Input("destCidr", required: true)]
        public Input<string> DestCidr { get; set; } = null!;

        /// <summary>
        /// The range of the destination port. Valid value: 80/80.
        /// </summary>
        [Input("destPortRange", required: true)]
        public Input<string> DestPortRange { get; set; } = null!;

        /// <summary>
        /// The direction of the ACL rule. Valid values: in|out.
        /// </summary>
        [Input("direction", required: true)]
        public Input<string> Direction { get; set; } = null!;

        /// <summary>
        /// The protocol used by the ACL rule. The value is not case sensitive.
        /// </summary>
        [Input("ipProtocol", required: true)]
        public Input<string> IpProtocol { get; set; } = null!;

        /// <summary>
        /// The policy used by the ACL rule. Valid values: accept|drop.
        /// </summary>
        [Input("policy", required: true)]
        public Input<string> Policy { get; set; } = null!;

        /// <summary>
        /// The priority of the ACL rule. Value range: 1 to 100.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The source address. It is an IPv4 address range in the CIDR format. Default value: 0.0.0.0/0.
        /// </summary>
        [Input("sourceCidr", required: true)]
        public Input<string> SourceCidr { get; set; } = null!;

        /// <summary>
        /// The range of the source port. Valid value: 80/80.
        /// </summary>
        [Input("sourcePortRange", required: true)]
        public Input<string> SourcePortRange { get; set; } = null!;

        public AclRuleArgs()
        {
        }
        public static new AclRuleArgs Empty => new AclRuleArgs();
    }

    public sealed class AclRuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the ACL.
        /// </summary>
        [Input("aclId")]
        public Input<string>? AclId { get; set; }

        /// <summary>
        /// The description of the ACL rule. It must be 1 to 512 characters in length.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The destination address. It is an IPv4 address range in CIDR format. Default value: 0.0.0.0/0.
        /// </summary>
        [Input("destCidr")]
        public Input<string>? DestCidr { get; set; }

        /// <summary>
        /// The range of the destination port. Valid value: 80/80.
        /// </summary>
        [Input("destPortRange")]
        public Input<string>? DestPortRange { get; set; }

        /// <summary>
        /// The direction of the ACL rule. Valid values: in|out.
        /// </summary>
        [Input("direction")]
        public Input<string>? Direction { get; set; }

        /// <summary>
        /// The protocol used by the ACL rule. The value is not case sensitive.
        /// </summary>
        [Input("ipProtocol")]
        public Input<string>? IpProtocol { get; set; }

        /// <summary>
        /// The policy used by the ACL rule. Valid values: accept|drop.
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        /// <summary>
        /// The priority of the ACL rule. Value range: 1 to 100.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The source address. It is an IPv4 address range in the CIDR format. Default value: 0.0.0.0/0.
        /// </summary>
        [Input("sourceCidr")]
        public Input<string>? SourceCidr { get; set; }

        /// <summary>
        /// The range of the source port. Valid value: 80/80.
        /// </summary>
        [Input("sourcePortRange")]
        public Input<string>? SourcePortRange { get; set; }

        public AclRuleState()
        {
        }
        public static new AclRuleState Empty => new AclRuleState();
    }
}

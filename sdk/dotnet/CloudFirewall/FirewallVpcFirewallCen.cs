// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CloudFirewall
{
    /// <summary>
    /// Provides a Cloud Firewall Vpc Firewall Cen resource.
    /// 
    /// For information about Cloud Firewall Vpc Firewall Cen and how to use it, see [What is Vpc Firewall Cen](https://www.alibabacloud.com/help/en/cloud-firewall/latest/createvpcfirewallcenconfigure).
    /// 
    /// &gt; **NOTE:** Available since v1.194.0.
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
    ///     // These resource primary keys should be replaced with your actual values.
    ///     var @default = new AliCloud.CloudFirewall.FirewallVpcFirewallCen("default", new()
    ///     {
    ///         CenId = "cen-xxx",
    ///         LocalVpc = new AliCloud.CloudFirewall.Inputs.FirewallVpcFirewallCenLocalVpcArgs
    ///         {
    ///             NetworkInstanceId = "vpc-xxx",
    ///         },
    ///         Status = "open",
    ///         MemberUid = "14151*****827022",
    ///         VpcRegion = "cn-hangzhou",
    ///         VpcFirewallName = "tf-vpc-firewall-name",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Cloud Firewall Vpc Firewall Cen can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:cloudfirewall/firewallVpcFirewallCen:FirewallVpcFirewallCen example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cloudfirewall/firewallVpcFirewallCen:FirewallVpcFirewallCen")]
    public partial class FirewallVpcFirewallCen : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the CEN instance.
        /// </summary>
        [Output("cenId")]
        public Output<string> CenId { get; private set; } = null!;

        /// <summary>
        /// Intercommunication type, value: expressconnect: Express Channel cen: Cloud Enterprise Network
        /// </summary>
        [Output("connectType")]
        public Output<string> ConnectType { get; private set; } = null!;

        /// <summary>
        /// The language type of the requested and received messages. Valid values:
        /// </summary>
        [Output("lang")]
        public Output<string?> Lang { get; private set; } = null!;

        /// <summary>
        /// The details of the VPC. See `local_vpc` below.
        /// </summary>
        [Output("localVpc")]
        public Output<Outputs.FirewallVpcFirewallCenLocalVpc> LocalVpc { get; private set; } = null!;

        /// <summary>
        /// The UID of the member account (other Alibaba Cloud account) of the current Alibaba cloud account.
        /// </summary>
        [Output("memberUid")]
        public Output<string?> MemberUid { get; private set; } = null!;

        /// <summary>
        /// Firewall switch status.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// VPC firewall ID
        /// </summary>
        [Output("vpcFirewallId")]
        public Output<string> VpcFirewallId { get; private set; } = null!;

        /// <summary>
        /// The name of the VPC firewall instance.
        /// </summary>
        [Output("vpcFirewallName")]
        public Output<string> VpcFirewallName { get; private set; } = null!;

        /// <summary>
        /// The ID of the region to which the VPC is created.
        /// </summary>
        [Output("vpcRegion")]
        public Output<string> VpcRegion { get; private set; } = null!;


        /// <summary>
        /// Create a FirewallVpcFirewallCen resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FirewallVpcFirewallCen(string name, FirewallVpcFirewallCenArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cloudfirewall/firewallVpcFirewallCen:FirewallVpcFirewallCen", name, args ?? new FirewallVpcFirewallCenArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FirewallVpcFirewallCen(string name, Input<string> id, FirewallVpcFirewallCenState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cloudfirewall/firewallVpcFirewallCen:FirewallVpcFirewallCen", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FirewallVpcFirewallCen resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FirewallVpcFirewallCen Get(string name, Input<string> id, FirewallVpcFirewallCenState? state = null, CustomResourceOptions? options = null)
        {
            return new FirewallVpcFirewallCen(name, id, state, options);
        }
    }

    public sealed class FirewallVpcFirewallCenArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the CEN instance.
        /// </summary>
        [Input("cenId", required: true)]
        public Input<string> CenId { get; set; } = null!;

        /// <summary>
        /// The language type of the requested and received messages. Valid values:
        /// </summary>
        [Input("lang")]
        public Input<string>? Lang { get; set; }

        /// <summary>
        /// The details of the VPC. See `local_vpc` below.
        /// </summary>
        [Input("localVpc", required: true)]
        public Input<Inputs.FirewallVpcFirewallCenLocalVpcArgs> LocalVpc { get; set; } = null!;

        /// <summary>
        /// The UID of the member account (other Alibaba Cloud account) of the current Alibaba cloud account.
        /// </summary>
        [Input("memberUid")]
        public Input<string>? MemberUid { get; set; }

        /// <summary>
        /// Firewall switch status.
        /// </summary>
        [Input("status", required: true)]
        public Input<string> Status { get; set; } = null!;

        /// <summary>
        /// The name of the VPC firewall instance.
        /// </summary>
        [Input("vpcFirewallName", required: true)]
        public Input<string> VpcFirewallName { get; set; } = null!;

        /// <summary>
        /// The ID of the region to which the VPC is created.
        /// </summary>
        [Input("vpcRegion", required: true)]
        public Input<string> VpcRegion { get; set; } = null!;

        public FirewallVpcFirewallCenArgs()
        {
        }
        public static new FirewallVpcFirewallCenArgs Empty => new FirewallVpcFirewallCenArgs();
    }

    public sealed class FirewallVpcFirewallCenState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the CEN instance.
        /// </summary>
        [Input("cenId")]
        public Input<string>? CenId { get; set; }

        /// <summary>
        /// Intercommunication type, value: expressconnect: Express Channel cen: Cloud Enterprise Network
        /// </summary>
        [Input("connectType")]
        public Input<string>? ConnectType { get; set; }

        /// <summary>
        /// The language type of the requested and received messages. Valid values:
        /// </summary>
        [Input("lang")]
        public Input<string>? Lang { get; set; }

        /// <summary>
        /// The details of the VPC. See `local_vpc` below.
        /// </summary>
        [Input("localVpc")]
        public Input<Inputs.FirewallVpcFirewallCenLocalVpcGetArgs>? LocalVpc { get; set; }

        /// <summary>
        /// The UID of the member account (other Alibaba Cloud account) of the current Alibaba cloud account.
        /// </summary>
        [Input("memberUid")]
        public Input<string>? MemberUid { get; set; }

        /// <summary>
        /// Firewall switch status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// VPC firewall ID
        /// </summary>
        [Input("vpcFirewallId")]
        public Input<string>? VpcFirewallId { get; set; }

        /// <summary>
        /// The name of the VPC firewall instance.
        /// </summary>
        [Input("vpcFirewallName")]
        public Input<string>? VpcFirewallName { get; set; }

        /// <summary>
        /// The ID of the region to which the VPC is created.
        /// </summary>
        [Input("vpcRegion")]
        public Input<string>? VpcRegion { get; set; }

        public FirewallVpcFirewallCenState()
        {
        }
        public static new FirewallVpcFirewallCenState Empty => new FirewallVpcFirewallCenState();
    }
}

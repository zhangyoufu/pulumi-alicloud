// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vpc
{
    /// <summary>
    /// Provides a VPC Ipv4 Cidr Block resource. VPC IPv4 additional network segment.
    /// 
    /// For information about VPC Ipv4 Cidr Block and how to use it, see [What is Ipv4 Cidr Block](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/associatevpccidrblock).
    /// 
    /// &gt; **NOTE:** Available since v1.185.0.
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
    ///     var name = config.Get("name") ?? "terraform-example";
    ///     var defaultvpc = new AliCloud.Vpc.Network("defaultvpc", new()
    ///     {
    ///         Description = name,
    ///     });
    /// 
    ///     var @default = new AliCloud.Vpc.Ipv4CidrBlock("default", new()
    ///     {
    ///         SecondaryCidrBlock = "192.168.0.0/16",
    ///         VpcId = defaultvpc.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// VPC Ipv4 Cidr Block can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:vpc/ipv4CidrBlock:Ipv4CidrBlock example &lt;vpc_id&gt;:&lt;secondary_cidr_block&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:vpc/ipv4CidrBlock:Ipv4CidrBlock")]
    public partial class Ipv4CidrBlock : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The IPv4 CIDR block. Take note of the following requirements:
        /// * You can specify one of the following standard IPv4 CIDR blocks or their subnets as the secondary IPv4 CIDR block: 192.168.0.0/16, 172.16.0.0/12, and 10.0.0.0/8.
        /// * You can also use a custom CIDR block other than 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, 169.254.0.0/16, or their subnets as the secondary IPv4 CIDR block of the VPC.
        /// * The CIDR block cannot start with 0. The subnet mask must be 8 to 28 bits in length.
        /// * The secondary CIDR block cannot overlap with the primary CIDR block or an existing secondary CIDR block.
        /// </summary>
        [Output("secondaryCidrBlock")]
        public Output<string> SecondaryCidrBlock { get; private set; } = null!;

        /// <summary>
        /// The ID of the VPC.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a Ipv4CidrBlock resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ipv4CidrBlock(string name, Ipv4CidrBlockArgs args, CustomResourceOptions? options = null)
            : base("alicloud:vpc/ipv4CidrBlock:Ipv4CidrBlock", name, args ?? new Ipv4CidrBlockArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Ipv4CidrBlock(string name, Input<string> id, Ipv4CidrBlockState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:vpc/ipv4CidrBlock:Ipv4CidrBlock", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Ipv4CidrBlock resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ipv4CidrBlock Get(string name, Input<string> id, Ipv4CidrBlockState? state = null, CustomResourceOptions? options = null)
        {
            return new Ipv4CidrBlock(name, id, state, options);
        }
    }

    public sealed class Ipv4CidrBlockArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IPv4 CIDR block. Take note of the following requirements:
        /// * You can specify one of the following standard IPv4 CIDR blocks or their subnets as the secondary IPv4 CIDR block: 192.168.0.0/16, 172.16.0.0/12, and 10.0.0.0/8.
        /// * You can also use a custom CIDR block other than 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, 169.254.0.0/16, or their subnets as the secondary IPv4 CIDR block of the VPC.
        /// * The CIDR block cannot start with 0. The subnet mask must be 8 to 28 bits in length.
        /// * The secondary CIDR block cannot overlap with the primary CIDR block or an existing secondary CIDR block.
        /// </summary>
        [Input("secondaryCidrBlock", required: true)]
        public Input<string> SecondaryCidrBlock { get; set; } = null!;

        /// <summary>
        /// The ID of the VPC.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public Ipv4CidrBlockArgs()
        {
        }
        public static new Ipv4CidrBlockArgs Empty => new Ipv4CidrBlockArgs();
    }

    public sealed class Ipv4CidrBlockState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IPv4 CIDR block. Take note of the following requirements:
        /// * You can specify one of the following standard IPv4 CIDR blocks or their subnets as the secondary IPv4 CIDR block: 192.168.0.0/16, 172.16.0.0/12, and 10.0.0.0/8.
        /// * You can also use a custom CIDR block other than 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, 169.254.0.0/16, or their subnets as the secondary IPv4 CIDR block of the VPC.
        /// * The CIDR block cannot start with 0. The subnet mask must be 8 to 28 bits in length.
        /// * The secondary CIDR block cannot overlap with the primary CIDR block or an existing secondary CIDR block.
        /// </summary>
        [Input("secondaryCidrBlock")]
        public Input<string>? SecondaryCidrBlock { get; set; }

        /// <summary>
        /// The ID of the VPC.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public Ipv4CidrBlockState()
        {
        }
        public static new Ipv4CidrBlockState Empty => new Ipv4CidrBlockState();
    }
}

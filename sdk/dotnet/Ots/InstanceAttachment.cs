// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ots
{
    /// <summary>
    /// This resource will help you to bind a VPC to an OTS instance.
    /// 
    /// &gt; **NOTE:** Available since v1.10.0.
    /// 
    /// ## Example Usage
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
    ///     var defaultInteger = new Random.Index.Integer("default", new()
    ///     {
    ///         Min = 10000,
    ///         Max = 99999,
    ///     });
    /// 
    ///     var defaultInstance = new AliCloud.Ots.Instance("default", new()
    ///     {
    ///         Name = $"{name}-{defaultInteger.Result}",
    ///         Description = name,
    ///         AccessedBy = "Vpc",
    ///         Tags = 
    ///         {
    ///             { "Created", "TF" },
    ///             { "For", "example" },
    ///         },
    ///     });
    /// 
    ///     var @default = AliCloud.GetZones.Invoke(new()
    ///     {
    ///         AvailableResourceCreation = "VSwitch",
    ///     });
    /// 
    ///     var defaultNetwork = new AliCloud.Vpc.Network("default", new()
    ///     {
    ///         VpcName = name,
    ///         CidrBlock = "10.4.0.0/16",
    ///     });
    /// 
    ///     var defaultSwitch = new AliCloud.Vpc.Switch("default", new()
    ///     {
    ///         VswitchName = name,
    ///         CidrBlock = "10.4.0.0/24",
    ///         VpcId = defaultNetwork.Id,
    ///         ZoneId = @default.Apply(@default =&gt; @default.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id)),
    ///     });
    /// 
    ///     var defaultInstanceAttachment = new AliCloud.Ots.InstanceAttachment("default", new()
    ///     {
    ///         InstanceName = defaultInstance.Name,
    ///         VpcName = "examplename",
    ///         VswitchId = defaultSwitch.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [AliCloudResourceType("alicloud:ots/instanceAttachment:InstanceAttachment")]
    public partial class InstanceAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the OTS instance.
        /// </summary>
        [Output("instanceName")]
        public Output<string> InstanceName { get; private set; } = null!;

        /// <summary>
        /// The ID of attaching VPC to instance.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        /// <summary>
        /// The name of attaching VPC to instance. It can only contain letters and numbers, must start with a letter, and is limited to 3-16 characters in length.
        /// </summary>
        [Output("vpcName")]
        public Output<string> VpcName { get; private set; } = null!;

        /// <summary>
        /// The ID of attaching VSwitch to instance.
        /// </summary>
        [Output("vswitchId")]
        public Output<string> VswitchId { get; private set; } = null!;


        /// <summary>
        /// Create a InstanceAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InstanceAttachment(string name, InstanceAttachmentArgs args, CustomResourceOptions? options = null)
            : base("alicloud:ots/instanceAttachment:InstanceAttachment", name, args ?? new InstanceAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InstanceAttachment(string name, Input<string> id, InstanceAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ots/instanceAttachment:InstanceAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InstanceAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InstanceAttachment Get(string name, Input<string> id, InstanceAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new InstanceAttachment(name, id, state, options);
        }
    }

    public sealed class InstanceAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the OTS instance.
        /// </summary>
        [Input("instanceName", required: true)]
        public Input<string> InstanceName { get; set; } = null!;

        /// <summary>
        /// The name of attaching VPC to instance. It can only contain letters and numbers, must start with a letter, and is limited to 3-16 characters in length.
        /// </summary>
        [Input("vpcName", required: true)]
        public Input<string> VpcName { get; set; } = null!;

        /// <summary>
        /// The ID of attaching VSwitch to instance.
        /// </summary>
        [Input("vswitchId", required: true)]
        public Input<string> VswitchId { get; set; } = null!;

        public InstanceAttachmentArgs()
        {
        }
        public static new InstanceAttachmentArgs Empty => new InstanceAttachmentArgs();
    }

    public sealed class InstanceAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the OTS instance.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        /// <summary>
        /// The ID of attaching VPC to instance.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// The name of attaching VPC to instance. It can only contain letters and numbers, must start with a letter, and is limited to 3-16 characters in length.
        /// </summary>
        [Input("vpcName")]
        public Input<string>? VpcName { get; set; }

        /// <summary>
        /// The ID of attaching VSwitch to instance.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        public InstanceAttachmentState()
        {
        }
        public static new InstanceAttachmentState Empty => new InstanceAttachmentState();
    }
}

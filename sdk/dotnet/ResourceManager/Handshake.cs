// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ResourceManager
{
    /// <summary>
    /// Provides a Resource Manager handshake resource. You can invite accounts to join a resource directory for unified management.
    /// For information about Resource Manager handshake and how to use it, see [What is Resource Manager handshake](https://www.alibabacloud.com/help/en/doc-detail/135287.htm).
    /// 
    /// &gt; **NOTE:** Available in v1.82.0+.
    /// 
    /// ## Example Usage
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
    ///     // Add a Resource Manager handshake.
    ///     var example = new AliCloud.ResourceManager.Handshake("example", new()
    ///     {
    ///         TargetEntity = "1182775234******",
    ///         TargetType = "Account",
    ///         Note = "test resource manager handshake",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Resource Manager handshake can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:resourcemanager/handshake:Handshake example h-QmdexeFm1kE*****
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:resourcemanager/handshake:Handshake")]
    public partial class Handshake : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The expiration time of the invitation.
        /// </summary>
        [Output("expireTime")]
        public Output<string> ExpireTime { get; private set; } = null!;

        /// <summary>
        /// Resource account master account ID.
        /// </summary>
        [Output("masterAccountId")]
        public Output<string> MasterAccountId { get; private set; } = null!;

        /// <summary>
        /// The name of the main account of the resource directory.
        /// </summary>
        [Output("masterAccountName")]
        public Output<string> MasterAccountName { get; private set; } = null!;

        /// <summary>
        /// The modification time of the invitation.
        /// </summary>
        [Output("modifyTime")]
        public Output<string> ModifyTime { get; private set; } = null!;

        /// <summary>
        /// Remarks. The maximum length is 1024 characters.
        /// </summary>
        [Output("note")]
        public Output<string?> Note { get; private set; } = null!;

        /// <summary>
        /// Resource directory ID.
        /// </summary>
        [Output("resourceDirectoryId")]
        public Output<string> ResourceDirectoryId { get; private set; } = null!;

        /// <summary>
        /// Invitation status. Valid values: `Pending` waiting for confirmation, `Accepted`, `Cancelled`, `Declined`, `Expired`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Invited account ID or login email.
        /// </summary>
        [Output("targetEntity")]
        public Output<string> TargetEntity { get; private set; } = null!;

        /// <summary>
        /// Type of account being invited. Valid values: `Account`, `Email`.
        /// </summary>
        [Output("targetType")]
        public Output<string> TargetType { get; private set; } = null!;


        /// <summary>
        /// Create a Handshake resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Handshake(string name, HandshakeArgs args, CustomResourceOptions? options = null)
            : base("alicloud:resourcemanager/handshake:Handshake", name, args ?? new HandshakeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Handshake(string name, Input<string> id, HandshakeState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:resourcemanager/handshake:Handshake", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Handshake resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Handshake Get(string name, Input<string> id, HandshakeState? state = null, CustomResourceOptions? options = null)
        {
            return new Handshake(name, id, state, options);
        }
    }

    public sealed class HandshakeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Remarks. The maximum length is 1024 characters.
        /// </summary>
        [Input("note")]
        public Input<string>? Note { get; set; }

        /// <summary>
        /// Invited account ID or login email.
        /// </summary>
        [Input("targetEntity", required: true)]
        public Input<string> TargetEntity { get; set; } = null!;

        /// <summary>
        /// Type of account being invited. Valid values: `Account`, `Email`.
        /// </summary>
        [Input("targetType", required: true)]
        public Input<string> TargetType { get; set; } = null!;

        public HandshakeArgs()
        {
        }
        public static new HandshakeArgs Empty => new HandshakeArgs();
    }

    public sealed class HandshakeState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The expiration time of the invitation.
        /// </summary>
        [Input("expireTime")]
        public Input<string>? ExpireTime { get; set; }

        /// <summary>
        /// Resource account master account ID.
        /// </summary>
        [Input("masterAccountId")]
        public Input<string>? MasterAccountId { get; set; }

        /// <summary>
        /// The name of the main account of the resource directory.
        /// </summary>
        [Input("masterAccountName")]
        public Input<string>? MasterAccountName { get; set; }

        /// <summary>
        /// The modification time of the invitation.
        /// </summary>
        [Input("modifyTime")]
        public Input<string>? ModifyTime { get; set; }

        /// <summary>
        /// Remarks. The maximum length is 1024 characters.
        /// </summary>
        [Input("note")]
        public Input<string>? Note { get; set; }

        /// <summary>
        /// Resource directory ID.
        /// </summary>
        [Input("resourceDirectoryId")]
        public Input<string>? ResourceDirectoryId { get; set; }

        /// <summary>
        /// Invitation status. Valid values: `Pending` waiting for confirmation, `Accepted`, `Cancelled`, `Declined`, `Expired`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Invited account ID or login email.
        /// </summary>
        [Input("targetEntity")]
        public Input<string>? TargetEntity { get; set; }

        /// <summary>
        /// Type of account being invited. Valid values: `Account`, `Email`.
        /// </summary>
        [Input("targetType")]
        public Input<string>? TargetType { get; set; }

        public HandshakeState()
        {
        }
        public static new HandshakeState Empty => new HandshakeState();
    }
}

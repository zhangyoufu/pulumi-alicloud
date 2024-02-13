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
    /// Provides a Sag Acl resource. Smart Access Gateway (SAG) provides the access control list (ACL) function in the form of whitelists and blacklists for different SAG instances.
    /// 
    /// For information about Sag Acl and how to use it, see [What is access control list (ACL)](https://www.alibabacloud.com/help/en/smart-access-gateway/latest/createacl).
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
    ///     var @default = new AliCloud.RocketMQ.Acl("default");
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// The Sag Acl can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:rocketmq/acl:Acl example acl-abc123456
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:rocketmq/acl:Acl")]
    public partial class Acl : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the ACL instance. The name can contain 2 to 128 characters including a-z, A-Z, 0-9, periods, underlines, and hyphens. The name must start with an English letter, but cannot start with http:// or https://.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a Acl resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Acl(string name, AclArgs? args = null, CustomResourceOptions? options = null)
            : base("alicloud:rocketmq/acl:Acl", name, args ?? new AclArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Acl(string name, Input<string> id, AclState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:rocketmq/acl:Acl", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Acl resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Acl Get(string name, Input<string> id, AclState? state = null, CustomResourceOptions? options = null)
        {
            return new Acl(name, id, state, options);
        }
    }

    public sealed class AclArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the ACL instance. The name can contain 2 to 128 characters including a-z, A-Z, 0-9, periods, underlines, and hyphens. The name must start with an English letter, but cannot start with http:// or https://.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public AclArgs()
        {
        }
        public static new AclArgs Empty => new AclArgs();
    }

    public sealed class AclState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the ACL instance. The name can contain 2 to 128 characters including a-z, A-Z, 0-9, periods, underlines, and hyphens. The name must start with an English letter, but cannot start with http:// or https://.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public AclState()
        {
        }
        public static new AclState Empty => new AclState();
    }
}

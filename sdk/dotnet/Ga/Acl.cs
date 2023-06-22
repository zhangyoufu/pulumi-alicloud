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
    /// Provides a Global Accelerator (GA) Acl resource.
    /// 
    /// For information about Global Accelerator (GA) Acl and how to use it, see [What is Acl](https://www.alibabacloud.com/help/en/global-accelerator/latest/api-doc-ga-2019-11-20-api-doc-createacl).
    /// 
    /// &gt; **NOTE:** Available since v1.150.0.
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
    ///     var defaultAcl = new AliCloud.Ga.Acl("defaultAcl", new()
    ///     {
    ///         AclName = "terraform-example",
    ///         AddressIpVersion = "IPv4",
    ///     });
    /// 
    ///     var defaultAclEntryAttachment = new AliCloud.Ga.AclEntryAttachment("defaultAclEntryAttachment", new()
    ///     {
    ///         AclId = defaultAcl.Id,
    ///         Entry = "192.168.1.1/32",
    ///         EntryDescription = "terraform-example",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Global Accelerator (GA) Acl can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:ga/acl:Acl example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:ga/acl:Acl")]
    public partial class Acl : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The entries of the Acl. See `acl_entries` below. **NOTE:** "Field 'acl_entries' has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource `alicloud.ga.AclEntryAttachment`."
        /// </summary>
        [Output("aclEntries")]
        public Output<ImmutableArray<Outputs.AclAclEntry>> AclEntries { get; private set; } = null!;

        /// <summary>
        /// The name of the ACL. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), hyphens (-) and underscores (_). It must start with a letter.
        /// </summary>
        [Output("aclName")]
        public Output<string?> AclName { get; private set; } = null!;

        /// <summary>
        /// The IP version. Valid values: `IPv4` and `IPv6`.
        /// </summary>
        [Output("addressIpVersion")]
        public Output<string> AddressIpVersion { get; private set; } = null!;

        /// <summary>
        /// The dry run.
        /// </summary>
        [Output("dryRun")]
        public Output<bool?> DryRun { get; private set; } = null!;

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a Acl resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Acl(string name, AclArgs args, CustomResourceOptions? options = null)
            : base("alicloud:ga/acl:Acl", name, args ?? new AclArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Acl(string name, Input<string> id, AclState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ga/acl:Acl", name, state, MakeResourceOptions(options, id))
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
        [Input("aclEntries")]
        private InputList<Inputs.AclAclEntryArgs>? _aclEntries;

        /// <summary>
        /// The entries of the Acl. See `acl_entries` below. **NOTE:** "Field 'acl_entries' has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource `alicloud.ga.AclEntryAttachment`."
        /// </summary>
        [Obsolete(@"Field 'acl_entries' has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource 'alicloud_ga_acl_entry_attachment'.")]
        public InputList<Inputs.AclAclEntryArgs> AclEntries
        {
            get => _aclEntries ?? (_aclEntries = new InputList<Inputs.AclAclEntryArgs>());
            set => _aclEntries = value;
        }

        /// <summary>
        /// The name of the ACL. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), hyphens (-) and underscores (_). It must start with a letter.
        /// </summary>
        [Input("aclName")]
        public Input<string>? AclName { get; set; }

        /// <summary>
        /// The IP version. Valid values: `IPv4` and `IPv6`.
        /// </summary>
        [Input("addressIpVersion", required: true)]
        public Input<string> AddressIpVersion { get; set; } = null!;

        /// <summary>
        /// The dry run.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        public AclArgs()
        {
        }
        public static new AclArgs Empty => new AclArgs();
    }

    public sealed class AclState : global::Pulumi.ResourceArgs
    {
        [Input("aclEntries")]
        private InputList<Inputs.AclAclEntryGetArgs>? _aclEntries;

        /// <summary>
        /// The entries of the Acl. See `acl_entries` below. **NOTE:** "Field 'acl_entries' has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource `alicloud.ga.AclEntryAttachment`."
        /// </summary>
        [Obsolete(@"Field 'acl_entries' has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource 'alicloud_ga_acl_entry_attachment'.")]
        public InputList<Inputs.AclAclEntryGetArgs> AclEntries
        {
            get => _aclEntries ?? (_aclEntries = new InputList<Inputs.AclAclEntryGetArgs>());
            set => _aclEntries = value;
        }

        /// <summary>
        /// The name of the ACL. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), hyphens (-) and underscores (_). It must start with a letter.
        /// </summary>
        [Input("aclName")]
        public Input<string>? AclName { get; set; }

        /// <summary>
        /// The IP version. Valid values: `IPv4` and `IPv6`.
        /// </summary>
        [Input("addressIpVersion")]
        public Input<string>? AddressIpVersion { get; set; }

        /// <summary>
        /// The dry run.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public AclState()
        {
        }
        public static new AclState Empty => new AclState();
    }
}

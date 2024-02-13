// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Dcdn
{
    /// <summary>
    /// Provides a DCDN Waf Policy Domain Attachment resource.
    /// 
    /// For information about DCDN Waf Policy Domain Attachment and how to use it, see [What is Waf Policy Domain Attachment](https://www.alibabacloud.com/help/en/dcdn/developer-reference/api-dcdn-2018-01-15-modifydcdnwafpolicydomains).
    /// 
    /// &gt; **NOTE:** Available since v1.186.0.
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
    /// using Random = Pulumi.Random;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var domainName = config.Get("domainName") ?? "tf-example.com";
    ///     var name = config.Get("name") ?? "tf_example";
    ///     var @default = new Random.RandomInteger("default", new()
    ///     {
    ///         Min = 10000,
    ///         Max = 99999,
    ///     });
    /// 
    ///     var exampleDomain = new AliCloud.Dcdn.Domain("exampleDomain", new()
    ///     {
    ///         DomainName = @default.Result.Apply(result =&gt; $"{domainName}-{result}"),
    ///         Scope = "overseas",
    ///         Sources = new[]
    ///         {
    ///             new AliCloud.Dcdn.Inputs.DomainSourceArgs
    ///             {
    ///                 Content = "1.1.1.1",
    ///                 Port = 80,
    ///                 Priority = "20",
    ///                 Type = "ipaddr",
    ///                 Weight = "10",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var exampleWafDomain = new AliCloud.Dcdn.WafDomain("exampleWafDomain", new()
    ///     {
    ///         DomainName = exampleDomain.DomainName,
    ///         ClientIpTag = "X-Forwarded-For",
    ///     });
    /// 
    ///     var exampleWafPolicy = new AliCloud.Dcdn.WafPolicy("exampleWafPolicy", new()
    ///     {
    ///         DefenseScene = "waf_group",
    ///         PolicyName = @default.Result.Apply(result =&gt; $"{name}_{result}"),
    ///         PolicyType = "custom",
    ///         Status = "on",
    ///     });
    /// 
    ///     var exampleWafPolicyDomainAttachment = new AliCloud.Dcdn.WafPolicyDomainAttachment("exampleWafPolicyDomainAttachment", new()
    ///     {
    ///         DomainName = exampleWafDomain.DomainName,
    ///         PolicyId = exampleWafPolicy.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// DCDN Waf Policy Domain Attachment can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:dcdn/wafPolicyDomainAttachment:WafPolicyDomainAttachment example policy_id:domain_name
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:dcdn/wafPolicyDomainAttachment:WafPolicyDomainAttachment")]
    public partial class WafPolicyDomainAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Access the accelerated domain name of the specified protection policy.
        /// </summary>
        [Output("domainName")]
        public Output<string> DomainName { get; private set; } = null!;

        /// <summary>
        /// The protection policy ID. Only one input is supported.
        /// </summary>
        [Output("policyId")]
        public Output<string> PolicyId { get; private set; } = null!;


        /// <summary>
        /// Create a WafPolicyDomainAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WafPolicyDomainAttachment(string name, WafPolicyDomainAttachmentArgs args, CustomResourceOptions? options = null)
            : base("alicloud:dcdn/wafPolicyDomainAttachment:WafPolicyDomainAttachment", name, args ?? new WafPolicyDomainAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WafPolicyDomainAttachment(string name, Input<string> id, WafPolicyDomainAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:dcdn/wafPolicyDomainAttachment:WafPolicyDomainAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing WafPolicyDomainAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WafPolicyDomainAttachment Get(string name, Input<string> id, WafPolicyDomainAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new WafPolicyDomainAttachment(name, id, state, options);
        }
    }

    public sealed class WafPolicyDomainAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access the accelerated domain name of the specified protection policy.
        /// </summary>
        [Input("domainName", required: true)]
        public Input<string> DomainName { get; set; } = null!;

        /// <summary>
        /// The protection policy ID. Only one input is supported.
        /// </summary>
        [Input("policyId", required: true)]
        public Input<string> PolicyId { get; set; } = null!;

        public WafPolicyDomainAttachmentArgs()
        {
        }
        public static new WafPolicyDomainAttachmentArgs Empty => new WafPolicyDomainAttachmentArgs();
    }

    public sealed class WafPolicyDomainAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access the accelerated domain name of the specified protection policy.
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        /// <summary>
        /// The protection policy ID. Only one input is supported.
        /// </summary>
        [Input("policyId")]
        public Input<string>? PolicyId { get; set; }

        public WafPolicyDomainAttachmentState()
        {
        }
        public static new WafPolicyDomainAttachmentState Empty => new WafPolicyDomainAttachmentState();
    }
}

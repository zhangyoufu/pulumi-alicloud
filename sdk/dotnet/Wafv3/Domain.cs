// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Wafv3
{
    /// <summary>
    /// Provides a Wafv3 Domain resource.
    /// 
    /// For information about Wafv3 Domain and how to use it, see [What is Domain](https://www.alibabacloud.com/help/en/web-application-firewall/latest/api-waf-openapi-2021-10-01-createdomain).
    /// 
    /// &gt; **NOTE:** Available since v1.200.0.
    /// 
    /// ## Import
    /// 
    /// Wafv3 Domain can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:wafv3/domain:Domain example &lt;instance_id&gt;:&lt;domain&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:wafv3/domain:Domain")]
    public partial class Domain : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The access type of the WAF instance. Value: **share** (default): CNAME access.
        /// </summary>
        [Output("accessType")]
        public Output<string?> AccessType { get; private set; } = null!;

        /// <summary>
        /// The name of the domain name to query.
        /// </summary>
        [Output("domain")]
        public Output<string> DomainName { get; private set; } = null!;

        /// <summary>
        /// WAF instance ID
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Configure listening information. See `listen` below.
        /// </summary>
        [Output("listen")]
        public Output<Outputs.DomainListen> Listen { get; private set; } = null!;

        /// <summary>
        /// Configure forwarding information. See `redirect` below.
        /// </summary>
        [Output("redirect")]
        public Output<Outputs.DomainRedirect> Redirect { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Output("resourceManagerResourceGroupId")]
        public Output<string> ResourceManagerResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a Domain resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Domain(string name, DomainArgs args, CustomResourceOptions? options = null)
            : base("alicloud:wafv3/domain:Domain", name, args ?? new DomainArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Domain(string name, Input<string> id, DomainState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:wafv3/domain:Domain", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Domain resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Domain Get(string name, Input<string> id, DomainState? state = null, CustomResourceOptions? options = null)
        {
            return new Domain(name, id, state, options);
        }
    }

    public sealed class DomainArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access type of the WAF instance. Value: **share** (default): CNAME access.
        /// </summary>
        [Input("accessType")]
        public Input<string>? AccessType { get; set; }

        /// <summary>
        /// The name of the domain name to query.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> DomainName { get; set; } = null!;

        /// <summary>
        /// WAF instance ID
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Configure listening information. See `listen` below.
        /// </summary>
        [Input("listen", required: true)]
        public Input<Inputs.DomainListenArgs> Listen { get; set; } = null!;

        /// <summary>
        /// Configure forwarding information. See `redirect` below.
        /// </summary>
        [Input("redirect", required: true)]
        public Input<Inputs.DomainRedirectArgs> Redirect { get; set; } = null!;

        public DomainArgs()
        {
        }
        public static new DomainArgs Empty => new DomainArgs();
    }

    public sealed class DomainState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access type of the WAF instance. Value: **share** (default): CNAME access.
        /// </summary>
        [Input("accessType")]
        public Input<string>? AccessType { get; set; }

        /// <summary>
        /// The name of the domain name to query.
        /// </summary>
        [Input("domain")]
        public Input<string>? DomainName { get; set; }

        /// <summary>
        /// WAF instance ID
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Configure listening information. See `listen` below.
        /// </summary>
        [Input("listen")]
        public Input<Inputs.DomainListenGetArgs>? Listen { get; set; }

        /// <summary>
        /// Configure forwarding information. See `redirect` below.
        /// </summary>
        [Input("redirect")]
        public Input<Inputs.DomainRedirectGetArgs>? Redirect { get; set; }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceManagerResourceGroupId")]
        public Input<string>? ResourceManagerResourceGroupId { get; set; }

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public DomainState()
        {
        }
        public static new DomainState Empty => new DomainState();
    }
}

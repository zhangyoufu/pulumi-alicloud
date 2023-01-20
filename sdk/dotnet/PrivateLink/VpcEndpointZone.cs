// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.PrivateLink
{
    /// <summary>
    /// Provides a Private Link Vpc Endpoint Zone resource.
    /// 
    /// For information about Private Link Vpc Endpoint Zone and how to use it, see [What is Vpc Endpoint Zone](https://help.aliyun.com/document_detail/183561.html).
    /// 
    /// &gt; **NOTE:** Available in v1.111.0+.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new AliCloud.PrivateLink.VpcEndpointZone("example", new()
    ///     {
    ///         EndpointId = "ep-gw8boxxxxx",
    ///         VswitchId = "vsw-rtycxxxxx",
    ///         ZoneId = "eu-central-1a",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Private Link Vpc Endpoint Zone can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:privatelink/vpcEndpointZone:VpcEndpointZone example &lt;endpoint_id&gt;:&lt;zone_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:privatelink/vpcEndpointZone:VpcEndpointZone")]
    public partial class VpcEndpointZone : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The dry run.
        /// </summary>
        [Output("dryRun")]
        public Output<bool?> DryRun { get; private set; } = null!;

        /// <summary>
        /// The ID of the Vpc Endpoint.
        /// </summary>
        [Output("endpointId")]
        public Output<string> EndpointId { get; private set; } = null!;

        /// <summary>
        /// Status.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The VSwitch id.
        /// </summary>
        [Output("vswitchId")]
        public Output<string> VswitchId { get; private set; } = null!;

        /// <summary>
        /// The Zone Id.
        /// </summary>
        [Output("zoneId")]
        public Output<string> ZoneId { get; private set; } = null!;


        /// <summary>
        /// Create a VpcEndpointZone resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcEndpointZone(string name, VpcEndpointZoneArgs args, CustomResourceOptions? options = null)
            : base("alicloud:privatelink/vpcEndpointZone:VpcEndpointZone", name, args ?? new VpcEndpointZoneArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcEndpointZone(string name, Input<string> id, VpcEndpointZoneState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:privatelink/vpcEndpointZone:VpcEndpointZone", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VpcEndpointZone resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcEndpointZone Get(string name, Input<string> id, VpcEndpointZoneState? state = null, CustomResourceOptions? options = null)
        {
            return new VpcEndpointZone(name, id, state, options);
        }
    }

    public sealed class VpcEndpointZoneArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The dry run.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The ID of the Vpc Endpoint.
        /// </summary>
        [Input("endpointId", required: true)]
        public Input<string> EndpointId { get; set; } = null!;

        /// <summary>
        /// The VSwitch id.
        /// </summary>
        [Input("vswitchId", required: true)]
        public Input<string> VswitchId { get; set; } = null!;

        /// <summary>
        /// The Zone Id.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public VpcEndpointZoneArgs()
        {
        }
        public static new VpcEndpointZoneArgs Empty => new VpcEndpointZoneArgs();
    }

    public sealed class VpcEndpointZoneState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The dry run.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The ID of the Vpc Endpoint.
        /// </summary>
        [Input("endpointId")]
        public Input<string>? EndpointId { get; set; }

        /// <summary>
        /// Status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The VSwitch id.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        /// <summary>
        /// The Zone Id.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public VpcEndpointZoneState()
        {
        }
        public static new VpcEndpointZoneState Empty => new VpcEndpointZoneState();
    }
}

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
    /// Provides a Private Link Vpc Endpoint Service resource.
    /// 
    /// For information about Private Link Vpc Endpoint Service and how to use it, see [What is Vpc Endpoint Service](https://help.aliyun.com/document_detail/183540.html).
    /// 
    /// &gt; **NOTE:** Available in v1.109.0+.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new AliCloud.PrivateLink.VpcEndpointService("example", new AliCloud.PrivateLink.VpcEndpointServiceArgs
    ///         {
    ///             AutoAcceptConnection = false,
    ///             ConnectBandwidth = 103,
    ///             ServiceDescription = "tftest",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Private Link Vpc Endpoint Service can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:privatelink/vpcEndpointService:VpcEndpointService example &lt;service_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:privatelink/vpcEndpointService:VpcEndpointService")]
    public partial class VpcEndpointService : Pulumi.CustomResource
    {
        /// <summary>
        /// Whether to automatically accept terminal node connections.
        /// </summary>
        [Output("autoAcceptConnection")]
        public Output<bool?> AutoAcceptConnection { get; private set; } = null!;

        /// <summary>
        /// The connection bandwidth.
        /// </summary>
        [Output("connectBandwidth")]
        public Output<int> ConnectBandwidth { get; private set; } = null!;

        /// <summary>
        /// Whether to pre-check this request only. Default to: `false`
        /// </summary>
        [Output("dryRun")]
        public Output<bool?> DryRun { get; private set; } = null!;

        /// <summary>
        /// The payer type. Valid Value: `EndpointService`, `Endpoint`. Default to: `Endpoint`.
        /// </summary>
        [Output("payer")]
        public Output<string?> Payer { get; private set; } = null!;

        /// <summary>
        /// The business status of Vpc Endpoint Service.
        /// </summary>
        [Output("serviceBusinessStatus")]
        public Output<string> ServiceBusinessStatus { get; private set; } = null!;

        /// <summary>
        /// The description of the terminal node service.
        /// </summary>
        [Output("serviceDescription")]
        public Output<string?> ServiceDescription { get; private set; } = null!;

        /// <summary>
        /// Service Domain.
        /// </summary>
        [Output("serviceDomain")]
        public Output<string> ServiceDomain { get; private set; } = null!;

        /// <summary>
        /// The status of Vpc Endpoint Service.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a VpcEndpointService resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcEndpointService(string name, VpcEndpointServiceArgs? args = null, CustomResourceOptions? options = null)
            : base("alicloud:privatelink/vpcEndpointService:VpcEndpointService", name, args ?? new VpcEndpointServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcEndpointService(string name, Input<string> id, VpcEndpointServiceState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:privatelink/vpcEndpointService:VpcEndpointService", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VpcEndpointService resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcEndpointService Get(string name, Input<string> id, VpcEndpointServiceState? state = null, CustomResourceOptions? options = null)
        {
            return new VpcEndpointService(name, id, state, options);
        }
    }

    public sealed class VpcEndpointServiceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to automatically accept terminal node connections.
        /// </summary>
        [Input("autoAcceptConnection")]
        public Input<bool>? AutoAcceptConnection { get; set; }

        /// <summary>
        /// The connection bandwidth.
        /// </summary>
        [Input("connectBandwidth")]
        public Input<int>? ConnectBandwidth { get; set; }

        /// <summary>
        /// Whether to pre-check this request only. Default to: `false`
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The payer type. Valid Value: `EndpointService`, `Endpoint`. Default to: `Endpoint`.
        /// </summary>
        [Input("payer")]
        public Input<string>? Payer { get; set; }

        /// <summary>
        /// The description of the terminal node service.
        /// </summary>
        [Input("serviceDescription")]
        public Input<string>? ServiceDescription { get; set; }

        public VpcEndpointServiceArgs()
        {
        }
    }

    public sealed class VpcEndpointServiceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to automatically accept terminal node connections.
        /// </summary>
        [Input("autoAcceptConnection")]
        public Input<bool>? AutoAcceptConnection { get; set; }

        /// <summary>
        /// The connection bandwidth.
        /// </summary>
        [Input("connectBandwidth")]
        public Input<int>? ConnectBandwidth { get; set; }

        /// <summary>
        /// Whether to pre-check this request only. Default to: `false`
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The payer type. Valid Value: `EndpointService`, `Endpoint`. Default to: `Endpoint`.
        /// </summary>
        [Input("payer")]
        public Input<string>? Payer { get; set; }

        /// <summary>
        /// The business status of Vpc Endpoint Service.
        /// </summary>
        [Input("serviceBusinessStatus")]
        public Input<string>? ServiceBusinessStatus { get; set; }

        /// <summary>
        /// The description of the terminal node service.
        /// </summary>
        [Input("serviceDescription")]
        public Input<string>? ServiceDescription { get; set; }

        /// <summary>
        /// Service Domain.
        /// </summary>
        [Input("serviceDomain")]
        public Input<string>? ServiceDomain { get; set; }

        /// <summary>
        /// The status of Vpc Endpoint Service.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public VpcEndpointServiceState()
        {
        }
    }
}

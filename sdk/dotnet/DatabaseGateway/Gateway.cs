// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.DatabaseGateway
{
    /// <summary>
    /// Provides a Database Gateway Gateway resource.
    /// 
    /// For information about Database Gateway Gateway and how to use it, see [What is Gateway](https://www.alibabacloud.com/help/doc-detail/123415.htm).
    /// 
    /// &gt; **NOTE:** Available in v1.135.0+.
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
    ///     var example = new AliCloud.DatabaseGateway.Gateway("example", new()
    ///     {
    ///         GatewayName = "example_value",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Database Gateway Gateway can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:databasegateway/gateway:Gateway example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:databasegateway/gateway:Gateway")]
    public partial class Gateway : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of Gateway.
        /// </summary>
        [Output("gatewayDesc")]
        public Output<string?> GatewayDesc { get; private set; } = null!;

        /// <summary>
        /// The name of the Gateway.
        /// </summary>
        [Output("gatewayName")]
        public Output<string> GatewayName { get; private set; } = null!;

        /// <summary>
        /// The status of gateway. Valid values: `EXCEPTION`, `NEW`, `RUNNING`, `STOPPED`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a Gateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Gateway(string name, GatewayArgs args, CustomResourceOptions? options = null)
            : base("alicloud:databasegateway/gateway:Gateway", name, args ?? new GatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Gateway(string name, Input<string> id, GatewayState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:databasegateway/gateway:Gateway", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Gateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Gateway Get(string name, Input<string> id, GatewayState? state = null, CustomResourceOptions? options = null)
        {
            return new Gateway(name, id, state, options);
        }
    }

    public sealed class GatewayArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of Gateway.
        /// </summary>
        [Input("gatewayDesc")]
        public Input<string>? GatewayDesc { get; set; }

        /// <summary>
        /// The name of the Gateway.
        /// </summary>
        [Input("gatewayName", required: true)]
        public Input<string> GatewayName { get; set; } = null!;

        public GatewayArgs()
        {
        }
        public static new GatewayArgs Empty => new GatewayArgs();
    }

    public sealed class GatewayState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of Gateway.
        /// </summary>
        [Input("gatewayDesc")]
        public Input<string>? GatewayDesc { get; set; }

        /// <summary>
        /// The name of the Gateway.
        /// </summary>
        [Input("gatewayName")]
        public Input<string>? GatewayName { get; set; }

        /// <summary>
        /// The status of gateway. Valid values: `EXCEPTION`, `NEW`, `RUNNING`, `STOPPED`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GatewayState()
        {
        }
        public static new GatewayState Empty => new GatewayState();
    }
}

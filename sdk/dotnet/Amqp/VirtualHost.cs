// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Amqp
{
    /// <summary>
    /// Provides a RabbitMQ (AMQP) Virtual Host resource.
    /// 
    /// For information about RabbitMQ (AMQP) Virtual Host and how to use it, see [What is Virtual Host](https://www.alibabacloud.com/help/en/message-queue-for-rabbitmq/latest/createvirtualhost).
    /// 
    /// &gt; **NOTE:** Available since v1.126.0.
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
    ///     var defaultInstance = new AliCloud.Amqp.Instance("defaultInstance", new()
    ///     {
    ///         InstanceType = "professional",
    ///         MaxTps = "1000",
    ///         QueueCapacity = "50",
    ///         SupportEip = true,
    ///         MaxEipTps = "128",
    ///         PaymentType = "Subscription",
    ///         Period = 1,
    ///     });
    /// 
    ///     var defaultVirtualHost = new AliCloud.Amqp.VirtualHost("defaultVirtualHost", new()
    ///     {
    ///         InstanceId = defaultInstance.Id,
    ///         VirtualHostName = "tf-example",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// RabbitMQ (AMQP) Virtual Host can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:amqp/virtualHost:VirtualHost example &lt;instance_id&gt;:&lt;virtual_host_name&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:amqp/virtualHost:VirtualHost")]
    public partial class VirtualHost : global::Pulumi.CustomResource
    {
        /// <summary>
        /// InstanceId.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// VirtualHostName.
        /// </summary>
        [Output("virtualHostName")]
        public Output<string> VirtualHostName { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualHost resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualHost(string name, VirtualHostArgs args, CustomResourceOptions? options = null)
            : base("alicloud:amqp/virtualHost:VirtualHost", name, args ?? new VirtualHostArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualHost(string name, Input<string> id, VirtualHostState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:amqp/virtualHost:VirtualHost", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VirtualHost resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualHost Get(string name, Input<string> id, VirtualHostState? state = null, CustomResourceOptions? options = null)
        {
            return new VirtualHost(name, id, state, options);
        }
    }

    public sealed class VirtualHostArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// InstanceId.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// VirtualHostName.
        /// </summary>
        [Input("virtualHostName", required: true)]
        public Input<string> VirtualHostName { get; set; } = null!;

        public VirtualHostArgs()
        {
        }
        public static new VirtualHostArgs Empty => new VirtualHostArgs();
    }

    public sealed class VirtualHostState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// InstanceId.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// VirtualHostName.
        /// </summary>
        [Input("virtualHostName")]
        public Input<string>? VirtualHostName { get; set; }

        public VirtualHostState()
        {
        }
        public static new VirtualHostState Empty => new VirtualHostState();
    }
}

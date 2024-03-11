// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Amqp
{
    public static class GetBindings
    {
        /// <summary>
        /// This data source provides the Amqp Bindings of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.135.0+.
        /// 
        /// ## Example Usage
        /// 
        /// Basic Usage
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
        ///     var examples = AliCloud.Amqp.GetBindings.Invoke(new()
        ///     {
        ///         InstanceId = "amqp-cn-xxxxx",
        ///         VirtualHostName = "my-vh",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetBindingsResult> InvokeAsync(GetBindingsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBindingsResult>("alicloud:amqp/getBindings:getBindings", args ?? new GetBindingsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Amqp Bindings of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.135.0+.
        /// 
        /// ## Example Usage
        /// 
        /// Basic Usage
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
        ///     var examples = AliCloud.Amqp.GetBindings.Invoke(new()
        ///     {
        ///         InstanceId = "amqp-cn-xxxxx",
        ///         VirtualHostName = "my-vh",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetBindingsResult> Invoke(GetBindingsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBindingsResult>("alicloud:amqp/getBindings:getBindings", args ?? new GetBindingsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBindingsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Instance Id.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Virtualhost Name.
        /// </summary>
        [Input("virtualHostName", required: true)]
        public string VirtualHostName { get; set; } = null!;

        public GetBindingsArgs()
        {
        }
        public static new GetBindingsArgs Empty => new GetBindingsArgs();
    }

    public sealed class GetBindingsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Instance Id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// Virtualhost Name.
        /// </summary>
        [Input("virtualHostName", required: true)]
        public Input<string> VirtualHostName { get; set; } = null!;

        public GetBindingsInvokeArgs()
        {
        }
        public static new GetBindingsInvokeArgs Empty => new GetBindingsInvokeArgs();
    }


    [OutputType]
    public sealed class GetBindingsResult
    {
        public readonly ImmutableArray<Outputs.GetBindingsBindingResult> Bindings;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string InstanceId;
        public readonly string? OutputFile;
        public readonly string VirtualHostName;

        [OutputConstructor]
        private GetBindingsResult(
            ImmutableArray<Outputs.GetBindingsBindingResult> bindings,

            string id,

            ImmutableArray<string> ids,

            string instanceId,

            string? outputFile,

            string virtualHostName)
        {
            Bindings = bindings;
            Id = id;
            Ids = ids;
            InstanceId = instanceId;
            OutputFile = outputFile;
            VirtualHostName = virtualHostName;
        }
    }
}

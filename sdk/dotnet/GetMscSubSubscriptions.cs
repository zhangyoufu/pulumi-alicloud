// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud
{
    public static class GetMscSubSubscriptions
    {
        /// <summary>
        /// This data source provides the Message Center Subscriptions of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.135.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
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
        ///         var @default = Output.Create(AliCloud.GetMscSubSubscriptions.InvokeAsync());
        ///         this.MscSubSubscriptionId1 = @default.Apply(@default =&gt; @default.Subscriptions[0].Id);
        ///     }
        /// 
        ///     [Output("mscSubSubscriptionId1")]
        ///     public Output&lt;string&gt; MscSubSubscriptionId1 { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetMscSubSubscriptionsResult> InvokeAsync(GetMscSubSubscriptionsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMscSubSubscriptionsResult>("alicloud:index/getMscSubSubscriptions:getMscSubSubscriptions", args ?? new GetMscSubSubscriptionsArgs(), options.WithVersion());
    }


    public sealed class GetMscSubSubscriptionsArgs : Pulumi.InvokeArgs
    {
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetMscSubSubscriptionsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetMscSubSubscriptionsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? OutputFile;
        public readonly ImmutableArray<Outputs.GetMscSubSubscriptionsSubscriptionResult> Subscriptions;

        [OutputConstructor]
        private GetMscSubSubscriptionsResult(
            string id,

            string? outputFile,

            ImmutableArray<Outputs.GetMscSubSubscriptionsSubscriptionResult> subscriptions)
        {
            Id = id;
            OutputFile = outputFile;
            Subscriptions = subscriptions;
        }
    }
}

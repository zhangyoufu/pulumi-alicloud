// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Nlb
{
    public static class GetListeners
    {
        /// <summary>
        /// This data source provides the Nlb Listeners of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.191.0+.
        /// </summary>
        public static Task<GetListenersResult> InvokeAsync(GetListenersArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetListenersResult>("alicloud:nlb/getListeners:getListeners", args ?? new GetListenersArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Nlb Listeners of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.191.0+.
        /// </summary>
        public static Output<GetListenersResult> Invoke(GetListenersInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetListenersResult>("alicloud:nlb/getListeners:getListeners", args ?? new GetListenersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetListenersArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Listener IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The listening protocol. Valid values: `TCP`, `UDP`, or `TCPSSL`.
        /// </summary>
        [Input("listenerProtocol")]
        public string? ListenerProtocol { get; set; }

        [Input("loadBalancerIds")]
        private List<string>? _loadBalancerIds;

        /// <summary>
        /// The ID of the NLB instance. You can specify at most 20 IDs.
        /// </summary>
        public List<string> LoadBalancerIds
        {
            get => _loadBalancerIds ?? (_loadBalancerIds = new List<string>());
            set => _loadBalancerIds = value;
        }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetListenersArgs()
        {
        }
        public static new GetListenersArgs Empty => new GetListenersArgs();
    }

    public sealed class GetListenersInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Listener IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The listening protocol. Valid values: `TCP`, `UDP`, or `TCPSSL`.
        /// </summary>
        [Input("listenerProtocol")]
        public Input<string>? ListenerProtocol { get; set; }

        [Input("loadBalancerIds")]
        private InputList<string>? _loadBalancerIds;

        /// <summary>
        /// The ID of the NLB instance. You can specify at most 20 IDs.
        /// </summary>
        public InputList<string> LoadBalancerIds
        {
            get => _loadBalancerIds ?? (_loadBalancerIds = new InputList<string>());
            set => _loadBalancerIds = value;
        }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetListenersInvokeArgs()
        {
        }
        public static new GetListenersInvokeArgs Empty => new GetListenersInvokeArgs();
    }


    [OutputType]
    public sealed class GetListenersResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? ListenerProtocol;
        public readonly ImmutableArray<Outputs.GetListenersListenerResult> Listeners;
        public readonly ImmutableArray<string> LoadBalancerIds;
        public readonly string? OutputFile;

        [OutputConstructor]
        private GetListenersResult(
            string id,

            ImmutableArray<string> ids,

            string? listenerProtocol,

            ImmutableArray<Outputs.GetListenersListenerResult> listeners,

            ImmutableArray<string> loadBalancerIds,

            string? outputFile)
        {
            Id = id;
            Ids = ids;
            ListenerProtocol = listenerProtocol;
            Listeners = listeners;
            LoadBalancerIds = loadBalancerIds;
            OutputFile = outputFile;
        }
    }
}

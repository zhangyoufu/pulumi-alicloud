// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Sae.Inputs
{

    public sealed class ApplicationPvtzDiscoverySvcGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enables the Kubernetes Service-based registration and discovery feature.
        /// </summary>
        [Input("enable")]
        public Input<bool>? Enable { get; set; }

        /// <summary>
        /// The ID of the namespace.
        /// </summary>
        [Input("namespaceId")]
        public Input<string>? NamespaceId { get; set; }

        [Input("portProtocols")]
        private InputList<Inputs.ApplicationPvtzDiscoverySvcPortProtocolGetArgs>? _portProtocols;

        /// <summary>
        /// The port number and protocol. See `port_protocols` below.
        /// </summary>
        public InputList<Inputs.ApplicationPvtzDiscoverySvcPortProtocolGetArgs> PortProtocols
        {
            get => _portProtocols ?? (_portProtocols = new InputList<Inputs.ApplicationPvtzDiscoverySvcPortProtocolGetArgs>());
            set => _portProtocols = value;
        }

        /// <summary>
        /// The name of the Service.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        public ApplicationPvtzDiscoverySvcGetArgs()
        {
        }
        public static new ApplicationPvtzDiscoverySvcGetArgs Empty => new ApplicationPvtzDiscoverySvcGetArgs();
    }
}

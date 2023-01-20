// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CS.Inputs
{

    public sealed class EdgeKubernetesCertificateAuthorityArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The path of client certificate, like `~/.kube/client-cert.pem`.
        /// </summary>
        [Input("clientCert")]
        public Input<string>? ClientCert { get; set; }

        /// <summary>
        /// The path of client key, like `~/.kube/client-key.pem`.
        /// </summary>
        [Input("clientKey")]
        public Input<string>? ClientKey { get; set; }

        /// <summary>
        /// The base64 encoded cluster certificate data required to communicate with your cluster. Add this to the certificate-authority-data section of the kubeconfig file for your cluster.
        /// </summary>
        [Input("clusterCert")]
        public Input<string>? ClusterCert { get; set; }

        public EdgeKubernetesCertificateAuthorityArgs()
        {
        }
        public static new EdgeKubernetesCertificateAuthorityArgs Empty => new EdgeKubernetesCertificateAuthorityArgs();
    }
}

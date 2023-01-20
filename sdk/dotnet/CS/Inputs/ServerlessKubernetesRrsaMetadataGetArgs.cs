// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CS.Inputs
{

    public sealed class ServerlessKubernetesRrsaMetadataGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the RRSA feature has been enabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The arn of OIDC provider that was registered in RAM.
        /// </summary>
        [Input("ramOidcProviderArn")]
        public Input<string>? RamOidcProviderArn { get; set; }

        /// <summary>
        /// The name of OIDC Provider that was registered in RAM.
        /// </summary>
        [Input("ramOidcProviderName")]
        public Input<string>? RamOidcProviderName { get; set; }

        /// <summary>
        /// The issuer URL of RRSA OIDC Token.
        /// </summary>
        [Input("rrsaOidcIssuerUrl")]
        public Input<string>? RrsaOidcIssuerUrl { get; set; }

        public ServerlessKubernetesRrsaMetadataGetArgs()
        {
        }
        public static new ServerlessKubernetesRrsaMetadataGetArgs Empty => new ServerlessKubernetesRrsaMetadataGetArgs();
    }
}

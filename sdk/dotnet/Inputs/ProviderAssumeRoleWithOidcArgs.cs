// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Inputs
{

    public sealed class ProviderAssumeRoleWithOidcArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARN of the OIDC IdP.
        /// </summary>
        [Input("oidcProviderArn", required: true)]
        public Input<string> OidcProviderArn { get; set; } = null!;

        [Input("oidcToken")]
        public Input<string>? OidcToken { get; set; }

        /// <summary>
        /// The file path of OIDC token that is issued by the external IdP.
        /// </summary>
        [Input("oidcTokenFile")]
        public Input<string>? OidcTokenFile { get; set; }

        /// <summary>
        /// The policy that specifies the permissions of the returned STS token. You can use this parameter to grant the STS token fewer permissions than the permissions granted to the RAM role.
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        /// <summary>
        /// ARN of a RAM role to assume prior to making API calls.
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        /// <summary>
        /// The custom name of the role session. Set this parameter based on your business requirements. In most cases, this parameter is set to the identity of the user who calls the operation, for example, the username.
        /// </summary>
        [Input("roleSessionName")]
        public Input<string>? RoleSessionName { get; set; }

        /// <summary>
        /// The validity period of the STS token. Unit: seconds. Default value: 3600. Minimum value: 900. Maximum value: the value of the MaxSessionDuration parameter when creating a ram role.
        /// </summary>
        [Input("sessionExpiration")]
        public Input<int>? SessionExpiration { get; set; }

        public ProviderAssumeRoleWithOidcArgs()
        {
        }
        public static new ProviderAssumeRoleWithOidcArgs Empty => new ProviderAssumeRoleWithOidcArgs();
    }
}

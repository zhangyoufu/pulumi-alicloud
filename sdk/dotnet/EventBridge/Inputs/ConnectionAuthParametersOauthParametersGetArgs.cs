// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.EventBridge.Inputs
{

    public sealed class ConnectionAuthParametersOauthParametersGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IP address of the authorized endpoint.
        /// </summary>
        [Input("authorizationEndpoint")]
        public Input<string>? AuthorizationEndpoint { get; set; }

        /// <summary>
        /// The parameters that are configured for the client. See `client_parameters` below.
        /// </summary>
        [Input("clientParameters")]
        public Input<Inputs.ConnectionAuthParametersOauthParametersClientParametersGetArgs>? ClientParameters { get; set; }

        /// <summary>
        /// The HTTP request method. Valid values: `GET`, `POST`, `HEAD`, `DELETE`, `PUT`, `PATCH`.
        /// </summary>
        [Input("httpMethod")]
        public Input<string>? HttpMethod { get; set; }

        /// <summary>
        /// The request parameters that are configured for OAuth authentication. See `oauth_http_parameters` below.
        /// </summary>
        [Input("oauthHttpParameters")]
        public Input<Inputs.ConnectionAuthParametersOauthParametersOauthHttpParametersGetArgs>? OauthHttpParameters { get; set; }

        public ConnectionAuthParametersOauthParametersGetArgs()
        {
        }
        public static new ConnectionAuthParametersOauthParametersGetArgs Empty => new ConnectionAuthParametersOauthParametersGetArgs();
    }
}

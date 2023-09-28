// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.EventBridge.Outputs
{

    [OutputType]
    public sealed class ConnectionAuthParametersOauthParametersOauthHttpParametersQueryStringParameter
    {
        /// <summary>
        /// Specifies whether to enable authentication.
        /// </summary>
        public readonly string? IsValueSecret;
        /// <summary>
        /// The key of the request header.
        /// </summary>
        public readonly string? Key;
        /// <summary>
        /// The value of the request header.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private ConnectionAuthParametersOauthParametersOauthHttpParametersQueryStringParameter(
            string? isValueSecret,

            string? key,

            string? value)
        {
            IsValueSecret = isValueSecret;
            Key = key;
            Value = value;
        }
    }
}

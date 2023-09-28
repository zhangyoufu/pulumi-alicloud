// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Alb.Outputs
{

    [OutputType]
    public sealed class RuleRuleActionRedirectConfig
    {
        /// <summary>
        /// The host name of the destination to which requests are directed. The host name must meet the following rules:  The host name must be 3 to 128 characters in length, and can contain letters, digits, hyphens (-), periods (.), asterisks (*), and question marks (?). The host name must contain at least one period (.), and cannot start or end with a period (.). The rightmost domain label can contain only letters, asterisks (*) and question marks (?) and cannot contain digits or hyphens (-). Other domain labels cannot start or end with a hyphen (-). You can include asterisks (*) and question marks (?) anywhere in a domain label. Default value: ${host}. You cannot use this value with other characters at the same time.
        /// </summary>
        public readonly string? Host;
        /// <summary>
        /// The HTTP status code of the response. The code must be an `HTTP_2xx`, `HTTP_4xx` or `HTTP_5xx.x` is a digit.
        /// </summary>
        public readonly string? HttpCode;
        /// <summary>
        /// The path of the destination to which requests are directed. Valid values: The path must be 1 to 128 characters in length, and start with a forward slash (/). The path can contain letters, digits, asterisks (*), question marks (?) and the following special characters: $ - _ . + / &amp; ~ @ :. It cannot contain the following special characters: " % # ; ! ( ) [ ] ^ , ”. The path is case-sensitive. Default value: ${path}. You can also reference ${host}, ${protocol}, and ${port}. Each variable can appear at most once. You can use the preceding variables at the same time, or use them with a valid string.
        /// </summary>
        public readonly string? Path;
        /// <summary>
        /// The port of the destination to which requests are redirected. Valid values: 1 to 63335. Default value: ${port}. You cannot use this value together with other characters at the same time.
        /// </summary>
        public readonly string? Port;
        /// <summary>
        /// The protocol of the requests to be redirected. Valid values: `HTTP` and `HTTPS`. Default value: `${protocol}`. You cannot use this value together with other characters at the same time. Note HTTPS listeners can redirect only HTTPS requests.
        /// </summary>
        public readonly string? Protocol;
        /// <summary>
        /// The query string of the request to be redirected. The query string must be 1 to 128 characters in length, can contain letters and printable characters. It cannot contain the following special characters: # [ ] { } \ | &lt; &gt; &amp;. Default value: ${query}. You can also reference ${host}, ${protocol}, and ${port}. Each variable can appear at most once. You can use the preceding variables at the same time, or use them together with a valid string.
        /// </summary>
        public readonly string? Query;

        [OutputConstructor]
        private RuleRuleActionRedirectConfig(
            string? host,

            string? httpCode,

            string? path,

            string? port,

            string? protocol,

            string? query)
        {
            Host = host;
            HttpCode = httpCode;
            Path = path;
            Port = port;
            Protocol = protocol;
            Query = query;
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Alb.Inputs
{

    public sealed class ServerGroupStickySessionConfigGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// the cookie that is configured on the server. **NOTE:** This parameter exists if the `StickySession`
        /// parameter is set to `On` and the `StickySessionType` parameter is set to `server`.
        /// </summary>
        [Input("cookie")]
        public Input<string>? Cookie { get; set; }

        /// <summary>
        /// The timeout period of a cookie. The timeout period of a cookie. Unit: seconds. Valid values: `1`
        /// to `86400`. Default value: `1000`.
        /// </summary>
        [Input("cookieTimeout")]
        public Input<int>? CookieTimeout { get; set; }

        /// <summary>
        /// Indicates whether sticky session is enabled. Values: `true` and `false`. Default
        /// value: `false`.  **NOTE:** This parameter exists if the `StickySession` parameter is set to `On`.
        /// </summary>
        [Input("stickySessionEnabled")]
        public Input<bool>? StickySessionEnabled { get; set; }

        /// <summary>
        /// The method that is used to handle a cookie. Values: `Server` and `Insert`.
        /// </summary>
        [Input("stickySessionType")]
        public Input<string>? StickySessionType { get; set; }

        public ServerGroupStickySessionConfigGetArgs()
        {
        }
        public static new ServerGroupStickySessionConfigGetArgs Empty => new ServerGroupStickySessionConfigGetArgs();
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.FC.Inputs
{

    public sealed class FunctionAsyncInvokeConfigDestinationConfigOnSuccessArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Alicloud Resource Name (ARN) of the destination resource. See the [Developer Guide](https://www.alibabacloud.com/help/doc-detail/181866.htm) for acceptable resource types and associated RAM permissions.
        /// </summary>
        [Input("destination", required: true)]
        public Input<string> Destination { get; set; } = null!;

        public FunctionAsyncInvokeConfigDestinationConfigOnSuccessArgs()
        {
        }
    }
}

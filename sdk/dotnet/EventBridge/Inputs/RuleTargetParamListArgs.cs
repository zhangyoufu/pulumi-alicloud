// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.EventBridge.Inputs
{

    public sealed class RuleTargetParamListArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The format of the event target parameter. Valid values: `ORIGINAL`, `TEMPLATE`, `JSONPATH`, `CONSTANT`.
        /// </summary>
        [Input("form", required: true)]
        public Input<string> Form { get; set; } = null!;

        /// <summary>
        /// The resource parameter of the event target. For more information, see [How to use it](https://www.alibabacloud.com/help/en/eventbridge/latest/event-target-parameters)
        /// </summary>
        [Input("resourceKey", required: true)]
        public Input<string> ResourceKey { get; set; } = null!;

        /// <summary>
        /// The template of the event target parameter.
        /// </summary>
        [Input("template")]
        public Input<string>? Template { get; set; }

        /// <summary>
        /// The value of the event target parameter.
        /// 
        /// &gt; **NOTE:** There exists a potential diff error that the backend service will return a default param as following:
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// 
        /// In order to fix the diff, from version 1.160.0, this resource has removed the param which `resource_key = "IsBase64Encode"` and `value = "false"`.
        /// If you want to set `resource_key = "IsBase64Encode"`, please avoid to set `value = "false"`.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public RuleTargetParamListArgs()
        {
        }
        public static new RuleTargetParamListArgs Empty => new RuleTargetParamListArgs();
    }
}

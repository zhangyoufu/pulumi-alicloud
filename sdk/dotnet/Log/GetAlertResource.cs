// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Log
{
    public static class GetAlertResource
    {
        /// <summary>
        /// Using this data source can init SLS Alert resources automatically.
        /// 
        /// For information about SLS Alert and how to use it, see [SLS Alert Overview](https://www.alibabacloud.com/help/en/doc-detail/209202.html)
        /// 
        /// &gt; **NOTE:** Available in v1.161.0+
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var exampleUser = AliCloud.Log.GetAlertResource.Invoke(new()
        ///     {
        ///         Lang = "cn",
        ///         Type = "user",
        ///     });
        /// 
        ///     var exampleProject = AliCloud.Log.GetAlertResource.Invoke(new()
        ///     {
        ///         Project = "test-alert-tf",
        ///         Type = "project",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetAlertResourceResult> InvokeAsync(GetAlertResourceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAlertResourceResult>("alicloud:log/getAlertResource:getAlertResource", args ?? new GetAlertResourceArgs(), options.WithDefaults());

        /// <summary>
        /// Using this data source can init SLS Alert resources automatically.
        /// 
        /// For information about SLS Alert and how to use it, see [SLS Alert Overview](https://www.alibabacloud.com/help/en/doc-detail/209202.html)
        /// 
        /// &gt; **NOTE:** Available in v1.161.0+
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var exampleUser = AliCloud.Log.GetAlertResource.Invoke(new()
        ///     {
        ///         Lang = "cn",
        ///         Type = "user",
        ///     });
        /// 
        ///     var exampleProject = AliCloud.Log.GetAlertResource.Invoke(new()
        ///     {
        ///         Project = "test-alert-tf",
        ///         Type = "project",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetAlertResourceResult> Invoke(GetAlertResourceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAlertResourceResult>("alicloud:log/getAlertResource:getAlertResource", args ?? new GetAlertResourceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAlertResourceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The lang of alert center resource when type is user.
        /// </summary>
        [Input("lang")]
        public string? Lang { get; set; }

        /// <summary>
        /// The project of alert resource when type is project.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        /// <summary>
        /// The type of alert resources, must be user or project, 'user' for init aliyuncloud account's alert center resource, including project named sls-alert-{uid}-{region} and some dashboards; 'project' for init project's alert resource, including logstore named internal-alert-history and alert dashboard.
        /// </summary>
        [Input("type", required: true)]
        public string Type { get; set; } = null!;

        public GetAlertResourceArgs()
        {
        }
        public static new GetAlertResourceArgs Empty => new GetAlertResourceArgs();
    }

    public sealed class GetAlertResourceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The lang of alert center resource when type is user.
        /// </summary>
        [Input("lang")]
        public Input<string>? Lang { get; set; }

        /// <summary>
        /// The project of alert resource when type is project.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The type of alert resources, must be user or project, 'user' for init aliyuncloud account's alert center resource, including project named sls-alert-{uid}-{region} and some dashboards; 'project' for init project's alert resource, including logstore named internal-alert-history and alert dashboard.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public GetAlertResourceInvokeArgs()
        {
        }
        public static new GetAlertResourceInvokeArgs Empty => new GetAlertResourceInvokeArgs();
    }


    [OutputType]
    public sealed class GetAlertResourceResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Lang;
        public readonly string? Project;
        public readonly string Type;

        [OutputConstructor]
        private GetAlertResourceResult(
            string id,

            string? lang,

            string? project,

            string type)
        {
            Id = id;
            Lang = lang;
            Project = project;
            Type = type;
        }
    }
}

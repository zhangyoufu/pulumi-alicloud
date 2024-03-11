// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ApiGateway
{
    public static class GetApis
    {
        /// <summary>
        /// This data source provides the apis of the current Alibaba Cloud user.
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
        ///     var dataApigatwayApis = AliCloud.ApiGateway.GetApis.Invoke(new()
        ///     {
        ///         OutputFile = "output_ApiGatawayApis",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstApiId"] = data.Alicloud_api_gateway_apis.Data_apigatway.Apis[0].Id,
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetApisResult> InvokeAsync(GetApisArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetApisResult>("alicloud:apigateway/getApis:getApis", args ?? new GetApisArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the apis of the current Alibaba Cloud user.
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
        ///     var dataApigatwayApis = AliCloud.ApiGateway.GetApis.Invoke(new()
        ///     {
        ///         OutputFile = "output_ApiGatawayApis",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstApiId"] = data.Alicloud_api_gateway_apis.Data_apigatway.Apis[0].Id,
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetApisResult> Invoke(GetApisInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetApisResult>("alicloud:apigateway/getApis:getApis", args ?? new GetApisInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetApisArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// (It has been deprecated from version 1.52.2, and use field 'ids' to replace.) ID of the specified API.
        /// </summary>
        [Input("apiId")]
        public string? ApiId { get; set; }

        /// <summary>
        /// ID of the specified group.
        /// </summary>
        [Input("groupId")]
        public string? GroupId { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of api IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter api gateway apis by name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetApisArgs()
        {
        }
        public static new GetApisArgs Empty => new GetApisArgs();
    }

    public sealed class GetApisInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// (It has been deprecated from version 1.52.2, and use field 'ids' to replace.) ID of the specified API.
        /// </summary>
        [Input("apiId")]
        public Input<string>? ApiId { get; set; }

        /// <summary>
        /// ID of the specified group.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of api IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter api gateway apis by name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetApisInvokeArgs()
        {
        }
        public static new GetApisInvokeArgs Empty => new GetApisInvokeArgs();
    }


    [OutputType]
    public sealed class GetApisResult
    {
        public readonly string? ApiId;
        /// <summary>
        /// A list of apis. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetApisApiResult> Apis;
        /// <summary>
        /// The group id that the apis belong to.
        /// </summary>
        public readonly string? GroupId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of api IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of api names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;

        [OutputConstructor]
        private GetApisResult(
            string? apiId,

            ImmutableArray<Outputs.GetApisApiResult> apis,

            string? groupId,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile)
        {
            ApiId = apiId;
            Apis = apis;
            GroupId = groupId;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
        }
    }
}

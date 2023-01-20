// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ThreatDetection
{
    public static class GetWebLockConfigs
    {
        /// <summary>
        /// This data source provides Threat Detection Web Lock Config available to the user.[What is Web Lock Config](https://www.alibabacloud.com/help/en/security-center/latest/api-doc-sas-2018-12-03-api-doc-modifyweblockstart)
        /// 
        /// &gt; **NOTE:** Available in 1.195.0+
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @default = AliCloud.ThreatDetection.GetWebLockConfigs.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             alicloud_threat_detection_web_lock_config.Default.Id,
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["alicloudThreatDetectionWebLockConfigExampleId"] = @default.Apply(getWebLockConfigsResult =&gt; getWebLockConfigsResult).Apply(@default =&gt; @default.Apply(getWebLockConfigsResult =&gt; getWebLockConfigsResult.Configs[0]?.Id)),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetWebLockConfigsResult> InvokeAsync(GetWebLockConfigsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWebLockConfigsResult>("alicloud:threatdetection/getWebLockConfigs:getWebLockConfigs", args ?? new GetWebLockConfigsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides Threat Detection Web Lock Config available to the user.[What is Web Lock Config](https://www.alibabacloud.com/help/en/security-center/latest/api-doc-sas-2018-12-03-api-doc-modifyweblockstart)
        /// 
        /// &gt; **NOTE:** Available in 1.195.0+
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @default = AliCloud.ThreatDetection.GetWebLockConfigs.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             alicloud_threat_detection_web_lock_config.Default.Id,
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["alicloudThreatDetectionWebLockConfigExampleId"] = @default.Apply(getWebLockConfigsResult =&gt; getWebLockConfigsResult).Apply(@default =&gt; @default.Apply(getWebLockConfigsResult =&gt; getWebLockConfigsResult.Configs[0]?.Id)),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetWebLockConfigsResult> Invoke(GetWebLockConfigsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWebLockConfigsResult>("alicloud:threatdetection/getWebLockConfigs:getWebLockConfigs", args ?? new GetWebLockConfigsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWebLockConfigsArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Web Lock Config IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The language of the content within the request and the response. Valid values: `zh`, `en`.
        /// </summary>
        [Input("lang")]
        public string? Lang { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        [Input("pageNumber")]
        public int? PageNumber { get; set; }

        [Input("pageSize")]
        public int? PageSize { get; set; }

        /// <summary>
        /// The string that allows you to search for servers in fuzzy match mode. You can enter a server name or IP address.
        /// </summary>
        [Input("remark")]
        public string? Remark { get; set; }

        /// <summary>
        /// The source IP address of the request.
        /// </summary>
        [Input("sourceIp")]
        public string? SourceIp { get; set; }

        /// <summary>
        /// The protection status of the server that you want to query. Valid values: `on`, `off`.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetWebLockConfigsArgs()
        {
        }
        public static new GetWebLockConfigsArgs Empty => new GetWebLockConfigsArgs();
    }

    public sealed class GetWebLockConfigsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Web Lock Config IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The language of the content within the request and the response. Valid values: `zh`, `en`.
        /// </summary>
        [Input("lang")]
        public Input<string>? Lang { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        [Input("pageNumber")]
        public Input<int>? PageNumber { get; set; }

        [Input("pageSize")]
        public Input<int>? PageSize { get; set; }

        /// <summary>
        /// The string that allows you to search for servers in fuzzy match mode. You can enter a server name or IP address.
        /// </summary>
        [Input("remark")]
        public Input<string>? Remark { get; set; }

        /// <summary>
        /// The source IP address of the request.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// The protection status of the server that you want to query. Valid values: `on`, `off`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetWebLockConfigsInvokeArgs()
        {
        }
        public static new GetWebLockConfigsInvokeArgs Empty => new GetWebLockConfigsInvokeArgs();
    }


    [OutputType]
    public sealed class GetWebLockConfigsResult
    {
        /// <summary>
        /// A list of Web Lock Config Entries. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetWebLockConfigsConfigResult> Configs;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of Web Lock Config IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? Lang;
        public readonly string? OutputFile;
        public readonly int? PageNumber;
        public readonly int? PageSize;
        public readonly string? Remark;
        public readonly string? SourceIp;
        public readonly string? Status;

        [OutputConstructor]
        private GetWebLockConfigsResult(
            ImmutableArray<Outputs.GetWebLockConfigsConfigResult> configs,

            string id,

            ImmutableArray<string> ids,

            string? lang,

            string? outputFile,

            int? pageNumber,

            int? pageSize,

            string? remark,

            string? sourceIp,

            string? status)
        {
            Configs = configs;
            Id = id;
            Ids = ids;
            Lang = lang;
            OutputFile = outputFile;
            PageNumber = pageNumber;
            PageSize = pageSize;
            Remark = remark;
            SourceIp = sourceIp;
            Status = status;
        }
    }
}

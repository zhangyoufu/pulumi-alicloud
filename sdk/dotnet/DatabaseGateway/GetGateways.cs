// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.DatabaseGateway
{
    public static class GetGateways
    {
        /// <summary>
        /// This data source provides the Database Gateway Gateways of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.135.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic Usage
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var ids = Output.Create(AliCloud.DatabaseGateway.GetGateways.InvokeAsync(new AliCloud.DatabaseGateway.GetGatewaysArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 "example_id",
        ///             },
        ///         }));
        ///         this.DatabaseGatewayGatewayId1 = ids.Apply(ids =&gt; ids.Gateways[0].Id);
        ///         var nameRegex = Output.Create(AliCloud.DatabaseGateway.GetGateways.InvokeAsync(new AliCloud.DatabaseGateway.GetGatewaysArgs
        ///         {
        ///             NameRegex = "^my-Gateway",
        ///         }));
        ///         this.DatabaseGatewayGatewayId2 = nameRegex.Apply(nameRegex =&gt; nameRegex.Gateways[0].Id);
        ///     }
        /// 
        ///     [Output("databaseGatewayGatewayId1")]
        ///     public Output&lt;string&gt; DatabaseGatewayGatewayId1 { get; set; }
        ///     [Output("databaseGatewayGatewayId2")]
        ///     public Output&lt;string&gt; DatabaseGatewayGatewayId2 { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetGatewaysResult> InvokeAsync(GetGatewaysArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGatewaysResult>("alicloud:databasegateway/getGateways:getGateways", args ?? new GetGatewaysArgs(), options.WithVersion());
    }


    public sealed class GetGatewaysArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Default to `false`. Set it to `true` can output more details about resource attributes.
        /// </summary>
        [Input("enableDetails")]
        public bool? EnableDetails { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Gateway IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Gateway name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The search key.
        /// </summary>
        [Input("searchKey")]
        public string? SearchKey { get; set; }

        /// <summary>
        /// The status of gateway. Valid values: `EXCEPTION`, `NEW`, `RUNNING`, `STOPPED`.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetGatewaysArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetGatewaysResult
    {
        public readonly bool? EnableDetails;
        public readonly ImmutableArray<Outputs.GetGatewaysGatewayResult> Gateways;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly string? SearchKey;
        public readonly string? Status;

        [OutputConstructor]
        private GetGatewaysResult(
            bool? enableDetails,

            ImmutableArray<Outputs.GetGatewaysGatewayResult> gateways,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? searchKey,

            string? status)
        {
            EnableDetails = enableDetails;
            Gateways = gateways;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            SearchKey = searchKey;
            Status = status;
        }
    }
}

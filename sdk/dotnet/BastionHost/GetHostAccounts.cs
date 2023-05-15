// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.BastionHost
{
    public static class GetHostAccounts
    {
        /// <summary>
        /// This data source provides the Bastionhost Host Accounts of the current Alibaba Cloud user.
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
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var ids = AliCloud.BastionHost.GetHostAccounts.Invoke(new()
        ///     {
        ///         HostId = "15",
        ///         InstanceId = "example_value",
        ///         Ids = new[]
        ///         {
        ///             "1",
        ///             "2",
        ///         },
        ///     });
        /// 
        ///     var nameRegex = AliCloud.BastionHost.GetHostAccounts.Invoke(new()
        ///     {
        ///         HostId = "15",
        ///         InstanceId = "example_value",
        ///         NameRegex = "^my-HostAccount",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["bastionhostHostAccountId1"] = ids.Apply(getHostAccountsResult =&gt; getHostAccountsResult.Accounts[0]?.Id),
        ///         ["bastionhostHostAccountId2"] = nameRegex.Apply(getHostAccountsResult =&gt; getHostAccountsResult.Accounts[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetHostAccountsResult> InvokeAsync(GetHostAccountsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetHostAccountsResult>("alicloud:bastionhost/getHostAccounts:getHostAccounts", args ?? new GetHostAccountsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Bastionhost Host Accounts of the current Alibaba Cloud user.
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
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var ids = AliCloud.BastionHost.GetHostAccounts.Invoke(new()
        ///     {
        ///         HostId = "15",
        ///         InstanceId = "example_value",
        ///         Ids = new[]
        ///         {
        ///             "1",
        ///             "2",
        ///         },
        ///     });
        /// 
        ///     var nameRegex = AliCloud.BastionHost.GetHostAccounts.Invoke(new()
        ///     {
        ///         HostId = "15",
        ///         InstanceId = "example_value",
        ///         NameRegex = "^my-HostAccount",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["bastionhostHostAccountId1"] = ids.Apply(getHostAccountsResult =&gt; getHostAccountsResult.Accounts[0]?.Id),
        ///         ["bastionhostHostAccountId2"] = nameRegex.Apply(getHostAccountsResult =&gt; getHostAccountsResult.Accounts[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetHostAccountsResult> Invoke(GetHostAccountsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetHostAccountsResult>("alicloud:bastionhost/getHostAccounts:getHostAccounts", args ?? new GetHostAccountsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetHostAccountsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the new hosting account's name, support the longest 128 characters.
        /// </summary>
        [Input("hostAccountName")]
        public string? HostAccountName { get; set; }

        /// <summary>
        /// Specifies the database where you want to create your hosting account's host ID.
        /// </summary>
        [Input("hostId", required: true)]
        public string HostId { get; set; } = null!;

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Host Account IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// Specifies the database where you want to create your hosting account's host bastion host ID of.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// A regex string to filter results by Host Account name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Specify the new hosting account of the agreement name. Valid values: USING SSH and RDP.
        /// </summary>
        [Input("protocolName")]
        public string? ProtocolName { get; set; }

        public GetHostAccountsArgs()
        {
        }
        public static new GetHostAccountsArgs Empty => new GetHostAccountsArgs();
    }

    public sealed class GetHostAccountsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the new hosting account's name, support the longest 128 characters.
        /// </summary>
        [Input("hostAccountName")]
        public Input<string>? HostAccountName { get; set; }

        /// <summary>
        /// Specifies the database where you want to create your hosting account's host ID.
        /// </summary>
        [Input("hostId", required: true)]
        public Input<string> HostId { get; set; } = null!;

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Host Account IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// Specifies the database where you want to create your hosting account's host bastion host ID of.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// A regex string to filter results by Host Account name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// Specify the new hosting account of the agreement name. Valid values: USING SSH and RDP.
        /// </summary>
        [Input("protocolName")]
        public Input<string>? ProtocolName { get; set; }

        public GetHostAccountsInvokeArgs()
        {
        }
        public static new GetHostAccountsInvokeArgs Empty => new GetHostAccountsInvokeArgs();
    }


    [OutputType]
    public sealed class GetHostAccountsResult
    {
        public readonly ImmutableArray<Outputs.GetHostAccountsAccountResult> Accounts;
        public readonly string? HostAccountName;
        public readonly string HostId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string InstanceId;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly string? ProtocolName;

        [OutputConstructor]
        private GetHostAccountsResult(
            ImmutableArray<Outputs.GetHostAccountsAccountResult> accounts,

            string? hostAccountName,

            string hostId,

            string id,

            ImmutableArray<string> ids,

            string instanceId,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? protocolName)
        {
            Accounts = accounts;
            HostAccountName = hostAccountName;
            HostId = hostId;
            Id = id;
            Ids = ids;
            InstanceId = instanceId;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            ProtocolName = protocolName;
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Gpdb
{
    public static class GetAccounts
    {
        /// <summary>
        /// This data source provides the Gpdb Accounts of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.142.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var ids = AliCloud.Gpdb.GetAccounts.Invoke(new()
        ///     {
        ///         DbInstanceId = "example_value",
        ///         Ids = new[]
        ///         {
        ///             "my-Account-1",
        ///             "my-Account-2",
        ///         },
        ///     });
        /// 
        ///     var nameRegex = AliCloud.Gpdb.GetAccounts.Invoke(new()
        ///     {
        ///         DbInstanceId = "example_value",
        ///         NameRegex = "^my-Account",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["gpdbAccountId1"] = ids.Apply(getAccountsResult =&gt; getAccountsResult.Accounts[0]?.Id),
        ///         ["gpdbAccountId2"] = nameRegex.Apply(getAccountsResult =&gt; getAccountsResult.Accounts[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAccountsResult> InvokeAsync(GetAccountsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAccountsResult>("alicloud:gpdb/getAccounts:getAccounts", args ?? new GetAccountsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Gpdb Accounts of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.142.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var ids = AliCloud.Gpdb.GetAccounts.Invoke(new()
        ///     {
        ///         DbInstanceId = "example_value",
        ///         Ids = new[]
        ///         {
        ///             "my-Account-1",
        ///             "my-Account-2",
        ///         },
        ///     });
        /// 
        ///     var nameRegex = AliCloud.Gpdb.GetAccounts.Invoke(new()
        ///     {
        ///         DbInstanceId = "example_value",
        ///         NameRegex = "^my-Account",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["gpdbAccountId1"] = ids.Apply(getAccountsResult =&gt; getAccountsResult.Accounts[0]?.Id),
        ///         ["gpdbAccountId2"] = nameRegex.Apply(getAccountsResult =&gt; getAccountsResult.Accounts[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAccountsResult> Invoke(GetAccountsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccountsResult>("alicloud:gpdb/getAccounts:getAccounts", args ?? new GetAccountsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAccountsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the instance.
        /// </summary>
        [Input("dbInstanceId", required: true)]
        public string DbInstanceId { get; set; } = null!;

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Account IDs. Its element value is same as Account Name.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Account name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The status of the account. Valid values: `Active`, `Creating` and `Deleting`.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetAccountsArgs()
        {
        }
        public static new GetAccountsArgs Empty => new GetAccountsArgs();
    }

    public sealed class GetAccountsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the instance.
        /// </summary>
        [Input("dbInstanceId", required: true)]
        public Input<string> DbInstanceId { get; set; } = null!;

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Account IDs. Its element value is same as Account Name.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Account name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The status of the account. Valid values: `Active`, `Creating` and `Deleting`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetAccountsInvokeArgs()
        {
        }
        public static new GetAccountsInvokeArgs Empty => new GetAccountsInvokeArgs();
    }


    [OutputType]
    public sealed class GetAccountsResult
    {
        public readonly ImmutableArray<Outputs.GetAccountsAccountResult> Accounts;
        public readonly string DbInstanceId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly string? Status;

        [OutputConstructor]
        private GetAccountsResult(
            ImmutableArray<Outputs.GetAccountsAccountResult> accounts,

            string dbInstanceId,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? status)
        {
            Accounts = accounts;
            DbInstanceId = dbInstanceId;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Status = status;
        }
    }
}

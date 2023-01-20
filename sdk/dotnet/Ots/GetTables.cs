// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ots
{
    public static class GetTables
    {
        /// <summary>
        /// This data source provides the ots tables of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.40.0+.
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
        ///     var tablesDs = AliCloud.Ots.GetTables.Invoke(new()
        ///     {
        ///         InstanceName = "sample-instance",
        ///         NameRegex = "sample-table",
        ///         OutputFile = "tables.txt",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstTableId"] = tablesDs.Apply(getTablesResult =&gt; getTablesResult.Tables[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetTablesResult> InvokeAsync(GetTablesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTablesResult>("alicloud:ots/getTables:getTables", args ?? new GetTablesArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the ots tables of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.40.0+.
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
        ///     var tablesDs = AliCloud.Ots.GetTables.Invoke(new()
        ///     {
        ///         InstanceName = "sample-instance",
        ///         NameRegex = "sample-table",
        ///         OutputFile = "tables.txt",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstTableId"] = tablesDs.Apply(getTablesResult =&gt; getTablesResult.Tables[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetTablesResult> Invoke(GetTablesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTablesResult>("alicloud:ots/getTables:getTables", args ?? new GetTablesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTablesArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of table IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The name of OTS instance.
        /// </summary>
        [Input("instanceName", required: true)]
        public string InstanceName { get; set; } = null!;

        /// <summary>
        /// A regex string to filter results by table name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetTablesArgs()
        {
        }
        public static new GetTablesArgs Empty => new GetTablesArgs();
    }

    public sealed class GetTablesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of table IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The name of OTS instance.
        /// </summary>
        [Input("instanceName", required: true)]
        public Input<string> InstanceName { get; set; } = null!;

        /// <summary>
        /// A regex string to filter results by table name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetTablesInvokeArgs()
        {
        }
        public static new GetTablesInvokeArgs Empty => new GetTablesInvokeArgs();
    }


    [OutputType]
    public sealed class GetTablesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of table IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// The OTS instance name.
        /// </summary>
        public readonly string InstanceName;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of table names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        /// <summary>
        /// A list of tables. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetTablesTableResult> Tables;

        [OutputConstructor]
        private GetTablesResult(
            string id,

            ImmutableArray<string> ids,

            string instanceName,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            ImmutableArray<Outputs.GetTablesTableResult> tables)
        {
            Id = id;
            Ids = ids;
            InstanceName = instanceName;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Tables = tables;
        }
    }
}

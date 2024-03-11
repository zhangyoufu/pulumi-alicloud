// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Dms
{
    public static class GetEnterpriseDatabases
    {
        /// <summary>
        /// This data source provides DMS Enterprise Database available to the user. [What is Database](https://www.alibabacloud.com/help/en/dms/developer-reference/api-dms-enterprise-2018-11-01-listdatabases).
        /// 
        /// &gt; **NOTE:** Available since v1.195.0.
        /// 
        /// ## Example Usage
        /// 
        /// Basic Usage
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
        ///     var @default = AliCloud.Dms.GetEnterpriseDatabases.Invoke(new()
        ///     {
        ///         NameRegex = "test2",
        ///         InstanceId = "2195118",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["alicloudDmsEnterpriseDatabaseExampleId"] = @default.Apply(@default =&gt; @default.Apply(getEnterpriseDatabasesResult =&gt; getEnterpriseDatabasesResult.Databases[0]?.Id)),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetEnterpriseDatabasesResult> InvokeAsync(GetEnterpriseDatabasesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetEnterpriseDatabasesResult>("alicloud:dms/getEnterpriseDatabases:getEnterpriseDatabases", args ?? new GetEnterpriseDatabasesArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides DMS Enterprise Database available to the user. [What is Database](https://www.alibabacloud.com/help/en/dms/developer-reference/api-dms-enterprise-2018-11-01-listdatabases).
        /// 
        /// &gt; **NOTE:** Available since v1.195.0.
        /// 
        /// ## Example Usage
        /// 
        /// Basic Usage
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
        ///     var @default = AliCloud.Dms.GetEnterpriseDatabases.Invoke(new()
        ///     {
        ///         NameRegex = "test2",
        ///         InstanceId = "2195118",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["alicloudDmsEnterpriseDatabaseExampleId"] = @default.Apply(@default =&gt; @default.Apply(getEnterpriseDatabasesResult =&gt; getEnterpriseDatabasesResult.Databases[0]?.Id)),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetEnterpriseDatabasesResult> Invoke(GetEnterpriseDatabasesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetEnterpriseDatabasesResult>("alicloud:dms/getEnterpriseDatabases:getEnterpriseDatabases", args ?? new GetEnterpriseDatabasesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEnterpriseDatabasesArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Database IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The instance ID of the target database.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// A regex string to filter the results by the database Schema Name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetEnterpriseDatabasesArgs()
        {
        }
        public static new GetEnterpriseDatabasesArgs Empty => new GetEnterpriseDatabasesArgs();
    }

    public sealed class GetEnterpriseDatabasesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Database IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The instance ID of the target database.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// A regex string to filter the results by the database Schema Name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetEnterpriseDatabasesInvokeArgs()
        {
        }
        public static new GetEnterpriseDatabasesInvokeArgs Empty => new GetEnterpriseDatabasesInvokeArgs();
    }


    [OutputType]
    public sealed class GetEnterpriseDatabasesResult
    {
        /// <summary>
        /// A list of Database Entries. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetEnterpriseDatabasesDatabaseResult> Databases;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of Database IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// The instance ID of the target database.
        /// </summary>
        public readonly string InstanceId;
        public readonly string? NameRegex;
        public readonly string? OutputFile;

        [OutputConstructor]
        private GetEnterpriseDatabasesResult(
            ImmutableArray<Outputs.GetEnterpriseDatabasesDatabaseResult> databases,

            string id,

            ImmutableArray<string> ids,

            string instanceId,

            string? nameRegex,

            string? outputFile)
        {
            Databases = databases;
            Id = id;
            Ids = ids;
            InstanceId = instanceId;
            NameRegex = nameRegex;
            OutputFile = outputFile;
        }
    }
}

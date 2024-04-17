// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Adb
{
    /// <summary>
    /// Provides a ADB Lake Account resource. Account of the DBClusterLakeVesion.
    /// 
    /// For information about ADB Lake Account and how to use it, see [What is Lake Account](https://www.alibabacloud.com/help/en/analyticdb-for-mysql/developer-reference/api-adb-2021-12-01-createaccount).
    /// For information about ADB Lake Account Privileges and how to use it, see [What are Lake Account Privileges](https://www.alibabacloud.com/help/en/analyticdb-for-mysql/developer-reference/api-adb-2021-12-01-modifyaccountprivileges/).
    /// 
    /// &gt; **NOTE:** Available since v1.214.0.
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
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "terraform-example";
    ///     var @default = AliCloud.GetZones.Invoke(new()
    ///     {
    ///         AvailableResourceCreation = "VSwitch",
    ///     });
    /// 
    ///     var VPCID = new AliCloud.Vpc.Network("VPCID", new()
    ///     {
    ///         VpcName = name,
    ///         CidrBlock = "172.16.0.0/12",
    ///     });
    /// 
    ///     var VSWITCHID = new AliCloud.Vpc.Switch("VSWITCHID", new()
    ///     {
    ///         VpcId = VPCID.Id,
    ///         ZoneId = "cn-hangzhou-k",
    ///         VswitchName = name,
    ///         CidrBlock = "172.16.0.0/24",
    ///     });
    /// 
    ///     var createInstance = new AliCloud.Adb.DBClusterLakeVersion("CreateInstance", new()
    ///     {
    ///         StorageResource = "0ACU",
    ///         ZoneId = "cn-hangzhou-k",
    ///         VpcId = VPCID.Id,
    ///         VswitchId = VSWITCHID.Id,
    ///         DbClusterDescription = name,
    ///         ComputeResource = "16ACU",
    ///         DbClusterVersion = "5.0",
    ///         PaymentType = "PayAsYouGo",
    ///         SecurityIps = "127.0.0.1",
    ///     });
    /// 
    ///     var defaultLakeAccount = new AliCloud.Adb.LakeAccount("default", new()
    ///     {
    ///         DbClusterId = createInstance.Id,
    ///         AccountType = "Super",
    ///         AccountName = "tfnormal",
    ///         AccountPassword = "normal@2023",
    ///         AccountPrivileges = new[]
    ///         {
    ///             new AliCloud.Adb.Inputs.LakeAccountAccountPrivilegeArgs
    ///             {
    ///                 PrivilegeType = "Database",
    ///                 PrivilegeObject = new AliCloud.Adb.Inputs.LakeAccountAccountPrivilegePrivilegeObjectArgs
    ///                 {
    ///                     Database = "MYSQL",
    ///                 },
    ///                 Privileges = new[]
    ///                 {
    ///                     "select",
    ///                     "update",
    ///                 },
    ///             },
    ///             new AliCloud.Adb.Inputs.LakeAccountAccountPrivilegeArgs
    ///             {
    ///                 PrivilegeType = "Table",
    ///                 PrivilegeObject = new AliCloud.Adb.Inputs.LakeAccountAccountPrivilegePrivilegeObjectArgs
    ///                 {
    ///                     Database = "INFORMATION_SCHEMA",
    ///                     Table = "ENGINES",
    ///                 },
    ///                 Privileges = new[]
    ///                 {
    ///                     "update",
    ///                 },
    ///             },
    ///             new AliCloud.Adb.Inputs.LakeAccountAccountPrivilegeArgs
    ///             {
    ///                 PrivilegeType = "Column",
    ///                 PrivilegeObject = new AliCloud.Adb.Inputs.LakeAccountAccountPrivilegePrivilegeObjectArgs
    ///                 {
    ///                     Table = "COLUMNS",
    ///                     Column = "PRIVILEGES",
    ///                     Database = "INFORMATION_SCHEMA",
    ///                 },
    ///                 Privileges = new[]
    ///                 {
    ///                     "update",
    ///                 },
    ///             },
    ///         },
    ///         AccountDescription = name,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// ADB Lake Account can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:adb/lakeAccount:LakeAccount example &lt;db_cluster_id&gt;:&lt;account_name&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:adb/lakeAccount:LakeAccount")]
    public partial class LakeAccount : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the account.
        /// </summary>
        [Output("accountDescription")]
        public Output<string?> AccountDescription { get; private set; } = null!;

        /// <summary>
        /// The name of the account.
        /// </summary>
        [Output("accountName")]
        public Output<string> AccountName { get; private set; } = null!;

        /// <summary>
        /// AccountPassword.
        /// </summary>
        [Output("accountPassword")]
        public Output<string> AccountPassword { get; private set; } = null!;

        /// <summary>
        /// List of permissions granted. See `account_privileges` below.
        /// </summary>
        [Output("accountPrivileges")]
        public Output<ImmutableArray<Outputs.LakeAccountAccountPrivilege>> AccountPrivileges { get; private set; } = null!;

        /// <summary>
        /// The type of the account.
        /// </summary>
        [Output("accountType")]
        public Output<string?> AccountType { get; private set; } = null!;

        /// <summary>
        /// The DBCluster ID.
        /// </summary>
        [Output("dbClusterId")]
        public Output<string> DbClusterId { get; private set; } = null!;

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a LakeAccount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LakeAccount(string name, LakeAccountArgs args, CustomResourceOptions? options = null)
            : base("alicloud:adb/lakeAccount:LakeAccount", name, args ?? new LakeAccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LakeAccount(string name, Input<string> id, LakeAccountState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:adb/lakeAccount:LakeAccount", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "accountPassword",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing LakeAccount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LakeAccount Get(string name, Input<string> id, LakeAccountState? state = null, CustomResourceOptions? options = null)
        {
            return new LakeAccount(name, id, state, options);
        }
    }

    public sealed class LakeAccountArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the account.
        /// </summary>
        [Input("accountDescription")]
        public Input<string>? AccountDescription { get; set; }

        /// <summary>
        /// The name of the account.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        [Input("accountPassword", required: true)]
        private Input<string>? _accountPassword;

        /// <summary>
        /// AccountPassword.
        /// </summary>
        public Input<string>? AccountPassword
        {
            get => _accountPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _accountPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("accountPrivileges")]
        private InputList<Inputs.LakeAccountAccountPrivilegeArgs>? _accountPrivileges;

        /// <summary>
        /// List of permissions granted. See `account_privileges` below.
        /// </summary>
        public InputList<Inputs.LakeAccountAccountPrivilegeArgs> AccountPrivileges
        {
            get => _accountPrivileges ?? (_accountPrivileges = new InputList<Inputs.LakeAccountAccountPrivilegeArgs>());
            set => _accountPrivileges = value;
        }

        /// <summary>
        /// The type of the account.
        /// </summary>
        [Input("accountType")]
        public Input<string>? AccountType { get; set; }

        /// <summary>
        /// The DBCluster ID.
        /// </summary>
        [Input("dbClusterId", required: true)]
        public Input<string> DbClusterId { get; set; } = null!;

        public LakeAccountArgs()
        {
        }
        public static new LakeAccountArgs Empty => new LakeAccountArgs();
    }

    public sealed class LakeAccountState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the account.
        /// </summary>
        [Input("accountDescription")]
        public Input<string>? AccountDescription { get; set; }

        /// <summary>
        /// The name of the account.
        /// </summary>
        [Input("accountName")]
        public Input<string>? AccountName { get; set; }

        [Input("accountPassword")]
        private Input<string>? _accountPassword;

        /// <summary>
        /// AccountPassword.
        /// </summary>
        public Input<string>? AccountPassword
        {
            get => _accountPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _accountPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("accountPrivileges")]
        private InputList<Inputs.LakeAccountAccountPrivilegeGetArgs>? _accountPrivileges;

        /// <summary>
        /// List of permissions granted. See `account_privileges` below.
        /// </summary>
        public InputList<Inputs.LakeAccountAccountPrivilegeGetArgs> AccountPrivileges
        {
            get => _accountPrivileges ?? (_accountPrivileges = new InputList<Inputs.LakeAccountAccountPrivilegeGetArgs>());
            set => _accountPrivileges = value;
        }

        /// <summary>
        /// The type of the account.
        /// </summary>
        [Input("accountType")]
        public Input<string>? AccountType { get; set; }

        /// <summary>
        /// The DBCluster ID.
        /// </summary>
        [Input("dbClusterId")]
        public Input<string>? DbClusterId { get; set; }

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public LakeAccountState()
        {
        }
        public static new LakeAccountState Empty => new LakeAccountState();
    }
}

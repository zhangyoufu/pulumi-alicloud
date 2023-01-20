// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Dms
{
    /// <summary>
    /// Provides a DMS Enterprise Logic Database resource.
    /// 
    /// For information about DMS Enterprise Logic Database and how to use it, see [What is Logic Database](https://www.alibabacloud.com/help/zh/data-management-service/latest/api-doc-dms-enterprise-2018-11-01-api-doc-createlogicdatabase).
    /// 
    /// &gt; **NOTE:** Available in v1.195.0+.
    /// 
    /// ## Example Usage
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
    ///     var @default = new AliCloud.Dms.EnterpriseLogicDatabase("default", new()
    ///     {
    ///         Alias = "TF_logic_db_test",
    ///         DatabaseIds = new[]
    ///         {
    ///             "35617919",
    ///             "35617920",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// DMS Enterprise Logic Database can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:dms/enterpriseLogicDatabase:EnterpriseLogicDatabase example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:dms/enterpriseLogicDatabase:EnterpriseLogicDatabase")]
    public partial class EnterpriseLogicDatabase : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Logical Library alias.
        /// </summary>
        [Output("alias")]
        public Output<string> Alias { get; private set; } = null!;

        /// <summary>
        /// Sub-Database ID
        /// </summary>
        [Output("databaseIds")]
        public Output<ImmutableArray<string>> DatabaseIds { get; private set; } = null!;

        /// <summary>
        /// Database type.
        /// </summary>
        [Output("dbType")]
        public Output<string> DbType { get; private set; } = null!;

        /// <summary>
        /// Environment type, return value is as follows:-product: production environment-dev: development environment-pre: Advance Environment-test: test environment-sit:SIT environment-uat:UAT environment-pet: Pressure measurement environment-stag:STAG environment
        /// </summary>
        [Output("envType")]
        public Output<string> EnvType { get; private set; } = null!;

        /// <summary>
        /// Whether it is a logical Library, the return value is true.
        /// </summary>
        [Output("logic")]
        public Output<bool> Logic { get; private set; } = null!;

        /// <summary>
        /// The ID of the logical Library.
        /// </summary>
        [Output("logicDatabaseId")]
        public Output<string> LogicDatabaseId { get; private set; } = null!;

        /// <summary>
        /// The user ID list of the logical library Owner.
        /// </summary>
        [Output("ownerIdLists")]
        public Output<ImmutableArray<string>> OwnerIdLists { get; private set; } = null!;

        /// <summary>
        /// The nickname list of the logical library Owner.
        /// </summary>
        [Output("ownerNameLists")]
        public Output<ImmutableArray<string>> OwnerNameLists { get; private set; } = null!;

        /// <summary>
        /// Logical Library name.
        /// </summary>
        [Output("schemaName")]
        public Output<string> SchemaName { get; private set; } = null!;

        /// <summary>
        /// Logical library search name.
        /// </summary>
        [Output("searchName")]
        public Output<string> SearchName { get; private set; } = null!;


        /// <summary>
        /// Create a EnterpriseLogicDatabase resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EnterpriseLogicDatabase(string name, EnterpriseLogicDatabaseArgs args, CustomResourceOptions? options = null)
            : base("alicloud:dms/enterpriseLogicDatabase:EnterpriseLogicDatabase", name, args ?? new EnterpriseLogicDatabaseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EnterpriseLogicDatabase(string name, Input<string> id, EnterpriseLogicDatabaseState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:dms/enterpriseLogicDatabase:EnterpriseLogicDatabase", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing EnterpriseLogicDatabase resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EnterpriseLogicDatabase Get(string name, Input<string> id, EnterpriseLogicDatabaseState? state = null, CustomResourceOptions? options = null)
        {
            return new EnterpriseLogicDatabase(name, id, state, options);
        }
    }

    public sealed class EnterpriseLogicDatabaseArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Logical Library alias.
        /// </summary>
        [Input("alias", required: true)]
        public Input<string> Alias { get; set; } = null!;

        [Input("databaseIds", required: true)]
        private InputList<string>? _databaseIds;

        /// <summary>
        /// Sub-Database ID
        /// </summary>
        public InputList<string> DatabaseIds
        {
            get => _databaseIds ?? (_databaseIds = new InputList<string>());
            set => _databaseIds = value;
        }

        /// <summary>
        /// The ID of the logical Library.
        /// </summary>
        [Input("logicDatabaseId")]
        public Input<string>? LogicDatabaseId { get; set; }

        public EnterpriseLogicDatabaseArgs()
        {
        }
        public static new EnterpriseLogicDatabaseArgs Empty => new EnterpriseLogicDatabaseArgs();
    }

    public sealed class EnterpriseLogicDatabaseState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Logical Library alias.
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        [Input("databaseIds")]
        private InputList<string>? _databaseIds;

        /// <summary>
        /// Sub-Database ID
        /// </summary>
        public InputList<string> DatabaseIds
        {
            get => _databaseIds ?? (_databaseIds = new InputList<string>());
            set => _databaseIds = value;
        }

        /// <summary>
        /// Database type.
        /// </summary>
        [Input("dbType")]
        public Input<string>? DbType { get; set; }

        /// <summary>
        /// Environment type, return value is as follows:-product: production environment-dev: development environment-pre: Advance Environment-test: test environment-sit:SIT environment-uat:UAT environment-pet: Pressure measurement environment-stag:STAG environment
        /// </summary>
        [Input("envType")]
        public Input<string>? EnvType { get; set; }

        /// <summary>
        /// Whether it is a logical Library, the return value is true.
        /// </summary>
        [Input("logic")]
        public Input<bool>? Logic { get; set; }

        /// <summary>
        /// The ID of the logical Library.
        /// </summary>
        [Input("logicDatabaseId")]
        public Input<string>? LogicDatabaseId { get; set; }

        [Input("ownerIdLists")]
        private InputList<string>? _ownerIdLists;

        /// <summary>
        /// The user ID list of the logical library Owner.
        /// </summary>
        public InputList<string> OwnerIdLists
        {
            get => _ownerIdLists ?? (_ownerIdLists = new InputList<string>());
            set => _ownerIdLists = value;
        }

        [Input("ownerNameLists")]
        private InputList<string>? _ownerNameLists;

        /// <summary>
        /// The nickname list of the logical library Owner.
        /// </summary>
        public InputList<string> OwnerNameLists
        {
            get => _ownerNameLists ?? (_ownerNameLists = new InputList<string>());
            set => _ownerNameLists = value;
        }

        /// <summary>
        /// Logical Library name.
        /// </summary>
        [Input("schemaName")]
        public Input<string>? SchemaName { get; set; }

        /// <summary>
        /// Logical library search name.
        /// </summary>
        [Input("searchName")]
        public Input<string>? SearchName { get; set; }

        public EnterpriseLogicDatabaseState()
        {
        }
        public static new EnterpriseLogicDatabaseState Empty => new EnterpriseLogicDatabaseState();
    }
}

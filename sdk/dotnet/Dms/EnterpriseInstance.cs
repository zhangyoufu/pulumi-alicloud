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
    /// Provides a DMS Enterprise Instance resource.
    /// 
    /// &gt; **NOTE:** API users must first register in DMS.
    /// **NOTE:** Available in 1.81.0+.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var @default = new AliCloud.Dms.EnterpriseInstance("default", new AliCloud.Dms.EnterpriseInstanceArgs
    ///         {
    ///             DatabasePassword = "Yourpassword123",
    ///             DatabaseUser = "your_user_name",
    ///             DbaUid = "1182725234xxxxxxx",
    ///             EcsRegion = "cn-shanghai",
    ///             EnvType = "test",
    ///             ExportTimeout = 600,
    ///             Host = "rm-uf648hgsxxxxxx.mysql.rds.aliyuncs.com",
    ///             InstanceAlias = "your_alias_name",
    ///             InstanceSource = "RDS",
    ///             InstanceType = "MySQL",
    ///             NetworkType = "VPC",
    ///             Port = 3306,
    ///             QueryTimeout = 60,
    ///             SafeRule = "自由操作",
    ///             Tid = 12345,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class EnterpriseInstance : Pulumi.CustomResource
    {
        /// <summary>
        /// Cross-database query datalink name.
        /// </summary>
        [Output("dataLinkName")]
        public Output<string?> DataLinkName { get; private set; } = null!;

        /// <summary>
        /// Database access password.
        /// </summary>
        [Output("databasePassword")]
        public Output<string> DatabasePassword { get; private set; } = null!;

        /// <summary>
        /// Database access account.
        /// </summary>
        [Output("databaseUser")]
        public Output<string> DatabaseUser { get; private set; } = null!;

        [Output("dbaId")]
        public Output<string> DbaId { get; private set; } = null!;

        /// <summary>
        /// The instance dba nickname.
        /// </summary>
        [Output("dbaNickName")]
        public Output<string> DbaNickName { get; private set; } = null!;

        /// <summary>
        /// The DBA of the instance is passed into the Alibaba Cloud uid of the DBA.
        /// </summary>
        [Output("dbaUid")]
        public Output<int> DbaUid { get; private set; } = null!;

        /// <summary>
        /// Whether to use online services, currently only supports MySQL and PolarDB. Valid values: `0` Not used, `1` Native online DDL priority, `2` DMS lock-free table structure change priority.
        /// </summary>
        [Output("ddlOnline")]
        public Output<int?> DdlOnline { get; private set; } = null!;

        /// <summary>
        /// ECS instance ID. The value of InstanceSource is the ECS self-built library. This value must be passed.
        /// </summary>
        [Output("ecsInstanceId")]
        public Output<string> EcsInstanceId { get; private set; } = null!;

        /// <summary>
        /// The region where the instance is located. This value must be passed when the value of InstanceSource is RDS, ECS self-built library, and VPC dedicated line IDC.
        /// </summary>
        [Output("ecsRegion")]
        public Output<string?> EcsRegion { get; private set; } = null!;

        /// <summary>
        /// Environment type. Valid values: `product` production environment, `dev` development environment, `pre` pre-release environment, `test` test environment, `sit` SIT environment, `uat` UAT environment, `pet` pressure test environment, `stag` STAG environment.
        /// </summary>
        [Output("envType")]
        public Output<string> EnvType { get; private set; } = null!;

        /// <summary>
        /// Export timeout, unit: s (seconds).
        /// </summary>
        [Output("exportTimeout")]
        public Output<int> ExportTimeout { get; private set; } = null!;

        /// <summary>
        /// Host address of the target database.
        /// </summary>
        [Output("host")]
        public Output<string> Host { get; private set; } = null!;

        /// <summary>
        /// Instance alias, to help users quickly distinguish positioning.
        /// </summary>
        [Output("instanceAlias")]
        public Output<string> InstanceAlias { get; private set; } = null!;

        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The source of the database instance. Valid values: `PUBLIC_OWN`, `RDS`, `ECS_OWN`, `VPC_IDC`.
        /// </summary>
        [Output("instanceSource")]
        public Output<string> InstanceSource { get; private set; } = null!;

        /// <summary>
        /// Database type. Valid values: `MySQL`, `SQLServer`, `PostgreSQL`, `Oracle,` `DRDS`, `OceanBase`, `Mongo`, `Redis`.
        /// </summary>
        [Output("instanceType")]
        public Output<string> InstanceType { get; private set; } = null!;

        /// <summary>
        /// Network type. Valid values: `CLASSIC`, `VPC`.
        /// </summary>
        [Output("networkType")]
        public Output<string> NetworkType { get; private set; } = null!;

        /// <summary>
        /// Access port of the target database.
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        /// <summary>
        /// Query timeout time, unit: s (seconds).
        /// </summary>
        [Output("queryTimeout")]
        public Output<int> QueryTimeout { get; private set; } = null!;

        /// <summary>
        /// The security rule of the instance is passed into the name of the security rule in the enterprise.
        /// </summary>
        [Output("safeRule")]
        public Output<string> SafeRule { get; private set; } = null!;

        [Output("safeRuleId")]
        public Output<string> SafeRuleId { get; private set; } = null!;

        /// <summary>
        /// The SID. This value must be passed when InstanceType is PostgreSQL or Oracle.
        /// </summary>
        [Output("sid")]
        public Output<string?> Sid { get; private set; } = null!;

        /// <summary>
        /// The instance status.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The tenant ID.
        /// </summary>
        [Output("tid")]
        public Output<int?> Tid { get; private set; } = null!;

        /// <summary>
        /// Whether to enable cross-instance query. Valid values: `0` not open, `1` open.
        /// </summary>
        [Output("useDsql")]
        public Output<int?> UseDsql { get; private set; } = null!;

        /// <summary>
        /// VPC ID. This value must be passed when the value of InstanceSource is VPC dedicated line IDC.
        /// </summary>
        [Output("vpcId")]
        public Output<string?> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a EnterpriseInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EnterpriseInstance(string name, EnterpriseInstanceArgs args, CustomResourceOptions? options = null)
            : base("alicloud:dms/enterpriseInstance:EnterpriseInstance", name, args ?? new EnterpriseInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EnterpriseInstance(string name, Input<string> id, EnterpriseInstanceState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:dms/enterpriseInstance:EnterpriseInstance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EnterpriseInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EnterpriseInstance Get(string name, Input<string> id, EnterpriseInstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new EnterpriseInstance(name, id, state, options);
        }
    }

    public sealed class EnterpriseInstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cross-database query datalink name.
        /// </summary>
        [Input("dataLinkName")]
        public Input<string>? DataLinkName { get; set; }

        /// <summary>
        /// Database access password.
        /// </summary>
        [Input("databasePassword", required: true)]
        public Input<string> DatabasePassword { get; set; } = null!;

        /// <summary>
        /// Database access account.
        /// </summary>
        [Input("databaseUser", required: true)]
        public Input<string> DatabaseUser { get; set; } = null!;

        [Input("dbaId")]
        public Input<string>? DbaId { get; set; }

        /// <summary>
        /// The DBA of the instance is passed into the Alibaba Cloud uid of the DBA.
        /// </summary>
        [Input("dbaUid", required: true)]
        public Input<int> DbaUid { get; set; } = null!;

        /// <summary>
        /// Whether to use online services, currently only supports MySQL and PolarDB. Valid values: `0` Not used, `1` Native online DDL priority, `2` DMS lock-free table structure change priority.
        /// </summary>
        [Input("ddlOnline")]
        public Input<int>? DdlOnline { get; set; }

        /// <summary>
        /// ECS instance ID. The value of InstanceSource is the ECS self-built library. This value must be passed.
        /// </summary>
        [Input("ecsInstanceId")]
        public Input<string>? EcsInstanceId { get; set; }

        /// <summary>
        /// The region where the instance is located. This value must be passed when the value of InstanceSource is RDS, ECS self-built library, and VPC dedicated line IDC.
        /// </summary>
        [Input("ecsRegion")]
        public Input<string>? EcsRegion { get; set; }

        /// <summary>
        /// Environment type. Valid values: `product` production environment, `dev` development environment, `pre` pre-release environment, `test` test environment, `sit` SIT environment, `uat` UAT environment, `pet` pressure test environment, `stag` STAG environment.
        /// </summary>
        [Input("envType", required: true)]
        public Input<string> EnvType { get; set; } = null!;

        /// <summary>
        /// Export timeout, unit: s (seconds).
        /// </summary>
        [Input("exportTimeout", required: true)]
        public Input<int> ExportTimeout { get; set; } = null!;

        /// <summary>
        /// Host address of the target database.
        /// </summary>
        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        /// <summary>
        /// Instance alias, to help users quickly distinguish positioning.
        /// </summary>
        [Input("instanceAlias", required: true)]
        public Input<string> InstanceAlias { get; set; } = null!;

        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The source of the database instance. Valid values: `PUBLIC_OWN`, `RDS`, `ECS_OWN`, `VPC_IDC`.
        /// </summary>
        [Input("instanceSource", required: true)]
        public Input<string> InstanceSource { get; set; } = null!;

        /// <summary>
        /// Database type. Valid values: `MySQL`, `SQLServer`, `PostgreSQL`, `Oracle,` `DRDS`, `OceanBase`, `Mongo`, `Redis`.
        /// </summary>
        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        /// <summary>
        /// Network type. Valid values: `CLASSIC`, `VPC`.
        /// </summary>
        [Input("networkType", required: true)]
        public Input<string> NetworkType { get; set; } = null!;

        /// <summary>
        /// Access port of the target database.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        /// <summary>
        /// Query timeout time, unit: s (seconds).
        /// </summary>
        [Input("queryTimeout", required: true)]
        public Input<int> QueryTimeout { get; set; } = null!;

        /// <summary>
        /// The security rule of the instance is passed into the name of the security rule in the enterprise.
        /// </summary>
        [Input("safeRule", required: true)]
        public Input<string> SafeRule { get; set; } = null!;

        [Input("safeRuleId")]
        public Input<string>? SafeRuleId { get; set; }

        /// <summary>
        /// The SID. This value must be passed when InstanceType is PostgreSQL or Oracle.
        /// </summary>
        [Input("sid")]
        public Input<string>? Sid { get; set; }

        /// <summary>
        /// The tenant ID.
        /// </summary>
        [Input("tid")]
        public Input<int>? Tid { get; set; }

        /// <summary>
        /// Whether to enable cross-instance query. Valid values: `0` not open, `1` open.
        /// </summary>
        [Input("useDsql")]
        public Input<int>? UseDsql { get; set; }

        /// <summary>
        /// VPC ID. This value must be passed when the value of InstanceSource is VPC dedicated line IDC.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public EnterpriseInstanceArgs()
        {
        }
    }

    public sealed class EnterpriseInstanceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cross-database query datalink name.
        /// </summary>
        [Input("dataLinkName")]
        public Input<string>? DataLinkName { get; set; }

        /// <summary>
        /// Database access password.
        /// </summary>
        [Input("databasePassword")]
        public Input<string>? DatabasePassword { get; set; }

        /// <summary>
        /// Database access account.
        /// </summary>
        [Input("databaseUser")]
        public Input<string>? DatabaseUser { get; set; }

        [Input("dbaId")]
        public Input<string>? DbaId { get; set; }

        /// <summary>
        /// The instance dba nickname.
        /// </summary>
        [Input("dbaNickName")]
        public Input<string>? DbaNickName { get; set; }

        /// <summary>
        /// The DBA of the instance is passed into the Alibaba Cloud uid of the DBA.
        /// </summary>
        [Input("dbaUid")]
        public Input<int>? DbaUid { get; set; }

        /// <summary>
        /// Whether to use online services, currently only supports MySQL and PolarDB. Valid values: `0` Not used, `1` Native online DDL priority, `2` DMS lock-free table structure change priority.
        /// </summary>
        [Input("ddlOnline")]
        public Input<int>? DdlOnline { get; set; }

        /// <summary>
        /// ECS instance ID. The value of InstanceSource is the ECS self-built library. This value must be passed.
        /// </summary>
        [Input("ecsInstanceId")]
        public Input<string>? EcsInstanceId { get; set; }

        /// <summary>
        /// The region where the instance is located. This value must be passed when the value of InstanceSource is RDS, ECS self-built library, and VPC dedicated line IDC.
        /// </summary>
        [Input("ecsRegion")]
        public Input<string>? EcsRegion { get; set; }

        /// <summary>
        /// Environment type. Valid values: `product` production environment, `dev` development environment, `pre` pre-release environment, `test` test environment, `sit` SIT environment, `uat` UAT environment, `pet` pressure test environment, `stag` STAG environment.
        /// </summary>
        [Input("envType")]
        public Input<string>? EnvType { get; set; }

        /// <summary>
        /// Export timeout, unit: s (seconds).
        /// </summary>
        [Input("exportTimeout")]
        public Input<int>? ExportTimeout { get; set; }

        /// <summary>
        /// Host address of the target database.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// Instance alias, to help users quickly distinguish positioning.
        /// </summary>
        [Input("instanceAlias")]
        public Input<string>? InstanceAlias { get; set; }

        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The source of the database instance. Valid values: `PUBLIC_OWN`, `RDS`, `ECS_OWN`, `VPC_IDC`.
        /// </summary>
        [Input("instanceSource")]
        public Input<string>? InstanceSource { get; set; }

        /// <summary>
        /// Database type. Valid values: `MySQL`, `SQLServer`, `PostgreSQL`, `Oracle,` `DRDS`, `OceanBase`, `Mongo`, `Redis`.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// Network type. Valid values: `CLASSIC`, `VPC`.
        /// </summary>
        [Input("networkType")]
        public Input<string>? NetworkType { get; set; }

        /// <summary>
        /// Access port of the target database.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Query timeout time, unit: s (seconds).
        /// </summary>
        [Input("queryTimeout")]
        public Input<int>? QueryTimeout { get; set; }

        /// <summary>
        /// The security rule of the instance is passed into the name of the security rule in the enterprise.
        /// </summary>
        [Input("safeRule")]
        public Input<string>? SafeRule { get; set; }

        [Input("safeRuleId")]
        public Input<string>? SafeRuleId { get; set; }

        /// <summary>
        /// The SID. This value must be passed when InstanceType is PostgreSQL or Oracle.
        /// </summary>
        [Input("sid")]
        public Input<string>? Sid { get; set; }

        /// <summary>
        /// The instance status.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The tenant ID.
        /// </summary>
        [Input("tid")]
        public Input<int>? Tid { get; set; }

        /// <summary>
        /// Whether to enable cross-instance query. Valid values: `0` not open, `1` open.
        /// </summary>
        [Input("useDsql")]
        public Input<int>? UseDsql { get; set; }

        /// <summary>
        /// VPC ID. This value must be passed when the value of InstanceSource is VPC dedicated line IDC.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public EnterpriseInstanceState()
        {
        }
    }
}

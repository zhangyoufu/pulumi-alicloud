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
    /// 
    /// &gt; **NOTE:** Available since v1.81.0.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "tf-example";
    ///     var current = AliCloud.GetAccount.Invoke();
    /// 
    ///     var defaultRegions = AliCloud.GetRegions.Invoke(new()
    ///     {
    ///         Current = true,
    ///     });
    /// 
    ///     var defaultUserTenants = AliCloud.Dms.GetUserTenants.Invoke(new()
    ///     {
    ///         Status = "ACTIVE",
    ///     });
    /// 
    ///     var defaultZones = AliCloud.Rds.GetZones.Invoke(new()
    ///     {
    ///         Engine = "MySQL",
    ///         EngineVersion = "8.0",
    ///         InstanceChargeType = "PostPaid",
    ///         Category = "HighAvailability",
    ///         DbInstanceStorageType = "cloud_essd",
    ///     });
    /// 
    ///     var defaultInstanceClasses = AliCloud.Rds.GetInstanceClasses.Invoke(new()
    ///     {
    ///         ZoneId = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///         Engine = "MySQL",
    ///         EngineVersion = "8.0",
    ///         Category = "HighAvailability",
    ///         DbInstanceStorageType = "cloud_essd",
    ///         InstanceChargeType = "PostPaid",
    ///     });
    /// 
    ///     var defaultNetwork = new AliCloud.Vpc.Network("defaultNetwork", new()
    ///     {
    ///         VpcName = name,
    ///         CidrBlock = "10.4.0.0/16",
    ///     });
    /// 
    ///     var defaultSwitch = new AliCloud.Vpc.Switch("defaultSwitch", new()
    ///     {
    ///         VswitchName = name,
    ///         CidrBlock = "10.4.0.0/24",
    ///         VpcId = defaultNetwork.Id,
    ///         ZoneId = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///     });
    /// 
    ///     var defaultSecurityGroup = new AliCloud.Ecs.SecurityGroup("defaultSecurityGroup", new()
    ///     {
    ///         VpcId = defaultNetwork.Id,
    ///     });
    /// 
    ///     var defaultInstance = new AliCloud.Rds.Instance("defaultInstance", new()
    ///     {
    ///         Engine = "MySQL",
    ///         EngineVersion = "8.0",
    ///         DbInstanceStorageType = "cloud_essd",
    ///         InstanceType = defaultInstanceClasses.Apply(getInstanceClassesResult =&gt; getInstanceClassesResult.InstanceClasses[0]?.InstanceClass),
    ///         InstanceStorage = defaultInstanceClasses.Apply(getInstanceClassesResult =&gt; getInstanceClassesResult.InstanceClasses[0]?.StorageRange?.Min),
    ///         VswitchId = defaultSwitch.Id,
    ///         InstanceName = name,
    ///         SecurityIps = new[]
    ///         {
    ///             "100.104.5.0/24",
    ///             "192.168.0.6",
    ///         },
    ///         Tags = 
    ///         {
    ///             { "Created", "TF" },
    ///             { "For", "example" },
    ///         },
    ///     });
    /// 
    ///     var defaultAccount = new AliCloud.Rds.Account("defaultAccount", new()
    ///     {
    ///         DbInstanceId = defaultInstance.Id,
    ///         AccountName = "tfexamplename",
    ///         AccountPassword = "Example12345",
    ///         AccountType = "Normal",
    ///     });
    /// 
    ///     var defaultEnterpriseInstance = new AliCloud.Dms.EnterpriseInstance("defaultEnterpriseInstance", new()
    ///     {
    ///         Tid = defaultUserTenants.Apply(getUserTenantsResult =&gt; getUserTenantsResult.Ids[0]),
    ///         InstanceType = "MySQL",
    ///         InstanceSource = "RDS",
    ///         NetworkType = "VPC",
    ///         EnvType = "dev",
    ///         Host = defaultInstance.ConnectionString,
    ///         Port = 3306,
    ///         DatabaseUser = defaultAccount.AccountName,
    ///         DatabasePassword = defaultAccount.AccountPassword,
    ///         InstanceName = name,
    ///         DbaUid = current.Apply(getAccountResult =&gt; getAccountResult.Id),
    ///         SafeRule = "自由操作",
    ///         QueryTimeout = 60,
    ///         ExportTimeout = 600,
    ///         EcsRegion = defaultRegions.Apply(getRegionsResult =&gt; getRegionsResult.Regions[0]?.Id),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// DMS Enterprise can be imported using host and port, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:dms/enterpriseInstance:EnterpriseInstance example rm-uf648hgs7874xxxx.mysql.rds.aliyuncs.com:3306
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:dms/enterpriseInstance:EnterpriseInstance")]
    public partial class EnterpriseInstance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Cross-database query datalink name.
        /// </summary>
        [Output("dataLinkName")]
        public Output<string> DataLinkName { get; private set; } = null!;

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

        /// <summary>
        /// The dba id of the database instance.
        /// </summary>
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
        /// Field `instance_alias` has been deprecated from version 1.100.0. Use `instance_name` instead.
        /// </summary>
        [Output("instanceAlias")]
        public Output<string> InstanceAlias { get; private set; } = null!;

        /// <summary>
        /// The instance id of the database instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Instance name, to help users quickly distinguish positioning.
        /// </summary>
        [Output("instanceName")]
        public Output<string> InstanceName { get; private set; } = null!;

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

        /// <summary>
        /// The safe rule id of the database instance.
        /// </summary>
        [Output("safeRuleId")]
        public Output<string> SafeRuleId { get; private set; } = null!;

        /// <summary>
        /// The SID. This value must be passed when InstanceType is PostgreSQL or Oracle.
        /// </summary>
        [Output("sid")]
        public Output<string?> Sid { get; private set; } = null!;

        /// <summary>
        /// Whether the instance ignores test connectivity. Valid values: `true`, `false`.
        /// </summary>
        [Output("skipTest")]
        public Output<bool?> SkipTest { get; private set; } = null!;

        /// <summary>
        /// It has been deprecated from provider version 1.100.0 and 'status' instead.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The instance status.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

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
                AdditionalSecretOutputs =
                {
                    "databasePassword",
                },
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

    public sealed class EnterpriseInstanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cross-database query datalink name.
        /// </summary>
        [Input("dataLinkName")]
        public Input<string>? DataLinkName { get; set; }

        [Input("databasePassword", required: true)]
        private Input<string>? _databasePassword;

        /// <summary>
        /// Database access password.
        /// </summary>
        public Input<string>? DatabasePassword
        {
            get => _databasePassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _databasePassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Database access account.
        /// </summary>
        [Input("databaseUser", required: true)]
        public Input<string> DatabaseUser { get; set; } = null!;

        /// <summary>
        /// The dba id of the database instance.
        /// </summary>
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
        /// Field `instance_alias` has been deprecated from version 1.100.0. Use `instance_name` instead.
        /// </summary>
        [Input("instanceAlias")]
        public Input<string>? InstanceAlias { get; set; }

        /// <summary>
        /// The instance id of the database instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Instance name, to help users quickly distinguish positioning.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

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

        /// <summary>
        /// The safe rule id of the database instance.
        /// </summary>
        [Input("safeRuleId")]
        public Input<string>? SafeRuleId { get; set; }

        /// <summary>
        /// The SID. This value must be passed when InstanceType is PostgreSQL or Oracle.
        /// </summary>
        [Input("sid")]
        public Input<string>? Sid { get; set; }

        /// <summary>
        /// Whether the instance ignores test connectivity. Valid values: `true`, `false`.
        /// </summary>
        [Input("skipTest")]
        public Input<bool>? SkipTest { get; set; }

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
        public static new EnterpriseInstanceArgs Empty => new EnterpriseInstanceArgs();
    }

    public sealed class EnterpriseInstanceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cross-database query datalink name.
        /// </summary>
        [Input("dataLinkName")]
        public Input<string>? DataLinkName { get; set; }

        [Input("databasePassword")]
        private Input<string>? _databasePassword;

        /// <summary>
        /// Database access password.
        /// </summary>
        public Input<string>? DatabasePassword
        {
            get => _databasePassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _databasePassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Database access account.
        /// </summary>
        [Input("databaseUser")]
        public Input<string>? DatabaseUser { get; set; }

        /// <summary>
        /// The dba id of the database instance.
        /// </summary>
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
        /// Field `instance_alias` has been deprecated from version 1.100.0. Use `instance_name` instead.
        /// </summary>
        [Input("instanceAlias")]
        public Input<string>? InstanceAlias { get; set; }

        /// <summary>
        /// The instance id of the database instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Instance name, to help users quickly distinguish positioning.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

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

        /// <summary>
        /// The safe rule id of the database instance.
        /// </summary>
        [Input("safeRuleId")]
        public Input<string>? SafeRuleId { get; set; }

        /// <summary>
        /// The SID. This value must be passed when InstanceType is PostgreSQL or Oracle.
        /// </summary>
        [Input("sid")]
        public Input<string>? Sid { get; set; }

        /// <summary>
        /// Whether the instance ignores test connectivity. Valid values: `true`, `false`.
        /// </summary>
        [Input("skipTest")]
        public Input<bool>? SkipTest { get; set; }

        /// <summary>
        /// It has been deprecated from provider version 1.100.0 and 'status' instead.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The instance status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

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
        public static new EnterpriseInstanceState Empty => new EnterpriseInstanceState();
    }
}

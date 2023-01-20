// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.KVStore
{
    /// <summary>
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
    ///     var config = new Config();
    ///     var creation = config.Get("creation") ?? "KVStore";
    ///     var multiAz = config.Get("multiAz") ?? "false";
    ///     var name = config.Get("name") ?? "kvstorebackuppolicyvpc";
    ///     var defaultZones = AliCloud.GetZones.Invoke(new()
    ///     {
    ///         AvailableResourceCreation = creation,
    ///     });
    /// 
    ///     var defaultNetwork = new AliCloud.Vpc.Network("defaultNetwork", new()
    ///     {
    ///         CidrBlock = "172.16.0.0/16",
    ///     });
    /// 
    ///     var defaultSwitch = new AliCloud.Vpc.Switch("defaultSwitch", new()
    ///     {
    ///         VpcId = defaultNetwork.Id,
    ///         CidrBlock = "172.16.0.0/24",
    ///         ZoneId = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///     });
    /// 
    ///     var defaultInstance = new AliCloud.KVStore.Instance("defaultInstance", new()
    ///     {
    ///         InstanceClass = "Memcache",
    ///         InstanceName = name,
    ///         VswitchId = defaultSwitch.Id,
    ///         PrivateIp = "172.16.0.10",
    ///         SecurityIps = new[]
    ///         {
    ///             "10.0.0.1",
    ///         },
    ///         InstanceType = "memcache.master.small.default",
    ///         EngineVersion = "2.8",
    ///     });
    /// 
    ///     var defaultBackupPolicy = new AliCloud.KVStore.BackupPolicy("defaultBackupPolicy", new()
    ///     {
    ///         InstanceId = defaultInstance.Id,
    ///         BackupPeriods = new[]
    ///         {
    ///             "Tuesday",
    ///             "Wednesday",
    ///         },
    ///         BackupTime = "10:00Z-11:00Z",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// KVStore backup policy can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:kvstore/backupPolicy:BackupPolicy example r-abc12345678
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:kvstore/backupPolicy:BackupPolicy")]
    public partial class BackupPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Backup Cycle. Allowed values: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday
        /// </summary>
        [Output("backupPeriods")]
        public Output<ImmutableArray<string>> BackupPeriods { get; private set; } = null!;

        /// <summary>
        /// Backup time, in the format of HH:mmZ- HH:mm Z
        /// </summary>
        [Output("backupTime")]
        public Output<string?> BackupTime { get; private set; } = null!;

        /// <summary>
        /// The id of ApsaraDB for Redis or Memcache intance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a BackupPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BackupPolicy(string name, BackupPolicyArgs args, CustomResourceOptions? options = null)
            : base("alicloud:kvstore/backupPolicy:BackupPolicy", name, args ?? new BackupPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BackupPolicy(string name, Input<string> id, BackupPolicyState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:kvstore/backupPolicy:BackupPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BackupPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BackupPolicy Get(string name, Input<string> id, BackupPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new BackupPolicy(name, id, state, options);
        }
    }

    public sealed class BackupPolicyArgs : global::Pulumi.ResourceArgs
    {
        [Input("backupPeriods")]
        private InputList<string>? _backupPeriods;

        /// <summary>
        /// Backup Cycle. Allowed values: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday
        /// </summary>
        public InputList<string> BackupPeriods
        {
            get => _backupPeriods ?? (_backupPeriods = new InputList<string>());
            set => _backupPeriods = value;
        }

        /// <summary>
        /// Backup time, in the format of HH:mmZ- HH:mm Z
        /// </summary>
        [Input("backupTime")]
        public Input<string>? BackupTime { get; set; }

        /// <summary>
        /// The id of ApsaraDB for Redis or Memcache intance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public BackupPolicyArgs()
        {
        }
        public static new BackupPolicyArgs Empty => new BackupPolicyArgs();
    }

    public sealed class BackupPolicyState : global::Pulumi.ResourceArgs
    {
        [Input("backupPeriods")]
        private InputList<string>? _backupPeriods;

        /// <summary>
        /// Backup Cycle. Allowed values: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday
        /// </summary>
        public InputList<string> BackupPeriods
        {
            get => _backupPeriods ?? (_backupPeriods = new InputList<string>());
            set => _backupPeriods = value;
        }

        /// <summary>
        /// Backup time, in the format of HH:mmZ- HH:mm Z
        /// </summary>
        [Input("backupTime")]
        public Input<string>? BackupTime { get; set; }

        /// <summary>
        /// The id of ApsaraDB for Redis or Memcache intance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        public BackupPolicyState()
        {
        }
        public static new BackupPolicyState Empty => new BackupPolicyState();
    }
}

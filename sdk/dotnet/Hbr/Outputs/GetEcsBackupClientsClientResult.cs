// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Hbr.Outputs
{

    [OutputType]
    public sealed class GetEcsBackupClientsClientResult
    {
        /// <summary>
        /// The system architecture of client, only the ECS File Backup Client is available. Valid values: `AMD64` , `386`.
        /// </summary>
        public readonly string ArchType;
        /// <summary>
        /// Client protected status.
        /// </summary>
        public readonly string BackupStatus;
        /// <summary>
        /// The type of client. Valid values: `ECS_CLIENT` (ECS File Backup Client).
        /// </summary>
        public readonly string ClientType;
        /// <summary>
        /// The version of client.
        /// </summary>
        public readonly string ClientVersion;
        /// <summary>
        /// The creation time of client. Unix time seconds.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The data plane access point type. Valid Values: `PUBLIC`, `VPC`, `CLASSIC`.
        /// </summary>
        public readonly string DataNetworkType;
        /// <summary>
        /// The data plane proxy settings. Valid Values: `DISABLE`, `USE_CONTROL_PROXY`, `CUSTOM`.
        /// </summary>
        public readonly string DataProxySetting;
        /// <summary>
        /// The first ID of the resource.
        /// </summary>
        public readonly string EcsBackupClientId;
        /// <summary>
        /// The name of ECS host.
        /// </summary>
        public readonly string Hostname;
        /// <summary>
        /// The ID of the Ecs Backup Client.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The ID of ECS instance.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// The name of ECS instance.
        /// </summary>
        public readonly string InstanceName;
        /// <summary>
        /// Client last heartbeat time. Unix Time Seconds.
        /// </summary>
        public readonly string LastHeartBeatTime;
        /// <summary>
        /// The latest client version.
        /// </summary>
        public readonly string MaxClientVersion;
        /// <summary>
        /// Number of CPU cores used by a single backup task, 0 means no restrictions.
        /// </summary>
        public readonly string MaxCpuCore;
        /// <summary>
        /// Number of concurrent jobs for a single backup task, 0 means no restrictions.
        /// </summary>
        public readonly string MaxWorker;
        /// <summary>
        /// The operating system type of client, only the ECS File Backup Client is available. Valid values: `windows`, `linux`.
        /// </summary>
        public readonly string OsType;
        /// <summary>
        /// Intranet IP address of the instance, only available for ECS file backup client.
        /// </summary>
        public readonly string PrivateIpv4;
        /// <summary>
        /// Custom data plane proxy server host address.
        /// </summary>
        public readonly string ProxyHost;
        /// <summary>
        /// Custom data plane proxy server password.
        /// </summary>
        public readonly string ProxyPassword;
        /// <summary>
        /// Custom data plane proxy server host port.
        /// </summary>
        public readonly string ProxyPort;
        /// <summary>
        /// Username of custom data plane proxy server.
        /// </summary>
        public readonly string ProxyUser;
        /// <summary>
        /// The status of the resource.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The update time of client. Unix Time Seconds.
        /// </summary>
        public readonly string UpdatedTime;
        /// <summary>
        /// Indicates whether to use the HTTPS protocol. Valid values: `true`, `false`.
        /// </summary>
        public readonly bool UseHttps;
        /// <summary>
        /// The ID of Zone.
        /// </summary>
        public readonly string ZoneId;

        [OutputConstructor]
        private GetEcsBackupClientsClientResult(
            string archType,

            string backupStatus,

            string clientType,

            string clientVersion,

            string createTime,

            string dataNetworkType,

            string dataProxySetting,

            string ecsBackupClientId,

            string hostname,

            string id,

            string instanceId,

            string instanceName,

            string lastHeartBeatTime,

            string maxClientVersion,

            string maxCpuCore,

            string maxWorker,

            string osType,

            string privateIpv4,

            string proxyHost,

            string proxyPassword,

            string proxyPort,

            string proxyUser,

            string status,

            string updatedTime,

            bool useHttps,

            string zoneId)
        {
            ArchType = archType;
            BackupStatus = backupStatus;
            ClientType = clientType;
            ClientVersion = clientVersion;
            CreateTime = createTime;
            DataNetworkType = dataNetworkType;
            DataProxySetting = dataProxySetting;
            EcsBackupClientId = ecsBackupClientId;
            Hostname = hostname;
            Id = id;
            InstanceId = instanceId;
            InstanceName = instanceName;
            LastHeartBeatTime = lastHeartBeatTime;
            MaxClientVersion = maxClientVersion;
            MaxCpuCore = maxCpuCore;
            MaxWorker = maxWorker;
            OsType = osType;
            PrivateIpv4 = privateIpv4;
            ProxyHost = proxyHost;
            ProxyPassword = proxyPassword;
            ProxyPort = proxyPort;
            ProxyUser = proxyUser;
            Status = status;
            UpdatedTime = updatedTime;
            UseHttps = useHttps;
            ZoneId = zoneId;
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Hbr
{
    public static class GetEcsBackupClients
    {
        /// <summary>
        /// This data source provides the Hbr Ecs Backup Clients of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.132.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic Usage
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var @default = Output.Create(AliCloud.Ecs.GetInstances.InvokeAsync(new AliCloud.Ecs.GetInstancesArgs
        ///         {
        ///             NameRegex = "ecs_instance_name",
        ///             Status = "Running",
        ///         }));
        ///         var ids = Output.Create(AliCloud.Hbr.GetEcsBackupClients.InvokeAsync(new AliCloud.Hbr.GetEcsBackupClientsArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 alicloud_hbr_ecs_backup_client.Default.Id,
        ///             },
        ///             InstanceIds = 
        ///             {
        ///                 alicloud_hbr_ecs_backup_client.Default.Instance_id,
        ///             },
        ///         }));
        ///         this.HbrEcsBackupClientId1 = ids.Apply(ids =&gt; ids.Clients[0].Id);
        ///     }
        /// 
        ///     [Output("hbrEcsBackupClientId1")]
        ///     public Output&lt;string&gt; HbrEcsBackupClientId1 { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetEcsBackupClientsResult> InvokeAsync(GetEcsBackupClientsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEcsBackupClientsResult>("alicloud:hbr/getEcsBackupClients:getEcsBackupClients", args ?? new GetEcsBackupClientsArgs(), options.WithVersion());
    }


    public sealed class GetEcsBackupClientsArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Ecs Backup Client IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        [Input("instanceIds")]
        private List<string>? _instanceIds;

        /// <summary>
        /// A list of ECS Instance IDs.
        /// </summary>
        public List<string> InstanceIds
        {
            get => _instanceIds ?? (_instanceIds = new List<string>());
            set => _instanceIds = value;
        }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetEcsBackupClientsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetEcsBackupClientsResult
    {
        public readonly ImmutableArray<Outputs.GetEcsBackupClientsClientResult> Clients;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly ImmutableArray<string> InstanceIds;
        public readonly string? OutputFile;
        public readonly string? Status;

        [OutputConstructor]
        private GetEcsBackupClientsResult(
            ImmutableArray<Outputs.GetEcsBackupClientsClientResult> clients,

            string id,

            ImmutableArray<string> ids,

            ImmutableArray<string> instanceIds,

            string? outputFile,

            string? status)
        {
            Clients = clients;
            Id = id;
            Ids = ids;
            InstanceIds = instanceIds;
            OutputFile = outputFile;
            Status = status;
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ecs
{
    public static partial class Invokes
    {
        /// <summary>
        /// This data source provides a list of key pairs in an Alibaba Cloud account according to the specified filters.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/key_pairs.html.markdown.
        /// </summary>
        [Obsolete("Use GetKeyPairs.InvokeAsync() instead")]
        public static Task<GetKeyPairsResult> GetKeyPairs(GetKeyPairsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetKeyPairsResult>("alicloud:ecs/getKeyPairs:getKeyPairs", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetKeyPairs
    {
        /// <summary>
        /// This data source provides a list of key pairs in an Alibaba Cloud account according to the specified filters.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/key_pairs.html.markdown.
        /// </summary>
        public static Task<GetKeyPairsResult> InvokeAsync(GetKeyPairsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetKeyPairsResult>("alicloud:ecs/getKeyPairs:getKeyPairs", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetKeyPairsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// A finger print used to retrieve specified key pair.
        /// </summary>
        [Input("fingerPrint")]
        public bool? FingerPrint { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of key pair IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to apply to the resulting key pairs.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The Id of resource group which the key pair belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public string? ResourceGroupId { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        public GetKeyPairsArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetKeyPairsResult
    {
        /// <summary>
        /// Finger print of the key pair.
        /// </summary>
        public readonly bool FingerPrint;
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// A list of key pairs. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetKeyPairsKeyPairsResult> KeyPairs;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of key pair names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        /// <summary>
        /// The Id of resource group.
        /// </summary>
        public readonly string? ResourceGroupId;
        /// <summary>
        /// (Optional, Available in v1.66.0+) A mapping of tags to assign to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Tags;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetKeyPairsResult(
            bool fingerPrint,
            ImmutableArray<string> ids,
            ImmutableArray<Outputs.GetKeyPairsKeyPairsResult> keyPairs,
            string? nameRegex,
            ImmutableArray<string> names,
            string? outputFile,
            string? resourceGroupId,
            ImmutableDictionary<string, object>? tags,
            string id)
        {
            FingerPrint = fingerPrint;
            Ids = ids;
            KeyPairs = keyPairs;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            ResourceGroupId = resourceGroupId;
            Tags = tags;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetKeyPairsKeyPairsInstancesResult
    {
        /// <summary>
        /// The ID of the availability zone where the ECS instance is located.
        /// </summary>
        public readonly string AvailabilityZone;
        public readonly string Description;
        public readonly string ImageId;
        /// <summary>
        /// The ID of the ECS instance.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// The name of the ECS instance.
        /// </summary>
        public readonly string InstanceName;
        public readonly string InstanceType;
        /// <summary>
        /// Name of the key pair.
        /// </summary>
        public readonly string KeyName;
        /// <summary>
        /// The private IP address of the ECS instance.
        /// </summary>
        public readonly string PrivateIp;
        /// <summary>
        /// The public IP address or EIP of the ECS instance.
        /// </summary>
        public readonly string PublicIp;
        public readonly string RegionId;
        public readonly string Status;
        /// <summary>
        /// The ID of the VSwitch attached to the ECS instance.
        /// </summary>
        public readonly string VswitchId;

        [OutputConstructor]
        private GetKeyPairsKeyPairsInstancesResult(
            string availabilityZone,
            string description,
            string imageId,
            string instanceId,
            string instanceName,
            string instanceType,
            string keyName,
            string privateIp,
            string publicIp,
            string regionId,
            string status,
            string vswitchId)
        {
            AvailabilityZone = availabilityZone;
            Description = description;
            ImageId = imageId;
            InstanceId = instanceId;
            InstanceName = instanceName;
            InstanceType = instanceType;
            KeyName = keyName;
            PrivateIp = privateIp;
            PublicIp = publicIp;
            RegionId = regionId;
            Status = status;
            VswitchId = vswitchId;
        }
    }

    [OutputType]
    public sealed class GetKeyPairsKeyPairsResult
    {
        /// <summary>
        /// A finger print used to retrieve specified key pair.
        /// </summary>
        public readonly string FingerPrint;
        /// <summary>
        /// ID of the key pair.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of ECS instances that has been bound this key pair.
        /// </summary>
        public readonly ImmutableArray<GetKeyPairsKeyPairsInstancesResult> Instances;
        /// <summary>
        /// Name of the key pair.
        /// </summary>
        public readonly string KeyName;
        /// <summary>
        /// The Id of resource group which the key pair belongs.
        /// </summary>
        public readonly string ResourceGroupId;
        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Tags;

        [OutputConstructor]
        private GetKeyPairsKeyPairsResult(
            string fingerPrint,
            string id,
            ImmutableArray<GetKeyPairsKeyPairsInstancesResult> instances,
            string keyName,
            string resourceGroupId,
            ImmutableDictionary<string, object>? tags)
        {
            FingerPrint = fingerPrint;
            Id = id;
            Instances = instances;
            KeyName = keyName;
            ResourceGroupId = resourceGroupId;
            Tags = tags;
        }
    }
    }
}

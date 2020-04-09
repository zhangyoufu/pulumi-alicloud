// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Slb
{
    public static partial class Invokes
    {
        /// <summary>
        /// This data source provides the CA certificate list.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/slb_ca_certificates.html.markdown.
        /// </summary>
        [Obsolete("Use GetCaCertificates.InvokeAsync() instead")]
        public static Task<GetCaCertificatesResult> GetCaCertificates(GetCaCertificatesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCaCertificatesResult>("alicloud:slb/getCaCertificates:getCaCertificates", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetCaCertificates
    {
        /// <summary>
        /// This data source provides the CA certificate list.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/slb_ca_certificates.html.markdown.
        /// </summary>
        public static Task<GetCaCertificatesResult> InvokeAsync(GetCaCertificatesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCaCertificatesResult>("alicloud:slb/getCaCertificates:getCaCertificates", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetCaCertificatesArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of ca certificates IDs to filter results.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by ca certificate name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The Id of resource group which ca certificates belongs.
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

        public GetCaCertificatesArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetCaCertificatesResult
    {
        /// <summary>
        /// A list of SLB ca certificates. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCaCertificatesCertificatesResult> Certificates;
        /// <summary>
        /// A list of SLB ca certificates IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of SLB ca certificates names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        /// <summary>
        /// The resource group Id of CA certificate.
        /// </summary>
        public readonly string? ResourceGroupId;
        /// <summary>
        /// (Available in v1.66.0+) A mapping of tags to assign to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Tags;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetCaCertificatesResult(
            ImmutableArray<Outputs.GetCaCertificatesCertificatesResult> certificates,
            ImmutableArray<string> ids,
            string? nameRegex,
            ImmutableArray<string> names,
            string? outputFile,
            string? resourceGroupId,
            ImmutableDictionary<string, object>? tags,
            string id)
        {
            Certificates = certificates;
            Ids = ids;
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
    public sealed class GetCaCertificatesCertificatesResult
    {
        /// <summary>
        /// CA certificate common name.
        /// </summary>
        public readonly string CommonName;
        /// <summary>
        /// CA certificate created time.
        /// </summary>
        public readonly string CreatedTime;
        /// <summary>
        /// CA certificate created timestamp.
        /// </summary>
        public readonly int CreatedTimestamp;
        /// <summary>
        /// CA certificate expired time.
        /// </summary>
        public readonly string ExpiredTime;
        /// <summary>
        /// CA certificate expired timestamp.
        /// </summary>
        public readonly int ExpiredTimestamp;
        /// <summary>
        /// CA certificate fingerprint.
        /// </summary>
        public readonly string Fingerprint;
        /// <summary>
        /// CA certificate ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// CA certificate name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The region Id of CA certificate.
        /// </summary>
        public readonly string RegionId;
        /// <summary>
        /// The Id of resource group which ca certificates belongs.
        /// </summary>
        public readonly string ResourceGroupId;
        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Tags;

        [OutputConstructor]
        private GetCaCertificatesCertificatesResult(
            string commonName,
            string createdTime,
            int createdTimestamp,
            string expiredTime,
            int expiredTimestamp,
            string fingerprint,
            string id,
            string name,
            string regionId,
            string resourceGroupId,
            ImmutableDictionary<string, object>? tags)
        {
            CommonName = commonName;
            CreatedTime = createdTime;
            CreatedTimestamp = createdTimestamp;
            ExpiredTime = expiredTime;
            ExpiredTimestamp = expiredTimestamp;
            Fingerprint = fingerprint;
            Id = id;
            Name = name;
            RegionId = regionId;
            ResourceGroupId = resourceGroupId;
            Tags = tags;
        }
    }
    }
}

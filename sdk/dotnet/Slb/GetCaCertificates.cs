// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Slb
{
    public static class GetCaCertificates
    {
        /// <summary>
        /// This data source provides the CA certificate list.
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCaCertificatesResult> InvokeAsync(GetCaCertificatesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCaCertificatesResult>("alicloud:slb/getCaCertificates:getCaCertificates", args ?? new GetCaCertificatesArgs(), options.WithVersion());
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
        public readonly ImmutableArray<Outputs.GetCaCertificatesCertificateResult> Certificates;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
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

        [OutputConstructor]
        private GetCaCertificatesResult(
            ImmutableArray<Outputs.GetCaCertificatesCertificateResult> certificates,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? resourceGroupId,

            ImmutableDictionary<string, object>? tags)
        {
            Certificates = certificates;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            ResourceGroupId = resourceGroupId;
            Tags = tags;
        }
    }
}

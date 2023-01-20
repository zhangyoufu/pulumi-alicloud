// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Dcdn
{
    public static class GetDomains
    {
        /// <summary>
        /// Provides a collection of DCDN Domains to the specified filters.
        /// 
        /// &gt; **NOTE:** Available in 1.94.0+.
        /// </summary>
        public static Task<GetDomainsResult> InvokeAsync(GetDomainsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDomainsResult>("alicloud:dcdn/getDomains:getDomains", args ?? new GetDomainsArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a collection of DCDN Domains to the specified filters.
        /// 
        /// &gt; **NOTE:** Available in 1.94.0+.
        /// </summary>
        public static Output<GetDomainsResult> Invoke(GetDomainsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDomainsResult>("alicloud:dcdn/getDomains:getDomains", args ?? new GetDomainsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDomainsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The end time of the update. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
        /// </summary>
        [Input("changeEndTime")]
        public string? ChangeEndTime { get; set; }

        /// <summary>
        /// The start time of the update. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
        /// </summary>
        [Input("changeStartTime")]
        public string? ChangeStartTime { get; set; }

        /// <summary>
        /// Specifies whether to display the domains in the checking, check_failed, or configure_failed status. Valid values: `true` or `false`.
        /// </summary>
        [Input("checkDomainShow")]
        public bool? CheckDomainShow { get; set; }

        /// <summary>
        /// The search method. Default value: `fuzzy_match`. Valid values: `fuzzy_match`, `pre_match`, `suf_match`, `full_match`.
        /// </summary>
        [Input("domainSearchType")]
        public string? DomainSearchType { get; set; }

        /// <summary>
        /// Default to `false`. Set it to true can output more details.
        /// </summary>
        [Input("enableDetails")]
        public bool? EnableDetails { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list ids of DCDN Domain.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by the DCDN Domain.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public string? ResourceGroupId { get; set; }

        [Input("securityToken")]
        public string? SecurityToken { get; set; }

        /// <summary>
        /// The status of DCDN Domain.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetDomainsArgs()
        {
        }
        public static new GetDomainsArgs Empty => new GetDomainsArgs();
    }

    public sealed class GetDomainsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The end time of the update. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
        /// </summary>
        [Input("changeEndTime")]
        public Input<string>? ChangeEndTime { get; set; }

        /// <summary>
        /// The start time of the update. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
        /// </summary>
        [Input("changeStartTime")]
        public Input<string>? ChangeStartTime { get; set; }

        /// <summary>
        /// Specifies whether to display the domains in the checking, check_failed, or configure_failed status. Valid values: `true` or `false`.
        /// </summary>
        [Input("checkDomainShow")]
        public Input<bool>? CheckDomainShow { get; set; }

        /// <summary>
        /// The search method. Default value: `fuzzy_match`. Valid values: `fuzzy_match`, `pre_match`, `suf_match`, `full_match`.
        /// </summary>
        [Input("domainSearchType")]
        public Input<string>? DomainSearchType { get; set; }

        /// <summary>
        /// Default to `false`. Set it to true can output more details.
        /// </summary>
        [Input("enableDetails")]
        public Input<bool>? EnableDetails { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list ids of DCDN Domain.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by the DCDN Domain.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("securityToken")]
        public Input<string>? SecurityToken { get; set; }

        /// <summary>
        /// The status of DCDN Domain.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetDomainsInvokeArgs()
        {
        }
        public static new GetDomainsInvokeArgs Empty => new GetDomainsInvokeArgs();
    }


    [OutputType]
    public sealed class GetDomainsResult
    {
        public readonly string? ChangeEndTime;
        public readonly string? ChangeStartTime;
        public readonly bool? CheckDomainShow;
        public readonly string? DomainSearchType;
        /// <summary>
        /// A list of domains. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDomainsDomainResult> Domains;
        public readonly bool? EnableDetails;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list ids of DCDN Domain.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of DCDN Domain names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        public readonly string? ResourceGroupId;
        public readonly string? SecurityToken;
        /// <summary>
        /// The status of DCDN Domain. Valid values: `online`, `offline`, `check_failed`, `checking`, `configure_failed`, `configuring`.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private GetDomainsResult(
            string? changeEndTime,

            string? changeStartTime,

            bool? checkDomainShow,

            string? domainSearchType,

            ImmutableArray<Outputs.GetDomainsDomainResult> domains,

            bool? enableDetails,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? resourceGroupId,

            string? securityToken,

            string? status)
        {
            ChangeEndTime = changeEndTime;
            ChangeStartTime = changeStartTime;
            CheckDomainShow = checkDomainShow;
            DomainSearchType = domainSearchType;
            Domains = domains;
            EnableDetails = enableDetails;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            ResourceGroupId = resourceGroupId;
            SecurityToken = securityToken;
            Status = status;
        }
    }
}

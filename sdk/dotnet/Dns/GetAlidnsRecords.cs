// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Dns
{
    public static class GetAlidnsRecords
    {
        /// <summary>
        /// This data source provides a list of Alidns Domain Records in an Alibaba Cloud account according to the specified filters.
        /// 
        /// &gt; **NOTE:**  Available in 1.86.0+.
        /// 
        /// ## Example Usage
        /// 
        /// ```terraform 
        /// data "alicloud_alidns_records" "records_ds" {
        ///   domain_name = "xiaozhu.top"
        ///   ids         = ["1978593525779****"]
        ///   type        = "A"
        ///   output_file = "records.txt"
        /// }
        /// 
        /// output "first_record_id" {
        ///   value = "${data.alicloud_alidns_records.records_ds.records.0.record_id}"
        /// }
        /// ```
        /// </summary>
        public static Task<GetAlidnsRecordsResult> InvokeAsync(GetAlidnsRecordsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAlidnsRecordsResult>("alicloud:dns/getAlidnsRecords:getAlidnsRecords", args ?? new GetAlidnsRecordsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides a list of Alidns Domain Records in an Alibaba Cloud account according to the specified filters.
        /// 
        /// &gt; **NOTE:**  Available in 1.86.0+.
        /// 
        /// ## Example Usage
        /// 
        /// ```terraform 
        /// data "alicloud_alidns_records" "records_ds" {
        ///   domain_name = "xiaozhu.top"
        ///   ids         = ["1978593525779****"]
        ///   type        = "A"
        ///   output_file = "records.txt"
        /// }
        /// 
        /// output "first_record_id" {
        ///   value = "${data.alicloud_alidns_records.records_ds.records.0.record_id}"
        /// }
        /// ```
        /// </summary>
        public static Output<GetAlidnsRecordsResult> Invoke(GetAlidnsRecordsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAlidnsRecordsResult>("alicloud:dns/getAlidnsRecords:getAlidnsRecords", args ?? new GetAlidnsRecordsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAlidnsRecordsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Sorting direction. Valid values: `DESC`,`ASC`. Default to `AESC`.
        /// </summary>
        [Input("direction")]
        public string? Direction { get; set; }

        /// <summary>
        /// The domain name associated to the records.
        /// </summary>
        [Input("domainName", required: true)]
        public string DomainName { get; set; } = null!;

        /// <summary>
        /// Domain name group ID.
        /// </summary>
        [Input("groupId")]
        public int? GroupId { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of record IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// Keywords.
        /// </summary>
        [Input("keyWord")]
        public string? KeyWord { get; set; }

        /// <summary>
        /// User language.
        /// </summary>
        [Input("lang")]
        public string? Lang { get; set; }

        /// <summary>
        /// ISP line. For checking all resolution lines enumeration please visit [Alibaba Cloud DNS doc](https://www.alibabacloud.com/help/en/alibaba-cloud-dns/latest/dns-lines)
        /// </summary>
        [Input("line")]
        public string? Line { get; set; }

        /// <summary>
        /// Sort by. Sort from newest to oldest according to the time added by resolution.
        /// </summary>
        [Input("orderBy")]
        public string? OrderBy { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The keywords recorded by the host are searched according to the `%!R(MISSING)RKeyWord%!`(MISSING) mode, and are not case sensitive.
        /// </summary>
        [Input("rrKeyWord")]
        public string? RrKeyWord { get; set; }

        /// <summary>
        /// Host record regex.
        /// </summary>
        [Input("rrRegex")]
        public string? RrRegex { get; set; }

        /// <summary>
        /// Search mode, Valid values: `LIKE`, `EXACT`, `ADVANCED`, `LIKE` (fuzzy), `EXACT` (accurate) search supports KeyWord field, `ADVANCED` (advanced) mode supports other fields.
        /// </summary>
        [Input("searchMode")]
        public string? SearchMode { get; set; }

        /// <summary>
        /// Record status. Valid values: `ENABLE` and `DISABLE`.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        /// <summary>
        /// Record type. Valid values: `A`, `NS`, `MX`, `TXT`, `CNAME`, `SRV`, `AAAA`, `REDIRECT_URL`, `FORWORD_URL` .
        /// </summary>
        [Input("type")]
        public string? Type { get; set; }

        /// <summary>
        /// Analyze type keywords, search by full match, not case sensitive.
        /// </summary>
        [Input("typeKeyWord")]
        public string? TypeKeyWord { get; set; }

        /// <summary>
        /// The keywords of the recorded value are searched according to the `%!V(MISSING)alueKeyWord%!`(MISSING) mode, and are not case sensitive.
        /// </summary>
        [Input("valueKeyWord")]
        public string? ValueKeyWord { get; set; }

        /// <summary>
        /// Host record value regex.
        /// </summary>
        [Input("valueRegex")]
        public string? ValueRegex { get; set; }

        public GetAlidnsRecordsArgs()
        {
        }
        public static new GetAlidnsRecordsArgs Empty => new GetAlidnsRecordsArgs();
    }

    public sealed class GetAlidnsRecordsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Sorting direction. Valid values: `DESC`,`ASC`. Default to `AESC`.
        /// </summary>
        [Input("direction")]
        public Input<string>? Direction { get; set; }

        /// <summary>
        /// The domain name associated to the records.
        /// </summary>
        [Input("domainName", required: true)]
        public Input<string> DomainName { get; set; } = null!;

        /// <summary>
        /// Domain name group ID.
        /// </summary>
        [Input("groupId")]
        public Input<int>? GroupId { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of record IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// Keywords.
        /// </summary>
        [Input("keyWord")]
        public Input<string>? KeyWord { get; set; }

        /// <summary>
        /// User language.
        /// </summary>
        [Input("lang")]
        public Input<string>? Lang { get; set; }

        /// <summary>
        /// ISP line. For checking all resolution lines enumeration please visit [Alibaba Cloud DNS doc](https://www.alibabacloud.com/help/en/alibaba-cloud-dns/latest/dns-lines)
        /// </summary>
        [Input("line")]
        public Input<string>? Line { get; set; }

        /// <summary>
        /// Sort by. Sort from newest to oldest according to the time added by resolution.
        /// </summary>
        [Input("orderBy")]
        public Input<string>? OrderBy { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The keywords recorded by the host are searched according to the `%!R(MISSING)RKeyWord%!`(MISSING) mode, and are not case sensitive.
        /// </summary>
        [Input("rrKeyWord")]
        public Input<string>? RrKeyWord { get; set; }

        /// <summary>
        /// Host record regex.
        /// </summary>
        [Input("rrRegex")]
        public Input<string>? RrRegex { get; set; }

        /// <summary>
        /// Search mode, Valid values: `LIKE`, `EXACT`, `ADVANCED`, `LIKE` (fuzzy), `EXACT` (accurate) search supports KeyWord field, `ADVANCED` (advanced) mode supports other fields.
        /// </summary>
        [Input("searchMode")]
        public Input<string>? SearchMode { get; set; }

        /// <summary>
        /// Record status. Valid values: `ENABLE` and `DISABLE`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Record type. Valid values: `A`, `NS`, `MX`, `TXT`, `CNAME`, `SRV`, `AAAA`, `REDIRECT_URL`, `FORWORD_URL` .
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Analyze type keywords, search by full match, not case sensitive.
        /// </summary>
        [Input("typeKeyWord")]
        public Input<string>? TypeKeyWord { get; set; }

        /// <summary>
        /// The keywords of the recorded value are searched according to the `%!V(MISSING)alueKeyWord%!`(MISSING) mode, and are not case sensitive.
        /// </summary>
        [Input("valueKeyWord")]
        public Input<string>? ValueKeyWord { get; set; }

        /// <summary>
        /// Host record value regex.
        /// </summary>
        [Input("valueRegex")]
        public Input<string>? ValueRegex { get; set; }

        public GetAlidnsRecordsInvokeArgs()
        {
        }
        public static new GetAlidnsRecordsInvokeArgs Empty => new GetAlidnsRecordsInvokeArgs();
    }


    [OutputType]
    public sealed class GetAlidnsRecordsResult
    {
        public readonly string? Direction;
        /// <summary>
        /// Name of the domain record belongs to.
        /// </summary>
        public readonly string DomainName;
        public readonly int? GroupId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of record IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? KeyWord;
        public readonly string? Lang;
        /// <summary>
        /// ISP line of the record.
        /// </summary>
        public readonly string? Line;
        public readonly string? OrderBy;
        public readonly string? OutputFile;
        /// <summary>
        /// A list of records. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAlidnsRecordsRecordResult> Records;
        public readonly string? RrKeyWord;
        public readonly string? RrRegex;
        public readonly string? SearchMode;
        /// <summary>
        /// Status of the record.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// Type of the record.
        /// </summary>
        public readonly string? Type;
        public readonly string? TypeKeyWord;
        public readonly string? ValueKeyWord;
        public readonly string? ValueRegex;

        [OutputConstructor]
        private GetAlidnsRecordsResult(
            string? direction,

            string domainName,

            int? groupId,

            string id,

            ImmutableArray<string> ids,

            string? keyWord,

            string? lang,

            string? line,

            string? orderBy,

            string? outputFile,

            ImmutableArray<Outputs.GetAlidnsRecordsRecordResult> records,

            string? rrKeyWord,

            string? rrRegex,

            string? searchMode,

            string? status,

            string? type,

            string? typeKeyWord,

            string? valueKeyWord,

            string? valueRegex)
        {
            Direction = direction;
            DomainName = domainName;
            GroupId = groupId;
            Id = id;
            Ids = ids;
            KeyWord = keyWord;
            Lang = lang;
            Line = line;
            OrderBy = orderBy;
            OutputFile = outputFile;
            Records = records;
            RrKeyWord = rrKeyWord;
            RrRegex = rrRegex;
            SearchMode = searchMode;
            Status = status;
            Type = type;
            TypeKeyWord = typeKeyWord;
            ValueKeyWord = valueKeyWord;
            ValueRegex = valueRegex;
        }
    }
}

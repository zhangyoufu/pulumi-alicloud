// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Quotas
{
    [AliCloudResourceType("alicloud:quotas/applicationInfo:ApplicationInfo")]
    public partial class ApplicationInfo : global::Pulumi.CustomResource
    {
        [Output("approveValue")]
        public Output<string> ApproveValue { get; private set; } = null!;

        [Output("auditMode")]
        public Output<string> AuditMode { get; private set; } = null!;

        [Output("auditReason")]
        public Output<string> AuditReason { get; private set; } = null!;

        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        [Output("desireValue")]
        public Output<double> DesireValue { get; private set; } = null!;

        [Output("dimensions")]
        public Output<ImmutableArray<Outputs.ApplicationInfoDimension>> Dimensions { get; private set; } = null!;

        [Output("effectiveTime")]
        public Output<string?> EffectiveTime { get; private set; } = null!;

        [Output("envLanguage")]
        public Output<string?> EnvLanguage { get; private set; } = null!;

        [Output("expireTime")]
        public Output<string?> ExpireTime { get; private set; } = null!;

        [Output("noticeType")]
        public Output<int> NoticeType { get; private set; } = null!;

        [Output("productCode")]
        public Output<string> ProductCode { get; private set; } = null!;

        [Output("quotaActionCode")]
        public Output<string> QuotaActionCode { get; private set; } = null!;

        [Output("quotaCategory")]
        public Output<string?> QuotaCategory { get; private set; } = null!;

        [Output("quotaDescription")]
        public Output<string> QuotaDescription { get; private set; } = null!;

        [Output("quotaName")]
        public Output<string> QuotaName { get; private set; } = null!;

        [Output("quotaUnit")]
        public Output<string> QuotaUnit { get; private set; } = null!;

        [Output("reason")]
        public Output<string> Reason { get; private set; } = null!;

        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a ApplicationInfo resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApplicationInfo(string name, ApplicationInfoArgs args, CustomResourceOptions? options = null)
            : base("alicloud:quotas/applicationInfo:ApplicationInfo", name, args ?? new ApplicationInfoArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApplicationInfo(string name, Input<string> id, ApplicationInfoState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:quotas/applicationInfo:ApplicationInfo", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ApplicationInfo resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApplicationInfo Get(string name, Input<string> id, ApplicationInfoState? state = null, CustomResourceOptions? options = null)
        {
            return new ApplicationInfo(name, id, state, options);
        }
    }

    public sealed class ApplicationInfoArgs : global::Pulumi.ResourceArgs
    {
        [Input("auditMode")]
        public Input<string>? AuditMode { get; set; }

        [Input("desireValue", required: true)]
        public Input<double> DesireValue { get; set; } = null!;

        [Input("dimensions")]
        private InputList<Inputs.ApplicationInfoDimensionArgs>? _dimensions;
        public InputList<Inputs.ApplicationInfoDimensionArgs> Dimensions
        {
            get => _dimensions ?? (_dimensions = new InputList<Inputs.ApplicationInfoDimensionArgs>());
            set => _dimensions = value;
        }

        [Input("effectiveTime")]
        public Input<string>? EffectiveTime { get; set; }

        [Input("envLanguage")]
        public Input<string>? EnvLanguage { get; set; }

        [Input("expireTime")]
        public Input<string>? ExpireTime { get; set; }

        [Input("noticeType")]
        public Input<int>? NoticeType { get; set; }

        [Input("productCode", required: true)]
        public Input<string> ProductCode { get; set; } = null!;

        [Input("quotaActionCode", required: true)]
        public Input<string> QuotaActionCode { get; set; } = null!;

        [Input("quotaCategory")]
        public Input<string>? QuotaCategory { get; set; }

        [Input("reason", required: true)]
        public Input<string> Reason { get; set; } = null!;

        public ApplicationInfoArgs()
        {
        }
        public static new ApplicationInfoArgs Empty => new ApplicationInfoArgs();
    }

    public sealed class ApplicationInfoState : global::Pulumi.ResourceArgs
    {
        [Input("approveValue")]
        public Input<string>? ApproveValue { get; set; }

        [Input("auditMode")]
        public Input<string>? AuditMode { get; set; }

        [Input("auditReason")]
        public Input<string>? AuditReason { get; set; }

        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        [Input("desireValue")]
        public Input<double>? DesireValue { get; set; }

        [Input("dimensions")]
        private InputList<Inputs.ApplicationInfoDimensionGetArgs>? _dimensions;
        public InputList<Inputs.ApplicationInfoDimensionGetArgs> Dimensions
        {
            get => _dimensions ?? (_dimensions = new InputList<Inputs.ApplicationInfoDimensionGetArgs>());
            set => _dimensions = value;
        }

        [Input("effectiveTime")]
        public Input<string>? EffectiveTime { get; set; }

        [Input("envLanguage")]
        public Input<string>? EnvLanguage { get; set; }

        [Input("expireTime")]
        public Input<string>? ExpireTime { get; set; }

        [Input("noticeType")]
        public Input<int>? NoticeType { get; set; }

        [Input("productCode")]
        public Input<string>? ProductCode { get; set; }

        [Input("quotaActionCode")]
        public Input<string>? QuotaActionCode { get; set; }

        [Input("quotaCategory")]
        public Input<string>? QuotaCategory { get; set; }

        [Input("quotaDescription")]
        public Input<string>? QuotaDescription { get; set; }

        [Input("quotaName")]
        public Input<string>? QuotaName { get; set; }

        [Input("quotaUnit")]
        public Input<string>? QuotaUnit { get; set; }

        [Input("reason")]
        public Input<string>? Reason { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        public ApplicationInfoState()
        {
        }
        public static new ApplicationInfoState Empty => new ApplicationInfoState();
    }
}

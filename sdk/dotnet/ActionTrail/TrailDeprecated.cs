// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ActionTrail
{
    /// <summary>
    /// &gt; **DEPRECATED:**  This resource has been renamed to alicloud.actiontrail.Trail from version 1.95.0.
    /// 
    /// Provides a new resource to manage [Action Trail](https://www.alibabacloud.com/help/doc-detail/28804.htm).
    /// 
    /// &gt; **NOTE:** Available in 1.35.0+
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
    ///     // Create a new action trail.
    ///     var foo = new AliCloud.ActionTrail.TrailDeprecated("foo", new()
    ///     {
    ///         EventRw = "Write-test",
    ///         OssBucketName = alicloud_oss_bucket.Bucket.Id,
    ///         RoleName = alicloud_ram_role_policy_attachment.Attach.Role_name,
    ///         OssKeyPrefix = "at-product-account-audit-B",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Action trail can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:actiontrail/trailDeprecated:TrailDeprecated foo abc12345678
    /// ```
    /// </summary>
    [Obsolete(@"Resource renamed to `Trail`")]
    [AliCloudResourceType("alicloud:actiontrail/trailDeprecated:TrailDeprecated")]
    public partial class TrailDeprecated : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Indicates whether the event is a read or a write event. Valid values: Read, Write, and All. Default value: Write.
        /// </summary>
        [Output("eventRw")]
        public Output<string?> EventRw { get; private set; } = null!;

        [Output("isOrganizationTrail")]
        public Output<bool?> IsOrganizationTrail { get; private set; } = null!;

        [Output("mnsTopicArn")]
        public Output<string?> MnsTopicArn { get; private set; } = null!;

        /// <summary>
        /// The name of the trail to be created, which must be unique for an account.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The OSS bucket to which the trail delivers logs. Ensure that this is an existing OSS bucket.
        /// </summary>
        [Output("ossBucketName")]
        public Output<string?> OssBucketName { get; private set; } = null!;

        /// <summary>
        /// The prefix of the specified OSS bucket name. This parameter can be left empty.
        /// </summary>
        [Output("ossKeyPrefix")]
        public Output<string?> OssKeyPrefix { get; private set; } = null!;

        [Output("ossWriteRoleArn")]
        public Output<string?> OssWriteRoleArn { get; private set; } = null!;

        /// <summary>
        /// The RAM role in ActionTrail permitted by the user.
        /// </summary>
        [Output("roleName")]
        public Output<string> RoleName { get; private set; } = null!;

        /// <summary>
        /// The unique ARN of the Log Service project.
        /// </summary>
        [Output("slsProjectArn")]
        public Output<string?> SlsProjectArn { get; private set; } = null!;

        /// <summary>
        /// The unique ARN of the Log Service role.
        /// 
        /// &gt; **NOTE:** `sls_project_arn` and `sls_write_role_arn` should be set or not set at the same time when actiontrail delivers logs.
        /// </summary>
        [Output("slsWriteRoleArn")]
        public Output<string> SlsWriteRoleArn { get; private set; } = null!;

        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        [Output("trailName")]
        public Output<string> TrailName { get; private set; } = null!;

        [Output("trailRegion")]
        public Output<string> TrailRegion { get; private set; } = null!;


        /// <summary>
        /// Create a TrailDeprecated resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TrailDeprecated(string name, TrailDeprecatedArgs? args = null, CustomResourceOptions? options = null)
            : base("alicloud:actiontrail/trailDeprecated:TrailDeprecated", name, args ?? new TrailDeprecatedArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TrailDeprecated(string name, Input<string> id, TrailDeprecatedState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:actiontrail/trailDeprecated:TrailDeprecated", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TrailDeprecated resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TrailDeprecated Get(string name, Input<string> id, TrailDeprecatedState? state = null, CustomResourceOptions? options = null)
        {
            return new TrailDeprecated(name, id, state, options);
        }
    }

    public sealed class TrailDeprecatedArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether the event is a read or a write event. Valid values: Read, Write, and All. Default value: Write.
        /// </summary>
        [Input("eventRw")]
        public Input<string>? EventRw { get; set; }

        [Input("isOrganizationTrail")]
        public Input<bool>? IsOrganizationTrail { get; set; }

        [Input("mnsTopicArn")]
        public Input<string>? MnsTopicArn { get; set; }

        /// <summary>
        /// The name of the trail to be created, which must be unique for an account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The OSS bucket to which the trail delivers logs. Ensure that this is an existing OSS bucket.
        /// </summary>
        [Input("ossBucketName")]
        public Input<string>? OssBucketName { get; set; }

        /// <summary>
        /// The prefix of the specified OSS bucket name. This parameter can be left empty.
        /// </summary>
        [Input("ossKeyPrefix")]
        public Input<string>? OssKeyPrefix { get; set; }

        [Input("ossWriteRoleArn")]
        public Input<string>? OssWriteRoleArn { get; set; }

        /// <summary>
        /// The RAM role in ActionTrail permitted by the user.
        /// </summary>
        [Input("roleName")]
        public Input<string>? RoleName { get; set; }

        /// <summary>
        /// The unique ARN of the Log Service project.
        /// </summary>
        [Input("slsProjectArn")]
        public Input<string>? SlsProjectArn { get; set; }

        /// <summary>
        /// The unique ARN of the Log Service role.
        /// 
        /// &gt; **NOTE:** `sls_project_arn` and `sls_write_role_arn` should be set or not set at the same time when actiontrail delivers logs.
        /// </summary>
        [Input("slsWriteRoleArn")]
        public Input<string>? SlsWriteRoleArn { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("trailName")]
        public Input<string>? TrailName { get; set; }

        [Input("trailRegion")]
        public Input<string>? TrailRegion { get; set; }

        public TrailDeprecatedArgs()
        {
        }
        public static new TrailDeprecatedArgs Empty => new TrailDeprecatedArgs();
    }

    public sealed class TrailDeprecatedState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether the event is a read or a write event. Valid values: Read, Write, and All. Default value: Write.
        /// </summary>
        [Input("eventRw")]
        public Input<string>? EventRw { get; set; }

        [Input("isOrganizationTrail")]
        public Input<bool>? IsOrganizationTrail { get; set; }

        [Input("mnsTopicArn")]
        public Input<string>? MnsTopicArn { get; set; }

        /// <summary>
        /// The name of the trail to be created, which must be unique for an account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The OSS bucket to which the trail delivers logs. Ensure that this is an existing OSS bucket.
        /// </summary>
        [Input("ossBucketName")]
        public Input<string>? OssBucketName { get; set; }

        /// <summary>
        /// The prefix of the specified OSS bucket name. This parameter can be left empty.
        /// </summary>
        [Input("ossKeyPrefix")]
        public Input<string>? OssKeyPrefix { get; set; }

        [Input("ossWriteRoleArn")]
        public Input<string>? OssWriteRoleArn { get; set; }

        /// <summary>
        /// The RAM role in ActionTrail permitted by the user.
        /// </summary>
        [Input("roleName")]
        public Input<string>? RoleName { get; set; }

        /// <summary>
        /// The unique ARN of the Log Service project.
        /// </summary>
        [Input("slsProjectArn")]
        public Input<string>? SlsProjectArn { get; set; }

        /// <summary>
        /// The unique ARN of the Log Service role.
        /// 
        /// &gt; **NOTE:** `sls_project_arn` and `sls_write_role_arn` should be set or not set at the same time when actiontrail delivers logs.
        /// </summary>
        [Input("slsWriteRoleArn")]
        public Input<string>? SlsWriteRoleArn { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("trailName")]
        public Input<string>? TrailName { get; set; }

        [Input("trailRegion")]
        public Input<string>? TrailRegion { get; set; }

        public TrailDeprecatedState()
        {
        }
        public static new TrailDeprecatedState Empty => new TrailDeprecatedState();
    }
}

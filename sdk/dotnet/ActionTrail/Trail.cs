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
    /// Provides a ActionTrail Trail resource. For information about alicloud actiontrail trail and how to use it, see [What is Resource Alicloud ActionTrail Trail](https://www.alibabacloud.com/help/en/actiontrail/latest/api-actiontrail-2020-07-06-createtrail).
    /// 
    /// &gt; **NOTE:** Available since v1.95.0.
    /// 
    /// &gt; **NOTE:** You can create a trail to deliver events to Log Service, Object Storage Service (OSS), or both. Before you call this operation to create a trail, make sure that the following requirements are met.
    /// - Deliver events to Log Service: A project is created in Log Service.
    /// - Deliver events to OSS: A bucket is created in OSS.
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
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "tf-example";
    ///     var exampleRegions = AliCloud.GetRegions.Invoke(new()
    ///     {
    ///         Current = true,
    ///     });
    /// 
    ///     var exampleAccount = AliCloud.GetAccount.Invoke();
    /// 
    ///     var exampleProject = new AliCloud.Log.Project("exampleProject", new()
    ///     {
    ///         Description = "tf actiontrail example",
    ///     });
    /// 
    ///     var exampleRoles = AliCloud.Ram.GetRoles.Invoke(new()
    ///     {
    ///         NameRegex = "AliyunServiceRoleForActionTrail",
    ///     });
    /// 
    ///     var exampleTrail = new AliCloud.ActionTrail.Trail("exampleTrail", new()
    ///     {
    ///         TrailName = name,
    ///         SlsWriteRoleArn = exampleRoles.Apply(getRolesResult =&gt; getRolesResult.Roles[0]?.Arn),
    ///         SlsProjectArn = Output.Tuple(exampleRegions, exampleAccount, exampleProject.Name).Apply(values =&gt;
    ///         {
    ///             var exampleRegions = values.Item1;
    ///             var exampleAccount = values.Item2;
    ///             var name = values.Item3;
    ///             return $"acs:log:{exampleRegions.Apply(getRegionsResult =&gt; getRegionsResult.Regions[0]?.Id)}:{exampleAccount.Apply(getAccountResult =&gt; getAccountResult.Id)}:project/{name}";
    ///         }),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Action trail can be imported using the id or trail_name, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:actiontrail/trail:Trail default abc12345678
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:actiontrail/trail:Trail")]
    public partial class Trail : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Indicates whether the event is a read or a write event. Valid values: `Read`, `Write`, and `All`. Default to `Write`.
        /// </summary>
        [Output("eventRw")]
        public Output<string?> EventRw { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to create a multi-account trail. Valid values:`true`: Create a multi-account trail.`false`: Create a single-account trail. It is the default value.
        /// </summary>
        [Output("isOrganizationTrail")]
        public Output<bool?> IsOrganizationTrail { get; private set; } = null!;

        /// <summary>
        /// Field `mns_topic_arn` has been deprecated from version 1.118.0.
        /// </summary>
        [Output("mnsTopicArn")]
        public Output<string?> MnsTopicArn { get; private set; } = null!;

        /// <summary>
        /// Field `name` has been deprecated from version 1.95.0. Use `trail_name` instead.
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

        /// <summary>
        /// The unique ARN of the Oss role.
        /// </summary>
        [Output("ossWriteRoleArn")]
        public Output<string?> OssWriteRoleArn { get; private set; } = null!;

        /// <summary>
        /// Field `name` has been deprecated from version 1.118.0.
        /// </summary>
        [Output("roleName")]
        public Output<string> RoleName { get; private set; } = null!;

        /// <summary>
        /// The unique ARN of the Log Service project. Ensure that `sls_project_arn` is valid .
        /// </summary>
        [Output("slsProjectArn")]
        public Output<string?> SlsProjectArn { get; private set; } = null!;

        /// <summary>
        /// The unique ARN of the Log Service role.
        /// </summary>
        [Output("slsWriteRoleArn")]
        public Output<string> SlsWriteRoleArn { get; private set; } = null!;

        /// <summary>
        /// The status of ActionTrail Trail. After creation, tracking is turned on by default, and you can set the status value to `Disable` to turn off tracking. Valid values: `Enable`, `Disable`. Default to `Enable`.
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        /// <summary>
        /// The name of the trail to be created, which must be unique for an account.
        /// </summary>
        [Output("trailName")]
        public Output<string> TrailName { get; private set; } = null!;

        /// <summary>
        /// The regions to which the trail is applied. Default to `All`.
        /// </summary>
        [Output("trailRegion")]
        public Output<string> TrailRegion { get; private set; } = null!;


        /// <summary>
        /// Create a Trail resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Trail(string name, TrailArgs? args = null, CustomResourceOptions? options = null)
            : base("alicloud:actiontrail/trail:Trail", name, args ?? new TrailArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Trail(string name, Input<string> id, TrailState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:actiontrail/trail:Trail", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Trail resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Trail Get(string name, Input<string> id, TrailState? state = null, CustomResourceOptions? options = null)
        {
            return new Trail(name, id, state, options);
        }
    }

    public sealed class TrailArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether the event is a read or a write event. Valid values: `Read`, `Write`, and `All`. Default to `Write`.
        /// </summary>
        [Input("eventRw")]
        public Input<string>? EventRw { get; set; }

        /// <summary>
        /// Specifies whether to create a multi-account trail. Valid values:`true`: Create a multi-account trail.`false`: Create a single-account trail. It is the default value.
        /// </summary>
        [Input("isOrganizationTrail")]
        public Input<bool>? IsOrganizationTrail { get; set; }

        /// <summary>
        /// Field `mns_topic_arn` has been deprecated from version 1.118.0.
        /// </summary>
        [Input("mnsTopicArn")]
        public Input<string>? MnsTopicArn { get; set; }

        /// <summary>
        /// Field `name` has been deprecated from version 1.95.0. Use `trail_name` instead.
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

        /// <summary>
        /// The unique ARN of the Oss role.
        /// </summary>
        [Input("ossWriteRoleArn")]
        public Input<string>? OssWriteRoleArn { get; set; }

        /// <summary>
        /// Field `name` has been deprecated from version 1.118.0.
        /// </summary>
        [Input("roleName")]
        public Input<string>? RoleName { get; set; }

        /// <summary>
        /// The unique ARN of the Log Service project. Ensure that `sls_project_arn` is valid .
        /// </summary>
        [Input("slsProjectArn")]
        public Input<string>? SlsProjectArn { get; set; }

        /// <summary>
        /// The unique ARN of the Log Service role.
        /// </summary>
        [Input("slsWriteRoleArn")]
        public Input<string>? SlsWriteRoleArn { get; set; }

        /// <summary>
        /// The status of ActionTrail Trail. After creation, tracking is turned on by default, and you can set the status value to `Disable` to turn off tracking. Valid values: `Enable`, `Disable`. Default to `Enable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The name of the trail to be created, which must be unique for an account.
        /// </summary>
        [Input("trailName")]
        public Input<string>? TrailName { get; set; }

        /// <summary>
        /// The regions to which the trail is applied. Default to `All`.
        /// </summary>
        [Input("trailRegion")]
        public Input<string>? TrailRegion { get; set; }

        public TrailArgs()
        {
        }
        public static new TrailArgs Empty => new TrailArgs();
    }

    public sealed class TrailState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether the event is a read or a write event. Valid values: `Read`, `Write`, and `All`. Default to `Write`.
        /// </summary>
        [Input("eventRw")]
        public Input<string>? EventRw { get; set; }

        /// <summary>
        /// Specifies whether to create a multi-account trail. Valid values:`true`: Create a multi-account trail.`false`: Create a single-account trail. It is the default value.
        /// </summary>
        [Input("isOrganizationTrail")]
        public Input<bool>? IsOrganizationTrail { get; set; }

        /// <summary>
        /// Field `mns_topic_arn` has been deprecated from version 1.118.0.
        /// </summary>
        [Input("mnsTopicArn")]
        public Input<string>? MnsTopicArn { get; set; }

        /// <summary>
        /// Field `name` has been deprecated from version 1.95.0. Use `trail_name` instead.
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

        /// <summary>
        /// The unique ARN of the Oss role.
        /// </summary>
        [Input("ossWriteRoleArn")]
        public Input<string>? OssWriteRoleArn { get; set; }

        /// <summary>
        /// Field `name` has been deprecated from version 1.118.0.
        /// </summary>
        [Input("roleName")]
        public Input<string>? RoleName { get; set; }

        /// <summary>
        /// The unique ARN of the Log Service project. Ensure that `sls_project_arn` is valid .
        /// </summary>
        [Input("slsProjectArn")]
        public Input<string>? SlsProjectArn { get; set; }

        /// <summary>
        /// The unique ARN of the Log Service role.
        /// </summary>
        [Input("slsWriteRoleArn")]
        public Input<string>? SlsWriteRoleArn { get; set; }

        /// <summary>
        /// The status of ActionTrail Trail. After creation, tracking is turned on by default, and you can set the status value to `Disable` to turn off tracking. Valid values: `Enable`, `Disable`. Default to `Enable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The name of the trail to be created, which must be unique for an account.
        /// </summary>
        [Input("trailName")]
        public Input<string>? TrailName { get; set; }

        /// <summary>
        /// The regions to which the trail is applied. Default to `All`.
        /// </summary>
        [Input("trailRegion")]
        public Input<string>? TrailRegion { get; set; }

        public TrailState()
        {
        }
        public static new TrailState Empty => new TrailState();
    }
}

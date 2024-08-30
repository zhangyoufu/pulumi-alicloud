// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Oos
{
    /// <summary>
    /// Provides a OOS Service Setting resource.
    /// 
    /// For information about OOS Service Setting and how to use it, see [What is Service Setting](https://www.alibabacloud.com/help/en/doc-detail/268700.html).
    /// 
    /// &gt; **NOTE:** Available in v1.147.0+.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
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
    ///     var name = config.Get("name") ?? "tf-testaccoossetting";
    ///     var @default = new AliCloud.Oss.Bucket("default", new()
    ///     {
    ///         BucketName = name,
    ///     });
    /// 
    ///     var defaultBucketAcl = new AliCloud.Oss.BucketAcl("default", new()
    ///     {
    ///         Bucket = @default.BucketName,
    ///         Acl = "public-read-write",
    ///     });
    /// 
    ///     var defaultProject = new AliCloud.Log.Project("default", new()
    ///     {
    ///         ProjectName = name,
    ///     });
    /// 
    ///     var defaultServiceSetting = new AliCloud.Oos.ServiceSetting("default", new()
    ///     {
    ///         DeliveryOssEnabled = true,
    ///         DeliveryOssKeyPrefix = "path1/",
    ///         DeliveryOssBucketName = @default.BucketName,
    ///         DeliverySlsEnabled = true,
    ///         DeliverySlsProjectName = defaultProject.ProjectName,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// OOS Service Setting can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:oos/serviceSetting:ServiceSetting example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:oos/serviceSetting:ServiceSetting")]
    public partial class ServiceSetting : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the OSS bucket. **NOTE:** When the `delivery_oss_enabled` is `true`, The `delivery_oss_bucket_name` is valid.
        /// </summary>
        [Output("deliveryOssBucketName")]
        public Output<string?> DeliveryOssBucketName { get; private set; } = null!;

        /// <summary>
        /// Is the recording function for the OSS delivery template enabled.
        /// </summary>
        [Output("deliveryOssEnabled")]
        public Output<bool?> DeliveryOssEnabled { get; private set; } = null!;

        /// <summary>
        /// The Directory of the OSS bucket. **NOTE:** When the `delivery_oss_enabled` is `true`, The `delivery_oss_bucket_name` is valid.
        /// </summary>
        [Output("deliveryOssKeyPrefix")]
        public Output<string?> DeliveryOssKeyPrefix { get; private set; } = null!;

        /// <summary>
        /// Is the execution record function to SLS delivery Template turned on.
        /// </summary>
        [Output("deliverySlsEnabled")]
        public Output<bool?> DeliverySlsEnabled { get; private set; } = null!;

        /// <summary>
        /// The name of SLS  Project. **NOTE:** When the `delivery_sls_enabled` is `true`, The `delivery_sls_project_name` is valid.
        /// </summary>
        [Output("deliverySlsProjectName")]
        public Output<string?> DeliverySlsProjectName { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceSetting resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceSetting(string name, ServiceSettingArgs? args = null, CustomResourceOptions? options = null)
            : base("alicloud:oos/serviceSetting:ServiceSetting", name, args ?? new ServiceSettingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceSetting(string name, Input<string> id, ServiceSettingState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:oos/serviceSetting:ServiceSetting", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceSetting resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceSetting Get(string name, Input<string> id, ServiceSettingState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceSetting(name, id, state, options);
        }
    }

    public sealed class ServiceSettingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the OSS bucket. **NOTE:** When the `delivery_oss_enabled` is `true`, The `delivery_oss_bucket_name` is valid.
        /// </summary>
        [Input("deliveryOssBucketName")]
        public Input<string>? DeliveryOssBucketName { get; set; }

        /// <summary>
        /// Is the recording function for the OSS delivery template enabled.
        /// </summary>
        [Input("deliveryOssEnabled")]
        public Input<bool>? DeliveryOssEnabled { get; set; }

        /// <summary>
        /// The Directory of the OSS bucket. **NOTE:** When the `delivery_oss_enabled` is `true`, The `delivery_oss_bucket_name` is valid.
        /// </summary>
        [Input("deliveryOssKeyPrefix")]
        public Input<string>? DeliveryOssKeyPrefix { get; set; }

        /// <summary>
        /// Is the execution record function to SLS delivery Template turned on.
        /// </summary>
        [Input("deliverySlsEnabled")]
        public Input<bool>? DeliverySlsEnabled { get; set; }

        /// <summary>
        /// The name of SLS  Project. **NOTE:** When the `delivery_sls_enabled` is `true`, The `delivery_sls_project_name` is valid.
        /// </summary>
        [Input("deliverySlsProjectName")]
        public Input<string>? DeliverySlsProjectName { get; set; }

        public ServiceSettingArgs()
        {
        }
        public static new ServiceSettingArgs Empty => new ServiceSettingArgs();
    }

    public sealed class ServiceSettingState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the OSS bucket. **NOTE:** When the `delivery_oss_enabled` is `true`, The `delivery_oss_bucket_name` is valid.
        /// </summary>
        [Input("deliveryOssBucketName")]
        public Input<string>? DeliveryOssBucketName { get; set; }

        /// <summary>
        /// Is the recording function for the OSS delivery template enabled.
        /// </summary>
        [Input("deliveryOssEnabled")]
        public Input<bool>? DeliveryOssEnabled { get; set; }

        /// <summary>
        /// The Directory of the OSS bucket. **NOTE:** When the `delivery_oss_enabled` is `true`, The `delivery_oss_bucket_name` is valid.
        /// </summary>
        [Input("deliveryOssKeyPrefix")]
        public Input<string>? DeliveryOssKeyPrefix { get; set; }

        /// <summary>
        /// Is the execution record function to SLS delivery Template turned on.
        /// </summary>
        [Input("deliverySlsEnabled")]
        public Input<bool>? DeliverySlsEnabled { get; set; }

        /// <summary>
        /// The name of SLS  Project. **NOTE:** When the `delivery_sls_enabled` is `true`, The `delivery_sls_project_name` is valid.
        /// </summary>
        [Input("deliverySlsProjectName")]
        public Input<string>? DeliverySlsProjectName { get; set; }

        public ServiceSettingState()
        {
        }
        public static new ServiceSettingState Empty => new ServiceSettingState();
    }
}

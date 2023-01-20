// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.FC
{
    /// <summary>
    /// ## Import
    /// 
    /// Function Compute Service can be imported using the id or name, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:fc/service:Service foo my-fc-service
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:fc/service:Service")]
    public partial class Service : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Function Compute Service description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Whether to allow the Service to access Internet. Default to "true".
        /// </summary>
        [Output("internetAccess")]
        public Output<bool?> InternetAccess { get; private set; } = null!;

        /// <summary>
        /// The date this resource was last modified.
        /// </summary>
        [Output("lastModified")]
        public Output<string> LastModified { get; private set; } = null!;

        /// <summary>
        /// Provide this to store your Function Compute Service logs. Fields documented below. See [Create a Service](https://www.alibabacloud.com/help/doc-detail/51924.htm). `log_config` requires the following: (**NOTE:** If both `project` and `logstore` are empty, log_config is considered to be empty or unset.)
        /// </summary>
        [Output("logConfig")]
        public Output<Outputs.ServiceLogConfig?> LogConfig { get; private set; } = null!;

        /// <summary>
        /// The Function Compute Service name. It is the only in one Alicloud account and is conflict with `name_prefix`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Setting a prefix to get a only name. It is conflict with `name`.
        /// </summary>
        [Output("namePrefix")]
        public Output<string?> NamePrefix { get; private set; } = null!;

        /// <summary>
        /// Provide [NAS configuration](https://www.alibabacloud.com/help/doc-detail/87401.htm) to allow Function Compute Service to access your NAS resources. `nas_config` requires the following:
        /// </summary>
        [Output("nasConfig")]
        public Output<Outputs.ServiceNasConfig?> NasConfig { get; private set; } = null!;

        /// <summary>
        /// Whether to publish creation/change as new Function Compute Service Version. Defaults to `false`.
        /// </summary>
        [Output("publish")]
        public Output<bool?> Publish { get; private set; } = null!;

        /// <summary>
        /// RAM role arn attached to the Function Compute Service. This governs both who / what can invoke your Function, as well as what resources our Function has access to. See [User Permissions](https://www.alibabacloud.com/help/doc-detail/52885.htm) for more details.
        /// </summary>
        [Output("role")]
        public Output<string?> Role { get; private set; } = null!;

        /// <summary>
        /// The Function Compute Service ID.
        /// </summary>
        [Output("serviceId")]
        public Output<string> ServiceId { get; private set; } = null!;

        /// <summary>
        /// Provide this to allow your Function Compute to report tracing information. Fields documented below. See [Function Compute Tracing Config](https://help.aliyun.com/document_detail/189805.html). `tracing_config` requires the following: (**NOTE:** If both `type` and `params` are empty, tracing_config is considered to be empty or unset.)
        /// </summary>
        [Output("tracingConfig")]
        public Output<Outputs.ServiceTracingConfig?> TracingConfig { get; private set; } = null!;

        /// <summary>
        /// The latest published version of your Function Compute Service.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;

        /// <summary>
        /// Provide this to allow your Function Compute Service to access your VPC. Fields documented below. See [Function Compute Service in VPC](https://www.alibabacloud.com/help/faq-detail/72959.htm). `vpc_config` requires the following: (**NOTE:** If both `vswitch_ids` and `security_group_id` are empty, vpc_config is considered to be empty or unset.)
        /// </summary>
        [Output("vpcConfig")]
        public Output<Outputs.ServiceVpcConfig?> VpcConfig { get; private set; } = null!;


        /// <summary>
        /// Create a Service resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Service(string name, ServiceArgs? args = null, CustomResourceOptions? options = null)
            : base("alicloud:fc/service:Service", name, args ?? new ServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Service(string name, Input<string> id, ServiceState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:fc/service:Service", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Service resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Service Get(string name, Input<string> id, ServiceState? state = null, CustomResourceOptions? options = null)
        {
            return new Service(name, id, state, options);
        }
    }

    public sealed class ServiceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Function Compute Service description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether to allow the Service to access Internet. Default to "true".
        /// </summary>
        [Input("internetAccess")]
        public Input<bool>? InternetAccess { get; set; }

        /// <summary>
        /// Provide this to store your Function Compute Service logs. Fields documented below. See [Create a Service](https://www.alibabacloud.com/help/doc-detail/51924.htm). `log_config` requires the following: (**NOTE:** If both `project` and `logstore` are empty, log_config is considered to be empty or unset.)
        /// </summary>
        [Input("logConfig")]
        public Input<Inputs.ServiceLogConfigArgs>? LogConfig { get; set; }

        /// <summary>
        /// The Function Compute Service name. It is the only in one Alicloud account and is conflict with `name_prefix`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Setting a prefix to get a only name. It is conflict with `name`.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// Provide [NAS configuration](https://www.alibabacloud.com/help/doc-detail/87401.htm) to allow Function Compute Service to access your NAS resources. `nas_config` requires the following:
        /// </summary>
        [Input("nasConfig")]
        public Input<Inputs.ServiceNasConfigArgs>? NasConfig { get; set; }

        /// <summary>
        /// Whether to publish creation/change as new Function Compute Service Version. Defaults to `false`.
        /// </summary>
        [Input("publish")]
        public Input<bool>? Publish { get; set; }

        /// <summary>
        /// RAM role arn attached to the Function Compute Service. This governs both who / what can invoke your Function, as well as what resources our Function has access to. See [User Permissions](https://www.alibabacloud.com/help/doc-detail/52885.htm) for more details.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// Provide this to allow your Function Compute to report tracing information. Fields documented below. See [Function Compute Tracing Config](https://help.aliyun.com/document_detail/189805.html). `tracing_config` requires the following: (**NOTE:** If both `type` and `params` are empty, tracing_config is considered to be empty or unset.)
        /// </summary>
        [Input("tracingConfig")]
        public Input<Inputs.ServiceTracingConfigArgs>? TracingConfig { get; set; }

        /// <summary>
        /// Provide this to allow your Function Compute Service to access your VPC. Fields documented below. See [Function Compute Service in VPC](https://www.alibabacloud.com/help/faq-detail/72959.htm). `vpc_config` requires the following: (**NOTE:** If both `vswitch_ids` and `security_group_id` are empty, vpc_config is considered to be empty or unset.)
        /// </summary>
        [Input("vpcConfig")]
        public Input<Inputs.ServiceVpcConfigArgs>? VpcConfig { get; set; }

        public ServiceArgs()
        {
        }
        public static new ServiceArgs Empty => new ServiceArgs();
    }

    public sealed class ServiceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Function Compute Service description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether to allow the Service to access Internet. Default to "true".
        /// </summary>
        [Input("internetAccess")]
        public Input<bool>? InternetAccess { get; set; }

        /// <summary>
        /// The date this resource was last modified.
        /// </summary>
        [Input("lastModified")]
        public Input<string>? LastModified { get; set; }

        /// <summary>
        /// Provide this to store your Function Compute Service logs. Fields documented below. See [Create a Service](https://www.alibabacloud.com/help/doc-detail/51924.htm). `log_config` requires the following: (**NOTE:** If both `project` and `logstore` are empty, log_config is considered to be empty or unset.)
        /// </summary>
        [Input("logConfig")]
        public Input<Inputs.ServiceLogConfigGetArgs>? LogConfig { get; set; }

        /// <summary>
        /// The Function Compute Service name. It is the only in one Alicloud account and is conflict with `name_prefix`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Setting a prefix to get a only name. It is conflict with `name`.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// Provide [NAS configuration](https://www.alibabacloud.com/help/doc-detail/87401.htm) to allow Function Compute Service to access your NAS resources. `nas_config` requires the following:
        /// </summary>
        [Input("nasConfig")]
        public Input<Inputs.ServiceNasConfigGetArgs>? NasConfig { get; set; }

        /// <summary>
        /// Whether to publish creation/change as new Function Compute Service Version. Defaults to `false`.
        /// </summary>
        [Input("publish")]
        public Input<bool>? Publish { get; set; }

        /// <summary>
        /// RAM role arn attached to the Function Compute Service. This governs both who / what can invoke your Function, as well as what resources our Function has access to. See [User Permissions](https://www.alibabacloud.com/help/doc-detail/52885.htm) for more details.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// The Function Compute Service ID.
        /// </summary>
        [Input("serviceId")]
        public Input<string>? ServiceId { get; set; }

        /// <summary>
        /// Provide this to allow your Function Compute to report tracing information. Fields documented below. See [Function Compute Tracing Config](https://help.aliyun.com/document_detail/189805.html). `tracing_config` requires the following: (**NOTE:** If both `type` and `params` are empty, tracing_config is considered to be empty or unset.)
        /// </summary>
        [Input("tracingConfig")]
        public Input<Inputs.ServiceTracingConfigGetArgs>? TracingConfig { get; set; }

        /// <summary>
        /// The latest published version of your Function Compute Service.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// Provide this to allow your Function Compute Service to access your VPC. Fields documented below. See [Function Compute Service in VPC](https://www.alibabacloud.com/help/faq-detail/72959.htm). `vpc_config` requires the following: (**NOTE:** If both `vswitch_ids` and `security_group_id` are empty, vpc_config is considered to be empty or unset.)
        /// </summary>
        [Input("vpcConfig")]
        public Input<Inputs.ServiceVpcConfigGetArgs>? VpcConfig { get; set; }

        public ServiceState()
        {
        }
        public static new ServiceState Empty => new ServiceState();
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.SimpleApplicationServer
{
    /// <summary>
    /// Provides a Simple Application Server Instance resource.
    /// 
    /// For information about Simple Application Server Instance and how to use it, see [What is Instance](https://www.alibabacloud.com/help/doc-detail/190440.htm).
    /// 
    /// &gt; **NOTE:** Available in v1.135.0+.
    /// 
    /// ## Import
    /// 
    /// Simple Application Server Instance can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:simpleapplicationserver/instance:Instance example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:simpleapplicationserver/instance:Instance")]
    public partial class Instance : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies whether to enable auto-renewal. Unit: months. Valid values: `true` and `false`.
        /// </summary>
        [Output("autoRenew")]
        public Output<bool?> AutoRenew { get; private set; } = null!;

        /// <summary>
        /// The auto renew period. Valid values: `1`,`3`, `6`, `12`, `24`, `36`. **NOTE:** The attribute `auto_renew` is valid when the attribute is `true`.
        /// </summary>
        [Output("autoRenewPeriod")]
        public Output<int?> AutoRenewPeriod { get; private set; } = null!;

        /// <summary>
        /// The size of the data disk. Unit: GB. Valid values: `0` to `16380`.
        /// </summary>
        [Output("dataDiskSize")]
        public Output<int?> DataDiskSize { get; private set; } = null!;

        /// <summary>
        /// The ID of the image.  You can use the `alicloud.simpleapplicationserver.getImages` to query the available images in the specified region. The value must be an integral multiple of 20.
        /// </summary>
        [Output("imageId")]
        public Output<string> ImageId { get; private set; } = null!;

        /// <summary>
        /// The name of the simple application server.
        /// </summary>
        [Output("instanceName")]
        public Output<string?> InstanceName { get; private set; } = null!;

        /// <summary>
        /// The password of the simple application server. The password must be 8 to 30 characters in length. It must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include: `( ) ~ ! @ # $ % ^ &amp; * - + = | { } [ ] : ; &lt; &gt; , . ? /`.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// The paymen type of the resource. Valid values: `Subscription`.
        /// </summary>
        [Output("paymentType")]
        public Output<string> PaymentType { get; private set; } = null!;

        /// <summary>
        /// The period. Unit: months. Valid values: `1`,`3`, `6`, `12`, `24`, `36`.
        /// </summary>
        [Output("period")]
        public Output<int> Period { get; private set; } = null!;

        /// <summary>
        /// The ID of the plan. You can use the `alicloud.simpleapplicationserver.getServerPlans`  to query all the plans provided by Simple Application Server in the specified region.
        /// </summary>
        [Output("planId")]
        public Output<string> PlanId { get; private set; } = null!;

        /// <summary>
        /// The status of the simple application server. Valid values: `Resetting`, `Running`, `Stopped`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("alicloud:simpleapplicationserver/instance:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:simpleapplicationserver/instance:Instance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Instance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Instance Get(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new Instance(name, id, state, options);
        }
    }

    public sealed class InstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether to enable auto-renewal. Unit: months. Valid values: `true` and `false`.
        /// </summary>
        [Input("autoRenew")]
        public Input<bool>? AutoRenew { get; set; }

        /// <summary>
        /// The auto renew period. Valid values: `1`,`3`, `6`, `12`, `24`, `36`. **NOTE:** The attribute `auto_renew` is valid when the attribute is `true`.
        /// </summary>
        [Input("autoRenewPeriod")]
        public Input<int>? AutoRenewPeriod { get; set; }

        /// <summary>
        /// The size of the data disk. Unit: GB. Valid values: `0` to `16380`.
        /// </summary>
        [Input("dataDiskSize")]
        public Input<int>? DataDiskSize { get; set; }

        /// <summary>
        /// The ID of the image.  You can use the `alicloud.simpleapplicationserver.getImages` to query the available images in the specified region. The value must be an integral multiple of 20.
        /// </summary>
        [Input("imageId", required: true)]
        public Input<string> ImageId { get; set; } = null!;

        /// <summary>
        /// The name of the simple application server.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        /// <summary>
        /// The password of the simple application server. The password must be 8 to 30 characters in length. It must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include: `( ) ~ ! @ # $ % ^ &amp; * - + = | { } [ ] : ; &lt; &gt; , . ? /`.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The paymen type of the resource. Valid values: `Subscription`.
        /// </summary>
        [Input("paymentType")]
        public Input<string>? PaymentType { get; set; }

        /// <summary>
        /// The period. Unit: months. Valid values: `1`,`3`, `6`, `12`, `24`, `36`.
        /// </summary>
        [Input("period", required: true)]
        public Input<int> Period { get; set; } = null!;

        /// <summary>
        /// The ID of the plan. You can use the `alicloud.simpleapplicationserver.getServerPlans`  to query all the plans provided by Simple Application Server in the specified region.
        /// </summary>
        [Input("planId", required: true)]
        public Input<string> PlanId { get; set; } = null!;

        /// <summary>
        /// The status of the simple application server. Valid values: `Resetting`, `Running`, `Stopped`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public InstanceArgs()
        {
        }
    }

    public sealed class InstanceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether to enable auto-renewal. Unit: months. Valid values: `true` and `false`.
        /// </summary>
        [Input("autoRenew")]
        public Input<bool>? AutoRenew { get; set; }

        /// <summary>
        /// The auto renew period. Valid values: `1`,`3`, `6`, `12`, `24`, `36`. **NOTE:** The attribute `auto_renew` is valid when the attribute is `true`.
        /// </summary>
        [Input("autoRenewPeriod")]
        public Input<int>? AutoRenewPeriod { get; set; }

        /// <summary>
        /// The size of the data disk. Unit: GB. Valid values: `0` to `16380`.
        /// </summary>
        [Input("dataDiskSize")]
        public Input<int>? DataDiskSize { get; set; }

        /// <summary>
        /// The ID of the image.  You can use the `alicloud.simpleapplicationserver.getImages` to query the available images in the specified region. The value must be an integral multiple of 20.
        /// </summary>
        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        /// <summary>
        /// The name of the simple application server.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        /// <summary>
        /// The password of the simple application server. The password must be 8 to 30 characters in length. It must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include: `( ) ~ ! @ # $ % ^ &amp; * - + = | { } [ ] : ; &lt; &gt; , . ? /`.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The paymen type of the resource. Valid values: `Subscription`.
        /// </summary>
        [Input("paymentType")]
        public Input<string>? PaymentType { get; set; }

        /// <summary>
        /// The period. Unit: months. Valid values: `1`,`3`, `6`, `12`, `24`, `36`.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// The ID of the plan. You can use the `alicloud.simpleapplicationserver.getServerPlans`  to query all the plans provided by Simple Application Server in the specified region.
        /// </summary>
        [Input("planId")]
        public Input<string>? PlanId { get; set; }

        /// <summary>
        /// The status of the simple application server. Valid values: `Resetting`, `Running`, `Stopped`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public InstanceState()
        {
        }
    }
}

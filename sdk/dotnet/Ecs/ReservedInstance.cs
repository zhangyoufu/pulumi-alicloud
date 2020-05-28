// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ecs
{
    /// <summary>
    /// Provides an Reserved Instance resource.
    /// 
    /// &gt; **NOTE:** Available in 1.65.0+
    /// 
    /// ## Example Usage
    /// 
    /// 
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var @default = new AliCloud.Ecs.ReservedInstance("default", new AliCloud.Ecs.ReservedInstanceArgs
    ///         {
    ///             InstanceType = "ecs.g6.large",
    ///             InstanceAmount = "1",
    ///             PeriodUnit = "Year",
    ///             OfferingType = "All Upfront",
    ///             Description = "ReservedInstance",
    ///             ZoneId = "cn-shanghai-g",
    ///             Scope = "Zone",
    ///             Period = "1",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class ReservedInstance : Pulumi.CustomResource
    {
        /// <summary>
        /// Description of the RI. 2 to 256 English or Chinese characters. It cannot start with http:// or https://.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Number of instances allocated to an RI (An RI is a coupon that includes one or more allocated instances.).
        /// </summary>
        [Output("instanceAmount")]
        public Output<int> InstanceAmount { get; private set; } = null!;

        /// <summary>
        /// Instance type of the RI. For more information, see [Instance type families](https://www.alibabacloud.com/help/doc-detail/25378.html).
        /// </summary>
        [Output("instanceType")]
        public Output<string> InstanceType { get; private set; } = null!;

        /// <summary>
        /// Name of the RI. The name must be a string of 2 to 128 characters in length and can contain letters, numbers, colons (:), underscores (_), and hyphens. It must start with a letter. It cannot start with http:// or https://.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Payment type of the RI. Optional values: `No Upfront`: No upfront payment is required., `Partial Upfront`: A portion of upfront payment is required.`All Upfront`: Full upfront payment is required.
        /// </summary>
        [Output("offeringType")]
        public Output<string?> OfferingType { get; private set; } = null!;

        [Output("period")]
        public Output<int?> Period { get; private set; } = null!;

        /// <summary>
        /// Term unit. Optional value: Year.
        /// </summary>
        [Output("periodUnit")]
        public Output<string?> PeriodUnit { get; private set; } = null!;

        /// <summary>
        /// The operating system type of the image used by the instance. Optional values: `Windows`, `Linux`. Default is `Linux`.
        /// </summary>
        [Output("platform")]
        public Output<string> Platform { get; private set; } = null!;

        /// <summary>
        /// Resource group ID.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// Scope of the RI. Optional values: `Region`: region-level, `Zone`: zone-level. Default is `Region`.
        /// </summary>
        [Output("scope")]
        public Output<string?> Scope { get; private set; } = null!;

        /// <summary>
        /// ID of the zone to which the RI belongs. When Scope is set to Zone, this parameter is required. For information about the zone list, see [DescribeZones](https://www.alibabacloud.com/help/doc-detail/25610.html).
        /// </summary>
        [Output("zoneId")]
        public Output<string?> ZoneId { get; private set; } = null!;


        /// <summary>
        /// Create a ReservedInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReservedInstance(string name, ReservedInstanceArgs args, CustomResourceOptions? options = null)
            : base("alicloud:ecs/reservedInstance:ReservedInstance", name, args ?? new ReservedInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ReservedInstance(string name, Input<string> id, ReservedInstanceState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ecs/reservedInstance:ReservedInstance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ReservedInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReservedInstance Get(string name, Input<string> id, ReservedInstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new ReservedInstance(name, id, state, options);
        }
    }

    public sealed class ReservedInstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the RI. 2 to 256 English or Chinese characters. It cannot start with http:// or https://.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Number of instances allocated to an RI (An RI is a coupon that includes one or more allocated instances.).
        /// </summary>
        [Input("instanceAmount")]
        public Input<int>? InstanceAmount { get; set; }

        /// <summary>
        /// Instance type of the RI. For more information, see [Instance type families](https://www.alibabacloud.com/help/doc-detail/25378.html).
        /// </summary>
        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        /// <summary>
        /// Name of the RI. The name must be a string of 2 to 128 characters in length and can contain letters, numbers, colons (:), underscores (_), and hyphens. It must start with a letter. It cannot start with http:// or https://.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Payment type of the RI. Optional values: `No Upfront`: No upfront payment is required., `Partial Upfront`: A portion of upfront payment is required.`All Upfront`: Full upfront payment is required.
        /// </summary>
        [Input("offeringType")]
        public Input<string>? OfferingType { get; set; }

        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// Term unit. Optional value: Year.
        /// </summary>
        [Input("periodUnit")]
        public Input<string>? PeriodUnit { get; set; }

        /// <summary>
        /// The operating system type of the image used by the instance. Optional values: `Windows`, `Linux`. Default is `Linux`.
        /// </summary>
        [Input("platform")]
        public Input<string>? Platform { get; set; }

        /// <summary>
        /// Resource group ID.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// Scope of the RI. Optional values: `Region`: region-level, `Zone`: zone-level. Default is `Region`.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        /// <summary>
        /// ID of the zone to which the RI belongs. When Scope is set to Zone, this parameter is required. For information about the zone list, see [DescribeZones](https://www.alibabacloud.com/help/doc-detail/25610.html).
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public ReservedInstanceArgs()
        {
        }
    }

    public sealed class ReservedInstanceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the RI. 2 to 256 English or Chinese characters. It cannot start with http:// or https://.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Number of instances allocated to an RI (An RI is a coupon that includes one or more allocated instances.).
        /// </summary>
        [Input("instanceAmount")]
        public Input<int>? InstanceAmount { get; set; }

        /// <summary>
        /// Instance type of the RI. For more information, see [Instance type families](https://www.alibabacloud.com/help/doc-detail/25378.html).
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// Name of the RI. The name must be a string of 2 to 128 characters in length and can contain letters, numbers, colons (:), underscores (_), and hyphens. It must start with a letter. It cannot start with http:// or https://.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Payment type of the RI. Optional values: `No Upfront`: No upfront payment is required., `Partial Upfront`: A portion of upfront payment is required.`All Upfront`: Full upfront payment is required.
        /// </summary>
        [Input("offeringType")]
        public Input<string>? OfferingType { get; set; }

        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// Term unit. Optional value: Year.
        /// </summary>
        [Input("periodUnit")]
        public Input<string>? PeriodUnit { get; set; }

        /// <summary>
        /// The operating system type of the image used by the instance. Optional values: `Windows`, `Linux`. Default is `Linux`.
        /// </summary>
        [Input("platform")]
        public Input<string>? Platform { get; set; }

        /// <summary>
        /// Resource group ID.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// Scope of the RI. Optional values: `Region`: region-level, `Zone`: zone-level. Default is `Region`.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        /// <summary>
        /// ID of the zone to which the RI belongs. When Scope is set to Zone, this parameter is required. For information about the zone list, see [DescribeZones](https://www.alibabacloud.com/help/doc-detail/25610.html).
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public ReservedInstanceState()
        {
        }
    }
}

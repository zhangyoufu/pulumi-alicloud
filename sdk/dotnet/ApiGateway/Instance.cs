// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ApiGateway
{
    /// <summary>
    /// Provides a Api Gateway Instance resource.
    /// 
    /// For information about Api Gateway Instance and how to use it, see [What is Instance](https://www.alibabacloud.com/help/en/api-gateway/product-overview/dedicated-instances).
    /// 
    /// &gt; **NOTE:** Available since v1.218.0.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "terraform-example";
    ///     var @default = new AliCloud.ApiGateway.Instance("default", new()
    ///     {
    ///         InstanceName = name,
    ///         InstanceSpec = "api.s1.small",
    ///         HttpsPolicy = "HTTPS2_TLS1_0",
    ///         ZoneId = "cn-hangzhou-MAZ6",
    ///         PaymentType = "PayAsYouGo",
    ///         UserVpcId = "1709116870",
    ///         InstanceType = "normal",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Api Gateway Instance can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:apigateway/instance:Instance example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:apigateway/instance:Instance")]
    public partial class Instance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Creation time.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The time of the instance package. Valid values:
        /// - PricingCycle is **Month**, indicating monthly payment. The value range is **1** to **9**.
        /// - PricingCycle is **Year**, indicating annual payment. The value range is **1** to **3**.
        /// 
        /// When the value of&gt; ChargeType is **PrePaid**, this parameter is available and must be passed in.
        /// </summary>
        [Output("duration")]
        public Output<int?> Duration { get; private set; } = null!;

        /// <summary>
        /// Does IPV6 Capability Support.
        /// </summary>
        [Output("egressIpv6Enable")]
        public Output<bool?> EgressIpv6Enable { get; private set; } = null!;

        /// <summary>
        /// Https policy.
        /// </summary>
        [Output("httpsPolicy")]
        public Output<string> HttpsPolicy { get; private set; } = null!;

        /// <summary>
        /// Instance name.
        /// </summary>
        [Output("instanceName")]
        public Output<string> InstanceName { get; private set; } = null!;

        /// <summary>
        /// Instance type.
        /// </summary>
        [Output("instanceSpec")]
        public Output<string> InstanceSpec { get; private set; } = null!;

        /// <summary>
        /// Instance type-normal: traditional exclusive instance.
        /// </summary>
        [Output("instanceType")]
        public Output<string> InstanceType { get; private set; } = null!;

        /// <summary>
        /// The payment type of the resource.
        /// </summary>
        [Output("paymentType")]
        public Output<string> PaymentType { get; private set; } = null!;

        /// <summary>
        /// The subscription instance is of the subscription year or month type. The value range is as follows:
        /// - **year**: year
        /// - **month**: month
        /// &gt; **NOTE:**  If the Payment type is PrePaid, this parameter is required.
        /// </summary>
        [Output("pricingCycle")]
        public Output<string?> PricingCycle { get; private set; } = null!;

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Does ipv6 support.
        /// </summary>
        [Output("supportIpv6")]
        public Output<bool?> SupportIpv6 { get; private set; } = null!;

        /// <summary>
        /// User's VpcID.
        /// </summary>
        [Output("userVpcId")]
        public Output<string?> UserVpcId { get; private set; } = null!;

        /// <summary>
        /// Whether the slb of the Vpc supports.
        /// </summary>
        [Output("vpcSlbIntranetEnable")]
        public Output<bool?> VpcSlbIntranetEnable { get; private set; } = null!;

        /// <summary>
        /// The zone where the instance is deployed.
        /// </summary>
        [Output("zoneId")]
        public Output<string> ZoneId { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("alicloud:apigateway/instance:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:apigateway/instance:Instance", name, state, MakeResourceOptions(options, id))
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

    public sealed class InstanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The time of the instance package. Valid values:
        /// - PricingCycle is **Month**, indicating monthly payment. The value range is **1** to **9**.
        /// - PricingCycle is **Year**, indicating annual payment. The value range is **1** to **3**.
        /// 
        /// When the value of&gt; ChargeType is **PrePaid**, this parameter is available and must be passed in.
        /// </summary>
        [Input("duration")]
        public Input<int>? Duration { get; set; }

        /// <summary>
        /// Does IPV6 Capability Support.
        /// </summary>
        [Input("egressIpv6Enable")]
        public Input<bool>? EgressIpv6Enable { get; set; }

        /// <summary>
        /// Https policy.
        /// </summary>
        [Input("httpsPolicy", required: true)]
        public Input<string> HttpsPolicy { get; set; } = null!;

        /// <summary>
        /// Instance name.
        /// </summary>
        [Input("instanceName", required: true)]
        public Input<string> InstanceName { get; set; } = null!;

        /// <summary>
        /// Instance type.
        /// </summary>
        [Input("instanceSpec", required: true)]
        public Input<string> InstanceSpec { get; set; } = null!;

        /// <summary>
        /// Instance type-normal: traditional exclusive instance.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// The payment type of the resource.
        /// </summary>
        [Input("paymentType", required: true)]
        public Input<string> PaymentType { get; set; } = null!;

        /// <summary>
        /// The subscription instance is of the subscription year or month type. The value range is as follows:
        /// - **year**: year
        /// - **month**: month
        /// &gt; **NOTE:**  If the Payment type is PrePaid, this parameter is required.
        /// </summary>
        [Input("pricingCycle")]
        public Input<string>? PricingCycle { get; set; }

        /// <summary>
        /// Does ipv6 support.
        /// </summary>
        [Input("supportIpv6")]
        public Input<bool>? SupportIpv6 { get; set; }

        /// <summary>
        /// User's VpcID.
        /// </summary>
        [Input("userVpcId")]
        public Input<string>? UserVpcId { get; set; }

        /// <summary>
        /// Whether the slb of the Vpc supports.
        /// </summary>
        [Input("vpcSlbIntranetEnable")]
        public Input<bool>? VpcSlbIntranetEnable { get; set; }

        /// <summary>
        /// The zone where the instance is deployed.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public InstanceArgs()
        {
        }
        public static new InstanceArgs Empty => new InstanceArgs();
    }

    public sealed class InstanceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Creation time.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The time of the instance package. Valid values:
        /// - PricingCycle is **Month**, indicating monthly payment. The value range is **1** to **9**.
        /// - PricingCycle is **Year**, indicating annual payment. The value range is **1** to **3**.
        /// 
        /// When the value of&gt; ChargeType is **PrePaid**, this parameter is available and must be passed in.
        /// </summary>
        [Input("duration")]
        public Input<int>? Duration { get; set; }

        /// <summary>
        /// Does IPV6 Capability Support.
        /// </summary>
        [Input("egressIpv6Enable")]
        public Input<bool>? EgressIpv6Enable { get; set; }

        /// <summary>
        /// Https policy.
        /// </summary>
        [Input("httpsPolicy")]
        public Input<string>? HttpsPolicy { get; set; }

        /// <summary>
        /// Instance name.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        /// <summary>
        /// Instance type.
        /// </summary>
        [Input("instanceSpec")]
        public Input<string>? InstanceSpec { get; set; }

        /// <summary>
        /// Instance type-normal: traditional exclusive instance.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// The payment type of the resource.
        /// </summary>
        [Input("paymentType")]
        public Input<string>? PaymentType { get; set; }

        /// <summary>
        /// The subscription instance is of the subscription year or month type. The value range is as follows:
        /// - **year**: year
        /// - **month**: month
        /// &gt; **NOTE:**  If the Payment type is PrePaid, this parameter is required.
        /// </summary>
        [Input("pricingCycle")]
        public Input<string>? PricingCycle { get; set; }

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Does ipv6 support.
        /// </summary>
        [Input("supportIpv6")]
        public Input<bool>? SupportIpv6 { get; set; }

        /// <summary>
        /// User's VpcID.
        /// </summary>
        [Input("userVpcId")]
        public Input<string>? UserVpcId { get; set; }

        /// <summary>
        /// Whether the slb of the Vpc supports.
        /// </summary>
        [Input("vpcSlbIntranetEnable")]
        public Input<bool>? VpcSlbIntranetEnable { get; set; }

        /// <summary>
        /// The zone where the instance is deployed.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public InstanceState()
        {
        }
        public static new InstanceState Empty => new InstanceState();
    }
}

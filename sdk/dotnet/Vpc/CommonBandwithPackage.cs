// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vpc
{
    /// <summary>
    /// ## Import
    /// 
    /// CBWP Common Bandwidth Package can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:vpc/commonBandwithPackage:CommonBandwithPackage example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:vpc/commonBandwithPackage:CommonBandwithPackage")]
    public partial class CommonBandwithPackage : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The maximum bandwidth of the Internet Shared Bandwidth instance. Unit: Mbit/s. Valid values: `1` to `1000`. Default value: `1`.
        /// </summary>
        [Output("bandwidth")]
        public Output<string> Bandwidth { get; private set; } = null!;

        /// <summary>
        /// The name of the Internet Shared Bandwidth instance. The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (\_), and hyphens (-). The name must start with a letter.
        /// </summary>
        [Output("bandwidthPackageName")]
        public Output<string> BandwidthPackageName { get; private set; } = null!;

        /// <summary>
        /// The creation time.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to enable deletion protection. Valid values:
        /// </summary>
        [Output("deletionProtection")]
        public Output<bool?> DeletionProtection { get; private set; } = null!;

        /// <summary>
        /// The description of the Internet Shared Bandwidth instance. The description must be 2 to 256 characters in length and start with a letter. The description cannot start with `http://` or `https://`.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to forcefully delete the Internet Shared Bandwidth instance. Valid values:
        /// </summary>
        [Output("force")]
        public Output<string?> Force { get; private set; } = null!;

        /// <summary>
        /// The billing method of the Internet Shared Bandwidth instance. Set the value to `PayByTraffic`, which specifies the pay-by-data-transfer billing method.
        /// </summary>
        [Output("internetChargeType")]
        public Output<string?> InternetChargeType { get; private set; } = null!;

        /// <summary>
        /// The line type. Valid values:
        /// - `BGP` All regions support BGP (Multi-ISP).
        /// - `BGP_PRO` BGP (Multi-ISP) Pro lines are available in the China (Hong Kong), Singapore, Japan (Tokyo), Philippines (Manila), Malaysia (Kuala Lumpur), Indonesia (Jakarta), and Thailand (Bangkok) regions.
        /// 
        /// If you are allowed to use single-ISP bandwidth, you can also use one of the following values:
        /// - `ChinaTelecom`
        /// - `ChinaUnicom`
        /// - `ChinaMobile`
        /// - `ChinaTelecom_L2`
        /// - `ChinaUnicom_L2`
        /// - `ChinaMobile_L2`
        /// 
        /// If your services are deployed in China East 1 Finance, this parameter is required and you must set the value to `BGP_FinanceCloud`.
        /// </summary>
        [Output("isp")]
        public Output<string> Isp { get; private set; } = null!;

        /// <summary>
        /// . Field 'name' has been deprecated from provider version 1.120.0. New field 'bandwidth_package_name' instead.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The billing type of the Internet Shared Bandwidth instance. Valid values: `PayAsYouGo`, `Subscription`.
        /// </summary>
        [Output("paymentType")]
        public Output<string> PaymentType { get; private set; } = null!;

        /// <summary>
        /// The percentage of the minimum bandwidth commitment. Set the parameter to `20`.
        /// 
        /// &gt; **NOTE:**  This parameter is available only on the Alibaba Cloud China site.
        /// </summary>
        [Output("ratio")]
        public Output<int> Ratio { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource group to which you want to move the resource.
        /// 
        /// &gt; **NOTE:**   You can use resource groups to facilitate resource grouping and permission management for an Alibaba Cloud. For more information, see [What is resource management?](https://www.alibabacloud.com/help/en/doc-detail/94475.html)
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// The edition of Anti-DDoS. If you do not set this parameter, Anti-DDoS Origin Basic is used. If you set the value to AntiDDoS_Enhanced, Anti-DDoS Pro(Premium) is used. It is valid when `internet_charge_type` is `PayBy95`.
        /// </summary>
        [Output("securityProtectionTypes")]
        public Output<ImmutableArray<string>> SecurityProtectionTypes { get; private set; } = null!;

        /// <summary>
        /// The status of the Internet Shared Bandwidth instance. Default value: `Available`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The tag of the resource
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The zone of the Internet Shared Bandwidth instance. This parameter is required if you create an Internet Shared Bandwidth instance for a cloud box. 
        /// 
        /// The following arguments will be discarded. Please use new fields as soon as possible:
        /// </summary>
        [Output("zone")]
        public Output<string?> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a CommonBandwithPackage resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CommonBandwithPackage(string name, CommonBandwithPackageArgs args, CustomResourceOptions? options = null)
            : base("alicloud:vpc/commonBandwithPackage:CommonBandwithPackage", name, args ?? new CommonBandwithPackageArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CommonBandwithPackage(string name, Input<string> id, CommonBandwithPackageState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:vpc/commonBandwithPackage:CommonBandwithPackage", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CommonBandwithPackage resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CommonBandwithPackage Get(string name, Input<string> id, CommonBandwithPackageState? state = null, CustomResourceOptions? options = null)
        {
            return new CommonBandwithPackage(name, id, state, options);
        }
    }

    public sealed class CommonBandwithPackageArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum bandwidth of the Internet Shared Bandwidth instance. Unit: Mbit/s. Valid values: `1` to `1000`. Default value: `1`.
        /// </summary>
        [Input("bandwidth", required: true)]
        public Input<string> Bandwidth { get; set; } = null!;

        /// <summary>
        /// The name of the Internet Shared Bandwidth instance. The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (\_), and hyphens (-). The name must start with a letter.
        /// </summary>
        [Input("bandwidthPackageName")]
        public Input<string>? BandwidthPackageName { get; set; }

        /// <summary>
        /// Specifies whether to enable deletion protection. Valid values:
        /// </summary>
        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        /// <summary>
        /// The description of the Internet Shared Bandwidth instance. The description must be 2 to 256 characters in length and start with a letter. The description cannot start with `http://` or `https://`.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies whether to forcefully delete the Internet Shared Bandwidth instance. Valid values:
        /// </summary>
        [Input("force")]
        public Input<string>? Force { get; set; }

        /// <summary>
        /// The billing method of the Internet Shared Bandwidth instance. Set the value to `PayByTraffic`, which specifies the pay-by-data-transfer billing method.
        /// </summary>
        [Input("internetChargeType")]
        public Input<string>? InternetChargeType { get; set; }

        /// <summary>
        /// The line type. Valid values:
        /// - `BGP` All regions support BGP (Multi-ISP).
        /// - `BGP_PRO` BGP (Multi-ISP) Pro lines are available in the China (Hong Kong), Singapore, Japan (Tokyo), Philippines (Manila), Malaysia (Kuala Lumpur), Indonesia (Jakarta), and Thailand (Bangkok) regions.
        /// 
        /// If you are allowed to use single-ISP bandwidth, you can also use one of the following values:
        /// - `ChinaTelecom`
        /// - `ChinaUnicom`
        /// - `ChinaMobile`
        /// - `ChinaTelecom_L2`
        /// - `ChinaUnicom_L2`
        /// - `ChinaMobile_L2`
        /// 
        /// If your services are deployed in China East 1 Finance, this parameter is required and you must set the value to `BGP_FinanceCloud`.
        /// </summary>
        [Input("isp")]
        public Input<string>? Isp { get; set; }

        /// <summary>
        /// . Field 'name' has been deprecated from provider version 1.120.0. New field 'bandwidth_package_name' instead.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The percentage of the minimum bandwidth commitment. Set the parameter to `20`.
        /// 
        /// &gt; **NOTE:**  This parameter is available only on the Alibaba Cloud China site.
        /// </summary>
        [Input("ratio")]
        public Input<int>? Ratio { get; set; }

        /// <summary>
        /// The ID of the resource group to which you want to move the resource.
        /// 
        /// &gt; **NOTE:**   You can use resource groups to facilitate resource grouping and permission management for an Alibaba Cloud. For more information, see [What is resource management?](https://www.alibabacloud.com/help/en/doc-detail/94475.html)
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("securityProtectionTypes")]
        private InputList<string>? _securityProtectionTypes;

        /// <summary>
        /// The edition of Anti-DDoS. If you do not set this parameter, Anti-DDoS Origin Basic is used. If you set the value to AntiDDoS_Enhanced, Anti-DDoS Pro(Premium) is used. It is valid when `internet_charge_type` is `PayBy95`.
        /// </summary>
        public InputList<string> SecurityProtectionTypes
        {
            get => _securityProtectionTypes ?? (_securityProtectionTypes = new InputList<string>());
            set => _securityProtectionTypes = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The tag of the resource
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The zone of the Internet Shared Bandwidth instance. This parameter is required if you create an Internet Shared Bandwidth instance for a cloud box. 
        /// 
        /// The following arguments will be discarded. Please use new fields as soon as possible:
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public CommonBandwithPackageArgs()
        {
        }
        public static new CommonBandwithPackageArgs Empty => new CommonBandwithPackageArgs();
    }

    public sealed class CommonBandwithPackageState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum bandwidth of the Internet Shared Bandwidth instance. Unit: Mbit/s. Valid values: `1` to `1000`. Default value: `1`.
        /// </summary>
        [Input("bandwidth")]
        public Input<string>? Bandwidth { get; set; }

        /// <summary>
        /// The name of the Internet Shared Bandwidth instance. The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (\_), and hyphens (-). The name must start with a letter.
        /// </summary>
        [Input("bandwidthPackageName")]
        public Input<string>? BandwidthPackageName { get; set; }

        /// <summary>
        /// The creation time.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Specifies whether to enable deletion protection. Valid values:
        /// </summary>
        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        /// <summary>
        /// The description of the Internet Shared Bandwidth instance. The description must be 2 to 256 characters in length and start with a letter. The description cannot start with `http://` or `https://`.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies whether to forcefully delete the Internet Shared Bandwidth instance. Valid values:
        /// </summary>
        [Input("force")]
        public Input<string>? Force { get; set; }

        /// <summary>
        /// The billing method of the Internet Shared Bandwidth instance. Set the value to `PayByTraffic`, which specifies the pay-by-data-transfer billing method.
        /// </summary>
        [Input("internetChargeType")]
        public Input<string>? InternetChargeType { get; set; }

        /// <summary>
        /// The line type. Valid values:
        /// - `BGP` All regions support BGP (Multi-ISP).
        /// - `BGP_PRO` BGP (Multi-ISP) Pro lines are available in the China (Hong Kong), Singapore, Japan (Tokyo), Philippines (Manila), Malaysia (Kuala Lumpur), Indonesia (Jakarta), and Thailand (Bangkok) regions.
        /// 
        /// If you are allowed to use single-ISP bandwidth, you can also use one of the following values:
        /// - `ChinaTelecom`
        /// - `ChinaUnicom`
        /// - `ChinaMobile`
        /// - `ChinaTelecom_L2`
        /// - `ChinaUnicom_L2`
        /// - `ChinaMobile_L2`
        /// 
        /// If your services are deployed in China East 1 Finance, this parameter is required and you must set the value to `BGP_FinanceCloud`.
        /// </summary>
        [Input("isp")]
        public Input<string>? Isp { get; set; }

        /// <summary>
        /// . Field 'name' has been deprecated from provider version 1.120.0. New field 'bandwidth_package_name' instead.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The billing type of the Internet Shared Bandwidth instance. Valid values: `PayAsYouGo`, `Subscription`.
        /// </summary>
        [Input("paymentType")]
        public Input<string>? PaymentType { get; set; }

        /// <summary>
        /// The percentage of the minimum bandwidth commitment. Set the parameter to `20`.
        /// 
        /// &gt; **NOTE:**  This parameter is available only on the Alibaba Cloud China site.
        /// </summary>
        [Input("ratio")]
        public Input<int>? Ratio { get; set; }

        /// <summary>
        /// The ID of the resource group to which you want to move the resource.
        /// 
        /// &gt; **NOTE:**   You can use resource groups to facilitate resource grouping and permission management for an Alibaba Cloud. For more information, see [What is resource management?](https://www.alibabacloud.com/help/en/doc-detail/94475.html)
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("securityProtectionTypes")]
        private InputList<string>? _securityProtectionTypes;

        /// <summary>
        /// The edition of Anti-DDoS. If you do not set this parameter, Anti-DDoS Origin Basic is used. If you set the value to AntiDDoS_Enhanced, Anti-DDoS Pro(Premium) is used. It is valid when `internet_charge_type` is `PayBy95`.
        /// </summary>
        public InputList<string> SecurityProtectionTypes
        {
            get => _securityProtectionTypes ?? (_securityProtectionTypes = new InputList<string>());
            set => _securityProtectionTypes = value;
        }

        /// <summary>
        /// The status of the Internet Shared Bandwidth instance. Default value: `Available`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The tag of the resource
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The zone of the Internet Shared Bandwidth instance. This parameter is required if you create an Internet Shared Bandwidth instance for a cloud box. 
        /// 
        /// The following arguments will be discarded. Please use new fields as soon as possible:
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public CommonBandwithPackageState()
        {
        }
        public static new CommonBandwithPackageState Empty => new CommonBandwithPackageState();
    }
}

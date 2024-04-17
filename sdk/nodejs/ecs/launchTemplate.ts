// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides an ECS Launch Template resource.
 *
 * For information about Launch Template and how to use it, see [Launch Template](https://www.alibabacloud.com/help/doc-detail/73916.html).
 *
 * > **DEPRECATED:**  This resource  has been deprecated from version `1.120.0`. Please use new resource alicloud_ecs_launch_template.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const images = alicloud.ecs.getImages({
 *     owners: "system",
 * });
 * const instances = alicloud.ecs.getInstances({});
 * const template = new alicloud.ecs.LaunchTemplate("template", {
 *     name: "tf-test-template",
 *     description: "test1",
 *     imageId: images.then(images => images.images?.[0]?.id),
 *     hostName: "tf-test-host",
 *     instanceChargeType: "PrePaid",
 *     instanceName: "tf-instance-name",
 *     instanceType: instances.then(instances => instances.instances?.[0]?.instanceType),
 *     internetChargeType: "PayByBandwidth",
 *     internetMaxBandwidthIn: 5,
 *     internetMaxBandwidthOut: 0,
 *     ioOptimized: "none",
 *     keyPairName: "test-key-pair",
 *     ramRoleName: "xxxxx",
 *     networkType: "vpc",
 *     securityEnhancementStrategy: "Active",
 *     spotPriceLimit: 5,
 *     spotStrategy: "SpotWithPriceLimit",
 *     securityGroupId: "sg-zxcvj0lasdf102350asdf9a",
 *     systemDiskCategory: "cloud_ssd",
 *     systemDiskDescription: "test disk",
 *     systemDiskName: "hello",
 *     systemDiskSize: 40,
 *     resourceGroupId: "rg-zkdfjahg9zxncv0",
 *     userdata: "xxxxxxxxxxxxxx",
 *     vswitchId: "sw-ljkngaksdjfj0nnasdf",
 *     vpcId: "vpc-asdfnbg0as8dfk1nb2",
 *     zoneId: "beijing-a",
 *     tags: {
 *         tag1: "hello",
 *         tag2: "world",
 *     },
 *     networkInterfaces: {
 *         name: "eth0",
 *         description: "hello1",
 *         primaryIp: "10.0.0.2",
 *         securityGroupId: "xxxx",
 *         vswitchId: "xxxxxxx",
 *     },
 *     dataDisks: [
 *         {
 *             name: "disk1",
 *             description: "test1",
 *         },
 *         {
 *             name: "disk2",
 *             description: "test2",
 *         },
 *     ],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Launch Template can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:ecs/launchTemplate:LaunchTemplate lt lt-abc1234567890000
 * ```
 */
export class LaunchTemplate extends pulumi.CustomResource {
    /**
     * Get an existing LaunchTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LaunchTemplateState, opts?: pulumi.CustomResourceOptions): LaunchTemplate {
        return new LaunchTemplate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ecs/launchTemplate:LaunchTemplate';

    /**
     * Returns true if the given object is an instance of LaunchTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LaunchTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LaunchTemplate.__pulumiType;
    }

    /**
     * Instance auto release time. The time is presented using the ISO8601 standard and in UTC time. The format is  YYYY-MM-DDTHH:MM:SSZ.
     */
    public readonly autoReleaseTime!: pulumi.Output<string | undefined>;
    /**
     * The list of data disks created with instance.
     */
    public readonly dataDisks!: pulumi.Output<outputs.ecs.LaunchTemplateDataDisk[] | undefined>;
    public readonly deploymentSetId!: pulumi.Output<string | undefined>;
    /**
     * The description of the data disk.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly enableVmOsConfig!: pulumi.Output<boolean | undefined>;
    /**
     * Instance host name.It cannot start or end with a period (.) or a hyphen (-) and it cannot have two or more consecutive periods (.) or hyphens (-).For Windows: The host name can be [2, 15] characters in length. It can contain A-Z, a-z, numbers, periods (.), and hyphens (-). It cannot only contain numbers. For other operating systems: The host name can be [2, 64] characters in length. It can be segments separated by periods (.). It can contain A-Z, a-z, numbers, and hyphens (-).
     */
    public readonly hostName!: pulumi.Output<string | undefined>;
    /**
     * Image ID.
     */
    public readonly imageId!: pulumi.Output<string | undefined>;
    public readonly imageOwnerAlias!: pulumi.Output<string | undefined>;
    /**
     * Billing methods. Optional values:
     * - PrePaid: Monthly, or annual subscription. Make sure that your registered credit card is invalid or you have insufficient balance in your PayPal account. Otherwise, InvalidPayMethod error may occur.
     * - PostPaid: Pay-As-You-Go.
     *
     * Default value: PostPaid.
     */
    public readonly instanceChargeType!: pulumi.Output<string | undefined>;
    /**
     * The name of the instance. The name is a string of 2 to 128 characters. It must begin with an English or a Chinese character. It can contain A-Z, a-z, Chinese characters, numbers, periods (.), colons (:), underscores (_), and hyphens (-).
     */
    public readonly instanceName!: pulumi.Output<string | undefined>;
    /**
     * Instance type. For more information, call resourceAlicloudInstances to obtain the latest instance type list.
     */
    public readonly instanceType!: pulumi.Output<string | undefined>;
    /**
     * Internet bandwidth billing method. Optional values: `PayByTraffic` | `PayByBandwidth`.
     */
    public readonly internetChargeType!: pulumi.Output<string | undefined>;
    /**
     * The maximum inbound bandwidth from the Internet network, measured in Mbit/s. Value range: [1, 200].
     */
    public readonly internetMaxBandwidthIn!: pulumi.Output<number>;
    /**
     * Maximum outbound bandwidth from the Internet, its unit of measurement is Mbit/s. Value range: [0, 100].
     */
    public readonly internetMaxBandwidthOut!: pulumi.Output<number | undefined>;
    /**
     * Whether it is an I/O-optimized instance or not. Optional values:
     * - none
     * - optimized
     */
    public readonly ioOptimized!: pulumi.Output<string | undefined>;
    /**
     * The name of the key pair.
     * - Ignore this parameter for Windows instances. It is null by default. Even if you enter this parameter, only the  Password content is used.
     * - The password logon method for Linux instances is set to forbidden upon initialization.
     */
    public readonly keyPairName!: pulumi.Output<string | undefined>;
    public readonly launchTemplateName!: pulumi.Output<string>;
    /**
     * The name of the data disk.
     *
     * @deprecated Field 'name' has been deprecated from provider version 1.120.0. New field 'launch_template_name' instead.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The list of network interfaces created with instance.
     */
    public readonly networkInterfaces!: pulumi.Output<outputs.ecs.LaunchTemplateNetworkInterfaces | undefined>;
    /**
     * Network type of the instance. Value options: `classic` | `vpc`.
     */
    public readonly networkType!: pulumi.Output<string | undefined>;
    public readonly passwordInherit!: pulumi.Output<boolean | undefined>;
    public readonly period!: pulumi.Output<number | undefined>;
    public readonly privateIpAddress!: pulumi.Output<string | undefined>;
    /**
     * The RAM role name of the instance. You can use the RAM API ListRoles to query instance RAM role names.
     */
    public readonly ramRoleName!: pulumi.Output<string | undefined>;
    public readonly resourceGroupId!: pulumi.Output<string | undefined>;
    /**
     * Whether or not to activate the security enhancement feature and install network security software free of charge. Optional values: Active | Deactive.
     */
    public readonly securityEnhancementStrategy!: pulumi.Output<string | undefined>;
    /**
     * The security group ID must be one in the same VPC.
     */
    public readonly securityGroupId!: pulumi.Output<string | undefined>;
    public readonly securityGroupIds!: pulumi.Output<string[] | undefined>;
    public readonly spotDuration!: pulumi.Output<string | undefined>;
    /**
     * Sets the maximum hourly instance price. Supports up to three decimal places.
     */
    public readonly spotPriceLimit!: pulumi.Output<number | undefined>;
    /**
     * The spot strategy for a Pay-As-You-Go instance. This parameter is valid and required only when InstanceChargeType is set to PostPaid. Value range:
     * - NoSpot: Normal Pay-As-You-Go instance.
     * - SpotWithPriceLimit: Sets the maximum price for a spot instance.
     * - SpotAsPriceGo: The system automatically calculates the price. The maximum value is the Pay-As-You-Go price.
     */
    public readonly spotStrategy!: pulumi.Output<string | undefined>;
    public readonly systemDisk!: pulumi.Output<outputs.ecs.LaunchTemplateSystemDisk>;
    /**
     * The category of the system disk. System disk type. Optional values:
     * - cloud: Basic cloud disk.
     * - cloud_efficiency: Ultra cloud disk.
     * - cloud_ssd: SSD cloud Disks.
     * - ephemeral_ssd: local SSD Disks
     * - cloud_essd: ESSD cloud Disks.
     *
     * @deprecated Field 'system_disk_category' has been deprecated from provider version 1.120.0. New field 'system_disk' instead.
     */
    public readonly systemDiskCategory!: pulumi.Output<string>;
    /**
     * System disk description. It cannot begin with http:// or https://.
     *
     * @deprecated Field 'system_disk_description' has been deprecated from provider version 1.120.0. New field 'system_disk' instead.
     */
    public readonly systemDiskDescription!: pulumi.Output<string>;
    /**
     * System disk name. The name is a string of 2 to 128 characters. It must begin with an English or a Chinese character. It can contain A-Z, a-z, Chinese characters, numbers, periods (.), colons (:), underscores (_), and hyphens (-).
     *
     * @deprecated Field 'system_disk_name' has been deprecated from provider version 1.120.0. New field 'system_disk' instead.
     */
    public readonly systemDiskName!: pulumi.Output<string>;
    /**
     * Size of the system disk, measured in GB. Value range: [20, 500].
     *
     * @deprecated Field 'system_disk_size' has been deprecated from provider version 1.120.0. New field 'system_disk' instead.
     */
    public readonly systemDiskSize!: pulumi.Output<number>;
    /**
     * A mapping of tags to assign to the resource.
     * - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    public readonly templateResourceGroupId!: pulumi.Output<string | undefined>;
    public readonly templateTags!: pulumi.Output<{[key: string]: any} | undefined>;
    public readonly userData!: pulumi.Output<string>;
    /**
     * User data of the instance, which is Base64-encoded. Size of the raw data cannot exceed 16 KB.
     *
     * @deprecated Field 'userdata' has been deprecated from provider version 1.120.0. New field 'user_data' instead.
     */
    public readonly userdata!: pulumi.Output<string>;
    public readonly versionDescription!: pulumi.Output<string | undefined>;
    public readonly vpcId!: pulumi.Output<string>;
    /**
     * The VSwitch ID for ENI. The instance must be in the same zone of the same VPC network as the ENI, but they may belong to different VSwitches.
     */
    public readonly vswitchId!: pulumi.Output<string | undefined>;
    /**
     * The zone ID of the instance.
     */
    public readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a LaunchTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: LaunchTemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LaunchTemplateArgs | LaunchTemplateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LaunchTemplateState | undefined;
            resourceInputs["autoReleaseTime"] = state ? state.autoReleaseTime : undefined;
            resourceInputs["dataDisks"] = state ? state.dataDisks : undefined;
            resourceInputs["deploymentSetId"] = state ? state.deploymentSetId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enableVmOsConfig"] = state ? state.enableVmOsConfig : undefined;
            resourceInputs["hostName"] = state ? state.hostName : undefined;
            resourceInputs["imageId"] = state ? state.imageId : undefined;
            resourceInputs["imageOwnerAlias"] = state ? state.imageOwnerAlias : undefined;
            resourceInputs["instanceChargeType"] = state ? state.instanceChargeType : undefined;
            resourceInputs["instanceName"] = state ? state.instanceName : undefined;
            resourceInputs["instanceType"] = state ? state.instanceType : undefined;
            resourceInputs["internetChargeType"] = state ? state.internetChargeType : undefined;
            resourceInputs["internetMaxBandwidthIn"] = state ? state.internetMaxBandwidthIn : undefined;
            resourceInputs["internetMaxBandwidthOut"] = state ? state.internetMaxBandwidthOut : undefined;
            resourceInputs["ioOptimized"] = state ? state.ioOptimized : undefined;
            resourceInputs["keyPairName"] = state ? state.keyPairName : undefined;
            resourceInputs["launchTemplateName"] = state ? state.launchTemplateName : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networkInterfaces"] = state ? state.networkInterfaces : undefined;
            resourceInputs["networkType"] = state ? state.networkType : undefined;
            resourceInputs["passwordInherit"] = state ? state.passwordInherit : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["privateIpAddress"] = state ? state.privateIpAddress : undefined;
            resourceInputs["ramRoleName"] = state ? state.ramRoleName : undefined;
            resourceInputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            resourceInputs["securityEnhancementStrategy"] = state ? state.securityEnhancementStrategy : undefined;
            resourceInputs["securityGroupId"] = state ? state.securityGroupId : undefined;
            resourceInputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            resourceInputs["spotDuration"] = state ? state.spotDuration : undefined;
            resourceInputs["spotPriceLimit"] = state ? state.spotPriceLimit : undefined;
            resourceInputs["spotStrategy"] = state ? state.spotStrategy : undefined;
            resourceInputs["systemDisk"] = state ? state.systemDisk : undefined;
            resourceInputs["systemDiskCategory"] = state ? state.systemDiskCategory : undefined;
            resourceInputs["systemDiskDescription"] = state ? state.systemDiskDescription : undefined;
            resourceInputs["systemDiskName"] = state ? state.systemDiskName : undefined;
            resourceInputs["systemDiskSize"] = state ? state.systemDiskSize : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["templateResourceGroupId"] = state ? state.templateResourceGroupId : undefined;
            resourceInputs["templateTags"] = state ? state.templateTags : undefined;
            resourceInputs["userData"] = state ? state.userData : undefined;
            resourceInputs["userdata"] = state ? state.userdata : undefined;
            resourceInputs["versionDescription"] = state ? state.versionDescription : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["vswitchId"] = state ? state.vswitchId : undefined;
            resourceInputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as LaunchTemplateArgs | undefined;
            resourceInputs["autoReleaseTime"] = args ? args.autoReleaseTime : undefined;
            resourceInputs["dataDisks"] = args ? args.dataDisks : undefined;
            resourceInputs["deploymentSetId"] = args ? args.deploymentSetId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enableVmOsConfig"] = args ? args.enableVmOsConfig : undefined;
            resourceInputs["hostName"] = args ? args.hostName : undefined;
            resourceInputs["imageId"] = args ? args.imageId : undefined;
            resourceInputs["imageOwnerAlias"] = args ? args.imageOwnerAlias : undefined;
            resourceInputs["instanceChargeType"] = args ? args.instanceChargeType : undefined;
            resourceInputs["instanceName"] = args ? args.instanceName : undefined;
            resourceInputs["instanceType"] = args ? args.instanceType : undefined;
            resourceInputs["internetChargeType"] = args ? args.internetChargeType : undefined;
            resourceInputs["internetMaxBandwidthIn"] = args ? args.internetMaxBandwidthIn : undefined;
            resourceInputs["internetMaxBandwidthOut"] = args ? args.internetMaxBandwidthOut : undefined;
            resourceInputs["ioOptimized"] = args ? args.ioOptimized : undefined;
            resourceInputs["keyPairName"] = args ? args.keyPairName : undefined;
            resourceInputs["launchTemplateName"] = args ? args.launchTemplateName : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkInterfaces"] = args ? args.networkInterfaces : undefined;
            resourceInputs["networkType"] = args ? args.networkType : undefined;
            resourceInputs["passwordInherit"] = args ? args.passwordInherit : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["privateIpAddress"] = args ? args.privateIpAddress : undefined;
            resourceInputs["ramRoleName"] = args ? args.ramRoleName : undefined;
            resourceInputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            resourceInputs["securityEnhancementStrategy"] = args ? args.securityEnhancementStrategy : undefined;
            resourceInputs["securityGroupId"] = args ? args.securityGroupId : undefined;
            resourceInputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            resourceInputs["spotDuration"] = args ? args.spotDuration : undefined;
            resourceInputs["spotPriceLimit"] = args ? args.spotPriceLimit : undefined;
            resourceInputs["spotStrategy"] = args ? args.spotStrategy : undefined;
            resourceInputs["systemDisk"] = args ? args.systemDisk : undefined;
            resourceInputs["systemDiskCategory"] = args ? args.systemDiskCategory : undefined;
            resourceInputs["systemDiskDescription"] = args ? args.systemDiskDescription : undefined;
            resourceInputs["systemDiskName"] = args ? args.systemDiskName : undefined;
            resourceInputs["systemDiskSize"] = args ? args.systemDiskSize : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["templateResourceGroupId"] = args ? args.templateResourceGroupId : undefined;
            resourceInputs["templateTags"] = args ? args.templateTags : undefined;
            resourceInputs["userData"] = args ? args.userData : undefined;
            resourceInputs["userdata"] = args ? args.userdata : undefined;
            resourceInputs["versionDescription"] = args ? args.versionDescription : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["vswitchId"] = args ? args.vswitchId : undefined;
            resourceInputs["zoneId"] = args ? args.zoneId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LaunchTemplate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LaunchTemplate resources.
 */
export interface LaunchTemplateState {
    /**
     * Instance auto release time. The time is presented using the ISO8601 standard and in UTC time. The format is  YYYY-MM-DDTHH:MM:SSZ.
     */
    autoReleaseTime?: pulumi.Input<string>;
    /**
     * The list of data disks created with instance.
     */
    dataDisks?: pulumi.Input<pulumi.Input<inputs.ecs.LaunchTemplateDataDisk>[]>;
    deploymentSetId?: pulumi.Input<string>;
    /**
     * The description of the data disk.
     */
    description?: pulumi.Input<string>;
    enableVmOsConfig?: pulumi.Input<boolean>;
    /**
     * Instance host name.It cannot start or end with a period (.) or a hyphen (-) and it cannot have two or more consecutive periods (.) or hyphens (-).For Windows: The host name can be [2, 15] characters in length. It can contain A-Z, a-z, numbers, periods (.), and hyphens (-). It cannot only contain numbers. For other operating systems: The host name can be [2, 64] characters in length. It can be segments separated by periods (.). It can contain A-Z, a-z, numbers, and hyphens (-).
     */
    hostName?: pulumi.Input<string>;
    /**
     * Image ID.
     */
    imageId?: pulumi.Input<string>;
    imageOwnerAlias?: pulumi.Input<string>;
    /**
     * Billing methods. Optional values:
     * - PrePaid: Monthly, or annual subscription. Make sure that your registered credit card is invalid or you have insufficient balance in your PayPal account. Otherwise, InvalidPayMethod error may occur.
     * - PostPaid: Pay-As-You-Go.
     *
     * Default value: PostPaid.
     */
    instanceChargeType?: pulumi.Input<string>;
    /**
     * The name of the instance. The name is a string of 2 to 128 characters. It must begin with an English or a Chinese character. It can contain A-Z, a-z, Chinese characters, numbers, periods (.), colons (:), underscores (_), and hyphens (-).
     */
    instanceName?: pulumi.Input<string>;
    /**
     * Instance type. For more information, call resourceAlicloudInstances to obtain the latest instance type list.
     */
    instanceType?: pulumi.Input<string>;
    /**
     * Internet bandwidth billing method. Optional values: `PayByTraffic` | `PayByBandwidth`.
     */
    internetChargeType?: pulumi.Input<string>;
    /**
     * The maximum inbound bandwidth from the Internet network, measured in Mbit/s. Value range: [1, 200].
     */
    internetMaxBandwidthIn?: pulumi.Input<number>;
    /**
     * Maximum outbound bandwidth from the Internet, its unit of measurement is Mbit/s. Value range: [0, 100].
     */
    internetMaxBandwidthOut?: pulumi.Input<number>;
    /**
     * Whether it is an I/O-optimized instance or not. Optional values:
     * - none
     * - optimized
     */
    ioOptimized?: pulumi.Input<string>;
    /**
     * The name of the key pair.
     * - Ignore this parameter for Windows instances. It is null by default. Even if you enter this parameter, only the  Password content is used.
     * - The password logon method for Linux instances is set to forbidden upon initialization.
     */
    keyPairName?: pulumi.Input<string>;
    launchTemplateName?: pulumi.Input<string>;
    /**
     * The name of the data disk.
     *
     * @deprecated Field 'name' has been deprecated from provider version 1.120.0. New field 'launch_template_name' instead.
     */
    name?: pulumi.Input<string>;
    /**
     * The list of network interfaces created with instance.
     */
    networkInterfaces?: pulumi.Input<inputs.ecs.LaunchTemplateNetworkInterfaces>;
    /**
     * Network type of the instance. Value options: `classic` | `vpc`.
     */
    networkType?: pulumi.Input<string>;
    passwordInherit?: pulumi.Input<boolean>;
    period?: pulumi.Input<number>;
    privateIpAddress?: pulumi.Input<string>;
    /**
     * The RAM role name of the instance. You can use the RAM API ListRoles to query instance RAM role names.
     */
    ramRoleName?: pulumi.Input<string>;
    resourceGroupId?: pulumi.Input<string>;
    /**
     * Whether or not to activate the security enhancement feature and install network security software free of charge. Optional values: Active | Deactive.
     */
    securityEnhancementStrategy?: pulumi.Input<string>;
    /**
     * The security group ID must be one in the same VPC.
     */
    securityGroupId?: pulumi.Input<string>;
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    spotDuration?: pulumi.Input<string>;
    /**
     * Sets the maximum hourly instance price. Supports up to three decimal places.
     */
    spotPriceLimit?: pulumi.Input<number>;
    /**
     * The spot strategy for a Pay-As-You-Go instance. This parameter is valid and required only when InstanceChargeType is set to PostPaid. Value range:
     * - NoSpot: Normal Pay-As-You-Go instance.
     * - SpotWithPriceLimit: Sets the maximum price for a spot instance.
     * - SpotAsPriceGo: The system automatically calculates the price. The maximum value is the Pay-As-You-Go price.
     */
    spotStrategy?: pulumi.Input<string>;
    systemDisk?: pulumi.Input<inputs.ecs.LaunchTemplateSystemDisk>;
    /**
     * The category of the system disk. System disk type. Optional values:
     * - cloud: Basic cloud disk.
     * - cloud_efficiency: Ultra cloud disk.
     * - cloud_ssd: SSD cloud Disks.
     * - ephemeral_ssd: local SSD Disks
     * - cloud_essd: ESSD cloud Disks.
     *
     * @deprecated Field 'system_disk_category' has been deprecated from provider version 1.120.0. New field 'system_disk' instead.
     */
    systemDiskCategory?: pulumi.Input<string>;
    /**
     * System disk description. It cannot begin with http:// or https://.
     *
     * @deprecated Field 'system_disk_description' has been deprecated from provider version 1.120.0. New field 'system_disk' instead.
     */
    systemDiskDescription?: pulumi.Input<string>;
    /**
     * System disk name. The name is a string of 2 to 128 characters. It must begin with an English or a Chinese character. It can contain A-Z, a-z, Chinese characters, numbers, periods (.), colons (:), underscores (_), and hyphens (-).
     *
     * @deprecated Field 'system_disk_name' has been deprecated from provider version 1.120.0. New field 'system_disk' instead.
     */
    systemDiskName?: pulumi.Input<string>;
    /**
     * Size of the system disk, measured in GB. Value range: [20, 500].
     *
     * @deprecated Field 'system_disk_size' has been deprecated from provider version 1.120.0. New field 'system_disk' instead.
     */
    systemDiskSize?: pulumi.Input<number>;
    /**
     * A mapping of tags to assign to the resource.
     * - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    templateResourceGroupId?: pulumi.Input<string>;
    templateTags?: pulumi.Input<{[key: string]: any}>;
    userData?: pulumi.Input<string>;
    /**
     * User data of the instance, which is Base64-encoded. Size of the raw data cannot exceed 16 KB.
     *
     * @deprecated Field 'userdata' has been deprecated from provider version 1.120.0. New field 'user_data' instead.
     */
    userdata?: pulumi.Input<string>;
    versionDescription?: pulumi.Input<string>;
    vpcId?: pulumi.Input<string>;
    /**
     * The VSwitch ID for ENI. The instance must be in the same zone of the same VPC network as the ENI, but they may belong to different VSwitches.
     */
    vswitchId?: pulumi.Input<string>;
    /**
     * The zone ID of the instance.
     */
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LaunchTemplate resource.
 */
export interface LaunchTemplateArgs {
    /**
     * Instance auto release time. The time is presented using the ISO8601 standard and in UTC time. The format is  YYYY-MM-DDTHH:MM:SSZ.
     */
    autoReleaseTime?: pulumi.Input<string>;
    /**
     * The list of data disks created with instance.
     */
    dataDisks?: pulumi.Input<pulumi.Input<inputs.ecs.LaunchTemplateDataDisk>[]>;
    deploymentSetId?: pulumi.Input<string>;
    /**
     * The description of the data disk.
     */
    description?: pulumi.Input<string>;
    enableVmOsConfig?: pulumi.Input<boolean>;
    /**
     * Instance host name.It cannot start or end with a period (.) or a hyphen (-) and it cannot have two or more consecutive periods (.) or hyphens (-).For Windows: The host name can be [2, 15] characters in length. It can contain A-Z, a-z, numbers, periods (.), and hyphens (-). It cannot only contain numbers. For other operating systems: The host name can be [2, 64] characters in length. It can be segments separated by periods (.). It can contain A-Z, a-z, numbers, and hyphens (-).
     */
    hostName?: pulumi.Input<string>;
    /**
     * Image ID.
     */
    imageId?: pulumi.Input<string>;
    imageOwnerAlias?: pulumi.Input<string>;
    /**
     * Billing methods. Optional values:
     * - PrePaid: Monthly, or annual subscription. Make sure that your registered credit card is invalid or you have insufficient balance in your PayPal account. Otherwise, InvalidPayMethod error may occur.
     * - PostPaid: Pay-As-You-Go.
     *
     * Default value: PostPaid.
     */
    instanceChargeType?: pulumi.Input<string>;
    /**
     * The name of the instance. The name is a string of 2 to 128 characters. It must begin with an English or a Chinese character. It can contain A-Z, a-z, Chinese characters, numbers, periods (.), colons (:), underscores (_), and hyphens (-).
     */
    instanceName?: pulumi.Input<string>;
    /**
     * Instance type. For more information, call resourceAlicloudInstances to obtain the latest instance type list.
     */
    instanceType?: pulumi.Input<string>;
    /**
     * Internet bandwidth billing method. Optional values: `PayByTraffic` | `PayByBandwidth`.
     */
    internetChargeType?: pulumi.Input<string>;
    /**
     * The maximum inbound bandwidth from the Internet network, measured in Mbit/s. Value range: [1, 200].
     */
    internetMaxBandwidthIn?: pulumi.Input<number>;
    /**
     * Maximum outbound bandwidth from the Internet, its unit of measurement is Mbit/s. Value range: [0, 100].
     */
    internetMaxBandwidthOut?: pulumi.Input<number>;
    /**
     * Whether it is an I/O-optimized instance or not. Optional values:
     * - none
     * - optimized
     */
    ioOptimized?: pulumi.Input<string>;
    /**
     * The name of the key pair.
     * - Ignore this parameter for Windows instances. It is null by default. Even if you enter this parameter, only the  Password content is used.
     * - The password logon method for Linux instances is set to forbidden upon initialization.
     */
    keyPairName?: pulumi.Input<string>;
    launchTemplateName?: pulumi.Input<string>;
    /**
     * The name of the data disk.
     *
     * @deprecated Field 'name' has been deprecated from provider version 1.120.0. New field 'launch_template_name' instead.
     */
    name?: pulumi.Input<string>;
    /**
     * The list of network interfaces created with instance.
     */
    networkInterfaces?: pulumi.Input<inputs.ecs.LaunchTemplateNetworkInterfaces>;
    /**
     * Network type of the instance. Value options: `classic` | `vpc`.
     */
    networkType?: pulumi.Input<string>;
    passwordInherit?: pulumi.Input<boolean>;
    period?: pulumi.Input<number>;
    privateIpAddress?: pulumi.Input<string>;
    /**
     * The RAM role name of the instance. You can use the RAM API ListRoles to query instance RAM role names.
     */
    ramRoleName?: pulumi.Input<string>;
    resourceGroupId?: pulumi.Input<string>;
    /**
     * Whether or not to activate the security enhancement feature and install network security software free of charge. Optional values: Active | Deactive.
     */
    securityEnhancementStrategy?: pulumi.Input<string>;
    /**
     * The security group ID must be one in the same VPC.
     */
    securityGroupId?: pulumi.Input<string>;
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    spotDuration?: pulumi.Input<string>;
    /**
     * Sets the maximum hourly instance price. Supports up to three decimal places.
     */
    spotPriceLimit?: pulumi.Input<number>;
    /**
     * The spot strategy for a Pay-As-You-Go instance. This parameter is valid and required only when InstanceChargeType is set to PostPaid. Value range:
     * - NoSpot: Normal Pay-As-You-Go instance.
     * - SpotWithPriceLimit: Sets the maximum price for a spot instance.
     * - SpotAsPriceGo: The system automatically calculates the price. The maximum value is the Pay-As-You-Go price.
     */
    spotStrategy?: pulumi.Input<string>;
    systemDisk?: pulumi.Input<inputs.ecs.LaunchTemplateSystemDisk>;
    /**
     * The category of the system disk. System disk type. Optional values:
     * - cloud: Basic cloud disk.
     * - cloud_efficiency: Ultra cloud disk.
     * - cloud_ssd: SSD cloud Disks.
     * - ephemeral_ssd: local SSD Disks
     * - cloud_essd: ESSD cloud Disks.
     *
     * @deprecated Field 'system_disk_category' has been deprecated from provider version 1.120.0. New field 'system_disk' instead.
     */
    systemDiskCategory?: pulumi.Input<string>;
    /**
     * System disk description. It cannot begin with http:// or https://.
     *
     * @deprecated Field 'system_disk_description' has been deprecated from provider version 1.120.0. New field 'system_disk' instead.
     */
    systemDiskDescription?: pulumi.Input<string>;
    /**
     * System disk name. The name is a string of 2 to 128 characters. It must begin with an English or a Chinese character. It can contain A-Z, a-z, Chinese characters, numbers, periods (.), colons (:), underscores (_), and hyphens (-).
     *
     * @deprecated Field 'system_disk_name' has been deprecated from provider version 1.120.0. New field 'system_disk' instead.
     */
    systemDiskName?: pulumi.Input<string>;
    /**
     * Size of the system disk, measured in GB. Value range: [20, 500].
     *
     * @deprecated Field 'system_disk_size' has been deprecated from provider version 1.120.0. New field 'system_disk' instead.
     */
    systemDiskSize?: pulumi.Input<number>;
    /**
     * A mapping of tags to assign to the resource.
     * - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    templateResourceGroupId?: pulumi.Input<string>;
    templateTags?: pulumi.Input<{[key: string]: any}>;
    userData?: pulumi.Input<string>;
    /**
     * User data of the instance, which is Base64-encoded. Size of the raw data cannot exceed 16 KB.
     *
     * @deprecated Field 'userdata' has been deprecated from provider version 1.120.0. New field 'user_data' instead.
     */
    userdata?: pulumi.Input<string>;
    versionDescription?: pulumi.Input<string>;
    vpcId?: pulumi.Input<string>;
    /**
     * The VSwitch ID for ENI. The instance must be in the same zone of the same VPC network as the ENI, but they may belong to different VSwitches.
     */
    vswitchId?: pulumi.Input<string>;
    /**
     * The zone ID of the instance.
     */
    zoneId?: pulumi.Input<string>;
}

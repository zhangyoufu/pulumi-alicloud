// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class ScalingConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing ScalingConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ScalingConfigurationState, opts?: pulumi.CustomResourceOptions): ScalingConfiguration {
        return new ScalingConfiguration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ess/scalingConfiguration:ScalingConfiguration';

    /**
     * Returns true if the given object is an instance of ScalingConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ScalingConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ScalingConfiguration.__pulumiType;
    }

    /**
     * Whether active current scaling configuration in the specified scaling group. Default to `false`.
     */
    public readonly active!: pulumi.Output<boolean>;
    /**
     * DataDisk mappings to attach to ecs instance. See Block datadisk below for details.
     */
    public readonly dataDisks!: pulumi.Output<outputs.ess.ScalingConfigurationDataDisk[] | undefined>;
    /**
     * Whether enable the specified scaling group(make it active) to which the current scaling configuration belongs.
     */
    public readonly enable!: pulumi.Output<boolean | undefined>;
    /**
     * The last scaling configuration will be deleted forcibly with deleting its scaling group. Default to false.
     */
    public readonly forceDelete!: pulumi.Output<boolean | undefined>;
    /**
     * ID of an image file, indicating the image resource selected when an instance is enabled.
     */
    public readonly imageId!: pulumi.Output<string>;
    /**
     * It has been deprecated from version 1.6.0. New resource `alicloud.ess.Attachment` replaces it.
     */
    public readonly instanceIds!: pulumi.Output<string[] | undefined>;
    /**
     * Name of an ECS instance. Default to "ESS-Instance". It is valid from version 1.7.1.
     */
    public readonly instanceName!: pulumi.Output<string | undefined>;
    /**
     * Resource type of an ECS instance.
     */
    public readonly instanceType!: pulumi.Output<string | undefined>;
    /**
     * Resource types of an ECS instance.
     */
    public readonly instanceTypes!: pulumi.Output<string[] | undefined>;
    /**
     * Network billing type, Values: PayByBandwidth or PayByTraffic. Default to `PayByBandwidth`.
     */
    public readonly internetChargeType!: pulumi.Output<string | undefined>;
    /**
     * Maximum incoming bandwidth from the public network, measured in Mbps (Mega bit per second). The value range is [1,200].
     */
    public readonly internetMaxBandwidthIn!: pulumi.Output<number>;
    /**
     * Maximum outgoing bandwidth from the public network, measured in Mbps (Mega bit per second). The value range for PayByBandwidth is [0,100].
     */
    public readonly internetMaxBandwidthOut!: pulumi.Output<number | undefined>;
    /**
     * It has been deprecated on instance resource. All the launched alicloud instances will be I/O optimized.
     */
    public readonly ioOptimized!: pulumi.Output<string | undefined>;
    /**
     * Whether to use outdated instance type. Default to false.
     */
    public readonly isOutdated!: pulumi.Output<boolean | undefined>;
    /**
     * The name of key pair that can login ECS instance successfully without password. If it is specified, the password would be invalid.
     */
    public readonly keyName!: pulumi.Output<string | undefined>;
    /**
     * An KMS encrypts password used to a db account. If the `password` is filled in, this field will be ignored.
     */
    public readonly kmsEncryptedPassword!: pulumi.Output<string | undefined>;
    /**
     * An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a db account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
     */
    public readonly kmsEncryptionContext!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Indicates whether to overwrite the existing data. Default to false.
     */
    public readonly override!: pulumi.Output<boolean | undefined>;
    /**
     * The password of the ECS instance. The password must be 8 to 30 characters in length. It must contains at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `() ~!@#$%^&*-_+=\|{}[]:;'<>,.?/`, The password of Windows-based instances cannot start with a forward slash (/).
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether to use the password that is predefined in the image. If the PasswordInherit parameter is set to true, the `password` and `kmsEncryptedPassword` will be ignored. You must ensure that the selected image has a password configured.
     */
    public readonly passwordInherit!: pulumi.Output<boolean | undefined>;
    /**
     * Instance RAM role name. The name is provided and maintained by RAM. You can use `alicloud.ram.Role` to create a new one.
     */
    public readonly roleName!: pulumi.Output<string | undefined>;
    /**
     * Name shown for the scheduled task. which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is ScalingConfigurationId.
     */
    public readonly scalingConfigurationName!: pulumi.Output<string>;
    /**
     * ID of the scaling group of a scaling configuration.
     */
    public readonly scalingGroupId!: pulumi.Output<string>;
    /**
     * ID of the security group used to create new instance. It is conflict with `securityGroupIds`.
     */
    public readonly securityGroupId!: pulumi.Output<string | undefined>;
    /**
     * List IDs of the security group used to create new instances. It is conflict with `securityGroupId`.
     */
    public readonly securityGroupIds!: pulumi.Output<string[] | undefined>;
    /**
     * The another scaling configuration which will be active automatically and replace current configuration when setting `active` to 'false'. It is invalid when `active` is 'true'.
     */
    public readonly substitute!: pulumi.Output<string>;
    /**
     * Category of the system disk. The parameter value options are `ephemeralSsd`, `cloudEfficiency`, `cloudSsd`, `cloudEssd` and `cloud`. `cloud` only is used to some no I/O optimized instance. Default to `cloudEfficiency`.
     */
    public readonly systemDiskCategory!: pulumi.Output<string | undefined>;
    /**
     * Size of system disk, in GiB. Optional values: cloud: 20-500, cloud_efficiency: 20-500, cloud_ssd: 20-500, ephemeral_ssd: 20-500 The default value is max{40, ImageSize}. If this parameter is set, the system disk size must be greater than or equal to max{40, ImageSize}.
     */
    public readonly systemDiskSize!: pulumi.Output<number | undefined>;
    /**
     * A mapping of tags to assign to the resource. It will be applied for ECS instances finally.
     * - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "http://", or "https://". It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "http://", or "https://" It can be a null string.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * User-defined data to customize the startup behaviors of the ECS instance and to pass data into the ECS instance.
     */
    public readonly userData!: pulumi.Output<string | undefined>;

    /**
     * Create a ScalingConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ScalingConfigurationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ScalingConfigurationArgs | ScalingConfigurationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ScalingConfigurationState | undefined;
            inputs["active"] = state ? state.active : undefined;
            inputs["dataDisks"] = state ? state.dataDisks : undefined;
            inputs["enable"] = state ? state.enable : undefined;
            inputs["forceDelete"] = state ? state.forceDelete : undefined;
            inputs["imageId"] = state ? state.imageId : undefined;
            inputs["instanceIds"] = state ? state.instanceIds : undefined;
            inputs["instanceName"] = state ? state.instanceName : undefined;
            inputs["instanceType"] = state ? state.instanceType : undefined;
            inputs["instanceTypes"] = state ? state.instanceTypes : undefined;
            inputs["internetChargeType"] = state ? state.internetChargeType : undefined;
            inputs["internetMaxBandwidthIn"] = state ? state.internetMaxBandwidthIn : undefined;
            inputs["internetMaxBandwidthOut"] = state ? state.internetMaxBandwidthOut : undefined;
            inputs["ioOptimized"] = state ? state.ioOptimized : undefined;
            inputs["isOutdated"] = state ? state.isOutdated : undefined;
            inputs["keyName"] = state ? state.keyName : undefined;
            inputs["kmsEncryptedPassword"] = state ? state.kmsEncryptedPassword : undefined;
            inputs["kmsEncryptionContext"] = state ? state.kmsEncryptionContext : undefined;
            inputs["override"] = state ? state.override : undefined;
            inputs["password"] = state ? state.password : undefined;
            inputs["passwordInherit"] = state ? state.passwordInherit : undefined;
            inputs["roleName"] = state ? state.roleName : undefined;
            inputs["scalingConfigurationName"] = state ? state.scalingConfigurationName : undefined;
            inputs["scalingGroupId"] = state ? state.scalingGroupId : undefined;
            inputs["securityGroupId"] = state ? state.securityGroupId : undefined;
            inputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            inputs["substitute"] = state ? state.substitute : undefined;
            inputs["systemDiskCategory"] = state ? state.systemDiskCategory : undefined;
            inputs["systemDiskSize"] = state ? state.systemDiskSize : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["userData"] = state ? state.userData : undefined;
        } else {
            const args = argsOrState as ScalingConfigurationArgs | undefined;
            if (!args || args.imageId === undefined) {
                throw new Error("Missing required property 'imageId'");
            }
            if (!args || args.scalingGroupId === undefined) {
                throw new Error("Missing required property 'scalingGroupId'");
            }
            inputs["active"] = args ? args.active : undefined;
            inputs["dataDisks"] = args ? args.dataDisks : undefined;
            inputs["enable"] = args ? args.enable : undefined;
            inputs["forceDelete"] = args ? args.forceDelete : undefined;
            inputs["imageId"] = args ? args.imageId : undefined;
            inputs["instanceIds"] = args ? args.instanceIds : undefined;
            inputs["instanceName"] = args ? args.instanceName : undefined;
            inputs["instanceType"] = args ? args.instanceType : undefined;
            inputs["instanceTypes"] = args ? args.instanceTypes : undefined;
            inputs["internetChargeType"] = args ? args.internetChargeType : undefined;
            inputs["internetMaxBandwidthIn"] = args ? args.internetMaxBandwidthIn : undefined;
            inputs["internetMaxBandwidthOut"] = args ? args.internetMaxBandwidthOut : undefined;
            inputs["ioOptimized"] = args ? args.ioOptimized : undefined;
            inputs["isOutdated"] = args ? args.isOutdated : undefined;
            inputs["keyName"] = args ? args.keyName : undefined;
            inputs["kmsEncryptedPassword"] = args ? args.kmsEncryptedPassword : undefined;
            inputs["kmsEncryptionContext"] = args ? args.kmsEncryptionContext : undefined;
            inputs["override"] = args ? args.override : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["passwordInherit"] = args ? args.passwordInherit : undefined;
            inputs["roleName"] = args ? args.roleName : undefined;
            inputs["scalingConfigurationName"] = args ? args.scalingConfigurationName : undefined;
            inputs["scalingGroupId"] = args ? args.scalingGroupId : undefined;
            inputs["securityGroupId"] = args ? args.securityGroupId : undefined;
            inputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            inputs["substitute"] = args ? args.substitute : undefined;
            inputs["systemDiskCategory"] = args ? args.systemDiskCategory : undefined;
            inputs["systemDiskSize"] = args ? args.systemDiskSize : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["userData"] = args ? args.userData : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ScalingConfiguration.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ScalingConfiguration resources.
 */
export interface ScalingConfigurationState {
    /**
     * Whether active current scaling configuration in the specified scaling group. Default to `false`.
     */
    readonly active?: pulumi.Input<boolean>;
    /**
     * DataDisk mappings to attach to ecs instance. See Block datadisk below for details.
     */
    readonly dataDisks?: pulumi.Input<pulumi.Input<inputs.ess.ScalingConfigurationDataDisk>[]>;
    /**
     * Whether enable the specified scaling group(make it active) to which the current scaling configuration belongs.
     */
    readonly enable?: pulumi.Input<boolean>;
    /**
     * The last scaling configuration will be deleted forcibly with deleting its scaling group. Default to false.
     */
    readonly forceDelete?: pulumi.Input<boolean>;
    /**
     * ID of an image file, indicating the image resource selected when an instance is enabled.
     */
    readonly imageId?: pulumi.Input<string>;
    /**
     * It has been deprecated from version 1.6.0. New resource `alicloud.ess.Attachment` replaces it.
     * @deprecated Field 'instance_ids' has been deprecated from provider version 1.6.0. New resource 'alicloud_ess_attachment' replaces it.
     */
    readonly instanceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of an ECS instance. Default to "ESS-Instance". It is valid from version 1.7.1.
     */
    readonly instanceName?: pulumi.Input<string>;
    /**
     * Resource type of an ECS instance.
     */
    readonly instanceType?: pulumi.Input<string>;
    /**
     * Resource types of an ECS instance.
     */
    readonly instanceTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Network billing type, Values: PayByBandwidth or PayByTraffic. Default to `PayByBandwidth`.
     */
    readonly internetChargeType?: pulumi.Input<string>;
    /**
     * Maximum incoming bandwidth from the public network, measured in Mbps (Mega bit per second). The value range is [1,200].
     */
    readonly internetMaxBandwidthIn?: pulumi.Input<number>;
    /**
     * Maximum outgoing bandwidth from the public network, measured in Mbps (Mega bit per second). The value range for PayByBandwidth is [0,100].
     */
    readonly internetMaxBandwidthOut?: pulumi.Input<number>;
    /**
     * It has been deprecated on instance resource. All the launched alicloud instances will be I/O optimized.
     * @deprecated Attribute io_optimized has been deprecated on instance resource. All the launched alicloud instances will be IO optimized. Suggest to remove it from your template.
     */
    readonly ioOptimized?: pulumi.Input<string>;
    /**
     * Whether to use outdated instance type. Default to false.
     */
    readonly isOutdated?: pulumi.Input<boolean>;
    /**
     * The name of key pair that can login ECS instance successfully without password. If it is specified, the password would be invalid.
     */
    readonly keyName?: pulumi.Input<string>;
    /**
     * An KMS encrypts password used to a db account. If the `password` is filled in, this field will be ignored.
     */
    readonly kmsEncryptedPassword?: pulumi.Input<string>;
    /**
     * An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a db account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
     */
    readonly kmsEncryptionContext?: pulumi.Input<{[key: string]: any}>;
    /**
     * Indicates whether to overwrite the existing data. Default to false.
     */
    readonly override?: pulumi.Input<boolean>;
    /**
     * The password of the ECS instance. The password must be 8 to 30 characters in length. It must contains at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `() ~!@#$%^&*-_+=\|{}[]:;'<>,.?/`, The password of Windows-based instances cannot start with a forward slash (/).
     */
    readonly password?: pulumi.Input<string>;
    /**
     * Specifies whether to use the password that is predefined in the image. If the PasswordInherit parameter is set to true, the `password` and `kmsEncryptedPassword` will be ignored. You must ensure that the selected image has a password configured.
     */
    readonly passwordInherit?: pulumi.Input<boolean>;
    /**
     * Instance RAM role name. The name is provided and maintained by RAM. You can use `alicloud.ram.Role` to create a new one.
     */
    readonly roleName?: pulumi.Input<string>;
    /**
     * Name shown for the scheduled task. which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is ScalingConfigurationId.
     */
    readonly scalingConfigurationName?: pulumi.Input<string>;
    /**
     * ID of the scaling group of a scaling configuration.
     */
    readonly scalingGroupId?: pulumi.Input<string>;
    /**
     * ID of the security group used to create new instance. It is conflict with `securityGroupIds`.
     */
    readonly securityGroupId?: pulumi.Input<string>;
    /**
     * List IDs of the security group used to create new instances. It is conflict with `securityGroupId`.
     */
    readonly securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The another scaling configuration which will be active automatically and replace current configuration when setting `active` to 'false'. It is invalid when `active` is 'true'.
     */
    readonly substitute?: pulumi.Input<string>;
    /**
     * Category of the system disk. The parameter value options are `ephemeralSsd`, `cloudEfficiency`, `cloudSsd`, `cloudEssd` and `cloud`. `cloud` only is used to some no I/O optimized instance. Default to `cloudEfficiency`.
     */
    readonly systemDiskCategory?: pulumi.Input<string>;
    /**
     * Size of system disk, in GiB. Optional values: cloud: 20-500, cloud_efficiency: 20-500, cloud_ssd: 20-500, ephemeral_ssd: 20-500 The default value is max{40, ImageSize}. If this parameter is set, the system disk size must be greater than or equal to max{40, ImageSize}.
     */
    readonly systemDiskSize?: pulumi.Input<number>;
    /**
     * A mapping of tags to assign to the resource. It will be applied for ECS instances finally.
     * - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "http://", or "https://". It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "http://", or "https://" It can be a null string.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * User-defined data to customize the startup behaviors of the ECS instance and to pass data into the ECS instance.
     */
    readonly userData?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ScalingConfiguration resource.
 */
export interface ScalingConfigurationArgs {
    /**
     * Whether active current scaling configuration in the specified scaling group. Default to `false`.
     */
    readonly active?: pulumi.Input<boolean>;
    /**
     * DataDisk mappings to attach to ecs instance. See Block datadisk below for details.
     */
    readonly dataDisks?: pulumi.Input<pulumi.Input<inputs.ess.ScalingConfigurationDataDisk>[]>;
    /**
     * Whether enable the specified scaling group(make it active) to which the current scaling configuration belongs.
     */
    readonly enable?: pulumi.Input<boolean>;
    /**
     * The last scaling configuration will be deleted forcibly with deleting its scaling group. Default to false.
     */
    readonly forceDelete?: pulumi.Input<boolean>;
    /**
     * ID of an image file, indicating the image resource selected when an instance is enabled.
     */
    readonly imageId: pulumi.Input<string>;
    /**
     * It has been deprecated from version 1.6.0. New resource `alicloud.ess.Attachment` replaces it.
     * @deprecated Field 'instance_ids' has been deprecated from provider version 1.6.0. New resource 'alicloud_ess_attachment' replaces it.
     */
    readonly instanceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of an ECS instance. Default to "ESS-Instance". It is valid from version 1.7.1.
     */
    readonly instanceName?: pulumi.Input<string>;
    /**
     * Resource type of an ECS instance.
     */
    readonly instanceType?: pulumi.Input<string>;
    /**
     * Resource types of an ECS instance.
     */
    readonly instanceTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Network billing type, Values: PayByBandwidth or PayByTraffic. Default to `PayByBandwidth`.
     */
    readonly internetChargeType?: pulumi.Input<string>;
    /**
     * Maximum incoming bandwidth from the public network, measured in Mbps (Mega bit per second). The value range is [1,200].
     */
    readonly internetMaxBandwidthIn?: pulumi.Input<number>;
    /**
     * Maximum outgoing bandwidth from the public network, measured in Mbps (Mega bit per second). The value range for PayByBandwidth is [0,100].
     */
    readonly internetMaxBandwidthOut?: pulumi.Input<number>;
    /**
     * It has been deprecated on instance resource. All the launched alicloud instances will be I/O optimized.
     * @deprecated Attribute io_optimized has been deprecated on instance resource. All the launched alicloud instances will be IO optimized. Suggest to remove it from your template.
     */
    readonly ioOptimized?: pulumi.Input<string>;
    /**
     * Whether to use outdated instance type. Default to false.
     */
    readonly isOutdated?: pulumi.Input<boolean>;
    /**
     * The name of key pair that can login ECS instance successfully without password. If it is specified, the password would be invalid.
     */
    readonly keyName?: pulumi.Input<string>;
    /**
     * An KMS encrypts password used to a db account. If the `password` is filled in, this field will be ignored.
     */
    readonly kmsEncryptedPassword?: pulumi.Input<string>;
    /**
     * An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a db account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
     */
    readonly kmsEncryptionContext?: pulumi.Input<{[key: string]: any}>;
    /**
     * Indicates whether to overwrite the existing data. Default to false.
     */
    readonly override?: pulumi.Input<boolean>;
    /**
     * The password of the ECS instance. The password must be 8 to 30 characters in length. It must contains at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `() ~!@#$%^&*-_+=\|{}[]:;'<>,.?/`, The password of Windows-based instances cannot start with a forward slash (/).
     */
    readonly password?: pulumi.Input<string>;
    /**
     * Specifies whether to use the password that is predefined in the image. If the PasswordInherit parameter is set to true, the `password` and `kmsEncryptedPassword` will be ignored. You must ensure that the selected image has a password configured.
     */
    readonly passwordInherit?: pulumi.Input<boolean>;
    /**
     * Instance RAM role name. The name is provided and maintained by RAM. You can use `alicloud.ram.Role` to create a new one.
     */
    readonly roleName?: pulumi.Input<string>;
    /**
     * Name shown for the scheduled task. which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is ScalingConfigurationId.
     */
    readonly scalingConfigurationName?: pulumi.Input<string>;
    /**
     * ID of the scaling group of a scaling configuration.
     */
    readonly scalingGroupId: pulumi.Input<string>;
    /**
     * ID of the security group used to create new instance. It is conflict with `securityGroupIds`.
     */
    readonly securityGroupId?: pulumi.Input<string>;
    /**
     * List IDs of the security group used to create new instances. It is conflict with `securityGroupId`.
     */
    readonly securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The another scaling configuration which will be active automatically and replace current configuration when setting `active` to 'false'. It is invalid when `active` is 'true'.
     */
    readonly substitute?: pulumi.Input<string>;
    /**
     * Category of the system disk. The parameter value options are `ephemeralSsd`, `cloudEfficiency`, `cloudSsd`, `cloudEssd` and `cloud`. `cloud` only is used to some no I/O optimized instance. Default to `cloudEfficiency`.
     */
    readonly systemDiskCategory?: pulumi.Input<string>;
    /**
     * Size of system disk, in GiB. Optional values: cloud: 20-500, cloud_efficiency: 20-500, cloud_ssd: 20-500, ephemeral_ssd: 20-500 The default value is max{40, ImageSize}. If this parameter is set, the system disk size must be greater than or equal to max{40, ImageSize}.
     */
    readonly systemDiskSize?: pulumi.Input<number>;
    /**
     * A mapping of tags to assign to the resource. It will be applied for ECS instances finally.
     * - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "http://", or "https://". It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "http://", or "https://" It can be a null string.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * User-defined data to customize the startup behaviors of the ECS instance and to pass data into the ECS instance.
     */
    readonly userData?: pulumi.Input<string>;
}

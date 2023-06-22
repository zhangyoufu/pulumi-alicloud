// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a ESS eci scaling configuration resource.
 *
 * For information about ess eci scaling configuration, see [CreateEciScalingConfiguration](https://www.alibabacloud.com/help/en/auto-scaling/latest/create-eci-scaling-configuration).
 *
 * > **NOTE:** Available since v1.164.0.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "terraform-example";
 * const defaultZones = alicloud.getZones({
 *     availableDiskCategory: "cloud_efficiency",
 *     availableResourceCreation: "VSwitch",
 * });
 * const defaultNetwork = new alicloud.vpc.Network("defaultNetwork", {
 *     vpcName: name,
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("defaultSwitch", {
 *     vpcId: defaultNetwork.id,
 *     cidrBlock: "172.16.0.0/24",
 *     zoneId: defaultZones.then(defaultZones => defaultZones.zones?.[0]?.id),
 *     vswitchName: name,
 * });
 * const defaultSecurityGroup = new alicloud.ecs.SecurityGroup("defaultSecurityGroup", {vpcId: defaultNetwork.id});
 * const defaultScalingGroup = new alicloud.ess.ScalingGroup("defaultScalingGroup", {
 *     minSize: 0,
 *     maxSize: 1,
 *     scalingGroupName: name,
 *     removalPolicies: [
 *         "OldestInstance",
 *         "NewestInstance",
 *     ],
 *     vswitchIds: [defaultSwitch.id],
 *     groupType: "ECI",
 * });
 * const defaultEciScalingConfiguration = new alicloud.ess.EciScalingConfiguration("defaultEciScalingConfiguration", {
 *     scalingGroupId: defaultScalingGroup.id,
 *     cpu: 2,
 *     memory: 4,
 *     securityGroupId: defaultSecurityGroup.id,
 *     forceDelete: true,
 *     active: true,
 *     containerGroupName: "container-group-1649839595174",
 *     containers: [{
 *         name: "container-1",
 *         image: "registry-vpc.cn-hangzhou.aliyuncs.com/eci_open/alpine:3.5",
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * ESS eci scaling configuration can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:ess/eciScalingConfiguration:EciScalingConfiguration example asc-abc123456
 * ```
 */
export class EciScalingConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing EciScalingConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EciScalingConfigurationState, opts?: pulumi.CustomResourceOptions): EciScalingConfiguration {
        return new EciScalingConfiguration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ess/eciScalingConfiguration:EciScalingConfiguration';

    /**
     * Returns true if the given object is an instance of EciScalingConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EciScalingConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EciScalingConfiguration.__pulumiType;
    }

    /**
     * Information about the Container Registry Enterprise Edition instance. See `acrRegistryInfos` below for details.
     */
    public readonly acrRegistryInfos!: pulumi.Output<outputs.ess.EciScalingConfigurationAcrRegistryInfo[] | undefined>;
    /**
     * Whether active current eci scaling configuration in the specified scaling group. Note that only
     * one configuration can be active. Default to `false`.
     */
    public readonly active!: pulumi.Output<boolean | undefined>;
    /**
     * Whether create eip automatically.
     */
    public readonly autoCreateEip!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the container group.
     */
    public readonly containerGroupName!: pulumi.Output<string | undefined>;
    /**
     * The list of containers. See `containers` below for details.
     */
    public readonly containers!: pulumi.Output<outputs.ess.EciScalingConfigurationContainer[] | undefined>;
    /**
     * The amount of CPU resources allocated to the container group.
     */
    public readonly cpu!: pulumi.Output<number | undefined>;
    /**
     * The description of data disk N. Valid values of N: 1 to 16. The description must be 2 to
     * 256 characters in length and cannot start with http:// or https://.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * dns policy of contain group.
     */
    public readonly dnsPolicy!: pulumi.Output<string | undefined>;
    /**
     * egress bandwidth.
     */
    public readonly egressBandwidth!: pulumi.Output<number | undefined>;
    /**
     * Eip bandwidth.
     */
    public readonly eipBandwidth!: pulumi.Output<number | undefined>;
    /**
     * Enable sls log service.
     */
    public readonly enableSls!: pulumi.Output<boolean | undefined>;
    /**
     * The eci scaling configuration will be deleted forcibly with deleting its scaling group.
     * Default to false.
     */
    public readonly forceDelete!: pulumi.Output<boolean | undefined>;
    /**
     * HostAliases. See `hostAliases` below.
     */
    public readonly hostAliases!: pulumi.Output<outputs.ess.EciScalingConfigurationHostAlias[] | undefined>;
    /**
     * Hostname of an ECI instance.
     */
    public readonly hostName!: pulumi.Output<string | undefined>;
    /**
     * The image registry credential.   See `imageRegistryCredentials` below for
     * details.
     */
    public readonly imageRegistryCredentials!: pulumi.Output<outputs.ess.EciScalingConfigurationImageRegistryCredential[] | undefined>;
    /**
     * Ingress bandwidth.
     */
    public readonly ingressBandwidth!: pulumi.Output<number | undefined>;
    /**
     * The list of initContainers. See `initContainers` below for details.
     */
    public readonly initContainers!: pulumi.Output<outputs.ess.EciScalingConfigurationInitContainer[] | undefined>;
    /**
     * The amount of memory resources allocated to the container group.
     */
    public readonly memory!: pulumi.Output<number | undefined>;
    /**
     * The RAM role that the container group assumes. ECI and ECS share the same RAM role.
     */
    public readonly ramRoleName!: pulumi.Output<string | undefined>;
    /**
     * ID of resource group.
     */
    public readonly resourceGroupId!: pulumi.Output<string | undefined>;
    /**
     * The restart policy of the container group. Default to `Always`.
     */
    public readonly restartPolicy!: pulumi.Output<string | undefined>;
    /**
     * Name shown for the scheduled task. which must contain 2-64 characters (
     * English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number,
     * underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is
     * EciScalingConfigurationId.
     */
    public readonly scalingConfigurationName!: pulumi.Output<string>;
    /**
     * ID of the scaling group of a eci scaling configuration.
     */
    public readonly scalingGroupId!: pulumi.Output<string>;
    /**
     * ID of the security group used to create new instance. It is conflict
     * with `securityGroupIds`.
     */
    public readonly securityGroupId!: pulumi.Output<string | undefined>;
    /**
     * The maximum price hourly for spot instance.
     */
    public readonly spotPriceLimit!: pulumi.Output<number | undefined>;
    /**
     * The spot strategy for a Pay-As-You-Go instance. Valid values: `NoSpot`, `SpotAsPriceGo`
     * , `SpotWithPriceLimit`.
     */
    public readonly spotStrategy!: pulumi.Output<string | undefined>;
    /**
     * A mapping of tags to assign to the resource. It will be applied for ECI instances finally.
     * - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "http://", or "https://". It cannot
     * be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "http://", or "https://" It can be
     * a null string.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The list of volumes. See `volumes` below for details.
     */
    public readonly volumes!: pulumi.Output<outputs.ess.EciScalingConfigurationVolume[] | undefined>;

    /**
     * Create a EciScalingConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EciScalingConfigurationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EciScalingConfigurationArgs | EciScalingConfigurationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EciScalingConfigurationState | undefined;
            resourceInputs["acrRegistryInfos"] = state ? state.acrRegistryInfos : undefined;
            resourceInputs["active"] = state ? state.active : undefined;
            resourceInputs["autoCreateEip"] = state ? state.autoCreateEip : undefined;
            resourceInputs["containerGroupName"] = state ? state.containerGroupName : undefined;
            resourceInputs["containers"] = state ? state.containers : undefined;
            resourceInputs["cpu"] = state ? state.cpu : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["dnsPolicy"] = state ? state.dnsPolicy : undefined;
            resourceInputs["egressBandwidth"] = state ? state.egressBandwidth : undefined;
            resourceInputs["eipBandwidth"] = state ? state.eipBandwidth : undefined;
            resourceInputs["enableSls"] = state ? state.enableSls : undefined;
            resourceInputs["forceDelete"] = state ? state.forceDelete : undefined;
            resourceInputs["hostAliases"] = state ? state.hostAliases : undefined;
            resourceInputs["hostName"] = state ? state.hostName : undefined;
            resourceInputs["imageRegistryCredentials"] = state ? state.imageRegistryCredentials : undefined;
            resourceInputs["ingressBandwidth"] = state ? state.ingressBandwidth : undefined;
            resourceInputs["initContainers"] = state ? state.initContainers : undefined;
            resourceInputs["memory"] = state ? state.memory : undefined;
            resourceInputs["ramRoleName"] = state ? state.ramRoleName : undefined;
            resourceInputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            resourceInputs["restartPolicy"] = state ? state.restartPolicy : undefined;
            resourceInputs["scalingConfigurationName"] = state ? state.scalingConfigurationName : undefined;
            resourceInputs["scalingGroupId"] = state ? state.scalingGroupId : undefined;
            resourceInputs["securityGroupId"] = state ? state.securityGroupId : undefined;
            resourceInputs["spotPriceLimit"] = state ? state.spotPriceLimit : undefined;
            resourceInputs["spotStrategy"] = state ? state.spotStrategy : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["volumes"] = state ? state.volumes : undefined;
        } else {
            const args = argsOrState as EciScalingConfigurationArgs | undefined;
            if ((!args || args.scalingGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scalingGroupId'");
            }
            resourceInputs["acrRegistryInfos"] = args ? args.acrRegistryInfos : undefined;
            resourceInputs["active"] = args ? args.active : undefined;
            resourceInputs["autoCreateEip"] = args ? args.autoCreateEip : undefined;
            resourceInputs["containerGroupName"] = args ? args.containerGroupName : undefined;
            resourceInputs["containers"] = args ? args.containers : undefined;
            resourceInputs["cpu"] = args ? args.cpu : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dnsPolicy"] = args ? args.dnsPolicy : undefined;
            resourceInputs["egressBandwidth"] = args ? args.egressBandwidth : undefined;
            resourceInputs["eipBandwidth"] = args ? args.eipBandwidth : undefined;
            resourceInputs["enableSls"] = args ? args.enableSls : undefined;
            resourceInputs["forceDelete"] = args ? args.forceDelete : undefined;
            resourceInputs["hostAliases"] = args ? args.hostAliases : undefined;
            resourceInputs["hostName"] = args ? args.hostName : undefined;
            resourceInputs["imageRegistryCredentials"] = args ? args.imageRegistryCredentials : undefined;
            resourceInputs["ingressBandwidth"] = args ? args.ingressBandwidth : undefined;
            resourceInputs["initContainers"] = args ? args.initContainers : undefined;
            resourceInputs["memory"] = args ? args.memory : undefined;
            resourceInputs["ramRoleName"] = args ? args.ramRoleName : undefined;
            resourceInputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            resourceInputs["restartPolicy"] = args ? args.restartPolicy : undefined;
            resourceInputs["scalingConfigurationName"] = args ? args.scalingConfigurationName : undefined;
            resourceInputs["scalingGroupId"] = args ? args.scalingGroupId : undefined;
            resourceInputs["securityGroupId"] = args ? args.securityGroupId : undefined;
            resourceInputs["spotPriceLimit"] = args ? args.spotPriceLimit : undefined;
            resourceInputs["spotStrategy"] = args ? args.spotStrategy : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["volumes"] = args ? args.volumes : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EciScalingConfiguration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EciScalingConfiguration resources.
 */
export interface EciScalingConfigurationState {
    /**
     * Information about the Container Registry Enterprise Edition instance. See `acrRegistryInfos` below for details.
     */
    acrRegistryInfos?: pulumi.Input<pulumi.Input<inputs.ess.EciScalingConfigurationAcrRegistryInfo>[]>;
    /**
     * Whether active current eci scaling configuration in the specified scaling group. Note that only
     * one configuration can be active. Default to `false`.
     */
    active?: pulumi.Input<boolean>;
    /**
     * Whether create eip automatically.
     */
    autoCreateEip?: pulumi.Input<boolean>;
    /**
     * The name of the container group.
     */
    containerGroupName?: pulumi.Input<string>;
    /**
     * The list of containers. See `containers` below for details.
     */
    containers?: pulumi.Input<pulumi.Input<inputs.ess.EciScalingConfigurationContainer>[]>;
    /**
     * The amount of CPU resources allocated to the container group.
     */
    cpu?: pulumi.Input<number>;
    /**
     * The description of data disk N. Valid values of N: 1 to 16. The description must be 2 to
     * 256 characters in length and cannot start with http:// or https://.
     */
    description?: pulumi.Input<string>;
    /**
     * dns policy of contain group.
     */
    dnsPolicy?: pulumi.Input<string>;
    /**
     * egress bandwidth.
     */
    egressBandwidth?: pulumi.Input<number>;
    /**
     * Eip bandwidth.
     */
    eipBandwidth?: pulumi.Input<number>;
    /**
     * Enable sls log service.
     */
    enableSls?: pulumi.Input<boolean>;
    /**
     * The eci scaling configuration will be deleted forcibly with deleting its scaling group.
     * Default to false.
     */
    forceDelete?: pulumi.Input<boolean>;
    /**
     * HostAliases. See `hostAliases` below.
     */
    hostAliases?: pulumi.Input<pulumi.Input<inputs.ess.EciScalingConfigurationHostAlias>[]>;
    /**
     * Hostname of an ECI instance.
     */
    hostName?: pulumi.Input<string>;
    /**
     * The image registry credential.   See `imageRegistryCredentials` below for
     * details.
     */
    imageRegistryCredentials?: pulumi.Input<pulumi.Input<inputs.ess.EciScalingConfigurationImageRegistryCredential>[]>;
    /**
     * Ingress bandwidth.
     */
    ingressBandwidth?: pulumi.Input<number>;
    /**
     * The list of initContainers. See `initContainers` below for details.
     */
    initContainers?: pulumi.Input<pulumi.Input<inputs.ess.EciScalingConfigurationInitContainer>[]>;
    /**
     * The amount of memory resources allocated to the container group.
     */
    memory?: pulumi.Input<number>;
    /**
     * The RAM role that the container group assumes. ECI and ECS share the same RAM role.
     */
    ramRoleName?: pulumi.Input<string>;
    /**
     * ID of resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The restart policy of the container group. Default to `Always`.
     */
    restartPolicy?: pulumi.Input<string>;
    /**
     * Name shown for the scheduled task. which must contain 2-64 characters (
     * English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number,
     * underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is
     * EciScalingConfigurationId.
     */
    scalingConfigurationName?: pulumi.Input<string>;
    /**
     * ID of the scaling group of a eci scaling configuration.
     */
    scalingGroupId?: pulumi.Input<string>;
    /**
     * ID of the security group used to create new instance. It is conflict
     * with `securityGroupIds`.
     */
    securityGroupId?: pulumi.Input<string>;
    /**
     * The maximum price hourly for spot instance.
     */
    spotPriceLimit?: pulumi.Input<number>;
    /**
     * The spot strategy for a Pay-As-You-Go instance. Valid values: `NoSpot`, `SpotAsPriceGo`
     * , `SpotWithPriceLimit`.
     */
    spotStrategy?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource. It will be applied for ECI instances finally.
     * - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "http://", or "https://". It cannot
     * be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "http://", or "https://" It can be
     * a null string.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The list of volumes. See `volumes` below for details.
     */
    volumes?: pulumi.Input<pulumi.Input<inputs.ess.EciScalingConfigurationVolume>[]>;
}

/**
 * The set of arguments for constructing a EciScalingConfiguration resource.
 */
export interface EciScalingConfigurationArgs {
    /**
     * Information about the Container Registry Enterprise Edition instance. See `acrRegistryInfos` below for details.
     */
    acrRegistryInfos?: pulumi.Input<pulumi.Input<inputs.ess.EciScalingConfigurationAcrRegistryInfo>[]>;
    /**
     * Whether active current eci scaling configuration in the specified scaling group. Note that only
     * one configuration can be active. Default to `false`.
     */
    active?: pulumi.Input<boolean>;
    /**
     * Whether create eip automatically.
     */
    autoCreateEip?: pulumi.Input<boolean>;
    /**
     * The name of the container group.
     */
    containerGroupName?: pulumi.Input<string>;
    /**
     * The list of containers. See `containers` below for details.
     */
    containers?: pulumi.Input<pulumi.Input<inputs.ess.EciScalingConfigurationContainer>[]>;
    /**
     * The amount of CPU resources allocated to the container group.
     */
    cpu?: pulumi.Input<number>;
    /**
     * The description of data disk N. Valid values of N: 1 to 16. The description must be 2 to
     * 256 characters in length and cannot start with http:// or https://.
     */
    description?: pulumi.Input<string>;
    /**
     * dns policy of contain group.
     */
    dnsPolicy?: pulumi.Input<string>;
    /**
     * egress bandwidth.
     */
    egressBandwidth?: pulumi.Input<number>;
    /**
     * Eip bandwidth.
     */
    eipBandwidth?: pulumi.Input<number>;
    /**
     * Enable sls log service.
     */
    enableSls?: pulumi.Input<boolean>;
    /**
     * The eci scaling configuration will be deleted forcibly with deleting its scaling group.
     * Default to false.
     */
    forceDelete?: pulumi.Input<boolean>;
    /**
     * HostAliases. See `hostAliases` below.
     */
    hostAliases?: pulumi.Input<pulumi.Input<inputs.ess.EciScalingConfigurationHostAlias>[]>;
    /**
     * Hostname of an ECI instance.
     */
    hostName?: pulumi.Input<string>;
    /**
     * The image registry credential.   See `imageRegistryCredentials` below for
     * details.
     */
    imageRegistryCredentials?: pulumi.Input<pulumi.Input<inputs.ess.EciScalingConfigurationImageRegistryCredential>[]>;
    /**
     * Ingress bandwidth.
     */
    ingressBandwidth?: pulumi.Input<number>;
    /**
     * The list of initContainers. See `initContainers` below for details.
     */
    initContainers?: pulumi.Input<pulumi.Input<inputs.ess.EciScalingConfigurationInitContainer>[]>;
    /**
     * The amount of memory resources allocated to the container group.
     */
    memory?: pulumi.Input<number>;
    /**
     * The RAM role that the container group assumes. ECI and ECS share the same RAM role.
     */
    ramRoleName?: pulumi.Input<string>;
    /**
     * ID of resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The restart policy of the container group. Default to `Always`.
     */
    restartPolicy?: pulumi.Input<string>;
    /**
     * Name shown for the scheduled task. which must contain 2-64 characters (
     * English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number,
     * underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is
     * EciScalingConfigurationId.
     */
    scalingConfigurationName?: pulumi.Input<string>;
    /**
     * ID of the scaling group of a eci scaling configuration.
     */
    scalingGroupId: pulumi.Input<string>;
    /**
     * ID of the security group used to create new instance. It is conflict
     * with `securityGroupIds`.
     */
    securityGroupId?: pulumi.Input<string>;
    /**
     * The maximum price hourly for spot instance.
     */
    spotPriceLimit?: pulumi.Input<number>;
    /**
     * The spot strategy for a Pay-As-You-Go instance. Valid values: `NoSpot`, `SpotAsPriceGo`
     * , `SpotWithPriceLimit`.
     */
    spotStrategy?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource. It will be applied for ECI instances finally.
     * - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "http://", or "https://". It cannot
     * be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "http://", or "https://" It can be
     * a null string.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The list of volumes. See `volumes` below for details.
     */
    volumes?: pulumi.Input<pulumi.Input<inputs.ess.EciScalingConfigurationVolume>[]>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a ECS Image Pipeline resource.
 *
 * For information about ECS Image Pipeline and how to use it, see [What is Image Pipeline](https://www.alibabacloud.com/help/en/doc-detail/200427.html).
 *
 * > **NOTE:** Available in v1.163.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const default = alicloud.resourcemanager.getResourceGroups({
 *     nameRegex: "default",
 * });
 * const defaultGetZones = alicloud.getZones({
 *     availableResourceCreation: "VSwitch",
 * });
 * const defaultGetImages = alicloud.ecs.getImages({
 *     nameRegex: "^ubuntu_18.*64",
 *     mostRecent: true,
 *     owners: "system",
 * });
 * const defaultGetInstanceTypes = defaultGetImages.then(defaultGetImages => alicloud.ecs.getInstanceTypes({
 *     imageId: defaultGetImages.ids?.[0],
 * }));
 * const defaultGetAccount = alicloud.getAccount({});
 * const defaultNetwork = new alicloud.vpc.Network("default", {
 *     vpcName: "terraform-example",
 *     cidrBlock: "172.17.3.0/24",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("default", {
 *     vswitchName: "terraform-example",
 *     cidrBlock: "172.17.3.0/24",
 *     vpcId: defaultNetwork.id,
 *     zoneId: defaultGetZones.then(defaultGetZones => defaultGetZones.zones?.[0]?.id),
 * });
 * const defaultEcsImagePipeline = new alicloud.ecs.EcsImagePipeline("default", {
 *     addAccounts: [defaultGetAccount.then(defaultGetAccount => defaultGetAccount.id)],
 *     baseImage: defaultGetImages.then(defaultGetImages => defaultGetImages.ids?.[0]),
 *     baseImageType: "IMAGE",
 *     buildContent: "RUN yum update -y",
 *     deleteInstanceOnFailure: false,
 *     imageName: "terraform-example",
 *     name: "terraform-example",
 *     description: "terraform-example",
 *     instanceType: defaultGetInstanceTypes.then(defaultGetInstanceTypes => defaultGetInstanceTypes.ids?.[0]),
 *     resourceGroupId: _default.then(_default => _default.groups?.[0]?.id),
 *     internetMaxBandwidthOut: 20,
 *     systemDiskSize: 40,
 *     toRegionIds: [
 *         "cn-qingdao",
 *         "cn-zhangjiakou",
 *     ],
 *     vswitchId: defaultSwitch.id,
 *     tags: {
 *         Created: "TF",
 *         For: "Acceptance-test",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * ECS Image Pipeline can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:ecs/ecsImagePipeline:EcsImagePipeline example <id>
 * ```
 */
export class EcsImagePipeline extends pulumi.CustomResource {
    /**
     * Get an existing EcsImagePipeline resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EcsImagePipelineState, opts?: pulumi.CustomResourceOptions): EcsImagePipeline {
        return new EcsImagePipeline(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ecs/ecsImagePipeline:EcsImagePipeline';

    /**
     * Returns true if the given object is an instance of EcsImagePipeline.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EcsImagePipeline {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EcsImagePipeline.__pulumiType;
    }

    /**
     * The ID of Alibaba Cloud account to which to share the created image.
     */
    public readonly addAccounts!: pulumi.Output<string[] | undefined>;
    /**
     * The source image. When you set `baseImageType` to `IMAGE`, set `baseImage` to the ID of a custom image. When you set `baseImageType` to `IMAGE_FAMILY`, set `baseImage` to the name of an image family.
     */
    public readonly baseImage!: pulumi.Output<string>;
    /**
     * The type of the source image. Valid values: `IMAGE`, `IMAGE_FAMILY`.
     * - IMAGE: custom image.
     * - IMAGE_FAMILY: image family.
     */
    public readonly baseImageType!: pulumi.Output<string>;
    /**
     * The content of the image template. The content cannot be greater than 16 KB in size, and can contain up to 127 commands.
     */
    public readonly buildContent!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether to release the intermediate instance if the image cannot be created.
     */
    public readonly deleteInstanceOnFailure!: pulumi.Output<boolean>;
    /**
     * The description of the image template. The description must be `2` to `256` characters in length and cannot start with `http://` or `https://`. **Note:** If the intermediate instance cannot be started, the instance is released by default.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name prefix of the image to be created. The prefix must be `2` to `64` characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), periods (.),and hyphens (-).
     */
    public readonly imageName!: pulumi.Output<string | undefined>;
    /**
     * The instance type of the instance. You can call the DescribeInstanceTypes operation to query instance types. If you do not specify this parameter, an instance type that provides the fewest vCPUs and memory resources is automatically selected. This configuration is subject to resource availability of instance types. For example, the `ecs.g6.large` instance type is selected by default. If available `ecs.g6.large` resources are insufficient, the `ecs.g6.xlarge` instance type is selected.
     */
    public readonly instanceType!: pulumi.Output<string | undefined>;
    /**
     * The size of the outbound public bandwidth for the intermediate instance. Unit: `Mbit/s`. Valid values: `0` to `100`. Default value: `0`.
     */
    public readonly internetMaxBandwidthOut!: pulumi.Output<number>;
    /**
     * The name of the image template. The name must be `2` to `128` characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), periods (.),and hyphens (-).
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the resource group.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * The size of the system disk of the intermediate instance. Unit: GiB. Valid values: `20` to `500`. Default value: `40`.
     */
    public readonly systemDiskSize!: pulumi.Output<number>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The ID of region to which to distribute the created image.
     */
    public readonly toRegionIds!: pulumi.Output<string[] | undefined>;
    /**
     * The ID of the vSwitch. If you do not specify this parameter, a virtual private cloud (VPC) and a vSwitch are created by default.
     */
    public readonly vswitchId!: pulumi.Output<string | undefined>;

    /**
     * Create a EcsImagePipeline resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EcsImagePipelineArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EcsImagePipelineArgs | EcsImagePipelineState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EcsImagePipelineState | undefined;
            resourceInputs["addAccounts"] = state ? state.addAccounts : undefined;
            resourceInputs["baseImage"] = state ? state.baseImage : undefined;
            resourceInputs["baseImageType"] = state ? state.baseImageType : undefined;
            resourceInputs["buildContent"] = state ? state.buildContent : undefined;
            resourceInputs["deleteInstanceOnFailure"] = state ? state.deleteInstanceOnFailure : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["imageName"] = state ? state.imageName : undefined;
            resourceInputs["instanceType"] = state ? state.instanceType : undefined;
            resourceInputs["internetMaxBandwidthOut"] = state ? state.internetMaxBandwidthOut : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            resourceInputs["systemDiskSize"] = state ? state.systemDiskSize : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["toRegionIds"] = state ? state.toRegionIds : undefined;
            resourceInputs["vswitchId"] = state ? state.vswitchId : undefined;
        } else {
            const args = argsOrState as EcsImagePipelineArgs | undefined;
            if ((!args || args.baseImage === undefined) && !opts.urn) {
                throw new Error("Missing required property 'baseImage'");
            }
            if ((!args || args.baseImageType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'baseImageType'");
            }
            resourceInputs["addAccounts"] = args ? args.addAccounts : undefined;
            resourceInputs["baseImage"] = args ? args.baseImage : undefined;
            resourceInputs["baseImageType"] = args ? args.baseImageType : undefined;
            resourceInputs["buildContent"] = args ? args.buildContent : undefined;
            resourceInputs["deleteInstanceOnFailure"] = args ? args.deleteInstanceOnFailure : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["imageName"] = args ? args.imageName : undefined;
            resourceInputs["instanceType"] = args ? args.instanceType : undefined;
            resourceInputs["internetMaxBandwidthOut"] = args ? args.internetMaxBandwidthOut : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            resourceInputs["systemDiskSize"] = args ? args.systemDiskSize : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["toRegionIds"] = args ? args.toRegionIds : undefined;
            resourceInputs["vswitchId"] = args ? args.vswitchId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EcsImagePipeline.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EcsImagePipeline resources.
 */
export interface EcsImagePipelineState {
    /**
     * The ID of Alibaba Cloud account to which to share the created image.
     */
    addAccounts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The source image. When you set `baseImageType` to `IMAGE`, set `baseImage` to the ID of a custom image. When you set `baseImageType` to `IMAGE_FAMILY`, set `baseImage` to the name of an image family.
     */
    baseImage?: pulumi.Input<string>;
    /**
     * The type of the source image. Valid values: `IMAGE`, `IMAGE_FAMILY`.
     * - IMAGE: custom image.
     * - IMAGE_FAMILY: image family.
     */
    baseImageType?: pulumi.Input<string>;
    /**
     * The content of the image template. The content cannot be greater than 16 KB in size, and can contain up to 127 commands.
     */
    buildContent?: pulumi.Input<string>;
    /**
     * Specifies whether to release the intermediate instance if the image cannot be created.
     */
    deleteInstanceOnFailure?: pulumi.Input<boolean>;
    /**
     * The description of the image template. The description must be `2` to `256` characters in length and cannot start with `http://` or `https://`. **Note:** If the intermediate instance cannot be started, the instance is released by default.
     */
    description?: pulumi.Input<string>;
    /**
     * The name prefix of the image to be created. The prefix must be `2` to `64` characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), periods (.),and hyphens (-).
     */
    imageName?: pulumi.Input<string>;
    /**
     * The instance type of the instance. You can call the DescribeInstanceTypes operation to query instance types. If you do not specify this parameter, an instance type that provides the fewest vCPUs and memory resources is automatically selected. This configuration is subject to resource availability of instance types. For example, the `ecs.g6.large` instance type is selected by default. If available `ecs.g6.large` resources are insufficient, the `ecs.g6.xlarge` instance type is selected.
     */
    instanceType?: pulumi.Input<string>;
    /**
     * The size of the outbound public bandwidth for the intermediate instance. Unit: `Mbit/s`. Valid values: `0` to `100`. Default value: `0`.
     */
    internetMaxBandwidthOut?: pulumi.Input<number>;
    /**
     * The name of the image template. The name must be `2` to `128` characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), periods (.),and hyphens (-).
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The size of the system disk of the intermediate instance. Unit: GiB. Valid values: `20` to `500`. Default value: `40`.
     */
    systemDiskSize?: pulumi.Input<number>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of region to which to distribute the created image.
     */
    toRegionIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the vSwitch. If you do not specify this parameter, a virtual private cloud (VPC) and a vSwitch are created by default.
     */
    vswitchId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EcsImagePipeline resource.
 */
export interface EcsImagePipelineArgs {
    /**
     * The ID of Alibaba Cloud account to which to share the created image.
     */
    addAccounts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The source image. When you set `baseImageType` to `IMAGE`, set `baseImage` to the ID of a custom image. When you set `baseImageType` to `IMAGE_FAMILY`, set `baseImage` to the name of an image family.
     */
    baseImage: pulumi.Input<string>;
    /**
     * The type of the source image. Valid values: `IMAGE`, `IMAGE_FAMILY`.
     * - IMAGE: custom image.
     * - IMAGE_FAMILY: image family.
     */
    baseImageType: pulumi.Input<string>;
    /**
     * The content of the image template. The content cannot be greater than 16 KB in size, and can contain up to 127 commands.
     */
    buildContent?: pulumi.Input<string>;
    /**
     * Specifies whether to release the intermediate instance if the image cannot be created.
     */
    deleteInstanceOnFailure?: pulumi.Input<boolean>;
    /**
     * The description of the image template. The description must be `2` to `256` characters in length and cannot start with `http://` or `https://`. **Note:** If the intermediate instance cannot be started, the instance is released by default.
     */
    description?: pulumi.Input<string>;
    /**
     * The name prefix of the image to be created. The prefix must be `2` to `64` characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), periods (.),and hyphens (-).
     */
    imageName?: pulumi.Input<string>;
    /**
     * The instance type of the instance. You can call the DescribeInstanceTypes operation to query instance types. If you do not specify this parameter, an instance type that provides the fewest vCPUs and memory resources is automatically selected. This configuration is subject to resource availability of instance types. For example, the `ecs.g6.large` instance type is selected by default. If available `ecs.g6.large` resources are insufficient, the `ecs.g6.xlarge` instance type is selected.
     */
    instanceType?: pulumi.Input<string>;
    /**
     * The size of the outbound public bandwidth for the intermediate instance. Unit: `Mbit/s`. Valid values: `0` to `100`. Default value: `0`.
     */
    internetMaxBandwidthOut?: pulumi.Input<number>;
    /**
     * The name of the image template. The name must be `2` to `128` characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), periods (.),and hyphens (-).
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The size of the system disk of the intermediate instance. Unit: GiB. Valid values: `20` to `500`. Default value: `40`.
     */
    systemDiskSize?: pulumi.Input<number>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of region to which to distribute the created image.
     */
    toRegionIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the vSwitch. If you do not specify this parameter, a virtual private cloud (VPC) and a vSwitch are created by default.
     */
    vswitchId?: pulumi.Input<string>;
}

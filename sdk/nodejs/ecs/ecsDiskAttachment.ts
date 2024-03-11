// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an Alicloud ECS Disk Attachment as a resource, to attach and detach disks from ECS Instances.
 *
 * For information about ECS Disk Attachment and how to use it, see [What is Disk Attachment](https://www.alibabacloud.com/help/en/doc-detail/25515.htm).
 *
 * > **NOTE:** Available since v1.122.0+.
 *
 * ## Example Usage
 *
 * Basic usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "tf-example";
 * const defaultZones = alicloud.getZones({
 *     availableResourceCreation: "Instance",
 * });
 * const defaultInstanceTypes = defaultZones.then(defaultZones => alicloud.ecs.getInstanceTypes({
 *     availabilityZone: defaultZones.zones?.[0]?.id,
 *     instanceTypeFamily: "ecs.sn1ne",
 * }));
 * const defaultNetwork = new alicloud.vpc.Network("defaultNetwork", {
 *     vpcName: name,
 *     cidrBlock: "10.4.0.0/16",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("defaultSwitch", {
 *     vpcId: defaultNetwork.id,
 *     cidrBlock: "10.4.0.0/24",
 *     zoneId: defaultZones.then(defaultZones => defaultZones.zones?.[0]?.id),
 * });
 * const defaultSecurityGroup = new alicloud.ecs.SecurityGroup("defaultSecurityGroup", {
 *     description: "New security group",
 *     vpcId: defaultNetwork.id,
 * });
 * const defaultImages = alicloud.ecs.getImages({
 *     nameRegex: "^ubuntu_[0-9]+_[0-9]+_x64*",
 *     mostRecent: true,
 *     owners: "system",
 * });
 * const defaultInstance = new alicloud.ecs.Instance("defaultInstance", {
 *     availabilityZone: defaultZones.then(defaultZones => defaultZones.zones?.[0]?.id),
 *     instanceName: name,
 *     hostName: name,
 *     imageId: defaultImages.then(defaultImages => defaultImages.images?.[0]?.id),
 *     instanceType: defaultInstanceTypes.then(defaultInstanceTypes => defaultInstanceTypes.instanceTypes?.[0]?.id),
 *     securityGroups: [defaultSecurityGroup.id],
 *     vswitchId: defaultSwitch.id,
 * });
 * const disk = alicloud.getZones({
 *     availableResourceCreation: "VSwitch",
 * });
 * const defaultEcsDisk = new alicloud.ecs.EcsDisk("defaultEcsDisk", {
 *     zoneId: disk.then(disk => disk.zones?.[0]?.id),
 *     category: "cloud_efficiency",
 *     deleteAutoSnapshot: true,
 *     description: "Test For Terraform",
 *     diskName: name,
 *     enableAutoSnapshot: true,
 *     encrypted: true,
 *     size: 500,
 *     tags: {
 *         Created: "TF",
 *         Environment: "Acceptance-test",
 *     },
 * });
 * const defaultEcsDiskAttachment = new alicloud.ecs.EcsDiskAttachment("defaultEcsDiskAttachment", {
 *     diskId: defaultEcsDisk.id,
 *     instanceId: defaultInstance.id,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * The disk attachment can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:ecs/ecsDiskAttachment:EcsDiskAttachment example d-abc12345678:i-abc12355
 * ```
 */
export class EcsDiskAttachment extends pulumi.CustomResource {
    /**
     * Get an existing EcsDiskAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EcsDiskAttachmentState, opts?: pulumi.CustomResourceOptions): EcsDiskAttachment {
        return new EcsDiskAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ecs/ecsDiskAttachment:EcsDiskAttachment';

    /**
     * Returns true if the given object is an instance of EcsDiskAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EcsDiskAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EcsDiskAttachment.__pulumiType;
    }

    /**
     * Whether to mount as a system disk. Default to: `false`.
     */
    public readonly bootable!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates whether the disk is released together with the instance. Default to: `false`.
     */
    public readonly deleteWithInstance!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the cloud disk device.
     */
    public /*out*/ readonly device!: pulumi.Output<string>;
    /**
     * ID of the Disk to be attached.
     */
    public readonly diskId!: pulumi.Output<string>;
    /**
     * ID of the Instance to attach to.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The name of key pair
     */
    public readonly keyPairName!: pulumi.Output<string | undefined>;
    /**
     * When mounting the system disk, setting the user name and password of the instance is only effective for the administrator and root user names, and other user names are not effective.
     */
    public readonly password!: pulumi.Output<string | undefined>;

    /**
     * Create a EcsDiskAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EcsDiskAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EcsDiskAttachmentArgs | EcsDiskAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EcsDiskAttachmentState | undefined;
            resourceInputs["bootable"] = state ? state.bootable : undefined;
            resourceInputs["deleteWithInstance"] = state ? state.deleteWithInstance : undefined;
            resourceInputs["device"] = state ? state.device : undefined;
            resourceInputs["diskId"] = state ? state.diskId : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["keyPairName"] = state ? state.keyPairName : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
        } else {
            const args = argsOrState as EcsDiskAttachmentArgs | undefined;
            if ((!args || args.diskId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'diskId'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["bootable"] = args ? args.bootable : undefined;
            resourceInputs["deleteWithInstance"] = args ? args.deleteWithInstance : undefined;
            resourceInputs["diskId"] = args ? args.diskId : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["keyPairName"] = args ? args.keyPairName : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["device"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EcsDiskAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EcsDiskAttachment resources.
 */
export interface EcsDiskAttachmentState {
    /**
     * Whether to mount as a system disk. Default to: `false`.
     */
    bootable?: pulumi.Input<boolean>;
    /**
     * Indicates whether the disk is released together with the instance. Default to: `false`.
     */
    deleteWithInstance?: pulumi.Input<boolean>;
    /**
     * The name of the cloud disk device.
     */
    device?: pulumi.Input<string>;
    /**
     * ID of the Disk to be attached.
     */
    diskId?: pulumi.Input<string>;
    /**
     * ID of the Instance to attach to.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The name of key pair
     */
    keyPairName?: pulumi.Input<string>;
    /**
     * When mounting the system disk, setting the user name and password of the instance is only effective for the administrator and root user names, and other user names are not effective.
     */
    password?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EcsDiskAttachment resource.
 */
export interface EcsDiskAttachmentArgs {
    /**
     * Whether to mount as a system disk. Default to: `false`.
     */
    bootable?: pulumi.Input<boolean>;
    /**
     * Indicates whether the disk is released together with the instance. Default to: `false`.
     */
    deleteWithInstance?: pulumi.Input<boolean>;
    /**
     * ID of the Disk to be attached.
     */
    diskId: pulumi.Input<string>;
    /**
     * ID of the Instance to attach to.
     */
    instanceId: pulumi.Input<string>;
    /**
     * The name of key pair
     */
    keyPairName?: pulumi.Input<string>;
    /**
     * When mounting the system disk, setting the user name and password of the instance is only effective for the administrator and root user names, and other user names are not effective.
     */
    password?: pulumi.Input<string>;
}

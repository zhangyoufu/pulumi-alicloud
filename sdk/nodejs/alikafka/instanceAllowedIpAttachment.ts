// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a AliKafka Instance Allowed Ip Attachment resource.
 *
 * For information about Ali Kafka Instance Allowed Ip Attachment and how to use it, see [What is Instance Allowed Ip Attachment](https://www.alibabacloud.com/help/en/message-queue-for-apache-kafka/latest/api-alikafka-2019-09-16-updateallowedip).
 *
 * > **NOTE:** Available since v1.163.0.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * import * as random from "@pulumi/random";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "tf-example";
 * const defaultRandomInteger = new random.RandomInteger("defaultRandomInteger", {
 *     min: 10000,
 *     max: 99999,
 * });
 * const defaultZones = alicloud.getZones({
 *     availableResourceCreation: "VSwitch",
 * });
 * const defaultNetwork = new alicloud.vpc.Network("defaultNetwork", {
 *     vpcName: name,
 *     cidrBlock: "10.4.0.0/16",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("defaultSwitch", {
 *     vswitchName: name,
 *     cidrBlock: "10.4.0.0/24",
 *     vpcId: defaultNetwork.id,
 *     zoneId: defaultZones.then(defaultZones => defaultZones.zones?.[0]?.id),
 * });
 * const defaultSecurityGroup = new alicloud.ecs.SecurityGroup("defaultSecurityGroup", {vpcId: defaultNetwork.id});
 * const defaultInstance = new alicloud.alikafka.Instance("defaultInstance", {
 *     partitionNum: 50,
 *     diskType: 1,
 *     diskSize: 500,
 *     deployType: 5,
 *     ioMax: 20,
 *     vswitchId: defaultSwitch.id,
 *     securityGroup: defaultSecurityGroup.id,
 * });
 * const defaultInstanceAllowedIpAttachment = new alicloud.alikafka.InstanceAllowedIpAttachment("defaultInstanceAllowedIpAttachment", {
 *     allowedIp: "114.237.9.78/32",
 *     allowedType: "vpc",
 *     instanceId: defaultInstance.id,
 *     portRange: "9092/9092",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * AliKafka Instance Allowed Ip Attachment can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:alikafka/instanceAllowedIpAttachment:InstanceAllowedIpAttachment example <instance_id>:<allowed_type>:<port_range>:<allowed_ip>
 * ```
 */
export class InstanceAllowedIpAttachment extends pulumi.CustomResource {
    /**
     * Get an existing InstanceAllowedIpAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceAllowedIpAttachmentState, opts?: pulumi.CustomResourceOptions): InstanceAllowedIpAttachment {
        return new InstanceAllowedIpAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:alikafka/instanceAllowedIpAttachment:InstanceAllowedIpAttachment';

    /**
     * Returns true if the given object is an instance of InstanceAllowedIpAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstanceAllowedIpAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstanceAllowedIpAttachment.__pulumiType;
    }

    /**
     * The allowed ip. It can be a CIDR block.
     */
    public readonly allowedIp!: pulumi.Output<string>;
    /**
     * The type of whitelist. Valid Value: `vpc`, `internet`. **NOTE:** From version 1.179.0, `allowedType` can be set to `internet`.
     */
    public readonly allowedType!: pulumi.Output<string>;
    /**
     * The ID of the instance.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The Port range.  Valid Value: `9092/9092`, `9093/9093`. **NOTE:** From version 1.179.0, `portRange` can be set to `9093/9093`.
     * - `9092/9092`: port range for a VPC whitelist.
     * - `9093/9093`: port range for an Internet whitelist.
     */
    public readonly portRange!: pulumi.Output<string>;

    /**
     * Create a InstanceAllowedIpAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceAllowedIpAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceAllowedIpAttachmentArgs | InstanceAllowedIpAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceAllowedIpAttachmentState | undefined;
            resourceInputs["allowedIp"] = state ? state.allowedIp : undefined;
            resourceInputs["allowedType"] = state ? state.allowedType : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["portRange"] = state ? state.portRange : undefined;
        } else {
            const args = argsOrState as InstanceAllowedIpAttachmentArgs | undefined;
            if ((!args || args.allowedIp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'allowedIp'");
            }
            if ((!args || args.allowedType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'allowedType'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.portRange === undefined) && !opts.urn) {
                throw new Error("Missing required property 'portRange'");
            }
            resourceInputs["allowedIp"] = args ? args.allowedIp : undefined;
            resourceInputs["allowedType"] = args ? args.allowedType : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["portRange"] = args ? args.portRange : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(InstanceAllowedIpAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstanceAllowedIpAttachment resources.
 */
export interface InstanceAllowedIpAttachmentState {
    /**
     * The allowed ip. It can be a CIDR block.
     */
    allowedIp?: pulumi.Input<string>;
    /**
     * The type of whitelist. Valid Value: `vpc`, `internet`. **NOTE:** From version 1.179.0, `allowedType` can be set to `internet`.
     */
    allowedType?: pulumi.Input<string>;
    /**
     * The ID of the instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The Port range.  Valid Value: `9092/9092`, `9093/9093`. **NOTE:** From version 1.179.0, `portRange` can be set to `9093/9093`.
     * - `9092/9092`: port range for a VPC whitelist.
     * - `9093/9093`: port range for an Internet whitelist.
     */
    portRange?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InstanceAllowedIpAttachment resource.
 */
export interface InstanceAllowedIpAttachmentArgs {
    /**
     * The allowed ip. It can be a CIDR block.
     */
    allowedIp: pulumi.Input<string>;
    /**
     * The type of whitelist. Valid Value: `vpc`, `internet`. **NOTE:** From version 1.179.0, `allowedType` can be set to `internet`.
     */
    allowedType: pulumi.Input<string>;
    /**
     * The ID of the instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * The Port range.  Valid Value: `9092/9092`, `9093/9093`. **NOTE:** From version 1.179.0, `portRange` can be set to `9093/9093`.
     * - `9092/9092`: port range for a VPC whitelist.
     * - `9093/9093`: port range for an Internet whitelist.
     */
    portRange: pulumi.Input<string>;
}

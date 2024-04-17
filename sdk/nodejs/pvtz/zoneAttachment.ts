// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * Using `vpcIds` to attach being in same region several vpc instances to a private zone
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const zone = new alicloud.pvtz.Zone("zone", {zoneName: "foo.example.com"});
 * const first = new alicloud.vpc.Network("first", {
 *     vpcName: "the-first-vpc",
 *     cidrBlock: "172.16.0.0/12",
 * });
 * const second = new alicloud.vpc.Network("second", {
 *     vpcName: "the-second-vpc",
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const zone_attachment = new alicloud.pvtz.ZoneAttachment("zone-attachment", {
 *     zoneId: zone.id,
 *     vpcIds: [
 *         first.id,
 *         second.id,
 *     ],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * Using `vpcs` to attach being in same region several vpc instances to a private zone
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const zone = new alicloud.pvtz.Zone("zone", {zoneName: "foo.example.com"});
 * const first = new alicloud.vpc.Network("first", {
 *     vpcName: "the-first-vpc",
 *     cidrBlock: "172.16.0.0/12",
 * });
 * const second = new alicloud.vpc.Network("second", {
 *     vpcName: "the-second-vpc",
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const zone_attachment = new alicloud.pvtz.ZoneAttachment("zone-attachment", {
 *     zoneId: zone.id,
 *     vpcs: [
 *         {
 *             vpcId: first.id,
 *         },
 *         {
 *             vpcId: second.id,
 *         },
 *     ],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * Using `vpcs` to attach being in different regions several vpc instances to a private zone
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const zone = new alicloud.pvtz.Zone("zone", {zoneName: "foo.example.com"});
 * const first = new alicloud.vpc.Network("first", {
 *     vpcName: "the-first-vpc",
 *     cidrBlock: "172.16.0.0/12",
 * });
 * const second = new alicloud.vpc.Network("second", {
 *     vpcName: "the-second-vpc",
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const third = new alicloud.vpc.Network("third", {
 *     vpcName: "the-third-vpc",
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const zone_attachment = new alicloud.pvtz.ZoneAttachment("zone-attachment", {
 *     zoneId: zone.id,
 *     vpcs: [
 *         {
 *             vpcId: first.id,
 *         },
 *         {
 *             vpcId: second.id,
 *         },
 *         {
 *             regionId: "eu-central-1",
 *             vpcId: third.id,
 *         },
 *     ],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Private Zone attachment can be imported using the id(same with `zone_id`), e.g.
 *
 * ```sh
 * $ pulumi import alicloud:pvtz/zoneAttachment:ZoneAttachment example abc123456
 * ```
 */
export class ZoneAttachment extends pulumi.CustomResource {
    /**
     * Get an existing ZoneAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ZoneAttachmentState, opts?: pulumi.CustomResourceOptions): ZoneAttachment {
        return new ZoneAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:pvtz/zoneAttachment:ZoneAttachment';

    /**
     * Returns true if the given object is an instance of ZoneAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ZoneAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ZoneAttachment.__pulumiType;
    }

    /**
     * The language of code.
     */
    public readonly lang!: pulumi.Output<string | undefined>;
    /**
     * The user custom IP address.
     */
    public readonly userClientIp!: pulumi.Output<string | undefined>;
    /**
     * The id List of the VPC with the same region, for example:["vpc-1","vpc-2"].
     */
    public readonly vpcIds!: pulumi.Output<string[]>;
    /**
     * See `vpcs` below.Recommend to use `vpcs`.
     */
    public readonly vpcs!: pulumi.Output<outputs.pvtz.ZoneAttachmentVpc[]>;
    /**
     * The name of the Private Zone Record.
     */
    public readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a ZoneAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ZoneAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ZoneAttachmentArgs | ZoneAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ZoneAttachmentState | undefined;
            resourceInputs["lang"] = state ? state.lang : undefined;
            resourceInputs["userClientIp"] = state ? state.userClientIp : undefined;
            resourceInputs["vpcIds"] = state ? state.vpcIds : undefined;
            resourceInputs["vpcs"] = state ? state.vpcs : undefined;
            resourceInputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as ZoneAttachmentArgs | undefined;
            if ((!args || args.zoneId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zoneId'");
            }
            resourceInputs["lang"] = args ? args.lang : undefined;
            resourceInputs["userClientIp"] = args ? args.userClientIp : undefined;
            resourceInputs["vpcIds"] = args ? args.vpcIds : undefined;
            resourceInputs["vpcs"] = args ? args.vpcs : undefined;
            resourceInputs["zoneId"] = args ? args.zoneId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ZoneAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ZoneAttachment resources.
 */
export interface ZoneAttachmentState {
    /**
     * The language of code.
     */
    lang?: pulumi.Input<string>;
    /**
     * The user custom IP address.
     */
    userClientIp?: pulumi.Input<string>;
    /**
     * The id List of the VPC with the same region, for example:["vpc-1","vpc-2"].
     */
    vpcIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * See `vpcs` below.Recommend to use `vpcs`.
     */
    vpcs?: pulumi.Input<pulumi.Input<inputs.pvtz.ZoneAttachmentVpc>[]>;
    /**
     * The name of the Private Zone Record.
     */
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ZoneAttachment resource.
 */
export interface ZoneAttachmentArgs {
    /**
     * The language of code.
     */
    lang?: pulumi.Input<string>;
    /**
     * The user custom IP address.
     */
    userClientIp?: pulumi.Input<string>;
    /**
     * The id List of the VPC with the same region, for example:["vpc-1","vpc-2"].
     */
    vpcIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * See `vpcs` below.Recommend to use `vpcs`.
     */
    vpcs?: pulumi.Input<pulumi.Input<inputs.pvtz.ZoneAttachmentVpc>[]>;
    /**
     * The name of the Private Zone Record.
     */
    zoneId: pulumi.Input<string>;
}

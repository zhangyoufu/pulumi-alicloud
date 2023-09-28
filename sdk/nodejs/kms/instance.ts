// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * KMS Instance can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:kms/instance:Instance example <id>
 * ```
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:kms/instance:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * Aucillary VPCs used to access this KMS instance. See `bindVpcs` below.
     */
    public readonly bindVpcs!: pulumi.Output<outputs.kms.InstanceBindVpc[] | undefined>;
    /**
     * KMS instance certificate chain in PEM format.
     */
    public /*out*/ readonly caCertificateChainPem!: pulumi.Output<string>;
    /**
     * The creation time of the resource.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The name of the resource.
     */
    public /*out*/ readonly instanceName!: pulumi.Output<string>;
    /**
     * Maximum number of stored keys.
     */
    public readonly keyNum!: pulumi.Output<number>;
    /**
     * KMS Instance commodity type (software/hardware). Currently, only version 3 is supported.
     */
    public readonly productVersion!: pulumi.Output<string | undefined>;
    /**
     * Automatic renewal period, in months.
     */
    public readonly renewPeriod!: pulumi.Output<number | undefined>;
    /**
     * Renewal options (manual renewal, automatic renewal, no renewal).
     */
    public readonly renewStatus!: pulumi.Output<string | undefined>;
    /**
     * Maximum number of Secrets.
     */
    public readonly secretNum!: pulumi.Output<number>;
    /**
     * The computation performance level of the KMS instance.
     */
    public readonly spec!: pulumi.Output<number>;
    /**
     * Instance status.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Instance VPC id.
     */
    public readonly vpcId!: pulumi.Output<string>;
    /**
     * The number of managed accesses. The maximum number of VPCs that can access this KMS instance.
     */
    public readonly vpcNum!: pulumi.Output<number>;
    /**
     * Instance bind vswitches.
     */
    public readonly vswitchIds!: pulumi.Output<string[]>;
    /**
     * zone id.
     */
    public readonly zoneIds!: pulumi.Output<string[]>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceState | undefined;
            resourceInputs["bindVpcs"] = state ? state.bindVpcs : undefined;
            resourceInputs["caCertificateChainPem"] = state ? state.caCertificateChainPem : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["instanceName"] = state ? state.instanceName : undefined;
            resourceInputs["keyNum"] = state ? state.keyNum : undefined;
            resourceInputs["productVersion"] = state ? state.productVersion : undefined;
            resourceInputs["renewPeriod"] = state ? state.renewPeriod : undefined;
            resourceInputs["renewStatus"] = state ? state.renewStatus : undefined;
            resourceInputs["secretNum"] = state ? state.secretNum : undefined;
            resourceInputs["spec"] = state ? state.spec : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["vpcNum"] = state ? state.vpcNum : undefined;
            resourceInputs["vswitchIds"] = state ? state.vswitchIds : undefined;
            resourceInputs["zoneIds"] = state ? state.zoneIds : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if ((!args || args.keyNum === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keyNum'");
            }
            if ((!args || args.secretNum === undefined) && !opts.urn) {
                throw new Error("Missing required property 'secretNum'");
            }
            if ((!args || args.spec === undefined) && !opts.urn) {
                throw new Error("Missing required property 'spec'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            if ((!args || args.vpcNum === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcNum'");
            }
            if ((!args || args.vswitchIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vswitchIds'");
            }
            if ((!args || args.zoneIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zoneIds'");
            }
            resourceInputs["bindVpcs"] = args ? args.bindVpcs : undefined;
            resourceInputs["keyNum"] = args ? args.keyNum : undefined;
            resourceInputs["productVersion"] = args ? args.productVersion : undefined;
            resourceInputs["renewPeriod"] = args ? args.renewPeriod : undefined;
            resourceInputs["renewStatus"] = args ? args.renewStatus : undefined;
            resourceInputs["secretNum"] = args ? args.secretNum : undefined;
            resourceInputs["spec"] = args ? args.spec : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["vpcNum"] = args ? args.vpcNum : undefined;
            resourceInputs["vswitchIds"] = args ? args.vswitchIds : undefined;
            resourceInputs["zoneIds"] = args ? args.zoneIds : undefined;
            resourceInputs["caCertificateChainPem"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["instanceName"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Instance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * Aucillary VPCs used to access this KMS instance. See `bindVpcs` below.
     */
    bindVpcs?: pulumi.Input<pulumi.Input<inputs.kms.InstanceBindVpc>[]>;
    /**
     * KMS instance certificate chain in PEM format.
     */
    caCertificateChainPem?: pulumi.Input<string>;
    /**
     * The creation time of the resource.
     */
    createTime?: pulumi.Input<string>;
    /**
     * The name of the resource.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * Maximum number of stored keys.
     */
    keyNum?: pulumi.Input<number>;
    /**
     * KMS Instance commodity type (software/hardware). Currently, only version 3 is supported.
     */
    productVersion?: pulumi.Input<string>;
    /**
     * Automatic renewal period, in months.
     */
    renewPeriod?: pulumi.Input<number>;
    /**
     * Renewal options (manual renewal, automatic renewal, no renewal).
     */
    renewStatus?: pulumi.Input<string>;
    /**
     * Maximum number of Secrets.
     */
    secretNum?: pulumi.Input<number>;
    /**
     * The computation performance level of the KMS instance.
     */
    spec?: pulumi.Input<number>;
    /**
     * Instance status.
     */
    status?: pulumi.Input<string>;
    /**
     * Instance VPC id.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * The number of managed accesses. The maximum number of VPCs that can access this KMS instance.
     */
    vpcNum?: pulumi.Input<number>;
    /**
     * Instance bind vswitches.
     */
    vswitchIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * zone id.
     */
    zoneIds?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * Aucillary VPCs used to access this KMS instance. See `bindVpcs` below.
     */
    bindVpcs?: pulumi.Input<pulumi.Input<inputs.kms.InstanceBindVpc>[]>;
    /**
     * Maximum number of stored keys.
     */
    keyNum: pulumi.Input<number>;
    /**
     * KMS Instance commodity type (software/hardware). Currently, only version 3 is supported.
     */
    productVersion?: pulumi.Input<string>;
    /**
     * Automatic renewal period, in months.
     */
    renewPeriod?: pulumi.Input<number>;
    /**
     * Renewal options (manual renewal, automatic renewal, no renewal).
     */
    renewStatus?: pulumi.Input<string>;
    /**
     * Maximum number of Secrets.
     */
    secretNum: pulumi.Input<number>;
    /**
     * The computation performance level of the KMS instance.
     */
    spec: pulumi.Input<number>;
    /**
     * Instance VPC id.
     */
    vpcId: pulumi.Input<string>;
    /**
     * The number of managed accesses. The maximum number of VPCs that can access this KMS instance.
     */
    vpcNum: pulumi.Input<number>;
    /**
     * Instance bind vswitches.
     */
    vswitchIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * zone id.
     */
    zoneIds: pulumi.Input<pulumi.Input<string>[]>;
}

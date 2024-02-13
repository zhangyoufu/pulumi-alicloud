// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a Private Zone Rule resource.
 *
 * For information about Private Zone Rule and how to use it, see [What is Rule](https://www.alibabacloud.com/help/en/privatezone/latest/add-forwarding-rule).
 *
 * > **NOTE:** Available since v1.143.0.
 *
 * ## Import
 *
 * Private Zone Rule can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:pvtz/rule:Rule example <id>
 * ```
 */
export class Rule extends pulumi.CustomResource {
    /**
     * Get an existing Rule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RuleState, opts?: pulumi.CustomResourceOptions): Rule {
        return new Rule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:pvtz/rule:Rule';

    /**
     * Returns true if the given object is an instance of Rule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Rule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Rule.__pulumiType;
    }

    /**
     * The ID of the Endpoint.
     */
    public readonly endpointId!: pulumi.Output<string>;
    /**
     * Forwarding target. See `forwardIps` below.
     */
    public readonly forwardIps!: pulumi.Output<outputs.pvtz.RuleForwardIp[]>;
    /**
     * The name of the resource.
     */
    public readonly ruleName!: pulumi.Output<string>;
    /**
     * The type of the rule. Valid values: `OUTBOUND`.
     */
    public readonly type!: pulumi.Output<string | undefined>;
    /**
     * The name of the forwarding zone.
     */
    public readonly zoneName!: pulumi.Output<string>;

    /**
     * Create a Rule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RuleArgs | RuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RuleState | undefined;
            resourceInputs["endpointId"] = state ? state.endpointId : undefined;
            resourceInputs["forwardIps"] = state ? state.forwardIps : undefined;
            resourceInputs["ruleName"] = state ? state.ruleName : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["zoneName"] = state ? state.zoneName : undefined;
        } else {
            const args = argsOrState as RuleArgs | undefined;
            if ((!args || args.endpointId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endpointId'");
            }
            if ((!args || args.forwardIps === undefined) && !opts.urn) {
                throw new Error("Missing required property 'forwardIps'");
            }
            if ((!args || args.ruleName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ruleName'");
            }
            if ((!args || args.zoneName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zoneName'");
            }
            resourceInputs["endpointId"] = args ? args.endpointId : undefined;
            resourceInputs["forwardIps"] = args ? args.forwardIps : undefined;
            resourceInputs["ruleName"] = args ? args.ruleName : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["zoneName"] = args ? args.zoneName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Rule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Rule resources.
 */
export interface RuleState {
    /**
     * The ID of the Endpoint.
     */
    endpointId?: pulumi.Input<string>;
    /**
     * Forwarding target. See `forwardIps` below.
     */
    forwardIps?: pulumi.Input<pulumi.Input<inputs.pvtz.RuleForwardIp>[]>;
    /**
     * The name of the resource.
     */
    ruleName?: pulumi.Input<string>;
    /**
     * The type of the rule. Valid values: `OUTBOUND`.
     */
    type?: pulumi.Input<string>;
    /**
     * The name of the forwarding zone.
     */
    zoneName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Rule resource.
 */
export interface RuleArgs {
    /**
     * The ID of the Endpoint.
     */
    endpointId: pulumi.Input<string>;
    /**
     * Forwarding target. See `forwardIps` below.
     */
    forwardIps: pulumi.Input<pulumi.Input<inputs.pvtz.RuleForwardIp>[]>;
    /**
     * The name of the resource.
     */
    ruleName: pulumi.Input<string>;
    /**
     * The type of the rule. Valid values: `OUTBOUND`.
     */
    type?: pulumi.Input<string>;
    /**
     * The name of the forwarding zone.
     */
    zoneName: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a Global Accelerator (GA) Forwarding Rule resource.
 *
 * For information about Global Accelerator (GA) Forwarding Rule and how to use it, see [What is Forwarding Rule](https://www.alibabacloud.com/help/zh/doc-detail/205815.htm).
 *
 * > **NOTE:** Available since v1.120.0.
 *
 * ## Import
 *
 * Ga Forwarding Rule can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:ga/forwardingRule:ForwardingRule example <id>
 * ```
 */
export class ForwardingRule extends pulumi.CustomResource {
    /**
     * Get an existing ForwardingRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ForwardingRuleState, opts?: pulumi.CustomResourceOptions): ForwardingRule {
        return new ForwardingRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ga/forwardingRule:ForwardingRule';

    /**
     * Returns true if the given object is an instance of ForwardingRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ForwardingRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ForwardingRule.__pulumiType;
    }

    /**
     * The ID of the Global Accelerator instance.
     */
    public readonly acceleratorId!: pulumi.Output<string>;
    /**
     * Forwarding Policy ID.
     */
    public /*out*/ readonly forwardingRuleId!: pulumi.Output<string>;
    /**
     * Forwarding policy name. The length of the name is 2-128 English or Chinese characters. It must start with uppercase and lowercase letters or Chinese characters. It can contain numbers, half width period (.), underscores (_) And dash (-).
     */
    public readonly forwardingRuleName!: pulumi.Output<string | undefined>;
    /**
     * Forwarding Policy Status.
     */
    public /*out*/ readonly forwardingRuleStatus!: pulumi.Output<string>;
    /**
     * The ID of the listener.
     */
    public readonly listenerId!: pulumi.Output<string>;
    /**
     * Forwarding policy priority.
     */
    public readonly priority!: pulumi.Output<number>;
    /**
     * Forward action. See `ruleActions` below.
     */
    public readonly ruleActions!: pulumi.Output<outputs.ga.ForwardingRuleRuleAction[]>;
    /**
     * Forwarding condition list. See `ruleConditions` below.
     */
    public readonly ruleConditions!: pulumi.Output<outputs.ga.ForwardingRuleRuleCondition[]>;

    /**
     * Create a ForwardingRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ForwardingRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ForwardingRuleArgs | ForwardingRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ForwardingRuleState | undefined;
            resourceInputs["acceleratorId"] = state ? state.acceleratorId : undefined;
            resourceInputs["forwardingRuleId"] = state ? state.forwardingRuleId : undefined;
            resourceInputs["forwardingRuleName"] = state ? state.forwardingRuleName : undefined;
            resourceInputs["forwardingRuleStatus"] = state ? state.forwardingRuleStatus : undefined;
            resourceInputs["listenerId"] = state ? state.listenerId : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["ruleActions"] = state ? state.ruleActions : undefined;
            resourceInputs["ruleConditions"] = state ? state.ruleConditions : undefined;
        } else {
            const args = argsOrState as ForwardingRuleArgs | undefined;
            if ((!args || args.acceleratorId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'acceleratorId'");
            }
            if ((!args || args.listenerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'listenerId'");
            }
            if ((!args || args.ruleActions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ruleActions'");
            }
            if ((!args || args.ruleConditions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ruleConditions'");
            }
            resourceInputs["acceleratorId"] = args ? args.acceleratorId : undefined;
            resourceInputs["forwardingRuleName"] = args ? args.forwardingRuleName : undefined;
            resourceInputs["listenerId"] = args ? args.listenerId : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["ruleActions"] = args ? args.ruleActions : undefined;
            resourceInputs["ruleConditions"] = args ? args.ruleConditions : undefined;
            resourceInputs["forwardingRuleId"] = undefined /*out*/;
            resourceInputs["forwardingRuleStatus"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ForwardingRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ForwardingRule resources.
 */
export interface ForwardingRuleState {
    /**
     * The ID of the Global Accelerator instance.
     */
    acceleratorId?: pulumi.Input<string>;
    /**
     * Forwarding Policy ID.
     */
    forwardingRuleId?: pulumi.Input<string>;
    /**
     * Forwarding policy name. The length of the name is 2-128 English or Chinese characters. It must start with uppercase and lowercase letters or Chinese characters. It can contain numbers, half width period (.), underscores (_) And dash (-).
     */
    forwardingRuleName?: pulumi.Input<string>;
    /**
     * Forwarding Policy Status.
     */
    forwardingRuleStatus?: pulumi.Input<string>;
    /**
     * The ID of the listener.
     */
    listenerId?: pulumi.Input<string>;
    /**
     * Forwarding policy priority.
     */
    priority?: pulumi.Input<number>;
    /**
     * Forward action. See `ruleActions` below.
     */
    ruleActions?: pulumi.Input<pulumi.Input<inputs.ga.ForwardingRuleRuleAction>[]>;
    /**
     * Forwarding condition list. See `ruleConditions` below.
     */
    ruleConditions?: pulumi.Input<pulumi.Input<inputs.ga.ForwardingRuleRuleCondition>[]>;
}

/**
 * The set of arguments for constructing a ForwardingRule resource.
 */
export interface ForwardingRuleArgs {
    /**
     * The ID of the Global Accelerator instance.
     */
    acceleratorId: pulumi.Input<string>;
    /**
     * Forwarding policy name. The length of the name is 2-128 English or Chinese characters. It must start with uppercase and lowercase letters or Chinese characters. It can contain numbers, half width period (.), underscores (_) And dash (-).
     */
    forwardingRuleName?: pulumi.Input<string>;
    /**
     * The ID of the listener.
     */
    listenerId: pulumi.Input<string>;
    /**
     * Forwarding policy priority.
     */
    priority?: pulumi.Input<number>;
    /**
     * Forward action. See `ruleActions` below.
     */
    ruleActions: pulumi.Input<pulumi.Input<inputs.ga.ForwardingRuleRuleAction>[]>;
    /**
     * Forwarding condition list. See `ruleConditions` below.
     */
    ruleConditions: pulumi.Input<pulumi.Input<inputs.ga.ForwardingRuleRuleCondition>[]>;
}

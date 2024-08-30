// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a VPC Traffic Mirror Filter Ingress Rule resource. Traffic mirror entry rule.
 *
 * For information about VPC Traffic Mirror Filter Ingress Rule and how to use it, see [What is Traffic Mirror Filter Ingress Rule](https://www.alibabacloud.com/help/doc-detail/261357.htm).
 *
 * > **NOTE:** Available since v1.141.0.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = new alicloud.vpc.TrafficMirrorFilter("example", {trafficMirrorFilterName: "example_value"});
 * const exampleTrafficMirrorFilterIngressRule = new alicloud.vpc.TrafficMirrorFilterIngressRule("example", {
 *     trafficMirrorFilterId: example.id,
 *     priority: 1,
 *     action: "accept",
 *     protocol: "UDP",
 *     destinationCidrBlock: "10.0.0.0/24",
 *     sourceCidrBlock: "10.0.0.0/24",
 *     destinationPortRange: "1/120",
 *     sourcePortRange: "1/120",
 * });
 * ```
 *
 * ## Import
 *
 * VPC Traffic Mirror Filter Ingress Rule can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:vpc/trafficMirrorFilterIngressRule:TrafficMirrorFilterIngressRule example <traffic_mirror_filter_id>:<traffic_mirror_filter_ingress_rule_id>
 * ```
 */
export class TrafficMirrorFilterIngressRule extends pulumi.CustomResource {
    /**
     * Get an existing TrafficMirrorFilterIngressRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TrafficMirrorFilterIngressRuleState, opts?: pulumi.CustomResourceOptions): TrafficMirrorFilterIngressRule {
        return new TrafficMirrorFilterIngressRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:vpc/trafficMirrorFilterIngressRule:TrafficMirrorFilterIngressRule';

    /**
     * Returns true if the given object is an instance of TrafficMirrorFilterIngressRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TrafficMirrorFilterIngressRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TrafficMirrorFilterIngressRule.__pulumiType;
    }

    /**
     * The collection policy of the inbound rule. Valid values: `accept` or `drop`. `accept`: collects network traffic. `drop`: does not collect network traffic.
     */
    public readonly action!: pulumi.Output<string>;
    /**
     * The destination CIDR block of the inbound traffic.
     */
    public readonly destinationCidrBlock!: pulumi.Output<string>;
    /**
     * The destination CIDR block of the inbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
     */
    public readonly destinationPortRange!: pulumi.Output<string>;
    /**
     * Whether to PreCheck this request only. Value:
     * - **true**: sends a check request and does not create inbound or outbound rules. Check items include whether required parameters are filled in, request format, and restrictions. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.
     * - **false** (default): Sends a normal request and directly creates an inbound or outbound direction rule after checking.
     */
    public readonly dryRun!: pulumi.Output<boolean | undefined>;
    /**
     * The priority of the inbound rule. A smaller value indicates a higher priority. The maximum value is `10`, which indicates that you can configure at most 10 inbound rules for a filter.
     */
    public readonly priority!: pulumi.Output<number>;
    /**
     * The transport protocol used by inbound traffic that needs to be mirrored. Valid values: `ALL`, `ICMP`, `TCP`, `UDP`.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * . Field 'rule_action' has been deprecated from provider version 1.211.0. New field 'action' instead.
     *
     * @deprecated Field 'rule_action' has been deprecated since provider version 1.211.0. New field 'action' instead.
     */
    public readonly ruleAction!: pulumi.Output<string>;
    /**
     * The source CIDR block of the inbound traffic.
     */
    public readonly sourceCidrBlock!: pulumi.Output<string>;
    /**
     * The source port range of the inbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
     */
    public readonly sourcePortRange!: pulumi.Output<string>;
    /**
     * The state of the inbound rule. `Creating`, `Created`, `Modifying` and `Deleting`.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The ID of the filter.
     *
     * The following arguments will be discarded. Please use new fields as soon as possible:
     */
    public readonly trafficMirrorFilterId!: pulumi.Output<string>;
    /**
     * The ID of the outbound rule.
     */
    public /*out*/ readonly trafficMirrorFilterIngressRuleId!: pulumi.Output<string>;

    /**
     * Create a TrafficMirrorFilterIngressRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TrafficMirrorFilterIngressRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TrafficMirrorFilterIngressRuleArgs | TrafficMirrorFilterIngressRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TrafficMirrorFilterIngressRuleState | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["destinationCidrBlock"] = state ? state.destinationCidrBlock : undefined;
            resourceInputs["destinationPortRange"] = state ? state.destinationPortRange : undefined;
            resourceInputs["dryRun"] = state ? state.dryRun : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["ruleAction"] = state ? state.ruleAction : undefined;
            resourceInputs["sourceCidrBlock"] = state ? state.sourceCidrBlock : undefined;
            resourceInputs["sourcePortRange"] = state ? state.sourcePortRange : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["trafficMirrorFilterId"] = state ? state.trafficMirrorFilterId : undefined;
            resourceInputs["trafficMirrorFilterIngressRuleId"] = state ? state.trafficMirrorFilterIngressRuleId : undefined;
        } else {
            const args = argsOrState as TrafficMirrorFilterIngressRuleArgs | undefined;
            if ((!args || args.destinationCidrBlock === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destinationCidrBlock'");
            }
            if ((!args || args.priority === undefined) && !opts.urn) {
                throw new Error("Missing required property 'priority'");
            }
            if ((!args || args.protocol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocol'");
            }
            if ((!args || args.sourceCidrBlock === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceCidrBlock'");
            }
            if ((!args || args.trafficMirrorFilterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'trafficMirrorFilterId'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["destinationCidrBlock"] = args ? args.destinationCidrBlock : undefined;
            resourceInputs["destinationPortRange"] = args ? args.destinationPortRange : undefined;
            resourceInputs["dryRun"] = args ? args.dryRun : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["ruleAction"] = args ? args.ruleAction : undefined;
            resourceInputs["sourceCidrBlock"] = args ? args.sourceCidrBlock : undefined;
            resourceInputs["sourcePortRange"] = args ? args.sourcePortRange : undefined;
            resourceInputs["trafficMirrorFilterId"] = args ? args.trafficMirrorFilterId : undefined;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["trafficMirrorFilterIngressRuleId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TrafficMirrorFilterIngressRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TrafficMirrorFilterIngressRule resources.
 */
export interface TrafficMirrorFilterIngressRuleState {
    /**
     * The collection policy of the inbound rule. Valid values: `accept` or `drop`. `accept`: collects network traffic. `drop`: does not collect network traffic.
     */
    action?: pulumi.Input<string>;
    /**
     * The destination CIDR block of the inbound traffic.
     */
    destinationCidrBlock?: pulumi.Input<string>;
    /**
     * The destination CIDR block of the inbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
     */
    destinationPortRange?: pulumi.Input<string>;
    /**
     * Whether to PreCheck this request only. Value:
     * - **true**: sends a check request and does not create inbound or outbound rules. Check items include whether required parameters are filled in, request format, and restrictions. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.
     * - **false** (default): Sends a normal request and directly creates an inbound or outbound direction rule after checking.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * The priority of the inbound rule. A smaller value indicates a higher priority. The maximum value is `10`, which indicates that you can configure at most 10 inbound rules for a filter.
     */
    priority?: pulumi.Input<number>;
    /**
     * The transport protocol used by inbound traffic that needs to be mirrored. Valid values: `ALL`, `ICMP`, `TCP`, `UDP`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * . Field 'rule_action' has been deprecated from provider version 1.211.0. New field 'action' instead.
     *
     * @deprecated Field 'rule_action' has been deprecated since provider version 1.211.0. New field 'action' instead.
     */
    ruleAction?: pulumi.Input<string>;
    /**
     * The source CIDR block of the inbound traffic.
     */
    sourceCidrBlock?: pulumi.Input<string>;
    /**
     * The source port range of the inbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
     */
    sourcePortRange?: pulumi.Input<string>;
    /**
     * The state of the inbound rule. `Creating`, `Created`, `Modifying` and `Deleting`.
     */
    status?: pulumi.Input<string>;
    /**
     * The ID of the filter.
     *
     * The following arguments will be discarded. Please use new fields as soon as possible:
     */
    trafficMirrorFilterId?: pulumi.Input<string>;
    /**
     * The ID of the outbound rule.
     */
    trafficMirrorFilterIngressRuleId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TrafficMirrorFilterIngressRule resource.
 */
export interface TrafficMirrorFilterIngressRuleArgs {
    /**
     * The collection policy of the inbound rule. Valid values: `accept` or `drop`. `accept`: collects network traffic. `drop`: does not collect network traffic.
     */
    action?: pulumi.Input<string>;
    /**
     * The destination CIDR block of the inbound traffic.
     */
    destinationCidrBlock: pulumi.Input<string>;
    /**
     * The destination CIDR block of the inbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
     */
    destinationPortRange?: pulumi.Input<string>;
    /**
     * Whether to PreCheck this request only. Value:
     * - **true**: sends a check request and does not create inbound or outbound rules. Check items include whether required parameters are filled in, request format, and restrictions. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.
     * - **false** (default): Sends a normal request and directly creates an inbound or outbound direction rule after checking.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * The priority of the inbound rule. A smaller value indicates a higher priority. The maximum value is `10`, which indicates that you can configure at most 10 inbound rules for a filter.
     */
    priority: pulumi.Input<number>;
    /**
     * The transport protocol used by inbound traffic that needs to be mirrored. Valid values: `ALL`, `ICMP`, `TCP`, `UDP`.
     */
    protocol: pulumi.Input<string>;
    /**
     * . Field 'rule_action' has been deprecated from provider version 1.211.0. New field 'action' instead.
     *
     * @deprecated Field 'rule_action' has been deprecated since provider version 1.211.0. New field 'action' instead.
     */
    ruleAction?: pulumi.Input<string>;
    /**
     * The source CIDR block of the inbound traffic.
     */
    sourceCidrBlock: pulumi.Input<string>;
    /**
     * The source port range of the inbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
     */
    sourcePortRange?: pulumi.Input<string>;
    /**
     * The ID of the filter.
     *
     * The following arguments will be discarded. Please use new fields as soon as possible:
     */
    trafficMirrorFilterId: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Sag Acl Rule resource. This topic describes how to configure an access control list (ACL) rule for a target Smart Access Gateway instance to permit or deny access to or from specified IP addresses in the ACL rule.
 *
 * For information about Sag Acl Rule and how to use it, see [What is access control list (ACL) rule](https://www.alibabacloud.com/help/doc-detail/111483.htm).
 *
 * > **NOTE:** Available in 1.60.0+
 *
 * > **NOTE:** Only the following regions support create Cloud Connect Network. [`cn-shanghai`, `cn-shanghai-finance-1`, `cn-hongkong`, `ap-southeast-1`, `ap-southeast-2`, `ap-southeast-3`, `ap-southeast-5`, `ap-northeast-1`, `eu-central-1`]
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultAcl = new alicloud.rocketmq.Acl("default", {
 *     sagCount: "0",
 * });
 * const defaultAclRule = new alicloud.rocketmq.AclRule("default", {
 *     aclId: defaultAcl.id,
 *     description: "tf-testSagAclRule",
 *     destCidr: "192.168.1.0/24",
 *     destPortRange: "-1/-1",
 *     direction: "in",
 *     ipProtocol: "ALL",
 *     policy: "accept",
 *     priority: 1,
 *     sourceCidr: "10.10.1.0/24",
 *     sourcePortRange: "-1/-1",
 * });
 * ```
 */
export class AclRule extends pulumi.CustomResource {
    /**
     * Get an existing AclRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AclRuleState, opts?: pulumi.CustomResourceOptions): AclRule {
        return new AclRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:rocketmq/aclRule:AclRule';

    /**
     * Returns true if the given object is an instance of AclRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AclRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AclRule.__pulumiType;
    }

    /**
     * The ID of the ACL.
     */
    public readonly aclId!: pulumi.Output<string>;
    /**
     * The description of the ACL rule. It must be 1 to 512 characters in length.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The destination address. It is an IPv4 address range in CIDR format. Default value: 0.0.0.0/0.
     */
    public readonly destCidr!: pulumi.Output<string>;
    /**
     * The range of the destination port. Valid value: 80/80.
     */
    public readonly destPortRange!: pulumi.Output<string>;
    /**
     * The direction of the ACL rule. Valid values: in|out.
     */
    public readonly direction!: pulumi.Output<string>;
    /**
     * The protocol used by the ACL rule. The value is not case sensitive.
     */
    public readonly ipProtocol!: pulumi.Output<string>;
    /**
     * The policy used by the ACL rule. Valid values: accept|drop.
     */
    public readonly policy!: pulumi.Output<string>;
    /**
     * The priority of the ACL rule. Value range: 1 to 100.
     */
    public readonly priority!: pulumi.Output<number | undefined>;
    /**
     * The source address. It is an IPv4 address range in the CIDR format. Default value: 0.0.0.0/0.
     */
    public readonly sourceCidr!: pulumi.Output<string>;
    /**
     * The range of the source port. Valid value: 80/80.
     */
    public readonly sourcePortRange!: pulumi.Output<string>;

    /**
     * Create a AclRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AclRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AclRuleArgs | AclRuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AclRuleState | undefined;
            inputs["aclId"] = state ? state.aclId : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["destCidr"] = state ? state.destCidr : undefined;
            inputs["destPortRange"] = state ? state.destPortRange : undefined;
            inputs["direction"] = state ? state.direction : undefined;
            inputs["ipProtocol"] = state ? state.ipProtocol : undefined;
            inputs["policy"] = state ? state.policy : undefined;
            inputs["priority"] = state ? state.priority : undefined;
            inputs["sourceCidr"] = state ? state.sourceCidr : undefined;
            inputs["sourcePortRange"] = state ? state.sourcePortRange : undefined;
        } else {
            const args = argsOrState as AclRuleArgs | undefined;
            if (!args || args.aclId === undefined) {
                throw new Error("Missing required property 'aclId'");
            }
            if (!args || args.destCidr === undefined) {
                throw new Error("Missing required property 'destCidr'");
            }
            if (!args || args.destPortRange === undefined) {
                throw new Error("Missing required property 'destPortRange'");
            }
            if (!args || args.direction === undefined) {
                throw new Error("Missing required property 'direction'");
            }
            if (!args || args.ipProtocol === undefined) {
                throw new Error("Missing required property 'ipProtocol'");
            }
            if (!args || args.policy === undefined) {
                throw new Error("Missing required property 'policy'");
            }
            if (!args || args.sourceCidr === undefined) {
                throw new Error("Missing required property 'sourceCidr'");
            }
            if (!args || args.sourcePortRange === undefined) {
                throw new Error("Missing required property 'sourcePortRange'");
            }
            inputs["aclId"] = args ? args.aclId : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["destCidr"] = args ? args.destCidr : undefined;
            inputs["destPortRange"] = args ? args.destPortRange : undefined;
            inputs["direction"] = args ? args.direction : undefined;
            inputs["ipProtocol"] = args ? args.ipProtocol : undefined;
            inputs["policy"] = args ? args.policy : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["sourceCidr"] = args ? args.sourceCidr : undefined;
            inputs["sourcePortRange"] = args ? args.sourcePortRange : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AclRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AclRule resources.
 */
export interface AclRuleState {
    /**
     * The ID of the ACL.
     */
    readonly aclId?: pulumi.Input<string>;
    /**
     * The description of the ACL rule. It must be 1 to 512 characters in length.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The destination address. It is an IPv4 address range in CIDR format. Default value: 0.0.0.0/0.
     */
    readonly destCidr?: pulumi.Input<string>;
    /**
     * The range of the destination port. Valid value: 80/80.
     */
    readonly destPortRange?: pulumi.Input<string>;
    /**
     * The direction of the ACL rule. Valid values: in|out.
     */
    readonly direction?: pulumi.Input<string>;
    /**
     * The protocol used by the ACL rule. The value is not case sensitive.
     */
    readonly ipProtocol?: pulumi.Input<string>;
    /**
     * The policy used by the ACL rule. Valid values: accept|drop.
     */
    readonly policy?: pulumi.Input<string>;
    /**
     * The priority of the ACL rule. Value range: 1 to 100.
     */
    readonly priority?: pulumi.Input<number>;
    /**
     * The source address. It is an IPv4 address range in the CIDR format. Default value: 0.0.0.0/0.
     */
    readonly sourceCidr?: pulumi.Input<string>;
    /**
     * The range of the source port. Valid value: 80/80.
     */
    readonly sourcePortRange?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AclRule resource.
 */
export interface AclRuleArgs {
    /**
     * The ID of the ACL.
     */
    readonly aclId: pulumi.Input<string>;
    /**
     * The description of the ACL rule. It must be 1 to 512 characters in length.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The destination address. It is an IPv4 address range in CIDR format. Default value: 0.0.0.0/0.
     */
    readonly destCidr: pulumi.Input<string>;
    /**
     * The range of the destination port. Valid value: 80/80.
     */
    readonly destPortRange: pulumi.Input<string>;
    /**
     * The direction of the ACL rule. Valid values: in|out.
     */
    readonly direction: pulumi.Input<string>;
    /**
     * The protocol used by the ACL rule. The value is not case sensitive.
     */
    readonly ipProtocol: pulumi.Input<string>;
    /**
     * The policy used by the ACL rule. Valid values: accept|drop.
     */
    readonly policy: pulumi.Input<string>;
    /**
     * The priority of the ACL rule. Value range: 1 to 100.
     */
    readonly priority?: pulumi.Input<number>;
    /**
     * The source address. It is an IPv4 address range in the CIDR format. Default value: 0.0.0.0/0.
     */
    readonly sourceCidr: pulumi.Input<string>;
    /**
     * The range of the source port. Valid value: 80/80.
     */
    readonly sourcePortRange: pulumi.Input<string>;
}

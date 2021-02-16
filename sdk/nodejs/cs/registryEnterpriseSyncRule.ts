// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This resource will help you to manager Container Registry Enterprise Edition sync rules.
 *
 * For information about Container Registry Enterprise Edition sync rules and how to use it, see [Create a Sync Rule](https://www.alibabacloud.com/help/doc-detail/145280.htm)
 *
 * > **NOTE:** Available in v1.90.0+.
 *
 * > **NOTE:** You need to set your registry password in Container Registry Enterprise Edition console before use this resource.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultRegistryEnterpriseSyncRule = new alicloud.cs.RegistryEnterpriseSyncRule("default", {
 *     instanceId: "my-source-instance-id",
 *     namespaceName: "my-source-namespace",
 *     repoName: "my-source-repo",
 *     tagFilter: ".*",
 *     targetInstanceId: "my-target-instance-id",
 *     targetNamespaceName: "my-target-namespace",
 *     targetRegionId: "cn-hangzhou",
 *     targetRepoName: "my-target-repo",
 * });
 * ```
 *
 * ## Import
 *
 * Container Registry Enterprise Edition sync rule can be imported using the id. Format to `{instance_id}:{namespace_name}:{rule_id}`, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:cs/registryEnterpriseSyncRule:RegistryEnterpriseSyncRule default `cri-xxx:my-namespace:crsr-yyy`
 * ```
 */
export class RegistryEnterpriseSyncRule extends pulumi.CustomResource {
    /**
     * Get an existing RegistryEnterpriseSyncRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RegistryEnterpriseSyncRuleState, opts?: pulumi.CustomResourceOptions): RegistryEnterpriseSyncRule {
        return new RegistryEnterpriseSyncRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cs/registryEnterpriseSyncRule:RegistryEnterpriseSyncRule';

    /**
     * Returns true if the given object is an instance of RegistryEnterpriseSyncRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegistryEnterpriseSyncRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegistryEnterpriseSyncRule.__pulumiType;
    }

    /**
     * ID of Container Registry Enterprise Edition source instance.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Name of Container Registry Enterprise Edition sync rule.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Name of Container Registry Enterprise Edition source namespace. It can contain 2 to 30 characters.
     */
    public readonly namespaceName!: pulumi.Output<string>;
    /**
     * Name of the source repository which should be set together with `targetRepoName`, if empty means that the synchronization scope is the entire namespace level.
     */
    public readonly repoName!: pulumi.Output<string | undefined>;
    /**
     * The uuid of Container Registry Enterprise Edition sync rule.
     */
    public /*out*/ readonly ruleId!: pulumi.Output<string>;
    /**
     * `FROM` or `TO`, the direction of synchronization. `FROM` means source instance, `TO` means target instance.
     */
    public /*out*/ readonly syncDirection!: pulumi.Output<string>;
    /**
     * `REPO` or `NAMESPACE`,the scope that the synchronization rule applies.
     */
    public /*out*/ readonly syncScope!: pulumi.Output<string>;
    /**
     * The regular expression used to filter image tags for synchronization in the source repository.
     */
    public readonly tagFilter!: pulumi.Output<string>;
    /**
     * ID of Container Registry Enterprise Edition target instance to be synchronized.
     */
    public readonly targetInstanceId!: pulumi.Output<string>;
    /**
     * Name of Container Registry Enterprise Edition target namespace to be synchronized. It can contain 2 to 30 characters.
     */
    public readonly targetNamespaceName!: pulumi.Output<string>;
    /**
     * The target region to be synchronized.
     */
    public readonly targetRegionId!: pulumi.Output<string>;
    /**
     * Name of the target repository.
     */
    public readonly targetRepoName!: pulumi.Output<string | undefined>;

    /**
     * Create a RegistryEnterpriseSyncRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegistryEnterpriseSyncRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RegistryEnterpriseSyncRuleArgs | RegistryEnterpriseSyncRuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RegistryEnterpriseSyncRuleState | undefined;
            inputs["instanceId"] = state ? state.instanceId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namespaceName"] = state ? state.namespaceName : undefined;
            inputs["repoName"] = state ? state.repoName : undefined;
            inputs["ruleId"] = state ? state.ruleId : undefined;
            inputs["syncDirection"] = state ? state.syncDirection : undefined;
            inputs["syncScope"] = state ? state.syncScope : undefined;
            inputs["tagFilter"] = state ? state.tagFilter : undefined;
            inputs["targetInstanceId"] = state ? state.targetInstanceId : undefined;
            inputs["targetNamespaceName"] = state ? state.targetNamespaceName : undefined;
            inputs["targetRegionId"] = state ? state.targetRegionId : undefined;
            inputs["targetRepoName"] = state ? state.targetRepoName : undefined;
        } else {
            const args = argsOrState as RegistryEnterpriseSyncRuleArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.namespaceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'namespaceName'");
            }
            if ((!args || args.tagFilter === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tagFilter'");
            }
            if ((!args || args.targetInstanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetInstanceId'");
            }
            if ((!args || args.targetNamespaceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetNamespaceName'");
            }
            if ((!args || args.targetRegionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetRegionId'");
            }
            inputs["instanceId"] = args ? args.instanceId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namespaceName"] = args ? args.namespaceName : undefined;
            inputs["repoName"] = args ? args.repoName : undefined;
            inputs["tagFilter"] = args ? args.tagFilter : undefined;
            inputs["targetInstanceId"] = args ? args.targetInstanceId : undefined;
            inputs["targetNamespaceName"] = args ? args.targetNamespaceName : undefined;
            inputs["targetRegionId"] = args ? args.targetRegionId : undefined;
            inputs["targetRepoName"] = args ? args.targetRepoName : undefined;
            inputs["ruleId"] = undefined /*out*/;
            inputs["syncDirection"] = undefined /*out*/;
            inputs["syncScope"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(RegistryEnterpriseSyncRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RegistryEnterpriseSyncRule resources.
 */
export interface RegistryEnterpriseSyncRuleState {
    /**
     * ID of Container Registry Enterprise Edition source instance.
     */
    readonly instanceId?: pulumi.Input<string>;
    /**
     * Name of Container Registry Enterprise Edition sync rule.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Name of Container Registry Enterprise Edition source namespace. It can contain 2 to 30 characters.
     */
    readonly namespaceName?: pulumi.Input<string>;
    /**
     * Name of the source repository which should be set together with `targetRepoName`, if empty means that the synchronization scope is the entire namespace level.
     */
    readonly repoName?: pulumi.Input<string>;
    /**
     * The uuid of Container Registry Enterprise Edition sync rule.
     */
    readonly ruleId?: pulumi.Input<string>;
    /**
     * `FROM` or `TO`, the direction of synchronization. `FROM` means source instance, `TO` means target instance.
     */
    readonly syncDirection?: pulumi.Input<string>;
    /**
     * `REPO` or `NAMESPACE`,the scope that the synchronization rule applies.
     */
    readonly syncScope?: pulumi.Input<string>;
    /**
     * The regular expression used to filter image tags for synchronization in the source repository.
     */
    readonly tagFilter?: pulumi.Input<string>;
    /**
     * ID of Container Registry Enterprise Edition target instance to be synchronized.
     */
    readonly targetInstanceId?: pulumi.Input<string>;
    /**
     * Name of Container Registry Enterprise Edition target namespace to be synchronized. It can contain 2 to 30 characters.
     */
    readonly targetNamespaceName?: pulumi.Input<string>;
    /**
     * The target region to be synchronized.
     */
    readonly targetRegionId?: pulumi.Input<string>;
    /**
     * Name of the target repository.
     */
    readonly targetRepoName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RegistryEnterpriseSyncRule resource.
 */
export interface RegistryEnterpriseSyncRuleArgs {
    /**
     * ID of Container Registry Enterprise Edition source instance.
     */
    readonly instanceId: pulumi.Input<string>;
    /**
     * Name of Container Registry Enterprise Edition sync rule.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Name of Container Registry Enterprise Edition source namespace. It can contain 2 to 30 characters.
     */
    readonly namespaceName: pulumi.Input<string>;
    /**
     * Name of the source repository which should be set together with `targetRepoName`, if empty means that the synchronization scope is the entire namespace level.
     */
    readonly repoName?: pulumi.Input<string>;
    /**
     * The regular expression used to filter image tags for synchronization in the source repository.
     */
    readonly tagFilter: pulumi.Input<string>;
    /**
     * ID of Container Registry Enterprise Edition target instance to be synchronized.
     */
    readonly targetInstanceId: pulumi.Input<string>;
    /**
     * Name of Container Registry Enterprise Edition target namespace to be synchronized. It can contain 2 to 30 characters.
     */
    readonly targetNamespaceName: pulumi.Input<string>;
    /**
     * The target region to be synchronized.
     */
    readonly targetRegionId: pulumi.Input<string>;
    /**
     * Name of the target repository.
     */
    readonly targetRepoName?: pulumi.Input<string>;
}

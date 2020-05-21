// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides a list of RAM policies in an Alibaba Cloud account according to the specified filters.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const policiesDs = pulumi.output(alicloud.ram.getPolicies({
 *     groupName: "group1",
 *     outputFile: "policies.txt",
 *     type: "System",
 *     userName: "user1",
 * }, { async: true }));
 *
 * export const firstPolicyName = policiesDs.policies[0].name;
 * ```
 */
export function getPolicies(args?: GetPoliciesArgs, opts?: pulumi.InvokeOptions): Promise<GetPoliciesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:ram/getPolicies:getPolicies", {
        "groupName": args.groupName,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "roleName": args.roleName,
        "type": args.type,
        "userName": args.userName,
    }, opts);
}

/**
 * A collection of arguments for invoking getPolicies.
 */
export interface GetPoliciesArgs {
    /**
     * Filter results by a specific group name. Returned policies are attached to the specified group.
     */
    readonly groupName?: string;
    /**
     * A regex string to filter resulting policies by name.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * Filter results by a specific role name. Returned policies are attached to the specified role.
     */
    readonly roleName?: string;
    /**
     * Filter results by a specific policy type. Valid values are `Custom` and `System`.
     */
    readonly type?: string;
    /**
     * Filter results by a specific user name. Returned policies are attached to the specified user.
     */
    readonly userName?: string;
}

/**
 * A collection of values returned by getPolicies.
 */
export interface GetPoliciesResult {
    readonly groupName?: string;
    readonly nameRegex?: string;
    /**
     * A list of ram group names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * A list of policies. Each element contains the following attributes:
     */
    readonly policies: outputs.ram.GetPoliciesPolicy[];
    readonly roleName?: string;
    /**
     * Type of the policy.
     */
    readonly type?: string;
    readonly userName?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}

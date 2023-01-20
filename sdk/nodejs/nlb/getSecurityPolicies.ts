// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Nlb Security Policies of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.187.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.nlb.getSecurityPolicies({});
 * export const nlbSecurityPolicyId1 = ids.then(ids => ids.policies?.[0]?.id);
 * const nameRegex = alicloud.nlb.getSecurityPolicies({
 *     nameRegex: "^my-SecurityPolicy",
 * });
 * export const nlbSecurityPolicyId2 = nameRegex.then(nameRegex => nameRegex.policies?.[0]?.id);
 * ```
 */
export function getSecurityPolicies(args?: GetSecurityPoliciesArgs, opts?: pulumi.InvokeOptions): Promise<GetSecurityPoliciesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:nlb/getSecurityPolicies:getSecurityPolicies", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "resourceGroupId": args.resourceGroupId,
        "securityPolicyNames": args.securityPolicyNames,
        "status": args.status,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getSecurityPolicies.
 */
export interface GetSecurityPoliciesArgs {
    /**
     * A list of Security Policy IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter results by Security Policy name.
     */
    nameRegex?: string;
    outputFile?: string;
    /**
     * The ID of the resource group.
     */
    resourceGroupId?: string;
    /**
     * The names of the TLS security policies.
     */
    securityPolicyNames?: string[];
    /**
     * The status of the resource.
     */
    status?: string;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getSecurityPolicies.
 */
export interface GetSecurityPoliciesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly policies: outputs.nlb.GetSecurityPoliciesPolicy[];
    readonly resourceGroupId?: string;
    readonly securityPolicyNames?: string[];
    readonly status?: string;
    readonly tags?: {[key: string]: any};
}
/**
 * This data source provides the Nlb Security Policies of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.187.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.nlb.getSecurityPolicies({});
 * export const nlbSecurityPolicyId1 = ids.then(ids => ids.policies?.[0]?.id);
 * const nameRegex = alicloud.nlb.getSecurityPolicies({
 *     nameRegex: "^my-SecurityPolicy",
 * });
 * export const nlbSecurityPolicyId2 = nameRegex.then(nameRegex => nameRegex.policies?.[0]?.id);
 * ```
 */
export function getSecurityPoliciesOutput(args?: GetSecurityPoliciesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSecurityPoliciesResult> {
    return pulumi.output(args).apply((a: any) => getSecurityPolicies(a, opts))
}

/**
 * A collection of arguments for invoking getSecurityPolicies.
 */
export interface GetSecurityPoliciesOutputArgs {
    /**
     * A list of Security Policy IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by Security Policy name.
     */
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
    /**
     * The ID of the resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The names of the TLS security policies.
     */
    securityPolicyNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The status of the resource.
     */
    status?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}

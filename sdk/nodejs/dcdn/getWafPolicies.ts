// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Dcdn Waf Policies of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.184.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.dcdn.getWafPolicies({});
 * export const dcdnWafPolicyId1 = ids.then(ids => ids.policies?.[0]?.id);
 * ```
 */
export function getWafPolicies(args?: GetWafPoliciesArgs, opts?: pulumi.InvokeOptions): Promise<GetWafPoliciesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:dcdn/getWafPolicies:getWafPolicies", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "queryArgs": args.queryArgs,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getWafPolicies.
 */
export interface GetWafPoliciesArgs {
    /**
     * A list of Waf Policy IDs.
     */
    ids?: string[];
    nameRegex?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * The query conditions. The value is a string in the JSON format. Format: `{"PolicyIds":"The ID of the proteuleIds":"Thection policy","R range of protection rule IDs","PolicyNameLike":"The name of the protection policy","DomainNames":"The protected domain names","PolicyType":"default","DefenseScenes":"wafGroup","PolicyStatus":"on","OrderBy":"GmtModified","Desc":"false"}`.
     */
    queryArgs?: string;
    /**
     * The status of the resource.
     */
    status?: string;
}

/**
 * A collection of values returned by getWafPolicies.
 */
export interface GetWafPoliciesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly policies: outputs.dcdn.GetWafPoliciesPolicy[];
    readonly queryArgs?: string;
    readonly status?: string;
}
/**
 * This data source provides the Dcdn Waf Policies of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.184.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.dcdn.getWafPolicies({});
 * export const dcdnWafPolicyId1 = ids.then(ids => ids.policies?.[0]?.id);
 * ```
 */
export function getWafPoliciesOutput(args?: GetWafPoliciesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetWafPoliciesResult> {
    return pulumi.output(args).apply((a: any) => getWafPolicies(a, opts))
}

/**
 * A collection of arguments for invoking getWafPolicies.
 */
export interface GetWafPoliciesOutputArgs {
    /**
     * A list of Waf Policy IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The query conditions. The value is a string in the JSON format. Format: `{"PolicyIds":"The ID of the proteuleIds":"Thection policy","R range of protection rule IDs","PolicyNameLike":"The name of the protection policy","DomainNames":"The protected domain names","PolicyType":"default","DefenseScenes":"wafGroup","PolicyStatus":"on","OrderBy":"GmtModified","Desc":"false"}`.
     */
    queryArgs?: pulumi.Input<string>;
    /**
     * The status of the resource.
     */
    status?: pulumi.Input<string>;
}

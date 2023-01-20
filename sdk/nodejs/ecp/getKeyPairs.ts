// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Ecp Key Pairs of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.130.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.ecp.getKeyPairs({});
 * export const ecpKeyPairId1 = ids.then(ids => ids.pairs?.[0]?.id);
 * const nameRegex = alicloud.ecp.getKeyPairs({
 *     nameRegex: "^my-KeyPair",
 * });
 * export const ecpKeyPairId2 = nameRegex.then(nameRegex => nameRegex.pairs?.[0]?.id);
 * ```
 */
export function getKeyPairs(args?: GetKeyPairsArgs, opts?: pulumi.InvokeOptions): Promise<GetKeyPairsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:ecp/getKeyPairs:getKeyPairs", {
        "ids": args.ids,
        "keyPairFingerPrint": args.keyPairFingerPrint,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getKeyPairs.
 */
export interface GetKeyPairsArgs {
    /**
     * A list of Key Pair IDs. Its element value is same as Key Pair Name.
     */
    ids?: string[];
    /**
     * The Private Key of the Fingerprint.
     */
    keyPairFingerPrint?: string;
    /**
     * A regex string to filter results by Key Pair name.
     */
    nameRegex?: string;
    outputFile?: string;
}

/**
 * A collection of values returned by getKeyPairs.
 */
export interface GetKeyPairsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly keyPairFingerPrint?: string;
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly pairs: outputs.ecp.GetKeyPairsPair[];
}
/**
 * This data source provides the Ecp Key Pairs of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.130.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.ecp.getKeyPairs({});
 * export const ecpKeyPairId1 = ids.then(ids => ids.pairs?.[0]?.id);
 * const nameRegex = alicloud.ecp.getKeyPairs({
 *     nameRegex: "^my-KeyPair",
 * });
 * export const ecpKeyPairId2 = nameRegex.then(nameRegex => nameRegex.pairs?.[0]?.id);
 * ```
 */
export function getKeyPairsOutput(args?: GetKeyPairsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetKeyPairsResult> {
    return pulumi.output(args).apply((a: any) => getKeyPairs(a, opts))
}

/**
 * A collection of arguments for invoking getKeyPairs.
 */
export interface GetKeyPairsOutputArgs {
    /**
     * A list of Key Pair IDs. Its element value is same as Key Pair Name.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Private Key of the Fingerprint.
     */
    keyPairFingerPrint?: pulumi.Input<string>;
    /**
     * A regex string to filter results by Key Pair name.
     */
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
}

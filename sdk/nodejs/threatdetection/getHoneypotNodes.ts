// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides Threat Detection Honeypot Node available to the user.[What is Honeypot Node](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-createhoneypotnode)
 *
 * > **NOTE:** Available in 1.195.0+
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const default = alicloud.threatdetection.getHoneypotNodes({
 *     ids: [alicloud_threat_detection_honeypot_node["default"].id],
 * });
 * export const alicloudThreatDetectionHoneypotNodeExampleId = _default.then(_default => _default.nodes?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getHoneypotNodes(args?: GetHoneypotNodesArgs, opts?: pulumi.InvokeOptions): Promise<GetHoneypotNodesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:threatdetection/getHoneypotNodes:getHoneypotNodes", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "nodeId": args.nodeId,
        "nodeName": args.nodeName,
        "outputFile": args.outputFile,
        "pageNumber": args.pageNumber,
        "pageSize": args.pageSize,
    }, opts);
}

/**
 * A collection of arguments for invoking getHoneypotNodes.
 */
export interface GetHoneypotNodesArgs {
    /**
     * A list of Honeypot Node IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter results by Honeypot Node name.
     */
    nameRegex?: string;
    /**
     * Honeypot management node id.
     */
    nodeId?: string;
    /**
     * The name of the management node.
     */
    nodeName?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    pageNumber?: number;
    pageSize?: number;
}

/**
 * A collection of values returned by getHoneypotNodes.
 */
export interface GetHoneypotNodesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of Honeypot Node IDs.
     */
    readonly ids: string[];
    readonly nameRegex?: string;
    /**
     * A list of Honeypot Node names.
     */
    readonly names: string[];
    /**
     * Honeypot management node id.
     */
    readonly nodeId?: string;
    /**
     * Management node name.
     */
    readonly nodeName?: string;
    /**
     * A list of Honeypot Node Entries. Each element contains the following attributes:
     */
    readonly nodes: outputs.threatdetection.GetHoneypotNodesNode[];
    readonly outputFile?: string;
    readonly pageNumber?: number;
    readonly pageSize?: number;
}
/**
 * This data source provides Threat Detection Honeypot Node available to the user.[What is Honeypot Node](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-createhoneypotnode)
 *
 * > **NOTE:** Available in 1.195.0+
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const default = alicloud.threatdetection.getHoneypotNodes({
 *     ids: [alicloud_threat_detection_honeypot_node["default"].id],
 * });
 * export const alicloudThreatDetectionHoneypotNodeExampleId = _default.then(_default => _default.nodes?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getHoneypotNodesOutput(args?: GetHoneypotNodesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetHoneypotNodesResult> {
    return pulumi.output(args).apply((a: any) => getHoneypotNodes(a, opts))
}

/**
 * A collection of arguments for invoking getHoneypotNodes.
 */
export interface GetHoneypotNodesOutputArgs {
    /**
     * A list of Honeypot Node IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by Honeypot Node name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * Honeypot management node id.
     */
    nodeId?: pulumi.Input<string>;
    /**
     * The name of the management node.
     */
    nodeName?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    pageNumber?: pulumi.Input<number>;
    pageSize?: pulumi.Input<number>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Ecd Commands of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.146.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultSimpleOfficeSite = new alicloud.eds.SimpleOfficeSite("default", {
 *     cidrBlock: "172.16.0.0/12",
 *     desktopAccessType: "Internet",
 *     officeSiteName: "your_office_site_name",
 * });
 * const default = alicloud.eds.getBundles({
 *     bundleType: "SYSTEM",
 *     nameRegex: "windows",
 * });
 * const defaultEcdPolicyGroup = new alicloud.eds.EcdPolicyGroup("default", {
 *     policyGroupName: "your_policy_group_name",
 *     clipboard: "readwrite",
 *     localDrive: "read",
 *     authorizeAccessPolicyRules: [{
 *         description: "example_value",
 *         cidrIp: "1.2.3.4/24",
 *     }],
 *     authorizeSecurityPolicyRules: [{
 *         type: "inflow",
 *         policy: "accept",
 *         description: "example_value",
 *         portRange: "80/80",
 *         ipProtocol: "TCP",
 *         priority: "1",
 *         cidrIp: "0.0.0.0/0",
 *     }],
 * });
 * const defaultDesktop = new alicloud.eds.Desktop("default", {
 *     officeSiteId: defaultSimpleOfficeSite.id,
 *     policyGroupId: defaultEcdPolicyGroup.id,
 *     bundleId: _default.then(_default => _default.bundles?.[0]?.id),
 *     desktopName: name,
 * });
 * const defaultCommand = new alicloud.eds.Command("default", {
 *     commandContent: "ipconfig",
 *     commandType: "RunPowerShellScript",
 *     desktopId: defaultDesktop.id,
 * });
 * const ids = alicloud.eds.getCommands({});
 * export const ecdCommandId1 = ids.then(ids => ids.commands?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getCommands(args?: GetCommandsArgs, opts?: pulumi.InvokeOptions): Promise<GetCommandsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:eds/getCommands:getCommands", {
        "commandType": args.commandType,
        "contentEncoding": args.contentEncoding,
        "desktopId": args.desktopId,
        "ids": args.ids,
        "outputFile": args.outputFile,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getCommands.
 */
export interface GetCommandsArgs {
    /**
     * The Script Type. Valid values: `RunBatScript`, `RunPowerShellScript`.
     */
    commandType?: string;
    /**
     * That Returns the Data Encoding Method. Valid values: `Base64`, `PlainText`.
     */
    contentEncoding?: string;
    /**
     * The desktop id of the Desktop.
     */
    desktopId?: string;
    /**
     * A list of Command IDs.
     */
    ids?: string[];
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * Script Is Executed in the Overall Implementation of the State. Valid values: `Pending`, `Failed`, `PartialFailed`, `Running`, `Stopped`, `Stopping`, `Finished`, `Success`.
     */
    status?: string;
}

/**
 * A collection of values returned by getCommands.
 */
export interface GetCommandsResult {
    readonly commandType?: string;
    readonly commands: outputs.eds.GetCommandsCommand[];
    readonly contentEncoding?: string;
    readonly desktopId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly outputFile?: string;
    readonly status?: string;
}
/**
 * This data source provides the Ecd Commands of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.146.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultSimpleOfficeSite = new alicloud.eds.SimpleOfficeSite("default", {
 *     cidrBlock: "172.16.0.0/12",
 *     desktopAccessType: "Internet",
 *     officeSiteName: "your_office_site_name",
 * });
 * const default = alicloud.eds.getBundles({
 *     bundleType: "SYSTEM",
 *     nameRegex: "windows",
 * });
 * const defaultEcdPolicyGroup = new alicloud.eds.EcdPolicyGroup("default", {
 *     policyGroupName: "your_policy_group_name",
 *     clipboard: "readwrite",
 *     localDrive: "read",
 *     authorizeAccessPolicyRules: [{
 *         description: "example_value",
 *         cidrIp: "1.2.3.4/24",
 *     }],
 *     authorizeSecurityPolicyRules: [{
 *         type: "inflow",
 *         policy: "accept",
 *         description: "example_value",
 *         portRange: "80/80",
 *         ipProtocol: "TCP",
 *         priority: "1",
 *         cidrIp: "0.0.0.0/0",
 *     }],
 * });
 * const defaultDesktop = new alicloud.eds.Desktop("default", {
 *     officeSiteId: defaultSimpleOfficeSite.id,
 *     policyGroupId: defaultEcdPolicyGroup.id,
 *     bundleId: _default.then(_default => _default.bundles?.[0]?.id),
 *     desktopName: name,
 * });
 * const defaultCommand = new alicloud.eds.Command("default", {
 *     commandContent: "ipconfig",
 *     commandType: "RunPowerShellScript",
 *     desktopId: defaultDesktop.id,
 * });
 * const ids = alicloud.eds.getCommands({});
 * export const ecdCommandId1 = ids.then(ids => ids.commands?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getCommandsOutput(args?: GetCommandsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCommandsResult> {
    return pulumi.output(args).apply((a: any) => getCommands(a, opts))
}

/**
 * A collection of arguments for invoking getCommands.
 */
export interface GetCommandsOutputArgs {
    /**
     * The Script Type. Valid values: `RunBatScript`, `RunPowerShellScript`.
     */
    commandType?: pulumi.Input<string>;
    /**
     * That Returns the Data Encoding Method. Valid values: `Base64`, `PlainText`.
     */
    contentEncoding?: pulumi.Input<string>;
    /**
     * The desktop id of the Desktop.
     */
    desktopId?: pulumi.Input<string>;
    /**
     * A list of Command IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * Script Is Executed in the Overall Implementation of the State. Valid values: `Pending`, `Failed`, `PartialFailed`, `Running`, `Stopped`, `Stopping`, `Finished`, `Success`.
     */
    status?: pulumi.Input<string>;
}

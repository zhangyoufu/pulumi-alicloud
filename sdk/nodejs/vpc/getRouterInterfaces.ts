// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides information about [router interfaces](https://www.alibabacloud.com/help/doc-detail/52412.htm)
 * that connect VPCs together.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const routerInterfacesDs = pulumi.output(alicloud.vpc.getRouterInterfaces({
 *     nameRegex: "^testenv",
 *     status: "Active",
 * }, { async: true }));
 * 
 * export const firstRouterInterfaceId = routerInterfacesDs.interfaces[0].id;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/router_interfaces.html.markdown.
 */
export function getRouterInterfaces(args?: GetRouterInterfacesArgs, opts?: pulumi.InvokeOptions): Promise<GetRouterInterfacesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:vpc/getRouterInterfaces:getRouterInterfaces", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "oppositeInterfaceId": args.oppositeInterfaceId,
        "oppositeInterfaceOwnerId": args.oppositeInterfaceOwnerId,
        "outputFile": args.outputFile,
        "role": args.role,
        "routerId": args.routerId,
        "routerType": args.routerType,
        "specification": args.specification,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getRouterInterfaces.
 */
export interface GetRouterInterfacesArgs {
    /**
     * A list of router interface IDs.
     */
    readonly ids?: string[];
    /**
     * A regex string used to filter by router interface name.
     */
    readonly nameRegex?: string;
    /**
     * ID of the peer router interface.
     */
    readonly oppositeInterfaceId?: string;
    /**
     * Account ID of the owner of the peer router interface.
     */
    readonly oppositeInterfaceOwnerId?: string;
    readonly outputFile?: string;
    /**
     * Role of the router interface. Valid values are `InitiatingSide` (connection initiator) and 
     * `AcceptingSide` (connection receiver). The value of this parameter must be `InitiatingSide` if the `routerType` is set to `VBR`.
     */
    readonly role?: string;
    /**
     * ID of the VRouter located in the local region.
     */
    readonly routerId?: string;
    /**
     * Router type in the local region. Valid values are `VRouter` and `VBR` (physical connection).
     */
    readonly routerType?: string;
    /**
     * Specification of the link, such as `Small.1` (10Mb), `Middle.1` (100Mb), `Large.2` (2Gb), ...etc.
     */
    readonly specification?: string;
    /**
     * Expected status. Valid values are `Active`, `Inactive` and `Idle`.
     */
    readonly status?: string;
}

/**
 * A collection of values returned by getRouterInterfaces.
 */
export interface GetRouterInterfacesResult {
    /**
     * A list of router interface IDs.
     */
    readonly ids: string[];
    /**
     * A list of router interfaces. Each element contains the following attributes:
     */
    readonly interfaces: outputs.vpc.GetRouterInterfacesInterface[];
    readonly nameRegex?: string;
    /**
     * A list of router interface names.
     */
    readonly names: string[];
    /**
     * Peer router interface ID.
     */
    readonly oppositeInterfaceId?: string;
    /**
     * Account ID of the owner of the peer router interface.
     */
    readonly oppositeInterfaceOwnerId?: string;
    readonly outputFile?: string;
    /**
     * Router interface role. Possible values: `InitiatingSide` and `AcceptingSide`.
     */
    readonly role?: string;
    /**
     * ID of the VRouter located in the local region.
     */
    readonly routerId?: string;
    /**
     * Router type in the local region. Possible values: `VRouter` and `VBR`.
     */
    readonly routerType?: string;
    /**
     * Router interface specification. Possible values: `Small.1`, `Middle.1`, `Large.2`, ...etc.
     */
    readonly specification?: string;
    /**
     * Router interface status. Possible values: `Active`, `Inactive` and `Idle`.
     */
    readonly status?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}

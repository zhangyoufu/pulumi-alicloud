// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides a list of Forward Entries owned by an Alibaba Cloud account.
 * 
 * > **NOTE:** Available in 1.37.0+.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const config = new pulumi.Config();
 * const name = config.get("name") || "forward-entry-config-example-name";
 * 
 * const defaultZones = alicloud.getZones({
 *     availableResourceCreation: "VSwitch",
 * });
 * const defaultNetwork = new alicloud.vpc.Network("default", {
 *     cidrBlock: "172.16.0.0/12",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("default", {
 *     availabilityZone: defaultZones.zones[0].id,
 *     cidrBlock: "172.16.0.0/21",
 *     vpcId: defaultNetwork.id,
 * });
 * const defaultNatGateway = new alicloud.vpc.NatGateway("default", {
 *     specification: "Small",
 *     vpcId: defaultNetwork.id,
 * });
 * const defaultEip = new alicloud.ecs.Eip("default", {});
 * const defaultEipAssociation = new alicloud.ecs.EipAssociation("default", {
 *     allocationId: defaultEip.id,
 *     instanceId: defaultNatGateway.id,
 * });
 * const defaultForwardEntry = new alicloud.vpc.ForwardEntry("default", {
 *     externalIp: defaultEip.ipAddress,
 *     externalPort: "80",
 *     forwardTableId: defaultNatGateway.forwardTableIds,
 *     internalIp: "172.16.0.3",
 *     internalPort: "8080",
 *     ipProtocol: "tcp",
 * });
 * const defaultForwardEntries = defaultForwardEntry.forwardTableId.apply(forwardTableId => alicloud.vpc.getForwardEntries({
 *     forwardTableId: forwardTableId,
 * }));
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/forward_entries.html.markdown.
 */
export function getForwardEntries(args: GetForwardEntriesArgs, opts?: pulumi.InvokeOptions): Promise<GetForwardEntriesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:vpc/getForwardEntries:getForwardEntries", {
        "externalIp": args.externalIp,
        "forwardTableId": args.forwardTableId,
        "ids": args.ids,
        "internalIp": args.internalIp,
        "nameRegex": args.nameRegex,
        "names": args.names,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getForwardEntries.
 */
export interface GetForwardEntriesArgs {
    /**
     * The public IP address.
     */
    readonly externalIp?: string;
    /**
     * The ID of the Forward table.
     */
    readonly forwardTableId: string;
    /**
     * A list of Forward Entries IDs.
     */
    readonly ids?: string[];
    /**
     * The private IP address.
     */
    readonly internalIp?: string;
    /**
     * A regex string to filter results by forward entry name.
     */
    readonly nameRegex?: string;
    /**
     * A list of Forward Entries names.
     */
    readonly names?: string[];
    readonly outputFile?: string;
}

/**
 * A collection of values returned by getForwardEntries.
 */
export interface GetForwardEntriesResult {
    /**
     * A list of Forward Entries. Each element contains the following attributes:
     */
    readonly entries: outputs.vpc.GetForwardEntriesEntry[];
    /**
     * The public IP address.
     */
    readonly externalIp?: string;
    readonly forwardTableId: string;
    /**
     * A list of Forward Entries IDs.
     */
    readonly ids: string[];
    /**
     * The private IP address.
     */
    readonly internalIp?: string;
    readonly nameRegex?: string;
    /**
     * A list of Forward Entries names.
     */
    readonly names?: string[];
    readonly outputFile?: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}

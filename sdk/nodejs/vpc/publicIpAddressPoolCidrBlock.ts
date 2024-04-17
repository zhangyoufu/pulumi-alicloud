// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a VPC Public Ip Address Pool Cidr Block resource.
 * > **NOTE:** Only users who have the required permissions can use the IP address pool feature of Elastic IP Address (EIP). To apply for the required permissions, [submit a ticket](https://smartservice.console.aliyun.com/service/create-ticket).
 *
 * For information about VPC Public Ip Address Pool Cidr Block and how to use it, see [What is Public Ip Address Pool Cidr Block](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/429100).
 *
 * > **NOTE:** Available since v1.189.0.
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
 * const config = new pulumi.Config();
 * const name = config.get("name") || "tf-example";
 * const default = alicloud.resourcemanager.getResourceGroups({
 *     status: "OK",
 * });
 * const defaultPublicIpAddressPool = new alicloud.vpc.PublicIpAddressPool("default", {
 *     description: name,
 *     publicIpAddressPoolName: name,
 *     isp: "BGP",
 *     resourceGroupId: _default.then(_default => _default.ids?.[0]),
 * });
 * const defaultPublicIpAddressPoolCidrBlock = new alicloud.vpc.PublicIpAddressPoolCidrBlock("default", {
 *     publicIpAddressPoolId: defaultPublicIpAddressPool.id,
 *     cidrBlock: "47.118.126.0/25",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * VPC Public Ip Address Pool Cidr Block can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:vpc/publicIpAddressPoolCidrBlock:PublicIpAddressPoolCidrBlock example <public_ip_address_pool_id>:<cidr_block>
 * ```
 */
export class PublicIpAddressPoolCidrBlock extends pulumi.CustomResource {
    /**
     * Get an existing PublicIpAddressPoolCidrBlock resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PublicIpAddressPoolCidrBlockState, opts?: pulumi.CustomResourceOptions): PublicIpAddressPoolCidrBlock {
        return new PublicIpAddressPoolCidrBlock(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:vpc/publicIpAddressPoolCidrBlock:PublicIpAddressPoolCidrBlock';

    /**
     * Returns true if the given object is an instance of PublicIpAddressPoolCidrBlock.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PublicIpAddressPoolCidrBlock {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PublicIpAddressPoolCidrBlock.__pulumiType;
    }

    /**
     * The CIDR block.
     */
    public readonly cidrBlock!: pulumi.Output<string>;
    /**
     * IP address and network segment mask. After you enter the mask, the system automatically allocates the IP address network segment. Value range: **24** to **28**.
     * > **NOTE:**  **CidrBlock** and **CidrMask** cannot be configured at the same time. Select one of them to configure.
     */
    public readonly cidrMask!: pulumi.Output<number | undefined>;
    /**
     * The creation time of the resource.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The ID of the VPC Public IP address pool.
     */
    public readonly publicIpAddressPoolId!: pulumi.Output<string>;
    /**
     * The status of the VPC Public Ip Address Pool Cidr Block.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a PublicIpAddressPoolCidrBlock resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PublicIpAddressPoolCidrBlockArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PublicIpAddressPoolCidrBlockArgs | PublicIpAddressPoolCidrBlockState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PublicIpAddressPoolCidrBlockState | undefined;
            resourceInputs["cidrBlock"] = state ? state.cidrBlock : undefined;
            resourceInputs["cidrMask"] = state ? state.cidrMask : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["publicIpAddressPoolId"] = state ? state.publicIpAddressPoolId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as PublicIpAddressPoolCidrBlockArgs | undefined;
            if ((!args || args.publicIpAddressPoolId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'publicIpAddressPoolId'");
            }
            resourceInputs["cidrBlock"] = args ? args.cidrBlock : undefined;
            resourceInputs["cidrMask"] = args ? args.cidrMask : undefined;
            resourceInputs["publicIpAddressPoolId"] = args ? args.publicIpAddressPoolId : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PublicIpAddressPoolCidrBlock.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PublicIpAddressPoolCidrBlock resources.
 */
export interface PublicIpAddressPoolCidrBlockState {
    /**
     * The CIDR block.
     */
    cidrBlock?: pulumi.Input<string>;
    /**
     * IP address and network segment mask. After you enter the mask, the system automatically allocates the IP address network segment. Value range: **24** to **28**.
     * > **NOTE:**  **CidrBlock** and **CidrMask** cannot be configured at the same time. Select one of them to configure.
     */
    cidrMask?: pulumi.Input<number>;
    /**
     * The creation time of the resource.
     */
    createTime?: pulumi.Input<string>;
    /**
     * The ID of the VPC Public IP address pool.
     */
    publicIpAddressPoolId?: pulumi.Input<string>;
    /**
     * The status of the VPC Public Ip Address Pool Cidr Block.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PublicIpAddressPoolCidrBlock resource.
 */
export interface PublicIpAddressPoolCidrBlockArgs {
    /**
     * The CIDR block.
     */
    cidrBlock?: pulumi.Input<string>;
    /**
     * IP address and network segment mask. After you enter the mask, the system automatically allocates the IP address network segment. Value range: **24** to **28**.
     * > **NOTE:**  **CidrBlock** and **CidrMask** cannot be configured at the same time. Select one of them to configure.
     */
    cidrMask?: pulumi.Input<number>;
    /**
     * The ID of the VPC Public IP address pool.
     */
    publicIpAddressPoolId: pulumi.Input<string>;
}

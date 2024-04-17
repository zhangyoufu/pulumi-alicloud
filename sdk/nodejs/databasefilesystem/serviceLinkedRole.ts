// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Using this data source can create Dbfs service-linked roles(SLR). Dbfs may need to access another Alibaba Cloud service to implement a specific feature. In this case, Dbfs must assume a specific service-linked role, which is a Resource Access Management (RAM) role, to obtain permissions to access another Alibaba Cloud service.
 *
 * For information about Dbfs service-linked roles(SLR) and how to use it, see [What is service-linked roles](https://www.alibabacloud.com/help/doc-detail/181425.htm).
 *
 * > **NOTE:** Available since v1.157.0.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const serviceLinkedRole = new alicloud.databasefilesystem.ServiceLinkedRole("service_linked_role", {productName: "AliyunServiceRoleForDbfs"});
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Dbfs service-linked roles(SLR) can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:databasefilesystem/serviceLinkedRole:ServiceLinkedRole example <product_name>
 * ```
 */
export class ServiceLinkedRole extends pulumi.CustomResource {
    /**
     * Get an existing ServiceLinkedRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceLinkedRoleState, opts?: pulumi.CustomResourceOptions): ServiceLinkedRole {
        return new ServiceLinkedRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:databasefilesystem/serviceLinkedRole:ServiceLinkedRole';

    /**
     * Returns true if the given object is an instance of ServiceLinkedRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceLinkedRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceLinkedRole.__pulumiType;
    }

    /**
     * The product name for SLR. Dbfs can automatically create the following service-linked roles: `AliyunServiceRoleForDbfs`.
     */
    public readonly productName!: pulumi.Output<string>;
    /**
     * The status of the service Associated role. Valid Values: `true`: Created. `false`: not created.
     */
    public /*out*/ readonly status!: pulumi.Output<boolean>;

    /**
     * Create a ServiceLinkedRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceLinkedRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceLinkedRoleArgs | ServiceLinkedRoleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceLinkedRoleState | undefined;
            resourceInputs["productName"] = state ? state.productName : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as ServiceLinkedRoleArgs | undefined;
            if ((!args || args.productName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'productName'");
            }
            resourceInputs["productName"] = args ? args.productName : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ServiceLinkedRole.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceLinkedRole resources.
 */
export interface ServiceLinkedRoleState {
    /**
     * The product name for SLR. Dbfs can automatically create the following service-linked roles: `AliyunServiceRoleForDbfs`.
     */
    productName?: pulumi.Input<string>;
    /**
     * The status of the service Associated role. Valid Values: `true`: Created. `false`: not created.
     */
    status?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a ServiceLinkedRole resource.
 */
export interface ServiceLinkedRoleArgs {
    /**
     * The product name for SLR. Dbfs can automatically create the following service-linked roles: `AliyunServiceRoleForDbfs`.
     */
    productName: pulumi.Input<string>;
}

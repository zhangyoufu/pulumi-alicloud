// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Cloud Monitor Service Enterprise Public resource. Hybrid Cloud Monitoring.
 *
 * For information about Cloud Monitor Service Enterprise Public and how to use it, see [What is Enterprise Public](https://www.alibabacloud.com/help/en/cms/user-guide/overview-3).
 *
 * > **NOTE:** Available since v1.215.0.
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
 * const name = config.get("name") || "terraform-example";
 * const _default = new alicloud.cloudmonitor.ServiceEnterprisePublic("default", {});
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Cloud Monitor Service Enterprise Public can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:cloudmonitor/serviceEnterprisePublic:ServiceEnterprisePublic example <id>
 * ```
 */
export class ServiceEnterprisePublic extends pulumi.CustomResource {
    /**
     * Get an existing ServiceEnterprisePublic resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceEnterprisePublicState, opts?: pulumi.CustomResourceOptions): ServiceEnterprisePublic {
        return new ServiceEnterprisePublic(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cloudmonitor/serviceEnterprisePublic:ServiceEnterprisePublic';

    /**
     * Returns true if the given object is an instance of ServiceEnterprisePublic.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceEnterprisePublic {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceEnterprisePublic.__pulumiType;
    }

    /**
     * The creation time of the resource.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;

    /**
     * Create a ServiceEnterprisePublic resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ServiceEnterprisePublicArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceEnterprisePublicArgs | ServiceEnterprisePublicState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceEnterprisePublicState | undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
        } else {
            const args = argsOrState as ServiceEnterprisePublicArgs | undefined;
            resourceInputs["createTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ServiceEnterprisePublic.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceEnterprisePublic resources.
 */
export interface ServiceEnterprisePublicState {
    /**
     * The creation time of the resource.
     */
    createTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceEnterprisePublic resource.
 */
export interface ServiceEnterprisePublicArgs {
}

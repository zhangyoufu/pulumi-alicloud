// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class RouteEntry extends pulumi.CustomResource {
    /**
     * Get an existing RouteEntry resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouteEntryState, opts?: pulumi.CustomResourceOptions): RouteEntry {
        return new RouteEntry(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:vpc/routeEntry:RouteEntry';

    /**
     * Returns true if the given object is an instance of RouteEntry.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RouteEntry {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RouteEntry.__pulumiType;
    }

    /**
     * The RouteEntry's target network segment.
     */
    public readonly destinationCidrblock!: pulumi.Output<string | undefined>;
    /**
     * The name of the route entry. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The route entry's next hop. ECS instance ID or VPC router interface ID.
     */
    public readonly nexthopId!: pulumi.Output<string | undefined>;
    /**
     * The next hop type. Available values:
     */
    public readonly nexthopType!: pulumi.Output<string | undefined>;
    /**
     * The ID of the route table.
     */
    public readonly routeTableId!: pulumi.Output<string>;
    /**
     * This argument has beeb deprecated. Please use other arguments to launch a custom route entry.
     *
     * @deprecated Attribute router_id has been deprecated and suggest removing it from your template.
     */
    public readonly routerId!: pulumi.Output<string>;

    /**
     * Create a RouteEntry resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouteEntryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouteEntryArgs | RouteEntryState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RouteEntryState | undefined;
            inputs["destinationCidrblock"] = state ? state.destinationCidrblock : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["nexthopId"] = state ? state.nexthopId : undefined;
            inputs["nexthopType"] = state ? state.nexthopType : undefined;
            inputs["routeTableId"] = state ? state.routeTableId : undefined;
            inputs["routerId"] = state ? state.routerId : undefined;
        } else {
            const args = argsOrState as RouteEntryArgs | undefined;
            if (!args || args.routeTableId === undefined) {
                throw new Error("Missing required property 'routeTableId'");
            }
            inputs["destinationCidrblock"] = args ? args.destinationCidrblock : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["nexthopId"] = args ? args.nexthopId : undefined;
            inputs["nexthopType"] = args ? args.nexthopType : undefined;
            inputs["routeTableId"] = args ? args.routeTableId : undefined;
            inputs["routerId"] = args ? args.routerId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(RouteEntry.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RouteEntry resources.
 */
export interface RouteEntryState {
    /**
     * The RouteEntry's target network segment.
     */
    readonly destinationCidrblock?: pulumi.Input<string>;
    /**
     * The name of the route entry. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The route entry's next hop. ECS instance ID or VPC router interface ID.
     */
    readonly nexthopId?: pulumi.Input<string>;
    /**
     * The next hop type. Available values:
     */
    readonly nexthopType?: pulumi.Input<string>;
    /**
     * The ID of the route table.
     */
    readonly routeTableId?: pulumi.Input<string>;
    /**
     * This argument has beeb deprecated. Please use other arguments to launch a custom route entry.
     *
     * @deprecated Attribute router_id has been deprecated and suggest removing it from your template.
     */
    readonly routerId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RouteEntry resource.
 */
export interface RouteEntryArgs {
    /**
     * The RouteEntry's target network segment.
     */
    readonly destinationCidrblock?: pulumi.Input<string>;
    /**
     * The name of the route entry. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The route entry's next hop. ECS instance ID or VPC router interface ID.
     */
    readonly nexthopId?: pulumi.Input<string>;
    /**
     * The next hop type. Available values:
     */
    readonly nexthopType?: pulumi.Input<string>;
    /**
     * The ID of the route table.
     */
    readonly routeTableId: pulumi.Input<string>;
    /**
     * This argument has beeb deprecated. Please use other arguments to launch a custom route entry.
     *
     * @deprecated Attribute router_id has been deprecated and suggest removing it from your template.
     */
    readonly routerId?: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Express Connect Physical Connection resource.
 *
 * For information about Express Connect Physical Connection and how to use it, see [What is Physical Connection](https://www.alibabacloud.com/help/doc-detail/44852.htm).
 *
 * > **NOTE:** Available since v1.132.0.
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
 * const hz = new alicloud.Provider("hz", {region: "cn-hangzhou"});
 * const sgp = new alicloud.Provider("sgp", {region: "ap-southeast-1"});
 * const domestic = new alicloud.expressconnect.PhysicalConnection("domestic", {
 *     accessPointId: "ap-cn-hangzhou-yh-B",
 *     lineOperator: "CT",
 *     peerLocation: "example_value",
 *     physicalConnectionName: "example_value",
 *     type: "VPC",
 *     description: "my domestic connection",
 *     portType: "1000Base-LX",
 *     bandwidth: "100",
 * }, {
 *     provider: alicloud.hz,
 * });
 * const international = new alicloud.expressconnect.PhysicalConnection("international", {
 *     accessPointId: "ap-sg-singpore-A",
 *     lineOperator: "Other",
 *     peerLocation: "example_value",
 *     physicalConnectionName: "example_value",
 *     type: "VPC",
 *     description: "my domestic connection",
 *     portType: "1000Base-LX",
 *     bandwidth: "100",
 * }, {
 *     provider: alicloud.sgp,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Express Connect Physical Connection can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:expressconnect/physicalConnection:PhysicalConnection example <id>
 * ```
 */
export class PhysicalConnection extends pulumi.CustomResource {
    /**
     * Get an existing PhysicalConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PhysicalConnectionState, opts?: pulumi.CustomResourceOptions): PhysicalConnection {
        return new PhysicalConnection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:expressconnect/physicalConnection:PhysicalConnection';

    /**
     * Returns true if the given object is an instance of PhysicalConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PhysicalConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PhysicalConnection.__pulumiType;
    }

    /**
     * The Physical Leased Line Access Point ID.
     */
    public readonly accessPointId!: pulumi.Output<string>;
    /**
     * On the Bandwidth of the ECC Service and Physical Connection.
     */
    public readonly bandwidth!: pulumi.Output<string>;
    /**
     * Operators for Physical Connection Circuit Provided Coding.
     */
    public readonly circuitCode!: pulumi.Output<string | undefined>;
    /**
     * The Physical Connection to Which the Description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Provides Access to the Physical Line Operator. Valid values:
     * * CT: China Telecom
     * * CU: China Unicom
     * * CM: china Mobile
     * * CO: Other Chinese
     * * Equinix: Equinix
     * * Other: Other Overseas.
     */
    public readonly lineOperator!: pulumi.Output<string>;
    /**
     * and an on-Premises Data Center Location.
     */
    public readonly peerLocation!: pulumi.Output<string | undefined>;
    /**
     * on Behalf of the Resource Name of the Resources-Attribute Field.
     */
    public readonly physicalConnectionName!: pulumi.Output<string | undefined>;
    /**
     * The Physical Leased Line Access Port Type. Valid value:
     * * 100Base-T: Fast Electrical Ports
     * * 1000Base-T: gigabit Electrical Ports
     * * 1000Base-LX: Gigabit Singlemode Optical Ports (10Km)
     * * 10GBase-T: Gigabit Electrical Port
     * * 10GBase-LR: Gigabit Singlemode Optical Ports (10Km).
     * * 40GBase-LR: 40 Gigabit Singlemode Optical Ports.
     * * 100GBase-LR: One hundred thousand Gigabit Singlemode Optical Ports.
     *
     * **NOTE:** From in v1.185.0+, The `40GBase-LR` and `100GBase-LR` is valid. and Set these values based on the water levels of background ports. For details about the water levels, contact the business manager.
     */
    public readonly portType!: pulumi.Output<string | undefined>;
    /**
     * Redundant Physical Connection to Which the ID.
     */
    public readonly redundantPhysicalConnectionId!: pulumi.Output<string | undefined>;
    /**
     * Resources on Behalf of a State of the Resource Attribute Field. Valid values: `Canceled`, `Enabled`, `Terminated`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Physical Private Line of Type. Default Value: VPC.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a PhysicalConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PhysicalConnectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PhysicalConnectionArgs | PhysicalConnectionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PhysicalConnectionState | undefined;
            resourceInputs["accessPointId"] = state ? state.accessPointId : undefined;
            resourceInputs["bandwidth"] = state ? state.bandwidth : undefined;
            resourceInputs["circuitCode"] = state ? state.circuitCode : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["lineOperator"] = state ? state.lineOperator : undefined;
            resourceInputs["peerLocation"] = state ? state.peerLocation : undefined;
            resourceInputs["physicalConnectionName"] = state ? state.physicalConnectionName : undefined;
            resourceInputs["portType"] = state ? state.portType : undefined;
            resourceInputs["redundantPhysicalConnectionId"] = state ? state.redundantPhysicalConnectionId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as PhysicalConnectionArgs | undefined;
            if ((!args || args.accessPointId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accessPointId'");
            }
            if ((!args || args.lineOperator === undefined) && !opts.urn) {
                throw new Error("Missing required property 'lineOperator'");
            }
            resourceInputs["accessPointId"] = args ? args.accessPointId : undefined;
            resourceInputs["bandwidth"] = args ? args.bandwidth : undefined;
            resourceInputs["circuitCode"] = args ? args.circuitCode : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["lineOperator"] = args ? args.lineOperator : undefined;
            resourceInputs["peerLocation"] = args ? args.peerLocation : undefined;
            resourceInputs["physicalConnectionName"] = args ? args.physicalConnectionName : undefined;
            resourceInputs["portType"] = args ? args.portType : undefined;
            resourceInputs["redundantPhysicalConnectionId"] = args ? args.redundantPhysicalConnectionId : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PhysicalConnection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PhysicalConnection resources.
 */
export interface PhysicalConnectionState {
    /**
     * The Physical Leased Line Access Point ID.
     */
    accessPointId?: pulumi.Input<string>;
    /**
     * On the Bandwidth of the ECC Service and Physical Connection.
     */
    bandwidth?: pulumi.Input<string>;
    /**
     * Operators for Physical Connection Circuit Provided Coding.
     */
    circuitCode?: pulumi.Input<string>;
    /**
     * The Physical Connection to Which the Description.
     */
    description?: pulumi.Input<string>;
    /**
     * Provides Access to the Physical Line Operator. Valid values:
     * * CT: China Telecom
     * * CU: China Unicom
     * * CM: china Mobile
     * * CO: Other Chinese
     * * Equinix: Equinix
     * * Other: Other Overseas.
     */
    lineOperator?: pulumi.Input<string>;
    /**
     * and an on-Premises Data Center Location.
     */
    peerLocation?: pulumi.Input<string>;
    /**
     * on Behalf of the Resource Name of the Resources-Attribute Field.
     */
    physicalConnectionName?: pulumi.Input<string>;
    /**
     * The Physical Leased Line Access Port Type. Valid value:
     * * 100Base-T: Fast Electrical Ports
     * * 1000Base-T: gigabit Electrical Ports
     * * 1000Base-LX: Gigabit Singlemode Optical Ports (10Km)
     * * 10GBase-T: Gigabit Electrical Port
     * * 10GBase-LR: Gigabit Singlemode Optical Ports (10Km).
     * * 40GBase-LR: 40 Gigabit Singlemode Optical Ports.
     * * 100GBase-LR: One hundred thousand Gigabit Singlemode Optical Ports.
     *
     * **NOTE:** From in v1.185.0+, The `40GBase-LR` and `100GBase-LR` is valid. and Set these values based on the water levels of background ports. For details about the water levels, contact the business manager.
     */
    portType?: pulumi.Input<string>;
    /**
     * Redundant Physical Connection to Which the ID.
     */
    redundantPhysicalConnectionId?: pulumi.Input<string>;
    /**
     * Resources on Behalf of a State of the Resource Attribute Field. Valid values: `Canceled`, `Enabled`, `Terminated`.
     */
    status?: pulumi.Input<string>;
    /**
     * Physical Private Line of Type. Default Value: VPC.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PhysicalConnection resource.
 */
export interface PhysicalConnectionArgs {
    /**
     * The Physical Leased Line Access Point ID.
     */
    accessPointId: pulumi.Input<string>;
    /**
     * On the Bandwidth of the ECC Service and Physical Connection.
     */
    bandwidth?: pulumi.Input<string>;
    /**
     * Operators for Physical Connection Circuit Provided Coding.
     */
    circuitCode?: pulumi.Input<string>;
    /**
     * The Physical Connection to Which the Description.
     */
    description?: pulumi.Input<string>;
    /**
     * Provides Access to the Physical Line Operator. Valid values:
     * * CT: China Telecom
     * * CU: China Unicom
     * * CM: china Mobile
     * * CO: Other Chinese
     * * Equinix: Equinix
     * * Other: Other Overseas.
     */
    lineOperator: pulumi.Input<string>;
    /**
     * and an on-Premises Data Center Location.
     */
    peerLocation?: pulumi.Input<string>;
    /**
     * on Behalf of the Resource Name of the Resources-Attribute Field.
     */
    physicalConnectionName?: pulumi.Input<string>;
    /**
     * The Physical Leased Line Access Port Type. Valid value:
     * * 100Base-T: Fast Electrical Ports
     * * 1000Base-T: gigabit Electrical Ports
     * * 1000Base-LX: Gigabit Singlemode Optical Ports (10Km)
     * * 10GBase-T: Gigabit Electrical Port
     * * 10GBase-LR: Gigabit Singlemode Optical Ports (10Km).
     * * 40GBase-LR: 40 Gigabit Singlemode Optical Ports.
     * * 100GBase-LR: One hundred thousand Gigabit Singlemode Optical Ports.
     *
     * **NOTE:** From in v1.185.0+, The `40GBase-LR` and `100GBase-LR` is valid. and Set these values based on the water levels of background ports. For details about the water levels, contact the business manager.
     */
    portType?: pulumi.Input<string>;
    /**
     * Redundant Physical Connection to Which the ID.
     */
    redundantPhysicalConnectionId?: pulumi.Input<string>;
    /**
     * Resources on Behalf of a State of the Resource Attribute Field. Valid values: `Canceled`, `Enabled`, `Terminated`.
     */
    status?: pulumi.Input<string>;
    /**
     * Physical Private Line of Type. Default Value: VPC.
     */
    type?: pulumi.Input<string>;
}

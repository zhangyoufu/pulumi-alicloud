// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.expressconnect;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PhysicalConnectionArgs extends com.pulumi.resources.ResourceArgs {

    public static final PhysicalConnectionArgs Empty = new PhysicalConnectionArgs();

    /**
     * The Physical Leased Line Access Point ID.
     * 
     */
    @Import(name="accessPointId", required=true)
    private Output<String> accessPointId;

    /**
     * @return The Physical Leased Line Access Point ID.
     * 
     */
    public Output<String> accessPointId() {
        return this.accessPointId;
    }

    /**
     * On the Bandwidth of the ECC Service and Physical Connection.
     * 
     */
    @Import(name="bandwidth")
    private @Nullable Output<String> bandwidth;

    /**
     * @return On the Bandwidth of the ECC Service and Physical Connection.
     * 
     */
    public Optional<Output<String>> bandwidth() {
        return Optional.ofNullable(this.bandwidth);
    }

    /**
     * Operators for Physical Connection Circuit Provided Coding.
     * 
     */
    @Import(name="circuitCode")
    private @Nullable Output<String> circuitCode;

    /**
     * @return Operators for Physical Connection Circuit Provided Coding.
     * 
     */
    public Optional<Output<String>> circuitCode() {
        return Optional.ofNullable(this.circuitCode);
    }

    /**
     * The Physical Connection to Which the Description.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The Physical Connection to Which the Description.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Provides Access to the Physical Line Operator. Valid values:
     * * CT: China Telecom
     * * CU: China Unicom
     * * CM: china Mobile
     * * CO: Other Chinese
     * * Equinix: Equinix
     * * Other: Other Overseas.
     * 
     */
    @Import(name="lineOperator", required=true)
    private Output<String> lineOperator;

    /**
     * @return Provides Access to the Physical Line Operator. Valid values:
     * * CT: China Telecom
     * * CU: China Unicom
     * * CM: china Mobile
     * * CO: Other Chinese
     * * Equinix: Equinix
     * * Other: Other Overseas.
     * 
     */
    public Output<String> lineOperator() {
        return this.lineOperator;
    }

    /**
     * and an on-Premises Data Center Location.
     * 
     */
    @Import(name="peerLocation", required=true)
    private Output<String> peerLocation;

    /**
     * @return and an on-Premises Data Center Location.
     * 
     */
    public Output<String> peerLocation() {
        return this.peerLocation;
    }

    /**
     * on Behalf of the Resource Name of the Resources-Attribute Field.
     * 
     */
    @Import(name="physicalConnectionName")
    private @Nullable Output<String> physicalConnectionName;

    /**
     * @return on Behalf of the Resource Name of the Resources-Attribute Field.
     * 
     */
    public Optional<Output<String>> physicalConnectionName() {
        return Optional.ofNullable(this.physicalConnectionName);
    }

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
     */
    @Import(name="portType")
    private @Nullable Output<String> portType;

    /**
     * @return The Physical Leased Line Access Port Type. Valid value:
     * * 100Base-T: Fast Electrical Ports
     * * 1000Base-T: gigabit Electrical Ports
     * * 1000Base-LX: Gigabit Singlemode Optical Ports (10Km)
     * * 10GBase-T: Gigabit Electrical Port
     * * 10GBase-LR: Gigabit Singlemode Optical Ports (10Km).
     * * 40GBase-LR: 40 Gigabit Singlemode Optical Ports.
     * * 100GBase-LR: One hundred thousand Gigabit Singlemode Optical Ports.
     * 
     */
    public Optional<Output<String>> portType() {
        return Optional.ofNullable(this.portType);
    }

    /**
     * Redundant Physical Connection to Which the ID.
     * 
     */
    @Import(name="redundantPhysicalConnectionId")
    private @Nullable Output<String> redundantPhysicalConnectionId;

    /**
     * @return Redundant Physical Connection to Which the ID.
     * 
     */
    public Optional<Output<String>> redundantPhysicalConnectionId() {
        return Optional.ofNullable(this.redundantPhysicalConnectionId);
    }

    /**
     * Resources on Behalf of a State of the Resource Attribute Field. Valid values: `Canceled`, `Enabled`, `Terminated`.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return Resources on Behalf of a State of the Resource Attribute Field. Valid values: `Canceled`, `Enabled`, `Terminated`.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * Physical Private Line of Type. Default Value: VPC.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return Physical Private Line of Type. Default Value: VPC.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private PhysicalConnectionArgs() {}

    private PhysicalConnectionArgs(PhysicalConnectionArgs $) {
        this.accessPointId = $.accessPointId;
        this.bandwidth = $.bandwidth;
        this.circuitCode = $.circuitCode;
        this.description = $.description;
        this.lineOperator = $.lineOperator;
        this.peerLocation = $.peerLocation;
        this.physicalConnectionName = $.physicalConnectionName;
        this.portType = $.portType;
        this.redundantPhysicalConnectionId = $.redundantPhysicalConnectionId;
        this.status = $.status;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PhysicalConnectionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PhysicalConnectionArgs $;

        public Builder() {
            $ = new PhysicalConnectionArgs();
        }

        public Builder(PhysicalConnectionArgs defaults) {
            $ = new PhysicalConnectionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessPointId The Physical Leased Line Access Point ID.
         * 
         * @return builder
         * 
         */
        public Builder accessPointId(Output<String> accessPointId) {
            $.accessPointId = accessPointId;
            return this;
        }

        /**
         * @param accessPointId The Physical Leased Line Access Point ID.
         * 
         * @return builder
         * 
         */
        public Builder accessPointId(String accessPointId) {
            return accessPointId(Output.of(accessPointId));
        }

        /**
         * @param bandwidth On the Bandwidth of the ECC Service and Physical Connection.
         * 
         * @return builder
         * 
         */
        public Builder bandwidth(@Nullable Output<String> bandwidth) {
            $.bandwidth = bandwidth;
            return this;
        }

        /**
         * @param bandwidth On the Bandwidth of the ECC Service and Physical Connection.
         * 
         * @return builder
         * 
         */
        public Builder bandwidth(String bandwidth) {
            return bandwidth(Output.of(bandwidth));
        }

        /**
         * @param circuitCode Operators for Physical Connection Circuit Provided Coding.
         * 
         * @return builder
         * 
         */
        public Builder circuitCode(@Nullable Output<String> circuitCode) {
            $.circuitCode = circuitCode;
            return this;
        }

        /**
         * @param circuitCode Operators for Physical Connection Circuit Provided Coding.
         * 
         * @return builder
         * 
         */
        public Builder circuitCode(String circuitCode) {
            return circuitCode(Output.of(circuitCode));
        }

        /**
         * @param description The Physical Connection to Which the Description.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The Physical Connection to Which the Description.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param lineOperator Provides Access to the Physical Line Operator. Valid values:
         * * CT: China Telecom
         * * CU: China Unicom
         * * CM: china Mobile
         * * CO: Other Chinese
         * * Equinix: Equinix
         * * Other: Other Overseas.
         * 
         * @return builder
         * 
         */
        public Builder lineOperator(Output<String> lineOperator) {
            $.lineOperator = lineOperator;
            return this;
        }

        /**
         * @param lineOperator Provides Access to the Physical Line Operator. Valid values:
         * * CT: China Telecom
         * * CU: China Unicom
         * * CM: china Mobile
         * * CO: Other Chinese
         * * Equinix: Equinix
         * * Other: Other Overseas.
         * 
         * @return builder
         * 
         */
        public Builder lineOperator(String lineOperator) {
            return lineOperator(Output.of(lineOperator));
        }

        /**
         * @param peerLocation and an on-Premises Data Center Location.
         * 
         * @return builder
         * 
         */
        public Builder peerLocation(Output<String> peerLocation) {
            $.peerLocation = peerLocation;
            return this;
        }

        /**
         * @param peerLocation and an on-Premises Data Center Location.
         * 
         * @return builder
         * 
         */
        public Builder peerLocation(String peerLocation) {
            return peerLocation(Output.of(peerLocation));
        }

        /**
         * @param physicalConnectionName on Behalf of the Resource Name of the Resources-Attribute Field.
         * 
         * @return builder
         * 
         */
        public Builder physicalConnectionName(@Nullable Output<String> physicalConnectionName) {
            $.physicalConnectionName = physicalConnectionName;
            return this;
        }

        /**
         * @param physicalConnectionName on Behalf of the Resource Name of the Resources-Attribute Field.
         * 
         * @return builder
         * 
         */
        public Builder physicalConnectionName(String physicalConnectionName) {
            return physicalConnectionName(Output.of(physicalConnectionName));
        }

        /**
         * @param portType The Physical Leased Line Access Port Type. Valid value:
         * * 100Base-T: Fast Electrical Ports
         * * 1000Base-T: gigabit Electrical Ports
         * * 1000Base-LX: Gigabit Singlemode Optical Ports (10Km)
         * * 10GBase-T: Gigabit Electrical Port
         * * 10GBase-LR: Gigabit Singlemode Optical Ports (10Km).
         * * 40GBase-LR: 40 Gigabit Singlemode Optical Ports.
         * * 100GBase-LR: One hundred thousand Gigabit Singlemode Optical Ports.
         * 
         * @return builder
         * 
         */
        public Builder portType(@Nullable Output<String> portType) {
            $.portType = portType;
            return this;
        }

        /**
         * @param portType The Physical Leased Line Access Port Type. Valid value:
         * * 100Base-T: Fast Electrical Ports
         * * 1000Base-T: gigabit Electrical Ports
         * * 1000Base-LX: Gigabit Singlemode Optical Ports (10Km)
         * * 10GBase-T: Gigabit Electrical Port
         * * 10GBase-LR: Gigabit Singlemode Optical Ports (10Km).
         * * 40GBase-LR: 40 Gigabit Singlemode Optical Ports.
         * * 100GBase-LR: One hundred thousand Gigabit Singlemode Optical Ports.
         * 
         * @return builder
         * 
         */
        public Builder portType(String portType) {
            return portType(Output.of(portType));
        }

        /**
         * @param redundantPhysicalConnectionId Redundant Physical Connection to Which the ID.
         * 
         * @return builder
         * 
         */
        public Builder redundantPhysicalConnectionId(@Nullable Output<String> redundantPhysicalConnectionId) {
            $.redundantPhysicalConnectionId = redundantPhysicalConnectionId;
            return this;
        }

        /**
         * @param redundantPhysicalConnectionId Redundant Physical Connection to Which the ID.
         * 
         * @return builder
         * 
         */
        public Builder redundantPhysicalConnectionId(String redundantPhysicalConnectionId) {
            return redundantPhysicalConnectionId(Output.of(redundantPhysicalConnectionId));
        }

        /**
         * @param status Resources on Behalf of a State of the Resource Attribute Field. Valid values: `Canceled`, `Enabled`, `Terminated`.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status Resources on Behalf of a State of the Resource Attribute Field. Valid values: `Canceled`, `Enabled`, `Terminated`.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param type Physical Private Line of Type. Default Value: VPC.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Physical Private Line of Type. Default Value: VPC.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public PhysicalConnectionArgs build() {
            $.accessPointId = Objects.requireNonNull($.accessPointId, "expected parameter 'accessPointId' to be non-null");
            $.lineOperator = Objects.requireNonNull($.lineOperator, "expected parameter 'lineOperator' to be non-null");
            $.peerLocation = Objects.requireNonNull($.peerLocation, "expected parameter 'peerLocation' to be non-null");
            return $;
        }
    }

}

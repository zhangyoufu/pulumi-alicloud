// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpn.outputs;

import com.pulumi.alicloud.vpn.outputs.GetGatewaysGateway;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetGatewaysResult {
    /**
     * @return The business status of the VPN gateway.
     * 
     */
    private @Nullable String businessStatus;
    /**
     * @return Whether the ipsec function is enabled.
     * 
     * @deprecated
     * Field &#39;enable_ipsec&#39; has been deprecated from provider version 1.193.0 and it will be removed in the future version.
     * 
     */
    @Deprecated /* Field 'enable_ipsec' has been deprecated from provider version 1.193.0 and it will be removed in the future version. */
    private @Nullable Boolean enableIpsec;
    /**
     * @return A list of VPN gateways. Each element contains the following attributes:
     * 
     */
    private List<GetGatewaysGateway> gateways;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return IDs of the VPN.
     * 
     */
    private List<String> ids;
    private @Nullable Boolean includeReservationData;
    private @Nullable String nameRegex;
    /**
     * @return names of the VPN.
     * 
     */
    private List<String> names;
    private @Nullable String outputFile;
    /**
     * @return The status of the VPN
     * 
     */
    private @Nullable String status;
    /**
     * @return ID of the VPC that the VPN belongs.
     * 
     */
    private @Nullable String vpcId;

    private GetGatewaysResult() {}
    /**
     * @return The business status of the VPN gateway.
     * 
     */
    public Optional<String> businessStatus() {
        return Optional.ofNullable(this.businessStatus);
    }
    /**
     * @return Whether the ipsec function is enabled.
     * 
     * @deprecated
     * Field &#39;enable_ipsec&#39; has been deprecated from provider version 1.193.0 and it will be removed in the future version.
     * 
     */
    @Deprecated /* Field 'enable_ipsec' has been deprecated from provider version 1.193.0 and it will be removed in the future version. */
    public Optional<Boolean> enableIpsec() {
        return Optional.ofNullable(this.enableIpsec);
    }
    /**
     * @return A list of VPN gateways. Each element contains the following attributes:
     * 
     */
    public List<GetGatewaysGateway> gateways() {
        return this.gateways;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return IDs of the VPN.
     * 
     */
    public List<String> ids() {
        return this.ids;
    }
    public Optional<Boolean> includeReservationData() {
        return Optional.ofNullable(this.includeReservationData);
    }
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }
    /**
     * @return names of the VPN.
     * 
     */
    public List<String> names() {
        return this.names;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    /**
     * @return The status of the VPN
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }
    /**
     * @return ID of the VPC that the VPN belongs.
     * 
     */
    public Optional<String> vpcId() {
        return Optional.ofNullable(this.vpcId);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetGatewaysResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String businessStatus;
        private @Nullable Boolean enableIpsec;
        private List<GetGatewaysGateway> gateways;
        private String id;
        private List<String> ids;
        private @Nullable Boolean includeReservationData;
        private @Nullable String nameRegex;
        private List<String> names;
        private @Nullable String outputFile;
        private @Nullable String status;
        private @Nullable String vpcId;
        public Builder() {}
        public Builder(GetGatewaysResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.businessStatus = defaults.businessStatus;
    	      this.enableIpsec = defaults.enableIpsec;
    	      this.gateways = defaults.gateways;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.includeReservationData = defaults.includeReservationData;
    	      this.nameRegex = defaults.nameRegex;
    	      this.names = defaults.names;
    	      this.outputFile = defaults.outputFile;
    	      this.status = defaults.status;
    	      this.vpcId = defaults.vpcId;
        }

        @CustomType.Setter
        public Builder businessStatus(@Nullable String businessStatus) {
            this.businessStatus = businessStatus;
            return this;
        }
        @CustomType.Setter
        public Builder enableIpsec(@Nullable Boolean enableIpsec) {
            this.enableIpsec = enableIpsec;
            return this;
        }
        @CustomType.Setter
        public Builder gateways(List<GetGatewaysGateway> gateways) {
            this.gateways = Objects.requireNonNull(gateways);
            return this;
        }
        public Builder gateways(GetGatewaysGateway... gateways) {
            return gateways(List.of(gateways));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder ids(List<String> ids) {
            this.ids = Objects.requireNonNull(ids);
            return this;
        }
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }
        @CustomType.Setter
        public Builder includeReservationData(@Nullable Boolean includeReservationData) {
            this.includeReservationData = includeReservationData;
            return this;
        }
        @CustomType.Setter
        public Builder nameRegex(@Nullable String nameRegex) {
            this.nameRegex = nameRegex;
            return this;
        }
        @CustomType.Setter
        public Builder names(List<String> names) {
            this.names = Objects.requireNonNull(names);
            return this;
        }
        public Builder names(String... names) {
            return names(List.of(names));
        }
        @CustomType.Setter
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
            return this;
        }
        @CustomType.Setter
        public Builder status(@Nullable String status) {
            this.status = status;
            return this;
        }
        @CustomType.Setter
        public Builder vpcId(@Nullable String vpcId) {
            this.vpcId = vpcId;
            return this;
        }
        public GetGatewaysResult build() {
            final var o = new GetGatewaysResult();
            o.businessStatus = businessStatus;
            o.enableIpsec = enableIpsec;
            o.gateways = gateways;
            o.id = id;
            o.ids = ids;
            o.includeReservationData = includeReservationData;
            o.nameRegex = nameRegex;
            o.names = names;
            o.outputFile = outputFile;
            o.status = status;
            o.vpcId = vpcId;
            return o;
        }
    }
}

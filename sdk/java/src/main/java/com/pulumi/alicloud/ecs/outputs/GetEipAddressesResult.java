// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs.outputs;

import com.pulumi.alicloud.ecs.outputs.GetEipAddressesAddress;
import com.pulumi.alicloud.ecs.outputs.GetEipAddressesEip;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetEipAddressesResult {
    private @Nullable String addressName;
    private List<GetEipAddressesAddress> addresses;
    private @Nullable String associatedInstanceId;
    private @Nullable String associatedInstanceType;
    private @Nullable Boolean dryRun;
    /**
     * @deprecated
     * Field &#39;eips&#39; has been deprecated from provider version 1.126.0 and it will be removed in the future version. Please use the new attribute &#39;addresses&#39; instead.
     * 
     */
    @Deprecated /* Field 'eips' has been deprecated from provider version 1.126.0 and it will be removed in the future version. Please use the new attribute 'addresses' instead. */
    private List<GetEipAddressesEip> eips;
    private @Nullable Boolean enableDetails;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private List<String> ids;
    private @Nullable Boolean includeReservationData;
    private @Nullable String ipAddress;
    /**
     * @deprecated
     * Field &#39;ip_addresses&#39; has been deprecated from provider version 1.126.0 and it will be removed in the future version. Please use the new attribute &#39;ip_address&#39; instead.
     * 
     */
    @Deprecated /* Field 'ip_addresses' has been deprecated from provider version 1.126.0 and it will be removed in the future version. Please use the new attribute 'ip_address' instead. */
    private @Nullable List<String> ipAddresses;
    private @Nullable String isp;
    private @Nullable String lockReason;
    private @Nullable String nameRegex;
    private List<String> names;
    private @Nullable String outputFile;
    private @Nullable String paymentType;
    private @Nullable String resourceGroupId;
    private @Nullable String segmentInstanceId;
    private @Nullable String status;
    private @Nullable Map<String,String> tags;

    private GetEipAddressesResult() {}
    public Optional<String> addressName() {
        return Optional.ofNullable(this.addressName);
    }
    public List<GetEipAddressesAddress> addresses() {
        return this.addresses;
    }
    public Optional<String> associatedInstanceId() {
        return Optional.ofNullable(this.associatedInstanceId);
    }
    public Optional<String> associatedInstanceType() {
        return Optional.ofNullable(this.associatedInstanceType);
    }
    public Optional<Boolean> dryRun() {
        return Optional.ofNullable(this.dryRun);
    }
    /**
     * @deprecated
     * Field &#39;eips&#39; has been deprecated from provider version 1.126.0 and it will be removed in the future version. Please use the new attribute &#39;addresses&#39; instead.
     * 
     */
    @Deprecated /* Field 'eips' has been deprecated from provider version 1.126.0 and it will be removed in the future version. Please use the new attribute 'addresses' instead. */
    public List<GetEipAddressesEip> eips() {
        return this.eips;
    }
    public Optional<Boolean> enableDetails() {
        return Optional.ofNullable(this.enableDetails);
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public List<String> ids() {
        return this.ids;
    }
    public Optional<Boolean> includeReservationData() {
        return Optional.ofNullable(this.includeReservationData);
    }
    public Optional<String> ipAddress() {
        return Optional.ofNullable(this.ipAddress);
    }
    /**
     * @deprecated
     * Field &#39;ip_addresses&#39; has been deprecated from provider version 1.126.0 and it will be removed in the future version. Please use the new attribute &#39;ip_address&#39; instead.
     * 
     */
    @Deprecated /* Field 'ip_addresses' has been deprecated from provider version 1.126.0 and it will be removed in the future version. Please use the new attribute 'ip_address' instead. */
    public List<String> ipAddresses() {
        return this.ipAddresses == null ? List.of() : this.ipAddresses;
    }
    public Optional<String> isp() {
        return Optional.ofNullable(this.isp);
    }
    public Optional<String> lockReason() {
        return Optional.ofNullable(this.lockReason);
    }
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }
    public List<String> names() {
        return this.names;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    public Optional<String> paymentType() {
        return Optional.ofNullable(this.paymentType);
    }
    public Optional<String> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }
    public Optional<String> segmentInstanceId() {
        return Optional.ofNullable(this.segmentInstanceId);
    }
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }
    public Map<String,String> tags() {
        return this.tags == null ? Map.of() : this.tags;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetEipAddressesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String addressName;
        private List<GetEipAddressesAddress> addresses;
        private @Nullable String associatedInstanceId;
        private @Nullable String associatedInstanceType;
        private @Nullable Boolean dryRun;
        private List<GetEipAddressesEip> eips;
        private @Nullable Boolean enableDetails;
        private String id;
        private List<String> ids;
        private @Nullable Boolean includeReservationData;
        private @Nullable String ipAddress;
        private @Nullable List<String> ipAddresses;
        private @Nullable String isp;
        private @Nullable String lockReason;
        private @Nullable String nameRegex;
        private List<String> names;
        private @Nullable String outputFile;
        private @Nullable String paymentType;
        private @Nullable String resourceGroupId;
        private @Nullable String segmentInstanceId;
        private @Nullable String status;
        private @Nullable Map<String,String> tags;
        public Builder() {}
        public Builder(GetEipAddressesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.addressName = defaults.addressName;
    	      this.addresses = defaults.addresses;
    	      this.associatedInstanceId = defaults.associatedInstanceId;
    	      this.associatedInstanceType = defaults.associatedInstanceType;
    	      this.dryRun = defaults.dryRun;
    	      this.eips = defaults.eips;
    	      this.enableDetails = defaults.enableDetails;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.includeReservationData = defaults.includeReservationData;
    	      this.ipAddress = defaults.ipAddress;
    	      this.ipAddresses = defaults.ipAddresses;
    	      this.isp = defaults.isp;
    	      this.lockReason = defaults.lockReason;
    	      this.nameRegex = defaults.nameRegex;
    	      this.names = defaults.names;
    	      this.outputFile = defaults.outputFile;
    	      this.paymentType = defaults.paymentType;
    	      this.resourceGroupId = defaults.resourceGroupId;
    	      this.segmentInstanceId = defaults.segmentInstanceId;
    	      this.status = defaults.status;
    	      this.tags = defaults.tags;
        }

        @CustomType.Setter
        public Builder addressName(@Nullable String addressName) {

            this.addressName = addressName;
            return this;
        }
        @CustomType.Setter
        public Builder addresses(List<GetEipAddressesAddress> addresses) {
            if (addresses == null) {
              throw new MissingRequiredPropertyException("GetEipAddressesResult", "addresses");
            }
            this.addresses = addresses;
            return this;
        }
        public Builder addresses(GetEipAddressesAddress... addresses) {
            return addresses(List.of(addresses));
        }
        @CustomType.Setter
        public Builder associatedInstanceId(@Nullable String associatedInstanceId) {

            this.associatedInstanceId = associatedInstanceId;
            return this;
        }
        @CustomType.Setter
        public Builder associatedInstanceType(@Nullable String associatedInstanceType) {

            this.associatedInstanceType = associatedInstanceType;
            return this;
        }
        @CustomType.Setter
        public Builder dryRun(@Nullable Boolean dryRun) {

            this.dryRun = dryRun;
            return this;
        }
        @CustomType.Setter
        public Builder eips(List<GetEipAddressesEip> eips) {
            if (eips == null) {
              throw new MissingRequiredPropertyException("GetEipAddressesResult", "eips");
            }
            this.eips = eips;
            return this;
        }
        public Builder eips(GetEipAddressesEip... eips) {
            return eips(List.of(eips));
        }
        @CustomType.Setter
        public Builder enableDetails(@Nullable Boolean enableDetails) {

            this.enableDetails = enableDetails;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetEipAddressesResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder ids(List<String> ids) {
            if (ids == null) {
              throw new MissingRequiredPropertyException("GetEipAddressesResult", "ids");
            }
            this.ids = ids;
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
        public Builder ipAddress(@Nullable String ipAddress) {

            this.ipAddress = ipAddress;
            return this;
        }
        @CustomType.Setter
        public Builder ipAddresses(@Nullable List<String> ipAddresses) {

            this.ipAddresses = ipAddresses;
            return this;
        }
        public Builder ipAddresses(String... ipAddresses) {
            return ipAddresses(List.of(ipAddresses));
        }
        @CustomType.Setter
        public Builder isp(@Nullable String isp) {

            this.isp = isp;
            return this;
        }
        @CustomType.Setter
        public Builder lockReason(@Nullable String lockReason) {

            this.lockReason = lockReason;
            return this;
        }
        @CustomType.Setter
        public Builder nameRegex(@Nullable String nameRegex) {

            this.nameRegex = nameRegex;
            return this;
        }
        @CustomType.Setter
        public Builder names(List<String> names) {
            if (names == null) {
              throw new MissingRequiredPropertyException("GetEipAddressesResult", "names");
            }
            this.names = names;
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
        public Builder paymentType(@Nullable String paymentType) {

            this.paymentType = paymentType;
            return this;
        }
        @CustomType.Setter
        public Builder resourceGroupId(@Nullable String resourceGroupId) {

            this.resourceGroupId = resourceGroupId;
            return this;
        }
        @CustomType.Setter
        public Builder segmentInstanceId(@Nullable String segmentInstanceId) {

            this.segmentInstanceId = segmentInstanceId;
            return this;
        }
        @CustomType.Setter
        public Builder status(@Nullable String status) {

            this.status = status;
            return this;
        }
        @CustomType.Setter
        public Builder tags(@Nullable Map<String,String> tags) {

            this.tags = tags;
            return this;
        }
        public GetEipAddressesResult build() {
            final var _resultValue = new GetEipAddressesResult();
            _resultValue.addressName = addressName;
            _resultValue.addresses = addresses;
            _resultValue.associatedInstanceId = associatedInstanceId;
            _resultValue.associatedInstanceType = associatedInstanceType;
            _resultValue.dryRun = dryRun;
            _resultValue.eips = eips;
            _resultValue.enableDetails = enableDetails;
            _resultValue.id = id;
            _resultValue.ids = ids;
            _resultValue.includeReservationData = includeReservationData;
            _resultValue.ipAddress = ipAddress;
            _resultValue.ipAddresses = ipAddresses;
            _resultValue.isp = isp;
            _resultValue.lockReason = lockReason;
            _resultValue.nameRegex = nameRegex;
            _resultValue.names = names;
            _resultValue.outputFile = outputFile;
            _resultValue.paymentType = paymentType;
            _resultValue.resourceGroupId = resourceGroupId;
            _resultValue.segmentInstanceId = segmentInstanceId;
            _resultValue.status = status;
            _resultValue.tags = tags;
            return _resultValue;
        }
    }
}

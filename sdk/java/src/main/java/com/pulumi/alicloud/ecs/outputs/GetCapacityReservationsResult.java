// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs.outputs;

import com.pulumi.alicloud.ecs.outputs.GetCapacityReservationsReservation;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetCapacityReservationsResult {
    private @Nullable List<String> capacityReservationIds;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return A list of Capacity Reservation IDs.
     * 
     */
    private List<String> ids;
    /**
     * @return Instance type. Currently, you can only set the capacity reservation service for one instance type.
     * 
     */
    private @Nullable String instanceType;
    private @Nullable String nameRegex;
    /**
     * @return A list of name of Capacity Reservations.
     * 
     */
    private List<String> names;
    private @Nullable String outputFile;
    /**
     * @return The payment type of the resource
     * 
     */
    private @Nullable String paymentType;
    /**
     * @return platform of the capacity reservation.
     * 
     */
    private @Nullable String platform;
    /**
     * @return A list of Capacity Reservation Entries. Each element contains the following attributes:
     * 
     */
    private List<GetCapacityReservationsReservation> reservations;
    /**
     * @return The resource group id
     * 
     */
    private @Nullable String resourceGroupId;
    /**
     * @return The status of the capacity reservation.
     * 
     */
    private @Nullable String status;
    /**
     * @return A mapping of tags to assign to the Capacity Reservation.
     * 
     */
    private @Nullable Map<String,String> tags;

    private GetCapacityReservationsResult() {}
    public List<String> capacityReservationIds() {
        return this.capacityReservationIds == null ? List.of() : this.capacityReservationIds;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return A list of Capacity Reservation IDs.
     * 
     */
    public List<String> ids() {
        return this.ids;
    }
    /**
     * @return Instance type. Currently, you can only set the capacity reservation service for one instance type.
     * 
     */
    public Optional<String> instanceType() {
        return Optional.ofNullable(this.instanceType);
    }
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }
    /**
     * @return A list of name of Capacity Reservations.
     * 
     */
    public List<String> names() {
        return this.names;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    /**
     * @return The payment type of the resource
     * 
     */
    public Optional<String> paymentType() {
        return Optional.ofNullable(this.paymentType);
    }
    /**
     * @return platform of the capacity reservation.
     * 
     */
    public Optional<String> platform() {
        return Optional.ofNullable(this.platform);
    }
    /**
     * @return A list of Capacity Reservation Entries. Each element contains the following attributes:
     * 
     */
    public List<GetCapacityReservationsReservation> reservations() {
        return this.reservations;
    }
    /**
     * @return The resource group id
     * 
     */
    public Optional<String> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }
    /**
     * @return The status of the capacity reservation.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }
    /**
     * @return A mapping of tags to assign to the Capacity Reservation.
     * 
     */
    public Map<String,String> tags() {
        return this.tags == null ? Map.of() : this.tags;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCapacityReservationsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> capacityReservationIds;
        private String id;
        private List<String> ids;
        private @Nullable String instanceType;
        private @Nullable String nameRegex;
        private List<String> names;
        private @Nullable String outputFile;
        private @Nullable String paymentType;
        private @Nullable String platform;
        private List<GetCapacityReservationsReservation> reservations;
        private @Nullable String resourceGroupId;
        private @Nullable String status;
        private @Nullable Map<String,String> tags;
        public Builder() {}
        public Builder(GetCapacityReservationsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.capacityReservationIds = defaults.capacityReservationIds;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.instanceType = defaults.instanceType;
    	      this.nameRegex = defaults.nameRegex;
    	      this.names = defaults.names;
    	      this.outputFile = defaults.outputFile;
    	      this.paymentType = defaults.paymentType;
    	      this.platform = defaults.platform;
    	      this.reservations = defaults.reservations;
    	      this.resourceGroupId = defaults.resourceGroupId;
    	      this.status = defaults.status;
    	      this.tags = defaults.tags;
        }

        @CustomType.Setter
        public Builder capacityReservationIds(@Nullable List<String> capacityReservationIds) {

            this.capacityReservationIds = capacityReservationIds;
            return this;
        }
        public Builder capacityReservationIds(String... capacityReservationIds) {
            return capacityReservationIds(List.of(capacityReservationIds));
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetCapacityReservationsResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder ids(List<String> ids) {
            if (ids == null) {
              throw new MissingRequiredPropertyException("GetCapacityReservationsResult", "ids");
            }
            this.ids = ids;
            return this;
        }
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }
        @CustomType.Setter
        public Builder instanceType(@Nullable String instanceType) {

            this.instanceType = instanceType;
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
              throw new MissingRequiredPropertyException("GetCapacityReservationsResult", "names");
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
        public Builder platform(@Nullable String platform) {

            this.platform = platform;
            return this;
        }
        @CustomType.Setter
        public Builder reservations(List<GetCapacityReservationsReservation> reservations) {
            if (reservations == null) {
              throw new MissingRequiredPropertyException("GetCapacityReservationsResult", "reservations");
            }
            this.reservations = reservations;
            return this;
        }
        public Builder reservations(GetCapacityReservationsReservation... reservations) {
            return reservations(List.of(reservations));
        }
        @CustomType.Setter
        public Builder resourceGroupId(@Nullable String resourceGroupId) {

            this.resourceGroupId = resourceGroupId;
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
        public GetCapacityReservationsResult build() {
            final var _resultValue = new GetCapacityReservationsResult();
            _resultValue.capacityReservationIds = capacityReservationIds;
            _resultValue.id = id;
            _resultValue.ids = ids;
            _resultValue.instanceType = instanceType;
            _resultValue.nameRegex = nameRegex;
            _resultValue.names = names;
            _resultValue.outputFile = outputFile;
            _resultValue.paymentType = paymentType;
            _resultValue.platform = platform;
            _resultValue.reservations = reservations;
            _resultValue.resourceGroupId = resourceGroupId;
            _resultValue.status = status;
            _resultValue.tags = tags;
            return _resultValue;
        }
    }
}

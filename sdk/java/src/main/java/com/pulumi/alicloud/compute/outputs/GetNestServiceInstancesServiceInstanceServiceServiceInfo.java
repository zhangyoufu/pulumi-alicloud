// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.compute.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetNestServiceInstancesServiceInstanceServiceServiceInfo {
    /**
     * @return The image of the service.
     * 
     */
    private String image;
    /**
     * @return The locale of the service.
     * 
     */
    private String locale;
    /**
     * @return The name of the filter. Valid Values: `Name`, `ServiceInstanceName`, `ServiceInstanceId`, `ServiceId`, `Version`, `Status`, `DeployType`, `ServiceType`, `OperationStartTimeBefore`, `OperationStartTimeAfter`, `OperationEndTimeBefore`, `OperationEndTimeAfter`, `OperatedServiceInstanceId`, `OperationServiceInstanceId`, `EnableInstanceOps`.
     * 
     */
    private String name;
    /**
     * @return The short description of the service.
     * 
     */
    private String shortDescription;

    private GetNestServiceInstancesServiceInstanceServiceServiceInfo() {}
    /**
     * @return The image of the service.
     * 
     */
    public String image() {
        return this.image;
    }
    /**
     * @return The locale of the service.
     * 
     */
    public String locale() {
        return this.locale;
    }
    /**
     * @return The name of the filter. Valid Values: `Name`, `ServiceInstanceName`, `ServiceInstanceId`, `ServiceId`, `Version`, `Status`, `DeployType`, `ServiceType`, `OperationStartTimeBefore`, `OperationStartTimeAfter`, `OperationEndTimeBefore`, `OperationEndTimeAfter`, `OperatedServiceInstanceId`, `OperationServiceInstanceId`, `EnableInstanceOps`.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return The short description of the service.
     * 
     */
    public String shortDescription() {
        return this.shortDescription;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetNestServiceInstancesServiceInstanceServiceServiceInfo defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String image;
        private String locale;
        private String name;
        private String shortDescription;
        public Builder() {}
        public Builder(GetNestServiceInstancesServiceInstanceServiceServiceInfo defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.image = defaults.image;
    	      this.locale = defaults.locale;
    	      this.name = defaults.name;
    	      this.shortDescription = defaults.shortDescription;
        }

        @CustomType.Setter
        public Builder image(String image) {
            this.image = Objects.requireNonNull(image);
            return this;
        }
        @CustomType.Setter
        public Builder locale(String locale) {
            this.locale = Objects.requireNonNull(locale);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder shortDescription(String shortDescription) {
            this.shortDescription = Objects.requireNonNull(shortDescription);
            return this;
        }
        public GetNestServiceInstancesServiceInstanceServiceServiceInfo build() {
            final var o = new GetNestServiceInstancesServiceInstanceServiceServiceInfo();
            o.image = image;
            o.locale = locale;
            o.name = name;
            o.shortDescription = shortDescription;
            return o;
        }
    }
}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ga.outputs;

import com.pulumi.alicloud.ga.outputs.GetBasicEndpointsEndpoint;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetBasicEndpointsResult {
    /**
     * @return The ID of the Basic Endpoint Group.
     * 
     */
    private String endpointGroupId;
    /**
     * @return The ID of the Basic Endpoint.
     * 
     */
    private @Nullable String endpointId;
    /**
     * @return The type of the Basic Endpoint.
     * 
     */
    private @Nullable String endpointType;
    /**
     * @return A list of Global Accelerator Basic Endpoints. Each element contains the following attributes:
     * 
     */
    private List<GetBasicEndpointsEndpoint> endpoints;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private List<String> ids;
    private @Nullable String name;
    private @Nullable String nameRegex;
    /**
     * @return A list of Global Accelerator Basic Endpoint names.
     * 
     */
    private List<String> names;
    private @Nullable String outputFile;
    /**
     * @return The status of the Basic Endpoint.
     * 
     */
    private @Nullable String status;

    private GetBasicEndpointsResult() {}
    /**
     * @return The ID of the Basic Endpoint Group.
     * 
     */
    public String endpointGroupId() {
        return this.endpointGroupId;
    }
    /**
     * @return The ID of the Basic Endpoint.
     * 
     */
    public Optional<String> endpointId() {
        return Optional.ofNullable(this.endpointId);
    }
    /**
     * @return The type of the Basic Endpoint.
     * 
     */
    public Optional<String> endpointType() {
        return Optional.ofNullable(this.endpointType);
    }
    /**
     * @return A list of Global Accelerator Basic Endpoints. Each element contains the following attributes:
     * 
     */
    public List<GetBasicEndpointsEndpoint> endpoints() {
        return this.endpoints;
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
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }
    /**
     * @return A list of Global Accelerator Basic Endpoint names.
     * 
     */
    public List<String> names() {
        return this.names;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    /**
     * @return The status of the Basic Endpoint.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBasicEndpointsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String endpointGroupId;
        private @Nullable String endpointId;
        private @Nullable String endpointType;
        private List<GetBasicEndpointsEndpoint> endpoints;
        private String id;
        private List<String> ids;
        private @Nullable String name;
        private @Nullable String nameRegex;
        private List<String> names;
        private @Nullable String outputFile;
        private @Nullable String status;
        public Builder() {}
        public Builder(GetBasicEndpointsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.endpointGroupId = defaults.endpointGroupId;
    	      this.endpointId = defaults.endpointId;
    	      this.endpointType = defaults.endpointType;
    	      this.endpoints = defaults.endpoints;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.name = defaults.name;
    	      this.nameRegex = defaults.nameRegex;
    	      this.names = defaults.names;
    	      this.outputFile = defaults.outputFile;
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder endpointGroupId(String endpointGroupId) {
            this.endpointGroupId = Objects.requireNonNull(endpointGroupId);
            return this;
        }
        @CustomType.Setter
        public Builder endpointId(@Nullable String endpointId) {
            this.endpointId = endpointId;
            return this;
        }
        @CustomType.Setter
        public Builder endpointType(@Nullable String endpointType) {
            this.endpointType = endpointType;
            return this;
        }
        @CustomType.Setter
        public Builder endpoints(List<GetBasicEndpointsEndpoint> endpoints) {
            this.endpoints = Objects.requireNonNull(endpoints);
            return this;
        }
        public Builder endpoints(GetBasicEndpointsEndpoint... endpoints) {
            return endpoints(List.of(endpoints));
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
        public Builder name(@Nullable String name) {
            this.name = name;
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
        public GetBasicEndpointsResult build() {
            final var o = new GetBasicEndpointsResult();
            o.endpointGroupId = endpointGroupId;
            o.endpointId = endpointId;
            o.endpointType = endpointType;
            o.endpoints = endpoints;
            o.id = id;
            o.ids = ids;
            o.name = name;
            o.nameRegex = nameRegex;
            o.names = names;
            o.outputFile = outputFile;
            o.status = status;
            return o;
        }
    }
}

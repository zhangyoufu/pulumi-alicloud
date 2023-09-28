// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.resourcemanager.outputs;

import com.pulumi.alicloud.resourcemanager.outputs.GetSharedTargetsTarget;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetSharedTargetsResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private List<String> ids;
    private @Nullable String outputFile;
    /**
     * @return The resource shared ID of resource manager.
     * 
     */
    private @Nullable String resourceShareId;
    /**
     * @return The status of shared target.
     * 
     */
    private @Nullable String status;
    /**
     * @return A list of Resource Manager Shared Targets. Each element contains the following attributes:
     * 
     */
    private List<GetSharedTargetsTarget> targets;

    private GetSharedTargetsResult() {}
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
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    /**
     * @return The resource shared ID of resource manager.
     * 
     */
    public Optional<String> resourceShareId() {
        return Optional.ofNullable(this.resourceShareId);
    }
    /**
     * @return The status of shared target.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }
    /**
     * @return A list of Resource Manager Shared Targets. Each element contains the following attributes:
     * 
     */
    public List<GetSharedTargetsTarget> targets() {
        return this.targets;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetSharedTargetsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private List<String> ids;
        private @Nullable String outputFile;
        private @Nullable String resourceShareId;
        private @Nullable String status;
        private List<GetSharedTargetsTarget> targets;
        public Builder() {}
        public Builder(GetSharedTargetsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.outputFile = defaults.outputFile;
    	      this.resourceShareId = defaults.resourceShareId;
    	      this.status = defaults.status;
    	      this.targets = defaults.targets;
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
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
            return this;
        }
        @CustomType.Setter
        public Builder resourceShareId(@Nullable String resourceShareId) {
            this.resourceShareId = resourceShareId;
            return this;
        }
        @CustomType.Setter
        public Builder status(@Nullable String status) {
            this.status = status;
            return this;
        }
        @CustomType.Setter
        public Builder targets(List<GetSharedTargetsTarget> targets) {
            this.targets = Objects.requireNonNull(targets);
            return this;
        }
        public Builder targets(GetSharedTargetsTarget... targets) {
            return targets(List.of(targets));
        }
        public GetSharedTargetsResult build() {
            final var o = new GetSharedTargetsResult();
            o.id = id;
            o.ids = ids;
            o.outputFile = outputFile;
            o.resourceShareId = resourceShareId;
            o.status = status;
            o.targets = targets;
            return o;
        }
    }
}

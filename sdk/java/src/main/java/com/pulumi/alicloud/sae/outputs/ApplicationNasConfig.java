// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.sae.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ApplicationNasConfig {
    /**
     * @return The domain name of the mount target.
     * 
     */
    private @Nullable String mountDomain;
    /**
     * @return The mount path of the container.
     * 
     */
    private @Nullable String mountPath;
    /**
     * @return The ID of the NAS file system.
     * 
     */
    private @Nullable String nasId;
    /**
     * @return The directory in the NAS file system.
     * 
     */
    private @Nullable String nasPath;
    /**
     * @return Specifies whether the application can read data from or write data to resources in the directory of the NAS. Valid values: `true` and `false`. If you set `read_only` to `false`, the application has the read and write permissions.
     * 
     */
    private @Nullable Boolean readOnly;

    private ApplicationNasConfig() {}
    /**
     * @return The domain name of the mount target.
     * 
     */
    public Optional<String> mountDomain() {
        return Optional.ofNullable(this.mountDomain);
    }
    /**
     * @return The mount path of the container.
     * 
     */
    public Optional<String> mountPath() {
        return Optional.ofNullable(this.mountPath);
    }
    /**
     * @return The ID of the NAS file system.
     * 
     */
    public Optional<String> nasId() {
        return Optional.ofNullable(this.nasId);
    }
    /**
     * @return The directory in the NAS file system.
     * 
     */
    public Optional<String> nasPath() {
        return Optional.ofNullable(this.nasPath);
    }
    /**
     * @return Specifies whether the application can read data from or write data to resources in the directory of the NAS. Valid values: `true` and `false`. If you set `read_only` to `false`, the application has the read and write permissions.
     * 
     */
    public Optional<Boolean> readOnly() {
        return Optional.ofNullable(this.readOnly);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ApplicationNasConfig defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String mountDomain;
        private @Nullable String mountPath;
        private @Nullable String nasId;
        private @Nullable String nasPath;
        private @Nullable Boolean readOnly;
        public Builder() {}
        public Builder(ApplicationNasConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.mountDomain = defaults.mountDomain;
    	      this.mountPath = defaults.mountPath;
    	      this.nasId = defaults.nasId;
    	      this.nasPath = defaults.nasPath;
    	      this.readOnly = defaults.readOnly;
        }

        @CustomType.Setter
        public Builder mountDomain(@Nullable String mountDomain) {
            this.mountDomain = mountDomain;
            return this;
        }
        @CustomType.Setter
        public Builder mountPath(@Nullable String mountPath) {
            this.mountPath = mountPath;
            return this;
        }
        @CustomType.Setter
        public Builder nasId(@Nullable String nasId) {
            this.nasId = nasId;
            return this;
        }
        @CustomType.Setter
        public Builder nasPath(@Nullable String nasPath) {
            this.nasPath = nasPath;
            return this;
        }
        @CustomType.Setter
        public Builder readOnly(@Nullable Boolean readOnly) {
            this.readOnly = readOnly;
            return this;
        }
        public ApplicationNasConfig build() {
            final var o = new ApplicationNasConfig();
            o.mountDomain = mountDomain;
            o.mountPath = mountPath;
            o.nasId = nasId;
            o.nasPath = nasPath;
            o.readOnly = readOnly;
            return o;
        }
    }
}

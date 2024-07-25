// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.maxcompute.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ProjectPropertiesEncryption {
    /**
     * @return The encryption algorithm supported by the key, including AES256, AESCTR, and RC4.
     * 
     */
    private @Nullable String algorithm;
    /**
     * @return Only enable function is supported. Value: (true).
     * 
     */
    private @Nullable Boolean enable;
    /**
     * @return The encryption algorithm Key, the Key type used by the project, including the Default Key (MaxCompute Default Key) and the self-contained Key (BYOK). The MaxCompute Default Key is the Default Key created inside MaxCompute.
     * 
     */
    private @Nullable String key;

    private ProjectPropertiesEncryption() {}
    /**
     * @return The encryption algorithm supported by the key, including AES256, AESCTR, and RC4.
     * 
     */
    public Optional<String> algorithm() {
        return Optional.ofNullable(this.algorithm);
    }
    /**
     * @return Only enable function is supported. Value: (true).
     * 
     */
    public Optional<Boolean> enable() {
        return Optional.ofNullable(this.enable);
    }
    /**
     * @return The encryption algorithm Key, the Key type used by the project, including the Default Key (MaxCompute Default Key) and the self-contained Key (BYOK). The MaxCompute Default Key is the Default Key created inside MaxCompute.
     * 
     */
    public Optional<String> key() {
        return Optional.ofNullable(this.key);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ProjectPropertiesEncryption defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String algorithm;
        private @Nullable Boolean enable;
        private @Nullable String key;
        public Builder() {}
        public Builder(ProjectPropertiesEncryption defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.algorithm = defaults.algorithm;
    	      this.enable = defaults.enable;
    	      this.key = defaults.key;
        }

        @CustomType.Setter
        public Builder algorithm(@Nullable String algorithm) {

            this.algorithm = algorithm;
            return this;
        }
        @CustomType.Setter
        public Builder enable(@Nullable Boolean enable) {

            this.enable = enable;
            return this;
        }
        @CustomType.Setter
        public Builder key(@Nullable String key) {

            this.key = key;
            return this;
        }
        public ProjectPropertiesEncryption build() {
            final var _resultValue = new ProjectPropertiesEncryption();
            _resultValue.algorithm = algorithm;
            _resultValue.enable = enable;
            _resultValue.key = key;
            return _resultValue;
        }
    }
}

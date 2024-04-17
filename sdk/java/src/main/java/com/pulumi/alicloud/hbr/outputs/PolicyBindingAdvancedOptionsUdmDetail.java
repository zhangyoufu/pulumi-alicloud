// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.hbr.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class PolicyBindingAdvancedOptionsUdmDetail {
    /**
     * @return Custom KMS key ID of encrypted copy.
     * 
     */
    private @Nullable String destinationKmsKeyId;
    /**
     * @return The list of backup disks. If it is empty, all disks are backed up.
     * 
     */
    private @Nullable List<String> diskIdLists;
    /**
     * @return List of cloud disk IDs that are not backed up.
     * 
     */
    private @Nullable List<String> excludeDiskIdLists;

    private PolicyBindingAdvancedOptionsUdmDetail() {}
    /**
     * @return Custom KMS key ID of encrypted copy.
     * 
     */
    public Optional<String> destinationKmsKeyId() {
        return Optional.ofNullable(this.destinationKmsKeyId);
    }
    /**
     * @return The list of backup disks. If it is empty, all disks are backed up.
     * 
     */
    public List<String> diskIdLists() {
        return this.diskIdLists == null ? List.of() : this.diskIdLists;
    }
    /**
     * @return List of cloud disk IDs that are not backed up.
     * 
     */
    public List<String> excludeDiskIdLists() {
        return this.excludeDiskIdLists == null ? List.of() : this.excludeDiskIdLists;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PolicyBindingAdvancedOptionsUdmDetail defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String destinationKmsKeyId;
        private @Nullable List<String> diskIdLists;
        private @Nullable List<String> excludeDiskIdLists;
        public Builder() {}
        public Builder(PolicyBindingAdvancedOptionsUdmDetail defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.destinationKmsKeyId = defaults.destinationKmsKeyId;
    	      this.diskIdLists = defaults.diskIdLists;
    	      this.excludeDiskIdLists = defaults.excludeDiskIdLists;
        }

        @CustomType.Setter
        public Builder destinationKmsKeyId(@Nullable String destinationKmsKeyId) {

            this.destinationKmsKeyId = destinationKmsKeyId;
            return this;
        }
        @CustomType.Setter
        public Builder diskIdLists(@Nullable List<String> diskIdLists) {

            this.diskIdLists = diskIdLists;
            return this;
        }
        public Builder diskIdLists(String... diskIdLists) {
            return diskIdLists(List.of(diskIdLists));
        }
        @CustomType.Setter
        public Builder excludeDiskIdLists(@Nullable List<String> excludeDiskIdLists) {

            this.excludeDiskIdLists = excludeDiskIdLists;
            return this;
        }
        public Builder excludeDiskIdLists(String... excludeDiskIdLists) {
            return excludeDiskIdLists(List.of(excludeDiskIdLists));
        }
        public PolicyBindingAdvancedOptionsUdmDetail build() {
            final var _resultValue = new PolicyBindingAdvancedOptionsUdmDetail();
            _resultValue.destinationKmsKeyId = destinationKmsKeyId;
            _resultValue.diskIdLists = diskIdLists;
            _resultValue.excludeDiskIdLists = excludeDiskIdLists;
            return _resultValue;
        }
    }
}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.apigateway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class AccessControlListAclEntry {
    /**
     * @return The description of the ACL.
     * 
     */
    private @Nullable String aclEntryComment;
    /**
     * @return The entries that you want to add to the ACL. You can add CIDR blocks. Separate multiple CIDR blocks with commas (,).
     * 
     */
    private @Nullable String aclEntryIp;

    private AccessControlListAclEntry() {}
    /**
     * @return The description of the ACL.
     * 
     */
    public Optional<String> aclEntryComment() {
        return Optional.ofNullable(this.aclEntryComment);
    }
    /**
     * @return The entries that you want to add to the ACL. You can add CIDR blocks. Separate multiple CIDR blocks with commas (,).
     * 
     */
    public Optional<String> aclEntryIp() {
        return Optional.ofNullable(this.aclEntryIp);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(AccessControlListAclEntry defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String aclEntryComment;
        private @Nullable String aclEntryIp;
        public Builder() {}
        public Builder(AccessControlListAclEntry defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.aclEntryComment = defaults.aclEntryComment;
    	      this.aclEntryIp = defaults.aclEntryIp;
        }

        @CustomType.Setter
        public Builder aclEntryComment(@Nullable String aclEntryComment) {

            this.aclEntryComment = aclEntryComment;
            return this;
        }
        @CustomType.Setter
        public Builder aclEntryIp(@Nullable String aclEntryIp) {

            this.aclEntryIp = aclEntryIp;
            return this;
        }
        public AccessControlListAclEntry build() {
            final var _resultValue = new AccessControlListAclEntry();
            _resultValue.aclEntryComment = aclEntryComment;
            _resultValue.aclEntryIp = aclEntryIp;
            return _resultValue;
        }
    }
}

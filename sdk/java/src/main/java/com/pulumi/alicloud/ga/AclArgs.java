// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ga;

import com.pulumi.alicloud.ga.inputs.AclAclEntryArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AclArgs extends com.pulumi.resources.ResourceArgs {

    public static final AclArgs Empty = new AclArgs();

    /**
     * The entries of the Acl. See `acl_entries` below. **NOTE:** &#34;Field `acl_entries` has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource `alicloud.ga.AclEntryAttachment`.&#34;
     * 
     * @deprecated
     * Field `acl_entries` has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource `alicloud.ga.AclEntryAttachment`.
     * 
     */
    @Deprecated /* Field `acl_entries` has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource `alicloud.ga.AclEntryAttachment`. */
    @Import(name="aclEntries")
    private @Nullable Output<List<AclAclEntryArgs>> aclEntries;

    /**
     * @return The entries of the Acl. See `acl_entries` below. **NOTE:** &#34;Field `acl_entries` has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource `alicloud.ga.AclEntryAttachment`.&#34;
     * 
     * @deprecated
     * Field `acl_entries` has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource `alicloud.ga.AclEntryAttachment`.
     * 
     */
    @Deprecated /* Field `acl_entries` has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource `alicloud.ga.AclEntryAttachment`. */
    public Optional<Output<List<AclAclEntryArgs>>> aclEntries() {
        return Optional.ofNullable(this.aclEntries);
    }

    /**
     * The name of the ACL. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), hyphens (-) and underscores (_). It must start with a letter.
     * 
     */
    @Import(name="aclName")
    private @Nullable Output<String> aclName;

    /**
     * @return The name of the ACL. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), hyphens (-) and underscores (_). It must start with a letter.
     * 
     */
    public Optional<Output<String>> aclName() {
        return Optional.ofNullable(this.aclName);
    }

    /**
     * The IP version. Valid values: `IPv4` and `IPv6`.
     * 
     */
    @Import(name="addressIpVersion", required=true)
    private Output<String> addressIpVersion;

    /**
     * @return The IP version. Valid values: `IPv4` and `IPv6`.
     * 
     */
    public Output<String> addressIpVersion() {
        return this.addressIpVersion;
    }

    /**
     * The dry run.
     * 
     */
    @Import(name="dryRun")
    private @Nullable Output<Boolean> dryRun;

    /**
     * @return The dry run.
     * 
     */
    public Optional<Output<Boolean>> dryRun() {
        return Optional.ofNullable(this.dryRun);
    }

    /**
     * The ID of the resource group. **Note:** Once you set a value of this property, you cannot set it to an empty string anymore.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group. **Note:** Once you set a value of this property, you cannot set it to an empty string anymore.
     * 
     */
    public Optional<Output<String>> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private AclArgs() {}

    private AclArgs(AclArgs $) {
        this.aclEntries = $.aclEntries;
        this.aclName = $.aclName;
        this.addressIpVersion = $.addressIpVersion;
        this.dryRun = $.dryRun;
        this.resourceGroupId = $.resourceGroupId;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AclArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AclArgs $;

        public Builder() {
            $ = new AclArgs();
        }

        public Builder(AclArgs defaults) {
            $ = new AclArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param aclEntries The entries of the Acl. See `acl_entries` below. **NOTE:** &#34;Field `acl_entries` has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource `alicloud.ga.AclEntryAttachment`.&#34;
         * 
         * @return builder
         * 
         * @deprecated
         * Field `acl_entries` has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource `alicloud.ga.AclEntryAttachment`.
         * 
         */
        @Deprecated /* Field `acl_entries` has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource `alicloud.ga.AclEntryAttachment`. */
        public Builder aclEntries(@Nullable Output<List<AclAclEntryArgs>> aclEntries) {
            $.aclEntries = aclEntries;
            return this;
        }

        /**
         * @param aclEntries The entries of the Acl. See `acl_entries` below. **NOTE:** &#34;Field `acl_entries` has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource `alicloud.ga.AclEntryAttachment`.&#34;
         * 
         * @return builder
         * 
         * @deprecated
         * Field `acl_entries` has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource `alicloud.ga.AclEntryAttachment`.
         * 
         */
        @Deprecated /* Field `acl_entries` has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource `alicloud.ga.AclEntryAttachment`. */
        public Builder aclEntries(List<AclAclEntryArgs> aclEntries) {
            return aclEntries(Output.of(aclEntries));
        }

        /**
         * @param aclEntries The entries of the Acl. See `acl_entries` below. **NOTE:** &#34;Field `acl_entries` has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource `alicloud.ga.AclEntryAttachment`.&#34;
         * 
         * @return builder
         * 
         * @deprecated
         * Field `acl_entries` has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource `alicloud.ga.AclEntryAttachment`.
         * 
         */
        @Deprecated /* Field `acl_entries` has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource `alicloud.ga.AclEntryAttachment`. */
        public Builder aclEntries(AclAclEntryArgs... aclEntries) {
            return aclEntries(List.of(aclEntries));
        }

        /**
         * @param aclName The name of the ACL. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), hyphens (-) and underscores (_). It must start with a letter.
         * 
         * @return builder
         * 
         */
        public Builder aclName(@Nullable Output<String> aclName) {
            $.aclName = aclName;
            return this;
        }

        /**
         * @param aclName The name of the ACL. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), hyphens (-) and underscores (_). It must start with a letter.
         * 
         * @return builder
         * 
         */
        public Builder aclName(String aclName) {
            return aclName(Output.of(aclName));
        }

        /**
         * @param addressIpVersion The IP version. Valid values: `IPv4` and `IPv6`.
         * 
         * @return builder
         * 
         */
        public Builder addressIpVersion(Output<String> addressIpVersion) {
            $.addressIpVersion = addressIpVersion;
            return this;
        }

        /**
         * @param addressIpVersion The IP version. Valid values: `IPv4` and `IPv6`.
         * 
         * @return builder
         * 
         */
        public Builder addressIpVersion(String addressIpVersion) {
            return addressIpVersion(Output.of(addressIpVersion));
        }

        /**
         * @param dryRun The dry run.
         * 
         * @return builder
         * 
         */
        public Builder dryRun(@Nullable Output<Boolean> dryRun) {
            $.dryRun = dryRun;
            return this;
        }

        /**
         * @param dryRun The dry run.
         * 
         * @return builder
         * 
         */
        public Builder dryRun(Boolean dryRun) {
            return dryRun(Output.of(dryRun));
        }

        /**
         * @param resourceGroupId The ID of the resource group. **Note:** Once you set a value of this property, you cannot set it to an empty string anymore.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable Output<String> resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param resourceGroupId The ID of the resource group. **Note:** Once you set a value of this property, you cannot set it to an empty string anymore.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(String resourceGroupId) {
            return resourceGroupId(Output.of(resourceGroupId));
        }

        /**
         * @param tags A mapping of tags to assign to the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A mapping of tags to assign to the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        public AclArgs build() {
            if ($.addressIpVersion == null) {
                throw new MissingRequiredPropertyException("AclArgs", "addressIpVersion");
            }
            return $;
        }
    }

}

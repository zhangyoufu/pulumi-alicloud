// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.resourcemanager.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PolicyVersionState extends com.pulumi.resources.ResourceArgs {

    public static final PolicyVersionState Empty = new PolicyVersionState();

    /**
     * Specifies whether to set the policy version as the default version. Default to `false`.
     * 
     * @deprecated
     * Field &#39;is_default_version&#39; has been deprecated from provider version 1.90.0
     * 
     */
    @Deprecated /* Field 'is_default_version' has been deprecated from provider version 1.90.0 */
    @Import(name="isDefaultVersion")
    private @Nullable Output<Boolean> isDefaultVersion;

    /**
     * @return Specifies whether to set the policy version as the default version. Default to `false`.
     * 
     * @deprecated
     * Field &#39;is_default_version&#39; has been deprecated from provider version 1.90.0
     * 
     */
    @Deprecated /* Field 'is_default_version' has been deprecated from provider version 1.90.0 */
    public Optional<Output<Boolean>> isDefaultVersion() {
        return Optional.ofNullable(this.isDefaultVersion);
    }

    /**
     * The content of the policy. The content must be 1 to 2,048 characters in length.
     * 
     */
    @Import(name="policyDocument")
    private @Nullable Output<String> policyDocument;

    /**
     * @return The content of the policy. The content must be 1 to 2,048 characters in length.
     * 
     */
    public Optional<Output<String>> policyDocument() {
        return Optional.ofNullable(this.policyDocument);
    }

    /**
     * The name of the policy. Name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
     * 
     */
    @Import(name="policyName")
    private @Nullable Output<String> policyName;

    /**
     * @return The name of the policy. Name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
     * 
     */
    public Optional<Output<String>> policyName() {
        return Optional.ofNullable(this.policyName);
    }

    private PolicyVersionState() {}

    private PolicyVersionState(PolicyVersionState $) {
        this.isDefaultVersion = $.isDefaultVersion;
        this.policyDocument = $.policyDocument;
        this.policyName = $.policyName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PolicyVersionState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PolicyVersionState $;

        public Builder() {
            $ = new PolicyVersionState();
        }

        public Builder(PolicyVersionState defaults) {
            $ = new PolicyVersionState(Objects.requireNonNull(defaults));
        }

        /**
         * @param isDefaultVersion Specifies whether to set the policy version as the default version. Default to `false`.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;is_default_version&#39; has been deprecated from provider version 1.90.0
         * 
         */
        @Deprecated /* Field 'is_default_version' has been deprecated from provider version 1.90.0 */
        public Builder isDefaultVersion(@Nullable Output<Boolean> isDefaultVersion) {
            $.isDefaultVersion = isDefaultVersion;
            return this;
        }

        /**
         * @param isDefaultVersion Specifies whether to set the policy version as the default version. Default to `false`.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;is_default_version&#39; has been deprecated from provider version 1.90.0
         * 
         */
        @Deprecated /* Field 'is_default_version' has been deprecated from provider version 1.90.0 */
        public Builder isDefaultVersion(Boolean isDefaultVersion) {
            return isDefaultVersion(Output.of(isDefaultVersion));
        }

        /**
         * @param policyDocument The content of the policy. The content must be 1 to 2,048 characters in length.
         * 
         * @return builder
         * 
         */
        public Builder policyDocument(@Nullable Output<String> policyDocument) {
            $.policyDocument = policyDocument;
            return this;
        }

        /**
         * @param policyDocument The content of the policy. The content must be 1 to 2,048 characters in length.
         * 
         * @return builder
         * 
         */
        public Builder policyDocument(String policyDocument) {
            return policyDocument(Output.of(policyDocument));
        }

        /**
         * @param policyName The name of the policy. Name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
         * 
         * @return builder
         * 
         */
        public Builder policyName(@Nullable Output<String> policyName) {
            $.policyName = policyName;
            return this;
        }

        /**
         * @param policyName The name of the policy. Name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
         * 
         * @return builder
         * 
         */
        public Builder policyName(String policyName) {
            return policyName(Output.of(policyName));
        }

        public PolicyVersionState build() {
            return $;
        }
    }

}

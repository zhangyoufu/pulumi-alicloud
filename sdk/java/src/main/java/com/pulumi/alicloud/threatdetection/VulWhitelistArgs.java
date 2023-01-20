// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.threatdetection;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VulWhitelistArgs extends com.pulumi.resources.ResourceArgs {

    public static final VulWhitelistArgs Empty = new VulWhitelistArgs();

    /**
     * Reason for adding whitelist.
     * 
     */
    @Import(name="reason")
    private @Nullable Output<String> reason;

    /**
     * @return Reason for adding whitelist.
     * 
     */
    public Optional<Output<String>> reason() {
        return Optional.ofNullable(this.reason);
    }

    /**
     * Set the effective range of the whitelist. see [how to use it](https://www.alibabacloud.com/help/en/security-center/latest/api-doc-sas-2018-12-03-api-doc-modifycreatevulwhitelist).
     * 
     */
    @Import(name="targetInfo")
    private @Nullable Output<String> targetInfo;

    /**
     * @return Set the effective range of the whitelist. see [how to use it](https://www.alibabacloud.com/help/en/security-center/latest/api-doc-sas-2018-12-03-api-doc-modifycreatevulwhitelist).
     * 
     */
    public Optional<Output<String>> targetInfo() {
        return Optional.ofNullable(this.targetInfo);
    }

    /**
     * Information about the vulnerability to be added to the whitelist. see [how to use it](https://www.alibabacloud.com/help/en/security-center/latest/api-doc-sas-2018-12-03-api-doc-modifycreatevulwhitelist).
     * 
     */
    @Import(name="whitelist", required=true)
    private Output<String> whitelist;

    /**
     * @return Information about the vulnerability to be added to the whitelist. see [how to use it](https://www.alibabacloud.com/help/en/security-center/latest/api-doc-sas-2018-12-03-api-doc-modifycreatevulwhitelist).
     * 
     */
    public Output<String> whitelist() {
        return this.whitelist;
    }

    private VulWhitelistArgs() {}

    private VulWhitelistArgs(VulWhitelistArgs $) {
        this.reason = $.reason;
        this.targetInfo = $.targetInfo;
        this.whitelist = $.whitelist;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VulWhitelistArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VulWhitelistArgs $;

        public Builder() {
            $ = new VulWhitelistArgs();
        }

        public Builder(VulWhitelistArgs defaults) {
            $ = new VulWhitelistArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param reason Reason for adding whitelist.
         * 
         * @return builder
         * 
         */
        public Builder reason(@Nullable Output<String> reason) {
            $.reason = reason;
            return this;
        }

        /**
         * @param reason Reason for adding whitelist.
         * 
         * @return builder
         * 
         */
        public Builder reason(String reason) {
            return reason(Output.of(reason));
        }

        /**
         * @param targetInfo Set the effective range of the whitelist. see [how to use it](https://www.alibabacloud.com/help/en/security-center/latest/api-doc-sas-2018-12-03-api-doc-modifycreatevulwhitelist).
         * 
         * @return builder
         * 
         */
        public Builder targetInfo(@Nullable Output<String> targetInfo) {
            $.targetInfo = targetInfo;
            return this;
        }

        /**
         * @param targetInfo Set the effective range of the whitelist. see [how to use it](https://www.alibabacloud.com/help/en/security-center/latest/api-doc-sas-2018-12-03-api-doc-modifycreatevulwhitelist).
         * 
         * @return builder
         * 
         */
        public Builder targetInfo(String targetInfo) {
            return targetInfo(Output.of(targetInfo));
        }

        /**
         * @param whitelist Information about the vulnerability to be added to the whitelist. see [how to use it](https://www.alibabacloud.com/help/en/security-center/latest/api-doc-sas-2018-12-03-api-doc-modifycreatevulwhitelist).
         * 
         * @return builder
         * 
         */
        public Builder whitelist(Output<String> whitelist) {
            $.whitelist = whitelist;
            return this;
        }

        /**
         * @param whitelist Information about the vulnerability to be added to the whitelist. see [how to use it](https://www.alibabacloud.com/help/en/security-center/latest/api-doc-sas-2018-12-03-api-doc-modifycreatevulwhitelist).
         * 
         * @return builder
         * 
         */
        public Builder whitelist(String whitelist) {
            return whitelist(Output.of(whitelist));
        }

        public VulWhitelistArgs build() {
            $.whitelist = Objects.requireNonNull($.whitelist, "expected parameter 'whitelist' to be non-null");
            return $;
        }
    }

}

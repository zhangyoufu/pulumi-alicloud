// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.resourcemanager;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AccountArgs extends com.pulumi.resources.ResourceArgs {

    public static final AccountArgs Empty = new AccountArgs();

    /**
     * The IDs of the check items that you can choose to ignore for the member deletion.
     * If you want to delete the account, please use datasource `alicloud.resourcemanager.getAccountDeletionCheckTask`
     * to get check ids and set them.
     * 
     */
    @Import(name="abandonAbleCheckIds")
    private @Nullable Output<List<String>> abandonAbleCheckIds;

    /**
     * @return The IDs of the check items that you can choose to ignore for the member deletion.
     * If you want to delete the account, please use datasource `alicloud.resourcemanager.getAccountDeletionCheckTask`
     * to get check ids and set them.
     * 
     */
    public Optional<Output<List<String>>> abandonAbleCheckIds() {
        return Optional.ofNullable(this.abandonAbleCheckIds);
    }

    /**
     * The name prefix of account.
     * 
     */
    @Import(name="accountNamePrefix")
    private @Nullable Output<String> accountNamePrefix;

    /**
     * @return The name prefix of account.
     * 
     */
    public Optional<Output<String>> accountNamePrefix() {
        return Optional.ofNullable(this.accountNamePrefix);
    }

    /**
     * Member name. The length is 2 ~ 50 characters or Chinese characters, which can include Chinese characters, English letters, numbers, underscores (_), dots (.) And dashes (-).
     * 
     */
    @Import(name="displayName", required=true)
    private Output<String> displayName;

    /**
     * @return Member name. The length is 2 ~ 50 characters or Chinese characters, which can include Chinese characters, English letters, numbers, underscores (_), dots (.) And dashes (-).
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }

    /**
     * The ID of the parent folder.
     * 
     */
    @Import(name="folderId")
    private @Nullable Output<String> folderId;

    /**
     * @return The ID of the parent folder.
     * 
     */
    public Optional<Output<String>> folderId() {
        return Optional.ofNullable(this.folderId);
    }

    /**
     * The ID of the billing account. If you leave this parameter empty, the current account is used as the billing account.
     * 
     */
    @Import(name="payerAccountId")
    private @Nullable Output<String> payerAccountId;

    /**
     * @return The ID of the billing account. If you leave this parameter empty, the current account is used as the billing account.
     * 
     */
    public Optional<Output<String>> payerAccountId() {
        return Optional.ofNullable(this.payerAccountId);
    }

    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,Object>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Optional<Output<Map<String,Object>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private AccountArgs() {}

    private AccountArgs(AccountArgs $) {
        this.abandonAbleCheckIds = $.abandonAbleCheckIds;
        this.accountNamePrefix = $.accountNamePrefix;
        this.displayName = $.displayName;
        this.folderId = $.folderId;
        this.payerAccountId = $.payerAccountId;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AccountArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AccountArgs $;

        public Builder() {
            $ = new AccountArgs();
        }

        public Builder(AccountArgs defaults) {
            $ = new AccountArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param abandonAbleCheckIds The IDs of the check items that you can choose to ignore for the member deletion.
         * If you want to delete the account, please use datasource `alicloud.resourcemanager.getAccountDeletionCheckTask`
         * to get check ids and set them.
         * 
         * @return builder
         * 
         */
        public Builder abandonAbleCheckIds(@Nullable Output<List<String>> abandonAbleCheckIds) {
            $.abandonAbleCheckIds = abandonAbleCheckIds;
            return this;
        }

        /**
         * @param abandonAbleCheckIds The IDs of the check items that you can choose to ignore for the member deletion.
         * If you want to delete the account, please use datasource `alicloud.resourcemanager.getAccountDeletionCheckTask`
         * to get check ids and set them.
         * 
         * @return builder
         * 
         */
        public Builder abandonAbleCheckIds(List<String> abandonAbleCheckIds) {
            return abandonAbleCheckIds(Output.of(abandonAbleCheckIds));
        }

        /**
         * @param abandonAbleCheckIds The IDs of the check items that you can choose to ignore for the member deletion.
         * If you want to delete the account, please use datasource `alicloud.resourcemanager.getAccountDeletionCheckTask`
         * to get check ids and set them.
         * 
         * @return builder
         * 
         */
        public Builder abandonAbleCheckIds(String... abandonAbleCheckIds) {
            return abandonAbleCheckIds(List.of(abandonAbleCheckIds));
        }

        /**
         * @param accountNamePrefix The name prefix of account.
         * 
         * @return builder
         * 
         */
        public Builder accountNamePrefix(@Nullable Output<String> accountNamePrefix) {
            $.accountNamePrefix = accountNamePrefix;
            return this;
        }

        /**
         * @param accountNamePrefix The name prefix of account.
         * 
         * @return builder
         * 
         */
        public Builder accountNamePrefix(String accountNamePrefix) {
            return accountNamePrefix(Output.of(accountNamePrefix));
        }

        /**
         * @param displayName Member name. The length is 2 ~ 50 characters or Chinese characters, which can include Chinese characters, English letters, numbers, underscores (_), dots (.) And dashes (-).
         * 
         * @return builder
         * 
         */
        public Builder displayName(Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName Member name. The length is 2 ~ 50 characters or Chinese characters, which can include Chinese characters, English letters, numbers, underscores (_), dots (.) And dashes (-).
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param folderId The ID of the parent folder.
         * 
         * @return builder
         * 
         */
        public Builder folderId(@Nullable Output<String> folderId) {
            $.folderId = folderId;
            return this;
        }

        /**
         * @param folderId The ID of the parent folder.
         * 
         * @return builder
         * 
         */
        public Builder folderId(String folderId) {
            return folderId(Output.of(folderId));
        }

        /**
         * @param payerAccountId The ID of the billing account. If you leave this parameter empty, the current account is used as the billing account.
         * 
         * @return builder
         * 
         */
        public Builder payerAccountId(@Nullable Output<String> payerAccountId) {
            $.payerAccountId = payerAccountId;
            return this;
        }

        /**
         * @param payerAccountId The ID of the billing account. If you leave this parameter empty, the current account is used as the billing account.
         * 
         * @return builder
         * 
         */
        public Builder payerAccountId(String payerAccountId) {
            return payerAccountId(Output.of(payerAccountId));
        }

        /**
         * @param tags A mapping of tags to assign to the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,Object>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A mapping of tags to assign to the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,Object> tags) {
            return tags(Output.of(tags));
        }

        public AccountArgs build() {
            $.displayName = Objects.requireNonNull($.displayName, "expected parameter 'displayName' to be non-null");
            return $;
        }
    }

}

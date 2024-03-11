// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.gpdb.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AccountState extends com.pulumi.resources.ResourceArgs {

    public static final AccountState Empty = new AccountState();

    /**
     * The description of the account.
     * * Starts with a letter.
     * * Does not start with `http://` or `https://`.
     * * Contains letters, underscores (_), hyphens (-), or digits.
     * * Be 2 to 256 characters in length.
     * 
     */
    @Import(name="accountDescription")
    private @Nullable Output<String> accountDescription;

    /**
     * @return The description of the account.
     * * Starts with a letter.
     * * Does not start with `http://` or `https://`.
     * * Contains letters, underscores (_), hyphens (-), or digits.
     * * Be 2 to 256 characters in length.
     * 
     */
    public Optional<Output<String>> accountDescription() {
        return Optional.ofNullable(this.accountDescription);
    }

    /**
     * The name of the account. The account name must be unique and meet the following requirements:
     * * Starts with a letter.
     * * Contains only lowercase letters, digits, or underscores (_).
     * * Be up to 16 characters in length.
     * * Contains no reserved keywords.
     * 
     */
    @Import(name="accountName")
    private @Nullable Output<String> accountName;

    /**
     * @return The name of the account. The account name must be unique and meet the following requirements:
     * * Starts with a letter.
     * * Contains only lowercase letters, digits, or underscores (_).
     * * Be up to 16 characters in length.
     * * Contains no reserved keywords.
     * 
     */
    public Optional<Output<String>> accountName() {
        return Optional.ofNullable(this.accountName);
    }

    /**
     * The password of the account. The password must be 8 to 32 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `! @ # $ %!^(MISSING) &amp; * ( ) _ + - =`.
     * 
     */
    @Import(name="accountPassword")
    private @Nullable Output<String> accountPassword;

    /**
     * @return The password of the account. The password must be 8 to 32 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `! @ # $ %!^(MISSING) &amp; * ( ) _ + - =`.
     * 
     */
    public Optional<Output<String>> accountPassword() {
        return Optional.ofNullable(this.accountPassword);
    }

    /**
     * The ID of the instance.
     * 
     */
    @Import(name="dbInstanceId")
    private @Nullable Output<String> dbInstanceId;

    /**
     * @return The ID of the instance.
     * 
     */
    public Optional<Output<String>> dbInstanceId() {
        return Optional.ofNullable(this.dbInstanceId);
    }

    /**
     * The status of the account. Valid values: `Active`, `Creating` and `Deleting`.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the account. Valid values: `Active`, `Creating` and `Deleting`.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    private AccountState() {}

    private AccountState(AccountState $) {
        this.accountDescription = $.accountDescription;
        this.accountName = $.accountName;
        this.accountPassword = $.accountPassword;
        this.dbInstanceId = $.dbInstanceId;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AccountState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AccountState $;

        public Builder() {
            $ = new AccountState();
        }

        public Builder(AccountState defaults) {
            $ = new AccountState(Objects.requireNonNull(defaults));
        }

        /**
         * @param accountDescription The description of the account.
         * * Starts with a letter.
         * * Does not start with `http://` or `https://`.
         * * Contains letters, underscores (_), hyphens (-), or digits.
         * * Be 2 to 256 characters in length.
         * 
         * @return builder
         * 
         */
        public Builder accountDescription(@Nullable Output<String> accountDescription) {
            $.accountDescription = accountDescription;
            return this;
        }

        /**
         * @param accountDescription The description of the account.
         * * Starts with a letter.
         * * Does not start with `http://` or `https://`.
         * * Contains letters, underscores (_), hyphens (-), or digits.
         * * Be 2 to 256 characters in length.
         * 
         * @return builder
         * 
         */
        public Builder accountDescription(String accountDescription) {
            return accountDescription(Output.of(accountDescription));
        }

        /**
         * @param accountName The name of the account. The account name must be unique and meet the following requirements:
         * * Starts with a letter.
         * * Contains only lowercase letters, digits, or underscores (_).
         * * Be up to 16 characters in length.
         * * Contains no reserved keywords.
         * 
         * @return builder
         * 
         */
        public Builder accountName(@Nullable Output<String> accountName) {
            $.accountName = accountName;
            return this;
        }

        /**
         * @param accountName The name of the account. The account name must be unique and meet the following requirements:
         * * Starts with a letter.
         * * Contains only lowercase letters, digits, or underscores (_).
         * * Be up to 16 characters in length.
         * * Contains no reserved keywords.
         * 
         * @return builder
         * 
         */
        public Builder accountName(String accountName) {
            return accountName(Output.of(accountName));
        }

        /**
         * @param accountPassword The password of the account. The password must be 8 to 32 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `! @ # $ %!^(MISSING) &amp; * ( ) _ + - =`.
         * 
         * @return builder
         * 
         */
        public Builder accountPassword(@Nullable Output<String> accountPassword) {
            $.accountPassword = accountPassword;
            return this;
        }

        /**
         * @param accountPassword The password of the account. The password must be 8 to 32 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `! @ # $ %!^(MISSING) &amp; * ( ) _ + - =`.
         * 
         * @return builder
         * 
         */
        public Builder accountPassword(String accountPassword) {
            return accountPassword(Output.of(accountPassword));
        }

        /**
         * @param dbInstanceId The ID of the instance.
         * 
         * @return builder
         * 
         */
        public Builder dbInstanceId(@Nullable Output<String> dbInstanceId) {
            $.dbInstanceId = dbInstanceId;
            return this;
        }

        /**
         * @param dbInstanceId The ID of the instance.
         * 
         * @return builder
         * 
         */
        public Builder dbInstanceId(String dbInstanceId) {
            return dbInstanceId(Output.of(dbInstanceId));
        }

        /**
         * @param status The status of the account. Valid values: `Active`, `Creating` and `Deleting`.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the account. Valid values: `Active`, `Creating` and `Deleting`.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        public AccountState build() {
            return $;
        }
    }

}

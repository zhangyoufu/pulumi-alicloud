// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.kms.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetSecretVersionsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetSecretVersionsPlainArgs Empty = new GetSecretVersionsPlainArgs();

    /**
     * Default to false and only output `secret_name`, `version_id`, `version_stages`. Set it to true can output more details.
     * 
     */
    @Import(name="enableDetails")
    private @Nullable Boolean enableDetails;

    /**
     * @return Default to false and only output `secret_name`, `version_id`, `version_stages`. Set it to true can output more details.
     * 
     */
    public Optional<Boolean> enableDetails() {
        return Optional.ofNullable(this.enableDetails);
    }

    /**
     * A list of KMS Secret Version ids.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of KMS Secret Version ids.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * Specifies whether to return deprecated secret versions. Default to `false`.
     * 
     */
    @Import(name="includeDeprecated")
    private @Nullable String includeDeprecated;

    /**
     * @return Specifies whether to return deprecated secret versions. Default to `false`.
     * 
     */
    public Optional<String> includeDeprecated() {
        return Optional.ofNullable(this.includeDeprecated);
    }

    /**
     * File name where to save data source results (after running `pulumi preview`).
     * 
     */
    @Import(name="outputFile")
    private @Nullable String outputFile;

    /**
     * @return File name where to save data source results (after running `pulumi preview`).
     * 
     */
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    /**
     * The name of the secret.
     * 
     */
    @Import(name="secretName", required=true)
    private String secretName;

    /**
     * @return The name of the secret.
     * 
     */
    public String secretName() {
        return this.secretName;
    }

    /**
     * The stage of the secret version.
     * 
     */
    @Import(name="versionStage")
    private @Nullable String versionStage;

    /**
     * @return The stage of the secret version.
     * 
     */
    public Optional<String> versionStage() {
        return Optional.ofNullable(this.versionStage);
    }

    private GetSecretVersionsPlainArgs() {}

    private GetSecretVersionsPlainArgs(GetSecretVersionsPlainArgs $) {
        this.enableDetails = $.enableDetails;
        this.ids = $.ids;
        this.includeDeprecated = $.includeDeprecated;
        this.outputFile = $.outputFile;
        this.secretName = $.secretName;
        this.versionStage = $.versionStage;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetSecretVersionsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetSecretVersionsPlainArgs $;

        public Builder() {
            $ = new GetSecretVersionsPlainArgs();
        }

        public Builder(GetSecretVersionsPlainArgs defaults) {
            $ = new GetSecretVersionsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param enableDetails Default to false and only output `secret_name`, `version_id`, `version_stages`. Set it to true can output more details.
         * 
         * @return builder
         * 
         */
        public Builder enableDetails(@Nullable Boolean enableDetails) {
            $.enableDetails = enableDetails;
            return this;
        }

        /**
         * @param ids A list of KMS Secret Version ids.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of KMS Secret Version ids.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param includeDeprecated Specifies whether to return deprecated secret versions. Default to `false`.
         * 
         * @return builder
         * 
         */
        public Builder includeDeprecated(@Nullable String includeDeprecated) {
            $.includeDeprecated = includeDeprecated;
            return this;
        }

        /**
         * @param outputFile File name where to save data source results (after running `pulumi preview`).
         * 
         * @return builder
         * 
         */
        public Builder outputFile(@Nullable String outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        /**
         * @param secretName The name of the secret.
         * 
         * @return builder
         * 
         */
        public Builder secretName(String secretName) {
            $.secretName = secretName;
            return this;
        }

        /**
         * @param versionStage The stage of the secret version.
         * 
         * @return builder
         * 
         */
        public Builder versionStage(@Nullable String versionStage) {
            $.versionStage = versionStage;
            return this;
        }

        public GetSecretVersionsPlainArgs build() {
            $.secretName = Objects.requireNonNull($.secretName, "expected parameter 'secretName' to be non-null");
            return $;
        }
    }

}

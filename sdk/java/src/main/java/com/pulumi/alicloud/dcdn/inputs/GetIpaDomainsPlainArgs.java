// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dcdn.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetIpaDomainsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetIpaDomainsPlainArgs Empty = new GetIpaDomainsPlainArgs();

    /**
     * The accelerated domain names.
     * 
     */
    @Import(name="domainName")
    private @Nullable String domainName;

    /**
     * @return The accelerated domain names.
     * 
     */
    public Optional<String> domainName() {
        return Optional.ofNullable(this.domainName);
    }

    /**
     * Default to `false`. Set it to `true` can output more details about resource attributes.
     * 
     */
    @Import(name="enableDetails")
    private @Nullable Boolean enableDetails;

    /**
     * @return Default to `false`. Set it to `true` can output more details about resource attributes.
     * 
     */
    public Optional<Boolean> enableDetails() {
        return Optional.ofNullable(this.enableDetails);
    }

    /**
     * A list of Ipa Domain IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of Ipa Domain IDs.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
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
     * The status of the accelerated domain name.
     * 
     */
    @Import(name="status")
    private @Nullable String status;

    /**
     * @return The status of the accelerated domain name.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    private GetIpaDomainsPlainArgs() {}

    private GetIpaDomainsPlainArgs(GetIpaDomainsPlainArgs $) {
        this.domainName = $.domainName;
        this.enableDetails = $.enableDetails;
        this.ids = $.ids;
        this.outputFile = $.outputFile;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetIpaDomainsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetIpaDomainsPlainArgs $;

        public Builder() {
            $ = new GetIpaDomainsPlainArgs();
        }

        public Builder(GetIpaDomainsPlainArgs defaults) {
            $ = new GetIpaDomainsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param domainName The accelerated domain names.
         * 
         * @return builder
         * 
         */
        public Builder domainName(@Nullable String domainName) {
            $.domainName = domainName;
            return this;
        }

        /**
         * @param enableDetails Default to `false`. Set it to `true` can output more details about resource attributes.
         * 
         * @return builder
         * 
         */
        public Builder enableDetails(@Nullable Boolean enableDetails) {
            $.enableDetails = enableDetails;
            return this;
        }

        /**
         * @param ids A list of Ipa Domain IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Ipa Domain IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
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
         * @param status The status of the accelerated domain name.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable String status) {
            $.status = status;
            return this;
        }

        public GetIpaDomainsPlainArgs build() {
            return $;
        }
    }

}

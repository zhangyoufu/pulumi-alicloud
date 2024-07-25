// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.directmail.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetDomainsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDomainsArgs Empty = new GetDomainsArgs();

    /**
     * Whether to query the detailed list of resource attributes. Default value: `false`.
     * 
     */
    @Import(name="enableDetails")
    private @Nullable Output<Boolean> enableDetails;

    /**
     * @return Whether to query the detailed list of resource attributes. Default value: `false`.
     * 
     */
    public Optional<Output<Boolean>> enableDetails() {
        return Optional.ofNullable(this.enableDetails);
    }

    /**
     * A list of Domain IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable Output<List<String>> ids;

    /**
     * @return A list of Domain IDs.
     * 
     */
    public Optional<Output<List<String>>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * The domain name. It must be 1 to 50 characters in length and can contain digits, letters, periods (.), and hyphens (-).
     * 
     */
    @Import(name="keyWord")
    private @Nullable Output<String> keyWord;

    /**
     * @return The domain name. It must be 1 to 50 characters in length and can contain digits, letters, periods (.), and hyphens (-).
     * 
     */
    public Optional<Output<String>> keyWord() {
        return Optional.ofNullable(this.keyWord);
    }

    /**
     * A regex string to filter results by Domain name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable Output<String> nameRegex;

    /**
     * @return A regex string to filter results by Domain name.
     * 
     */
    public Optional<Output<String>> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }

    /**
     * File name where to save data source results (after running `pulumi preview`).
     * 
     */
    @Import(name="outputFile")
    private @Nullable Output<String> outputFile;

    /**
     * @return File name where to save data source results (after running `pulumi preview`).
     * 
     */
    public Optional<Output<String>> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    /**
     * The status of the domain name. Valid values:
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the domain name. Valid values:
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    private GetDomainsArgs() {}

    private GetDomainsArgs(GetDomainsArgs $) {
        this.enableDetails = $.enableDetails;
        this.ids = $.ids;
        this.keyWord = $.keyWord;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDomainsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDomainsArgs $;

        public Builder() {
            $ = new GetDomainsArgs();
        }

        public Builder(GetDomainsArgs defaults) {
            $ = new GetDomainsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param enableDetails Whether to query the detailed list of resource attributes. Default value: `false`.
         * 
         * @return builder
         * 
         */
        public Builder enableDetails(@Nullable Output<Boolean> enableDetails) {
            $.enableDetails = enableDetails;
            return this;
        }

        /**
         * @param enableDetails Whether to query the detailed list of resource attributes. Default value: `false`.
         * 
         * @return builder
         * 
         */
        public Builder enableDetails(Boolean enableDetails) {
            return enableDetails(Output.of(enableDetails));
        }

        /**
         * @param ids A list of Domain IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable Output<List<String>> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Domain IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(List<String> ids) {
            return ids(Output.of(ids));
        }

        /**
         * @param ids A list of Domain IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param keyWord The domain name. It must be 1 to 50 characters in length and can contain digits, letters, periods (.), and hyphens (-).
         * 
         * @return builder
         * 
         */
        public Builder keyWord(@Nullable Output<String> keyWord) {
            $.keyWord = keyWord;
            return this;
        }

        /**
         * @param keyWord The domain name. It must be 1 to 50 characters in length and can contain digits, letters, periods (.), and hyphens (-).
         * 
         * @return builder
         * 
         */
        public Builder keyWord(String keyWord) {
            return keyWord(Output.of(keyWord));
        }

        /**
         * @param nameRegex A regex string to filter results by Domain name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable Output<String> nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter results by Domain name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(String nameRegex) {
            return nameRegex(Output.of(nameRegex));
        }

        /**
         * @param outputFile File name where to save data source results (after running `pulumi preview`).
         * 
         * @return builder
         * 
         */
        public Builder outputFile(@Nullable Output<String> outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        /**
         * @param outputFile File name where to save data source results (after running `pulumi preview`).
         * 
         * @return builder
         * 
         */
        public Builder outputFile(String outputFile) {
            return outputFile(Output.of(outputFile));
        }

        /**
         * @param status The status of the domain name. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the domain name. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        public GetDomainsArgs build() {
            return $;
        }
    }

}

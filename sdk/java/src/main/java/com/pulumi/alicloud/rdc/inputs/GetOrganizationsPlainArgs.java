// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.rdc.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetOrganizationsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetOrganizationsPlainArgs Empty = new GetOrganizationsPlainArgs();

    /**
     * A list of Organization IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of Organization IDs.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * A regex string to filter results by Organization name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable String nameRegex;

    /**
     * @return A regex string to filter results by Organization name.
     * 
     */
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
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
     * User pk, not required, only required when the ak used by the calling interface is inconsistent with the user pk
     * 
     */
    @Import(name="realPk")
    private @Nullable String realPk;

    /**
     * @return User pk, not required, only required when the ak used by the calling interface is inconsistent with the user pk
     * 
     */
    public Optional<String> realPk() {
        return Optional.ofNullable(this.realPk);
    }

    private GetOrganizationsPlainArgs() {}

    private GetOrganizationsPlainArgs(GetOrganizationsPlainArgs $) {
        this.ids = $.ids;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
        this.realPk = $.realPk;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetOrganizationsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetOrganizationsPlainArgs $;

        public Builder() {
            $ = new GetOrganizationsPlainArgs();
        }

        public Builder(GetOrganizationsPlainArgs defaults) {
            $ = new GetOrganizationsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ids A list of Organization IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Organization IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param nameRegex A regex string to filter results by Organization name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable String nameRegex) {
            $.nameRegex = nameRegex;
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
         * @param realPk User pk, not required, only required when the ak used by the calling interface is inconsistent with the user pk
         * 
         * @return builder
         * 
         */
        public Builder realPk(@Nullable String realPk) {
            $.realPk = realPk;
            return this;
        }

        public GetOrganizationsPlainArgs build() {
            return $;
        }
    }

}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.alb.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetAscriptsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAscriptsArgs Empty = new GetAscriptsArgs();

    /**
     * Script name.
     * 
     */
    @Import(name="ascriptName")
    private @Nullable Output<String> ascriptName;

    /**
     * @return Script name.
     * 
     */
    public Optional<Output<String>> ascriptName() {
        return Optional.ofNullable(this.ascriptName);
    }

    @Import(name="enableDetails")
    private @Nullable Output<Boolean> enableDetails;

    public Optional<Output<Boolean>> enableDetails() {
        return Optional.ofNullable(this.enableDetails);
    }

    /**
     * A list of AScript IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable Output<List<String>> ids;

    /**
     * @return A list of AScript IDs.
     * 
     */
    public Optional<Output<List<String>>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * Listener ID of script attribution
     * 
     */
    @Import(name="listenerId")
    private @Nullable Output<String> listenerId;

    /**
     * @return Listener ID of script attribution
     * 
     */
    public Optional<Output<String>> listenerId() {
        return Optional.ofNullable(this.listenerId);
    }

    /**
     * A regex string to filter results by Group Metric Rule name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable Output<String> nameRegex;

    /**
     * @return A regex string to filter results by Group Metric Rule name.
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

    private GetAscriptsArgs() {}

    private GetAscriptsArgs(GetAscriptsArgs $) {
        this.ascriptName = $.ascriptName;
        this.enableDetails = $.enableDetails;
        this.ids = $.ids;
        this.listenerId = $.listenerId;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAscriptsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAscriptsArgs $;

        public Builder() {
            $ = new GetAscriptsArgs();
        }

        public Builder(GetAscriptsArgs defaults) {
            $ = new GetAscriptsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ascriptName Script name.
         * 
         * @return builder
         * 
         */
        public Builder ascriptName(@Nullable Output<String> ascriptName) {
            $.ascriptName = ascriptName;
            return this;
        }

        /**
         * @param ascriptName Script name.
         * 
         * @return builder
         * 
         */
        public Builder ascriptName(String ascriptName) {
            return ascriptName(Output.of(ascriptName));
        }

        public Builder enableDetails(@Nullable Output<Boolean> enableDetails) {
            $.enableDetails = enableDetails;
            return this;
        }

        public Builder enableDetails(Boolean enableDetails) {
            return enableDetails(Output.of(enableDetails));
        }

        /**
         * @param ids A list of AScript IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable Output<List<String>> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of AScript IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(List<String> ids) {
            return ids(Output.of(ids));
        }

        /**
         * @param ids A list of AScript IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param listenerId Listener ID of script attribution
         * 
         * @return builder
         * 
         */
        public Builder listenerId(@Nullable Output<String> listenerId) {
            $.listenerId = listenerId;
            return this;
        }

        /**
         * @param listenerId Listener ID of script attribution
         * 
         * @return builder
         * 
         */
        public Builder listenerId(String listenerId) {
            return listenerId(Output.of(listenerId));
        }

        /**
         * @param nameRegex A regex string to filter results by Group Metric Rule name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable Output<String> nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter results by Group Metric Rule name.
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

        public GetAscriptsArgs build() {
            return $;
        }
    }

}

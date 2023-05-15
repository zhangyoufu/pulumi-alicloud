// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eventbridge.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetRulesArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetRulesArgs Empty = new GetRulesArgs();

    /**
     * The name of event bus.
     * 
     */
    @Import(name="eventBusName", required=true)
    private Output<String> eventBusName;

    /**
     * @return The name of event bus.
     * 
     */
    public Output<String> eventBusName() {
        return this.eventBusName;
    }

    /**
     * A list of Rule IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable Output<List<String>> ids;

    /**
     * @return A list of Rule IDs.
     * 
     */
    public Optional<Output<List<String>>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * A regex string to filter results by Rule name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable Output<String> nameRegex;

    /**
     * @return A regex string to filter results by Rule name.
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
     * The rule name prefix.
     * 
     */
    @Import(name="ruleNamePrefix")
    private @Nullable Output<String> ruleNamePrefix;

    /**
     * @return The rule name prefix.
     * 
     */
    public Optional<Output<String>> ruleNamePrefix() {
        return Optional.ofNullable(this.ruleNamePrefix);
    }

    /**
     * Rule status, either Enable or Disable.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return Rule status, either Enable or Disable.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    private GetRulesArgs() {}

    private GetRulesArgs(GetRulesArgs $) {
        this.eventBusName = $.eventBusName;
        this.ids = $.ids;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
        this.ruleNamePrefix = $.ruleNamePrefix;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRulesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRulesArgs $;

        public Builder() {
            $ = new GetRulesArgs();
        }

        public Builder(GetRulesArgs defaults) {
            $ = new GetRulesArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param eventBusName The name of event bus.
         * 
         * @return builder
         * 
         */
        public Builder eventBusName(Output<String> eventBusName) {
            $.eventBusName = eventBusName;
            return this;
        }

        /**
         * @param eventBusName The name of event bus.
         * 
         * @return builder
         * 
         */
        public Builder eventBusName(String eventBusName) {
            return eventBusName(Output.of(eventBusName));
        }

        /**
         * @param ids A list of Rule IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable Output<List<String>> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Rule IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(List<String> ids) {
            return ids(Output.of(ids));
        }

        /**
         * @param ids A list of Rule IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param nameRegex A regex string to filter results by Rule name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable Output<String> nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter results by Rule name.
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
         * @param ruleNamePrefix The rule name prefix.
         * 
         * @return builder
         * 
         */
        public Builder ruleNamePrefix(@Nullable Output<String> ruleNamePrefix) {
            $.ruleNamePrefix = ruleNamePrefix;
            return this;
        }

        /**
         * @param ruleNamePrefix The rule name prefix.
         * 
         * @return builder
         * 
         */
        public Builder ruleNamePrefix(String ruleNamePrefix) {
            return ruleNamePrefix(Output.of(ruleNamePrefix));
        }

        /**
         * @param status Rule status, either Enable or Disable.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status Rule status, either Enable or Disable.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        public GetRulesArgs build() {
            $.eventBusName = Objects.requireNonNull($.eventBusName, "expected parameter 'eventBusName' to be non-null");
            return $;
        }
    }

}

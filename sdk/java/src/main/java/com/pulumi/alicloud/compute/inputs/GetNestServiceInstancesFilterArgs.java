// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.compute.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetNestServiceInstancesFilterArgs extends com.pulumi.resources.ResourceArgs {

    public static final GetNestServiceInstancesFilterArgs Empty = new GetNestServiceInstancesFilterArgs();

    /**
     * The name of the filter. Valid Values: `Name`, `ServiceInstanceName`, `ServiceInstanceId`, `ServiceId`, `Version`, `Status`, `DeployType`, `ServiceType`, `OperationStartTimeBefore`, `OperationStartTimeAfter`, `OperationEndTimeBefore`, `OperationEndTimeAfter`, `OperatedServiceInstanceId`, `OperationServiceInstanceId`, `EnableInstanceOps`.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the filter. Valid Values: `Name`, `ServiceInstanceName`, `ServiceInstanceId`, `ServiceId`, `Version`, `Status`, `DeployType`, `ServiceType`, `OperationStartTimeBefore`, `OperationStartTimeAfter`, `OperationEndTimeBefore`, `OperationEndTimeAfter`, `OperatedServiceInstanceId`, `OperationServiceInstanceId`, `EnableInstanceOps`.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Set of values that are accepted for the given field.
     * 
     */
    @Import(name="values")
    private @Nullable Output<List<String>> values;

    /**
     * @return Set of values that are accepted for the given field.
     * 
     */
    public Optional<Output<List<String>>> values() {
        return Optional.ofNullable(this.values);
    }

    private GetNestServiceInstancesFilterArgs() {}

    private GetNestServiceInstancesFilterArgs(GetNestServiceInstancesFilterArgs $) {
        this.name = $.name;
        this.values = $.values;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetNestServiceInstancesFilterArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetNestServiceInstancesFilterArgs $;

        public Builder() {
            $ = new GetNestServiceInstancesFilterArgs();
        }

        public Builder(GetNestServiceInstancesFilterArgs defaults) {
            $ = new GetNestServiceInstancesFilterArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The name of the filter. Valid Values: `Name`, `ServiceInstanceName`, `ServiceInstanceId`, `ServiceId`, `Version`, `Status`, `DeployType`, `ServiceType`, `OperationStartTimeBefore`, `OperationStartTimeAfter`, `OperationEndTimeBefore`, `OperationEndTimeAfter`, `OperatedServiceInstanceId`, `OperationServiceInstanceId`, `EnableInstanceOps`.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the filter. Valid Values: `Name`, `ServiceInstanceName`, `ServiceInstanceId`, `ServiceId`, `Version`, `Status`, `DeployType`, `ServiceType`, `OperationStartTimeBefore`, `OperationStartTimeAfter`, `OperationEndTimeBefore`, `OperationEndTimeAfter`, `OperatedServiceInstanceId`, `OperationServiceInstanceId`, `EnableInstanceOps`.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param values Set of values that are accepted for the given field.
         * 
         * @return builder
         * 
         */
        public Builder values(@Nullable Output<List<String>> values) {
            $.values = values;
            return this;
        }

        /**
         * @param values Set of values that are accepted for the given field.
         * 
         * @return builder
         * 
         */
        public Builder values(List<String> values) {
            return values(Output.of(values));
        }

        /**
         * @param values Set of values that are accepted for the given field.
         * 
         * @return builder
         * 
         */
        public Builder values(String... values) {
            return values(List.of(values));
        }

        public GetNestServiceInstancesFilterArgs build() {
            return $;
        }
    }

}

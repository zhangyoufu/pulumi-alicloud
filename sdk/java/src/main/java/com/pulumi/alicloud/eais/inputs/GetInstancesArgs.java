// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eais.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetInstancesArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetInstancesArgs Empty = new GetInstancesArgs();

    /**
     * A list of Instance IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable Output<List<String>> ids;

    /**
     * @return A list of Instance IDs.
     * 
     */
    public Optional<Output<List<String>>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * The type of the resource. Valid values: `eais.ei-a6.4xlarge`, `eais.ei-a6.2xlarge`, `eais.ei-a6.xlarge`, `eais.ei-a6.large`, `eais.ei-a6.medium`.
     * 
     */
    @Import(name="instanceType")
    private @Nullable Output<String> instanceType;

    /**
     * @return The type of the resource. Valid values: `eais.ei-a6.4xlarge`, `eais.ei-a6.2xlarge`, `eais.ei-a6.xlarge`, `eais.ei-a6.large`, `eais.ei-a6.medium`.
     * 
     */
    public Optional<Output<String>> instanceType() {
        return Optional.ofNullable(this.instanceType);
    }

    /**
     * A regex string to filter results by Instance name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable Output<String> nameRegex;

    /**
     * @return A regex string to filter results by Instance name.
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
     * The status of the resource. Valid values: `Attaching`, `Available`, `Detaching`, `InUse`, `Starting`, `Unavailable`.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the resource. Valid values: `Attaching`, `Available`, `Detaching`, `InUse`, `Starting`, `Unavailable`.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    private GetInstancesArgs() {}

    private GetInstancesArgs(GetInstancesArgs $) {
        this.ids = $.ids;
        this.instanceType = $.instanceType;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetInstancesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetInstancesArgs $;

        public Builder() {
            $ = new GetInstancesArgs();
        }

        public Builder(GetInstancesArgs defaults) {
            $ = new GetInstancesArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ids A list of Instance IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable Output<List<String>> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Instance IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(List<String> ids) {
            return ids(Output.of(ids));
        }

        /**
         * @param ids A list of Instance IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param instanceType The type of the resource. Valid values: `eais.ei-a6.4xlarge`, `eais.ei-a6.2xlarge`, `eais.ei-a6.xlarge`, `eais.ei-a6.large`, `eais.ei-a6.medium`.
         * 
         * @return builder
         * 
         */
        public Builder instanceType(@Nullable Output<String> instanceType) {
            $.instanceType = instanceType;
            return this;
        }

        /**
         * @param instanceType The type of the resource. Valid values: `eais.ei-a6.4xlarge`, `eais.ei-a6.2xlarge`, `eais.ei-a6.xlarge`, `eais.ei-a6.large`, `eais.ei-a6.medium`.
         * 
         * @return builder
         * 
         */
        public Builder instanceType(String instanceType) {
            return instanceType(Output.of(instanceType));
        }

        /**
         * @param nameRegex A regex string to filter results by Instance name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable Output<String> nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter results by Instance name.
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
         * @param status The status of the resource. Valid values: `Attaching`, `Available`, `Detaching`, `InUse`, `Starting`, `Unavailable`.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the resource. Valid values: `Attaching`, `Available`, `Detaching`, `InUse`, `Starting`, `Unavailable`.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        public GetInstancesArgs build() {
            return $;
        }
    }

}

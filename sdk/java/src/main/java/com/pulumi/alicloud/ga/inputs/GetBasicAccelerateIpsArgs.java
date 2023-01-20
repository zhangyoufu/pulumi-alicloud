// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ga.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetBasicAccelerateIpsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetBasicAccelerateIpsArgs Empty = new GetBasicAccelerateIpsArgs();

    /**
     * The address of the Basic Accelerate IP.
     * 
     */
    @Import(name="accelerateIpAddress")
    private @Nullable Output<String> accelerateIpAddress;

    /**
     * @return The address of the Basic Accelerate IP.
     * 
     */
    public Optional<Output<String>> accelerateIpAddress() {
        return Optional.ofNullable(this.accelerateIpAddress);
    }

    /**
     * The id of the Basic Accelerate IP.
     * 
     */
    @Import(name="accelerateIpId")
    private @Nullable Output<String> accelerateIpId;

    /**
     * @return The id of the Basic Accelerate IP.
     * 
     */
    public Optional<Output<String>> accelerateIpId() {
        return Optional.ofNullable(this.accelerateIpId);
    }

    /**
     * A list of Global Accelerator Basic Accelerate IP IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable Output<List<String>> ids;

    /**
     * @return A list of Global Accelerator Basic Accelerate IP IDs.
     * 
     */
    public Optional<Output<List<String>>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * The ID of the Basic Ip Set.
     * 
     */
    @Import(name="ipSetId", required=true)
    private Output<String> ipSetId;

    /**
     * @return The ID of the Basic Ip Set.
     * 
     */
    public Output<String> ipSetId() {
        return this.ipSetId;
    }

    @Import(name="outputFile")
    private @Nullable Output<String> outputFile;

    public Optional<Output<String>> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    /**
     * The status of the Global Accelerator Basic Accelerate IP instance. Valid Value: `active`, `binding`, `bound`, `unbinding`, `deleting`.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the Global Accelerator Basic Accelerate IP instance. Valid Value: `active`, `binding`, `bound`, `unbinding`, `deleting`.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    private GetBasicAccelerateIpsArgs() {}

    private GetBasicAccelerateIpsArgs(GetBasicAccelerateIpsArgs $) {
        this.accelerateIpAddress = $.accelerateIpAddress;
        this.accelerateIpId = $.accelerateIpId;
        this.ids = $.ids;
        this.ipSetId = $.ipSetId;
        this.outputFile = $.outputFile;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetBasicAccelerateIpsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetBasicAccelerateIpsArgs $;

        public Builder() {
            $ = new GetBasicAccelerateIpsArgs();
        }

        public Builder(GetBasicAccelerateIpsArgs defaults) {
            $ = new GetBasicAccelerateIpsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param accelerateIpAddress The address of the Basic Accelerate IP.
         * 
         * @return builder
         * 
         */
        public Builder accelerateIpAddress(@Nullable Output<String> accelerateIpAddress) {
            $.accelerateIpAddress = accelerateIpAddress;
            return this;
        }

        /**
         * @param accelerateIpAddress The address of the Basic Accelerate IP.
         * 
         * @return builder
         * 
         */
        public Builder accelerateIpAddress(String accelerateIpAddress) {
            return accelerateIpAddress(Output.of(accelerateIpAddress));
        }

        /**
         * @param accelerateIpId The id of the Basic Accelerate IP.
         * 
         * @return builder
         * 
         */
        public Builder accelerateIpId(@Nullable Output<String> accelerateIpId) {
            $.accelerateIpId = accelerateIpId;
            return this;
        }

        /**
         * @param accelerateIpId The id of the Basic Accelerate IP.
         * 
         * @return builder
         * 
         */
        public Builder accelerateIpId(String accelerateIpId) {
            return accelerateIpId(Output.of(accelerateIpId));
        }

        /**
         * @param ids A list of Global Accelerator Basic Accelerate IP IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable Output<List<String>> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Global Accelerator Basic Accelerate IP IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(List<String> ids) {
            return ids(Output.of(ids));
        }

        /**
         * @param ids A list of Global Accelerator Basic Accelerate IP IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param ipSetId The ID of the Basic Ip Set.
         * 
         * @return builder
         * 
         */
        public Builder ipSetId(Output<String> ipSetId) {
            $.ipSetId = ipSetId;
            return this;
        }

        /**
         * @param ipSetId The ID of the Basic Ip Set.
         * 
         * @return builder
         * 
         */
        public Builder ipSetId(String ipSetId) {
            return ipSetId(Output.of(ipSetId));
        }

        public Builder outputFile(@Nullable Output<String> outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        public Builder outputFile(String outputFile) {
            return outputFile(Output.of(outputFile));
        }

        /**
         * @param status The status of the Global Accelerator Basic Accelerate IP instance. Valid Value: `active`, `binding`, `bound`, `unbinding`, `deleting`.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the Global Accelerator Basic Accelerate IP instance. Valid Value: `active`, `binding`, `bound`, `unbinding`, `deleting`.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        public GetBasicAccelerateIpsArgs build() {
            $.ipSetId = Objects.requireNonNull($.ipSetId, "expected parameter 'ipSetId' to be non-null");
            return $;
        }
    }

}

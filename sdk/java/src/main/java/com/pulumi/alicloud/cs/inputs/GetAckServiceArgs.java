// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cs.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetAckServiceArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAckServiceArgs Empty = new GetAckServiceArgs();

    /**
     * Setting the value to `On` to enable the service. If has been enabled, return the result. Valid values: `On` or `Off`. Default to `Off`.
     * 
     */
    @Import(name="enable")
    private @Nullable Output<String> enable;

    /**
     * @return Setting the value to `On` to enable the service. If has been enabled, return the result. Valid values: `On` or `Off`. Default to `Off`.
     * 
     */
    public Optional<Output<String>> enable() {
        return Optional.ofNullable(this.enable);
    }

    /**
     * Types of services opened. Valid values: `propayasgo`: Container service ack Pro managed version, `edgepayasgo`: Edge container service, `gspayasgo`: Gene computing services.
     * 
     * &gt; **NOTE:** Setting `enable = &#34;On&#34;` to open the Container Service (CS) service that means you have read and agreed the [Container Service (CS) Terms of Service](https://help.aliyun.com/document_detail/157971.html). The service can not closed once it is opened.
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return Types of services opened. Valid values: `propayasgo`: Container service ack Pro managed version, `edgepayasgo`: Edge container service, `gspayasgo`: Gene computing services.
     * 
     * &gt; **NOTE:** Setting `enable = &#34;On&#34;` to open the Container Service (CS) service that means you have read and agreed the [Container Service (CS) Terms of Service](https://help.aliyun.com/document_detail/157971.html). The service can not closed once it is opened.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    private GetAckServiceArgs() {}

    private GetAckServiceArgs(GetAckServiceArgs $) {
        this.enable = $.enable;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAckServiceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAckServiceArgs $;

        public Builder() {
            $ = new GetAckServiceArgs();
        }

        public Builder(GetAckServiceArgs defaults) {
            $ = new GetAckServiceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param enable Setting the value to `On` to enable the service. If has been enabled, return the result. Valid values: `On` or `Off`. Default to `Off`.
         * 
         * @return builder
         * 
         */
        public Builder enable(@Nullable Output<String> enable) {
            $.enable = enable;
            return this;
        }

        /**
         * @param enable Setting the value to `On` to enable the service. If has been enabled, return the result. Valid values: `On` or `Off`. Default to `Off`.
         * 
         * @return builder
         * 
         */
        public Builder enable(String enable) {
            return enable(Output.of(enable));
        }

        /**
         * @param type Types of services opened. Valid values: `propayasgo`: Container service ack Pro managed version, `edgepayasgo`: Edge container service, `gspayasgo`: Gene computing services.
         * 
         * &gt; **NOTE:** Setting `enable = &#34;On&#34;` to open the Container Service (CS) service that means you have read and agreed the [Container Service (CS) Terms of Service](https://help.aliyun.com/document_detail/157971.html). The service can not closed once it is opened.
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Types of services opened. Valid values: `propayasgo`: Container service ack Pro managed version, `edgepayasgo`: Edge container service, `gspayasgo`: Gene computing services.
         * 
         * &gt; **NOTE:** Setting `enable = &#34;On&#34;` to open the Container Service (CS) service that means you have read and agreed the [Container Service (CS) Terms of Service](https://help.aliyun.com/document_detail/157971.html). The service can not closed once it is opened.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public GetAckServiceArgs build() {
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            return $;
        }
    }

}

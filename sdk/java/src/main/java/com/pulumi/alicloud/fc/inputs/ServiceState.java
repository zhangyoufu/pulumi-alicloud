// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.fc.inputs;

import com.pulumi.alicloud.fc.inputs.ServiceLogConfigArgs;
import com.pulumi.alicloud.fc.inputs.ServiceNasConfigArgs;
import com.pulumi.alicloud.fc.inputs.ServiceTracingConfigArgs;
import com.pulumi.alicloud.fc.inputs.ServiceVpcConfigArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceState extends com.pulumi.resources.ResourceArgs {

    public static final ServiceState Empty = new ServiceState();

    /**
     * The Function Compute Service description.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The Function Compute Service description.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Whether to allow the Service to access Internet. Default to &#34;true&#34;.
     * 
     */
    @Import(name="internetAccess")
    private @Nullable Output<Boolean> internetAccess;

    /**
     * @return Whether to allow the Service to access Internet. Default to &#34;true&#34;.
     * 
     */
    public Optional<Output<Boolean>> internetAccess() {
        return Optional.ofNullable(this.internetAccess);
    }

    /**
     * The date this resource was last modified.
     * 
     */
    @Import(name="lastModified")
    private @Nullable Output<String> lastModified;

    /**
     * @return The date this resource was last modified.
     * 
     */
    public Optional<Output<String>> lastModified() {
        return Optional.ofNullable(this.lastModified);
    }

    /**
     * Provide this to store your Function Compute Service logs. Fields documented below. See [Create a Service](https://www.alibabacloud.com/help/doc-detail/51924.htm). `log_config` requires the following: (**NOTE:** If both `project` and `logstore` are empty, log_config is considered to be empty or unset.)
     * 
     */
    @Import(name="logConfig")
    private @Nullable Output<ServiceLogConfigArgs> logConfig;

    /**
     * @return Provide this to store your Function Compute Service logs. Fields documented below. See [Create a Service](https://www.alibabacloud.com/help/doc-detail/51924.htm). `log_config` requires the following: (**NOTE:** If both `project` and `logstore` are empty, log_config is considered to be empty or unset.)
     * 
     */
    public Optional<Output<ServiceLogConfigArgs>> logConfig() {
        return Optional.ofNullable(this.logConfig);
    }

    /**
     * The Function Compute Service name. It is the only in one Alicloud account and is conflict with `name_prefix`.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The Function Compute Service name. It is the only in one Alicloud account and is conflict with `name_prefix`.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Setting a prefix to get a only name. It is conflict with `name`.
     * 
     */
    @Import(name="namePrefix")
    private @Nullable Output<String> namePrefix;

    /**
     * @return Setting a prefix to get a only name. It is conflict with `name`.
     * 
     */
    public Optional<Output<String>> namePrefix() {
        return Optional.ofNullable(this.namePrefix);
    }

    /**
     * Provide [NAS configuration](https://www.alibabacloud.com/help/doc-detail/87401.htm) to allow Function Compute Service to access your NAS resources. `nas_config` requires the following:
     * 
     */
    @Import(name="nasConfig")
    private @Nullable Output<ServiceNasConfigArgs> nasConfig;

    /**
     * @return Provide [NAS configuration](https://www.alibabacloud.com/help/doc-detail/87401.htm) to allow Function Compute Service to access your NAS resources. `nas_config` requires the following:
     * 
     */
    public Optional<Output<ServiceNasConfigArgs>> nasConfig() {
        return Optional.ofNullable(this.nasConfig);
    }

    /**
     * Whether to publish creation/change as new Function Compute Service Version. Defaults to `false`.
     * 
     */
    @Import(name="publish")
    private @Nullable Output<Boolean> publish;

    /**
     * @return Whether to publish creation/change as new Function Compute Service Version. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> publish() {
        return Optional.ofNullable(this.publish);
    }

    /**
     * RAM role arn attached to the Function Compute Service. This governs both who / what can invoke your Function, as well as what resources our Function has access to. See [User Permissions](https://www.alibabacloud.com/help/doc-detail/52885.htm) for more details.
     * 
     */
    @Import(name="role")
    private @Nullable Output<String> role;

    /**
     * @return RAM role arn attached to the Function Compute Service. This governs both who / what can invoke your Function, as well as what resources our Function has access to. See [User Permissions](https://www.alibabacloud.com/help/doc-detail/52885.htm) for more details.
     * 
     */
    public Optional<Output<String>> role() {
        return Optional.ofNullable(this.role);
    }

    /**
     * The Function Compute Service ID.
     * 
     */
    @Import(name="serviceId")
    private @Nullable Output<String> serviceId;

    /**
     * @return The Function Compute Service ID.
     * 
     */
    public Optional<Output<String>> serviceId() {
        return Optional.ofNullable(this.serviceId);
    }

    /**
     * Provide this to allow your Function Compute to report tracing information. Fields documented below. See [Function Compute Tracing Config](https://help.aliyun.com/document_detail/189805.html). `tracing_config` requires the following: (**NOTE:** If both `type` and `params` are empty, tracing_config is considered to be empty or unset.)
     * 
     */
    @Import(name="tracingConfig")
    private @Nullable Output<ServiceTracingConfigArgs> tracingConfig;

    /**
     * @return Provide this to allow your Function Compute to report tracing information. Fields documented below. See [Function Compute Tracing Config](https://help.aliyun.com/document_detail/189805.html). `tracing_config` requires the following: (**NOTE:** If both `type` and `params` are empty, tracing_config is considered to be empty or unset.)
     * 
     */
    public Optional<Output<ServiceTracingConfigArgs>> tracingConfig() {
        return Optional.ofNullable(this.tracingConfig);
    }

    /**
     * The latest published version of your Function Compute Service.
     * 
     */
    @Import(name="version")
    private @Nullable Output<String> version;

    /**
     * @return The latest published version of your Function Compute Service.
     * 
     */
    public Optional<Output<String>> version() {
        return Optional.ofNullable(this.version);
    }

    /**
     * Provide this to allow your Function Compute Service to access your VPC. Fields documented below. See [Function Compute Service in VPC](https://www.alibabacloud.com/help/faq-detail/72959.htm). `vpc_config` requires the following: (**NOTE:** If both `vswitch_ids` and `security_group_id` are empty, vpc_config is considered to be empty or unset.)
     * 
     */
    @Import(name="vpcConfig")
    private @Nullable Output<ServiceVpcConfigArgs> vpcConfig;

    /**
     * @return Provide this to allow your Function Compute Service to access your VPC. Fields documented below. See [Function Compute Service in VPC](https://www.alibabacloud.com/help/faq-detail/72959.htm). `vpc_config` requires the following: (**NOTE:** If both `vswitch_ids` and `security_group_id` are empty, vpc_config is considered to be empty or unset.)
     * 
     */
    public Optional<Output<ServiceVpcConfigArgs>> vpcConfig() {
        return Optional.ofNullable(this.vpcConfig);
    }

    private ServiceState() {}

    private ServiceState(ServiceState $) {
        this.description = $.description;
        this.internetAccess = $.internetAccess;
        this.lastModified = $.lastModified;
        this.logConfig = $.logConfig;
        this.name = $.name;
        this.namePrefix = $.namePrefix;
        this.nasConfig = $.nasConfig;
        this.publish = $.publish;
        this.role = $.role;
        this.serviceId = $.serviceId;
        this.tracingConfig = $.tracingConfig;
        this.version = $.version;
        this.vpcConfig = $.vpcConfig;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceState $;

        public Builder() {
            $ = new ServiceState();
        }

        public Builder(ServiceState defaults) {
            $ = new ServiceState(Objects.requireNonNull(defaults));
        }

        /**
         * @param description The Function Compute Service description.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The Function Compute Service description.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param internetAccess Whether to allow the Service to access Internet. Default to &#34;true&#34;.
         * 
         * @return builder
         * 
         */
        public Builder internetAccess(@Nullable Output<Boolean> internetAccess) {
            $.internetAccess = internetAccess;
            return this;
        }

        /**
         * @param internetAccess Whether to allow the Service to access Internet. Default to &#34;true&#34;.
         * 
         * @return builder
         * 
         */
        public Builder internetAccess(Boolean internetAccess) {
            return internetAccess(Output.of(internetAccess));
        }

        /**
         * @param lastModified The date this resource was last modified.
         * 
         * @return builder
         * 
         */
        public Builder lastModified(@Nullable Output<String> lastModified) {
            $.lastModified = lastModified;
            return this;
        }

        /**
         * @param lastModified The date this resource was last modified.
         * 
         * @return builder
         * 
         */
        public Builder lastModified(String lastModified) {
            return lastModified(Output.of(lastModified));
        }

        /**
         * @param logConfig Provide this to store your Function Compute Service logs. Fields documented below. See [Create a Service](https://www.alibabacloud.com/help/doc-detail/51924.htm). `log_config` requires the following: (**NOTE:** If both `project` and `logstore` are empty, log_config is considered to be empty or unset.)
         * 
         * @return builder
         * 
         */
        public Builder logConfig(@Nullable Output<ServiceLogConfigArgs> logConfig) {
            $.logConfig = logConfig;
            return this;
        }

        /**
         * @param logConfig Provide this to store your Function Compute Service logs. Fields documented below. See [Create a Service](https://www.alibabacloud.com/help/doc-detail/51924.htm). `log_config` requires the following: (**NOTE:** If both `project` and `logstore` are empty, log_config is considered to be empty or unset.)
         * 
         * @return builder
         * 
         */
        public Builder logConfig(ServiceLogConfigArgs logConfig) {
            return logConfig(Output.of(logConfig));
        }

        /**
         * @param name The Function Compute Service name. It is the only in one Alicloud account and is conflict with `name_prefix`.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The Function Compute Service name. It is the only in one Alicloud account and is conflict with `name_prefix`.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param namePrefix Setting a prefix to get a only name. It is conflict with `name`.
         * 
         * @return builder
         * 
         */
        public Builder namePrefix(@Nullable Output<String> namePrefix) {
            $.namePrefix = namePrefix;
            return this;
        }

        /**
         * @param namePrefix Setting a prefix to get a only name. It is conflict with `name`.
         * 
         * @return builder
         * 
         */
        public Builder namePrefix(String namePrefix) {
            return namePrefix(Output.of(namePrefix));
        }

        /**
         * @param nasConfig Provide [NAS configuration](https://www.alibabacloud.com/help/doc-detail/87401.htm) to allow Function Compute Service to access your NAS resources. `nas_config` requires the following:
         * 
         * @return builder
         * 
         */
        public Builder nasConfig(@Nullable Output<ServiceNasConfigArgs> nasConfig) {
            $.nasConfig = nasConfig;
            return this;
        }

        /**
         * @param nasConfig Provide [NAS configuration](https://www.alibabacloud.com/help/doc-detail/87401.htm) to allow Function Compute Service to access your NAS resources. `nas_config` requires the following:
         * 
         * @return builder
         * 
         */
        public Builder nasConfig(ServiceNasConfigArgs nasConfig) {
            return nasConfig(Output.of(nasConfig));
        }

        /**
         * @param publish Whether to publish creation/change as new Function Compute Service Version. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder publish(@Nullable Output<Boolean> publish) {
            $.publish = publish;
            return this;
        }

        /**
         * @param publish Whether to publish creation/change as new Function Compute Service Version. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder publish(Boolean publish) {
            return publish(Output.of(publish));
        }

        /**
         * @param role RAM role arn attached to the Function Compute Service. This governs both who / what can invoke your Function, as well as what resources our Function has access to. See [User Permissions](https://www.alibabacloud.com/help/doc-detail/52885.htm) for more details.
         * 
         * @return builder
         * 
         */
        public Builder role(@Nullable Output<String> role) {
            $.role = role;
            return this;
        }

        /**
         * @param role RAM role arn attached to the Function Compute Service. This governs both who / what can invoke your Function, as well as what resources our Function has access to. See [User Permissions](https://www.alibabacloud.com/help/doc-detail/52885.htm) for more details.
         * 
         * @return builder
         * 
         */
        public Builder role(String role) {
            return role(Output.of(role));
        }

        /**
         * @param serviceId The Function Compute Service ID.
         * 
         * @return builder
         * 
         */
        public Builder serviceId(@Nullable Output<String> serviceId) {
            $.serviceId = serviceId;
            return this;
        }

        /**
         * @param serviceId The Function Compute Service ID.
         * 
         * @return builder
         * 
         */
        public Builder serviceId(String serviceId) {
            return serviceId(Output.of(serviceId));
        }

        /**
         * @param tracingConfig Provide this to allow your Function Compute to report tracing information. Fields documented below. See [Function Compute Tracing Config](https://help.aliyun.com/document_detail/189805.html). `tracing_config` requires the following: (**NOTE:** If both `type` and `params` are empty, tracing_config is considered to be empty or unset.)
         * 
         * @return builder
         * 
         */
        public Builder tracingConfig(@Nullable Output<ServiceTracingConfigArgs> tracingConfig) {
            $.tracingConfig = tracingConfig;
            return this;
        }

        /**
         * @param tracingConfig Provide this to allow your Function Compute to report tracing information. Fields documented below. See [Function Compute Tracing Config](https://help.aliyun.com/document_detail/189805.html). `tracing_config` requires the following: (**NOTE:** If both `type` and `params` are empty, tracing_config is considered to be empty or unset.)
         * 
         * @return builder
         * 
         */
        public Builder tracingConfig(ServiceTracingConfigArgs tracingConfig) {
            return tracingConfig(Output.of(tracingConfig));
        }

        /**
         * @param version The latest published version of your Function Compute Service.
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable Output<String> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version The latest published version of your Function Compute Service.
         * 
         * @return builder
         * 
         */
        public Builder version(String version) {
            return version(Output.of(version));
        }

        /**
         * @param vpcConfig Provide this to allow your Function Compute Service to access your VPC. Fields documented below. See [Function Compute Service in VPC](https://www.alibabacloud.com/help/faq-detail/72959.htm). `vpc_config` requires the following: (**NOTE:** If both `vswitch_ids` and `security_group_id` are empty, vpc_config is considered to be empty or unset.)
         * 
         * @return builder
         * 
         */
        public Builder vpcConfig(@Nullable Output<ServiceVpcConfigArgs> vpcConfig) {
            $.vpcConfig = vpcConfig;
            return this;
        }

        /**
         * @param vpcConfig Provide this to allow your Function Compute Service to access your VPC. Fields documented below. See [Function Compute Service in VPC](https://www.alibabacloud.com/help/faq-detail/72959.htm). `vpc_config` requires the following: (**NOTE:** If both `vswitch_ids` and `security_group_id` are empty, vpc_config is considered to be empty or unset.)
         * 
         * @return builder
         * 
         */
        public Builder vpcConfig(ServiceVpcConfigArgs vpcConfig) {
            return vpcConfig(Output.of(vpcConfig));
        }

        public ServiceState build() {
            return $;
        }
    }

}

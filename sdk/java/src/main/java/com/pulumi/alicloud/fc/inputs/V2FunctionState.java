// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.fc.inputs;

import com.pulumi.alicloud.fc.inputs.V2FunctionCodeArgs;
import com.pulumi.alicloud.fc.inputs.V2FunctionCustomContainerConfigArgs;
import com.pulumi.alicloud.fc.inputs.V2FunctionCustomDnsArgs;
import com.pulumi.alicloud.fc.inputs.V2FunctionCustomHealthCheckConfigArgs;
import com.pulumi.alicloud.fc.inputs.V2FunctionCustomRuntimeConfigArgs;
import com.pulumi.alicloud.fc.inputs.V2FunctionInstanceLifecycleConfigArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Double;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class V2FunctionState extends com.pulumi.resources.ResourceArgs {

    public static final V2FunctionState Empty = new V2FunctionState();

    @Import(name="caPort")
    private @Nullable Output<Integer> caPort;

    public Optional<Output<Integer>> caPort() {
        return Optional.ofNullable(this.caPort);
    }

    @Import(name="code")
    private @Nullable Output<V2FunctionCodeArgs> code;

    public Optional<Output<V2FunctionCodeArgs>> code() {
        return Optional.ofNullable(this.code);
    }

    @Import(name="codeChecksum")
    private @Nullable Output<String> codeChecksum;

    public Optional<Output<String>> codeChecksum() {
        return Optional.ofNullable(this.codeChecksum);
    }

    @Import(name="cpu")
    private @Nullable Output<Double> cpu;

    public Optional<Output<Double>> cpu() {
        return Optional.ofNullable(this.cpu);
    }

    @Import(name="createTime")
    private @Nullable Output<String> createTime;

    public Optional<Output<String>> createTime() {
        return Optional.ofNullable(this.createTime);
    }

    @Import(name="customContainerConfig")
    private @Nullable Output<V2FunctionCustomContainerConfigArgs> customContainerConfig;

    public Optional<Output<V2FunctionCustomContainerConfigArgs>> customContainerConfig() {
        return Optional.ofNullable(this.customContainerConfig);
    }

    @Import(name="customDns")
    private @Nullable Output<V2FunctionCustomDnsArgs> customDns;

    public Optional<Output<V2FunctionCustomDnsArgs>> customDns() {
        return Optional.ofNullable(this.customDns);
    }

    @Import(name="customHealthCheckConfig")
    private @Nullable Output<V2FunctionCustomHealthCheckConfigArgs> customHealthCheckConfig;

    public Optional<Output<V2FunctionCustomHealthCheckConfigArgs>> customHealthCheckConfig() {
        return Optional.ofNullable(this.customHealthCheckConfig);
    }

    @Import(name="customRuntimeConfig")
    private @Nullable Output<V2FunctionCustomRuntimeConfigArgs> customRuntimeConfig;

    public Optional<Output<V2FunctionCustomRuntimeConfigArgs>> customRuntimeConfig() {
        return Optional.ofNullable(this.customRuntimeConfig);
    }

    @Import(name="description")
    private @Nullable Output<String> description;

    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    @Import(name="diskSize")
    private @Nullable Output<Integer> diskSize;

    public Optional<Output<Integer>> diskSize() {
        return Optional.ofNullable(this.diskSize);
    }

    @Import(name="environmentVariables")
    private @Nullable Output<Map<String,Object>> environmentVariables;

    public Optional<Output<Map<String,Object>>> environmentVariables() {
        return Optional.ofNullable(this.environmentVariables);
    }

    @Import(name="functionName")
    private @Nullable Output<String> functionName;

    public Optional<Output<String>> functionName() {
        return Optional.ofNullable(this.functionName);
    }

    @Import(name="gpuMemorySize")
    private @Nullable Output<Integer> gpuMemorySize;

    public Optional<Output<Integer>> gpuMemorySize() {
        return Optional.ofNullable(this.gpuMemorySize);
    }

    @Import(name="handler")
    private @Nullable Output<String> handler;

    public Optional<Output<String>> handler() {
        return Optional.ofNullable(this.handler);
    }

    @Import(name="initializationTimeout")
    private @Nullable Output<Integer> initializationTimeout;

    public Optional<Output<Integer>> initializationTimeout() {
        return Optional.ofNullable(this.initializationTimeout);
    }

    @Import(name="initializer")
    private @Nullable Output<String> initializer;

    public Optional<Output<String>> initializer() {
        return Optional.ofNullable(this.initializer);
    }

    @Import(name="instanceConcurrency")
    private @Nullable Output<Integer> instanceConcurrency;

    public Optional<Output<Integer>> instanceConcurrency() {
        return Optional.ofNullable(this.instanceConcurrency);
    }

    @Import(name="instanceLifecycleConfig")
    private @Nullable Output<V2FunctionInstanceLifecycleConfigArgs> instanceLifecycleConfig;

    public Optional<Output<V2FunctionInstanceLifecycleConfigArgs>> instanceLifecycleConfig() {
        return Optional.ofNullable(this.instanceLifecycleConfig);
    }

    @Import(name="instanceType")
    private @Nullable Output<String> instanceType;

    public Optional<Output<String>> instanceType() {
        return Optional.ofNullable(this.instanceType);
    }

    @Import(name="layers")
    private @Nullable Output<List<String>> layers;

    public Optional<Output<List<String>>> layers() {
        return Optional.ofNullable(this.layers);
    }

    @Import(name="memorySize")
    private @Nullable Output<Integer> memorySize;

    public Optional<Output<Integer>> memorySize() {
        return Optional.ofNullable(this.memorySize);
    }

    @Import(name="runtime")
    private @Nullable Output<String> runtime;

    public Optional<Output<String>> runtime() {
        return Optional.ofNullable(this.runtime);
    }

    @Import(name="serviceName")
    private @Nullable Output<String> serviceName;

    public Optional<Output<String>> serviceName() {
        return Optional.ofNullable(this.serviceName);
    }

    @Import(name="timeout")
    private @Nullable Output<Integer> timeout;

    public Optional<Output<Integer>> timeout() {
        return Optional.ofNullable(this.timeout);
    }

    private V2FunctionState() {}

    private V2FunctionState(V2FunctionState $) {
        this.caPort = $.caPort;
        this.code = $.code;
        this.codeChecksum = $.codeChecksum;
        this.cpu = $.cpu;
        this.createTime = $.createTime;
        this.customContainerConfig = $.customContainerConfig;
        this.customDns = $.customDns;
        this.customHealthCheckConfig = $.customHealthCheckConfig;
        this.customRuntimeConfig = $.customRuntimeConfig;
        this.description = $.description;
        this.diskSize = $.diskSize;
        this.environmentVariables = $.environmentVariables;
        this.functionName = $.functionName;
        this.gpuMemorySize = $.gpuMemorySize;
        this.handler = $.handler;
        this.initializationTimeout = $.initializationTimeout;
        this.initializer = $.initializer;
        this.instanceConcurrency = $.instanceConcurrency;
        this.instanceLifecycleConfig = $.instanceLifecycleConfig;
        this.instanceType = $.instanceType;
        this.layers = $.layers;
        this.memorySize = $.memorySize;
        this.runtime = $.runtime;
        this.serviceName = $.serviceName;
        this.timeout = $.timeout;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(V2FunctionState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private V2FunctionState $;

        public Builder() {
            $ = new V2FunctionState();
        }

        public Builder(V2FunctionState defaults) {
            $ = new V2FunctionState(Objects.requireNonNull(defaults));
        }

        public Builder caPort(@Nullable Output<Integer> caPort) {
            $.caPort = caPort;
            return this;
        }

        public Builder caPort(Integer caPort) {
            return caPort(Output.of(caPort));
        }

        public Builder code(@Nullable Output<V2FunctionCodeArgs> code) {
            $.code = code;
            return this;
        }

        public Builder code(V2FunctionCodeArgs code) {
            return code(Output.of(code));
        }

        public Builder codeChecksum(@Nullable Output<String> codeChecksum) {
            $.codeChecksum = codeChecksum;
            return this;
        }

        public Builder codeChecksum(String codeChecksum) {
            return codeChecksum(Output.of(codeChecksum));
        }

        public Builder cpu(@Nullable Output<Double> cpu) {
            $.cpu = cpu;
            return this;
        }

        public Builder cpu(Double cpu) {
            return cpu(Output.of(cpu));
        }

        public Builder createTime(@Nullable Output<String> createTime) {
            $.createTime = createTime;
            return this;
        }

        public Builder createTime(String createTime) {
            return createTime(Output.of(createTime));
        }

        public Builder customContainerConfig(@Nullable Output<V2FunctionCustomContainerConfigArgs> customContainerConfig) {
            $.customContainerConfig = customContainerConfig;
            return this;
        }

        public Builder customContainerConfig(V2FunctionCustomContainerConfigArgs customContainerConfig) {
            return customContainerConfig(Output.of(customContainerConfig));
        }

        public Builder customDns(@Nullable Output<V2FunctionCustomDnsArgs> customDns) {
            $.customDns = customDns;
            return this;
        }

        public Builder customDns(V2FunctionCustomDnsArgs customDns) {
            return customDns(Output.of(customDns));
        }

        public Builder customHealthCheckConfig(@Nullable Output<V2FunctionCustomHealthCheckConfigArgs> customHealthCheckConfig) {
            $.customHealthCheckConfig = customHealthCheckConfig;
            return this;
        }

        public Builder customHealthCheckConfig(V2FunctionCustomHealthCheckConfigArgs customHealthCheckConfig) {
            return customHealthCheckConfig(Output.of(customHealthCheckConfig));
        }

        public Builder customRuntimeConfig(@Nullable Output<V2FunctionCustomRuntimeConfigArgs> customRuntimeConfig) {
            $.customRuntimeConfig = customRuntimeConfig;
            return this;
        }

        public Builder customRuntimeConfig(V2FunctionCustomRuntimeConfigArgs customRuntimeConfig) {
            return customRuntimeConfig(Output.of(customRuntimeConfig));
        }

        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        public Builder description(String description) {
            return description(Output.of(description));
        }

        public Builder diskSize(@Nullable Output<Integer> diskSize) {
            $.diskSize = diskSize;
            return this;
        }

        public Builder diskSize(Integer diskSize) {
            return diskSize(Output.of(diskSize));
        }

        public Builder environmentVariables(@Nullable Output<Map<String,Object>> environmentVariables) {
            $.environmentVariables = environmentVariables;
            return this;
        }

        public Builder environmentVariables(Map<String,Object> environmentVariables) {
            return environmentVariables(Output.of(environmentVariables));
        }

        public Builder functionName(@Nullable Output<String> functionName) {
            $.functionName = functionName;
            return this;
        }

        public Builder functionName(String functionName) {
            return functionName(Output.of(functionName));
        }

        public Builder gpuMemorySize(@Nullable Output<Integer> gpuMemorySize) {
            $.gpuMemorySize = gpuMemorySize;
            return this;
        }

        public Builder gpuMemorySize(Integer gpuMemorySize) {
            return gpuMemorySize(Output.of(gpuMemorySize));
        }

        public Builder handler(@Nullable Output<String> handler) {
            $.handler = handler;
            return this;
        }

        public Builder handler(String handler) {
            return handler(Output.of(handler));
        }

        public Builder initializationTimeout(@Nullable Output<Integer> initializationTimeout) {
            $.initializationTimeout = initializationTimeout;
            return this;
        }

        public Builder initializationTimeout(Integer initializationTimeout) {
            return initializationTimeout(Output.of(initializationTimeout));
        }

        public Builder initializer(@Nullable Output<String> initializer) {
            $.initializer = initializer;
            return this;
        }

        public Builder initializer(String initializer) {
            return initializer(Output.of(initializer));
        }

        public Builder instanceConcurrency(@Nullable Output<Integer> instanceConcurrency) {
            $.instanceConcurrency = instanceConcurrency;
            return this;
        }

        public Builder instanceConcurrency(Integer instanceConcurrency) {
            return instanceConcurrency(Output.of(instanceConcurrency));
        }

        public Builder instanceLifecycleConfig(@Nullable Output<V2FunctionInstanceLifecycleConfigArgs> instanceLifecycleConfig) {
            $.instanceLifecycleConfig = instanceLifecycleConfig;
            return this;
        }

        public Builder instanceLifecycleConfig(V2FunctionInstanceLifecycleConfigArgs instanceLifecycleConfig) {
            return instanceLifecycleConfig(Output.of(instanceLifecycleConfig));
        }

        public Builder instanceType(@Nullable Output<String> instanceType) {
            $.instanceType = instanceType;
            return this;
        }

        public Builder instanceType(String instanceType) {
            return instanceType(Output.of(instanceType));
        }

        public Builder layers(@Nullable Output<List<String>> layers) {
            $.layers = layers;
            return this;
        }

        public Builder layers(List<String> layers) {
            return layers(Output.of(layers));
        }

        public Builder layers(String... layers) {
            return layers(List.of(layers));
        }

        public Builder memorySize(@Nullable Output<Integer> memorySize) {
            $.memorySize = memorySize;
            return this;
        }

        public Builder memorySize(Integer memorySize) {
            return memorySize(Output.of(memorySize));
        }

        public Builder runtime(@Nullable Output<String> runtime) {
            $.runtime = runtime;
            return this;
        }

        public Builder runtime(String runtime) {
            return runtime(Output.of(runtime));
        }

        public Builder serviceName(@Nullable Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        public Builder timeout(@Nullable Output<Integer> timeout) {
            $.timeout = timeout;
            return this;
        }

        public Builder timeout(Integer timeout) {
            return timeout(Output.of(timeout));
        }

        public V2FunctionState build() {
            return $;
        }
    }

}

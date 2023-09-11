// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eci.outputs;

import com.pulumi.alicloud.eci.outputs.ContainerGroupInitContainerEnvironmentVar;
import com.pulumi.alicloud.eci.outputs.ContainerGroupInitContainerPort;
import com.pulumi.alicloud.eci.outputs.ContainerGroupInitContainerVolumeMount;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Double;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ContainerGroupInitContainer {
    /**
     * @return The arguments passed to the commands.
     * 
     */
    private @Nullable List<String> args;
    /**
     * @return The commands run by the init container.
     * 
     */
    private @Nullable List<String> commands;
    /**
     * @return The amount of CPU resources allocated to the container. Default value: `0`.
     * 
     */
    private @Nullable Double cpu;
    /**
     * @return The structure of environmentVars. See `environment_vars` below.
     * 
     */
    private @Nullable List<ContainerGroupInitContainerEnvironmentVar> environmentVars;
    /**
     * @return The number GPUs. Default value: `0`.
     * 
     */
    private @Nullable Integer gpu;
    /**
     * @return The image of the container.
     * 
     */
    private @Nullable String image;
    /**
     * @return The restart policy of the image. Default value: `IfNotPresent`. Valid values: `Always`, `IfNotPresent`, `Never`.
     * 
     */
    private @Nullable String imagePullPolicy;
    /**
     * @return The amount of memory resources allocated to the container. Default value: `0`.
     * 
     */
    private @Nullable Double memory;
    /**
     * @return The name of the mounted volume.
     * 
     */
    private @Nullable String name;
    /**
     * @return The structure of port. See `ports` below.
     * 
     */
    private @Nullable List<ContainerGroupInitContainerPort> ports;
    /**
     * @return Indicates whether the container passed the readiness probe.
     * 
     */
    private @Nullable Boolean ready;
    /**
     * @return The number of times that the container restarted.
     * 
     */
    private @Nullable Integer restartCount;
    /**
     * @return The structure of volumeMounts. See `volume_mounts` below.
     * 
     */
    private @Nullable List<ContainerGroupInitContainerVolumeMount> volumeMounts;
    /**
     * @return The working directory of the container.
     * 
     */
    private @Nullable String workingDir;

    private ContainerGroupInitContainer() {}
    /**
     * @return The arguments passed to the commands.
     * 
     */
    public List<String> args() {
        return this.args == null ? List.of() : this.args;
    }
    /**
     * @return The commands run by the init container.
     * 
     */
    public List<String> commands() {
        return this.commands == null ? List.of() : this.commands;
    }
    /**
     * @return The amount of CPU resources allocated to the container. Default value: `0`.
     * 
     */
    public Optional<Double> cpu() {
        return Optional.ofNullable(this.cpu);
    }
    /**
     * @return The structure of environmentVars. See `environment_vars` below.
     * 
     */
    public List<ContainerGroupInitContainerEnvironmentVar> environmentVars() {
        return this.environmentVars == null ? List.of() : this.environmentVars;
    }
    /**
     * @return The number GPUs. Default value: `0`.
     * 
     */
    public Optional<Integer> gpu() {
        return Optional.ofNullable(this.gpu);
    }
    /**
     * @return The image of the container.
     * 
     */
    public Optional<String> image() {
        return Optional.ofNullable(this.image);
    }
    /**
     * @return The restart policy of the image. Default value: `IfNotPresent`. Valid values: `Always`, `IfNotPresent`, `Never`.
     * 
     */
    public Optional<String> imagePullPolicy() {
        return Optional.ofNullable(this.imagePullPolicy);
    }
    /**
     * @return The amount of memory resources allocated to the container. Default value: `0`.
     * 
     */
    public Optional<Double> memory() {
        return Optional.ofNullable(this.memory);
    }
    /**
     * @return The name of the mounted volume.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return The structure of port. See `ports` below.
     * 
     */
    public List<ContainerGroupInitContainerPort> ports() {
        return this.ports == null ? List.of() : this.ports;
    }
    /**
     * @return Indicates whether the container passed the readiness probe.
     * 
     */
    public Optional<Boolean> ready() {
        return Optional.ofNullable(this.ready);
    }
    /**
     * @return The number of times that the container restarted.
     * 
     */
    public Optional<Integer> restartCount() {
        return Optional.ofNullable(this.restartCount);
    }
    /**
     * @return The structure of volumeMounts. See `volume_mounts` below.
     * 
     */
    public List<ContainerGroupInitContainerVolumeMount> volumeMounts() {
        return this.volumeMounts == null ? List.of() : this.volumeMounts;
    }
    /**
     * @return The working directory of the container.
     * 
     */
    public Optional<String> workingDir() {
        return Optional.ofNullable(this.workingDir);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ContainerGroupInitContainer defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> args;
        private @Nullable List<String> commands;
        private @Nullable Double cpu;
        private @Nullable List<ContainerGroupInitContainerEnvironmentVar> environmentVars;
        private @Nullable Integer gpu;
        private @Nullable String image;
        private @Nullable String imagePullPolicy;
        private @Nullable Double memory;
        private @Nullable String name;
        private @Nullable List<ContainerGroupInitContainerPort> ports;
        private @Nullable Boolean ready;
        private @Nullable Integer restartCount;
        private @Nullable List<ContainerGroupInitContainerVolumeMount> volumeMounts;
        private @Nullable String workingDir;
        public Builder() {}
        public Builder(ContainerGroupInitContainer defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.args = defaults.args;
    	      this.commands = defaults.commands;
    	      this.cpu = defaults.cpu;
    	      this.environmentVars = defaults.environmentVars;
    	      this.gpu = defaults.gpu;
    	      this.image = defaults.image;
    	      this.imagePullPolicy = defaults.imagePullPolicy;
    	      this.memory = defaults.memory;
    	      this.name = defaults.name;
    	      this.ports = defaults.ports;
    	      this.ready = defaults.ready;
    	      this.restartCount = defaults.restartCount;
    	      this.volumeMounts = defaults.volumeMounts;
    	      this.workingDir = defaults.workingDir;
        }

        @CustomType.Setter
        public Builder args(@Nullable List<String> args) {
            this.args = args;
            return this;
        }
        public Builder args(String... args) {
            return args(List.of(args));
        }
        @CustomType.Setter
        public Builder commands(@Nullable List<String> commands) {
            this.commands = commands;
            return this;
        }
        public Builder commands(String... commands) {
            return commands(List.of(commands));
        }
        @CustomType.Setter
        public Builder cpu(@Nullable Double cpu) {
            this.cpu = cpu;
            return this;
        }
        @CustomType.Setter
        public Builder environmentVars(@Nullable List<ContainerGroupInitContainerEnvironmentVar> environmentVars) {
            this.environmentVars = environmentVars;
            return this;
        }
        public Builder environmentVars(ContainerGroupInitContainerEnvironmentVar... environmentVars) {
            return environmentVars(List.of(environmentVars));
        }
        @CustomType.Setter
        public Builder gpu(@Nullable Integer gpu) {
            this.gpu = gpu;
            return this;
        }
        @CustomType.Setter
        public Builder image(@Nullable String image) {
            this.image = image;
            return this;
        }
        @CustomType.Setter
        public Builder imagePullPolicy(@Nullable String imagePullPolicy) {
            this.imagePullPolicy = imagePullPolicy;
            return this;
        }
        @CustomType.Setter
        public Builder memory(@Nullable Double memory) {
            this.memory = memory;
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder ports(@Nullable List<ContainerGroupInitContainerPort> ports) {
            this.ports = ports;
            return this;
        }
        public Builder ports(ContainerGroupInitContainerPort... ports) {
            return ports(List.of(ports));
        }
        @CustomType.Setter
        public Builder ready(@Nullable Boolean ready) {
            this.ready = ready;
            return this;
        }
        @CustomType.Setter
        public Builder restartCount(@Nullable Integer restartCount) {
            this.restartCount = restartCount;
            return this;
        }
        @CustomType.Setter
        public Builder volumeMounts(@Nullable List<ContainerGroupInitContainerVolumeMount> volumeMounts) {
            this.volumeMounts = volumeMounts;
            return this;
        }
        public Builder volumeMounts(ContainerGroupInitContainerVolumeMount... volumeMounts) {
            return volumeMounts(List.of(volumeMounts));
        }
        @CustomType.Setter
        public Builder workingDir(@Nullable String workingDir) {
            this.workingDir = workingDir;
            return this;
        }
        public ContainerGroupInitContainer build() {
            final var o = new ContainerGroupInitContainer();
            o.args = args;
            o.commands = commands;
            o.cpu = cpu;
            o.environmentVars = environmentVars;
            o.gpu = gpu;
            o.image = image;
            o.imagePullPolicy = imagePullPolicy;
            o.memory = memory;
            o.name = name;
            o.ports = ports;
            o.ready = ready;
            o.restartCount = restartCount;
            o.volumeMounts = volumeMounts;
            o.workingDir = workingDir;
            return o;
        }
    }
}

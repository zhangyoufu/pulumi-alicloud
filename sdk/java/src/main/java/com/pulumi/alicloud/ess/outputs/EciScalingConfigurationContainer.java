// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ess.outputs;

import com.pulumi.alicloud.ess.outputs.EciScalingConfigurationContainerEnvironmentVar;
import com.pulumi.alicloud.ess.outputs.EciScalingConfigurationContainerPort;
import com.pulumi.alicloud.ess.outputs.EciScalingConfigurationContainerVolumeMount;
import com.pulumi.core.annotations.CustomType;
import java.lang.Double;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class EciScalingConfigurationContainer {
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
     * @return The amount of CPU resources allocated to the container.
     * 
     */
    private @Nullable Double cpu;
    /**
     * @return The structure of environmentVars.
     * See `environment_vars` below for details.
     * 
     */
    private @Nullable List<EciScalingConfigurationContainerEnvironmentVar> environmentVars;
    /**
     * @return The number GPUs.
     * 
     */
    private @Nullable Integer gpu;
    /**
     * @return The image of the container.
     * 
     */
    private @Nullable String image;
    /**
     * @return The restart policy of the image.
     * 
     */
    private @Nullable String imagePullPolicy;
    /**
     * @return Commands that you want to run in containers when you use the CLI to perform liveness probes.
     * 
     */
    private @Nullable List<String> livenessProbeExecCommands;
    /**
     * @return The minimum number of consecutive failures for the liveness probe to be considered failed after having been successful. Default value: 3.
     * 
     */
    private @Nullable Integer livenessProbeFailureThreshold;
    /**
     * @return The path to which HTTP GET requests are sent when you use HTTP requests to perform liveness probes.
     * 
     */
    private @Nullable String livenessProbeHttpGetPath;
    /**
     * @return The port to which HTTP GET requests are sent when you use HTTP requests to perform liveness probes.
     * 
     */
    private @Nullable Integer livenessProbeHttpGetPort;
    /**
     * @return The protocol type of HTTP GET requests when you use HTTP requests for liveness probes.Valid values:HTTP and HTTPS.
     * 
     */
    private @Nullable String livenessProbeHttpGetScheme;
    /**
     * @return The number of seconds after container has started before liveness probes are initiated.
     * 
     */
    private @Nullable Integer livenessProbeInitialDelaySeconds;
    /**
     * @return The interval at which the liveness probe is performed. Unit: seconds. Default value: 10. Minimum value: 1.
     * 
     */
    private @Nullable Integer livenessProbePeriodSeconds;
    /**
     * @return The minimum number of consecutive successes for the liveness probe to be considered successful after having failed. Default value: 1. Set the value to 1.
     * 
     */
    private @Nullable Integer livenessProbeSuccessThreshold;
    /**
     * @return The port detected by TCP sockets when you use TCP sockets to perform liveness probes.
     * 
     */
    private @Nullable Integer livenessProbeTcpSocketPort;
    /**
     * @return The timeout period for the liveness probe. Unit: seconds. Default value: 1. Minimum value: 1.
     * 
     */
    private @Nullable Integer livenessProbeTimeoutSeconds;
    /**
     * @return The amount of memory resources allocated to the container.
     * 
     */
    private @Nullable Double memory;
    /**
     * @return The name of the mounted volume.
     * 
     */
    private @Nullable String name;
    /**
     * @return The structure of port. See `ports` below for details.
     * 
     */
    private @Nullable List<EciScalingConfigurationContainerPort> ports;
    /**
     * @return Commands that you want to run in containers when you use the CLI to perform readiness probes.
     * 
     */
    private @Nullable List<String> readinessProbeExecCommands;
    /**
     * @return The minimum number of consecutive failures for the readiness probe to be considered failed after having been successful. Default value: 3.
     * 
     */
    private @Nullable Integer readinessProbeFailureThreshold;
    /**
     * @return The path to which HTTP GET requests are sent when you use HTTP requests to perform readiness probes.
     * 
     */
    private @Nullable String readinessProbeHttpGetPath;
    /**
     * @return The port to which HTTP GET requests are sent when you use HTTP requests to perform readiness probes.
     * 
     */
    private @Nullable Integer readinessProbeHttpGetPort;
    /**
     * @return The protocol type of HTTP GET requests when you use HTTP requests for readiness probes. Valid values: HTTP and HTTPS.
     * 
     */
    private @Nullable String readinessProbeHttpGetScheme;
    /**
     * @return The number of seconds after container N has started before readiness probes are initiated.
     * 
     */
    private @Nullable Integer readinessProbeInitialDelaySeconds;
    /**
     * @return The interval at which the readiness probe is performed. Unit: seconds. Default value: 10. Minimum value: 1.
     * 
     */
    private @Nullable Integer readinessProbePeriodSeconds;
    /**
     * @return The minimum number of consecutive successes for the readiness probe to be considered successful after having failed. Default value: 1. Set the value to 1.
     * 
     */
    private @Nullable Integer readinessProbeSuccessThreshold;
    /**
     * @return The port detected by Transmission Control Protocol (TCP) sockets when you use TCP sockets to perform readiness probes.
     * 
     */
    private @Nullable Integer readinessProbeTcpSocketPort;
    /**
     * @return The timeout period for the readiness probe. Unit: seconds. Default value: 1. Minimum value: 1.
     * 
     */
    private @Nullable Integer readinessProbeTimeoutSeconds;
    /**
     * @return The structure of volumeMounts.
     * See `volume_mounts` below for details.
     * 
     */
    private @Nullable List<EciScalingConfigurationContainerVolumeMount> volumeMounts;
    /**
     * @return The working directory of the container.
     * 
     */
    private @Nullable String workingDir;

    private EciScalingConfigurationContainer() {}
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
     * @return The amount of CPU resources allocated to the container.
     * 
     */
    public Optional<Double> cpu() {
        return Optional.ofNullable(this.cpu);
    }
    /**
     * @return The structure of environmentVars.
     * See `environment_vars` below for details.
     * 
     */
    public List<EciScalingConfigurationContainerEnvironmentVar> environmentVars() {
        return this.environmentVars == null ? List.of() : this.environmentVars;
    }
    /**
     * @return The number GPUs.
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
     * @return The restart policy of the image.
     * 
     */
    public Optional<String> imagePullPolicy() {
        return Optional.ofNullable(this.imagePullPolicy);
    }
    /**
     * @return Commands that you want to run in containers when you use the CLI to perform liveness probes.
     * 
     */
    public List<String> livenessProbeExecCommands() {
        return this.livenessProbeExecCommands == null ? List.of() : this.livenessProbeExecCommands;
    }
    /**
     * @return The minimum number of consecutive failures for the liveness probe to be considered failed after having been successful. Default value: 3.
     * 
     */
    public Optional<Integer> livenessProbeFailureThreshold() {
        return Optional.ofNullable(this.livenessProbeFailureThreshold);
    }
    /**
     * @return The path to which HTTP GET requests are sent when you use HTTP requests to perform liveness probes.
     * 
     */
    public Optional<String> livenessProbeHttpGetPath() {
        return Optional.ofNullable(this.livenessProbeHttpGetPath);
    }
    /**
     * @return The port to which HTTP GET requests are sent when you use HTTP requests to perform liveness probes.
     * 
     */
    public Optional<Integer> livenessProbeHttpGetPort() {
        return Optional.ofNullable(this.livenessProbeHttpGetPort);
    }
    /**
     * @return The protocol type of HTTP GET requests when you use HTTP requests for liveness probes.Valid values:HTTP and HTTPS.
     * 
     */
    public Optional<String> livenessProbeHttpGetScheme() {
        return Optional.ofNullable(this.livenessProbeHttpGetScheme);
    }
    /**
     * @return The number of seconds after container has started before liveness probes are initiated.
     * 
     */
    public Optional<Integer> livenessProbeInitialDelaySeconds() {
        return Optional.ofNullable(this.livenessProbeInitialDelaySeconds);
    }
    /**
     * @return The interval at which the liveness probe is performed. Unit: seconds. Default value: 10. Minimum value: 1.
     * 
     */
    public Optional<Integer> livenessProbePeriodSeconds() {
        return Optional.ofNullable(this.livenessProbePeriodSeconds);
    }
    /**
     * @return The minimum number of consecutive successes for the liveness probe to be considered successful after having failed. Default value: 1. Set the value to 1.
     * 
     */
    public Optional<Integer> livenessProbeSuccessThreshold() {
        return Optional.ofNullable(this.livenessProbeSuccessThreshold);
    }
    /**
     * @return The port detected by TCP sockets when you use TCP sockets to perform liveness probes.
     * 
     */
    public Optional<Integer> livenessProbeTcpSocketPort() {
        return Optional.ofNullable(this.livenessProbeTcpSocketPort);
    }
    /**
     * @return The timeout period for the liveness probe. Unit: seconds. Default value: 1. Minimum value: 1.
     * 
     */
    public Optional<Integer> livenessProbeTimeoutSeconds() {
        return Optional.ofNullable(this.livenessProbeTimeoutSeconds);
    }
    /**
     * @return The amount of memory resources allocated to the container.
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
     * @return The structure of port. See `ports` below for details.
     * 
     */
    public List<EciScalingConfigurationContainerPort> ports() {
        return this.ports == null ? List.of() : this.ports;
    }
    /**
     * @return Commands that you want to run in containers when you use the CLI to perform readiness probes.
     * 
     */
    public List<String> readinessProbeExecCommands() {
        return this.readinessProbeExecCommands == null ? List.of() : this.readinessProbeExecCommands;
    }
    /**
     * @return The minimum number of consecutive failures for the readiness probe to be considered failed after having been successful. Default value: 3.
     * 
     */
    public Optional<Integer> readinessProbeFailureThreshold() {
        return Optional.ofNullable(this.readinessProbeFailureThreshold);
    }
    /**
     * @return The path to which HTTP GET requests are sent when you use HTTP requests to perform readiness probes.
     * 
     */
    public Optional<String> readinessProbeHttpGetPath() {
        return Optional.ofNullable(this.readinessProbeHttpGetPath);
    }
    /**
     * @return The port to which HTTP GET requests are sent when you use HTTP requests to perform readiness probes.
     * 
     */
    public Optional<Integer> readinessProbeHttpGetPort() {
        return Optional.ofNullable(this.readinessProbeHttpGetPort);
    }
    /**
     * @return The protocol type of HTTP GET requests when you use HTTP requests for readiness probes. Valid values: HTTP and HTTPS.
     * 
     */
    public Optional<String> readinessProbeHttpGetScheme() {
        return Optional.ofNullable(this.readinessProbeHttpGetScheme);
    }
    /**
     * @return The number of seconds after container N has started before readiness probes are initiated.
     * 
     */
    public Optional<Integer> readinessProbeInitialDelaySeconds() {
        return Optional.ofNullable(this.readinessProbeInitialDelaySeconds);
    }
    /**
     * @return The interval at which the readiness probe is performed. Unit: seconds. Default value: 10. Minimum value: 1.
     * 
     */
    public Optional<Integer> readinessProbePeriodSeconds() {
        return Optional.ofNullable(this.readinessProbePeriodSeconds);
    }
    /**
     * @return The minimum number of consecutive successes for the readiness probe to be considered successful after having failed. Default value: 1. Set the value to 1.
     * 
     */
    public Optional<Integer> readinessProbeSuccessThreshold() {
        return Optional.ofNullable(this.readinessProbeSuccessThreshold);
    }
    /**
     * @return The port detected by Transmission Control Protocol (TCP) sockets when you use TCP sockets to perform readiness probes.
     * 
     */
    public Optional<Integer> readinessProbeTcpSocketPort() {
        return Optional.ofNullable(this.readinessProbeTcpSocketPort);
    }
    /**
     * @return The timeout period for the readiness probe. Unit: seconds. Default value: 1. Minimum value: 1.
     * 
     */
    public Optional<Integer> readinessProbeTimeoutSeconds() {
        return Optional.ofNullable(this.readinessProbeTimeoutSeconds);
    }
    /**
     * @return The structure of volumeMounts.
     * See `volume_mounts` below for details.
     * 
     */
    public List<EciScalingConfigurationContainerVolumeMount> volumeMounts() {
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

    public static Builder builder(EciScalingConfigurationContainer defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> args;
        private @Nullable List<String> commands;
        private @Nullable Double cpu;
        private @Nullable List<EciScalingConfigurationContainerEnvironmentVar> environmentVars;
        private @Nullable Integer gpu;
        private @Nullable String image;
        private @Nullable String imagePullPolicy;
        private @Nullable List<String> livenessProbeExecCommands;
        private @Nullable Integer livenessProbeFailureThreshold;
        private @Nullable String livenessProbeHttpGetPath;
        private @Nullable Integer livenessProbeHttpGetPort;
        private @Nullable String livenessProbeHttpGetScheme;
        private @Nullable Integer livenessProbeInitialDelaySeconds;
        private @Nullable Integer livenessProbePeriodSeconds;
        private @Nullable Integer livenessProbeSuccessThreshold;
        private @Nullable Integer livenessProbeTcpSocketPort;
        private @Nullable Integer livenessProbeTimeoutSeconds;
        private @Nullable Double memory;
        private @Nullable String name;
        private @Nullable List<EciScalingConfigurationContainerPort> ports;
        private @Nullable List<String> readinessProbeExecCommands;
        private @Nullable Integer readinessProbeFailureThreshold;
        private @Nullable String readinessProbeHttpGetPath;
        private @Nullable Integer readinessProbeHttpGetPort;
        private @Nullable String readinessProbeHttpGetScheme;
        private @Nullable Integer readinessProbeInitialDelaySeconds;
        private @Nullable Integer readinessProbePeriodSeconds;
        private @Nullable Integer readinessProbeSuccessThreshold;
        private @Nullable Integer readinessProbeTcpSocketPort;
        private @Nullable Integer readinessProbeTimeoutSeconds;
        private @Nullable List<EciScalingConfigurationContainerVolumeMount> volumeMounts;
        private @Nullable String workingDir;
        public Builder() {}
        public Builder(EciScalingConfigurationContainer defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.args = defaults.args;
    	      this.commands = defaults.commands;
    	      this.cpu = defaults.cpu;
    	      this.environmentVars = defaults.environmentVars;
    	      this.gpu = defaults.gpu;
    	      this.image = defaults.image;
    	      this.imagePullPolicy = defaults.imagePullPolicy;
    	      this.livenessProbeExecCommands = defaults.livenessProbeExecCommands;
    	      this.livenessProbeFailureThreshold = defaults.livenessProbeFailureThreshold;
    	      this.livenessProbeHttpGetPath = defaults.livenessProbeHttpGetPath;
    	      this.livenessProbeHttpGetPort = defaults.livenessProbeHttpGetPort;
    	      this.livenessProbeHttpGetScheme = defaults.livenessProbeHttpGetScheme;
    	      this.livenessProbeInitialDelaySeconds = defaults.livenessProbeInitialDelaySeconds;
    	      this.livenessProbePeriodSeconds = defaults.livenessProbePeriodSeconds;
    	      this.livenessProbeSuccessThreshold = defaults.livenessProbeSuccessThreshold;
    	      this.livenessProbeTcpSocketPort = defaults.livenessProbeTcpSocketPort;
    	      this.livenessProbeTimeoutSeconds = defaults.livenessProbeTimeoutSeconds;
    	      this.memory = defaults.memory;
    	      this.name = defaults.name;
    	      this.ports = defaults.ports;
    	      this.readinessProbeExecCommands = defaults.readinessProbeExecCommands;
    	      this.readinessProbeFailureThreshold = defaults.readinessProbeFailureThreshold;
    	      this.readinessProbeHttpGetPath = defaults.readinessProbeHttpGetPath;
    	      this.readinessProbeHttpGetPort = defaults.readinessProbeHttpGetPort;
    	      this.readinessProbeHttpGetScheme = defaults.readinessProbeHttpGetScheme;
    	      this.readinessProbeInitialDelaySeconds = defaults.readinessProbeInitialDelaySeconds;
    	      this.readinessProbePeriodSeconds = defaults.readinessProbePeriodSeconds;
    	      this.readinessProbeSuccessThreshold = defaults.readinessProbeSuccessThreshold;
    	      this.readinessProbeTcpSocketPort = defaults.readinessProbeTcpSocketPort;
    	      this.readinessProbeTimeoutSeconds = defaults.readinessProbeTimeoutSeconds;
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
        public Builder environmentVars(@Nullable List<EciScalingConfigurationContainerEnvironmentVar> environmentVars) {
            this.environmentVars = environmentVars;
            return this;
        }
        public Builder environmentVars(EciScalingConfigurationContainerEnvironmentVar... environmentVars) {
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
        public Builder livenessProbeExecCommands(@Nullable List<String> livenessProbeExecCommands) {
            this.livenessProbeExecCommands = livenessProbeExecCommands;
            return this;
        }
        public Builder livenessProbeExecCommands(String... livenessProbeExecCommands) {
            return livenessProbeExecCommands(List.of(livenessProbeExecCommands));
        }
        @CustomType.Setter
        public Builder livenessProbeFailureThreshold(@Nullable Integer livenessProbeFailureThreshold) {
            this.livenessProbeFailureThreshold = livenessProbeFailureThreshold;
            return this;
        }
        @CustomType.Setter
        public Builder livenessProbeHttpGetPath(@Nullable String livenessProbeHttpGetPath) {
            this.livenessProbeHttpGetPath = livenessProbeHttpGetPath;
            return this;
        }
        @CustomType.Setter
        public Builder livenessProbeHttpGetPort(@Nullable Integer livenessProbeHttpGetPort) {
            this.livenessProbeHttpGetPort = livenessProbeHttpGetPort;
            return this;
        }
        @CustomType.Setter
        public Builder livenessProbeHttpGetScheme(@Nullable String livenessProbeHttpGetScheme) {
            this.livenessProbeHttpGetScheme = livenessProbeHttpGetScheme;
            return this;
        }
        @CustomType.Setter
        public Builder livenessProbeInitialDelaySeconds(@Nullable Integer livenessProbeInitialDelaySeconds) {
            this.livenessProbeInitialDelaySeconds = livenessProbeInitialDelaySeconds;
            return this;
        }
        @CustomType.Setter
        public Builder livenessProbePeriodSeconds(@Nullable Integer livenessProbePeriodSeconds) {
            this.livenessProbePeriodSeconds = livenessProbePeriodSeconds;
            return this;
        }
        @CustomType.Setter
        public Builder livenessProbeSuccessThreshold(@Nullable Integer livenessProbeSuccessThreshold) {
            this.livenessProbeSuccessThreshold = livenessProbeSuccessThreshold;
            return this;
        }
        @CustomType.Setter
        public Builder livenessProbeTcpSocketPort(@Nullable Integer livenessProbeTcpSocketPort) {
            this.livenessProbeTcpSocketPort = livenessProbeTcpSocketPort;
            return this;
        }
        @CustomType.Setter
        public Builder livenessProbeTimeoutSeconds(@Nullable Integer livenessProbeTimeoutSeconds) {
            this.livenessProbeTimeoutSeconds = livenessProbeTimeoutSeconds;
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
        public Builder ports(@Nullable List<EciScalingConfigurationContainerPort> ports) {
            this.ports = ports;
            return this;
        }
        public Builder ports(EciScalingConfigurationContainerPort... ports) {
            return ports(List.of(ports));
        }
        @CustomType.Setter
        public Builder readinessProbeExecCommands(@Nullable List<String> readinessProbeExecCommands) {
            this.readinessProbeExecCommands = readinessProbeExecCommands;
            return this;
        }
        public Builder readinessProbeExecCommands(String... readinessProbeExecCommands) {
            return readinessProbeExecCommands(List.of(readinessProbeExecCommands));
        }
        @CustomType.Setter
        public Builder readinessProbeFailureThreshold(@Nullable Integer readinessProbeFailureThreshold) {
            this.readinessProbeFailureThreshold = readinessProbeFailureThreshold;
            return this;
        }
        @CustomType.Setter
        public Builder readinessProbeHttpGetPath(@Nullable String readinessProbeHttpGetPath) {
            this.readinessProbeHttpGetPath = readinessProbeHttpGetPath;
            return this;
        }
        @CustomType.Setter
        public Builder readinessProbeHttpGetPort(@Nullable Integer readinessProbeHttpGetPort) {
            this.readinessProbeHttpGetPort = readinessProbeHttpGetPort;
            return this;
        }
        @CustomType.Setter
        public Builder readinessProbeHttpGetScheme(@Nullable String readinessProbeHttpGetScheme) {
            this.readinessProbeHttpGetScheme = readinessProbeHttpGetScheme;
            return this;
        }
        @CustomType.Setter
        public Builder readinessProbeInitialDelaySeconds(@Nullable Integer readinessProbeInitialDelaySeconds) {
            this.readinessProbeInitialDelaySeconds = readinessProbeInitialDelaySeconds;
            return this;
        }
        @CustomType.Setter
        public Builder readinessProbePeriodSeconds(@Nullable Integer readinessProbePeriodSeconds) {
            this.readinessProbePeriodSeconds = readinessProbePeriodSeconds;
            return this;
        }
        @CustomType.Setter
        public Builder readinessProbeSuccessThreshold(@Nullable Integer readinessProbeSuccessThreshold) {
            this.readinessProbeSuccessThreshold = readinessProbeSuccessThreshold;
            return this;
        }
        @CustomType.Setter
        public Builder readinessProbeTcpSocketPort(@Nullable Integer readinessProbeTcpSocketPort) {
            this.readinessProbeTcpSocketPort = readinessProbeTcpSocketPort;
            return this;
        }
        @CustomType.Setter
        public Builder readinessProbeTimeoutSeconds(@Nullable Integer readinessProbeTimeoutSeconds) {
            this.readinessProbeTimeoutSeconds = readinessProbeTimeoutSeconds;
            return this;
        }
        @CustomType.Setter
        public Builder volumeMounts(@Nullable List<EciScalingConfigurationContainerVolumeMount> volumeMounts) {
            this.volumeMounts = volumeMounts;
            return this;
        }
        public Builder volumeMounts(EciScalingConfigurationContainerVolumeMount... volumeMounts) {
            return volumeMounts(List.of(volumeMounts));
        }
        @CustomType.Setter
        public Builder workingDir(@Nullable String workingDir) {
            this.workingDir = workingDir;
            return this;
        }
        public EciScalingConfigurationContainer build() {
            final var o = new EciScalingConfigurationContainer();
            o.args = args;
            o.commands = commands;
            o.cpu = cpu;
            o.environmentVars = environmentVars;
            o.gpu = gpu;
            o.image = image;
            o.imagePullPolicy = imagePullPolicy;
            o.livenessProbeExecCommands = livenessProbeExecCommands;
            o.livenessProbeFailureThreshold = livenessProbeFailureThreshold;
            o.livenessProbeHttpGetPath = livenessProbeHttpGetPath;
            o.livenessProbeHttpGetPort = livenessProbeHttpGetPort;
            o.livenessProbeHttpGetScheme = livenessProbeHttpGetScheme;
            o.livenessProbeInitialDelaySeconds = livenessProbeInitialDelaySeconds;
            o.livenessProbePeriodSeconds = livenessProbePeriodSeconds;
            o.livenessProbeSuccessThreshold = livenessProbeSuccessThreshold;
            o.livenessProbeTcpSocketPort = livenessProbeTcpSocketPort;
            o.livenessProbeTimeoutSeconds = livenessProbeTimeoutSeconds;
            o.memory = memory;
            o.name = name;
            o.ports = ports;
            o.readinessProbeExecCommands = readinessProbeExecCommands;
            o.readinessProbeFailureThreshold = readinessProbeFailureThreshold;
            o.readinessProbeHttpGetPath = readinessProbeHttpGetPath;
            o.readinessProbeHttpGetPort = readinessProbeHttpGetPort;
            o.readinessProbeHttpGetScheme = readinessProbeHttpGetScheme;
            o.readinessProbeInitialDelaySeconds = readinessProbeInitialDelaySeconds;
            o.readinessProbePeriodSeconds = readinessProbePeriodSeconds;
            o.readinessProbeSuccessThreshold = readinessProbeSuccessThreshold;
            o.readinessProbeTcpSocketPort = readinessProbeTcpSocketPort;
            o.readinessProbeTimeoutSeconds = readinessProbeTimeoutSeconds;
            o.volumeMounts = volumeMounts;
            o.workingDir = workingDir;
            return o;
        }
    }
}

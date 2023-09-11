// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Eci.Inputs
{

    public sealed class ContainerGroupInitContainerGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("args")]
        private InputList<string>? _args;

        /// <summary>
        /// The arguments passed to the commands.
        /// </summary>
        public InputList<string> Args
        {
            get => _args ?? (_args = new InputList<string>());
            set => _args = value;
        }

        [Input("commands")]
        private InputList<string>? _commands;

        /// <summary>
        /// The commands run by the init container.
        /// </summary>
        public InputList<string> Commands
        {
            get => _commands ?? (_commands = new InputList<string>());
            set => _commands = value;
        }

        /// <summary>
        /// The amount of CPU resources allocated to the container. Default value: `0`.
        /// </summary>
        [Input("cpu")]
        public Input<double>? Cpu { get; set; }

        [Input("environmentVars")]
        private InputList<Inputs.ContainerGroupInitContainerEnvironmentVarGetArgs>? _environmentVars;

        /// <summary>
        /// The structure of environmentVars. See `environment_vars` below.
        /// </summary>
        public InputList<Inputs.ContainerGroupInitContainerEnvironmentVarGetArgs> EnvironmentVars
        {
            get => _environmentVars ?? (_environmentVars = new InputList<Inputs.ContainerGroupInitContainerEnvironmentVarGetArgs>());
            set => _environmentVars = value;
        }

        /// <summary>
        /// The number GPUs. Default value: `0`.
        /// </summary>
        [Input("gpu")]
        public Input<int>? Gpu { get; set; }

        /// <summary>
        /// The image of the container.
        /// </summary>
        [Input("image")]
        public Input<string>? Image { get; set; }

        /// <summary>
        /// The restart policy of the image. Default value: `IfNotPresent`. Valid values: `Always`, `IfNotPresent`, `Never`.
        /// </summary>
        [Input("imagePullPolicy")]
        public Input<string>? ImagePullPolicy { get; set; }

        /// <summary>
        /// The amount of memory resources allocated to the container. Default value: `0`.
        /// </summary>
        [Input("memory")]
        public Input<double>? Memory { get; set; }

        /// <summary>
        /// The name of the mounted volume.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("ports")]
        private InputList<Inputs.ContainerGroupInitContainerPortGetArgs>? _ports;

        /// <summary>
        /// The structure of port. See `ports` below.
        /// </summary>
        public InputList<Inputs.ContainerGroupInitContainerPortGetArgs> Ports
        {
            get => _ports ?? (_ports = new InputList<Inputs.ContainerGroupInitContainerPortGetArgs>());
            set => _ports = value;
        }

        /// <summary>
        /// Indicates whether the container passed the readiness probe.
        /// </summary>
        [Input("ready")]
        public Input<bool>? Ready { get; set; }

        /// <summary>
        /// The number of times that the container restarted.
        /// </summary>
        [Input("restartCount")]
        public Input<int>? RestartCount { get; set; }

        [Input("volumeMounts")]
        private InputList<Inputs.ContainerGroupInitContainerVolumeMountGetArgs>? _volumeMounts;

        /// <summary>
        /// The structure of volumeMounts. See `volume_mounts` below.
        /// </summary>
        public InputList<Inputs.ContainerGroupInitContainerVolumeMountGetArgs> VolumeMounts
        {
            get => _volumeMounts ?? (_volumeMounts = new InputList<Inputs.ContainerGroupInitContainerVolumeMountGetArgs>());
            set => _volumeMounts = value;
        }

        /// <summary>
        /// The working directory of the container.
        /// </summary>
        [Input("workingDir")]
        public Input<string>? WorkingDir { get; set; }

        public ContainerGroupInitContainerGetArgs()
        {
        }
        public static new ContainerGroupInitContainerGetArgs Empty => new ContainerGroupInitContainerGetArgs();
    }
}

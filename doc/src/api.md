# API

## Objects

WebGPU objects are referenced through blind handles, which are homogenous to pointers. The following objects are defined by the API:


 - [`WGPUAdapter`](../object/Adapter)
 - [`WGPUBindGroup`](../object/BindGroup)
 - [`WGPUBindGroupLayout`](../object/BindGroupLayout)
 - [`WGPUBuffer`](../object/Buffer)
 - [`WGPUCommandBuffer`](../object/CommandBuffer)
 - [`WGPUCommandEncoder`](../object/CommandEncoder)
 - [`WGPUComputePassEncoder`](../object/ComputePassEncoder)
 - [`WGPUComputePipeline`](../object/ComputePipeline)
 - [`WGPUDevice`](../object/Device)
 - [`WGPUInstance`](../object/Instance)
 - [`WGPUPipelineLayout`](../object/PipelineLayout)
 - [`WGPUQuerySet`](../object/QuerySet)
 - [`WGPUQueue`](../object/Queue)
 - [`WGPURenderBundle`](../object/RenderBundle)
 - [`WGPURenderBundleEncoder`](../object/RenderBundleEncoder)
 - [`WGPURenderPassEncoder`](../object/RenderPassEncoder)
 - [`WGPURenderPipeline`](../object/RenderPipeline)
 - [`WGPUSampler`](../object/Sampler)
 - [`WGPUShaderModule`](../object/ShaderModule)
 - [`WGPUSurface`](../object/Surface)
 - [`WGPUTexture`](../object/Texture)
 - [`WGPUTextureView`](../object/TextureView)

## Structures


 - [`WGPUAdapterInfo`](AdapterInfo)
 - [`WGPUAdapterProperties`](AdapterProperties)
 - [`WGPUBindGroupEntry`](BindGroupEntry)
 - [`WGPUBlendComponent`](BlendComponent)
 - [`WGPUBufferBindingLayout`](BufferBindingLayout)
 - [`WGPUBufferDescriptor`](BufferDescriptor)
 - [`WGPUColor`](Color)
 - [`WGPUCommandBufferDescriptor`](CommandBufferDescriptor)
 - [`WGPUCommandEncoderDescriptor`](CommandEncoderDescriptor)
 - [`WGPUCompilationMessage`](CompilationMessage)
 - [`WGPUComputePassTimestampWrites`](ComputePassTimestampWrites)
 - [`WGPUConstantEntry`](ConstantEntry)
 - [`WGPUExtent3D`](Extent3D)
 - [`WGPUInstanceDescriptor`](InstanceDescriptor)
 - [`WGPULimits`](Limits)
 - [`WGPUMultisampleState`](MultisampleState)
 - [`WGPUOrigin3D`](Origin3D)
 - [`WGPUPipelineLayoutDescriptor`](PipelineLayoutDescriptor)
 - [`WGPUPrimitiveDepthClipControl`](PrimitiveDepthClipControl)
 - [`WGPUPrimitiveState`](PrimitiveState)
 - [`WGPUQuerySetDescriptor`](QuerySetDescriptor)
 - [`WGPUQueueDescriptor`](QueueDescriptor)
 - [`WGPURenderBundleDescriptor`](RenderBundleDescriptor)
 - [`WGPURenderBundleEncoderDescriptor`](RenderBundleEncoderDescriptor)
 - [`WGPURenderPassDepthStencilAttachment`](RenderPassDepthStencilAttachment)
 - [`WGPURenderPassDescriptorMaxDrawCount`](RenderPassDescriptorMaxDrawCount)
 - [`WGPURenderPassTimestampWrites`](RenderPassTimestampWrites)
 - [`WGPURequestAdapterOptions`](RequestAdapterOptions)
 - [`WGPUSamplerBindingLayout`](SamplerBindingLayout)
 - [`WGPUSamplerDescriptor`](SamplerDescriptor)
 - [`WGPUShaderModuleCompilationHint`](ShaderModuleCompilationHint)
 - [`WGPUShaderModuleSPIRVDescriptor`](ShaderModuleSPIRVDescriptor)
 - [`WGPUShaderModuleWGSLDescriptor`](ShaderModuleWGSLDescriptor)
 - [`WGPUStencilFaceState`](StencilFaceState)
 - [`WGPUStorageTextureBindingLayout`](StorageTextureBindingLayout)
 - [`WGPUSurfaceCapabilities`](SurfaceCapabilities)
 - [`WGPUSurfaceConfiguration`](SurfaceConfiguration)
 - [`WGPUSurfaceDescriptor`](SurfaceDescriptor)
 - [`WGPUSurfaceDescriptorFromAndroidNativeWindow`](SurfaceDescriptorFromAndroidNativeWindow)
 - [`WGPUSurfaceDescriptorFromCanvasHTMLSelector`](SurfaceDescriptorFromCanvasHTMLSelector)
 - [`WGPUSurfaceDescriptorFromMetalLayer`](SurfaceDescriptorFromMetalLayer)
 - [`WGPUSurfaceDescriptorFromWaylandSurface`](SurfaceDescriptorFromWaylandSurface)
 - [`WGPUSurfaceDescriptorFromWindowsHWND`](SurfaceDescriptorFromWindowsHWND)
 - [`WGPUSurfaceDescriptorFromXcbWindow`](SurfaceDescriptorFromXcbWindow)
 - [`WGPUSurfaceDescriptorFromXlibWindow`](SurfaceDescriptorFromXlibWindow)
 - [`WGPUSurfaceTexture`](SurfaceTexture)
 - [`WGPUTextureBindingLayout`](TextureBindingLayout)
 - [`WGPUTextureDataLayout`](TextureDataLayout)
 - [`WGPUTextureViewDescriptor`](TextureViewDescriptor)
 - [`WGPUVertexAttribute`](VertexAttribute)
 - [`WGPUBindGroupDescriptor`](BindGroupDescriptor)
 - [`WGPUBindGroupLayoutEntry`](BindGroupLayoutEntry)
 - [`WGPUBlendState`](BlendState)
 - [`WGPUCompilationInfo`](CompilationInfo)
 - [`WGPUComputePassDescriptor`](ComputePassDescriptor)
 - [`WGPUDepthStencilState`](DepthStencilState)
 - [`WGPUImageCopyBuffer`](ImageCopyBuffer)
 - [`WGPUImageCopyTexture`](ImageCopyTexture)
 - [`WGPUProgrammableStageDescriptor`](ProgrammableStageDescriptor)
 - [`WGPURenderPassColorAttachment`](RenderPassColorAttachment)
 - [`WGPURequiredLimits`](RequiredLimits)
 - [`WGPUShaderModuleDescriptor`](ShaderModuleDescriptor)
 - [`WGPUSupportedLimits`](SupportedLimits)
 - [`WGPUTextureDescriptor`](TextureDescriptor)
 - [`WGPUVertexBufferLayout`](VertexBufferLayout)
 - [`WGPUBindGroupLayoutDescriptor`](BindGroupLayoutDescriptor)
 - [`WGPUColorTargetState`](ColorTargetState)
 - [`WGPUComputePipelineDescriptor`](ComputePipelineDescriptor)
 - [`WGPUDeviceDescriptor`](DeviceDescriptor)
 - [`WGPURenderPassDescriptor`](RenderPassDescriptor)
 - [`WGPUVertexState`](VertexState)
 - [`WGPUFragmentState`](FragmentState)
 - [`WGPURenderPipelineDescriptor`](RenderPipelineDescriptor)


## Enumerations


 - [`WGPUAdapterType`](../enum/AdapterType)
 - [`WGPUAddressMode`](../enum/AddressMode)
 - [`WGPUBackendType`](../enum/BackendType)
 - [`WGPUBlendFactor`](../enum/BlendFactor)
 - [`WGPUBlendOperation`](../enum/BlendOperation)
 - [`WGPUBufferBindingType`](../enum/BufferBindingType)
 - [`WGPUBufferMapAsyncStatus`](../enum/BufferMapAsyncStatus)
 - [`WGPUBufferMapState`](../enum/BufferMapState)
 - [`WGPUCompareFunction`](../enum/CompareFunction)
 - [`WGPUCompilationInfoRequestStatus`](../enum/CompilationInfoRequestStatus)
 - [`WGPUCompilationMessageType`](../enum/CompilationMessageType)
 - [`WGPUCompositeAlphaMode`](../enum/CompositeAlphaMode)
 - [`WGPUCreatePipelineAsyncStatus`](../enum/CreatePipelineAsyncStatus)
 - [`WGPUCullMode`](../enum/CullMode)
 - [`WGPUDeviceLostReason`](../enum/DeviceLostReason)
 - [`WGPUErrorFilter`](../enum/ErrorFilter)
 - [`WGPUErrorType`](../enum/ErrorType)
 - [`WGPUFeatureName`](../enum/FeatureName)
 - [`WGPUFilterMode`](../enum/FilterMode)
 - [`WGPUFrontFace`](../enum/FrontFace)
 - [`WGPUIndexFormat`](../enum/IndexFormat)
 - [`WGPULoadOp`](../enum/LoadOp)
 - [`WGPUMipmapFilterMode`](../enum/MipmapFilterMode)
 - [`WGPUPowerPreference`](../enum/PowerPreference)
 - [`WGPUPresentMode`](../enum/PresentMode)
 - [`WGPUPrimitiveTopology`](../enum/PrimitiveTopology)
 - [`WGPUQueryType`](../enum/QueryType)
 - [`WGPUQueueWorkDoneStatus`](../enum/QueueWorkDoneStatus)
 - [`WGPURequestAdapterStatus`](../enum/RequestAdapterStatus)
 - [`WGPURequestDeviceStatus`](../enum/RequestDeviceStatus)
 - [`WGPUSType`](../enum/SType)
 - [`WGPUSamplerBindingType`](../enum/SamplerBindingType)
 - [`WGPUStencilOperation`](../enum/StencilOperation)
 - [`WGPUStorageTextureAccess`](../enum/StorageTextureAccess)
 - [`WGPUStoreOp`](../enum/StoreOp)
 - [`WGPUSurfaceGetCurrentTextureStatus`](../enum/SurfaceGetCurrentTextureStatus)
 - [`WGPUTextureAspect`](../enum/TextureAspect)
 - [`WGPUTextureDimension`](../enum/TextureDimension)
 - [`WGPUTextureFormat`](../enum/TextureFormat)
 - [`WGPUTextureSampleType`](../enum/TextureSampleType)
 - [`WGPUTextureViewDimension`](../enum/TextureViewDimension)
 - [`WGPUVertexFormat`](../enum/VertexFormat)
 - [`WGPUVertexStepMode`](../enum/VertexStepMode)
 - [`WGPUWGSLFeatureName`](../enum/WGSLFeatureName)


## Constants

TODO

 - `WGPU_ARRAY_LAYER_COUNT_UNDEFINED` = `0xffffffffUL`

TODO

 - `WGPU_COPY_STRIDE_UNDEFINED` = `0xffffffffUL`

TODO

 - `WGPU_DEPTH_SLICE_UNDEFINED` = `0xffffffffUL`

TODO

 - `WGPU_LIMIT_U32_UNDEFINED` = `0xffffffffUL`

TODO

 - `WGPU_LIMIT_U64_UNDEFINED` = `0xffffffffffffffffULL`

TODO

 - `WGPU_MIP_LEVEL_COUNT_UNDEFINED` = `0xffffffffUL`

TODO

 - `WGPU_QUERY_SET_INDEX_UNDEFINED` = `0xffffffffUL`

TODO

 - `WGPU_WHOLE_MAP_SIZE` = `SIZE_MAX`

TODO


 - `WGPU_WHOLE_SIZE` = `0xffffffffffffffffULL`



## Typedefs

 - `typedef uint32_t WGPUFlags`
 - `typedef uint32_t WGPUBool`



## Callback types



TODO

 - `typedef void (*WGPUAdapterRequestAdapterInfoCallback)(struct WGPUAdapterInfo adapterInfo, WGPU_NULLABLE void * userdata)`

TODO

 - `typedef void (*WGPUAdapterRequestDeviceCallback)(WGPURequestDeviceStatus status, WGPUDevice device, char const * message, WGPU_NULLABLE void * userdata)`

TODO

 - `typedef void (*WGPUBufferMapAsyncCallback)(WGPUBufferMapAsyncStatus status, WGPU_NULLABLE void * userdata)`

TODO

 - `typedef void (*WGPUDeviceCreateComputePipelineAsyncCallback)(WGPUCreatePipelineAsyncStatus status, WGPUComputePipeline pipeline, char const * message, WGPU_NULLABLE void * userdata)`

TODO

 - `typedef void (*WGPUDeviceCreateRenderPipelineAsyncCallback)(WGPUCreatePipelineAsyncStatus status, WGPURenderPipeline pipeline, char const * message, WGPU_NULLABLE void * userdata)`

TODO

 - `typedef void (*WGPUInstanceRequestAdapterCallback)(WGPURequestAdapterStatus status, WGPUAdapter adapter, char const * message, WGPU_NULLABLE void * userdata)`

TODO

 - `typedef void (*WGPUQueueOnSubmittedWorkDoneCallback)(WGPUQueueWorkDoneStatus status, WGPU_NULLABLE void * userdata)`

TODO

 - `typedef void (*WGPUShaderModuleGetCompilationInfoCallback)(WGPUCompilationInfoRequestStatus status, struct WGPUCompilationInfo const * compilationInfo, WGPU_NULLABLE void * userdata)`


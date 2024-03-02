# API

## Objects

WebGPU objects are referenced through blind handles, which are homogenous to pointers. The following objects are defined by the API:


 - [`WGPUAdapter`](../object#WGPUAdapter)
 - [`WGPUBindGroup`](../object#WGPUBindGroup)
 - [`WGPUBindGroupLayout`](../object#WGPUBindGroupLayout)
 - [`WGPUBuffer`](../object#WGPUBuffer)
 - [`WGPUCommandBuffer`](../object#WGPUCommandBuffer)
 - [`WGPUCommandEncoder`](../object#WGPUCommandEncoder)
 - [`WGPUComputePassEncoder`](../object#WGPUComputePassEncoder)
 - [`WGPUComputePipeline`](../object#WGPUComputePipeline)
 - [`WGPUDevice`](../object#WGPUDevice)
 - [`WGPUInstance`](../object#WGPUInstance)
 - [`WGPUPipelineLayout`](../object#WGPUPipelineLayout)
 - [`WGPUQuerySet`](../object#WGPUQuerySet)
 - [`WGPUQueue`](../object#WGPUQueue)
 - [`WGPURenderBundle`](../object#WGPURenderBundle)
 - [`WGPURenderBundleEncoder`](../object#WGPURenderBundleEncoder)
 - [`WGPURenderPassEncoder`](../object#WGPURenderPassEncoder)
 - [`WGPURenderPipeline`](../object#WGPURenderPipeline)
 - [`WGPUSampler`](../object#WGPUSampler)
 - [`WGPUShaderModule`](../object#WGPUShaderModule)
 - [`WGPUSurface`](../object#WGPUSurface)
 - [`WGPUTexture`](../object#WGPUTexture)
 - [`WGPUTextureView`](../object#WGPUTextureView)

## Structures


 - `WGPUAdapterInfo`
 - `WGPUAdapterProperties`
 - `WGPUBindGroupEntry`
 - `WGPUBlendComponent`
 - `WGPUBufferBindingLayout`
 - `WGPUBufferDescriptor`
 - `WGPUColor`
 - `WGPUCommandBufferDescriptor`
 - `WGPUCommandEncoderDescriptor`
 - `WGPUCompilationMessage`
 - `WGPUComputePassTimestampWrites`
 - `WGPUConstantEntry`
 - `WGPUExtent3D`
 - `WGPUInstanceDescriptor`
 - `WGPULimits`
 - `WGPUMultisampleState`
 - `WGPUOrigin3D`
 - `WGPUPipelineLayoutDescriptor`
 - `WGPUPrimitiveDepthClipControl`
 - `WGPUPrimitiveState`
 - `WGPUQuerySetDescriptor`
 - `WGPUQueueDescriptor`
 - `WGPURenderBundleDescriptor`
 - `WGPURenderBundleEncoderDescriptor`
 - `WGPURenderPassDepthStencilAttachment`
 - `WGPURenderPassDescriptorMaxDrawCount`
 - `WGPURenderPassTimestampWrites`
 - `WGPURequestAdapterOptions`
 - `WGPUSamplerBindingLayout`
 - `WGPUSamplerDescriptor`
 - `WGPUShaderModuleCompilationHint`
 - `WGPUShaderModuleSPIRVDescriptor`
 - `WGPUShaderModuleWGSLDescriptor`
 - `WGPUStencilFaceState`
 - `WGPUStorageTextureBindingLayout`
 - `WGPUSurfaceCapabilities`
 - `WGPUSurfaceConfiguration`
 - `WGPUSurfaceDescriptor`
 - `WGPUSurfaceDescriptorFromAndroidNativeWindow`
 - `WGPUSurfaceDescriptorFromCanvasHTMLSelector`
 - `WGPUSurfaceDescriptorFromMetalLayer`
 - `WGPUSurfaceDescriptorFromWaylandSurface`
 - `WGPUSurfaceDescriptorFromWindowsHWND`
 - `WGPUSurfaceDescriptorFromXcbWindow`
 - `WGPUSurfaceDescriptorFromXlibWindow`
 - `WGPUSurfaceTexture`
 - `WGPUTextureBindingLayout`
 - `WGPUTextureDataLayout`
 - `WGPUTextureViewDescriptor`
 - `WGPUVertexAttribute`
 - `WGPUBindGroupDescriptor`
 - `WGPUBindGroupLayoutEntry`
 - `WGPUBlendState`
 - `WGPUCompilationInfo`
 - `WGPUComputePassDescriptor`
 - `WGPUDepthStencilState`
 - `WGPUImageCopyBuffer`
 - `WGPUImageCopyTexture`
 - `WGPUProgrammableStageDescriptor`
 - `WGPURenderPassColorAttachment`
 - `WGPURequiredLimits`
 - `WGPUShaderModuleDescriptor`
 - `WGPUSupportedLimits`
 - `WGPUTextureDescriptor`
 - `WGPUVertexBufferLayout`
 - `WGPUBindGroupLayoutDescriptor`
 - `WGPUColorTargetState`
 - `WGPUComputePipelineDescriptor`
 - `WGPUDeviceDescriptor`
 - `WGPURenderPassDescriptor`
 - `WGPUVertexState`
 - `WGPUFragmentState`
 - `WGPURenderPipelineDescriptor`


## Enumerations


 - [`WGPUAdapterType`](../enum#WGPUAdapterType)
 - [`WGPUAddressMode`](../enum#WGPUAddressMode)
 - [`WGPUBackendType`](../enum#WGPUBackendType)
 - [`WGPUBlendFactor`](../enum#WGPUBlendFactor)
 - [`WGPUBlendOperation`](../enum#WGPUBlendOperation)
 - [`WGPUBufferBindingType`](../enum#WGPUBufferBindingType)
 - [`WGPUBufferMapAsyncStatus`](../enum#WGPUBufferMapAsyncStatus)
 - [`WGPUBufferMapState`](../enum#WGPUBufferMapState)
 - [`WGPUCompareFunction`](../enum#WGPUCompareFunction)
 - [`WGPUCompilationInfoRequestStatus`](../enum#WGPUCompilationInfoRequestStatus)
 - [`WGPUCompilationMessageType`](../enum#WGPUCompilationMessageType)
 - [`WGPUCompositeAlphaMode`](../enum#WGPUCompositeAlphaMode)
 - [`WGPUCreatePipelineAsyncStatus`](../enum#WGPUCreatePipelineAsyncStatus)
 - [`WGPUCullMode`](../enum#WGPUCullMode)
 - [`WGPUDeviceLostReason`](../enum#WGPUDeviceLostReason)
 - [`WGPUErrorFilter`](../enum#WGPUErrorFilter)
 - [`WGPUErrorType`](../enum#WGPUErrorType)
 - [`WGPUFeatureName`](../enum#WGPUFeatureName)
 - [`WGPUFilterMode`](../enum#WGPUFilterMode)
 - [`WGPUFrontFace`](../enum#WGPUFrontFace)
 - [`WGPUIndexFormat`](../enum#WGPUIndexFormat)
 - [`WGPULoadOp`](../enum#WGPULoadOp)
 - [`WGPUMipmapFilterMode`](../enum#WGPUMipmapFilterMode)
 - [`WGPUPowerPreference`](../enum#WGPUPowerPreference)
 - [`WGPUPresentMode`](../enum#WGPUPresentMode)
 - [`WGPUPrimitiveTopology`](../enum#WGPUPrimitiveTopology)
 - [`WGPUQueryType`](../enum#WGPUQueryType)
 - [`WGPUQueueWorkDoneStatus`](../enum#WGPUQueueWorkDoneStatus)
 - [`WGPURequestAdapterStatus`](../enum#WGPURequestAdapterStatus)
 - [`WGPURequestDeviceStatus`](../enum#WGPURequestDeviceStatus)
 - [`WGPUSType`](../enum#WGPUSType)
 - [`WGPUSamplerBindingType`](../enum#WGPUSamplerBindingType)
 - [`WGPUStencilOperation`](../enum#WGPUStencilOperation)
 - [`WGPUStorageTextureAccess`](../enum#WGPUStorageTextureAccess)
 - [`WGPUStoreOp`](../enum#WGPUStoreOp)
 - [`WGPUSurfaceGetCurrentTextureStatus`](../enum#WGPUSurfaceGetCurrentTextureStatus)
 - [`WGPUTextureAspect`](../enum#WGPUTextureAspect)
 - [`WGPUTextureDimension`](../enum#WGPUTextureDimension)
 - [`WGPUTextureFormat`](../enum#WGPUTextureFormat)
 - [`WGPUTextureSampleType`](../enum#WGPUTextureSampleType)
 - [`WGPUTextureViewDimension`](../enum#WGPUTextureViewDimension)
 - [`WGPUVertexFormat`](../enum#WGPUVertexFormat)
 - [`WGPUVertexStepMode`](../enum#WGPUVertexStepMode)
 - [`WGPUWGSLFeatureName`](../enum#WGPUWGSLFeatureName)


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


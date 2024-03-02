# Objects

WebGPU objects are referenced through blind handles, which are homogenous to pointers. The following objects are defined and detailed below:

 - [`WGPUAdapter`](#WGPUAdapter)
 - [`WGPUBindGroup`](#WGPUBindGroup)
 - [`WGPUBindGroupLayout`](#WGPUBindGroupLayout)
 - [`WGPUBuffer`](#WGPUBuffer)
 - [`WGPUCommandBuffer`](#WGPUCommandBuffer)
 - [`WGPUCommandEncoder`](#WGPUCommandEncoder)
 - [`WGPUComputePassEncoder`](#WGPUComputePassEncoder)
 - [`WGPUComputePipeline`](#WGPUComputePipeline)
 - [`WGPUDevice`](#WGPUDevice)
 - [`WGPUInstance`](#WGPUInstance)
 - [`WGPUPipelineLayout`](#WGPUPipelineLayout)
 - [`WGPUQuerySet`](#WGPUQuerySet)
 - [`WGPUQueue`](#WGPUQueue)
 - [`WGPURenderBundle`](#WGPURenderBundle)
 - [`WGPURenderBundleEncoder`](#WGPURenderBundleEncoder)
 - [`WGPURenderPassEncoder`](#WGPURenderPassEncoder)
 - [`WGPURenderPipeline`](#WGPURenderPipeline)
 - [`WGPUSampler`](#WGPUSampler)
 - [`WGPUShaderModule`](#WGPUShaderModule)
 - [`WGPUSurface`](#WGPUSurface)
 - [`WGPUTexture`](#WGPUTexture)
 - [`WGPUTextureView`](#WGPUTextureView)

**NB** Additionally, all objects provide a `Release` and a `Reference` methods:

 - `void wgpu<ObjectName>Reference(WGPU<ObjectName>)`
 - `void wgpu<ObjectName>Release(WGPU<ObjectName>)`



## Adapter { #WGPUAdapter }


This is a test.




### `wgpuAdapterEnumerateFeatures` { #wgpuAdapterEnumerateFeatures }

**Prototype:** `size_t wgpuAdapterEnumerateFeatures(WGPUAdapter adapter, WGPUFeatureName * features)`


TODO


**Arguments:**

 - TODO



**Returns:** `size_t` 
TODO





### `wgpuAdapterGetLimits` { #wgpuAdapterGetLimits }

**Prototype:** `WGPUBool wgpuAdapterGetLimits(WGPUAdapter adapter, WGPUSupportedLimits * limits)`


Get the limits that the adapter supports. This is used to drive the
choice of required limits in {device_descriptor.required_limits}.


**Arguments:**

 - TODO



**Returns:** `WGPUBool` 
True iff limits were successfully retrieved.





### `wgpuAdapterGetProperties` { #wgpuAdapterGetProperties }

**Prototype:** `void wgpuAdapterGetProperties(WGPUAdapter adapter, WGPUAdapterProperties * properties)`


TODO


**Arguments:**

 - TODO




### `wgpuAdapterHasFeature` { #wgpuAdapterHasFeature }

**Prototype:** `WGPUBool wgpuAdapterHasFeature(WGPUAdapter adapter, WGPUFeatureName feature)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUBool` 
TODO





### `wgpuAdapterRequestAdapterInfo` { #wgpuAdapterRequestAdapterInfo }

**Prototype:** `void wgpuAdapterRequestAdapterInfo(WGPUAdapter adapter, WGPUAdapterRequestAdapterInfoCallback callback, WGPU_NULLABLE void * userdata)`


TODO


**Arguments:**

 - TODO




### `wgpuAdapterRequestDevice` { #wgpuAdapterRequestDevice }

**Prototype:** `void wgpuAdapterRequestDevice(WGPUAdapter adapter, WGPU_NULLABLE WGPUDeviceDescriptor const * descriptor, WGPUAdapterRequestDeviceCallback callback, WGPU_NULLABLE void * userdata)`


TODO


**Arguments:**

 - TODO






## BindGroup { #WGPUBindGroup }


TODO




### `wgpuBindGroupSetLabel` { #wgpuBindGroupSetLabel }

**Prototype:** `void wgpuBindGroupSetLabel(WGPUBindGroup bindGroup, char const * label)`


TODO


**Arguments:**

 - TODO






## BindGroupLayout { #WGPUBindGroupLayout }


TODO




### `wgpuBindGroupLayoutSetLabel` { #wgpuBindGroupLayoutSetLabel }

**Prototype:** `void wgpuBindGroupLayoutSetLabel(WGPUBindGroupLayout bindGroupLayout, char const * label)`


TODO


**Arguments:**

 - TODO






## Buffer { #WGPUBuffer }


TODO




### `wgpuBufferDestroy` { #wgpuBufferDestroy }

**Prototype:** `void wgpuBufferDestroy(WGPUBuffer buffer)`


TODO



**Arguments:**

 - TODO




### `wgpuBufferGetConstMappedRange` { #wgpuBufferGetConstMappedRange }

**Prototype:** `void const * wgpuBufferGetConstMappedRange(WGPUBuffer buffer, size_t offset, size_t size)`


TODO


**Arguments:**

 - TODO



**Returns:** `void const *` 
TODO





### `wgpuBufferGetMapState` { #wgpuBufferGetMapState }

**Prototype:** `WGPUBufferMapState wgpuBufferGetMapState(WGPUBuffer buffer)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUBufferMapState` 
TODO





### `wgpuBufferGetMappedRange` { #wgpuBufferGetMappedRange }

**Prototype:** `void * wgpuBufferGetMappedRange(WGPUBuffer buffer, size_t offset, size_t size)`


TODO


**Arguments:**

 - TODO



**Returns:** `void *` 
TODO





### `wgpuBufferGetSize` { #wgpuBufferGetSize }

**Prototype:** `uint64_t wgpuBufferGetSize(WGPUBuffer buffer)`


TODO


**Arguments:**

 - TODO



**Returns:** `uint64_t` 
TODO





### `wgpuBufferGetUsage` { #wgpuBufferGetUsage }

**Prototype:** `WGPUBufferUsageFlags wgpuBufferGetUsage(WGPUBuffer buffer)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUBufferUsageFlags` 
TODO





### `wgpuBufferMapAsync` { #wgpuBufferMapAsync }

**Prototype:** `void wgpuBufferMapAsync(WGPUBuffer buffer, WGPUMapModeFlags mode, size_t offset, size_t size, WGPUBufferMapAsyncCallback callback, WGPU_NULLABLE void * userdata)`


TODO


**Arguments:**

 - TODO




### `wgpuBufferSetLabel` { #wgpuBufferSetLabel }

**Prototype:** `void wgpuBufferSetLabel(WGPUBuffer buffer, char const * label)`


TODO


**Arguments:**

 - TODO




### `wgpuBufferUnmap` { #wgpuBufferUnmap }

**Prototype:** `void wgpuBufferUnmap(WGPUBuffer buffer)`


TODO


**Arguments:**

 - TODO






## CommandBuffer { #WGPUCommandBuffer }


TODO




### `wgpuCommandBufferSetLabel` { #wgpuCommandBufferSetLabel }

**Prototype:** `void wgpuCommandBufferSetLabel(WGPUCommandBuffer commandBuffer, char const * label)`


TODO


**Arguments:**

 - TODO






## CommandEncoder { #WGPUCommandEncoder }


TODO




### `wgpuCommandEncoderBeginComputePass` { #wgpuCommandEncoderBeginComputePass }

**Prototype:** `WGPUComputePassEncoder wgpuCommandEncoderBeginComputePass(WGPUCommandEncoder commandEncoder, WGPU_NULLABLE WGPUComputePassDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUComputePassEncoder` 
TODO





### `wgpuCommandEncoderBeginRenderPass` { #wgpuCommandEncoderBeginRenderPass }

**Prototype:** `WGPURenderPassEncoder wgpuCommandEncoderBeginRenderPass(WGPUCommandEncoder commandEncoder, WGPURenderPassDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPURenderPassEncoder` 
TODO





### `wgpuCommandEncoderClearBuffer` { #wgpuCommandEncoderClearBuffer }

**Prototype:** `void wgpuCommandEncoderClearBuffer(WGPUCommandEncoder commandEncoder, WGPUBuffer buffer, uint64_t offset, uint64_t size)`


TODO


**Arguments:**

 - TODO




### `wgpuCommandEncoderCopyBufferToBuffer` { #wgpuCommandEncoderCopyBufferToBuffer }

**Prototype:** `void wgpuCommandEncoderCopyBufferToBuffer(WGPUCommandEncoder commandEncoder, WGPUBuffer source, uint64_t sourceOffset, WGPUBuffer destination, uint64_t destinationOffset, uint64_t size)`


TODO


**Arguments:**

 - TODO




### `wgpuCommandEncoderCopyBufferToTexture` { #wgpuCommandEncoderCopyBufferToTexture }

**Prototype:** `void wgpuCommandEncoderCopyBufferToTexture(WGPUCommandEncoder commandEncoder, WGPUImageCopyBuffer const * source, WGPUImageCopyTexture const * destination, WGPUExtent3D const * copySize)`


TODO


**Arguments:**

 - TODO




### `wgpuCommandEncoderCopyTextureToBuffer` { #wgpuCommandEncoderCopyTextureToBuffer }

**Prototype:** `void wgpuCommandEncoderCopyTextureToBuffer(WGPUCommandEncoder commandEncoder, WGPUImageCopyTexture const * source, WGPUImageCopyBuffer const * destination, WGPUExtent3D const * copySize)`


TODO


**Arguments:**

 - TODO




### `wgpuCommandEncoderCopyTextureToTexture` { #wgpuCommandEncoderCopyTextureToTexture }

**Prototype:** `void wgpuCommandEncoderCopyTextureToTexture(WGPUCommandEncoder commandEncoder, WGPUImageCopyTexture const * source, WGPUImageCopyTexture const * destination, WGPUExtent3D const * copySize)`


TODO


**Arguments:**

 - TODO




### `wgpuCommandEncoderFinish` { #wgpuCommandEncoderFinish }

**Prototype:** `WGPUCommandBuffer wgpuCommandEncoderFinish(WGPUCommandEncoder commandEncoder, WGPU_NULLABLE WGPUCommandBufferDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUCommandBuffer` 
TODO





### `wgpuCommandEncoderInsertDebugMarker` { #wgpuCommandEncoderInsertDebugMarker }

**Prototype:** `void wgpuCommandEncoderInsertDebugMarker(WGPUCommandEncoder commandEncoder, char const * markerLabel)`


TODO


**Arguments:**

 - TODO




### `wgpuCommandEncoderPopDebugGroup` { #wgpuCommandEncoderPopDebugGroup }

**Prototype:** `void wgpuCommandEncoderPopDebugGroup(WGPUCommandEncoder commandEncoder)`


TODO


**Arguments:**

 - TODO




### `wgpuCommandEncoderPushDebugGroup` { #wgpuCommandEncoderPushDebugGroup }

**Prototype:** `void wgpuCommandEncoderPushDebugGroup(WGPUCommandEncoder commandEncoder, char const * groupLabel)`


TODO


**Arguments:**

 - TODO




### `wgpuCommandEncoderResolveQuerySet` { #wgpuCommandEncoderResolveQuerySet }

**Prototype:** `void wgpuCommandEncoderResolveQuerySet(WGPUCommandEncoder commandEncoder, WGPUQuerySet querySet, uint32_t firstQuery, uint32_t queryCount, WGPUBuffer destination, uint64_t destinationOffset)`


TODO


**Arguments:**

 - TODO




### `wgpuCommandEncoderSetLabel` { #wgpuCommandEncoderSetLabel }

**Prototype:** `void wgpuCommandEncoderSetLabel(WGPUCommandEncoder commandEncoder, char const * label)`


TODO


**Arguments:**

 - TODO




### `wgpuCommandEncoderWriteTimestamp` { #wgpuCommandEncoderWriteTimestamp }

**Prototype:** `void wgpuCommandEncoderWriteTimestamp(WGPUCommandEncoder commandEncoder, WGPUQuerySet querySet, uint32_t queryIndex)`


TODO


**Arguments:**

 - TODO






## ComputePassEncoder { #WGPUComputePassEncoder }


TODO




### `wgpuComputePassEncoderDispatchWorkgroups` { #wgpuComputePassEncoderDispatchWorkgroups }

**Prototype:** `void wgpuComputePassEncoderDispatchWorkgroups(WGPUComputePassEncoder computePassEncoder, uint32_t workgroupCountX, uint32_t workgroupCountY, uint32_t workgroupCountZ)`


TODO


**Arguments:**

 - TODO




### `wgpuComputePassEncoderDispatchWorkgroupsIndirect` { #wgpuComputePassEncoderDispatchWorkgroupsIndirect }

**Prototype:** `void wgpuComputePassEncoderDispatchWorkgroupsIndirect(WGPUComputePassEncoder computePassEncoder, WGPUBuffer indirectBuffer, uint64_t indirectOffset)`


TODO


**Arguments:**

 - TODO




### `wgpuComputePassEncoderEnd` { #wgpuComputePassEncoderEnd }

**Prototype:** `void wgpuComputePassEncoderEnd(WGPUComputePassEncoder computePassEncoder)`


TODO


**Arguments:**

 - TODO




### `wgpuComputePassEncoderInsertDebugMarker` { #wgpuComputePassEncoderInsertDebugMarker }

**Prototype:** `void wgpuComputePassEncoderInsertDebugMarker(WGPUComputePassEncoder computePassEncoder, char const * markerLabel)`


TODO


**Arguments:**

 - TODO




### `wgpuComputePassEncoderPopDebugGroup` { #wgpuComputePassEncoderPopDebugGroup }

**Prototype:** `void wgpuComputePassEncoderPopDebugGroup(WGPUComputePassEncoder computePassEncoder)`


TODO


**Arguments:**

 - TODO




### `wgpuComputePassEncoderPushDebugGroup` { #wgpuComputePassEncoderPushDebugGroup }

**Prototype:** `void wgpuComputePassEncoderPushDebugGroup(WGPUComputePassEncoder computePassEncoder, char const * groupLabel)`


TODO


**Arguments:**

 - TODO




### `wgpuComputePassEncoderSetBindGroup` { #wgpuComputePassEncoderSetBindGroup }

**Prototype:** `void wgpuComputePassEncoderSetBindGroup(WGPUComputePassEncoder computePassEncoder, uint32_t groupIndex, WGPU_NULLABLE WGPUBindGroup group, size_t dynamicOffsetCount, uint32_t const * dynamicOffsets)`


TODO


**Arguments:**

 - TODO




### `wgpuComputePassEncoderSetLabel` { #wgpuComputePassEncoderSetLabel }

**Prototype:** `void wgpuComputePassEncoderSetLabel(WGPUComputePassEncoder computePassEncoder, char const * label)`


TODO


**Arguments:**

 - TODO




### `wgpuComputePassEncoderSetPipeline` { #wgpuComputePassEncoderSetPipeline }

**Prototype:** `void wgpuComputePassEncoderSetPipeline(WGPUComputePassEncoder computePassEncoder, WGPUComputePipeline pipeline)`


TODO


**Arguments:**

 - TODO






## ComputePipeline { #WGPUComputePipeline }


TODO




### `wgpuComputePipelineGetBindGroupLayout` { #wgpuComputePipelineGetBindGroupLayout }

**Prototype:** `WGPUBindGroupLayout wgpuComputePipelineGetBindGroupLayout(WGPUComputePipeline computePipeline, uint32_t groupIndex)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUBindGroupLayout` 
TODO





### `wgpuComputePipelineSetLabel` { #wgpuComputePipelineSetLabel }

**Prototype:** `void wgpuComputePipelineSetLabel(WGPUComputePipeline computePipeline, char const * label)`


TODO


**Arguments:**

 - TODO






## Device { #WGPUDevice }


TODO




### `wgpuDeviceCreateBindGroup` { #wgpuDeviceCreateBindGroup }

**Prototype:** `WGPUBindGroup wgpuDeviceCreateBindGroup(WGPUDevice device, WGPUBindGroupDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUBindGroup` 
TODO





### `wgpuDeviceCreateBindGroupLayout` { #wgpuDeviceCreateBindGroupLayout }

**Prototype:** `WGPUBindGroupLayout wgpuDeviceCreateBindGroupLayout(WGPUDevice device, WGPUBindGroupLayoutDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUBindGroupLayout` 
TODO





### `wgpuDeviceCreateBuffer` { #wgpuDeviceCreateBuffer }

**Prototype:** `WGPUBuffer wgpuDeviceCreateBuffer(WGPUDevice device, WGPUBufferDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUBuffer` 
TODO





### `wgpuDeviceCreateCommandEncoder` { #wgpuDeviceCreateCommandEncoder }

**Prototype:** `WGPUCommandEncoder wgpuDeviceCreateCommandEncoder(WGPUDevice device, WGPU_NULLABLE WGPUCommandEncoderDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUCommandEncoder` 
TODO





### `wgpuDeviceCreateComputePipeline` { #wgpuDeviceCreateComputePipeline }

**Prototype:** `WGPUComputePipeline wgpuDeviceCreateComputePipeline(WGPUDevice device, WGPUComputePipelineDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUComputePipeline` 
TODO





### `wgpuDeviceCreateComputePipelineAsync` { #wgpuDeviceCreateComputePipelineAsync }

**Prototype:** `void wgpuDeviceCreateComputePipelineAsync(WGPUDevice device, WGPUComputePipelineDescriptor const * descriptor, WGPUDeviceCreateComputePipelineAsyncCallback callback, WGPU_NULLABLE void * userdata)`


TODO


**Arguments:**

 - TODO




### `wgpuDeviceCreatePipelineLayout` { #wgpuDeviceCreatePipelineLayout }

**Prototype:** `WGPUPipelineLayout wgpuDeviceCreatePipelineLayout(WGPUDevice device, WGPUPipelineLayoutDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUPipelineLayout` 
TODO





### `wgpuDeviceCreateQuerySet` { #wgpuDeviceCreateQuerySet }

**Prototype:** `WGPUQuerySet wgpuDeviceCreateQuerySet(WGPUDevice device, WGPUQuerySetDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUQuerySet` 
TODO





### `wgpuDeviceCreateRenderBundleEncoder` { #wgpuDeviceCreateRenderBundleEncoder }

**Prototype:** `WGPURenderBundleEncoder wgpuDeviceCreateRenderBundleEncoder(WGPUDevice device, WGPURenderBundleEncoderDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPURenderBundleEncoder` 
TODO





### `wgpuDeviceCreateRenderPipeline` { #wgpuDeviceCreateRenderPipeline }

**Prototype:** `WGPURenderPipeline wgpuDeviceCreateRenderPipeline(WGPUDevice device, WGPURenderPipelineDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPURenderPipeline` 
TODO





### `wgpuDeviceCreateRenderPipelineAsync` { #wgpuDeviceCreateRenderPipelineAsync }

**Prototype:** `void wgpuDeviceCreateRenderPipelineAsync(WGPUDevice device, WGPURenderPipelineDescriptor const * descriptor, WGPUDeviceCreateRenderPipelineAsyncCallback callback, WGPU_NULLABLE void * userdata)`


TODO


**Arguments:**

 - TODO




### `wgpuDeviceCreateSampler` { #wgpuDeviceCreateSampler }

**Prototype:** `WGPUSampler wgpuDeviceCreateSampler(WGPUDevice device, WGPU_NULLABLE WGPUSamplerDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUSampler` 
TODO





### `wgpuDeviceCreateShaderModule` { #wgpuDeviceCreateShaderModule }

**Prototype:** `WGPUShaderModule wgpuDeviceCreateShaderModule(WGPUDevice device, WGPUShaderModuleDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUShaderModule` 
TODO





### `wgpuDeviceCreateTexture` { #wgpuDeviceCreateTexture }

**Prototype:** `WGPUTexture wgpuDeviceCreateTexture(WGPUDevice device, WGPUTextureDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUTexture` 
TODO





### `wgpuDeviceDestroy` { #wgpuDeviceDestroy }

**Prototype:** `void wgpuDeviceDestroy(WGPUDevice device)`


TODO


**Arguments:**

 - TODO




### `wgpuDeviceEnumerateFeatures` { #wgpuDeviceEnumerateFeatures }

**Prototype:** `size_t wgpuDeviceEnumerateFeatures(WGPUDevice device, WGPUFeatureName * features)`


TODO


**Arguments:**

 - TODO



**Returns:** `size_t` 
TODO





### `wgpuDeviceGetLimits` { #wgpuDeviceGetLimits }

**Prototype:** `WGPUBool wgpuDeviceGetLimits(WGPUDevice device, WGPUSupportedLimits * limits)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUBool` 
TODO





### `wgpuDeviceGetQueue` { #wgpuDeviceGetQueue }

**Prototype:** `WGPUQueue wgpuDeviceGetQueue(WGPUDevice device)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUQueue` 
TODO





### `wgpuDeviceHasFeature` { #wgpuDeviceHasFeature }

**Prototype:** `WGPUBool wgpuDeviceHasFeature(WGPUDevice device, WGPUFeatureName feature)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUBool` 
TODO





### `wgpuDevicePopErrorScope` { #wgpuDevicePopErrorScope }

**Prototype:** `void wgpuDevicePopErrorScope(WGPUDevice device, WGPUErrorCallback callback, void * userdata)`


TODO


**Arguments:**

 - TODO




### `wgpuDevicePushErrorScope` { #wgpuDevicePushErrorScope }

**Prototype:** `void wgpuDevicePushErrorScope(WGPUDevice device, WGPUErrorFilter filter)`


TODO


**Arguments:**

 - TODO




### `wgpuDeviceSetLabel` { #wgpuDeviceSetLabel }

**Prototype:** `void wgpuDeviceSetLabel(WGPUDevice device, char const * label)`


TODO


**Arguments:**

 - TODO




### `wgpuDeviceSetUncapturedErrorCallback` { #wgpuDeviceSetUncapturedErrorCallback }

**Prototype:** `void wgpuDeviceSetUncapturedErrorCallback(WGPUDevice device, WGPUErrorCallback callback, void * userdata)`


TODO


**Arguments:**

 - TODO






## Instance { #WGPUInstance }


TODO




### `wgpuInstanceCreateSurface` { #wgpuInstanceCreateSurface }

**Prototype:** `WGPUSurface wgpuInstanceCreateSurface(WGPUInstance instance, WGPUSurfaceDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUSurface` 
TODO





### `wgpuInstanceHasWGSLLanguageFeature` { #wgpuInstanceHasWGSLLanguageFeature }

**Prototype:** `WGPUBool wgpuInstanceHasWGSLLanguageFeature(WGPUInstance instance, WGPUWGSLFeatureName feature)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUBool` 
TODO





### `wgpuInstanceProcessEvents` { #wgpuInstanceProcessEvents }

**Prototype:** `void wgpuInstanceProcessEvents(WGPUInstance instance)`


TODO


**Arguments:**

 - TODO




### `wgpuInstanceRequestAdapter` { #wgpuInstanceRequestAdapter }

**Prototype:** `void wgpuInstanceRequestAdapter(WGPUInstance instance, WGPU_NULLABLE WGPURequestAdapterOptions const * options, WGPUInstanceRequestAdapterCallback callback, WGPU_NULLABLE void * userdata)`


TODO


**Arguments:**

 - TODO






## PipelineLayout { #WGPUPipelineLayout }


TODO




### `wgpuPipelineLayoutSetLabel` { #wgpuPipelineLayoutSetLabel }

**Prototype:** `void wgpuPipelineLayoutSetLabel(WGPUPipelineLayout pipelineLayout, char const * label)`


TODO


**Arguments:**

 - TODO






## QuerySet { #WGPUQuerySet }


TODO




### `wgpuQuerySetDestroy` { #wgpuQuerySetDestroy }

**Prototype:** `void wgpuQuerySetDestroy(WGPUQuerySet querySet)`


TODO



**Arguments:**

 - TODO




### `wgpuQuerySetGetCount` { #wgpuQuerySetGetCount }

**Prototype:** `uint32_t wgpuQuerySetGetCount(WGPUQuerySet querySet)`


TODO


**Arguments:**

 - TODO



**Returns:** `uint32_t` 
TODO





### `wgpuQuerySetGetType` { #wgpuQuerySetGetType }

**Prototype:** `WGPUQueryType wgpuQuerySetGetType(WGPUQuerySet querySet)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUQueryType` 
TODO





### `wgpuQuerySetSetLabel` { #wgpuQuerySetSetLabel }

**Prototype:** `void wgpuQuerySetSetLabel(WGPUQuerySet querySet, char const * label)`


TODO


**Arguments:**

 - TODO






## Queue { #WGPUQueue }


TODO




### `wgpuQueueOnSubmittedWorkDone` { #wgpuQueueOnSubmittedWorkDone }

**Prototype:** `void wgpuQueueOnSubmittedWorkDone(WGPUQueue queue, WGPUQueueOnSubmittedWorkDoneCallback callback, WGPU_NULLABLE void * userdata)`


TODO


**Arguments:**

 - TODO




### `wgpuQueueSetLabel` { #wgpuQueueSetLabel }

**Prototype:** `void wgpuQueueSetLabel(WGPUQueue queue, char const * label)`


TODO


**Arguments:**

 - TODO




### `wgpuQueueSubmit` { #wgpuQueueSubmit }

**Prototype:** `void wgpuQueueSubmit(WGPUQueue queue, size_t commandCount, WGPUCommandBuffer const * commands)`


TODO


**Arguments:**

 - TODO




### `wgpuQueueWriteBuffer` { #wgpuQueueWriteBuffer }

**Prototype:** `void wgpuQueueWriteBuffer(WGPUQueue queue, WGPUBuffer buffer, uint64_t bufferOffset, void const * data, size_t size)`


TODO


**Arguments:**

 - TODO




### `wgpuQueueWriteTexture` { #wgpuQueueWriteTexture }

**Prototype:** `void wgpuQueueWriteTexture(WGPUQueue queue, WGPUImageCopyTexture const * destination, void const * data, size_t dataSize, WGPUTextureDataLayout const * dataLayout, WGPUExtent3D const * writeSize)`


TODO


**Arguments:**

 - TODO






## RenderBundle { #WGPURenderBundle }


TODO




### `wgpuRenderBundleSetLabel` { #wgpuRenderBundleSetLabel }

**Prototype:** `void wgpuRenderBundleSetLabel(WGPURenderBundle renderBundle, char const * label)`


TODO


**Arguments:**

 - TODO






## RenderBundleEncoder { #WGPURenderBundleEncoder }


TODO




### `wgpuRenderBundleEncoderDraw` { #wgpuRenderBundleEncoderDraw }

**Prototype:** `void wgpuRenderBundleEncoderDraw(WGPURenderBundleEncoder renderBundleEncoder, uint32_t vertexCount, uint32_t instanceCount, uint32_t firstVertex, uint32_t firstInstance)`


TODO


**Arguments:**

 - TODO




### `wgpuRenderBundleEncoderDrawIndexed` { #wgpuRenderBundleEncoderDrawIndexed }

**Prototype:** `void wgpuRenderBundleEncoderDrawIndexed(WGPURenderBundleEncoder renderBundleEncoder, uint32_t indexCount, uint32_t instanceCount, uint32_t firstIndex, int32_t baseVertex, uint32_t firstInstance)`


TODO


**Arguments:**

 - TODO




### `wgpuRenderBundleEncoderDrawIndexedIndirect` { #wgpuRenderBundleEncoderDrawIndexedIndirect }

**Prototype:** `void wgpuRenderBundleEncoderDrawIndexedIndirect(WGPURenderBundleEncoder renderBundleEncoder, WGPUBuffer indirectBuffer, uint64_t indirectOffset)`


TODO


**Arguments:**

 - TODO




### `wgpuRenderBundleEncoderDrawIndirect` { #wgpuRenderBundleEncoderDrawIndirect }

**Prototype:** `void wgpuRenderBundleEncoderDrawIndirect(WGPURenderBundleEncoder renderBundleEncoder, WGPUBuffer indirectBuffer, uint64_t indirectOffset)`


TODO


**Arguments:**

 - TODO




### `wgpuRenderBundleEncoderFinish` { #wgpuRenderBundleEncoderFinish }

**Prototype:** `WGPURenderBundle wgpuRenderBundleEncoderFinish(WGPURenderBundleEncoder renderBundleEncoder, WGPU_NULLABLE WGPURenderBundleDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPURenderBundle` 
TODO





### `wgpuRenderBundleEncoderInsertDebugMarker` { #wgpuRenderBundleEncoderInsertDebugMarker }

**Prototype:** `void wgpuRenderBundleEncoderInsertDebugMarker(WGPURenderBundleEncoder renderBundleEncoder, char const * markerLabel)`


TODO


**Arguments:**

 - TODO




### `wgpuRenderBundleEncoderPopDebugGroup` { #wgpuRenderBundleEncoderPopDebugGroup }

**Prototype:** `void wgpuRenderBundleEncoderPopDebugGroup(WGPURenderBundleEncoder renderBundleEncoder)`


TODO


**Arguments:**

 - TODO




### `wgpuRenderBundleEncoderPushDebugGroup` { #wgpuRenderBundleEncoderPushDebugGroup }

**Prototype:** `void wgpuRenderBundleEncoderPushDebugGroup(WGPURenderBundleEncoder renderBundleEncoder, char const * groupLabel)`


TODO


**Arguments:**

 - TODO




### `wgpuRenderBundleEncoderSetBindGroup` { #wgpuRenderBundleEncoderSetBindGroup }

**Prototype:** `void wgpuRenderBundleEncoderSetBindGroup(WGPURenderBundleEncoder renderBundleEncoder, uint32_t groupIndex, WGPU_NULLABLE WGPUBindGroup group, size_t dynamicOffsetCount, uint32_t const * dynamicOffsets)`


TODO


**Arguments:**

 - TODO




### `wgpuRenderBundleEncoderSetIndexBuffer` { #wgpuRenderBundleEncoderSetIndexBuffer }

**Prototype:** `void wgpuRenderBundleEncoderSetIndexBuffer(WGPURenderBundleEncoder renderBundleEncoder, WGPUBuffer buffer, WGPUIndexFormat format, uint64_t offset, uint64_t size)`


TODO


**Arguments:**

 - TODO




### `wgpuRenderBundleEncoderSetLabel` { #wgpuRenderBundleEncoderSetLabel }

**Prototype:** `void wgpuRenderBundleEncoderSetLabel(WGPURenderBundleEncoder renderBundleEncoder, char const * label)`


TODO


**Arguments:**

 - TODO




### `wgpuRenderBundleEncoderSetPipeline` { #wgpuRenderBundleEncoderSetPipeline }

**Prototype:** `void wgpuRenderBundleEncoderSetPipeline(WGPURenderBundleEncoder renderBundleEncoder, WGPURenderPipeline pipeline)`


TODO


**Arguments:**

 - TODO




### `wgpuRenderBundleEncoderSetVertexBuffer` { #wgpuRenderBundleEncoderSetVertexBuffer }

**Prototype:** `void wgpuRenderBundleEncoderSetVertexBuffer(WGPURenderBundleEncoder renderBundleEncoder, uint32_t slot, WGPU_NULLABLE WGPUBuffer buffer, uint64_t offset, uint64_t size)`


TODO


**Arguments:**

 - TODO






## RenderPassEncoder { #WGPURenderPassEncoder }


TODO




### `wgpuRenderPassEncoderBeginOcclusionQuery` { #wgpuRenderPassEncoderBeginOcclusionQuery }

**Prototype:** `void wgpuRenderPassEncoderBeginOcclusionQuery(WGPURenderPassEncoder renderPassEncoder, uint32_t queryIndex)`


TODO


**Arguments:**

 - TODO




### `wgpuRenderPassEncoderDraw` { #wgpuRenderPassEncoderDraw }

**Prototype:** `void wgpuRenderPassEncoderDraw(WGPURenderPassEncoder renderPassEncoder, uint32_t vertexCount, uint32_t instanceCount, uint32_t firstVertex, uint32_t firstInstance)`


TODO


**Arguments:**

 - TODO




### `wgpuRenderPassEncoderDrawIndexed` { #wgpuRenderPassEncoderDrawIndexed }

**Prototype:** `void wgpuRenderPassEncoderDrawIndexed(WGPURenderPassEncoder renderPassEncoder, uint32_t indexCount, uint32_t instanceCount, uint32_t firstIndex, int32_t baseVertex, uint32_t firstInstance)`


TODO


**Arguments:**

 - TODO




### `wgpuRenderPassEncoderDrawIndexedIndirect` { #wgpuRenderPassEncoderDrawIndexedIndirect }

**Prototype:** `void wgpuRenderPassEncoderDrawIndexedIndirect(WGPURenderPassEncoder renderPassEncoder, WGPUBuffer indirectBuffer, uint64_t indirectOffset)`


TODO


**Arguments:**

 - TODO




### `wgpuRenderPassEncoderDrawIndirect` { #wgpuRenderPassEncoderDrawIndirect }

**Prototype:** `void wgpuRenderPassEncoderDrawIndirect(WGPURenderPassEncoder renderPassEncoder, WGPUBuffer indirectBuffer, uint64_t indirectOffset)`


TODO


**Arguments:**

 - TODO




### `wgpuRenderPassEncoderEnd` { #wgpuRenderPassEncoderEnd }

**Prototype:** `void wgpuRenderPassEncoderEnd(WGPURenderPassEncoder renderPassEncoder)`


TODO


**Arguments:**

 - TODO




### `wgpuRenderPassEncoderEndOcclusionQuery` { #wgpuRenderPassEncoderEndOcclusionQuery }

**Prototype:** `void wgpuRenderPassEncoderEndOcclusionQuery(WGPURenderPassEncoder renderPassEncoder)`


TODO


**Arguments:**

 - TODO




### `wgpuRenderPassEncoderExecuteBundles` { #wgpuRenderPassEncoderExecuteBundles }

**Prototype:** `void wgpuRenderPassEncoderExecuteBundles(WGPURenderPassEncoder renderPassEncoder, size_t bundleCount, WGPURenderBundle const * bundles)`


TODO


**Arguments:**

 - TODO




### `wgpuRenderPassEncoderInsertDebugMarker` { #wgpuRenderPassEncoderInsertDebugMarker }

**Prototype:** `void wgpuRenderPassEncoderInsertDebugMarker(WGPURenderPassEncoder renderPassEncoder, char const * markerLabel)`


TODO


**Arguments:**

 - TODO




### `wgpuRenderPassEncoderPopDebugGroup` { #wgpuRenderPassEncoderPopDebugGroup }

**Prototype:** `void wgpuRenderPassEncoderPopDebugGroup(WGPURenderPassEncoder renderPassEncoder)`


TODO


**Arguments:**

 - TODO




### `wgpuRenderPassEncoderPushDebugGroup` { #wgpuRenderPassEncoderPushDebugGroup }

**Prototype:** `void wgpuRenderPassEncoderPushDebugGroup(WGPURenderPassEncoder renderPassEncoder, char const * groupLabel)`


TODO


**Arguments:**

 - TODO




### `wgpuRenderPassEncoderSetBindGroup` { #wgpuRenderPassEncoderSetBindGroup }

**Prototype:** `void wgpuRenderPassEncoderSetBindGroup(WGPURenderPassEncoder renderPassEncoder, uint32_t groupIndex, WGPU_NULLABLE WGPUBindGroup group, size_t dynamicOffsetCount, uint32_t const * dynamicOffsets)`


TODO


**Arguments:**

 - TODO




### `wgpuRenderPassEncoderSetBlendConstant` { #wgpuRenderPassEncoderSetBlendConstant }

**Prototype:** `void wgpuRenderPassEncoderSetBlendConstant(WGPURenderPassEncoder renderPassEncoder, WGPUColor const * color)`


TODO


**Arguments:**

 - TODO




### `wgpuRenderPassEncoderSetIndexBuffer` { #wgpuRenderPassEncoderSetIndexBuffer }

**Prototype:** `void wgpuRenderPassEncoderSetIndexBuffer(WGPURenderPassEncoder renderPassEncoder, WGPUBuffer buffer, WGPUIndexFormat format, uint64_t offset, uint64_t size)`


TODO


**Arguments:**

 - TODO




### `wgpuRenderPassEncoderSetLabel` { #wgpuRenderPassEncoderSetLabel }

**Prototype:** `void wgpuRenderPassEncoderSetLabel(WGPURenderPassEncoder renderPassEncoder, char const * label)`


TODO


**Arguments:**

 - TODO




### `wgpuRenderPassEncoderSetPipeline` { #wgpuRenderPassEncoderSetPipeline }

**Prototype:** `void wgpuRenderPassEncoderSetPipeline(WGPURenderPassEncoder renderPassEncoder, WGPURenderPipeline pipeline)`


TODO


**Arguments:**

 - TODO




### `wgpuRenderPassEncoderSetScissorRect` { #wgpuRenderPassEncoderSetScissorRect }

**Prototype:** `void wgpuRenderPassEncoderSetScissorRect(WGPURenderPassEncoder renderPassEncoder, uint32_t x, uint32_t y, uint32_t width, uint32_t height)`


TODO


**Arguments:**

 - TODO




### `wgpuRenderPassEncoderSetStencilReference` { #wgpuRenderPassEncoderSetStencilReference }

**Prototype:** `void wgpuRenderPassEncoderSetStencilReference(WGPURenderPassEncoder renderPassEncoder, uint32_t reference)`


TODO


**Arguments:**

 - TODO




### `wgpuRenderPassEncoderSetVertexBuffer` { #wgpuRenderPassEncoderSetVertexBuffer }

**Prototype:** `void wgpuRenderPassEncoderSetVertexBuffer(WGPURenderPassEncoder renderPassEncoder, uint32_t slot, WGPU_NULLABLE WGPUBuffer buffer, uint64_t offset, uint64_t size)`


TODO


**Arguments:**

 - TODO




### `wgpuRenderPassEncoderSetViewport` { #wgpuRenderPassEncoderSetViewport }

**Prototype:** `void wgpuRenderPassEncoderSetViewport(WGPURenderPassEncoder renderPassEncoder, float x, float y, float width, float height, float minDepth, float maxDepth)`


TODO


**Arguments:**

 - TODO






## RenderPipeline { #WGPURenderPipeline }


TODO




### `wgpuRenderPipelineGetBindGroupLayout` { #wgpuRenderPipelineGetBindGroupLayout }

**Prototype:** `WGPUBindGroupLayout wgpuRenderPipelineGetBindGroupLayout(WGPURenderPipeline renderPipeline, uint32_t groupIndex)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUBindGroupLayout` 
TODO





### `wgpuRenderPipelineSetLabel` { #wgpuRenderPipelineSetLabel }

**Prototype:** `void wgpuRenderPipelineSetLabel(WGPURenderPipeline renderPipeline, char const * label)`


TODO


**Arguments:**

 - TODO






## Sampler { #WGPUSampler }


TODO




### `wgpuSamplerSetLabel` { #wgpuSamplerSetLabel }

**Prototype:** `void wgpuSamplerSetLabel(WGPUSampler sampler, char const * label)`


TODO


**Arguments:**

 - TODO






## ShaderModule { #WGPUShaderModule }


TODO




### `wgpuShaderModuleGetCompilationInfo` { #wgpuShaderModuleGetCompilationInfo }

**Prototype:** `void wgpuShaderModuleGetCompilationInfo(WGPUShaderModule shaderModule, WGPUShaderModuleGetCompilationInfoCallback callback, WGPU_NULLABLE void * userdata)`


TODO


**Arguments:**

 - TODO




### `wgpuShaderModuleSetLabel` { #wgpuShaderModuleSetLabel }

**Prototype:** `void wgpuShaderModuleSetLabel(WGPUShaderModule shaderModule, char const * label)`


TODO


**Arguments:**

 - TODO






## Surface { #WGPUSurface }


TODO




### `wgpuSurfaceConfigure` { #wgpuSurfaceConfigure }

**Prototype:** `void wgpuSurfaceConfigure(WGPUSurface surface, WGPUSurfaceConfiguration const * config)`


TODO


**Arguments:**

 - TODO




### `wgpuSurfaceGetCapabilities` { #wgpuSurfaceGetCapabilities }

**Prototype:** `void wgpuSurfaceGetCapabilities(WGPUSurface surface, WGPUAdapter adapter, WGPUSurfaceCapabilities * capabilities)`


TODO


**Arguments:**

 - TODO




### `wgpuSurfaceGetCurrentTexture` { #wgpuSurfaceGetCurrentTexture }

**Prototype:** `void wgpuSurfaceGetCurrentTexture(WGPUSurface surface, WGPUSurfaceTexture * surfaceTexture)`


TODO


**Arguments:**

 - TODO




### `wgpuSurfaceGetPreferredFormat` { #wgpuSurfaceGetPreferredFormat }

**Prototype:** `WGPUTextureFormat wgpuSurfaceGetPreferredFormat(WGPUSurface surface, WGPUAdapter adapter)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUTextureFormat` 
TODO





### `wgpuSurfacePresent` { #wgpuSurfacePresent }

**Prototype:** `void wgpuSurfacePresent(WGPUSurface surface)`


TODO


**Arguments:**

 - TODO




### `wgpuSurfaceSetLabel` { #wgpuSurfaceSetLabel }

**Prototype:** `void wgpuSurfaceSetLabel(WGPUSurface surface, char const * label)`


TODO


**Arguments:**

 - TODO




### `wgpuSurfaceUnconfigure` { #wgpuSurfaceUnconfigure }

**Prototype:** `void wgpuSurfaceUnconfigure(WGPUSurface surface)`


TODO


**Arguments:**

 - TODO







## Texture { #WGPUTexture }


TODO




### `wgpuTextureCreateView` { #wgpuTextureCreateView }

**Prototype:** `WGPUTextureView wgpuTextureCreateView(WGPUTexture texture, WGPU_NULLABLE WGPUTextureViewDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUTextureView` 
TODO





### `wgpuTextureDestroy` { #wgpuTextureDestroy }

**Prototype:** `void wgpuTextureDestroy(WGPUTexture texture)`


TODO



**Arguments:**

 - TODO




### `wgpuTextureGetDepthOrArrayLayers` { #wgpuTextureGetDepthOrArrayLayers }

**Prototype:** `uint32_t wgpuTextureGetDepthOrArrayLayers(WGPUTexture texture)`


TODO


**Arguments:**

 - TODO



**Returns:** `uint32_t` 
TODO





### `wgpuTextureGetDimension` { #wgpuTextureGetDimension }

**Prototype:** `WGPUTextureDimension wgpuTextureGetDimension(WGPUTexture texture)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUTextureDimension` 
TODO





### `wgpuTextureGetFormat` { #wgpuTextureGetFormat }

**Prototype:** `WGPUTextureFormat wgpuTextureGetFormat(WGPUTexture texture)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUTextureFormat` 
TODO





### `wgpuTextureGetHeight` { #wgpuTextureGetHeight }

**Prototype:** `uint32_t wgpuTextureGetHeight(WGPUTexture texture)`


TODO


**Arguments:**

 - TODO



**Returns:** `uint32_t` 
TODO





### `wgpuTextureGetMipLevelCount` { #wgpuTextureGetMipLevelCount }

**Prototype:** `uint32_t wgpuTextureGetMipLevelCount(WGPUTexture texture)`


TODO


**Arguments:**

 - TODO



**Returns:** `uint32_t` 
TODO





### `wgpuTextureGetSampleCount` { #wgpuTextureGetSampleCount }

**Prototype:** `uint32_t wgpuTextureGetSampleCount(WGPUTexture texture)`


TODO


**Arguments:**

 - TODO



**Returns:** `uint32_t` 
TODO





### `wgpuTextureGetUsage` { #wgpuTextureGetUsage }

**Prototype:** `WGPUTextureUsageFlags wgpuTextureGetUsage(WGPUTexture texture)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUTextureUsageFlags` 
TODO





### `wgpuTextureGetWidth` { #wgpuTextureGetWidth }

**Prototype:** `uint32_t wgpuTextureGetWidth(WGPUTexture texture)`


TODO


**Arguments:**

 - TODO



**Returns:** `uint32_t` 
TODO





### `wgpuTextureSetLabel` { #wgpuTextureSetLabel }

**Prototype:** `void wgpuTextureSetLabel(WGPUTexture texture, char const * label)`


TODO


**Arguments:**

 - TODO






## TextureView { #WGPUTextureView }


TODO




### `wgpuTextureViewSetLabel` { #wgpuTextureViewSetLabel }

**Prototype:** `void wgpuTextureViewSetLabel(WGPUTextureView textureView, char const * label)`


TODO


**Arguments:**

 - TODO







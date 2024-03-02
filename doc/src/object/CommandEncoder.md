

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





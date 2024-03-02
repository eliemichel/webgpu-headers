

# CommandEncoder { #WGPUCommandEncoder }


TODO




<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## BeginComputePass { #wgpuCommandEncoderBeginComputePass }

```C
WGPUComputePassEncoder wgpuCommandEncoderBeginComputePass(WGPUCommandEncoder commandEncoder, WGPU_NULLABLE WGPUComputePassDescriptor const * descriptor)
```


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUComputePassEncoder` 
TODO





<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## BeginRenderPass { #wgpuCommandEncoderBeginRenderPass }

```C
WGPURenderPassEncoder wgpuCommandEncoderBeginRenderPass(WGPUCommandEncoder commandEncoder, WGPURenderPassDescriptor const * descriptor)
```


TODO


**Arguments:**

 - TODO



**Returns:** `WGPURenderPassEncoder` 
TODO





<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## ClearBuffer { #wgpuCommandEncoderClearBuffer }

```C
void wgpuCommandEncoderClearBuffer(WGPUCommandEncoder commandEncoder, WGPUBuffer buffer, uint64_t offset, uint64_t size)
```


TODO


**Arguments:**

 - TODO




<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## CopyBufferToBuffer { #wgpuCommandEncoderCopyBufferToBuffer }

```C
void wgpuCommandEncoderCopyBufferToBuffer(WGPUCommandEncoder commandEncoder, WGPUBuffer source, uint64_t sourceOffset, WGPUBuffer destination, uint64_t destinationOffset, uint64_t size)
```


TODO


**Arguments:**

 - TODO




<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## CopyBufferToTexture { #wgpuCommandEncoderCopyBufferToTexture }

```C
void wgpuCommandEncoderCopyBufferToTexture(WGPUCommandEncoder commandEncoder, WGPUImageCopyBuffer const * source, WGPUImageCopyTexture const * destination, WGPUExtent3D const * copySize)
```


TODO


**Arguments:**

 - TODO




<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## CopyTextureToBuffer { #wgpuCommandEncoderCopyTextureToBuffer }

```C
void wgpuCommandEncoderCopyTextureToBuffer(WGPUCommandEncoder commandEncoder, WGPUImageCopyTexture const * source, WGPUImageCopyBuffer const * destination, WGPUExtent3D const * copySize)
```


TODO


**Arguments:**

 - TODO




<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## CopyTextureToTexture { #wgpuCommandEncoderCopyTextureToTexture }

```C
void wgpuCommandEncoderCopyTextureToTexture(WGPUCommandEncoder commandEncoder, WGPUImageCopyTexture const * source, WGPUImageCopyTexture const * destination, WGPUExtent3D const * copySize)
```


TODO


**Arguments:**

 - TODO




<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## Finish { #wgpuCommandEncoderFinish }

```C
WGPUCommandBuffer wgpuCommandEncoderFinish(WGPUCommandEncoder commandEncoder, WGPU_NULLABLE WGPUCommandBufferDescriptor const * descriptor)
```


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUCommandBuffer` 
TODO





<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## InsertDebugMarker { #wgpuCommandEncoderInsertDebugMarker }

```C
void wgpuCommandEncoderInsertDebugMarker(WGPUCommandEncoder commandEncoder, char const * markerLabel)
```


TODO


**Arguments:**

 - TODO




<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## PopDebugGroup { #wgpuCommandEncoderPopDebugGroup }

```C
void wgpuCommandEncoderPopDebugGroup(WGPUCommandEncoder commandEncoder)
```


TODO


**Arguments:**

 - TODO




<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## PushDebugGroup { #wgpuCommandEncoderPushDebugGroup }

```C
void wgpuCommandEncoderPushDebugGroup(WGPUCommandEncoder commandEncoder, char const * groupLabel)
```


TODO


**Arguments:**

 - TODO




<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## ResolveQuerySet { #wgpuCommandEncoderResolveQuerySet }

```C
void wgpuCommandEncoderResolveQuerySet(WGPUCommandEncoder commandEncoder, WGPUQuerySet querySet, uint32_t firstQuery, uint32_t queryCount, WGPUBuffer destination, uint64_t destinationOffset)
```


TODO


**Arguments:**

 - TODO




<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## SetLabel { #wgpuCommandEncoderSetLabel }

```C
void wgpuCommandEncoderSetLabel(WGPUCommandEncoder commandEncoder, char const * label)
```


TODO


**Arguments:**

 - TODO




<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## WriteTimestamp { #wgpuCommandEncoderWriteTimestamp }

```C
void wgpuCommandEncoderWriteTimestamp(WGPUCommandEncoder commandEncoder, WGPUQuerySet querySet, uint32_t queryIndex)
```


TODO


**Arguments:**

 - TODO



